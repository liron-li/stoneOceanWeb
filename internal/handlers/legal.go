package handlers

import (
	"net/http"
	"strings"

	"stone-ocean-web/internal/i18n"

	"github.com/gin-gonic/gin"
)

type legalPage struct {
	Kind        string
	Eyebrow     string
	Title       string
	Description string
	Updated     string
	Intro       string
	Sections    []legalSection
}

type legalSection struct {
	Title      string
	Paragraphs []string
	Bullets    []string
}

func Privacy(c *gin.Context) {
	locale := i18n.MatchAcceptLanguage(c.GetHeader("Accept-Language"))
	c.Redirect(http.StatusFound, i18n.PrivacyPath(locale))
}

func Terms(c *gin.Context) {
	locale := i18n.MatchAcceptLanguage(c.GetHeader("Accept-Language"))
	c.Redirect(http.StatusFound, i18n.TermsPath(locale))
}

func LocalizedPrivacy(c *gin.Context) {
	locale := c.Param("locale")
	if !i18n.Supported(locale) {
		c.Redirect(http.StatusMovedPermanently, i18n.PrivacyPath(i18n.DefaultLocale))
		return
	}

	renderLegal(c, locale, "privacy")
}

func LocalizedTerms(c *gin.Context) {
	locale := c.Param("locale")
	if !i18n.Supported(locale) {
		c.Redirect(http.StatusMovedPermanently, i18n.TermsPath(i18n.DefaultLocale))
		return
	}

	renderLegal(c, locale, "terms")
}

func renderLegal(c *gin.Context, locale, kind string) {
	page := legalContent(locale, kind)
	baseURL := requestBaseURL(c)
	path := legalPath(locale, kind)
	suffix := legalSuffix(kind)

	c.HTML(http.StatusOK, "legal.tmpl", gin.H{
		"Title":               page.Title,
		"Description":         page.Description,
		"Locale":              locale,
		"HTMLLang":            i18n.HTMLLang(locale),
		"Canonical":           strings.TrimRight(baseURL, "/") + path,
		"DefaultURL":          strings.TrimRight(baseURL, "/") + suffix,
		"Alternates":          i18n.AlternatesForPath(baseURL, suffix),
		"Languages":           i18n.LanguagesForPath(locale, suffix),
		"HomePath":            i18n.Path(locale),
		"CheckoutPath":        i18n.CheckoutPath(locale),
		"LicenseRecoveryPath": i18n.LicenseRecoveryPath(locale),
		"PrivacyPath":         i18n.PrivacyPath(locale),
		"TermsPath":           i18n.TermsPath(locale),
		"Page":                page,
		"T": func(key string) string {
			return i18n.T(locale, key)
		},
	})
}

func legalPath(locale, kind string) string {
	if kind == "terms" {
		return i18n.TermsPath(locale)
	}
	return i18n.PrivacyPath(locale)
}

func legalSuffix(kind string) string {
	if kind == "terms" {
		return "/terms"
	}
	return "/privacy"
}

func legalContent(locale, kind string) legalPage {
	if kind == "terms" {
		if locale == "zh" {
			return zhTerms
		}
		return enTerms
	}
	if locale == "zh" {
		return zhPrivacy
	}
	return enPrivacy
}

