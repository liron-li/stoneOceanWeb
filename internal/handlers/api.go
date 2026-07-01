package handlers

import (
	"context"
	"crypto/rand"
	"crypto/sha256"
	"encoding/hex"
	"errors"
	"fmt"
	"math/big"
	"net/http"
	"strings"
	"time"

	"stone-ocean-web/internal/events"
	"stone-ocean-web/internal/store"

	"github.com/gin-gonic/gin"
)

type API struct {
	store                 *store.Store
	events                *events.Bus
	recoveryVerifyLimiter *ipRateLimiter
}

// checkoutRequest 是前端创建订单时提交的 JSON 请求体。
type checkoutRequest struct {
	Email         string              `json:"email"`
	License       store.LicenseKind   `json:"license"`
	PaymentMethod store.PaymentMethod `json:"paymentMethod"`
	Locale        string              `json:"locale"`
}

// recoveryCodeRequest 是发送找回验证码接口使用的 JSON 请求体。
type recoveryCodeRequest struct {
	Email  string `json:"email"`
	Locale string `json:"locale"`
}

// recoveryCodeVerifyRequest 是校验找回验证码接口使用的 JSON 请求体。
type recoveryCodeVerifyRequest struct {
	Email string `json:"email"`
	Code  string `json:"code"`
}

// NewAPI 创建一组依赖 Store 的 API 处理器。
func NewAPI(appStore *store.Store) *API {
	return NewAPIWithEvents(appStore, nil)
}

func NewAPIWithEvents(appStore *store.Store, eventBus *events.Bus) *API {
	return &API{
		store:                 appStore,
		events:                eventBus,
		recoveryVerifyLimiter: newIPRateLimiter(5, 10*time.Minute),
	}
}

// CreateCheckout 根据邮箱、授权类型和支付方式创建待支付订单。
func (api *API) CreateCheckout(c *gin.Context) {
	if !api.ready(c) {
		return
	}

	var req checkoutRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid checkout request."})
		return
	}

	// 前端只提交授权类型，后端再映射为数据库中的套餐 code。
	planCode, ok := planCodeForLicense(req.License)
	if !ok {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid license type."})
		return
	}

	result, err := api.store.CreateCheckoutOrder(c.Request.Context(), store.CreateCheckoutOrderInput{
		Email:         req.Email,
		PlanCode:      planCode,
		PaymentMethod: req.PaymentMethod,
		Locale:        req.Locale,
	})
	if err != nil {
		api.writeStoreError(c, err)
		return
	}

	// 当前支付链接指向本地确认接口，方便前端演示完整支付到发码流程。
	c.JSON(http.StatusOK, gin.H{
		"orderNo":       result.Order.OrderNo,
		"paymentNo":     result.Payment.PaymentNo,
		"amountCents":   result.Order.AmountCents,
		"currency":      result.Order.Currency,
		"status":        result.Order.Status,
		"paymentMethod": result.Payment.Method,
		"paymentUrl":    "/api/payments/" + result.Payment.PaymentNo + "/confirm",
	})
}

// ConfirmPayment 将指定支付单标记为已支付，并返回生成或已有的授权码。
func (api *API) ConfirmPayment(c *gin.Context) {
	if !api.ready(c) {
		return
	}

	paymentNo := c.Param("paymentNo")
	alreadyPaid := api.paymentAlreadyDelivered(c.Request.Context(), paymentNo)
	license, err := api.store.MarkPaymentPaid(c.Request.Context(), paymentNo, "frontend-test-confirm", time.Now())
	if err != nil {
		api.writeStoreError(c, err)
		return
	}
	if !alreadyPaid {
		api.publishPaymentPaid(license)
	}

	c.JSON(http.StatusOK, gin.H{
		"license": licenseResponse(license),
	})
}

// PaymentResult 查询支付单当前状态；支付成功后会一并返回授权码。
func (api *API) PaymentResult(c *gin.Context) {
	if !api.ready(c) {
		return
	}

	paymentNo := c.Param("paymentNo")
	result, err := api.store.FindPaymentResult(c.Request.Context(), paymentNo)
	if err != nil {
		api.writeStoreError(c, err)
		return
	}

	response := gin.H{
		"paymentNo":     result.Payment.PaymentNo,
		"orderNo":       result.Order.OrderNo,
		"status":        result.Payment.Status,
		"amountCents":   result.Order.AmountCents,
		"currency":      result.Order.Currency,
		"paymentMethod": result.Payment.Method,
	}
	if result.Payment.PaidAt != nil {
		response["paidAt"] = result.Payment.PaidAt.Format(time.RFC3339)
	}
	if result.License != nil {
		response["license"] = licenseResponse(result.License)
	}

	c.JSON(http.StatusOK, response)
}

// SendLicenseRecoveryCode 根据购买邮箱发送找回验证码，不直接返回授权码。
func (api *API) SendLicenseRecoveryCode(c *gin.Context) {
	if !api.ready(c) {
		return
	}

	var req recoveryCodeRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid recovery request."})
		return
	}

	licenses, err := api.store.FindLicensesByEmail(c.Request.Context(), req.Email)
	if err != nil {
		api.writeStoreError(c, err)
		return
	}

	// 无论是否存在购买记录都返回通用成功，避免接口泄露邮箱是否购买过。
	if len(licenses) == 0 {
		c.JSON(http.StatusOK, gin.H{"ok": true})
		return
	}

	code, err := newRecoveryCode()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Something went wrong. Please try again."})
		return
	}
	tokenHash, err := recoveryCodeHash(req.Email, code)
	if err != nil {
		api.writeStoreError(c, err)
		return
	}
	expiresAt := time.Now().Add(10 * time.Minute)
	if _, err := api.store.CreateRecoveryToken(c.Request.Context(), req.Email, tokenHash, expiresAt); err != nil {
		api.writeStoreError(c, err)
		return
	}

	api.publishLicenseRecoveryCode(req.Email, code, req.Locale, expiresAt)
	c.JSON(http.StatusOK, gin.H{"ok": true})
}

