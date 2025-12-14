package middleware

import (
	"net/http"
	"time"

	"github.com/aequo-labs/webserver-core/pkg/logging"
)

// AuditConfig holds configuration for audit logging
type AuditConfig struct {
	// LogAllRequests logs every request, not just auth-related ones
	LogAllRequests bool
	// LogHeaders includes request headers in audit log (be careful with sensitive data)
	LogHeaders bool
	// SensitiveEndpoints are endpoints that should always be logged
	SensitiveEndpoints []string
}

// DefaultAuditConfig returns sensible defaults
func DefaultAuditConfig() AuditConfig {
	return AuditConfig{
		LogAllRequests: false,
		LogHeaders:     false,
		SensitiveEndpoints: []string{
			"/api/licenses/create",
			"/api/licenses/activate",
			"/api/licenses/deactivate",
			"/api/licenses/revoke",
			"/api/webhooks",
		},
	}
}

// responseWriter wraps http.ResponseWriter to capture status code
type responseWriter struct {
	http.ResponseWriter
	statusCode int
}

func (rw *responseWriter) WriteHeader(code int) {
	rw.statusCode = code
	rw.ResponseWriter.WriteHeader(code)
}

// AuditMiddleware logs security-relevant requests
func AuditMiddleware(logger *logging.Logger, config AuditConfig) func(http.Handler) http.Handler {
	// Build a map for faster lookup
	sensitiveMap := make(map[string]bool)
	for _, endpoint := range config.SensitiveEndpoints {
		sensitiveMap[endpoint] = true
	}

	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			start := time.Now()
			ip := GetClientIP(r)
			path := r.URL.Path
			method := r.Method

			// Wrap response writer to capture status code
			wrapped := &responseWriter{ResponseWriter: w, statusCode: http.StatusOK}

			// Serve the request
			next.ServeHTTP(wrapped, r)

			duration := time.Since(start)
			status := wrapped.statusCode

			// Determine if we should log this request
			shouldLog := config.LogAllRequests ||
				sensitiveMap[path] ||
				status == http.StatusUnauthorized ||
				status == http.StatusForbidden ||
				status == http.StatusTooManyRequests

			if shouldLog {
				// Build audit log entry
				userAgent := r.Header.Get("User-Agent")

				if status >= 400 {
					logger.Warnf("AUDIT: %s %s from %s - %d (%s) [%s]",
						method, path, ip, status, http.StatusText(status), duration)
				} else {
					logger.Infof("AUDIT: %s %s from %s - %d (%s) [%s]",
						method, path, ip, status, http.StatusText(status), duration)
				}

				// Log additional details for sensitive endpoints
				if sensitiveMap[path] && config.LogHeaders {
					logger.Debugf("AUDIT DETAIL: User-Agent: %s, Content-Type: %s",
						userAgent, r.Header.Get("Content-Type"))
				}
			}
		})
	}
}
