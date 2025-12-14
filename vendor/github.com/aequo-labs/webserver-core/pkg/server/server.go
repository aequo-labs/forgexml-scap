package server

import (
	"crypto/tls"
	"net/http"

	"github.com/aequo-labs/webserver-core/pkg/https"
	"github.com/aequo-labs/webserver-core/pkg/logging"
	"github.com/aequo-labs/webserver-core/pkg/middleware"
)

// Server represents the base HTTP server with core functionality
type Server struct {
	router           *http.ServeMux
	middlewares      []func(http.Handler) http.Handler
	middlewareConfig *middleware.MiddlewareConfig
	httpsConfig      *https.Config
	logger           *logging.Logger
}

// New creates a new base server instance
func New() *Server {
	logger, _ := logging.NewLogger(logging.DefaultConfig())
	return &Server{
		router:           http.NewServeMux(),
		middlewares:      make([]func(http.Handler) http.Handler, 0),
		middlewareConfig: middleware.NewMiddlewareConfig(),
		httpsConfig:      https.NewConfig(),
		logger:           logger,
	}
}

// NewWithConfig creates a new server with custom configurations
func NewWithConfig(middlewareConfig *middleware.MiddlewareConfig, httpsConfig *https.Config) *Server {
	logger, _ := logging.NewLogger(logging.DefaultConfig())
	s := &Server{
		router:      http.NewServeMux(),
		middlewares: make([]func(http.Handler) http.Handler, 0),
		httpsConfig: httpsConfig,
		logger:      logger,
	}

	if middlewareConfig != nil {
		s.middlewareConfig = middlewareConfig
	} else {
		s.middlewareConfig = middleware.NewMiddlewareConfig()
	}

	if httpsConfig == nil {
		s.httpsConfig = https.NewConfig()
	}

	return s
}

// Router returns the underlying ServeMux router for adding routes
func (s *Server) Router() *http.ServeMux {
	return s.router
}

// MiddlewareConfig returns the current middleware configuration
func (s *Server) MiddlewareConfig() *middleware.MiddlewareConfig {
	return s.middlewareConfig
}

// HTTPSConfig returns the current HTTPS configuration
func (s *Server) HTTPSConfig() *https.Config {
	return s.httpsConfig
}

// SetLogger sets a custom logger for the server
func (s *Server) SetLogger(logger *logging.Logger) {
	s.logger = logger
}

// Logger returns the server's logger
func (s *Server) Logger() *logging.Logger {
	return s.logger
}

// SetupOptionalMiddleware sets up the conditional middleware based on configuration
func (s *Server) SetupOptionalMiddleware() {
	// Add gzip compression middleware first (conditional based on config)
	s.middlewares = append(s.middlewares, middleware.ConditionalGzipMiddleware(&s.middlewareConfig.Gzip))

	// Add minification middleware (conditional based on config)
	// Applied after gzip wrapper so minified content gets compressed
	s.middlewares = append(s.middlewares, middleware.ConditionalMinifyMiddleware(&s.middlewareConfig.Minification))
}

// UseMiddleware adds a middleware to the chain
func (s *Server) UseMiddleware(mw func(http.Handler) http.Handler) {
	s.middlewares = append(s.middlewares, mw)
}

// Use is an alias for UseMiddleware for compatibility
func (s *Server) Use(mw func(http.Handler) http.Handler) {
	s.UseMiddleware(mw)
}

// Handler returns the final handler with all middleware applied
func (s *Server) Handler() http.Handler {
	var handler http.Handler = s.router

	// Apply middleware in reverse order so first added is outermost
	for i := len(s.middlewares) - 1; i >= 0; i-- {
		handler = s.middlewares[i](handler)
	}

	return handler
}

// EnableMinification enables HTML minification
func (s *Server) EnableMinification() {
	s.middlewareConfig.Minification.Enabled = true
}

// DisableMinification disables HTML minification
func (s *Server) DisableMinification() {
	s.middlewareConfig.Minification.Enabled = false
}

