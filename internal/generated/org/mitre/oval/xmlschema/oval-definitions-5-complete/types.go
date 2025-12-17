// Package oval_definitions_5_complete - Complete variant with all substitution group members
// Generated from schema: http://oval.mitre.org/XMLSchema/oval-definitions-5
// This package imports base + all platform packages to avoid import cycles

package oval_definitions_5_complete

import (
	"encoding/xml"

	oval_definitions_5 "github.com/aequo-labs/forgexml-scap/internal/generated/org/mitre/oval/xmlschema/oval-definitions-5"
	oval_definitions_5_aix "github.com/aequo-labs/forgexml-scap/internal/generated/org/mitre/oval/xmlschema/oval-definitions-5-aix"
	oval_definitions_5_android "github.com/aequo-labs/forgexml-scap/internal/generated/org/mitre/oval/xmlschema/oval-definitions-5-android"
	oval_definitions_5_apache "github.com/aequo-labs/forgexml-scap/internal/generated/org/mitre/oval/xmlschema/oval-definitions-5-apache"
	oval_definitions_5_apple_ios "github.com/aequo-labs/forgexml-scap/internal/generated/org/mitre/oval/xmlschema/oval-definitions-5-apple_ios"
	oval_definitions_5_asa "github.com/aequo-labs/forgexml-scap/internal/generated/org/mitre/oval/xmlschema/oval-definitions-5-asa"
	oval_definitions_5_catos "github.com/aequo-labs/forgexml-scap/internal/generated/org/mitre/oval/xmlschema/oval-definitions-5-catos"
	oval_definitions_5_esx "github.com/aequo-labs/forgexml-scap/internal/generated/org/mitre/oval/xmlschema/oval-definitions-5-esx"
	oval_definitions_5_freebsd "github.com/aequo-labs/forgexml-scap/internal/generated/org/mitre/oval/xmlschema/oval-definitions-5-freebsd"
	oval_definitions_5_hpux "github.com/aequo-labs/forgexml-scap/internal/generated/org/mitre/oval/xmlschema/oval-definitions-5-hpux"
	oval_definitions_5_independent "github.com/aequo-labs/forgexml-scap/internal/generated/org/mitre/oval/xmlschema/oval-definitions-5-independent"
	oval_definitions_5_ios "github.com/aequo-labs/forgexml-scap/internal/generated/org/mitre/oval/xmlschema/oval-definitions-5-ios"
	oval_definitions_5_iosxe "github.com/aequo-labs/forgexml-scap/internal/generated/org/mitre/oval/xmlschema/oval-definitions-5-iosxe"
	oval_definitions_5_junos "github.com/aequo-labs/forgexml-scap/internal/generated/org/mitre/oval/xmlschema/oval-definitions-5-junos"
	oval_definitions_5_linux "github.com/aequo-labs/forgexml-scap/internal/generated/org/mitre/oval/xmlschema/oval-definitions-5-linux"
	oval_definitions_5_macos "github.com/aequo-labs/forgexml-scap/internal/generated/org/mitre/oval/xmlschema/oval-definitions-5-macos"
	oval_definitions_5_netconf "github.com/aequo-labs/forgexml-scap/internal/generated/org/mitre/oval/xmlschema/oval-definitions-5-netconf"
	oval_definitions_5_pixos "github.com/aequo-labs/forgexml-scap/internal/generated/org/mitre/oval/xmlschema/oval-definitions-5-pixos"
	oval_definitions_5_sharepoint "github.com/aequo-labs/forgexml-scap/internal/generated/org/mitre/oval/xmlschema/oval-definitions-5-sharepoint"
	oval_definitions_5_solaris "github.com/aequo-labs/forgexml-scap/internal/generated/org/mitre/oval/xmlschema/oval-definitions-5-solaris"
	oval_definitions_5_unix "github.com/aequo-labs/forgexml-scap/internal/generated/org/mitre/oval/xmlschema/oval-definitions-5-unix"
	oval_definitions_5_windows "github.com/aequo-labs/forgexml-scap/internal/generated/org/mitre/oval/xmlschema/oval-definitions-5-windows"
)

// Re-export base package types and functions
var Schemas = oval_definitions_5.Schemas
var ExtractElementPrefixes = oval_definitions_5.ExtractElementPrefixes
var ExtractElementsWithXmlns = oval_definitions_5.ExtractElementsWithXmlns

// Re-export base package element types
type NotesElement = oval_definitions_5.NotesElement
type ObjectElement = oval_definitions_5.ObjectElement
type StateElement = oval_definitions_5.StateElement
type Constant_variableElement = oval_definitions_5.Constant_variableElement
type External_variableElement = oval_definitions_5.External_variableElement
type Local_variableElement = oval_definitions_5.Local_variableElement
type DefinitionElement = oval_definitions_5.DefinitionElement
type TestElement = oval_definitions_5.TestElement
type SetElement = oval_definitions_5.SetElement
type FilterElement = oval_definitions_5.FilterElement
type VariableElement = oval_definitions_5.VariableElement

// Oval_definitionsElement - Complete version using complete types
// Embeds base to get all fields/methods, overrides specific fields with complete types
type Oval_definitionsElement struct {
	oval_definitions_5.Oval_definitionsElement
	// Tests overrides base field to use complete type
	Tests *TestsType `xml:"tests,omitempty"`
	// Objects overrides base field to use complete type
	Objects *ObjectsType `xml:"objects,omitempty"`
	// States overrides base field to use complete type
	States *StatesType `xml:"states,omitempty"`
}

