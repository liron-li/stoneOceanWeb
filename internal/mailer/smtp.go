package mailer

import (
	"context"
	"crypto/tls"
	"errors"
	"fmt"
	"html"
	"io"
	"mime"
	"net"
	"net/mail"
	"net/smtp"
	"strconv"
	"strings"
	"time"

	"stone-ocean-web/internal/config"
	"stone-ocean-web/internal/store"
)

type SMTPMailer struct {
	cfg config.EmailConfig
}

func NewSMTPMailer(cfg config.EmailConfig) *SMTPMailer {
	return &SMTPMailer{cfg: cfg}
}

func (m *SMTPMailer) SendLicense(ctx context.Context, license *store.License) error {
	if m == nil || !m.cfg.Enabled {
		return nil
	}
	if license == nil {
		return nil
	}
	if err := m.validate(); err != nil {
		return err
	}

	to := strings.TrimSpace(license.Order.Email)
	if to == "" && license.Customer.Email != "" {
		to = strings.TrimSpace(license.Customer.Email)
	}
	if to == "" {
		return errors.New("license email recipient is empty")
	}

	message, err := m.buildLicenseMessage(to, license)
	if err != nil {
		return err
	}
	return m.send(ctx, []string{to}, message)
}

func (m *SMTPMailer) SendRecoveryCode(ctx context.Context, email, code, locale string, expiresAt time.Time) error {
	if m == nil || !m.cfg.Enabled {
		return nil
	}
	if err := m.validate(); err != nil {
		return err
	}

	to := strings.TrimSpace(email)
	if to == "" {
		return errors.New("recovery code email recipient is empty")
	}

	message, err := m.buildRecoveryCodeMessage(to, code, locale, expiresAt)
	if err != nil {
		return err
	}
	return m.send(ctx, []string{to}, message)
}

func (m *SMTPMailer) validate() error {
	if strings.TrimSpace(m.cfg.Host) == "" {
		return errors.New("email host is required")
	}
	if strings.TrimSpace(m.cfg.Port) == "" {
		return errors.New("email port is required")
	}
	if strings.TrimSpace(m.cfg.Username) == "" {
		return errors.New("email username is required")
	}
	if strings.TrimSpace(m.cfg.Password) == "" {
		return errors.New("email password is required")
	}
	if strings.TrimSpace(m.cfg.FromAddress) == "" {
		return errors.New("email from address is required")
	}
	return nil
}

func (m *SMTPMailer) buildLicenseMessage(to string, license *store.License) ([]byte, error) {
	from := mail.Address{Name: m.cfg.FromName, Address: m.cfg.FromAddress}
	recipient := mail.Address{Address: to}
	copy := licenseEmailCopy(license.Order.Locale)
	boundary := "recoverease-license-boundary"

	headers := map[string]string{
		"From":         from.String(),
		"To":           recipient.String(),
		"Subject":      mime.QEncoding.Encode("UTF-8", copy.Subject),
		"Date":         time.Now().Format(time.RFC1123Z),
		"MIME-Version": "1.0",
		"Content-Type": fmt.Sprintf(`multipart/alternative; boundary="%s"`, boundary),
	}
	if strings.TrimSpace(m.cfg.ReplyTo) != "" {
		headers["Reply-To"] = strings.TrimSpace(m.cfg.ReplyTo)
	}

	var builder strings.Builder
	for _, key := range []string{"From", "To", "Reply-To", "Subject", "Date", "MIME-Version", "Content-Type"} {
		value, ok := headers[key]
		if !ok {
			continue
		}
		builder.WriteString(key)
		builder.WriteString(": ")
		builder.WriteString(value)
		builder.WriteString("\r\n")
	}
	builder.WriteString("\r\n")
	builder.WriteString("--")
	builder.WriteString(boundary)
	builder.WriteString("\r\n")
	builder.WriteString("Content-Type: text/plain; charset=\"UTF-8\"\r\n")
	builder.WriteString("Content-Transfer-Encoding: 8bit\r\n\r\n")
	builder.WriteString(licenseEmailTextBody(license, copy))
	builder.WriteString("\r\n--")
	builder.WriteString(boundary)
	builder.WriteString("\r\n")
	builder.WriteString("Content-Type: text/html; charset=\"UTF-8\"\r\n")
	builder.WriteString("Content-Transfer-Encoding: 8bit\r\n\r\n")
	builder.WriteString(licenseEmailHTMLBody(license, copy))
	builder.WriteString("\r\n--")
	builder.WriteString(boundary)
	builder.WriteString("--\r\n")
	return []byte(builder.String()), nil
}

