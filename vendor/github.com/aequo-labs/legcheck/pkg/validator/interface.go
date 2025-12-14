// Package validator defines the interface for license validation.
package validator

import (
	"github.com/aequo-labs/legcheck/pkg/license"
)

// Validator defines the interface for license validation
type Validator interface {
	// ValidateLicense validates a license and returns whether it's valid
	ValidateLicense(license *license.License) (bool, error)

	// ValidateLicenseXML validates license XML data
	ValidateLicenseXML(xmlData string) (bool, *license.License, error)

	// LoadLicenseFromFile loads a license from a file path
	LoadLicenseFromFile(filePath string) (*license.License, error)

	// HasFeature checks if a license has a specific feature
	HasFeature(license *license.License, featureName string) (bool, string)

	// VerifyFeature verifies a license has a specific feature and it's valid
	VerifyFeature(license *license.License, featureName string) (bool, error)

	// IsExpired checks if a license is expired
	IsExpired(license *license.License) bool

	// IsRevoked checks if a license is revoked
	IsRevoked(license *license.License) bool

	// GetStatus returns the 4-state license status
	GetStatus(license *license.License) string
}
