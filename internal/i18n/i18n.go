package i18n

import (
	"strconv"
	"strings"
)

const DefaultLocale = "en"

type Language struct {
	Code     string
	Name     string
	HTMLLang string
	HrefLang string
	Path     string
	Active   bool
}

var languageDefinitions = []Language{
	{Code: "zh", Name: "中文", HTMLLang: "zh-CN", HrefLang: "zh-CN", Path: "/zh"},
	{Code: "en", Name: "English", HTMLLang: "en", HrefLang: "en", Path: "/en"},
	{Code: "ja", Name: "日本語", HTMLLang: "ja", HrefLang: "ja", Path: "/ja"},
	{Code: "ko", Name: "한국어", HTMLLang: "ko", HrefLang: "ko", Path: "/ko"},
	{Code: "de", Name: "Deutsch", HTMLLang: "de", HrefLang: "de", Path: "/de"},
	{Code: "fr", Name: "Français", HTMLLang: "fr", HrefLang: "fr", Path: "/fr"},
	{Code: "es", Name: "Español", HTMLLang: "es", HrefLang: "es", Path: "/es"},
	{Code: "pt", Name: "Português", HTMLLang: "pt", HrefLang: "pt", Path: "/pt"},
	{Code: "ru", Name: "Русский", HTMLLang: "ru", HrefLang: "ru", Path: "/ru"},
}

var en = map[string]string{
	"meta.title":                "RecoverEase - Password Recovery Software",
	"meta.description":          "RecoverEase helps you regain access to password-protected archives, documents, and encrypted drives.",
	"brand.aria":                "RecoverEase home",
	"nav.aria":                  "Main navigation",
	"nav.formats":               "Formats",
	"nav.modes":                 "Modes",
	"nav.pricing":               "Pricing",
	"nav.download":              "Download",
	"controls.language":         "Language",
	"controls.theme":            "Switch to dark theme",
	"cta.freeTrial":             "Free Trial",
	"cta.buyPro":                "Buy Pro",
	"hero.eyebrow":              "Password Recovery Software",
	"hero.title":                "Regain Access to Encrypted Files",
	"hero.subtitle":             "RecoverEase helps you regain access to password-protected archives, documents, and encrypted drives.",
	"hero.actions":              "Primary actions",
	"hero.trust":                "Product highlights",
	"trust.offline":             "Offline processing",
	"trust.noUpload":            "No file uploads",
	"trust.gpu":                 "GPU acceleration",
	"appPreview.aria":           "RecoverEase software interface preview",
	"app.nav.recovery":          "Recovery",
	"app.nav.files":             "Files",
	"app.nav.hints":             "Hints",
	"app.nav.reports":           "Reports",
	"app.currentTask":           "Current task",
	"app.archiveRecovery":       "Archive recovery",
	"app.smartMode":             "Smart Mode",
	"app.encryptedDetected":     "Encrypted archive detected",
	"app.passwordHints":         "Password hints",
	"app.hintSummer":            "summer",
	"app.hintYear":              "2024",
	"app.hintCompany":           "company",
	"app.gpuSpeed":              "GPU speed",
	"app.local":                 "Local",
	"app.privacyMode":           "Privacy mode",
	"formats.eyebrow":           "Supported Formats",
	"formats.title":             "Support for Common Encrypted Files and Disk Containers",
	"formats.subtitle":          "From home backups to business archives, RecoverEase handles protected files through one clear workflow.",
	"formats.archive.title":     "Archives",
	"formats.archive.desc":      "Recover access to common RAR, ZIP, and 7Z archive scenarios.",
	"formats.office.title":      "Office Documents",
	"formats.office.desc":       "Work with password-protected Microsoft Word, Excel, and PowerPoint files.",
	"formats.pdf.title":         "PDF Files",
	"formats.pdf.desc":          "Recover PDF open passwords and regain access to important documents.",
	"formats.disk.title":        "Disk Encryption",
	"formats.disk.desc":         "Professional recovery tasks for BitLocker and VeraCrypt containers.",
	"modes.eyebrow":             "Recovery Modes",
	"modes.title":               "From Quick Attempts to Expert Strategies",
	"modes.subtitle":            "Choose the right recovery method based on password hints, file type, and time budget.",
	"modes.quick.title":         "Quick Recovery",
	"modes.quick.desc":          "Try common passwords and lightweight rules first.",
	"modes.quick.fit":           "Best for: home users and recently forgotten passwords",
	"modes.smart.title":         "Smart Recovery",
	"modes.smart.desc":          "Automatically builds a strategy from file signals, password libraries, and your hints.",
	"modes.smart.fit":           "Best for: most recovery tasks",
	"modes.deep.title":          "Deep Recovery",
	"modes.deep.desc":           "Expands the search scope for complex passwords or limited clues.",
	"modes.deep.fit":            "Best for: business archives and long-forgotten passwords",
	"modes.expert.title":        "Expert Recovery",
	"modes.expert.desc":         "Customize length, character sets, rules, and task queues.",
	"modes.expert.fit":          "Best for: IT professionals",
	"labels.recommended":        "Recommended",
	"labels.bestValue":          "Best Value",
	"advantages.eyebrow":        "Advantages",
	"advantages.title":          "Professional, Trustworthy, and Privacy-Focused",
	"advantages.subtitle":       "RecoverEase focuses on local recovery workflows for personal files, small business assets, and IT support.",
	"advantages.gpu.title":      "GPU Acceleration",
	"advantages.gpu.desc":       "Use your local graphics card to improve recovery speed.",
	"advantages.hints.title":    "Smart Password Hints",
	"advantages.hints.desc":     "Turn names, dates, keywords, and patterns into candidate strategies.",
	"advantages.offline.title":  "Offline Processing",
	"advantages.offline.desc":   "Core recovery tasks run locally without cloud queues.",
	"advantages.privacy.title":  "Privacy First",
	"advantages.privacy.desc":   "Sensitive files are not uploaded, and the task process stays under your control.",
	"advantages.noUpload.title": "No File Uploads",
	"advantages.noUpload.desc":  "Your files remain on your computer at all times.",
	"workflow.eyebrow":          "Workflow",
	"workflow.title":            "Start Recovery in Four Steps",
	"workflow.step1.title":      "Choose a File",
	"workflow.step1.desc":       "Add an archive, document, or encrypted container that needs recovery.",
	"workflow.step2.title":      "Pick a Mode",
	"workflow.step2.desc":       "Choose quick, smart, deep, or expert strategy for the task.",
	"workflow.step3.title":      "Add Hints",
	"workflow.step3.desc":       "Enter likely names, dates, phrases, or password patterns.",
	"workflow.step4.title":      "Start Recovery",
	"workflow.step4.desc":       "Track real-time progress, speed, and results.",
	"pricing.eyebrow":           "Pricing",
	"pricing.title":             "Simple, Transparent Licensing",
	"pricing.free.title":        "Free",
	"pricing.free.item1":        "File analysis",
	"pricing.free.item2":        "Quick recovery",
	"pricing.free.item3":        "Common password dictionary",
	"pricing.pro.title":         "Pro",
	"pricing.pro.note":          "one-time lifetime license",
	"pricing.pro.item1":         "Unlimited recovery",
	"pricing.pro.item2":         "Deep recovery",
	"pricing.pro.item3":         "GPU acceleration",
	"pricing.pro.item4":         "Batch tasks",
	"pricing.pro.item5":         "Priority support",
	"faq.eyebrow":               "FAQ",
	"faq.title":                 "Frequently Asked Questions",
	"faq.allPasswords.q":        "Can every password be recovered?",
	"faq.allPasswords.a":        "No. Results depend on password complexity, available hints, file type, and the amount of compute time you can spend.",
	"faq.duration.q":            "How long does recovery take?",
	"faq.duration.a":            "Simple passwords may take minutes. Complex passwords can require hours, days, or longer.",
	"faq.local.q":               "Will my files leave my computer?",
	"faq.local.a":               "No. RecoverEase is built around local processing, so recovery tasks run on your computer.",
	"faq.internet.q":            "Is an internet connection required?",
	"faq.internet.a":            "The recovery process does not depend on the internet. Connectivity is only used for updates, license activation, or support.",
	"faq.safety.q":              "Is my data safe?",
	"faq.safety.a":              "The software does not upload your files, and sensitive recovery tasks can run in an offline environment.",
	"download.eyebrow":          "Download",
	"download.title":            "Start Regaining Access to Your Files",
	"download.subtitle":         "A modern password recovery tool for home users, small businesses, and IT professionals.",
	"download.windows":          "Download for Windows",
	"download.mac":              "Mac Coming Soon",
	"footer.copyright":          "© 2026 RecoverEase",
	"footer.aria":               "Footer navigation",
	"footer.privacy":            "Privacy Policy",
	"footer.terms":              "Terms of Service",
	"footer.contact":            "Contact Us",
	"footer.licenseRecovery":    "Recover license key",
}