func (m *SMTPMailer) buildRecoveryCodeMessage(to, code, locale string, expiresAt time.Time) ([]byte, error) {
	from := mail.Address{Name: m.cfg.FromName, Address: m.cfg.FromAddress}
	recipient := mail.Address{Address: to}
	copy := recoveryCodeEmailCopy(locale)
	boundary := "recoverease-recovery-code-boundary"

	headers := map[string]string{
		"From":         from.String(),
		"To":           recipient.String(),
		"Subject":      mime.QEncoding.Encode("UTF-8", copy.Subject),
		"Date":         time.Now().Format(time.RFC1123Z),
		"MIME-Version": "1.0",
		"Content-Type": fmt.Sprintf(`multipart/alternative; boundary="%s"`, boundary),
	}
	if strings.TrimSpace(m.cfg.ReplyTo) != "" {
		headers["Reply-To"] = strings.TrimSpace(m.cfg.ReplyTo)
	}

	var builder strings.Builder
	for _, key := range []string{"From", "To", "Reply-To", "Subject", "Date", "MIME-Version", "Content-Type"} {
		value, ok := headers[key]
		if !ok {
			continue
		}
		builder.WriteString(key)
		builder.WriteString(": ")
		builder.WriteString(value)
		builder.WriteString("\r\n")
	}
	builder.WriteString("\r\n")
	builder.WriteString("--")
	builder.WriteString(boundary)
	builder.WriteString("\r\n")
	builder.WriteString("Content-Type: text/plain; charset=\"UTF-8\"\r\n")
	builder.WriteString("Content-Transfer-Encoding: 8bit\r\n\r\n")
	builder.WriteString(recoveryCodeEmailTextBody(code, expiresAt, copy))
	builder.WriteString("\r\n--")
	builder.WriteString(boundary)
	builder.WriteString("\r\n")
	builder.WriteString("Content-Type: text/html; charset=\"UTF-8\"\r\n")
	builder.WriteString("Content-Transfer-Encoding: 8bit\r\n\r\n")
	builder.WriteString(recoveryCodeEmailHTMLBody(code, expiresAt, copy))
	builder.WriteString("\r\n--")
	builder.WriteString(boundary)
	builder.WriteString("--\r\n")
	return []byte(builder.String()), nil
}

type emailCopy struct {
	Subject        string
	Preheader      string
	Title          string
	Subtitle       string
	LicenseKey     string
	Order          string
	Plan           string
	Issued         string
	Expires        string
	Lifetime       string
	PlainGreeting  string
	PlainThanks    string
	ResultPageNote string
	Team           string
}

type recoveryCodeCopy struct {
	Subject       string
	Preheader     string
	Title         string
	Subtitle      string
	CodeLabel     string
	ExpiresLabel  string
	IgnoreNotice  string
	PlainGreeting string
	PlainIntro    string
	PlainExpires  string
	Team          string
}