var enPrivacy = legalPage{
	Kind:        "privacy",
	Eyebrow:     "Legal",
	Title:       "Privacy Policy",
	Description: "How RecoverEase collects, uses, protects, and shares information.",
	Updated:     "Last updated: July 1, 2026",
	Intro:       "RecoverEase is designed around local password recovery. Your protected files and recovery work normally stay on your own device. This policy explains what information we collect through the website, checkout, license recovery, license activation, product updates, and support.",
	Sections: []legalSection{
		{
			Title: "Information we collect",
			Bullets: []string{
				"Purchase and contact information, such as your email address, order number, license key status, and support messages.",
				"Payment-related records from payment providers, such as transaction status, amount, currency, and fraud-prevention signals. We do not store full payment card numbers.",
				"License and product information, such as activation status, software version, device or installation identifiers, update checks, and timestamps.",
				"Website and technical data, such as IP address, browser, device type, language preference, referring pages, logs, and approximate location derived from IP address.",
				"Recovery task data stays local by default. RecoverEase does not upload your protected files, passwords, or password hints for recovery unless you deliberately send information to support.",
			},
		},
		{
			Title: "How we use information",
			Bullets: []string{
				"To provide the website, checkout, license delivery, license recovery, activation, updates, and customer support.",
				"To process payments, prevent fraud, troubleshoot problems, and maintain service security.",
				"To send transactional emails, including purchase confirmations, license keys, verification codes, and important service notices.",
				"To improve product reliability, website performance, and user experience using aggregated or minimized technical data.",
				"To comply with legal obligations, enforce our terms, and protect the rights, property, and safety of users and RecoverEase.",
			},
		},
		{
			Title: "Legal bases and user choices",
			Paragraphs: []string{
				"Where applicable, we process personal information to perform a contract with you, comply with legal obligations, pursue legitimate interests such as security and product improvement, or based on your consent. You can decline optional communications and may contact us to exercise privacy rights available in your region.",
			},
		},
		{
			Title: "How we share information",
			Paragraphs: []string{
				"We do not sell your personal information. We share information only when needed to operate the service or as permitted by law.",
			},
			Bullets: []string{
				"Payment processors, email providers, hosting providers, analytics/security tools, and other vendors that process information for us.",
				"Professional advisers, authorities, or third parties when required by law, to protect rights and safety, or in connection with a business transfer.",
				"Support information you choose to provide may be reviewed by our support team or service providers helping us resolve your request.",
			},
		},
		{
			Title: "Data retention",
			Paragraphs: []string{
				"We keep personal information only as long as reasonably necessary for the purposes described in this policy, including providing licenses and support, maintaining security records, complying with tax/accounting requirements, resolving disputes, and enforcing agreements. When information is no longer needed, we delete, de-identify, or securely retain it as required by law.",
			},
		},
		{
			Title: "Security",
			Paragraphs: []string{
				"We use reasonable administrative, technical, and organizational safeguards designed to protect information we keep. No method of transmission or storage is perfectly secure, so we cannot guarantee absolute security. You should keep your license key, email account, and device secure.",
			},
		},
		{
			Title: "Your rights",
			Paragraphs: []string{
				"Depending on your location, you may have rights to access, correct, delete, export, or restrict the use of your personal information; object to certain processing; withdraw consent; or opt out of sale or sharing where applicable. We will not discriminate against you for exercising privacy rights. We may need to verify your identity before fulfilling a request.",
			},
		},
		{
			Title: "International users",
			Paragraphs: []string{
				"RecoverEase may process information in countries other than where you live. When required, we use appropriate safeguards for cross-border transfers.",
			},
		},
		{
			Title: "Children",
			Paragraphs: []string{
				"RecoverEase is not directed to children under 13, and we do not knowingly collect personal information from children.",
			},
		},
		{
			Title: "Changes and contact",
			Paragraphs: []string{
				"We may update this policy from time to time. The updated date will show when the policy was last revised. For privacy requests or questions, contact support@recoverease.com.",
			},
		},
	},
}

