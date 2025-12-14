// Loading overlay functionality
(function() {
    // Global flag to disable loading overlay (can be set by pages that don't want it)
    let overlayDisabled = false;
    
    // Function to create the loading overlay if it doesn't exist
    function createOverlay() {
        if (!document.getElementById('loading-overlay')) {
            const overlay = document.createElement('div');
            overlay.id = 'loading-overlay';
            overlay.style.display = 'none';
            overlay.style.pointerEvents = 'none';
            overlay.style.position = 'fixed';
            overlay.style.top = '0';
            overlay.style.left = '0';
            overlay.style.width = '100%';
            overlay.style.height = '100%';
            overlay.style.backgroundColor = 'rgba(0,0,0,0.5)';
            overlay.style.zIndex = '9999';
            
            const spinnerContainer = document.createElement('div');
            spinnerContainer.style.position = 'absolute';
            spinnerContainer.style.top = '50%';
            spinnerContainer.style.left = '50%';
            spinnerContainer.style.transform = 'translate(-50%, -50%)';
            spinnerContainer.style.textAlign = 'center';
            spinnerContainer.style.color = 'white';
            
            const spinner = document.createElement('div');
            spinner.className = 'loading-spinner';
            spinner.style.margin = '0 auto';
            
            const loadingText = document.createElement('p');
            loadingText.textContent = 'Loading...';
            loadingText.style.marginTop = '20px';
            loadingText.style.fontSize = '18px';
            
            spinnerContainer.appendChild(spinner);
            spinnerContainer.appendChild(loadingText);
            overlay.appendChild(spinnerContainer);
            
            document.body.appendChild(overlay);
        }
    }

    // Function to hide the loading overlay
    function hideOverlay() {
        const overlay = document.getElementById('loading-overlay');
        if (overlay) {
            overlay.style.display = 'none';
            overlay.style.pointerEvents = 'none';
        }
        
        // Clear any existing timer
        const existingTimer = sessionStorage.getItem('loadingTimer');
        if (existingTimer) {
            clearTimeout(parseInt(existingTimer));
            sessionStorage.removeItem('loadingTimer');
        }
    }
    
    // Function to show the loading overlay (respects disabled flag)
    function showOverlay() {
        if (overlayDisabled) {
            return null;
        }
        
        const overlay = document.getElementById('loading-overlay');
        if (overlay) {
            overlay.style.display = 'block';
            overlay.style.pointerEvents = 'auto';
        }
        
        return null;
    }

    // Create overlay when DOM is ready
    if (document.readyState === 'loading') {
        document.addEventListener('DOMContentLoaded', createOverlay);
    } else {
        createOverlay();
    }

    // Hide overlay on page load - multiple events for redundancy
    window.addEventListener('load', hideOverlay);
    window.addEventListener('pageshow', hideOverlay);
    window.addEventListener('popstate', hideOverlay);
    
    // Add click event listeners to all navigation links
    document.addEventListener('DOMContentLoaded', function() {
        // Check if page has data-no-loading-overlay attribute on body
        if (document.body.hasAttribute('data-no-loading-overlay')) {
            overlayDisabled = true;
        }
        
        // Use capturing phase to intercept download link clicks early
        document.addEventListener('click', function(e) {
            // Check if overlay is disabled
            if (overlayDisabled) {
                return;
            }
            
            // Find the actual link element (might be nested inside the link)
            let target = e.target;
            while (target && target.tagName !== 'A') {
                target = target.parentElement;
            }
            
            // If this is a download link or has data-no-loading attribute, ensure overlay is hidden
            if (target && target.tagName === 'A') {
                if (target.hasAttribute('download') || target.hasAttribute('data-no-loading')) {
                    hideOverlay();
                    e.__skipLoading = true;
                    return;
                }
            }
            
            // Check for data-no-loading on any clicked element or its parents
            let checkTarget = e.target;
            while (checkTarget) {
                if (checkTarget.hasAttribute && checkTarget.hasAttribute('data-no-loading')) {
                    hideOverlay();
                    e.__skipLoading = true;
                    return;
                }
                checkTarget = checkTarget.parentElement;
            }
        }, true); // Use capture phase
        
        const links = document.querySelectorAll('a:not([target="_blank"]):not([href^="#"]):not([href^="javascript:"]):not([download]):not([data-no-loading])');
        links.forEach(link => {
            link.addEventListener('click', function(e) {
                // Check if overlay is disabled or this event should skip loading
                if (overlayDisabled || e.__skipLoading) {
                    hideOverlay();
                    return true;
                }
                
                // Double-check: Don't show loading overlay for special links
                const href = this.getAttribute('href');
                if (this.getAttribute('target') === '_blank' || 
                    (href && href.startsWith('#')) ||
                    this.hasAttribute('download') ||
                    this.hasAttribute('data-no-loading')) {
                    hideOverlay();
                    return true;
                }
                
                // Show overlay for navigation
                showOverlay();
                
                return true;
            });
        });
        
        // Add submit event listeners to all forms
        const forms = document.querySelectorAll('form:not([target="_blank"]):not([data-no-loading])');
        forms.forEach(form => {
            form.addEventListener('submit', function(e) {
                // Check if overlay is disabled
                if (overlayDisabled) {
                    return true;
                }
                
                // Don't show loading overlay for forms that should open in a new tab
                if (this.getAttribute('target') === '_blank' || this.hasAttribute('data-no-loading')) {
                    return true;
                }
                
                // Show overlay for form submission
                showOverlay();
                
                return true;
            });
        });
    });
    
    // Clear any existing loading timer when the page loads
    document.addEventListener('DOMContentLoaded', function() {
        const existingTimer = sessionStorage.getItem('loadingTimer');
        if (existingTimer) {
            clearTimeout(parseInt(existingTimer));
            sessionStorage.removeItem('loadingTimer');
        }
        
        // Ensure overlay is hidden
        hideOverlay();
    });
    
    // Make functions available globally
    window.showLoadingOverlay = function() {
        return showOverlay();
    };
    
    window.hideLoadingOverlay = function(timerId) {
        // Clear the timeout if it hasn't fired yet
        if (timerId) {
            clearTimeout(timerId);
        }
        
        // Hide the overlay
        hideOverlay();
    };
    
    // Function to disable loading overlay for the current page
    window.disableLoadingOverlay = function() {
        overlayDisabled = true;
        hideOverlay();
    };
    
    // Function to re-enable loading overlay
    window.enableLoadingOverlay = function() {
        overlayDisabled = false;
    };
    
    // Function to check if overlay is disabled
    window.isLoadingOverlayDisabled = function() {
        return overlayDisabled;
    };
})();