func recoveryCodeEmailCopy(locale string) recoveryCodeCopy {
	switch store.NormalizeLocale(locale) {
	case "zh":
		return recoveryCodeCopy{
			Subject:       "您的 RecoverEase 找回验证码",
			Preheader:     "使用此验证码查看您的历史激活码。",
			Title:         "验证您的购买邮箱",
			Subtitle:      "请输入以下验证码以查看该邮箱下的历史激活码。",
			CodeLabel:     "验证码",
			ExpiresLabel:  "有效期至",
			IgnoreNotice:  "如果不是您发起的请求，可以忽略这封邮件。",
			PlainGreeting: "您好，",
			PlainIntro:    "请使用以下验证码完成 RecoverEase 激活码找回：",
			PlainExpires:  "验证码有效期至",
			Team:          "RecoverEase 团队",
		}
	case "ja":
		return recoveryCodeCopy{
			Subject:       "RecoverEase 復元確認コード",
			Preheader:     "このコードで過去のライセンスキーを表示できます。",
			Title:         "購入メールを確認してください",
			Subtitle:      "以下の確認コードを入力すると、このメールのライセンス履歴を表示できます。",
			CodeLabel:     "確認コード",
			ExpiresLabel:  "有効期限",
			IgnoreNotice:  "このリクエストに心当たりがない場合は、このメールを無視してください。",
			PlainGreeting: "こんにちは、",
			PlainIntro:    "RecoverEase のライセンス復元には次の確認コードを使用してください。",
			PlainExpires:  "コードの有効期限",
			Team:          "RecoverEase チーム",
		}
	case "ko":
		return recoveryCodeCopy{
			Subject:       "RecoverEase 복구 인증 코드",
			Preheader:     "이 코드로 이전 라이선스 키를 확인할 수 있습니다.",
			Title:         "구매 이메일을 확인하세요",
			Subtitle:      "아래 인증 코드를 입력하면 이 이메일의 라이선스 기록을 볼 수 있습니다.",
			CodeLabel:     "인증 코드",
			ExpiresLabel:  "만료 시간",
			IgnoreNotice:  "요청한 적이 없다면 이 이메일을 무시하세요.",
			PlainGreeting: "안녕하세요,",
			PlainIntro:    "RecoverEase 라이선스 복구에 다음 인증 코드를 사용하세요.",
			PlainExpires:  "코드 만료 시간",
			Team:          "RecoverEase 팀",
		}
	case "de":
		return recoveryCodeCopy{
			Subject:       "Ihr RecoverEase-Bestaetigungscode",
			Preheader:     "Mit diesem Code sehen Sie Ihre bisherigen Lizenzschluessel.",
			Title:         "Kauf-E-Mail bestaetigen",
			Subtitle:      "Geben Sie diesen Code ein, um die Lizenzhistorie dieser E-Mail anzuzeigen.",
			CodeLabel:     "Bestaetigungscode",
			ExpiresLabel:  "Gueltig bis",
			IgnoreNotice:  "Wenn Sie diese Anfrage nicht gestellt haben, koennen Sie diese E-Mail ignorieren.",
			PlainGreeting: "Hallo,",
			PlainIntro:    "Verwenden Sie diesen Code, um Ihre RecoverEase-Lizenzschluessel wiederherzustellen:",
			PlainExpires:  "Code gueltig bis",
			Team:          "RecoverEase Team",
		}
	case "fr":
		return recoveryCodeCopy{
			Subject:       "Votre code de verification RecoverEase",
			Preheader:     "Ce code permet d'afficher vos anciennes cles de licence.",
			Title:         "Verifiez votre e-mail d'achat",
			Subtitle:      "Saisissez ce code pour afficher l'historique des licences de cet e-mail.",
			CodeLabel:     "Code de verification",
			ExpiresLabel:  "Valable jusqu'a",
			IgnoreNotice:  "Si vous n'etes pas a l'origine de cette demande, ignorez cet e-mail.",
			PlainGreeting: "Bonjour,",
			PlainIntro:    "Utilisez ce code pour recuperer vos cles RecoverEase :",
			PlainExpires:  "Code valable jusqu'a",
			Team:          "Equipe RecoverEase",
		}
	case "es":
		return recoveryCodeCopy{
			Subject:       "Tu codigo de verificacion de RecoverEase",
			Preheader:     "Usa este codigo para ver tus claves de licencia anteriores.",
			Title:         "Verifica tu correo de compra",
			Subtitle:      "Introduce este codigo para ver el historial de licencias de este correo.",
			CodeLabel:     "Codigo de verificacion",
			ExpiresLabel:  "Valido hasta",
			IgnoreNotice:  "Si no solicitaste esto, puedes ignorar este correo.",
			PlainGreeting: "Hola,",
			PlainIntro:    "Usa este codigo para recuperar tus claves de RecoverEase:",
			PlainExpires:  "Codigo valido hasta",
			Team:          "Equipo de RecoverEase",
		}
	case "pt":
		return recoveryCodeCopy{
			Subject:       "Seu codigo de verificacao RecoverEase",
			Preheader:     "Use este codigo para ver suas chaves de licenca anteriores.",
			Title:         "Verifique o e-mail da compra",
			Subtitle:      "Digite este codigo para ver o historico de licencas deste e-mail.",
			CodeLabel:     "Codigo de verificacao",
			ExpiresLabel:  "Valido ate",
			IgnoreNotice:  "Se voce nao fez esta solicitacao, ignore este e-mail.",
			PlainGreeting: "Ola,",
			PlainIntro:    "Use este codigo para recuperar suas chaves RecoverEase:",
			PlainExpires:  "Codigo valido ate",
			Team:          "Equipe RecoverEase",
		}
	case "ru":
		return recoveryCodeCopy{
			Subject:       "Код подтверждения RecoverEase",
			Preheader:     "Используйте этот код, чтобы посмотреть прошлые лицензионные ключи.",
			Title:         "Подтвердите e-mail покупки",
			Subtitle:      "Введите этот код, чтобы увидеть историю лицензий для этого e-mail.",
			CodeLabel:     "Код подтверждения",
			ExpiresLabel:  "Действует до",
			IgnoreNotice:  "Если вы не запрашивали код, просто проигнорируйте это письмо.",
			PlainGreeting: "Здравствуйте,",
			PlainIntro:    "Используйте этот код для восстановления ключей RecoverEase:",
			PlainExpires:  "Код действует до",
			Team:          "Команда RecoverEase",
		}
	default:
		return recoveryCodeCopy{
			Subject:       "Your RecoverEase verification code",
			Preheader:     "Use this code to view your historical license keys.",
			Title:         "Verify your purchase email",
			Subtitle:      "Enter this code to view the license history for this email address.",
			CodeLabel:     "Verification code",
			ExpiresLabel:  "Expires",
			IgnoreNotice:  "If you did not request this code, you can ignore this email.",
			PlainGreeting: "Hi,",
			PlainIntro:    "Use this code to recover your RecoverEase license keys:",
			PlainExpires:  "Code expires",
			Team:          "RecoverEase Team",
		}
	}
}

