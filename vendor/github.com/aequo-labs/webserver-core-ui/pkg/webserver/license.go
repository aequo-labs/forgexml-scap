package webserver

import (
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"path/filepath"
	"runtime"
	"strings"
	"sync"
	"time"

	"github.com/aequo-labs/legcheck/pkg/license"
)

// LicenseConfig contains configuration for license management
type LicenseConfig struct {
	ProductName     string // Product name to validate against
	LicenseServer   string // URL of the indie-legs license server
	LicenseAPIUser  string // API username for license server
	LicenseAPIPass  string // API password for license server
	LicenseFilePath string // Path to store the license file (default: data/license.xml)

	// Proxy mode fields - for UIs that manage licenses on a backend server
	ProxyMode bool   // If true, proxy license operations to ProxyURL instead of local management
	ProxyURL  string // Backend server URL for proxy mode (e.g., http://localhost:8080)
}

// LicenseStatus represents the current license status
type LicenseStatus struct {
	Licensed     bool     `json:"licensed"`
	Activated    bool     `json:"activated"`
	Expired      bool     `json:"expired"`
	GracePeriod  bool     `json:"graceperiod"`
	LicenseID    string   `json:"licenseId,omitempty"`
	CustomerName string   `json:"customerName,omitempty"`
	ProductName  string   `json:"productName,omitempty"`
	LicenseType  string   `json:"licenseType,omitempty"`
	ExpiryDate   string   `json:"expiryDate,omitempty"`
	MachineID    string   `json:"machineId,omitempty"`
	Features     []string `json:"features,omitempty"`
	Message      string   `json:"message,omitempty"`
}

// LicenseManager handles license operations
type LicenseManager struct {
	config    *LicenseConfig
	license   *license.License
	machineID string
	activated bool
	mu        sync.RWMutex
	logger    interface{ Debug(string, ...interface{}) }
}

// NewLicenseManager creates a new license manager
func NewLicenseManager(config *LicenseConfig, logger interface{ Debug(string, ...interface{}) }) *LicenseManager {
	if config.LicenseFilePath == "" {
		config.LicenseFilePath = "data/license.xml"
	}

	lm := &LicenseManager{
		config:    config,
		machineID: generateMachineID(),
		logger:    logger,
	}

	// Try to load existing license
	lm.loadLicense()

	return lm
}

// generateMachineID creates a unique machine identifier
func generateMachineID() string {
	// Combine hostname, OS, and architecture for a semi-unique ID
	hostname, _ := os.Hostname()
	data := fmt.Sprintf("%s-%s-%s", hostname, runtime.GOOS, runtime.GOARCH)

	// Add MAC address or other hardware info if available
	// For now, use a hash of the basic info
	hash := sha256.Sum256([]byte(data))
	return hex.EncodeToString(hash[:])[:32]
}

// loadLicense loads the license from file
func (lm *LicenseManager) loadLicense() error {
	lm.mu.Lock()
	defer lm.mu.Unlock()

	lic, err := license.LoadFromFile(lm.config.LicenseFilePath)
	if err != nil {
		if lm.logger != nil {
			lm.logger.Debug("No license file found or invalid", "path", lm.config.LicenseFilePath, "error", err)
		}
		return err
	}

	// Validate the license
	if lm.config.ProductName != "" && lic.Product.Name != lm.config.ProductName {
		if lm.logger != nil {
			lm.logger.Debug("License product mismatch", "expected", lm.config.ProductName, "got", lic.Product.Name)
		}
		return fmt.Errorf("license is for product '%s', not '%s'", lic.Product.Name, lm.config.ProductName)
	}

	lm.license = lic
	if lm.logger != nil {
		lm.logger.Debug("License loaded", "id", lic.ID, "customer", lic.CustomerName, "expires", lic.ExpiryDate)
	}

	return nil
}

