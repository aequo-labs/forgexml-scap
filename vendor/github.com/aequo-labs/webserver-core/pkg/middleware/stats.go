package middleware

import (
	"bufio"
	"net/http"
	"os"
	"runtime"
	"strconv"
	"strings"
	"sync"
	"sync/atomic"
	"time"

	"github.com/dustin/go-humanize"
)

// RequestStats tracks request statistics for monitoring
type RequestStats struct {
	mu sync.RWMutex

	// Request counters
	totalRequests  atomic.Uint64
	requestsPerMin atomic.Uint64
	currentMinute  time.Time

	// Error counters
	total401s       atomic.Uint64
	total404s       atomic.Uint64
	total500s       atomic.Uint64
	recentErrors    []ErrorRecord
	maxRecentErrors int

	// Timing
	startTime time.Time
	stopChan  chan struct{}

	// Per-minute history for graphs (last 60 minutes)
	requestHistory []MinuteStats
	historyMu      sync.RWMutex

	// Resource history for charts (sampled every 10 seconds, keep last 60 samples = 10 min)
	resourceHistory []ResourceSample
	resourceMu      sync.RWMutex
}

// MinuteStats holds stats for a single minute
type MinuteStats struct {
	Timestamp time.Time `json:"timestamp"`
	Requests  uint64    `json:"requests"`
	Errors401 uint64    `json:"errors_401"`
	Errors404 uint64    `json:"errors_404"`
	Errors500 uint64    `json:"errors_500"`
}

// ResourceSample holds a point-in-time CPU/memory sample
type ResourceSample struct {
	Timestamp  int64   `json:"timestamp"` // Unix timestamp in seconds
	CPUPercent float64 `json:"cpu"`
	MemoryMB   float64 `json:"memory"`
}

// ErrorRecord represents a single error event
type ErrorRecord struct {
	Timestamp  time.Time `json:"timestamp"`
	StatusCode int       `json:"status_code"`
	Path       string    `json:"path"`
	Method     string    `json:"method"`
	ClientIP   string    `json:"client_ip"`
	UserAgent  string    `json:"user_agent,omitempty"`
}

// StatsSnapshot represents a point-in-time snapshot of stats
type StatsSnapshot struct {
	// Uptime
	Uptime        string `json:"uptime"`
	UptimeSeconds int64  `json:"uptime_seconds"`
	StartTime     string `json:"start_time"`

	// Request counts
	TotalRequests  uint64 `json:"total_requests"`
	RequestsPerMin uint64 `json:"requests_per_minute"`

	// Error counts
	Total401Errors uint64 `json:"total_401_errors"`
	Total404Errors uint64 `json:"total_404_errors"`
	Total500Errors uint64 `json:"total_500_errors"`

	// Recent errors (last N)
	RecentErrors []ErrorRecord `json:"recent_errors,omitempty"`

	// System stats - Memory
	NumGoroutines int    `json:"num_goroutines"`
	MemoryAlloc   string `json:"memory_alloc"`
	MemoryTotal   string `json:"memory_total"`
	MemorySys     string `json:"memory_sys"`
	NumGC         uint32 `json:"num_gc"`

	// System stats - CPU & Load
	NumCPU    int     `json:"num_cpu"`
	LoadAvg1  float64 `json:"load_avg_1"`
	LoadAvg5  float64 `json:"load_avg_5"`
	LoadAvg15 float64 `json:"load_avg_15"`
	CPUUsage  float64 `json:"cpu_usage_percent"`

	// Historical data for graphs
	History []MinuteStats `json:"history,omitempty"`

	// Resource history for CPU/memory charts (last 10 minutes, sampled every 10s)
	ResourceHistory []ResourceSample `json:"resource_history,omitempty"`
}

// NewRequestStats creates a new stats tracker
func NewRequestStats() *RequestStats {
	rs := &RequestStats{
		startTime:       time.Now(),
		currentMinute:   time.Now().Truncate(time.Minute),
		maxRecentErrors: 50,
		recentErrors:    make([]ErrorRecord, 0, 50),
		requestHistory:  make([]MinuteStats, 0, 60),
		resourceHistory: make([]ResourceSample, 0, 60),
		stopChan:        make(chan struct{}),
	}

	// Start background goroutine to rotate minute stats
	go rs.rotateLoop()

	// Start background goroutine to sample resources
	go rs.resourceSampleLoop()

	return rs
}

// rotateLoop handles per-minute stat rotation
func (rs *RequestStats) rotateLoop() {
	ticker := time.NewTicker(time.Minute)
	defer ticker.Stop()

	for {
		select {
		case <-ticker.C:
			rs.rotateMinute()
		case <-rs.stopChan:
			return
		}
	}
}