func licenseEmailCopy(locale string) emailCopy {
	switch store.NormalizeLocale(locale) {
	case "zh":
		return emailCopy{
			Subject:        "您的 RecoverEase 激活码",
			Preheader:      "您的 RecoverEase Pro 激活码已准备好。",
			Title:          "您的激活码已准备好",
			Subtitle:       "感谢购买 RecoverEase Pro。请妥善保存这封邮件。",
			LicenseKey:     "激活码",
			Order:          "订单号",
			Plan:           "授权方案",
			Issued:         "发放时间",
			Expires:        "有效期至",
			Lifetime:       "永久有效",
			PlainGreeting:  "您好，",
			PlainThanks:    "感谢您购买 RecoverEase Pro。",
			ResultPageNote: "您也可以在支付结果页查看该激活码，之后也能用购买邮箱找回。",
			Team:           "RecoverEase 团队",
		}
	case "ja":
		return emailCopy{
			Subject:        "RecoverEase ライセンスキー",
			Preheader:      "RecoverEase Pro のライセンスキーが発行されました。",
			Title:          "ライセンスキーが発行されました",
			Subtitle:       "RecoverEase Pro のご購入ありがとうございます。このメールを大切に保管してください。",
			LicenseKey:     "ライセンスキー",
			Order:          "注文",
			Plan:           "プラン",
			Issued:         "発行日",
			Expires:        "有効期限",
			Lifetime:       "無期限",
			PlainGreeting:  "こんにちは。",
			PlainThanks:    "RecoverEase Pro のご購入ありがとうございます。",
			ResultPageNote: "支払い結果ページでもライセンスキーを確認できます。後から購入時のメールで再取得することもできます。",
			Team:           "RecoverEase チーム",
		}
	case "ko":
		return emailCopy{
			Subject:        "RecoverEase 라이선스 키",
			Preheader:      "RecoverEase Pro 라이선스 키가 준비되었습니다.",
			Title:          "라이선스 키가 준비되었습니다",
			Subtitle:       "RecoverEase Pro를 구매해 주셔서 감사합니다. 이 이메일을 보관해 주세요.",
			LicenseKey:     "라이선스 키",
			Order:          "주문",
			Plan:           "플랜",
			Issued:         "발급일",
			Expires:        "만료일",
			Lifetime:       "평생",
			PlainGreeting:  "안녕하세요.",
			PlainThanks:    "RecoverEase Pro를 구매해 주셔서 감사합니다.",
			ResultPageNote: "결제 결과 페이지에서도 라이선스 키를 확인할 수 있으며, 나중에 구매 이메일로 다시 찾을 수 있습니다.",
			Team:           "RecoverEase 팀",
		}
	case "de":
		return emailCopy{
			Subject:        "Ihr RecoverEase-Lizenzschlüssel",
			Preheader:      "Ihr RecoverEase Pro Lizenzschlüssel ist bereit.",
			Title:          "Ihr Lizenzschlüssel ist bereit",
			Subtitle:       "Vielen Dank für den Kauf von RecoverEase Pro. Bewahren Sie diese E-Mail gut auf.",
			LicenseKey:     "Lizenzschlüssel",
			Order:          "Bestellung",
			Plan:           "Plan",
			Issued:         "Ausgestellt",
			Expires:        "Gültig bis",
			Lifetime:       "Lebenslang",
			PlainGreeting:  "Hallo,",
			PlainThanks:    "Vielen Dank für den Kauf von RecoverEase Pro.",
			ResultPageNote: "Sie können diesen Lizenzschlüssel auch auf der Zahlungsseite ansehen oder später mit Ihrer Kauf-E-Mail wiederherstellen.",
			Team:           "RecoverEase Team",
		}
	case "fr":
		return emailCopy{
			Subject:        "Votre clé de licence RecoverEase",
			Preheader:      "Votre clé RecoverEase Pro est prête.",
			Title:          "Votre clé de licence est prête",
			Subtitle:       "Merci pour votre achat de RecoverEase Pro. Conservez cet e-mail.",
			LicenseKey:     "Clé de licence",
			Order:          "Commande",
			Plan:           "Formule",
			Issued:         "Émise le",
			Expires:        "Expire le",
			Lifetime:       "À vie",
			PlainGreeting:  "Bonjour,",
			PlainThanks:    "Merci pour votre achat de RecoverEase Pro.",
			ResultPageNote: "Vous pouvez aussi consulter cette clé sur la page de résultat du paiement ou la retrouver plus tard avec votre e-mail d'achat.",
			Team:           "Équipe RecoverEase",
		}
	case "es":
		return emailCopy{
			Subject:        "Tu clave de licencia de RecoverEase",
			Preheader:      "Tu clave de RecoverEase Pro está lista.",
			Title:          "Tu clave de licencia está lista",
			Subtitle:       "Gracias por comprar RecoverEase Pro. Guarda este correo para tus registros.",
			LicenseKey:     "Clave de licencia",
			Order:          "Pedido",
			Plan:           "Plan",
			Issued:         "Emitida",
			Expires:        "Caduca",
			Lifetime:       "De por vida",
			PlainGreeting:  "Hola,",
			PlainThanks:    "Gracias por comprar RecoverEase Pro.",
			ResultPageNote: "También puedes ver esta clave en la página de resultado del pago o recuperarla más tarde con tu correo de compra.",
			Team:           "Equipo de RecoverEase",
		}
	case "pt":
		return emailCopy{
			Subject:        "Sua chave de licença RecoverEase",
			Preheader:      "Sua chave do RecoverEase Pro está pronta.",
			Title:          "Sua chave de licença está pronta",
			Subtitle:       "Obrigado por comprar o RecoverEase Pro. Guarde este e-mail.",
			LicenseKey:     "Chave de licença",
			Order:          "Pedido",
			Plan:           "Plano",
			Issued:         "Emitida em",
			Expires:        "Expira em",
			Lifetime:       "Vitalícia",
			PlainGreeting:  "Olá,",
			PlainThanks:    "Obrigado por comprar o RecoverEase Pro.",
			ResultPageNote: "Você também pode ver esta chave na página de resultado do pagamento ou recuperá-la depois com o e-mail da compra.",
			Team:           "Equipe RecoverEase",
		}
	case "ru":
		return emailCopy{
			Subject:        "Ваш лицензионный ключ RecoverEase",
			Preheader:      "Ваш ключ RecoverEase Pro готов.",
			Title:          "Ваш лицензионный ключ готов",
			Subtitle:       "Спасибо за покупку RecoverEase Pro. Сохраните это письмо.",
			LicenseKey:     "Лицензионный ключ",
			Order:          "Заказ",
			Plan:           "План",
			Issued:         "Выдан",
			Expires:        "Действует до",
			Lifetime:       "Бессрочно",
			PlainGreeting:  "Здравствуйте,",
			PlainThanks:    "Спасибо за покупку RecoverEase Pro.",
			ResultPageNote: "Вы также можете посмотреть этот ключ на странице результата оплаты или восстановить его позже по e-mail покупки.",
			Team:           "Команда RecoverEase",
		}
	default:
		return emailCopy{
			Subject:        "Your RecoverEase license key",
			Preheader:      "Your RecoverEase Pro license key is ready.",
			Title:          "Your license key is ready",
			Subtitle:       "Thanks for your purchase. Keep this email for your records.",
			LicenseKey:     "License key",
			Order:          "Order",
			Plan:           "Plan",
			Issued:         "Issued",
			Expires:        "Expires",
			Lifetime:       "Lifetime",
			PlainGreeting:  "Hi,",
			PlainThanks:    "Thank you for purchasing RecoverEase Pro.",
			ResultPageNote: "You can also view this license key on the payment result page or recover it later with your purchase email.",
			Team:           "RecoverEase Team",
		}
	}
}

