/**
 * Authentication JavaScript module for webserver-core-ui
 * Provides client-side auth status checking, login/logout UI updates,
 * and CSRF token management.
 */

(function() {
    'use strict';

    // Auth state
    let authState = {
        authenticated: false,
        username: '',
        email: '',
        setupNeeded: false,
        authMethod: '',
        oauthEnabled: false,
        localEnabled: false
    };

    // CSRF token (if available)
    let csrfToken = '';

    /**
     * Initialize auth module
     */
    function init() {
        // Get CSRF token from meta tag if present
        const csrfMeta = document.querySelector('meta[name="csrf-token"]');
        if (csrfMeta) {
            csrfToken = csrfMeta.getAttribute('content');
        }

        // Check auth status on page load
        checkAuthStatus();

        // Setup logout handlers
        setupLogoutHandlers();

        // Setup periodic auth check (every 5 minutes)
        setInterval(checkAuthStatus, 5 * 60 * 1000);
    }

    /**
     * Check authentication status from server
     */
    async function checkAuthStatus() {
        try {
            const response = await fetch('/api/auth/status', {
                method: 'GET',
                credentials: 'same-origin',
                headers: {
                    'Accept': 'application/json'
                }
            });

            if (response.ok) {
                const data = await response.json();
                updateAuthState(data);
                updateUI();
            }
        } catch (error) {
            console.error('Failed to check auth status:', error);
        }
    }

    /**
     * Update internal auth state
     */
    function updateAuthState(data) {
        authState = {
            authenticated: data.authenticated || false,
            username: data.username || '',
            email: data.email || '',
            name: data.name || '',
            setupNeeded: data.setupNeeded || false,
            authMethod: data.authMethod || '',
            oauthEnabled: data.oauthEnabled || false,
            localEnabled: data.localEnabled || false
        };
    }

    /**
     * Update UI elements based on auth state
     */
    function updateUI() {
        // Update user icon/button in navbar
        const userIcon = document.querySelector('.navbar-user-icon, .user-menu');
        if (userIcon) {
            if (authState.authenticated) {
                userIcon.classList.add('is-authenticated');
                userIcon.classList.remove('is-unauthenticated');
                
                // Update username display if present
                const usernameDisplay = userIcon.querySelector('.username-display');
                if (usernameDisplay) {
                    usernameDisplay.textContent = authState.username || authState.name || authState.email || 'User';
                }
            } else {
                userIcon.classList.remove('is-authenticated');
                userIcon.classList.add('is-unauthenticated');
            }
        }

        // Update login/logout buttons
        const loginBtn = document.querySelector('.login-btn, [data-auth="login"]');
        const logoutBtn = document.querySelector('.logout-btn, [data-auth="logout"]');

        if (loginBtn) {
            if (authState.authenticated) {
                loginBtn.classList.add('d-none');
            } else {
                loginBtn.classList.remove('d-none');
            }
        }
        if (logoutBtn) {
            if (authState.authenticated) {
                logoutBtn.classList.remove('d-none');
            } else {
                logoutBtn.classList.add('d-none');
            }
        }

        // Dispatch custom event for other scripts to react
        document.dispatchEvent(new CustomEvent('authStateChanged', {
            detail: authState
        }));
    }

    /**
     * Setup logout button handlers
     */
    function setupLogoutHandlers() {
        document.addEventListener('click', function(e) {
            const logoutBtn = e.target.closest('[data-auth="logout"], .logout-btn');
            if (logoutBtn) {
                e.preventDefault();
                logout();
            }
        });
    }

    /**
     * Perform logout
     */
    async function logout() {
        try {
            const headers = {
                'Accept': 'application/json',
                'Content-Type': 'application/json'
            };
            
            if (csrfToken) {
                headers['X-CSRF-Token'] = csrfToken;
            }

            const response = await fetch('/api/auth/logout', {
                method: 'POST',
                credentials: 'same-origin',
                headers: headers
            });

            if (response.ok) {
                const data = await response.json();
                // Redirect to logout destination
                window.location.href = data.redirect || '/login';
            } else {
                // Fallback: just redirect to login
                window.location.href = '/login';
            }
        } catch (error) {
            console.error('Logout failed:', error);
            // Fallback: redirect to login anyway
            window.location.href = '/login';
        }
    }

    /**
     * Perform login via API
     */
    async function login(username, password, rememberMe) {
        const headers = {
            'Accept': 'application/json',
            'Content-Type': 'application/json'
        };
        
        if (csrfToken) {
            headers['X-CSRF-Token'] = csrfToken;
        }

        const response = await fetch('/api/auth/login', {
            method: 'POST',
            credentials: 'same-origin',
            headers: headers,
            body: JSON.stringify({
                username: username,
                password: password,
                rememberMe: rememberMe || false
            })
        });

        const data = await response.json();
        
        if (data.success) {
            // Update local state
            await checkAuthStatus();
            return { success: true, redirect: data.redirect };
        } else {
            return { success: false, error: data.error || 'Login failed' };
        }
    }

    /**
     * Get current auth state
     */
    function getAuthState() {
        return { ...authState };
    }

    /**
     * Check if user is authenticated
     */
    function isAuthenticated() {
        return authState.authenticated;
    }

    /**
     * Get CSRF token for forms
     */
    function getCsrfToken() {
        return csrfToken;
    }

    /**
     * Add CSRF token to a form
     */
    function addCsrfToForm(form) {
        if (!csrfToken) return;
        
        let input = form.querySelector('input[name="_csrf"]');
        if (!input) {
            input = document.createElement('input');
            input.type = 'hidden';
            input.name = '_csrf';
            form.appendChild(input);
        }
        input.value = csrfToken;
    }

    /**
     * Add CSRF token to fetch headers
     */
    function getAuthHeaders(additionalHeaders) {
        const headers = {
            'Accept': 'application/json',
            'Content-Type': 'application/json',
            ...additionalHeaders
        };
        
        if (csrfToken) {
            headers['X-CSRF-Token'] = csrfToken;
        }
        
        return headers;
    }

    // Initialize on DOM ready
    if (document.readyState === 'loading') {
        document.addEventListener('DOMContentLoaded', init);
    } else {
        init();
    }

    // Expose public API
    window.Auth = {
        checkStatus: checkAuthStatus,
        login: login,
        logout: logout,
        getState: getAuthState,
        isAuthenticated: isAuthenticated,
        getCsrfToken: getCsrfToken,
        addCsrfToForm: addCsrfToForm,
        getAuthHeaders: getAuthHeaders
    };

})();