// GetStatus returns the current license status
func (lm *LicenseManager) GetStatus() LicenseStatus {
	lm.mu.RLock()
	defer lm.mu.RUnlock()

	status := LicenseStatus{
		MachineID: lm.machineID,
	}

	if lm.license == nil {
		status.Licensed = false
		status.Message = "No license installed"
		return status
	}

	status.Licensed = true
	status.LicenseID = lm.license.ID
	status.CustomerName = lm.license.CustomerName
	status.ProductName = lm.license.Product.Name
	status.LicenseType = lm.license.Product.Type
	status.ExpiryDate = lm.license.ExpiryDate.Format(time.RFC3339)
	status.Activated = lm.activated

	// Check expiry and grace period
	if time.Now().After(lm.license.ExpiryDate) {
		status.Expired = true
		status.Message = "License has expired"
	} else {
		daysLeft := int(time.Until(lm.license.ExpiryDate).Hours() / 24)
		if daysLeft <= 7 {
			// Grace period: 7 days or less until expiry
			status.GracePeriod = true
			status.Message = fmt.Sprintf("License expires in %d days - renew soon!", daysLeft)
		} else if daysLeft <= 30 {
			status.Message = fmt.Sprintf("License expires in %d days", daysLeft)
		} else {
			status.Message = fmt.Sprintf("Valid until %s", lm.license.ExpiryDate.Format("January 2, 2006"))
		}
	}

	// Extract features
	for _, f := range lm.license.Features {
		status.Features = append(status.Features, f.Name)
	}

	return status
}

// InstallLicense downloads and installs a license by key
func (lm *LicenseManager) InstallLicense(licenseKey string) error {
	if lm.config.LicenseServer == "" {
		return fmt.Errorf("license server not configured")
	}

	// Fetch license XML from server
	url := fmt.Sprintf("%s/api/licenses/%s/xml", strings.TrimSuffix(lm.config.LicenseServer, "/"), licenseKey)

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return fmt.Errorf("failed to create request: %w", err)
	}

	if lm.config.LicenseAPIUser != "" {
		req.SetBasicAuth(lm.config.LicenseAPIUser, lm.config.LicenseAPIPass)
	}

	client := &http.Client{Timeout: 30 * time.Second}
	resp, err := client.Do(req)
	if err != nil {
		return fmt.Errorf("failed to fetch license: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		body, _ := io.ReadAll(resp.Body)
		return fmt.Errorf("license server returned %d: %s", resp.StatusCode, string(body))
	}

	var result struct {
		Success   bool   `json:"success"`
		LicenseID string `json:"licenseId"`
		XML       string `json:"xml"`
		Error     string `json:"error"`
	}

	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return fmt.Errorf("failed to decode response: %w", err)
	}

	if !result.Success {
		return fmt.Errorf("license server error: %s", result.Error)
	}

	// Parse and validate the license
	lic, err := license.FromXML(result.XML)
	if err != nil {
		return fmt.Errorf("invalid license format: %w", err)
	}

	// Validate product name if configured
	if lm.config.ProductName != "" && lic.Product.Name != lm.config.ProductName {
		return fmt.Errorf("license is for product '%s', not '%s'", lic.Product.Name, lm.config.ProductName)
	}

	// Save license to file
	dir := filepath.Dir(lm.config.LicenseFilePath)
	if err := os.MkdirAll(dir, 0755); err != nil {
		return fmt.Errorf("failed to create license directory: %w", err)
	}

	if err := os.WriteFile(lm.config.LicenseFilePath, []byte(result.XML), 0644); err != nil {
		return fmt.Errorf("failed to save license: %w", err)
	}

	// Update internal state
	lm.mu.Lock()
	lm.license = lic
	lm.activated = false // Reset activation status for new license
	lm.mu.Unlock()

	if lm.logger != nil {
		lm.logger.Debug("License installed", "id", lic.ID, "customer", lic.CustomerName)
	}

	return nil
}

// Activate activates the license for this machine
func (lm *LicenseManager) Activate() error {
	lm.mu.RLock()
	if lm.license == nil {
		lm.mu.RUnlock()
		return fmt.Errorf("no license installed")
	}
	licenseID := lm.license.ID
	lm.mu.RUnlock()

	if lm.config.LicenseServer == "" {
		return fmt.Errorf("license server not configured")
	}

	url := fmt.Sprintf("%s/api/licenses/activate", strings.TrimSuffix(lm.config.LicenseServer, "/"))

	payload := map[string]string{
		"licenseId": licenseID,
		"machineId": lm.machineID,
	}
	payloadBytes, _ := json.Marshal(payload)

	req, err := http.NewRequest("POST", url, strings.NewReader(string(payloadBytes)))
	if err != nil {
		return fmt.Errorf("failed to create request: %w", err)
	}

	req.Header.Set("Content-Type", "application/json")
	// Activation endpoint is typically public (no auth required)

	client := &http.Client{Timeout: 30 * time.Second}
	resp, err := client.Do(req)
	if err != nil {
		return fmt.Errorf("failed to activate: %w", err)
	}
	defer resp.Body.Close()

	var result struct {
		Success     bool   `json:"success"`
		Message     string `json:"message"`
		Error       string `json:"error"`
		ActivatedAt string `json:"activatedAt"`
	}

	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return fmt.Errorf("failed to decode response: %w", err)
	}

	if !result.Success {
		errMsg := result.Error
		if errMsg == "" {
			errMsg = result.Message
		}
		return fmt.Errorf("activation failed: %s", errMsg)
	}

	lm.mu.Lock()
	lm.activated = true
	lm.mu.Unlock()

	if lm.logger != nil {
		lm.logger.Debug("License activated", "licenseId", licenseID, "machineId", lm.machineID)
	}

	return nil
}

