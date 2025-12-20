package middleware

import (
	"fmt"
	"net/http"
	"os"
)

// License represents a software license
type License interface {
	IsValid() bool
	HasFeature(featureName string) bool
}

// LicenseVerifier defines the interface for license verification
type LicenseVerifier interface {
	VerifyLicense(licenseFile string) (License, error)
}

// findLicenseFileFn is the function type for finding license files
type findLicenseFileFn func() (string, error)

// findLicenseFile attempts to find the license file in common locations
var findLicenseFile findLicenseFileFn = func() (string, error) {
	// Common locations to check for the license file
	locations := []string{
		"license.xml",          // Current directory
		"./config/license.xml", // Config subdirectory
		"../license.xml",       // Parent directory
	}

	// Check each location
	for _, location := range locations {
		if _, err := os.Stat(location); err == nil {
			return location, nil
		}
	}

	return "", fmt.Errorf("license file not found in common locations")
}

// RequireLicenseFeature creates a middleware that checks if a feature is enabled in the license
func RequireLicenseFeature(verifier LicenseVerifier, featureName string) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			// Find and verify license
			licenseFile, err := findLicenseFile()
			if err != nil {
				http.Error(w, "License not found", http.StatusForbidden)
				return
			}

			lic, err := verifier.VerifyLicense(licenseFile)
			if err != nil {
				http.Error(w, "Invalid license", http.StatusForbidden)
				return
			}

			// Check if license is valid and feature is enabled
			if !lic.IsValid() || !lic.HasFeature(featureName) {
				http.Error(w, fmt.Sprintf("Feature '%s' not available in your license", featureName), http.StatusForbidden)
				return
			}

			// Feature is enabled, proceed to the next handler
			next.ServeHTTP(w, r)
		})
	}
}

// RequireAnyLicenseFeature creates a middleware that checks if any of the specified features are enabled in the license
func RequireAnyLicenseFeature(verifier LicenseVerifier, featureNames ...string) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			// Find and verify license
			licenseFile, err := findLicenseFile()
			if err != nil {
				http.Error(w, "License not found", http.StatusForbidden)
				return
			}

			lic, err := verifier.VerifyLicense(licenseFile)
			if err != nil {
				http.Error(w, "Invalid license", http.StatusForbidden)
				return
			}

			// Check if license is valid
			if !lic.IsValid() {
				http.Error(w, "Invalid license", http.StatusForbidden)
				return
			}

			// Check if any of the features are enabled
			for _, featureName := range featureNames {
				if lic.HasFeature(featureName) {
					// At least one feature is enabled, proceed to the next handler
					next.ServeHTTP(w, r)
					return
				}
			}

			// None of the features are enabled
			http.Error(w, "Required features not available in your license", http.StatusForbidden)
		})
	}
}
