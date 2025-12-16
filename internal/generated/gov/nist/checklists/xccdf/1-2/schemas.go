// Package xccdf1_2 schema metadata and embedded XSD files
// This file is auto-generated - do not edit manually

package xccdf1_2

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
	Namespace: "http://checklists.nist.gov/xccdf/1.2",
	SourceXSD: "xccdf_1.2.xsd",
	SchemaFiles: []string{
		"/home/mmcnew/repos/forgexml-scap/schemas/xccdf/common/cpe-1.0.xsd",
		"/home/mmcnew/repos/forgexml-scap/schemas/xccdf/common/cpe-language_2.0.xsd",
		"/home/mmcnew/repos/forgexml-scap/schemas/xccdf/common/platform-0.2.3.xsd",
		"/home/mmcnew/repos/forgexml-scap/schemas/xccdf/common/simpledc20021212.xsd",
		"/home/mmcnew/repos/forgexml-scap/schemas/xccdf/common/xccdfp-1.1.xsd",
		"/home/mmcnew/repos/forgexml-scap/schemas/xccdf/common/xml.xsd",
		"/home/mmcnew/repos/forgexml-scap/schemas/xccdf/schemas/common/cpe-1.0.xsd",
		"/home/mmcnew/repos/forgexml-scap/schemas/xccdf/schemas/common/cpe-language_2.0.xsd",
		"/home/mmcnew/repos/forgexml-scap/schemas/xccdf/schemas/common/platform-0.2.3.xsd",
		"/home/mmcnew/repos/forgexml-scap/schemas/xccdf/schemas/common/simpledc20021212.xsd",
		"/home/mmcnew/repos/forgexml-scap/schemas/xccdf/schemas/common/xccdfp-1.1.xsd",
		"/home/mmcnew/repos/forgexml-scap/schemas/xccdf/schemas/common/xml.xsd",
		"/home/mmcnew/repos/forgexml-scap/schemas/xccdf/schemas/xccdf/1.1/cpe-1.0.xsd",
		"/home/mmcnew/repos/forgexml-scap/schemas/xccdf/schemas/xccdf/1.1/cpe-language_2.0.xsd",
		"/home/mmcnew/repos/forgexml-scap/schemas/xccdf/schemas/xccdf/1.1/platform-0.2.3.xsd",
		"/home/mmcnew/repos/forgexml-scap/schemas/xccdf/schemas/xccdf/1.1/simpledc20021212.xsd",
		"/home/mmcnew/repos/forgexml-scap/schemas/xccdf/schemas/xccdf/1.1/xccdf-1.1.4.xsd",
		"/home/mmcnew/repos/forgexml-scap/schemas/xccdf/schemas/xccdf/1.1/xccdfp-1.1.xsd",
		"/home/mmcnew/repos/forgexml-scap/schemas/xccdf/schemas/xccdf/1.1/xml.xsd",
		"/home/mmcnew/repos/forgexml-scap/schemas/xccdf/xccdf-1.0-complete.xsd",
		"/home/mmcnew/repos/forgexml-scap/schemas/xccdf/xccdf-1.0.xsd",
		"/home/mmcnew/repos/forgexml-scap/schemas/xccdf/xccdf-1.1.2.xsd",
		"/home/mmcnew/repos/forgexml-scap/schemas/xccdf/xccdf-1.1.3.xsd",
		"/home/mmcnew/repos/forgexml-scap/schemas/xccdf/xccdf-1.1.xsd",
		"/home/mmcnew/repos/forgexml-scap/schemas/xccdf/xccdf_1.2.xsd",
	},
	GeneratedBy: "forgexml v1.0.0",
}
