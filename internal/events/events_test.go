package events

import (
	"context"
	"testing"
	"time"

	"stone-ocean-web/internal/store"
)

type recordingMailer struct {
	called chan *store.License
}

func (m recordingMailer) SendLicense(ctx context.Context, license *store.License) error {
	select {
	case m.called <- license:
	case <-ctx.Done():
		return ctx.Err()
	}
	return nil
}

type recordingRecoveryCodeMailer struct {
	called chan LicenseRecoveryCodeEvent
}

func (m recordingRecoveryCodeMailer) SendRecoveryCode(ctx context.Context, email, code, locale string, expiresAt time.Time) error {
	select {
	case m.called <- LicenseRecoveryCodeEvent{Email: email, Code: code, Locale: locale, ExpiresAt: expiresAt}:
	case <-ctx.Done():
		return ctx.Err()
	}
	return nil
}

func TestPaymentPaidEventSendsLicenseEmail(t *testing.T) {
	called := make(chan *store.License, 1)
	bus := NewBus(nil)
	bus.AddPaymentPaidListener(NewLicenseEmailListener(recordingMailer{called: called}))

	license := &store.License{LicenseKey: "RE-TEST"}
	bus.PublishPaymentPaid(PaymentPaidEvent{License: license})

	select {
	case got := <-called:
		if got != license {
			t.Fatal("listener received a different license pointer")
		}
	case <-time.After(time.Second):
		t.Fatal("listener was not called")
	}
}

func TestLicenseRecoveryCodeEventSendsEmail(t *testing.T) {
	called := make(chan LicenseRecoveryCodeEvent, 1)
	bus := NewBus(nil)
	bus.AddLicenseRecoveryCodeListener(NewLicenseRecoveryCodeEmailListener(recordingRecoveryCodeMailer{called: called}))

	expiresAt := time.Now().Add(10 * time.Minute)
	bus.PublishLicenseRecoveryCode(LicenseRecoveryCodeEvent{
		Email:     "buyer@example.com",
		Code:      "123456",
		Locale:    "zh",
		ExpiresAt: expiresAt,
	})

	select {
	case got := <-called:
		if got.Email != "buyer@example.com" || got.Code != "123456" || got.Locale != "zh" || !got.ExpiresAt.Equal(expiresAt) {
			t.Fatalf("listener received %#v", got)
		}
	case <-time.After(time.Second):
		t.Fatal("listener was not called")
	}
}
