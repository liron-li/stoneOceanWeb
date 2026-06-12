package router

import (
	"net/http"

	"stone-ocean-web/internal/handlers"

	"github.com/gin-gonic/gin"
)

func New() *gin.Engine {
	r := gin.Default()

	r.Static("/static", "./web/static")
	r.LoadHTMLGlob("web/templates/*")

	r.GET("/", handlers.Home)
	r.GET("/checkout", handlers.Checkout)
	r.GET("/healthz", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"status": "ok"})
	})
	r.GET("/robots.txt", handlers.Robots)
	r.GET("/sitemap.xml", handlers.Sitemap)
	r.GET("/:locale/checkout", handlers.LocalizedCheckout)
	r.GET("/:locale", handlers.LocalizedHome)

	return r
}
