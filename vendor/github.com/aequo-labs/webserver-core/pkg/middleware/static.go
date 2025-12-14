package middleware

import (
	"io/fs"
	"net/http"
	"path/filepath"
	"strings"
)

// StaticFileHandler creates an HTTP handler for serving static files with proper MIME types
func StaticFileHandler(fileSystem fs.FS) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Get file extension and set appropriate Content-Type
		ext := strings.ToLower(filepath.Ext(r.URL.Path))
		switch ext {
		case ".css":
			w.Header().Set("Content-Type", "text/css; charset=utf-8")
		case ".js":
			w.Header().Set("Content-Type", "application/javascript; charset=utf-8")
		case ".woff2":
			w.Header().Set("Content-Type", "font/woff2")
		case ".woff":
			w.Header().Set("Content-Type", "font/woff")
		case ".ttf":
			w.Header().Set("Content-Type", "font/ttf")
		case ".png":
			w.Header().Set("Content-Type", "image/png")
		case ".jpg", ".jpeg":
			w.Header().Set("Content-Type", "image/jpeg")
		case ".svg":
			w.Header().Set("Content-Type", "image/svg+xml")
		case ".ico":
			w.Header().Set("Content-Type", "image/x-icon")
		}
		
		// Serve the file using the standard file server
		http.FileServer(http.FS(fileSystem)).ServeHTTP(w, r)
	})
}