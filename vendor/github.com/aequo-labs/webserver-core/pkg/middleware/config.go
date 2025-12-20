package middleware

import "net/http"

// MiddlewareConfig holds configuration for all middleware
type MiddlewareConfig struct {
	Minification MinifyConfig
	Gzip         GzipConfig
}

// GzipConfig holds configuration for gzip compression middleware
type GzipConfig struct {
	Enabled bool
}

// NewMiddlewareConfig returns a default middleware configuration
func NewMiddlewareConfig() *MiddlewareConfig {
	return &MiddlewareConfig{
		Minification: MinifyConfig{Enabled: false}, // Default disabled for debugging safety
		Gzip:         GzipConfig{Enabled: false},   // Default disabled for debugging safety
	}
}

// ConditionalGzipMiddleware returns gzip middleware that can be toggled at runtime
func ConditionalGzipMiddleware(config *GzipConfig) func(next http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			if config.Enabled {
				GzipMiddleware(next).ServeHTTP(w, r)
			} else {
				next.ServeHTTP(w, r)
			}
		})
	}
}

// ConditionalMinifyMiddleware returns minify middleware that can be toggled at runtime
func ConditionalMinifyMiddleware(config *MinifyConfig) func(next http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			if config.Enabled {
				MinifyMiddleware(*config)(next).ServeHTTP(w, r)
			} else {
				next.ServeHTTP(w, r)
			}
		})
	}
}
