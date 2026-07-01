package mailer

import (
	"strings"
	"testing"
	"time"

	"stone-ocean-web/internal/config"
	"stone-ocean-web/internal/store"
)

func TestBuildLicenseMessageIncludesStyledHTML(t *testing.T) {
	mailer := NewSMTPMailer(config.EmailConfig{
		FromName:    "RecoverEase",
		FromAddress: "mailer@example.com",
	})
	license := &store.License{
		LicenseKey: "RE-TEST-KEY",
		Order: store.Order{
			OrderNo: "order_test",
			Email:   "buyer@example.com",
		},
		LicensePlan: store.LicensePlan{
			Name: "RecoverEase Pro Lifetime",
		},
		IssuedAt: time.Date(2026, 7, 1, 14, 21, 0, 0, time.UTC),
	}

	message, err := mailer.buildLicenseMessage("buyer@example.com", license)
	if err != nil {
		t.Fatalf("buildLicenseMessage() error = %v", err)
	}

	body := string(message)
	for _, want := range []string{
		`Content-Type: multipart/alternative`,
		`Content-Type: text/plain; charset="UTF-8"`,
		`Content-Type: text/html; charset="UTF-8"`,
		"Your license key is ready",
		"RE-TEST-KEY",
		"RecoverEase Pro Lifetime",
	} {
		if !strings.Contains(body, want) {
			t.Fatalf("message does not contain %q", want)
		}
	}
}

func TestBuildLicenseMessageUsesOrderLocale(t *testing.T) {
	mailer := NewSMTPMailer(config.EmailConfig{
		FromName:    "RecoverEase",
		FromAddress: "mailer@example.com",
	})
	license := &store.License{
		LicenseKey: "RE-ZH-KEY",
		Order: store.Order{
			OrderNo: "order_zh",
			Email:   "buyer@example.com",
			Locale:  "zh",
		},
		LicensePlan: store.LicensePlan{
			Name: "RecoverEase Pro Lifetime",
		},
		IssuedAt: time.Date(2026, 7, 1, 14, 21, 0, 0, time.UTC),
	}

	message, err := mailer.buildLicenseMessage("buyer@example.com", license)
	if err != nil {
		t.Fatalf("buildLicenseMessage() error = %v", err)
	}

	body := string(message)
	for _, want := range []string{
		"您的激活码已准备好",
		"激活码",
		"订单号",
		"永久有效",
	} {
		if !strings.Contains(body, want) {
			t.Fatalf("message does not contain localized text %q", want)
		}
	}
}

func TestBuildRecoveryCodeMessageUsesLocale(t *testing.T) {
	mailer := NewSMTPMailer(config.EmailConfig{
		FromName:    "RecoverEase",
		FromAddress: "mailer@example.com",
	})
	expiresAt := time.Date(2026, 7, 1, 15, 0, 0, 0, time.UTC)

	message, err := mailer.buildRecoveryCodeMessage("buyer@example.com", "123456", "zh", expiresAt)
	if err != nil {
		t.Fatalf("buildRecoveryCodeMessage() error = %v", err)
	}

	body := string(message)
	for _, want := range []string{
		`Content-Type: multipart/alternative`,
		`Content-Type: text/plain; charset="UTF-8"`,
		`Content-Type: text/html; charset="UTF-8"`,
		"123456",
		"验证您的购买邮箱",
		"验证码",
	} {
		if !strings.Contains(body, want) {
			t.Fatalf("message does not contain %q", want)
		}
	}
}
