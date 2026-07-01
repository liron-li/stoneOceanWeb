package store

import (
	"context"
	"crypto/rand"
	"encoding/hex"
	"errors"
	"fmt"
	"net/mail"
	"strings"
	"time"

	"gorm.io/gorm"
)

const (
	DefaultMonthlyPlanCode  = "recoverease-pro-monthly"
	DefaultLifetimePlanCode = "recoverease-pro-lifetime"
)

var (
	ErrInvalidInput  = errors.New("invalid input")
	ErrNotFound      = errors.New("not found")
	ErrTokenExpired  = errors.New("recovery token expired")
	ErrTokenUsed     = errors.New("recovery token already used")
	ErrPaymentMethod = errors.New("unsupported payment method")
)

type Store struct {
	db *gorm.DB
}

type CreateCheckoutOrderInput struct {
	Email         string
	PlanCode      string
	PaymentMethod PaymentMethod
}

type CreateCheckoutOrderResult struct {
	Customer Customer
	Order    Order
	Payment  Payment
}

type PaymentResult struct {
	Payment Payment
	Order   Order
	License *License
}

func New(db *gorm.DB) *Store {
	return &Store{db: db}
}

func (s *Store) DB() *gorm.DB {
	return s.db
}

func AutoMigrate(db *gorm.DB) error {
	return db.AutoMigrate(
		&Customer{},
		&LicensePlan{},
		&Order{},
		&Payment{},
		&License{},
		&RecoveryToken{},
	)
}

func SeedDefaultPlans(ctx context.Context, db *gorm.DB) error {
	monthlyDays := 30
	plans := []LicensePlan{
		{
			Code:         DefaultMonthlyPlanCode,
			Name:         "RecoverEase Pro Monthly",
			Kind:         LicenseKindMonthly,
			DurationDays: &monthlyDays,
			PriceCents:   900,
			Currency:     "USD",
			IsActive:     true,
		},
		{
			Code:       DefaultLifetimePlanCode,
			Name:       "RecoverEase Pro Lifetime",
			Kind:       LicenseKindLifetime,
			PriceCents: 2900,
			Currency:   "USD",
			IsActive:   true,
		},
	}

	for _, plan := range plans {
		var existing LicensePlan
		err := db.WithContext(ctx).Where("code = ?", plan.Code).First(&existing).Error
		if errors.Is(err, gorm.ErrRecordNotFound) {
			if err := db.WithContext(ctx).Create(&plan).Error; err != nil {
				return err
			}
			continue
		}
		if err != nil {
			return err
		}

		plan.ID = existing.ID
		if err := db.WithContext(ctx).Model(&existing).Updates(plan).Error; err != nil {
			return err
		}
	}

	return nil
}

func SeedDemoData(ctx context.Context, s *Store) error {
	samples := []CreateCheckoutOrderInput{
		{
			Email:         "demo-lifetime@recoverease.test",
			PlanCode:      DefaultLifetimePlanCode,
			PaymentMethod: PaymentMethodUSDT,
		},
		{
			Email:         "demo-monthly@recoverease.test",
			PlanCode:      DefaultMonthlyPlanCode,
			PaymentMethod: PaymentMethodCard,
		},
	}

	for index, sample := range samples {
		licenses, err := s.FindLicensesByEmail(ctx, sample.Email)
		if err != nil {
			return err
		}
		if len(licenses) > 0 {
			continue
		}

		result, err := s.CreateCheckoutOrder(ctx, sample)
		if err != nil {
			return err
		}
		paidAt := time.Now().Add(-time.Duration(index+1) * 24 * time.Hour)
		if _, err := s.MarkPaymentPaid(ctx, result.Payment.PaymentNo, "demo-seed", paidAt); err != nil {
			return err
		}
	}

	return nil
}

func NormalizeEmail(email string) (string, error) {
	email = strings.TrimSpace(email)
	if email == "" {
		return "", fmt.Errorf("%w: email is required", ErrInvalidInput)
	}

	address, err := mail.ParseAddress(email)
	if err != nil || !strings.EqualFold(address.Address, email) {
		return "", fmt.Errorf("%w: invalid email", ErrInvalidInput)
	}

	return strings.ToLower(address.Address), nil
}

