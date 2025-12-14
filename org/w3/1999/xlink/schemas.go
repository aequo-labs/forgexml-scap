// Package pkg_1999xlink schema metadata and embedded XSD files
// This file is auto-generated - do not edit manually

package pkg_1999xlink

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
	Namespace:   "http://www.w3.org/1999/xlink",
	SourceXSD:   "xlink.xsd",
	SchemaFiles: []string{
		"/home/mmcnew/repos/forgexml-scap/schemas/arf/asset-identification_1.1.0.xsd",
		"/home/mmcnew/repos/forgexml-scap/schemas/arf/asset-reporting-format_1.0.0-ea1.xsd",
		"/home/mmcnew/repos/forgexml-scap/schemas/arf/asset-reporting-format_1.1.0.xsd",
		"/home/mmcnew/repos/forgexml-scap/schemas/arf/cpe-naming_2.3.xsd",
		"/home/mmcnew/repos/forgexml-scap/schemas/arf/reporting-core_1.1.0.xsd",
		"/home/mmcnew/repos/forgexml-scap/schemas/arf/xAL.xsd",
		"/home/mmcnew/repos/forgexml-scap/schemas/arf/xNL.xsd",
		"/home/mmcnew/repos/forgexml-scap/schemas/arf/xlink.xsd",
		"/home/mmcnew/repos/forgexml-scap/schemas/arf/xml.xsd",
	},
	GeneratedBy: "forgexml v1.0.0",
}
