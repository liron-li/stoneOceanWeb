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
