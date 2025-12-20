package https

import (
	"crypto/rand"
	"crypto/rsa"
	"crypto/tls"
	"crypto/x509"
	"crypto/x509/pkix"
	"encoding/pem"
	"fmt"
	"math/big"
	"net"
	"os"
	"path/filepath"
	"time"
)

const (
	// DefaultCertDir is the default directory where HTTPS certificates are stored
	DefaultCertDir = "certs/https"

	// DefaultKeySize is the default RSA key size in bits for HTTPS certificates
	DefaultKeySize = 2048

	// DefaultCertValidity is the default HTTPS certificate validity period in years
	DefaultCertValidity = 2
)

// Config holds the configuration for HTTPS
type Config struct {
	CertFile     string
	KeyFile      string
	Enabled      bool
	Organization string
	CertDir      string
}

// NewConfig creates a new HTTPS configuration with defaults
func NewConfig() *Config {
	return &Config{
		CertFile:     filepath.Join(DefaultCertDir, "cert.pem"),
		KeyFile:      filepath.Join(DefaultCertDir, "key.pem"),
		Enabled:      false,
		Organization: "WebServer-Core",
		CertDir:      DefaultCertDir,
	}
}

// NewConfigWithOptions creates a new HTTPS configuration with custom options
func NewConfigWithOptions(certDir, organization string) *Config {
	return &Config{
		CertFile:     filepath.Join(certDir, "cert.pem"),
		KeyFile:      filepath.Join(certDir, "key.pem"),
		Enabled:      false,
		Organization: organization,
		CertDir:      certDir,
	}
}

// CertificatesExist checks if the HTTPS certificates exist
func (c *Config) CertificatesExist() bool {
	_, certErr := os.Stat(c.CertFile)
	_, keyErr := os.Stat(c.KeyFile)
	return certErr == nil && keyErr == nil
}

// GenerateCertificates generates new HTTPS certificates
func (c *Config) GenerateCertificates() error {
	// Create the certificates directory if it doesn't exist
	if err := os.MkdirAll(c.CertDir, 0755); err != nil {
		return fmt.Errorf("failed to create HTTPS certificates directory: %w", err)
	}

	// Generate a new private key
	privateKey, err := rsa.GenerateKey(rand.Reader, DefaultKeySize)
	if err != nil {
		return fmt.Errorf("failed to generate RSA key pair: %w", err)
	}

	// Generate a serial number for the certificate
	serialNumber, err := rand.Int(rand.Reader, new(big.Int).Lsh(big.NewInt(1), 128))
	if err != nil {
		return fmt.Errorf("failed to generate serial number: %w", err)
	}

	// Get the local hostname
	hostname, err := os.Hostname()
	if err != nil {
		hostname = "localhost"
	}

	// Prepare certificate template
	now := time.Now()
	template := x509.Certificate{
		SerialNumber: serialNumber,
		Subject: pkix.Name{
			Organization: []string{c.Organization},
			CommonName:   hostname,
		},
		NotBefore:             now,
		NotAfter:              now.AddDate(DefaultCertValidity, 0, 0),
		KeyUsage:              x509.KeyUsageDigitalSignature | x509.KeyUsageKeyEncipherment,
		ExtKeyUsage:           []x509.ExtKeyUsage{x509.ExtKeyUsageServerAuth},
		BasicConstraintsValid: true,
		IsCA:                  false,
	}

	// Add localhost and local IP addresses to the certificate
	template.DNSNames = append(template.DNSNames, "localhost")
	addrs, err := net.InterfaceAddrs()
	if err == nil {
		for _, addr := range addrs {
			if ipnet, ok := addr.(*net.IPNet); ok && !ipnet.IP.IsLoopback() {
				if ipnet.IP.To4() != nil {
					template.IPAddresses = append(template.IPAddresses, ipnet.IP)
				}
			}
		}
	}
	// Always add loopback addresses
	template.IPAddresses = append(template.IPAddresses, net.ParseIP("127.0.0.1"), net.ParseIP("::1"))

	// Create the certificate
	certBytes, err := x509.CreateCertificate(rand.Reader, &template, &template, &privateKey.PublicKey, privateKey)
	if err != nil {
		return fmt.Errorf("failed to create certificate: %w", err)
	}

	// Save the private key to a file
	keyFile, err := os.Create(c.KeyFile)
	if err != nil {
		return fmt.Errorf("failed to create private key file: %w", err)
	}
	defer keyFile.Close()

	keyPEM := &pem.Block{
		Type:  "RSA PRIVATE KEY",
		Bytes: x509.MarshalPKCS1PrivateKey(privateKey),
	}
	if err := pem.Encode(keyFile, keyPEM); err != nil {
		return fmt.Errorf("failed to encode private key to PEM: %w", err)
	}

	// Save the certificate to a file
	certFile, err := os.Create(c.CertFile)
	if err != nil {
		return fmt.Errorf("failed to create certificate file: %w", err)
	}
	defer certFile.Close()

	certPEM := &pem.Block{
		Type:  "CERTIFICATE",
		Bytes: certBytes,
	}
	if err := pem.Encode(certFile, certPEM); err != nil {
		return fmt.Errorf("failed to encode certificate to PEM: %w", err)
	}

	return nil
}

// LoadTLSConfig loads the TLS configuration for HTTPS
func (c *Config) LoadTLSConfig() (*tls.Config, error) {
	cert, err := tls.LoadX509KeyPair(c.CertFile, c.KeyFile)
	if err != nil {
		return nil, fmt.Errorf("failed to load X509 key pair: %w", err)
	}

	return &tls.Config{
		Certificates: []tls.Certificate{cert},
		MinVersion:   tls.VersionTLS12,
	}, nil
}

// EnsureCertificates ensures that HTTPS certificates exist, generating them if necessary
func (c *Config) EnsureCertificates() error {
	if !c.CertificatesExist() {
		return c.GenerateCertificates()
	}
	return nil
}