var translations = map[string]map[string]string{
	"en": en,
	"zh": merge(en, map[string]string{
		"meta.title": "RecoverEase - 密码恢复软件", "meta.description": "RecoverEase 帮助用户找回受密码保护的压缩包、文档和加密硬盘访问权限。", "brand.aria": "RecoverEase 首页", "nav.aria": "主导航", "nav.formats": "支持格式", "nav.modes": "恢复模式", "nav.pricing": "定价", "nav.download": "下载", "controls.language": "语言", "controls.theme": "切换暗色主题", "cta.freeTrial": "免费试用", "cta.buyPro": "购买专业版", "hero.eyebrow": "密码恢复软件", "hero.title": "找回加密文件访问权限", "hero.subtitle": "RecoverEase 帮助您重新获得对受密码保护的压缩包、文档和加密硬盘的访问。", "hero.actions": "主要操作", "hero.trust": "产品特点", "trust.offline": "离线处理", "trust.noUpload": "不上传文件", "trust.gpu": "GPU 加速", "appPreview.aria": "RecoverEase 软件界面预览", "app.nav.recovery": "恢复", "app.nav.files": "文件", "app.nav.hints": "线索", "app.nav.reports": "报告", "app.currentTask": "当前任务", "app.archiveRecovery": "压缩包恢复", "app.smartMode": "智能模式", "app.encryptedDetected": "检测到加密压缩包", "app.passwordHints": "密码线索", "app.hintSummer": "夏天", "app.hintYear": "2024", "app.hintCompany": "公司", "app.gpuSpeed": "GPU 速度", "app.local": "本地", "app.privacyMode": "隐私模式", "formats.eyebrow": "支持格式", "formats.title": "覆盖常见加密文件与磁盘容器", "formats.subtitle": "从家庭备份到企业资料归档，RecoverEase 用统一流程处理不同类型的受保护文件。", "formats.archive.title": "压缩文件", "formats.archive.desc": "支持 RAR、ZIP、7Z 等常见压缩包恢复场景。", "formats.office.title": "Office 文档", "formats.office.desc": "处理 Microsoft Word、Excel、PowerPoint 受密码保护文件。", "formats.pdf.title": "PDF 文件", "formats.pdf.desc": "帮助找回 PDF 打开密码，恢复重要资料访问权限。", "formats.disk.title": "磁盘加密", "formats.disk.desc": "面向 BitLocker 与 VeraCrypt 容器的专业恢复任务。", "modes.eyebrow": "恢复模式", "modes.title": "从简单尝试到专家级策略", "modes.subtitle": "根据您掌握的密码线索、文件类型和时间预算，选择合适的恢复方式。", "modes.quick.title": "快速恢复", "modes.quick.desc": "使用常见密码和轻量规则快速尝试。", "modes.quick.fit": "适合：家庭用户、刚忘记的密码", "modes.smart.title": "智能恢复", "modes.smart.desc": "结合文件特征、常用密码库和您提供的线索自动生成策略。", "modes.smart.fit": "适合：大多数恢复任务", "modes.deep.title": "深度恢复", "modes.deep.desc": "扩大搜索范围，适合密码复杂、线索较少的情况。", "modes.deep.fit": "适合：企业归档、长期遗忘密码", "modes.expert.title": "高级/专家模式", "modes.expert.desc": "自定义长度、字符集、规则和任务队列。", "modes.expert.fit": "适合：IT 专业人士", "labels.recommended": "推荐", "labels.bestValue": "最优选择", "advantages.eyebrow": "软件优势", "advantages.title": "专业可信，同时尊重隐私", "advantages.subtitle": "RecoverEase 专注本地恢复流程，适合个人文件、小型企业资产和 IT 支持场景。", "advantages.gpu.title": "GPU 高速加速", "advantages.gpu.desc": "充分利用本机显卡提升恢复速度。", "advantages.hints.title": "智能密码线索", "advantages.hints.desc": "把姓名、日期、关键词等线索转化为候选策略。", "advantages.offline.title": "离线处理", "advantages.offline.desc": "核心恢复任务在本机完成，无需云端排队。", "advantages.privacy.title": "注重隐私", "advantages.privacy.desc": "敏感文件不上传，任务过程更可控。", "advantages.noUpload.title": "不上传文件", "advantages.noUpload.desc": "文件始终保留在您的电脑中。", "workflow.eyebrow": "恢复流程", "workflow.title": "四步开始恢复", "workflow.step1.title": "选择文件", "workflow.step1.desc": "添加需要恢复访问权限的压缩包、文档或加密容器。", "workflow.step2.title": "选择恢复模式", "workflow.step2.desc": "从快速、智能、深度或专家模式中选择任务策略。", "workflow.step3.title": "添加密码线索", "workflow.step3.desc": "输入可能出现的姓名、日期、词组或格式规律。", "workflow.step4.title": "开始恢复", "workflow.step4.desc": "查看实时进度、速度和任务结果。", "pricing.eyebrow": "定价", "pricing.title": "简单透明的授权方案", "pricing.free.title": "免费版", "pricing.free.item1": "文件分析", "pricing.free.item2": "快速恢复", "pricing.free.item3": "常用密码字典", "pricing.pro.title": "专业版", "pricing.pro.note": "一次性永久授权", "pricing.pro.item1": "无限恢复", "pricing.pro.item2": "深度恢复", "pricing.pro.item3": "GPU 加速", "pricing.pro.item4": "批量任务", "pricing.pro.item5": "优先客服支持", "faq.eyebrow": "常见问题", "faq.title": "常见问题", "faq.allPasswords.q": "所有密码都能找回吗？", "faq.allPasswords.a": "不能保证。恢复结果取决于密码复杂度、可用线索、文件类型和可投入的计算时间。", "faq.duration.q": "破解需要多久？", "faq.duration.a": "简单密码可能几分钟内完成，复杂密码可能需要数小时、数天甚至更久。", "faq.local.q": "文件会离开我的电脑吗？", "faq.local.a": "不会。RecoverEase 以本地处理为核心，恢复任务在您的电脑上运行。", "faq.internet.q": "需要联网吗？", "faq.internet.a": "恢复过程不依赖联网。联网只用于下载更新、激活授权或联系客服。", "faq.safety.q": "数据安全吗？", "faq.safety.a": "软件不上传您的文件，您可以在离线环境中执行敏感恢复任务。", "download.eyebrow": "下载", "download.title": "开始恢复您的文件访问权限", "download.subtitle": "适用于家庭用户、小型企业和 IT 专业人士的现代密码恢复工具。", "download.windows": "Windows 下载", "download.mac": "Mac 下载（即将上线）", "footer.copyright": "© 2026 RecoverEase", "footer.aria": "页脚导航", "footer.privacy": "隐私政策", "footer.terms": "服务条款", "footer.contact": "联系我们",
	}),
	"ja": merge(en, map[string]string{
		"meta.title": "RecoverEase - パスワード復元ソフト", "meta.description": "RecoverEase は、パスワード保護された圧縮ファイル、文書、暗号化ドライブへのアクセス回復を支援します。", "brand.aria": "RecoverEase ホーム", "nav.aria": "メインナビゲーション", "nav.formats": "対応形式", "nav.modes": "復元モード", "nav.pricing": "価格", "nav.download": "ダウンロード", "controls.language": "言語", "controls.theme": "ダークテーマに切り替え", "cta.freeTrial": "無料で試す", "cta.buyPro": "Pro を購入", "hero.eyebrow": "パスワード復元ソフト", "hero.title": "暗号化ファイルへのアクセスを取り戻す", "hero.subtitle": "RecoverEase は、パスワード保護された圧縮ファイル、文書、暗号化ドライブへのアクセス回復を支援します。", "hero.actions": "主な操作", "hero.trust": "製品の特長", "trust.offline": "オフライン処理", "trust.noUpload": "ファイルをアップロードしない", "trust.gpu": "GPU 高速化", "appPreview.aria": "RecoverEase ソフトウェア画面プレビュー", "app.nav.recovery": "復元", "app.nav.files": "ファイル", "app.nav.hints": "ヒント", "app.nav.reports": "レポート", "app.currentTask": "現在のタスク", "app.archiveRecovery": "アーカイブ復元", "app.smartMode": "スマートモード", "app.encryptedDetected": "暗号化アーカイブを検出", "app.passwordHints": "パスワードのヒント", "app.hintSummer": "夏", "app.hintYear": "2024", "app.hintCompany": "会社", "app.gpuSpeed": "GPU 速度", "app.local": "ローカル", "app.privacyMode": "プライバシーモード", "formats.eyebrow": "対応形式", "formats.title": "一般的な暗号化ファイルとディスクコンテナに対応", "formats.subtitle": "家庭のバックアップからビジネスアーカイブまで、RecoverEase は保護されたファイルを明確な手順で処理します。", "formats.archive.title": "圧縮ファイル", "formats.archive.desc": "RAR、ZIP、7Z などの一般的な圧縮ファイルに対応します。", "formats.office.title": "Office 文書", "formats.office.desc": "Microsoft Word、Excel、PowerPoint のパスワード保護ファイルを処理します。", "formats.pdf.title": "PDF ファイル", "formats.pdf.desc": "PDF のオープンパスワード回復を支援します。", "formats.disk.title": "ディスク暗号化", "formats.disk.desc": "BitLocker と VeraCrypt コンテナ向けの専門的な復元タスク。", "modes.eyebrow": "復元モード", "modes.title": "簡単な試行から専門的な戦略まで", "modes.subtitle": "パスワードの手がかり、ファイル形式、時間の余裕に合わせて復元方法を選べます。", "modes.quick.title": "クイック復元", "modes.quick.desc": "一般的なパスワードと軽量ルールを先に試します。", "modes.quick.fit": "対象：家庭ユーザー、最近忘れたパスワード", "modes.smart.title": "スマート復元", "modes.smart.desc": "ファイル情報、パスワード辞書、手がかりから戦略を自動作成します。", "modes.smart.fit": "対象：ほとんどの復元タスク", "modes.deep.title": "ディープ復元", "modes.deep.desc": "複雑なパスワードや少ない手がかりに向けて検索範囲を広げます。", "modes.deep.fit": "対象：業務アーカイブ、長期間忘れたパスワード", "modes.expert.title": "エキスパート復元", "modes.expert.desc": "長さ、文字セット、ルール、タスクキューをカスタマイズできます。", "modes.expert.fit": "対象：IT プロフェッショナル", "labels.recommended": "おすすめ", "labels.bestValue": "最適な選択", "advantages.eyebrow": "利点", "advantages.title": "プロ品質、信頼性、プライバシー重視", "advantages.subtitle": "RecoverEase は個人ファイル、小規模ビジネス資産、IT サポート向けのローカル復元に集中します。", "advantages.gpu.title": "GPU 高速化", "advantages.gpu.desc": "ローカル GPU を活用して復元速度を高めます。", "advantages.hints.title": "スマートなパスワードヒント", "advantages.hints.desc": "名前、日付、キーワード、パターンを候補戦略へ変換します。", "advantages.offline.title": "オフライン処理", "advantages.offline.desc": "主要な復元タスクはクラウド待ちなしでローカル実行されます。", "advantages.privacy.title": "プライバシー優先", "advantages.privacy.desc": "機密ファイルはアップロードされず、処理を手元で管理できます。", "advantages.noUpload.title": "ファイルをアップロードしない", "advantages.noUpload.desc": "ファイルは常にお使いのコンピューター上に残ります。", "workflow.eyebrow": "流れ", "workflow.title": "4 ステップで復元開始", "workflow.step1.title": "ファイルを選択", "workflow.step1.desc": "復元したい圧縮ファイル、文書、暗号化コンテナを追加します。", "workflow.step2.title": "モードを選択", "workflow.step2.desc": "クイック、スマート、ディープ、エキスパートから選びます。", "workflow.step3.title": "ヒントを追加", "workflow.step3.desc": "名前、日付、語句、パスワード形式を入力します。", "workflow.step4.title": "復元を開始", "workflow.step4.desc": "進行状況、速度、結果をリアルタイムで確認できます。", "pricing.eyebrow": "価格", "pricing.title": "シンプルで透明なライセンス", "pricing.free.title": "無料版", "pricing.free.item1": "ファイル分析", "pricing.free.item2": "クイック復元", "pricing.free.item3": "一般的なパスワード辞書", "pricing.pro.title": "Pro", "pricing.pro.note": "買い切り永久ライセンス", "pricing.pro.item1": "無制限復元", "pricing.pro.item2": "ディープ復元", "pricing.pro.item3": "GPU 高速化", "pricing.pro.item4": "一括タスク", "pricing.pro.item5": "優先サポート", "faq.eyebrow": "FAQ", "faq.title": "よくある質問", "faq.allPasswords.q": "すべてのパスワードを復元できますか？", "faq.allPasswords.a": "保証はできません。結果はパスワードの複雑さ、手がかり、ファイル形式、計算時間に依存します。", "faq.duration.q": "復元にはどのくらい時間がかかりますか？", "faq.duration.a": "単純なパスワードは数分、複雑なものは数時間、数日、またはそれ以上かかる場合があります。", "faq.local.q": "ファイルはコンピューター外へ送られますか？", "faq.local.a": "いいえ。RecoverEase はローカル処理を中心に設計されています。", "faq.internet.q": "インターネット接続は必要ですか？", "faq.internet.a": "復元処理にインターネットは不要です。更新、ライセンス認証、サポート時のみ使用します。", "faq.safety.q": "データは安全ですか？", "faq.safety.a": "ソフトウェアはファイルをアップロードせず、オフライン環境でも機密タスクを実行できます。", "download.eyebrow": "ダウンロード", "download.title": "ファイルへのアクセス回復を始めましょう", "download.subtitle": "家庭ユーザー、小規模ビジネス、IT プロ向けのモダンなパスワード復元ツール。", "download.windows": "Windows 版をダウンロード", "download.mac": "Mac 版は近日公開", "footer.copyright": "© 2026 RecoverEase", "footer.aria": "フッターナビゲーション", "footer.privacy": "プライバシーポリシー", "footer.terms": "利用規約", "footer.contact": "お問い合わせ",
	}),
	"ko": merge(en, map[string]string{
		"meta.title": "RecoverEase - 비밀번호 복구 소프트웨어", "meta.description": "RecoverEase는 비밀번호로 보호된 압축 파일, 문서, 암호화 드라이브 접근 권한을 되찾도록 돕습니다.", "brand.aria": "RecoverEase 홈", "nav.aria": "기본 탐색", "nav.formats": "지원 형식", "nav.modes": "복구 모드", "nav.pricing": "가격", "nav.download": "다운로드", "controls.language": "언어", "controls.theme": "다크 테마로 전환", "cta.freeTrial": "무료 체험", "cta.buyPro": "Pro 구매", "hero.eyebrow": "비밀번호 복구 소프트웨어", "hero.title": "암호화된 파일 접근 권한 복구", "hero.subtitle": "RecoverEase는 비밀번호로 보호된 압축 파일, 문서, 암호화 드라이브에 다시 접근하도록 돕습니다.", "hero.actions": "주요 작업", "hero.trust": "제품 특징", "trust.offline": "오프라인 처리", "trust.noUpload": "파일 업로드 없음", "trust.gpu": "GPU 가속", "appPreview.aria": "RecoverEase 소프트웨어 화면 미리보기", "app.nav.recovery": "복구", "app.nav.files": "파일", "app.nav.hints": "단서", "app.nav.reports": "보고서", "app.currentTask": "현재 작업", "app.archiveRecovery": "압축 파일 복구", "app.smartMode": "스마트 모드", "app.encryptedDetected": "암호화된 압축 파일 감지됨", "app.passwordHints": "비밀번호 단서", "app.hintSummer": "여름", "app.hintYear": "2024", "app.hintCompany": "회사", "app.gpuSpeed": "GPU 속도", "app.local": "로컬", "app.privacyMode": "개인정보 모드", "formats.eyebrow": "지원 형식", "formats.title": "일반 암호화 파일 및 디스크 컨테이너 지원", "formats.subtitle": "가정용 백업부터 비즈니스 아카이브까지 RecoverEase는 보호된 파일을 명확한 흐름으로 처리합니다.", "formats.archive.title": "압축 파일", "formats.archive.desc": "RAR, ZIP, 7Z 등 일반 압축 파일 복구를 지원합니다.", "formats.office.title": "Office 문서", "formats.office.desc": "Microsoft Word, Excel, PowerPoint 보호 파일을 처리합니다.", "formats.pdf.title": "PDF 파일", "formats.pdf.desc": "PDF 열기 비밀번호 복구를 도와줍니다.", "formats.disk.title": "디스크 암호화", "formats.disk.desc": "BitLocker 및 VeraCrypt 컨테이너용 전문 복구 작업.", "modes.eyebrow": "복구 모드", "modes.title": "빠른 시도부터 전문가 전략까지", "modes.subtitle": "비밀번호 단서, 파일 형식, 시간 예산에 맞는 복구 방법을 선택하세요.", "modes.quick.title": "빠른 복구", "modes.quick.desc": "일반적인 비밀번호와 가벼운 규칙을 먼저 시도합니다.", "modes.quick.fit": "적합: 가정 사용자, 최근 잊어버린 비밀번호", "modes.smart.title": "스마트 복구", "modes.smart.desc": "파일 정보, 비밀번호 사전, 입력한 단서를 바탕으로 전략을 자동 생성합니다.", "modes.smart.fit": "적합: 대부분의 복구 작업", "modes.deep.title": "심층 복구", "modes.deep.desc": "복잡한 비밀번호나 단서가 적은 경우 검색 범위를 넓힙니다.", "modes.deep.fit": "적합: 기업 아카이브, 오래전에 잊은 비밀번호", "modes.expert.title": "전문가 복구", "modes.expert.desc": "길이, 문자 집합, 규칙, 작업 대기열을 사용자 지정합니다.", "modes.expert.fit": "적합: IT 전문가", "labels.recommended": "추천", "labels.bestValue": "최고 가치", "advantages.eyebrow": "장점", "advantages.title": "전문적이고 신뢰할 수 있으며 개인정보 중심", "advantages.subtitle": "RecoverEase는 개인 파일, 소규모 비즈니스 자산, IT 지원을 위한 로컬 복구 흐름에 집중합니다.", "advantages.gpu.title": "GPU 고속 가속", "advantages.gpu.desc": "로컬 그래픽 카드를 활용해 복구 속도를 높입니다.", "advantages.hints.title": "스마트 비밀번호 단서", "advantages.hints.desc": "이름, 날짜, 키워드, 패턴을 후보 전략으로 변환합니다.", "advantages.offline.title": "오프라인 처리", "advantages.offline.desc": "핵심 복구 작업은 클라우드 대기 없이 로컬에서 실행됩니다.", "advantages.privacy.title": "개인정보 우선", "advantages.privacy.desc": "민감한 파일은 업로드되지 않으며 작업 과정을 직접 제어할 수 있습니다.", "advantages.noUpload.title": "파일 업로드 없음", "advantages.noUpload.desc": "파일은 항상 사용자의 컴퓨터에 남아 있습니다.", "workflow.eyebrow": "작업 흐름", "workflow.title": "4단계로 복구 시작", "workflow.step1.title": "파일 선택", "workflow.step1.desc": "복구할 압축 파일, 문서 또는 암호화 컨테이너를 추가합니다.", "workflow.step2.title": "모드 선택", "workflow.step2.desc": "빠른, 스마트, 심층 또는 전문가 전략을 선택합니다.", "workflow.step3.title": "단서 추가", "workflow.step3.desc": "가능한 이름, 날짜, 문구 또는 비밀번호 패턴을 입력합니다.", "workflow.step4.title": "복구 시작", "workflow.step4.desc": "진행률, 속도, 결과를 실시간으로 확인합니다.", "pricing.eyebrow": "가격", "pricing.title": "단순하고 투명한 라이선스", "pricing.free.title": "무료", "pricing.free.item1": "파일 분석", "pricing.free.item2": "빠른 복구", "pricing.free.item3": "일반 비밀번호 사전", "pricing.pro.title": "Pro", "pricing.pro.note": "일회성 평생 라이선스", "pricing.pro.item1": "무제한 복구", "pricing.pro.item2": "심층 복구", "pricing.pro.item3": "GPU 가속", "pricing.pro.item4": "일괄 작업", "pricing.pro.item5": "우선 지원", "faq.eyebrow": "FAQ", "faq.title": "자주 묻는 질문", "faq.allPasswords.q": "모든 비밀번호를 복구할 수 있나요?", "faq.allPasswords.a": "보장할 수 없습니다. 결과는 비밀번호 복잡도, 사용 가능한 단서, 파일 형식, 투입 가능한 계산 시간에 따라 달라집니다.", "faq.duration.q": "복구에는 얼마나 걸리나요?", "faq.duration.a": "간단한 비밀번호는 몇 분, 복잡한 비밀번호는 몇 시간, 며칠 또는 그 이상 걸릴 수 있습니다.", "faq.local.q": "파일이 제 컴퓨터를 떠나나요?", "faq.local.a": "아니요. RecoverEase는 로컬 처리를 중심으로 설계되어 복구 작업이 컴퓨터에서 실행됩니다.", "faq.internet.q": "인터넷 연결이 필요한가요?", "faq.internet.a": "복구 과정은 인터넷에 의존하지 않습니다. 업데이트, 라이선스 활성화, 지원에만 연결이 사용됩니다.", "faq.safety.q": "데이터는 안전한가요?", "faq.safety.a": "소프트웨어는 파일을 업로드하지 않으며 오프라인 환경에서도 민감한 복구 작업을 실행할 수 있습니다.", "download.eyebrow": "다운로드", "download.title": "파일 접근 권한 복구 시작", "download.subtitle": "가정 사용자, 소규모 비즈니스, IT 전문가를 위한 현대적인 비밀번호 복구 도구입니다.", "download.windows": "Windows 다운로드", "download.mac": "Mac 곧 출시", "footer.copyright": "© 2026 RecoverEase", "footer.aria": "푸터 탐색", "footer.privacy": "개인정보 처리방침", "footer.terms": "서비스 약관", "footer.contact": "문의하기",
	}),
}

