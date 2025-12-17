// Package xmlschemaoval_variables_5 schema metadata and embedded XSD files
// This file is auto-generated - do not edit manually

package xmlschemaoval_variables_5

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
	Namespace:   "http://oval.mitre.org/XMLSchema/oval-variables-5",
	SourceXSD:   "oval-variables-schema.xsd",
	SchemaFiles: []string{
		"/home/mmcnew/repos/forgexml-scap/schemas/oval/5.11.2/aix-definitions-schema.xsd",
		"/home/mmcnew/repos/forgexml-scap/schemas/oval/5.11.2/aix-system-characteristics-schema.xsd",
		"/home/mmcnew/repos/forgexml-scap/schemas/oval/5.11.2/all-oval-definitions.xsd",
		"/home/mmcnew/repos/forgexml-scap/schemas/oval/5.11.2/android-definitions-schema.xsd",
		"/home/mmcnew/repos/forgexml-scap/schemas/oval/5.11.2/android-system-characteristics-schema.xsd",
		"/home/mmcnew/repos/forgexml-scap/schemas/oval/5.11.2/apache-definitions-schema.xsd",
		"/home/mmcnew/repos/forgexml-scap/schemas/oval/5.11.2/apache-system-characteristics-schema.xsd",
		"/home/mmcnew/repos/forgexml-scap/schemas/oval/5.11.2/apple-ios-definitions-schema.xsd",
		"/home/mmcnew/repos/forgexml-scap/schemas/oval/5.11.2/apple-ios-system-characteristics-schema.xsd",
		"/home/mmcnew/repos/forgexml-scap/schemas/oval/5.11.2/asa-definitions-schema.xsd",
		"/home/mmcnew/repos/forgexml-scap/schemas/oval/5.11.2/asa-system-characteristics-schema.xsd",
		"/home/mmcnew/repos/forgexml-scap/schemas/oval/5.11.2/catos-definitions-schema.xsd",
		"/home/mmcnew/repos/forgexml-scap/schemas/oval/5.11.2/catos-system-characteristics-schema.xsd",
		"/home/mmcnew/repos/forgexml-scap/schemas/oval/5.11.2/esx-definitions-schema.xsd",
		"/home/mmcnew/repos/forgexml-scap/schemas/oval/5.11.2/esx-system-characteristics-schema.xsd",
		"/home/mmcnew/repos/forgexml-scap/schemas/oval/5.11.2/evaluation-ids.xsd",
		"/home/mmcnew/repos/forgexml-scap/schemas/oval/5.11.2/freebsd-definitions-schema.xsd",
		"/home/mmcnew/repos/forgexml-scap/schemas/oval/5.11.2/freebsd-system-characteristics-schema.xsd",
		"/home/mmcnew/repos/forgexml-scap/schemas/oval/5.11.2/hpux-definitions-schema.xsd",
		"/home/mmcnew/repos/forgexml-scap/schemas/oval/5.11.2/hpux-system-characteristics-schema.xsd",
		"/home/mmcnew/repos/forgexml-scap/schemas/oval/5.11.2/independent-definitions-schema.xsd",
		"/home/mmcnew/repos/forgexml-scap/schemas/oval/5.11.2/independent-system-characteristics-schema.xsd",
		"/home/mmcnew/repos/forgexml-scap/schemas/oval/5.11.2/ios-definitions-schema.xsd",
		"/home/mmcnew/repos/forgexml-scap/schemas/oval/5.11.2/ios-system-characteristics-schema.xsd",
		"/home/mmcnew/repos/forgexml-scap/schemas/oval/5.11.2/iosxe-definitions-schema.xsd",
		"/home/mmcnew/repos/forgexml-scap/schemas/oval/5.11.2/iosxe-system-characteristics-schema.xsd",
		"/home/mmcnew/repos/forgexml-scap/schemas/oval/5.11.2/junos-definitions-schema.xsd",
		"/home/mmcnew/repos/forgexml-scap/schemas/oval/5.11.2/junos-system-characteristics-schema.xsd",
		"/home/mmcnew/repos/forgexml-scap/schemas/oval/5.11.2/linux-definitions-schema.xsd",
		"/home/mmcnew/repos/forgexml-scap/schemas/oval/5.11.2/linux-system-characteristics-schema.xsd",
		"/home/mmcnew/repos/forgexml-scap/schemas/oval/5.11.2/macos-definitions-schema.xsd",
		"/home/mmcnew/repos/forgexml-scap/schemas/oval/5.11.2/macos-system-characteristics-schema.xsd",
		"/home/mmcnew/repos/forgexml-scap/schemas/oval/5.11.2/netconf-definitions-schema.xsd",
		"/home/mmcnew/repos/forgexml-scap/schemas/oval/5.11.2/netconf-system-characteristics-schema.xsd",
		"/home/mmcnew/repos/forgexml-scap/schemas/oval/5.11.2/oval-common-schema.xsd",
		"/home/mmcnew/repos/forgexml-scap/schemas/oval/5.11.2/oval-definitions-schema.xsd",
		"/home/mmcnew/repos/forgexml-scap/schemas/oval/5.11.2/oval-directives-schema.xsd",
		"/home/mmcnew/repos/forgexml-scap/schemas/oval/5.11.2/oval-results-schema.xsd",
		"/home/mmcnew/repos/forgexml-scap/schemas/oval/5.11.2/oval-system-characteristics-schema.xsd",
		"/home/mmcnew/repos/forgexml-scap/schemas/oval/5.11.2/oval-variables-schema.xsd",
		"/home/mmcnew/repos/forgexml-scap/schemas/oval/5.11.2/pixos-definitions-schema.xsd",
		"/home/mmcnew/repos/forgexml-scap/schemas/oval/5.11.2/pixos-system-characteristics-schema.xsd",
		"/home/mmcnew/repos/forgexml-scap/schemas/oval/5.11.2/sharepoint-definitions-schema.xsd",
		"/home/mmcnew/repos/forgexml-scap/schemas/oval/5.11.2/sharepoint-system-characteristics-schema.xsd",
		"/home/mmcnew/repos/forgexml-scap/schemas/oval/5.11.2/solaris-definitions-schema.xsd",
		"/home/mmcnew/repos/forgexml-scap/schemas/oval/5.11.2/solaris-system-characteristics-schema.xsd",
		"/home/mmcnew/repos/forgexml-scap/schemas/oval/5.11.2/unix-definitions-schema.xsd",
		"/home/mmcnew/repos/forgexml-scap/schemas/oval/5.11.2/unix-system-characteristics-schema.xsd",
		"/home/mmcnew/repos/forgexml-scap/schemas/oval/5.11.2/windows-definitions-schema.xsd",
		"/home/mmcnew/repos/forgexml-scap/schemas/oval/5.11.2/windows-system-characteristics-schema.xsd",
		"/home/mmcnew/repos/forgexml-scap/schemas/oval/5.11.2/xmldsig-core-schema.xsd",
	},
	GeneratedBy: "forgexml v1.0.0",
}
