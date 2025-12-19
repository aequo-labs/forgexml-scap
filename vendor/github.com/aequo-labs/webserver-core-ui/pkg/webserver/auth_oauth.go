package webserver

import (
	"context"
	"crypto/rand"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"time"

	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
)

// oauthStateKey is the session key for OAuth state
const oauthStateKey = "oauth_state"

// getOAuthConfig returns the OAuth2 config for the configured provider
func (s *UIServer) getOAuthConfig() *oauth2.Config {
	if s.authManager == nil || s.authManager.config.OAuth == nil {
		return nil
	}

	oauthCfg := s.authManager.config.OAuth

	switch oauthCfg.Provider {
	case "google":
		return &oauth2.Config{
			ClientID:     oauthCfg.ClientID,
			ClientSecret: oauthCfg.ClientSecret,
			RedirectURL:  oauthCfg.RedirectURL,
			Scopes:       oauthCfg.Scopes,
			Endpoint:     google.Endpoint,
		}
	case "custom":
		if oauthCfg.Custom == nil {
			return nil
		}
		return &oauth2.Config{
			ClientID:     oauthCfg.ClientID,
			ClientSecret: oauthCfg.ClientSecret,
			RedirectURL:  oauthCfg.RedirectURL,
			Scopes:       oauthCfg.Scopes,
			Endpoint: oauth2.Endpoint{
				AuthURL:  oauthCfg.Custom.AuthURL,
				TokenURL: oauthCfg.Custom.TokenURL,
			},
		}
	default:
		return nil
	}
}

// generateOAuthState creates a random state string for CSRF protection
func generateOAuthState() string {
	b := make([]byte, 32)
	rand.Read(b)
	return base64.URLEncoding.EncodeToString(b)
}

// RegisterOAuthRoutes registers OAuth routes - call this after SetAuthManager if OAuth is enabled
func (s *UIServer) RegisterOAuthRoutes() {
	if s.authManager == nil || !s.authManager.IsOAuthEnabled() {
		return
	}

	s.Router().HandleFunc("GET /api/oauth/login", s.oauthLoginHandler)
	s.Router().HandleFunc("GET /api/oauth/callback", s.oauthCallbackHandler)

	s.logger.Debug("OAuth routes registered", "provider", s.authManager.GetOAuthProvider())
}

// oauthLoginHandler initiates the OAuth flow
func (s *UIServer) oauthLoginHandler(w http.ResponseWriter, r *http.Request) {
	if s.authManager == nil || !s.authManager.IsOAuthEnabled() {
		http.Error(w, "OAuth not configured", http.StatusServiceUnavailable)
		return
	}

	oauthConfig := s.getOAuthConfig()
	if oauthConfig == nil {
		http.Error(w, "OAuth configuration error", http.StatusInternalServerError)
		return
	}

	// Generate and store state for CSRF protection
	state := generateOAuthState()

	// Store state in session
	if s.authManager.session != nil {
		session, err := s.authManager.session.Get(r)
		if err == nil {
			session.Values[oauthStateKey] = state
			s.authManager.session.Save(r, w, session)
		}
	}

	// Redirect to OAuth provider
	url := oauthConfig.AuthCodeURL(state)
	http.Redirect(w, r, url, http.StatusTemporaryRedirect)
}

