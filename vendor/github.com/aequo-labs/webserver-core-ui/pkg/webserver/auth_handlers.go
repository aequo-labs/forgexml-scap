package webserver

import (
	"encoding/json"
	"net/http"
	"time"

	"golang.org/x/crypto/bcrypt"
)

// authStatusHandler returns the current authentication status
func (s *UIServer) authStatusHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")

	status := AuthStatus{
		Authenticated: false,
		SetupNeeded:   false,
		AuthMethod:    "",
		OAuthEnabled:  false,
		LocalEnabled:  false,
	}

	if s.authManager == nil || !s.authManager.IsEnabled() {
		// Auth not configured - report as authenticated (no auth required)
		status.Authenticated = true
		json.NewEncoder(w).Encode(status)
		return
	}

	status.AuthMethod = s.authManager.config.AuthMethod
	status.OAuthEnabled = s.authManager.IsOAuthEnabled()
	status.LocalEnabled = s.authManager.IsLocalAuthEnabled()
	status.SetupNeeded = s.authManager.IsSetupNeeded()

	// Check if user is authenticated
	if s.authManager.session != nil {
		if sessionData, ok := s.authManager.session.GetAuthenticatedUser(r); ok {
			status.Authenticated = true
			status.Username = sessionData.Username
			status.Email = sessionData.Email
			status.Name = sessionData.Name
			status.Provider = sessionData.Provider
		}
	}

	json.NewEncoder(w).Encode(status)
}

// loginPageHandler renders the login page
func (s *UIServer) loginPageHandler(w http.ResponseWriter, r *http.Request) {
	// Check if already authenticated
	if s.IsAuthenticated(r) {
		http.Redirect(w, r, s.authManager.config.LoginRedirect, http.StatusSeeOther)
		return
	}

	// Check if setup is needed
	if s.authManager.IsSetupNeeded() {
		http.Redirect(w, r, "/setup", http.StatusSeeOther)
		return
	}

	data := s.GetBasePageDataWithRequest(r, "login")
	data.Title = "Login"

	// Add auth-specific data for template
	type LoginPageData struct {
		PageData
		ErrorMessage  string
		OAuthEnabled  bool
		LocalEnabled  bool
		OAuthProvider string
	}

	loginData := LoginPageData{
		PageData:      data,
		ErrorMessage:  r.URL.Query().Get("error"),
		OAuthEnabled:  s.authManager.IsOAuthEnabled(),
		LocalEnabled:  s.authManager.IsLocalAuthEnabled(),
		OAuthProvider: s.authManager.GetOAuthProvider(),
	}

	// Allow application to enrich page data
	s.EnrichPageData(&loginData.PageData)

	if err := s.RenderPageWithContent(w, "login-content", loginData); err != nil {
		s.logger.Error("Failed to render login page:", err)
	}
}

// loginHandler handles local login form submission
func (s *UIServer) loginHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")

	if s.authManager == nil || !s.authManager.IsEnabled() {
		w.WriteHeader(http.StatusServiceUnavailable)
		json.NewEncoder(w).Encode(map[string]interface{}{
			"success": false,
			"error":   "Authentication not configured",
		})
		return
	}

	if !s.authManager.IsLocalAuthEnabled() {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]interface{}{
			"success": false,
			"error":   "Local authentication not enabled",
		})
		return
	}

	// Parse request body
	var request struct {
		Username   string `json:"username"`
		Password   string `json:"password"`
		RememberMe bool   `json:"rememberMe"`
	}

	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		// Try form data
		if err := r.ParseForm(); err != nil {
			w.WriteHeader(http.StatusBadRequest)
			json.NewEncoder(w).Encode(map[string]interface{}{
				"success": false,
				"error":   "Invalid request",
			})
			return
		}
		request.Username = r.FormValue("username")
		request.Password = r.FormValue("password")
		request.RememberMe = r.FormValue("rememberMe") == "true" || r.FormValue("rememberMe") == "on"
	}

	// Validate credentials
	var authenticated bool
	var user *User

	if s.authManager.config.UserStore != nil {
		// Multi-user mode
		user, err := s.authManager.config.UserStore.GetUserByUsername(request.Username)
		if err != nil || user == nil {
			w.WriteHeader(http.StatusUnauthorized)
			json.NewEncoder(w).Encode(map[string]interface{}{
				"success": false,
				"error":   "Invalid username or password",
			})
			return
		}
		authenticated = s.authManager.config.UserStore.ValidatePassword(user, request.Password)
	} else {
		// Single-user mode
		localAuth := s.authManager.config.LocalAuth
		if localAuth == nil {
			w.WriteHeader(http.StatusServiceUnavailable)
			json.NewEncoder(w).Encode(map[string]interface{}{
				"success": false,
				"error":   "Local authentication not configured",
			})
			return
		}

		if request.Username != localAuth.Username {
			w.WriteHeader(http.StatusUnauthorized)
			json.NewEncoder(w).Encode(map[string]interface{}{
				"success": false,
				"error":   "Invalid username or password",
			})
			return
		}

		// Verify password with bcrypt
		err := bcrypt.CompareHashAndPassword([]byte(localAuth.PasswordHash), []byte(request.Password))
		authenticated = err == nil
		if authenticated {
			user = &User{
				Username: localAuth.Username,
				Provider: "local",
			}
		}
	}

	if !authenticated {
		w.WriteHeader(http.StatusUnauthorized)
		json.NewEncoder(w).Encode(map[string]interface{}{
			"success": false,
			"error":   "Invalid username or password",
		})
		return
	}

	// Create session
	sessionData := &SessionData{
		Authenticated: true,
		UserID:        user.ID,
		Username:      user.Username,
		Email:         user.Email,
		Name:          user.Name,
		Provider:      "local",
		RememberMe:    request.RememberMe,
		CreatedAt:     time.Now(),
	}

	if request.RememberMe && s.authManager.config.LocalAuth != nil {
		days := s.authManager.config.LocalAuth.RememberMeDays
		if days == 0 {
			days = 30
		}
		sessionData.ExpiresAt = time.Now().AddDate(0, 0, days)
	} else {
		sessionData.ExpiresAt = time.Now().Add(time.Duration(s.authManager.config.SessionMaxAge) * time.Second)
	}

	if err := s.authManager.session.SetSessionData(r, w, sessionData); err != nil {
		s.logger.Error("Failed to create session:", err)
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(map[string]interface{}{
			"success": false,
			"error":   "Failed to create session",
		})
		return
	}

	s.logger.Debug("User logged in", "username", user.Username, "rememberMe", request.RememberMe)

	json.NewEncoder(w).Encode(map[string]interface{}{
		"success":  true,
		"redirect": s.authManager.config.LoginRedirect,
	})
}

