package webserver

import (
	"context"
	"net/http"
	"strings"
)

// authContextKey is a custom type for context keys to avoid collisions
type authContextKey string

const (
	// ContextKeyUser is the context key for the authenticated user
	ContextKeyUser authContextKey = "authUser"
	// ContextKeySessionData is the context key for session data
	ContextKeySessionData authContextKey = "sessionData"
)

// AuthMiddleware creates middleware that enforces authentication
func (am *AuthManager) AuthMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// If auth is not enabled, pass through
		if !am.IsEnabled() {
			next.ServeHTTP(w, r)
			return
		}

		// Check if this is a public path
		if am.IsPublicPath(r.URL.Path) {
			next.ServeHTTP(w, r)
			return
		}

		// Check if setup is needed
		if am.IsSetupNeeded() {
			http.Redirect(w, r, "/setup", http.StatusSeeOther)
			return
		}

		// Check authentication
		if am.session == nil {
			// No session manager, can't authenticate
			http.Redirect(w, r, am.config.UnauthorizedRedirect, http.StatusSeeOther)
			return
		}

		sessionData, authenticated := am.session.GetAuthenticatedUser(r)
		if !authenticated {
			// Not authenticated - redirect based on auth method
			redirectURL := am.getLoginRedirect()
			http.Redirect(w, r, redirectURL, http.StatusSeeOther)
			return
		}

		// Add session data to context
		ctx := context.WithValue(r.Context(), ContextKeySessionData, sessionData)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}

// SessionMiddleware adds session data to context for all requests (even unauthenticated)
// This is useful for templates that need to show auth status
func (am *AuthManager) SessionMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// If auth is not enabled or no session manager, pass through
		if !am.IsEnabled() || am.session == nil {
			next.ServeHTTP(w, r)
			return
		}

		// Try to get session data (may be nil if not authenticated)
		sessionData, _ := am.session.GetSessionData(r)
		if sessionData != nil {
			ctx := context.WithValue(r.Context(), ContextKeySessionData, sessionData)
			r = r.WithContext(ctx)
		}

		next.ServeHTTP(w, r)
	})
}

// RequireAuth is a middleware that requires authentication for specific handlers
// Use this for individual routes instead of global middleware
func (am *AuthManager) RequireAuth(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// If auth is not enabled, pass through
		if !am.IsEnabled() {
			next(w, r)
			return
		}

		// Check authentication
		if am.session == nil || !am.session.IsAuthenticated(r) {
			// Check if this is an API request
			if strings.HasPrefix(r.URL.Path, "/api/") {
				w.Header().Set("Content-Type", "application/json")
				w.WriteHeader(http.StatusUnauthorized)
				w.Write([]byte(`{"error":"unauthorized","message":"Authentication required"}`))
				return
			}
			// Regular request - redirect to login
			http.Redirect(w, r, am.getLoginRedirect(), http.StatusSeeOther)
			return
		}

		// Add session data to context
		sessionData, _ := am.session.GetAuthenticatedUser(r)
		if sessionData != nil {
			ctx := context.WithValue(r.Context(), ContextKeySessionData, sessionData)
			r = r.WithContext(ctx)
		}

		next(w, r)
	}
}

// getLoginRedirect returns the appropriate login redirect URL
func (am *AuthManager) getLoginRedirect() string {
	switch am.config.AuthMethod {
	case "oauth":
		return "/api/oauth/login"
	case "local":
		return "/login"
	case "both":
		return "/login" // Show login page with both options
	default:
		return "/login"
	}
}

// GetSessionDataFromContext retrieves session data from the request context
func GetSessionDataFromContext(ctx context.Context) *SessionData {
	if data, ok := ctx.Value(ContextKeySessionData).(*SessionData); ok {
		return data
	}
	return nil
}

// GetUsernameFromContext retrieves the username from the request context
func GetUsernameFromContext(ctx context.Context) string {
	if data := GetSessionDataFromContext(ctx); data != nil {
		return data.Username
	}
	return ""
}

// GetUserEmailFromContext retrieves the user email from the request context
func GetUserEmailFromContext(ctx context.Context) string {
	if data := GetSessionDataFromContext(ctx); data != nil {
		return data.Email
	}
	return ""
}

// IsAuthenticatedFromContext checks if the request is authenticated from context
func IsAuthenticatedFromContext(ctx context.Context) bool {
	if data := GetSessionDataFromContext(ctx); data != nil {
		return data.Authenticated
	}
	return false
}
