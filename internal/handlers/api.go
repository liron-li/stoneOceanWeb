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

type checkoutRequest struct {
	Email         string              `json:"email"`
	License       store.LicenseKind   `json:"license"`
	PaymentMethod store.PaymentMethod `json:"paymentMethod"`
}

type recoveryRequest struct {
	Email string `json:"email"`
}

func NewAPI(appStore *store.Store) *API {
	return &API{store: appStore}
}

func (api *API) CreateCheckout(c *gin.Context) {
	if !api.ready(c) {
		return
	}

	var req checkoutRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid checkout request."})
		return
	}

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
