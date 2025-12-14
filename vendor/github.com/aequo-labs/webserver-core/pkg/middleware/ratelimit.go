package middleware

import (
	"net/http"
	"sync"
	"time"
)

// RateLimiter implements a token bucket rate limiter per IP
type RateLimiter struct {
	mu       sync.RWMutex
	clients  map[string]*clientLimiter
	rate     int           // requests per window
	window   time.Duration // time window
	cleanup  time.Duration // cleanup interval for old entries
	stopChan chan struct{}
}

type clientLimiter struct {
	tokens    int
	lastReset time.Time
	failures  int       // track failed auth attempts
	lockUntil time.Time // lockout time for brute force protection
}

// RateLimitConfig holds configuration for rate limiting
type RateLimitConfig struct {
	RequestsPerMinute int           // max requests per minute per IP
	BurstSize         int           // allow burst above limit temporarily
	LockoutThreshold  int           // failed attempts before lockout
	LockoutDuration   time.Duration // how long to lock out after threshold
	CleanupInterval   time.Duration // how often to clean old entries
}

// DefaultRateLimitConfig returns sensible defaults
func DefaultRateLimitConfig() RateLimitConfig {
	return RateLimitConfig{
		RequestsPerMinute: 60,
		BurstSize:         10,
		LockoutThreshold:  5,
		LockoutDuration:   15 * time.Minute,
		CleanupInterval:   5 * time.Minute,
	}
}

// NewRateLimiter creates a new rate limiter
func NewRateLimiter(config RateLimitConfig) *RateLimiter {
	rl := &RateLimiter{
		clients:  make(map[string]*clientLimiter),
		rate:     config.RequestsPerMinute + config.BurstSize,
		window:   time.Minute,
		cleanup:  config.CleanupInterval,
		stopChan: make(chan struct{}),
	}

	// Start cleanup goroutine
	go rl.cleanupLoop()

	return rl
}

// cleanupLoop removes old client entries periodically
func (rl *RateLimiter) cleanupLoop() {
	ticker := time.NewTicker(rl.cleanup)
	defer ticker.Stop()

	for {
		select {
		case <-ticker.C:
			rl.mu.Lock()
			now := time.Now()
			for ip, client := range rl.clients {
				// Remove entries that haven't been seen in 2x the window
				if now.Sub(client.lastReset) > rl.window*2 && now.After(client.lockUntil) {
					delete(rl.clients, ip)
				}
			}
			rl.mu.Unlock()
		case <-rl.stopChan:
			return
		}
	}
}

// Stop stops the cleanup goroutine
func (rl *RateLimiter) Stop() {
	close(rl.stopChan)
}

// GetClientIP extracts the client IP from the request
func GetClientIP(r *http.Request) string {
	// Check X-Forwarded-For header first (for proxies)
	if xff := r.Header.Get("X-Forwarded-For"); xff != "" {
		// Take the first IP in the chain
		for i := 0; i < len(xff); i++ {
			if xff[i] == ',' {
				return xff[:i]
			}
		}
		return xff
	}

	// Check X-Real-IP header
	if xri := r.Header.Get("X-Real-IP"); xri != "" {
		return xri
	}

	// Fall back to RemoteAddr
	// Remove port number
	addr := r.RemoteAddr
	for i := len(addr) - 1; i >= 0; i-- {
		if addr[i] == ':' {
			return addr[:i]
		}
	}
	return addr
}

// Allow checks if a request should be allowed
func (rl *RateLimiter) Allow(ip string) bool {
	rl.mu.Lock()
	defer rl.mu.Unlock()

	now := time.Now()
	client, exists := rl.clients[ip]

	if !exists {
		rl.clients[ip] = &clientLimiter{
			tokens:    rl.rate - 1,
			lastReset: now,
		}
		return true
	}

	// Check if locked out
	if now.Before(client.lockUntil) {
		return false
	}

	// Reset tokens if window has passed
	if now.Sub(client.lastReset) >= rl.window {
		client.tokens = rl.rate
		client.lastReset = now
	}

	// Check if tokens available
	if client.tokens > 0 {
		client.tokens--
		return true
	}

	return false
}

// RecordFailure records a failed authentication attempt
func (rl *RateLimiter) RecordFailure(ip string, threshold int, lockoutDuration time.Duration) bool {
	rl.mu.Lock()
	defer rl.mu.Unlock()

	client, exists := rl.clients[ip]
	if !exists {
		client = &clientLimiter{
			tokens:    rl.rate,
			lastReset: time.Now(),
		}
		rl.clients[ip] = client
	}

	client.failures++

	if client.failures >= threshold {
		client.lockUntil = time.Now().Add(lockoutDuration)
		client.failures = 0 // Reset counter after lockout
		return true         // Locked out
	}

	return false
}

// ResetFailures resets the failure counter for an IP (called on successful auth)
func (rl *RateLimiter) ResetFailures(ip string) {
	rl.mu.Lock()
	defer rl.mu.Unlock()

	if client, exists := rl.clients[ip]; exists {
		client.failures = 0
	}
}

// IsLockedOut checks if an IP is currently locked out
func (rl *RateLimiter) IsLockedOut(ip string) bool {
	rl.mu.RLock()
	defer rl.mu.RUnlock()

	if client, exists := rl.clients[ip]; exists {
		return time.Now().Before(client.lockUntil)
	}
	return false
}

// RateLimitMiddleware returns a middleware that rate limits requests
func RateLimitMiddleware(limiter *RateLimiter) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			ip := GetClientIP(r)

			// Check if locked out
			if limiter.IsLockedOut(ip) {
				w.Header().Set("Retry-After", "900") // 15 minutes
				http.Error(w, "Too many failed attempts. Please try again later.", http.StatusTooManyRequests)
				return
			}

			// Check rate limit
			if !limiter.Allow(ip) {
				w.Header().Set("Retry-After", "60")
				http.Error(w, "Rate limit exceeded. Please slow down.", http.StatusTooManyRequests)
				return
			}

			next.ServeHTTP(w, r)
		})
	}
}
