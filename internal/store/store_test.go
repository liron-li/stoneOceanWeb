package store

import (
	"context"
	"errors"
	"testing"
	"time"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func TestCheckoutPaymentAndLicenseFlow(t *testing.T) {
	store := newTestStore(t)
	ctx := context.Background()

	result, err := store.CreateCheckoutOrder(ctx, CreateCheckoutOrderInput{
		Email:         "User@Example.COM",
		PlanCode:      DefaultLifetimePlanCode,
		PaymentMethod: PaymentMethodUSDT,
	})
	if err != nil {
		t.Fatalf("CreateCheckoutOrder() error = %v", err)
	}
	if result.Customer.EmailNormalized != "user@example.com" {
		t.Fatalf("EmailNormalized = %q, want user@example.com", result.Customer.EmailNormalized)
	}
	if result.Order.AmountCents != 2900 {
		t.Fatalf("Order.AmountCents = %d, want 2900", result.Order.AmountCents)
	}
	if result.Payment.Status != PaymentStatusPending {
		t.Fatalf("Payment.Status = %q, want %q", result.Payment.Status, PaymentStatusPending)
	}

	paidAt := time.Date(2026, 6, 26, 12, 0, 0, 0, time.UTC)
	license, err := store.MarkPaymentPaid(ctx, result.Payment.PaymentNo, "provider-ref-1", paidAt)
	if err != nil {
		t.Fatalf("MarkPaymentPaid() error = %v", err)
	}
	if license.Status != LicenseStatusActive {
		t.Fatalf("License.Status = %q, want %q", license.Status, LicenseStatusActive)
	}
	if license.ExpiresAt != nil {
		t.Fatalf("lifetime license ExpiresAt = %v, want nil", license.ExpiresAt)
	}
	if license.LicenseKey == "" {
		t.Fatal("LicenseKey is empty")
	}

	sameLicense, err := store.MarkPaymentPaid(ctx, result.Payment.PaymentNo, "provider-ref-1", paidAt)
	if err != nil {
		t.Fatalf("second MarkPaymentPaid() error = %v", err)
	}
	if sameLicense.ID != license.ID {
		t.Fatalf("second MarkPaymentPaid() license ID = %d, want %d", sameLicense.ID, license.ID)
	}

	licenses, err := store.FindLicensesByEmail(ctx, "user@example.com")
	if err != nil {
		t.Fatalf("FindLicensesByEmail() error = %v", err)
	}
	if len(licenses) != 1 {
		t.Fatalf("FindLicensesByEmail() length = %d, want 1", len(licenses))
	}
}

func TestMonthlyLicenseExpiresAfterDuration(t *testing.T) {
	store := newTestStore(t)
	ctx := context.Background()

	result, err := store.CreateCheckoutOrder(ctx, CreateCheckoutOrderInput{
		Email:         "monthly@example.com",
		PlanCode:      DefaultMonthlyPlanCode,
		PaymentMethod: PaymentMethodCard,
	})
	if err != nil {
		t.Fatalf("CreateCheckoutOrder() error = %v", err)
	}
	if result.Order.AmountCents != 900 {
		t.Fatalf("Order.AmountCents = %d, want 900", result.Order.AmountCents)
	}

	paidAt := time.Date(2026, 6, 26, 12, 0, 0, 0, time.UTC)
	license, err := store.MarkPaymentPaid(ctx, result.Payment.PaymentNo, "provider-ref-2", paidAt)
	if err != nil {
		t.Fatalf("MarkPaymentPaid() error = %v", err)
	}
	if license.ExpiresAt == nil {
		t.Fatal("monthly license ExpiresAt is nil")
	}
	want := paidAt.AddDate(0, 0, 30)
	if !license.ExpiresAt.Equal(want) {
		t.Fatalf("ExpiresAt = %s, want %s", license.ExpiresAt, want)
	}
}

func TestFindPaymentResultReturnsLicenseAfterPayment(t *testing.T) {
	store := newTestStore(t)
	ctx := context.Background()

	result, err := store.CreateCheckoutOrder(ctx, CreateCheckoutOrderInput{
		Email:         "result@example.com",
		PlanCode:      DefaultLifetimePlanCode,
		PaymentMethod: PaymentMethodUSDT,
		Locale:        "zh-CN",
	})
	if err != nil {
		t.Fatalf("CreateCheckoutOrder() error = %v", err)
	}
	if result.Order.Locale != "zh" {
		t.Fatalf("Order.Locale = %q, want zh", result.Order.Locale)
	}

	pending, err := store.FindPaymentResult(ctx, result.Payment.PaymentNo)
	if err != nil {
		t.Fatalf("pending FindPaymentResult() error = %v", err)
	}
	if pending.Payment.Status != PaymentStatusPending {
		t.Fatalf("pending Payment.Status = %q, want %q", pending.Payment.Status, PaymentStatusPending)
	}
	if pending.License != nil {
		t.Fatal("pending result has a license, want nil")
	}

	paidAt := time.Date(2026, 6, 26, 12, 0, 0, 0, time.UTC)
	license, err := store.MarkPaymentPaid(ctx, result.Payment.PaymentNo, "provider-ref-result", paidAt)
	if err != nil {
		t.Fatalf("MarkPaymentPaid() error = %v", err)
	}

	paid, err := store.FindPaymentResult(ctx, result.Payment.PaymentNo)
	if err != nil {
		t.Fatalf("paid FindPaymentResult() error = %v", err)
	}
	if paid.Payment.Status != PaymentStatusPaid {
		t.Fatalf("paid Payment.Status = %q, want %q", paid.Payment.Status, PaymentStatusPaid)
	}
	if paid.License == nil {
		t.Fatal("paid result License is nil")
	}
	if paid.License.LicenseKey != license.LicenseKey {
		t.Fatalf("LicenseKey = %q, want %q", paid.License.LicenseKey, license.LicenseKey)
	}
}

func TestRecoveryTokenLifecycle(t *testing.T) {
	store := newTestStore(t)
	ctx := context.Background()
	now := time.Now()

	token, err := store.CreateRecoveryToken(ctx, "Buyer@Example.com", "hash-1", now.Add(time.Hour))
	if err != nil {
		t.Fatalf("CreateRecoveryToken() error = %v", err)
	}
	if token.EmailNormalized != "buyer@example.com" {
		t.Fatalf("EmailNormalized = %q, want buyer@example.com", token.EmailNormalized)
	}

	consumed, err := store.ConsumeRecoveryToken(ctx, "hash-1", now)
	if err != nil {
		t.Fatalf("ConsumeRecoveryToken() error = %v", err)
	}
	if consumed.UsedAt == nil {
		t.Fatal("UsedAt is nil after consuming token")
	}

	_, err = store.ConsumeRecoveryToken(ctx, "hash-1", now)
	if !errors.Is(err, ErrTokenUsed) {
		t.Fatalf("second ConsumeRecoveryToken() error = %v, want ErrTokenUsed", err)
	}

	_, err = store.CreateRecoveryToken(ctx, "buyer@example.com", "hash-2", time.Now().Add(-time.Minute))
	if !errors.Is(err, ErrInvalidInput) {
		t.Fatalf("expired CreateRecoveryToken() error = %v, want ErrInvalidInput", err)
	}
}

func TestLicenseActivationLifecycle(t *testing.T) {
	store := newTestStore(t)
	ctx := context.Background()
	checkout, err := store.CreateCheckoutOrder(ctx, CreateCheckoutOrderInput{
		Email:         "activation@example.com",
		PlanCode:      DefaultLifetimePlanCode,
		PaymentMethod: PaymentMethodCard,
	})
	if err != nil {
		t.Fatalf("CreateCheckoutOrder() error = %v", err)
	}
	license, err := store.MarkPaymentPaid(ctx, checkout.Payment.PaymentNo, "provider-activation", time.Now())
	if err != nil {
		t.Fatalf("MarkPaymentPaid() error = %v", err)
	}

	first, err := store.ActivateLicense(ctx, LicenseActivationInput{
		LicenseKey:          license.LicenseKey,
		DeviceIDHash:        "device-hash-1",
		DeviceName:          "Workstation",
		Platform:            "windows",
		AppVersion:          "1.0.0",
		ActivationTokenHash: "token-hash-1",
	})
	if err != nil {
		t.Fatalf("ActivateLicense() error = %v", err)
	}
	if first.Activation.Status != ActivationStatusActive {
		t.Fatalf("Activation.Status = %q, want %q", first.Activation.Status, ActivationStatusActive)
	}

	sameDevice, err := store.ActivateLicense(ctx, LicenseActivationInput{
		LicenseKey:          license.LicenseKey,
		DeviceIDHash:        "device-hash-1",
		DeviceName:          "Workstation",
		Platform:            "windows",
		AppVersion:          "1.0.1",
		ActivationTokenHash: "token-hash-1b",
	})
	if err != nil {
		t.Fatalf("same-device ActivateLicense() error = %v", err)
	}
	if sameDevice.Activation.ID != first.Activation.ID {
		t.Fatalf("same-device activation ID = %d, want %d", sameDevice.Activation.ID, first.Activation.ID)
	}

	second, err := store.ActivateLicense(ctx, LicenseActivationInput{
		LicenseKey:          license.LicenseKey,
		DeviceIDHash:        "device-hash-2",
		DeviceName:          "Laptop",
		Platform:            "windows",
		AppVersion:          "1.0.0",
		ActivationTokenHash: "token-hash-2",
	})
	if err != nil {
		t.Fatalf("second ActivateLicense() error = %v", err)
	}
	if second.Activation.ID == first.Activation.ID {
		t.Fatal("second activation reused the first activation ID")
	}

	_, err = store.ActivateLicense(ctx, LicenseActivationInput{
		LicenseKey:          license.LicenseKey,
		DeviceIDHash:        "device-hash-3",
		DeviceName:          "Third PC",
		Platform:            "windows",
		AppVersion:          "1.0.0",
		ActivationTokenHash: "token-hash-3",
	})
	if !errors.Is(err, ErrActivationLimit) {
		t.Fatalf("third ActivateLicense() error = %v, want ErrActivationLimit", err)
	}

	validated, err := store.ValidateActivation(ctx, LicenseValidationInput{
		ActivationTokenHash: "token-hash-1b",
		DeviceIDHash:        "device-hash-1",
		AppVersion:          "1.0.2",
	})
	if err != nil {
		t.Fatalf("ValidateActivation() error = %v", err)
	}
	if validated.Activation.AppVersion != "1.0.2" {
		t.Fatalf("AppVersion = %q, want 1.0.2", validated.Activation.AppVersion)
	}

	deactivated, err := store.DeactivateActivation(ctx, LicenseValidationInput{
		ActivationTokenHash: "token-hash-1b",
		DeviceIDHash:        "device-hash-1",
	})
	if err != nil {
		t.Fatalf("DeactivateActivation() error = %v", err)
	}
	if deactivated.Activation.Status != ActivationStatusDeactivated {
		t.Fatalf("deactivated Status = %q, want %q", deactivated.Activation.Status, ActivationStatusDeactivated)
	}

	third, err := store.ActivateLicense(ctx, LicenseActivationInput{
		LicenseKey:          license.LicenseKey,
		DeviceIDHash:        "device-hash-3",
		DeviceName:          "Third PC",
		Platform:            "windows",
		AppVersion:          "1.0.0",
		ActivationTokenHash: "token-hash-3",
	})
	if err != nil {
		t.Fatalf("third ActivateLicense() after deactivation error = %v", err)
	}
	if third.Activation.DeviceIDHash != "device-hash-3" {
		t.Fatalf("third DeviceIDHash = %q, want device-hash-3", third.Activation.DeviceIDHash)
	}
}

func TestExpiredLicenseCannotActivate(t *testing.T) {
	store := newTestStore(t)
	ctx := context.Background()
	checkout, err := store.CreateCheckoutOrder(ctx, CreateCheckoutOrderInput{
		Email:         "expired@example.com",
		PlanCode:      DefaultMonthlyPlanCode,
		PaymentMethod: PaymentMethodCard,
	})
	if err != nil {
		t.Fatalf("CreateCheckoutOrder() error = %v", err)
	}
	paidAt := time.Date(2026, 1, 1, 0, 0, 0, 0, time.UTC)
	license, err := store.MarkPaymentPaid(ctx, checkout.Payment.PaymentNo, "provider-expired", paidAt)
	if err != nil {
		t.Fatalf("MarkPaymentPaid() error = %v", err)
	}

	_, err = store.ActivateLicense(ctx, LicenseActivationInput{
		LicenseKey:          license.LicenseKey,
		DeviceIDHash:        "expired-device",
		ActivationTokenHash: "expired-token",
		Now:                 paidAt.AddDate(0, 0, 31),
	})
	if !errors.Is(err, ErrLicenseExpired) {
		t.Fatalf("expired ActivateLicense() error = %v, want ErrLicenseExpired", err)
	}
}

func newTestStore(t *testing.T) *Store {
	t.Helper()

	db, err := gorm.Open(sqlite.Open("file::memory:?cache=shared"), &gorm.Config{})
	if err != nil {
		t.Fatalf("gorm.Open() error = %v", err)
	}
	if err := AutoMigrate(db); err != nil {
		t.Fatalf("AutoMigrate() error = %v", err)
	}
	if err := SeedDefaultPlans(context.Background(), db); err != nil {
		t.Fatalf("SeedDefaultPlans() error = %v", err)
	}

	return New(db)
}
