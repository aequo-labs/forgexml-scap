// Package language2_0 schema metadata and embedded XSD files
// This file is auto-generated - do not edit manually

package language2_0

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
	Namespace:   "http://cpe.mitre.org/language/2.0",
	SourceXSD:   "cpe-language_2.0.xsd",
	SchemaFiles: []string{
		"/home/mmcnew/repos/forgexml-scap/schemas/xccdf/common/cpe-1.0.xsd",
		"/home/mmcnew/repos/forgexml-scap/schemas/xccdf/common/cpe-language_2.0.xsd",
		"/home/mmcnew/repos/forgexml-scap/schemas/xccdf/common/platform-0.2.3.xsd",
		"/home/mmcnew/repos/forgexml-scap/schemas/xccdf/common/simpledc20021212.xsd",
		"/home/mmcnew/repos/forgexml-scap/schemas/xccdf/common/xccdfp-1.1.xsd",
		"/home/mmcnew/repos/forgexml-scap/schemas/xccdf/common/xml.xsd",
	},
	GeneratedBy: "forgexml v1.0.0",
}
