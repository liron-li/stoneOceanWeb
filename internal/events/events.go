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

type PaymentPaidListener interface {
	HandlePaymentPaid(context.Context, PaymentPaidEvent) error
}

type Bus struct {
	paymentPaidListeners []PaymentPaidListener
	logger               *log.Logger
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

type LicenseMailer interface {
	SendLicense(context.Context, *store.License) error
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