// Deactivate deactivates the license for this machine
func (lm *LicenseManager) Deactivate() error {
	lm.mu.RLock()
	if lm.license == nil {
		lm.mu.RUnlock()
		return fmt.Errorf("no license installed")
	}
	licenseID := lm.license.ID
	lm.mu.RUnlock()

	if lm.config.LicenseServer == "" {
		return fmt.Errorf("license server not configured")
	}

	url := fmt.Sprintf("%s/api/licenses/deactivate", strings.TrimSuffix(lm.config.LicenseServer, "/"))

	payload := map[string]string{
		"licenseId": licenseID,
		"machineId": lm.machineID,
	}
	payloadBytes, _ := json.Marshal(payload)

	req, err := http.NewRequest("POST", url, strings.NewReader(string(payloadBytes)))
	if err != nil {
		return fmt.Errorf("failed to create request: %w", err)
	}

	req.Header.Set("Content-Type", "application/json")
	// Deactivation endpoint is typically public (no auth required)

	client := &http.Client{Timeout: 30 * time.Second}
	resp, err := client.Do(req)
	if err != nil {
		return fmt.Errorf("failed to deactivate: %w", err)
	}
	defer resp.Body.Close()

	var result struct {
		Success bool   `json:"success"`
		Message string `json:"message"`
		Error   string `json:"error"`
	}

	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return fmt.Errorf("failed to decode response: %w", err)
	}

	if !result.Success {
		errMsg := result.Error
		if errMsg == "" {
			errMsg = result.Message
		}
		return fmt.Errorf("deactivation failed: %s", errMsg)
	}

	lm.mu.Lock()
	lm.activated = false
	lm.mu.Unlock()

	if lm.logger != nil {
		lm.logger.Debug("License deactivated", "licenseId", licenseID, "machineId", lm.machineID)
	}

	return nil
}

// IsLicensed returns true if a valid license is installed
func (lm *LicenseManager) IsLicensed() bool {
	lm.mu.RLock()
	defer lm.mu.RUnlock()

	if lm.license == nil {
		return false
	}

	// Check expiry
	return time.Now().Before(lm.license.ExpiryDate)
}

// GetLicense returns the current license (may be nil)
func (lm *LicenseManager) GetLicense() *license.License {
	lm.mu.RLock()
	defer lm.mu.RUnlock()
	return lm.license
}

// GetMachineID returns the machine ID
func (lm *LicenseManager) GetMachineID() string {
	return lm.machineID
}

// IsProxyMode returns true if the license manager is operating in proxy mode
func (lm *LicenseManager) IsProxyMode() bool {
	return lm.config.ProxyMode
}

// GetProxyURL returns the proxy URL for backend requests
func (lm *LicenseManager) GetProxyURL() string {
	return lm.config.ProxyURL
}

// ProxyGetStatus fetches license status from the backend server
func (lm *LicenseManager) ProxyGetStatus() (LicenseStatus, error) {
	if lm.config.ProxyURL == "" {
		return LicenseStatus{}, fmt.Errorf("proxy URL not configured")
	}

	url := fmt.Sprintf("%s/api/license/status", strings.TrimSuffix(lm.config.ProxyURL, "/"))

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return LicenseStatus{}, fmt.Errorf("failed to create request: %w", err)
	}

	if lm.config.LicenseAPIUser != "" {
		req.SetBasicAuth(lm.config.LicenseAPIUser, lm.config.LicenseAPIPass)
	}

	client := &http.Client{Timeout: 30 * time.Second}
	resp, err := client.Do(req)
	if err != nil {
		return LicenseStatus{}, fmt.Errorf("failed to fetch status: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		body, _ := io.ReadAll(resp.Body)
		return LicenseStatus{}, fmt.Errorf("backend returned %d: %s", resp.StatusCode, string(body))
	}

	var status LicenseStatus
	if err := json.NewDecoder(resp.Body).Decode(&status); err != nil {
		return LicenseStatus{}, fmt.Errorf("failed to decode response: %w", err)
	}

	return status, nil
}