// rotateMinute saves current minute stats and resets counters
func (rs *RequestStats) rotateMinute() {
	rs.historyMu.Lock()
	defer rs.historyMu.Unlock()

	// Save current minute stats
	stats := MinuteStats{
		Timestamp: rs.currentMinute,
		Requests:  rs.requestsPerMin.Load(),
		// Note: we'd need separate per-minute error counters for accurate per-minute errors
	}

	rs.requestHistory = append(rs.requestHistory, stats)

	// Keep only last 60 minutes
	if len(rs.requestHistory) > 60 {
		rs.requestHistory = rs.requestHistory[1:]
	}

	// Reset per-minute counter
	rs.requestsPerMin.Store(0)
	rs.currentMinute = time.Now().Truncate(time.Minute)
}

// resourceSampleLoop samples CPU and memory every 10 seconds
func (rs *RequestStats) resourceSampleLoop() {
	ticker := time.NewTicker(10 * time.Second)
	defer ticker.Stop()

	// Take initial sample
	rs.sampleResources()

	for {
		select {
		case <-ticker.C:
			rs.sampleResources()
		case <-rs.stopChan:
			return
		}
	}
}

// sampleResources takes a CPU/memory sample and adds to history
func (rs *RequestStats) sampleResources() {
	var memStats runtime.MemStats
	runtime.ReadMemStats(&memStats)

	sample := ResourceSample{
		Timestamp:  time.Now().Unix(),
		CPUPercent: getCPUUsage(),
		MemoryMB:   float64(memStats.Alloc) / 1024 / 1024,
	}

	rs.resourceMu.Lock()
	rs.resourceHistory = append(rs.resourceHistory, sample)
	// Keep last 60 samples (10 minutes at 10-second intervals)
	if len(rs.resourceHistory) > 60 {
		rs.resourceHistory = rs.resourceHistory[1:]
	}
	rs.resourceMu.Unlock()
}

// Stop stops the background goroutines
func (rs *RequestStats) Stop() {
	close(rs.stopChan)
}

// RecordRequest increments request counters
func (rs *RequestStats) RecordRequest() {
	rs.totalRequests.Add(1)
	rs.requestsPerMin.Add(1)
}

// Record401 records an unauthorized access attempt
func (rs *RequestStats) Record401(r *http.Request) {
	rs.total401s.Add(1)
	rs.recordError(401, r)
}

// Record404 records a not found error
func (rs *RequestStats) Record404(r *http.Request) {
	rs.total404s.Add(1)
	rs.recordError(404, r)
}

// Record500 records a server error
func (rs *RequestStats) Record500(r *http.Request) {
	rs.total500s.Add(1)
	rs.recordError(500, r)
}

// recordError adds an error to the recent errors list
func (rs *RequestStats) recordError(statusCode int, r *http.Request) {
	record := ErrorRecord{
		Timestamp:  time.Now(),
		StatusCode: statusCode,
		Path:       r.URL.Path,
		Method:     r.Method,
		ClientIP:   GetClientIP(r),
		UserAgent:  r.Header.Get("User-Agent"),
	}

	rs.mu.Lock()
	defer rs.mu.Unlock()

	rs.recentErrors = append(rs.recentErrors, record)

	// Trim to max size
	if len(rs.recentErrors) > rs.maxRecentErrors {
		rs.recentErrors = rs.recentErrors[1:]
	}
}

// GetSnapshot returns current stats snapshot
func (rs *RequestStats) GetSnapshot(includeHistory bool) StatsSnapshot {
	var memStats runtime.MemStats
	runtime.ReadMemStats(&memStats)

	rs.mu.RLock()
	recentErrors := make([]ErrorRecord, len(rs.recentErrors))
	copy(recentErrors, rs.recentErrors)
	rs.mu.RUnlock()

	uptime := time.Since(rs.startTime)

	// Get load averages (Linux only, returns 0 on other platforms)
	load1, load5, load15 := getLoadAverage()

	// Get CPU usage
	cpuUsage := getCPUUsage()

	snapshot := StatsSnapshot{
		Uptime:         humanize.RelTime(rs.startTime, time.Now(), "", ""),
		UptimeSeconds:  int64(uptime.Seconds()),
		StartTime:      rs.startTime.Format(time.RFC3339),
		TotalRequests:  rs.totalRequests.Load(),
		RequestsPerMin: rs.requestsPerMin.Load(),
		Total401Errors: rs.total401s.Load(),
		Total404Errors: rs.total404s.Load(),
		Total500Errors: rs.total500s.Load(),
		RecentErrors:   recentErrors,
		NumGoroutines:  runtime.NumGoroutine(),
		MemoryAlloc:    humanize.Bytes(memStats.Alloc),
		MemoryTotal:    humanize.Bytes(memStats.TotalAlloc),
		MemorySys:      humanize.Bytes(memStats.Sys),
		NumGC:          memStats.NumGC,
		NumCPU:         runtime.NumCPU(),
		LoadAvg1:       load1,
		LoadAvg5:       load5,
		LoadAvg15:      load15,
		CPUUsage:       cpuUsage,
	}

	if includeHistory {
		rs.historyMu.RLock()
		snapshot.History = make([]MinuteStats, len(rs.requestHistory))
		copy(snapshot.History, rs.requestHistory)
		rs.historyMu.RUnlock()
	}

	// Always include resource history for charts
	rs.resourceMu.RLock()
	snapshot.ResourceHistory = make([]ResourceSample, len(rs.resourceHistory))
	copy(snapshot.ResourceHistory, rs.resourceHistory)
	rs.resourceMu.RUnlock()

	return snapshot
}

