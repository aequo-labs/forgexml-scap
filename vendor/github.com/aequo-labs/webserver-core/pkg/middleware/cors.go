package middleware

import (
	"net/http"
	"strings"
)

// CORSConfig holds configuration for CORS middleware
type CORSConfig struct {
	// AllowedOrigins is a comma-separated list of allowed origins, or "*" for all
	AllowedOrigins string
	// AllowedMethods are the HTTP methods allowed for CORS requests
	AllowedMethods string
	// AllowedHeaders are the headers allowed in CORS requests
	AllowedHeaders string
	// AllowCredentials indicates if credentials (cookies, auth) are allowed
	AllowCredentials bool
	// MaxAge is the max time (seconds) to cache preflight responses (default: 86400 = 24h)
	MaxAge string
}

// DefaultCORSConfig returns sensible defaults for CORS
func DefaultCORSConfig() CORSConfig {
	return CORSConfig{
		AllowedOrigins:   "http://localhost:3000,http://localhost:3001,http://localhost:5173,http://localhost:5174",
		AllowedMethods:   "GET, POST, PUT, DELETE, OPTIONS, PATCH",
		AllowedHeaders:   "Accept, Content-Type, Content-Length, Accept-Encoding, Authorization, X-CSRF-Token",
		AllowCredentials: true,
		MaxAge:           "86400", // 24 hours
	}
}

// CORSMiddleware returns a middleware function that handles CORS headers
func CORSMiddleware(config CORSConfig) func(http.Handler) http.Handler {
	// Parse allowed origins once at setup time
	allowedOriginsMap := make(map[string]bool)
	allowAll := false
	for _, origin := range strings.Split(config.AllowedOrigins, ",") {
		origin = strings.TrimSpace(origin)
		if origin == "*" {
			allowAll = true
		}
		allowedOriginsMap[origin] = true
	}

	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			// Get the origin from the request
			origin := r.Header.Get("Origin")

			// Check if the origin is allowed
			isAllowed := false
			if origin != "" {
				if allowAll || allowedOriginsMap[origin] {
					isAllowed = true
				}
			}

			// Set CORS headers if origin is allowed
			if isAllowed {
				w.Header().Set("Access-Control-Allow-Origin", origin)
				if config.AllowCredentials {
					w.Header().Set("Access-Control-Allow-Credentials", "true")
				}
			}

			// Always set these headers for preflight support
			w.Header().Set("Access-Control-Allow-Methods", config.AllowedMethods)
			w.Header().Set("Access-Control-Allow-Headers", config.AllowedHeaders)
			w.Header().Set("Access-Control-Max-Age", config.MaxAge)

			// Handle preflight requests
			if r.Method == "OPTIONS" {
				w.WriteHeader(http.StatusNoContent)
				return
			}

			next.ServeHTTP(w, r)
		})
	}
}

// SimpleCORSMiddleware creates a CORS middleware with the specified allowed origins
// Use "*" to allow all origins (not recommended for production with credentials)
func SimpleCORSMiddleware(allowedOrigins string) func(http.Handler) http.Handler {
	config := DefaultCORSConfig()
	config.AllowedOrigins = allowedOrigins
	return CORSMiddleware(config)
}