func init() {
	addLocale("de", map[string]string{
		"meta.title": "RecoverEase - Passwort-Wiederherstellungssoftware", "meta.description": "RecoverEase hilft beim Wiederzugriff auf passwortgeschützte Archive, Dokumente und verschlüsselte Laufwerke.", "brand.aria": "RecoverEase Startseite", "nav.aria": "Hauptnavigation", "nav.formats": "Formate", "nav.modes": "Modi", "nav.pricing": "Preise", "nav.download": "Download", "controls.language": "Sprache", "controls.theme": "Dunkles Design aktivieren", "cta.freeTrial": "Kostenlos testen", "cta.buyPro": "Pro kaufen", "hero.eyebrow": "Passwort-Wiederherstellungssoftware", "hero.title": "Zugriff auf verschlüsselte Dateien zurückerlangen", "hero.subtitle": "RecoverEase hilft Ihnen, wieder auf passwortgeschützte Archive, Dokumente und verschlüsselte Laufwerke zuzugreifen.", "hero.actions": "Primäre Aktionen", "hero.trust": "Produktmerkmale", "trust.offline": "Offline-Verarbeitung", "trust.noUpload": "Keine Datei-Uploads", "trust.gpu": "GPU-Beschleunigung", "appPreview.aria": "Vorschau der RecoverEase-Oberfläche", "app.nav.recovery": "Wiederherstellung", "app.nav.files": "Dateien", "app.nav.hints": "Hinweise", "app.nav.reports": "Berichte", "app.currentTask": "Aktuelle Aufgabe", "app.archiveRecovery": "Archiv-Wiederherstellung", "app.smartMode": "Smart-Modus", "app.encryptedDetected": "Verschlüsseltes Archiv erkannt", "app.passwordHints": "Passwort-Hinweise", "app.hintSummer": "sommer", "app.hintYear": "2024", "app.hintCompany": "firma", "app.gpuSpeed": "GPU-Geschwindigkeit", "app.local": "Lokal", "app.privacyMode": "Privatsphäre-Modus", "formats.eyebrow": "Unterstützte Formate", "formats.title": "Unterstützung für gängige verschlüsselte Dateien und Container", "formats.subtitle": "Von privaten Backups bis zu Unternehmensarchiven verarbeitet RecoverEase geschützte Dateien in einem klaren Ablauf.", "formats.archive.title": "Archive", "formats.archive.desc": "Unterstützt RAR-, ZIP- und 7Z-Szenarien.", "formats.office.title": "Office-Dokumente", "formats.office.desc": "Verarbeitet geschützte Microsoft Word-, Excel- und PowerPoint-Dateien.", "formats.pdf.title": "PDF-Dateien", "formats.pdf.desc": "Hilft bei der Wiederherstellung von PDF-Öffnungspasswörtern.", "formats.disk.title": "Datenträgerverschlüsselung", "formats.disk.desc": "Professionelle Aufgaben für BitLocker- und VeraCrypt-Container.", "modes.eyebrow": "Wiederherstellungsmodi", "modes.title": "Von schnellen Versuchen bis zu Expertenstrategien", "modes.subtitle": "Wählen Sie die passende Methode nach Hinweisen, Dateityp und Zeitbudget.", "modes.quick.title": "Schnelle Wiederherstellung", "modes.quick.desc": "Testet zuerst häufige Passwörter und leichte Regeln.", "modes.quick.fit": "Geeignet für: private Nutzer und kürzlich vergessene Passwörter", "modes.smart.title": "Intelligente Wiederherstellung", "modes.smart.desc": "Erstellt automatisch eine Strategie aus Dateisignalen, Wörterbüchern und Ihren Hinweisen.", "modes.smart.fit": "Geeignet für: die meisten Aufgaben", "modes.deep.title": "Tiefe Wiederherstellung", "modes.deep.desc": "Erweitert den Suchbereich für komplexe Passwörter oder wenige Hinweise.", "modes.deep.fit": "Geeignet für: Unternehmensarchive und lange vergessene Passwörter", "modes.expert.title": "Expertenmodus", "modes.expert.desc": "Passen Sie Länge, Zeichensätze, Regeln und Aufgabenwarteschlangen an.", "modes.expert.fit": "Geeignet für: IT-Profis", "labels.recommended": "Empfohlen", "labels.bestValue": "Bester Wert", "advantages.eyebrow": "Vorteile", "advantages.title": "Professionell, vertrauenswürdig und datenschutzorientiert", "advantages.subtitle": "RecoverEase konzentriert sich auf lokale Wiederherstellung für persönliche Dateien, kleine Unternehmen und IT-Support.", "advantages.gpu.title": "GPU-Beschleunigung", "advantages.gpu.desc": "Nutzen Sie Ihre lokale Grafikkarte, um die Wiederherstellung zu beschleunigen.", "advantages.hints.title": "Intelligente Passwort-Hinweise", "advantages.hints.desc": "Verwandelt Namen, Daten, Schlüsselwörter und Muster in Kandidatenstrategien.", "advantages.offline.title": "Offline-Verarbeitung", "advantages.offline.desc": "Kernaufgaben laufen lokal ohne Cloud-Warteschlangen.", "advantages.privacy.title": "Datenschutz zuerst", "advantages.privacy.desc": "Sensible Dateien werden nicht hochgeladen und der Prozess bleibt unter Ihrer Kontrolle.", "advantages.noUpload.title": "Keine Datei-Uploads", "advantages.noUpload.desc": "Ihre Dateien bleiben jederzeit auf Ihrem Computer.", "workflow.eyebrow": "Ablauf", "workflow.title": "Wiederherstellung in vier Schritten", "workflow.step1.title": "Datei auswählen", "workflow.step1.desc": "Fügen Sie ein Archiv, Dokument oder einen verschlüsselten Container hinzu.", "workflow.step2.title": "Modus wählen", "workflow.step2.desc": "Wählen Sie eine schnelle, smarte, tiefe oder Expertenstrategie.", "workflow.step3.title": "Hinweise hinzufügen", "workflow.step3.desc": "Geben Sie wahrscheinliche Namen, Daten, Phrasen oder Muster ein.", "workflow.step4.title": "Wiederherstellung starten", "workflow.step4.desc": "Verfolgen Sie Fortschritt, Geschwindigkeit und Ergebnisse in Echtzeit.", "pricing.eyebrow": "Preise", "pricing.title": "Einfache, transparente Lizenzierung", "pricing.free.title": "Kostenlos", "pricing.free.item1": "Dateianalyse", "pricing.free.item2": "Schnelle Wiederherstellung", "pricing.free.item3": "Häufiges Passwort-Wörterbuch", "pricing.pro.title": "Pro", "pricing.pro.note": "einmalige lebenslange Lizenz", "pricing.pro.item1": "Unbegrenzte Wiederherstellung", "pricing.pro.item2": "Tiefe Wiederherstellung", "pricing.pro.item3": "GPU-Beschleunigung", "pricing.pro.item4": "Stapelaufgaben", "pricing.pro.item5": "Priorisierter Support", "faq.eyebrow": "FAQ", "faq.title": "Häufige Fragen", "faq.allPasswords.q": "Kann jedes Passwort wiederhergestellt werden?", "faq.allPasswords.a": "Nein. Das Ergebnis hängt von Komplexität, Hinweisen, Dateityp und verfügbarer Rechenzeit ab.", "faq.duration.q": "Wie lange dauert die Wiederherstellung?", "faq.duration.a": "Einfache Passwörter können Minuten dauern, komplexe Passwörter Stunden, Tage oder länger.", "faq.local.q": "Verlassen meine Dateien den Computer?", "faq.local.a": "Nein. RecoverEase ist auf lokale Verarbeitung ausgelegt.", "faq.internet.q": "Ist Internet erforderlich?", "faq.internet.a": "Die Wiederherstellung benötigt kein Internet. Verbindung wird nur für Updates, Aktivierung oder Support genutzt.", "faq.safety.q": "Sind meine Daten sicher?", "faq.safety.a": "Die Software lädt Ihre Dateien nicht hoch und kann sensible Aufgaben offline ausführen.", "download.eyebrow": "Herunterladen", "download.title": "Beginnen Sie mit der Wiederherstellung Ihres Dateizugriffs", "download.subtitle": "Ein modernes Passwort-Wiederherstellungstool für private Nutzer, kleine Unternehmen und IT-Profis.", "download.windows": "Für Windows herunterladen", "download.mac": "Mac bald verfügbar", "footer.copyright": "© 2026 RecoverEase", "footer.aria": "Fußnavigation", "footer.privacy": "Datenschutz", "footer.terms": "Nutzungsbedingungen", "footer.contact": "Kontakt",
	})
	addLocale("fr", map[string]string{
		"meta.title": "RecoverEase - Logiciel de récupération de mots de passe", "meta.description": "RecoverEase vous aide à retrouver l'accès aux archives, documents et disques chiffrés protégés par mot de passe.", "brand.aria": "Accueil RecoverEase", "nav.aria": "Navigation principale", "nav.formats": "Formats pris en charge", "nav.modes": "Modes", "nav.pricing": "Tarifs", "nav.download": "Télécharger", "controls.language": "Langue", "controls.theme": "Activer le thème sombre", "cta.freeTrial": "Essai gratuit", "cta.buyPro": "Acheter Pro", "hero.eyebrow": "Logiciel de récupération de mots de passe", "hero.title": "Retrouvez l'accès à vos fichiers chiffrés", "hero.subtitle": "RecoverEase vous aide à retrouver l'accès aux archives, documents et disques chiffrés protégés par mot de passe.", "hero.actions": "Actions principales", "hero.trust": "Points forts du produit", "trust.offline": "Traitement hors ligne", "trust.noUpload": "Aucun envoi de fichier", "trust.gpu": "Accélération GPU", "appPreview.aria": "Aperçu de l'interface RecoverEase", "app.nav.recovery": "Récupération", "app.nav.files": "Fichiers", "app.nav.hints": "Indices", "app.nav.reports": "Rapports", "app.currentTask": "Tâche actuelle", "app.archiveRecovery": "Récupération d'archive", "app.smartMode": "Mode intelligent", "app.encryptedDetected": "Archive chiffrée détectée", "app.passwordHints": "Indices de mot de passe", "app.hintSummer": "été", "app.hintYear": "2024", "app.hintCompany": "entreprise", "app.gpuSpeed": "Vitesse GPU", "app.local": "Local", "app.privacyMode": "Mode confidentialité", "formats.eyebrow": "Formats pris en charge", "formats.title": "Prise en charge des fichiers chiffrés et conteneurs courants", "formats.subtitle": "Des sauvegardes personnelles aux archives d'entreprise, RecoverEase traite les fichiers protégés dans un flux clair.", "formats.archive.title": "Archives", "formats.archive.desc": "Prend en charge les scénarios RAR, ZIP et 7Z courants.", "formats.office.title": "Documents Office", "formats.office.desc": "Traite les fichiers Microsoft Word, Excel et PowerPoint protégés.", "formats.pdf.title": "Fichiers PDF", "formats.pdf.desc": "Aide à récupérer les mots de passe d'ouverture PDF.", "formats.disk.title": "Chiffrement de disque", "formats.disk.desc": "Tâches professionnelles pour conteneurs BitLocker et VeraCrypt.", "modes.eyebrow": "Modes de récupération", "modes.title": "Des essais rapides aux stratégies expertes", "modes.subtitle": "Choisissez la bonne méthode selon les indices, le type de fichier et le temps disponible.", "modes.quick.title": "Récupération rapide", "modes.quick.desc": "Essaie d'abord les mots de passe courants et les règles légères.", "modes.quick.fit": "Idéal pour : particuliers et mots de passe récemment oubliés", "modes.smart.title": "Récupération intelligente", "modes.smart.desc": "Crée automatiquement une stratégie à partir du fichier, des dictionnaires et de vos indices.", "modes.smart.fit": "Idéal pour : la plupart des tâches", "modes.deep.title": "Récupération approfondie", "modes.deep.desc": "Élargit la recherche pour les mots de passe complexes ou peu d'indices.", "modes.deep.fit": "Idéal pour : archives d'entreprise et anciens mots de passe", "modes.expert.title": "Mode expert", "modes.expert.desc": "Personnalisez longueur, jeux de caractères, règles et files de tâches.", "modes.expert.fit": "Idéal pour : professionnels IT", "labels.recommended": "Recommandé", "labels.bestValue": "Meilleur choix", "advantages.eyebrow": "Avantages", "advantages.title": "Professionnel, fiable et respectueux de la vie privée", "advantages.subtitle": "RecoverEase se concentre sur la récupération locale pour fichiers personnels, petites entreprises et support IT.", "advantages.gpu.title": "Accélération GPU", "advantages.gpu.desc": "Utilisez votre carte graphique locale pour améliorer la vitesse.", "advantages.hints.title": "Indices intelligents", "advantages.hints.desc": "Transforme noms, dates, mots-clés et modèles en stratégies candidates.", "advantages.offline.title": "Traitement hors ligne", "advantages.offline.desc": "Les tâches principales s'exécutent localement sans file d'attente cloud.", "advantages.privacy.title": "Confidentialité d'abord", "advantages.privacy.desc": "Les fichiers sensibles ne sont pas envoyés et le processus reste sous votre contrôle.", "advantages.noUpload.title": "Aucun envoi de fichier", "advantages.noUpload.desc": "Vos fichiers restent toujours sur votre ordinateur.", "workflow.eyebrow": "Flux", "workflow.title": "Démarrez en quatre étapes", "workflow.step1.title": "Choisir un fichier", "workflow.step1.desc": "Ajoutez une archive, un document ou un conteneur chiffré.", "workflow.step2.title": "Choisir un mode", "workflow.step2.desc": "Choisissez une stratégie rapide, intelligente, approfondie ou experte.", "workflow.step3.title": "Ajouter des indices", "workflow.step3.desc": "Saisissez noms, dates, phrases ou modèles probables.", "workflow.step4.title": "Lancer la récupération", "workflow.step4.desc": "Suivez la progression, la vitesse et les résultats en temps réel.", "pricing.eyebrow": "Tarifs", "pricing.title": "Licence simple et transparente", "pricing.free.title": "Gratuit", "pricing.free.item1": "Analyse de fichier", "pricing.free.item2": "Récupération rapide", "pricing.free.item3": "Dictionnaire courant", "pricing.pro.title": "Pro", "pricing.pro.note": "licence à vie en paiement unique", "pricing.pro.item1": "Récupération illimitée", "pricing.pro.item2": "Récupération approfondie", "pricing.pro.item3": "Accélération GPU", "pricing.pro.item4": "Tâches en lot", "pricing.pro.item5": "Support prioritaire", "faq.eyebrow": "FAQ", "faq.title": "Questions fréquentes", "faq.allPasswords.q": "Tous les mots de passe peuvent-ils être récupérés ?", "faq.allPasswords.a": "Non. Le résultat dépend de la complexité, des indices, du type de fichier et du temps de calcul.", "faq.duration.q": "Combien de temps cela prend-il ?", "faq.duration.a": "Un mot de passe simple peut prendre quelques minutes ; un mot de passe complexe peut prendre des heures, des jours ou plus.", "faq.local.q": "Mes fichiers quittent-ils mon ordinateur ?", "faq.local.a": "Non. RecoverEase est conçu autour du traitement local.", "faq.internet.q": "Internet est-il nécessaire ?", "faq.internet.a": "La récupération ne dépend pas d'Internet. La connexion sert aux mises à jour, à l'activation ou au support.", "faq.safety.q": "Mes données sont-elles sûres ?", "faq.safety.a": "Le logiciel n'envoie pas vos fichiers et peut fonctionner hors ligne.", "download.eyebrow": "Télécharger", "download.title": "Commencez à récupérer l'accès à vos fichiers", "download.subtitle": "Un outil moderne pour particuliers, petites entreprises et professionnels IT.", "download.windows": "Télécharger pour Windows", "download.mac": "Mac bientôt disponible", "footer.copyright": "© 2026 RecoverEase", "footer.aria": "Navigation de pied de page", "footer.privacy": "Politique de confidentialité", "footer.terms": "Conditions d'utilisation", "footer.contact": "Nous contacter",
	})
	addLocale("es", map[string]string{
		"meta.title": "RecoverEase - Software de recuperación de contraseñas", "meta.description": "RecoverEase ayuda a recuperar el acceso a archivos comprimidos, documentos y unidades cifradas protegidos por contraseña.", "brand.aria": "Inicio de RecoverEase", "nav.aria": "Navegación principal", "nav.formats": "Formatos", "nav.modes": "Modos", "nav.pricing": "Precios", "nav.download": "Descargar", "controls.language": "Idioma", "controls.theme": "Cambiar a tema oscuro", "cta.freeTrial": "Prueba gratis", "cta.buyPro": "Comprar Pro", "hero.eyebrow": "Software de recuperación de contraseñas", "hero.title": "Recupera el acceso a archivos cifrados", "hero.subtitle": "RecoverEase te ayuda a volver a acceder a archivos comprimidos, documentos y unidades cifradas protegidos por contraseña.", "hero.actions": "Acciones principales", "hero.trust": "Aspectos del producto", "trust.offline": "Procesamiento sin conexión", "trust.noUpload": "Sin subir archivos", "trust.gpu": "Aceleración GPU", "appPreview.aria": "Vista previa de la interfaz de RecoverEase", "app.nav.recovery": "Recuperación", "app.nav.files": "Archivos", "app.nav.hints": "Pistas", "app.nav.reports": "Informes", "app.currentTask": "Tarea actual", "app.archiveRecovery": "Recuperación de archivo", "app.smartMode": "Modo inteligente", "app.encryptedDetected": "Archivo cifrado detectado", "app.passwordHints": "Pistas de contraseña", "app.hintSummer": "verano", "app.hintYear": "2024", "app.hintCompany": "empresa", "app.gpuSpeed": "Velocidad GPU", "app.local": "Local", "app.privacyMode": "Modo privacidad", "formats.eyebrow": "Formatos compatibles", "formats.title": "Compatible con archivos cifrados y contenedores comunes", "formats.subtitle": "Desde copias domésticas hasta archivos empresariales, RecoverEase procesa archivos protegidos con un flujo claro.", "formats.archive.title": "Archivos comprimidos", "formats.archive.desc": "Compatible con escenarios RAR, ZIP y 7Z.", "formats.office.title": "Documentos Office", "formats.office.desc": "Trabaja con archivos Microsoft Word, Excel y PowerPoint protegidos.", "formats.pdf.title": "Archivos PDF", "formats.pdf.desc": "Ayuda a recuperar contraseñas de apertura de PDF.", "formats.disk.title": "Cifrado de disco", "formats.disk.desc": "Tareas profesionales para contenedores BitLocker y VeraCrypt.", "modes.eyebrow": "Modos de recuperación", "modes.title": "De intentos rápidos a estrategias expertas", "modes.subtitle": "Elige el método adecuado según pistas, tipo de archivo y tiempo disponible.", "modes.quick.title": "Recuperación rápida", "modes.quick.desc": "Prueba primero contraseñas comunes y reglas ligeras.", "modes.quick.fit": "Ideal para: usuarios domésticos y contraseñas olvidadas recientemente", "modes.smart.title": "Recuperación inteligente", "modes.smart.desc": "Crea automáticamente una estrategia a partir del archivo, diccionarios y tus pistas.", "modes.smart.fit": "Ideal para: la mayoría de tareas", "modes.deep.title": "Recuperación profunda", "modes.deep.desc": "Amplía la búsqueda para contraseñas complejas o pocas pistas.", "modes.deep.fit": "Ideal para: archivos empresariales y contraseñas antiguas", "modes.expert.title": "Modo experto", "modes.expert.desc": "Personaliza longitud, conjuntos de caracteres, reglas y colas.", "modes.expert.fit": "Ideal para: profesionales de IT", "labels.recommended": "Recomendado", "labels.bestValue": "Mejor opción", "advantages.eyebrow": "Ventajas", "advantages.title": "Profesional, fiable y centrado en la privacidad", "advantages.subtitle": "RecoverEase se centra en flujos locales para archivos personales, pequeñas empresas y soporte IT.", "advantages.gpu.title": "Aceleración GPU", "advantages.gpu.desc": "Usa tu tarjeta gráfica local para mejorar la velocidad.", "advantages.hints.title": "Pistas inteligentes", "advantages.hints.desc": "Convierte nombres, fechas, palabras clave y patrones en estrategias candidatas.", "advantages.offline.title": "Procesamiento sin conexión", "advantages.offline.desc": "Las tareas principales se ejecutan localmente sin colas en la nube.", "advantages.privacy.title": "Privacidad primero", "advantages.privacy.desc": "Los archivos sensibles no se suben y el proceso queda bajo tu control.", "advantages.noUpload.title": "Sin subir archivos", "advantages.noUpload.desc": "Tus archivos permanecen siempre en tu ordenador.", "workflow.eyebrow": "Flujo", "workflow.title": "Empieza en cuatro pasos", "workflow.step1.title": "Elegir archivo", "workflow.step1.desc": "Añade un archivo comprimido, documento o contenedor cifrado.", "workflow.step2.title": "Elegir modo", "workflow.step2.desc": "Elige estrategia rápida, inteligente, profunda o experta.", "workflow.step3.title": "Añadir pistas", "workflow.step3.desc": "Introduce nombres, fechas, frases o patrones probables.", "workflow.step4.title": "Iniciar recuperación", "workflow.step4.desc": "Sigue progreso, velocidad y resultados en tiempo real.", "pricing.eyebrow": "Precios", "pricing.title": "Licencia simple y transparente", "pricing.free.title": "Gratis", "pricing.free.item1": "Análisis de archivo", "pricing.free.item2": "Recuperación rápida", "pricing.free.item3": "Diccionario común", "pricing.pro.title": "Pro", "pricing.pro.note": "licencia vitalicia de pago único", "pricing.pro.item1": "Recuperación ilimitada", "pricing.pro.item2": "Recuperación profunda", "pricing.pro.item3": "Aceleración GPU", "pricing.pro.item4": "Tareas por lotes", "pricing.pro.item5": "Soporte prioritario", "faq.eyebrow": "FAQ", "faq.title": "Preguntas frecuentes", "faq.allPasswords.q": "¿Se pueden recuperar todas las contraseñas?", "faq.allPasswords.a": "No. El resultado depende de la complejidad, las pistas, el tipo de archivo y el tiempo de cálculo.", "faq.duration.q": "¿Cuánto tarda?", "faq.duration.a": "Una contraseña simple puede tardar minutos; una compleja puede tardar horas, días o más.", "faq.local.q": "¿Mis archivos salen de mi ordenador?", "faq.local.a": "No. RecoverEase está diseñado para procesamiento local.", "faq.internet.q": "¿Se necesita internet?", "faq.internet.a": "La recuperación no depende de internet. La conexión solo se usa para actualizaciones, activación o soporte.", "faq.safety.q": "¿Mis datos están seguros?", "faq.safety.a": "El software no sube tus archivos y puede ejecutar tareas sensibles sin conexión.", "download.eyebrow": "Descargar", "download.title": "Empieza a recuperar el acceso a tus archivos", "download.subtitle": "Una herramienta moderna para usuarios domésticos, pequeñas empresas y profesionales IT.", "download.windows": "Descargar para Windows", "download.mac": "Mac próximamente", "footer.copyright": "© 2026 RecoverEase", "footer.aria": "Navegación de pie de página", "footer.privacy": "Política de privacidad", "footer.terms": "Términos del servicio", "footer.contact": "Contacto",
	})
	addLocale("pt", map[string]string{
		"meta.title": "RecoverEase - Software de recuperação de senhas", "meta.description": "RecoverEase ajuda você a recuperar o acesso a arquivos compactados, documentos e unidades criptografadas protegidos por senha.", "brand.aria": "Página inicial do RecoverEase", "nav.aria": "Navegação principal", "nav.formats": "Formatos", "nav.modes": "Modos", "nav.pricing": "Preços", "nav.download": "Baixar", "controls.language": "Idioma", "controls.theme": "Alternar para tema escuro", "cta.freeTrial": "Teste grátis", "cta.buyPro": "Comprar Pro", "hero.eyebrow": "Software de recuperação de senhas", "hero.title": "Recupere o acesso a arquivos criptografados", "hero.subtitle": "RecoverEase ajuda você a acessar novamente arquivos compactados, documentos e unidades criptografadas protegidos por senha.", "hero.actions": "Ações principais", "hero.trust": "Destaques do produto", "trust.offline": "Processamento offline", "trust.noUpload": "Sem envio de arquivos", "trust.gpu": "Aceleração por GPU", "appPreview.aria": "Prévia da interface do RecoverEase", "app.nav.recovery": "Recuperação", "app.nav.files": "Arquivos", "app.nav.hints": "Pistas", "app.nav.reports": "Relatórios", "app.currentTask": "Tarefa atual", "app.archiveRecovery": "Recuperação de arquivo", "app.smartMode": "Modo inteligente", "app.encryptedDetected": "Arquivo criptografado detectado", "app.passwordHints": "Pistas de senha", "app.hintSummer": "verão", "app.hintYear": "2024", "app.hintCompany": "empresa", "app.gpuSpeed": "Velocidade GPU", "app.local": "Local", "app.privacyMode": "Modo privacidade", "formats.eyebrow": "Formatos compatíveis", "formats.title": "Suporte a arquivos criptografados e contêineres comuns", "formats.subtitle": "De backups domésticos a arquivos empresariais, o RecoverEase processa arquivos protegidos em um fluxo claro.", "formats.archive.title": "Arquivos compactados", "formats.archive.desc": "Compatível com cenários RAR, ZIP e 7Z.", "formats.office.title": "Documentos Office", "formats.office.desc": "Trabalha com arquivos Microsoft Word, Excel e PowerPoint protegidos.", "formats.pdf.title": "Arquivos PDF", "formats.pdf.desc": "Ajuda a recuperar senhas de abertura de PDF.", "formats.disk.title": "Criptografia de disco", "formats.disk.desc": "Tarefas profissionais para contêineres BitLocker e VeraCrypt.", "modes.eyebrow": "Modos de recuperação", "modes.title": "De tentativas rápidas a estratégias avançadas", "modes.subtitle": "Escolha o método certo com base nas pistas, tipo de arquivo e tempo disponível.", "modes.quick.title": "Recuperação rápida", "modes.quick.desc": "Tenta primeiro senhas comuns e regras leves.", "modes.quick.fit": "Ideal para: usuários domésticos e senhas esquecidas recentemente", "modes.smart.title": "Recuperação inteligente", "modes.smart.desc": "Cria automaticamente uma estratégia com sinais do arquivo, dicionários e suas pistas.", "modes.smart.fit": "Ideal para: a maioria das tarefas", "modes.deep.title": "Recuperação profunda", "modes.deep.desc": "Amplia a busca para senhas complexas ou poucas pistas.", "modes.deep.fit": "Ideal para: arquivos empresariais e senhas antigas", "modes.expert.title": "Modo especialista", "modes.expert.desc": "Personalize comprimento, conjuntos de caracteres, regras e filas.", "modes.expert.fit": "Ideal para: profissionais de TI", "labels.recommended": "Recomendado", "labels.bestValue": "Melhor valor", "advantages.eyebrow": "Vantagens", "advantages.title": "Profissional, confiável e focado em privacidade", "advantages.subtitle": "O RecoverEase foca em recuperação local para arquivos pessoais, pequenas empresas e suporte de TI.", "advantages.gpu.title": "Aceleração por GPU", "advantages.gpu.desc": "Use sua placa gráfica local para melhorar a velocidade.", "advantages.hints.title": "Pistas inteligentes", "advantages.hints.desc": "Transforma nomes, datas, palavras-chave e padrões em estratégias candidatas.", "advantages.offline.title": "Processamento offline", "advantages.offline.desc": "As tarefas principais rodam localmente sem filas na nuvem.", "advantages.privacy.title": "Privacidade em primeiro lugar", "advantages.privacy.desc": "Arquivos sensíveis não são enviados e o processo fica sob seu controle.", "advantages.noUpload.title": "Sem envio de arquivos", "advantages.noUpload.desc": "Seus arquivos permanecem sempre no seu computador.", "workflow.eyebrow": "Fluxo", "workflow.title": "Comece em quatro etapas", "workflow.step1.title": "Escolher arquivo", "workflow.step1.desc": "Adicione um arquivo compactado, documento ou contêiner criptografado.", "workflow.step2.title": "Escolher modo", "workflow.step2.desc": "Escolha estratégia rápida, inteligente, profunda ou especialista.", "workflow.step3.title": "Adicionar pistas", "workflow.step3.desc": "Informe nomes, datas, frases ou padrões prováveis.", "workflow.step4.title": "Iniciar recuperação", "workflow.step4.desc": "Acompanhe progresso, velocidade e resultados em tempo real.", "pricing.eyebrow": "Preços", "pricing.title": "Licenciamento simples e transparente", "pricing.free.title": "Grátis", "pricing.free.item1": "Análise de arquivo", "pricing.free.item2": "Recuperação rápida", "pricing.free.item3": "Dicionário comum", "pricing.pro.title": "Pro", "pricing.pro.note": "licença vitalícia em pagamento único", "pricing.pro.item1": "Recuperação ilimitada", "pricing.pro.item2": "Recuperação profunda", "pricing.pro.item3": "Aceleração GPU", "pricing.pro.item4": "Tarefas em lote", "pricing.pro.item5": "Suporte prioritário", "faq.eyebrow": "FAQ", "faq.title": "Perguntas frequentes", "faq.allPasswords.q": "Todas as senhas podem ser recuperadas?", "faq.allPasswords.a": "Não. O resultado depende da complexidade, pistas disponíveis, tipo de arquivo e tempo de computação.", "faq.duration.q": "Quanto tempo leva?", "faq.duration.a": "Senhas simples podem levar minutos; senhas complexas podem levar horas, dias ou mais.", "faq.local.q": "Meus arquivos saem do computador?", "faq.local.a": "Não. O RecoverEase foi criado para processamento local.", "faq.internet.q": "É preciso internet?", "faq.internet.a": "A recuperação não depende da internet. A conexão é usada apenas para atualizações, ativação ou suporte.", "faq.safety.q": "Meus dados estão seguros?", "faq.safety.a": "O software não envia seus arquivos e pode executar tarefas sensíveis offline.", "download.eyebrow": "Baixar", "download.title": "Comece a recuperar o acesso aos seus arquivos", "download.subtitle": "Uma ferramenta moderna para usuários domésticos, pequenas empresas e profissionais de TI.", "download.windows": "Baixar para Windows", "download.mac": "Mac em breve", "footer.copyright": "© 2026 RecoverEase", "footer.aria": "Navegação do rodapé", "footer.privacy": "Política de privacidade", "footer.terms": "Termos de serviço", "footer.contact": "Contato",
	})
	addLocale("ru", map[string]string{
		"meta.title": "RecoverEase - программа для восстановления паролей", "meta.description": "RecoverEase помогает вернуть доступ к архивам, документам и зашифрованным дискам, защищенным паролем.", "brand.aria": "Главная RecoverEase", "nav.aria": "Основная навигация", "nav.formats": "Форматы", "nav.modes": "Режимы", "nav.pricing": "Цены", "nav.download": "Скачать", "controls.language": "Язык", "controls.theme": "Включить темную тему", "cta.freeTrial": "Бесплатная версия", "cta.buyPro": "Купить Pro", "hero.eyebrow": "Программа для восстановления паролей", "hero.title": "Верните доступ к зашифрованным файлам", "hero.subtitle": "RecoverEase помогает снова получить доступ к архивам, документам и зашифрованным дискам, защищенным паролем.", "hero.actions": "Основные действия", "hero.trust": "Преимущества продукта", "trust.offline": "Офлайн-обработка", "trust.noUpload": "Без загрузки файлов", "trust.gpu": "GPU-ускорение", "appPreview.aria": "Предпросмотр интерфейса RecoverEase", "app.nav.recovery": "Восстановление", "app.nav.files": "Файлы", "app.nav.hints": "Подсказки", "app.nav.reports": "Отчеты", "app.currentTask": "Текущая задача", "app.archiveRecovery": "Восстановление архива", "app.smartMode": "Умный режим", "app.encryptedDetected": "Обнаружен зашифрованный архив", "app.passwordHints": "Подсказки пароля", "app.hintSummer": "лето", "app.hintYear": "2024", "app.hintCompany": "компания", "app.gpuSpeed": "Скорость GPU", "app.local": "Локально", "app.privacyMode": "Режим приватности", "formats.eyebrow": "Поддерживаемые форматы", "formats.title": "Поддержка популярных зашифрованных файлов и контейнеров", "formats.subtitle": "От домашних резервных копий до бизнес-архивов RecoverEase обрабатывает защищенные файлы по понятному сценарию.", "formats.archive.title": "Архивы", "formats.archive.desc": "Поддерживаются распространенные сценарии RAR, ZIP и 7Z.", "formats.office.title": "Документы Office", "formats.office.desc": "Работа с защищенными файлами Microsoft Word, Excel и PowerPoint.", "formats.pdf.title": "PDF-файлы", "formats.pdf.desc": "Помогает восстановить пароль открытия PDF.", "formats.disk.title": "Шифрование дисков", "formats.disk.desc": "Профессиональные задачи для контейнеров BitLocker и VeraCrypt.", "modes.eyebrow": "Режимы восстановления", "modes.title": "От быстрых попыток до экспертных стратегий", "modes.subtitle": "Выберите метод по подсказкам, типу файла и доступному времени.", "modes.quick.title": "Быстрое восстановление", "modes.quick.desc": "Сначала проверяет частые пароли и легкие правила.", "modes.quick.fit": "Подходит: домашним пользователям и недавно забытым паролям", "modes.smart.title": "Умное восстановление", "modes.smart.desc": "Автоматически строит стратегию по данным файла, словарям и подсказкам.", "modes.smart.fit": "Подходит: большинству задач", "modes.deep.title": "Глубокое восстановление", "modes.deep.desc": "Расширяет поиск для сложных паролей или малого числа подсказок.", "modes.deep.fit": "Подходит: бизнес-архивам и давно забытым паролям", "modes.expert.title": "Экспертный режим", "modes.expert.desc": "Настройте длину, наборы символов, правила и очереди задач.", "modes.expert.fit": "Подходит: IT-специалистам", "labels.recommended": "Рекомендуется", "labels.bestValue": "Лучший выбор", "advantages.eyebrow": "Преимущества", "advantages.title": "Профессионально, надежно и с уважением к приватности", "advantages.subtitle": "RecoverEase ориентирован на локальное восстановление личных файлов, активов малого бизнеса и IT-поддержки.", "advantages.gpu.title": "GPU-ускорение", "advantages.gpu.desc": "Используйте локальную видеокарту для ускорения восстановления.", "advantages.hints.title": "Умные подсказки пароля", "advantages.hints.desc": "Преобразует имена, даты, ключевые слова и шаблоны в стратегии поиска.", "advantages.offline.title": "Офлайн-обработка", "advantages.offline.desc": "Основные задачи выполняются локально без облачных очередей.", "advantages.privacy.title": "Приватность прежде всего", "advantages.privacy.desc": "Чувствительные файлы не загружаются, а процесс остается под вашим контролем.", "advantages.noUpload.title": "Без загрузки файлов", "advantages.noUpload.desc": "Ваши файлы всегда остаются на вашем компьютере.", "workflow.eyebrow": "Процесс", "workflow.title": "Начните восстановление за четыре шага", "workflow.step1.title": "Выберите файл", "workflow.step1.desc": "Добавьте архив, документ или зашифрованный контейнер.", "workflow.step2.title": "Выберите режим", "workflow.step2.desc": "Выберите быструю, умную, глубокую или экспертную стратегию.", "workflow.step3.title": "Добавьте подсказки", "workflow.step3.desc": "Введите вероятные имена, даты, фразы или шаблоны.", "workflow.step4.title": "Запустите восстановление", "workflow.step4.desc": "Отслеживайте прогресс, скорость и результаты в реальном времени.", "pricing.eyebrow": "Цены", "pricing.title": "Простая и прозрачная лицензия", "pricing.free.title": "Бесплатно", "pricing.free.item1": "Анализ файла", "pricing.free.item2": "Быстрое восстановление", "pricing.free.item3": "Словарь частых паролей", "pricing.pro.title": "Pro", "pricing.pro.note": "пожизненная лицензия одним платежом", "pricing.pro.item1": "Неограниченное восстановление", "pricing.pro.item2": "Глубокое восстановление", "pricing.pro.item3": "GPU-ускорение", "pricing.pro.item4": "Пакетные задачи", "pricing.pro.item5": "Приоритетная поддержка", "faq.eyebrow": "FAQ", "faq.title": "Частые вопросы", "faq.allPasswords.q": "Можно ли восстановить любой пароль?", "faq.allPasswords.a": "Нет. Результат зависит от сложности пароля, подсказок, типа файла и доступного времени вычислений.", "faq.duration.q": "Сколько это занимает?", "faq.duration.a": "Простые пароли могут занять минуты, сложные - часы, дни или дольше.", "faq.local.q": "Покидают ли файлы мой компьютер?", "faq.local.a": "Нет. RecoverEase построен вокруг локальной обработки.", "faq.internet.q": "Нужен ли интернет?", "faq.internet.a": "Восстановление не зависит от интернета. Подключение нужно только для обновлений, активации или поддержки.", "faq.safety.q": "Безопасны ли мои данные?", "faq.safety.a": "Программа не загружает ваши файлы и может выполнять чувствительные задачи офлайн.", "download.eyebrow": "Скачать", "download.title": "Начните возвращать доступ к файлам", "download.subtitle": "Современный инструмент для домашних пользователей, малого бизнеса и IT-специалистов.", "download.windows": "Скачать для Windows", "download.mac": "Mac скоро", "footer.copyright": "© 2026 RecoverEase", "footer.aria": "Навигация в подвале", "footer.privacy": "Политика конфиденциальности", "footer.terms": "Условия использования", "footer.contact": "Связаться с нами",
	})
	applyRecoveryModeOverrides()
	applyDownloadCtaOverrides()
	applyFooterContent()
	applyCheckoutContent()
}