// ToggleMinification toggles HTML minification on/off
func (s *Server) ToggleMinification() bool {
	s.middlewareConfig.Minification.Enabled = !s.middlewareConfig.Minification.Enabled
	return s.middlewareConfig.Minification.Enabled
}

// IsMinificationEnabled returns whether HTML minification is enabled
func (s *Server) IsMinificationEnabled() bool {
	return s.middlewareConfig.Minification.Enabled
}

// EnableGzip enables gzip compression
func (s *Server) EnableGzip() {
	s.middlewareConfig.Gzip.Enabled = true
}

// DisableGzip disables gzip compression
func (s *Server) DisableGzip() {
	s.middlewareConfig.Gzip.Enabled = false
}

// ToggleGzip toggles gzip compression on/off
func (s *Server) ToggleGzip() bool {
	s.middlewareConfig.Gzip.Enabled = !s.middlewareConfig.Gzip.Enabled
	return s.middlewareConfig.Gzip.Enabled
}

// IsGzipEnabled returns whether gzip compression is enabled
func (s *Server) IsGzipEnabled() bool {
	return s.middlewareConfig.Gzip.Enabled
}

// ListenAndServe starts the HTTP server on the specified address
func (s *Server) ListenAndServe(addr string) error {
	s.logger.Info("Server starting", "addr", addr)
	return http.ListenAndServe(addr, s.Handler())
}

// ListenAndServeTLS starts the HTTPS server on the specified address
func (s *Server) ListenAndServeTLS(addr string) error {
	// Ensure certificates exist
	if err := s.httpsConfig.EnsureCertificates(); err != nil {
		return err
	}

	// Load TLS configuration
	tlsConfig, err := s.httpsConfig.LoadTLSConfig()
	if err != nil {
		return err
	}

	server := &http.Server{
		Addr:      addr,
		Handler:   s.Handler(),
		TLSConfig: tlsConfig,
	}

	s.logger.Info("HTTPS Server starting", "addr", addr)
	return server.ListenAndServeTLS(s.httpsConfig.CertFile, s.httpsConfig.KeyFile)
}

// Serve starts the server with automatic HTTPS redirect if configured
func (s *Server) Serve(httpAddr, httpsAddr string) error {
	if s.httpsConfig.Enabled {
		// Start HTTP server that redirects to HTTPS
		go func() {
			redirectServer := &http.Server{
				Addr: httpAddr,
				Handler: http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
					http.Redirect(w, r, "https://"+r.Host+r.RequestURI, http.StatusMovedPermanently)
				}),
			}
			s.logger.Info("HTTP Redirect server starting", "addr", httpAddr)
			if err := redirectServer.ListenAndServe(); err != nil {
				s.logger.Error("HTTP Redirect server error", "error", err)
			}
		}()

		// Start HTTPS server
		return s.ListenAndServeTLS(httpsAddr)
	}

	// Start HTTP server only
	return s.ListenAndServe(httpAddr)
}

// ServeWithTLS runs the server with a custom TLS configuration
func (s *Server) ServeWithTLS(addr string, tlsConfig *tls.Config) error {
	server := &http.Server{
		Addr:      addr,
		Handler:   s.Handler(),
		TLSConfig: tlsConfig,
	}

	s.logger.Info("Custom TLS Server starting", "addr", addr)
	return server.ListenAndServeTLS("", "")
}

// HandleFunc registers a handler function for the given pattern.
// Pattern format for Go 1.22+: "METHOD /path" or "METHOD /path/{param}"
// Examples:
//   - "GET /api/users"
//   - "POST /api/users"
//   - "GET /api/users/{id}"
//   - "/static/" (matches all methods)
func (s *Server) HandleFunc(pattern string, handler func(http.ResponseWriter, *http.Request)) {
	s.router.HandleFunc(pattern, handler)
}

// Handle registers a handler for the given pattern.
func (s *Server) Handle(pattern string, handler http.Handler) {
	s.router.Handle(pattern, handler)
}
