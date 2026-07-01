package events

import (
	"context"
	"log"
	"time"

	"stone-ocean-web/internal/store"
)

type PaymentPaidEvent struct {
	License *store.License
}

type LicenseRecoveryCodeEvent struct {
	Email     string
	Code      string
	Locale    string
	ExpiresAt time.Time
}

type PaymentPaidListener interface {
	HandlePaymentPaid(context.Context, PaymentPaidEvent) error
}

type LicenseRecoveryCodeListener interface {
	HandleLicenseRecoveryCode(context.Context, LicenseRecoveryCodeEvent) error
}

type Bus struct {
	paymentPaidListeners         []PaymentPaidListener
	licenseRecoveryCodeListeners []LicenseRecoveryCodeListener
	logger                       *log.Logger
}

func NewBus(logger *log.Logger) *Bus {
	if logger == nil {
		logger = log.Default()
	}
	return &Bus{logger: logger}
}

func (b *Bus) AddPaymentPaidListener(listener PaymentPaidListener) {
	if listener == nil {
		return
	}
	b.paymentPaidListeners = append(b.paymentPaidListeners, listener)
}

func (b *Bus) AddLicenseRecoveryCodeListener(listener LicenseRecoveryCodeListener) {
	if listener == nil {
		return
	}
	b.licenseRecoveryCodeListeners = append(b.licenseRecoveryCodeListeners, listener)
}

func (b *Bus) PublishPaymentPaid(event PaymentPaidEvent) {
	if b == nil || len(b.paymentPaidListeners) == 0 {
		return
	}

	for _, listener := range b.paymentPaidListeners {
		listener := listener
		go func() {
			ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
			defer cancel()

			if err := listener.HandlePaymentPaid(ctx, event); err != nil {
				b.logger.Printf("payment paid listener failed: %v", err)
			}
		}()
	}
}

func (b *Bus) PublishLicenseRecoveryCode(event LicenseRecoveryCodeEvent) {
	if b == nil || len(b.licenseRecoveryCodeListeners) == 0 {
		return
	}

	for _, listener := range b.licenseRecoveryCodeListeners {
		listener := listener
		go func() {
			ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
			defer cancel()

			if err := listener.HandleLicenseRecoveryCode(ctx, event); err != nil {
				b.logger.Printf("license recovery code listener failed: %v", err)
			}
		}()
	}
}

type LicenseMailer interface {
	SendLicense(context.Context, *store.License) error
}

type RecoveryCodeMailer interface {
	SendRecoveryCode(context.Context, string, string, string, time.Time) error
}

type LicenseEmailListener struct {
	mailer LicenseMailer
}

func NewLicenseEmailListener(mailer LicenseMailer) *LicenseEmailListener {
	return &LicenseEmailListener{mailer: mailer}
}

func (l *LicenseEmailListener) HandlePaymentPaid(ctx context.Context, event PaymentPaidEvent) error {
	if l == nil || l.mailer == nil || event.License == nil {
		return nil
	}
	return l.mailer.SendLicense(ctx, event.License)
}

type LicenseRecoveryCodeEmailListener struct {
	mailer RecoveryCodeMailer
}

func NewLicenseRecoveryCodeEmailListener(mailer RecoveryCodeMailer) *LicenseRecoveryCodeEmailListener {
	return &LicenseRecoveryCodeEmailListener{mailer: mailer}
}

func (l *LicenseRecoveryCodeEmailListener) HandleLicenseRecoveryCode(ctx context.Context, event LicenseRecoveryCodeEvent) error {
	if l == nil || l.mailer == nil || event.Email == "" || event.Code == "" {
		return nil
	}
	return l.mailer.SendRecoveryCode(ctx, event.Email, event.Code, event.Locale, event.ExpiresAt)
}
