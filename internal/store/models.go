package store

import "time"

type LicenseKind string
type OrderStatus string
type PaymentMethod string
type PaymentStatus string
type LicenseStatus string

const (
	LicenseKindMonthly  LicenseKind = "monthly"
	LicenseKindLifetime LicenseKind = "lifetime"

	OrderStatusPending  OrderStatus = "pending"
	OrderStatusPaid     OrderStatus = "paid"
	OrderStatusCanceled OrderStatus = "canceled"
	OrderStatusRefunded OrderStatus = "refunded"

	PaymentMethodUSDT PaymentMethod = "usdt"
	PaymentMethodCard PaymentMethod = "card"

	PaymentStatusPending PaymentStatus = "pending"
	PaymentStatusPaid    PaymentStatus = "paid"
	PaymentStatusFailed  PaymentStatus = "failed"

	LicenseStatusActive  LicenseStatus = "active"
	LicenseStatusExpired LicenseStatus = "expired"
	LicenseStatusRevoked LicenseStatus = "revoked"
)

type Customer struct {
	ID              uint       `gorm:"primaryKey"`
	Email           string     `gorm:"size:320;not null"`
	EmailNormalized string     `gorm:"size:320;not null;uniqueIndex"`
	LastSeenAt      *time.Time `gorm:"index"`
	CreatedAt       time.Time
	UpdatedAt       time.Time
}

type LicensePlan struct {
	ID           uint        `gorm:"primaryKey"`
	Code         string      `gorm:"size:64;not null;uniqueIndex"`
	Name         string      `gorm:"size:120;not null"`
	Kind         LicenseKind `gorm:"size:32;not null;index"`
	DurationDays *int
	PriceCents   int64  `gorm:"not null"`
	Currency     string `gorm:"size:8;not null;default:USD"`
	IsActive     bool   `gorm:"not null;default:true;index"`
	CreatedAt    time.Time
	UpdatedAt    time.Time
}

type Order struct {
	ID            uint          `gorm:"primaryKey"`
	OrderNo       string        `gorm:"size:40;not null;uniqueIndex"`
	CustomerID    uint          `gorm:"not null;index"`
	Customer      Customer      `gorm:"constraint:OnUpdate:CASCADE,OnDelete:RESTRICT"`
	LicensePlanID uint          `gorm:"not null;index"`
	LicensePlan   LicensePlan   `gorm:"constraint:OnUpdate:CASCADE,OnDelete:RESTRICT"`
	Email         string        `gorm:"size:320;not null"`
	LicenseKind   LicenseKind   `gorm:"size:32;not null"`
	AmountCents   int64         `gorm:"not null"`
	Currency      string        `gorm:"size:8;not null"`
	Status        OrderStatus   `gorm:"size:32;not null;index"`
	PaymentMethod PaymentMethod `gorm:"size:32;not null"`
	PaidAt        *time.Time    `gorm:"index"`
	CreatedAt     time.Time
	UpdatedAt     time.Time
}

type Payment struct {
	ID          uint          `gorm:"primaryKey"`
	PaymentNo   string        `gorm:"size:40;not null;uniqueIndex"`
	OrderID     uint          `gorm:"not null;index"`
	Order       Order         `gorm:"constraint:OnUpdate:CASCADE,OnDelete:RESTRICT"`
	Method      PaymentMethod `gorm:"size:32;not null;index"`
	Provider    string        `gorm:"size:64"`
	ProviderRef string        `gorm:"size:128;index"`
	Status      PaymentStatus `gorm:"size:32;not null;index"`
	AmountCents int64         `gorm:"not null"`
	Currency    string        `gorm:"size:8;not null"`
	PaidAt      *time.Time    `gorm:"index"`
	CreatedAt   time.Time
	UpdatedAt   time.Time
}

type License struct {
	ID            uint          `gorm:"primaryKey"`
	LicenseKey    string        `gorm:"size:80;not null;uniqueIndex"`
	CustomerID    uint          `gorm:"not null;index"`
	Customer      Customer      `gorm:"constraint:OnUpdate:CASCADE,OnDelete:RESTRICT"`
	OrderID       uint          `gorm:"not null;uniqueIndex"`
	Order         Order         `gorm:"constraint:OnUpdate:CASCADE,OnDelete:RESTRICT"`
	LicensePlanID uint          `gorm:"not null;index"`
	LicensePlan   LicensePlan   `gorm:"constraint:OnUpdate:CASCADE,OnDelete:RESTRICT"`
	Status        LicenseStatus `gorm:"size:32;not null;index"`
	IssuedAt      time.Time     `gorm:"not null;index"`
	ExpiresAt     *time.Time    `gorm:"index"`
	RevokedAt     *time.Time
	CreatedAt     time.Time
	UpdatedAt     time.Time
}

type RecoveryToken struct {
	ID              uint       `gorm:"primaryKey"`
	Email           string     `gorm:"size:320;not null"`
	EmailNormalized string     `gorm:"size:320;not null;index"`
	TokenHash       string     `gorm:"size:128;not null;uniqueIndex"`
	Purpose         string     `gorm:"size:64;not null;default:license_recovery"`
	ExpiresAt       time.Time  `gorm:"not null;index"`
	UsedAt          *time.Time `gorm:"index"`
	CreatedAt       time.Time
	UpdatedAt       time.Time
}
