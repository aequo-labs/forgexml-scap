package middleware

import (
	"net/http"
	"strings"
)

// Body size limits
const (
	MaxBodySize       = 1 << 20   // 1MB for normal API requests
	MaxUploadBodySize = 100 << 20 // 100MB for file uploads
)

// BodyLimitConfig holds configuration for body size limiting
type BodyLimitConfig struct {
	DefaultMaxBytes int64    // Default max body size
	UploadMaxBytes  int64    // Max body size for upload endpoints
	UploadPrefixes  []string // URL prefixes that use upload limit
}

// DefaultBodyLimitConfig returns sensible defaults
func DefaultBodyLimitConfig() BodyLimitConfig {
	return BodyLimitConfig{
		DefaultMaxBytes: MaxBodySize,
		UploadMaxBytes:  MaxUploadBodySize,
		UploadPrefixes: []string{
			"/api/admin/products/", // Product version uploads
		},
	}
}

// BodyLimitMiddleware limits the size of incoming request bodies
// with different limits for upload endpoints
func BodyLimitMiddleware(config BodyLimitConfig) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			// Skip for requests without bodies
			if r.Body == nil {
				next.ServeHTTP(w, r)
				return
			}

			// Determine max body size based on path
			maxBytes := config.DefaultMaxBytes
			path := r.URL.Path

			for _, prefix := range config.UploadPrefixes {
				if strings.HasPrefix(path, prefix) {
					maxBytes = config.UploadMaxBytes
					break
				}
			}

			// Wrap the body with a max bytes reader
			r.Body = http.MaxBytesReader(w, r.Body, maxBytes)

			next.ServeHTTP(w, r)
		})
	}
}
