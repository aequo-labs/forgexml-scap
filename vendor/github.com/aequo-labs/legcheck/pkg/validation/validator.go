// Package validation provides license validation functionality.
package validation

import (
	"crypto"
	"crypto/rsa"
	"crypto/sha256"
	"crypto/x509"
	"encoding/base64"
	"encoding/pem"
	"fmt"
	"time"

	"github.com/aequo-labs/legcheck/pkg/license"
	"github.com/aequo-labs/legcheck/pkg/validator"
)

// EmbeddedPublicKey - RSA public key for license verification
// Available to all components (license-server-ui, go-algo, etc.)
// Can be overridden at build time with:
//   go build -ldflags "-X 'github.com/aequo-labs/legcheck/pkg/validation.EmbeddedPublicKey=YOUR_PEM_KEY'"
var EmbeddedPublicKey = `-----BEGIN PUBLIC KEY-----
MIIBIjANBgkqhkiG9w0BAQEFAAOCAQ8AMIIBCgKCAQEAvSYWAkZkr9qe/EegJXC0
ctxdnoyRUptksZibUQhqFc+th/m4ncQO6P+IBXBGlukRJKeZHrFaQjK9+R5J/18W
xuFa5OdSNFNsn9xmmlZMycWIN8flTmhfjDmbxR9X9G4f5S71TIlCrFxBkUaU9Cn7
qv8itG0nsgkj8C1r0Lix7X6ZdsQjS0XOqVnPbJpAjX6j3DuAN9flb4lsrxkzRLqs
LUYFdyX4puqTGbM7mGinXok0p5Y3LIgdNebjX9jXSh+C1IE0BwhFMJR4JPzIHIB5
tCFFWM5wHwhaEqQUNzmMdDTeNdl9MIv20+8mDPGAFK5LA8uKL8iOWnEVbzmS8pC6
MQIDAQAB
-----END PUBLIC KEY-----`

// ConcreteValidator implements the Validator interface
type ConcreteValidator struct {
	publicKey *rsa.PublicKey
}

// NewValidator creates a new license validator with the embedded public key
func NewValidator() (validator.Validator, error) {
	block, _ := pem.Decode([]byte(EmbeddedPublicKey))
	if block == nil {
		return nil, fmt.Errorf("failed to parse embedded public key PEM block")
	}

	pub, err := x509.ParsePKIXPublicKey(block.Bytes)
	if err != nil {
		return nil, fmt.Errorf("failed to parse embedded public key: %w", err)
	}

	publicKey, ok := pub.(*rsa.PublicKey)
	if !ok {
		return nil, fmt.Errorf("embedded key is not an RSA public key")
	}

	return &ConcreteValidator{
		publicKey: publicKey,
	}, nil
}

// NewValidatorFromPEM creates a new license validator from PEM-encoded public key
func NewValidatorFromPEM(pemData string) (validator.Validator, error) {
	block, _ := pem.Decode([]byte(pemData))
	if block == nil {
		return nil, fmt.Errorf("failed to parse PEM block")
	}

	pub, err := x509.ParsePKIXPublicKey(block.Bytes)
	if err != nil {
		return nil, fmt.Errorf("failed to parse public key: %w", err)
	}

	publicKey, ok := pub.(*rsa.PublicKey)
	if !ok {
		return nil, fmt.Errorf("not an RSA public key")
	}

	return &ConcreteValidator{
		publicKey: publicKey,
	}, nil
}

// LoadLicenseFromFile loads a license from a file
func (v *ConcreteValidator) LoadLicenseFromFile(filePath string) (*license.License, error) {
	return license.LoadFromFile(filePath)
}

