package middleware

import (
	"crypto/subtle"
	"encoding/base64"
	"net/http"
	"strings"
)

// BasicAuthConfig holds configuration for basic authentication
type BasicAuthConfig struct {
	Username      string
	Password      string
	Realm         string
	SkipPaths     []string              // Paths to skip authentication
	SkipPrefixes  []string              // Path prefixes to skip authentication
	OnAuthFailure func(r *http.Request) // Called when authentication fails (for rate limiting/lockout)
	OnAuthSuccess func(r *http.Request) // Called when authentication succeeds (to reset failure counters)
}

// BasicAuthMiddleware provides HTTP Basic Authentication
func BasicAuthMiddleware(config BasicAuthConfig) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			// Check if this path should skip authentication
			if shouldSkipAuth(r.URL.Path, config.SkipPaths, config.SkipPrefixes) {
				next.ServeHTTP(w, r)
				return
			}

			// Get the Authorization header
			auth := r.Header.Get("Authorization")
			if auth == "" {
				requireAuth(w, config.Realm)
				return
			}

			// Check if it's Basic auth
			if !strings.HasPrefix(auth, "Basic ") {
				requireAuth(w, config.Realm)
				return
			}

			// Decode the base64 encoded credentials
			payload, err := base64.StdEncoding.DecodeString(auth[6:])
			if err != nil {
				requireAuth(w, config.Realm)
				return
			}

			// Split username and password
			pair := strings.SplitN(string(payload), ":", 2)
			if len(pair) != 2 {
				requireAuth(w, config.Realm)
				return
			}

			username := pair[0]
			password := pair[1]

			// Use constant-time comparison to prevent timing attacks
			if subtle.ConstantTimeCompare([]byte(username), []byte(config.Username)) == 1 &&
				subtle.ConstantTimeCompare([]byte(password), []byte(config.Password)) == 1 {
				// Authentication successful - call success callback if set
				if config.OnAuthSuccess != nil {
					config.OnAuthSuccess(r)
				}
				next.ServeHTTP(w, r)
				return
			}

			// Authentication failed - call failure callback if set
			if config.OnAuthFailure != nil {
				config.OnAuthFailure(r)
			}
			requireAuth(w, config.Realm)
		})
	}
}

// shouldSkipAuth checks if the given path should skip authentication
func shouldSkipAuth(path string, skipPaths, skipPrefixes []string) bool {
	// Check exact paths
	for _, skipPath := range skipPaths {
		if path == skipPath {
			return true
		}
	}

	// Check prefixes
	for _, prefix := range skipPrefixes {
		if strings.HasPrefix(path, prefix) {
			return true
		}
	}

	return false
}

// requireAuth sends a 401 Unauthorized response with WWW-Authenticate header
func requireAuth(w http.ResponseWriter, realm string) {
	if realm == "" {
		realm = "Protected Area"
	}
	w.Header().Set("WWW-Authenticate", `Basic realm="`+realm+`"`)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusUnauthorized)
	w.Write([]byte(`{"success": false, "error": "Authentication required", "code": 401}`))
}

// APIKeyAuthConfig holds configuration for API key authentication
type APIKeyAuthConfig struct {
	APIKey       string
	HeaderName   string   // Header name to look for API key (default: "X-API-Key")
	SkipPaths    []string // Paths to skip authentication
	SkipPrefixes []string // Path prefixes to skip authentication
}

// APIKeyAuthMiddleware provides API Key authentication
func APIKeyAuthMiddleware(config APIKeyAuthConfig) func(http.Handler) http.Handler {
	headerName := config.HeaderName
	if headerName == "" {
		headerName = "X-API-Key"
	}

	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			// Check if this path should skip authentication
			if shouldSkipAuth(r.URL.Path, config.SkipPaths, config.SkipPrefixes) {
				next.ServeHTTP(w, r)
				return
			}

			// Get the API key from header
			apiKey := r.Header.Get(headerName)
			if apiKey == "" {
				requireAPIKey(w, headerName)
				return
			}

			// Use constant-time comparison to prevent timing attacks
			if subtle.ConstantTimeCompare([]byte(apiKey), []byte(config.APIKey)) == 1 {
				// Authentication successful
				next.ServeHTTP(w, r)
				return
			}

			// Authentication failed
			requireAPIKey(w, headerName)
		})
	}
}

// requireAPIKey sends a 401 Unauthorized response for API key auth
func requireAPIKey(w http.ResponseWriter, headerName string) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusUnauthorized)
	w.Write([]byte(`{"success": false, "error": "Valid API key required in ` + headerName + ` header", "code": 401}`))
}