// ProxyInstallLicense proxies a license install request to the backend server
func (lm *LicenseManager) ProxyInstallLicense(licenseKey string) error {
	if lm.config.ProxyURL == "" {
		return fmt.Errorf("proxy URL not configured")
	}

	url := fmt.Sprintf("%s/api/license/install", strings.TrimSuffix(lm.config.ProxyURL, "/"))

	payload := map[string]string{"licenseKey": licenseKey}
	payloadBytes, _ := json.Marshal(payload)

	req, err := http.NewRequest("POST", url, strings.NewReader(string(payloadBytes)))
	if err != nil {
		return fmt.Errorf("failed to create request: %w", err)
	}

	req.Header.Set("Content-Type", "application/json")
	if lm.config.LicenseAPIUser != "" {
		req.SetBasicAuth(lm.config.LicenseAPIUser, lm.config.LicenseAPIPass)
	}

	client := &http.Client{Timeout: 30 * time.Second}
	resp, err := client.Do(req)
	if err != nil {
		return fmt.Errorf("failed to install license: %w", err)
	}
	defer resp.Body.Close()

	var result struct {
		Success bool   `json:"success"`
		Message string `json:"message"`
		Error   string `json:"error"`
	}

	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return fmt.Errorf("failed to decode response: %w", err)
	}

	if !result.Success {
		errMsg := result.Error
		if errMsg == "" {
			errMsg = result.Message
		}
		return fmt.Errorf("install failed: %s", errMsg)
	}

	return nil
}

// ProxyActivate proxies a license activation request to the backend server
func (lm *LicenseManager) ProxyActivate() error {
	if lm.config.ProxyURL == "" {
		return fmt.Errorf("proxy URL not configured")
	}

	url := fmt.Sprintf("%s/api/license/activate", strings.TrimSuffix(lm.config.ProxyURL, "/"))

	req, err := http.NewRequest("POST", url, nil)
	if err != nil {
		return fmt.Errorf("failed to create request: %w", err)
	}

	req.Header.Set("Content-Type", "application/json")
	if lm.config.LicenseAPIUser != "" {
		req.SetBasicAuth(lm.config.LicenseAPIUser, lm.config.LicenseAPIPass)
	}

	client := &http.Client{Timeout: 30 * time.Second}
	resp, err := client.Do(req)
	if err != nil {
		return fmt.Errorf("failed to activate: %w", err)
	}
	defer resp.Body.Close()

	var result struct {
		Success bool   `json:"success"`
		Message string `json:"message"`
		Error   string `json:"error"`
	}

	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return fmt.Errorf("failed to decode response: %w", err)
	}

	if !result.Success {
		errMsg := result.Error
		if errMsg == "" {
			errMsg = result.Message
		}
		return fmt.Errorf("activation failed: %s", errMsg)
	}

	return nil
}

// ProxyDeactivate proxies a license deactivation request to the backend server
func (lm *LicenseManager) ProxyDeactivate() error {
	if lm.config.ProxyURL == "" {
		return fmt.Errorf("proxy URL not configured")
	}

	url := fmt.Sprintf("%s/api/license/deactivate", strings.TrimSuffix(lm.config.ProxyURL, "/"))

	req, err := http.NewRequest("POST", url, nil)
	if err != nil {
		return fmt.Errorf("failed to create request: %w", err)
	}

	req.Header.Set("Content-Type", "application/json")
	if lm.config.LicenseAPIUser != "" {
		req.SetBasicAuth(lm.config.LicenseAPIUser, lm.config.LicenseAPIPass)
	}

	client := &http.Client{Timeout: 30 * time.Second}
	resp, err := client.Do(req)
	if err != nil {
		return fmt.Errorf("failed to deactivate: %w", err)
	}
	defer resp.Body.Close()

	var result struct {
		Success bool   `json:"success"`
		Message string `json:"message"`
		Error   string `json:"error"`
	}

	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return fmt.Errorf("failed to decode response: %w", err)
	}

	if !result.Success {
		errMsg := result.Error
		if errMsg == "" {
			errMsg = result.Message
		}
		return fmt.Errorf("deactivation failed: %s", errMsg)
	}

	return nil
}
