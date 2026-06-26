package handlers

import (
	"encoding/xml"
	"net/http"
	"strings"
	"time"

	"stone-ocean-web/internal/i18n"

	"github.com/gin-gonic/gin"
)

func Home(c *gin.Context) {
	locale := i18n.MatchAcceptLanguage(c.GetHeader("Accept-Language"))
	c.Redirect(http.StatusFound, i18n.Path(locale))
}

func Checkout(c *gin.Context) {
	locale := i18n.MatchAcceptLanguage(c.GetHeader("Accept-Language"))
	c.Redirect(http.StatusFound, i18n.CheckoutPath(locale))
}

func LicenseRecovery(c *gin.Context) {
	locale := i18n.MatchAcceptLanguage(c.GetHeader("Accept-Language"))
	c.Redirect(http.StatusFound, i18n.LicenseRecoveryPath(locale))
}

func LocalizedHome(c *gin.Context) {
	locale := c.Param("locale")
	if !i18n.Supported(locale) {
		c.Redirect(http.StatusMovedPermanently, i18n.Path(i18n.DefaultLocale))
		return
	}

	renderHome(c, locale)
}

func LocalizedCheckout(c *gin.Context) {
	locale := c.Param("locale")
	if !i18n.Supported(locale) {
		c.Redirect(http.StatusMovedPermanently, i18n.CheckoutPath(i18n.DefaultLocale))
		return
	}

	renderCheckout(c, locale)
}

func LocalizedLicenseRecovery(c *gin.Context) {
	locale := c.Param("locale")
	if !i18n.Supported(locale) {
		c.Redirect(http.StatusMovedPermanently, i18n.LicenseRecoveryPath(i18n.DefaultLocale))
		return
	}

	renderLicenseRecovery(c, locale)
}

func Robots(c *gin.Context) {
	baseURL := requestBaseURL(c)
	c.String(http.StatusOK, "User-agent: *\nAllow: /\nSitemap: %s/sitemap.xml\n", strings.TrimRight(baseURL, "/"))
}

func Sitemap(c *gin.Context) {
	baseURL := requestBaseURL(c)
	now := time.Now().Format("2006-01-02")
	urls := make([]sitemapURL, 0, len(i18n.Languages(i18n.DefaultLocale)))

	for _, lang := range i18n.Languages(i18n.DefaultLocale) {
		urls = append(urls, sitemapURL{
			Loc:     strings.TrimRight(baseURL, "/") + lang.Path,
			LastMod: now,
		})
		urls = append(urls, sitemapURL{
			Loc:     strings.TrimRight(baseURL, "/") + i18n.CheckoutPath(lang.Code),
			LastMod: now,
		})
		urls = append(urls, sitemapURL{
			Loc:     strings.TrimRight(baseURL, "/") + i18n.LicenseRecoveryPath(lang.Code),
			LastMod: now,
		})
	}

	c.XML(http.StatusOK, sitemap{
		XMLName: xml.Name{Local: "urlset"},
		Xmlns:   "http://www.sitemaps.org/schemas/sitemap/0.9",
		URLs:    urls,
	})
}

func renderHome(c *gin.Context, locale string) {
	baseURL := requestBaseURL(c)
	path := i18n.Path(locale)

	c.HTML(http.StatusOK, "index.tmpl", gin.H{
		"Title":        i18n.T(locale, "meta.title"),
		"Description":  i18n.T(locale, "meta.description"),
		"Locale":       locale,
		"HTMLLang":     i18n.HTMLLang(locale),
		"Canonical":    strings.TrimRight(baseURL, "/") + path,
		"DefaultURL":   strings.TrimRight(baseURL, "/") + "/",
		"Alternates":   i18n.Alternates(baseURL),
		"Languages":    i18n.Languages(locale),
		"HomePath":     path,
		"CheckoutPath": i18n.CheckoutPath(locale),
		"T": func(key string) string {
			return i18n.T(locale, key)
		},
	})
}

func renderCheckout(c *gin.Context, locale string) {
	baseURL := requestBaseURL(c)
	path := i18n.CheckoutPath(locale)

	c.HTML(http.StatusOK, "checkout.tmpl", gin.H{
		"Title":               i18n.T(locale, "checkout.meta.title"),
		"Description":         i18n.T(locale, "checkout.meta.description"),
		"Locale":              locale,
		"HTMLLang":            i18n.HTMLLang(locale),
		"Canonical":           strings.TrimRight(baseURL, "/") + path,
		"DefaultURL":          strings.TrimRight(baseURL, "/") + "/checkout",
		"Alternates":          i18n.AlternatesForPath(baseURL, "/checkout"),
		"Languages":           i18n.LanguagesForPath(locale, "/checkout"),
		"HomePath":            i18n.Path(locale),
		"CheckoutPath":        path,
		"LicenseRecoveryPath": i18n.LicenseRecoveryPath(locale),
		"T": func(key string) string {
			return i18n.T(locale, key)
		},
	})
}

func renderLicenseRecovery(c *gin.Context, locale string) {
	baseURL := requestBaseURL(c)
	path := i18n.LicenseRecoveryPath(locale)

	c.HTML(http.StatusOK, "license_recovery.tmpl", gin.H{
		"Title":               i18n.T(locale, "recovery.meta.title"),
		"Description":         i18n.T(locale, "recovery.meta.description"),
		"Locale":              locale,
		"HTMLLang":            i18n.HTMLLang(locale),
		"Canonical":           strings.TrimRight(baseURL, "/") + path,
		"DefaultURL":          strings.TrimRight(baseURL, "/") + "/license-recovery",
		"Alternates":          i18n.AlternatesForPath(baseURL, "/license-recovery"),
		"Languages":           i18n.LanguagesForPath(locale, "/license-recovery"),
		"HomePath":            i18n.Path(locale),
		"CheckoutPath":        i18n.CheckoutPath(locale),
		"LicenseRecoveryPath": path,
		"T": func(key string) string {
			return i18n.T(locale, key)
		},
	})
}

func requestBaseURL(c *gin.Context) string {
	scheme := c.GetHeader("X-Forwarded-Proto")
	if scheme == "" {
		scheme = "http"
	}

	host := c.Request.Host
	if forwardedHost := c.GetHeader("X-Forwarded-Host"); forwardedHost != "" {
		host = forwardedHost
	}

	return scheme + "://" + host
}

type sitemap struct {
	XMLName xml.Name     `xml:"urlset"`
	Xmlns   string       `xml:"xmlns,attr"`
	URLs    []sitemapURL `xml:"url"`
}

type sitemapURL struct {
	Loc     string `xml:"loc"`
	LastMod string `xml:"lastmod"`
}
