// auth.js - Common authentication functionality

// Auto-logout timer for inactive sessions
let logoutTimer;
const LOGOUT_WARNING_TIME = 23 * 60 * 1000; // 23 minutes (warn 1 minute before 24h expiry)
const LOGOUT_TIME = 24 * 60 * 60 * 1000; // 24 hours

// Initialize authentication features
function initAuth() {
    // Set up auto-logout timer
    resetLogoutTimer();
    
    // Reset timer on user activity
    ['mousedown', 'mousemove', 'keypress', 'scroll', 'touchstart'].forEach(event => {
        document.addEventListener(event, resetLogoutTimer, true);
    });
    
    // Handle logout buttons
    document.addEventListener('click', function(e) {
        if (e.target.matches('[data-logout]') || e.target.closest('[data-logout]')) {
            e.preventDefault();
            logout();
        }
    });
}

// Reset the logout timer
function resetLogoutTimer() {
    clearTimeout(logoutTimer);
    
    // Show warning before logout
    logoutTimer = setTimeout(() => {
        showLogoutWarning();
    }, LOGOUT_WARNING_TIME);
}

// Show logout warning
function showLogoutWarning() {
    if (confirm('Your session will expire in 1 minute. Do you want to stay logged in?')) {
        // User wants to stay logged in - extend session
        extendSession();
    } else {
        // User accepted logout or didn't respond
        logout();
    }
}

// Extend the current session
function extendSession() {
    fetch('/api/extend-session', {
        method: 'POST',
        credentials: 'same-origin',
        headers: {
            'Content-Type': 'application/json',
        }
    })
    .then(response => {
        if (response.ok) {
            resetLogoutTimer();
            showNotification('Session extended successfully', 'is-success');
        } else {
            showNotification('Failed to extend session', 'is-warning');
            logout();
        }
    })
    .catch(error => {
        console.error('Error extending session:', error);
        showNotification('Connection error', 'is-danger');
        logout();
    });
}

// Logout the user
function logout() {
    // Clear any timers
    clearTimeout(logoutTimer);
    
    // Make logout request
    fetch('/logout', {
        method: 'POST',
        credentials: 'same-origin',
        headers: {
            'Content-Type': 'application/json',
        }
    })
    .then(response => {
        // Redirect to login page regardless of response
        window.location.href = '/login';
    })
    .catch(error => {
        console.error('Logout error:', error);
        // Redirect anyway
        window.location.href = '/login';
    });
}

// Show notification to user
function showNotification(message, type = 'is-info') {
    // Create notification element
    const notification = document.createElement('div');
    notification.className = `notification ${type} is-light`;
    notification.innerHTML = `
        <button class="delete"></button>
        ${message}
    `;
    
    // Add to page
    const container = document.querySelector('.main-content') || document.body;
    container.insertBefore(notification, container.firstChild);
    
    // Handle delete button
    const deleteBtn = notification.querySelector('.delete');
    deleteBtn.addEventListener('click', () => {
        notification.remove();
    });
    
    // Auto-remove after 5 seconds
    setTimeout(() => {
        if (notification.parentNode) {
            notification.remove();
        }
    }, 5000);
}

// CSRF token management
function getCSRFToken() {
    const meta = document.querySelector('meta[name="csrf-token"]');
    return meta ? meta.getAttribute('content') : '';
}

// Add CSRF token to fetch requests
function authFetch(url, options = {}) {
    const csrfToken = getCSRFToken();
    
    if (csrfToken) {
        options.headers = options.headers || {};
        options.headers['X-CSRF-Token'] = csrfToken;
    }
    
    options.credentials = options.credentials || 'same-origin';
    
    return fetch(url, options);
}

// Initialize when DOM is ready
document.addEventListener('DOMContentLoaded', initAuth);