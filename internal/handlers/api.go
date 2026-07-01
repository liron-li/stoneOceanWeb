package handlers

import (
	"errors"
	"net/http"
	"time"

	"stone-ocean-web/internal/store"

	"github.com/gin-gonic/gin"
)

type API struct {
	store *store.Store
}

// checkoutRequest 是前端创建订单时提交的 JSON 请求体。
type checkoutRequest struct {
	Email         string              `json:"email"`
	License       store.LicenseKind   `json:"license"`
	PaymentMethod store.PaymentMethod `json:"paymentMethod"`
}

// recoveryRequest 是找回授权码接口使用的 JSON 请求体。
type recoveryRequest struct {
	Email string `json:"email"`
}

// NewAPI 创建一组依赖 Store 的 API 处理器。
func NewAPI(appStore *store.Store) *API {
	return &API{store: appStore}
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
	license, err := api.store.MarkPaymentPaid(c.Request.Context(), paymentNo, "frontend-test-confirm", time.Now())
	if err != nil {
		api.writeStoreError(c, err)
		return
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

// RecoverLicenses 根据购买邮箱查询历史授权码。
func (api *API) RecoverLicenses(c *gin.Context) {
	if !api.ready(c) {
		return
	}

	var req recoveryRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid recovery request."})
		return
	}

	licenses, err := api.store.FindLicensesByEmail(c.Request.Context(), req.Email)
	if err != nil {
		api.writeStoreError(c, err)
		return
	}

	// 将数据库模型转换为前端需要的轻量 JSON 列表。
	items := make([]gin.H, 0, len(licenses))
	for i := range licenses {
		items = append(items, licenseResponse(&licenses[i]))
	}

	c.JSON(http.StatusOK, gin.H{
		"licenses": items,
	})
}

// ready 统一检查 API 是否已经接入可用数据库。
func (api *API) ready(c *gin.Context) bool {
	if api == nil || api.store == nil {
		c.JSON(http.StatusServiceUnavailable, gin.H{"error": "Database is not enabled."})
		return false
	}
	return true
}

// writeStoreError 将 store 层错误转换成对外 HTTP 状态码和通用提示。
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

// planCodeForLicense 把页面上的授权类型映射到默认套餐 code。
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

// licenseResponse 统一授权码接口响应字段，时间字段使用 RFC3339 格式。
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