// responseRecorder wraps http.ResponseWriter to capture status code
type responseRecorder struct {
	http.ResponseWriter
	statusCode int
	written    bool
}

func (rr *responseRecorder) WriteHeader(code int) {
	if !rr.written {
		rr.statusCode = code
		rr.written = true
	}
	rr.ResponseWriter.WriteHeader(code)
}

func (rr *responseRecorder) Write(b []byte) (int, error) {
	if !rr.written {
		rr.statusCode = http.StatusOK
		rr.written = true
	}
	return rr.ResponseWriter.Write(b)
}

// StatsMiddleware returns a middleware that tracks request statistics
func StatsMiddleware(stats *RequestStats) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			// Record the request
			stats.RecordRequest()

			// Wrap response writer to capture status code
			recorder := &responseRecorder{
				ResponseWriter: w,
				statusCode:     http.StatusOK,
			}

			// Call next handler
			next.ServeHTTP(recorder, r)

			// Record errors based on status code
			switch recorder.statusCode {
			case http.StatusUnauthorized:
				stats.Record401(r)
			case http.StatusNotFound:
				stats.Record404(r)
			case http.StatusInternalServerError:
				stats.Record500(r)
			}
		})
	}
}

// getLoadAverage reads load averages from /proc/loadavg (Linux only)
func getLoadAverage() (load1, load5, load15 float64) {
	file, err := os.Open("/proc/loadavg")
	if err != nil {
		return 0, 0, 0
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	if scanner.Scan() {
		fields := strings.Fields(scanner.Text())
		if len(fields) >= 3 {
			load1, _ = strconv.ParseFloat(fields[0], 64)
			load5, _ = strconv.ParseFloat(fields[1], 64)
			load15, _ = strconv.ParseFloat(fields[2], 64)
		}
	}
	return load1, load5, load15
}

// cpuStats holds CPU timing info for calculating usage
type cpuStats struct {
	user, nice, system, idle, iowait, irq, softirq, steal uint64
}

// readCPUStats reads CPU stats from /proc/stat
func readCPUStats() cpuStats {
	file, err := os.Open("/proc/stat")
	if err != nil {
		return cpuStats{}
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		if strings.HasPrefix(line, "cpu ") {
			fields := strings.Fields(line)
			if len(fields) >= 8 {
				var stats cpuStats
				stats.user, _ = strconv.ParseUint(fields[1], 10, 64)
				stats.nice, _ = strconv.ParseUint(fields[2], 10, 64)
				stats.system, _ = strconv.ParseUint(fields[3], 10, 64)
				stats.idle, _ = strconv.ParseUint(fields[4], 10, 64)
				stats.iowait, _ = strconv.ParseUint(fields[5], 10, 64)
				stats.irq, _ = strconv.ParseUint(fields[6], 10, 64)
				stats.softirq, _ = strconv.ParseUint(fields[7], 10, 64)
				if len(fields) >= 9 {
					stats.steal, _ = strconv.ParseUint(fields[8], 10, 64)
				}
				return stats
			}
		}
	}
	return cpuStats{}
}

// getCPUUsage calculates CPU usage percentage over a brief sample period
func getCPUUsage() float64 {
	stats1 := readCPUStats()
	time.Sleep(100 * time.Millisecond)
	stats2 := readCPUStats()

	// Calculate totals
	total1 := stats1.user + stats1.nice + stats1.system + stats1.idle + stats1.iowait + stats1.irq + stats1.softirq + stats1.steal
	total2 := stats2.user + stats2.nice + stats2.system + stats2.idle + stats2.iowait + stats2.irq + stats2.softirq + stats2.steal

	totalDiff := float64(total2 - total1)
	if totalDiff == 0 {
		return 0
	}

	idleDiff := float64((stats2.idle + stats2.iowait) - (stats1.idle + stats1.iowait))
	usage := (1.0 - idleDiff/totalDiff) * 100.0

	// Round to 2 decimal places
	return float64(int(usage*100)) / 100
}
