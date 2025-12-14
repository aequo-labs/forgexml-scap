package middleware

import (
	"bufio"
	"bytes"
	"net"
	"net/http"
	"strings"

	"github.com/tdewolff/minify/v2"
	"github.com/tdewolff/minify/v2/html"
)

// MinifyConfig holds configuration for the minify middleware
type MinifyConfig struct {
	Enabled bool
}

// minifyResponseWriter wraps the original http.ResponseWriter
// and provides HTML minification for the response
type minifyResponseWriter struct {
	http.ResponseWriter
	buf           *bytes.Buffer
	minifier      *minify.M
	statusCode    int
	headerWritten bool
}

// WriteHeader captures the status code
func (w *minifyResponseWriter) WriteHeader(statusCode int) {
	w.statusCode = statusCode
	// Don't actually write the header yet, as we need to buffer the response
}

// Write captures the response body for minification
func (w *minifyResponseWriter) Write(b []byte) (int, error) {
	// Only buffer if we haven't written the header yet
	return w.buf.Write(b)
}

// Flush implements the http.Flusher interface
func (w *minifyResponseWriter) Flush() {
	// If the underlying ResponseWriter supports Flush, call it
	if flusher, ok := w.ResponseWriter.(http.Flusher); ok {
		flusher.Flush()
	}
}

// Hijack implements the http.Hijacker interface
func (w *minifyResponseWriter) Hijack() (net.Conn, *bufio.ReadWriter, error) {
	// If the underlying ResponseWriter supports Hijack, call it
	if hijacker, ok := w.ResponseWriter.(http.Hijacker); ok {
		return hijacker.Hijack()
	}
	return nil, nil, http.ErrNotSupported
}

// Push implements the http.Pusher interface
func (w *minifyResponseWriter) Push(target string, opts *http.PushOptions) error {
	// If the underlying ResponseWriter supports Push, call it
	if pusher, ok := w.ResponseWriter.(http.Pusher); ok {
		return pusher.Push(target, opts)
	}
	return http.ErrNotSupported
}

// Close minifies the buffered content and writes it to the original ResponseWriter
func (w *minifyResponseWriter) Close() error {
	// Set content type if not already set
	contentType := w.Header().Get("Content-Type")
	if contentType == "" {
		contentType = http.DetectContentType(w.buf.Bytes())
		w.Header().Set("Content-Type", contentType)
	}

	// Write the status code if it was set
	if w.statusCode != 0 {
		w.ResponseWriter.WriteHeader(w.statusCode)
	}

	// Only minify HTML content
	if strings.Contains(contentType, "text/html") {
		minified, err := w.minifier.String("text/html", w.buf.String())
		if err != nil {
			// If minification fails, write the original content
			_, writeErr := w.ResponseWriter.Write(w.buf.Bytes())
			return writeErr
		}
		_, err = w.ResponseWriter.Write([]byte(minified))
		return err
	} else if strings.Contains(contentType, "application/json") {
		// For JSON content, just pass through without minification
		// This is important for API responses
		_, err := w.ResponseWriter.Write(w.buf.Bytes())
		return err
	}

	// For non-HTML content, write the original content
	_, err := w.ResponseWriter.Write(w.buf.Bytes())
	return err
}

// MinifyMiddleware creates middleware that minifies HTML responses
func MinifyMiddleware(config MinifyConfig) func(http.Handler) http.Handler {
	m := minify.New()
	// Configure HTML minifier with more conservative settings
	m.Add("text/html", &html.Minifier{
		KeepDefaultAttrVals: true,
		KeepDocumentTags:    true,
		KeepEndTags:         true,
	})

	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			if !config.Enabled {
				// If minification is disabled, just pass through
				next.ServeHTTP(w, r)
				return
			}

			// Create a buffering minify writer
			mw := &minifyResponseWriter{
				ResponseWriter: w,
				buf:            &bytes.Buffer{},
				minifier:       m,
			}

			// Call the next handler with our wrapped response writer
			next.ServeHTTP(mw, r)

			// Minify and write the response
			mw.Close()
		})
	}
}