var zhPrivacy = legalPage{
	Kind:        "privacy",
	Eyebrow:     "法律信息",
	Title:       "隐私政策",
	Description: "RecoverEase 如何收集、使用、保护和共享信息。",
	Updated:     "最后更新：2026 年 7 月 1 日",
	Intro:       "RecoverEase 围绕本地密码恢复流程设计。通常情况下，您的受保护文件和恢复任务会保留在自己的设备上。本政策说明我们通过网站、结账、激活码找回、授权激活、产品更新和客服支持收集与处理哪些信息。",
	Sections: []legalSection{
		{
			Title: "我们收集的信息",
			Bullets: []string{
				"购买与联系信息，例如邮箱地址、订单号、激活码状态和客服沟通内容。",
				"支付服务商返回的支付相关记录，例如交易状态、金额、币种和风控信号。我们不保存完整银行卡号。",
				"授权与产品信息，例如激活状态、软件版本、设备或安装标识、更新检查记录和时间戳。",
				"网站与技术信息，例如 IP 地址、浏览器、设备类型、语言偏好、来源页面、日志，以及根据 IP 推断的大致地区。",
				"恢复任务数据默认保留在本地。除非您主动发送给客服，RecoverEase 不会为了恢复任务上传您的受保护文件、密码或密码线索。",
			},
		},
		{
			Title: "我们如何使用信息",
			Bullets: []string{
				"提供网站、结账、激活码交付、激活码找回、授权激活、产品更新和客服支持。",
				"处理支付、防范欺诈、排查问题，并维护服务安全。",
				"发送交易类邮件，包括购买确认、激活码、验证码和重要服务通知。",
				"使用汇总化或最小化的技术数据改进产品稳定性、网站性能和用户体验。",
				"履行法律义务、执行服务条款，并保护用户和 RecoverEase 的权利、财产与安全。",
			},
		},
		{
			Title: "处理依据与用户选择",
			Paragraphs: []string{
				"在适用法律要求下，我们会基于履行合同、遵守法律义务、维护安全和改进产品等正当利益，或基于您的同意处理个人信息。您可以拒绝可选营销沟通，也可以联系我们行使所在地区法律赋予的隐私权利。",
			},
		},
		{
			Title: "我们如何共享信息",
			Paragraphs: []string{
				"我们不会出售您的个人信息。只有在运营服务所需或法律允许的情况下，我们才会共享信息。",
			},
			Bullets: []string{
				"支付处理、邮件发送、托管、分析/安全工具等代表我们处理信息的服务商。",
				"在法律要求、保护权利与安全，或业务转让相关场景下，向专业顾问、主管机关或第三方披露必要信息。",
				"您主动提供给客服的支持材料，可能由客服团队或协助处理问题的服务商查看。",
			},
		},
		{
			Title: "数据保留",
			Paragraphs: []string{
				"我们只会在实现本政策所述目的所需的合理期限内保留个人信息，包括提供授权和支持、维护安全记录、满足税务/会计要求、解决争议和执行协议。当信息不再需要时，我们会删除、去标识化，或依法安全保留。",
			},
		},
		{
			Title: "安全措施",
			Paragraphs: []string{
				"我们采用合理的管理、技术和组织措施保护所保留的信息。但任何传输或存储方式都无法保证绝对安全。您也应妥善保护自己的激活码、邮箱账号和设备。",
			},
		},
		{
			Title: "您的权利",
			Paragraphs: []string{
				"根据您所在地区的法律，您可能有权访问、更正、删除、导出或限制使用个人信息，反对某些处理，撤回同意，或在适用情况下选择退出出售或共享。我们不会因您行使隐私权利而歧视您。处理请求前，我们可能需要验证您的身份。",
			},
		},
		{
			Title: "国际用户",
			Paragraphs: []string{
				"RecoverEase 可能会在您所在国家或地区以外处理信息。法律要求时，我们会为跨境传输采取适当保护措施。",
			},
		},
		{
			Title: "儿童隐私",
			Paragraphs: []string{
				"RecoverEase 不面向 13 岁以下儿童，我们也不会有意收集儿童的个人信息。",
			},
		},
		{
			Title: "政策变更与联系",
			Paragraphs: []string{
				"我们可能会不时更新本政策。页面上的更新日期表示最近修订时间。如需提出隐私请求或咨询问题，请联系 support@recoverease.com。",
			},
		},
	},
}