func licenseEmailTextBody(license *store.License, copy emailCopy) string {
	expires := copy.Lifetime
	if license.ExpiresAt != nil {
		expires = license.ExpiresAt.Format("2006-01-02")
	}

	return fmt.Sprintf(`%s

%s

%s:
%s

%s: %s
%s: %s
%s: %s
%s: %s

%s

%s
`,
		copy.PlainGreeting,
		copy.PlainThanks,
		copy.LicenseKey,
		license.LicenseKey,
		copy.Order,
		license.Order.OrderNo,
		copy.Plan,
		license.LicensePlan.Name,
		copy.Issued,
		license.IssuedAt.Format("2006-01-02 15:04 MST"),
		copy.Expires,
		expires,
		copy.ResultPageNote,
		copy.Team,
	)
}

func licenseEmailHTMLBody(license *store.License, copy emailCopy) string {
	expires := copy.Lifetime
	if license.ExpiresAt != nil {
		expires = license.ExpiresAt.Format("Jan 2, 2006")
	}

	preheader := html.EscapeString(copy.Preheader)
	title := html.EscapeString(copy.Title)
	subtitle := html.EscapeString(copy.Subtitle)
	licenseKeyLabel := html.EscapeString(copy.LicenseKey)
	orderLabel := html.EscapeString(copy.Order)
	planLabel := html.EscapeString(copy.Plan)
	issuedLabel := html.EscapeString(copy.Issued)
	expiresLabel := html.EscapeString(copy.Expires)
	resultPageNote := html.EscapeString(copy.ResultPageNote)
	team := html.EscapeString(copy.Team)
	licenseKey := html.EscapeString(license.LicenseKey)
	orderNo := html.EscapeString(license.Order.OrderNo)
	planName := html.EscapeString(license.LicensePlan.Name)
	issuedAt := html.EscapeString(license.IssuedAt.Format("Jan 2, 2006, 15:04 MST"))
	expires = html.EscapeString(expires)

	return fmt.Sprintf(`<!doctype html>
<html>
  <body style="margin:0;background:#f4f7fb;color:#111827;font-family:Inter,Segoe UI,Arial,sans-serif;">
    <span style="display:none!important;visibility:hidden;opacity:0;height:0;width:0;overflow:hidden;">%s</span>
    <table role="presentation" width="100%%" cellspacing="0" cellpadding="0" style="background:#f4f7fb;padding:32px 16px;">
      <tr>
        <td align="center">
          <table role="presentation" width="100%%" cellspacing="0" cellpadding="0" style="max-width:640px;background:#ffffff;border:1px solid #e6edf5;border-radius:18px;overflow:hidden;box-shadow:0 18px 50px rgba(15,23,42,0.10);">
            <tr>
              <td style="padding:28px 32px;background:#eef6ff;border-bottom:1px solid #dbeafe;">
                <div style="font-size:13px;font-weight:800;letter-spacing:.08em;text-transform:uppercase;color:#156dff;">RecoverEase Pro</div>
                <h1 style="margin:10px 0 0;font-size:26px;line-height:1.25;color:#0d1321;">%s</h1>
                <p style="margin:12px 0 0;font-size:15px;line-height:1.7;color:#526070;">%s</p>
              </td>
            </tr>
            <tr>
              <td style="padding:30px 32px;">
                <div style="margin-bottom:22px;padding:22px;border:1px solid #8bbcff;border-radius:14px;background:#f8fbff;">
                  <div style="margin-bottom:10px;font-size:13px;font-weight:800;color:#526070;text-transform:uppercase;letter-spacing:.06em;">%s</div>
                  <div style="font-family:SFMono-Regular,Consolas,Liberation Mono,monospace;font-size:22px;line-height:1.5;font-weight:800;color:#071b4f;word-break:break-all;">%s</div>
                </div>
                <table role="presentation" width="100%%" cellspacing="0" cellpadding="0" style="border-collapse:collapse;">
                  <tr>
                    <td style="padding:12px 0;border-bottom:1px solid #e6edf5;color:#667085;font-size:14px;">%s</td>
                    <td align="right" style="padding:12px 0;border-bottom:1px solid #e6edf5;color:#0d1321;font-size:14px;font-weight:700;">%s</td>
                  </tr>
                  <tr>
                    <td style="padding:12px 0;border-bottom:1px solid #e6edf5;color:#667085;font-size:14px;">%s</td>
                    <td align="right" style="padding:12px 0;border-bottom:1px solid #e6edf5;color:#0d1321;font-size:14px;font-weight:700;">%s</td>
                  </tr>
                  <tr>
                    <td style="padding:12px 0;border-bottom:1px solid #e6edf5;color:#667085;font-size:14px;">%s</td>
                    <td align="right" style="padding:12px 0;border-bottom:1px solid #e6edf5;color:#0d1321;font-size:14px;font-weight:700;">%s</td>
                  </tr>
                  <tr>
                    <td style="padding:12px 0;color:#667085;font-size:14px;">%s</td>
                    <td align="right" style="padding:12px 0;color:#0d1321;font-size:14px;font-weight:700;">%s</td>
                  </tr>
                </table>
                <p style="margin:24px 0 0;font-size:14px;line-height:1.7;color:#526070;">%s</p>
              </td>
            </tr>
            <tr>
              <td style="padding:18px 32px;background:#f8fafc;border-top:1px solid #e6edf5;color:#667085;font-size:13px;line-height:1.6;">
                %s
              </td>
            </tr>
          </table>
        </td>
      </tr>
    </table>
  </body>
</html>`,
		preheader,
		title,
		subtitle,
		licenseKeyLabel,
		licenseKey,
		orderLabel,
		orderNo,
		planLabel,
		planName,
		issuedLabel,
		issuedAt,
		expiresLabel,
		expires,
		resultPageNote,
		team,
	)
}

