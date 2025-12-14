package middleware

import (
	"net/http"
	"net/url"
	"strings"
)

// DecodePathValue decodes a URL path parameter value, handling special characters
// like "+" which have special meaning in URLs. Use this when retrieving path values
// from r.PathValue() that may contain encoded characters.
func DecodePathValue(value string) string {
	if value == "" {
		return value
	}

	// First replace any "+" with "%2B" to ensure they're properly handled
	// This is needed because "+" has special meaning in URLs (space)
	v := strings.ReplaceAll(value, "+", "%2B")

	// Then use url.PathUnescape to handle all other encoded characters
	decoded, err := url.PathUnescape(v)
	if err != nil {
		return value // Keep original on error
	}
	return decoded
}

// GetDecodedPathValue is a convenience function that gets a path value from the request
// and decodes it in one call. Use this instead of r.PathValue() when the value may
// contain URL-encoded characters.
func GetDecodedPathValue(r *http.Request, name string) string {
	return DecodePathValue(r.PathValue(name))
}

// DecodeURLParamMiddleware is deprecated - with http.ServeMux, use GetDecodedPathValue()
// or DecodePathValue() directly in handlers instead.
// This middleware is kept for backward compatibility but does nothing with http.ServeMux.
func DecodeURLParamMiddleware(next http.Handler) http.Handler {
	// With http.ServeMux, path values are accessed via r.PathValue() and cannot be
	// modified via middleware. Use GetDecodedPathValue() in handlers instead.
	return next
}