var enTerms = legalPage{
	Kind:        "terms",
	Eyebrow:     "Legal",
	Title:       "Terms of Service",
	Description: "Terms governing your use of the RecoverEase website, software, checkout, licenses, and support.",
	Updated:     "Last updated: July 1, 2026",
	Intro:       "These Terms of Service govern your access to and use of the RecoverEase website, software, checkout, license keys, updates, documentation, and support. By using RecoverEase, you agree to these terms.",
	Sections: []legalSection{
		{
			Title: "Eligibility and acceptance",
			Paragraphs: []string{
				"You may use RecoverEase only if you can form a binding agreement and comply with applicable law. If you use RecoverEase on behalf of an organization, you represent that you have authority to bind that organization.",
			},
		},
		{
			Title: "Software license",
			Paragraphs: []string{
				"Subject to these terms and payment of applicable fees, RecoverEase grants you a limited, non-exclusive, non-transferable license to install and use the software for legitimate password recovery tasks. The software is licensed, not sold. RecoverEase and its licensors retain all rights not expressly granted.",
			},
			Bullets: []string{
				"You may not copy, resell, sublicense, rent, lease, or distribute the software except as expressly allowed.",
				"You may not reverse engineer, bypass licensing controls, remove notices, or use the software to build a competing product, except where law expressly permits.",
				"You are responsible for keeping your license key confidential and for activity using your license.",
			},
		},
		{
			Title: "Lawful and authorized use",
			Paragraphs: []string{
				"RecoverEase is intended to help recover access to files, documents, disks, or systems that you own or are authorized to access. You must not use RecoverEase for unauthorized access, credential theft, circumvention of systems you do not control, malware, harassment, or any unlawful activity.",
			},
		},
		{
			Title: "Purchases, delivery, and refunds",
			Paragraphs: []string{
				"Prices, taxes, features, and license terms are shown at checkout and may change for future purchases. Payments are processed by third-party payment providers. After payment confirmation, your license key may be delivered by email and/or shown on the payment result page.",
				"Refunds are handled according to the refund information presented at purchase, applicable law, and support review. If you experience an activation or delivery issue, contact support@recoverease.com so we can help resolve it.",
			},
		},
		{
			Title: "Updates and availability",
			Paragraphs: []string{
				"We may provide updates, patches, or changes to features from time to time. Some features may require internet access for activation, updates, or support. We may modify, suspend, or discontinue parts of the website or service where reasonably necessary.",
			},
		},
		{
			Title: "Support and user materials",
			Paragraphs: []string{
				"Support is provided through the channels and plans available at the time. If you send files, screenshots, logs, or other materials to support, you confirm that you have the right to share them and authorize us to use them to investigate and resolve your request.",
			},
		},
		{
			Title: "Privacy",
			Paragraphs: []string{
				"Our Privacy Policy explains how we collect and use information. By using RecoverEase, you acknowledge that information will be handled according to that policy.",
			},
		},
		{
			Title: "No recovery guarantee",
			Paragraphs: []string{
				"Password recovery results depend on password complexity, file type, available hints, hardware, time, and other factors. RecoverEase does not guarantee that every password can be recovered or that a task will complete within a particular time.",
			},
		},
		{
			Title: "Disclaimers",
			Paragraphs: []string{
				"To the maximum extent permitted by law, RecoverEase is provided as is and as available, without warranties of merchantability, fitness for a particular purpose, non-infringement, uninterrupted operation, or error-free performance. Mandatory consumer rights are not affected.",
			},
		},
		{
			Title: "Limitation of liability",
			Paragraphs: []string{
				"To the maximum extent permitted by law, RecoverEase will not be liable for indirect, incidental, special, consequential, exemplary, or punitive damages, or for lost profits, lost data, business interruption, or unauthorized access to your systems. Our total liability for any claim is limited to the amount you paid for the product or service giving rise to the claim in the 12 months before the event, unless applicable law requires otherwise.",
			},
		},
		{
			Title: "Termination",
			Paragraphs: []string{
				"We may suspend or terminate access to the website, support, updates, or licensing services if you materially breach these terms, use RecoverEase unlawfully, create security risk, or fail to pay applicable fees. You may stop using RecoverEase at any time.",
			},
		},
		{
			Title: "Changes and contact",
			Paragraphs: []string{
				"We may update these terms from time to time. The updated date will show when the terms were last revised. For questions, contact support@recoverease.com.",
			},
		},
	},
}

