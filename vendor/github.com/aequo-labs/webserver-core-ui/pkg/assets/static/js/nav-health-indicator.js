/**
 * Unified Nav Health Indicator Component
 * Provides consistent health status monitoring across all WebServer-Core applications
 */
(function() {
    'use strict';
    
    let healthCheckInterval;
    const HEALTH_CHECK_INTERVAL = 30000; // 30 seconds
    
    /**
     * Updates the navigation health indicator status
     * @param {string} status - 'online', 'offline', 'checking', 'error'
     * @param {string} message - Optional message for error states
     * @param {Date} timestamp - When the check was performed
     */
    function updateNavHealthStatus(status, message, timestamp) {
        const indicator = document.getElementById('nav-health-status');
        
        if (!indicator) return;
        
        // Remove all status classes
        indicator.classList.remove('online', 'offline', 'checking', 'error');
        
        // Format timestamp for tooltip
        const timeStr = timestamp ? ' (Last checked: ' + timestamp.toLocaleTimeString() + ')' : '';
        
        // Update based on status
        switch(status) {
            case 'online':
                indicator.classList.add('online');
                indicator.title = 'Server is online' + timeStr;
                break;
                
            case 'offline':
                indicator.classList.add('offline');
                indicator.title = 'Server is offline: ' + (message || 'Connection failed') + timeStr;
                break;
                
            case 'checking':
                indicator.classList.add('checking');
                indicator.title = 'Checking server connection...';
                break;
                
            case 'error':
                indicator.classList.add('error');
                indicator.title = 'Server error: ' + (message || 'Authentication failed') + timeStr;
                break;
        }
        
        // Dispatch custom event for other components to listen to
        window.dispatchEvent(new CustomEvent('nav-health-status-update', {
            detail: { status, message, timestamp }
        }));
    }
    
    /**
     * Performs health check against /api/health endpoint
     */
    async function performHealthCheck() {
        updateNavHealthStatus('checking');
        
        try {
            const response = await fetch('/api/health', {
                method: 'GET',
                headers: {
                    'Accept': 'application/json',
                },
                timeout: 10000, // 10 second timeout
            });
            
            const timestamp = new Date();
            
            if (response.ok) {
                const data = await response.json();
                updateNavHealthStatus('online', data.message || 'Healthy', timestamp);
            } else {
                const errorText = await response.text();
                updateNavHealthStatus('error', `HTTP ${response.status}: ${errorText}`, timestamp);
            }
        } catch (error) {
            const timestamp = new Date();
            if (error.name === 'AbortError') {
                updateNavHealthStatus('offline', 'Request timeout', timestamp);
            } else {
                updateNavHealthStatus('offline', error.message || 'Connection failed', timestamp);
            }
        }
    }
    
    /**
     * Starts the health monitoring
     */
    function startHealthMonitoring() {
        // Perform initial check
        performHealthCheck();
        
        // Set up periodic checks
        healthCheckInterval = setInterval(performHealthCheck, HEALTH_CHECK_INTERVAL);
    }
    
    /**
     * Stops the health monitoring
     */
    function stopHealthMonitoring() {
        if (healthCheckInterval) {
            clearInterval(healthCheckInterval);
            healthCheckInterval = null;
        }
    }
    
    /**
     * Initialize the health indicator when DOM is ready
     */
    function init() {
        if (document.readyState === 'loading') {
            document.addEventListener('DOMContentLoaded', startHealthMonitoring);
        } else {
            startHealthMonitoring();
        }
        
        // Clean up on page unload
        window.addEventListener('beforeunload', stopHealthMonitoring);
        
        // Listen for visibility changes to pause/resume monitoring
        document.addEventListener('visibilitychange', function() {
            if (document.hidden) {
                stopHealthMonitoring();
            } else {
                startHealthMonitoring();
            }
        });
    }
    
    // Expose public API
    window.NavHealthIndicator = {
        init: init,
        updateStatus: updateNavHealthStatus,
        performCheck: performHealthCheck,
        start: startHealthMonitoring,
        stop: stopHealthMonitoring
    };
    
    // Auto-initialize
    init();
})();