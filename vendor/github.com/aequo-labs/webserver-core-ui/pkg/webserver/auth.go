package webserver

import (
	"sync"
	"time"
)

// AuthConfig contains configuration for authentication
type AuthConfig struct {
	// Enabled controls whether authentication is required
	Enabled bool

	// AuthMethod specifies which authentication methods are available
	// Options: "local", "oauth", "both"
	AuthMethod string

	// SessionSecret is the secret key for session encryption
	SessionSecret string

	// SessionName is the cookie name for sessions (default: "auth-session")
	SessionName string

	// SessionMaxAge is how long sessions last in seconds (default: 86400 = 24 hours)
	SessionMaxAge int

	// LocalAuth configures local username/password authentication
	LocalAuth *LocalAuthConfig

	// OAuth configures OAuth2 authentication
	OAuth *OAuthConfig

	// SetupEnabled allows first-time setup wizard
	SetupEnabled bool

	// SetupCompleteCheck is a function that returns true if setup is complete
	// If nil, setup is always considered complete
	SetupCompleteCheck func() bool

	// PublicPaths are paths that don't require authentication
	PublicPaths []string

	// LoginRedirect is where to redirect after login (default: "/")
	LoginRedirect string

	// LogoutRedirect is where to redirect after logout (default: "/login")
	LogoutRedirect string

	// UnauthorizedRedirect is where to redirect unauthorized requests (default: "/login")
	UnauthorizedRedirect string

	// UserStore is an optional interface for multi-user support
	// If nil, uses single-user mode with LocalAuth credentials
	UserStore UserStore
}

// LocalAuthConfig configures local username/password authentication
type LocalAuthConfig struct {
	// Username for single-user mode
	Username string

	// PasswordHash is the bcrypt hash of the password for single-user mode
	PasswordHash string

	// RememberMeDays is how long "remember me" sessions last (default: 30)
	RememberMeDays int
}

// OAuthConfig configures OAuth2 authentication
type OAuthConfig struct {
	// Provider is the OAuth provider name ("google" or "custom")
	Provider string

	// ClientID is the OAuth client ID
	ClientID string

	// ClientSecret is the OAuth client secret
	ClientSecret string

	// RedirectURL is the OAuth callback URL
	RedirectURL string

	// Scopes are the OAuth scopes to request
	Scopes []string

	// Custom provider configuration (for Provider="custom")
	Custom *CustomOAuthConfig
}

// CustomOAuthConfig configures a custom OAuth2 provider
type CustomOAuthConfig struct {
	// Name is the display name for the provider
	Name string

	// AuthURL is the authorization endpoint
	AuthURL string

	// TokenURL is the token endpoint
	TokenURL string

	// UserInfoURL is the user info endpoint
	UserInfoURL string

	// UserIDField is the JSON field for user ID (default: "id")
	UserIDField string

	// UserNameField is the JSON field for username (default: "name")
	UserNameField string

	// UserEmailField is the JSON field for email (default: "email")
	UserEmailField string
}

// User represents an authenticated user
type User struct {
	ID       string
	Username string
	Email    string
	Name     string
	Provider string // "local" or OAuth provider name
	Metadata map[string]interface{}
}

// UserStore is an interface for multi-user support
// Implement this interface to enable multi-user authentication
type UserStore interface {
	// GetUserByUsername retrieves a user by username
	GetUserByUsername(username string) (*User, error)

	// GetUserByEmail retrieves a user by email
	GetUserByEmail(email string) (*User, error)

	// GetUserByID retrieves a user by ID
	GetUserByID(id string) (*User, error)

	// ValidatePassword checks if the password is correct for the user
	ValidatePassword(user *User, password string) bool

	// CreateUser creates a new user
	CreateUser(user *User, passwordHash string) error

	// UpdateUser updates an existing user
	UpdateUser(user *User) error

	// GetOrCreateOAuthUser gets or creates a user from OAuth data
	GetOrCreateOAuthUser(provider string, providerUserID string, email string, name string) (*User, error)
}

// AuthStatus represents the current authentication status
type AuthStatus struct {
	Authenticated bool   `json:"authenticated"`
	Username      string `json:"username,omitempty"`
	Email         string `json:"email,omitempty"`
	Name          string `json:"name,omitempty"`
	Provider      string `json:"provider,omitempty"`
	SetupNeeded   bool   `json:"setupNeeded"`
	AuthMethod    string `json:"authMethod,omitempty"`
	OAuthEnabled  bool   `json:"oauthEnabled"`
	LocalEnabled  bool   `json:"localEnabled"`
}