// logoutHandler handles logout requests
func (s *UIServer) logoutHandler(w http.ResponseWriter, r *http.Request) {
	if s.authManager != nil && s.authManager.session != nil {
		if err := s.authManager.session.ClearSession(r, w); err != nil {
			s.logger.Error("Failed to clear session:", err)
		}
	}

	// Check if this is an API request
	if r.Header.Get("Accept") == "application/json" || r.Header.Get("Content-Type") == "application/json" {
		w.Header().Set("Content-Type", "application/json; charset=utf-8")
		json.NewEncoder(w).Encode(map[string]interface{}{
			"success":  true,
			"redirect": s.authManager.config.LogoutRedirect,
		})
		return
	}

	// Regular request - redirect
	redirectURL := "/"
	if s.authManager != nil {
		redirectURL = s.authManager.config.LogoutRedirect
	}
	http.Redirect(w, r, redirectURL, http.StatusSeeOther)
}

// setupPageHandler renders the first-time setup page
func (s *UIServer) setupPageHandler(w http.ResponseWriter, r *http.Request) {
	// Check if setup is already complete
	if s.authManager != nil && !s.authManager.IsSetupNeeded() {
		http.Redirect(w, r, "/login", http.StatusSeeOther)
		return
	}

	data := s.GetBasePageData("setup")
	data.Title = "Setup"

	type SetupPageData struct {
		PageData
		ErrorMessage string
	}

	setupData := SetupPageData{
		PageData:     data,
		ErrorMessage: r.URL.Query().Get("error"),
	}

	if err := s.RenderPageWithContent(w, "setup-content", setupData); err != nil {
		s.logger.Error("Failed to render setup page:", err)
	}
}

// setupHandler handles the setup form submission
func (s *UIServer) setupHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")

	// Check if setup is already complete
	if s.authManager != nil && !s.authManager.IsSetupNeeded() {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]interface{}{
			"success": false,
			"error":   "Setup already complete",
		})
		return
	}

	// Parse request
	var request struct {
		Username string `json:"username"`
		Password string `json:"password"`
	}

	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		if err := r.ParseForm(); err != nil {
			w.WriteHeader(http.StatusBadRequest)
			json.NewEncoder(w).Encode(map[string]interface{}{
				"success": false,
				"error":   "Invalid request",
			})
			return
		}
		request.Username = r.FormValue("username")
		request.Password = r.FormValue("password")
	}

	if request.Username == "" || request.Password == "" {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]interface{}{
			"success": false,
			"error":   "Username and password are required",
		})
		return
	}

	// Hash password
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(request.Password), bcrypt.DefaultCost)
	if err != nil {
		s.logger.Error("Failed to hash password:", err)
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(map[string]interface{}{
			"success": false,
			"error":   "Failed to process password",
		})
		return
	}

	// Update config with new credentials
	if s.authManager.config.LocalAuth == nil {
		s.authManager.config.LocalAuth = &LocalAuthConfig{}
	}
	s.authManager.config.LocalAuth.Username = request.Username
	s.authManager.config.LocalAuth.PasswordHash = string(hashedPassword)

	// The application should persist these credentials
	// This is just updating the in-memory config

	s.logger.Debug("Setup completed", "username", request.Username)

	// Auto-login after setup
	sessionData := &SessionData{
		Authenticated: true,
		Username:      request.Username,
		Provider:      "local",
		CreatedAt:     time.Now(),
		ExpiresAt:     time.Now().Add(time.Duration(s.authManager.config.SessionMaxAge) * time.Second),
	}

	if s.authManager.session != nil {
		if err := s.authManager.session.SetSessionData(r, w, sessionData); err != nil {
			s.logger.Error("Failed to create session after setup:", err)
		}
	}

	json.NewEncoder(w).Encode(map[string]interface{}{
		"success":  true,
		"message":  "Setup complete",
		"redirect": s.authManager.config.LoginRedirect,
	})
}

// HashPassword creates a bcrypt hash of a password
// This is a helper function for applications to use when setting up credentials
func HashPassword(password string) (string, error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}
	return string(hashedPassword), nil
}

// VerifyPassword checks if a password matches a bcrypt hash
func VerifyPassword(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}
