package assets

import (
	"embed"
	"fmt"
	"io"
	"io/fs"
	"net/http"
	"path/filepath"
	"strings"
)

// Manager handles static asset management and serving
type Manager struct {
	staticFS  fs.FS
	urlPrefix string
	cacheTime int
	embedFS   *embed.FS
}

// Config holds configuration for the asset manager
type Config struct {
	URLPrefix string
	CacheTime int // Cache time in seconds
	StaticDir string
}

// NewManager creates a new asset manager
func NewManager(config Config) *Manager {
	return &Manager{
		urlPrefix: config.URLPrefix,
		cacheTime: config.CacheTime,
	}
}

// NewManagerWithEmbed creates a new asset manager with embedded files
func NewManagerWithEmbed(config Config, embedFS *embed.FS) *Manager {
	return &Manager{
		urlPrefix: config.URLPrefix,
		cacheTime: config.CacheTime,
		embedFS:   embedFS,
	}
}

// SetFileSystem sets the filesystem for serving assets
func (m *Manager) SetFileSystem(fs fs.FS) {
	m.staticFS = fs
}

// ServeHTTP serves static assets
func (m *Manager) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	// Remove the URL prefix from the request path
	path := strings.TrimPrefix(r.URL.Path, m.urlPrefix)
	if path == "" {
		path = "index.html"
	}

	// Clean the path to prevent directory traversal
	path = filepath.Clean(path)
	if strings.HasPrefix(path, "..") {
		http.Error(w, "Invalid path", http.StatusBadRequest)
		return
	}

	// Determine which filesystem to use
	var fileSystem fs.FS
	if m.embedFS != nil {
		fileSystem = m.embedFS
	} else if m.staticFS != nil {
		fileSystem = m.staticFS
	} else {
		http.Error(w, "No filesystem configured", http.StatusInternalServerError)
		return
	}

	// Try to open the file
	file, err := fileSystem.Open(path)
	if err != nil {
		http.NotFound(w, r)
		return
	}
	defer file.Close()

	// Get file info
	stat, err := file.Stat()
	if err != nil {
		http.Error(w, "File stat error", http.StatusInternalServerError)
		return
	}

	// Set content type based on file extension
	m.setContentType(w, path)

	// Set cache headers if cache time is configured
	if m.cacheTime > 0 {
		w.Header().Set("Cache-Control", fmt.Sprintf("public, max-age=%d", m.cacheTime))
	}

	// Serve the file
	http.ServeContent(w, r, stat.Name(), stat.ModTime(), file.(io.ReadSeeker))
}

// setContentType sets the appropriate content type based on file extension
func (m *Manager) setContentType(w http.ResponseWriter, path string) {
	ext := strings.ToLower(filepath.Ext(path))

	switch ext {
	case ".css":
		w.Header().Set("Content-Type", "text/css; charset=utf-8")
	case ".js":
		w.Header().Set("Content-Type", "application/javascript; charset=utf-8")
	case ".json":
		w.Header().Set("Content-Type", "application/json; charset=utf-8")
	case ".png":
		w.Header().Set("Content-Type", "image/png")
	case ".jpg", ".jpeg":
		w.Header().Set("Content-Type", "image/jpeg")
	case ".gif":
		w.Header().Set("Content-Type", "image/gif")
	case ".svg":
		w.Header().Set("Content-Type", "image/svg+xml")
	case ".ico":
		w.Header().Set("Content-Type", "image/x-icon")
	case ".woff":
		w.Header().Set("Content-Type", "font/woff")
	case ".woff2":
		w.Header().Set("Content-Type", "font/woff2")
	case ".ttf":
		w.Header().Set("Content-Type", "font/ttf")
	case ".eot":
		w.Header().Set("Content-Type", "application/vnd.ms-fontobject")
	default:
		w.Header().Set("Content-Type", "application/octet-stream")
	}
}

// GetAssetURL returns the full URL for an asset
func (m *Manager) GetAssetURL(path string) string {
	return m.urlPrefix + "/" + strings.TrimPrefix(path, "/")
}
