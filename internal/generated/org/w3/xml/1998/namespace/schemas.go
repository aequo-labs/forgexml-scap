// Package pkg_1998namespace schema metadata and embedded XSD files
// This file is auto-generated - do not edit manually

package pkg_1998namespace

import "embed"

//go:embed schemas/*.xsd
var Schemas embed.FS

// PackageMetadata contains information about this generated package
type PackageMetadata struct {
	Namespace   string   // XSD target namespace
	SourceXSD   string   // Primary XSD source file
	SchemaFiles []string // All XSD files used to generate this package
	GeneratedBy string   // Tool version
}

// Metadata provides schema information for this package
var Metadata = PackageMetadata{
	Namespace:   "http://www.w3.org/XML/1998/namespace",
	SourceXSD:   "xml.xsd",
	SchemaFiles: []string{
		"/home/mmcnew/repos/forgexml-scap/schemas/common/xml.xsd",
	},
	GeneratedBy: "forgexml v1.0.0",
}
