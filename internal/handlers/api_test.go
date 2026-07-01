package handlers

import (
	"bytes"
	"context"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"stone-ocean-web/internal/events"
	"stone-ocean-web/internal/store"

	"github.com/gin-gonic/gin"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

type recoveryCodeRecorder struct {
	called chan events.LicenseRecoveryCodeEvent
}

func (r recoveryCodeRecorder) HandleLicenseRecoveryCode(ctx context.Context, event events.LicenseRecoveryCodeEvent) error {
	select {
	case r.called <- event:
	case <-ctx.Done():
		return ctx.Err()
	}
	return nil
}

func TestLicenseRecoveryRequiresVerificationCode(t *testing.T) {
	gin.SetMode(gin.TestMode)

	appStore := newAPITestStore(t)
	checkout, err := appStore.CreateCheckoutOrder(context.Background(), store.CreateCheckoutOrderInput{
		Email:         "buyer@example.com",
		PlanCode:      store.DefaultLifetimePlanCode,
		PaymentMethod: store.PaymentMethodCard,
	})
	if err != nil {
		t.Fatalf("CreateCheckoutOrder() error = %v", err)
	}
	license, err := appStore.MarkPaymentPaid(context.Background(), checkout.Payment.PaymentNo, "provider-test", time.Now())
	if err != nil {
		t.Fatalf("MarkPaymentPaid() error = %v", err)
	}

	called := make(chan events.LicenseRecoveryCodeEvent, 1)
	eventBus := events.NewBus(nil)
	eventBus.AddLicenseRecoveryCodeListener(recoveryCodeRecorder{called: called})

	api := NewAPIWithEvents(appStore, eventBus)
	router := gin.New()
	router.POST("/api/license-recovery/verification-code", api.SendLicenseRecoveryCode)
	router.POST("/api/license-recovery/verification-code/verify", api.VerifyLicenseRecoveryCode)

	firstResponse := postJSONForTest(t, router, "/api/license-recovery/verification-code", map[string]string{
		"email":  "buyer@example.com",
		"locale": "en",
	})
	if _, ok := firstResponse["licenses"]; ok {
		t.Fatal("first recovery step returned licenses before verification")
	}
	if firstResponse["ok"] != true {
		t.Fatalf("first recovery step ok = %#v, want true", firstResponse["ok"])
	}

	var event events.LicenseRecoveryCodeEvent
	select {
	case event = <-called:
	case <-time.After(time.Second):
		t.Fatal("recovery code event was not published")
	}

	verifyResponse := postJSONForTest(t, router, "/api/license-recovery/verification-code/verify", map[string]string{
		"email": "buyer@example.com",
		"code":  event.Code,
	})
	items, ok := verifyResponse["licenses"].([]any)
	if !ok || len(items) != 1 {
		t.Fatalf("verified response licenses = %#v, want one item", verifyResponse["licenses"])
	}
	item, ok := items[0].(map[string]any)
	if !ok {
		t.Fatalf("license response item = %#v", items[0])
	}
	if item["licenseKey"] != license.LicenseKey {
		t.Fatalf("licenseKey = %#v, want %q", item["licenseKey"], license.LicenseKey)
	}
}

func TestLicenseRecoveryVerifyRateLimitsFailedAttemptsByIP(t *testing.T) {
	gin.SetMode(gin.TestMode)

	appStore := newAPITestStore(t)
	checkout, err := appStore.CreateCheckoutOrder(context.Background(), store.CreateCheckoutOrderInput{
		Email:         "limited@example.com",
		PlanCode:      store.DefaultLifetimePlanCode,
		PaymentMethod: store.PaymentMethodCard,
	})
	if err != nil {
		t.Fatalf("CreateCheckoutOrder() error = %v", err)
	}
	if _, err := appStore.MarkPaymentPaid(context.Background(), checkout.Payment.PaymentNo, "provider-test", time.Now()); err != nil {
		t.Fatalf("MarkPaymentPaid() error = %v", err)
	}

	called := make(chan events.LicenseRecoveryCodeEvent, 1)
	eventBus := events.NewBus(nil)
	eventBus.AddLicenseRecoveryCodeListener(recoveryCodeRecorder{called: called})

	api := NewAPIWithEvents(appStore, eventBus)
	router := gin.New()
	router.POST("/api/license-recovery/verification-code", api.SendLicenseRecoveryCode)
	router.POST("/api/license-recovery/verification-code/verify", api.VerifyLicenseRecoveryCode)

	postJSONForTest(t, router, "/api/license-recovery/verification-code", map[string]string{
		"email":  "limited@example.com",
		"locale": "en",
	})

	for i := 0; i < 5; i++ {
		response := postJSONWithStatusForTest(t, router, "/api/license-recovery/verification-code/verify", map[string]string{
			"email": "limited@example.com",
			"code":  "000000",
		}, http.StatusBadRequest)
		if response["error"] == "" {
			t.Fatalf("attempt %d returned empty error", i+1)
		}
	}

	response := postJSONWithStatusForTest(t, router, "/api/license-recovery/verification-code/verify", map[string]string{
		"email": "limited@example.com",
		"code":  "000000",
	}, http.StatusTooManyRequests)
	if response["error"] == "" {
		t.Fatal("rate limited response returned empty error")
	}
}

func postJSONForTest(t *testing.T, handler http.Handler, path string, body any) map[string]any {
	t.Helper()

	return postJSONWithStatusForTest(t, handler, path, body, http.StatusOK)
}

func postJSONWithStatusForTest(t *testing.T, handler http.Handler, path string, body any, wantStatus int) map[string]any {
	t.Helper()

	payload, err := json.Marshal(body)
	if err != nil {
		t.Fatalf("json.Marshal() error = %v", err)
	}
	request := httptest.NewRequest(http.MethodPost, path, bytes.NewReader(payload))
	request.Header.Set("Content-Type", "application/json")
	response := httptest.NewRecorder()
	handler.ServeHTTP(response, request)

	if response.Code != wantStatus {
		t.Fatalf("POST %s status = %d, want %d, body = %s", path, response.Code, wantStatus, response.Body.String())
	}

	var result map[string]any
	if err := json.Unmarshal(response.Body.Bytes(), &result); err != nil {
		t.Fatalf("json.Unmarshal() error = %v", err)
	}
	return result
}

func newAPITestStore(t *testing.T) *store.Store {
	t.Helper()

	db, err := gorm.Open(sqlite.Open("file::memory:?cache=shared"), &gorm.Config{})
	if err != nil {
		t.Fatalf("gorm.Open() error = %v", err)
	}
	if err := store.AutoMigrate(db); err != nil {
		t.Fatalf("AutoMigrate() error = %v", err)
	}
	if err := store.SeedDefaultPlans(context.Background(), db); err != nil {
		t.Fatalf("SeedDefaultPlans() error = %v", err)
	}

	return store.New(db)
}