func (s *Store) CreateCheckoutOrder(ctx context.Context, input CreateCheckoutOrderInput) (*CreateCheckoutOrderResult, error) {
	normalizedEmail, err := NormalizeEmail(input.Email)
	if err != nil {
		return nil, err
	}
	if input.PlanCode == "" {
		return nil, fmt.Errorf("%w: plan code is required", ErrInvalidInput)
	}
	if input.PaymentMethod != PaymentMethodUSDT && input.PaymentMethod != PaymentMethodCard {
		return nil, ErrPaymentMethod
	}

	result := &CreateCheckoutOrderResult{}
	err = s.db.WithContext(ctx).Transaction(func(tx *gorm.DB) error {
		var plan LicensePlan
		if err := tx.Where("code = ? AND is_active = ?", input.PlanCode, true).First(&plan).Error; err != nil {
			if errors.Is(err, gorm.ErrRecordNotFound) {
				return ErrNotFound
			}
			return err
		}

		now := time.Now()
		customer, err := findOrCreateCustomer(ctx, tx, strings.TrimSpace(input.Email), normalizedEmail, now)
		if err != nil {
			return err
		}

		order := Order{
			OrderNo:       newPublicID("order"),
			CustomerID:    customer.ID,
			LicensePlanID: plan.ID,
			Email:         customer.Email,
			LicenseKind:   plan.Kind,
			AmountCents:   plan.PriceCents,
			Currency:      plan.Currency,
			Status:        OrderStatusPending,
			PaymentMethod: input.PaymentMethod,
		}
		if err := tx.Create(&order).Error; err != nil {
			return err
		}

		payment := Payment{
			PaymentNo:   newPublicID("pay"),
			OrderID:     order.ID,
			Method:      input.PaymentMethod,
			Status:      PaymentStatusPending,
			AmountCents: order.AmountCents,
			Currency:    order.Currency,
		}
		if err := tx.Create(&payment).Error; err != nil {
			return err
		}

		order.Customer = customer
		order.LicensePlan = plan
		payment.Order = order
		result.Customer = customer
		result.Order = order
		result.Payment = payment
		return nil
	})
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (s *Store) MarkPaymentPaid(ctx context.Context, paymentNo string, providerRef string, paidAt time.Time) (*License, error) {
	if strings.TrimSpace(paymentNo) == "" {
		return nil, fmt.Errorf("%w: payment number is required", ErrInvalidInput)
	}
	if paidAt.IsZero() {
		paidAt = time.Now()
	}

	var license License
	err := s.db.WithContext(ctx).Transaction(func(tx *gorm.DB) error {
		var payment Payment
		err := tx.
			Preload("Order.Customer").
			Preload("Order.LicensePlan").
			Where("payment_no = ?", paymentNo).
			First(&payment).Error
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return ErrNotFound
		}
		if err != nil {
			return err
		}

		updates := map[string]any{
			"status":       PaymentStatusPaid,
			"provider_ref": providerRef,
			"paid_at":      &paidAt,
		}
		if err := tx.Model(&payment).Updates(updates).Error; err != nil {
			return err
		}

		orderUpdates := map[string]any{
			"status":  OrderStatusPaid,
			"paid_at": &paidAt,
		}
		if err := tx.Model(&payment.Order).Updates(orderUpdates).Error; err != nil {
			return err
		}

		err = tx.
			Preload("Customer").
			Preload("Order").
			Preload("LicensePlan").
			Where("order_id = ?", payment.OrderID).
			First(&license).Error
		if err == nil {
			return nil
		}
		if !errors.Is(err, gorm.ErrRecordNotFound) {
			return err
		}

		var expiresAt *time.Time
		if payment.Order.LicensePlan.DurationDays != nil {
			expires := paidAt.AddDate(0, 0, *payment.Order.LicensePlan.DurationDays)
			expiresAt = &expires
		}

		license = License{
			LicenseKey:    newLicenseKey(),
			CustomerID:    payment.Order.CustomerID,
			OrderID:       payment.Order.ID,
			LicensePlanID: payment.Order.LicensePlanID,
			Status:        LicenseStatusActive,
			IssuedAt:      paidAt,
			ExpiresAt:     expiresAt,
		}
		if err := tx.Create(&license).Error; err != nil {
			return err
		}

		return tx.
			Preload("Customer").
			Preload("Order").
			Preload("LicensePlan").
			First(&license, license.ID).Error
	})
	if err != nil {
		return nil, err
	}

	return &license, nil
}

