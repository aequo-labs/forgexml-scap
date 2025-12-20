package middleware

import (
	"compress/gzip"
	"io"
	"net/http"
	"strings"
)

// List of content types to compress
var compressibleTypes = map[string]bool{
	"text/html":                true,
	"text/css":                 true,
	"text/javascript":          true,
	"application/javascript":   true,
	"application/json":         true,
	"application/x-javascript": true,
	"text/plain":               true, // Added for plain text
	"image/jpeg":               true, // Added for JPEG images
}

// gzipResponseWriter wraps the original http.ResponseWriter
// and provides gzip compression for the response
type gzipResponseWriter struct {
	io.Writer
	http.ResponseWriter
	sniffDone bool
}

// Write implements the http.ResponseWriter interface
func (w *gzipResponseWriter) Write(b []byte) (int, error) {
	// If content type is not set, infer it from the data
	if !w.sniffDone {
		w.sniffDone = true
		if w.Header().Get("Content-Type") == "" {
			w.Header().Set("Content-Type", http.DetectContentType(b))
		}
	}
	return w.Writer.Write(b)
}

// WriteHeader implements the http.ResponseWriter interface
func (w *gzipResponseWriter) WriteHeader(statusCode int) {
	w.ResponseWriter.WriteHeader(statusCode)
}

// Flush implements the http.Flusher interface
func (w *gzipResponseWriter) Flush() {
	// Implement Flush for http.Flusher
	if f, ok := w.ResponseWriter.(http.Flusher); ok {
		f.Flush()
	}
	// Also flush the gzip writer
	if gzw, ok := w.Writer.(*gzip.Writer); ok {
		gzw.Flush()
	}
}

// shouldCompress checks if the content type should be compressed
func shouldCompress(contentType string) bool {
	// Extract the base content type (remove charset, etc.)
	if idx := strings.IndexByte(contentType, ';'); idx >= 0 {
		contentType = contentType[:idx]
	}
	contentType = strings.TrimSpace(contentType)
	return compressibleTypes[contentType]
}

// contentTypeChecker is a ResponseWriter wrapper that checks content type
type contentTypeChecker struct {
	http.ResponseWriter
	headerChecked  bool
	shouldCompress bool
	writer         io.Writer
}

// WriteHeader checks the content type before writing the header
func (c *contentTypeChecker) WriteHeader(statusCode int) {
	if !c.headerChecked {
		c.headerChecked = true
		contentType := c.Header().Get("Content-Type")
		c.shouldCompress = shouldCompress(contentType)

		if c.shouldCompress {
			c.Header().Set("Content-Encoding", "gzip")
			c.Header().Del("Content-Length")
		}
	}
	c.ResponseWriter.WriteHeader(statusCode)
}

// Write checks if the response should be compressed
func (c *contentTypeChecker) Write(b []byte) (int, error) {
	if !c.headerChecked {
		c.headerChecked = true
		if c.Header().Get("Content-Type") == "" {
			contentType := http.DetectContentType(b)
			c.Header().Set("Content-Type", contentType)
			c.shouldCompress = shouldCompress(contentType)
		} else {
			contentType := c.Header().Get("Content-Type")
			c.shouldCompress = shouldCompress(contentType)
		}

		if c.shouldCompress {
			c.Header().Set("Content-Encoding", "gzip")
			c.Header().Del("Content-Length")
		}
	}

	if c.shouldCompress {
		return c.writer.Write(b)
	}
	return c.ResponseWriter.Write(b)
}

// Flush implements the http.Flusher interface
func (c *contentTypeChecker) Flush() {
	if f, ok := c.ResponseWriter.(http.Flusher); ok {
		f.Flush()
	}

	if c.shouldCompress {
		if gzw, ok := c.writer.(*gzip.Writer); ok {
			gzw.Flush()
		}
	}
}

// GzipMiddleware compresses HTTP responses for clients that support it
// using the highest compression level (9)
func GzipMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Check if client accepts gzip encoding
		if !strings.Contains(r.Header.Get("Accept-Encoding"), "gzip") {
			next.ServeHTTP(w, r)
			return
		}

		// Create gzip writer with maximum compression level
		gz, err := gzip.NewWriterLevel(w, gzip.BestCompression) // Level 9
		if err != nil {
			next.ServeHTTP(w, r)
			return
		}
		defer gz.Close()

		// Set appropriate headers for compression
		w.Header().Add("Vary", "Accept-Encoding, Content-Type")

		// Create a content type checker wrapper
		checker := &contentTypeChecker{
			ResponseWriter: w,
			headerChecked:  false,
			shouldCompress: false,
			writer:         gz,
		}

		// Call the next handler with our wrapped response writer
		next.ServeHTTP(checker, r)
	})
}