var zhTerms = legalPage{
	Kind:        "terms",
	Eyebrow:     "法律信息",
	Title:       "服务条款",
	Description: "适用于 RecoverEase 网站、软件、结账、授权和支持服务的使用条款。",
	Updated:     "最后更新：2026 年 7 月 1 日",
	Intro:       "本服务条款适用于您访问和使用 RecoverEase 网站、软件、结账流程、激活码、更新、文档和客服支持。使用 RecoverEase 即表示您同意本条款。",
	Sections: []legalSection{
		{
			Title: "资格与接受",
			Paragraphs: []string{
				"您只能在具备订立有效协议能力并遵守适用法律的情况下使用 RecoverEase。如果您代表组织使用 RecoverEase，即表示您有权代表该组织接受本条款。",
			},
		},
		{
			Title: "软件授权",
			Paragraphs: []string{
				"在您遵守本条款并支付适用费用的前提下，RecoverEase 授予您有限的、非独占的、不可转让的软件安装和使用授权，用于合法的密码恢复任务。本软件是授权使用而非出售。未明确授予的权利均由 RecoverEase 及其许可方保留。",
			},
			Bullets: []string{
				"除非明确允许，您不得复制、转售、再授权、出租、租赁或分发本软件。",
				"除非法律明确允许，您不得反向工程、绕过授权控制、移除声明，或使用本软件开发竞争产品。",
				"您应妥善保管激活码，并对使用该授权产生的活动负责。",
			},
		},
		{
			Title: "合法且经授权的使用",
			Paragraphs: []string{
				"RecoverEase 旨在帮助您恢复自己拥有或已获授权访问的文件、文档、磁盘或系统。您不得将 RecoverEase 用于未授权访问、窃取凭据、绕过不受您控制的系统、恶意软件、骚扰或任何违法活动。",
			},
		},
		{
			Title: "购买、交付与退款",
			Paragraphs: []string{
				"价格、税费、功能和授权期限以结账页面展示为准，并可能对未来购买发生变化。支付由第三方支付服务商处理。支付确认后，激活码可能通过邮件发送，也可能显示在支付结果页。",
				"退款将依据购买时展示的退款说明、适用法律以及客服审核处理。如果您遇到激活或交付问题，请联系 support@recoverease.com，我们会协助排查和解决。",
			},
		},
		{
			Title: "更新与可用性",
			Paragraphs: []string{
				"我们可能不时提供更新、补丁或功能调整。部分功能可能需要联网完成激活、更新或客服支持。在合理必要的情况下，我们可能修改、暂停或停止网站或服务的部分内容。",
			},
		},
		{
			Title: "客服支持与用户材料",
			Paragraphs: []string{
				"客服支持按当时可用的渠道和方案提供。如果您向客服发送文件、截图、日志或其他材料，即表示您有权分享这些材料，并授权我们为调查和解决您的请求而使用这些材料。",
			},
		},
		{
			Title: "隐私",
			Paragraphs: []string{
				"我们的隐私政策说明了我们如何收集和使用信息。使用 RecoverEase 即表示您知悉相关信息将按照隐私政策处理。",
			},
		},
		{
			Title: "不保证一定恢复成功",
			Paragraphs: []string{
				"密码恢复结果取决于密码复杂度、文件类型、可用线索、硬件性能、投入时间和其他因素。RecoverEase 不保证所有密码都能恢复，也不保证任务能在特定时间内完成。",
			},
		},
		{
			Title: "免责声明",
			Paragraphs: []string{
				"在法律允许的最大范围内，RecoverEase 按现状和可用状态提供，不作关于适销性、特定用途适用性、不侵权、持续可用或无错误运行的保证。法律强制规定的消费者权利不受影响。",
			},
		},
		{
			Title: "责任限制",
			Paragraphs: []string{
				"在法律允许的最大范围内，RecoverEase 不对间接、附带、特殊、后果性、惩罚性或惩戒性损害，利润损失、数据丢失、业务中断或系统被未授权访问承担责任。除非适用法律另有要求，我们对任何索赔的总责任以相关事件发生前 12 个月内您为引发索赔的产品或服务实际支付的金额为限。",
			},
		},
		{
			Title: "终止",
			Paragraphs: []string{
				"如果您严重违反本条款、违法使用 RecoverEase、造成安全风险或未支付适用费用，我们可能暂停或终止您访问网站、支持、更新或授权服务的权限。您也可以随时停止使用 RecoverEase。",
			},
		},
		{
			Title: "条款变更与联系",
			Paragraphs: []string{
				"我们可能不时更新本条款。页面上的更新日期表示最近修订时间。如有问题，请联系 support@recoverease.com。",
			},
		},
	},
}
