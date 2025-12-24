package middleware

import (
	"fmt"
	"net/http"
	"strings"
)

// contains checks if a string contains a substring
func contains(s, substr string) bool {
	return strings.Contains(s, substr)
}

// SecurityHeadersConfig holds configuration for security headers middleware
type SecurityHeadersConfig struct {
	// UseTLS indicates if TLS is enabled (affects HSTS header)
	UseTLS bool
	// ContentSecurityPolicy allows custom CSP, empty uses default
	ContentSecurityPolicy string
	// PermissionsPolicy allows custom permissions policy, empty uses default
	PermissionsPolicy string
	// EnableHSTS enables Strict-Transport-Security header (requires TLS)
	EnableHSTS bool
	// HSTSMaxAge sets the max-age for HSTS in seconds (default: 31536000 = 1 year)
	HSTSMaxAge int
	// HSTSIncludeSubDomains includes subdomains in HSTS
	HSTSIncludeSubDomains bool
}

// DefaultSecurityHeadersConfig returns sensible defaults for security headers
// All assets should be served from the embedded filesystem ('self'), not external CDNs
func DefaultSecurityHeadersConfig() SecurityHeadersConfig {
	return SecurityHeadersConfig{
		UseTLS: false,
		// CSP restricts all resources to 'self' (embedded filesystem)
		// 'unsafe-inline' is needed for script-src and style-src to support Go templates
		ContentSecurityPolicy: "default-src 'self'; " +
			"script-src 'self' 'unsafe-inline'; " +
			"style-src 'self' 'unsafe-inline'; " +
			"img-src 'self' data:; " +
			"font-src 'self' data:; " +
			"connect-src 'self'; " +
			"frame-ancestors 'self'",
		PermissionsPolicy:     "geolocation=(), microphone=(), camera=()",
		EnableHSTS:            true,
		HSTSMaxAge:            31536000, // 1 year
		HSTSIncludeSubDomains: true,
	}
}

// isStaticAsset checks if the request path is for a static asset
func isStaticAsset(path string) bool {
	// Check for common static asset paths and extensions
	if strings.HasPrefix(path, "/static/") {
		return true
	}
	// Check file extensions for assets that might be served from other paths
	lowerPath := strings.ToLower(path)
	staticExtensions := []string{".css", ".js", ".woff", ".woff2", ".ttf", ".eot", ".svg", ".png", ".jpg", ".jpeg", ".gif", ".ico", ".webp"}
	for _, ext := range staticExtensions {
		if strings.HasSuffix(lowerPath, ext) {
			return true
		}
	}
	return false
}

// isAPIEndpoint checks if the request path is for a JSON API endpoint
// CSP headers are not needed for API responses (they don't execute scripts)
func isAPIEndpoint(path string) bool {
	return strings.HasPrefix(path, "/api/") || strings.HasPrefix(path, "/analytics/")
}

// needsCSPHeader determines if a request needs Content-Security-Policy header
// CSP is only needed for HTML pages that can execute scripts, not for:
// - Static assets (CSS, JS, fonts, images)
// - API endpoints (JSON responses)
func needsCSPHeader(path string) bool {
	return !isStaticAsset(path) && !isAPIEndpoint(path)
}

// SecurityHeadersMiddleware adds security headers to all responses
func SecurityHeadersMiddleware(config SecurityHeadersConfig) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			// Core security headers - apply to all responses
			w.Header().Set("X-Content-Type-Options", "nosniff")
			w.Header().Set("Referrer-Policy", "strict-origin-when-cross-origin")

			// Check if this is a static asset request
			isStatic := isStaticAsset(r.URL.Path)

			// Cache-Control: different policies for static vs dynamic content
			if isStatic {
				// Static assets can be cached (1 year with immutable for versioned assets)
				w.Header().Set("Cache-Control", "public, max-age=31536000, immutable")
			} else {
				// Dynamic content should not be cached
				w.Header().Set("Cache-Control", "no-cache")
			}

			// CSP and Permissions-Policy only for HTML pages, not static assets or API endpoints
			if needsCSPHeader(r.URL.Path) {
				// Content Security Policy with frame-ancestors (replaces X-Frame-Options)
				csp := config.ContentSecurityPolicy
				if csp != "" {
					// Add frame-ancestors if not already present
					if !contains(csp, "frame-ancestors") {
						csp += "; frame-ancestors 'self'"
					}
					w.Header().Set("Content-Security-Policy", csp)
				}

				// Permissions Policy
				if config.PermissionsPolicy != "" {
					w.Header().Set("Permissions-Policy", config.PermissionsPolicy)
				}
			}

			// HSTS header only when using TLS (applies to all responses)
			if config.UseTLS && config.EnableHSTS {
				hstsValue := fmt.Sprintf("max-age=%d", config.HSTSMaxAge)
				if config.HSTSIncludeSubDomains {
					hstsValue += "; includeSubDomains"
				}
				w.Header().Set("Strict-Transport-Security", hstsValue)
			}

			next.ServeHTTP(w, r)
		})
	}
}

// SimpleSecurityHeadersMiddleware adds security headers with default configuration
// Use this for quick setup, or SecurityHeadersMiddleware for custom configuration
func SimpleSecurityHeadersMiddleware(useTLS bool) func(http.Handler) http.Handler {
	config := DefaultSecurityHeadersConfig()
	config.UseTLS = useTLS
	return SecurityHeadersMiddleware(config)
}