// ValidateLicense validates a license
func (v *ConcreteValidator) ValidateLicense(lic *license.License) (bool, error) {
	if v.publicKey == nil {
		return false, fmt.Errorf("public key not available")
	}

	if lic == nil {
		return false, fmt.Errorf("license is nil")
	}

	// Check required fields
	if lic.ID == "" {
		return false, fmt.Errorf("license ID cannot be empty")
	}
	if lic.CustomerName == "" {
		return false, fmt.Errorf("customer name cannot be empty")
	}
	if lic.CustomerEmail == "" {
		return false, fmt.Errorf("customer email cannot be empty")
	}

	// Check if issue date is zero/invalid
	if lic.IssueDate.IsZero() {
		return false, fmt.Errorf("license has invalid issue date")
	}

	// Check if expiry date is zero/invalid
	if lic.ExpiryDate.IsZero() {
		return false, fmt.Errorf("license has invalid expiry date")
	}

	// Check if issue date is after expiry date
	if lic.IssueDate.After(lic.ExpiryDate) {
		return false, fmt.Errorf("license issue date cannot be after expiry date")
	}

	// Check if the license is expired
	if time.Now().After(lic.ExpiryDate) {
		return false, fmt.Errorf("license expired on %s", lic.ExpiryDate.Format(time.RFC3339))
	}

	// Check if signature is empty
	if lic.Signature == "" {
		return false, fmt.Errorf("license has no signature")
	}

	// Verify the signature using RSA
	tempLicense := *lic
	tempLicense.Signature = ""
	xmlData, err := tempLicense.ToXML()
	if err != nil {
		return false, fmt.Errorf("failed to marshal license to XML: %w", err)
	}

	// Create SHA-256 hash of the XML data
	hash := sha256.Sum256([]byte(xmlData))

	// Decode the base64-encoded signature
	signatureBytes, err := base64.StdEncoding.DecodeString(lic.Signature)
	if err != nil {
		return false, fmt.Errorf("failed to decode signature: %w", err)
	}

	// Check signature length (RSA-2048 signatures should be 256 bytes)
	expectedSigLength := v.publicKey.Size()
	if len(signatureBytes) != expectedSigLength {
		return false, fmt.Errorf("invalid signature length: expected %d bytes, got %d bytes", expectedSigLength, len(signatureBytes))
	}

	// Verify the signature using RSA
	err = rsa.VerifyPKCS1v15(v.publicKey, crypto.SHA256, hash[:], signatureBytes)
	if err != nil {
		// Signature verification failed - return false but no error
		return false, nil
	}

	return true, nil
}

// ValidateLicenseXML validates license XML data
func (v *ConcreteValidator) ValidateLicenseXML(xmlData string) (bool, *license.License, error) {
	lic, err := license.FromXML(xmlData)
	if err != nil {
		return false, nil, fmt.Errorf("failed to parse XML license: %w", err)
	}

	valid, err := v.ValidateLicense(lic)
	return valid, lic, err
}

// HasFeature checks if a license has a specific feature
func (v *ConcreteValidator) HasFeature(lic *license.License, featureName string) (bool, string) {
	for _, feature := range lic.Features {
		if feature.Name == featureName {
			return true, feature.Value
		}
	}
	return false, ""
}

// VerifyFeature verifies a license has a specific feature and it's valid
func (v *ConcreteValidator) VerifyFeature(lic *license.License, featureName string) (bool, error) {
	// First validate the license
	valid, err := v.ValidateLicense(lic)
	if err != nil || !valid {
		return false, err
	}

	// Check for the feature
	hasFeature, _ := v.HasFeature(lic, featureName)
	return hasFeature, nil
}

// IsExpired checks if a license is expired
func (v *ConcreteValidator) IsExpired(lic *license.License) bool {
	if lic == nil {
		return true
	}
	return time.Now().After(lic.ExpiryDate)
}

// IsRevoked checks if a license is revoked
func (v *ConcreteValidator) IsRevoked(lic *license.License) bool {
	if lic == nil {
		return true
	}
	return lic.RevokedAt != nil
}

// GetStatus returns the 4-state license status: valid, activated, invalid, expired
func (v *ConcreteValidator) GetStatus(lic *license.License) string {
	if lic == nil {
		return "invalid"
	}

	if v.IsExpired(lic) {
		return "expired"
	}
	if v.IsRevoked(lic) {
		return "invalid"
	}
	if lic.IsActivated() {
		return "activated"
	}
	return "valid"
}

// CreateDefaultValidator creates a validator with the embedded public key.
// This is an alias for NewValidator() for backwards compatibility.
func CreateDefaultValidator() (validator.Validator, error) {
	return NewValidator()
}