// AuthManager handles authentication operations
type AuthManager struct {
	config  *AuthConfig
	mu      sync.RWMutex
	logger  authLogger
	session *SessionManager
}

// authLogger is a minimal logging interface compatible with logging.Logger
type authLogger interface {
	Debug(msg string, keysAndValues ...interface{})
	Warn(v ...interface{})
	Error(v ...interface{})
}

// NewAuthManager creates a new auth manager
func NewAuthManager(config *AuthConfig, logger authLogger) *AuthManager {
	// Set defaults
	if config.SessionName == "" {
		config.SessionName = "auth-session"
	}
	if config.SessionMaxAge == 0 {
		config.SessionMaxAge = 86400 // 24 hours
	}
	if config.LoginRedirect == "" {
		config.LoginRedirect = "/"
	}
	if config.LogoutRedirect == "" {
		config.LogoutRedirect = "/login"
	}
	if config.UnauthorizedRedirect == "" {
		config.UnauthorizedRedirect = "/login"
	}
	if config.LocalAuth != nil && config.LocalAuth.RememberMeDays == 0 {
		config.LocalAuth.RememberMeDays = 30
	}

	// Add default public paths
	defaultPublicPaths := []string{
		"/login",
		"/api/auth/login",
		"/api/auth/status",
		"/api/oauth/login",
		"/api/oauth/callback",
		"/setup",
		"/api/setup",
		"/static/",
		"/health",
	}
	config.PublicPaths = append(config.PublicPaths, defaultPublicPaths...)

	am := &AuthManager{
		config: config,
		logger: logger,
	}

	// Initialize session manager
	if config.Enabled && config.SessionSecret != "" {
		am.session = NewSessionManager(config.SessionSecret, config.SessionName, config.SessionMaxAge)
	}

	return am
}

// IsEnabled returns whether authentication is enabled
func (am *AuthManager) IsEnabled() bool {
	am.mu.RLock()
	defer am.mu.RUnlock()
	return am.config.Enabled
}

// GetConfig returns the auth configuration
func (am *AuthManager) GetConfig() *AuthConfig {
	am.mu.RLock()
	defer am.mu.RUnlock()
	return am.config
}

// GetSessionManager returns the session manager
func (am *AuthManager) GetSessionManager() *SessionManager {
	return am.session
}

// IsSetupNeeded returns whether first-time setup is needed
func (am *AuthManager) IsSetupNeeded() bool {
	am.mu.RLock()
	defer am.mu.RUnlock()

	if !am.config.SetupEnabled {
		return false
	}

	if am.config.SetupCompleteCheck != nil {
		return !am.config.SetupCompleteCheck()
	}

	// Default: setup is needed if no credentials are configured
	if am.config.LocalAuth == nil {
		return true
	}
	return am.config.LocalAuth.Username == "" || am.config.LocalAuth.PasswordHash == ""
}

// IsLocalAuthEnabled returns whether local authentication is enabled
func (am *AuthManager) IsLocalAuthEnabled() bool {
	am.mu.RLock()
	defer am.mu.RUnlock()
	return am.config.AuthMethod == "local" || am.config.AuthMethod == "both"
}

// IsOAuthEnabled returns whether OAuth authentication is enabled
func (am *AuthManager) IsOAuthEnabled() bool {
	am.mu.RLock()
	defer am.mu.RUnlock()
	return (am.config.AuthMethod == "oauth" || am.config.AuthMethod == "both") && am.config.OAuth != nil
}

// IsPublicPath checks if a path is public (doesn't require authentication)
func (am *AuthManager) IsPublicPath(path string) bool {
	am.mu.RLock()
	defer am.mu.RUnlock()

	for _, publicPath := range am.config.PublicPaths {
		if path == publicPath {
			return true
		}
		// Check prefix match for paths ending with /
		if len(publicPath) > 0 && publicPath[len(publicPath)-1] == '/' {
			if len(path) >= len(publicPath) && path[:len(publicPath)] == publicPath {
				return true
			}
		}
	}
	return false
}

// GetOAuthProvider returns the OAuth provider name
func (am *AuthManager) GetOAuthProvider() string {
	am.mu.RLock()
	defer am.mu.RUnlock()

	if am.config.OAuth == nil {
		return ""
	}
	if am.config.OAuth.Provider == "custom" && am.config.OAuth.Custom != nil {
		return am.config.OAuth.Custom.Name
	}
	return am.config.OAuth.Provider
}

// SessionData represents data stored in a session
type SessionData struct {
	Authenticated bool
	UserID        string
	Username      string
	Email         string
	Name          string
	Provider      string
	RememberMe    bool
	CreatedAt     time.Time
	ExpiresAt     time.Time
}