// StatesType - Complete version with all substitution group members
// This type is standalone without embedding to avoid UnknownElements stripping xmlns
type StatesType struct {
	// Apple_iosProfile_state from namespace http://oval.mitre.org/XMLSchema/oval-definitions-5#apple_ios
	Apple_iosProfile_state []oval_definitions_5_apple_ios.Profile_stateElement `xml:"http://oval.mitre.org/XMLSchema/oval-definitions-5#apple_ios profile_state,omitempty"`
	// Apple_iosGlobalrestrictions_state from namespace http://oval.mitre.org/XMLSchema/oval-definitions-5#apple_ios
	Apple_iosGlobalrestrictions_state []oval_definitions_5_apple_ios.Globalrestrictions_stateElement `xml:"http://oval.mitre.org/XMLSchema/oval-definitions-5#apple_ios globalrestrictions_state,omitempty"`
	// Apple_iosPasscodepolicy_state from namespace http://oval.mitre.org/XMLSchema/oval-definitions-5#apple_ios
	Apple_iosPasscodepolicy_state []oval_definitions_5_apple_ios.Passcodepolicy_stateElement `xml:"http://oval.mitre.org/XMLSchema/oval-definitions-5#apple_ios passcodepolicy_state,omitempty"`
	// HpuxPatch53_state from namespace http://oval.mitre.org/XMLSchema/oval-definitions-5#hpux
	HpuxPatch53_state []oval_definitions_5_hpux.Patch53_stateElement `xml:"http://oval.mitre.org/XMLSchema/oval-definitions-5#hpux patch53_state,omitempty"`
	// HpuxSwlist_state from namespace http://oval.mitre.org/XMLSchema/oval-definitions-5#hpux
	HpuxSwlist_state []oval_definitions_5_hpux.Swlist_stateElement `xml:"http://oval.mitre.org/XMLSchema/oval-definitions-5#hpux swlist_state,omitempty"`
	// HpuxGetconf_state from namespace http://oval.mitre.org/XMLSchema/oval-definitions-5#hpux
	HpuxGetconf_state []oval_definitions_5_hpux.Getconf_stateElement `xml:"http://oval.mitre.org/XMLSchema/oval-definitions-5#hpux getconf_state,omitempty"`
	// HpuxNdd_state from namespace http://oval.mitre.org/XMLSchema/oval-definitions-5#hpux
	HpuxNdd_state []oval_definitions_5_hpux.Ndd_stateElement `xml:"http://oval.mitre.org/XMLSchema/oval-definitions-5#hpux ndd_state,omitempty"`
	// HpuxPatch_state from namespace http://oval.mitre.org/XMLSchema/oval-definitions-5#hpux
	HpuxPatch_state []oval_definitions_5_hpux.Patch_stateElement `xml:"http://oval.mitre.org/XMLSchema/oval-definitions-5#hpux patch_state,omitempty"`
	// HpuxTrusted_state from namespace http://oval.mitre.org/XMLSchema/oval-definitions-5#hpux
	HpuxTrusted_state []oval_definitions_5_hpux.Trusted_stateElement `xml:"http://oval.mitre.org/XMLSchema/oval-definitions-5#hpux trusted_state,omitempty"`
	// JunosXml_config_state from namespace http://oval.mitre.org/XMLSchema/oval-definitions-5#junos
	JunosXml_config_state []oval_definitions_5_junos.Xml_config_stateElement `xml:"http://oval.mitre.org/XMLSchema/oval-definitions-5#junos xml_config_state,omitempty"`
	// JunosVersion_state from namespace http://oval.mitre.org/XMLSchema/oval-definitions-5#junos
	JunosVersion_state []oval_definitions_5_junos.Version_stateElement `xml:"http://oval.mitre.org/XMLSchema/oval-definitions-5#junos version_state,omitempty"`
	// JunosXml_show_state from namespace http://oval.mitre.org/XMLSchema/oval-definitions-5#junos
	JunosXml_show_state []oval_definitions_5_junos.Xml_show_stateElement `xml:"http://oval.mitre.org/XMLSchema/oval-definitions-5#junos xml_show_state,omitempty"`
	// JunosShow_state from namespace http://oval.mitre.org/XMLSchema/oval-definitions-5#junos
	JunosShow_state []oval_definitions_5_junos.Show_stateElement `xml:"http://oval.mitre.org/XMLSchema/oval-definitions-5#junos show_state,omitempty"`
	// NetconfConfig_state from namespace http://oval.mitre.org/XMLSchema/oval-definitions-5#netconf
	NetconfConfig_state []oval_definitions_5_netconf.Config_stateElement `xml:"http://oval.mitre.org/XMLSchema/oval-definitions-5#netconf config_state,omitempty"`
	// PixosVersion_state from namespace http://oval.mitre.org/XMLSchema/oval-definitions-5#pixos
	PixosVersion_state []oval_definitions_5_pixos.Version_stateElement `xml:"http://oval.mitre.org/XMLSchema/oval-definitions-5#pixos version_state,omitempty"`
	// PixosLine_state from namespace http://oval.mitre.org/XMLSchema/oval-definitions-5#pixos
	PixosLine_state []oval_definitions_5_pixos.Line_stateElement `xml:"http://oval.mitre.org/XMLSchema/oval-definitions-5#pixos line_state,omitempty"`
	// WindowsUserright_state from namespace http://oval.mitre.org/XMLSchema/oval-definitions-5#windows
	WindowsUserright_state []oval_definitions_5_windows.Userright_stateElement `xml:"http://oval.mitre.org/XMLSchema/oval-definitions-5#windows userright_state,omitempty"`
	// WindowsPort_state from namespace http://oval.mitre.org/XMLSchema/oval-definitions-5#windows
	WindowsPort_state []oval_definitions_5_windows.Port_stateElement `xml:"http://oval.mitre.org/XMLSchema/oval-definitions-5#windows port_state,omitempty"`
	// WindowsAuditeventpolicy_state from namespace http://oval.mitre.org/XMLSchema/oval-definitions-5#windows
	WindowsAuditeventpolicy_state []oval_definitions_5_windows.Auditeventpolicy_stateElement `xml:"http://oval.mitre.org/XMLSchema/oval-definitions-5#windows auditeventpolicy_state,omitempty"`
	// WindowsFileauditedpermissions_state from namespace http://oval.mitre.org/XMLSchema/oval-definitions-5#windows
	WindowsFileauditedpermissions_state []oval_definitions_5_windows.Fileauditedpermissions_stateElement `xml:"http://oval.mitre.org/XMLSchema/oval-definitions-5#windows fileauditedpermissions_state,omitempty"`
	// WindowsPeheader_state from namespace http://oval.mitre.org/XMLSchema/oval-definitions-5#windows
	WindowsPeheader_state []oval_definitions_5_windows.Peheader_stateElement `xml:"http://oval.mitre.org/XMLSchema/oval-definitions-5#windows peheader_state,omitempty"`
	// WindowsJunction_state from namespace http://oval.mitre.org/XMLSchema/oval-definitions-5#windows
	WindowsJunction_state []oval_definitions_5_windows.Junction_stateElement `xml:"http://oval.mitre.org/XMLSchema/oval-definitions-5#windows junction_state,omitempty"`
	// WindowsWuaupdatesearcher_state from namespace http://oval.mitre.org/XMLSchema/oval-definitions-5#windows
	WindowsWuaupdatesearcher_state []oval_definitions_5_windows.Wuaupdatesearcher_stateElement `xml:"http://oval.mitre.org/XMLSchema/oval-definitions-5#windows wuaupdatesearcher_state,omitempty"`
	// WindowsProcess_state from namespace http://oval.mitre.org/XMLSchema/oval-definitions-5#windows
	WindowsProcess_state []oval_definitions_5_windows.Process_stateElement `xml:"http://oval.mitre.org/XMLSchema/oval-definitions-5#windows process_state,omitempty"`
	// WindowsDnscache_state from namespace http://oval.mitre.org/XMLSchema/oval-definitions-5#windows
	WindowsDnscache_state []oval_definitions_5_windows.Dnscache_stateElement `xml:"http://oval.mitre.org/XMLSchema/oval-definitions-5#windows dnscache_state,omitempty"`
	// WindowsAuditeventpolicysubcategories_state from namespace http://oval.mitre.org/XMLSchema/oval-definitions-5#windows
	WindowsAuditeventpolicysubcategories_state []oval_definitions_5_windows.Auditeventpolicysubcategories_stateElement `xml:"http://oval.mitre.org/XMLSchema/oval-definitions-5#windows auditeventpolicysubcategories_state,omitempty"`
	// WindowsUac_state from namespace http://oval.mitre.org/XMLSchema/oval-definitions-5#windows
	WindowsUac_state []oval_definitions_5_windows.Uac_stateElement `xml:"http://oval.mitre.org/XMLSchema/oval-definitions-5#windows uac_state,omitempty"`
	// WindowsSid_state from namespace http://oval.mitre.org/XMLSchema/oval-definitions-5#windows
	WindowsSid_state []oval_definitions_5_windows.Sid_stateElement `xml:"http://oval.mitre.org/XMLSchema/oval-definitions-5#windows sid_state,omitempty"`
	// WindowsVolume_state from namespace http://oval.mitre.org/XMLSchema/oval-definitions-5#windows
	WindowsVolume_state []oval_definitions_5_windows.Volume_stateElement `xml:"http://oval.mitre.org/XMLSchema/oval-definitions-5#windows volume_state,omitempty"`
	// WindowsFileeffectiverights_state from namespace http://oval.mitre.org/XMLSchema/oval-definitions-5#windows
	WindowsFileeffectiverights_state []oval_definitions_5_windows.Fileeffectiverights_stateElement `xml:"http://oval.mitre.org/XMLSchema/oval-definitions-5#windows fileeffectiverights_state,omitempty"`
	// WindowsSystemmetric_state from namespace http://oval.mitre.org/XMLSchema/oval-definitions-5#windows
	WindowsSystemmetric_state []oval_definitions_5_windows.Systemmetric_stateElement `xml:"http://oval.mitre.org/XMLSchema/oval-definitions-5#windows systemmetric_state,omitempty"`
	// WindowsPasswordpolicy_state from namespace http://oval.mitre.org/XMLSchema/oval-definitions-5#windows
	WindowsPasswordpolicy_state []oval_definitions_5_windows.Passwordpolicy_stateElement `xml:"http://oval.mitre.org/XMLSchema/oval-definitions-5#windows passwordpolicy_state,omitempty"`
	// WindowsRegkeyauditedpermissions53_state from namespace http://oval.mitre.org/XMLSchema/oval-definitions-5#windows
	WindowsRegkeyauditedpermissions53_state []oval_definitions_5_windows.Regkeyauditedpermissions53_stateElement `xml:"http://oval.mitre.org/XMLSchema/oval-definitions-5#windows regkeyauditedpermissions53_state,omitempty"`
	// WindowsFile_state from namespace http://oval.mitre.org/XMLSchema/oval-definitions-5#windows
	WindowsFile_state []oval_definitions_5_windows.File_stateElement `xml:"http://oval.mitre.org/XMLSchema/oval-definitions-5#windows file_state,omitempty"`
	// WindowsActivedirectory57_state from namespace http://oval.mitre.org/XMLSchema/oval-definitions-5#windows
	WindowsActivedirectory57_state []oval_definitions_5_windows.Activedirectory57_stateElement `xml:"http://oval.mitre.org/XMLSchema/oval-definitions-5#windows activedirectory57_state,omitempty"`
	// WindowsInterface_state from namespace http://oval.mitre.org/XMLSchema/oval-definitions-5#windows
	WindowsInterface_state []oval_definitions_5_windows.Interface_stateElement `xml:"http://oval.mitre.org/XMLSchema/oval-definitions-5#windows interface_state,omitempty"`
	// WindowsFileauditedpermissions53_state from namespace http://oval.mitre.org/XMLSchema/oval-definitions-5#windows
	WindowsFileauditedpermissions53_state []oval_definitions_5_windows.Fileauditedpermissions53_stateElement `xml:"http://oval.mitre.org/XMLSchema/oval-definitions-5#windows fileauditedpermissions53_state,omitempty"`
	// WindowsGroup_sid_state from namespace http://oval.mitre.org/XMLSchema/oval-definitions-5#windows
	WindowsGroup_sid_state []oval_definitions_5_windows.Group_sid_stateElement `xml:"http://oval.mitre.org/XMLSchema/oval-definitions-5#windows group_sid_state,omitempty"`
	// WindowsSid_sid_state from namespace http://oval.mitre.org/XMLSchema/oval-definitions-5#windows
	WindowsSid_sid_state []oval_definitions_5_windows.Sid_sid_stateElement `xml:"http://oval.mitre.org/XMLSchema/oval-definitions-5#windows sid_sid_state,omitempty"`
	// WindowsUser_sid55_state from namespace http://oval.mitre.org/XMLSchema/oval-definitions-5#windows
	WindowsUser_sid55_state []oval_definitions_5_windows.User_sid55_stateElement `xml:"http://oval.mitre.org/XMLSchema/oval-definitions-5#windows user_sid55_state,omitempty"`
	// WindowsFileeffectiverights53_state from namespace http://oval.mitre.org/XMLSchema/oval-definitions-5#windows
	WindowsFileeffectiverights53_state []oval_definitions_5_windows.Fileeffectiverights53_stateElement `xml:"http://oval.mitre.org/XMLSchema/oval-definitions-5#windows fileeffectiverights53_state,omitempty"`
	// WindowsProcess58_state from namespace http://oval.mitre.org/XMLSchema/oval-definitions-5#windows
	WindowsProcess58_state []oval_definitions_5_windows.Process58_stateElement `xml:"http://oval.mitre.org/XMLSchema/oval-definitions-5#windows process58_state,omitempty"`
	// WindowsCmdlet_state from namespace http://oval.mitre.org/XMLSchema/oval-definitions-5#windows
	WindowsCmdlet_state []oval_definitions_5_windows.Cmdlet_stateElement `xml:"http://oval.mitre.org/XMLSchema/oval-definitions-5#windows cmdlet_state,omitempty"`
	// WindowsRegkeyeffectiverights_state from namespace http://oval.mitre.org/XMLSchema/oval-definitions-5#windows
	WindowsRegkeyeffectiverights_state []oval_definitions_5_windows.Regkeyeffectiverights_stateElement `xml:"http://oval.mitre.org/XMLSchema/oval-definitions-5#windows regkeyeffectiverights_state,omitempty"`
	// WindowsRegkeyauditedpermissions_state from namespace http://oval.mitre.org/XMLSchema/oval-definitions-5#windows
	WindowsRegkeyauditedpermissions_state []oval_definitions_5_windows.Regkeyauditedpermissions_stateElement `xml:"http://oval.mitre.org/XMLSchema/oval-definitions-5#windows regkeyauditedpermissions_state,omitempty"`
	// WindowsSharedresourceeffectiverights_state from namespace http://oval.mitre.org/XMLSchema/oval-definitions-5#windows
	WindowsSharedresourceeffectiverights_state []oval_definitions_5_windows.Sharedresourceeffectiverights_stateElement `xml:"http://oval.mitre.org/XMLSchema/oval-definitions-5#windows sharedresourceeffectiverights_state,omitempty"`
	// WindowsRegkeyeffectiverights53_state from namespace http://oval.mitre.org/XMLSchema/oval-definitions-5#windows
	WindowsRegkeyeffectiverights53_state []oval_definitions_5_windows.Regkeyeffectiverights53_stateElement `xml:"http://oval.mitre.org/XMLSchema/oval-definitions-5#windows regkeyeffectiverights53_state,omitempty"`
	// WindowsServiceeffectiverights_state from namespace http://oval.mitre.org/XMLSchema/oval-definitions-5#windows
	WindowsServiceeffectiverights_state []oval_definitions_5_windows.Serviceeffectiverights_stateElement `xml:"http://oval.mitre.org/XMLSchema/oval-definitions-5#windows serviceeffectiverights_state,omitempty"`
	// WindowsPrintereffectiverights_state from namespace http://oval.mitre.org/XMLSchema/oval-definitions-5#windows
	WindowsPrintereffectiverights_state []oval_definitions_5_windows.Printereffectiverights_stateElement `xml:"http://oval.mitre.org/XMLSchema/oval-definitions-5#windows printereffectiverights_state,omitempty"`
	// WindowsLockoutpolicy_state from namespace http://oval.mitre.org/XMLSchema/oval-definitions-5#windows
	WindowsLockoutpolicy_state []oval_definitions_5_windows.Lockoutpolicy_stateElement `xml:"http://oval.mitre.org/XMLSchema/oval-definitions-5#windows lockoutpolicy_state,omitempty"`
	// WindowsService_state from namespace http://oval.mitre.org/XMLSchema/oval-definitions-5#windows
	WindowsService_state []oval_definitions_5_windows.Service_stateElement `xml:"http://oval.mitre.org/XMLSchema/oval-definitions-5#windows service_state,omitempty"`
	// WindowsNtuser_state from namespace http://oval.mitre.org/XMLSchema/oval-definitions-5#windows
	WindowsNtuser_state []oval_definitions_5_windows.Ntuser_stateElement `xml:"http://oval.mitre.org/XMLSchema/oval-definitions-5#windows ntuser_state,omitempty"`
	// WindowsSharedresource_state from namespace http://oval.mitre.org/XMLSchema/oval-definitions-5#windows
	WindowsSharedresource_state []oval_definitions_5_windows.Sharedresource_stateElement `xml:"http://oval.mitre.org/XMLSchema/oval-definitions-5#windows sharedresource_state,omitempty"`
	// WindowsUser_sid_state from namespace http://oval.mitre.org/XMLSchema/oval-definitions-5#windows
	WindowsUser_sid_state []oval_definitions_5_windows.User_sid_stateElement `xml:"http://oval.mitre.org/XMLSchema/oval-definitions-5#windows user_sid_state,omitempty"`
	// WindowsActivedirectory_state from namespace http://oval.mitre.org/XMLSchema/oval-definitions-5#windows
	WindowsActivedirectory_state []oval_definitions_5_windows.Activedirectory_stateElement `xml:"http://oval.mitre.org/XMLSchema/oval-definitions-5#windows activedirectory_state,omitempty"`
	// WindowsLicense_state from namespace http://oval.mitre.org/XMLSchema/oval-definitions-5#windows
	WindowsLicense_state []oval_definitions_5_windows.License_stateElement `xml:"http://oval.mitre.org/XMLSchema/oval-definitions-5#windows license_state,omitempty"`
	// WindowsRegistry_state from namespace http://oval.mitre.org/XMLSchema/oval-definitions-5#windows
	WindowsRegistry_state []oval_definitions_5_windows.Registry_stateElement `xml:"http://oval.mitre.org/XMLSchema/oval-definitions-5#windows registry_state,omitempty"`
	// WindowsWmi_state from namespace http://oval.mitre.org/XMLSchema/oval-definitions-5#windows
	WindowsWmi_state []oval_definitions_5_windows.Wmi_stateElement `xml:"http://oval.mitre.org/XMLSchema/oval-definitions-5#windows wmi_state,omitempty"`
	// WindowsGroup_state from namespace http://oval.mitre.org/XMLSchema/oval-definitions-5#windows
	WindowsGroup_state []oval_definitions_5_windows.Group_stateElement `xml:"http://oval.mitre.org/XMLSchema/oval-definitions-5#windows group_state,omitempty"`
	// WindowsMetabase_state from namespace http://oval.mitre.org/XMLSchema/oval-definitions-5#windows
	WindowsMetabase_state []oval_definitions_5_windows.Metabase_stateElement `xml:"http://oval.mitre.org/XMLSchema/oval-definitions-5#windows metabase_state,omitempty"`
	// WindowsSharedresourceauditedpermissions_state from namespace http://oval.mitre.org/XMLSchema/oval-definitions-5#windows
	WindowsSharedresourceauditedpermissions_state []oval_definitions_5_windows.Sharedresourceauditedpermissions_stateElement `xml:"http://oval.mitre.org/XMLSchema/oval-definitions-5#windows sharedresourceauditedpermissions_state,omitempty"`
	// WindowsUser_state from namespace http://oval.mitre.org/XMLSchema/oval-definitions-5#windows
	WindowsUser_state []oval_definitions_5_windows.User_stateElement `xml:"http://oval.mitre.org/XMLSchema/oval-definitions-5#windows user_state,omitempty"`
	// WindowsWmi57_state from namespace http://oval.mitre.org/XMLSchema/oval-definitions-5#windows
	WindowsWmi57_state []oval_definitions_5_windows.Wmi57_stateElement `xml:"http://oval.mitre.org/XMLSchema/oval-definitions-5#windows wmi57_state,omitempty"`
	// WindowsAccesstoken_state from namespace http://oval.mitre.org/XMLSchema/oval-definitions-5#windows
	WindowsAccesstoken_state []oval_definitions_5_windows.Accesstoken_stateElement `xml:"http://oval.mitre.org/XMLSchema/oval-definitions-5#windows accesstoken_state,omitempty"`
	// LinuxPartition_state from namespace http://oval.mitre.org/XMLSchema/oval-definitions-5#linux
	LinuxPartition_state []oval_definitions_5_linux.Partition_stateElement `xml:"http://oval.mitre.org/XMLSchema/oval-definitions-5#linux partition_state,omitempty"`
	// LinuxRpminfo_state from namespace http://oval.mitre.org/XMLSchema/oval-definitions-5#linux
	LinuxRpminfo_state []oval_definitions_5_linux.Rpminfo_stateElement `xml:"http://oval.mitre.org/XMLSchema/oval-definitions-5#linux rpminfo_state,omitempty"`
	// LinuxSlackwarepkginfo_state from namespace http://oval.mitre.org/XMLSchema/oval-definitions-5#linux
	LinuxSlackwarepkginfo_state []oval_definitions_5_linux.Slackwarepkginfo_stateElement `xml:"http://oval.mitre.org/XMLSchema/oval-definitions-5#linux slackwarepkginfo_state,omitempty"`
	// LinuxSystemdunitdependency_state from namespace http://oval.mitre.org/XMLSchema/oval-definitions-5#linux
	LinuxSystemdunitdependency_state []oval_definitions_5_linux.Systemdunitdependency_stateElement `xml:"http://oval.mitre.org/XMLSchema/oval-definitions-5#linux systemdunitdependency_state,omitempty"`
	// LinuxDpkginfo_state from namespace http://oval.mitre.org/XMLSchema/oval-definitions-5#linux
	LinuxDpkginfo_state []oval_definitions_5_linux.Dpkginfo_stateElement `xml:"http://oval.mitre.org/XMLSchema/oval-definitions-5#linux dpkginfo_state,omitempty"`
	// LinuxSelinuxsecuritycontext_state from namespace http://oval.mitre.org/XMLSchema/oval-definitions-5#linux
	LinuxSelinuxsecuritycontext_state []oval_definitions_5_linux.Selinuxsecuritycontext_stateElement `xml:"http://oval.mitre.org/XMLSchema/oval-definitions-5#linux selinuxsecuritycontext_state,omitempty"`
	// LinuxApparmorstatus_state from namespace http://oval.mitre.org/XMLSchema/oval-definitions-5#linux
	LinuxApparmorstatus_state []oval_definitions_5_linux.Apparmorstatus_stateElement `xml:"http://oval.mitre.org/XMLSchema/oval-definitions-5#linux apparmorstatus_state,omitempty"`
	// LinuxRpmverifyfile_state from namespace http://oval.mitre.org/XMLSchema/oval-definitions-5#linux
	LinuxRpmverifyfile_state []oval_definitions_5_linux.Rpmverifyfile_stateElement `xml:"http://oval.mitre.org/XMLSchema/oval-definitions-5#linux rpmverifyfile_state,omitempty"`
	// LinuxSystemdunitproperty_state from namespace http://oval.mitre.org/XMLSchema/oval-definitions-5#linux
	LinuxSystemdunitproperty_state []oval_definitions_5_linux.Systemdunitproperty_stateElement `xml:"http://oval.mitre.org/XMLSchema/oval-definitions-5#linux systemdunitproperty_state,omitempty"`
	// LinuxIflisteners_state from namespace http://oval.mitre.org/XMLSchema/oval-definitions-5#linux
	LinuxIflisteners_state []oval_definitions_5_linux.Iflisteners_stateElement `xml:"http://oval.mitre.org/XMLSchema/oval-definitions-5#linux iflisteners_state,omitempty"`
	// LinuxRpmverifypackage_state from namespace http://oval.mitre.org/XMLSchema/oval-definitions-5#linux
	LinuxRpmverifypackage_state []oval_definitions_5_linux.Rpmverifypackage_stateElement `xml:"http://oval.mitre.org/XMLSchema/oval-definitions-5#linux rpmverifypackage_state,omitempty"`
	// LinuxRpmverify_state from namespace http://oval.mitre.org/XMLSchema/oval-definitions-5#linux
	LinuxRpmverify_state []oval_definitions_5_linux.Rpmverify_stateElement `xml:"http://oval.mitre.org/XMLSchema/oval-definitions-5#linux rpmverify_state,omitempty"`
	// LinuxSelinuxboolean_state from namespace http://oval.mitre.org/XMLSchema/oval-definitions-5#linux
	LinuxSelinuxboolean_state []oval_definitions_5_linux.Selinuxboolean_stateElement `xml:"http://oval.mitre.org/XMLSchema/oval-definitions-5#linux selinuxboolean_state,omitempty"`
	// LinuxInetlisteningservers_state from namespace http://oval.mitre.org/XMLSchema/oval-definitions-5#linux
	LinuxInetlisteningservers_state []oval_definitions_5_linux.Inetlisteningservers_stateElement `xml:"http://oval.mitre.org/XMLSchema/oval-definitions-5#linux inetlisteningservers_state,omitempty"`
	// AsaClass_map_state from namespace http://oval.mitre.org/XMLSchema/oval-definitions-5#asa
	AsaClass_map_state []oval_definitions_5_asa.Class_map_stateElement `xml:"http://oval.mitre.org/XMLSchema/oval-definitions-5#asa class_map_state,omitempty"`
	// AsaService_policy_state from namespace http://oval.mitre.org/XMLSchema/oval-definitions-5#asa
	AsaService_policy_state []oval_definitions_5_asa.Service_policy_stateElement `xml:"http://oval.mitre.org/XMLSchema/oval-definitions-5#asa service_policy_state,omitempty"`
	// AsaSnmp_group_state from namespace http://oval.mitre.org/XMLSchema/oval-definitions-5#asa
	AsaSnmp_group_state []oval_definitions_5_asa.Snmp_group_stateElement `xml:"http://oval.mitre.org/XMLSchema/oval-definitions-5#asa snmp_group_state,omitempty"`
	// AsaVersion_state from namespace http://oval.mitre.org/XMLSchema/oval-definitions-5#asa
	AsaVersion_state []oval_definitions_5_asa.Version_stateElement `xml:"http://oval.mitre.org/XMLSchema/oval-definitions-5#asa version_state,omitempty"`
	// AsaSnmp_host_state from namespace http://oval.mitre.org/XMLSchema/oval-definitions-5#asa
	AsaSnmp_host_state []oval_definitions_5_asa.Snmp_host_stateElement `xml:"http://oval.mitre.org/XMLSchema/oval-definitions-5#asa snmp_host_state,omitempty"`
	// AsaTcp_map_state from namespace http://oval.mitre.org/XMLSchema/oval-definitions-5#asa
	AsaTcp_map_state []oval_definitions_5_asa.Tcp_map_stateElement `xml:"http://oval.mitre.org/XMLSchema/oval-definitions-5#asa tcp_map_state,omitempty"`
	// AsaAcl_state from namespace http://oval.mitre.org/XMLSchema/oval-definitions-5#asa
	AsaAcl_state []oval_definitions_5_asa.Acl_stateElement `xml:"http://oval.mitre.org/XMLSchema/oval-definitions-5#asa acl_state,omitempty"`
	// AsaPolicy_map_state from namespace http://oval.mitre.org/XMLSchema/oval-definitions-5#asa
	AsaPolicy_map_state []oval_definitions_5_asa.Policy_map_stateElement `xml:"http://oval.mitre.org/XMLSchema/oval-definitions-5#asa policy_map_state,omitempty"`
	// AsaSnmp_user_state from namespace http://oval.mitre.org/XMLSchema/oval-definitions-5#asa
	AsaSnmp_user_state []oval_definitions_5_asa.Snmp_user_stateElement `xml:"http://oval.mitre.org/XMLSchema/oval-definitions-5#asa snmp_user_state,omitempty"`
	// AsaLine_state from namespace http://oval.mitre.org/XMLSchema/oval-definitions-5#asa
	AsaLine_state []oval_definitions_5_asa.Line_stateElement `xml:"http://oval.mitre.org/XMLSchema/oval-definitions-5#asa line_state,omitempty"`
	// AsaInterface_state from namespace http://oval.mitre.org/XMLSchema/oval-definitions-5#asa
	AsaInterface_state []oval_definitions_5_asa.Interface_stateElement `xml:"http://oval.mitre.org/XMLSchema/oval-definitions-5#asa interface_state,omitempty"`
	// IosxeSnmphost_state from namespace http://oval.mitre.org/XMLSchema/oval-definitions-5#iosxe
	IosxeSnmphost_state []oval_definitions_5_iosxe.Snmphost_stateElement `xml:"http://oval.mitre.org/XMLSchema/oval-definitions-5#iosxe snmphost_state,omitempty"`
	// IosxeInterface_state from namespace http://oval.mitre.org/XMLSchema/oval-definitions-5#iosxe
	IosxeInterface_state []oval_definitions_5_iosxe.Interface_stateElement `xml:"http://oval.mitre.org/XMLSchema/oval-definitions-5#iosxe interface_state,omitempty"`
	// IosxeSection_state from namespace http://oval.mitre.org/XMLSchema/oval-definitions-5#iosxe
	IosxeSection_state []oval_definitions_5_iosxe.Section_stateElement `xml:"http://oval.mitre.org/XMLSchema/oval-definitions-5#iosxe section_state,omitempty"`
	// IosxeSnmpgroup_state from namespace http://oval.mitre.org/XMLSchema/oval-definitions-5#iosxe
	IosxeSnmpgroup_state []oval_definitions_5_iosxe.Snmpgroup_stateElement `xml:"http://oval.mitre.org/XMLSchema/oval-definitions-5#iosxe snmpgroup_state,omitempty"`
	// IosxeSnmpuser_state from namespace http://oval.mitre.org/XMLSchema/oval-definitions-5#iosxe
	IosxeSnmpuser_state []oval_definitions_5_iosxe.Snmpuser_stateElement `xml:"http://oval.mitre.org/XMLSchema/oval-definitions-5#iosxe snmpuser_state,omitempty"`
	// IosxeGlobal_state from namespace http://oval.mitre.org/XMLSchema/oval-definitions-5#iosxe
	IosxeGlobal_state []oval_definitions_5_iosxe.Global_stateElement `xml:"http://oval.mitre.org/XMLSchema/oval-definitions-5#iosxe global_state,omitempty"`
	// IosxeSnmpcommunity_state from namespace http://oval.mitre.org/XMLSchema/oval-definitions-5#iosxe
	IosxeSnmpcommunity_state []oval_definitions_5_iosxe.Snmpcommunity_stateElement `xml:"http://oval.mitre.org/XMLSchema/oval-definitions-5#iosxe snmpcommunity_state,omitempty"`
	// IosxeSnmpview_state from namespace http://oval.mitre.org/XMLSchema/oval-definitions-5#iosxe
	IosxeSnmpview_state []oval_definitions_5_iosxe.Snmpview_stateElement `xml:"http://oval.mitre.org/XMLSchema/oval-definitions-5#iosxe snmpview_state,omitempty"`
	// IosxeLine_state from namespace http://oval.mitre.org/XMLSchema/oval-definitions-5#iosxe
	IosxeLine_state []oval_definitions_5_iosxe.Line_stateElement `xml:"http://oval.mitre.org/XMLSchema/oval-definitions-5#iosxe line_state,omitempty"`
	// IosxeVersion_state from namespace http://oval.mitre.org/XMLSchema/oval-definitions-5#iosxe
	IosxeVersion_state []oval_definitions_5_iosxe.Version_stateElement `xml:"http://oval.mitre.org/XMLSchema/oval-definitions-5#iosxe version_state,omitempty"`
	// IosxeBgpneighbor_state from namespace http://oval.mitre.org/XMLSchema/oval-definitions-5#iosxe
	IosxeBgpneighbor_state []oval_definitions_5_iosxe.Bgpneighbor_stateElement `xml:"http://oval.mitre.org/XMLSchema/oval-definitions-5#iosxe bgpneighbor_state,omitempty"`
	// IosxeRoutingprotocolauthintf_state from namespace http://oval.mitre.org/XMLSchema/oval-definitions-5#iosxe
	IosxeRoutingprotocolauthintf_state []oval_definitions_5_iosxe.Routingprotocolauthintf_stateElement `xml:"http://oval.mitre.org/XMLSchema/oval-definitions-5#iosxe routingprotocolauthintf_state,omitempty"`
	// IosxeRouter_state from namespace http://oval.mitre.org/XMLSchema/oval-definitions-5#iosxe
	IosxeRouter_state []oval_definitions_5_iosxe.Router_stateElement `xml:"http://oval.mitre.org/XMLSchema/oval-definitions-5#iosxe router_state,omitempty"`
	// IosxeAcl_state from namespace http://oval.mitre.org/XMLSchema/oval-definitions-5#iosxe
	IosxeAcl_state []oval_definitions_5_iosxe.Acl_stateElement `xml:"http://oval.mitre.org/XMLSchema/oval-definitions-5#iosxe acl_state,omitempty"`
	// MacosSoftwareupdate_state from namespace http://oval.mitre.org/XMLSchema/oval-definitions-5#macos
	MacosSoftwareupdate_state []oval_definitions_5_macos.Softwareupdate_stateElement `xml:"http://oval.mitre.org/XMLSchema/oval-definitions-5#macos softwareupdate_state,omitempty"`
	// MacosDiskutil_state from namespace http://oval.mitre.org/XMLSchema/oval-definitions-5#macos
	MacosDiskutil_state []oval_definitions_5_macos.Diskutil_stateElement `xml:"http://oval.mitre.org/XMLSchema/oval-definitions-5#macos diskutil_state,omitempty"`
	// MacosPwpolicy_state from namespace http://oval.mitre.org/XMLSchema/oval-definitions-5#macos
	MacosPwpolicy_state []oval_definitions_5_macos.Pwpolicy_stateElement `xml:"http://oval.mitre.org/XMLSchema/oval-definitions-5#macos pwpolicy_state,omitempty"`
	// MacosPwpolicy59_state from namespace http://oval.mitre.org/XMLSchema/oval-definitions-5#macos
	MacosPwpolicy59_state []oval_definitions_5_macos.Pwpolicy59_stateElement `xml:"http://oval.mitre.org/XMLSchema/oval-definitions-5#macos pwpolicy59_state,omitempty"`
	// MacosRlimit_state from namespace http://oval.mitre.org/XMLSchema/oval-definitions-5#macos
	MacosRlimit_state []oval_definitions_5_macos.Rlimit_stateElement `xml:"http://oval.mitre.org/XMLSchema/oval-definitions-5#macos rlimit_state,omitempty"`
	// MacosPlist510_state from namespace http://oval.mitre.org/XMLSchema/oval-definitions-5#macos
	MacosPlist510_state []oval_definitions_5_macos.Plist510_stateElement `xml:"http://oval.mitre.org/XMLSchema/oval-definitions-5#macos plist510_state,omitempty"`
	// MacosLaunchd_state from namespace http://oval.mitre.org/XMLSchema/oval-definitions-5#macos
	MacosLaunchd_state []oval_definitions_5_macos.Launchd_stateElement `xml:"http://oval.mitre.org/XMLSchema/oval-definitions-5#macos launchd_state,omitempty"`
	// MacosAuthorizationdb_state from namespace http://oval.mitre.org/XMLSchema/oval-definitions-5#macos
	MacosAuthorizationdb_state []oval_definitions_5_macos.Authorizationdb_stateElement `xml:"http://oval.mitre.org/XMLSchema/oval-definitions-5#macos authorizationdb_state,omitempty"`
	// MacosSystemsetup_state from namespace http://oval.mitre.org/XMLSchema/oval-definitions-5#macos
	MacosSystemsetup_state []oval_definitions_5_macos.Systemsetup_stateElement `xml:"http://oval.mitre.org/XMLSchema/oval-definitions-5#macos systemsetup_state,omitempty"`
	// MacosCorestorage_state from namespace http://oval.mitre.org/XMLSchema/oval-definitions-5#macos
	MacosCorestorage_state []oval_definitions_5_macos.Corestorage_stateElement `xml:"http://oval.mitre.org/XMLSchema/oval-definitions-5#macos corestorage_state,omitempty"`
	// MacosPlist_state from namespace http://oval.mitre.org/XMLSchema/oval-definitions-5#macos
	MacosPlist_state []oval_definitions_5_macos.Plist_stateElement `xml:"http://oval.mitre.org/XMLSchema/oval-definitions-5#macos plist_state,omitempty"`
	// MacosPlist511_state from namespace http://oval.mitre.org/XMLSchema/oval-definitions-5#macos
	MacosPlist511_state []oval_definitions_5_macos.Plist511_stateElement `xml:"http://oval.mitre.org/XMLSchema/oval-definitions-5#macos plist511_state,omitempty"`
	// MacosSystemprofiler_state from namespace http://oval.mitre.org/XMLSchema/oval-definitions-5#macos
	MacosSystemprofiler_state []oval_definitions_5_macos.Systemprofiler_stateElement `xml:"http://oval.mitre.org/XMLSchema/oval-definitions-5#macos systemprofiler_state,omitempty"`
	// MacosInetlisteningserver510_state from namespace http://oval.mitre.org/XMLSchema/oval-definitions-5#macos
	MacosInetlisteningserver510_state []oval_definitions_5_macos.Inetlisteningserver510_stateElement `xml:"http://oval.mitre.org/XMLSchema/oval-definitions-5#macos inetlisteningserver510_state,omitempty"`
	// MacosNvram_state from namespace http://oval.mitre.org/XMLSchema/oval-definitions-5#macos
	MacosNvram_state []oval_definitions_5_macos.Nvram_stateElement `xml:"http://oval.mitre.org/XMLSchema/oval-definitions-5#macos nvram_state,omitempty"`
	// MacosAccountinfo_state from namespace http://oval.mitre.org/XMLSchema/oval-definitions-5#macos
	MacosAccountinfo_state []oval_definitions_5_macos.Accountinfo_stateElement `xml:"http://oval.mitre.org/XMLSchema/oval-definitions-5#macos accountinfo_state,omitempty"`
	// MacosGatekeeper_state from namespace http://oval.mitre.org/XMLSchema/oval-definitions-5#macos
	MacosGatekeeper_state []oval_definitions_5_macos.Gatekeeper_stateElement `xml:"http://oval.mitre.org/XMLSchema/oval-definitions-5#macos gatekeeper_state,omitempty"`
	// MacosInetlisteningservers_state from namespace http://oval.mitre.org/XMLSchema/oval-definitions-5#macos
	MacosInetlisteningservers_state []oval_definitions_5_macos.Inetlisteningservers_stateElement `xml:"http://oval.mitre.org/XMLSchema/oval-definitions-5#macos inetlisteningservers_state,omitempty"`
	// MacosKeychain_state from namespace http://oval.mitre.org/XMLSchema/oval-definitions-5#macos
	MacosKeychain_state []oval_definitions_5_macos.Keychain_stateElement `xml:"http://oval.mitre.org/XMLSchema/oval-definitions-5#macos keychain_state,omitempty"`
	// AixInterim_fix_state from namespace http://oval.mitre.org/XMLSchema/oval-definitions-5#aix
	AixInterim_fix_state []oval_definitions_5_aix.Interim_fix_stateElement `xml:"http://oval.mitre.org/XMLSchema/oval-definitions-5#aix interim_fix_state,omitempty"`
	// AixNo_state from namespace http://oval.mitre.org/XMLSchema/oval-definitions-5#aix
	AixNo_state []oval_definitions_5_aix.No_stateElement `xml:"http://oval.mitre.org/XMLSchema/oval-definitions-5#aix no_state,omitempty"`
	// AixOslevel_state from namespace http://oval.mitre.org/XMLSchema/oval-definitions-5#aix
	AixOslevel_state []oval_definitions_5_aix.Oslevel_stateElement `xml:"http://oval.mitre.org/XMLSchema/oval-definitions-5#aix oslevel_state,omitempty"`
	// AixFileset_state from namespace http://oval.mitre.org/XMLSchema/oval-definitions-5#aix
	AixFileset_state []oval_definitions_5_aix.Fileset_stateElement `xml:"http://oval.mitre.org/XMLSchema/oval-definitions-5#aix fileset_state,omitempty"`
	// AixFix_state from namespace http://oval.mitre.org/XMLSchema/oval-definitions-5#aix
	AixFix_state []oval_definitions_5_aix.Fix_stateElement `xml:"http://oval.mitre.org/XMLSchema/oval-definitions-5#aix fix_state,omitempty"`
	// EsxPatch56_state from namespace http://oval.mitre.org/XMLSchema/oval-definitions-5#esx
	EsxPatch56_state []oval_definitions_5_esx.Patch56_stateElement `xml:"http://oval.mitre.org/XMLSchema/oval-definitions-5#esx patch56_state,omitempty"`
	// EsxPatch_state from namespace http://oval.mitre.org/XMLSchema/oval-definitions-5#esx
	EsxPatch_state []oval_definitions_5_esx.Patch_stateElement `xml:"http://oval.mitre.org/XMLSchema/oval-definitions-5#esx patch_state,omitempty"`
	// EsxVersion_state from namespace http://oval.mitre.org/XMLSchema/oval-definitions-5#esx
	EsxVersion_state []oval_definitions_5_esx.Version_stateElement `xml:"http://oval.mitre.org/XMLSchema/oval-definitions-5#esx version_state,omitempty"`
	// EsxVisdkmanagedobject_state from namespace http://oval.mitre.org/XMLSchema/oval-definitions-5#esx
	EsxVisdkmanagedobject_state []oval_definitions_5_esx.Visdkmanagedobject_stateElement `xml:"http://oval.mitre.org/XMLSchema/oval-definitions-5#esx visdkmanagedobject_state,omitempty"`
	// UnixInterface_state from namespace http://oval.mitre.org/XMLSchema/oval-definitions-5#unix
	UnixInterface_state []oval_definitions_5_unix.Interface_stateElement `xml:"http://oval.mitre.org/XMLSchema/oval-definitions-5#unix interface_state,omitempty"`
	// UnixSymlink_state from namespace http://oval.mitre.org/XMLSchema/oval-definitions-5#unix
	UnixSymlink_state []oval_definitions_5_unix.Symlink_stateElement `xml:"http://oval.mitre.org/XMLSchema/oval-definitions-5#unix symlink_state,omitempty"`
	// UnixSysctl_state from namespace http://oval.mitre.org/XMLSchema/oval-definitions-5#unix
	UnixSysctl_state []oval_definitions_5_unix.Sysctl_stateElement `xml:"http://oval.mitre.org/XMLSchema/oval-definitions-5#unix sysctl_state,omitempty"`
	// UnixFileextendedattribute_state from namespace http://oval.mitre.org/XMLSchema/oval-definitions-5#unix
	UnixFileextendedattribute_state []oval_definitions_5_unix.Fileextendedattribute_stateElement `xml:"http://oval.mitre.org/XMLSchema/oval-definitions-5#unix fileextendedattribute_state,omitempty"`
	// UnixProcess_state from namespace http://oval.mitre.org/XMLSchema/oval-definitions-5#unix
	UnixProcess_state []oval_definitions_5_unix.Process_stateElement `xml:"http://oval.mitre.org/XMLSchema/oval-definitions-5#unix process_state,omitempty"`
	// UnixRunlevel_state from namespace http://oval.mitre.org/XMLSchema/oval-definitions-5#unix
	UnixRunlevel_state []oval_definitions_5_unix.Runlevel_stateElement `xml:"http://oval.mitre.org/XMLSchema/oval-definitions-5#unix runlevel_state,omitempty"`
	// UnixPassword_state from namespace http://oval.mitre.org/XMLSchema/oval-definitions-5#unix
	UnixPassword_state []oval_definitions_5_unix.Password_stateElement `xml:"http://oval.mitre.org/XMLSchema/oval-definitions-5#unix password_state,omitempty"`
	// UnixSccs_state from namespace http://oval.mitre.org/XMLSchema/oval-definitions-5#unix
	UnixSccs_state []oval_definitions_5_unix.Sccs_stateElement `xml:"http://oval.mitre.org/XMLSchema/oval-definitions-5#unix sccs_state,omitempty"`
	// UnixUname_state from namespace http://oval.mitre.org/XMLSchema/oval-definitions-5#unix
	UnixUname_state []oval_definitions_5_unix.Uname_stateElement `xml:"http://oval.mitre.org/XMLSchema/oval-definitions-5#unix uname_state,omitempty"`
	// UnixXinetd_state from namespace http://oval.mitre.org/XMLSchema/oval-definitions-5#unix
	UnixXinetd_state []oval_definitions_5_unix.Xinetd_stateElement `xml:"http://oval.mitre.org/XMLSchema/oval-definitions-5#unix xinetd_state,omitempty"`
	// UnixProcess58_state from namespace http://oval.mitre.org/XMLSchema/oval-definitions-5#unix
	UnixProcess58_state []oval_definitions_5_unix.Process58_stateElement `xml:"http://oval.mitre.org/XMLSchema/oval-definitions-5#unix process58_state,omitempty"`
	// UnixShadow_state from namespace http://oval.mitre.org/XMLSchema/oval-definitions-5#unix
	UnixShadow_state []oval_definitions_5_unix.Shadow_stateElement `xml:"http://oval.mitre.org/XMLSchema/oval-definitions-5#unix shadow_state,omitempty"`
	// UnixFile_state from namespace http://oval.mitre.org/XMLSchema/oval-definitions-5#unix
	UnixFile_state []oval_definitions_5_unix.File_stateElement `xml:"http://oval.mitre.org/XMLSchema/oval-definitions-5#unix file_state,omitempty"`
	// UnixGconf_state from namespace http://oval.mitre.org/XMLSchema/oval-definitions-5#unix
	UnixGconf_state []oval_definitions_5_unix.Gconf_stateElement `xml:"http://oval.mitre.org/XMLSchema/oval-definitions-5#unix gconf_state,omitempty"`
	// UnixInetd_state from namespace http://oval.mitre.org/XMLSchema/oval-definitions-5#unix
	UnixInetd_state []oval_definitions_5_unix.Inetd_stateElement `xml:"http://oval.mitre.org/XMLSchema/oval-definitions-5#unix inetd_state,omitempty"`
	// UnixRoutingtable_state from namespace http://oval.mitre.org/XMLSchema/oval-definitions-5#unix
	UnixRoutingtable_state []oval_definitions_5_unix.Routingtable_stateElement `xml:"http://oval.mitre.org/XMLSchema/oval-definitions-5#unix routingtable_state,omitempty"`
	// UnixDnscache_state from namespace http://oval.mitre.org/XMLSchema/oval-definitions-5#unix
	UnixDnscache_state []oval_definitions_5_unix.Dnscache_stateElement `xml:"http://oval.mitre.org/XMLSchema/oval-definitions-5#unix dnscache_state,omitempty"`
	// FreebsdPortinfo_state from namespace http://oval.mitre.org/XMLSchema/oval-definitions-5#freebsd
	FreebsdPortinfo_state []oval_definitions_5_freebsd.Portinfo_stateElement `xml:"http://oval.mitre.org/XMLSchema/oval-definitions-5#freebsd portinfo_state,omitempty"`
	// IosAcl_state from namespace http://oval.mitre.org/XMLSchema/oval-definitions-5#ios
	IosAcl_state []oval_definitions_5_ios.Acl_stateElement `xml:"http://oval.mitre.org/XMLSchema/oval-definitions-5#ios acl_state,omitempty"`
	// IosLine_state from namespace http://oval.mitre.org/XMLSchema/oval-definitions-5#ios
	IosLine_state []oval_definitions_5_ios.Line_stateElement `xml:"http://oval.mitre.org/XMLSchema/oval-definitions-5#ios line_state,omitempty"`
	// IosSnmp_state from namespace http://oval.mitre.org/XMLSchema/oval-definitions-5#ios
	IosSnmp_state []oval_definitions_5_ios.Snmp_stateElement `xml:"http://oval.mitre.org/XMLSchema/oval-definitions-5#ios snmp_state,omitempty"`
	// IosGlobal_state from namespace http://oval.mitre.org/XMLSchema/oval-definitions-5#ios
	IosGlobal_state []oval_definitions_5_ios.Global_stateElement `xml:"http://oval.mitre.org/XMLSchema/oval-definitions-5#ios global_state,omitempty"`
	// IosSection_state from namespace http://oval.mitre.org/XMLSchema/oval-definitions-5#ios
	IosSection_state []oval_definitions_5_ios.Section_stateElement `xml:"http://oval.mitre.org/XMLSchema/oval-definitions-5#ios section_state,omitempty"`
	// IosSnmpview_state from namespace http://oval.mitre.org/XMLSchema/oval-definitions-5#ios
	IosSnmpview_state []oval_definitions_5_ios.Snmpview_stateElement `xml:"http://oval.mitre.org/XMLSchema/oval-definitions-5#ios snmpview_state,omitempty"`
	// IosSnmpgroup_state from namespace http://oval.mitre.org/XMLSchema/oval-definitions-5#ios
	IosSnmpgroup_state []oval_definitions_5_ios.Snmpgroup_stateElement `xml:"http://oval.mitre.org/XMLSchema/oval-definitions-5#ios snmpgroup_state,omitempty"`
	// IosBgpneighbor_state from namespace http://oval.mitre.org/XMLSchema/oval-definitions-5#ios
	IosBgpneighbor_state []oval_definitions_5_ios.Bgpneighbor_stateElement `xml:"http://oval.mitre.org/XMLSchema/oval-definitions-5#ios bgpneighbor_state,omitempty"`
	// IosVersion_state from namespace http://oval.mitre.org/XMLSchema/oval-definitions-5#ios
	IosVersion_state []oval_definitions_5_ios.Version_stateElement `xml:"http://oval.mitre.org/XMLSchema/oval-definitions-5#ios version_state,omitempty"`
	// IosSnmpuser_state from namespace http://oval.mitre.org/XMLSchema/oval-definitions-5#ios
	IosSnmpuser_state []oval_definitions_5_ios.Snmpuser_stateElement `xml:"http://oval.mitre.org/XMLSchema/oval-definitions-5#ios snmpuser_state,omitempty"`
	// IosTclsh_state from namespace http://oval.mitre.org/XMLSchema/oval-definitions-5#ios
	IosTclsh_state []oval_definitions_5_ios.Tclsh_stateElement `xml:"http://oval.mitre.org/XMLSchema/oval-definitions-5#ios tclsh_state,omitempty"`
	// IosVersion55_state from namespace http://oval.mitre.org/XMLSchema/oval-definitions-5#ios
	IosVersion55_state []oval_definitions_5_ios.Version55_stateElement `xml:"http://oval.mitre.org/XMLSchema/oval-definitions-5#ios version55_state,omitempty"`
	// IosInterface_state from namespace http://oval.mitre.org/XMLSchema/oval-definitions-5#ios
	IosInterface_state []oval_definitions_5_ios.Interface_stateElement `xml:"http://oval.mitre.org/XMLSchema/oval-definitions-5#ios interface_state,omitempty"`
	// IosRouter_state from namespace http://oval.mitre.org/XMLSchema/oval-definitions-5#ios
	IosRouter_state []oval_definitions_5_ios.Router_stateElement `xml:"http://oval.mitre.org/XMLSchema/oval-definitions-5#ios router_state,omitempty"`
	// IosSnmpcommunity_state from namespace http://oval.mitre.org/XMLSchema/oval-definitions-5#ios
	IosSnmpcommunity_state []oval_definitions_5_ios.Snmpcommunity_stateElement `xml:"http://oval.mitre.org/XMLSchema/oval-definitions-5#ios snmpcommunity_state,omitempty"`
	// IosSnmphost_state from namespace http://oval.mitre.org/XMLSchema/oval-definitions-5#ios
	IosSnmphost_state []oval_definitions_5_ios.Snmphost_stateElement `xml:"http://oval.mitre.org/XMLSchema/oval-definitions-5#ios snmphost_state,omitempty"`
	// IosRoutingprotocolauthintf_state from namespace http://oval.mitre.org/XMLSchema/oval-definitions-5#ios
	IosRoutingprotocolauthintf_state []oval_definitions_5_ios.Routingprotocolauthintf_stateElement `xml:"http://oval.mitre.org/XMLSchema/oval-definitions-5#ios routingprotocolauthintf_state,omitempty"`
	// ApacheHttpd_state from namespace http://oval.mitre.org/XMLSchema/oval-definitions-5#apache
	ApacheHttpd_state []oval_definitions_5_apache.Httpd_stateElement `xml:"http://oval.mitre.org/XMLSchema/oval-definitions-5#apache httpd_state,omitempty"`
	// SharepointSpjobdefinition510_state from namespace http://oval.mitre.org/XMLSchema/oval-definitions-5#sharepoint
	SharepointSpjobdefinition510_state []oval_definitions_5_sharepoint.Spjobdefinition510_stateElement `xml:"http://oval.mitre.org/XMLSchema/oval-definitions-5#sharepoint spjobdefinition510_state,omitempty"`
	// SharepointSpwebapplication_state from namespace http://oval.mitre.org/XMLSchema/oval-definitions-5#sharepoint
	SharepointSpwebapplication_state []oval_definitions_5_sharepoint.Spwebapplication_stateElement `xml:"http://oval.mitre.org/XMLSchema/oval-definitions-5#sharepoint spwebapplication_state,omitempty"`
	// SharepointSplist_state from namespace http://oval.mitre.org/XMLSchema/oval-definitions-5#sharepoint
	SharepointSplist_state []oval_definitions_5_sharepoint.Splist_stateElement `xml:"http://oval.mitre.org/XMLSchema/oval-definitions-5#sharepoint splist_state,omitempty"`
	// SharepointSpantivirussettings_state from namespace http://oval.mitre.org/XMLSchema/oval-definitions-5#sharepoint
	SharepointSpantivirussettings_state []oval_definitions_5_sharepoint.Spantivirussettings_stateElement `xml:"http://oval.mitre.org/XMLSchema/oval-definitions-5#sharepoint spantivirussettings_state,omitempty"`
	// SharepointSpsiteadministration_state from namespace http://oval.mitre.org/XMLSchema/oval-definitions-5#sharepoint
	SharepointSpsiteadministration_state []oval_definitions_5_sharepoint.Spsiteadministration_stateElement `xml:"http://oval.mitre.org/XMLSchema/oval-definitions-5#sharepoint spsiteadministration_state,omitempty"`
	// SharepointSpsite_state from namespace http://oval.mitre.org/XMLSchema/oval-definitions-5#sharepoint
	SharepointSpsite_state []oval_definitions_5_sharepoint.Spsite_stateElement `xml:"http://oval.mitre.org/XMLSchema/oval-definitions-5#sharepoint spsite_state,omitempty"`
	// SharepointSpjobdefinition_state from namespace http://oval.mitre.org/XMLSchema/oval-definitions-5#sharepoint
	SharepointSpjobdefinition_state []oval_definitions_5_sharepoint.Spjobdefinition_stateElement `xml:"http://oval.mitre.org/XMLSchema/oval-definitions-5#sharepoint spjobdefinition_state,omitempty"`
	// SharepointBestbet_state from namespace http://oval.mitre.org/XMLSchema/oval-definitions-5#sharepoint
	SharepointBestbet_state []oval_definitions_5_sharepoint.Bestbet_stateElement `xml:"http://oval.mitre.org/XMLSchema/oval-definitions-5#sharepoint bestbet_state,omitempty"`
	// SharepointSppolicy_state from namespace http://oval.mitre.org/XMLSchema/oval-definitions-5#sharepoint
	SharepointSppolicy_state []oval_definitions_5_sharepoint.Sppolicy_stateElement `xml:"http://oval.mitre.org/XMLSchema/oval-definitions-5#sharepoint sppolicy_state,omitempty"`
	// SharepointSpcrawlrule_state from namespace http://oval.mitre.org/XMLSchema/oval-definitions-5#sharepoint
	SharepointSpcrawlrule_state []oval_definitions_5_sharepoint.Spcrawlrule_stateElement `xml:"http://oval.mitre.org/XMLSchema/oval-definitions-5#sharepoint spcrawlrule_state,omitempty"`
	// SharepointInfopolicycoll_state from namespace http://oval.mitre.org/XMLSchema/oval-definitions-5#sharepoint
	SharepointInfopolicycoll_state []oval_definitions_5_sharepoint.Infopolicycoll_stateElement `xml:"http://oval.mitre.org/XMLSchema/oval-definitions-5#sharepoint infopolicycoll_state,omitempty"`
	// SharepointSpdiagnosticsservice_state from namespace http://oval.mitre.org/XMLSchema/oval-definitions-5#sharepoint
	SharepointSpdiagnosticsservice_state []oval_definitions_5_sharepoint.Spdiagnosticsservice_stateElement `xml:"http://oval.mitre.org/XMLSchema/oval-definitions-5#sharepoint spdiagnosticsservice_state,omitempty"`
	// SharepointSpdiagnosticslevel_state from namespace http://oval.mitre.org/XMLSchema/oval-definitions-5#sharepoint
	SharepointSpdiagnosticslevel_state []oval_definitions_5_sharepoint.Spdiagnosticslevel_stateElement `xml:"http://oval.mitre.org/XMLSchema/oval-definitions-5#sharepoint spdiagnosticslevel_state,omitempty"`
	// SharepointSpweb_state from namespace http://oval.mitre.org/XMLSchema/oval-definitions-5#sharepoint
	SharepointSpweb_state []oval_definitions_5_sharepoint.Spweb_stateElement `xml:"http://oval.mitre.org/XMLSchema/oval-definitions-5#sharepoint spweb_state,omitempty"`
	// SharepointSppolicyfeature_state from namespace http://oval.mitre.org/XMLSchema/oval-definitions-5#sharepoint
	SharepointSppolicyfeature_state []oval_definitions_5_sharepoint.Sppolicyfeature_stateElement `xml:"http://oval.mitre.org/XMLSchema/oval-definitions-5#sharepoint sppolicyfeature_state,omitempty"`
	// SharepointSpgroup_state from namespace http://oval.mitre.org/XMLSchema/oval-definitions-5#sharepoint
	SharepointSpgroup_state []oval_definitions_5_sharepoint.Spgroup_stateElement `xml:"http://oval.mitre.org/XMLSchema/oval-definitions-5#sharepoint spgroup_state,omitempty"`
	// CatosVersion_state from namespace http://oval.mitre.org/XMLSchema/oval-definitions-5#catos
	CatosVersion_state []oval_definitions_5_catos.Version_stateElement `xml:"http://oval.mitre.org/XMLSchema/oval-definitions-5#catos version_state,omitempty"`
	// CatosLine_state from namespace http://oval.mitre.org/XMLSchema/oval-definitions-5#catos
	CatosLine_state []oval_definitions_5_catos.Line_stateElement `xml:"http://oval.mitre.org/XMLSchema/oval-definitions-5#catos line_state,omitempty"`
	// CatosModule_state from namespace http://oval.mitre.org/XMLSchema/oval-definitions-5#catos
	CatosModule_state []oval_definitions_5_catos.Module_stateElement `xml:"http://oval.mitre.org/XMLSchema/oval-definitions-5#catos module_state,omitempty"`
	// CatosVersion55_state from namespace http://oval.mitre.org/XMLSchema/oval-definitions-5#catos
	CatosVersion55_state []oval_definitions_5_catos.Version55_stateElement `xml:"http://oval.mitre.org/XMLSchema/oval-definitions-5#catos version55_state,omitempty"`
	// AndroidBluetooth_state from namespace http://oval.mitre.org/XMLSchema/oval-definitions-5#android
	AndroidBluetooth_state []oval_definitions_5_android.Bluetooth_stateElement `xml:"http://oval.mitre.org/XMLSchema/oval-definitions-5#android bluetooth_state,omitempty"`
	// AndroidNetwork_state from namespace http://oval.mitre.org/XMLSchema/oval-definitions-5#android
	AndroidNetwork_state []oval_definitions_5_android.Network_stateElement `xml:"http://oval.mitre.org/XMLSchema/oval-definitions-5#android network_state,omitempty"`
	// AndroidTelephony_state from namespace http://oval.mitre.org/XMLSchema/oval-definitions-5#android
	AndroidTelephony_state []oval_definitions_5_android.Telephony_stateElement `xml:"http://oval.mitre.org/XMLSchema/oval-definitions-5#android telephony_state,omitempty"`
	// AndroidCertificate_state from namespace http://oval.mitre.org/XMLSchema/oval-definitions-5#android
	AndroidCertificate_state []oval_definitions_5_android.Certificate_stateElement `xml:"http://oval.mitre.org/XMLSchema/oval-definitions-5#android certificate_state,omitempty"`
	// AndroidCamera_state from namespace http://oval.mitre.org/XMLSchema/oval-definitions-5#android
	AndroidCamera_state []oval_definitions_5_android.Camera_stateElement `xml:"http://oval.mitre.org/XMLSchema/oval-definitions-5#android camera_state,omitempty"`
	// AndroidEncryption_state from namespace http://oval.mitre.org/XMLSchema/oval-definitions-5#android
	AndroidEncryption_state []oval_definitions_5_android.Encryption_stateElement `xml:"http://oval.mitre.org/XMLSchema/oval-definitions-5#android encryption_state,omitempty"`
	// AndroidPassword_state from namespace http://oval.mitre.org/XMLSchema/oval-definitions-5#android
	AndroidPassword_state []oval_definitions_5_android.Password_stateElement `xml:"http://oval.mitre.org/XMLSchema/oval-definitions-5#android password_state,omitempty"`
	// AndroidWifinetwork_state from namespace http://oval.mitre.org/XMLSchema/oval-definitions-5#android
	AndroidWifinetwork_state []oval_definitions_5_android.Wifinetwork_stateElement `xml:"http://oval.mitre.org/XMLSchema/oval-definitions-5#android wifinetwork_state,omitempty"`
	// AndroidLocationservice_state from namespace http://oval.mitre.org/XMLSchema/oval-definitions-5#android
	AndroidLocationservice_state []oval_definitions_5_android.Locationservice_stateElement `xml:"http://oval.mitre.org/XMLSchema/oval-definitions-5#android locationservice_state,omitempty"`
	// AndroidWifi_state from namespace http://oval.mitre.org/XMLSchema/oval-definitions-5#android
	AndroidWifi_state []oval_definitions_5_android.Wifi_stateElement `xml:"http://oval.mitre.org/XMLSchema/oval-definitions-5#android wifi_state,omitempty"`
	// AndroidAppmanager_state from namespace http://oval.mitre.org/XMLSchema/oval-definitions-5#android
	AndroidAppmanager_state []oval_definitions_5_android.Appmanager_stateElement `xml:"http://oval.mitre.org/XMLSchema/oval-definitions-5#android appmanager_state,omitempty"`
	// AndroidDevicesettings_state from namespace http://oval.mitre.org/XMLSchema/oval-definitions-5#android
	AndroidDevicesettings_state []oval_definitions_5_android.Devicesettings_stateElement `xml:"http://oval.mitre.org/XMLSchema/oval-definitions-5#android devicesettings_state,omitempty"`
	// AndroidSystemdetails_state from namespace http://oval.mitre.org/XMLSchema/oval-definitions-5#android
	AndroidSystemdetails_state []oval_definitions_5_android.Systemdetails_stateElement `xml:"http://oval.mitre.org/XMLSchema/oval-definitions-5#android systemdetails_state,omitempty"`
	// IndependentFilehash58_state from namespace http://oval.mitre.org/XMLSchema/oval-definitions-5#independent
	IndependentFilehash58_state []oval_definitions_5_independent.Filehash58_stateElement `xml:"http://oval.mitre.org/XMLSchema/oval-definitions-5#independent filehash58_state,omitempty"`
	// IndependentEnvironmentvariable_state from namespace http://oval.mitre.org/XMLSchema/oval-definitions-5#independent
	IndependentEnvironmentvariable_state []oval_definitions_5_independent.Environmentvariable_stateElement `xml:"http://oval.mitre.org/XMLSchema/oval-definitions-5#independent environmentvariable_state,omitempty"`
	// IndependentFilehash_state from namespace http://oval.mitre.org/XMLSchema/oval-definitions-5#independent
	IndependentFilehash_state []oval_definitions_5_independent.Filehash_stateElement `xml:"http://oval.mitre.org/XMLSchema/oval-definitions-5#independent filehash_state,omitempty"`
	// IndependentTextfilecontent54_state from namespace http://oval.mitre.org/XMLSchema/oval-definitions-5#independent
	IndependentTextfilecontent54_state []oval_definitions_5_independent.Textfilecontent54_stateElement `xml:"http://oval.mitre.org/XMLSchema/oval-definitions-5#independent textfilecontent54_state,omitempty"`
	// IndependentTextfilecontent_state from namespace http://oval.mitre.org/XMLSchema/oval-definitions-5#independent
	IndependentTextfilecontent_state []oval_definitions_5_independent.Textfilecontent_stateElement `xml:"http://oval.mitre.org/XMLSchema/oval-definitions-5#independent textfilecontent_state,omitempty"`
	// IndependentXmlfilecontent_state from namespace http://oval.mitre.org/XMLSchema/oval-definitions-5#independent
	IndependentXmlfilecontent_state []oval_definitions_5_independent.Xmlfilecontent_stateElement `xml:"http://oval.mitre.org/XMLSchema/oval-definitions-5#independent xmlfilecontent_state,omitempty"`
	// IndependentEnvironmentvariable58_state from namespace http://oval.mitre.org/XMLSchema/oval-definitions-5#independent
	IndependentEnvironmentvariable58_state []oval_definitions_5_independent.Environmentvariable58_stateElement `xml:"http://oval.mitre.org/XMLSchema/oval-definitions-5#independent environmentvariable58_state,omitempty"`
	// IndependentFamily_state from namespace http://oval.mitre.org/XMLSchema/oval-definitions-5#independent
	IndependentFamily_state []oval_definitions_5_independent.Family_stateElement `xml:"http://oval.mitre.org/XMLSchema/oval-definitions-5#independent family_state,omitempty"`
	// IndependentVariable_state from namespace http://oval.mitre.org/XMLSchema/oval-definitions-5#independent
	IndependentVariable_state []oval_definitions_5_independent.Variable_stateElement `xml:"http://oval.mitre.org/XMLSchema/oval-definitions-5#independent variable_state,omitempty"`
	// IndependentLdap57_state from namespace http://oval.mitre.org/XMLSchema/oval-definitions-5#independent
	IndependentLdap57_state []oval_definitions_5_independent.Ldap57_stateElement `xml:"http://oval.mitre.org/XMLSchema/oval-definitions-5#independent ldap57_state,omitempty"`
	// IndependentSql57_state from namespace http://oval.mitre.org/XMLSchema/oval-definitions-5#independent
	IndependentSql57_state []oval_definitions_5_independent.Sql57_stateElement `xml:"http://oval.mitre.org/XMLSchema/oval-definitions-5#independent sql57_state,omitempty"`
	// IndependentLdap_state from namespace http://oval.mitre.org/XMLSchema/oval-definitions-5#independent
	IndependentLdap_state []oval_definitions_5_independent.Ldap_stateElement `xml:"http://oval.mitre.org/XMLSchema/oval-definitions-5#independent ldap_state,omitempty"`
	// IndependentSql_state from namespace http://oval.mitre.org/XMLSchema/oval-definitions-5#independent
	IndependentSql_state []oval_definitions_5_independent.Sql_stateElement `xml:"http://oval.mitre.org/XMLSchema/oval-definitions-5#independent sql_state,omitempty"`
	// SolarisPackagecheck_state from namespace http://oval.mitre.org/XMLSchema/oval-definitions-5#solaris
	SolarisPackagecheck_state []oval_definitions_5_solaris.Packagecheck_stateElement `xml:"http://oval.mitre.org/XMLSchema/oval-definitions-5#solaris packagecheck_state,omitempty"`
	// SolarisSmf_state from namespace http://oval.mitre.org/XMLSchema/oval-definitions-5#solaris
	SolarisSmf_state []oval_definitions_5_solaris.Smf_stateElement `xml:"http://oval.mitre.org/XMLSchema/oval-definitions-5#solaris smf_state,omitempty"`
	// SolarisFacet_state from namespace http://oval.mitre.org/XMLSchema/oval-definitions-5#solaris
	SolarisFacet_state []oval_definitions_5_solaris.Facet_stateElement `xml:"http://oval.mitre.org/XMLSchema/oval-definitions-5#solaris facet_state,omitempty"`
	// SolarisImage_state from namespace http://oval.mitre.org/XMLSchema/oval-definitions-5#solaris
	SolarisImage_state []oval_definitions_5_solaris.Image_stateElement `xml:"http://oval.mitre.org/XMLSchema/oval-definitions-5#solaris image_state,omitempty"`
	// SolarisNdd_state from namespace http://oval.mitre.org/XMLSchema/oval-definitions-5#solaris
	SolarisNdd_state []oval_definitions_5_solaris.Ndd_stateElement `xml:"http://oval.mitre.org/XMLSchema/oval-definitions-5#solaris ndd_state,omitempty"`
	// SolarisPatch_state from namespace http://oval.mitre.org/XMLSchema/oval-definitions-5#solaris
	SolarisPatch_state []oval_definitions_5_solaris.Patch_stateElement `xml:"http://oval.mitre.org/XMLSchema/oval-definitions-5#solaris patch_state,omitempty"`
	// SolarisSmfproperty_state from namespace http://oval.mitre.org/XMLSchema/oval-definitions-5#solaris
	SolarisSmfproperty_state []oval_definitions_5_solaris.Smfproperty_stateElement `xml:"http://oval.mitre.org/XMLSchema/oval-definitions-5#solaris smfproperty_state,omitempty"`
	// SolarisVirtualizationinfo_state from namespace http://oval.mitre.org/XMLSchema/oval-definitions-5#solaris
	SolarisVirtualizationinfo_state []oval_definitions_5_solaris.Virtualizationinfo_stateElement `xml:"http://oval.mitre.org/XMLSchema/oval-definitions-5#solaris virtualizationinfo_state,omitempty"`
	// SolarisPackagepublisher_state from namespace http://oval.mitre.org/XMLSchema/oval-definitions-5#solaris
	SolarisPackagepublisher_state []oval_definitions_5_solaris.Packagepublisher_stateElement `xml:"http://oval.mitre.org/XMLSchema/oval-definitions-5#solaris packagepublisher_state,omitempty"`
	// SolarisVariant_state from namespace http://oval.mitre.org/XMLSchema/oval-definitions-5#solaris
	SolarisVariant_state []oval_definitions_5_solaris.Variant_stateElement `xml:"http://oval.mitre.org/XMLSchema/oval-definitions-5#solaris variant_state,omitempty"`
	// SolarisIsainfo_state from namespace http://oval.mitre.org/XMLSchema/oval-definitions-5#solaris
	SolarisIsainfo_state []oval_definitions_5_solaris.Isainfo_stateElement `xml:"http://oval.mitre.org/XMLSchema/oval-definitions-5#solaris isainfo_state,omitempty"`
	// SolarisPackageavoidlist_state from namespace http://oval.mitre.org/XMLSchema/oval-definitions-5#solaris
	SolarisPackageavoidlist_state []oval_definitions_5_solaris.Packageavoidlist_stateElement `xml:"http://oval.mitre.org/XMLSchema/oval-definitions-5#solaris packageavoidlist_state,omitempty"`
	// SolarisPackagefreezelist_state from namespace http://oval.mitre.org/XMLSchema/oval-definitions-5#solaris
	SolarisPackagefreezelist_state []oval_definitions_5_solaris.Packagefreezelist_stateElement `xml:"http://oval.mitre.org/XMLSchema/oval-definitions-5#solaris packagefreezelist_state,omitempty"`
	// SolarisPackage_state from namespace http://oval.mitre.org/XMLSchema/oval-definitions-5#solaris
	SolarisPackage_state []oval_definitions_5_solaris.Package_stateElement `xml:"http://oval.mitre.org/XMLSchema/oval-definitions-5#solaris package_state,omitempty"`
	// SolarisPackage511_state from namespace http://oval.mitre.org/XMLSchema/oval-definitions-5#solaris
	SolarisPackage511_state []oval_definitions_5_solaris.Package511_stateElement `xml:"http://oval.mitre.org/XMLSchema/oval-definitions-5#solaris package511_state,omitempty"`
	// UnknownAttrs captures any attributes not defined in XSD
	UnknownAttrs []xml.Attr `xml:",any,attr,omitempty"`
}

