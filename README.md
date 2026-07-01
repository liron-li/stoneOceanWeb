# RecoverEase Web

RecoverEase Web 是 RecoverEase 密码恢复工具的官网项目，基于 Gin 开发。当前包含多语言官网页面、结算页、激活码找回页，以及一套基于 GORM 的订单、支付、授权和激活码数据层。

## 技术栈

- Go 1.22+
- Gin
- GORM
- MySQL
- HTML 模板、CSS、原生 JavaScript

## 目录结构

```text
cmd/web              Web 服务入口
cmd/dbsetup          数据库迁移和测试数据初始化命令
internal/config      .env 配置读取
internal/handlers    页面处理器和 JSON API 处理器
internal/i18n        多语言文案和语言路径
internal/router      Gin 路由注册
internal/store       GORM 模型和数据库操作
web/static           CSS 和 JavaScript 静态资源
web/templates        HTML 模板
```

## 配置

复制示例配置文件：

```powershell
Copy-Item .env.example .env
```

常用配置项：

```env
APP_PORT=8080

DB_ENABLED=true
DB_DRIVER=mysql
DB_HOST=127.0.0.1
DB_PORT=3306
DB_USER=root
DB_PASSWORD=change-me
DB_NAME=recoverease
DB_CHARSET=utf8mb4
DB_PARSE_TIME=true
DB_LOC=Local
DB_AUTO_MIGRATE=true
DB_SEED_PLANS=true
DB_SEED_DEMO_DATA=false
```

说明：

- `APP_PORT`：Web 服务端口。
- `DB_ENABLED`：是否启用数据库。为 `false` 时官网页面可以打开，但结算和找回接口会返回数据库未启用。
- `DB_AUTO_MIGRATE`：启动时是否自动执行 GORM 迁移。
- `DB_SEED_PLANS`：是否写入默认授权套餐。
- `DB_SEED_DEMO_DATA`：是否写入演示购买数据，生产环境建议保持 `false`。

真实 `.env` 已被 `.gitignore` 忽略，不应提交到仓库。

## 数据库初始化

先在 MySQL 中创建数据库，例如：

```sql
CREATE DATABASE recoverease CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci;
```

然后执行：

```powershell
go run ./cmd/dbsetup
```

该命令会执行 GORM 迁移，写入默认授权套餐，并生成演示购买数据。

会创建的表：

- `customers`
- `license_plans`
- `orders`
- `payments`
- `licenses`
- `recovery_tokens`

默认授权套餐：

- `recoverease-pro-monthly`：一个月授权，9 美元
- `recoverease-pro-lifetime`：永久授权，29 美元

演示找回邮箱：

- `demo-lifetime@recoverease.test`
- `demo-monthly@recoverease.test`

## 本地运行

```powershell
go run ./cmd/web
```

打开页面：

```text
http://localhost:8080/zh
http://localhost:8080/zh/checkout
http://localhost:8080/zh/license-recovery
```

## 打包

```powershell
go build -o web.exe ./cmd/web
```

打包后的程序需要在项目根目录运行，因为它会读取 `web/templates`、`web/static` 和 `.env`。

```powershell
.\web.exe
```

## API 接口

创建结算订单：

```http
POST /api/checkout
Content-Type: application/json

{
  "email": "buyer@example.com",
  "license": "lifetime",
  "paymentMethod": "usdt"
}
```

当前测试流程中确认支付：

```http
POST /api/payments/{paymentNo}/confirm
```

通过购买邮箱找回激活码，先发送验证码：

```http
POST /api/license-recovery/verification-code
Content-Type: application/json

{
  "email": "buyer@example.com"
}
```

再用验证码查询激活码：

```http
POST /api/license-recovery/verification-code/verify
Content-Type: application/json

{
  "email": "buyer@example.com",
  "code": "123456"
}
```

当前前端会把确认支付接口当作测试支付流程使用：点击“立即支付”后创建订单，然后立即确认支付并生成激活码。后续接入真实支付服务时，应将这一步替换为支付服务商跳转和 webhook 回调确认。

## 测试

```powershell
go test ./...
```

说明：`internal/store` 的测试使用内存 SQLite；生产 Web 运行时只使用 MySQL。

## GoLand 注意事项

如果 GoLand 出现下面这类错误：

```text
CreateProcess error=193, %1 不是有效的 Win32 应用程序
compile: version "go1.24.1" does not match go tool version "go1.25.9"
```

检查运行配置：

- 不要在 Environment variables 中手动设置 `GOROOT`。
- 保证 GoLand 使用同一套完整 Go SDK / toolchain。
- 如果 GoLand 一直使用旧的临时配置，删除该临时配置，或使用项目中的 `RecoverEase Web` 运行配置。
- 必要时执行 `File > Invalidate Caches... > Just Restart` 清理缓存。