var recoveryModeOverrides = map[string]map[string]string{
	"en": {
		"modes.title":          "Two Recovery Modes, Clear Control",
		"modes.subtitle":       "Use Smart Recovery for automatic common-password attempts, or Advanced Recovery to define length, character sets, and required strings yourself.",
		"modes.smart.desc":     "Automatically tries common passwords and practical rules, then uses your hints to prioritize likely candidates.",
		"modes.smart.fit":      "Best for: most users and first recovery attempts",
		"modes.advanced.title": "Advanced Recovery",
		"modes.advanced.desc":  "Set password length, character ranges, included strings, and other constraints to narrow the search.",
		"modes.advanced.fit":   "Best for: users who know part of the password structure",
		"workflow.step2.desc":  "Choose Smart Recovery for automatic attempts or Advanced Recovery for custom length and character settings.",
		"pricing.free.item2":   "Smart Recovery",
		"pricing.pro.item2":    "Advanced Recovery",
	},
	"zh": {
		"modes.title":          "两种恢复模式，路径更清晰",
		"modes.subtitle":       "智能恢复会自动尝试常见密码；高级恢复让您自行设置密码长度、字符集和包含字符串等条件。",
		"modes.smart.desc":     "自动尝试常见密码和实用规则，并结合线索优先测试更可能的候选密码。",
		"modes.smart.fit":      "适合：大多数用户和首次恢复尝试",
		"modes.advanced.title": "高级恢复",
		"modes.advanced.desc":  "自定义密码长度、字符范围、必须包含的字符串等条件，缩小搜索范围。",
		"modes.advanced.fit":   "适合：知道部分密码结构的用户",
		"workflow.step2.desc":  "选择智能恢复自动尝试，或使用高级恢复自定义长度、字符集和字符串条件。",
		"pricing.free.item2":   "智能恢复",
		"pricing.pro.item2":    "高级恢复",
	},
	"ja": {
		"modes.title":          "2つの復旧モードで迷わず開始",
		"modes.subtitle":       "スマート復旧は一般的なパスワードを自動で試行し、高度な復旧では長さ、文字種、含める文字列を自分で設定できます。",
		"modes.smart.desc":     "一般的なパスワードと実用的なルールを自動で試し、ヒントを使って可能性の高い候補を優先します。",
		"modes.smart.fit":      "最適: 多くのユーザーと初回の復旧",
		"modes.advanced.title": "高度な復旧",
		"modes.advanced.desc":  "パスワード長、文字範囲、含める文字列などを指定して検索範囲を絞り込みます。",
		"modes.advanced.fit":   "最適: パスワード構造の一部を覚えているユーザー",
		"workflow.step2.desc":  "自動試行のスマート復旧、または長さや文字条件を設定する高度な復旧を選びます。",
		"pricing.free.item2":   "スマート復旧",
		"pricing.pro.item2":    "高度な復旧",
	},
	"ko": {
		"modes.title":          "두 가지 복구 모드로 더 명확하게",
		"modes.subtitle":       "스마트 복구는 일반적인 비밀번호를 자동으로 시도하고, 고급 복구는 길이, 문자 집합, 포함 문자열을 직접 설정합니다.",
		"modes.smart.desc":     "일반적인 비밀번호와 실용 규칙을 자동으로 시도하고, 힌트를 바탕으로 가능성 높은 후보를 우선 테스트합니다.",
		"modes.smart.fit":      "적합: 대부분의 사용자와 첫 복구 시도",
		"modes.advanced.title": "고급 복구",
		"modes.advanced.desc":  "비밀번호 길이, 문자 범위, 포함해야 할 문자열 등 조건을 설정해 검색 범위를 좁힙니다.",
		"modes.advanced.fit":   "적합: 비밀번호 구조 일부를 알고 있는 사용자",
		"workflow.step2.desc":  "자동 시도는 스마트 복구를, 길이와 문자 조건 설정은 고급 복구를 선택합니다.",
		"pricing.free.item2":   "스마트 복구",
		"pricing.pro.item2":    "고급 복구",
	},
	"de": {
		"modes.title":          "Zwei Wiederherstellungsmodi, klare Kontrolle",
		"modes.subtitle":       "Smart Recovery versucht gängige Passwörter automatisch; Advanced Recovery lässt Sie Länge, Zeichensätze und erforderliche Zeichenfolgen selbst festlegen.",
		"modes.smart.desc":     "Versucht automatisch gängige Passwörter und praktische Regeln und priorisiert mit Ihren Hinweisen wahrscheinliche Kandidaten.",
		"modes.smart.fit":      "Geeignet für: die meisten Nutzer und erste Wiederherstellungsversuche",
		"modes.advanced.title": "Advanced Recovery",
		"modes.advanced.desc":  "Legen Sie Passwortlänge, Zeichenbereiche, enthaltene Zeichenfolgen und weitere Grenzen fest.",
		"modes.advanced.fit":   "Geeignet für: Nutzer, die Teile der Passwortstruktur kennen",
		"workflow.step2.desc":  "Wählen Sie Smart Recovery für automatische Versuche oder Advanced Recovery für eigene Längen- und Zeicheneinstellungen.",
		"pricing.free.item2":   "Smart Recovery",
		"pricing.pro.item2":    "Advanced Recovery",
	},
	"fr": {
		"modes.title":          "Deux modes de récupération, plus de clarté",
		"modes.subtitle":       "La récupération intelligente essaie automatiquement les mots de passe courants ; la récupération avancée vous laisse définir longueur, jeux de caractères et chaînes requises.",
		"modes.smart.desc":     "Essaie automatiquement les mots de passe courants et des règles pratiques, puis utilise vos indices pour prioriser les candidats probables.",
		"modes.smart.fit":      "Idéal pour : la plupart des utilisateurs et les premières tentatives",
		"modes.advanced.title": "Récupération avancée",
		"modes.advanced.desc":  "Définissez la longueur, les plages de caractères, les chaînes à inclure et d'autres contraintes.",
		"modes.advanced.fit":   "Idéal pour : les utilisateurs qui connaissent une partie de la structure",
		"workflow.step2.desc":  "Choisissez la récupération intelligente pour les essais automatiques ou la récupération avancée pour les réglages personnalisés.",
		"pricing.free.item2":   "Récupération intelligente",
		"pricing.pro.item2":    "Récupération avancée",
	},
	"es": {
		"modes.title":          "Dos modos de recuperación, control claro",
		"modes.subtitle":       "La recuperación inteligente prueba contraseñas comunes automáticamente; la avanzada permite definir longitud, conjuntos de caracteres y cadenas requeridas.",
		"modes.smart.desc":     "Prueba automáticamente contraseñas comunes y reglas prácticas, y usa tus pistas para priorizar candidatos probables.",
		"modes.smart.fit":      "Ideal para: la mayoría de usuarios y primeros intentos",
		"modes.advanced.title": "Recuperación avanzada",
		"modes.advanced.desc":  "Configura longitud, rangos de caracteres, cadenas incluidas y otras restricciones.",
		"modes.advanced.fit":   "Ideal para: usuarios que conocen parte de la estructura",
		"workflow.step2.desc":  "Elige recuperación inteligente para intentos automáticos o avanzada para ajustes personalizados.",
		"pricing.free.item2":   "Recuperación inteligente",
		"pricing.pro.item2":    "Recuperación avanzada",
	},
	"pt": {
		"modes.title":          "Dois modos de recuperação, controle claro",
		"modes.subtitle":       "A recuperação inteligente tenta senhas comuns automaticamente; a avançada permite definir comprimento, conjuntos de caracteres e strings obrigatórias.",
		"modes.smart.desc":     "Tenta automaticamente senhas comuns e regras práticas, usando suas pistas para priorizar candidatos prováveis.",
		"modes.smart.fit":      "Ideal para: a maioria dos usuários e primeiras tentativas",
		"modes.advanced.title": "Recuperação avançada",
		"modes.advanced.desc":  "Defina comprimento, faixas de caracteres, strings incluídas e outras restrições.",
		"modes.advanced.fit":   "Ideal para: usuários que conhecem parte da estrutura",
		"workflow.step2.desc":  "Escolha recuperação inteligente para tentativas automáticas ou avançada para ajustes personalizados.",
		"pricing.free.item2":   "Recuperação inteligente",
		"pricing.pro.item2":    "Recuperação avançada",
	},
	"ru": {
		"modes.title":          "Два режима восстановления, больше контроля",
		"modes.subtitle":       "Умное восстановление автоматически пробует частые пароли; расширенное позволяет задать длину, наборы символов и обязательные строки.",
		"modes.smart.desc":     "Автоматически пробует частые пароли и практичные правила, используя подсказки для приоритета вероятных вариантов.",
		"modes.smart.fit":      "Подходит: большинству пользователей и первым попыткам",
		"modes.advanced.title": "Расширенное восстановление",
		"modes.advanced.desc":  "Настройте длину пароля, диапазоны символов, обязательные строки и другие ограничения.",
		"modes.advanced.fit":   "Подходит: тем, кто знает часть структуры пароля",
		"workflow.step2.desc":  "Выберите умное восстановление для автоматических попыток или расширенное для собственных настроек.",
		"pricing.free.item2":   "Умное восстановление",
		"pricing.pro.item2":    "Расширенное восстановление",
	},
}