// TestsType - Complete version with all substitution group members
// This type is standalone without embedding to avoid UnknownElements stripping xmlns
type TestsType struct {
	// IndependentFamily_test from namespace http://oval.mitre.org/XMLSchema/oval-definitions-5#independent
	IndependentFamily_test []oval_definitions_5_independent.Family_testElement `xml:"http://oval.mitre.org/XMLSchema/oval-definitions-5#independent family_test,omitempty"`
	// IndependentEnvironmentvariable58_test from namespace http://oval.mitre.org/XMLSchema/oval-definitions-5#independent
	IndependentEnvironmentvariable58_test []oval_definitions_5_independent.Environmentvariable58_testElement `xml:"http://oval.mitre.org/XMLSchema/oval-definitions-5#independent environmentvariable58_test,omitempty"`
	// IndependentFilehash58_test from namespace http://oval.mitre.org/XMLSchema/oval-definitions-5#independent
	IndependentFilehash58_test []oval_definitions_5_independent.Filehash58_testElement `xml:"http://oval.mitre.org/XMLSchema/oval-definitions-5#independent filehash58_test,omitempty"`
	// IndependentUnknown_test from namespace http://oval.mitre.org/XMLSchema/oval-definitions-5#independent
	IndependentUnknown_test []oval_definitions_5_independent.Unknown_testElement `xml:"http://oval.mitre.org/XMLSchema/oval-definitions-5#independent unknown_test,omitempty"`
	// IndependentTextfilecontent_test from namespace http://oval.mitre.org/XMLSchema/oval-definitions-5#independent
	IndependentTextfilecontent_test []oval_definitions_5_independent.Textfilecontent_testElement `xml:"http://oval.mitre.org/XMLSchema/oval-definitions-5#independent textfilecontent_test,omitempty"`
	// IndependentVariable_test from namespace http://oval.mitre.org/XMLSchema/oval-definitions-5#independent
	IndependentVariable_test []oval_definitions_5_independent.Variable_testElement `xml:"http://oval.mitre.org/XMLSchema/oval-definitions-5#independent variable_test,omitempty"`
	// IndependentTextfilecontent54_test from namespace http://oval.mitre.org/XMLSchema/oval-definitions-5#independent
	IndependentTextfilecontent54_test []oval_definitions_5_independent.Textfilecontent54_testElement `xml:"http://oval.mitre.org/XMLSchema/oval-definitions-5#independent textfilecontent54_test,omitempty"`
	// IndependentEnvironmentvariable_test from namespace http://oval.mitre.org/XMLSchema/oval-definitions-5#independent
	IndependentEnvironmentvariable_test []oval_definitions_5_independent.Environmentvariable_testElement `xml:"http://oval.mitre.org/XMLSchema/oval-definitions-5#independent environmentvariable_test,omitempty"`
	// IndependentLdap_test from namespace http://oval.mitre.org/XMLSchema/oval-definitions-5#independent
	IndependentLdap_test []oval_definitions_5_independent.Ldap_testElement `xml:"http://oval.mitre.org/XMLSchema/oval-definitions-5#independent ldap_test,omitempty"`
	// IndependentSql_test from namespace http://oval.mitre.org/XMLSchema/oval-definitions-5#independent
	IndependentSql_test []oval_definitions_5_independent.Sql_testElement `xml:"http://oval.mitre.org/XMLSchema/oval-definitions-5#independent sql_test,omitempty"`
	// IndependentSql57_test from namespace http://oval.mitre.org/XMLSchema/oval-definitions-5#independent
	IndependentSql57_test []oval_definitions_5_independent.Sql57_testElement `xml:"http://oval.mitre.org/XMLSchema/oval-definitions-5#independent sql57_test,omitempty"`
	// IndependentXmlfilecontent_test from namespace http://oval.mitre.org/XMLSchema/oval-definitions-5#independent
	IndependentXmlfilecontent_test []oval_definitions_5_independent.Xmlfilecontent_testElement `xml:"http://oval.mitre.org/XMLSchema/oval-definitions-5#independent xmlfilecontent_test,omitempty"`
	// IndependentFilehash_test from namespace http://oval.mitre.org/XMLSchema/oval-definitions-5#independent
	IndependentFilehash_test []oval_definitions_5_independent.Filehash_testElement `xml:"http://oval.mitre.org/XMLSchema/oval-definitions-5#independent filehash_test,omitempty"`
	// IndependentLdap57_test from namespace http://oval.mitre.org/XMLSchema/oval-definitions-5#independent
	IndependentLdap57_test []oval_definitions_5_independent.Ldap57_testElement `xml:"http://oval.mitre.org/XMLSchema/oval-definitions-5#independent ldap57_test,omitempty"`
	// SolarisSmfproperty_test from namespace http://oval.mitre.org/XMLSchema/oval-definitions-5#solaris
	SolarisSmfproperty_test []oval_definitions_5_solaris.Smfproperty_testElement `xml:"http://oval.mitre.org/XMLSchema/oval-definitions-5#solaris smfproperty_test,omitempty"`
	// SolarisPackagefreezelist_test from namespace http://oval.mitre.org/XMLSchema/oval-definitions-5#solaris
	SolarisPackagefreezelist_test []oval_definitions_5_solaris.Packagefreezelist_testElement `xml:"http://oval.mitre.org/XMLSchema/oval-definitions-5#solaris packagefreezelist_test,omitempty"`
	// SolarisPatch_test from namespace http://oval.mitre.org/XMLSchema/oval-definitions-5#solaris
	SolarisPatch_test []oval_definitions_5_solaris.Patch_testElement `xml:"http://oval.mitre.org/XMLSchema/oval-definitions-5#solaris patch_test,omitempty"`
	// SolarisSmf_test from namespace http://oval.mitre.org/XMLSchema/oval-definitions-5#solaris
	SolarisSmf_test []oval_definitions_5_solaris.Smf_testElement `xml:"http://oval.mitre.org/XMLSchema/oval-definitions-5#solaris smf_test,omitempty"`
	// SolarisPackagecheck_test from namespace http://oval.mitre.org/XMLSchema/oval-definitions-5#solaris
	SolarisPackagecheck_test []oval_definitions_5_solaris.Packagecheck_testElement `xml:"http://oval.mitre.org/XMLSchema/oval-definitions-5#solaris packagecheck_test,omitempty"`
	// SolarisPatch54_test from namespace http://oval.mitre.org/XMLSchema/oval-definitions-5#solaris
	SolarisPatch54_test []oval_definitions_5_solaris.Patch54_testElement `xml:"http://oval.mitre.org/XMLSchema/oval-definitions-5#solaris patch54_test,omitempty"`
	// SolarisImage_test from namespace http://oval.mitre.org/XMLSchema/oval-definitions-5#solaris
	SolarisImage_test []oval_definitions_5_solaris.Image_testElement `xml:"http://oval.mitre.org/XMLSchema/oval-definitions-5#solaris image_test,omitempty"`
	// SolarisPackage511_test from namespace http://oval.mitre.org/XMLSchema/oval-definitions-5#solaris
	SolarisPackage511_test []oval_definitions_5_solaris.Package511_testElement `xml:"http://oval.mitre.org/XMLSchema/oval-definitions-5#solaris package511_test,omitempty"`
	// SolarisPackageavoidlist_test from namespace http://oval.mitre.org/XMLSchema/oval-definitions-5#solaris
	SolarisPackageavoidlist_test []oval_definitions_5_solaris.Packageavoidlist_testElement `xml:"http://oval.mitre.org/XMLSchema/oval-definitions-5#solaris packageavoidlist_test,omitempty"`
	// SolarisPackagepublisher_test from namespace http://oval.mitre.org/XMLSchema/oval-definitions-5#solaris
	SolarisPackagepublisher_test []oval_definitions_5_solaris.Packagepublisher_testElement `xml:"http://oval.mitre.org/XMLSchema/oval-definitions-5#solaris packagepublisher_test,omitempty"`
	// SolarisVariant_test from namespace http://oval.mitre.org/XMLSchema/oval-definitions-5#solaris
	SolarisVariant_test []oval_definitions_5_solaris.Variant_testElement `xml:"http://oval.mitre.org/XMLSchema/oval-definitions-5#solaris variant_test,omitempty"`
	// SolarisVirtualizationinfo_test from namespace http://oval.mitre.org/XMLSchema/oval-definitions-5#solaris
	SolarisVirtualizationinfo_test []oval_definitions_5_solaris.Virtualizationinfo_testElement `xml:"http://oval.mitre.org/XMLSchema/oval-definitions-5#solaris virtualizationinfo_test,omitempty"`
	// SolarisFacet_test from namespace http://oval.mitre.org/XMLSchema/oval-definitions-5#solaris
	SolarisFacet_test []oval_definitions_5_solaris.Facet_testElement `xml:"http://oval.mitre.org/XMLSchema/oval-definitions-5#solaris facet_test,omitempty"`
	// SolarisIsainfo_test from namespace http://oval.mitre.org/XMLSchema/oval-definitions-5#solaris
	SolarisIsainfo_test []oval_definitions_5_solaris.Isainfo_testElement `xml:"http://oval.mitre.org/XMLSchema/oval-definitions-5#solaris isainfo_test,omitempty"`
	// SolarisNdd_test from namespace http://oval.mitre.org/XMLSchema/oval-definitions-5#solaris
	SolarisNdd_test []oval_definitions_5_solaris.Ndd_testElement `xml:"http://oval.mitre.org/XMLSchema/oval-definitions-5#solaris ndd_test,omitempty"`
	// SolarisPackage_test from namespace http://oval.mitre.org/XMLSchema/oval-definitions-5#solaris
	SolarisPackage_test []oval_definitions_5_solaris.Package_testElement `xml:"http://oval.mitre.org/XMLSchema/oval-definitions-5#solaris package_test,omitempty"`
	// Apple_iosProfile_test from namespace http://oval.mitre.org/XMLSchema/oval-definitions-5#apple_ios
	Apple_iosProfile_test []oval_definitions_5_apple_ios.Profile_testElement `xml:"http://oval.mitre.org/XMLSchema/oval-definitions-5#apple_ios profile_test,omitempty"`
	// Apple_iosGlobalrestrictions_test from namespace http://oval.mitre.org/XMLSchema/oval-definitions-5#apple_ios
	Apple_iosGlobalrestrictions_test []oval_definitions_5_apple_ios.Globalrestrictions_testElement `xml:"http://oval.mitre.org/XMLSchema/oval-definitions-5#apple_ios globalrestrictions_test,omitempty"`
	// Apple_iosPasscodepolicy_test from namespace http://oval.mitre.org/XMLSchema/oval-definitions-5#apple_ios
	Apple_iosPasscodepolicy_test []oval_definitions_5_apple_ios.Passcodepolicy_testElement `xml:"http://oval.mitre.org/XMLSchema/oval-definitions-5#apple_ios passcodepolicy_test,omitempty"`
	// HpuxPatch53_test from namespace http://oval.mitre.org/XMLSchema/oval-definitions-5#hpux
	HpuxPatch53_test []oval_definitions_5_hpux.Patch53_testElement `xml:"http://oval.mitre.org/XMLSchema/oval-definitions-5#hpux patch53_test,omitempty"`
	// HpuxPatch_test from namespace http://oval.mitre.org/XMLSchema/oval-definitions-5#hpux
	HpuxPatch_test []oval_definitions_5_hpux.Patch_testElement `xml:"http://oval.mitre.org/XMLSchema/oval-definitions-5#hpux patch_test,omitempty"`
	// HpuxNdd_test from namespace http://oval.mitre.org/XMLSchema/oval-definitions-5#hpux
	HpuxNdd_test []oval_definitions_5_hpux.Ndd_testElement `xml:"http://oval.mitre.org/XMLSchema/oval-definitions-5#hpux ndd_test,omitempty"`
	// HpuxTrusted_test from namespace http://oval.mitre.org/XMLSchema/oval-definitions-5#hpux
	HpuxTrusted_test []oval_definitions_5_hpux.Trusted_testElement `xml:"http://oval.mitre.org/XMLSchema/oval-definitions-5#hpux trusted_test,omitempty"`
	// HpuxGetconf_test from namespace http://oval.mitre.org/XMLSchema/oval-definitions-5#hpux
	HpuxGetconf_test []oval_definitions_5_hpux.Getconf_testElement `xml:"http://oval.mitre.org/XMLSchema/oval-definitions-5#hpux getconf_test,omitempty"`
	// HpuxSwlist_test from namespace http://oval.mitre.org/XMLSchema/oval-definitions-5#hpux
	HpuxSwlist_test []oval_definitions_5_hpux.Swlist_testElement `xml:"http://oval.mitre.org/XMLSchema/oval-definitions-5#hpux swlist_test,omitempty"`
	// JunosXml_config_test from namespace http://oval.mitre.org/XMLSchema/oval-definitions-5#junos
	JunosXml_config_test []oval_definitions_5_junos.Xml_config_testElement `xml:"http://oval.mitre.org/XMLSchema/oval-definitions-5#junos xml_config_test,omitempty"`
	// JunosVersion_test from namespace http://oval.mitre.org/XMLSchema/oval-definitions-5#junos
	JunosVersion_test []oval_definitions_5_junos.Version_testElement `xml:"http://oval.mitre.org/XMLSchema/oval-definitions-5#junos version_test,omitempty"`
	// JunosXml_show_test from namespace http://oval.mitre.org/XMLSchema/oval-definitions-5#junos
	JunosXml_show_test []oval_definitions_5_junos.Xml_show_testElement `xml:"http://oval.mitre.org/XMLSchema/oval-definitions-5#junos xml_show_test,omitempty"`
	// JunosShow_test from namespace http://oval.mitre.org/XMLSchema/oval-definitions-5#junos
	JunosShow_test []oval_definitions_5_junos.Show_testElement `xml:"http://oval.mitre.org/XMLSchema/oval-definitions-5#junos show_test,omitempty"`
	// NetconfConfig_test from namespace http://oval.mitre.org/XMLSchema/oval-definitions-5#netconf
	NetconfConfig_test []oval_definitions_5_netconf.Config_testElement `xml:"http://oval.mitre.org/XMLSchema/oval-definitions-5#netconf config_test,omitempty"`
	// PixosLine_test from namespace http://oval.mitre.org/XMLSchema/oval-definitions-5#pixos
	PixosLine_test []oval_definitions_5_pixos.Line_testElement `xml:"http://oval.mitre.org/XMLSchema/oval-definitions-5#pixos line_test,omitempty"`
	// PixosVersion_test from namespace http://oval.mitre.org/XMLSchema/oval-definitions-5#pixos
	PixosVersion_test []oval_definitions_5_pixos.Version_testElement `xml:"http://oval.mitre.org/XMLSchema/oval-definitions-5#pixos version_test,omitempty"`
	// WindowsSid_sid_test from namespace http://oval.mitre.org/XMLSchema/oval-definitions-5#windows
	WindowsSid_sid_test []oval_definitions_5_windows.Sid_sid_testElement `xml:"http://oval.mitre.org/XMLSchema/oval-definitions-5#windows sid_sid_test,omitempty"`
	// WindowsPasswordpolicy_test from namespace http://oval.mitre.org/XMLSchema/oval-definitions-5#windows
	WindowsPasswordpolicy_test []oval_definitions_5_windows.Passwordpolicy_testElement `xml:"http://oval.mitre.org/XMLSchema/oval-definitions-5#windows passwordpolicy_test,omitempty"`
	// WindowsPeheader_test from namespace http://oval.mitre.org/XMLSchema/oval-definitions-5#windows
	WindowsPeheader_test []oval_definitions_5_windows.Peheader_testElement `xml:"http://oval.mitre.org/XMLSchema/oval-definitions-5#windows peheader_test,omitempty"`
	// WindowsWmi57_test from namespace http://oval.mitre.org/XMLSchema/oval-definitions-5#windows
	WindowsWmi57_test []oval_definitions_5_windows.Wmi57_testElement `xml:"http://oval.mitre.org/XMLSchema/oval-definitions-5#windows wmi57_test,omitempty"`
	// WindowsActivedirectory_test from namespace http://oval.mitre.org/XMLSchema/oval-definitions-5#windows
	WindowsActivedirectory_test []oval_definitions_5_windows.Activedirectory_testElement `xml:"http://oval.mitre.org/XMLSchema/oval-definitions-5#windows activedirectory_test,omitempty"`
	// WindowsServiceeffectiverights_test from namespace http://oval.mitre.org/XMLSchema/oval-definitions-5#windows
	WindowsServiceeffectiverights_test []oval_definitions_5_windows.Serviceeffectiverights_testElement `xml:"http://oval.mitre.org/XMLSchema/oval-definitions-5#windows serviceeffectiverights_test,omitempty"`
	// WindowsAuditeventpolicy_test from namespace http://oval.mitre.org/XMLSchema/oval-definitions-5#windows
	WindowsAuditeventpolicy_test []oval_definitions_5_windows.Auditeventpolicy_testElement `xml:"http://oval.mitre.org/XMLSchema/oval-definitions-5#windows auditeventpolicy_test,omitempty"`
	// WindowsJunction_test from namespace http://oval.mitre.org/XMLSchema/oval-definitions-5#windows
	WindowsJunction_test []oval_definitions_5_windows.Junction_testElement `xml:"http://oval.mitre.org/XMLSchema/oval-definitions-5#windows junction_test,omitempty"`
	// WindowsSharedresource_test from namespace http://oval.mitre.org/XMLSchema/oval-definitions-5#windows
	WindowsSharedresource_test []oval_definitions_5_windows.Sharedresource_testElement `xml:"http://oval.mitre.org/XMLSchema/oval-definitions-5#windows sharedresource_test,omitempty"`
	// WindowsCmdlet_test from namespace http://oval.mitre.org/XMLSchema/oval-definitions-5#windows
	WindowsCmdlet_test []oval_definitions_5_windows.Cmdlet_testElement `xml:"http://oval.mitre.org/XMLSchema/oval-definitions-5#windows cmdlet_test,omitempty"`
	// WindowsUac_test from namespace http://oval.mitre.org/XMLSchema/oval-definitions-5#windows
	WindowsUac_test []oval_definitions_5_windows.Uac_testElement `xml:"http://oval.mitre.org/XMLSchema/oval-definitions-5#windows uac_test,omitempty"`
	// WindowsVolume_test from namespace http://oval.mitre.org/XMLSchema/oval-definitions-5#windows
	WindowsVolume_test []oval_definitions_5_windows.Volume_testElement `xml:"http://oval.mitre.org/XMLSchema/oval-definitions-5#windows volume_test,omitempty"`
	// WindowsSid_test from namespace http://oval.mitre.org/XMLSchema/oval-definitions-5#windows
	WindowsSid_test []oval_definitions_5_windows.Sid_testElement `xml:"http://oval.mitre.org/XMLSchema/oval-definitions-5#windows sid_test,omitempty"`
	// WindowsAuditeventpolicysubcategories_test from namespace http://oval.mitre.org/XMLSchema/oval-definitions-5#windows
	WindowsAuditeventpolicysubcategories_test []oval_definitions_5_windows.Auditeventpolicysubcategories_testElement `xml:"http://oval.mitre.org/XMLSchema/oval-definitions-5#windows auditeventpolicysubcategories_test,omitempty"`
	// WindowsInterface_test from namespace http://oval.mitre.org/XMLSchema/oval-definitions-5#windows
	WindowsInterface_test []oval_definitions_5_windows.Interface_testElement `xml:"http://oval.mitre.org/XMLSchema/oval-definitions-5#windows interface_test,omitempty"`
	// WindowsRegkeyauditedpermissions53_test from namespace http://oval.mitre.org/XMLSchema/oval-definitions-5#windows
	WindowsRegkeyauditedpermissions53_test []oval_definitions_5_windows.Regkeyauditedpermissions53_testElement `xml:"http://oval.mitre.org/XMLSchema/oval-definitions-5#windows regkeyauditedpermissions53_test,omitempty"`
	// WindowsWuaupdatesearcher_test from namespace http://oval.mitre.org/XMLSchema/oval-definitions-5#windows
	WindowsWuaupdatesearcher_test []oval_definitions_5_windows.Wuaupdatesearcher_testElement `xml:"http://oval.mitre.org/XMLSchema/oval-definitions-5#windows wuaupdatesearcher_test,omitempty"`
	// WindowsFileeffectiverights53_test from namespace http://oval.mitre.org/XMLSchema/oval-definitions-5#windows
	WindowsFileeffectiverights53_test []oval_definitions_5_windows.Fileeffectiverights53_testElement `xml:"http://oval.mitre.org/XMLSchema/oval-definitions-5#windows fileeffectiverights53_test,omitempty"`
	// WindowsSharedresourceeffectiverights_test from namespace http://oval.mitre.org/XMLSchema/oval-definitions-5#windows
	WindowsSharedresourceeffectiverights_test []oval_definitions_5_windows.Sharedresourceeffectiverights_testElement `xml:"http://oval.mitre.org/XMLSchema/oval-definitions-5#windows sharedresourceeffectiverights_test,omitempty"`
	// WindowsUser_sid_test from namespace http://oval.mitre.org/XMLSchema/oval-definitions-5#windows
	WindowsUser_sid_test []oval_definitions_5_windows.User_sid_testElement `xml:"http://oval.mitre.org/XMLSchema/oval-definitions-5#windows user_sid_test,omitempty"`
	// WindowsActivedirectory57_test from namespace http://oval.mitre.org/XMLSchema/oval-definitions-5#windows
	WindowsActivedirectory57_test []oval_definitions_5_windows.Activedirectory57_testElement `xml:"http://oval.mitre.org/XMLSchema/oval-definitions-5#windows activedirectory57_test,omitempty"`
	// WindowsFileauditedpermissions_test from namespace http://oval.mitre.org/XMLSchema/oval-definitions-5#windows
	WindowsFileauditedpermissions_test []oval_definitions_5_windows.Fileauditedpermissions_testElement `xml:"http://oval.mitre.org/XMLSchema/oval-definitions-5#windows fileauditedpermissions_test,omitempty"`
	// WindowsFileeffectiverights_test from namespace http://oval.mitre.org/XMLSchema/oval-definitions-5#windows
	WindowsFileeffectiverights_test []oval_definitions_5_windows.Fileeffectiverights_testElement `xml:"http://oval.mitre.org/XMLSchema/oval-definitions-5#windows fileeffectiverights_test,omitempty"`
	// WindowsGroup_sid_test from namespace http://oval.mitre.org/XMLSchema/oval-definitions-5#windows
	WindowsGroup_sid_test []oval_definitions_5_windows.Group_sid_testElement `xml:"http://oval.mitre.org/XMLSchema/oval-definitions-5#windows group_sid_test,omitempty"`
	// WindowsUser_test from namespace http://oval.mitre.org/XMLSchema/oval-definitions-5#windows
	WindowsUser_test []oval_definitions_5_windows.User_testElement `xml:"http://oval.mitre.org/XMLSchema/oval-definitions-5#windows user_test,omitempty"`
	// WindowsLockoutpolicy_test from namespace http://oval.mitre.org/XMLSchema/oval-definitions-5#windows
	WindowsLockoutpolicy_test []oval_definitions_5_windows.Lockoutpolicy_testElement `xml:"http://oval.mitre.org/XMLSchema/oval-definitions-5#windows lockoutpolicy_test,omitempty"`
	// WindowsService_test from namespace http://oval.mitre.org/XMLSchema/oval-definitions-5#windows
	WindowsService_test []oval_definitions_5_windows.Service_testElement `xml:"http://oval.mitre.org/XMLSchema/oval-definitions-5#windows service_test,omitempty"`
	// WindowsFileauditedpermissions53_test from namespace http://oval.mitre.org/XMLSchema/oval-definitions-5#windows
	WindowsFileauditedpermissions53_test []oval_definitions_5_windows.Fileauditedpermissions53_testElement `xml:"http://oval.mitre.org/XMLSchema/oval-definitions-5#windows fileauditedpermissions53_test,omitempty"`
	// WindowsUser_sid55_test from namespace http://oval.mitre.org/XMLSchema/oval-definitions-5#windows
	WindowsUser_sid55_test []oval_definitions_5_windows.User_sid55_testElement `xml:"http://oval.mitre.org/XMLSchema/oval-definitions-5#windows user_sid55_test,omitempty"`
	// WindowsGroup_test from namespace http://oval.mitre.org/XMLSchema/oval-definitions-5#windows
	WindowsGroup_test []oval_definitions_5_windows.Group_testElement `xml:"http://oval.mitre.org/XMLSchema/oval-definitions-5#windows group_test,omitempty"`
	// WindowsPort_test from namespace http://oval.mitre.org/XMLSchema/oval-definitions-5#windows
	WindowsPort_test []oval_definitions_5_windows.Port_testElement `xml:"http://oval.mitre.org/XMLSchema/oval-definitions-5#windows port_test,omitempty"`
	// WindowsRegistry_test from namespace http://oval.mitre.org/XMLSchema/oval-definitions-5#windows
	WindowsRegistry_test []oval_definitions_5_windows.Registry_testElement `xml:"http://oval.mitre.org/XMLSchema/oval-definitions-5#windows registry_test,omitempty"`
	// WindowsRegkeyeffectiverights53_test from namespace http://oval.mitre.org/XMLSchema/oval-definitions-5#windows
	WindowsRegkeyeffectiverights53_test []oval_definitions_5_windows.Regkeyeffectiverights53_testElement `xml:"http://oval.mitre.org/XMLSchema/oval-definitions-5#windows regkeyeffectiverights53_test,omitempty"`
	// WindowsSystemmetric_test from namespace http://oval.mitre.org/XMLSchema/oval-definitions-5#windows
	WindowsSystemmetric_test []oval_definitions_5_windows.Systemmetric_testElement `xml:"http://oval.mitre.org/XMLSchema/oval-definitions-5#windows systemmetric_test,omitempty"`
	// WindowsSharedresourceauditedpermissions_test from namespace http://oval.mitre.org/XMLSchema/oval-definitions-5#windows
	WindowsSharedresourceauditedpermissions_test []oval_definitions_5_windows.Sharedresourceauditedpermissions_testElement `xml:"http://oval.mitre.org/XMLSchema/oval-definitions-5#windows sharedresourceauditedpermissions_test,omitempty"`
	// WindowsDnscache_test from namespace http://oval.mitre.org/XMLSchema/oval-definitions-5#windows
	WindowsDnscache_test []oval_definitions_5_windows.Dnscache_testElement `xml:"http://oval.mitre.org/XMLSchema/oval-definitions-5#windows dnscache_test,omitempty"`
	// WindowsLicense_test from namespace http://oval.mitre.org/XMLSchema/oval-definitions-5#windows
	WindowsLicense_test []oval_definitions_5_windows.License_testElement `xml:"http://oval.mitre.org/XMLSchema/oval-definitions-5#windows license_test,omitempty"`
	// WindowsProcess_test from namespace http://oval.mitre.org/XMLSchema/oval-definitions-5#windows
	WindowsProcess_test []oval_definitions_5_windows.Process_testElement `xml:"http://oval.mitre.org/XMLSchema/oval-definitions-5#windows process_test,omitempty"`
	// WindowsPrintereffectiverights_test from namespace http://oval.mitre.org/XMLSchema/oval-definitions-5#windows
	WindowsPrintereffectiverights_test []oval_definitions_5_windows.Printereffectiverights_testElement `xml:"http://oval.mitre.org/XMLSchema/oval-definitions-5#windows printereffectiverights_test,omitempty"`
	// WindowsUserright_test from namespace http://oval.mitre.org/XMLSchema/oval-definitions-5#windows
	WindowsUserright_test []oval_definitions_5_windows.Userright_testElement `xml:"http://oval.mitre.org/XMLSchema/oval-definitions-5#windows userright_test,omitempty"`
	// WindowsWmi_test from namespace http://oval.mitre.org/XMLSchema/oval-definitions-5#windows
	WindowsWmi_test []oval_definitions_5_windows.Wmi_testElement `xml:"http://oval.mitre.org/XMLSchema/oval-definitions-5#windows wmi_test,omitempty"`
	// WindowsMetabase_test from namespace http://oval.mitre.org/XMLSchema/oval-definitions-5#windows
	WindowsMetabase_test []oval_definitions_5_windows.Metabase_testElement `xml:"http://oval.mitre.org/XMLSchema/oval-definitions-5#windows metabase_test,omitempty"`
	// WindowsNtuser_test from namespace http://oval.mitre.org/XMLSchema/oval-definitions-5#windows
	WindowsNtuser_test []oval_definitions_5_windows.Ntuser_testElement `xml:"http://oval.mitre.org/XMLSchema/oval-definitions-5#windows ntuser_test,omitempty"`
	// WindowsProcess58_test from namespace http://oval.mitre.org/XMLSchema/oval-definitions-5#windows
	WindowsProcess58_test []oval_definitions_5_windows.Process58_testElement `xml:"http://oval.mitre.org/XMLSchema/oval-definitions-5#windows process58_test,omitempty"`
	// WindowsRegkeyauditedpermissions_test from namespace http://oval.mitre.org/XMLSchema/oval-definitions-5#windows
	WindowsRegkeyauditedpermissions_test []oval_definitions_5_windows.Regkeyauditedpermissions_testElement `xml:"http://oval.mitre.org/XMLSchema/oval-definitions-5#windows regkeyauditedpermissions_test,omitempty"`
	// WindowsRegkeyeffectiverights_test from namespace http://oval.mitre.org/XMLSchema/oval-definitions-5#windows
	WindowsRegkeyeffectiverights_test []oval_definitions_5_windows.Regkeyeffectiverights_testElement `xml:"http://oval.mitre.org/XMLSchema/oval-definitions-5#windows regkeyeffectiverights_test,omitempty"`
	// WindowsFile_test from namespace http://oval.mitre.org/XMLSchema/oval-definitions-5#windows
	WindowsFile_test []oval_definitions_5_windows.File_testElement `xml:"http://oval.mitre.org/XMLSchema/oval-definitions-5#windows file_test,omitempty"`
	// WindowsAccesstoken_test from namespace http://oval.mitre.org/XMLSchema/oval-definitions-5#windows
	WindowsAccesstoken_test []oval_definitions_5_windows.Accesstoken_testElement `xml:"http://oval.mitre.org/XMLSchema/oval-definitions-5#windows accesstoken_test,omitempty"`
	// LinuxSystemdunitproperty_test from namespace http://oval.mitre.org/XMLSchema/oval-definitions-5#linux
	LinuxSystemdunitproperty_test []oval_definitions_5_linux.Systemdunitproperty_testElement `xml:"http://oval.mitre.org/XMLSchema/oval-definitions-5#linux systemdunitproperty_test,omitempty"`
	// LinuxIflisteners_test from namespace http://oval.mitre.org/XMLSchema/oval-definitions-5#linux
	LinuxIflisteners_test []oval_definitions_5_linux.Iflisteners_testElement `xml:"http://oval.mitre.org/XMLSchema/oval-definitions-5#linux iflisteners_test,omitempty"`
	// LinuxSlackwarepkginfo_test from namespace http://oval.mitre.org/XMLSchema/oval-definitions-5#linux
	LinuxSlackwarepkginfo_test []oval_definitions_5_linux.Slackwarepkginfo_testElement `xml:"http://oval.mitre.org/XMLSchema/oval-definitions-5#linux slackwarepkginfo_test,omitempty"`
	// LinuxSelinuxboolean_test from namespace http://oval.mitre.org/XMLSchema/oval-definitions-5#linux
	LinuxSelinuxboolean_test []oval_definitions_5_linux.Selinuxboolean_testElement `xml:"http://oval.mitre.org/XMLSchema/oval-definitions-5#linux selinuxboolean_test,omitempty"`
	// LinuxInetlisteningservers_test from namespace http://oval.mitre.org/XMLSchema/oval-definitions-5#linux
	LinuxInetlisteningservers_test []oval_definitions_5_linux.Inetlisteningservers_testElement `xml:"http://oval.mitre.org/XMLSchema/oval-definitions-5#linux inetlisteningservers_test,omitempty"`
	// LinuxRpmverify_test from namespace http://oval.mitre.org/XMLSchema/oval-definitions-5#linux
	LinuxRpmverify_test []oval_definitions_5_linux.Rpmverify_testElement `xml:"http://oval.mitre.org/XMLSchema/oval-definitions-5#linux rpmverify_test,omitempty"`
	// LinuxDpkginfo_test from namespace http://oval.mitre.org/XMLSchema/oval-definitions-5#linux
	LinuxDpkginfo_test []oval_definitions_5_linux.Dpkginfo_testElement `xml:"http://oval.mitre.org/XMLSchema/oval-definitions-5#linux dpkginfo_test,omitempty"`
	// LinuxRpmverifyfile_test from namespace http://oval.mitre.org/XMLSchema/oval-definitions-5#linux
	LinuxRpmverifyfile_test []oval_definitions_5_linux.Rpmverifyfile_testElement `xml:"http://oval.mitre.org/XMLSchema/oval-definitions-5#linux rpmverifyfile_test,omitempty"`
	// LinuxSystemdunitdependency_test from namespace http://oval.mitre.org/XMLSchema/oval-definitions-5#linux
	LinuxSystemdunitdependency_test []oval_definitions_5_linux.Systemdunitdependency_testElement `xml:"http://oval.mitre.org/XMLSchema/oval-definitions-5#linux systemdunitdependency_test,omitempty"`
	// LinuxRpminfo_test from namespace http://oval.mitre.org/XMLSchema/oval-definitions-5#linux
	LinuxRpminfo_test []oval_definitions_5_linux.Rpminfo_testElement `xml:"http://oval.mitre.org/XMLSchema/oval-definitions-5#linux rpminfo_test,omitempty"`
	// LinuxSelinuxsecuritycontext_test from namespace http://oval.mitre.org/XMLSchema/oval-definitions-5#linux
	LinuxSelinuxsecuritycontext_test []oval_definitions_5_linux.Selinuxsecuritycontext_testElement `xml:"http://oval.mitre.org/XMLSchema/oval-definitions-5#linux selinuxsecuritycontext_test,omitempty"`
	// LinuxApparmorstatus_test from namespace http://oval.mitre.org/XMLSchema/oval-definitions-5#linux
	LinuxApparmorstatus_test []oval_definitions_5_linux.Apparmorstatus_testElement `xml:"http://oval.mitre.org/XMLSchema/oval-definitions-5#linux apparmorstatus_test,omitempty"`
	// LinuxPartition_test from namespace http://oval.mitre.org/XMLSchema/oval-definitions-5#linux
	LinuxPartition_test []oval_definitions_5_linux.Partition_testElement `xml:"http://oval.mitre.org/XMLSchema/oval-definitions-5#linux partition_test,omitempty"`
	// LinuxRpmverifypackage_test from namespace http://oval.mitre.org/XMLSchema/oval-definitions-5#linux
	LinuxRpmverifypackage_test []oval_definitions_5_linux.Rpmverifypackage_testElement `xml:"http://oval.mitre.org/XMLSchema/oval-definitions-5#linux rpmverifypackage_test,omitempty"`
	// AsaService_policy_test from namespace http://oval.mitre.org/XMLSchema/oval-definitions-5#asa
	AsaService_policy_test []oval_definitions_5_asa.Service_policy_testElement `xml:"http://oval.mitre.org/XMLSchema/oval-definitions-5#asa service_policy_test,omitempty"`
	// AsaInterface_test from namespace http://oval.mitre.org/XMLSchema/oval-definitions-5#asa
	AsaInterface_test []oval_definitions_5_asa.Interface_testElement `xml:"http://oval.mitre.org/XMLSchema/oval-definitions-5#asa interface_test,omitempty"`
	// AsaLine_test from namespace http://oval.mitre.org/XMLSchema/oval-definitions-5#asa
	AsaLine_test []oval_definitions_5_asa.Line_testElement `xml:"http://oval.mitre.org/XMLSchema/oval-definitions-5#asa line_test,omitempty"`
	// AsaTcp_map_test from namespace http://oval.mitre.org/XMLSchema/oval-definitions-5#asa
	AsaTcp_map_test []oval_definitions_5_asa.Tcp_map_testElement `xml:"http://oval.mitre.org/XMLSchema/oval-definitions-5#asa tcp_map_test,omitempty"`
	// AsaVersion_test from namespace http://oval.mitre.org/XMLSchema/oval-definitions-5#asa
	AsaVersion_test []oval_definitions_5_asa.Version_testElement `xml:"http://oval.mitre.org/XMLSchema/oval-definitions-5#asa version_test,omitempty"`
	// AsaClass_map_test from namespace http://oval.mitre.org/XMLSchema/oval-definitions-5#asa
	AsaClass_map_test []oval_definitions_5_asa.Class_map_testElement `xml:"http://oval.mitre.org/XMLSchema/oval-definitions-5#asa class_map_test,omitempty"`
	// AsaSnmp_host_test from namespace http://oval.mitre.org/XMLSchema/oval-definitions-5#asa
	AsaSnmp_host_test []oval_definitions_5_asa.Snmp_host_testElement `xml:"http://oval.mitre.org/XMLSchema/oval-definitions-5#asa snmp_host_test,omitempty"`
	// AsaSnmp_group_test from namespace http://oval.mitre.org/XMLSchema/oval-definitions-5#asa
	AsaSnmp_group_test []oval_definitions_5_asa.Snmp_group_testElement `xml:"http://oval.mitre.org/XMLSchema/oval-definitions-5#asa snmp_group_test,omitempty"`
	// AsaSnmp_user_test from namespace http://oval.mitre.org/XMLSchema/oval-definitions-5#asa
	AsaSnmp_user_test []oval_definitions_5_asa.Snmp_user_testElement `xml:"http://oval.mitre.org/XMLSchema/oval-definitions-5#asa snmp_user_test,omitempty"`
	// AsaPolicy_map_test from namespace http://oval.mitre.org/XMLSchema/oval-definitions-5#asa
	AsaPolicy_map_test []oval_definitions_5_asa.Policy_map_testElement `xml:"http://oval.mitre.org/XMLSchema/oval-definitions-5#asa policy_map_test,omitempty"`
	// AsaAcl_test from namespace http://oval.mitre.org/XMLSchema/oval-definitions-5#asa
	AsaAcl_test []oval_definitions_5_asa.Acl_testElement `xml:"http://oval.mitre.org/XMLSchema/oval-definitions-5#asa acl_test,omitempty"`
	// IosxeLine_test from namespace http://oval.mitre.org/XMLSchema/oval-definitions-5#iosxe
	IosxeLine_test []oval_definitions_5_iosxe.Line_testElement `xml:"http://oval.mitre.org/XMLSchema/oval-definitions-5#iosxe line_test,omitempty"`
	// IosxeSnmpgroup_test from namespace http://oval.mitre.org/XMLSchema/oval-definitions-5#iosxe
	IosxeSnmpgroup_test []oval_definitions_5_iosxe.Snmpgroup_testElement `xml:"http://oval.mitre.org/XMLSchema/oval-definitions-5#iosxe snmpgroup_test,omitempty"`
	// IosxeVersion_test from namespace http://oval.mitre.org/XMLSchema/oval-definitions-5#iosxe
	IosxeVersion_test []oval_definitions_5_iosxe.Version_testElement `xml:"http://oval.mitre.org/XMLSchema/oval-definitions-5#iosxe version_test,omitempty"`
	// IosxeBgpneighbor_test from namespace http://oval.mitre.org/XMLSchema/oval-definitions-5#iosxe
	IosxeBgpneighbor_test []oval_definitions_5_iosxe.Bgpneighbor_testElement `xml:"http://oval.mitre.org/XMLSchema/oval-definitions-5#iosxe bgpneighbor_test,omitempty"`
	// IosxeSnmpcommunity_test from namespace http://oval.mitre.org/XMLSchema/oval-definitions-5#iosxe
	IosxeSnmpcommunity_test []oval_definitions_5_iosxe.Snmpcommunity_testElement `xml:"http://oval.mitre.org/XMLSchema/oval-definitions-5#iosxe snmpcommunity_test,omitempty"`
	// IosxeRouter_test from namespace http://oval.mitre.org/XMLSchema/oval-definitions-5#iosxe
	IosxeRouter_test []oval_definitions_5_iosxe.Router_testElement `xml:"http://oval.mitre.org/XMLSchema/oval-definitions-5#iosxe router_test,omitempty"`
	// IosxeSnmpview_test from namespace http://oval.mitre.org/XMLSchema/oval-definitions-5#iosxe
	IosxeSnmpview_test []oval_definitions_5_iosxe.Snmpview_testElement `xml:"http://oval.mitre.org/XMLSchema/oval-definitions-5#iosxe snmpview_test,omitempty"`
	// IosxeInterface_test from namespace http://oval.mitre.org/XMLSchema/oval-definitions-5#iosxe
	IosxeInterface_test []oval_definitions_5_iosxe.Interface_testElement `xml:"http://oval.mitre.org/XMLSchema/oval-definitions-5#iosxe interface_test,omitempty"`
	// IosxeSection_test from namespace http://oval.mitre.org/XMLSchema/oval-definitions-5#iosxe
	IosxeSection_test []oval_definitions_5_iosxe.Section_testElement `xml:"http://oval.mitre.org/XMLSchema/oval-definitions-5#iosxe section_test,omitempty"`
	// IosxeAcl_test from namespace http://oval.mitre.org/XMLSchema/oval-definitions-5#iosxe
	IosxeAcl_test []oval_definitions_5_iosxe.Acl_testElement `xml:"http://oval.mitre.org/XMLSchema/oval-definitions-5#iosxe acl_test,omitempty"`
	// IosxeGlobal_test from namespace http://oval.mitre.org/XMLSchema/oval-definitions-5#iosxe
	IosxeGlobal_test []oval_definitions_5_iosxe.Global_testElement `xml:"http://oval.mitre.org/XMLSchema/oval-definitions-5#iosxe global_test,omitempty"`
	// IosxeRoutingprotocolauthintf_test from namespace http://oval.mitre.org/XMLSchema/oval-definitions-5#iosxe
	IosxeRoutingprotocolauthintf_test []oval_definitions_5_iosxe.Routingprotocolauthintf_testElement `xml:"http://oval.mitre.org/XMLSchema/oval-definitions-5#iosxe routingprotocolauthintf_test,omitempty"`
	// IosxeSnmphost_test from namespace http://oval.mitre.org/XMLSchema/oval-definitions-5#iosxe
	IosxeSnmphost_test []oval_definitions_5_iosxe.Snmphost_testElement `xml:"http://oval.mitre.org/XMLSchema/oval-definitions-5#iosxe snmphost_test,omitempty"`
	// IosxeSnmpuser_test from namespace http://oval.mitre.org/XMLSchema/oval-definitions-5#iosxe
	IosxeSnmpuser_test []oval_definitions_5_iosxe.Snmpuser_testElement `xml:"http://oval.mitre.org/XMLSchema/oval-definitions-5#iosxe snmpuser_test,omitempty"`
	// MacosKeychain_test from namespace http://oval.mitre.org/XMLSchema/oval-definitions-5#macos
	MacosKeychain_test []oval_definitions_5_macos.Keychain_testElement `xml:"http://oval.mitre.org/XMLSchema/oval-definitions-5#macos keychain_test,omitempty"`
	// MacosInetlisteningserver510_test from namespace http://oval.mitre.org/XMLSchema/oval-definitions-5#macos
	MacosInetlisteningserver510_test []oval_definitions_5_macos.Inetlisteningserver510_testElement `xml:"http://oval.mitre.org/XMLSchema/oval-definitions-5#macos inetlisteningserver510_test,omitempty"`
	// MacosGatekeeper_test from namespace http://oval.mitre.org/XMLSchema/oval-definitions-5#macos
	MacosGatekeeper_test []oval_definitions_5_macos.Gatekeeper_testElement `xml:"http://oval.mitre.org/XMLSchema/oval-definitions-5#macos gatekeeper_test,omitempty"`
	// MacosAccountinfo_test from namespace http://oval.mitre.org/XMLSchema/oval-definitions-5#macos
	MacosAccountinfo_test []oval_definitions_5_macos.Accountinfo_testElement `xml:"http://oval.mitre.org/XMLSchema/oval-definitions-5#macos accountinfo_test,omitempty"`
	// MacosPwpolicy59_test from namespace http://oval.mitre.org/XMLSchema/oval-definitions-5#macos
	MacosPwpolicy59_test []oval_definitions_5_macos.Pwpolicy59_testElement `xml:"http://oval.mitre.org/XMLSchema/oval-definitions-5#macos pwpolicy59_test,omitempty"`
	// MacosSystemsetup_test from namespace http://oval.mitre.org/XMLSchema/oval-definitions-5#macos
	MacosSystemsetup_test []oval_definitions_5_macos.Systemsetup_testElement `xml:"http://oval.mitre.org/XMLSchema/oval-definitions-5#macos systemsetup_test,omitempty"`
	// MacosPwpolicy_test from namespace http://oval.mitre.org/XMLSchema/oval-definitions-5#macos
	MacosPwpolicy_test []oval_definitions_5_macos.Pwpolicy_testElement `xml:"http://oval.mitre.org/XMLSchema/oval-definitions-5#macos pwpolicy_test,omitempty"`
	// MacosCorestorage_test from namespace http://oval.mitre.org/XMLSchema/oval-definitions-5#macos
	MacosCorestorage_test []oval_definitions_5_macos.Corestorage_testElement `xml:"http://oval.mitre.org/XMLSchema/oval-definitions-5#macos corestorage_test,omitempty"`
	// MacosNvram_test from namespace http://oval.mitre.org/XMLSchema/oval-definitions-5#macos
	MacosNvram_test []oval_definitions_5_macos.Nvram_testElement `xml:"http://oval.mitre.org/XMLSchema/oval-definitions-5#macos nvram_test,omitempty"`
	// MacosPlist511_test from namespace http://oval.mitre.org/XMLSchema/oval-definitions-5#macos
	MacosPlist511_test []oval_definitions_5_macos.Plist511_testElement `xml:"http://oval.mitre.org/XMLSchema/oval-definitions-5#macos plist511_test,omitempty"`
	// MacosDiskutil_test from namespace http://oval.mitre.org/XMLSchema/oval-definitions-5#macos
	MacosDiskutil_test []oval_definitions_5_macos.Diskutil_testElement `xml:"http://oval.mitre.org/XMLSchema/oval-definitions-5#macos diskutil_test,omitempty"`
	// MacosInetlisteningservers_test from namespace http://oval.mitre.org/XMLSchema/oval-definitions-5#macos
	MacosInetlisteningservers_test []oval_definitions_5_macos.Inetlisteningservers_testElement `xml:"http://oval.mitre.org/XMLSchema/oval-definitions-5#macos inetlisteningservers_test,omitempty"`
	// MacosLaunchd_test from namespace http://oval.mitre.org/XMLSchema/oval-definitions-5#macos
	MacosLaunchd_test []oval_definitions_5_macos.Launchd_testElement `xml:"http://oval.mitre.org/XMLSchema/oval-definitions-5#macos launchd_test,omitempty"`
	// MacosSystemprofiler_test from namespace http://oval.mitre.org/XMLSchema/oval-definitions-5#macos
	MacosSystemprofiler_test []oval_definitions_5_macos.Systemprofiler_testElement `xml:"http://oval.mitre.org/XMLSchema/oval-definitions-5#macos systemprofiler_test,omitempty"`
	// MacosPlist510_test from namespace http://oval.mitre.org/XMLSchema/oval-definitions-5#macos
	MacosPlist510_test []oval_definitions_5_macos.Plist510_testElement `xml:"http://oval.mitre.org/XMLSchema/oval-definitions-5#macos plist510_test,omitempty"`
	// MacosPlist_test from namespace http://oval.mitre.org/XMLSchema/oval-definitions-5#macos
	MacosPlist_test []oval_definitions_5_macos.Plist_testElement `xml:"http://oval.mitre.org/XMLSchema/oval-definitions-5#macos plist_test,omitempty"`
	// MacosAuthorizationdb_test from namespace http://oval.mitre.org/XMLSchema/oval-definitions-5#macos
	MacosAuthorizationdb_test []oval_definitions_5_macos.Authorizationdb_testElement `xml:"http://oval.mitre.org/XMLSchema/oval-definitions-5#macos authorizationdb_test,omitempty"`
	// MacosRlimit_test from namespace http://oval.mitre.org/XMLSchema/oval-definitions-5#macos
	MacosRlimit_test []oval_definitions_5_macos.Rlimit_testElement `xml:"http://oval.mitre.org/XMLSchema/oval-definitions-5#macos rlimit_test,omitempty"`
	// MacosSoftwareupdate_test from namespace http://oval.mitre.org/XMLSchema/oval-definitions-5#macos
	MacosSoftwareupdate_test []oval_definitions_5_macos.Softwareupdate_testElement `xml:"http://oval.mitre.org/XMLSchema/oval-definitions-5#macos softwareupdate_test,omitempty"`
	// AixInterim_fix_test from namespace http://oval.mitre.org/XMLSchema/oval-definitions-5#aix
	AixInterim_fix_test []oval_definitions_5_aix.Interim_fix_testElement `xml:"http://oval.mitre.org/XMLSchema/oval-definitions-5#aix interim_fix_test,omitempty"`
	// AixNo_test from namespace http://oval.mitre.org/XMLSchema/oval-definitions-5#aix
	AixNo_test []oval_definitions_5_aix.No_testElement `xml:"http://oval.mitre.org/XMLSchema/oval-definitions-5#aix no_test,omitempty"`
	// AixOslevel_test from namespace http://oval.mitre.org/XMLSchema/oval-definitions-5#aix
	AixOslevel_test []oval_definitions_5_aix.Oslevel_testElement `xml:"http://oval.mitre.org/XMLSchema/oval-definitions-5#aix oslevel_test,omitempty"`
	// AixFileset_test from namespace http://oval.mitre.org/XMLSchema/oval-definitions-5#aix
	AixFileset_test []oval_definitions_5_aix.Fileset_testElement `xml:"http://oval.mitre.org/XMLSchema/oval-definitions-5#aix fileset_test,omitempty"`
	// AixFix_test from namespace http://oval.mitre.org/XMLSchema/oval-definitions-5#aix
	AixFix_test []oval_definitions_5_aix.Fix_testElement `xml:"http://oval.mitre.org/XMLSchema/oval-definitions-5#aix fix_test,omitempty"`
	// EsxVisdkmanagedobject_test from namespace http://oval.mitre.org/XMLSchema/oval-definitions-5#esx
	EsxVisdkmanagedobject_test []oval_definitions_5_esx.Visdkmanagedobject_testElement `xml:"http://oval.mitre.org/XMLSchema/oval-definitions-5#esx visdkmanagedobject_test,omitempty"`
	// EsxPatch56_test from namespace http://oval.mitre.org/XMLSchema/oval-definitions-5#esx
	EsxPatch56_test []oval_definitions_5_esx.Patch56_testElement `xml:"http://oval.mitre.org/XMLSchema/oval-definitions-5#esx patch56_test,omitempty"`
	// EsxVersion_test from namespace http://oval.mitre.org/XMLSchema/oval-definitions-5#esx
	EsxVersion_test []oval_definitions_5_esx.Version_testElement `xml:"http://oval.mitre.org/XMLSchema/oval-definitions-5#esx version_test,omitempty"`
	// EsxPatch_test from namespace http://oval.mitre.org/XMLSchema/oval-definitions-5#esx
	EsxPatch_test []oval_definitions_5_esx.Patch_testElement `xml:"http://oval.mitre.org/XMLSchema/oval-definitions-5#esx patch_test,omitempty"`
	// UnixUname_test from namespace http://oval.mitre.org/XMLSchema/oval-definitions-5#unix
	UnixUname_test []oval_definitions_5_unix.Uname_testElement `xml:"http://oval.mitre.org/XMLSchema/oval-definitions-5#unix uname_test,omitempty"`
	// UnixShadow_test from namespace http://oval.mitre.org/XMLSchema/oval-definitions-5#unix
	UnixShadow_test []oval_definitions_5_unix.Shadow_testElement `xml:"http://oval.mitre.org/XMLSchema/oval-definitions-5#unix shadow_test,omitempty"`
	// UnixFileextendedattribute_test from namespace http://oval.mitre.org/XMLSchema/oval-definitions-5#unix
	UnixFileextendedattribute_test []oval_definitions_5_unix.Fileextendedattribute_testElement `xml:"http://oval.mitre.org/XMLSchema/oval-definitions-5#unix fileextendedattribute_test,omitempty"`
	// UnixPassword_test from namespace http://oval.mitre.org/XMLSchema/oval-definitions-5#unix
	UnixPassword_test []oval_definitions_5_unix.Password_testElement `xml:"http://oval.mitre.org/XMLSchema/oval-definitions-5#unix password_test,omitempty"`
	// UnixRoutingtable_test from namespace http://oval.mitre.org/XMLSchema/oval-definitions-5#unix
	UnixRoutingtable_test []oval_definitions_5_unix.Routingtable_testElement `xml:"http://oval.mitre.org/XMLSchema/oval-definitions-5#unix routingtable_test,omitempty"`
	// UnixInterface_test from namespace http://oval.mitre.org/XMLSchema/oval-definitions-5#unix
	UnixInterface_test []oval_definitions_5_unix.Interface_testElement `xml:"http://oval.mitre.org/XMLSchema/oval-definitions-5#unix interface_test,omitempty"`
	// UnixSymlink_test from namespace http://oval.mitre.org/XMLSchema/oval-definitions-5#unix
	UnixSymlink_test []oval_definitions_5_unix.Symlink_testElement `xml:"http://oval.mitre.org/XMLSchema/oval-definitions-5#unix symlink_test,omitempty"`
	// UnixProcess_test from namespace http://oval.mitre.org/XMLSchema/oval-definitions-5#unix
	UnixProcess_test []oval_definitions_5_unix.Process_testElement `xml:"http://oval.mitre.org/XMLSchema/oval-definitions-5#unix process_test,omitempty"`
	// UnixXinetd_test from namespace http://oval.mitre.org/XMLSchema/oval-definitions-5#unix
	UnixXinetd_test []oval_definitions_5_unix.Xinetd_testElement `xml:"http://oval.mitre.org/XMLSchema/oval-definitions-5#unix xinetd_test,omitempty"`
	// UnixFile_test from namespace http://oval.mitre.org/XMLSchema/oval-definitions-5#unix
	UnixFile_test []oval_definitions_5_unix.File_testElement `xml:"http://oval.mitre.org/XMLSchema/oval-definitions-5#unix file_test,omitempty"`
	// UnixSysctl_test from namespace http://oval.mitre.org/XMLSchema/oval-definitions-5#unix
	UnixSysctl_test []oval_definitions_5_unix.Sysctl_testElement `xml:"http://oval.mitre.org/XMLSchema/oval-definitions-5#unix sysctl_test,omitempty"`
	// UnixGconf_test from namespace http://oval.mitre.org/XMLSchema/oval-definitions-5#unix
	UnixGconf_test []oval_definitions_5_unix.Gconf_testElement `xml:"http://oval.mitre.org/XMLSchema/oval-definitions-5#unix gconf_test,omitempty"`
	// UnixSccs_test from namespace http://oval.mitre.org/XMLSchema/oval-definitions-5#unix
	UnixSccs_test []oval_definitions_5_unix.Sccs_testElement `xml:"http://oval.mitre.org/XMLSchema/oval-definitions-5#unix sccs_test,omitempty"`
	// UnixProcess58_test from namespace http://oval.mitre.org/XMLSchema/oval-definitions-5#unix
	UnixProcess58_test []oval_definitions_5_unix.Process58_testElement `xml:"http://oval.mitre.org/XMLSchema/oval-definitions-5#unix process58_test,omitempty"`
	// UnixDnscache_test from namespace http://oval.mitre.org/XMLSchema/oval-definitions-5#unix
	UnixDnscache_test []oval_definitions_5_unix.Dnscache_testElement `xml:"http://oval.mitre.org/XMLSchema/oval-definitions-5#unix dnscache_test,omitempty"`
	// UnixInetd_test from namespace http://oval.mitre.org/XMLSchema/oval-definitions-5#unix
	UnixInetd_test []oval_definitions_5_unix.Inetd_testElement `xml:"http://oval.mitre.org/XMLSchema/oval-definitions-5#unix inetd_test,omitempty"`
	// UnixRunlevel_test from namespace http://oval.mitre.org/XMLSchema/oval-definitions-5#unix
	UnixRunlevel_test []oval_definitions_5_unix.Runlevel_testElement `xml:"http://oval.mitre.org/XMLSchema/oval-definitions-5#unix runlevel_test,omitempty"`
	// FreebsdPortinfo_test from namespace http://oval.mitre.org/XMLSchema/oval-definitions-5#freebsd
	FreebsdPortinfo_test []oval_definitions_5_freebsd.Portinfo_testElement `xml:"http://oval.mitre.org/XMLSchema/oval-definitions-5#freebsd portinfo_test,omitempty"`
	// IosSnmp_test from namespace http://oval.mitre.org/XMLSchema/oval-definitions-5#ios
	IosSnmp_test []oval_definitions_5_ios.Snmp_testElement `xml:"http://oval.mitre.org/XMLSchema/oval-definitions-5#ios snmp_test,omitempty"`
	// IosSnmphost_test from namespace http://oval.mitre.org/XMLSchema/oval-definitions-5#ios
	IosSnmphost_test []oval_definitions_5_ios.Snmphost_testElement `xml:"http://oval.mitre.org/XMLSchema/oval-definitions-5#ios snmphost_test,omitempty"`
	// IosTclsh_test from namespace http://oval.mitre.org/XMLSchema/oval-definitions-5#ios
	IosTclsh_test []oval_definitions_5_ios.Tclsh_testElement `xml:"http://oval.mitre.org/XMLSchema/oval-definitions-5#ios tclsh_test,omitempty"`
	// IosSnmpuser_test from namespace http://oval.mitre.org/XMLSchema/oval-definitions-5#ios
	IosSnmpuser_test []oval_definitions_5_ios.Snmpuser_testElement `xml:"http://oval.mitre.org/XMLSchema/oval-definitions-5#ios snmpuser_test,omitempty"`
	// IosVersion55_test from namespace http://oval.mitre.org/XMLSchema/oval-definitions-5#ios
	IosVersion55_test []oval_definitions_5_ios.Version55_testElement `xml:"http://oval.mitre.org/XMLSchema/oval-definitions-5#ios version55_test,omitempty"`
	// IosVersion_test from namespace http://oval.mitre.org/XMLSchema/oval-definitions-5#ios
	IosVersion_test []oval_definitions_5_ios.Version_testElement `xml:"http://oval.mitre.org/XMLSchema/oval-definitions-5#ios version_test,omitempty"`
	// IosSnmpcommunity_test from namespace http://oval.mitre.org/XMLSchema/oval-definitions-5#ios
	IosSnmpcommunity_test []oval_definitions_5_ios.Snmpcommunity_testElement `xml:"http://oval.mitre.org/XMLSchema/oval-definitions-5#ios snmpcommunity_test,omitempty"`
	// IosGlobal_test from namespace http://oval.mitre.org/XMLSchema/oval-definitions-5#ios
	IosGlobal_test []oval_definitions_5_ios.Global_testElement `xml:"http://oval.mitre.org/XMLSchema/oval-definitions-5#ios global_test,omitempty"`
	// IosInterface_test from namespace http://oval.mitre.org/XMLSchema/oval-definitions-5#ios
	IosInterface_test []oval_definitions_5_ios.Interface_testElement `xml:"http://oval.mitre.org/XMLSchema/oval-definitions-5#ios interface_test,omitempty"`
	// IosSection_test from namespace http://oval.mitre.org/XMLSchema/oval-definitions-5#ios
	IosSection_test []oval_definitions_5_ios.Section_testElement `xml:"http://oval.mitre.org/XMLSchema/oval-definitions-5#ios section_test,omitempty"`
	// IosSnmpview_test from namespace http://oval.mitre.org/XMLSchema/oval-definitions-5#ios
	IosSnmpview_test []oval_definitions_5_ios.Snmpview_testElement `xml:"http://oval.mitre.org/XMLSchema/oval-definitions-5#ios snmpview_test,omitempty"`
	// IosBgpneighbor_test from namespace http://oval.mitre.org/XMLSchema/oval-definitions-5#ios
	IosBgpneighbor_test []oval_definitions_5_ios.Bgpneighbor_testElement `xml:"http://oval.mitre.org/XMLSchema/oval-definitions-5#ios bgpneighbor_test,omitempty"`
	// IosRoutingprotocolauthintf_test from namespace http://oval.mitre.org/XMLSchema/oval-definitions-5#ios
	IosRoutingprotocolauthintf_test []oval_definitions_5_ios.Routingprotocolauthintf_testElement `xml:"http://oval.mitre.org/XMLSchema/oval-definitions-5#ios routingprotocolauthintf_test,omitempty"`
	// IosLine_test from namespace http://oval.mitre.org/XMLSchema/oval-definitions-5#ios
	IosLine_test []oval_definitions_5_ios.Line_testElement `xml:"http://oval.mitre.org/XMLSchema/oval-definitions-5#ios line_test,omitempty"`
	// IosRouter_test from namespace http://oval.mitre.org/XMLSchema/oval-definitions-5#ios
	IosRouter_test []oval_definitions_5_ios.Router_testElement `xml:"http://oval.mitre.org/XMLSchema/oval-definitions-5#ios router_test,omitempty"`
	// IosSnmpgroup_test from namespace http://oval.mitre.org/XMLSchema/oval-definitions-5#ios
	IosSnmpgroup_test []oval_definitions_5_ios.Snmpgroup_testElement `xml:"http://oval.mitre.org/XMLSchema/oval-definitions-5#ios snmpgroup_test,omitempty"`
	// IosAcl_test from namespace http://oval.mitre.org/XMLSchema/oval-definitions-5#ios
	IosAcl_test []oval_definitions_5_ios.Acl_testElement `xml:"http://oval.mitre.org/XMLSchema/oval-definitions-5#ios acl_test,omitempty"`
	// ApacheHttpd_test from namespace http://oval.mitre.org/XMLSchema/oval-definitions-5#apache
	ApacheHttpd_test []oval_definitions_5_apache.Httpd_testElement `xml:"http://oval.mitre.org/XMLSchema/oval-definitions-5#apache httpd_test,omitempty"`
	// SharepointSpdiagnosticslevel_test from namespace http://oval.mitre.org/XMLSchema/oval-definitions-5#sharepoint
	SharepointSpdiagnosticslevel_test []oval_definitions_5_sharepoint.Spdiagnosticslevel_testElement `xml:"http://oval.mitre.org/XMLSchema/oval-definitions-5#sharepoint spdiagnosticslevel_test,omitempty"`
	// SharepointSppolicy_test from namespace http://oval.mitre.org/XMLSchema/oval-definitions-5#sharepoint
	SharepointSppolicy_test []oval_definitions_5_sharepoint.Sppolicy_testElement `xml:"http://oval.mitre.org/XMLSchema/oval-definitions-5#sharepoint sppolicy_test,omitempty"`
	// SharepointSpwebapplication_test from namespace http://oval.mitre.org/XMLSchema/oval-definitions-5#sharepoint
	SharepointSpwebapplication_test []oval_definitions_5_sharepoint.Spwebapplication_testElement `xml:"http://oval.mitre.org/XMLSchema/oval-definitions-5#sharepoint spwebapplication_test,omitempty"`
	// SharepointSpgroup_test from namespace http://oval.mitre.org/XMLSchema/oval-definitions-5#sharepoint
	SharepointSpgroup_test []oval_definitions_5_sharepoint.Spgroup_testElement `xml:"http://oval.mitre.org/XMLSchema/oval-definitions-5#sharepoint spgroup_test,omitempty"`
	// SharepointSplist_test from namespace http://oval.mitre.org/XMLSchema/oval-definitions-5#sharepoint
	SharepointSplist_test []oval_definitions_5_sharepoint.Splist_testElement `xml:"http://oval.mitre.org/XMLSchema/oval-definitions-5#sharepoint splist_test,omitempty"`
	// SharepointSppolicyfeature_test from namespace http://oval.mitre.org/XMLSchema/oval-definitions-5#sharepoint
	SharepointSppolicyfeature_test []oval_definitions_5_sharepoint.Sppolicyfeature_testElement `xml:"http://oval.mitre.org/XMLSchema/oval-definitions-5#sharepoint sppolicyfeature_test,omitempty"`
	// SharepointSpjobdefinition510_test from namespace http://oval.mitre.org/XMLSchema/oval-definitions-5#sharepoint
	SharepointSpjobdefinition510_test []oval_definitions_5_sharepoint.Spjobdefinition510_testElement `xml:"http://oval.mitre.org/XMLSchema/oval-definitions-5#sharepoint spjobdefinition510_test,omitempty"`
	// SharepointSpweb_test from namespace http://oval.mitre.org/XMLSchema/oval-definitions-5#sharepoint
	SharepointSpweb_test []oval_definitions_5_sharepoint.Spweb_testElement `xml:"http://oval.mitre.org/XMLSchema/oval-definitions-5#sharepoint spweb_test,omitempty"`
	// SharepointSpantivirussettings_test from namespace http://oval.mitre.org/XMLSchema/oval-definitions-5#sharepoint
	SharepointSpantivirussettings_test []oval_definitions_5_sharepoint.Spantivirussettings_testElement `xml:"http://oval.mitre.org/XMLSchema/oval-definitions-5#sharepoint spantivirussettings_test,omitempty"`
	// SharepointSpsite_test from namespace http://oval.mitre.org/XMLSchema/oval-definitions-5#sharepoint
	SharepointSpsite_test []oval_definitions_5_sharepoint.Spsite_testElement `xml:"http://oval.mitre.org/XMLSchema/oval-definitions-5#sharepoint spsite_test,omitempty"`
	// SharepointSpcrawlrule_test from namespace http://oval.mitre.org/XMLSchema/oval-definitions-5#sharepoint
	SharepointSpcrawlrule_test []oval_definitions_5_sharepoint.Spcrawlrule_testElement `xml:"http://oval.mitre.org/XMLSchema/oval-definitions-5#sharepoint spcrawlrule_test,omitempty"`
	// SharepointBestbet_test from namespace http://oval.mitre.org/XMLSchema/oval-definitions-5#sharepoint
	SharepointBestbet_test []oval_definitions_5_sharepoint.Bestbet_testElement `xml:"http://oval.mitre.org/XMLSchema/oval-definitions-5#sharepoint bestbet_test,omitempty"`
	// SharepointInfopolicycoll_test from namespace http://oval.mitre.org/XMLSchema/oval-definitions-5#sharepoint
	SharepointInfopolicycoll_test []oval_definitions_5_sharepoint.Infopolicycoll_testElement `xml:"http://oval.mitre.org/XMLSchema/oval-definitions-5#sharepoint infopolicycoll_test,omitempty"`
	// SharepointSpdiagnosticsservice_test from namespace http://oval.mitre.org/XMLSchema/oval-definitions-5#sharepoint
	SharepointSpdiagnosticsservice_test []oval_definitions_5_sharepoint.Spdiagnosticsservice_testElement `xml:"http://oval.mitre.org/XMLSchema/oval-definitions-5#sharepoint spdiagnosticsservice_test,omitempty"`
	// SharepointSpsiteadministration_test from namespace http://oval.mitre.org/XMLSchema/oval-definitions-5#sharepoint
	SharepointSpsiteadministration_test []oval_definitions_5_sharepoint.Spsiteadministration_testElement `xml:"http://oval.mitre.org/XMLSchema/oval-definitions-5#sharepoint spsiteadministration_test,omitempty"`
	// SharepointSpjobdefinition_test from namespace http://oval.mitre.org/XMLSchema/oval-definitions-5#sharepoint
	SharepointSpjobdefinition_test []oval_definitions_5_sharepoint.Spjobdefinition_testElement `xml:"http://oval.mitre.org/XMLSchema/oval-definitions-5#sharepoint spjobdefinition_test,omitempty"`
	// CatosModule_test from namespace http://oval.mitre.org/XMLSchema/oval-definitions-5#catos
	CatosModule_test []oval_definitions_5_catos.Module_testElement `xml:"http://oval.mitre.org/XMLSchema/oval-definitions-5#catos module_test,omitempty"`
	// CatosVersion55_test from namespace http://oval.mitre.org/XMLSchema/oval-definitions-5#catos
	CatosVersion55_test []oval_definitions_5_catos.Version55_testElement `xml:"http://oval.mitre.org/XMLSchema/oval-definitions-5#catos version55_test,omitempty"`
	// CatosVersion_test from namespace http://oval.mitre.org/XMLSchema/oval-definitions-5#catos
	CatosVersion_test []oval_definitions_5_catos.Version_testElement `xml:"http://oval.mitre.org/XMLSchema/oval-definitions-5#catos version_test,omitempty"`
	// CatosLine_test from namespace http://oval.mitre.org/XMLSchema/oval-definitions-5#catos
	CatosLine_test []oval_definitions_5_catos.Line_testElement `xml:"http://oval.mitre.org/XMLSchema/oval-definitions-5#catos line_test,omitempty"`
	// AndroidNetwork_test from namespace http://oval.mitre.org/XMLSchema/oval-definitions-5#android
	AndroidNetwork_test []oval_definitions_5_android.Network_testElement `xml:"http://oval.mitre.org/XMLSchema/oval-definitions-5#android network_test,omitempty"`
	// AndroidPassword_test from namespace http://oval.mitre.org/XMLSchema/oval-definitions-5#android
	AndroidPassword_test []oval_definitions_5_android.Password_testElement `xml:"http://oval.mitre.org/XMLSchema/oval-definitions-5#android password_test,omitempty"`
	// AndroidCertificate_test from namespace http://oval.mitre.org/XMLSchema/oval-definitions-5#android
	AndroidCertificate_test []oval_definitions_5_android.Certificate_testElement `xml:"http://oval.mitre.org/XMLSchema/oval-definitions-5#android certificate_test,omitempty"`
	// AndroidSystemdetails_test from namespace http://oval.mitre.org/XMLSchema/oval-definitions-5#android
	AndroidSystemdetails_test []oval_definitions_5_android.Systemdetails_testElement `xml:"http://oval.mitre.org/XMLSchema/oval-definitions-5#android systemdetails_test,omitempty"`
	// AndroidWifi_test from namespace http://oval.mitre.org/XMLSchema/oval-definitions-5#android
	AndroidWifi_test []oval_definitions_5_android.Wifi_testElement `xml:"http://oval.mitre.org/XMLSchema/oval-definitions-5#android wifi_test,omitempty"`
	// AndroidBluetooth_test from namespace http://oval.mitre.org/XMLSchema/oval-definitions-5#android
	AndroidBluetooth_test []oval_definitions_5_android.Bluetooth_testElement `xml:"http://oval.mitre.org/XMLSchema/oval-definitions-5#android bluetooth_test,omitempty"`
	// AndroidCamera_test from namespace http://oval.mitre.org/XMLSchema/oval-definitions-5#android
	AndroidCamera_test []oval_definitions_5_android.Camera_testElement `xml:"http://oval.mitre.org/XMLSchema/oval-definitions-5#android camera_test,omitempty"`
	// AndroidWifinetwork_test from namespace http://oval.mitre.org/XMLSchema/oval-definitions-5#android
	AndroidWifinetwork_test []oval_definitions_5_android.Wifinetwork_testElement `xml:"http://oval.mitre.org/XMLSchema/oval-definitions-5#android wifinetwork_test,omitempty"`
	// AndroidTelephony_test from namespace http://oval.mitre.org/XMLSchema/oval-definitions-5#android
	AndroidTelephony_test []oval_definitions_5_android.Telephony_testElement `xml:"http://oval.mitre.org/XMLSchema/oval-definitions-5#android telephony_test,omitempty"`
	// AndroidEncryption_test from namespace http://oval.mitre.org/XMLSchema/oval-definitions-5#android
	AndroidEncryption_test []oval_definitions_5_android.Encryption_testElement `xml:"http://oval.mitre.org/XMLSchema/oval-definitions-5#android encryption_test,omitempty"`
	// AndroidDevicesettings_test from namespace http://oval.mitre.org/XMLSchema/oval-definitions-5#android
	AndroidDevicesettings_test []oval_definitions_5_android.Devicesettings_testElement `xml:"http://oval.mitre.org/XMLSchema/oval-definitions-5#android devicesettings_test,omitempty"`
	// AndroidAppmanager_test from namespace http://oval.mitre.org/XMLSchema/oval-definitions-5#android
	AndroidAppmanager_test []oval_definitions_5_android.Appmanager_testElement `xml:"http://oval.mitre.org/XMLSchema/oval-definitions-5#android appmanager_test,omitempty"`
	// AndroidLocationservice_test from namespace http://oval.mitre.org/XMLSchema/oval-definitions-5#android
	AndroidLocationservice_test []oval_definitions_5_android.Locationservice_testElement `xml:"http://oval.mitre.org/XMLSchema/oval-definitions-5#android locationservice_test,omitempty"`
	// UnknownAttrs captures any attributes not defined in XSD
	UnknownAttrs []xml.Attr `xml:",any,attr,omitempty"`
}

