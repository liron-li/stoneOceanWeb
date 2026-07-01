package router

import (
	"net/http"

	"stone-ocean-web/internal/events"
	"stone-ocean-web/internal/handlers"
	"stone-ocean-web/internal/store"

	"github.com/gin-gonic/gin"
)

func New(appStore *store.Store) *gin.Engine {
	return NewWithEvents(appStore, nil)
}

func NewWithEvents(appStore *store.Store, eventBus *events.Bus) *gin.Engine {
	r := gin.Default()
	api := handlers.NewAPIWithEvents(appStore, eventBus)

	r.Static("/static", "./web/static")
	r.LoadHTMLGlob("web/templates/*")

	r.GET("/", handlers.Home)
	r.GET("/checkout", handlers.Checkout)
	r.GET("/checkout/success", handlers.CheckoutSuccess)
	r.GET("/license-recovery", handlers.LicenseRecovery)
	r.GET("/healthz", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"status": "ok"})
	})
	r.POST("/api/checkout", api.CreateCheckout)
	r.POST("/api/payments/:paymentNo/confirm", api.ConfirmPayment)
	r.GET("/api/payments/:paymentNo/result", api.PaymentResult)
	r.POST("/api/license-recovery", api.RecoverLicenses)
	r.GET("/robots.txt", handlers.Robots)
	r.GET("/sitemap.xml", handlers.Sitemap)
	r.GET("/:locale/checkout", handlers.LocalizedCheckout)
	r.GET("/:locale/checkout/success", handlers.LocalizedCheckoutSuccess)
	r.GET("/:locale/license-recovery", handlers.LocalizedLicenseRecovery)
	r.GET("/:locale", handlers.LocalizedHome)

	return r
}