// oauthCallbackHandler handles the OAuth callback
func (s *UIServer) oauthCallbackHandler(w http.ResponseWriter, r *http.Request) {
	if s.authManager == nil || !s.authManager.IsOAuthEnabled() {
		http.Error(w, "OAuth not configured", http.StatusServiceUnavailable)
		return
	}

	// Verify state
	state := r.URL.Query().Get("state")
	if s.authManager.session != nil {
		session, err := s.authManager.session.Get(r)
		if err != nil {
			s.logger.Error("Failed to get session for OAuth callback:", err)
			http.Redirect(w, r, "/login?error=session_error", http.StatusSeeOther)
			return
		}

		expectedState, ok := session.Values[oauthStateKey].(string)
		if !ok || state != expectedState {
			s.logger.Warn("OAuth state mismatch - expected:", expectedState, "got:", state)
			http.Redirect(w, r, "/login?error=invalid_state", http.StatusSeeOther)
			return
		}

		// Clear the state
		delete(session.Values, oauthStateKey)
		s.authManager.session.Save(r, w, session)
	}

	// Check for errors from OAuth provider
	if errMsg := r.URL.Query().Get("error"); errMsg != "" {
		s.logger.Warn("OAuth error from provider:", errMsg, r.URL.Query().Get("error_description"))
		http.Redirect(w, r, "/login?error=oauth_denied", http.StatusSeeOther)
		return
	}

	// Exchange code for token
	code := r.URL.Query().Get("code")
	if code == "" {
		http.Redirect(w, r, "/login?error=no_code", http.StatusSeeOther)
		return
	}

	oauthConfig := s.getOAuthConfig()
	if oauthConfig == nil {
		http.Error(w, "OAuth configuration error", http.StatusInternalServerError)
		return
	}

	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	token, err := oauthConfig.Exchange(ctx, code)
	if err != nil {
		s.logger.Error("OAuth code exchange failed:", err)
		http.Redirect(w, r, "/login?error=exchange_failed", http.StatusSeeOther)
		return
	}

	// Get user info
	userInfo, err := s.getOAuthUserInfo(ctx, token)
	if err != nil {
		s.logger.Error("Failed to get OAuth user info:", err)
		http.Redirect(w, r, "/login?error=userinfo_failed", http.StatusSeeOther)
		return
	}

	// Create or get user
	var user *User
	provider := s.authManager.config.OAuth.Provider
	if s.authManager.config.OAuth.Custom != nil {
		provider = s.authManager.config.OAuth.Custom.Name
	}

	if s.authManager.config.UserStore != nil {
		// Multi-user mode - get or create user
		user, err = s.authManager.config.UserStore.GetOrCreateOAuthUser(
			provider,
			userInfo.ID,
			userInfo.Email,
			userInfo.Name,
		)
		if err != nil {
			s.logger.Error("Failed to get/create OAuth user:", err)
			http.Redirect(w, r, "/login?error=user_creation_failed", http.StatusSeeOther)
			return
		}
	} else {
		// Single-user mode - create user from OAuth info
		user = &User{
			ID:       userInfo.ID,
			Username: userInfo.Email,
			Email:    userInfo.Email,
			Name:     userInfo.Name,
			Provider: provider,
		}
	}

	// Create session
	sessionData := &SessionData{
		Authenticated: true,
		UserID:        user.ID,
		Username:      user.Username,
		Email:         user.Email,
		Name:          user.Name,
		Provider:      provider,
		CreatedAt:     time.Now(),
		ExpiresAt:     time.Now().Add(time.Duration(s.authManager.config.SessionMaxAge) * time.Second),
	}

	if s.authManager.session != nil {
		if err := s.authManager.session.SetSessionData(r, w, sessionData); err != nil {
			s.logger.Error("Failed to create session after OAuth:", err)
			http.Redirect(w, r, "/login?error=session_failed", http.StatusSeeOther)
			return
		}
	}

	s.logger.Debug("OAuth login successful", "email", user.Email, "provider", provider)

	// Redirect to login redirect URL
	http.Redirect(w, r, s.authManager.config.LoginRedirect, http.StatusSeeOther)
}

// oauthUserInfo represents user info from OAuth provider
type oauthUserInfo struct {
	ID    string
	Email string
	Name  string
}

// getOAuthUserInfo fetches user info from the OAuth provider
func (s *UIServer) getOAuthUserInfo(ctx context.Context, token *oauth2.Token) (*oauthUserInfo, error) {
	oauthCfg := s.authManager.config.OAuth

	var userInfoURL string
	var idField, nameField, emailField string

	switch oauthCfg.Provider {
	case "google":
		userInfoURL = "https://www.googleapis.com/oauth2/v2/userinfo"
		idField = "id"
		nameField = "name"
		emailField = "email"
	case "custom":
		if oauthCfg.Custom == nil {
			return nil, fmt.Errorf("custom OAuth config not set")
		}
		userInfoURL = oauthCfg.Custom.UserInfoURL
		idField = oauthCfg.Custom.UserIDField
		if idField == "" {
			idField = "id"
		}
		nameField = oauthCfg.Custom.UserNameField
		if nameField == "" {
			nameField = "name"
		}
		emailField = oauthCfg.Custom.UserEmailField
		if emailField == "" {
			emailField = "email"
		}
	default:
		return nil, fmt.Errorf("unsupported OAuth provider: %s", oauthCfg.Provider)
	}

	// Create HTTP client with token
	client := oauth2.NewClient(ctx, oauth2.StaticTokenSource(token))

	resp, err := client.Get(userInfoURL)
	if err != nil {
		return nil, fmt.Errorf("failed to get user info: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		body, _ := io.ReadAll(resp.Body)
		return nil, fmt.Errorf("user info request failed with status %d: %s", resp.StatusCode, string(body))
	}

	var data map[string]interface{}
	if err := json.NewDecoder(resp.Body).Decode(&data); err != nil {
		return nil, fmt.Errorf("failed to decode user info: %w", err)
	}

	userInfo := &oauthUserInfo{}

	if id, ok := data[idField]; ok {
		userInfo.ID = fmt.Sprintf("%v", id)
	}
	if name, ok := data[nameField].(string); ok {
		userInfo.Name = name
	}
	if email, ok := data[emailField].(string); ok {
		userInfo.Email = email
	}

	if userInfo.Email == "" && userInfo.ID == "" {
		return nil, fmt.Errorf("no user identifier found in OAuth response")
	}

	return userInfo, nil
}
