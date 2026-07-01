package router

import (
	"net/http"
	"net/http/httptest"
	"os"
	"path/filepath"
	"strings"
	"testing"

	"github.com/gin-gonic/gin"
)

func TestSwaggerOpenAPIYAMLIsServed(t *testing.T) {
	gin.SetMode(gin.TestMode)
	previousWD, err := os.Getwd()
	if err != nil {
		t.Fatalf("os.Getwd() error = %v", err)
	}
	if err := os.Chdir(filepath.Join("..", "..")); err != nil {
		t.Fatalf("os.Chdir() error = %v", err)
	}
	t.Cleanup(func() {
		_ = os.Chdir(previousWD)
	})

	router := NewWithEvents(nil, nil)
	request := httptest.NewRequest(http.MethodGet, "/swagger/openapi.yaml", nil)
	response := httptest.NewRecorder()
	router.ServeHTTP(response, request)

	if response.Code != http.StatusOK {
		t.Fatalf("GET /swagger/openapi.yaml status = %d, want %d", response.Code, http.StatusOK)
	}
	if !strings.Contains(response.Body.String(), "openapi: 3.0.3") {
		t.Fatalf("GET /swagger/openapi.yaml body does not look like OpenAPI YAML")
	}
}
