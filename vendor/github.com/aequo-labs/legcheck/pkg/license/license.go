// Package license provides core license types for license validation.
// This package contains only pure data structures with no external dependencies.
package license

import (
	"encoding/xml"
	"fmt"
	"os"
	"strconv"
	"time"
)

// Feature represents a license feature with name and value
type Feature struct {
	Name  string `json:"name"`
	Value string `json:"value"`
}

// Product represents a product covered by a license
type Product struct {
	Name    string `json:"name"`
	Version string `json:"version"`
	Type    string `json:"type"`
}

// License represents a software license
type License struct {
	ID            string     `json:"licenseId"`
	CustomerName  string     `json:"customerName"`
	CustomerEmail string     `json:"customerEmail"`
	IssueDate     time.Time  `json:"issueDate"`
	ExpiryDate    time.Time  `json:"expiryDate"`
	RevokedAt     *time.Time `json:"revokedAt,omitempty"` // Optional revocation timestamp
	Product       Product    `json:"product"`
	Features      []Feature  `json:"features"`
	Signature     string     `json:"signature"`

	// Activation data
	MaxActivations  int                     `json:"maxActivations"` // 0 = unlimited, >0 = limited
	Activations     []BasicActivationRecord `json:"activations,omitempty"`
	ActivationCount int                     `json:"activationCount"`
	LastActivatedAt *time.Time              `json:"lastActivatedAt,omitempty"`
}

// BasicActivationRecord represents a simple license activation record
type BasicActivationRecord struct {
	LicenseID   string    `json:"licenseId"`
	MachineID   string    `json:"machineId"`
	ActivatedAt time.Time `json:"activatedAt"`
	IsActive    bool      `json:"isActive"`
}

// IsExpired checks if the license has expired
func (l *License) IsExpired() bool {
	return time.Now().After(l.ExpiryDate)
}

// IsValid checks if the license is currently valid (not expired and not revoked)
func (l *License) IsValid() bool {
	return !l.IsExpired() && l.RevokedAt == nil
}

// IsActivated checks if the license has any active activations
func (l *License) IsActivated() bool {
	for _, activation := range l.Activations {
		if activation.IsActive {
			return true
		}
	}
	return false
}

// GetStatus returns the 4-state license status: valid, activated, invalid, expired
func (l *License) GetStatus() string {
	if l.IsExpired() {
		return "expired"
	}
	if l.RevokedAt != nil {
		return "invalid"
	}
	if l.IsActivated() {
		return "activated"
	}
	return "valid"
}

// HasFeature checks if the license has a specific feature
func (l *License) HasFeature(featureName string) bool {
	for _, feature := range l.Features {
		if feature.Name == featureName {
			return true
		}
	}
	return false
}

// GetFeatureValue returns the value of a feature if it exists
func (l *License) GetFeatureValue(featureName string) (string, bool) {
	for _, feature := range l.Features {
		if feature.Name == featureName {
			return feature.Value, true
		}
	}
	return "", false
}

// GetGracePeriodDays returns the grace period in days from the license features.
// Grace period allows the application to continue functioning for a specified
// number of days after the license expires, giving customers time to renew.
// Returns 0 if no grace period is configured.
func (l *License) GetGracePeriodDays() int {
	value, found := l.GetFeatureValue("grace_days")
	if !found {
		return 0
	}
	days, err := strconv.Atoi(value)
	if err != nil || days < 0 {
		return 0
	}
	return days
}

// IsInGracePeriod checks if the license is currently within its grace period.
// Returns true if the license has expired but is still within the grace period.
// Returns false if the license is not expired, or if it's past the grace period.
func (l *License) IsInGracePeriod() bool {
	if !l.IsExpired() {
		return false // Not expired yet, so not in grace period
	}
	graceDays := l.GetGracePeriodDays()
	if graceDays == 0 {
		return false // No grace period configured
	}
	graceEnd := l.ExpiryDate.Add(time.Duration(graceDays) * 24 * time.Hour)
	return time.Now().Before(graceEnd)
}

// IsFullyExpired checks if the license has expired AND the grace period has passed.
// This is the point at which the application should stop functioning.
func (l *License) IsFullyExpired() bool {
	if !l.IsExpired() {
		return false
	}
	graceDays := l.GetGracePeriodDays()
	if graceDays == 0 {
		return true // No grace period, so fully expired
	}
	graceEnd := l.ExpiryDate.Add(time.Duration(graceDays) * 24 * time.Hour)
	return time.Now().After(graceEnd)
}

