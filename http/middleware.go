package http

import (
	"net/http"
)

// Middleware is an alias for a function that takes a handler and returns one, too.
type Middleware = func(http.Handler) http.Handler

// NoClickjacking middleware sets headers to disallow frame embedding and XSS protection for older browsers.
// See https://developer.mozilla.org/en-US/docs/Web/HTTP/Headers/X-Frame-Options
// See https://developer.mozilla.org/en-US/docs/Web/HTTP/Headers/X-XSS-Protection
func NoClickjacking(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("X-Frame-Options", "deny")
		w.Header().Set("X-XSS-Protection", "1; mode=block")
		next.ServeHTTP(w, r)
	})
}

// StrictContentSecurityPolicy sets best practice CSP headers.
// This disallows all external img, script, and style links, and disallows all objects (flash etc.).
// See https://infosec.mozilla.org/guidelines/web_security#content-security-policy
func StrictContentSecurityPolicy(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Security-Policy",
			"default-src 'none'; "+
				"connect-src 'self'; "+
				"font-src 'self'; "+
				"img-src 'self'; "+
				"manifest-src 'self'; "+
				"media-src 'self'; "+
				"script-src 'self'; "+
				"style-src 'self';")
		next.ServeHTTP(w, r)
	})
}