var downloadCtaLabels = map[string]string{
	"en": "Download Now",
	"zh": "立即下载",
	"ja": "今すぐダウンロード",
	"ko": "지금 다운로드",
	"de": "Jetzt herunterladen",
	"fr": "Telecharger",
	"es": "Descargar ahora",
	"pt": "Baixar agora",
	"ru": "Скачать сейчас",
}

func applyDownloadCtaOverrides() {
	for code, label := range downloadCtaLabels {
		if locale, ok := translations[code]; ok {
			locale["cta.freeTrial"] = label
		}
	}
}

var footerContent = map[string]map[string]string{
	"en": {"footer.licenseRecovery": "Recover license key"},
	"zh": {"footer.licenseRecovery": "找回激活码"},
	"ja": {"footer.licenseRecovery": "ライセンスキーを復元"},
	"ko": {"footer.licenseRecovery": "라이선스 키 찾기"},
	"de": {"footer.licenseRecovery": "Lizenzschlüssel wiederherstellen"},
	"fr": {"footer.licenseRecovery": "Récupérer la clé"},
	"es": {"footer.licenseRecovery": "Recuperar clave"},
	"pt": {"footer.licenseRecovery": "Recuperar chave"},
	"ru": {"footer.licenseRecovery": "Восстановить ключ"},
}

func applyFooterContent() {
	for code, labels := range footerContent {
		if locale, ok := translations[code]; ok {
			for key, value := range labels {
				locale[key] = value
			}
		}
	}
}