// DaysUntilFullExpiry returns the number of days until the license is fully expired
// (including grace period). Returns negative value if already fully expired.
func (l *License) DaysUntilFullExpiry() int {
	graceDays := l.GetGracePeriodDays()
	graceEnd := l.ExpiryDate.Add(time.Duration(graceDays) * 24 * time.Hour)
	duration := time.Until(graceEnd)
	return int(duration.Hours() / 24)
}

// XML Structures for license serialization

// LicenseXML is the XML representation of a license
type LicenseXML struct {
	XMLName       xml.Name    `xml:"License"`
	ID            string      `xml:"ID,attr"`
	CustomerName  string      `xml:"Customer>Name"`
	CustomerEmail string      `xml:"Customer>Email"`
	IssueDate     string      `xml:"IssueDate"`
	ExpiryDate    string      `xml:"ExpiryDate"`
	Product       ProductXML  `xml:"Product"`
	Features      FeaturesXML `xml:"Features"`
	Signature     string      `xml:"Signature"`
}

// ProductXML is the XML representation of a product
type ProductXML struct {
	Name    string `xml:"Name,attr"`
	Version string `xml:"Version,attr"`
	Type    string `xml:"Type,attr"`
}

// FeaturesXML is the XML representation of a collection of features
type FeaturesXML struct {
	Items []FeatureXML `xml:"Feature"`
}

// FeatureXML is the XML representation of a feature
type FeatureXML struct {
	Name  string `xml:"Name,attr"`
	Value string `xml:"Value,attr"`
}

// ToXML converts a License to its XML representation
func (l *License) ToXML() (string, error) {
	features := make([]FeatureXML, len(l.Features))
	for i, f := range l.Features {
		features[i] = FeatureXML{
			Name:  f.Name,
			Value: f.Value,
		}
	}

	licenseXML := LicenseXML{
		ID:            l.ID,
		CustomerName:  l.CustomerName,
		CustomerEmail: l.CustomerEmail,
		IssueDate:     l.IssueDate.Format(time.RFC3339),
		ExpiryDate:    l.ExpiryDate.Format(time.RFC3339),
		Product: ProductXML{
			Name:    l.Product.Name,
			Version: l.Product.Version,
			Type:    l.Product.Type,
		},
		Features: FeaturesXML{
			Items: features,
		},
		Signature: l.Signature,
	}

	// Marshal to XML
	xmlBytes, err := xml.MarshalIndent(licenseXML, "", "  ")
	if err != nil {
		return "", fmt.Errorf("failed to marshal license to XML: %w", err)
	}

	// Add XML header
	xmlString := xml.Header + string(xmlBytes)
	return xmlString, nil
}

// FromXML converts XML to a License
func FromXML(xmlString string) (*License, error) {
	var licenseXML LicenseXML
	if err := xml.Unmarshal([]byte(xmlString), &licenseXML); err != nil {
		return nil, fmt.Errorf("failed to unmarshal XML to license: %w", err)
	}

	// Parse dates
	issueDate, err := time.Parse(time.RFC3339, licenseXML.IssueDate)
	if err != nil {
		return nil, fmt.Errorf("failed to parse issue date: %w", err)
	}

	expiryDate, err := time.Parse(time.RFC3339, licenseXML.ExpiryDate)
	if err != nil {
		return nil, fmt.Errorf("failed to parse expiry date: %w", err)
	}

	// Convert features
	features := make([]Feature, len(licenseXML.Features.Items))
	for i, f := range licenseXML.Features.Items {
		features[i] = Feature{
			Name:  f.Name,
			Value: f.Value,
		}
	}

	// Create license
	license := &License{
		ID:            licenseXML.ID,
		CustomerName:  licenseXML.CustomerName,
		CustomerEmail: licenseXML.CustomerEmail,
		IssueDate:     issueDate,
		ExpiryDate:    expiryDate,
		Product: Product{
			Name:    licenseXML.Product.Name,
			Version: licenseXML.Product.Version,
			Type:    licenseXML.Product.Type,
		},
		Features:  features,
		Signature: licenseXML.Signature,
	}

	return license, nil
}

// LoadFromFile loads a license from an XML file
func LoadFromFile(filename string) (*License, error) {
	data, err := os.ReadFile(filename)
	if err != nil {
		return nil, fmt.Errorf("failed to read license file: %w", err)
	}

	license, err := FromXML(string(data))
	if err != nil {
		return nil, fmt.Errorf("failed to parse license XML: %w", err)
	}

	return license, nil
}

// SaveToFile saves a license to an XML file
func (l *License) SaveToFile(filename string) error {
	xmlData, err := l.ToXML()
	if err != nil {
		return fmt.Errorf("failed to marshal license to XML: %w", err)
	}

	err = os.WriteFile(filename, []byte(xmlData), 0644)
	if err != nil {
		return fmt.Errorf("failed to write license to file: %w", err)
	}

	return nil
}