// ObjectsType - Complete version with all substitution group members
// This type is standalone without embedding to avoid UnknownElements stripping xmlns
type ObjectsType struct {
	// LinuxRpmverifyfile_object from namespace http://oval.mitre.org/XMLSchema/oval-definitions-5#linux
	LinuxRpmverifyfile_object []oval_definitions_5_linux.Rpmverifyfile_objectElement `xml:"http://oval.mitre.org/XMLSchema/oval-definitions-5#linux rpmverifyfile_object,omitempty"`
	// LinuxApparmorstatus_object from namespace http://oval.mitre.org/XMLSchema/oval-definitions-5#linux
	LinuxApparmorstatus_object []oval_definitions_5_linux.Apparmorstatus_objectElement `xml:"http://oval.mitre.org/XMLSchema/oval-definitions-5#linux apparmorstatus_object,omitempty"`
	// LinuxDpkginfo_object from namespace http://oval.mitre.org/XMLSchema/oval-definitions-5#linux
	LinuxDpkginfo_object []oval_definitions_5_linux.Dpkginfo_objectElement `xml:"http://oval.mitre.org/XMLSchema/oval-definitions-5#linux dpkginfo_object,omitempty"`
	// LinuxRpmverify_object from namespace http://oval.mitre.org/XMLSchema/oval-definitions-5#linux
	LinuxRpmverify_object []oval_definitions_5_linux.Rpmverify_objectElement `xml:"http://oval.mitre.org/XMLSchema/oval-definitions-5#linux rpmverify_object,omitempty"`
	// LinuxSelinuxboolean_object from namespace http://oval.mitre.org/XMLSchema/oval-definitions-5#linux
	LinuxSelinuxboolean_object []oval_definitions_5_linux.Selinuxboolean_objectElement `xml:"http://oval.mitre.org/XMLSchema/oval-definitions-5#linux selinuxboolean_object,omitempty"`
	// LinuxInetlisteningservers_object from namespace http://oval.mitre.org/XMLSchema/oval-definitions-5#linux
	LinuxInetlisteningservers_object []oval_definitions_5_linux.Inetlisteningservers_objectElement `xml:"http://oval.mitre.org/XMLSchema/oval-definitions-5#linux inetlisteningservers_object,omitempty"`
	// LinuxRpminfo_object from namespace http://oval.mitre.org/XMLSchema/oval-definitions-5#linux
	LinuxRpminfo_object []oval_definitions_5_linux.Rpminfo_objectElement `xml:"http://oval.mitre.org/XMLSchema/oval-definitions-5#linux rpminfo_object,omitempty"`
	// LinuxSlackwarepkginfo_object from namespace http://oval.mitre.org/XMLSchema/oval-definitions-5#linux
	LinuxSlackwarepkginfo_object []oval_definitions_5_linux.Slackwarepkginfo_objectElement `xml:"http://oval.mitre.org/XMLSchema/oval-definitions-5#linux slackwarepkginfo_object,omitempty"`
	// LinuxSystemdunitdependency_object from namespace http://oval.mitre.org/XMLSchema/oval-definitions-5#linux
	LinuxSystemdunitdependency_object []oval_definitions_5_linux.Systemdunitdependency_objectElement `xml:"http://oval.mitre.org/XMLSchema/oval-definitions-5#linux systemdunitdependency_object,omitempty"`
	// LinuxIflisteners_object from namespace http://oval.mitre.org/XMLSchema/oval-definitions-5#linux
	LinuxIflisteners_object []oval_definitions_5_linux.Iflisteners_objectElement `xml:"http://oval.mitre.org/XMLSchema/oval-definitions-5#linux iflisteners_object,omitempty"`
	// LinuxPartition_object from namespace http://oval.mitre.org/XMLSchema/oval-definitions-5#linux
	LinuxPartition_object []oval_definitions_5_linux.Partition_objectElement `xml:"http://oval.mitre.org/XMLSchema/oval-definitions-5#linux partition_object,omitempty"`
	// LinuxRpmverifypackage_object from namespace http://oval.mitre.org/XMLSchema/oval-definitions-5#linux
	LinuxRpmverifypackage_object []oval_definitions_5_linux.Rpmverifypackage_objectElement `xml:"http://oval.mitre.org/XMLSchema/oval-definitions-5#linux rpmverifypackage_object,omitempty"`
	// LinuxSelinuxsecuritycontext_object from namespace http://oval.mitre.org/XMLSchema/oval-definitions-5#linux
	LinuxSelinuxsecuritycontext_object []oval_definitions_5_linux.Selinuxsecuritycontext_objectElement `xml:"http://oval.mitre.org/XMLSchema/oval-definitions-5#linux selinuxsecuritycontext_object,omitempty"`
	// LinuxSystemdunitproperty_object from namespace http://oval.mitre.org/XMLSchema/oval-definitions-5#linux
	LinuxSystemdunitproperty_object []oval_definitions_5_linux.Systemdunitproperty_objectElement `xml:"http://oval.mitre.org/XMLSchema/oval-definitions-5#linux systemdunitproperty_object,omitempty"`
	// AsaInterface_object from namespace http://oval.mitre.org/XMLSchema/oval-definitions-5#asa
	AsaInterface_object []oval_definitions_5_asa.Interface_objectElement `xml:"http://oval.mitre.org/XMLSchema/oval-definitions-5#asa interface_object,omitempty"`
	// AsaVersion_object from namespace http://oval.mitre.org/XMLSchema/oval-definitions-5#asa
	AsaVersion_object []oval_definitions_5_asa.Version_objectElement `xml:"http://oval.mitre.org/XMLSchema/oval-definitions-5#asa version_object,omitempty"`
	// AsaTcp_map_object from namespace http://oval.mitre.org/XMLSchema/oval-definitions-5#asa
	AsaTcp_map_object []oval_definitions_5_asa.Tcp_map_objectElement `xml:"http://oval.mitre.org/XMLSchema/oval-definitions-5#asa tcp_map_object,omitempty"`
	// AsaPolicy_map_object from namespace http://oval.mitre.org/XMLSchema/oval-definitions-5#asa
	AsaPolicy_map_object []oval_definitions_5_asa.Policy_map_objectElement `xml:"http://oval.mitre.org/XMLSchema/oval-definitions-5#asa policy_map_object,omitempty"`
	// AsaService_policy_object from namespace http://oval.mitre.org/XMLSchema/oval-definitions-5#asa
	AsaService_policy_object []oval_definitions_5_asa.Service_policy_objectElement `xml:"http://oval.mitre.org/XMLSchema/oval-definitions-5#asa service_policy_object,omitempty"`
	// AsaSnmp_group_object from namespace http://oval.mitre.org/XMLSchema/oval-definitions-5#asa
	AsaSnmp_group_object []oval_definitions_5_asa.Snmp_group_objectElement `xml:"http://oval.mitre.org/XMLSchema/oval-definitions-5#asa snmp_group_object,omitempty"`
	// AsaLine_object from namespace http://oval.mitre.org/XMLSchema/oval-definitions-5#asa
	AsaLine_object []oval_definitions_5_asa.Line_objectElement `xml:"http://oval.mitre.org/XMLSchema/oval-definitions-5#asa line_object,omitempty"`
	// AsaSnmp_host_object from namespace http://oval.mitre.org/XMLSchema/oval-definitions-5#asa
	AsaSnmp_host_object []oval_definitions_5_asa.Snmp_host_objectElement `xml:"http://oval.mitre.org/XMLSchema/oval-definitions-5#asa snmp_host_object,omitempty"`
	// AsaSnmp_user_object from namespace http://oval.mitre.org/XMLSchema/oval-definitions-5#asa
	AsaSnmp_user_object []oval_definitions_5_asa.Snmp_user_objectElement `xml:"http://oval.mitre.org/XMLSchema/oval-definitions-5#asa snmp_user_object,omitempty"`
	// AsaClass_map_object from namespace http://oval.mitre.org/XMLSchema/oval-definitions-5#asa
	AsaClass_map_object []oval_definitions_5_asa.Class_map_objectElement `xml:"http://oval.mitre.org/XMLSchema/oval-definitions-5#asa class_map_object,omitempty"`
	// AsaAcl_object from namespace http://oval.mitre.org/XMLSchema/oval-definitions-5#asa
	AsaAcl_object []oval_definitions_5_asa.Acl_objectElement `xml:"http://oval.mitre.org/XMLSchema/oval-definitions-5#asa acl_object,omitempty"`
	// IosxeBgpneighbor_object from namespace http://oval.mitre.org/XMLSchema/oval-definitions-5#iosxe
	IosxeBgpneighbor_object []oval_definitions_5_iosxe.Bgpneighbor_objectElement `xml:"http://oval.mitre.org/XMLSchema/oval-definitions-5#iosxe bgpneighbor_object,omitempty"`
	// IosxeSnmpgroup_object from namespace http://oval.mitre.org/XMLSchema/oval-definitions-5#iosxe
	IosxeSnmpgroup_object []oval_definitions_5_iosxe.Snmpgroup_objectElement `xml:"http://oval.mitre.org/XMLSchema/oval-definitions-5#iosxe snmpgroup_object,omitempty"`
	// IosxeInterface_object from namespace http://oval.mitre.org/XMLSchema/oval-definitions-5#iosxe
	IosxeInterface_object []oval_definitions_5_iosxe.Interface_objectElement `xml:"http://oval.mitre.org/XMLSchema/oval-definitions-5#iosxe interface_object,omitempty"`
	// IosxeRoutingprotocolauthintf_object from namespace http://oval.mitre.org/XMLSchema/oval-definitions-5#iosxe
	IosxeRoutingprotocolauthintf_object []oval_definitions_5_iosxe.Routingprotocolauthintf_objectElement `xml:"http://oval.mitre.org/XMLSchema/oval-definitions-5#iosxe routingprotocolauthintf_object,omitempty"`
	// IosxeSnmpuser_object from namespace http://oval.mitre.org/XMLSchema/oval-definitions-5#iosxe
	IosxeSnmpuser_object []oval_definitions_5_iosxe.Snmpuser_objectElement `xml:"http://oval.mitre.org/XMLSchema/oval-definitions-5#iosxe snmpuser_object,omitempty"`
	// IosxeVersion_object from namespace http://oval.mitre.org/XMLSchema/oval-definitions-5#iosxe
	IosxeVersion_object []oval_definitions_5_iosxe.Version_objectElement `xml:"http://oval.mitre.org/XMLSchema/oval-definitions-5#iosxe version_object,omitempty"`
	// IosxeSnmphost_object from namespace http://oval.mitre.org/XMLSchema/oval-definitions-5#iosxe
	IosxeSnmphost_object []oval_definitions_5_iosxe.Snmphost_objectElement `xml:"http://oval.mitre.org/XMLSchema/oval-definitions-5#iosxe snmphost_object,omitempty"`
	// IosxeSnmpcommunity_object from namespace http://oval.mitre.org/XMLSchema/oval-definitions-5#iosxe
	IosxeSnmpcommunity_object []oval_definitions_5_iosxe.Snmpcommunity_objectElement `xml:"http://oval.mitre.org/XMLSchema/oval-definitions-5#iosxe snmpcommunity_object,omitempty"`
	// IosxeSnmpview_object from namespace http://oval.mitre.org/XMLSchema/oval-definitions-5#iosxe
	IosxeSnmpview_object []oval_definitions_5_iosxe.Snmpview_objectElement `xml:"http://oval.mitre.org/XMLSchema/oval-definitions-5#iosxe snmpview_object,omitempty"`
	// IosxeRouter_object from namespace http://oval.mitre.org/XMLSchema/oval-definitions-5#iosxe
	IosxeRouter_object []oval_definitions_5_iosxe.Router_objectElement `xml:"http://oval.mitre.org/XMLSchema/oval-definitions-5#iosxe router_object,omitempty"`
	// IosxeGlobal_object from namespace http://oval.mitre.org/XMLSchema/oval-definitions-5#iosxe
	IosxeGlobal_object []oval_definitions_5_iosxe.Global_objectElement `xml:"http://oval.mitre.org/XMLSchema/oval-definitions-5#iosxe global_object,omitempty"`
	// IosxeSection_object from namespace http://oval.mitre.org/XMLSchema/oval-definitions-5#iosxe
	IosxeSection_object []oval_definitions_5_iosxe.Section_objectElement `xml:"http://oval.mitre.org/XMLSchema/oval-definitions-5#iosxe section_object,omitempty"`
	// IosxeAcl_object from namespace http://oval.mitre.org/XMLSchema/oval-definitions-5#iosxe
	IosxeAcl_object []oval_definitions_5_iosxe.Acl_objectElement `xml:"http://oval.mitre.org/XMLSchema/oval-definitions-5#iosxe acl_object,omitempty"`
	// IosxeLine_object from namespace http://oval.mitre.org/XMLSchema/oval-definitions-5#iosxe
	IosxeLine_object []oval_definitions_5_iosxe.Line_objectElement `xml:"http://oval.mitre.org/XMLSchema/oval-definitions-5#iosxe line_object,omitempty"`
	// MacosRlimit_object from namespace http://oval.mitre.org/XMLSchema/oval-definitions-5#macos
	MacosRlimit_object []oval_definitions_5_macos.Rlimit_objectElement `xml:"http://oval.mitre.org/XMLSchema/oval-definitions-5#macos rlimit_object,omitempty"`
	// MacosAuthorizationdb_object from namespace http://oval.mitre.org/XMLSchema/oval-definitions-5#macos
	MacosAuthorizationdb_object []oval_definitions_5_macos.Authorizationdb_objectElement `xml:"http://oval.mitre.org/XMLSchema/oval-definitions-5#macos authorizationdb_object,omitempty"`
	// MacosPwpolicy_object from namespace http://oval.mitre.org/XMLSchema/oval-definitions-5#macos
	MacosPwpolicy_object []oval_definitions_5_macos.Pwpolicy_objectElement `xml:"http://oval.mitre.org/XMLSchema/oval-definitions-5#macos pwpolicy_object,omitempty"`
	// MacosAccountinfo_object from namespace http://oval.mitre.org/XMLSchema/oval-definitions-5#macos
	MacosAccountinfo_object []oval_definitions_5_macos.Accountinfo_objectElement `xml:"http://oval.mitre.org/XMLSchema/oval-definitions-5#macos accountinfo_object,omitempty"`
	// MacosSystemsetup_object from namespace http://oval.mitre.org/XMLSchema/oval-definitions-5#macos
	MacosSystemsetup_object []oval_definitions_5_macos.Systemsetup_objectElement `xml:"http://oval.mitre.org/XMLSchema/oval-definitions-5#macos systemsetup_object,omitempty"`
	// MacosInetlisteningserver510_object from namespace http://oval.mitre.org/XMLSchema/oval-definitions-5#macos
	MacosInetlisteningserver510_object []oval_definitions_5_macos.Inetlisteningserver510_objectElement `xml:"http://oval.mitre.org/XMLSchema/oval-definitions-5#macos inetlisteningserver510_object,omitempty"`
	// MacosNvram_object from namespace http://oval.mitre.org/XMLSchema/oval-definitions-5#macos
	MacosNvram_object []oval_definitions_5_macos.Nvram_objectElement `xml:"http://oval.mitre.org/XMLSchema/oval-definitions-5#macos nvram_object,omitempty"`
	// MacosInetlisteningservers_object from namespace http://oval.mitre.org/XMLSchema/oval-definitions-5#macos
	MacosInetlisteningservers_object []oval_definitions_5_macos.Inetlisteningservers_objectElement `xml:"http://oval.mitre.org/XMLSchema/oval-definitions-5#macos inetlisteningservers_object,omitempty"`
	// MacosPlist510_object from namespace http://oval.mitre.org/XMLSchema/oval-definitions-5#macos
	MacosPlist510_object []oval_definitions_5_macos.Plist510_objectElement `xml:"http://oval.mitre.org/XMLSchema/oval-definitions-5#macos plist510_object,omitempty"`
	// MacosSoftwareupdate_object from namespace http://oval.mitre.org/XMLSchema/oval-definitions-5#macos
	MacosSoftwareupdate_object []oval_definitions_5_macos.Softwareupdate_objectElement `xml:"http://oval.mitre.org/XMLSchema/oval-definitions-5#macos softwareupdate_object,omitempty"`
	// MacosLaunchd_object from namespace http://oval.mitre.org/XMLSchema/oval-definitions-5#macos
	MacosLaunchd_object []oval_definitions_5_macos.Launchd_objectElement `xml:"http://oval.mitre.org/XMLSchema/oval-definitions-5#macos launchd_object,omitempty"`
	// MacosSystemprofiler_object from namespace http://oval.mitre.org/XMLSchema/oval-definitions-5#macos
	MacosSystemprofiler_object []oval_definitions_5_macos.Systemprofiler_objectElement `xml:"http://oval.mitre.org/XMLSchema/oval-definitions-5#macos systemprofiler_object,omitempty"`
	// MacosDiskutil_object from namespace http://oval.mitre.org/XMLSchema/oval-definitions-5#macos
	MacosDiskutil_object []oval_definitions_5_macos.Diskutil_objectElement `xml:"http://oval.mitre.org/XMLSchema/oval-definitions-5#macos diskutil_object,omitempty"`
	// MacosKeychain_object from namespace http://oval.mitre.org/XMLSchema/oval-definitions-5#macos
	MacosKeychain_object []oval_definitions_5_macos.Keychain_objectElement `xml:"http://oval.mitre.org/XMLSchema/oval-definitions-5#macos keychain_object,omitempty"`
	// MacosCorestorage_object from namespace http://oval.mitre.org/XMLSchema/oval-definitions-5#macos
	MacosCorestorage_object []oval_definitions_5_macos.Corestorage_objectElement `xml:"http://oval.mitre.org/XMLSchema/oval-definitions-5#macos corestorage_object,omitempty"`
	// MacosPlist_object from namespace http://oval.mitre.org/XMLSchema/oval-definitions-5#macos
	MacosPlist_object []oval_definitions_5_macos.Plist_objectElement `xml:"http://oval.mitre.org/XMLSchema/oval-definitions-5#macos plist_object,omitempty"`
	// MacosPlist511_object from namespace http://oval.mitre.org/XMLSchema/oval-definitions-5#macos
	MacosPlist511_object []oval_definitions_5_macos.Plist511_objectElement `xml:"http://oval.mitre.org/XMLSchema/oval-definitions-5#macos plist511_object,omitempty"`
	// MacosGatekeeper_object from namespace http://oval.mitre.org/XMLSchema/oval-definitions-5#macos
	MacosGatekeeper_object []oval_definitions_5_macos.Gatekeeper_objectElement `xml:"http://oval.mitre.org/XMLSchema/oval-definitions-5#macos gatekeeper_object,omitempty"`
	// MacosPwpolicy59_object from namespace http://oval.mitre.org/XMLSchema/oval-definitions-5#macos
	MacosPwpolicy59_object []oval_definitions_5_macos.Pwpolicy59_objectElement `xml:"http://oval.mitre.org/XMLSchema/oval-definitions-5#macos pwpolicy59_object,omitempty"`
	// AixInterim_fix_object from namespace http://oval.mitre.org/XMLSchema/oval-definitions-5#aix
	AixInterim_fix_object []oval_definitions_5_aix.Interim_fix_objectElement `xml:"http://oval.mitre.org/XMLSchema/oval-definitions-5#aix interim_fix_object,omitempty"`
	// AixFileset_object from namespace http://oval.mitre.org/XMLSchema/oval-definitions-5#aix
	AixFileset_object []oval_definitions_5_aix.Fileset_objectElement `xml:"http://oval.mitre.org/XMLSchema/oval-definitions-5#aix fileset_object,omitempty"`
	// AixFix_object from namespace http://oval.mitre.org/XMLSchema/oval-definitions-5#aix
	AixFix_object []oval_definitions_5_aix.Fix_objectElement `xml:"http://oval.mitre.org/XMLSchema/oval-definitions-5#aix fix_object,omitempty"`
	// AixNo_object from namespace http://oval.mitre.org/XMLSchema/oval-definitions-5#aix
	AixNo_object []oval_definitions_5_aix.No_objectElement `xml:"http://oval.mitre.org/XMLSchema/oval-definitions-5#aix no_object,omitempty"`
	// AixOslevel_object from namespace http://oval.mitre.org/XMLSchema/oval-definitions-5#aix
	AixOslevel_object []oval_definitions_5_aix.Oslevel_objectElement `xml:"http://oval.mitre.org/XMLSchema/oval-definitions-5#aix oslevel_object,omitempty"`
	// EsxPatch56_object from namespace http://oval.mitre.org/XMLSchema/oval-definitions-5#esx
	EsxPatch56_object []oval_definitions_5_esx.Patch56_objectElement `xml:"http://oval.mitre.org/XMLSchema/oval-definitions-5#esx patch56_object,omitempty"`
	// EsxPatch_object from namespace http://oval.mitre.org/XMLSchema/oval-definitions-5#esx
	EsxPatch_object []oval_definitions_5_esx.Patch_objectElement `xml:"http://oval.mitre.org/XMLSchema/oval-definitions-5#esx patch_object,omitempty"`
	// EsxVersion_object from namespace http://oval.mitre.org/XMLSchema/oval-definitions-5#esx
	EsxVersion_object []oval_definitions_5_esx.Version_objectElement `xml:"http://oval.mitre.org/XMLSchema/oval-definitions-5#esx version_object,omitempty"`
	// EsxVisdkmanagedobject_object from namespace http://oval.mitre.org/XMLSchema/oval-definitions-5#esx
	EsxVisdkmanagedobject_object []oval_definitions_5_esx.Visdkmanagedobject_objectElement `xml:"http://oval.mitre.org/XMLSchema/oval-definitions-5#esx visdkmanagedobject_object,omitempty"`
	// UnixProcess_object from namespace http://oval.mitre.org/XMLSchema/oval-definitions-5#unix
	UnixProcess_object []oval_definitions_5_unix.Process_objectElement `xml:"http://oval.mitre.org/XMLSchema/oval-definitions-5#unix process_object,omitempty"`
	// UnixSccs_object from namespace http://oval.mitre.org/XMLSchema/oval-definitions-5#unix
	UnixSccs_object []oval_definitions_5_unix.Sccs_objectElement `xml:"http://oval.mitre.org/XMLSchema/oval-definitions-5#unix sccs_object,omitempty"`
	// UnixUname_object from namespace http://oval.mitre.org/XMLSchema/oval-definitions-5#unix
	UnixUname_object []oval_definitions_5_unix.Uname_objectElement `xml:"http://oval.mitre.org/XMLSchema/oval-definitions-5#unix uname_object,omitempty"`
	// UnixRunlevel_object from namespace http://oval.mitre.org/XMLSchema/oval-definitions-5#unix
	UnixRunlevel_object []oval_definitions_5_unix.Runlevel_objectElement `xml:"http://oval.mitre.org/XMLSchema/oval-definitions-5#unix runlevel_object,omitempty"`
	// UnixXinetd_object from namespace http://oval.mitre.org/XMLSchema/oval-definitions-5#unix
	UnixXinetd_object []oval_definitions_5_unix.Xinetd_objectElement `xml:"http://oval.mitre.org/XMLSchema/oval-definitions-5#unix xinetd_object,omitempty"`
	// UnixProcess58_object from namespace http://oval.mitre.org/XMLSchema/oval-definitions-5#unix
	UnixProcess58_object []oval_definitions_5_unix.Process58_objectElement `xml:"http://oval.mitre.org/XMLSchema/oval-definitions-5#unix process58_object,omitempty"`
	// UnixInetd_object from namespace http://oval.mitre.org/XMLSchema/oval-definitions-5#unix
	UnixInetd_object []oval_definitions_5_unix.Inetd_objectElement `xml:"http://oval.mitre.org/XMLSchema/oval-definitions-5#unix inetd_object,omitempty"`
	// UnixFile_object from namespace http://oval.mitre.org/XMLSchema/oval-definitions-5#unix
	UnixFile_object []oval_definitions_5_unix.File_objectElement `xml:"http://oval.mitre.org/XMLSchema/oval-definitions-5#unix file_object,omitempty"`
	// UnixDnscache_object from namespace http://oval.mitre.org/XMLSchema/oval-definitions-5#unix
	UnixDnscache_object []oval_definitions_5_unix.Dnscache_objectElement `xml:"http://oval.mitre.org/XMLSchema/oval-definitions-5#unix dnscache_object,omitempty"`
	// UnixSymlink_object from namespace http://oval.mitre.org/XMLSchema/oval-definitions-5#unix
	UnixSymlink_object []oval_definitions_5_unix.Symlink_objectElement `xml:"http://oval.mitre.org/XMLSchema/oval-definitions-5#unix symlink_object,omitempty"`
	// UnixGconf_object from namespace http://oval.mitre.org/XMLSchema/oval-definitions-5#unix
	UnixGconf_object []oval_definitions_5_unix.Gconf_objectElement `xml:"http://oval.mitre.org/XMLSchema/oval-definitions-5#unix gconf_object,omitempty"`
	// UnixPassword_object from namespace http://oval.mitre.org/XMLSchema/oval-definitions-5#unix
	UnixPassword_object []oval_definitions_5_unix.Password_objectElement `xml:"http://oval.mitre.org/XMLSchema/oval-definitions-5#unix password_object,omitempty"`
	// UnixShadow_object from namespace http://oval.mitre.org/XMLSchema/oval-definitions-5#unix
	UnixShadow_object []oval_definitions_5_unix.Shadow_objectElement `xml:"http://oval.mitre.org/XMLSchema/oval-definitions-5#unix shadow_object,omitempty"`
	// UnixInterface_object from namespace http://oval.mitre.org/XMLSchema/oval-definitions-5#unix
	UnixInterface_object []oval_definitions_5_unix.Interface_objectElement `xml:"http://oval.mitre.org/XMLSchema/oval-definitions-5#unix interface_object,omitempty"`
	// UnixRoutingtable_object from namespace http://oval.mitre.org/XMLSchema/oval-definitions-5#unix
	UnixRoutingtable_object []oval_definitions_5_unix.Routingtable_objectElement `xml:"http://oval.mitre.org/XMLSchema/oval-definitions-5#unix routingtable_object,omitempty"`
	// UnixSysctl_object from namespace http://oval.mitre.org/XMLSchema/oval-definitions-5#unix
	UnixSysctl_object []oval_definitions_5_unix.Sysctl_objectElement `xml:"http://oval.mitre.org/XMLSchema/oval-definitions-5#unix sysctl_object,omitempty"`
	// UnixFileextendedattribute_object from namespace http://oval.mitre.org/XMLSchema/oval-definitions-5#unix
	UnixFileextendedattribute_object []oval_definitions_5_unix.Fileextendedattribute_objectElement `xml:"http://oval.mitre.org/XMLSchema/oval-definitions-5#unix fileextendedattribute_object,omitempty"`
	// FreebsdPortinfo_object from namespace http://oval.mitre.org/XMLSchema/oval-definitions-5#freebsd
	FreebsdPortinfo_object []oval_definitions_5_freebsd.Portinfo_objectElement `xml:"http://oval.mitre.org/XMLSchema/oval-definitions-5#freebsd portinfo_object,omitempty"`
	// IosVersion_object from namespace http://oval.mitre.org/XMLSchema/oval-definitions-5#ios
	IosVersion_object []oval_definitions_5_ios.Version_objectElement `xml:"http://oval.mitre.org/XMLSchema/oval-definitions-5#ios version_object,omitempty"`
	// IosAcl_object from namespace http://oval.mitre.org/XMLSchema/oval-definitions-5#ios
	IosAcl_object []oval_definitions_5_ios.Acl_objectElement `xml:"http://oval.mitre.org/XMLSchema/oval-definitions-5#ios acl_object,omitempty"`
	// IosSnmphost_object from namespace http://oval.mitre.org/XMLSchema/oval-definitions-5#ios
	IosSnmphost_object []oval_definitions_5_ios.Snmphost_objectElement `xml:"http://oval.mitre.org/XMLSchema/oval-definitions-5#ios snmphost_object,omitempty"`
	// IosGlobal_object from namespace http://oval.mitre.org/XMLSchema/oval-definitions-5#ios
	IosGlobal_object []oval_definitions_5_ios.Global_objectElement `xml:"http://oval.mitre.org/XMLSchema/oval-definitions-5#ios global_object,omitempty"`
	// IosInterface_object from namespace http://oval.mitre.org/XMLSchema/oval-definitions-5#ios
	IosInterface_object []oval_definitions_5_ios.Interface_objectElement `xml:"http://oval.mitre.org/XMLSchema/oval-definitions-5#ios interface_object,omitempty"`
	// IosSnmpcommunity_object from namespace http://oval.mitre.org/XMLSchema/oval-definitions-5#ios
	IosSnmpcommunity_object []oval_definitions_5_ios.Snmpcommunity_objectElement `xml:"http://oval.mitre.org/XMLSchema/oval-definitions-5#ios snmpcommunity_object,omitempty"`
	// IosSnmpview_object from namespace http://oval.mitre.org/XMLSchema/oval-definitions-5#ios
	IosSnmpview_object []oval_definitions_5_ios.Snmpview_objectElement `xml:"http://oval.mitre.org/XMLSchema/oval-definitions-5#ios snmpview_object,omitempty"`
	// IosLine_object from namespace http://oval.mitre.org/XMLSchema/oval-definitions-5#ios
	IosLine_object []oval_definitions_5_ios.Line_objectElement `xml:"http://oval.mitre.org/XMLSchema/oval-definitions-5#ios line_object,omitempty"`
	// IosRouter_object from namespace http://oval.mitre.org/XMLSchema/oval-definitions-5#ios
	IosRouter_object []oval_definitions_5_ios.Router_objectElement `xml:"http://oval.mitre.org/XMLSchema/oval-definitions-5#ios router_object,omitempty"`
	// IosRoutingprotocolauthintf_object from namespace http://oval.mitre.org/XMLSchema/oval-definitions-5#ios
	IosRoutingprotocolauthintf_object []oval_definitions_5_ios.Routingprotocolauthintf_objectElement `xml:"http://oval.mitre.org/XMLSchema/oval-definitions-5#ios routingprotocolauthintf_object,omitempty"`
	// IosSection_object from namespace http://oval.mitre.org/XMLSchema/oval-definitions-5#ios
	IosSection_object []oval_definitions_5_ios.Section_objectElement `xml:"http://oval.mitre.org/XMLSchema/oval-definitions-5#ios section_object,omitempty"`
	// IosBgpneighbor_object from namespace http://oval.mitre.org/XMLSchema/oval-definitions-5#ios
	IosBgpneighbor_object []oval_definitions_5_ios.Bgpneighbor_objectElement `xml:"http://oval.mitre.org/XMLSchema/oval-definitions-5#ios bgpneighbor_object,omitempty"`
	// IosSnmp_object from namespace http://oval.mitre.org/XMLSchema/oval-definitions-5#ios
	IosSnmp_object []oval_definitions_5_ios.Snmp_objectElement `xml:"http://oval.mitre.org/XMLSchema/oval-definitions-5#ios snmp_object,omitempty"`
	// IosSnmpuser_object from namespace http://oval.mitre.org/XMLSchema/oval-definitions-5#ios
	IosSnmpuser_object []oval_definitions_5_ios.Snmpuser_objectElement `xml:"http://oval.mitre.org/XMLSchema/oval-definitions-5#ios snmpuser_object,omitempty"`
	// IosVersion55_object from namespace http://oval.mitre.org/XMLSchema/oval-definitions-5#ios
	IosVersion55_object []oval_definitions_5_ios.Version55_objectElement `xml:"http://oval.mitre.org/XMLSchema/oval-definitions-5#ios version55_object,omitempty"`
	// IosTclsh_object from namespace http://oval.mitre.org/XMLSchema/oval-definitions-5#ios
	IosTclsh_object []oval_definitions_5_ios.Tclsh_objectElement `xml:"http://oval.mitre.org/XMLSchema/oval-definitions-5#ios tclsh_object,omitempty"`
	// IosSnmpgroup_object from namespace http://oval.mitre.org/XMLSchema/oval-definitions-5#ios
	IosSnmpgroup_object []oval_definitions_5_ios.Snmpgroup_objectElement `xml:"http://oval.mitre.org/XMLSchema/oval-definitions-5#ios snmpgroup_object,omitempty"`
	// ApacheHttpd_object from namespace http://oval.mitre.org/XMLSchema/oval-definitions-5#apache
	ApacheHttpd_object []oval_definitions_5_apache.Httpd_objectElement `xml:"http://oval.mitre.org/XMLSchema/oval-definitions-5#apache httpd_object,omitempty"`
	// SharepointInfopolicycoll_object from namespace http://oval.mitre.org/XMLSchema/oval-definitions-5#sharepoint
	SharepointInfopolicycoll_object []oval_definitions_5_sharepoint.Infopolicycoll_objectElement `xml:"http://oval.mitre.org/XMLSchema/oval-definitions-5#sharepoint infopolicycoll_object,omitempty"`
	// SharepointSpsite_object from namespace http://oval.mitre.org/XMLSchema/oval-definitions-5#sharepoint
	SharepointSpsite_object []oval_definitions_5_sharepoint.Spsite_objectElement `xml:"http://oval.mitre.org/XMLSchema/oval-definitions-5#sharepoint spsite_object,omitempty"`
	// SharepointSppolicyfeature_object from namespace http://oval.mitre.org/XMLSchema/oval-definitions-5#sharepoint
	SharepointSppolicyfeature_object []oval_definitions_5_sharepoint.Sppolicyfeature_objectElement `xml:"http://oval.mitre.org/XMLSchema/oval-definitions-5#sharepoint sppolicyfeature_object,omitempty"`
	// SharepointSppolicy_object from namespace http://oval.mitre.org/XMLSchema/oval-definitions-5#sharepoint
	SharepointSppolicy_object []oval_definitions_5_sharepoint.Sppolicy_objectElement `xml:"http://oval.mitre.org/XMLSchema/oval-definitions-5#sharepoint sppolicy_object,omitempty"`
	// SharepointSpjobdefinition_object from namespace http://oval.mitre.org/XMLSchema/oval-definitions-5#sharepoint
	SharepointSpjobdefinition_object []oval_definitions_5_sharepoint.Spjobdefinition_objectElement `xml:"http://oval.mitre.org/XMLSchema/oval-definitions-5#sharepoint spjobdefinition_object,omitempty"`
	// SharepointSpdiagnosticsservice_object from namespace http://oval.mitre.org/XMLSchema/oval-definitions-5#sharepoint
	SharepointSpdiagnosticsservice_object []oval_definitions_5_sharepoint.Spdiagnosticsservice_objectElement `xml:"http://oval.mitre.org/XMLSchema/oval-definitions-5#sharepoint spdiagnosticsservice_object,omitempty"`
	// SharepointSpdiagnosticslevel_object from namespace http://oval.mitre.org/XMLSchema/oval-definitions-5#sharepoint
	SharepointSpdiagnosticslevel_object []oval_definitions_5_sharepoint.Spdiagnosticslevel_objectElement `xml:"http://oval.mitre.org/XMLSchema/oval-definitions-5#sharepoint spdiagnosticslevel_object,omitempty"`
	// SharepointSplist_object from namespace http://oval.mitre.org/XMLSchema/oval-definitions-5#sharepoint
	SharepointSplist_object []oval_definitions_5_sharepoint.Splist_objectElement `xml:"http://oval.mitre.org/XMLSchema/oval-definitions-5#sharepoint splist_object,omitempty"`
	// SharepointSpweb_object from namespace http://oval.mitre.org/XMLSchema/oval-definitions-5#sharepoint
	SharepointSpweb_object []oval_definitions_5_sharepoint.Spweb_objectElement `xml:"http://oval.mitre.org/XMLSchema/oval-definitions-5#sharepoint spweb_object,omitempty"`
	// SharepointSpcrawlrule_object from namespace http://oval.mitre.org/XMLSchema/oval-definitions-5#sharepoint
	SharepointSpcrawlrule_object []oval_definitions_5_sharepoint.Spcrawlrule_objectElement `xml:"http://oval.mitre.org/XMLSchema/oval-definitions-5#sharepoint spcrawlrule_object,omitempty"`
	// SharepointSpwebapplication_object from namespace http://oval.mitre.org/XMLSchema/oval-definitions-5#sharepoint
	SharepointSpwebapplication_object []oval_definitions_5_sharepoint.Spwebapplication_objectElement `xml:"http://oval.mitre.org/XMLSchema/oval-definitions-5#sharepoint spwebapplication_object,omitempty"`
	// SharepointSpantivirussettings_object from namespace http://oval.mitre.org/XMLSchema/oval-definitions-5#sharepoint
	SharepointSpantivirussettings_object []oval_definitions_5_sharepoint.Spantivirussettings_objectElement `xml:"http://oval.mitre.org/XMLSchema/oval-definitions-5#sharepoint spantivirussettings_object,omitempty"`
	// SharepointSpsiteadministration_object from namespace http://oval.mitre.org/XMLSchema/oval-definitions-5#sharepoint
	SharepointSpsiteadministration_object []oval_definitions_5_sharepoint.Spsiteadministration_objectElement `xml:"http://oval.mitre.org/XMLSchema/oval-definitions-5#sharepoint spsiteadministration_object,omitempty"`
	// SharepointSpjobdefinition510_object from namespace http://oval.mitre.org/XMLSchema/oval-definitions-5#sharepoint
	SharepointSpjobdefinition510_object []oval_definitions_5_sharepoint.Spjobdefinition510_objectElement `xml:"http://oval.mitre.org/XMLSchema/oval-definitions-5#sharepoint spjobdefinition510_object,omitempty"`
	// SharepointSpgroup_object from namespace http://oval.mitre.org/XMLSchema/oval-definitions-5#sharepoint
	SharepointSpgroup_object []oval_definitions_5_sharepoint.Spgroup_objectElement `xml:"http://oval.mitre.org/XMLSchema/oval-definitions-5#sharepoint spgroup_object,omitempty"`
	// SharepointBestbet_object from namespace http://oval.mitre.org/XMLSchema/oval-definitions-5#sharepoint
	SharepointBestbet_object []oval_definitions_5_sharepoint.Bestbet_objectElement `xml:"http://oval.mitre.org/XMLSchema/oval-definitions-5#sharepoint bestbet_object,omitempty"`
	// CatosLine_object from namespace http://oval.mitre.org/XMLSchema/oval-definitions-5#catos
	CatosLine_object []oval_definitions_5_catos.Line_objectElement `xml:"http://oval.mitre.org/XMLSchema/oval-definitions-5#catos line_object,omitempty"`
	// CatosModule_object from namespace http://oval.mitre.org/XMLSchema/oval-definitions-5#catos
	CatosModule_object []oval_definitions_5_catos.Module_objectElement `xml:"http://oval.mitre.org/XMLSchema/oval-definitions-5#catos module_object,omitempty"`
	// CatosVersion55_object from namespace http://oval.mitre.org/XMLSchema/oval-definitions-5#catos
	CatosVersion55_object []oval_definitions_5_catos.Version55_objectElement `xml:"http://oval.mitre.org/XMLSchema/oval-definitions-5#catos version55_object,omitempty"`
	// CatosVersion_object from namespace http://oval.mitre.org/XMLSchema/oval-definitions-5#catos
	CatosVersion_object []oval_definitions_5_catos.Version_objectElement `xml:"http://oval.mitre.org/XMLSchema/oval-definitions-5#catos version_object,omitempty"`
	// AndroidLocationservice_object from namespace http://oval.mitre.org/XMLSchema/oval-definitions-5#android
	AndroidLocationservice_object []oval_definitions_5_android.Locationservice_objectElement `xml:"http://oval.mitre.org/XMLSchema/oval-definitions-5#android locationservice_object,omitempty"`
	// AndroidWifinetwork_object from namespace http://oval.mitre.org/XMLSchema/oval-definitions-5#android
	AndroidWifinetwork_object []oval_definitions_5_android.Wifinetwork_objectElement `xml:"http://oval.mitre.org/XMLSchema/oval-definitions-5#android wifinetwork_object,omitempty"`
	// AndroidDevicesettings_object from namespace http://oval.mitre.org/XMLSchema/oval-definitions-5#android
	AndroidDevicesettings_object []oval_definitions_5_android.Devicesettings_objectElement `xml:"http://oval.mitre.org/XMLSchema/oval-definitions-5#android devicesettings_object,omitempty"`
	// AndroidAppmanager_object from namespace http://oval.mitre.org/XMLSchema/oval-definitions-5#android
	AndroidAppmanager_object []oval_definitions_5_android.Appmanager_objectElement `xml:"http://oval.mitre.org/XMLSchema/oval-definitions-5#android appmanager_object,omitempty"`
	// AndroidNetwork_object from namespace http://oval.mitre.org/XMLSchema/oval-definitions-5#android
	AndroidNetwork_object []oval_definitions_5_android.Network_objectElement `xml:"http://oval.mitre.org/XMLSchema/oval-definitions-5#android network_object,omitempty"`
	// AndroidPassword_object from namespace http://oval.mitre.org/XMLSchema/oval-definitions-5#android
	AndroidPassword_object []oval_definitions_5_android.Password_objectElement `xml:"http://oval.mitre.org/XMLSchema/oval-definitions-5#android password_object,omitempty"`
	// AndroidCertificate_object from namespace http://oval.mitre.org/XMLSchema/oval-definitions-5#android
	AndroidCertificate_object []oval_definitions_5_android.Certificate_objectElement `xml:"http://oval.mitre.org/XMLSchema/oval-definitions-5#android certificate_object,omitempty"`
	// AndroidCamera_object from namespace http://oval.mitre.org/XMLSchema/oval-definitions-5#android
	AndroidCamera_object []oval_definitions_5_android.Camera_objectElement `xml:"http://oval.mitre.org/XMLSchema/oval-definitions-5#android camera_object,omitempty"`
	// AndroidSystemdetails_object from namespace http://oval.mitre.org/XMLSchema/oval-definitions-5#android
	AndroidSystemdetails_object []oval_definitions_5_android.Systemdetails_objectElement `xml:"http://oval.mitre.org/XMLSchema/oval-definitions-5#android systemdetails_object,omitempty"`
	// AndroidEncryption_object from namespace http://oval.mitre.org/XMLSchema/oval-definitions-5#android
	AndroidEncryption_object []oval_definitions_5_android.Encryption_objectElement `xml:"http://oval.mitre.org/XMLSchema/oval-definitions-5#android encryption_object,omitempty"`
	// AndroidWifi_object from namespace http://oval.mitre.org/XMLSchema/oval-definitions-5#android
	AndroidWifi_object []oval_definitions_5_android.Wifi_objectElement `xml:"http://oval.mitre.org/XMLSchema/oval-definitions-5#android wifi_object,omitempty"`
	// AndroidTelephony_object from namespace http://oval.mitre.org/XMLSchema/oval-definitions-5#android
	AndroidTelephony_object []oval_definitions_5_android.Telephony_objectElement `xml:"http://oval.mitre.org/XMLSchema/oval-definitions-5#android telephony_object,omitempty"`
	// AndroidBluetooth_object from namespace http://oval.mitre.org/XMLSchema/oval-definitions-5#android
	AndroidBluetooth_object []oval_definitions_5_android.Bluetooth_objectElement `xml:"http://oval.mitre.org/XMLSchema/oval-definitions-5#android bluetooth_object,omitempty"`
	// IndependentFilehash_object from namespace http://oval.mitre.org/XMLSchema/oval-definitions-5#independent
	IndependentFilehash_object []oval_definitions_5_independent.Filehash_objectElement `xml:"http://oval.mitre.org/XMLSchema/oval-definitions-5#independent filehash_object,omitempty"`
	// IndependentTextfilecontent_object from namespace http://oval.mitre.org/XMLSchema/oval-definitions-5#independent
	IndependentTextfilecontent_object []oval_definitions_5_independent.Textfilecontent_objectElement `xml:"http://oval.mitre.org/XMLSchema/oval-definitions-5#independent textfilecontent_object,omitempty"`
	// IndependentEnvironmentvariable_object from namespace http://oval.mitre.org/XMLSchema/oval-definitions-5#independent
	IndependentEnvironmentvariable_object []oval_definitions_5_independent.Environmentvariable_objectElement `xml:"http://oval.mitre.org/XMLSchema/oval-definitions-5#independent environmentvariable_object,omitempty"`
	// IndependentEnvironmentvariable58_object from namespace http://oval.mitre.org/XMLSchema/oval-definitions-5#independent
	IndependentEnvironmentvariable58_object []oval_definitions_5_independent.Environmentvariable58_objectElement `xml:"http://oval.mitre.org/XMLSchema/oval-definitions-5#independent environmentvariable58_object,omitempty"`
	// IndependentLdap_object from namespace http://oval.mitre.org/XMLSchema/oval-definitions-5#independent
	IndependentLdap_object []oval_definitions_5_independent.Ldap_objectElement `xml:"http://oval.mitre.org/XMLSchema/oval-definitions-5#independent ldap_object,omitempty"`
	// IndependentSql_object from namespace http://oval.mitre.org/XMLSchema/oval-definitions-5#independent
	IndependentSql_object []oval_definitions_5_independent.Sql_objectElement `xml:"http://oval.mitre.org/XMLSchema/oval-definitions-5#independent sql_object,omitempty"`
	// IndependentTextfilecontent54_object from namespace http://oval.mitre.org/XMLSchema/oval-definitions-5#independent
	IndependentTextfilecontent54_object []oval_definitions_5_independent.Textfilecontent54_objectElement `xml:"http://oval.mitre.org/XMLSchema/oval-definitions-5#independent textfilecontent54_object,omitempty"`
	// IndependentVariable_object from namespace http://oval.mitre.org/XMLSchema/oval-definitions-5#independent
	IndependentVariable_object []oval_definitions_5_independent.Variable_objectElement `xml:"http://oval.mitre.org/XMLSchema/oval-definitions-5#independent variable_object,omitempty"`
	// IndependentFamily_object from namespace http://oval.mitre.org/XMLSchema/oval-definitions-5#independent
	IndependentFamily_object []oval_definitions_5_independent.Family_objectElement `xml:"http://oval.mitre.org/XMLSchema/oval-definitions-5#independent family_object,omitempty"`
	// IndependentSql57_object from namespace http://oval.mitre.org/XMLSchema/oval-definitions-5#independent
	IndependentSql57_object []oval_definitions_5_independent.Sql57_objectElement `xml:"http://oval.mitre.org/XMLSchema/oval-definitions-5#independent sql57_object,omitempty"`
	// IndependentFilehash58_object from namespace http://oval.mitre.org/XMLSchema/oval-definitions-5#independent
	IndependentFilehash58_object []oval_definitions_5_independent.Filehash58_objectElement `xml:"http://oval.mitre.org/XMLSchema/oval-definitions-5#independent filehash58_object,omitempty"`
	// IndependentLdap57_object from namespace http://oval.mitre.org/XMLSchema/oval-definitions-5#independent
	IndependentLdap57_object []oval_definitions_5_independent.Ldap57_objectElement `xml:"http://oval.mitre.org/XMLSchema/oval-definitions-5#independent ldap57_object,omitempty"`
	// IndependentXmlfilecontent_object from namespace http://oval.mitre.org/XMLSchema/oval-definitions-5#independent
	IndependentXmlfilecontent_object []oval_definitions_5_independent.Xmlfilecontent_objectElement `xml:"http://oval.mitre.org/XMLSchema/oval-definitions-5#independent xmlfilecontent_object,omitempty"`
	// SolarisPackagecheck_object from namespace http://oval.mitre.org/XMLSchema/oval-definitions-5#solaris
	SolarisPackagecheck_object []oval_definitions_5_solaris.Packagecheck_objectElement `xml:"http://oval.mitre.org/XMLSchema/oval-definitions-5#solaris packagecheck_object,omitempty"`
	// SolarisNdd_object from namespace http://oval.mitre.org/XMLSchema/oval-definitions-5#solaris
	SolarisNdd_object []oval_definitions_5_solaris.Ndd_objectElement `xml:"http://oval.mitre.org/XMLSchema/oval-definitions-5#solaris ndd_object,omitempty"`
	// SolarisPackage_object from namespace http://oval.mitre.org/XMLSchema/oval-definitions-5#solaris
	SolarisPackage_object []oval_definitions_5_solaris.Package_objectElement `xml:"http://oval.mitre.org/XMLSchema/oval-definitions-5#solaris package_object,omitempty"`
	// SolarisVirtualizationinfo_object from namespace http://oval.mitre.org/XMLSchema/oval-definitions-5#solaris
	SolarisVirtualizationinfo_object []oval_definitions_5_solaris.Virtualizationinfo_objectElement `xml:"http://oval.mitre.org/XMLSchema/oval-definitions-5#solaris virtualizationinfo_object,omitempty"`
	// SolarisFacet_object from namespace http://oval.mitre.org/XMLSchema/oval-definitions-5#solaris
	SolarisFacet_object []oval_definitions_5_solaris.Facet_objectElement `xml:"http://oval.mitre.org/XMLSchema/oval-definitions-5#solaris facet_object,omitempty"`
	// SolarisPackagepublisher_object from namespace http://oval.mitre.org/XMLSchema/oval-definitions-5#solaris
	SolarisPackagepublisher_object []oval_definitions_5_solaris.Packagepublisher_objectElement `xml:"http://oval.mitre.org/XMLSchema/oval-definitions-5#solaris packagepublisher_object,omitempty"`
	// SolarisPackage511_object from namespace http://oval.mitre.org/XMLSchema/oval-definitions-5#solaris
	SolarisPackage511_object []oval_definitions_5_solaris.Package511_objectElement `xml:"http://oval.mitre.org/XMLSchema/oval-definitions-5#solaris package511_object,omitempty"`
	// SolarisVariant_object from namespace http://oval.mitre.org/XMLSchema/oval-definitions-5#solaris
	SolarisVariant_object []oval_definitions_5_solaris.Variant_objectElement `xml:"http://oval.mitre.org/XMLSchema/oval-definitions-5#solaris variant_object,omitempty"`
	// SolarisPackagefreezelist_object from namespace http://oval.mitre.org/XMLSchema/oval-definitions-5#solaris
	SolarisPackagefreezelist_object []oval_definitions_5_solaris.Packagefreezelist_objectElement `xml:"http://oval.mitre.org/XMLSchema/oval-definitions-5#solaris packagefreezelist_object,omitempty"`
	// SolarisPatch_object from namespace http://oval.mitre.org/XMLSchema/oval-definitions-5#solaris
	SolarisPatch_object []oval_definitions_5_solaris.Patch_objectElement `xml:"http://oval.mitre.org/XMLSchema/oval-definitions-5#solaris patch_object,omitempty"`
	// SolarisImage_object from namespace http://oval.mitre.org/XMLSchema/oval-definitions-5#solaris
	SolarisImage_object []oval_definitions_5_solaris.Image_objectElement `xml:"http://oval.mitre.org/XMLSchema/oval-definitions-5#solaris image_object,omitempty"`
	// SolarisIsainfo_object from namespace http://oval.mitre.org/XMLSchema/oval-definitions-5#solaris
	SolarisIsainfo_object []oval_definitions_5_solaris.Isainfo_objectElement `xml:"http://oval.mitre.org/XMLSchema/oval-definitions-5#solaris isainfo_object,omitempty"`
	// SolarisPackageavoidlist_object from namespace http://oval.mitre.org/XMLSchema/oval-definitions-5#solaris
	SolarisPackageavoidlist_object []oval_definitions_5_solaris.Packageavoidlist_objectElement `xml:"http://oval.mitre.org/XMLSchema/oval-definitions-5#solaris packageavoidlist_object,omitempty"`
	// SolarisPatch54_object from namespace http://oval.mitre.org/XMLSchema/oval-definitions-5#solaris
	SolarisPatch54_object []oval_definitions_5_solaris.Patch54_objectElement `xml:"http://oval.mitre.org/XMLSchema/oval-definitions-5#solaris patch54_object,omitempty"`
	// SolarisSmf_object from namespace http://oval.mitre.org/XMLSchema/oval-definitions-5#solaris
	SolarisSmf_object []oval_definitions_5_solaris.Smf_objectElement `xml:"http://oval.mitre.org/XMLSchema/oval-definitions-5#solaris smf_object,omitempty"`
	// SolarisSmfproperty_object from namespace http://oval.mitre.org/XMLSchema/oval-definitions-5#solaris
	SolarisSmfproperty_object []oval_definitions_5_solaris.Smfproperty_objectElement `xml:"http://oval.mitre.org/XMLSchema/oval-definitions-5#solaris smfproperty_object,omitempty"`
	// Apple_iosPasscodepolicy_object from namespace http://oval.mitre.org/XMLSchema/oval-definitions-5#apple_ios
	Apple_iosPasscodepolicy_object []oval_definitions_5_apple_ios.Passcodepolicy_objectElement `xml:"http://oval.mitre.org/XMLSchema/oval-definitions-5#apple_ios passcodepolicy_object,omitempty"`
	// Apple_iosProfile_object from namespace http://oval.mitre.org/XMLSchema/oval-definitions-5#apple_ios
	Apple_iosProfile_object []oval_definitions_5_apple_ios.Profile_objectElement `xml:"http://oval.mitre.org/XMLSchema/oval-definitions-5#apple_ios profile_object,omitempty"`
	// Apple_iosGlobalrestrictions_object from namespace http://oval.mitre.org/XMLSchema/oval-definitions-5#apple_ios
	Apple_iosGlobalrestrictions_object []oval_definitions_5_apple_ios.Globalrestrictions_objectElement `xml:"http://oval.mitre.org/XMLSchema/oval-definitions-5#apple_ios globalrestrictions_object,omitempty"`
	// HpuxSwlist_object from namespace http://oval.mitre.org/XMLSchema/oval-definitions-5#hpux
	HpuxSwlist_object []oval_definitions_5_hpux.Swlist_objectElement `xml:"http://oval.mitre.org/XMLSchema/oval-definitions-5#hpux swlist_object,omitempty"`
	// HpuxPatch_object from namespace http://oval.mitre.org/XMLSchema/oval-definitions-5#hpux
	HpuxPatch_object []oval_definitions_5_hpux.Patch_objectElement `xml:"http://oval.mitre.org/XMLSchema/oval-definitions-5#hpux patch_object,omitempty"`
	// HpuxGetconf_object from namespace http://oval.mitre.org/XMLSchema/oval-definitions-5#hpux
	HpuxGetconf_object []oval_definitions_5_hpux.Getconf_objectElement `xml:"http://oval.mitre.org/XMLSchema/oval-definitions-5#hpux getconf_object,omitempty"`
	// HpuxNdd_object from namespace http://oval.mitre.org/XMLSchema/oval-definitions-5#hpux
	HpuxNdd_object []oval_definitions_5_hpux.Ndd_objectElement `xml:"http://oval.mitre.org/XMLSchema/oval-definitions-5#hpux ndd_object,omitempty"`
	// HpuxTrusted_object from namespace http://oval.mitre.org/XMLSchema/oval-definitions-5#hpux
	HpuxTrusted_object []oval_definitions_5_hpux.Trusted_objectElement `xml:"http://oval.mitre.org/XMLSchema/oval-definitions-5#hpux trusted_object,omitempty"`
	// HpuxPatch53_object from namespace http://oval.mitre.org/XMLSchema/oval-definitions-5#hpux
	HpuxPatch53_object []oval_definitions_5_hpux.Patch53_objectElement `xml:"http://oval.mitre.org/XMLSchema/oval-definitions-5#hpux patch53_object,omitempty"`
	// JunosShow_object from namespace http://oval.mitre.org/XMLSchema/oval-definitions-5#junos
	JunosShow_object []oval_definitions_5_junos.Show_objectElement `xml:"http://oval.mitre.org/XMLSchema/oval-definitions-5#junos show_object,omitempty"`
	// JunosVersion_object from namespace http://oval.mitre.org/XMLSchema/oval-definitions-5#junos
	JunosVersion_object []oval_definitions_5_junos.Version_objectElement `xml:"http://oval.mitre.org/XMLSchema/oval-definitions-5#junos version_object,omitempty"`
	// JunosXml_show_object from namespace http://oval.mitre.org/XMLSchema/oval-definitions-5#junos
	JunosXml_show_object []oval_definitions_5_junos.Xml_show_objectElement `xml:"http://oval.mitre.org/XMLSchema/oval-definitions-5#junos xml_show_object,omitempty"`
	// JunosXml_config_object from namespace http://oval.mitre.org/XMLSchema/oval-definitions-5#junos
	JunosXml_config_object []oval_definitions_5_junos.Xml_config_objectElement `xml:"http://oval.mitre.org/XMLSchema/oval-definitions-5#junos xml_config_object,omitempty"`
	// NetconfConfig_object from namespace http://oval.mitre.org/XMLSchema/oval-definitions-5#netconf
	NetconfConfig_object []oval_definitions_5_netconf.Config_objectElement `xml:"http://oval.mitre.org/XMLSchema/oval-definitions-5#netconf config_object,omitempty"`
	// PixosVersion_object from namespace http://oval.mitre.org/XMLSchema/oval-definitions-5#pixos
	PixosVersion_object []oval_definitions_5_pixos.Version_objectElement `xml:"http://oval.mitre.org/XMLSchema/oval-definitions-5#pixos version_object,omitempty"`
	// PixosLine_object from namespace http://oval.mitre.org/XMLSchema/oval-definitions-5#pixos
	PixosLine_object []oval_definitions_5_pixos.Line_objectElement `xml:"http://oval.mitre.org/XMLSchema/oval-definitions-5#pixos line_object,omitempty"`
	// WindowsUac_object from namespace http://oval.mitre.org/XMLSchema/oval-definitions-5#windows
	WindowsUac_object []oval_definitions_5_windows.Uac_objectElement `xml:"http://oval.mitre.org/XMLSchema/oval-definitions-5#windows uac_object,omitempty"`
	// WindowsUser_object from namespace http://oval.mitre.org/XMLSchema/oval-definitions-5#windows
	WindowsUser_object []oval_definitions_5_windows.User_objectElement `xml:"http://oval.mitre.org/XMLSchema/oval-definitions-5#windows user_object,omitempty"`
	// WindowsFile_object from namespace http://oval.mitre.org/XMLSchema/oval-definitions-5#windows
	WindowsFile_object []oval_definitions_5_windows.File_objectElement `xml:"http://oval.mitre.org/XMLSchema/oval-definitions-5#windows file_object,omitempty"`
	// WindowsActivedirectory_object from namespace http://oval.mitre.org/XMLSchema/oval-definitions-5#windows
	WindowsActivedirectory_object []oval_definitions_5_windows.Activedirectory_objectElement `xml:"http://oval.mitre.org/XMLSchema/oval-definitions-5#windows activedirectory_object,omitempty"`
	// WindowsFileauditedpermissions_object from namespace http://oval.mitre.org/XMLSchema/oval-definitions-5#windows
	WindowsFileauditedpermissions_object []oval_definitions_5_windows.Fileauditedpermissions_objectElement `xml:"http://oval.mitre.org/XMLSchema/oval-definitions-5#windows fileauditedpermissions_object,omitempty"`
	// WindowsDnscache_object from namespace http://oval.mitre.org/XMLSchema/oval-definitions-5#windows
	WindowsDnscache_object []oval_definitions_5_windows.Dnscache_objectElement `xml:"http://oval.mitre.org/XMLSchema/oval-definitions-5#windows dnscache_object,omitempty"`
	// WindowsRegkeyeffectiverights53_object from namespace http://oval.mitre.org/XMLSchema/oval-definitions-5#windows
	WindowsRegkeyeffectiverights53_object []oval_definitions_5_windows.Regkeyeffectiverights53_objectElement `xml:"http://oval.mitre.org/XMLSchema/oval-definitions-5#windows regkeyeffectiverights53_object,omitempty"`
	// WindowsFileeffectiverights_object from namespace http://oval.mitre.org/XMLSchema/oval-definitions-5#windows
	WindowsFileeffectiverights_object []oval_definitions_5_windows.Fileeffectiverights_objectElement `xml:"http://oval.mitre.org/XMLSchema/oval-definitions-5#windows fileeffectiverights_object,omitempty"`
	// WindowsServiceeffectiverights_object from namespace http://oval.mitre.org/XMLSchema/oval-definitions-5#windows
	WindowsServiceeffectiverights_object []oval_definitions_5_windows.Serviceeffectiverights_objectElement `xml:"http://oval.mitre.org/XMLSchema/oval-definitions-5#windows serviceeffectiverights_object,omitempty"`
	// WindowsSharedresourceeffectiverights_object from namespace http://oval.mitre.org/XMLSchema/oval-definitions-5#windows
	WindowsSharedresourceeffectiverights_object []oval_definitions_5_windows.Sharedresourceeffectiverights_objectElement `xml:"http://oval.mitre.org/XMLSchema/oval-definitions-5#windows sharedresourceeffectiverights_object,omitempty"`
	// WindowsJunction_object from namespace http://oval.mitre.org/XMLSchema/oval-definitions-5#windows
	WindowsJunction_object []oval_definitions_5_windows.Junction_objectElement `xml:"http://oval.mitre.org/XMLSchema/oval-definitions-5#windows junction_object,omitempty"`
	// WindowsLicense_object from namespace http://oval.mitre.org/XMLSchema/oval-definitions-5#windows
	WindowsLicense_object []oval_definitions_5_windows.License_objectElement `xml:"http://oval.mitre.org/XMLSchema/oval-definitions-5#windows license_object,omitempty"`
	// WindowsWmi57_object from namespace http://oval.mitre.org/XMLSchema/oval-definitions-5#windows
	WindowsWmi57_object []oval_definitions_5_windows.Wmi57_objectElement `xml:"http://oval.mitre.org/XMLSchema/oval-definitions-5#windows wmi57_object,omitempty"`
	// WindowsSharedresource_object from namespace http://oval.mitre.org/XMLSchema/oval-definitions-5#windows
	WindowsSharedresource_object []oval_definitions_5_windows.Sharedresource_objectElement `xml:"http://oval.mitre.org/XMLSchema/oval-definitions-5#windows sharedresource_object,omitempty"`
	// WindowsRegkeyauditedpermissions_object from namespace http://oval.mitre.org/XMLSchema/oval-definitions-5#windows
	WindowsRegkeyauditedpermissions_object []oval_definitions_5_windows.Regkeyauditedpermissions_objectElement `xml:"http://oval.mitre.org/XMLSchema/oval-definitions-5#windows regkeyauditedpermissions_object,omitempty"`
	// WindowsUser_sid_object from namespace http://oval.mitre.org/XMLSchema/oval-definitions-5#windows
	WindowsUser_sid_object []oval_definitions_5_windows.User_sid_objectElement `xml:"http://oval.mitre.org/XMLSchema/oval-definitions-5#windows user_sid_object,omitempty"`
	// WindowsInterface_object from namespace http://oval.mitre.org/XMLSchema/oval-definitions-5#windows
	WindowsInterface_object []oval_definitions_5_windows.Interface_objectElement `xml:"http://oval.mitre.org/XMLSchema/oval-definitions-5#windows interface_object,omitempty"`
	// WindowsFileeffectiverights53_object from namespace http://oval.mitre.org/XMLSchema/oval-definitions-5#windows
	WindowsFileeffectiverights53_object []oval_definitions_5_windows.Fileeffectiverights53_objectElement `xml:"http://oval.mitre.org/XMLSchema/oval-definitions-5#windows fileeffectiverights53_object,omitempty"`
	// WindowsWmi_object from namespace http://oval.mitre.org/XMLSchema/oval-definitions-5#windows
	WindowsWmi_object []oval_definitions_5_windows.Wmi_objectElement `xml:"http://oval.mitre.org/XMLSchema/oval-definitions-5#windows wmi_object,omitempty"`
	// WindowsProcess_object from namespace http://oval.mitre.org/XMLSchema/oval-definitions-5#windows
	WindowsProcess_object []oval_definitions_5_windows.Process_objectElement `xml:"http://oval.mitre.org/XMLSchema/oval-definitions-5#windows process_object,omitempty"`
	// WindowsSystemmetric_object from namespace http://oval.mitre.org/XMLSchema/oval-definitions-5#windows
	WindowsSystemmetric_object []oval_definitions_5_windows.Systemmetric_objectElement `xml:"http://oval.mitre.org/XMLSchema/oval-definitions-5#windows systemmetric_object,omitempty"`
	// WindowsUserright_object from namespace http://oval.mitre.org/XMLSchema/oval-definitions-5#windows
	WindowsUserright_object []oval_definitions_5_windows.Userright_objectElement `xml:"http://oval.mitre.org/XMLSchema/oval-definitions-5#windows userright_object,omitempty"`
	// WindowsPort_object from namespace http://oval.mitre.org/XMLSchema/oval-definitions-5#windows
	WindowsPort_object []oval_definitions_5_windows.Port_objectElement `xml:"http://oval.mitre.org/XMLSchema/oval-definitions-5#windows port_object,omitempty"`
	// WindowsRegkeyauditedpermissions53_object from namespace http://oval.mitre.org/XMLSchema/oval-definitions-5#windows
	WindowsRegkeyauditedpermissions53_object []oval_definitions_5_windows.Regkeyauditedpermissions53_objectElement `xml:"http://oval.mitre.org/XMLSchema/oval-definitions-5#windows regkeyauditedpermissions53_object,omitempty"`
	// WindowsUser_sid55_object from namespace http://oval.mitre.org/XMLSchema/oval-definitions-5#windows
	WindowsUser_sid55_object []oval_definitions_5_windows.User_sid55_objectElement `xml:"http://oval.mitre.org/XMLSchema/oval-definitions-5#windows user_sid55_object,omitempty"`
	// WindowsAuditeventpolicysubcategories_object from namespace http://oval.mitre.org/XMLSchema/oval-definitions-5#windows
	WindowsAuditeventpolicysubcategories_object []oval_definitions_5_windows.Auditeventpolicysubcategories_objectElement `xml:"http://oval.mitre.org/XMLSchema/oval-definitions-5#windows auditeventpolicysubcategories_object,omitempty"`
	// WindowsGroup_sid_object from namespace http://oval.mitre.org/XMLSchema/oval-definitions-5#windows
	WindowsGroup_sid_object []oval_definitions_5_windows.Group_sid_objectElement `xml:"http://oval.mitre.org/XMLSchema/oval-definitions-5#windows group_sid_object,omitempty"`
	// WindowsSharedresourceauditedpermissions_object from namespace http://oval.mitre.org/XMLSchema/oval-definitions-5#windows
	WindowsSharedresourceauditedpermissions_object []oval_definitions_5_windows.Sharedresourceauditedpermissions_objectElement `xml:"http://oval.mitre.org/XMLSchema/oval-definitions-5#windows sharedresourceauditedpermissions_object,omitempty"`
	// WindowsMetabase_object from namespace http://oval.mitre.org/XMLSchema/oval-definitions-5#windows
	WindowsMetabase_object []oval_definitions_5_windows.Metabase_objectElement `xml:"http://oval.mitre.org/XMLSchema/oval-definitions-5#windows metabase_object,omitempty"`
	// WindowsRegistry_object from namespace http://oval.mitre.org/XMLSchema/oval-definitions-5#windows
	WindowsRegistry_object []oval_definitions_5_windows.Registry_objectElement `xml:"http://oval.mitre.org/XMLSchema/oval-definitions-5#windows registry_object,omitempty"`
	// WindowsNtuser_object from namespace http://oval.mitre.org/XMLSchema/oval-definitions-5#windows
	WindowsNtuser_object []oval_definitions_5_windows.Ntuser_objectElement `xml:"http://oval.mitre.org/XMLSchema/oval-definitions-5#windows ntuser_object,omitempty"`
	// WindowsPrintereffectiverights_object from namespace http://oval.mitre.org/XMLSchema/oval-definitions-5#windows
	WindowsPrintereffectiverights_object []oval_definitions_5_windows.Printereffectiverights_objectElement `xml:"http://oval.mitre.org/XMLSchema/oval-definitions-5#windows printereffectiverights_object,omitempty"`
	// WindowsSid_object from namespace http://oval.mitre.org/XMLSchema/oval-definitions-5#windows
	WindowsSid_object []oval_definitions_5_windows.Sid_objectElement `xml:"http://oval.mitre.org/XMLSchema/oval-definitions-5#windows sid_object,omitempty"`
	// WindowsLockoutpolicy_object from namespace http://oval.mitre.org/XMLSchema/oval-definitions-5#windows
	WindowsLockoutpolicy_object []oval_definitions_5_windows.Lockoutpolicy_objectElement `xml:"http://oval.mitre.org/XMLSchema/oval-definitions-5#windows lockoutpolicy_object,omitempty"`
	// WindowsSid_sid_object from namespace http://oval.mitre.org/XMLSchema/oval-definitions-5#windows
	WindowsSid_sid_object []oval_definitions_5_windows.Sid_sid_objectElement `xml:"http://oval.mitre.org/XMLSchema/oval-definitions-5#windows sid_sid_object,omitempty"`
	// WindowsVolume_object from namespace http://oval.mitre.org/XMLSchema/oval-definitions-5#windows
	WindowsVolume_object []oval_definitions_5_windows.Volume_objectElement `xml:"http://oval.mitre.org/XMLSchema/oval-definitions-5#windows volume_object,omitempty"`
	// WindowsActivedirectory57_object from namespace http://oval.mitre.org/XMLSchema/oval-definitions-5#windows
	WindowsActivedirectory57_object []oval_definitions_5_windows.Activedirectory57_objectElement `xml:"http://oval.mitre.org/XMLSchema/oval-definitions-5#windows activedirectory57_object,omitempty"`
	// WindowsWuaupdatesearcher_object from namespace http://oval.mitre.org/XMLSchema/oval-definitions-5#windows
	WindowsWuaupdatesearcher_object []oval_definitions_5_windows.Wuaupdatesearcher_objectElement `xml:"http://oval.mitre.org/XMLSchema/oval-definitions-5#windows wuaupdatesearcher_object,omitempty"`
	// WindowsAccesstoken_object from namespace http://oval.mitre.org/XMLSchema/oval-definitions-5#windows
	WindowsAccesstoken_object []oval_definitions_5_windows.Accesstoken_objectElement `xml:"http://oval.mitre.org/XMLSchema/oval-definitions-5#windows accesstoken_object,omitempty"`
	// WindowsFileauditedpermissions53_object from namespace http://oval.mitre.org/XMLSchema/oval-definitions-5#windows
	WindowsFileauditedpermissions53_object []oval_definitions_5_windows.Fileauditedpermissions53_objectElement `xml:"http://oval.mitre.org/XMLSchema/oval-definitions-5#windows fileauditedpermissions53_object,omitempty"`
	// WindowsAuditeventpolicy_object from namespace http://oval.mitre.org/XMLSchema/oval-definitions-5#windows
	WindowsAuditeventpolicy_object []oval_definitions_5_windows.Auditeventpolicy_objectElement `xml:"http://oval.mitre.org/XMLSchema/oval-definitions-5#windows auditeventpolicy_object,omitempty"`
	// WindowsGroup_object from namespace http://oval.mitre.org/XMLSchema/oval-definitions-5#windows
	WindowsGroup_object []oval_definitions_5_windows.Group_objectElement `xml:"http://oval.mitre.org/XMLSchema/oval-definitions-5#windows group_object,omitempty"`
	// WindowsPasswordpolicy_object from namespace http://oval.mitre.org/XMLSchema/oval-definitions-5#windows
	WindowsPasswordpolicy_object []oval_definitions_5_windows.Passwordpolicy_objectElement `xml:"http://oval.mitre.org/XMLSchema/oval-definitions-5#windows passwordpolicy_object,omitempty"`
	// WindowsRegkeyeffectiverights_object from namespace http://oval.mitre.org/XMLSchema/oval-definitions-5#windows
	WindowsRegkeyeffectiverights_object []oval_definitions_5_windows.Regkeyeffectiverights_objectElement `xml:"http://oval.mitre.org/XMLSchema/oval-definitions-5#windows regkeyeffectiverights_object,omitempty"`
	// WindowsPeheader_object from namespace http://oval.mitre.org/XMLSchema/oval-definitions-5#windows
	WindowsPeheader_object []oval_definitions_5_windows.Peheader_objectElement `xml:"http://oval.mitre.org/XMLSchema/oval-definitions-5#windows peheader_object,omitempty"`
	// WindowsCmdlet_object from namespace http://oval.mitre.org/XMLSchema/oval-definitions-5#windows
	WindowsCmdlet_object []oval_definitions_5_windows.Cmdlet_objectElement `xml:"http://oval.mitre.org/XMLSchema/oval-definitions-5#windows cmdlet_object,omitempty"`
	// WindowsProcess58_object from namespace http://oval.mitre.org/XMLSchema/oval-definitions-5#windows
	WindowsProcess58_object []oval_definitions_5_windows.Process58_objectElement `xml:"http://oval.mitre.org/XMLSchema/oval-definitions-5#windows process58_object,omitempty"`
	// WindowsService_object from namespace http://oval.mitre.org/XMLSchema/oval-definitions-5#windows
	WindowsService_object []oval_definitions_5_windows.Service_objectElement `xml:"http://oval.mitre.org/XMLSchema/oval-definitions-5#windows service_object,omitempty"`
	// UnknownAttrs captures any attributes not defined in XSD
	UnknownAttrs []xml.Attr `xml:",any,attr,omitempty"`
}