var checkoutContent = map[string]map[string]string{
	"en": {
		"checkout.meta.title":               "RecoverEase Pro Checkout",
		"checkout.meta.description":         "Complete your RecoverEase Pro purchase and get a lifetime license for advanced password recovery.",
		"checkout.backHome":                 "Back to website",
		"checkout.recovery.nav":             "Recover license key",
		"checkout.eyebrow":                  "Secure Checkout",
		"checkout.title":                    "Get your RecoverEase Pro license key",
		"checkout.subtitle":                 "Enter the email for license delivery, then continue to the secure payment page.",
		"checkout.aria":                     "RecoverEase Pro checkout",
		"checkout.account.title":            "Email for license delivery",
		"checkout.account.desc":             "Your license key and purchase confirmation will be sent to this email address after payment.",
		"checkout.account.badges":           "License delivery notes",
		"checkout.account.noLogin":          "No account required",
		"checkout.account.recoverable":      "License key sent by email",
		"checkout.email":                    "Email address",
		"checkout.emailPlaceholder":         "Enter your purchase email",
		"checkout.emailHint":                "Use an email you can access later if you need your license key resent.",
		"checkout.name":                     "Receipt name",
		"checkout.payment.title":            "Payment method",
		"checkout.payment.desc":             "Choose USDT or credit card, then continue to the secure payment page.",
		"checkout.payment.usdt":             "USDT",
		"checkout.payment.card":             "Credit card",
		"checkout.cardNumber":               "Card number",
		"checkout.expiry":                   "Expiry date",
		"checkout.cvc":                      "Security code",
		"checkout.billing.title":            "Receipt details",
		"checkout.billing.desc":             "Used for tax calculation, payment records, and receipt delivery where required.",
		"checkout.country":                  "Country or region",
		"checkout.country.us":               "United States",
		"checkout.country.cn":               "China",
		"checkout.country.eu":               "European Union",
		"checkout.country.other":            "Other",
		"checkout.license.title":            "License type",
		"checkout.license.monthly.title":    "One-month license",
		"checkout.license.monthly.desc":     "1 user, valid for one month",
		"checkout.license.monthly.price":    "$9",
		"checkout.license.lifetime.title":   "Lifetime license",
		"checkout.license.lifetime.desc":    "1 user, one-time purchase",
		"checkout.license.lifetime.price":   "$29",
		"checkout.summary.title":            "License summary",
		"checkout.summary.license":          "Lifetime license, 1 user",
		"checkout.summary.total":            "Amount due",
		"checkout.pay":                      "Pay now",
		"checkout.terms":                    "You will be redirected to the payment provider to complete payment. After confirmation, your license key appears on the payment result page.",
		"checkout.success.meta.title":       "RecoverEase Payment Result",
		"checkout.success.meta.description": "View your RecoverEase payment result and license key.",
		"checkout.success.eyebrow":          "Payment Result",
		"checkout.success.title":            "Your license key is ready",
		"checkout.success.subtitle":         "Keep this page open until you copy your license key. You can also recover it later with your purchase email.",
		"checkout.success.aria":             "RecoverEase payment result",
		"checkout.success.card.title":       "License key",
		"checkout.success.loading":          "Checking payment result...",
		"checkout.success.missing":          "Payment number is missing. Please return to checkout and try again.",
		"checkout.success.pending":          "Payment is not confirmed yet. Please refresh this page after the payment provider confirms it.",
		"checkout.success.failed":           "Unable to load payment result. Please try again later.",
		"checkout.success.copy":             "Copy",
		"checkout.success.copied":           "Copied",
		"checkout.success.order":            "Order",
		"checkout.success.plan":             "Plan",
		"checkout.success.issuedAt":         "Issued",
		"checkout.success.expiresAt":        "Expires",
		"checkout.success.lifetime":         "Lifetime",
		"checkout.success.backCheckout":     "Back to checkout",
		"checkout.success.recovery":         "Recover license key",
		"checkout.recovery.title":           "Already purchased?",
		"checkout.recovery.desc":            "Enter the purchase email to receive a secure link for your license key and purchase record.",
		"checkout.recovery.email":           "Purchase email",
		"checkout.recovery.submit":          "Recover license key",
		"recovery.meta.title":               "Recover RecoverEase License Key",
		"recovery.meta.description":         "Recover your RecoverEase Pro license key by purchase email.",
		"recovery.backCheckout":             "Buy Pro",
		"recovery.eyebrow":                  "License Recovery",
		"recovery.title":                    "Recover your license key",
		"recovery.subtitle":                 "Enter the email used at purchase. We will send a secure link to view your license key and purchase record.",
		"recovery.aria":                     "RecoverEase license recovery",
		"recovery.form.title":               "Purchase email",
		"recovery.form.desc":                "No password is needed. The secure link is sent only to the purchase email.",
		"recovery.email":                    "Email address",
		"recovery.emailHint":                "Use the same email entered during checkout.",
		"recovery.submit":                   "Send recovery link",
		"recovery.note":                     "If the email matches a purchase, the recovery link will arrive in a few minutes.",
		"checkout.assurance.delivery.title": "Instant license delivery",
		"checkout.assurance.delivery.desc":  "See your license key after payment and keep a copy in your inbox.",
		"checkout.assurance.privacy.title":  "Local recovery workflow",
		"checkout.assurance.privacy.desc":   "RecoverEase does not upload your protected files for recovery.",
		"checkout.assurance.secure.title":   "Encrypted payment",
		"checkout.assurance.secure.desc":    "Payment details are submitted through a secure payment flow.",
	},
	"zh": {
		"checkout.meta.title":               "RecoverEase Pro 结算",
		"checkout.meta.description":         "完成 RecoverEase Pro 购买，获取高级密码恢复功能的一次性永久授权。",
		"checkout.backHome":                 "返回官网",
		"checkout.recovery.nav":             "找回激活码",
		"checkout.eyebrow":                  "安全结算",
		"checkout.title":                    "获取 RecoverEase Pro 激活码",
		"checkout.subtitle":                 "填写接收激活码的邮箱，然后前往安全支付页完成付款。",
		"checkout.aria":                     "RecoverEase Pro 结算页",
		"checkout.account.title":            "用于接收激活码的邮箱",
		"checkout.account.desc":             "支付完成后，激活码和购买确认邮件会发送到这个邮箱地址。",
		"checkout.account.badges":           "授权交付说明",
		"checkout.account.noLogin":          "无需注册账号",
		"checkout.account.recoverable":      "激活码邮件交付",
		"checkout.email":                    "邮箱地址",
		"checkout.emailPlaceholder":         "请输入购买时使用的邮箱",
		"checkout.emailHint":                "请填写以后还能访问的邮箱，方便重新发送激活码。",
		"checkout.name":                     "收据姓名",
		"checkout.payment.title":            "选择支付方式",
		"checkout.payment.desc":             "选择 USDT 或信用卡，然后前往安全支付页完成付款。",
		"checkout.payment.usdt":             "USDT",
		"checkout.payment.card":             "信用卡",
		"checkout.cardNumber":               "银行卡号",
		"checkout.expiry":                   "有效期",
		"checkout.cvc":                      "安全码",
		"checkout.billing.title":            "收据信息",
		"checkout.billing.desc":             "用于税费计算、支付记录以及必要的收据交付。",
		"checkout.country":                  "国家或地区",
		"checkout.country.us":               "美国",
		"checkout.country.cn":               "中国",
		"checkout.country.eu":               "欧盟",
		"checkout.country.other":            "其他",
		"checkout.license.title":            "授权类型",
		"checkout.license.monthly.title":    "一个月授权",
		"checkout.license.monthly.desc":     "1 位用户，有效期一个月",
		"checkout.license.monthly.price":    "$9",
		"checkout.license.lifetime.title":   "永久授权",
		"checkout.license.lifetime.desc":    "1 位用户，一次性购买",
		"checkout.license.lifetime.price":   "$29",
		"checkout.summary.title":            "授权摘要",
		"checkout.summary.license":          "永久授权，1 位用户",
		"checkout.summary.total":            "应付金额",
		"checkout.pay":                      "立即支付",
		"checkout.terms":                    "您将跳转到支付服务商页面完成付款。支付确认后，激活码会显示在支付结果页。",
		"checkout.success.meta.title":       "RecoverEase 支付结果",
		"checkout.success.meta.description": "查看 RecoverEase 支付结果和激活码。",
		"checkout.success.eyebrow":          "支付结果",
		"checkout.success.title":            "您的激活码已生成",
		"checkout.success.subtitle":         "请在复制激活码前保留此页面。之后也可以用购买邮箱找回激活码。",
		"checkout.success.aria":             "RecoverEase 支付结果",
		"checkout.success.card.title":       "激活码",
		"checkout.success.loading":          "正在查询支付结果...",
		"checkout.success.missing":          "缺少支付单号，请返回结算页重新尝试。",
		"checkout.success.pending":          "支付尚未确认。请在支付服务商确认后刷新此页面。",
		"checkout.success.failed":           "暂时无法加载支付结果，请稍后重试。",
		"checkout.success.copy":             "复制",
		"checkout.success.copied":           "已复制",
		"checkout.success.order":            "订单号",
		"checkout.success.plan":             "授权方案",
		"checkout.success.issuedAt":         "发放时间",
		"checkout.success.expiresAt":        "有效期至",
		"checkout.success.lifetime":         "永久有效",
		"checkout.success.backCheckout":     "返回结算页",
		"checkout.success.recovery":         "找回激活码",
		"checkout.recovery.title":           "已经购买？",
		"checkout.recovery.desc":            "输入购买时使用的邮箱，接收安全链接以找回激活码和购买记录。",
		"checkout.recovery.email":           "购买邮箱",
		"checkout.recovery.submit":          "找回激活码",
		"recovery.meta.title":               "找回 RecoverEase 激活码",
		"recovery.meta.description":         "通过购买邮箱找回 RecoverEase Pro 激活码。",
		"recovery.backCheckout":             "购买专业版",
		"recovery.eyebrow":                  "激活码找回",
		"recovery.title":                    "找回您的激活码",
		"recovery.subtitle":                 "输入购买时使用的邮箱，我们会发送安全链接，用于查看激活码和购买记录。",
		"recovery.aria":                     "RecoverEase 激活码找回",
		"recovery.form.title":               "购买邮箱",
		"recovery.form.desc":                "无需密码。安全链接只会发送到购买时使用的邮箱。",
		"recovery.email":                    "邮箱地址",
		"recovery.emailHint":                "请填写结算时使用的同一个邮箱。",
		"recovery.submit":                   "发送找回链接",
		"recovery.note":                     "如果该邮箱存在购买记录，找回链接会在几分钟内发送。",
		"checkout.assurance.delivery.title": "授权码即时交付",
		"checkout.assurance.delivery.desc":  "支付确认后可立即查看激活码，并在邮箱中保留副本。",
		"checkout.assurance.privacy.title":  "本地恢复流程",
		"checkout.assurance.privacy.desc":   "RecoverEase 不会上传您的受保护文件进行恢复。",
		"checkout.assurance.secure.title":   "加密支付",
		"checkout.assurance.secure.desc":    "支付信息会通过安全支付流程提交。",
	},
	"ja": {
		"checkout.meta.title":               "RecoverEase Pro 決済",
		"checkout.meta.description":         "RecoverEase Pro の購入を完了し、高度なパスワード復元向けライセンスを取得します。",
		"checkout.backHome":                 "サイトへ戻る",
		"checkout.recovery.nav":             "ライセンスキーを復元",
		"checkout.eyebrow":                  "安全な決済",
		"checkout.title":                    "RecoverEase Pro ライセンスキーを取得",
		"checkout.subtitle":                 "ライセンスを受け取るメールアドレスを入力し、安全な支払いページへ進みます。",
		"checkout.aria":                     "RecoverEase Pro 決済",
		"checkout.account.title":            "ライセンス受信用メール",
		"checkout.account.desc":             "支払い完了後、ライセンスキーと購入確認メールがこのアドレスへ送信されます。",
		"checkout.account.badges":           "ライセンス配信情報",
		"checkout.account.noLogin":          "アカウント不要",
		"checkout.account.recoverable":      "ライセンスキーをメールで送信",
		"checkout.email":                    "メールアドレス",
		"checkout.emailPlaceholder":         "購入時に使用するメールを入力",
		"checkout.emailHint":                "後でライセンスキーを再送できるよう、利用可能なメールを入力してください。",
		"checkout.name":                     "領収書名",
		"checkout.payment.title":            "支払い方法",
		"checkout.payment.desc":             "USDT またはクレジットカードを選び、安全な支払いページへ進みます。",
		"checkout.payment.usdt":             "USDT",
		"checkout.payment.card":             "クレジットカード",
		"checkout.cardNumber":               "カード番号",
		"checkout.expiry":                   "有効期限",
		"checkout.cvc":                      "セキュリティコード",
		"checkout.billing.title":            "領収書情報",
		"checkout.billing.desc":             "必要に応じて税計算、支払い記録、領収書送付に使用します。",
		"checkout.country":                  "国または地域",
		"checkout.country.us":               "アメリカ合衆国",
		"checkout.country.cn":               "中国",
		"checkout.country.eu":               "欧州連合",
		"checkout.country.other":            "その他",
		"checkout.license.title":            "ライセンス種類",
		"checkout.license.monthly.title":    "1か月ライセンス",
		"checkout.license.monthly.desc":     "1ユーザー、1か月有効",
		"checkout.license.monthly.price":    "$9",
		"checkout.license.lifetime.title":   "永久ライセンス",
		"checkout.license.lifetime.desc":    "1ユーザー、買い切り",
		"checkout.license.lifetime.price":   "$29",
		"checkout.summary.title":            "ライセンス概要",
		"checkout.summary.license":          "永久ライセンス、1ユーザー",
		"checkout.summary.total":            "お支払い金額",
		"checkout.pay":                      "今すぐ支払う",
		"checkout.terms":                    "支払いサービスのページへ移動して決済を完了します。確認後、ライセンスキーがメールで送信されます。",
		"checkout.recovery.title":           "購入済みですか？",
		"checkout.recovery.desc":            "購入時のメールを入力すると、ライセンスキーと購入記録を確認する安全なリンクを受け取れます。",
		"checkout.recovery.email":           "購入時のメール",
		"checkout.recovery.submit":          "ライセンスキーを復元",
		"recovery.meta.title":               "RecoverEase ライセンスキー復元",
		"recovery.meta.description":         "購入時のメールで RecoverEase Pro のライセンスキーを復元します。",
		"recovery.backCheckout":             "Pro を購入",
		"recovery.eyebrow":                  "ライセンス復元",
		"recovery.title":                    "ライセンスキーを復元",
		"recovery.subtitle":                 "購入時に使用したメールを入力してください。ライセンスキーと購入記録を確認する安全なリンクを送信します。",
		"recovery.aria":                     "RecoverEase ライセンス復元",
		"recovery.form.title":               "購入時のメール",
		"recovery.form.desc":                "パスワードは不要です。安全なリンクは購入時のメールにのみ送信されます。",
		"recovery.email":                    "メールアドレス",
		"recovery.emailHint":                "決済時に入力した同じメールを使用してください。",
		"recovery.submit":                   "復元リンクを送信",
		"recovery.note":                     "メールが購入記録と一致する場合、復元リンクが数分以内に届きます。",
		"checkout.assurance.delivery.title": "即時ライセンス配信",
		"checkout.assurance.delivery.desc":  "支払い確認後すぐにライセンスキーを確認でき、メールにも控えが残ります。",
		"checkout.assurance.privacy.title":  "ローカル復元ワークフロー",
		"checkout.assurance.privacy.desc":   "RecoverEase は復元のために保護されたファイルをアップロードしません。",
		"checkout.assurance.secure.title":   "暗号化された支払い",
		"checkout.assurance.secure.desc":    "支払い情報は安全な支払いフローで送信されます。",
	},
	"ko": {
		"checkout.meta.title":               "RecoverEase Pro 결제",
		"checkout.meta.description":         "RecoverEase Pro 구매를 완료하고 고급 비밀번호 복구용 라이선스를 받으세요.",
		"checkout.backHome":                 "웹사이트로 돌아가기",
		"checkout.recovery.nav":             "라이선스 키 찾기",
		"checkout.eyebrow":                  "안전 결제",
		"checkout.title":                    "RecoverEase Pro 라이선스 키 받기",
		"checkout.subtitle":                 "라이선스를 받을 이메일을 입력한 뒤 안전한 결제 페이지로 이동하세요.",
		"checkout.aria":                     "RecoverEase Pro 결제",
		"checkout.account.title":            "라이선스 수신 이메일",
		"checkout.account.desc":             "결제 완료 후 라이선스 키와 구매 확인 메일이 이 주소로 발송됩니다.",
		"checkout.account.badges":           "라이선스 전달 안내",
		"checkout.account.noLogin":          "계정 필요 없음",
		"checkout.account.recoverable":      "라이선스 키 이메일 발송",
		"checkout.email":                    "이메일 주소",
		"checkout.emailPlaceholder":         "구매에 사용할 이메일 입력",
		"checkout.emailHint":                "나중에 라이선스 키를 다시 받을 수 있도록 접근 가능한 이메일을 입력하세요.",
		"checkout.name":                     "영수증 이름",
		"checkout.payment.title":            "결제 방법",
		"checkout.payment.desc":             "USDT 또는 신용카드를 선택한 뒤 안전한 결제 페이지로 이동하세요.",
		"checkout.payment.usdt":             "USDT",
		"checkout.payment.card":             "신용카드",
		"checkout.cardNumber":               "카드 번호",
		"checkout.expiry":                   "유효 기간",
		"checkout.cvc":                      "보안 코드",
		"checkout.billing.title":            "영수증 정보",
		"checkout.billing.desc":             "필요한 경우 세금 계산, 결제 기록, 영수증 전달에 사용됩니다.",
		"checkout.country":                  "국가 또는 지역",
		"checkout.country.us":               "미국",
		"checkout.country.cn":               "중국",
		"checkout.country.eu":               "유럽 연합",
		"checkout.country.other":            "기타",
		"checkout.license.title":            "라이선스 유형",
		"checkout.license.monthly.title":    "1개월 라이선스",
		"checkout.license.monthly.desc":     "사용자 1명, 1개월 유효",
		"checkout.license.monthly.price":    "$9",
		"checkout.license.lifetime.title":   "영구 라이선스",
		"checkout.license.lifetime.desc":    "사용자 1명, 1회 구매",
		"checkout.license.lifetime.price":   "$29",
		"checkout.summary.title":            "라이선스 요약",
		"checkout.summary.license":          "영구 라이선스, 사용자 1명",
		"checkout.summary.total":            "결제 금액",
		"checkout.pay":                      "지금 결제",
		"checkout.terms":                    "결제 서비스 페이지로 이동해 결제를 완료합니다. 확인 후 라이선스 키가 이메일로 발송됩니다.",
		"checkout.recovery.title":           "이미 구매하셨나요?",
		"checkout.recovery.desc":            "구매 이메일을 입력하면 라이선스 키와 구매 기록을 확인할 안전한 링크를 받을 수 있습니다.",
		"checkout.recovery.email":           "구매 이메일",
		"checkout.recovery.submit":          "라이선스 키 찾기",
		"recovery.meta.title":               "RecoverEase 라이선스 키 찾기",
		"recovery.meta.description":         "구매 이메일로 RecoverEase Pro 라이선스 키를 찾습니다.",
		"recovery.backCheckout":             "Pro 구매",
		"recovery.eyebrow":                  "라이선스 찾기",
		"recovery.title":                    "라이선스 키 찾기",
		"recovery.subtitle":                 "구매 시 사용한 이메일을 입력하세요. 라이선스 키와 구매 기록을 확인할 안전한 링크를 보내드립니다.",
		"recovery.aria":                     "RecoverEase 라이선스 찾기",
		"recovery.form.title":               "구매 이메일",
		"recovery.form.desc":                "비밀번호는 필요 없습니다. 안전한 링크는 구매 이메일로만 전송됩니다.",
		"recovery.email":                    "이메일 주소",
		"recovery.emailHint":                "결제 시 입력한 이메일을 사용하세요.",
		"recovery.submit":                   "찾기 링크 보내기",
		"recovery.note":                     "해당 이메일에 구매 기록이 있으면 몇 분 안에 찾기 링크가 도착합니다.",
		"checkout.assurance.delivery.title": "즉시 라이선스 전달",
		"checkout.assurance.delivery.desc":  "결제 확인 후 라이선스 키를 바로 확인하고 이메일에 사본을 보관할 수 있습니다.",
		"checkout.assurance.privacy.title":  "로컬 복구 워크플로",
		"checkout.assurance.privacy.desc":   "RecoverEase는 복구를 위해 보호된 파일을 업로드하지 않습니다.",
		"checkout.assurance.secure.title":   "암호화된 결제",
		"checkout.assurance.secure.desc":    "결제 정보는 안전한 결제 흐름을 통해 제출됩니다.",
	},
	"de": {
		"checkout.meta.title":               "RecoverEase Pro Kasse",
		"checkout.meta.description":         "Schließen Sie den Kauf von RecoverEase Pro ab und erhalten Sie eine Lizenz für erweiterte Passwortwiederherstellung.",
		"checkout.backHome":                 "Zur Website",
		"checkout.recovery.nav":             "Lizenzschlüssel wiederherstellen",
		"checkout.eyebrow":                  "Sichere Kasse",
		"checkout.title":                    "RecoverEase Pro Lizenzschlüssel erhalten",
		"checkout.subtitle":                 "Geben Sie die E-Mail-Adresse für die Lizenzzustellung ein und gehen Sie dann zur sicheren Zahlungsseite.",
		"checkout.aria":                     "RecoverEase Pro Kasse",
		"checkout.account.title":            "E-Mail für Lizenzzustellung",
		"checkout.account.desc":             "Nach der Zahlung werden Lizenzschlüssel und Kaufbestätigung an diese E-Mail-Adresse gesendet.",
		"checkout.account.badges":           "Hinweise zur Lizenzzustellung",
		"checkout.account.noLogin":          "Kein Konto erforderlich",
		"checkout.account.recoverable":      "Lizenzschlüssel per E-Mail",
		"checkout.email":                    "E-Mail-Adresse",
		"checkout.emailPlaceholder":         "Kauf-E-Mail eingeben",
		"checkout.emailHint":                "Nutzen Sie eine E-Mail, auf die Sie später zugreifen können, falls der Lizenzschlüssel erneut gesendet werden muss.",
		"checkout.name":                     "Name für Beleg",
		"checkout.payment.title":            "Zahlungsmethode",
		"checkout.payment.desc":             "Wählen Sie USDT oder Kreditkarte und fahren Sie mit der sicheren Zahlungsseite fort.",
		"checkout.payment.usdt":             "USDT",
		"checkout.payment.card":             "Kreditkarte",
		"checkout.cardNumber":               "Kartennummer",
		"checkout.expiry":                   "Ablaufdatum",
		"checkout.cvc":                      "Sicherheitscode",
		"checkout.billing.title":            "Belegdaten",
		"checkout.billing.desc":             "Wird bei Bedarf für Steuerberechnung, Zahlungsunterlagen und Belegzustellung verwendet.",
		"checkout.country":                  "Land oder Region",
		"checkout.country.us":               "Vereinigte Staaten",
		"checkout.country.cn":               "China",
		"checkout.country.eu":               "Europäische Union",
		"checkout.country.other":            "Andere",
		"checkout.license.title":            "Lizenztyp",
		"checkout.license.monthly.title":    "Ein-Monats-Lizenz",
		"checkout.license.monthly.desc":     "1 Benutzer, einen Monat gültig",
		"checkout.license.monthly.price":    "$9",
		"checkout.license.lifetime.title":   "Lebenslange Lizenz",
		"checkout.license.lifetime.desc":    "1 Benutzer, einmaliger Kauf",
		"checkout.license.lifetime.price":   "$29",
		"checkout.summary.title":            "Lizenzübersicht",
		"checkout.summary.license":          "Lebenslange Lizenz, 1 Benutzer",
		"checkout.summary.total":            "Fälliger Betrag",
		"checkout.pay":                      "Jetzt bezahlen",
		"checkout.terms":                    "Sie werden zur Zahlungsseite weitergeleitet. Nach der Bestätigung wird der Lizenzschlüssel an Ihre E-Mail gesendet.",
		"checkout.recovery.title":           "Bereits gekauft?",
		"checkout.recovery.desc":            "Geben Sie die Kauf-E-Mail ein, um einen sicheren Link für Lizenzschlüssel und Kaufdatensatz zu erhalten.",
		"checkout.recovery.email":           "Kauf-E-Mail",
		"checkout.recovery.submit":          "Lizenzschlüssel wiederherstellen",
		"recovery.meta.title":               "RecoverEase Lizenzschlüssel wiederherstellen",
		"recovery.meta.description":         "Stellen Sie Ihren RecoverEase Pro Lizenzschlüssel über die Kauf-E-Mail wieder her.",
		"recovery.backCheckout":             "Pro kaufen",
		"recovery.eyebrow":                  "Lizenzwiederherstellung",
		"recovery.title":                    "Lizenzschlüssel wiederherstellen",
		"recovery.subtitle":                 "Geben Sie die beim Kauf verwendete E-Mail ein. Wir senden einen sicheren Link zum Anzeigen Ihres Lizenzschlüssels und Kaufdatensatzes.",
		"recovery.aria":                     "RecoverEase Lizenzwiederherstellung",
		"recovery.form.title":               "Kauf-E-Mail",
		"recovery.form.desc":                "Kein Passwort erforderlich. Der sichere Link wird nur an die Kauf-E-Mail gesendet.",
		"recovery.email":                    "E-Mail-Adresse",
		"recovery.emailHint":                "Verwenden Sie dieselbe E-Mail wie beim Checkout.",
		"recovery.submit":                   "Wiederherstellungslink senden",
		"recovery.note":                     "Wenn die E-Mail zu einem Kauf passt, trifft der Link in wenigen Minuten ein.",
		"checkout.assurance.delivery.title": "Sofortige Lizenzzustellung",
		"checkout.assurance.delivery.desc":  "Nach Zahlungsbestätigung sehen Sie den Lizenzschlüssel sofort und behalten eine Kopie im Postfach.",
		"checkout.assurance.privacy.title":  "Lokaler Wiederherstellungsablauf",
		"checkout.assurance.privacy.desc":   "RecoverEase lädt Ihre geschützten Dateien nicht zur Wiederherstellung hoch.",
		"checkout.assurance.secure.title":   "Verschlüsselte Zahlung",
		"checkout.assurance.secure.desc":    "Zahlungsdaten werden über einen sicheren Zahlungsablauf übermittelt.",
	},
	"fr": {
		"checkout.meta.title":               "Paiement RecoverEase Pro",
		"checkout.meta.description":         "Finalisez votre achat RecoverEase Pro et obtenez une licence pour la récupération avancée de mots de passe.",
		"checkout.backHome":                 "Retour au site",
		"checkout.recovery.nav":             "Récupérer la clé",
		"checkout.eyebrow":                  "Paiement sécurisé",
		"checkout.title":                    "Obtenir votre clé RecoverEase Pro",
		"checkout.subtitle":                 "Saisissez l'e-mail de livraison de la licence, puis continuez vers la page de paiement sécurisée.",
		"checkout.aria":                     "Paiement RecoverEase Pro",
		"checkout.account.title":            "E-mail de livraison de la licence",
		"checkout.account.desc":             "Après paiement, la clé de licence et la confirmation d'achat seront envoyées à cette adresse.",
		"checkout.account.badges":           "Notes de livraison de licence",
		"checkout.account.noLogin":          "Aucun compte requis",
		"checkout.account.recoverable":      "Clé envoyée par e-mail",
		"checkout.email":                    "Adresse e-mail",
		"checkout.emailPlaceholder":         "Saisissez l'e-mail d'achat",
		"checkout.emailHint":                "Utilisez un e-mail accessible plus tard si vous devez recevoir à nouveau votre clé.",
		"checkout.name":                     "Nom du reçu",
		"checkout.payment.title":            "Mode de paiement",
		"checkout.payment.desc":             "Choisissez USDT ou carte bancaire, puis continuez vers la page de paiement sécurisée.",
		"checkout.payment.usdt":             "USDT",
		"checkout.payment.card":             "Carte bancaire",
		"checkout.cardNumber":               "Numéro de carte",
		"checkout.expiry":                   "Date d'expiration",
		"checkout.cvc":                      "Code de sécurité",
		"checkout.billing.title":            "Informations de reçu",
		"checkout.billing.desc":             "Utilisées si nécessaire pour les taxes, les enregistrements de paiement et l'envoi du reçu.",
		"checkout.country":                  "Pays ou région",
		"checkout.country.us":               "États-Unis",
		"checkout.country.cn":               "Chine",
		"checkout.country.eu":               "Union européenne",
		"checkout.country.other":            "Autre",
		"checkout.license.title":            "Type de licence",
		"checkout.license.monthly.title":    "Licence d'un mois",
		"checkout.license.monthly.desc":     "1 utilisateur, valable un mois",
		"checkout.license.monthly.price":    "$9",
		"checkout.license.lifetime.title":   "Licence à vie",
		"checkout.license.lifetime.desc":    "1 utilisateur, achat unique",
		"checkout.license.lifetime.price":   "$29",
		"checkout.summary.title":            "Résumé de la licence",
		"checkout.summary.license":          "Licence à vie, 1 utilisateur",
		"checkout.summary.total":            "Montant dû",
		"checkout.pay":                      "Payer maintenant",
		"checkout.terms":                    "Vous serez redirigé vers le prestataire de paiement. Après confirmation, la clé sera envoyée par e-mail.",
		"checkout.recovery.title":           "Déjà acheté ?",
		"checkout.recovery.desc":            "Saisissez l'e-mail d'achat pour recevoir un lien sécurisé vers votre clé et votre historique.",
		"checkout.recovery.email":           "E-mail d'achat",
		"checkout.recovery.submit":          "Récupérer la clé",
		"recovery.meta.title":               "Récupérer la clé RecoverEase",
		"recovery.meta.description":         "Récupérez votre clé RecoverEase Pro avec l'e-mail d'achat.",
		"recovery.backCheckout":             "Acheter Pro",
		"recovery.eyebrow":                  "Récupération de licence",
		"recovery.title":                    "Récupérer votre clé",
		"recovery.subtitle":                 "Saisissez l'e-mail utilisé lors de l'achat. Nous enverrons un lien sécurisé pour afficher votre clé et votre historique.",
		"recovery.aria":                     "Récupération de licence RecoverEase",
		"recovery.form.title":               "E-mail d'achat",
		"recovery.form.desc":                "Aucun mot de passe n'est requis. Le lien sécurisé est envoyé uniquement à l'e-mail d'achat.",
		"recovery.email":                    "Adresse e-mail",
		"recovery.emailHint":                "Utilisez le même e-mail que lors du paiement.",
		"recovery.submit":                   "Envoyer le lien",
		"recovery.note":                     "Si l'e-mail correspond à un achat, le lien arrivera dans quelques minutes.",
		"checkout.assurance.delivery.title": "Livraison instantanée",
		"checkout.assurance.delivery.desc":  "Voyez votre clé après paiement et conservez une copie dans votre boîte mail.",
		"checkout.assurance.privacy.title":  "Flux de récupération local",
		"checkout.assurance.privacy.desc":   "RecoverEase n'importe pas vos fichiers protégés pour la récupération.",
		"checkout.assurance.secure.title":   "Paiement chiffré",
		"checkout.assurance.secure.desc":    "Les informations de paiement passent par un flux sécurisé.",
	},
	"es": {
		"checkout.meta.title":               "Pago de RecoverEase Pro",
		"checkout.meta.description":         "Completa tu compra de RecoverEase Pro y obtén una licencia para recuperación avanzada de contraseñas.",
		"checkout.backHome":                 "Volver al sitio",
		"checkout.recovery.nav":             "Recuperar clave",
		"checkout.eyebrow":                  "Pago seguro",
		"checkout.title":                    "Obtén tu clave de RecoverEase Pro",
		"checkout.subtitle":                 "Introduce el correo para recibir la licencia y continúa a la página de pago segura.",
		"checkout.aria":                     "Pago de RecoverEase Pro",
		"checkout.account.title":            "Correo para recibir la licencia",
		"checkout.account.desc":             "Tras el pago, la clave de licencia y la confirmación de compra se enviarán a este correo.",
		"checkout.account.badges":           "Notas de entrega de licencia",
		"checkout.account.noLogin":          "No requiere cuenta",
		"checkout.account.recoverable":      "Clave enviada por correo",
		"checkout.email":                    "Correo electrónico",
		"checkout.emailPlaceholder":         "Introduce el correo de compra",
		"checkout.emailHint":                "Usa un correo al que puedas acceder más tarde si necesitas reenviar la clave.",
		"checkout.name":                     "Nombre del recibo",
		"checkout.payment.title":            "Método de pago",
		"checkout.payment.desc":             "Elige USDT o tarjeta de crédito y continúa a la página de pago segura.",
		"checkout.payment.usdt":             "USDT",
		"checkout.payment.card":             "Tarjeta de crédito",
		"checkout.cardNumber":               "Número de tarjeta",
		"checkout.expiry":                   "Fecha de vencimiento",
		"checkout.cvc":                      "Código de seguridad",
		"checkout.billing.title":            "Datos del recibo",
		"checkout.billing.desc":             "Se usa cuando sea necesario para impuestos, registros de pago y entrega del recibo.",
		"checkout.country":                  "País o región",
		"checkout.country.us":               "Estados Unidos",
		"checkout.country.cn":               "China",
		"checkout.country.eu":               "Unión Europea",
		"checkout.country.other":            "Otro",
		"checkout.license.title":            "Tipo de licencia",
		"checkout.license.monthly.title":    "Licencia de un mes",
		"checkout.license.monthly.desc":     "1 usuario, válida por un mes",
		"checkout.license.monthly.price":    "$9",
		"checkout.license.lifetime.title":   "Licencia permanente",
		"checkout.license.lifetime.desc":    "1 usuario, compra única",
		"checkout.license.lifetime.price":   "$29",
		"checkout.summary.title":            "Resumen de licencia",
		"checkout.summary.license":          "Licencia permanente, 1 usuario",
		"checkout.summary.total":            "Importe a pagar",
		"checkout.pay":                      "Pagar ahora",
		"checkout.terms":                    "Serás redirigido al proveedor de pago. Tras la confirmación, la clave se enviará a tu correo.",
		"checkout.recovery.title":           "¿Ya compraste?",
		"checkout.recovery.desc":            "Introduce el correo de compra para recibir un enlace seguro a tu clave y registro.",
		"checkout.recovery.email":           "Correo de compra",
		"checkout.recovery.submit":          "Recuperar clave",
		"recovery.meta.title":               "Recuperar clave de RecoverEase",
		"recovery.meta.description":         "Recupera tu clave de RecoverEase Pro con el correo de compra.",
		"recovery.backCheckout":             "Comprar Pro",
		"recovery.eyebrow":                  "Recuperación de licencia",
		"recovery.title":                    "Recupera tu clave",
		"recovery.subtitle":                 "Introduce el correo usado en la compra. Te enviaremos un enlace seguro para ver tu clave y registro.",
		"recovery.aria":                     "Recuperación de licencia RecoverEase",
		"recovery.form.title":               "Correo de compra",
		"recovery.form.desc":                "No se necesita contraseña. El enlace seguro solo se envía al correo de compra.",
		"recovery.email":                    "Correo electrónico",
		"recovery.emailHint":                "Usa el mismo correo introducido en el pago.",
		"recovery.submit":                   "Enviar enlace",
		"recovery.note":                     "Si el correo coincide con una compra, el enlace llegará en unos minutos.",
		"checkout.assurance.delivery.title": "Entrega instantánea",
		"checkout.assurance.delivery.desc":  "Consulta tu clave tras el pago y guarda una copia en tu correo.",
		"checkout.assurance.privacy.title":  "Flujo de recuperación local",
		"checkout.assurance.privacy.desc":   "RecoverEase no sube tus archivos protegidos para recuperarlos.",
		"checkout.assurance.secure.title":   "Pago cifrado",
		"checkout.assurance.secure.desc":    "La información de pago se envía mediante un flujo seguro.",
	},
	"pt": {
		"checkout.meta.title":               "Checkout RecoverEase Pro",
		"checkout.meta.description":         "Conclua a compra do RecoverEase Pro e obtenha uma licença para recuperação avançada de senhas.",
		"checkout.backHome":                 "Voltar ao site",
		"checkout.recovery.nav":             "Recuperar chave",
		"checkout.eyebrow":                  "Checkout seguro",
		"checkout.title":                    "Obtenha sua chave do RecoverEase Pro",
		"checkout.subtitle":                 "Informe o e-mail para receber a licença e continue para a página de pagamento segura.",
		"checkout.aria":                     "Checkout RecoverEase Pro",
		"checkout.account.title":            "E-mail para entrega da licença",
		"checkout.account.desc":             "Após o pagamento, a chave de licença e a confirmação de compra serão enviadas para este e-mail.",
		"checkout.account.badges":           "Notas de entrega da licença",
		"checkout.account.noLogin":          "Sem conta obrigatória",
		"checkout.account.recoverable":      "Chave enviada por e-mail",
		"checkout.email":                    "Endereço de e-mail",
		"checkout.emailPlaceholder":         "Informe o e-mail da compra",
		"checkout.emailHint":                "Use um e-mail que você consiga acessar depois, caso precise reenviar a chave.",
		"checkout.name":                     "Nome no recibo",
		"checkout.payment.title":            "Método de pagamento",
		"checkout.payment.desc":             "Escolha USDT ou cartão de crédito e continue para a página de pagamento segura.",
		"checkout.payment.usdt":             "USDT",
		"checkout.payment.card":             "Cartão de crédito",
		"checkout.cardNumber":               "Número do cartão",
		"checkout.expiry":                   "Data de validade",
		"checkout.cvc":                      "Código de segurança",
		"checkout.billing.title":            "Dados do recibo",
		"checkout.billing.desc":             "Usado quando necessário para impostos, registros de pagamento e entrega do recibo.",
		"checkout.country":                  "País ou região",
		"checkout.country.us":               "Estados Unidos",
		"checkout.country.cn":               "China",
		"checkout.country.eu":               "União Europeia",
		"checkout.country.other":            "Outro",
		"checkout.license.title":            "Tipo de licença",
		"checkout.license.monthly.title":    "Licença de um mês",
		"checkout.license.monthly.desc":     "1 usuário, válida por um mês",
		"checkout.license.monthly.price":    "$9",
		"checkout.license.lifetime.title":   "Licença vitalícia",
		"checkout.license.lifetime.desc":    "1 usuário, compra única",
		"checkout.license.lifetime.price":   "$29",
		"checkout.summary.title":            "Resumo da licença",
		"checkout.summary.license":          "Licença vitalícia, 1 usuário",
		"checkout.summary.total":            "Valor a pagar",
		"checkout.pay":                      "Pagar agora",
		"checkout.terms":                    "Você será redirecionado ao provedor de pagamento. Após a confirmação, a chave será enviada ao seu e-mail.",
		"checkout.recovery.title":           "Já comprou?",
		"checkout.recovery.desc":            "Informe o e-mail da compra para receber um link seguro para sua chave e registro.",
		"checkout.recovery.email":           "E-mail da compra",
		"checkout.recovery.submit":          "Recuperar chave",
		"recovery.meta.title":               "Recuperar chave RecoverEase",
		"recovery.meta.description":         "Recupere sua chave RecoverEase Pro pelo e-mail da compra.",
		"recovery.backCheckout":             "Comprar Pro",
		"recovery.eyebrow":                  "Recuperação de licença",
		"recovery.title":                    "Recupere sua chave",
		"recovery.subtitle":                 "Informe o e-mail usado na compra. Enviaremos um link seguro para ver sua chave e registro.",
		"recovery.aria":                     "Recuperação de licença RecoverEase",
		"recovery.form.title":               "E-mail da compra",
		"recovery.form.desc":                "Nenhuma senha é necessária. O link seguro é enviado apenas ao e-mail da compra.",
		"recovery.email":                    "Endereço de e-mail",
		"recovery.emailHint":                "Use o mesmo e-mail informado no checkout.",
		"recovery.submit":                   "Enviar link",
		"recovery.note":                     "Se o e-mail corresponder a uma compra, o link chegará em alguns minutos.",
		"checkout.assurance.delivery.title": "Entrega instantânea",
		"checkout.assurance.delivery.desc":  "Veja sua chave após o pagamento e mantenha uma cópia no e-mail.",
		"checkout.assurance.privacy.title":  "Fluxo local de recuperação",
		"checkout.assurance.privacy.desc":   "O RecoverEase não envia seus arquivos protegidos para recuperação.",
		"checkout.assurance.secure.title":   "Pagamento criptografado",
		"checkout.assurance.secure.desc":    "As informações de pagamento são enviadas por um fluxo seguro.",
	},
	"ru": {
		"checkout.meta.title":               "Оплата RecoverEase Pro",
		"checkout.meta.description":         "Завершите покупку RecoverEase Pro и получите лицензию для расширенного восстановления паролей.",
		"checkout.backHome":                 "Вернуться на сайт",
		"checkout.recovery.nav":             "Восстановить ключ",
		"checkout.eyebrow":                  "Безопасная оплата",
		"checkout.title":                    "Получите ключ RecoverEase Pro",
		"checkout.subtitle":                 "Введите e-mail для доставки лицензии, затем перейдите на безопасную страницу оплаты.",
		"checkout.aria":                     "Оплата RecoverEase Pro",
		"checkout.account.title":            "E-mail для доставки лицензии",
		"checkout.account.desc":             "После оплаты ключ лицензии и подтверждение покупки будут отправлены на этот адрес.",
		"checkout.account.badges":           "Информация о доставке лицензии",
		"checkout.account.noLogin":          "Аккаунт не нужен",
		"checkout.account.recoverable":      "Ключ отправляется по e-mail",
		"checkout.email":                    "Адрес e-mail",
		"checkout.emailPlaceholder":         "Введите e-mail покупки",
		"checkout.emailHint":                "Используйте e-mail, к которому сможете получить доступ позже для повторной отправки ключа.",
		"checkout.name":                     "Имя для квитанции",
		"checkout.payment.title":            "Способ оплаты",
		"checkout.payment.desc":             "Выберите USDT или кредитную карту, затем перейдите на безопасную страницу оплаты.",
		"checkout.payment.usdt":             "USDT",
		"checkout.payment.card":             "Кредитная карта",
		"checkout.cardNumber":               "Номер карты",
		"checkout.expiry":                   "Срок действия",
		"checkout.cvc":                      "Код безопасности",
		"checkout.billing.title":            "Данные квитанции",
		"checkout.billing.desc":             "Используются при необходимости для налогов, платежных записей и доставки квитанции.",
		"checkout.country":                  "Страна или регион",
		"checkout.country.us":               "США",
		"checkout.country.cn":               "Китай",
		"checkout.country.eu":               "Европейский союз",
		"checkout.country.other":            "Другое",
		"checkout.license.title":            "Тип лицензии",
		"checkout.license.monthly.title":    "Лицензия на месяц",
		"checkout.license.monthly.desc":     "1 пользователь, действует один месяц",
		"checkout.license.monthly.price":    "$9",
		"checkout.license.lifetime.title":   "Пожизненная лицензия",
		"checkout.license.lifetime.desc":    "1 пользователь, разовая покупка",
		"checkout.license.lifetime.price":   "$29",
		"checkout.summary.title":            "Сводка лицензии",
		"checkout.summary.license":          "Пожизненная лицензия, 1 пользователь",
		"checkout.summary.total":            "К оплате",
		"checkout.pay":                      "Оплатить сейчас",
		"checkout.terms":                    "Вы будете перенаправлены к платежному провайдеру. После подтверждения ключ будет отправлен на ваш e-mail.",
		"checkout.recovery.title":           "Уже покупали?",
		"checkout.recovery.desc":            "Введите e-mail покупки, чтобы получить безопасную ссылку на ключ и запись покупки.",
		"checkout.recovery.email":           "E-mail покупки",
		"checkout.recovery.submit":          "Восстановить ключ",
		"recovery.meta.title":               "Восстановить ключ RecoverEase",
		"recovery.meta.description":         "Восстановите ключ RecoverEase Pro по e-mail покупки.",
		"recovery.backCheckout":             "Купить Pro",
		"recovery.eyebrow":                  "Восстановление лицензии",
		"recovery.title":                    "Восстановите ключ лицензии",
		"recovery.subtitle":                 "Введите e-mail, использованный при покупке. Мы отправим безопасную ссылку для просмотра ключа и записи покупки.",
		"recovery.aria":                     "Восстановление лицензии RecoverEase",
		"recovery.form.title":               "E-mail покупки",
		"recovery.form.desc":                "Пароль не нужен. Безопасная ссылка отправляется только на e-mail покупки.",
		"recovery.email":                    "Адрес e-mail",
		"recovery.emailHint":                "Используйте тот же e-mail, который указали при оформлении.",
		"recovery.submit":                   "Отправить ссылку",
		"recovery.note":                     "Если e-mail соответствует покупке, ссылка придет в течение нескольких минут.",
		"checkout.assurance.delivery.title": "Мгновенная доставка лицензии",
		"checkout.assurance.delivery.desc":  "После подтверждения оплаты вы увидите ключ и сохраните копию в почте.",
		"checkout.assurance.privacy.title":  "Локальный процесс восстановления",
		"checkout.assurance.privacy.desc":   "RecoverEase не загружает ваши защищенные файлы для восстановления.",
		"checkout.assurance.secure.title":   "Зашифрованная оплата",
		"checkout.assurance.secure.desc":    "Платежные данные отправляются через безопасный платежный процесс.",
	},
}