// VerifyLicenseRecoveryCode 校验邮箱验证码；验证通过后返回该邮箱下的授权码列表。
func (api *API) VerifyLicenseRecoveryCode(c *gin.Context) {
	if !api.ready(c) {
		return
	}

	now := time.Now()
	clientIP := c.ClientIP()
	if !api.recoveryVerifyLimiter.Allow(clientIP, now) {
		c.JSON(http.StatusTooManyRequests, gin.H{"error": "Too many verification attempts. Please try again later."})
		return
	}

	var req recoveryCodeVerifyRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		api.recoveryVerifyLimiter.RecordFailure(clientIP, now)
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid recovery request."})
		return
	}

	tokenHash, err := recoveryCodeHash(req.Email, req.Code)
	if err != nil {
		api.recoveryVerifyLimiter.RecordFailure(clientIP, now)
		api.writeStoreError(c, err)
		return
	}
	token, err := api.store.ConsumeRecoveryToken(c.Request.Context(), tokenHash, now)
	if err != nil {
		switch {
		case errors.Is(err, store.ErrNotFound), errors.Is(err, store.ErrTokenExpired), errors.Is(err, store.ErrTokenUsed):
			api.recoveryVerifyLimiter.RecordFailure(clientIP, now)
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid or expired verification code."})
		default:
			api.writeStoreError(c, err)
		}
		return
	}
	api.recoveryVerifyLimiter.Reset(clientIP)

	licenses, err := api.store.FindLicensesByEmail(c.Request.Context(), token.Email)
	if err != nil {
		api.writeStoreError(c, err)
		return
	}

	items := make([]gin.H, 0, len(licenses))
	for i := range licenses {
		items = append(items, licenseResponse(&licenses[i]))
	}

	c.JSON(http.StatusOK, gin.H{
		"licenses": items,
	})
}

func (api *API) ready(c *gin.Context) bool {
	if api == nil || api.store == nil {
		c.JSON(http.StatusServiceUnavailable, gin.H{"error": "Database is not enabled."})
		return false
	}
	return true
}

func (api *API) writeStoreError(c *gin.Context, err error) {
	switch {
	case errors.Is(err, store.ErrInvalidInput), errors.Is(err, store.ErrPaymentMethod):
		c.JSON(http.StatusBadRequest, gin.H{"error": "Please check the information and try again."})
	case errors.Is(err, store.ErrNotFound):
		c.JSON(http.StatusNotFound, gin.H{"error": "Record not found."})
	default:
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Something went wrong. Please try again."})
	}
}

func (api *API) publishPaymentPaid(license *store.License) {
	if api == nil || api.events == nil || license == nil {
		return
	}
	api.events.PublishPaymentPaid(events.PaymentPaidEvent{License: license})
}

func (api *API) publishLicenseRecoveryCode(email, code, locale string, expiresAt time.Time) {
	if api == nil || api.events == nil {
		return
	}
	api.events.PublishLicenseRecoveryCode(events.LicenseRecoveryCodeEvent{
		Email:     email,
		Code:      code,
		Locale:    store.NormalizeLocale(locale),
		ExpiresAt: expiresAt,
	})
}

func (api *API) paymentAlreadyDelivered(ctx context.Context, paymentNo string) bool {
	result, err := api.store.FindPaymentResult(ctx, paymentNo)
	if err != nil {
		return false
	}
	return result.Payment.Status == store.PaymentStatusPaid && result.License != nil
}

func planCodeForLicense(kind store.LicenseKind) (string, bool) {
	switch kind {
	case store.LicenseKindMonthly:
		return store.DefaultMonthlyPlanCode, true
	case store.LicenseKindLifetime:
		return store.DefaultLifetimePlanCode, true
	default:
		return "", false
	}
}

func newRecoveryCode() (string, error) {
	value, err := rand.Int(rand.Reader, big.NewInt(1000000))
	if err != nil {
		return "", err
	}
	return fmt.Sprintf("%06d", value.Int64()), nil
}

func recoveryCodeHash(email, code string) (string, error) {
	normalizedEmail, err := store.NormalizeEmail(email)
	if err != nil {
		return "", err
	}
	code = strings.TrimSpace(code)
	if len(code) != 6 {
		return "", fmt.Errorf("%w: code must be six digits", store.ErrInvalidInput)
	}
	for _, char := range code {
		if char < '0' || char > '9' {
			return "", fmt.Errorf("%w: code must be six digits", store.ErrInvalidInput)
		}
	}

	sum := sha256.Sum256([]byte(normalizedEmail + ":" + code))
	return hex.EncodeToString(sum[:]), nil
}

func licenseResponse(license *store.License) gin.H {
	var expiresAt any
	if license.ExpiresAt != nil {
		expiresAt = license.ExpiresAt.Format(time.RFC3339)
	}

	return gin.H{
		"licenseKey": license.LicenseKey,
		"status":     license.Status,
		"kind":       license.LicensePlan.Kind,
		"plan":       license.LicensePlan.Name,
		"orderNo":    license.Order.OrderNo,
		"issuedAt":   license.IssuedAt.Format(time.RFC3339),
		"expiresAt":  expiresAt,
	}
}