func recoveryCodeEmailTextBody(code string, expiresAt time.Time, copy recoveryCodeCopy) string {
	return fmt.Sprintf(`%s

%s

%s:
%s

%s: %s

%s

%s
`,
		copy.PlainGreeting,
		copy.PlainIntro,
		copy.CodeLabel,
		code,
		copy.PlainExpires,
		expiresAt.Format("2006-01-02 15:04 MST"),
		copy.IgnoreNotice,
		copy.Team,
	)
}

func recoveryCodeEmailHTMLBody(code string, expiresAt time.Time, copy recoveryCodeCopy) string {
	preheader := html.EscapeString(copy.Preheader)
	title := html.EscapeString(copy.Title)
	subtitle := html.EscapeString(copy.Subtitle)
	codeLabel := html.EscapeString(copy.CodeLabel)
	expiresLabel := html.EscapeString(copy.ExpiresLabel)
	ignoreNotice := html.EscapeString(copy.IgnoreNotice)
	team := html.EscapeString(copy.Team)
	escapedCode := html.EscapeString(code)
	expires := html.EscapeString(expiresAt.Format("2006-01-02 15:04 MST"))

	return fmt.Sprintf(`<!doctype html>
<html>
  <body style="margin:0;background:#f4f7fb;color:#111827;font-family:Inter,Segoe UI,Arial,sans-serif;">
    <span style="display:none!important;visibility:hidden;opacity:0;height:0;width:0;overflow:hidden;">%s</span>
    <table role="presentation" width="100%%" cellspacing="0" cellpadding="0" style="background:#f4f7fb;padding:32px 16px;">
      <tr>
        <td align="center">
          <table role="presentation" width="100%%" cellspacing="0" cellpadding="0" style="max-width:600px;background:#ffffff;border:1px solid #e6edf5;border-radius:18px;overflow:hidden;box-shadow:0 18px 50px rgba(15,23,42,0.10);">
            <tr>
              <td style="padding:28px 32px;background:#eef6ff;border-bottom:1px solid #dbeafe;">
                <div style="font-size:13px;font-weight:800;letter-spacing:.08em;text-transform:uppercase;color:#156dff;">RecoverEase</div>
                <h1 style="margin:10px 0 0;font-size:26px;line-height:1.25;color:#0d1321;">%s</h1>
                <p style="margin:12px 0 0;font-size:15px;line-height:1.7;color:#526070;">%s</p>
              </td>
            </tr>
            <tr>
              <td style="padding:30px 32px;">
                <div style="margin-bottom:22px;padding:22px;border:1px solid #8bbcff;border-radius:14px;background:#f8fbff;text-align:center;">
                  <div style="margin-bottom:10px;font-size:13px;font-weight:800;color:#526070;text-transform:uppercase;letter-spacing:.06em;">%s</div>
                  <div style="font-family:SFMono-Regular,Consolas,Liberation Mono,monospace;font-size:34px;line-height:1.2;font-weight:900;color:#071b4f;letter-spacing:.14em;">%s</div>
                </div>
                <table role="presentation" width="100%%" cellspacing="0" cellpadding="0" style="border-collapse:collapse;">
                  <tr>
                    <td style="padding:12px 0;border-bottom:1px solid #e6edf5;color:#667085;font-size:14px;">%s</td>
                    <td align="right" style="padding:12px 0;border-bottom:1px solid #e6edf5;color:#0d1321;font-size:14px;font-weight:700;">%s</td>
                  </tr>
                </table>
                <p style="margin:24px 0 0;font-size:14px;line-height:1.7;color:#526070;">%s</p>
              </td>
            </tr>
            <tr>
              <td style="padding:18px 32px;background:#f8fafc;border-top:1px solid #e6edf5;color:#667085;font-size:13px;line-height:1.6;">
                %s
              </td>
            </tr>
          </table>
        </td>
      </tr>
    </table>
  </body>
</html>`,
		preheader,
		title,
		subtitle,
		codeLabel,
		escapedCode,
		expiresLabel,
		expires,
		ignoreNotice,
		team,
	)
}