var checkoutSuccessContent = map[string]map[string]string{
	"en": {
		"checkout.terms":                    "You will be redirected to the payment provider to complete payment. After confirmation, your license key appears on the payment result page.",
		"checkout.success.meta.title":       "RecoverEase Payment Result",
		"checkout.success.meta.description": "View your RecoverEase payment result and license key.",
		"checkout.success.eyebrow":          "Payment Result",
		"checkout.success.title":            "Your license key is ready",
		"checkout.success.subtitle":         "Keep this page open until you copy your license key. You can also recover it later with your purchase email.",
		"checkout.success.aria":             "RecoverEase payment result",
		"checkout.success.card.title":       "License key",
		"checkout.success.loading":          "Checking payment result...",
		"checkout.success.missing":          "Payment number is missing. Please return to checkout and try again.",
		"checkout.success.pending":          "Payment is not confirmed yet. Please refresh this page after the payment provider confirms it.",
		"checkout.success.failed":           "Unable to load payment result. Please try again later.",
		"checkout.success.copy":             "Copy",
		"checkout.success.copied":           "Copied",
		"checkout.success.order":            "Order",
		"checkout.success.plan":             "Plan",
		"checkout.success.issuedAt":         "Issued",
		"checkout.success.expiresAt":        "Expires",
		"checkout.success.lifetime":         "Lifetime",
		"checkout.success.backCheckout":     "Back to checkout",
		"checkout.success.recovery":         "Recover license key",
	},
	"zh": {
		"checkout.terms":                    "您将跳转到支付服务商页面完成付款。支付确认后，激活码会显示在支付结果页。",
		"checkout.success.meta.title":       "RecoverEase 支付结果",
		"checkout.success.meta.description": "查看 RecoverEase 支付结果和激活码。",
		"checkout.success.eyebrow":          "支付结果",
		"checkout.success.title":            "您的激活码已生成",
		"checkout.success.subtitle":         "请在复制激活码前保留此页面。之后也可以用购买邮箱找回激活码。",
		"checkout.success.aria":             "RecoverEase 支付结果",
		"checkout.success.card.title":       "激活码",
		"checkout.success.loading":          "正在查询支付结果...",
		"checkout.success.missing":          "缺少支付单号，请返回结算页重新尝试。",
		"checkout.success.pending":          "支付尚未确认。请在支付服务商确认后刷新此页面。",
		"checkout.success.failed":           "暂时无法加载支付结果，请稍后重试。",
		"checkout.success.copy":             "复制",
		"checkout.success.copied":           "已复制",
		"checkout.success.order":            "订单号",
		"checkout.success.plan":             "授权方案",
		"checkout.success.issuedAt":         "发放时间",
		"checkout.success.expiresAt":        "有效期至",
		"checkout.success.lifetime":         "永久有效",
		"checkout.success.backCheckout":     "返回结算页",
		"checkout.success.recovery":         "找回激活码",
	},
	"ja": {
		"checkout.terms":                    "支払いサービスのページへ移動して決済を完了します。確認後、ライセンスキーは支払い結果ページに表示されます。",
		"checkout.success.meta.title":       "RecoverEase 支払い結果",
		"checkout.success.meta.description": "RecoverEase の支払い結果とライセンスキーを確認します。",
		"checkout.success.eyebrow":          "支払い結果",
		"checkout.success.title":            "ライセンスキーが発行されました",
		"checkout.success.subtitle":         "ライセンスキーをコピーするまで、このページを開いたままにしてください。後から購入時のメールでも復元できます。",
		"checkout.success.aria":             "RecoverEase 支払い結果",
		"checkout.success.card.title":       "ライセンスキー",
		"checkout.success.loading":          "支払い結果を確認しています...",
		"checkout.success.missing":          "支払い番号がありません。決済ページに戻ってもう一度お試しください。",
		"checkout.success.pending":          "支払いはまだ確認されていません。支払いサービスで確認後、このページを更新してください。",
		"checkout.success.failed":           "支払い結果を読み込めません。しばらくしてからもう一度お試しください。",
		"checkout.success.copy":             "コピー",
		"checkout.success.copied":           "コピー済み",
		"checkout.success.order":            "注文",
		"checkout.success.plan":             "プラン",
		"checkout.success.issuedAt":         "発行日",
		"checkout.success.expiresAt":        "有効期限",
		"checkout.success.lifetime":         "無期限",
		"checkout.success.backCheckout":     "決済へ戻る",
		"checkout.success.recovery":         "ライセンスキーを復元",
	},
	"ko": {
		"checkout.terms":                    "결제 제공업체 페이지로 이동해 결제를 완료합니다. 확인 후 라이선스 키가 결제 결과 페이지에 표시됩니다.",
		"checkout.success.meta.title":       "RecoverEase 결제 결과",
		"checkout.success.meta.description": "RecoverEase 결제 결과와 라이선스 키를 확인하세요.",
		"checkout.success.eyebrow":          "결제 결과",
		"checkout.success.title":            "라이선스 키가 준비되었습니다",
		"checkout.success.subtitle":         "라이선스 키를 복사할 때까지 이 페이지를 열어 두세요. 나중에 구매 이메일로도 다시 찾을 수 있습니다.",
		"checkout.success.aria":             "RecoverEase 결제 결과",
		"checkout.success.card.title":       "라이선스 키",
		"checkout.success.loading":          "결제 결과를 확인하는 중...",
		"checkout.success.missing":          "결제 번호가 없습니다. 결제 페이지로 돌아가 다시 시도하세요.",
		"checkout.success.pending":          "결제가 아직 확인되지 않았습니다. 결제 제공업체 확인 후 이 페이지를 새로고침하세요.",
		"checkout.success.failed":           "결제 결과를 불러올 수 없습니다. 잠시 후 다시 시도하세요.",
		"checkout.success.copy":             "복사",
		"checkout.success.copied":           "복사됨",
		"checkout.success.order":            "주문",
		"checkout.success.plan":             "플랜",
		"checkout.success.issuedAt":         "발급일",
		"checkout.success.expiresAt":        "만료일",
		"checkout.success.lifetime":         "평생",
		"checkout.success.backCheckout":     "결제로 돌아가기",
		"checkout.success.recovery":         "라이선스 키 찾기",
	},
	"de": {
		"checkout.terms":                    "Sie werden zum Zahlungsanbieter weitergeleitet. Nach der Bestätigung erscheint Ihr Lizenzschlüssel auf der Zahlungsseite.",
		"checkout.success.meta.title":       "RecoverEase Zahlungsergebnis",
		"checkout.success.meta.description": "Sehen Sie Ihr RecoverEase Zahlungsergebnis und Ihren Lizenzschlüssel.",
		"checkout.success.eyebrow":          "Zahlungsergebnis",
		"checkout.success.title":            "Ihr Lizenzschlüssel ist bereit",
		"checkout.success.subtitle":         "Lassen Sie diese Seite geöffnet, bis Sie den Lizenzschlüssel kopiert haben. Später können Sie ihn mit Ihrer Kauf-E-Mail wiederherstellen.",
		"checkout.success.aria":             "RecoverEase Zahlungsergebnis",
		"checkout.success.card.title":       "Lizenzschlüssel",
		"checkout.success.loading":          "Zahlungsergebnis wird geprüft...",
		"checkout.success.missing":          "Die Zahlungsnummer fehlt. Bitte kehren Sie zur Kasse zurück und versuchen Sie es erneut.",
		"checkout.success.pending":          "Die Zahlung wurde noch nicht bestätigt. Aktualisieren Sie diese Seite nach der Bestätigung durch den Anbieter.",
		"checkout.success.failed":           "Das Zahlungsergebnis konnte nicht geladen werden. Bitte versuchen Sie es später erneut.",
		"checkout.success.copy":             "Kopieren",
		"checkout.success.copied":           "Kopiert",
		"checkout.success.order":            "Bestellung",
		"checkout.success.plan":             "Plan",
		"checkout.success.issuedAt":         "Ausgestellt",
		"checkout.success.expiresAt":        "Gültig bis",
		"checkout.success.lifetime":         "Lebenslang",
		"checkout.success.backCheckout":     "Zurück zur Kasse",
		"checkout.success.recovery":         "Lizenzschlüssel wiederherstellen",
	},
	"fr": {
		"checkout.terms":                    "Vous serez redirigé vers le prestataire de paiement. Après confirmation, votre clé apparaîtra sur la page de résultat du paiement.",
		"checkout.success.meta.title":       "Résultat du paiement RecoverEase",
		"checkout.success.meta.description": "Consultez le résultat de votre paiement RecoverEase et votre clé de licence.",
		"checkout.success.eyebrow":          "Résultat du paiement",
		"checkout.success.title":            "Votre clé de licence est prête",
		"checkout.success.subtitle":         "Gardez cette page ouverte jusqu'à ce que vous copiiez votre clé. Vous pourrez aussi la récupérer plus tard avec votre e-mail d'achat.",
		"checkout.success.aria":             "Résultat du paiement RecoverEase",
		"checkout.success.card.title":       "Clé de licence",
		"checkout.success.loading":          "Vérification du résultat du paiement...",
		"checkout.success.missing":          "Le numéro de paiement est manquant. Revenez au paiement et réessayez.",
		"checkout.success.pending":          "Le paiement n'est pas encore confirmé. Actualisez cette page après confirmation du prestataire.",
		"checkout.success.failed":           "Impossible de charger le résultat du paiement. Réessayez plus tard.",
		"checkout.success.copy":             "Copier",
		"checkout.success.copied":           "Copié",
		"checkout.success.order":            "Commande",
		"checkout.success.plan":             "Formule",
		"checkout.success.issuedAt":         "Émise le",
		"checkout.success.expiresAt":        "Expire le",
		"checkout.success.lifetime":         "À vie",
		"checkout.success.backCheckout":     "Retour au paiement",
		"checkout.success.recovery":         "Récupérer la clé",
	},
	"es": {
		"checkout.terms":                    "Serás redirigido al proveedor de pago. Tras la confirmación, tu clave aparecerá en la página de resultado del pago.",
		"checkout.success.meta.title":       "Resultado del pago de RecoverEase",
		"checkout.success.meta.description": "Consulta el resultado de pago de RecoverEase y tu clave de licencia.",
		"checkout.success.eyebrow":          "Resultado del pago",
		"checkout.success.title":            "Tu clave de licencia está lista",
		"checkout.success.subtitle":         "Mantén esta página abierta hasta copiar tu clave. También podrás recuperarla después con tu correo de compra.",
		"checkout.success.aria":             "Resultado del pago de RecoverEase",
		"checkout.success.card.title":       "Clave de licencia",
		"checkout.success.loading":          "Comprobando el resultado del pago...",
		"checkout.success.missing":          "Falta el número de pago. Vuelve al checkout e inténtalo de nuevo.",
		"checkout.success.pending":          "El pago aún no está confirmado. Actualiza esta página cuando el proveedor lo confirme.",
		"checkout.success.failed":           "No se pudo cargar el resultado del pago. Inténtalo más tarde.",
		"checkout.success.copy":             "Copiar",
		"checkout.success.copied":           "Copiado",
		"checkout.success.order":            "Pedido",
		"checkout.success.plan":             "Plan",
		"checkout.success.issuedAt":         "Emitida",
		"checkout.success.expiresAt":        "Caduca",
		"checkout.success.lifetime":         "De por vida",
		"checkout.success.backCheckout":     "Volver al checkout",
		"checkout.success.recovery":         "Recuperar clave",
	},
	"pt": {
		"checkout.terms":                    "Você será redirecionado ao provedor de pagamento. Após a confirmação, sua chave aparecerá na página de resultado do pagamento.",
		"checkout.success.meta.title":       "Resultado do pagamento RecoverEase",
		"checkout.success.meta.description": "Veja o resultado do pagamento RecoverEase e sua chave de licença.",
		"checkout.success.eyebrow":          "Resultado do pagamento",
		"checkout.success.title":            "Sua chave de licença está pronta",
		"checkout.success.subtitle":         "Mantenha esta página aberta até copiar sua chave. Você também poderá recuperá-la depois com o e-mail da compra.",
		"checkout.success.aria":             "Resultado do pagamento RecoverEase",
		"checkout.success.card.title":       "Chave de licença",
		"checkout.success.loading":          "Verificando resultado do pagamento...",
		"checkout.success.missing":          "O número do pagamento está ausente. Volte ao checkout e tente novamente.",
		"checkout.success.pending":          "O pagamento ainda não foi confirmado. Atualize esta página após a confirmação do provedor.",
		"checkout.success.failed":           "Não foi possível carregar o resultado do pagamento. Tente novamente mais tarde.",
		"checkout.success.copy":             "Copiar",
		"checkout.success.copied":           "Copiado",
		"checkout.success.order":            "Pedido",
		"checkout.success.plan":             "Plano",
		"checkout.success.issuedAt":         "Emitida em",
		"checkout.success.expiresAt":        "Expira em",
		"checkout.success.lifetime":         "Vitalícia",
		"checkout.success.backCheckout":     "Voltar ao checkout",
		"checkout.success.recovery":         "Recuperar chave",
	},
	"ru": {
		"checkout.terms":                    "Вы будете перенаправлены к платежному провайдеру. После подтверждения ключ появится на странице результата оплаты.",
		"checkout.success.meta.title":       "Результат оплаты RecoverEase",
		"checkout.success.meta.description": "Посмотрите результат оплаты RecoverEase и лицензионный ключ.",
		"checkout.success.eyebrow":          "Результат оплаты",
		"checkout.success.title":            "Ваш лицензионный ключ готов",
		"checkout.success.subtitle":         "Не закрывайте страницу, пока не скопируете ключ. Позже его можно восстановить по e-mail покупки.",
		"checkout.success.aria":             "Результат оплаты RecoverEase",
		"checkout.success.card.title":       "Лицензионный ключ",
		"checkout.success.loading":          "Проверяем результат оплаты...",
		"checkout.success.missing":          "Номер платежа отсутствует. Вернитесь к оплате и попробуйте снова.",
		"checkout.success.pending":          "Платеж еще не подтвержден. Обновите страницу после подтверждения провайдером.",
		"checkout.success.failed":           "Не удалось загрузить результат оплаты. Попробуйте позже.",
		"checkout.success.copy":             "Копировать",
		"checkout.success.copied":           "Скопировано",
		"checkout.success.order":            "Заказ",
		"checkout.success.plan":             "План",
		"checkout.success.issuedAt":         "Выдан",
		"checkout.success.expiresAt":        "Действует до",
		"checkout.success.lifetime":         "Бессрочно",
		"checkout.success.backCheckout":     "Вернуться к оплате",
		"checkout.success.recovery":         "Восстановить ключ",
	},
}