func (s *Store) FindPaymentResult(ctx context.Context, paymentNo string) (*PaymentResult, error) {
	if strings.TrimSpace(paymentNo) == "" {
		return nil, fmt.Errorf("%w: payment number is required", ErrInvalidInput)
	}

	var payment Payment
	err := s.db.WithContext(ctx).
		Preload("Order.Customer").
		Preload("Order.LicensePlan").
		Where("payment_no = ?", strings.TrimSpace(paymentNo)).
		First(&payment).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, ErrNotFound
	}
	if err != nil {
		return nil, err
	}

	result := &PaymentResult{
		Payment: payment,
		Order:   payment.Order,
	}
	if payment.Status != PaymentStatusPaid {
		return result, nil
	}

	var license License
	err = s.db.WithContext(ctx).
		Preload("Customer").
		Preload("Order").
		Preload("LicensePlan").
		Where("order_id = ?", payment.OrderID).
		First(&license).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return result, nil
	}
	if err != nil {
		return nil, err
	}

	result.License = &license
	return result, nil
}

func (s *Store) FindLicensesByEmail(ctx context.Context, email string) ([]License, error) {
	normalizedEmail, err := NormalizeEmail(email)
	if err != nil {
		return nil, err
	}

	var licenses []License
	err = s.db.WithContext(ctx).
		Joins("Customer").
		Preload("Customer").
		Preload("Order").
		Preload("LicensePlan").
		Where("Customer.email_normalized = ?", normalizedEmail).
		Order("licenses.created_at DESC").
		Find(&licenses).Error
	if err != nil {
		return nil, err
	}

	return licenses, nil
}

func (s *Store) CreateRecoveryToken(ctx context.Context, email string, tokenHash string, expiresAt time.Time) (*RecoveryToken, error) {
	normalizedEmail, err := NormalizeEmail(email)
	if err != nil {
		return nil, err
	}
	if strings.TrimSpace(tokenHash) == "" {
		return nil, fmt.Errorf("%w: token hash is required", ErrInvalidInput)
	}
	if expiresAt.IsZero() || !expiresAt.After(time.Now()) {
		return nil, fmt.Errorf("%w: token expiry must be in the future", ErrInvalidInput)
	}

	token := RecoveryToken{
		Email:           strings.TrimSpace(email),
		EmailNormalized: normalizedEmail,
		TokenHash:       tokenHash,
		Purpose:         "license_recovery",
		ExpiresAt:       expiresAt,
	}
	if err := s.db.WithContext(ctx).Create(&token).Error; err != nil {
		return nil, err
	}

	return &token, nil
}

func (s *Store) ConsumeRecoveryToken(ctx context.Context, tokenHash string, now time.Time) (*RecoveryToken, error) {
	if strings.TrimSpace(tokenHash) == "" {
		return nil, fmt.Errorf("%w: token hash is required", ErrInvalidInput)
	}
	if now.IsZero() {
		now = time.Now()
	}

	var token RecoveryToken
	err := s.db.WithContext(ctx).Transaction(func(tx *gorm.DB) error {
		if err := tx.Where("token_hash = ?", tokenHash).First(&token).Error; err != nil {
			if errors.Is(err, gorm.ErrRecordNotFound) {
				return ErrNotFound
			}
			return err
		}
		if token.UsedAt != nil {
			return ErrTokenUsed
		}
		if !token.ExpiresAt.After(now) {
			return ErrTokenExpired
		}

		token.UsedAt = &now
		return tx.Save(&token).Error
	})
	if err != nil {
		return nil, err
	}

	return &token, nil
}

func findOrCreateCustomer(ctx context.Context, db *gorm.DB, email string, normalizedEmail string, seenAt time.Time) (Customer, error) {
	var customer Customer
	err := db.WithContext(ctx).Where("email_normalized = ?", normalizedEmail).First(&customer).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		customer = Customer{
			Email:           email,
			EmailNormalized: normalizedEmail,
			LastSeenAt:      &seenAt,
		}
		return customer, db.WithContext(ctx).Create(&customer).Error
	}
	if err != nil {
		return customer, err
	}

	customer.Email = email
	customer.LastSeenAt = &seenAt
	err = db.WithContext(ctx).Model(&customer).Updates(map[string]any{
		"email":        customer.Email,
		"last_seen_at": customer.LastSeenAt,
	}).Error
	return customer, err
}

func newPublicID(prefix string) string {
	return prefix + "_" + randomHex(12)
}

func newLicenseKey() string {
	raw := strings.ToUpper(randomHex(16))
	return fmt.Sprintf("RE-%s-%s-%s-%s", raw[0:8], raw[8:16], raw[16:24], raw[24:32])
}

func randomHex(size int) string {
	buf := make([]byte, size)
	if _, err := rand.Read(buf); err != nil {
		panic(err)
	}
	return hex.EncodeToString(buf)
}
