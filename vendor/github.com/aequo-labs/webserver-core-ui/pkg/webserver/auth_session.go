package webserver

import (
	"encoding/gob"
	"net/http"
	"time"

	"github.com/gorilla/sessions"
)

func init() {
	// Register types for gob encoding
	gob.Register(SessionData{})
	gob.Register(time.Time{})
}

// SessionManager handles session operations
type SessionManager struct {
	store      *sessions.CookieStore
	name       string
	maxAge     int
	cookiePath string
}

// NewSessionManager creates a new session manager
func NewSessionManager(secret string, name string, maxAge int) *SessionManager {
	store := sessions.NewCookieStore([]byte(secret))
	store.Options = &sessions.Options{
		Path:     "/",
		MaxAge:   maxAge,
		HttpOnly: true,
		Secure:   false, // Set to true in production with HTTPS
		SameSite: http.SameSiteLaxMode,
	}

	return &SessionManager{
		store:      store,
		name:       name,
		maxAge:     maxAge,
		cookiePath: "/",
	}
}

// GetStore returns the underlying session store
func (sm *SessionManager) GetStore() *sessions.CookieStore {
	return sm.store
}

// Get retrieves a session from the request
func (sm *SessionManager) Get(r *http.Request) (*sessions.Session, error) {
	return sm.store.Get(r, sm.name)
}

// Save saves the session to the response
func (sm *SessionManager) Save(r *http.Request, w http.ResponseWriter, session *sessions.Session) error {
	return sm.store.Save(r, w, session)
}

// GetSessionData retrieves the session data from the request
func (sm *SessionManager) GetSessionData(r *http.Request) (*SessionData, error) {
	session, err := sm.Get(r)
	if err != nil {
		return nil, err
	}

	// Check if we have session data stored
	if data, ok := session.Values["data"].(SessionData); ok {
		return &data, nil
	}

	// Try to construct from individual values (backwards compatibility)
	authenticated, _ := session.Values["authenticated"].(bool)
	if !authenticated {
		return nil, nil
	}

	data := &SessionData{
		Authenticated: authenticated,
	}
	if username, ok := session.Values["username"].(string); ok {
		data.Username = username
	}
	if email, ok := session.Values["email"].(string); ok {
		data.Email = email
	}
	if name, ok := session.Values["name"].(string); ok {
		data.Name = name
	}
	if provider, ok := session.Values["provider"].(string); ok {
		data.Provider = provider
	}
	if userID, ok := session.Values["userId"].(string); ok {
		data.UserID = userID
	}

	return data, nil
}

// SetSessionData sets the session data
func (sm *SessionManager) SetSessionData(r *http.Request, w http.ResponseWriter, data *SessionData) error {
	session, err := sm.Get(r)
	if err != nil {
		return err
	}

	// Store as structured data
	session.Values["data"] = *data

	// Also store individual values for backwards compatibility
	session.Values["authenticated"] = data.Authenticated
	session.Values["username"] = data.Username
	session.Values["email"] = data.Email
	session.Values["name"] = data.Name
	session.Values["provider"] = data.Provider
	session.Values["userId"] = data.UserID

	// Adjust max age for remember me
	if data.RememberMe {
		session.Options.MaxAge = 30 * 24 * 60 * 60 // 30 days
	}

	return sm.Save(r, w, session)
}

// ClearSession clears the session data
func (sm *SessionManager) ClearSession(r *http.Request, w http.ResponseWriter) error {
	session, err := sm.Get(r)
	if err != nil {
		return err
	}

	// Clear all values
	session.Values = make(map[interface{}]interface{})
	session.Options.MaxAge = -1 // Delete the cookie

	return sm.Save(r, w, session)
}

// IsAuthenticated checks if the session is authenticated
func (sm *SessionManager) IsAuthenticated(r *http.Request) bool {
	data, err := sm.GetSessionData(r)
	if err != nil {
		return false
	}
	return data != nil && data.Authenticated
}

// GetAuthenticatedUser returns the authenticated user info from session
func (sm *SessionManager) GetAuthenticatedUser(r *http.Request) (*SessionData, bool) {
	data, err := sm.GetSessionData(r)
	if err != nil || data == nil {
		return nil, false
	}
	if !data.Authenticated {
		return nil, false
	}
	return data, true
}

// SetThemePreference stores the theme preference in the session
func (sm *SessionManager) SetThemePreference(r *http.Request, w http.ResponseWriter, theme string) error {
	session, err := sm.Get(r)
	if err != nil {
		return err
	}
	session.Values["theme"] = theme
	return sm.Save(r, w, session)
}

// GetThemePreference retrieves the theme preference from the session
func (sm *SessionManager) GetThemePreference(r *http.Request) string {
	session, err := sm.Get(r)
	if err != nil {
		return "light"
	}
	if theme, ok := session.Values["theme"].(string); ok && theme != "" {
		return theme
	}
	return "light"
}

// SetFlash stores a flash message in the session
func (sm *SessionManager) SetFlash(r *http.Request, w http.ResponseWriter, flashType string, message string) error {
	session, err := sm.Get(r)
	if err != nil {
		return err
	}
	session.AddFlash(message, flashType)
	return sm.Save(r, w, session)
}

// GetFlashes retrieves and clears flash messages of a specific type
func (sm *SessionManager) GetFlashes(r *http.Request, w http.ResponseWriter, flashType string) []string {
	session, err := sm.Get(r)
	if err != nil {
		return nil
	}
	flashes := session.Flashes(flashType)
	if len(flashes) > 0 {
		_ = sm.Save(r, w, session)
	}
	messages := make([]string, 0, len(flashes))
	for _, f := range flashes {
		if msg, ok := f.(string); ok {
			messages = append(messages, msg)
		}
	}
	return messages
}