func (m *SMTPMailer) send(ctx context.Context, recipients []string, message []byte) error {
	port, err := strconv.Atoi(strings.TrimSpace(m.cfg.Port))
	if err != nil {
		return fmt.Errorf("invalid email port: %w", err)
	}

	address := net.JoinHostPort(strings.TrimSpace(m.cfg.Host), strings.TrimSpace(m.cfg.Port))
	auth := smtp.PlainAuth("", strings.TrimSpace(m.cfg.Username), strings.TrimSpace(m.cfg.Password), strings.TrimSpace(m.cfg.Host))
	if port == 465 {
		return m.sendImplicitTLS(ctx, address, auth, recipients, message)
	}

	return smtp.SendMail(address, auth, strings.TrimSpace(m.cfg.FromAddress), recipients, message)
}

func (m *SMTPMailer) sendImplicitTLS(ctx context.Context, address string, auth smtp.Auth, recipients []string, message []byte) error {
	dialer := &net.Dialer{Timeout: 10 * time.Second}
	conn, err := tls.DialWithDialer(dialer, "tcp", address, &tls.Config{
		ServerName: strings.TrimSpace(m.cfg.Host),
		MinVersion: tls.VersionTLS12,
	})
	if err != nil {
		return err
	}
	defer conn.Close()

	if deadline, ok := ctx.Deadline(); ok {
		_ = conn.SetDeadline(deadline)
	}

	client, err := smtp.NewClient(conn, strings.TrimSpace(m.cfg.Host))
	if err != nil {
		return err
	}
	defer client.Close()

	if err := client.Auth(auth); err != nil {
		return err
	}
	if err := client.Mail(strings.TrimSpace(m.cfg.FromAddress)); err != nil {
		return err
	}
	for _, recipient := range recipients {
		if err := client.Rcpt(recipient); err != nil {
			return err
		}
	}

	writer, err := client.Data()
	if err != nil {
		return err
	}
	if _, err := writer.Write(message); err != nil {
		_ = writer.Close()
		return err
	}
	if err := writer.Close(); err != nil {
		return err
	}
	if err := client.Quit(); err != nil && !errors.Is(err, io.EOF) {
		return err
	}
	return nil
}
