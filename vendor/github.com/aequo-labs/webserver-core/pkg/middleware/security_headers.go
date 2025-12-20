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
func DefaultSecurityHeadersConfig() SecurityHeadersConfig {
	return SecurityHeadersConfig{
		UseTLS:                false,
		ContentSecurityPolicy: "default-src 'self'; script-src 'self'; style-src 'self' 'unsafe-inline'; img-src 'self' data:; font-src 'self'; frame-ancestors 'self'",
		PermissionsPolicy:     "geolocation=(), microphone=(), camera=()",
		EnableHSTS:            true,
		HSTSMaxAge:            31536000, // 1 year
		HSTSIncludeSubDomains: true,
	}
}

// SecurityHeadersMiddleware adds security headers to all responses
func SecurityHeadersMiddleware(config SecurityHeadersConfig) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			// Core security headers
			w.Header().Set("X-Content-Type-Options", "nosniff")
			w.Header().Set("Referrer-Policy", "strict-origin-when-cross-origin")

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

			// HSTS header only when using TLS
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