var recoveryVerificationContent = map[string]map[string]string{
	"en": {
		"checkout.recovery.desc":    "Enter the purchase email to receive a verification code before viewing your license key and purchase record.",
		"recovery.subtitle":         "Enter the email used at purchase. We will send a verification code before showing your license history.",
		"recovery.form.desc":        "No password is needed. The verification code is sent only to the purchase email.",
		"recovery.submit":           "Send verification code",
		"recovery.note":             "If the email matches a purchase, the verification code will arrive in a few minutes.",
		"recovery.emailPlaceholder": "Enter your purchase email",
		"recovery.code":             "Verification code",
		"recovery.codePlaceholder":  "Enter the 6-digit email code",
		"recovery.codeHint":         "Enter the 6-digit code sent to your purchase email.",
	},
	"zh": {
		"checkout.recovery.desc":    "输入购买时使用的邮箱，先接收验证码，验证后查看激活码和购买记录。",
		"recovery.subtitle":         "输入购买时使用的邮箱，我们会先发送验证码，验证通过后再显示历史激活码。",
		"recovery.form.desc":        "无需密码。验证码只会发送到购买时使用的邮箱。",
		"recovery.submit":           "发送验证码",
		"recovery.note":             "如果该邮箱存在购买记录，验证码会在几分钟内发送。",
		"recovery.emailPlaceholder": "请输入购买时使用的邮箱",
		"recovery.code":             "验证码",
		"recovery.codePlaceholder":  "请输入邮件中的 6 位验证码",
		"recovery.codeHint":         "请输入发送到购买邮箱的 6 位验证码。",
	},
	"ja": {
		"checkout.recovery.desc":    "購入時のメールを入力し、確認コードで認証してからライセンスキーと購入記録を表示します。",
		"recovery.subtitle":         "購入時に使用したメールを入力してください。確認コードで認証後、ライセンス履歴を表示します。",
		"recovery.form.desc":        "パスワードは不要です。確認コードは購入時のメールにのみ送信されます。",
		"recovery.submit":           "確認コードを送信",
		"recovery.note":             "メールが購入記録と一致する場合、確認コードが数分以内に届きます。",
		"recovery.emailPlaceholder": "購入時のメールを入力",
		"recovery.code":             "確認コード",
		"recovery.codePlaceholder":  "メールの6桁コードを入力",
		"recovery.codeHint":         "購入メールに届いた 6 桁のコードを入力してください。",
	},
	"ko": {
		"checkout.recovery.desc":    "구매 이메일을 입력하고 인증 코드를 받은 뒤 라이선스 키와 구매 기록을 확인하세요.",
		"recovery.subtitle":         "구매 시 사용한 이메일을 입력하세요. 인증 후 라이선스 기록을 보여드립니다.",
		"recovery.form.desc":        "비밀번호는 필요 없습니다. 인증 코드는 구매 이메일로만 전송됩니다.",
		"recovery.submit":           "인증 코드 보내기",
		"recovery.note":             "해당 이메일에 구매 기록이 있으면 몇 분 안에 인증 코드가 도착합니다.",
		"recovery.emailPlaceholder": "구매 이메일 입력",
		"recovery.code":             "인증 코드",
		"recovery.codePlaceholder":  "이메일의 6자리 코드 입력",
		"recovery.codeHint":         "구매 이메일로 받은 6자리 코드를 입력하세요.",
	},
	"de": {
		"checkout.recovery.desc":    "Geben Sie die Kauf-E-Mail ein und bestaetigen Sie sie per Code, bevor Lizenzschluessel und Kaufdaten angezeigt werden.",
		"recovery.subtitle":         "Geben Sie die beim Kauf verwendete E-Mail ein. Nach der Code-Bestaetigung zeigen wir Ihre Lizenzhistorie.",
		"recovery.form.desc":        "Kein Passwort erforderlich. Der Bestaetigungscode wird nur an die Kauf-E-Mail gesendet.",
		"recovery.submit":           "Bestaetigungscode senden",
		"recovery.note":             "Wenn die E-Mail zu einem Kauf passt, trifft der Code in wenigen Minuten ein.",
		"recovery.emailPlaceholder": "Kauf-E-Mail eingeben",
		"recovery.code":             "Bestaetigungscode",
		"recovery.codePlaceholder":  "6-stelligen E-Mail-Code eingeben",
		"recovery.codeHint":         "Geben Sie den 6-stelligen Code aus der Kauf-E-Mail ein.",
	},
	"fr": {
		"checkout.recovery.desc":    "Saisissez l'e-mail d'achat et confirmez-le avec un code avant d'afficher vos cles et votre historique.",
		"recovery.subtitle":         "Saisissez l'e-mail utilise lors de l'achat. Nous afficherons l'historique apres verification du code.",
		"recovery.form.desc":        "Aucun mot de passe n'est requis. Le code est envoye uniquement a l'e-mail d'achat.",
		"recovery.submit":           "Envoyer le code",
		"recovery.note":             "Si l'e-mail correspond a un achat, le code arrivera dans quelques minutes.",
		"recovery.emailPlaceholder": "Saisissez l'e-mail d'achat",
		"recovery.code":             "Code de verification",
		"recovery.codePlaceholder":  "Saisissez le code e-mail a 6 chiffres",
		"recovery.codeHint":         "Saisissez le code a 6 chiffres envoye a l'e-mail d'achat.",
	},
	"es": {
		"checkout.recovery.desc":    "Introduce el correo de compra y verificalo con un codigo antes de ver tu clave e historial.",
		"recovery.subtitle":         "Introduce el correo usado en la compra. Mostraremos el historial despues de verificar el codigo.",
		"recovery.form.desc":        "No se necesita contraseña. El codigo solo se envia al correo de compra.",
		"recovery.submit":           "Enviar codigo",
		"recovery.note":             "Si el correo coincide con una compra, el codigo llegara en unos minutos.",
		"recovery.emailPlaceholder": "Introduce el correo de compra",
		"recovery.code":             "Codigo de verificacion",
		"recovery.codePlaceholder":  "Introduce el codigo de 6 digitos",
		"recovery.codeHint":         "Introduce el codigo de 6 digitos enviado al correo de compra.",
	},
	"pt": {
		"checkout.recovery.desc":    "Informe o e-mail da compra e confirme com um codigo antes de ver sua chave e registro.",
		"recovery.subtitle":         "Informe o e-mail usado na compra. Mostraremos o historico depois de verificar o codigo.",
		"recovery.form.desc":        "Nenhuma senha e necessaria. O codigo e enviado apenas ao e-mail da compra.",
		"recovery.submit":           "Enviar codigo",
		"recovery.note":             "Se o e-mail corresponder a uma compra, o codigo chegara em alguns minutos.",
		"recovery.emailPlaceholder": "Informe o e-mail da compra",
		"recovery.code":             "Codigo de verificacao",
		"recovery.codePlaceholder":  "Digite o codigo de 6 digitos",
		"recovery.codeHint":         "Digite o codigo de 6 digitos enviado ao e-mail da compra.",
	},
	"ru": {
		"checkout.recovery.desc":    "Введите e-mail покупки и подтвердите его кодом, чтобы увидеть ключ и историю покупки.",
		"recovery.subtitle":         "Введите e-mail, использованный при покупке. После проверки кода мы покажем историю лицензий.",
		"recovery.form.desc":        "Пароль не нужен. Код подтверждения отправляется только на e-mail покупки.",
		"recovery.submit":           "Отправить код",
		"recovery.note":             "Если e-mail соответствует покупке, код придет в течение нескольких минут.",
		"recovery.emailPlaceholder": "Введите e-mail покупки",
		"recovery.code":             "Код подтверждения",
		"recovery.codePlaceholder":  "Введите 6-значный код из письма",
		"recovery.codeHint":         "Введите 6-значный код, отправленный на e-mail покупки.",
	},
}

func applyCheckoutContent() {
	for code, labels := range checkoutContent {
		if locale, ok := translations[code]; ok {
			for key, value := range labels {
				locale[key] = value
			}
		}
	}
	for code, labels := range checkoutSuccessContent {
		if locale, ok := translations[code]; ok {
			for key, value := range labels {
				locale[key] = value
			}
		}
	}
	for code, labels := range recoveryVerificationContent {
		if locale, ok := translations[code]; ok {
			for key, value := range labels {
				locale[key] = value
			}
		}
	}
}

func applyRecoveryModeOverrides() {
	for code, overrides := range recoveryModeOverrides {
		if locale, ok := translations[code]; ok {
			for key, value := range overrides {
				locale[key] = value
			}
		}
	}
}

func addLocale(code string, overrides map[string]string) {
	translations[code] = merge(en, overrides)
}

func Languages(activeCode string) []Language {
	return LanguagesForPath(activeCode, "")
}

func LanguagesForPath(activeCode, suffix string) []Language {
	langs := make([]Language, len(languageDefinitions))
	for i, lang := range languageDefinitions {
		lang.Active = lang.Code == activeCode
		if suffix != "" {
			lang.Path = strings.TrimRight(lang.Path, "/") + suffix
		}
		langs[i] = lang
	}
	return langs
}

func Alternates(baseURL string) []Language {
	return AlternatesForPath(baseURL, "")
}

func AlternatesForPath(baseURL, suffix string) []Language {
	langs := make([]Language, len(languageDefinitions))
	for i, lang := range languageDefinitions {
		path := lang.Path
		if suffix != "" {
			path = strings.TrimRight(path, "/") + suffix
		}
		lang.Path = absoluteURL(baseURL, path)
		langs[i] = lang
	}
	return langs
}

func Supported(code string) bool {
	_, ok := translations[code]
	return ok
}

func Normalize(code string) string {
	code = strings.ToLower(strings.TrimSpace(code))
	if Supported(code) {
		return code
	}
	return DefaultLocale
}

func MatchAcceptLanguage(header string) string {
	bestLocale := DefaultLocale
	bestQ := -1.0

	for _, item := range strings.Split(header, ",") {
		tag, q := parseAcceptLanguageItem(item)
		if q <= 0 || q <= bestQ {
			continue
		}

		if locale, ok := matchLanguageTag(tag); ok {
			bestLocale = locale
			bestQ = q
		}
	}

	return bestLocale
}

func parseAcceptLanguageItem(item string) (string, float64) {
	parts := strings.Split(item, ";")
	tag := strings.ToLower(strings.TrimSpace(parts[0]))
	q := 1.0

	for _, param := range parts[1:] {
		keyValue := strings.SplitN(strings.TrimSpace(param), "=", 2)
		if len(keyValue) != 2 || strings.ToLower(strings.TrimSpace(keyValue[0])) != "q" {
			continue
		}

		if parsed, err := strconv.ParseFloat(strings.TrimSpace(keyValue[1]), 64); err == nil {
			q = parsed
		}
	}

	return tag, q
}

func matchLanguageTag(tag string) (string, bool) {
	if tag == "" || tag == "*" {
		return "", false
	}

	tag = strings.ReplaceAll(tag, "_", "-")
	if Supported(tag) {
		return tag, true
	}

	base, _, _ := strings.Cut(tag, "-")
	if Supported(base) {
		return base, true
	}

	return "", false
}

func HTMLLang(code string) string {
	code = Normalize(code)
	for _, lang := range languageDefinitions {
		if lang.Code == code {
			return lang.HTMLLang
		}
	}
	return "zh-CN"
}

func Path(code string) string {
	code = Normalize(code)
	for _, lang := range languageDefinitions {
		if lang.Code == code {
			return lang.Path
		}
	}
	return "/"
}

func CheckoutPath(code string) string {
	return strings.TrimRight(Path(code), "/") + "/checkout"
}

func CheckoutSuccessPath(code string) string {
	return strings.TrimRight(CheckoutPath(code), "/") + "/success"
}

func LicenseRecoveryPath(code string) string {
	return strings.TrimRight(Path(code), "/") + "/license-recovery"
}

func PrivacyPath(code string) string {
	return strings.TrimRight(Path(code), "/") + "/privacy"
}

func TermsPath(code string) string {
	return strings.TrimRight(Path(code), "/") + "/terms"
}

func T(code, key string) string {
	code = Normalize(code)
	if value, ok := translations[code][key]; ok {
		return value
	}
	if value, ok := en[key]; ok {
		return value
	}
	return key
}

func Map(code string) map[string]string {
	code = Normalize(code)
	return translations[code]
}

func merge(base, overrides map[string]string) map[string]string {
	result := make(map[string]string, len(base)+len(overrides))
	for key, value := range base {
		result[key] = value
	}
	for key, value := range overrides {
		result[key] = value
	}
	return result
}

func absoluteURL(baseURL, path string) string {
	return strings.TrimRight(baseURL, "/") + path
}
