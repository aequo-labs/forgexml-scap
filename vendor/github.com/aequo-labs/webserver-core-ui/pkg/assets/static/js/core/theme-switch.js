// theme-switch.js - Enhanced theme switching with server-side persistence

// Theme configuration management
let themeConfig = {
    siteTheme: 'light',
    domains: {}
};

// Initialize theme configuration
function initThemeConfig() {
    const stored = localStorage.getItem('themeConfig');
    if (stored) {
        try {
            themeConfig = JSON.parse(stored);
        } catch (e) {
            console.warn('Invalid theme config in localStorage, using defaults');
        }
    }
    
    // Ensure required properties exist
    themeConfig.siteTheme = themeConfig.siteTheme || 'light';
    themeConfig.domains = themeConfig.domains || {};
}

// Save theme configuration
function saveThemeConfig() {
    localStorage.setItem('themeConfig', JSON.stringify(themeConfig));
}

// Get system preference
function getSystemThemePreference() {
    return window.matchMedia('(prefers-color-scheme: dark)').matches ? 'dark' : 'light';
}

// Apply theme to the page
function applyTheme(theme) {
    // Update CSS files
    const lightTheme = document.getElementById('light-theme');
    const darkTheme = document.getElementById('dark-theme');
    
    if (lightTheme) lightTheme.disabled = (theme === 'dark');
    if (darkTheme) darkTheme.disabled = (theme === 'light');
    
    // Update body and html classes
    document.documentElement.classList.remove('light-mode', 'dark-mode');
    document.documentElement.classList.add(theme + '-mode');
    document.body.classList.remove('light-mode', 'dark-mode');
    document.body.classList.add(theme + '-mode');
    
    // Update theme config
    themeConfig.siteTheme = theme;
    saveThemeConfig();
    
    // Update toggle button icons
    updateToggleButton(theme);
    
    // Notify server of theme preference
    updateServerThemePreference(theme);
    
    // Dispatch event for other scripts
    document.dispatchEvent(new CustomEvent('themeChanged', { 
        detail: { theme, config: themeConfig } 
    }));
}

// Update toggle button appearance
// Update toggle button appearance
// Update toggle button appearance
function updateToggleButton(theme) {
    const toggleBtn = document.getElementById('theme-toggle');
    if (!toggleBtn) return;
    
    const lightIcon = toggleBtn.querySelector('.light-mode-icon');
    const darkIcon = toggleBtn.querySelector('.dark-mode-icon');
    
    if (lightIcon && darkIcon) {
        if (theme === 'dark') {
            lightIcon.classList.remove('d-none');
            darkIcon.classList.add('d-none');
        } else {
            lightIcon.classList.add('d-none');
            darkIcon.classList.remove('d-none');
        }
    }
}
function toggleTheme() {
    const currentTheme = themeConfig.siteTheme;
    const newTheme = currentTheme === 'light' ? 'dark' : 'light';
    applyTheme(newTheme);
}

// Update server-side theme preference
function updateServerThemePreference(theme) {
    // Always try to update server theme preference
    fetch('/api/user/theme-preference', {
        method: 'POST',
        headers: {
            'Content-Type': 'application/json'
        },
        credentials: 'same-origin',
        body: JSON.stringify({ theme: theme })
    }).catch(error => {
        console.warn('Failed to update server theme preference:', error);
    });
}

// Initialize theme system
function initTheme() {
    initThemeConfig();
    
    // Determine initial theme
    let initialTheme = themeConfig.siteTheme;
    
    // Check if server provided a preference
    const serverTheme = document.documentElement.getAttribute('data-server-theme');
    if (serverTheme && ['light', 'dark'].includes(serverTheme)) {
        initialTheme = serverTheme;
    }
    
    // Fall back to system preference if no saved preference
    if (!themeConfig.siteTheme || themeConfig.siteTheme === 'auto') {
        initialTheme = getSystemThemePreference();
    }
    
    // Apply initial theme
    applyTheme(initialTheme);
    
    // Set up toggle button
    const toggleBtn = document.getElementById('theme-toggle');
    if (toggleBtn) {
        toggleBtn.addEventListener('click', toggleTheme);
    }
    
    // Set up about button
    const aboutBtn = document.getElementById('about-button');
    if (aboutBtn) {
        aboutBtn.addEventListener('click', () => {
            window.location.href = '/about';
        });
    }
    
    // Set up settings button (uses data attribute to prevent hover URL text)
    const settingsBtn = document.getElementById('settings-button');
    if (settingsBtn) {
        settingsBtn.addEventListener('click', () => {
            const settingsUrl = settingsBtn.getAttribute('data-settings-url') || '/config';
            window.location.href = settingsUrl;
        });
    }
    
    // Listen for system theme changes
    window.matchMedia('(prefers-color-scheme: dark)').addEventListener('change', (e) => {
        // Only auto-switch if user hasn't explicitly set a preference
        if (!localStorage.getItem('themeConfig')) {
            const systemTheme = e.matches ? 'dark' : 'light';
            applyTheme(systemTheme);
        }
    });
    
    // Listen for keyboard shortcut (Ctrl/Cmd + Shift + T)
    document.addEventListener('keydown', (e) => {
        if ((e.ctrlKey || e.metaKey) && e.shiftKey && e.key === 'T') {
            e.preventDefault();
            toggleTheme();
        }
    });
}

// Expose functions for use by other scripts
window.themeSwitch = {
    toggle: toggleTheme,
    apply: applyTheme,
    getConfig: () => themeConfig,
    getCurrentTheme: () => themeConfig.siteTheme
};

// Initialize when DOM is ready
document.addEventListener('DOMContentLoaded', initTheme);
