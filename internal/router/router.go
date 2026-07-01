package router

import (
	"net/http"

	"stone-ocean-web/internal/events"
	"stone-ocean-web/internal/handlers"
	"stone-ocean-web/internal/store"

	"github.com/gin-gonic/gin"
)

func New(appStore *store.Store) *gin.Engine {
	return NewWithEvents(appStore, nil)
}

func NewWithEvents(appStore *store.Store, eventBus *events.Bus) *gin.Engine {
	r := gin.Default()
	api := handlers.NewAPIWithEvents(appStore, eventBus)

	// 挂载前端静态资源目录。
	r.Static("/static", "./web/static")
	// 加载页面模板文件。
	r.LoadHTMLGlob("web/templates/*")

	// 默认首页，内部会跳转到默认语言首页。
	r.GET("/", handlers.Home)
	// 默认语言的购买结算页。
	r.GET("/checkout", handlers.Checkout)
	// 默认语言的支付结果页，用于展示支付状态和激活码。
	r.GET("/checkout/success", handlers.CheckoutSuccess)
	// 默认语言的激活码找回页。
	r.GET("/license-recovery", handlers.LicenseRecovery)
	// 默认语言的隐私政策页。
	r.GET("/privacy", handlers.Privacy)
	// 默认语言的服务条款页。
	r.GET("/terms", handlers.Terms)
	// 健康检查接口，用于部署或监控探活。
	r.GET("/healthz", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"status": "ok"})
	})
	// 创建待支付订单。
	r.POST("/api/checkout", api.CreateCheckout)
	// 测试支付确认接口，将支付单标记为已支付并生成激活码。
	r.POST("/api/payments/:paymentNo/confirm", api.ConfirmPayment)
	// 查询支付结果；支付成功后会返回激活码。
	r.GET("/api/payments/:paymentNo/result", api.PaymentResult)
	// 发送激活码找回验证码到购买邮箱。
	r.POST("/api/license-recovery/verification-code", api.SendLicenseRecoveryCode)
	// 校验找回验证码；验证通过后返回该邮箱下的激活码。
	r.POST("/api/license-recovery/verification-code/verify", api.VerifyLicenseRecoveryCode)
	// 搜索引擎爬虫规则文件。
	r.GET("/robots.txt", handlers.Robots)
	// 站点地图文件。
	r.GET("/sitemap.xml", handlers.Sitemap)
	// 指定语言的购买结算页。
	r.GET("/:locale/checkout", handlers.LocalizedCheckout)
	// 指定语言的支付结果页。
	r.GET("/:locale/checkout/success", handlers.LocalizedCheckoutSuccess)
	// 指定语言的激活码找回页。
	r.GET("/:locale/license-recovery", handlers.LocalizedLicenseRecovery)
	// 指定语言的隐私政策页。
	r.GET("/:locale/privacy", handlers.LocalizedPrivacy)
	// 指定语言的服务条款页。
	r.GET("/:locale/terms", handlers.LocalizedTerms)
	// 指定语言首页。
	r.GET("/:locale", handlers.LocalizedHome)

	return r
}
