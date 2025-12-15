# Schema Diagrams

Auto-generated diagrams from XSD schema analysis.

## Schema Statistics

- **Tables**: 409
- **Relationships**: 145
- **Root Entities**: 291

## Entity-Relationship Diagram

Shows database tables with their columns and relationships.

```mermaid
erDiagram
    AddressDetails {
        int id PK
        string address_type "nullable"
        string current_status "nullable"
        string valid_from_date "nullable"
        string valid_to_date "nullable"
        string usage "nullable"
        string address_details_key "nullable"
        string postal_service_elements "nullable"
        string address
        int address_lines_id
        string country
    }
    AddressElementType {
        int id PK
        string type "nullable"
    }
    AddressIdentifierElementType {
        int id PK
        string identifier_type "nullable"
        string type "nullable"
    }
    AddressLatitudeDirectionElementType {
        int id PK
        string type "nullable"
    }
    AddressLatitudeElementType {
        int id PK
        string type "nullable"
    }
    AddressLine {
        int id PK
        string type "nullable"
    }
    AddressLineElementType {
        int id PK
        string type "nullable"
    }
    AddressLinesType {
        int id PK
    }
    AddressLongitudeDirectionElementType {
        int id PK
        string type "nullable"
    }
    AddressLongitudeElementType {
        int id PK
        string type "nullable"
    }
    AddresseeIndicatorElementType {
        int id PK
        string code "nullable"
    }
    AdministrativeArea {
        int id PK
        string type "nullable"
        string usage_type "nullable"
        string indicator "nullable"
        string administrative_area_name "nullable"
        string sub_administrative_area "nullable"
    }
    AdministrativeAreaElementType {
        int id PK
        string type "nullable"
        string usage_type "nullable"
        string indicator "nullable"
        string administrative_area_name "nullable"
        string sub_administrative_area "nullable"
    }
    AdministrativeAreaNameElementType {
        int id PK
        string type "nullable"
    }
    AffectedType {
        int id PK
        string family
        string platform "nullable"
        string product "nullable"
    }
    AliasElementType {
        int id PK
        string type "nullable"
        string name_type "nullable"
        string code "nullable"
    }
    ArcType {
        int id PK
    }
    ArithmeticFunctionType {
        int id PK
        string arithmetic_operation
    }
    AssetIdentificationAssetElementType {
        int id PK
        string xsd_id
    }
    RelationshipsContainerType {
        int id PK
        string relationships "nullable"
    }
    AssetsType {
        int id PK
        string asset
        int parent_id "nullable"
    }
    AssetIdentificationType {
        int id PK
        string asset_ref
        int parent_id "nullable"
    }
    AssetReportCollection {
        int id PK
        string xsd_id "nullable"
        string report_requests "nullable"
        string assets "nullable"
        string reports
        string extended_infos "nullable"
        int parent_id "nullable"
    }
    AssetReportCollectionElementType {
        int id PK
        string xsd_id "nullable"
        string report_requests "nullable"
        string assets "nullable"
        string reports
        string extended_infos "nullable"
        int parent_id "nullable"
    }
    AssetReportingFormAssetElementType {
        int id PK
        string xsd_id
    }
    AssetType {
        int id PK
        string extended_information "nullable"
    }
    AssetsElementType {
        int id PK
        string asset
    }
    BarcodeElementType {
        int id PK
        string type "nullable"
    }
    BeginFunctionType {
        int id PK
        string character
    }
    Benchmark {
        int id PK
        string xsd_id
        bool resolved "nullable"
        string style "nullable"
        string style_href "nullable"
        int version_id
        int signature_id "nullable"
    }
    VersionType {
        int id PK
        datetime time "nullable"
        string update "nullable"
    }
    BenchmarkElementType {
        int id PK
        string xsd_id
        bool resolved "nullable"
        string style "nullable"
        string style_href "nullable"
        int version_id
        int signature_id "nullable"
    }
    BenchmarkReferenceType {
        int id PK
        string href
        string xsd_id "nullable"
    }
    BirthdateElementType {
        int id PK
    }
    BuildingNameType {
        int id PK
        string type "nullable"
        string type_occurrence "nullable"
    }
    CPE2idrefType {
        int id PK
        string idref
    }
    CanonicalizationMethodType {
        int id PK
        string algorithm
    }
    CheckContentRefType {
        int id PK
        string href
        string name "nullable"
    }
    CheckContentType {
        int id PK
    }
    CheckExportType {
        int id PK
        string value_id
        string export_name
    }
    CheckImportType {
        int id PK
        string import_name
        string import_xpath "nullable"
    }
    CidrElementType {
        int id PK
        int parent_id "nullable"
    }
    CircuitNameElementType {
        int id PK
    }
    ItAssetType {
        int id PK
        int parent_id "nullable"
    }
    CircuitType {
        int id PK
        string circuit_name "nullable"
        int parent_id "nullable"
    }
    ComplexCheckType {
        int id PK
        string operator
        bool negate "nullable"
        int check_id
        int complex_check_id
    }
    ComplexValueType {
        int id PK
        string item "nullable"
    }
    ComputingDeviceType {
        int id PK
        string distinguished_name "nullable"
        string connections "nullable"
        string hostname "nullable"
        string motherboard_guid "nullable"
        int parent_id "nullable"
    }
    ConcatFunctionType {
        int id PK
    }
    ConnectionsElementType {
        int id PK
    }
    VariableType {
        int id PK
        string xsd_id
        int version
        string datatype
        string comment
        bool deprecated "nullable"
    }
    ConstantVariable {
        int id PK
        int parent_id "nullable"
    }
    ConstantVariableElementType {
        int id PK
        int parent_id "nullable"
    }
    ContentElementType {
        int id PK
        datetime data_valid_start_date "nullable"
        datetime data_valid_end_date "nullable"
    }
    ContentElementType1 {
        int id PK
    }
    CountFunctionType {
        int id PK
    }
    CountryElementType {
        int id PK
        string country_name_code "nullable"
    }
    CountryName {
        int id PK
        string type "nullable"
    }
    CountryNameCodeElementType {
        int id PK
        string scheme "nullable"
    }
    CountryNameElementType {
        int id PK
        string type "nullable"
    }
    Cpe {
        int id PK
        int parent_id "nullable"
    }
    CpeElementType {
        int id PK
        int parent_id "nullable"
    }
    CriterionType {
        int id PK
        bool applicability_check "nullable"
        string test_ref
        bool negate "nullable"
        string comment "nullable"
    }
    ExtendDefinitionType {
        int id PK
        bool applicability_check "nullable"
        string definition_ref
        bool negate "nullable"
        string comment "nullable"
    }
    CriteriaType {
        int id PK
        bool applicability_check "nullable"
        string operator "nullable"
        bool negate "nullable"
        string comment "nullable"
        int criteria_id
        int criterion_id
        int extend_definition_id
    }
    DSAKeyValueType {
        int id PK
        int g_id "nullable"
        int y_id
        int j_id "nullable"
        int p_id
        int q_id
        int seed_id
        int pgen_counter_id
    }
    DataType {
        int id PK
        int parent_id "nullable"
    }
    DatabaseType {
        int id PK
        string instance_name "nullable"
        int parent_id "nullable"
    }
    DcStatusType {
        int id PK
    }
    DefinitionType {
        int id PK
        string xsd_id
        int version
        string class
        bool deprecated "nullable"
        int metadata_id
        int criteria_id "nullable"
    }
    DefinitionsType {
        int id PK
    }
    Department {
        int id PK
        string type "nullable"
        string department_name "nullable"
        int mail_stop_id "nullable"
    }
    MailStopType {
        int id PK
        string type "nullable"
        string mail_stop_name "nullable"
        string mail_stop_number "nullable"
    }
    DepartmentElementType {
        int id PK
        string type "nullable"
        string department_name "nullable"
        int mail_stop_id "nullable"
    }
    DepartmentNameElementType {
        int id PK
        string type "nullable"
    }
    DependencyNameElementType {
        int id PK
        string dependency_type "nullable"
        int parent_id "nullable"
    }
    DependentLocalityNameElementType {
        int id PK
        string type "nullable"
    }
    DependentLocalityNumberElementType {
        int id PK
        string name_number_occurrence "nullable"
    }
    DependentLocalityType {
        int id PK
        string type "nullable"
        string usage_type "nullable"
        string connector "nullable"
        string indicator "nullable"
        string dependent_locality_name "nullable"
        string dependent_locality_number "nullable"
        int dependent_locality_id "nullable"
        int large_mail_user_id
        int postal_route_id
    }
    ThoroughfarePreDirectionType {
        int id PK
        string type "nullable"
    }
    ThoroughfareLeadingTypeType {
        int id PK
        string type "nullable"
    }
    ThoroughfareTrailingTypeType {
        int id PK
        string type "nullable"
    }
    ThoroughfarePostDirectionType {
        int id PK
        string type "nullable"
    }
    DependentThoroughfareElementType {
        int id PK
        string type "nullable"
        int thoroughfare_pre_direction_id "nullable"
        int thoroughfare_leading_type_id "nullable"
        int thoroughfare_trailing_type_id "nullable"
        int thoroughfare_post_direction_id "nullable"
    }
    DeprecatedInfoType {
        int id PK
        string version
        string reason
        string comment "nullable"
    }
    Dictionary20CheckType {
        int id PK
        string system
        string href "nullable"
    }
    SchemaVersionType {
        int id PK
        string platform "nullable"
        int parent_id "nullable"
    }
    Dictionary20GeneratorType {
        int id PK
        string product_name "nullable"
        string product_version "nullable"
        float schema_version
        datetime timestamp
    }
    ReferencesType {
        int id PK
        string reference
    }
    Dictionary20ItemType {
        int id PK
        string name
        bool deprecated "nullable"
        string deprecated_by "nullable"
        datetime deprecation_date "nullable"
        int references_id "nullable"
    }
    Dictionary20NotesType {
        int id PK
        string note
    }
    Dictionary20TextType {
        int id PK
    }
    DigestMethodType {
        int id PK
        string algorithm
    }
    DistinguishedNameElementType {
        int id PK
    }
    DocumentRootElementType {
        int id PK
    }
    ElementMapItemType {
        int id PK
        string target_namespace "nullable"
    }
    ElementMapType {
        int id PK
        int test_id
        int object_id "nullable"
        int state_id "nullable"
        int item_id "nullable"
    }
    EmailAddress {
        int id PK
    }
    EmailAddressElementType {
        int id PK
    }
    EndFunctionType {
        int id PK
        string character
    }
    EndorsementLineCodeElementType {
        int id PK
        string type "nullable"
    }
    EntityComplexBaseType {
        int id PK
    }
    EntitySimpleBaseType {
        int id PK
    }
    EntityObjectAnySimpleType {
        int id PK
        string datatype "nullable"
        int parent_id "nullable"
    }
    EntityObjectBinaryType {
        int id PK
        string datatype
        int parent_id "nullable"
    }
    EntityObjectBoolType {
        int id PK
        string datatype
        int parent_id "nullable"
    }
    EntityObjectFieldType {
        int id PK
        string name
        string entity_check "nullable"
    }
    EntityObjectFloatType {
        int id PK
        string datatype
        int parent_id "nullable"
    }
    EntityObjectIPAddressStringType {
        int id PK
        string datatype "nullable"
        int parent_id "nullable"
    }
    EntityObjectIPAddressType {
        int id PK
        string datatype
        int parent_id "nullable"
    }
    EntityObjectIntType {
        int id PK
        string datatype
        int parent_id "nullable"
    }
    EntityObjectRecordType {
        int id PK
        int parent_id "nullable"
    }
    EntityObjectStringType {
        int id PK
        string datatype "nullable"
        int parent_id "nullable"
    }
    EntityObjectVersionType {
        int id PK
        string datatype
        int parent_id "nullable"
    }
    EntityStateSimpleBaseType {
        int id PK
        string entity_check "nullable"
        string check_existence "nullable"
        int parent_id "nullable"
    }
    EntityStateAnySimpleType {
        int id PK
        string datatype "nullable"
        int parent_id "nullable"
    }
    EntityStateBinaryType {
        int id PK
        string datatype
        int parent_id "nullable"
    }
    EntityStateBoolType {
        int id PK
        string datatype
        int parent_id "nullable"
    }
    EntityStateComplexBaseType {
        int id PK
        string entity_check "nullable"
        string check_existence "nullable"
        int parent_id "nullable"
    }
    EntityStateDebianEVRStringType {
        int id PK
        string datatype
        int parent_id "nullable"
    }
    EntityStateEVRStringType {
        int id PK
        string datatype
        int parent_id "nullable"
    }
    EntityStateFieldType {
        int id PK
        string name
        string entity_check "nullable"
    }
    EntityStateFileSetRevisionType {
        int id PK
        string datatype
        int parent_id "nullable"
    }
    EntityStateFloatType {
        int id PK
        string datatype
        int parent_id "nullable"
    }
    EntityStateIOSVersionType {
        int id PK
        string datatype "nullable"
        int parent_id "nullable"
    }
    EntityStateIPAddressStringType {
        int id PK
        string datatype "nullable"
        int parent_id "nullable"
    }
    EntityStateIPAddressType {
        int id PK
        string datatype
        int parent_id "nullable"
    }
    EntityStateIntType {
        int id PK
        string datatype
        int parent_id "nullable"
    }
    EntityStateRecordType {
        int id PK
        int parent_id "nullable"
    }
    EntityStateStringType {
        int id PK
        string datatype "nullable"
        int parent_id "nullable"
    }
    EntityStateVersionType {
        int id PK
        string datatype
        int parent_id "nullable"
    }
    EscapeRegexFunctionType {
        int id PK
    }
    Extended {
        int id PK
    }
    ExtendedInfoElementType {
        int id PK
        string xsd_id
    }
    ExtendedInformationElementType {
        int id PK
    }
    ExtendedInfosElementType {
        int id PK
        string extended_info
    }
    ExternalVariable {
        int id PK
        int possible_value_id
        int possible_restriction_id
        int parent_id "nullable"
    }
    PossibleValueType {
        int id PK
        string hint
    }
    PossibleRestrictionType {
        int id PK
        string operator "nullable"
        string hint
    }
    ExternalVariableElementType {
        int id PK
        int possible_value_id
        int possible_restriction_id
        int parent_id "nullable"
    }
    FactRefType {
        int id PK
        string name
    }
    FactType {
        int id PK
        string name
        string type "nullable"
    }
    Filter {
        int id PK
        string action "nullable"
        int parent_id "nullable"
    }
    FilterElementType {
        int id PK
        string action "nullable"
        int parent_id "nullable"
    }
    FirmNameElementType {
        int id PK
        string type "nullable"
    }
    FirmType {
        int id PK
        string type "nullable"
        string firm_name "nullable"
        int mail_stop_id "nullable"
    }
    FirstNameElementType {
        int id PK
        string type "nullable"
        string name_type "nullable"
        string code "nullable"
    }
    IdrefType {
        int id PK
        string idref
    }
    SubType {
        int id PK
        string use "nullable"
        int parent_id "nullable"
    }
    HtmlTextWithSubType {
        int id PK
        bool override "nullable"
        int sub_id
    }
    FixTextType {
        int id PK
        string fixref "nullable"
        bool reboot "nullable"
        string strategy "nullable"
        string disruption "nullable"
        string complexity "nullable"
        int parent_id "nullable"
    }
    InstanceFixType {
        int id PK
        string context "nullable"
    }
    FixType {
        int id PK
        string xsd_id "nullable"
        bool reboot "nullable"
        string strategy "nullable"
        string disruption "nullable"
        string complexity "nullable"
        string system "nullable"
        string platform "nullable"
        int sub_id
        int instance_id
    }
    FormerNameElementType {
        int id PK
        string valid_from "nullable"
        string valid_to "nullable"
        int parent_id "nullable"
    }
    Fqdn {
        int id PK
    }
    FqdnElementType {
        int id PK
    }
    Function {
        int id PK
        string code "nullable"
    }
    GeneralSuffixElementType {
        int id PK
        string type "nullable"
        string code "nullable"
    }
    GenerationIdentifierElementType {
        int id PK
        string type "nullable"
        string code "nullable"
    }
    GlobToRegexFunctionType {
        int id PK
        bool glob_noescape "nullable"
    }
    SelectableItemType {
        int id PK
        bool selected "nullable"
        string weight "nullable"
        int parent_id "nullable"
    }
    GroupType {
        int id PK
        string xsd_id
        int signature_id "nullable"
        int parent_id "nullable"
    }
    HostElementType {
        int id PK
    }
    HostnameElementType {
        int id PK
        int parent_id "nullable"
    }
    HtmlTextType {
        int id PK
        bool override "nullable"
    }
    IdentType {
        int id PK
        string system
    }
    IdentityType {
        int id PK
        bool authenticated
        bool privileged
    }
    IdrefListType {
        int id PK
        string idref
    }
    InstallationIdElementType {
        int id PK
    }
    InstanceNameElementType {
        int id PK
    }
    InstanceResultType {
        int id PK
        string context "nullable"
        string parent_context "nullable"
    }
    IpAddressType {
        int id PK
        string ip_v4 "nullable"
        string ip_v6 "nullable"
    }
    IpNetRangeElementType {
        int id PK
        int ip_net_range_start_id
        int ip_net_range_end_id
    }
    IpV4ElementType {
        int id PK
        int parent_id "nullable"
    }
    IpV6ElementType {
        int id PK
        int parent_id "nullable"
    }
    JointPersonName {
        int id PK
        string joint_name_connector "nullable"
        string code "nullable"
    }
    JointPersonNameElementType {
        int id PK
        string joint_name_connector "nullable"
        string code "nullable"
    }
    KeyInfoType {
        int id PK
        string xsd_id "nullable"
    }
    KeyLineCodeElementType {
        int id PK
        string type "nullable"
    }
    KeyValueType {
        int id PK
    }
    KnownAsElementType {
        int id PK
        string valid_from "nullable"
        string valid_to "nullable"
        int parent_id "nullable"
    }
    Language20TextType {
        int id PK
    }
    LargeMailUserIdentifierElementType {
        int id PK
        string type "nullable"
        string indicator "nullable"
    }
    LargeMailUserNameElementType {
        int id PK
        string type "nullable"
        string code "nullable"
    }
    LargeMailUserType {
        int id PK
        string type "nullable"
        string large_mail_user_name "nullable"
        string large_mail_user_identifier "nullable"
    }
    LastNameElementType {
        int id PK
        string type "nullable"
        string name_type "nullable"
        string code "nullable"
    }
    LicenseElementType {
        int id PK
    }
    ListType {
        int id PK
        int generator_id "nullable"
    }
    LiteralComponentType {
        int id PK
        string datatype "nullable"
    }
    LocalVariable {
        int id PK
        int parent_id "nullable"
    }
    LocalVariableElementType {
        int id PK
        int parent_id "nullable"
    }
    LocaleElementType {
        int id PK
        int parent_id "nullable"
    }
    Locality {
        int id PK
        string type "nullable"
        string usage_type "nullable"
        string indicator "nullable"
        string locality_name "nullable"
        int dependent_locality_id "nullable"
        int large_mail_user_id
        int postal_route_id
    }
    LocalityElementType {
        int id PK
        string type "nullable"
        string usage_type "nullable"
        string indicator "nullable"
        string locality_name "nullable"
        int dependent_locality_id "nullable"
        int large_mail_user_id
        int postal_route_id
    }
    LocalityNameElementType {
        int id PK
        string type "nullable"
    }
    LocationPoint {
        int id PK
        string latitude
        string longitude
        float elevation "nullable"
        string radius "nullable"
    }
    LocationPointElementType {
        int id PK
        string latitude
        string longitude
        float elevation "nullable"
        string radius "nullable"
    }
    LocationRegion {
        int id PK
    }
    LocationRegionElementType {
        int id PK
    }
    Locations {
        int id PK
    }
    LocationsElementType {
        int id PK
    }
    LocatorType {
        int id PK
    }
    LogicalTestType {
        int id PK
        string operator
        bool negate
    }
    MacAddressElementType {
        int id PK
        int parent_id "nullable"
    }
    MailStopNameElementType {
        int id PK
        string type "nullable"
    }
    MailStopNumberElementType {
        int id PK
        string name_number_separator "nullable"
    }
    ManifestType {
        int id PK
        string xsd_id "nullable"
    }
    MiddleNameElementType {
        int id PK
        string type "nullable"
        string name_type "nullable"
        string code "nullable"
    }
    Model {
        int id PK
        string system
    }
    ModelElementType {
        int id PK
        string system
    }
    MotherboardGuidElementType {
        int id PK
    }
    NameDetailsElementType {
        int id PK
        string name_details_key "nullable"
        string addressee_indicator "nullable"
        string dependency_name "nullable"
        int parent_id "nullable"
    }
    NameLineType {
        int id PK
        string type "nullable"
        string name_type "nullable"
        string code "nullable"
    }
    NamePrefixElementType {
        int id PK
        string type "nullable"
        string name_type "nullable"
        string code "nullable"
    }
    NetworkInterfaceType {
        int id PK
        string mac_address "nullable"
        string url "nullable"
        int subnet_mask_id "nullable"
        int default_route_id "nullable"
    }
    NetworkNameElementType {
        int id PK
    }
    NetworkType {
        int id PK
        string network_name "nullable"
        string ip_net_range
        string cidr
        int parent_id "nullable"
    }
    Notes {
        int id PK
        string note "nullable"
        int parent_id "nullable"
    }
    NotesElementType {
        int id PK
        string note "nullable"
        int parent_id "nullable"
    }
    NoticeType {
        int id PK
        string xsd_id "nullable"
    }
    Ns09XmldsigObjectType {
        int id PK
        string xsd_id "nullable"
        string mime_type "nullable"
        string encoding "nullable"
    }
    Ns09XmldsigReferenceType {
        int id PK
        string xsd_id "nullable"
        string u_r_i "nullable"
        string type "nullable"
    }
    Ns09XmldsigSignatureType {
        int id PK
        string xsd_id "nullable"
    }
    OasisNamesTcCiqXNameDetails {
        int id PK
        string name_details_key "nullable"
        string addressee_indicator "nullable"
        string dependency_name "nullable"
        int parent_id "nullable"
    }
    OasisNamesTcCiqXOrganisationNameDetails {
        int id PK
        string organisation_former_name "nullable"
        string organisation_known_as "nullable"
        int parent_id "nullable"
    }
    OasisNamesTcCiqXPersonName {
        int id PK
        string former_name "nullable"
        string known_as "nullable"
        int parent_id "nullable"
    }
    ObjectComponentType {
        int id PK
        string object_ref
        string item_field
        string record_field "nullable"
    }
    ObjectRef {
        int id PK
        string ref_id "nullable"
    }
    ObjectRefElementType {
        int id PK
        string ref_id "nullable"
    }
    ObjectRefType {
        int id PK
        string object_ref
    }
    ObjectsType {
        int id PK
    }
    OrganisationFormerNameElementType {
        int id PK
        string valid_from "nullable"
        string valid_to "nullable"
        int parent_id "nullable"
    }
    OrganisationKnownAsElementType {
        int id PK
        string valid_from "nullable"
        string valid_to "nullable"
        int parent_id "nullable"
    }
    OrganisationNameDetailsElementType {
        int id PK
        string organisation_former_name "nullable"
        string organisation_known_as "nullable"
        int parent_id "nullable"
    }
    OrganisationNameElementType {
        int id PK
        string type "nullable"
        string name_type "nullable"
        string code "nullable"
    }
    OrganisationTypeElementType {
        int id PK
        string type "nullable"
        string name_type "nullable"
        string code "nullable"
    }
    OrganizationType {
        int id PK
        int parent_id "nullable"
    }
    OtherNameElementType {
        int id PK
        string type "nullable"
        string name_type "nullable"
        string code "nullable"
    }
    OvalDefinitions {
        int id PK
        int generator_id
        int definitions_id "nullable"
        int tests_id "nullable"
        int objects_id "nullable"
        int states_id "nullable"
        int variables_id "nullable"
    }
    TestsType {
        int id PK
    }
    StatesType {
        int id PK
    }
    VariablesType {
        int id PK
    }
    OvalDefinitionsElementType {
        int id PK
        int generator_id
        int definitions_id "nullable"
        int tests_id "nullable"
        int objects_id "nullable"
        int states_id "nullable"
        int variables_id "nullable"
    }
    OvalMitreOrgOvalGeneratorType {
        int id PK
        string product_name "nullable"
        string product_version "nullable"
        datetime timestamp
    }
    OvalMitreOrgOvalMessageType {
        int id PK
        string level "nullable"
    }
    OvalMitreOrgOvalMetadataType {
        int id PK
        string title
        string description
    }
    OvalMitreOrgOvalNotesType {
        int id PK
        string note "nullable"
    }
    OvalMitreOrgOvalObjectType {
        int id PK
        string xsd_id
        int version
        string comment "nullable"
        bool deprecated "nullable"
    }
    OvalMitreOrgOvalReferenceType {
        int id PK
        string source
        string ref_id
        string ref_url "nullable"
    }
    OvalMitreOrgOvalValueType {
        int id PK
    }
    OverrideType {
        int id PK
        datetime time
        string authority
        int old_result_id
        int new_result_id
        int remark_id
    }
    OverrideableCPE2idrefType {
        int id PK
        bool override "nullable"
        int parent_id "nullable"
    }
    PGPDataType {
        int id PK
        binary p_g_p_key_i_d
        binary p_g_p_key_packet "nullable"
    }
    ParamType {
        int id PK
        string name
    }
    PersonNameElementType {
        int id PK
        string former_name "nullable"
        string known_as "nullable"
        int parent_id "nullable"
    }
    PersonType {
        int id PK
        string birthdate "nullable"
        int parent_id "nullable"
    }
    PlainTextType {
        int id PK
        string xsd_id
    }
    PlatformSpecification {
        int id PK
    }
    PlatformSpecificationElementType {
        int id PK
    }
    PlatformType {
        int id PK
        string xsd_id
        int logical_test_id
    }
    PortElementType {
        int id PK
        int parent_id "nullable"
    }
    PortRangeElementType {
        int id PK
        string lower_bound
        string upper_bound
    }
    PostBox {
        int id PK
        string type "nullable"
        string indicator "nullable"
        string post_box_number
        string post_box_number_prefix "nullable"
        string post_box_number_suffix "nullable"
        string post_box_number_extension "nullable"
        int firm_id "nullable"
    }
    PostBoxElementType {
        int id PK
        string type "nullable"
        string indicator "nullable"
        string post_box_number
        string post_box_number_prefix "nullable"
        string post_box_number_suffix "nullable"
        string post_box_number_extension "nullable"
        int firm_id "nullable"
    }
    PostBoxNumberElementType {
        int id PK
    }
    PostBoxNumberExtensionElementType {
        int id PK
        string number_extension_separator "nullable"
    }
    PostBoxNumberPrefixElementType {
        int id PK
        string number_prefix_separator "nullable"
    }
    PostBoxNumberSuffixElementType {
        int id PK
        string number_suffix_separator "nullable"
    }
    PostOffice {
        int id PK
        string type "nullable"
        string indicator "nullable"
        int postal_route_id "nullable"
        string post_office_name "nullable"
        string post_office_number "nullable"
    }
    PostalRouteType {
        int id PK
        string type "nullable"
        string postal_route_name
        string postal_route_number
    }
    PostOfficeElementType {
        int id PK
        string type "nullable"
        string indicator "nullable"
        int postal_route_id "nullable"
        string post_office_name "nullable"
        string post_office_number "nullable"
    }
    PostOfficeNameElementType {
        int id PK
        string type "nullable"
    }
    PostOfficeNumberElementType {
        int id PK
        string indicator "nullable"
        string indicator_occurrence "nullable"
    }
    PostTownElementType {
        int id PK
        string type "nullable"
        string post_town_name "nullable"
        string post_town_suffix "nullable"
    }
    PostTownNameElementType {
        int id PK
        string type "nullable"
    }
    PostTownSuffixElementType {
        int id PK
    }
    PostalCode {
        int id PK
        string type "nullable"
        string postal_code_number "nullable"
        string postal_code_number_extension "nullable"
        string post_town "nullable"
    }
    PostalCodeElementType {
        int id PK
        string type "nullable"
        string postal_code_number "nullable"
        string postal_code_number_extension "nullable"
        string post_town "nullable"
    }
    PostalCodeNumberElementType {
        int id PK
        string type "nullable"
    }
    PostalCodeNumberExtensionElementType {
        int id PK
        string type "nullable"
        string number_extension_separator "nullable"
    }
    PostalRouteNameElementType {
        int id PK
        string type "nullable"
    }
    PostalRouteNumberElementType {
        int id PK
    }
    PostalServiceElementsElementType {
        int id PK
        string type "nullable"
        string address_identifier "nullable"
        string endorsement_line_code "nullable"
        string key_line_code "nullable"
        string barcode "nullable"
        string sorting_code "nullable"
        string address_latitude "nullable"
        string address_latitude_direction "nullable"
        string address_longitude "nullable"
        string address_longitude_direction "nullable"
        string supplementary_postal_service_data "nullable"
    }
    PrecedingTitleElementType {
        int id PK
        string type "nullable"
        string code "nullable"
    }
    Premise {
        int id PK
        string type "nullable"
        string premise_dependency "nullable"
        string premise_dependency_type "nullable"
        string premise_thoroughfare_connector "nullable"
        string premise_name "nullable"
        int mail_stop_id "nullable"
        string premise_location
        string premise_number_range
        int firm_id "nullable"
    }
    PremiseElementType {
        int id PK
        string type "nullable"
        string premise_dependency "nullable"
        string premise_dependency_type "nullable"
        string premise_thoroughfare_connector "nullable"
        string premise_name "nullable"
        int mail_stop_id "nullable"
        string premise_location
        string premise_number_range
        int firm_id "nullable"
    }
    PremiseLocationElementType {
        int id PK
    }
    PremiseNameElementType {
        int id PK
        string type "nullable"
        string type_occurrence "nullable"
    }
    PremiseNumber {
        int id PK
        string number_type "nullable"
        string type "nullable"
        string indicator "nullable"
        string indicator_occurrence "nullable"
        string number_type_occurrence "nullable"
    }
    PremiseNumberElementType {
        int id PK
        string number_type "nullable"
        string type "nullable"
        string indicator "nullable"
        string indicator_occurrence "nullable"
        string number_type_occurrence "nullable"
    }
    PremiseNumberPrefix {
        int id PK
        string number_prefix_separator "nullable"
        string type "nullable"
    }
    PremiseNumberPrefixElementType {
        int id PK
        string number_prefix_separator "nullable"
        string type "nullable"
    }
    PremiseNumberRangeElementType {
        int id PK
        string range_type "nullable"
        string indicator "nullable"
        string separator "nullable"
        string type "nullable"
        string indicator_occurence "nullable"
        string number_range_occurence "nullable"
        string premise_number_range_from
        string premise_number_range_to
    }
    PremiseNumberRangeFromElementType {
        int id PK
    }
    PremiseNumberRangeToElementType {
        int id PK
    }
    PremiseNumberSuffix {
        int id PK
        string number_suffix_separator "nullable"
        string type "nullable"
    }
    PremiseNumberSuffixElementType {
        int id PK
        string number_suffix_separator "nullable"
        string type "nullable"
    }
    ProfileNoteType {
        int id PK
        string tag
        int sub_id
    }
    ProfileRefineRuleType {
        int id PK
        string idref
        string weight "nullable"
        string selector "nullable"
        string severity "nullable"
        string role "nullable"
    }
    ProfileRefineValueType {
        int id PK
        string idref
        string selector "nullable"
        string operator "nullable"
    }
    ProfileSelectType {
        int id PK
        string idref
        bool selected
    }
    ProfileSetComplexValueType {
        int id PK
        string idref
        int parent_id "nullable"
    }
    ProfileSetValueType {
        int id PK
        string idref
    }
    ProfileType {
        int id PK
        string xsd_id
        bool prohibit_changes "nullable"
        bool abstract "nullable"
        string note_tag "nullable"
        string extends "nullable"
        int version_id "nullable"
        int signature_id "nullable"
        int select_id "nullable"
        int set_complex_value_id "nullable"
        int set_value_id "nullable"
        int refine_value_id "nullable"
        int refine_rule_id "nullable"
    }
    ProtocolElementType {
        int id PK
    }
    RSAKeyValueType {
        int id PK
        int modulus_id
        int exponent_id
    }
    ReferenceElementType {
        int id PK
        string href "nullable"
    }
    RegexCaptureFunctionType {
        int id PK
        string pattern "nullable"
    }
    RelationshipType {
        int id PK
        string type
        string scope "nullable"
        string subject
        string ref
    }
    RelationshipsElementType {
        int id PK
    }
    RemoteResource {
        int id PK
    }
    RemoteResourceElementType {
        int id PK
    }
    ReportRequestType {
        int id PK
        string xsd_id
        string content
    }
    ReportRequestsElementType {
        int id PK
    }
    ReportType {
        int id PK
        string xsd_id
        string content
    }
    ReportsElementType {
        int id PK
    }
    ResourceType {
        int id PK
    }
    RestrictionType {
        int id PK
        string operation
    }
    RetrievalMethodType {
        int id PK
        string u_r_i "nullable"
        string type "nullable"
    }
    RuleResultType {
        int id PK
        string idref
        string role "nullable"
        string severity "nullable"
        datetime time "nullable"
        string version "nullable"
        string weight "nullable"
        int result_id
        int complex_check_id "nullable"
    }
    RuleType {
        int id PK
        string xsd_id
        string role "nullable"
        string severity "nullable"
        bool multiple "nullable"
        string impact_metric "nullable"
        int signature_id "nullable"
        int complex_check_id "nullable"
        int parent_id "nullable"
    }
    SPKIDataType {
        int id PK
        binary s_p_k_i_sexp
    }
    ScoreType {
        int id PK
        string system "nullable"
        float maximum "nullable"
    }
    SelChoicesType {
        int id PK
        bool must_match "nullable"
        string selector "nullable"
        string choice
        int complex_choice_id
    }
    SelComplexValueType {
        int id PK
        string selector "nullable"
        int parent_id "nullable"
    }
    SelNumType {
        int id PK
        string selector "nullable"
    }
    SelStringType {
        int id PK
        string selector "nullable"
    }
    ServiceType {
        int id PK
        string host "nullable"
        string port "nullable"
        string port_range "nullable"
        string protocol "nullable"
        int parent_id "nullable"
    }
    Set {
        int id PK
        string set_operator "nullable"
    }
    SetElementType {
        int id PK
        string set_operator "nullable"
    }
    SignatureMethodType {
        int id PK
        string algorithm
        int h_m_a_c_output_length_id "nullable"
    }
    SignaturePropertiesType {
        int id PK
        string xsd_id "nullable"
    }
    SignaturePropertyType {
        int id PK
        string target
        string xsd_id "nullable"
    }
    SignatureValueType {
        int id PK
        string xsd_id "nullable"
    }
    SignedInfoType {
        int id PK
        string xsd_id "nullable"
    }
    Simple {
        int id PK
    }
    SoftwareType {
        int id PK
        string installation_id "nullable"
        string license "nullable"
        int parent_id "nullable"
    }
    SortingCodeElementType {
        int id PK
        string type "nullable"
    }
    SplitFunctionType {
        int id PK
        string delimiter
    }
    StateRefType {
        int id PK
        string state_ref
    }
    StateType {
        int id PK
        string xsd_id
        int version
        string operator "nullable"
        string comment "nullable"
        bool deprecated "nullable"
    }
    Status {
        int id PK
        datetime date "nullable"
        int parent_id "nullable"
    }
    StatusElementType {
        int id PK
        datetime date "nullable"
        int parent_id "nullable"
    }
    SubAdministrativeAreaElementType {
        int id PK
        string type "nullable"
        string usage_type "nullable"
        string indicator "nullable"
        string sub_administrative_area_name "nullable"
    }
    SubAdministrativeAreaNameElementType {
        int id PK
        string type "nullable"
    }
    SubPremiseLocationElementType {
        int id PK
    }
    SubPremiseNameElementType {
        int id PK
        string type "nullable"
        string type_occurrence "nullable"
    }
    SubPremiseNumberElementType {
        int id PK
        string indicator "nullable"
        string indicator_occurrence "nullable"
        string number_type_occurrence "nullable"
        string premise_number_separator "nullable"
        string type "nullable"
    }
    SubPremiseNumberPrefixElementType {
        int id PK
        string number_prefix_separator "nullable"
        string type "nullable"
    }
    SubPremiseNumberSuffixElementType {
        int id PK
        string number_suffix_separator "nullable"
        string type "nullable"
    }
    SubPremiseType {
        int id PK
        string type "nullable"
        string sub_premise_name "nullable"
        string sub_premise_number_prefix "nullable"
        string sub_premise_number_suffix "nullable"
        int firm_id "nullable"
        int mail_stop_id "nullable"
        int sub_premise_id "nullable"
        string sub_premise_location
        string sub_premise_number "nullable"
    }
    SubstringFunctionType {
        int id PK
        int substring_start
        int substring_length
    }
    SuffixElementType {
        int id PK
        string type "nullable"
        string code "nullable"
    }
    SupplementaryPostalServiceDataElementType {
        int id PK
        string type "nullable"
    }
    SyntheticId {
        int id PK
        string resource
        string xsd_id
    }
    SyntheticIdElementType {
        int id PK
        string resource
        string xsd_id
    }
    SystemNameElementType {
        int id PK
    }
    SystemType {
        int id PK
        string system_name "nullable"
        string version "nullable"
        int parent_id "nullable"
    }
    TailoringBenchmarkReferenceType {
        int id PK
        string version "nullable"
        int parent_id "nullable"
    }
    TailoringReferenceType {
        int id PK
        string href
        string xsd_id
        string version
        datetime time
    }
    TailoringVersionType {
        int id PK
        datetime time
    }
    TailoringType {
        int id PK
        string xsd_id
        int benchmark_id "nullable"
        int version_id
        int signature_id "nullable"
    }
    TargetFactsType {
        int id PK
    }
    TargetIdRefType {
        int id PK
        string system
        string href
        string name "nullable"
    }
    TelephoneNumber {
        int id PK
        int parent_id "nullable"
    }
    TelephoneNumberElementType {
        int id PK
        int parent_id "nullable"
    }
    TestResultType {
        int id PK
        string xsd_id
        datetime start_time "nullable"
        datetime end_time
        string test_system "nullable"
        string version "nullable"
        int benchmark_id "nullable"
        int tailoring_file_id "nullable"
        string organization "nullable"
        int identity_id "nullable"
        int profile_id "nullable"
        string target
        string target_address "nullable"
        int target_facts_id "nullable"
        int signature_id "nullable"
        int target_id_ref_id
        int set_value_id
        int set_complex_value_id
    }
    TestType {
        int id PK
        string xsd_id
        int version
        string check_existence "nullable"
        string check
        string state_operator "nullable"
        string comment
        bool deprecated "nullable"
    }
    TextWithSubType {
        int id PK
        bool override "nullable"
    }
    Thoroughfare {
        int id PK
        string type "nullable"
        string dependent_thoroughfares "nullable"
        string dependent_thoroughfares_indicator "nullable"
        string dependent_thoroughfares_connector "nullable"
        string dependent_thoroughfares_type "nullable"
        int thoroughfare_pre_direction_id "nullable"
        int thoroughfare_leading_type_id "nullable"
        int thoroughfare_trailing_type_id "nullable"
        int thoroughfare_post_direction_id "nullable"
        string dependent_thoroughfare "nullable"
        string thoroughfare_number_range
        int dependent_locality_id
        int firm_id
    }
    ThoroughfareElementType {
        int id PK
        string type "nullable"
        string dependent_thoroughfares "nullable"
        string dependent_thoroughfares_indicator "nullable"
        string dependent_thoroughfares_connector "nullable"
        string dependent_thoroughfares_type "nullable"
        int thoroughfare_pre_direction_id "nullable"
        int thoroughfare_leading_type_id "nullable"
        int thoroughfare_trailing_type_id "nullable"
        int thoroughfare_post_direction_id "nullable"
        string dependent_thoroughfare "nullable"
        string thoroughfare_number_range
        int dependent_locality_id
        int firm_id
    }
    ThoroughfareNameType {
        int id PK
        string type "nullable"
    }
    ThoroughfareNumber {
        int id PK
        string number_type "nullable"
        string type "nullable"
        string indicator "nullable"
        string indicator_occurrence "nullable"
        string number_occurrence "nullable"
    }
    ThoroughfareNumberElementType {
        int id PK
        string number_type "nullable"
        string type "nullable"
        string indicator "nullable"
        string indicator_occurrence "nullable"
        string number_occurrence "nullable"
    }
    ThoroughfareNumberFromElementType {
        int id PK
    }
    ThoroughfareNumberPrefix {
        int id PK
        string number_prefix_separator "nullable"
        string type "nullable"
    }
    ThoroughfareNumberPrefixElementType {
        int id PK
        string number_prefix_separator "nullable"
        string type "nullable"
    }
    ThoroughfareNumberRangeElementType {
        int id PK
        string range_type "nullable"
        string indicator "nullable"
        string separator "nullable"
        string indicator_occurrence "nullable"
        string number_range_occurrence "nullable"
        string type "nullable"
        string thoroughfare_number_from
        string thoroughfare_number_to
    }
    ThoroughfareNumberSuffix {
        int id PK
        string number_suffix_separator "nullable"
        string type "nullable"
    }
    ThoroughfareNumberSuffixElementType {
        int id PK
        string number_suffix_separator "nullable"
        string type "nullable"
    }
    ThoroughfareNumberToElementType {
        int id PK
    }
    TimeDifferenceFunctionType {
        int id PK
        string format_1 "nullable"
        string format_2 "nullable"
    }
    TitleElementType {
        int id PK
        string type "nullable"
        string code "nullable"
    }
    TitleEltType {
        int id PK
    }
    TransformType {
        int id PK
        string algorithm
        string x_path
    }
    TransformsType {
        int id PK
    }
    UniqueFunctionType {
        int id PK
    }
    UriRefType {
        int id PK
        string uri
    }
    UrlElementType {
        int id PK
    }
    VariableComponentType {
        int id PK
        string var_ref
    }
    VersionElementType {
        int id PK
    }
    WarningType {
        int id PK
        string category "nullable"
        int parent_id "nullable"
    }
    WebsiteType {
        int id PK
        string document_root "nullable"
        string locale "nullable"
        int parent_id "nullable"
    }
    WebsiteUrl {
        int id PK
    }
    WebsiteUrlElementType {
        int id PK
    }
    X509DataType {
        int id PK
        int x509_issuer_serial_id
        binary x509_s_k_i
        string x509_subject_name
        binary x509_certificate
        binary x509_c_r_l
    }
    X509IssuerSerialType {
        int id PK
        string x509_issuer_name
        int x509_serial_number
    }
    XAL {
        int id PK
        string version "nullable"
    }
    XALElementType {
        int id PK
        string version "nullable"
    }
    XNL {
        int id PK
        string version "nullable"
    }
    XNLElementType {
        int id PK
        string version "nullable"
    }
    Xccdf12CheckType {
        int id PK
        string system
        bool negate "nullable"
        string xsd_id "nullable"
        string selector "nullable"
        bool multi_check "nullable"
        int check_content_id "nullable"
    }
    Xccdf12ItemType {
        int id PK
        bool abstract "nullable"
        string cluster_id "nullable"
        string extends "nullable"
        bool hidden "nullable"
        bool prohibit_changes "nullable"
        string xsd_id "nullable"
        int version_id "nullable"
    }
    Xccdf12MessageType {
        int id PK
        string severity
    }
    Xccdf12MetadataType {
        int id PK
    }
    Xccdf12ReferenceType {
        int id PK
        string href "nullable"
        bool override "nullable"
    }
    Xccdf12SignatureType {
        int id PK
    }
    Xccdf12TextType {
        int id PK
        bool override "nullable"
    }
    Xccdf12ValueType {
        int id PK
        string xsd_id
        string type "nullable"
        string operator "nullable"
        bool interactive "nullable"
        string interface_hint "nullable"
        int signature_id "nullable"
        int value_id
        int complex_value_id
        int default_id
        int complex_default_id
        int parent_id "nullable"
    }

    RelationshipsContainerType |o--o{ AssetsType : "References parent type relationships-container-type"
    RelationshipsContainerType ||--|| AssetsType : "extends"
    AssetsType |o--o{ AssetIdentificationType : "References parent type assets-type"
    AssetsType ||--|| AssetIdentificationType : "extends"
    RelationshipsContainerType |o--o{ AssetReportCollection : "References parent type relationships-container-type"
    RelationshipsContainerType ||--|| AssetReportCollection : "extends"
    RelationshipsContainerType |o--o{ AssetReportCollectionElementType : "References parent type relationships-container-type"
    RelationshipsContainerType ||--|| AssetReportCollectionElementType : "extends"
    VersionType ||--o{ BenchmarkElementType : "References versionType"
    CidrType |o--o{ CidrElementType : "References parent type cidr-type"
    CidrType ||--|| CidrElementType : "extends"
    AssetType |o--o{ ItAssetType : "References parent type asset-type"
    AssetType ||--|| ItAssetType : "extends"
    ItAssetType |o--o{ CircuitType : "References parent type it-asset-type"
    ItAssetType ||--|| CircuitType : "extends"
    ComplexCheckType ||--o{ ComplexCheckType : "References complexCheckType"
    ItAssetType |o--o{ ComputingDeviceType : "References parent type it-asset-type"
    ItAssetType ||--|| ComputingDeviceType : "extends"
    VariableType |o--o{ ConstantVariable : "References parent type VariableType"
    VariableType ||--|| ConstantVariable : "extends"
    VariableType |o--o{ ConstantVariableElementType : "References parent type VariableType"
    VariableType ||--|| ConstantVariableElementType : "extends"
    CpeType |o--o{ Cpe : "References parent type cpe-type"
    CpeType ||--|| Cpe : "extends"
    CpeType |o--o{ CpeElementType : "References parent type cpe-type"
    CpeType ||--|| CpeElementType : "extends"
    CriteriaType ||--o{ CriteriaType : "References CriteriaType"
    CriterionType ||--o{ CriteriaType : "References CriterionType"
    ExtendDefinitionType ||--o{ CriteriaType : "References ExtendDefinitionType"
    AssetType |o--o{ DataType : "References parent type asset-type"
    AssetType ||--|| DataType : "extends"
    ItAssetType |o--o{ DatabaseType : "References parent type it-asset-type"
    ItAssetType ||--|| DatabaseType : "extends"
    CriteriaType |o--o{ DefinitionType : "References CriteriaType"
    MailStopType |o--o{ DepartmentElementType : "References MailStopType"
    NameDetails |o--o{ DependencyNameElementType : "References parent type NameDetails"
    NameDetails ||--|| DependencyNameElementType : "extends"
    DependentLocalityType |o--o{ DependentLocalityType : "References DependentLocalityType"
    ThoroughfarePreDirectionType |o--o{ DependentThoroughfareElementType : "References ThoroughfarePreDirectionType"
    ThoroughfareLeadingTypeType |o--o{ DependentThoroughfareElementType : "References ThoroughfareLeadingTypeType"
    ThoroughfareTrailingTypeType |o--o{ DependentThoroughfareElementType : "References ThoroughfareTrailingTypeType"
    ThoroughfarePostDirectionType |o--o{ DependentThoroughfareElementType : "References ThoroughfarePostDirectionType"
    SchemaVersionPattern |o--o{ SchemaVersionType : "References parent type SchemaVersionPattern"
    SchemaVersionPattern ||--|| SchemaVersionType : "extends"
    SchemaVersionType ||--o{ Dictionary20GeneratorType : "References SchemaVersionType"
    ReferencesType |o--o{ Dictionary20ItemType : "References ReferencesType"
    ElementMapItemType ||--o{ ElementMapType : "References ElementMapItemType"
    ElementMapItemType |o--o{ ElementMapType : "References ElementMapItemType"
    ElementMapItemType |o--o{ ElementMapType : "References ElementMapItemType"
    ElementMapItemType |o--o{ ElementMapType : "References ElementMapItemType"
    EntitySimpleBaseType |o--o{ EntityObjectAnySimpleType : "References parent type EntitySimpleBaseType"
    EntitySimpleBaseType ||--|| EntityObjectAnySimpleType : "extends"
    EntitySimpleBaseType |o--o{ EntityObjectBinaryType : "References parent type EntitySimpleBaseType"
    EntitySimpleBaseType ||--|| EntityObjectBinaryType : "extends"
    EntitySimpleBaseType |o--o{ EntityObjectBoolType : "References parent type EntitySimpleBaseType"
    EntitySimpleBaseType ||--|| EntityObjectBoolType : "extends"
    EntitySimpleBaseType |o--o{ EntityObjectFloatType : "References parent type EntitySimpleBaseType"
    EntitySimpleBaseType ||--|| EntityObjectFloatType : "extends"
    EntitySimpleBaseType |o--o{ EntityObjectIPAddressStringType : "References parent type EntitySimpleBaseType"
    EntitySimpleBaseType ||--|| EntityObjectIPAddressStringType : "extends"
    EntitySimpleBaseType |o--o{ EntityObjectIPAddressType : "References parent type EntitySimpleBaseType"
    EntitySimpleBaseType ||--|| EntityObjectIPAddressType : "extends"
    EntitySimpleBaseType |o--o{ EntityObjectIntType : "References parent type EntitySimpleBaseType"
    EntitySimpleBaseType ||--|| EntityObjectIntType : "extends"
    EntityComplexBaseType |o--o{ EntityObjectRecordType : "References parent type EntityComplexBaseType"
    EntityComplexBaseType ||--|| EntityObjectRecordType : "extends"
    EntitySimpleBaseType |o--o{ EntityObjectStringType : "References parent type EntitySimpleBaseType"
    EntitySimpleBaseType ||--|| EntityObjectStringType : "extends"
    EntitySimpleBaseType |o--o{ EntityObjectVersionType : "References parent type EntitySimpleBaseType"
    EntitySimpleBaseType ||--|| EntityObjectVersionType : "extends"
    EntitySimpleBaseType |o--o{ EntityStateSimpleBaseType : "References parent type EntitySimpleBaseType"
    EntitySimpleBaseType ||--|| EntityStateSimpleBaseType : "extends"
    EntityStateSimpleBaseType |o--o{ EntityStateAnySimpleType : "References parent type EntityStateSimpleBaseType"
    EntityStateSimpleBaseType ||--|| EntityStateAnySimpleType : "extends"
    EntityStateSimpleBaseType |o--o{ EntityStateBinaryType : "References parent type EntityStateSimpleBaseType"
    EntityStateSimpleBaseType ||--|| EntityStateBinaryType : "extends"
    EntityStateSimpleBaseType |o--o{ EntityStateBoolType : "References parent type EntityStateSimpleBaseType"
    EntityStateSimpleBaseType ||--|| EntityStateBoolType : "extends"
    EntityComplexBaseType |o--o{ EntityStateComplexBaseType : "References parent type EntityComplexBaseType"
    EntityComplexBaseType ||--|| EntityStateComplexBaseType : "extends"
    EntityStateSimpleBaseType |o--o{ EntityStateDebianEVRStringType : "References parent type EntityStateSimpleBaseType"
    EntityStateSimpleBaseType ||--|| EntityStateDebianEVRStringType : "extends"
    EntityStateSimpleBaseType |o--o{ EntityStateEVRStringType : "References parent type EntityStateSimpleBaseType"
    EntityStateSimpleBaseType ||--|| EntityStateEVRStringType : "extends"
    EntityStateSimpleBaseType |o--o{ EntityStateFileSetRevisionType : "References parent type EntityStateSimpleBaseType"
    EntityStateSimpleBaseType ||--|| EntityStateFileSetRevisionType : "extends"
    EntityStateSimpleBaseType |o--o{ EntityStateFloatType : "References parent type EntityStateSimpleBaseType"
    EntityStateSimpleBaseType ||--|| EntityStateFloatType : "extends"
    EntityStateSimpleBaseType |o--o{ EntityStateIOSVersionType : "References parent type EntityStateSimpleBaseType"
    EntityStateSimpleBaseType ||--|| EntityStateIOSVersionType : "extends"
    EntityStateSimpleBaseType |o--o{ EntityStateIPAddressStringType : "References parent type EntityStateSimpleBaseType"
    EntityStateSimpleBaseType ||--|| EntityStateIPAddressStringType : "extends"
    EntityStateSimpleBaseType |o--o{ EntityStateIPAddressType : "References parent type EntityStateSimpleBaseType"
    EntityStateSimpleBaseType ||--|| EntityStateIPAddressType : "extends"
    EntityStateSimpleBaseType |o--o{ EntityStateIntType : "References parent type EntityStateSimpleBaseType"
    EntityStateSimpleBaseType ||--|| EntityStateIntType : "extends"
    EntityStateComplexBaseType |o--o{ EntityStateRecordType : "References parent type EntityStateComplexBaseType"
    EntityStateComplexBaseType ||--|| EntityStateRecordType : "extends"
    EntityStateSimpleBaseType |o--o{ EntityStateStringType : "References parent type EntityStateSimpleBaseType"
    EntityStateSimpleBaseType ||--|| EntityStateStringType : "extends"
    EntityStateSimpleBaseType |o--o{ EntityStateVersionType : "References parent type EntityStateSimpleBaseType"
    EntityStateSimpleBaseType ||--|| EntityStateVersionType : "extends"
    VariableType |o--o{ ExternalVariable : "References parent type VariableType"
    VariableType ||--|| ExternalVariable : "extends"
    VariableType |o--o{ ExternalVariableElementType : "References parent type VariableType"
    PossibleValueType ||--o{ ExternalVariableElementType : "References PossibleValueType"
    PossibleRestrictionType ||--o{ ExternalVariableElementType : "References PossibleRestrictionType"
    VariableType ||--|| ExternalVariableElementType : "extends"
    StateIDPattern |o--o{ Filter : "References parent type StateIDPattern"
    StateIDPattern ||--|| Filter : "extends"
    StateIDPattern |o--o{ FilterElementType : "References parent type StateIDPattern"
    StateIDPattern ||--|| FilterElementType : "extends"
    MailStopType |o--o{ FirmType : "References MailStopType"
    IdrefType |o--o{ SubType : "References parent type idrefType"
    IdrefType ||--|| SubType : "extends"
    SubType ||--o{ HtmlTextWithSubType : "References subType"
    HtmlTextWithSubType |o--o{ FixTextType : "References parent type htmlTextWithSubType"
    HtmlTextWithSubType ||--|| FixTextType : "extends"
    SubType ||--o{ FixType : "References subType"
    InstanceFixType ||--o{ FixType : "References instanceFixType"
    PersonName |o--o{ FormerNameElementType : "References parent type PersonName"
    PersonName ||--|| FormerNameElementType : "extends"
    ItemType |o--o{ SelectableItemType : "References parent type itemType"
    ItemType ||--|| SelectableItemType : "extends"
    SelectableItemType |o--o{ GroupType : "References parent type selectableItemType"
    SelectableItemType ||--|| GroupType : "extends"
    HostnameType |o--o{ HostnameElementType : "References parent type hostname-type"
    HostnameType ||--|| HostnameElementType : "extends"
    IpAddressType ||--o{ IpNetRangeElementType : "References ip-address-type"
    IpAddressType ||--o{ IpNetRangeElementType : "References ip-address-type"
    Ipv4Type |o--o{ IpV4ElementType : "References parent type ipv4-type"
    Ipv4Type ||--|| IpV4ElementType : "extends"
    Ipv6Type |o--o{ IpV6ElementType : "References parent type ipv6-type"
    Ipv6Type ||--|| IpV6ElementType : "extends"
    PersonName |o--o{ KnownAsElementType : "References parent type PersonName"
    PersonName ||--|| KnownAsElementType : "extends"
    VariableType |o--o{ LocalVariable : "References parent type VariableType"
    VariableType ||--|| LocalVariable : "extends"
    VariableType |o--o{ LocalVariableElementType : "References parent type VariableType"
    VariableType ||--|| LocalVariableElementType : "extends"
    LocaleType |o--o{ LocaleElementType : "References parent type locale-type"
    LocaleType ||--|| LocaleElementType : "extends"
    DependentLocalityType |o--o{ LocalityElementType : "References DependentLocalityType"
    MacAddressType |o--o{ MacAddressElementType : "References parent type mac-address-type"
    MacAddressType ||--|| MacAddressElementType : "extends"
    NameDetails |o--o{ NameDetailsElementType : "References parent type NameDetails"
    NameDetails ||--|| NameDetailsElementType : "extends"
    IpAddressType |o--o{ NetworkInterfaceType : "References ip-address-type"
    IpAddressType |o--o{ NetworkInterfaceType : "References ip-address-type"
    ItAssetType |o--o{ NetworkType : "References parent type it-asset-type"
    ItAssetType ||--|| NetworkType : "extends"
    NotesType |o--o{ Notes : "References parent type NotesType"
    NotesType ||--|| Notes : "extends"
    NotesType |o--o{ NotesElementType : "References parent type NotesType"
    NotesType ||--|| NotesElementType : "extends"
    NameDetails |o--o{ OasisNamesTcCiqXNameDetails : "References parent type NameDetails"
    NameDetails ||--|| OasisNamesTcCiqXNameDetails : "extends"
    OrganisationNameDetails |o--o{ OasisNamesTcCiqXOrganisationNameDetails : "References parent type OrganisationNameDetails"
    OrganisationNameDetails ||--|| OasisNamesTcCiqXOrganisationNameDetails : "extends"
    PersonName |o--o{ OasisNamesTcCiqXPersonName : "References parent type PersonName"
    PersonName ||--|| OasisNamesTcCiqXPersonName : "extends"
    OrganisationNameDetails |o--o{ OrganisationFormerNameElementType : "References parent type OrganisationNameDetails"
    OrganisationNameDetails ||--|| OrganisationFormerNameElementType : "extends"
    OrganisationNameDetails |o--o{ OrganisationKnownAsElementType : "References parent type OrganisationNameDetails"
    OrganisationNameDetails ||--|| OrganisationKnownAsElementType : "extends"
    OrganisationNameDetails |o--o{ OrganisationNameDetailsElementType : "References parent type OrganisationNameDetails"
    OrganisationNameDetails ||--|| OrganisationNameDetailsElementType : "extends"
    AssetType |o--o{ OrganizationType : "References parent type asset-type"
    AssetType ||--|| OrganizationType : "extends"
    DefinitionsType |o--o{ OvalDefinitionsElementType : "References DefinitionsType"
    TestsType |o--o{ OvalDefinitionsElementType : "References TestsType"
    ObjectsType |o--o{ OvalDefinitionsElementType : "References ObjectsType"
    StatesType |o--o{ OvalDefinitionsElementType : "References StatesType"
    VariablesType |o--o{ OvalDefinitionsElementType : "References VariablesType"
    CPE2idrefType |o--o{ OverrideableCPE2idrefType : "References parent type CPE2idrefType"
    CPE2idrefType ||--|| OverrideableCPE2idrefType : "extends"
    PersonName |o--o{ PersonNameElementType : "References parent type PersonName"
    PersonName ||--|| PersonNameElementType : "extends"
    AssetType |o--o{ PersonType : "References parent type asset-type"
    AssetType ||--|| PersonType : "extends"
    LogicalTestType ||--o{ PlatformType : "References LogicalTestType"
    PortType |o--o{ PortElementType : "References parent type port-type"
    PortType ||--|| PortElementType : "extends"
    FirmType |o--o{ PostBoxElementType : "References FirmType"
    PostalRouteType |o--o{ PostOfficeElementType : "References PostalRouteType"
    MailStopType |o--o{ PremiseElementType : "References MailStopType"
    SubType ||--o{ ProfileNoteType : "References subType"
    ComplexValueType |o--o{ ProfileSetComplexValueType : "References parent type complexValueType"
    ComplexValueType ||--|| ProfileSetComplexValueType : "extends"
    VersionType |o--o{ ProfileType : "References versionType"
    SelectableItemType |o--o{ RuleType : "References parent type selectableItemType"
    SelectableItemType ||--|| RuleType : "extends"
    ComplexValueType ||--o{ SelChoicesType : "References complexValueType"
    ComplexValueType |o--o{ SelComplexValueType : "References parent type complexValueType"
    ComplexValueType ||--|| SelComplexValueType : "extends"
    ItAssetType |o--o{ ServiceType : "References parent type it-asset-type"
    ItAssetType ||--|| ServiceType : "extends"
    ItAssetType |o--o{ SoftwareType : "References parent type it-asset-type"
    ItAssetType ||--|| SoftwareType : "extends"
    StatusType |o--o{ Status : "References parent type statusType"
    StatusType ||--|| Status : "extends"
    StatusType |o--o{ StatusElementType : "References parent type statusType"
    StatusType ||--|| StatusElementType : "extends"
    FirmType |o--o{ SubPremiseType : "References FirmType"
    MailStopType |o--o{ SubPremiseType : "References MailStopType"
    SubPremiseType |o--o{ SubPremiseType : "References SubPremiseType"
    ItAssetType |o--o{ SystemType : "References parent type it-asset-type"
    ItAssetType ||--|| SystemType : "extends"
    BenchmarkReferenceType |o--o{ TailoringBenchmarkReferenceType : "References parent type benchmarkReferenceType"
    BenchmarkReferenceType ||--|| TailoringBenchmarkReferenceType : "extends"
    TailoringBenchmarkReferenceType |o--o{ TailoringType : "References tailoringBenchmarkReferenceType"
    TailoringVersionType ||--o{ TailoringType : "References tailoringVersionType"
    TelephoneNumberType |o--o{ TelephoneNumber : "References parent type telephone-number-type"
    TelephoneNumberType ||--|| TelephoneNumber : "extends"
    TelephoneNumberType |o--o{ TelephoneNumberElementType : "References parent type telephone-number-type"
    TelephoneNumberType ||--|| TelephoneNumberElementType : "extends"
    BenchmarkReferenceType |o--o{ TestResultType : "References benchmarkReferenceType"
    TailoringReferenceType |o--o{ TestResultType : "References tailoringReferenceType"
    IdentityType |o--o{ TestResultType : "References identityType"
    IdrefType |o--o{ TestResultType : "References idrefType"
    TargetFactsType |o--o{ TestResultType : "References targetFactsType"
    ThoroughfarePreDirectionType |o--o{ ThoroughfareElementType : "References ThoroughfarePreDirectionType"
    ThoroughfareLeadingTypeType |o--o{ ThoroughfareElementType : "References ThoroughfareLeadingTypeType"
    ThoroughfareTrailingTypeType |o--o{ ThoroughfareElementType : "References ThoroughfareTrailingTypeType"
    ThoroughfarePostDirectionType |o--o{ ThoroughfareElementType : "References ThoroughfarePostDirectionType"
    HtmlTextWithSubType |o--o{ WarningType : "References parent type htmlTextWithSubType"
    HtmlTextWithSubType ||--|| WarningType : "extends"
    ItAssetType |o--o{ WebsiteType : "References parent type it-asset-type"
    ItAssetType ||--|| WebsiteType : "extends"
    CheckContentType |o--o{ Xccdf12CheckType : "References checkContentType"
    VersionType |o--o{ Xccdf12ItemType : "References versionType"
    ItemType |o--o{ Xccdf12ValueType : "References parent type itemType"
    ItemType ||--|| Xccdf12ValueType : "extends"
```

## Class Diagram

Shows Go struct hierarchy and associations.

```mermaid
classDiagram
    class AddressDetails {
        #int64 ID
        +string AddressType
        +string CurrentStatus
        +string ValidFromDate
        +string ValidToDate
        +string Usage
        +string AddressDetailsKey
        +string PostalServiceElements
        +string Address
        +*int64 AddressLinesID
        +string Country
    }
    class AddressElementType {
        #int64 ID
        +string Type
    }
    class AddressIdentifierElementType {
        #int64 ID
        +string IdentifierType
        +string Type
    }
    class AddressLatitudeDirectionElementType {
        #int64 ID
        +string Type
    }
    class AddressLatitudeElementType {
        #int64 ID
        +string Type
    }
    class AddressLine {
        #int64 ID
        +string Type
    }
    class AddressLineElementType {
        #int64 ID
        +string Type
    }
    class AddressLinesType {
        #int64 ID
    }
    class AddressLongitudeDirectionElementType {
        #int64 ID
        +string Type
    }
    class AddressLongitudeElementType {
        #int64 ID
        +string Type
    }
    class AddresseeIndicatorElementType {
        #int64 ID
        +string Code
    }
    class AdministrativeArea {
        #int64 ID
        +string Type
        +string UsageType
        +string Indicator
        +string AdministrativeAreaName
        +string SubAdministrativeArea
    }
    class AdministrativeAreaElementType {
        #int64 ID
        +string Type
        +string UsageType
        +string Indicator
        +string AdministrativeAreaName
        +string SubAdministrativeArea
    }
    class AdministrativeAreaNameElementType {
        #int64 ID
        +string Type
    }
    class AffectedType {
        #int64 ID
        +string Family
        +string Platform
        +string Product
    }
    class AliasElementType {
        #int64 ID
        +string Type
        +string NameType
        +string Code
    }
    class ArcType {
        #int64 ID
    }
    class ArithmeticFunctionType {
        #int64 ID
        +string Arithmetic_operation
    }
    class AssetIdentificationAssetElementType {
        #int64 ID
        +string XsdId
    }
    class RelationshipsContainerType {
        #int64 ID
        +string Relationships
    }
    class AssetsType {
        #int64 ID
        +string Asset
        +*int64 ParentID
    }
    class AssetIdentificationType {
        #int64 ID
        +string AssetRef
        +*int64 ParentID
    }
    class AssetReportCollection {
        #int64 ID
        +string XsdId
        +string ReportRequests
        +string Assets
        +string Reports
        +string ExtendedInfos
        +*int64 ParentID
    }
    class AssetReportCollectionElementType {
        #int64 ID
        +string XsdId
        +string ReportRequests
        +string Assets
        +string Reports
        +string ExtendedInfos
        +*int64 ParentID
    }
    class AssetReportingFormAssetElementType {
        #int64 ID
        +string XsdId
    }
    class AssetType {
        #int64 ID
        +string ExtendedInformation
    }
    class AssetsElementType {
        #int64 ID
        +string Asset
    }
    class BarcodeElementType {
        #int64 ID
        +string Type
    }
    class BeginFunctionType {
        #int64 ID
        +string Character
    }
    class Benchmark {
        #int64 ID
        +string XsdId
        +bool Resolved
        +string Style
        +string StyleHref
        +*int64 VersionID
        +*int64 SignatureID
    }
    class VersionType {
        #int64 ID
        +time.Time Time
        +string Update
    }
    class BenchmarkElementType {
        #int64 ID
        +string XsdId
        +bool Resolved
        +string Style
        +string StyleHref
        +*int64 VersionID
        +*int64 SignatureID
    }
    class BenchmarkReferenceType {
        #int64 ID
        +string Href
        +string XsdId
    }
    class BirthdateElementType {
        #int64 ID
    }
    class BuildingNameType {
        #int64 ID
        +string Type
        +string TypeOccurrence
    }
    class CPE2idrefType {
        #int64 ID
        +string Idref
    }
    class CanonicalizationMethodType {
        #int64 ID
        +string Algorithm
    }
    class CheckContentRefType {
        #int64 ID
        +string Href
        +string Name
    }
    class CheckContentType {
        #int64 ID
    }
    class CheckExportType {
        #int64 ID
        +string ValueId
        +string ExportName
    }
    class CheckImportType {
        #int64 ID
        +string ImportName
        +string ImportXpath
    }
    class CidrElementType {
        #int64 ID
        +*int64 ParentID
    }
    class CircuitNameElementType {
        #int64 ID
    }
    class ItAssetType {
        #int64 ID
        +*int64 ParentID
    }
    class CircuitType {
        #int64 ID
        +string CircuitName
        +*int64 ParentID
    }
    class ComplexCheckType {
        #int64 ID
        +string Operator
        +bool Negate
        +*int64 CheckID
        +*int64 ComplexCheckID
    }
    class ComplexValueType {
        #int64 ID
        +string Item
    }
    class ComputingDeviceType {
        #int64 ID
        +string DistinguishedName
        +string Connections
        +string Hostname
        +string MotherboardGuid
        +*int64 ParentID
    }
    class ConcatFunctionType {
        #int64 ID
    }
    class ConnectionsElementType {
        #int64 ID
    }
    class VariableType {
        #int64 ID
        +string XsdId
        +uint64 Version
        +string Datatype
        +string Comment
        +bool Deprecated
    }
    class ConstantVariable {
        #int64 ID
        +*int64 ParentID
    }
    class ConstantVariableElementType {
        #int64 ID
        +*int64 ParentID
    }
    class ContentElementType {
        #int64 ID
        +time.Time DataValidStartDate
        +time.Time DataValidEndDate
    }
    class ContentElementType1 {
        #int64 ID
    }
    class CountFunctionType {
        #int64 ID
    }
    class CountryElementType {
        #int64 ID
        +string CountryNameCode
    }
    class CountryName {
        #int64 ID
        +string Type
    }
    class CountryNameCodeElementType {
        #int64 ID
        +string Scheme
    }
    class CountryNameElementType {
        #int64 ID
        +string Type
    }
    class Cpe {
        #int64 ID
        +*int64 ParentID
    }
    class CpeElementType {
        #int64 ID
        +*int64 ParentID
    }
    class CriterionType {
        #int64 ID
        +bool Applicability_check
        +string Test_ref
        +bool Negate
        +string Comment
    }
    class ExtendDefinitionType {
        #int64 ID
        +bool Applicability_check
        +string Definition_ref
        +bool Negate
        +string Comment
    }
    class CriteriaType {
        #int64 ID
        +bool Applicability_check
        +string Operator
        +bool Negate
        +string Comment
        +*int64 CriteriaID
        +*int64 CriterionID
        +*int64 Extend_definitionID
    }
    class DSAKeyValueType {
        #int64 ID
        +*int64 GID
        +*int64 YID
        +*int64 JID
        +*int64 PID
        +*int64 QID
        +*int64 SeedID
        +*int64 PgenCounterID
    }
    class DataType {
        #int64 ID
        +*int64 ParentID
    }
    class DatabaseType {
        #int64 ID
        +string InstanceName
        +*int64 ParentID
    }
    class DcStatusType {
        #int64 ID
    }
    class DefinitionType {
        #int64 ID
        +string XsdId
        +uint64 Version
        +string Class
        +bool Deprecated
        +*int64 MetadataID
        +*int64 CriteriaID
    }
    class DefinitionsType {
        #int64 ID
    }
    class Department {
        #int64 ID
        +string Type
        +string DepartmentName
        +*int64 MailStopID
    }
    class MailStopType {
        #int64 ID
        +string Type
        +string MailStopName
        +string MailStopNumber
    }
    class DepartmentElementType {
        #int64 ID
        +string Type
        +string DepartmentName
        +*int64 MailStopID
    }
    class DepartmentNameElementType {
        #int64 ID
        +string Type
    }
    class DependencyNameElementType {
        #int64 ID
        +string DependencyType
        +*int64 ParentID
    }
    class DependentLocalityNameElementType {
        #int64 ID
        +string Type
    }
    class DependentLocalityNumberElementType {
        #int64 ID
        +string NameNumberOccurrence
    }
    class DependentLocalityType {
        #int64 ID
        +string Type
        +string UsageType
        +string Connector
        +string Indicator
        +string DependentLocalityName
        +string DependentLocalityNumber
        +*int64 DependentLocalityID
        +*int64 LargeMailUserID
        +*int64 PostalRouteID
    }
    class ThoroughfarePreDirectionType {
        #int64 ID
        +string Type
    }
    class ThoroughfareLeadingTypeType {
        #int64 ID
        +string Type
    }
    class ThoroughfareTrailingTypeType {
        #int64 ID
        +string Type
    }
    class ThoroughfarePostDirectionType {
        #int64 ID
        +string Type
    }
    class DependentThoroughfareElementType {
        #int64 ID
        +string Type
        +*int64 ThoroughfarePreDirectionID
        +*int64 ThoroughfareLeadingTypeID
        +*int64 ThoroughfareTrailingTypeID
        +*int64 ThoroughfarePostDirectionID
    }
    class DeprecatedInfoType {
        #int64 ID
        +string Version
        +string Reason
        +string Comment
    }
    class Dictionary20CheckType {
        #int64 ID
        +string System
        +string Href
    }
    class SchemaVersionType {
        #int64 ID
        +string Platform
        +*int64 ParentID
    }
    class Dictionary20GeneratorType {
        #int64 ID
        +string Product_name
        +string Product_version
        +string Schema_version
        +time.Time Timestamp
    }
    class ReferencesType {
        #int64 ID
        +string Reference
    }
    class Dictionary20ItemType {
        #int64 ID
        +string Name
        +bool Deprecated
        +string Deprecated_by
        +time.Time Deprecation_date
        +*int64 ReferencesID
    }
    class Dictionary20NotesType {
        #int64 ID
        +string Note
    }
    class Dictionary20TextType {
        #int64 ID
    }
    class DigestMethodType {
        #int64 ID
        +string Algorithm
    }
    class DistinguishedNameElementType {
        #int64 ID
    }
    class DocumentRootElementType {
        #int64 ID
    }
    class ElementMapItemType {
        #int64 ID
        +string Target_namespace
    }
    class ElementMapType {
        #int64 ID
        +*int64 TestID
        +*int64 ObjectID
        +*int64 StateID
        +*int64 ItemID
    }
    class EmailAddress {
        #int64 ID
    }
    class EmailAddressElementType {
        #int64 ID
    }
    class EndFunctionType {
        #int64 ID
        +string Character
    }
    class EndorsementLineCodeElementType {
        #int64 ID
        +string Type
    }
    class EntityComplexBaseType {
        #int64 ID
    }
    class EntitySimpleBaseType {
        #int64 ID
    }
    class EntityObjectAnySimpleType {
        #int64 ID
        +string Datatype
        +*int64 ParentID
    }
    class EntityObjectBinaryType {
        #int64 ID
        +string Datatype
        +*int64 ParentID
    }
    class EntityObjectBoolType {
        #int64 ID
        +string Datatype
        +*int64 ParentID
    }
    class EntityObjectFieldType {
        #int64 ID
        +string Name
        +string Entity_check
    }
    class EntityObjectFloatType {
        #int64 ID
        +string Datatype
        +*int64 ParentID
    }
    class EntityObjectIPAddressStringType {
        #int64 ID
        +string Datatype
        +*int64 ParentID
    }
    class EntityObjectIPAddressType {
        #int64 ID
        +string Datatype
        +*int64 ParentID
    }
    class EntityObjectIntType {
        #int64 ID
        +string Datatype
        +*int64 ParentID
    }
    class EntityObjectRecordType {
        #int64 ID
        +*int64 ParentID
    }
    class EntityObjectStringType {
        #int64 ID
        +string Datatype
        +*int64 ParentID
    }
    class EntityObjectVersionType {
        #int64 ID
        +string Datatype
        +*int64 ParentID
    }
    class EntityStateSimpleBaseType {
        #int64 ID
        +string Entity_check
        +string Check_existence
        +*int64 ParentID
    }
    class EntityStateAnySimpleType {
        #int64 ID
        +string Datatype
        +*int64 ParentID
    }
    class EntityStateBinaryType {
        #int64 ID
        +string Datatype
        +*int64 ParentID
    }
    class EntityStateBoolType {
        #int64 ID
        +string Datatype
        +*int64 ParentID
    }
    class EntityStateComplexBaseType {
        #int64 ID
        +string Entity_check
        +string Check_existence
        +*int64 ParentID
    }
    class EntityStateDebianEVRStringType {
        #int64 ID
        +string Datatype
        +*int64 ParentID
    }
    class EntityStateEVRStringType {
        #int64 ID
        +string Datatype
        +*int64 ParentID
    }
    class EntityStateFieldType {
        #int64 ID
        +string Name
        +string Entity_check
    }
    class EntityStateFileSetRevisionType {
        #int64 ID
        +string Datatype
        +*int64 ParentID
    }
    class EntityStateFloatType {
        #int64 ID
        +string Datatype
        +*int64 ParentID
    }
    class EntityStateIOSVersionType {
        #int64 ID
        +string Datatype
        +*int64 ParentID
    }
    class EntityStateIPAddressStringType {
        #int64 ID
        +string Datatype
        +*int64 ParentID
    }
    class EntityStateIPAddressType {
        #int64 ID
        +string Datatype
        +*int64 ParentID
    }
    class EntityStateIntType {
        #int64 ID
        +string Datatype
        +*int64 ParentID
    }
    class EntityStateRecordType {
        #int64 ID
        +*int64 ParentID
    }
    class EntityStateStringType {
        #int64 ID
        +string Datatype
        +*int64 ParentID
    }
    class EntityStateVersionType {
        #int64 ID
        +string Datatype
        +*int64 ParentID
    }
    class EscapeRegexFunctionType {
        #int64 ID
    }
    class Extended {
        #int64 ID
    }
    class ExtendedInfoElementType {
        #int64 ID
        +string XsdId
    }
    class ExtendedInformationElementType {
        #int64 ID
    }
    class ExtendedInfosElementType {
        #int64 ID
        +string ExtendedInfo
    }
    class ExternalVariable {
        #int64 ID
        +*int64 Possible_valueID
        +*int64 Possible_restrictionID
        +*int64 ParentID
    }
    class PossibleValueType {
        #int64 ID
        +string Hint
    }
    class PossibleRestrictionType {
        #int64 ID
        +string Operator
        +string Hint
    }
    class ExternalVariableElementType {
        #int64 ID
        +*int64 Possible_valueID
        +*int64 Possible_restrictionID
        +*int64 ParentID
    }
    class FactRefType {
        #int64 ID
        +string Name
    }
    class FactType {
        #int64 ID
        +string Name
        +string Type
    }
    class Filter {
        #int64 ID
        +string Action
        +*int64 ParentID
    }
    class FilterElementType {
        #int64 ID
        +string Action
        +*int64 ParentID
    }
    class FirmNameElementType {
        #int64 ID
        +string Type
    }
    class FirmType {
        #int64 ID
        +string Type
        +string FirmName
        +*int64 MailStopID
    }
    class FirstNameElementType {
        #int64 ID
        +string Type
        +string NameType
        +string Code
    }
    class IdrefType {
        #int64 ID
        +string Idref
    }
    class SubType {
        #int64 ID
        +string Use
        +*int64 ParentID
    }
    class HtmlTextWithSubType {
        #int64 ID
        +bool Override
        +*int64 SubID
    }
    class FixTextType {
        #int64 ID
        +string Fixref
        +bool Reboot
        +string Strategy
        +string Disruption
        +string Complexity
        +*int64 ParentID
    }
    class InstanceFixType {
        #int64 ID
        +string Context
    }
    class FixType {
        #int64 ID
        +string XsdId
        +bool Reboot
        +string Strategy
        +string Disruption
        +string Complexity
        +string System
        +string Platform
        +*int64 SubID
        +*int64 InstanceID
    }
    class FormerNameElementType {
        #int64 ID
        +string ValidFrom
        +string ValidTo
        +*int64 ParentID
    }
    class Fqdn {
        #int64 ID
    }
    class FqdnElementType {
        #int64 ID
    }
    class Function {
        #int64 ID
        +string Code
    }
    class GeneralSuffixElementType {
        #int64 ID
        +string Type
        +string Code
    }
    class GenerationIdentifierElementType {
        #int64 ID
        +string Type
        +string Code
    }
    class GlobToRegexFunctionType {
        #int64 ID
        +bool Glob_noescape
    }
    class SelectableItemType {
        #int64 ID
        +bool Selected
        +string Weight
        +*int64 ParentID
    }
    class GroupType {
        #int64 ID
        +string XsdId
        +*int64 SignatureID
        +*int64 ParentID
    }
    class HostElementType {
        #int64 ID
    }
    class HostnameElementType {
        #int64 ID
        +*int64 ParentID
    }
    class HtmlTextType {
        #int64 ID
        +bool Override
    }
    class IdentType {
        #int64 ID
        +string System
    }
    class IdentityType {
        #int64 ID
        +bool Authenticated
        +bool Privileged
    }
    class IdrefListType {
        #int64 ID
        +[]string Idref
    }
    class InstallationIdElementType {
        #int64 ID
    }
    class InstanceNameElementType {
        #int64 ID
    }
    class InstanceResultType {
        #int64 ID
        +string Context
        +string ParentContext
    }
    class IpAddressType {
        #int64 ID
        +string IpV4
        +string IpV6
    }
    class IpNetRangeElementType {
        #int64 ID
        +*int64 IpNetRangeStartID
        +*int64 IpNetRangeEndID
    }
    class IpV4ElementType {
        #int64 ID
        +*int64 ParentID
    }
    class IpV6ElementType {
        #int64 ID
        +*int64 ParentID
    }
    class JointPersonName {
        #int64 ID
        +string JointNameConnector
        +string Code
    }
    class JointPersonNameElementType {
        #int64 ID
        +string JointNameConnector
        +string Code
    }
    class KeyInfoType {
        #int64 ID
        +string XsdId
    }
    class KeyLineCodeElementType {
        #int64 ID
        +string Type
    }
    class KeyValueType {
        #int64 ID
    }
    class KnownAsElementType {
        #int64 ID
        +string ValidFrom
        +string ValidTo
        +*int64 ParentID
    }
    class Language20TextType {
        #int64 ID
    }
    class LargeMailUserIdentifierElementType {
        #int64 ID
        +string Type
        +string Indicator
    }
    class LargeMailUserNameElementType {
        #int64 ID
        +string Type
        +string Code
    }
    class LargeMailUserType {
        #int64 ID
        +string Type
        +string LargeMailUserName
        +string LargeMailUserIdentifier
    }
    class LastNameElementType {
        #int64 ID
        +string Type
        +string NameType
        +string Code
    }
    class LicenseElementType {
        #int64 ID
    }
    class ListType {
        #int64 ID
        +*int64 GeneratorID
    }
    class LiteralComponentType {
        #int64 ID
        +string Datatype
    }
    class LocalVariable {
        #int64 ID
        +*int64 ParentID
    }
    class LocalVariableElementType {
        #int64 ID
        +*int64 ParentID
    }
    class LocaleElementType {
        #int64 ID
        +*int64 ParentID
    }
    class Locality {
        #int64 ID
        +string Type
        +string UsageType
        +string Indicator
        +string LocalityName
        +*int64 DependentLocalityID
        +*int64 LargeMailUserID
        +*int64 PostalRouteID
    }
    class LocalityElementType {
        #int64 ID
        +string Type
        +string UsageType
        +string Indicator
        +string LocalityName
        +*int64 DependentLocalityID
        +*int64 LargeMailUserID
        +*int64 PostalRouteID
    }
    class LocalityNameElementType {
        #int64 ID
        +string Type
    }
    class LocationPoint {
        #int64 ID
        +string Latitude
        +string Longitude
        +float64 Elevation
        +string Radius
    }
    class LocationPointElementType {
        #int64 ID
        +string Latitude
        +string Longitude
        +float64 Elevation
        +string Radius
    }
    class LocationRegion {
        #int64 ID
    }
    class LocationRegionElementType {
        #int64 ID
    }
    class Locations {
        #int64 ID
    }
    class LocationsElementType {
        #int64 ID
    }
    class LocatorType {
        #int64 ID
    }
    class LogicalTestType {
        #int64 ID
        +string Operator
        +bool Negate
    }
    class MacAddressElementType {
        #int64 ID
        +*int64 ParentID
    }
    class MailStopNameElementType {
        #int64 ID
        +string Type
    }
    class MailStopNumberElementType {
        #int64 ID
        +string NameNumberSeparator
    }
    class ManifestType {
        #int64 ID
        +string XsdId
    }
    class MiddleNameElementType {
        #int64 ID
        +string Type
        +string NameType
        +string Code
    }
    class Model {
        #int64 ID
        +string System
    }
    class ModelElementType {
        #int64 ID
        +string System
    }
    class MotherboardGuidElementType {
        #int64 ID
    }
    class NameDetailsElementType {
        #int64 ID
        +string NameDetailsKey
        +string AddresseeIndicator
        +string DependencyName
        +*int64 ParentID
    }
    class NameLineType {
        #int64 ID
        +string Type
        +string NameType
        +string Code
    }
    class NamePrefixElementType {
        #int64 ID
        +string Type
        +string NameType
        +string Code
    }
    class NetworkInterfaceType {
        #int64 ID
        +string MacAddress
        +string Url
        +*int64 SubnetMaskID
        +*int64 DefaultRouteID
    }
    class NetworkNameElementType {
        #int64 ID
    }
    class NetworkType {
        #int64 ID
        +string NetworkName
        +string IpNetRange
        +string Cidr
        +*int64 ParentID
    }
    class Notes {
        #int64 ID
        +string Note
        +*int64 ParentID
    }
    class NotesElementType {
        #int64 ID
        +string Note
        +*int64 ParentID
    }
    class NoticeType {
        #int64 ID
        +string XsdId
    }
    class Ns09XmldsigObjectType {
        #int64 ID
        +string XsdId
        +string MimeType
        +string Encoding
    }
    class Ns09XmldsigReferenceType {
        #int64 ID
        +string XsdId
        +string URI
        +string Type
    }
    class Ns09XmldsigSignatureType {
        #int64 ID
        +string XsdId
    }
    class OasisNamesTcCiqXNameDetails {
        #int64 ID
        +string NameDetailsKey
        +string AddresseeIndicator
        +string DependencyName
        +*int64 ParentID
    }
    class OasisNamesTcCiqXOrganisationNameDetails {
        #int64 ID
        +string OrganisationFormerName
        +string OrganisationKnownAs
        +*int64 ParentID
    }
    class OasisNamesTcCiqXPersonName {
        #int64 ID
        +string FormerName
        +string KnownAs
        +*int64 ParentID
    }
    class ObjectComponentType {
        #int64 ID
        +string Object_ref
        +string Item_field
        +string Record_field
    }
    class ObjectRef {
        #int64 ID
        +string RefId
    }
    class ObjectRefElementType {
        #int64 ID
        +string RefId
    }
    class ObjectRefType {
        #int64 ID
        +string Object_ref
    }
    class ObjectsType {
        #int64 ID
    }
    class OrganisationFormerNameElementType {
        #int64 ID
        +string ValidFrom
        +string ValidTo
        +*int64 ParentID
    }
    class OrganisationKnownAsElementType {
        #int64 ID
        +string ValidFrom
        +string ValidTo
        +*int64 ParentID
    }
    class OrganisationNameDetailsElementType {
        #int64 ID
        +string OrganisationFormerName
        +string OrganisationKnownAs
        +*int64 ParentID
    }
    class OrganisationNameElementType {
        #int64 ID
        +string Type
        +string NameType
        +string Code
    }
    class OrganisationTypeElementType {
        #int64 ID
        +string Type
        +string NameType
        +string Code
    }
    class OrganizationType {
        #int64 ID
        +*int64 ParentID
    }
    class OtherNameElementType {
        #int64 ID
        +string Type
        +string NameType
        +string Code
    }
    class OvalDefinitions {
        #int64 ID
        +*int64 GeneratorID
        +*int64 DefinitionsID
        +*int64 TestsID
        +*int64 ObjectsID
        +*int64 StatesID
        +*int64 VariablesID
    }
    class TestsType {
        #int64 ID
    }
    class StatesType {
        #int64 ID
    }
    class VariablesType {
        #int64 ID
    }
    class OvalDefinitionsElementType {
        #int64 ID
        +*int64 GeneratorID
        +*int64 DefinitionsID
        +*int64 TestsID
        +*int64 ObjectsID
        +*int64 StatesID
        +*int64 VariablesID
    }
    class OvalMitreOrgOvalGeneratorType {
        #int64 ID
        +string Product_name
        +string Product_version
        +time.Time Timestamp
    }
    class OvalMitreOrgOvalMessageType {
        #int64 ID
        +string Level
    }
    class OvalMitreOrgOvalMetadataType {
        #int64 ID
        +string Title
        +string Description
    }
    class OvalMitreOrgOvalNotesType {
        #int64 ID
        +string Note
    }
    class OvalMitreOrgOvalObjectType {
        #int64 ID
        +string XsdId
        +uint64 Version
        +string Comment
        +bool Deprecated
    }
    class OvalMitreOrgOvalReferenceType {
        #int64 ID
        +string Source
        +string Ref_id
        +string Ref_url
    }
    class OvalMitreOrgOvalValueType {
        #int64 ID
    }
    class OverrideType {
        #int64 ID
        +time.Time Time
        +string Authority
        +*int64 OldResultID
        +*int64 NewResultID
        +*int64 RemarkID
    }
    class OverrideableCPE2idrefType {
        #int64 ID
        +bool Override
        +*int64 ParentID
    }
    class PGPDataType {
        #int64 ID
        +[]byte PGPKeyID
        +[]byte PGPKeyPacket
    }
    class ParamType {
        #int64 ID
        +string Name
    }
    class PersonNameElementType {
        #int64 ID
        +string FormerName
        +string KnownAs
        +*int64 ParentID
    }
    class PersonType {
        #int64 ID
        +string Birthdate
        +*int64 ParentID
    }
    class PlainTextType {
        #int64 ID
        +string XsdId
    }
    class PlatformSpecification {
        #int64 ID
    }
    class PlatformSpecificationElementType {
        #int64 ID
    }
    class PlatformType {
        #int64 ID
        +string XsdId
        +*int64 LogicalTestID
    }
    class PortElementType {
        #int64 ID
        +*int64 ParentID
    }
    class PortRangeElementType {
        #int64 ID
        +string LowerBound
        +string UpperBound
    }
    class PostBox {
        #int64 ID
        +string Type
        +string Indicator
        +string PostBoxNumber
        +string PostBoxNumberPrefix
        +string PostBoxNumberSuffix
        +string PostBoxNumberExtension
        +*int64 FirmID
    }
    class PostBoxElementType {
        #int64 ID
        +string Type
        +string Indicator
        +string PostBoxNumber
        +string PostBoxNumberPrefix
        +string PostBoxNumberSuffix
        +string PostBoxNumberExtension
        +*int64 FirmID
    }
    class PostBoxNumberElementType {
        #int64 ID
    }
    class PostBoxNumberExtensionElementType {
        #int64 ID
        +string NumberExtensionSeparator
    }
    class PostBoxNumberPrefixElementType {
        #int64 ID
        +string NumberPrefixSeparator
    }
    class PostBoxNumberSuffixElementType {
        #int64 ID
        +string NumberSuffixSeparator
    }
    class PostOffice {
        #int64 ID
        +string Type
        +string Indicator
        +*int64 PostalRouteID
        +string PostOfficeName
        +string PostOfficeNumber
    }
    class PostalRouteType {
        #int64 ID
        +string Type
        +string PostalRouteName
        +string PostalRouteNumber
    }
    class PostOfficeElementType {
        #int64 ID
        +string Type
        +string Indicator
        +*int64 PostalRouteID
        +string PostOfficeName
        +string PostOfficeNumber
    }
    class PostOfficeNameElementType {
        #int64 ID
        +string Type
    }
    class PostOfficeNumberElementType {
        #int64 ID
        +string Indicator
        +string IndicatorOccurrence
    }
    class PostTownElementType {
        #int64 ID
        +string Type
        +string PostTownName
        +string PostTownSuffix
    }
    class PostTownNameElementType {
        #int64 ID
        +string Type
    }
    class PostTownSuffixElementType {
        #int64 ID
    }
    class PostalCode {
        #int64 ID
        +string Type
        +string PostalCodeNumber
        +string PostalCodeNumberExtension
        +string PostTown
    }
    class PostalCodeElementType {
        #int64 ID
        +string Type
        +string PostalCodeNumber
        +string PostalCodeNumberExtension
        +string PostTown
    }
    class PostalCodeNumberElementType {
        #int64 ID
        +string Type
    }
    class PostalCodeNumberExtensionElementType {
        #int64 ID
        +string Type
        +string NumberExtensionSeparator
    }
    class PostalRouteNameElementType {
        #int64 ID
        +string Type
    }
    class PostalRouteNumberElementType {
        #int64 ID
    }
    class PostalServiceElementsElementType {
        #int64 ID
        +string Type
        +string AddressIdentifier
        +string EndorsementLineCode
        +string KeyLineCode
        +string Barcode
        +string SortingCode
        +string AddressLatitude
        +string AddressLatitudeDirection
        +string AddressLongitude
        +string AddressLongitudeDirection
        +string SupplementaryPostalServiceData
    }
    class PrecedingTitleElementType {
        #int64 ID
        +string Type
        +string Code
    }
    class Premise {
        #int64 ID
        +string Type
        +string PremiseDependency
        +string PremiseDependencyType
        +string PremiseThoroughfareConnector
        +string PremiseName
        +*int64 MailStopID
        +string PremiseLocation
        +string PremiseNumberRange
        +*int64 FirmID
    }
    class PremiseElementType {
        #int64 ID
        +string Type
        +string PremiseDependency
        +string PremiseDependencyType
        +string PremiseThoroughfareConnector
        +string PremiseName
        +*int64 MailStopID
        +string PremiseLocation
        +string PremiseNumberRange
        +*int64 FirmID
    }
    class PremiseLocationElementType {
        #int64 ID
    }
    class PremiseNameElementType {
        #int64 ID
        +string Type
        +string TypeOccurrence
    }
    class PremiseNumber {
        #int64 ID
        +string NumberType
        +string Type
        +string Indicator
        +string IndicatorOccurrence
        +string NumberTypeOccurrence
    }
    class PremiseNumberElementType {
        #int64 ID
        +string NumberType
        +string Type
        +string Indicator
        +string IndicatorOccurrence
        +string NumberTypeOccurrence
    }
    class PremiseNumberPrefix {
        #int64 ID
        +string NumberPrefixSeparator
        +string Type
    }
    class PremiseNumberPrefixElementType {
        #int64 ID
        +string NumberPrefixSeparator
        +string Type
    }
    class PremiseNumberRangeElementType {
        #int64 ID
        +string RangeType
        +string Indicator
        +string Separator
        +string Type
        +string IndicatorOccurence
        +string NumberRangeOccurence
        +string PremiseNumberRangeFrom
        +string PremiseNumberRangeTo
    }
    class PremiseNumberRangeFromElementType {
        #int64 ID
    }
    class PremiseNumberRangeToElementType {
        #int64 ID
    }
    class PremiseNumberSuffix {
        #int64 ID
        +string NumberSuffixSeparator
        +string Type
    }
    class PremiseNumberSuffixElementType {
        #int64 ID
        +string NumberSuffixSeparator
        +string Type
    }
    class ProfileNoteType {
        #int64 ID
        +string Tag
        +*int64 SubID
    }
    class ProfileRefineRuleType {
        #int64 ID
        +string Idref
        +string Weight
        +string Selector
        +string Severity
        +string Role
    }
    class ProfileRefineValueType {
        #int64 ID
        +string Idref
        +string Selector
        +string Operator
    }
    class ProfileSelectType {
        #int64 ID
        +string Idref
        +bool Selected
    }
    class ProfileSetComplexValueType {
        #int64 ID
        +string Idref
        +*int64 ParentID
    }
    class ProfileSetValueType {
        #int64 ID
        +string Idref
    }
    class ProfileType {
        #int64 ID
        +string XsdId
        +bool ProhibitChanges
        +bool Abstract
        +string NoteTag
        +string Extends
        +*int64 VersionID
        +*int64 SignatureID
        +*int64 SelectID
        +*int64 SetComplexValueID
        +*int64 SetValueID
        +*int64 RefineValueID
        +*int64 RefineRuleID
    }
    class ProtocolElementType {
        #int64 ID
    }
    class RSAKeyValueType {
        #int64 ID
        +*int64 ModulusID
        +*int64 ExponentID
    }
    class ReferenceElementType {
        #int64 ID
        +string Href
    }
    class RegexCaptureFunctionType {
        #int64 ID
        +string Pattern
    }
    class RelationshipType {
        #int64 ID
        +string Type
        +string Scope
        +string Subject
        +string Ref
    }
    class RelationshipsElementType {
        #int64 ID
    }
    class RemoteResource {
        #int64 ID
    }
    class RemoteResourceElementType {
        #int64 ID
    }
    class ReportRequestType {
        #int64 ID
        +string XsdId
        +string Content
    }
    class ReportRequestsElementType {
        #int64 ID
    }
    class ReportType {
        #int64 ID
        +string XsdId
        +string Content
    }
    class ReportsElementType {
        #int64 ID
    }
    class ResourceType {
        #int64 ID
    }
    class RestrictionType {
        #int64 ID
        +string Operation
    }
    class RetrievalMethodType {
        #int64 ID
        +string URI
        +string Type
    }
    class RuleResultType {
        #int64 ID
        +string Idref
        +string Role
        +string Severity
        +time.Time Time
        +string Version
        +string Weight
        +*int64 ResultID
        +*int64 ComplexCheckID
    }
    class RuleType {
        #int64 ID
        +string XsdId
        +string Role
        +string Severity
        +bool Multiple
        +string ImpactMetric
        +*int64 SignatureID
        +*int64 ComplexCheckID
        +*int64 ParentID
    }
    class SPKIDataType {
        #int64 ID
        +[]byte SPKISexp
    }
    class ScoreType {
        #int64 ID
        +string System
        +string Maximum
    }
    class SelChoicesType {
        #int64 ID
        +bool MustMatch
        +string Selector
        +string Choice
        +*int64 ComplexChoiceID
    }
    class SelComplexValueType {
        #int64 ID
        +string Selector
        +*int64 ParentID
    }
    class SelNumType {
        #int64 ID
        +string Selector
    }
    class SelStringType {
        #int64 ID
        +string Selector
    }
    class ServiceType {
        #int64 ID
        +string Host
        +string Port
        +string PortRange
        +string Protocol
        +*int64 ParentID
    }
    class Set {
        #int64 ID
        +string Set_operator
    }
    class SetElementType {
        #int64 ID
        +string Set_operator
    }
    class SignatureMethodType {
        #int64 ID
        +string Algorithm
        +*int64 HMACOutputLengthID
    }
    class SignaturePropertiesType {
        #int64 ID
        +string XsdId
    }
    class SignaturePropertyType {
        #int64 ID
        +string Target
        +string XsdId
    }
    class SignatureValueType {
        #int64 ID
        +string XsdId
    }
    class SignedInfoType {
        #int64 ID
        +string XsdId
    }
    class Simple {
        #int64 ID
    }
    class SoftwareType {
        #int64 ID
        +string InstallationId
        +string License
        +*int64 ParentID
    }
    class SortingCodeElementType {
        #int64 ID
        +string Type
    }
    class SplitFunctionType {
        #int64 ID
        +string Delimiter
    }
    class StateRefType {
        #int64 ID
        +string State_ref
    }
    class StateType {
        #int64 ID
        +string XsdId
        +uint64 Version
        +string Operator
        +string Comment
        +bool Deprecated
    }
    class Status {
        #int64 ID
        +time.Time Date
        +*int64 ParentID
    }
    class StatusElementType {
        #int64 ID
        +time.Time Date
        +*int64 ParentID
    }
    class SubAdministrativeAreaElementType {
        #int64 ID
        +string Type
        +string UsageType
        +string Indicator
        +string SubAdministrativeAreaName
    }
    class SubAdministrativeAreaNameElementType {
        #int64 ID
        +string Type
    }
    class SubPremiseLocationElementType {
        #int64 ID
    }
    class SubPremiseNameElementType {
        #int64 ID
        +string Type
        +string TypeOccurrence
    }
    class SubPremiseNumberElementType {
        #int64 ID
        +string Indicator
        +string IndicatorOccurrence
        +string NumberTypeOccurrence
        +string PremiseNumberSeparator
        +string Type
    }
    class SubPremiseNumberPrefixElementType {
        #int64 ID
        +string NumberPrefixSeparator
        +string Type
    }
    class SubPremiseNumberSuffixElementType {
        #int64 ID
        +string NumberSuffixSeparator
        +string Type
    }
    class SubPremiseType {
        #int64 ID
        +string Type
        +string SubPremiseName
        +string SubPremiseNumberPrefix
        +string SubPremiseNumberSuffix
        +*int64 FirmID
        +*int64 MailStopID
        +*int64 SubPremiseID
        +string SubPremiseLocation
        +string SubPremiseNumber
    }
    class SubstringFunctionType {
        #int64 ID
        +int32 Substring_start
        +int32 Substring_length
    }
    class SuffixElementType {
        #int64 ID
        +string Type
        +string Code
    }
    class SupplementaryPostalServiceDataElementType {
        #int64 ID
        +string Type
    }
    class SyntheticId {
        #int64 ID
        +string Resource
        +string XsdId
    }
    class SyntheticIdElementType {
        #int64 ID
        +string Resource
        +string XsdId
    }
    class SystemNameElementType {
        #int64 ID
    }
    class SystemType {
        #int64 ID
        +string SystemName
        +string Version
        +*int64 ParentID
    }
    class TailoringBenchmarkReferenceType {
        #int64 ID
        +string Version
        +*int64 ParentID
    }
    class TailoringReferenceType {
        #int64 ID
        +string Href
        +string XsdId
        +string Version
        +time.Time Time
    }
    class TailoringVersionType {
        #int64 ID
        +time.Time Time
    }
    class TailoringType {
        #int64 ID
        +string XsdId
        +*int64 BenchmarkID
        +*int64 VersionID
        +*int64 SignatureID
    }
    class TargetFactsType {
        #int64 ID
    }
    class TargetIdRefType {
        #int64 ID
        +string System
        +string Href
        +string Name
    }
    class TelephoneNumber {
        #int64 ID
        +*int64 ParentID
    }
    class TelephoneNumberElementType {
        #int64 ID
        +*int64 ParentID
    }
    class TestResultType {
        #int64 ID
        +string XsdId
        +time.Time StartTime
        +time.Time EndTime
        +string TestSystem
        +string Version
        +*int64 BenchmarkID
        +*int64 TailoringFileID
        +string Organization
        +*int64 IdentityID
        +*int64 ProfileID
        +string Target
        +string TargetAddress
        +*int64 TargetFactsID
        +*int64 SignatureID
        +*int64 TargetIdRefID
        +*int64 SetValueID
        +*int64 SetComplexValueID
    }
    class TestType {
        #int64 ID
        +string XsdId
        +uint64 Version
        +string Check_existence
        +string Check
        +string State_operator
        +string Comment
        +bool Deprecated
    }
    class TextWithSubType {
        #int64 ID
        +bool Override
    }
    class Thoroughfare {
        #int64 ID
        +string Type
        +string DependentThoroughfares
        +string DependentThoroughfaresIndicator
        +string DependentThoroughfaresConnector
        +string DependentThoroughfaresType
        +*int64 ThoroughfarePreDirectionID
        +*int64 ThoroughfareLeadingTypeID
        +*int64 ThoroughfareTrailingTypeID
        +*int64 ThoroughfarePostDirectionID
        +string DependentThoroughfare
        +string ThoroughfareNumberRange
        +*int64 DependentLocalityID
        +*int64 FirmID
    }
    class ThoroughfareElementType {
        #int64 ID
        +string Type
        +string DependentThoroughfares
        +string DependentThoroughfaresIndicator
        +string DependentThoroughfaresConnector
        +string DependentThoroughfaresType
        +*int64 ThoroughfarePreDirectionID
        +*int64 ThoroughfareLeadingTypeID
        +*int64 ThoroughfareTrailingTypeID
        +*int64 ThoroughfarePostDirectionID
        +string DependentThoroughfare
        +string ThoroughfareNumberRange
        +*int64 DependentLocalityID
        +*int64 FirmID
    }
    class ThoroughfareNameType {
        #int64 ID
        +string Type
    }
    class ThoroughfareNumber {
        #int64 ID
        +string NumberType
        +string Type
        +string Indicator
        +string IndicatorOccurrence
        +string NumberOccurrence
    }
    class ThoroughfareNumberElementType {
        #int64 ID
        +string NumberType
        +string Type
        +string Indicator
        +string IndicatorOccurrence
        +string NumberOccurrence
    }
    class ThoroughfareNumberFromElementType {
        #int64 ID
    }
    class ThoroughfareNumberPrefix {
        #int64 ID
        +string NumberPrefixSeparator
        +string Type
    }
    class ThoroughfareNumberPrefixElementType {
        #int64 ID
        +string NumberPrefixSeparator
        +string Type
    }
    class ThoroughfareNumberRangeElementType {
        #int64 ID
        +string RangeType
        +string Indicator
        +string Separator
        +string IndicatorOccurrence
        +string NumberRangeOccurrence
        +string Type
        +string ThoroughfareNumberFrom
        +string ThoroughfareNumberTo
    }
    class ThoroughfareNumberSuffix {
        #int64 ID
        +string NumberSuffixSeparator
        +string Type
    }
    class ThoroughfareNumberSuffixElementType {
        #int64 ID
        +string NumberSuffixSeparator
        +string Type
    }
    class ThoroughfareNumberToElementType {
        #int64 ID
    }
    class TimeDifferenceFunctionType {
        #int64 ID
        +string Format_1
        +string Format_2
    }
    class TitleElementType {
        #int64 ID
        +string Type
        +string Code
    }
    class TitleEltType {
        #int64 ID
    }
    class TransformType {
        #int64 ID
        +string Algorithm
        +string XPath
    }
    class TransformsType {
        #int64 ID
    }
    class UniqueFunctionType {
        #int64 ID
    }
    class UriRefType {
        #int64 ID
        +string Uri
    }
    class UrlElementType {
        #int64 ID
    }
    class VariableComponentType {
        #int64 ID
        +string Var_ref
    }
    class VersionElementType {
        #int64 ID
    }
    class WarningType {
        #int64 ID
        +string Category
        +*int64 ParentID
    }
    class WebsiteType {
        #int64 ID
        +string DocumentRoot
        +string Locale
        +*int64 ParentID
    }
    class WebsiteUrl {
        #int64 ID
    }
    class WebsiteUrlElementType {
        #int64 ID
    }
    class X509DataType {
        #int64 ID
        +*int64 X509IssuerSerialID
        +[]byte X509SKI
        +string X509SubjectName
        +[]byte X509Certificate
        +[]byte X509CRL
    }
    class X509IssuerSerialType {
        #int64 ID
        +string X509IssuerName
        +int64 X509SerialNumber
    }
    class XAL {
        #int64 ID
        +string Version
    }
    class XALElementType {
        #int64 ID
        +string Version
    }
    class XNL {
        #int64 ID
        +string Version
    }
    class XNLElementType {
        #int64 ID
        +string Version
    }
    class Xccdf12CheckType {
        #int64 ID
        +string System
        +bool Negate
        +string XsdId
        +string Selector
        +bool MultiCheck
        +*int64 CheckContentID
    }
    class Xccdf12ItemType {
        #int64 ID
        +bool Abstract
        +string ClusterId
        +string Extends
        +bool Hidden
        +bool ProhibitChanges
        +string XsdId
        +*int64 VersionID
    }
    class Xccdf12MessageType {
        #int64 ID
        +string Severity
    }
    class Xccdf12MetadataType {
        #int64 ID
    }
    class Xccdf12ReferenceType {
        #int64 ID
        +string Href
        +bool Override
    }
    class Xccdf12SignatureType {
        #int64 ID
    }
    class Xccdf12TextType {
        #int64 ID
        +bool Override
    }
    class Xccdf12ValueType {
        #int64 ID
        +string XsdId
        +string Type
        +string Operator
        +bool Interactive
        +string InterfaceHint
        +*int64 SignatureID
        +*int64 ValueID
        +*int64 ComplexValueID
        +*int64 DefaultID
        +*int64 ComplexDefaultID
        +*int64 ParentID
    }

    RelationshipsContainerType <|-- AssetsType
    AssetsType <|-- AssetIdentificationType
    RelationshipsContainerType <|-- AssetReportCollection
    RelationshipsContainerType <|-- AssetReportCollectionElementType
    VersionType "1" o-- "*" BenchmarkElementType
    CidrType <|-- CidrElementType
    AssetType <|-- ItAssetType
    ItAssetType <|-- CircuitType
    ComplexCheckType "1" o-- "*" ComplexCheckType
    ItAssetType <|-- ComputingDeviceType
    VariableType <|-- ConstantVariable
    VariableType <|-- ConstantVariableElementType
    CpeType <|-- Cpe
    CpeType <|-- CpeElementType
    CriteriaType "1" o-- "*" CriteriaType
    CriterionType "1" o-- "*" CriteriaType
    ExtendDefinitionType "1" o-- "*" CriteriaType
    AssetType <|-- DataType
    ItAssetType <|-- DatabaseType
    CriteriaType "1" o-- "*" DefinitionType
    MailStopType "1" o-- "*" DepartmentElementType
    NameDetails <|-- DependencyNameElementType
    DependentLocalityType "1" o-- "*" DependentLocalityType
    ThoroughfarePreDirectionType "1" o-- "*" DependentThoroughfareElementType
    ThoroughfareLeadingTypeType "1" o-- "*" DependentThoroughfareElementType
    ThoroughfareTrailingTypeType "1" o-- "*" DependentThoroughfareElementType
    ThoroughfarePostDirectionType "1" o-- "*" DependentThoroughfareElementType
    SchemaVersionPattern <|-- SchemaVersionType
    SchemaVersionType "1" o-- "*" Dictionary20GeneratorType
    ReferencesType "1" o-- "*" Dictionary20ItemType
    ElementMapItemType "1" o-- "*" ElementMapType
    ElementMapItemType "1" o-- "*" ElementMapType
    ElementMapItemType "1" o-- "*" ElementMapType
    ElementMapItemType "1" o-- "*" ElementMapType
    EntitySimpleBaseType <|-- EntityObjectAnySimpleType
    EntitySimpleBaseType <|-- EntityObjectBinaryType
    EntitySimpleBaseType <|-- EntityObjectBoolType
    EntitySimpleBaseType <|-- EntityObjectFloatType
    EntitySimpleBaseType <|-- EntityObjectIPAddressStringType
    EntitySimpleBaseType <|-- EntityObjectIPAddressType
    EntitySimpleBaseType <|-- EntityObjectIntType
    EntityComplexBaseType <|-- EntityObjectRecordType
    EntitySimpleBaseType <|-- EntityObjectStringType
    EntitySimpleBaseType <|-- EntityObjectVersionType
    EntitySimpleBaseType <|-- EntityStateSimpleBaseType
    EntityStateSimpleBaseType <|-- EntityStateAnySimpleType
    EntityStateSimpleBaseType <|-- EntityStateBinaryType
    EntityStateSimpleBaseType <|-- EntityStateBoolType
    EntityComplexBaseType <|-- EntityStateComplexBaseType
    EntityStateSimpleBaseType <|-- EntityStateDebianEVRStringType
    EntityStateSimpleBaseType <|-- EntityStateEVRStringType
    EntityStateSimpleBaseType <|-- EntityStateFileSetRevisionType
    EntityStateSimpleBaseType <|-- EntityStateFloatType
    EntityStateSimpleBaseType <|-- EntityStateIOSVersionType
    EntityStateSimpleBaseType <|-- EntityStateIPAddressStringType
    EntityStateSimpleBaseType <|-- EntityStateIPAddressType
    EntityStateSimpleBaseType <|-- EntityStateIntType
    EntityStateComplexBaseType <|-- EntityStateRecordType
    EntityStateSimpleBaseType <|-- EntityStateStringType
    EntityStateSimpleBaseType <|-- EntityStateVersionType
    VariableType <|-- ExternalVariable
    VariableType <|-- ExternalVariableElementType
    PossibleValueType "1" o-- "*" ExternalVariableElementType
    PossibleRestrictionType "1" o-- "*" ExternalVariableElementType
    StateIDPattern <|-- Filter
    StateIDPattern <|-- FilterElementType
    MailStopType "1" o-- "*" FirmType
    IdrefType <|-- SubType
    SubType "1" o-- "*" HtmlTextWithSubType
    HtmlTextWithSubType <|-- FixTextType
    SubType "1" o-- "*" FixType
    InstanceFixType "1" o-- "*" FixType
    PersonName <|-- FormerNameElementType
    ItemType <|-- SelectableItemType
    SelectableItemType <|-- GroupType
    HostnameType <|-- HostnameElementType
    IpAddressType "1" o-- "*" IpNetRangeElementType
    IpAddressType "1" o-- "*" IpNetRangeElementType
    Ipv4Type <|-- IpV4ElementType
    Ipv6Type <|-- IpV6ElementType
    PersonName <|-- KnownAsElementType
    VariableType <|-- LocalVariable
    VariableType <|-- LocalVariableElementType
    LocaleType <|-- LocaleElementType
    DependentLocalityType "1" o-- "*" LocalityElementType
    MacAddressType <|-- MacAddressElementType
    NameDetails <|-- NameDetailsElementType
    IpAddressType "1" o-- "*" NetworkInterfaceType
    IpAddressType "1" o-- "*" NetworkInterfaceType
    ItAssetType <|-- NetworkType
    NotesType <|-- Notes
    NotesType <|-- NotesElementType
    NameDetails <|-- OasisNamesTcCiqXNameDetails
    OrganisationNameDetails <|-- OasisNamesTcCiqXOrganisationNameDetails
    PersonName <|-- OasisNamesTcCiqXPersonName
    OrganisationNameDetails <|-- OrganisationFormerNameElementType
    OrganisationNameDetails <|-- OrganisationKnownAsElementType
    OrganisationNameDetails <|-- OrganisationNameDetailsElementType
    AssetType <|-- OrganizationType
    DefinitionsType "1" o-- "*" OvalDefinitionsElementType
    TestsType "1" o-- "*" OvalDefinitionsElementType
    ObjectsType "1" o-- "*" OvalDefinitionsElementType
    StatesType "1" o-- "*" OvalDefinitionsElementType
    VariablesType "1" o-- "*" OvalDefinitionsElementType
    CPE2idrefType <|-- OverrideableCPE2idrefType
    PersonName <|-- PersonNameElementType
    AssetType <|-- PersonType
    LogicalTestType "1" o-- "*" PlatformType
    PortType <|-- PortElementType
    FirmType "1" o-- "*" PostBoxElementType
    PostalRouteType "1" o-- "*" PostOfficeElementType
    MailStopType "1" o-- "*" PremiseElementType
    SubType "1" o-- "*" ProfileNoteType
    ComplexValueType <|-- ProfileSetComplexValueType
    VersionType "1" o-- "*" ProfileType
    SelectableItemType <|-- RuleType
    ComplexValueType "1" o-- "*" SelChoicesType
    ComplexValueType <|-- SelComplexValueType
    ItAssetType <|-- ServiceType
    ItAssetType <|-- SoftwareType
    StatusType <|-- Status
    StatusType <|-- StatusElementType
    FirmType "1" o-- "*" SubPremiseType
    MailStopType "1" o-- "*" SubPremiseType
    SubPremiseType "1" o-- "*" SubPremiseType
    ItAssetType <|-- SystemType
    BenchmarkReferenceType <|-- TailoringBenchmarkReferenceType
    TailoringBenchmarkReferenceType "1" o-- "*" TailoringType
    TailoringVersionType "1" o-- "*" TailoringType
    TelephoneNumberType <|-- TelephoneNumber
    TelephoneNumberType <|-- TelephoneNumberElementType
    BenchmarkReferenceType "1" o-- "*" TestResultType
    TailoringReferenceType "1" o-- "*" TestResultType
    IdentityType "1" o-- "*" TestResultType
    IdrefType "1" o-- "*" TestResultType
    TargetFactsType "1" o-- "*" TestResultType
    ThoroughfarePreDirectionType "1" o-- "*" ThoroughfareElementType
    ThoroughfareLeadingTypeType "1" o-- "*" ThoroughfareElementType
    ThoroughfareTrailingTypeType "1" o-- "*" ThoroughfareElementType
    ThoroughfarePostDirectionType "1" o-- "*" ThoroughfareElementType
    HtmlTextWithSubType <|-- WarningType
    ItAssetType <|-- WebsiteType
    CheckContentType "1" o-- "*" Xccdf12CheckType
    VersionType "1" o-- "*" Xccdf12ItemType
    ItemType <|-- Xccdf12ValueType
```

## Data Flow Diagram

Shows parent-child data flow between entities.

```mermaid
flowchart TD
    addressdetails[["AddressDetails"]]
    addresselementtype[["AddressElementType"]]
    addressidentifierelementtype[["AddressIdentifierElementType"]]
    addresslatitudedirectionelementtype[["AddressLatitudeDirectionElementType"]]
    addresslatitudeelementtype[["AddressLatitudeElementType"]]
    addressline[["AddressLine"]]
    addresslineelementtype[["AddressLineElementType"]]
    addresslinestype[["AddressLinesType"]]
    addresslongitudedirectionelementtype[["AddressLongitudeDirectionElementType"]]
    addresslongitudeelementtype[["AddressLongitudeElementType"]]
    addresseeindicatorelementtype[["AddresseeIndicatorElementType"]]
    administrativearea[["AdministrativeArea"]]
    administrativeareaelementtype[["AdministrativeAreaElementType"]]
    administrativeareanameelementtype[["AdministrativeAreaNameElementType"]]
    affectedtype[["AffectedType"]]
    aliaselementtype[["AliasElementType"]]
    arctype[["ArcType"]]
    arithmeticfunctiontype[["ArithmeticFunctionType"]]
    assetidentificationassetelementtype[["AssetIdentificationAssetElementType"]]
    relationshipscontainertype[["RelationshipsContainerType"]]
    assetstype["AssetsType"]
    assetidentificationtype["AssetIdentificationType"]
    assetreportcollection["AssetReportCollection"]
    assetreportcollectionelementtype["AssetReportCollectionElementType"]
    assetreportingformassetelementtype[["AssetReportingFormAssetElementType"]]
    assettype[["AssetType"]]
    assetselementtype[["AssetsElementType"]]
    barcodeelementtype[["BarcodeElementType"]]
    beginfunctiontype[["BeginFunctionType"]]
    benchmark[["Benchmark"]]
    versiontype[["VersionType"]]
    benchmarkelementtype["BenchmarkElementType"]
    benchmarkreferencetype[["BenchmarkReferenceType"]]
    birthdateelementtype[["BirthdateElementType"]]
    buildingnametype[["BuildingNameType"]]
    cpe2idreftype[["CPE2idrefType"]]
    canonicalizationmethodtype[["CanonicalizationMethodType"]]
    checkcontentreftype[["CheckContentRefType"]]
    checkcontenttype[["CheckContentType"]]
    checkexporttype[["CheckExportType"]]
    checkimporttype[["CheckImportType"]]
    cidrelementtype["CidrElementType"]
    circuitnameelementtype[["CircuitNameElementType"]]
    itassettype["ItAssetType"]
    circuittype["CircuitType"]
    complexchecktype["ComplexCheckType"]
    complexvaluetype[["ComplexValueType"]]
    computingdevicetype["ComputingDeviceType"]
    concatfunctiontype[["ConcatFunctionType"]]
    connectionselementtype[["ConnectionsElementType"]]
    variabletype[["VariableType"]]
    constantvariable["ConstantVariable"]
    constantvariableelementtype["ConstantVariableElementType"]
    contentelementtype[["ContentElementType"]]
    contentelementtype1[["ContentElementType1"]]
    countfunctiontype[["CountFunctionType"]]
    countryelementtype[["CountryElementType"]]
    countryname[["CountryName"]]
    countrynamecodeelementtype[["CountryNameCodeElementType"]]
    countrynameelementtype[["CountryNameElementType"]]
    cpe["Cpe"]
    cpeelementtype["CpeElementType"]
    criteriontype[["CriterionType"]]
    extenddefinitiontype[["ExtendDefinitionType"]]
    criteriatype["CriteriaType"]
    dsakeyvaluetype[["DSAKeyValueType"]]
    datatype["DataType"]
    databasetype["DatabaseType"]
    dcstatustype[["DcStatusType"]]
    definitiontype["DefinitionType"]
    definitionstype[["DefinitionsType"]]
    department[["Department"]]
    mailstoptype[["MailStopType"]]
    departmentelementtype["DepartmentElementType"]
    departmentnameelementtype[["DepartmentNameElementType"]]
    dependencynameelementtype["DependencyNameElementType"]
    dependentlocalitynameelementtype[["DependentLocalityNameElementType"]]
    dependentlocalitynumberelementtype[["DependentLocalityNumberElementType"]]
    dependentlocalitytype["DependentLocalityType"]
    thoroughfarepredirectiontype[["ThoroughfarePreDirectionType"]]
    thoroughfareleadingtypetype[["ThoroughfareLeadingTypeType"]]
    thoroughfaretrailingtypetype[["ThoroughfareTrailingTypeType"]]
    thoroughfarepostdirectiontype[["ThoroughfarePostDirectionType"]]
    dependentthoroughfareelementtype["DependentThoroughfareElementType"]
    deprecatedinfotype[["DeprecatedInfoType"]]
    dictionary20checktype[["Dictionary20CheckType"]]
    schemaversiontype["SchemaVersionType"]
    dictionary20generatortype["Dictionary20GeneratorType"]
    referencestype[["ReferencesType"]]
    dictionary20itemtype["Dictionary20ItemType"]
    dictionary20notestype[["Dictionary20NotesType"]]
    dictionary20texttype[["Dictionary20TextType"]]
    digestmethodtype[["DigestMethodType"]]
    distinguishednameelementtype[["DistinguishedNameElementType"]]
    documentrootelementtype[["DocumentRootElementType"]]
    elementmapitemtype[["ElementMapItemType"]]
    elementmaptype["ElementMapType"]
    emailaddress[["EmailAddress"]]
    emailaddresselementtype[["EmailAddressElementType"]]
    endfunctiontype[["EndFunctionType"]]
    endorsementlinecodeelementtype[["EndorsementLineCodeElementType"]]
    entitycomplexbasetype[["EntityComplexBaseType"]]
    entitysimplebasetype[["EntitySimpleBaseType"]]
    entityobjectanysimpletype["EntityObjectAnySimpleType"]
    entityobjectbinarytype["EntityObjectBinaryType"]
    entityobjectbooltype["EntityObjectBoolType"]
    entityobjectfieldtype[["EntityObjectFieldType"]]
    entityobjectfloattype["EntityObjectFloatType"]
    entityobjectipaddressstringtype["EntityObjectIPAddressStringType"]
    entityobjectipaddresstype["EntityObjectIPAddressType"]
    entityobjectinttype["EntityObjectIntType"]
    entityobjectrecordtype["EntityObjectRecordType"]
    entityobjectstringtype["EntityObjectStringType"]
    entityobjectversiontype["EntityObjectVersionType"]
    entitystatesimplebasetype["EntityStateSimpleBaseType"]
    entitystateanysimpletype["EntityStateAnySimpleType"]
    entitystatebinarytype["EntityStateBinaryType"]
    entitystatebooltype["EntityStateBoolType"]
    entitystatecomplexbasetype["EntityStateComplexBaseType"]
    entitystatedebianevrstringtype["EntityStateDebianEVRStringType"]
    entitystateevrstringtype["EntityStateEVRStringType"]
    entitystatefieldtype[["EntityStateFieldType"]]
    entitystatefilesetrevisiontype["EntityStateFileSetRevisionType"]
    entitystatefloattype["EntityStateFloatType"]
    entitystateiosversiontype["EntityStateIOSVersionType"]
    entitystateipaddressstringtype["EntityStateIPAddressStringType"]
    entitystateipaddresstype["EntityStateIPAddressType"]
    entitystateinttype["EntityStateIntType"]
    entitystaterecordtype["EntityStateRecordType"]
    entitystatestringtype["EntityStateStringType"]
    entitystateversiontype["EntityStateVersionType"]
    escaperegexfunctiontype[["EscapeRegexFunctionType"]]
    extended[["Extended"]]
    extendedinfoelementtype[["ExtendedInfoElementType"]]
    extendedinformationelementtype[["ExtendedInformationElementType"]]
    extendedinfoselementtype[["ExtendedInfosElementType"]]
    externalvariable["ExternalVariable"]
    possiblevaluetype[["PossibleValueType"]]
    possiblerestrictiontype[["PossibleRestrictionType"]]
    externalvariableelementtype["ExternalVariableElementType"]
    factreftype[["FactRefType"]]
    facttype[["FactType"]]
    filter["Filter"]
    filterelementtype["FilterElementType"]
    firmnameelementtype[["FirmNameElementType"]]
    firmtype["FirmType"]
    firstnameelementtype[["FirstNameElementType"]]
    idreftype[["IdrefType"]]
    subtype["SubType"]
    htmltextwithsubtype["HtmlTextWithSubType"]
    fixtexttype["FixTextType"]
    instancefixtype[["InstanceFixType"]]
    fixtype["FixType"]
    formernameelementtype["FormerNameElementType"]
    fqdn[["Fqdn"]]
    fqdnelementtype[["FqdnElementType"]]
    function[["Function"]]
    generalsuffixelementtype[["GeneralSuffixElementType"]]
    generationidentifierelementtype[["GenerationIdentifierElementType"]]
    globtoregexfunctiontype[["GlobToRegexFunctionType"]]
    selectableitemtype["SelectableItemType"]
    grouptype["GroupType"]
    hostelementtype[["HostElementType"]]
    hostnameelementtype["HostnameElementType"]
    htmltexttype[["HtmlTextType"]]
    identtype[["IdentType"]]
    identitytype[["IdentityType"]]
    idreflisttype[["IdrefListType"]]
    installationidelementtype[["InstallationIdElementType"]]
    instancenameelementtype[["InstanceNameElementType"]]
    instanceresulttype[["InstanceResultType"]]
    ipaddresstype[["IpAddressType"]]
    ipnetrangeelementtype["IpNetRangeElementType"]
    ipv4elementtype["IpV4ElementType"]
    ipv6elementtype["IpV6ElementType"]
    jointpersonname[["JointPersonName"]]
    jointpersonnameelementtype[["JointPersonNameElementType"]]
    keyinfotype[["KeyInfoType"]]
    keylinecodeelementtype[["KeyLineCodeElementType"]]
    keyvaluetype[["KeyValueType"]]
    knownaselementtype["KnownAsElementType"]
    language20texttype[["Language20TextType"]]
    largemailuseridentifierelementtype[["LargeMailUserIdentifierElementType"]]
    largemailusernameelementtype[["LargeMailUserNameElementType"]]
    largemailusertype[["LargeMailUserType"]]
    lastnameelementtype[["LastNameElementType"]]
    licenseelementtype[["LicenseElementType"]]
    listtype[["ListType"]]
    literalcomponenttype[["LiteralComponentType"]]
    localvariable["LocalVariable"]
    localvariableelementtype["LocalVariableElementType"]
    localeelementtype["LocaleElementType"]
    locality[["Locality"]]
    localityelementtype["LocalityElementType"]
    localitynameelementtype[["LocalityNameElementType"]]
    locationpoint[["LocationPoint"]]
    locationpointelementtype[["LocationPointElementType"]]
    locationregion[["LocationRegion"]]
    locationregionelementtype[["LocationRegionElementType"]]
    locations[["Locations"]]
    locationselementtype[["LocationsElementType"]]
    locatortype[["LocatorType"]]
    logicaltesttype[["LogicalTestType"]]
    macaddresselementtype["MacAddressElementType"]
    mailstopnameelementtype[["MailStopNameElementType"]]
    mailstopnumberelementtype[["MailStopNumberElementType"]]
    manifesttype[["ManifestType"]]
    middlenameelementtype[["MiddleNameElementType"]]
    model[["Model"]]
    modelelementtype[["ModelElementType"]]
    motherboardguidelementtype[["MotherboardGuidElementType"]]
    namedetailselementtype["NameDetailsElementType"]
    namelinetype[["NameLineType"]]
    nameprefixelementtype[["NamePrefixElementType"]]
    networkinterfacetype["NetworkInterfaceType"]
    networknameelementtype[["NetworkNameElementType"]]
    networktype["NetworkType"]
    notes["Notes"]
    noteselementtype["NotesElementType"]
    noticetype[["NoticeType"]]
    ns09xmldsigobjecttype[["Ns09XmldsigObjectType"]]
    ns09xmldsigreferencetype[["Ns09XmldsigReferenceType"]]
    ns09xmldsigsignaturetype[["Ns09XmldsigSignatureType"]]
    oasisnamestcciqxnamedetails["OasisNamesTcCiqXNameDetails"]
    oasisnamestcciqxorganisationnamedetails["OasisNamesTcCiqXOrganisationNameDetails"]
    oasisnamestcciqxpersonname["OasisNamesTcCiqXPersonName"]
    objectcomponenttype[["ObjectComponentType"]]
    objectref[["ObjectRef"]]
    objectrefelementtype[["ObjectRefElementType"]]
    objectreftype[["ObjectRefType"]]
    objectstype[["ObjectsType"]]
    organisationformernameelementtype["OrganisationFormerNameElementType"]
    organisationknownaselementtype["OrganisationKnownAsElementType"]
    organisationnamedetailselementtype["OrganisationNameDetailsElementType"]
    organisationnameelementtype[["OrganisationNameElementType"]]
    organisationtypeelementtype[["OrganisationTypeElementType"]]
    organizationtype["OrganizationType"]
    othernameelementtype[["OtherNameElementType"]]
    ovaldefinitions[["OvalDefinitions"]]
    teststype[["TestsType"]]
    statestype[["StatesType"]]
    variablestype[["VariablesType"]]
    ovaldefinitionselementtype["OvalDefinitionsElementType"]
    ovalmitreorgovalgeneratortype[["OvalMitreOrgOvalGeneratorType"]]
    ovalmitreorgovalmessagetype[["OvalMitreOrgOvalMessageType"]]
    ovalmitreorgovalmetadatatype[["OvalMitreOrgOvalMetadataType"]]
    ovalmitreorgovalnotestype[["OvalMitreOrgOvalNotesType"]]
    ovalmitreorgovalobjecttype[["OvalMitreOrgOvalObjectType"]]
    ovalmitreorgovalreferencetype[["OvalMitreOrgOvalReferenceType"]]
    ovalmitreorgovalvaluetype[["OvalMitreOrgOvalValueType"]]
    overridetype[["OverrideType"]]
    overrideablecpe2idreftype["OverrideableCPE2idrefType"]
    pgpdatatype[["PGPDataType"]]
    paramtype[["ParamType"]]
    personnameelementtype["PersonNameElementType"]
    persontype["PersonType"]
    plaintexttype[["PlainTextType"]]
    platformspecification[["PlatformSpecification"]]
    platformspecificationelementtype[["PlatformSpecificationElementType"]]
    platformtype["PlatformType"]
    portelementtype["PortElementType"]
    portrangeelementtype[["PortRangeElementType"]]
    postbox[["PostBox"]]
    postboxelementtype["PostBoxElementType"]
    postboxnumberelementtype[["PostBoxNumberElementType"]]
    postboxnumberextensionelementtype[["PostBoxNumberExtensionElementType"]]
    postboxnumberprefixelementtype[["PostBoxNumberPrefixElementType"]]
    postboxnumbersuffixelementtype[["PostBoxNumberSuffixElementType"]]
    postoffice[["PostOffice"]]
    postalroutetype[["PostalRouteType"]]
    postofficeelementtype["PostOfficeElementType"]
    postofficenameelementtype[["PostOfficeNameElementType"]]
    postofficenumberelementtype[["PostOfficeNumberElementType"]]
    posttownelementtype[["PostTownElementType"]]
    posttownnameelementtype[["PostTownNameElementType"]]
    posttownsuffixelementtype[["PostTownSuffixElementType"]]
    postalcode[["PostalCode"]]
    postalcodeelementtype[["PostalCodeElementType"]]
    postalcodenumberelementtype[["PostalCodeNumberElementType"]]
    postalcodenumberextensionelementtype[["PostalCodeNumberExtensionElementType"]]
    postalroutenameelementtype[["PostalRouteNameElementType"]]
    postalroutenumberelementtype[["PostalRouteNumberElementType"]]
    postalserviceelementselementtype[["PostalServiceElementsElementType"]]
    precedingtitleelementtype[["PrecedingTitleElementType"]]
    premise[["Premise"]]
    premiseelementtype["PremiseElementType"]
    premiselocationelementtype[["PremiseLocationElementType"]]
    premisenameelementtype[["PremiseNameElementType"]]
    premisenumber[["PremiseNumber"]]
    premisenumberelementtype[["PremiseNumberElementType"]]
    premisenumberprefix[["PremiseNumberPrefix"]]
    premisenumberprefixelementtype[["PremiseNumberPrefixElementType"]]
    premisenumberrangeelementtype[["PremiseNumberRangeElementType"]]
    premisenumberrangefromelementtype[["PremiseNumberRangeFromElementType"]]
    premisenumberrangetoelementtype[["PremiseNumberRangeToElementType"]]
    premisenumbersuffix[["PremiseNumberSuffix"]]
    premisenumbersuffixelementtype[["PremiseNumberSuffixElementType"]]
    profilenotetype["ProfileNoteType"]
    profilerefineruletype[["ProfileRefineRuleType"]]
    profilerefinevaluetype[["ProfileRefineValueType"]]
    profileselecttype[["ProfileSelectType"]]
    profilesetcomplexvaluetype["ProfileSetComplexValueType"]
    profilesetvaluetype[["ProfileSetValueType"]]
    profiletype["ProfileType"]
    protocolelementtype[["ProtocolElementType"]]
    rsakeyvaluetype[["RSAKeyValueType"]]
    referenceelementtype[["ReferenceElementType"]]
    regexcapturefunctiontype[["RegexCaptureFunctionType"]]
    relationshiptype[["RelationshipType"]]
    relationshipselementtype[["RelationshipsElementType"]]
    remoteresource[["RemoteResource"]]
    remoteresourceelementtype[["RemoteResourceElementType"]]
    reportrequesttype[["ReportRequestType"]]
    reportrequestselementtype[["ReportRequestsElementType"]]
    reporttype[["ReportType"]]
    reportselementtype[["ReportsElementType"]]
    resourcetype[["ResourceType"]]
    restrictiontype[["RestrictionType"]]
    retrievalmethodtype[["RetrievalMethodType"]]
    ruleresulttype[["RuleResultType"]]
    ruletype["RuleType"]
    spkidatatype[["SPKIDataType"]]
    scoretype[["ScoreType"]]
    selchoicestype["SelChoicesType"]
    selcomplexvaluetype["SelComplexValueType"]
    selnumtype[["SelNumType"]]
    selstringtype[["SelStringType"]]
    servicetype["ServiceType"]
    set[["Set"]]
    setelementtype[["SetElementType"]]
    signaturemethodtype[["SignatureMethodType"]]
    signaturepropertiestype[["SignaturePropertiesType"]]
    signaturepropertytype[["SignaturePropertyType"]]
    signaturevaluetype[["SignatureValueType"]]
    signedinfotype[["SignedInfoType"]]
    simple[["Simple"]]
    softwaretype["SoftwareType"]
    sortingcodeelementtype[["SortingCodeElementType"]]
    splitfunctiontype[["SplitFunctionType"]]
    statereftype[["StateRefType"]]
    statetype[["StateType"]]
    status["Status"]
    statuselementtype["StatusElementType"]
    subadministrativeareaelementtype[["SubAdministrativeAreaElementType"]]
    subadministrativeareanameelementtype[["SubAdministrativeAreaNameElementType"]]
    subpremiselocationelementtype[["SubPremiseLocationElementType"]]
    subpremisenameelementtype[["SubPremiseNameElementType"]]
    subpremisenumberelementtype[["SubPremiseNumberElementType"]]
    subpremisenumberprefixelementtype[["SubPremiseNumberPrefixElementType"]]
    subpremisenumbersuffixelementtype[["SubPremiseNumberSuffixElementType"]]
    subpremisetype["SubPremiseType"]
    substringfunctiontype[["SubstringFunctionType"]]
    suffixelementtype[["SuffixElementType"]]
    supplementarypostalservicedataelementtype[["SupplementaryPostalServiceDataElementType"]]
    syntheticid[["SyntheticId"]]
    syntheticidelementtype[["SyntheticIdElementType"]]
    systemnameelementtype[["SystemNameElementType"]]
    systemtype["SystemType"]
    tailoringbenchmarkreferencetype["TailoringBenchmarkReferenceType"]
    tailoringreferencetype[["TailoringReferenceType"]]
    tailoringversiontype[["TailoringVersionType"]]
    tailoringtype["TailoringType"]
    targetfactstype[["TargetFactsType"]]
    targetidreftype[["TargetIdRefType"]]
    telephonenumber["TelephoneNumber"]
    telephonenumberelementtype["TelephoneNumberElementType"]
    testresulttype["TestResultType"]
    testtype[["TestType"]]
    textwithsubtype[["TextWithSubType"]]
    thoroughfare[["Thoroughfare"]]
    thoroughfareelementtype["ThoroughfareElementType"]
    thoroughfarenametype[["ThoroughfareNameType"]]
    thoroughfarenumber[["ThoroughfareNumber"]]
    thoroughfarenumberelementtype[["ThoroughfareNumberElementType"]]
    thoroughfarenumberfromelementtype[["ThoroughfareNumberFromElementType"]]
    thoroughfarenumberprefix[["ThoroughfareNumberPrefix"]]
    thoroughfarenumberprefixelementtype[["ThoroughfareNumberPrefixElementType"]]
    thoroughfarenumberrangeelementtype[["ThoroughfareNumberRangeElementType"]]
    thoroughfarenumbersuffix[["ThoroughfareNumberSuffix"]]
    thoroughfarenumbersuffixelementtype[["ThoroughfareNumberSuffixElementType"]]
    thoroughfarenumbertoelementtype[["ThoroughfareNumberToElementType"]]
    timedifferencefunctiontype[["TimeDifferenceFunctionType"]]
    titleelementtype[["TitleElementType"]]
    titleelttype[["TitleEltType"]]
    transformtype[["TransformType"]]
    transformstype[["TransformsType"]]
    uniquefunctiontype[["UniqueFunctionType"]]
    urireftype[["UriRefType"]]
    urlelementtype[["UrlElementType"]]
    variablecomponenttype[["VariableComponentType"]]
    versionelementtype[["VersionElementType"]]
    warningtype["WarningType"]
    websitetype["WebsiteType"]
    websiteurl[["WebsiteUrl"]]
    websiteurlelementtype[["WebsiteUrlElementType"]]
    x509datatype[["X509DataType"]]
    x509issuerserialtype[["X509IssuerSerialType"]]
    xal[["XAL"]]
    xalelementtype[["XALElementType"]]
    xnl[["XNL"]]
    xnlelementtype[["XNLElementType"]]
    xccdf12checktype["Xccdf12CheckType"]
    xccdf12itemtype["Xccdf12ItemType"]
    xccdf12messagetype[["Xccdf12MessageType"]]
    xccdf12metadatatype[["Xccdf12MetadataType"]]
    xccdf12referencetype[["Xccdf12ReferenceType"]]
    xccdf12signaturetype[["Xccdf12SignatureType"]]
    xccdf12texttype[["Xccdf12TextType"]]
    xccdf12valuetype["Xccdf12ValueType"]

    relationshipscontainertype --> assetstype
    relationshipscontainertype -.->|extends| assetstype
    assetstype --> assetidentificationtype
    assetstype -.->|extends| assetidentificationtype
    relationshipscontainertype --> assetreportcollection
    relationshipscontainertype -.->|extends| assetreportcollection
    relationshipscontainertype --> assetreportcollectionelementtype
    relationshipscontainertype -.->|extends| assetreportcollectionelementtype
    versiontype --> benchmarkelementtype
    cidrtype --> cidrelementtype
    cidrtype -.->|extends| cidrelementtype
    assettype --> itassettype
    assettype -.->|extends| itassettype
    itassettype --> circuittype
    itassettype -.->|extends| circuittype
    complexchecktype --> complexchecktype
    itassettype --> computingdevicetype
    itassettype -.->|extends| computingdevicetype
    variabletype --> constantvariable
    variabletype -.->|extends| constantvariable
    variabletype --> constantvariableelementtype
    variabletype -.->|extends| constantvariableelementtype
    cpetype --> cpe
    cpetype -.->|extends| cpe
    cpetype --> cpeelementtype
    cpetype -.->|extends| cpeelementtype
    criteriatype --> criteriatype
    criteriontype --> criteriatype
    extenddefinitiontype --> criteriatype
    assettype --> datatype
    assettype -.->|extends| datatype
    itassettype --> databasetype
    itassettype -.->|extends| databasetype
    criteriatype --> definitiontype
    mailstoptype --> departmentelementtype
    namedetails --> dependencynameelementtype
    namedetails -.->|extends| dependencynameelementtype
    dependentlocalitytype --> dependentlocalitytype
    thoroughfarepredirectiontype --> dependentthoroughfareelementtype
    thoroughfareleadingtypetype --> dependentthoroughfareelementtype
    thoroughfaretrailingtypetype --> dependentthoroughfareelementtype
    thoroughfarepostdirectiontype --> dependentthoroughfareelementtype
    schemaversionpattern --> schemaversiontype
    schemaversionpattern -.->|extends| schemaversiontype
    schemaversiontype --> dictionary20generatortype
    referencestype --> dictionary20itemtype
    elementmapitemtype --> elementmaptype
    elementmapitemtype --> elementmaptype
    elementmapitemtype --> elementmaptype
    elementmapitemtype --> elementmaptype
    entitysimplebasetype --> entityobjectanysimpletype
    entitysimplebasetype -.->|extends| entityobjectanysimpletype
    entitysimplebasetype --> entityobjectbinarytype
    entitysimplebasetype -.->|extends| entityobjectbinarytype
    entitysimplebasetype --> entityobjectbooltype
    entitysimplebasetype -.->|extends| entityobjectbooltype
    entitysimplebasetype --> entityobjectfloattype
    entitysimplebasetype -.->|extends| entityobjectfloattype
    entitysimplebasetype --> entityobjectipaddressstringtype
    entitysimplebasetype -.->|extends| entityobjectipaddressstringtype
    entitysimplebasetype --> entityobjectipaddresstype
    entitysimplebasetype -.->|extends| entityobjectipaddresstype
    entitysimplebasetype --> entityobjectinttype
    entitysimplebasetype -.->|extends| entityobjectinttype
    entitycomplexbasetype --> entityobjectrecordtype
    entitycomplexbasetype -.->|extends| entityobjectrecordtype
    entitysimplebasetype --> entityobjectstringtype
    entitysimplebasetype -.->|extends| entityobjectstringtype
    entitysimplebasetype --> entityobjectversiontype
    entitysimplebasetype -.->|extends| entityobjectversiontype
    entitysimplebasetype --> entitystatesimplebasetype
    entitysimplebasetype -.->|extends| entitystatesimplebasetype
    entitystatesimplebasetype --> entitystateanysimpletype
    entitystatesimplebasetype -.->|extends| entitystateanysimpletype
    entitystatesimplebasetype --> entitystatebinarytype
    entitystatesimplebasetype -.->|extends| entitystatebinarytype
    entitystatesimplebasetype --> entitystatebooltype
    entitystatesimplebasetype -.->|extends| entitystatebooltype
    entitycomplexbasetype --> entitystatecomplexbasetype
    entitycomplexbasetype -.->|extends| entitystatecomplexbasetype
    entitystatesimplebasetype --> entitystatedebianevrstringtype
    entitystatesimplebasetype -.->|extends| entitystatedebianevrstringtype
    entitystatesimplebasetype --> entitystateevrstringtype
    entitystatesimplebasetype -.->|extends| entitystateevrstringtype
    entitystatesimplebasetype --> entitystatefilesetrevisiontype
    entitystatesimplebasetype -.->|extends| entitystatefilesetrevisiontype
    entitystatesimplebasetype --> entitystatefloattype
    entitystatesimplebasetype -.->|extends| entitystatefloattype
    entitystatesimplebasetype --> entitystateiosversiontype
    entitystatesimplebasetype -.->|extends| entitystateiosversiontype
    entitystatesimplebasetype --> entitystateipaddressstringtype
    entitystatesimplebasetype -.->|extends| entitystateipaddressstringtype
    entitystatesimplebasetype --> entitystateipaddresstype
    entitystatesimplebasetype -.->|extends| entitystateipaddresstype
    entitystatesimplebasetype --> entitystateinttype
    entitystatesimplebasetype -.->|extends| entitystateinttype
    entitystatecomplexbasetype --> entitystaterecordtype
    entitystatecomplexbasetype -.->|extends| entitystaterecordtype
    entitystatesimplebasetype --> entitystatestringtype
    entitystatesimplebasetype -.->|extends| entitystatestringtype
    entitystatesimplebasetype --> entitystateversiontype
    entitystatesimplebasetype -.->|extends| entitystateversiontype
    variabletype --> externalvariable
    variabletype -.->|extends| externalvariable
    variabletype --> externalvariableelementtype
    possiblevaluetype --> externalvariableelementtype
    possiblerestrictiontype --> externalvariableelementtype
    variabletype -.->|extends| externalvariableelementtype
    stateidpattern --> filter
    stateidpattern -.->|extends| filter
    stateidpattern --> filterelementtype
    stateidpattern -.->|extends| filterelementtype
    mailstoptype --> firmtype
    idreftype --> subtype
    idreftype -.->|extends| subtype
    subtype --> htmltextwithsubtype
    htmltextwithsubtype --> fixtexttype
    htmltextwithsubtype -.->|extends| fixtexttype
    subtype --> fixtype
    instancefixtype --> fixtype
    personname --> formernameelementtype
    personname -.->|extends| formernameelementtype
    itemtype --> selectableitemtype
    itemtype -.->|extends| selectableitemtype
    selectableitemtype --> grouptype
    selectableitemtype -.->|extends| grouptype
    hostnametype --> hostnameelementtype
    hostnametype -.->|extends| hostnameelementtype
    ipaddresstype --> ipnetrangeelementtype
    ipaddresstype --> ipnetrangeelementtype
    ipv4type --> ipv4elementtype
    ipv4type -.->|extends| ipv4elementtype
    ipv6type --> ipv6elementtype
    ipv6type -.->|extends| ipv6elementtype
    personname --> knownaselementtype
    personname -.->|extends| knownaselementtype
    variabletype --> localvariable
    variabletype -.->|extends| localvariable
    variabletype --> localvariableelementtype
    variabletype -.->|extends| localvariableelementtype
    localetype --> localeelementtype
    localetype -.->|extends| localeelementtype
    dependentlocalitytype --> localityelementtype
    macaddresstype --> macaddresselementtype
    macaddresstype -.->|extends| macaddresselementtype
    namedetails --> namedetailselementtype
    namedetails -.->|extends| namedetailselementtype
    ipaddresstype --> networkinterfacetype
    ipaddresstype --> networkinterfacetype
    itassettype --> networktype
    itassettype -.->|extends| networktype
    notestype --> notes
    notestype -.->|extends| notes
    notestype --> noteselementtype
    notestype -.->|extends| noteselementtype
    namedetails --> oasisnamestcciqxnamedetails
    namedetails -.->|extends| oasisnamestcciqxnamedetails
    organisationnamedetails --> oasisnamestcciqxorganisationnamedetails
    organisationnamedetails -.->|extends| oasisnamestcciqxorganisationnamedetails
    personname --> oasisnamestcciqxpersonname
    personname -.->|extends| oasisnamestcciqxpersonname
    organisationnamedetails --> organisationformernameelementtype
    organisationnamedetails -.->|extends| organisationformernameelementtype
    organisationnamedetails --> organisationknownaselementtype
    organisationnamedetails -.->|extends| organisationknownaselementtype
    organisationnamedetails --> organisationnamedetailselementtype
    organisationnamedetails -.->|extends| organisationnamedetailselementtype
    assettype --> organizationtype
    assettype -.->|extends| organizationtype
    definitionstype --> ovaldefinitionselementtype
    teststype --> ovaldefinitionselementtype
    objectstype --> ovaldefinitionselementtype
    statestype --> ovaldefinitionselementtype
    variablestype --> ovaldefinitionselementtype
    cpe2idreftype --> overrideablecpe2idreftype
    cpe2idreftype -.->|extends| overrideablecpe2idreftype
    personname --> personnameelementtype
    personname -.->|extends| personnameelementtype
    assettype --> persontype
    assettype -.->|extends| persontype
    logicaltesttype --> platformtype
    porttype --> portelementtype
    porttype -.->|extends| portelementtype
    firmtype --> postboxelementtype
    postalroutetype --> postofficeelementtype
    mailstoptype --> premiseelementtype
    subtype --> profilenotetype
    complexvaluetype --> profilesetcomplexvaluetype
    complexvaluetype -.->|extends| profilesetcomplexvaluetype
    versiontype --> profiletype
    selectableitemtype --> ruletype
    selectableitemtype -.->|extends| ruletype
    complexvaluetype --> selchoicestype
    complexvaluetype --> selcomplexvaluetype
    complexvaluetype -.->|extends| selcomplexvaluetype
    itassettype --> servicetype
    itassettype -.->|extends| servicetype
    itassettype --> softwaretype
    itassettype -.->|extends| softwaretype
    statustype --> status
    statustype -.->|extends| status
    statustype --> statuselementtype
    statustype -.->|extends| statuselementtype
    firmtype --> subpremisetype
    mailstoptype --> subpremisetype
    subpremisetype --> subpremisetype
    itassettype --> systemtype
    itassettype -.->|extends| systemtype
    benchmarkreferencetype --> tailoringbenchmarkreferencetype
    benchmarkreferencetype -.->|extends| tailoringbenchmarkreferencetype
    tailoringbenchmarkreferencetype --> tailoringtype
    tailoringversiontype --> tailoringtype
    telephonenumbertype --> telephonenumber
    telephonenumbertype -.->|extends| telephonenumber
    telephonenumbertype --> telephonenumberelementtype
    telephonenumbertype -.->|extends| telephonenumberelementtype
    benchmarkreferencetype --> testresulttype
    tailoringreferencetype --> testresulttype
    identitytype --> testresulttype
    idreftype --> testresulttype
    targetfactstype --> testresulttype
    thoroughfarepredirectiontype --> thoroughfareelementtype
    thoroughfareleadingtypetype --> thoroughfareelementtype
    thoroughfaretrailingtypetype --> thoroughfareelementtype
    thoroughfarepostdirectiontype --> thoroughfareelementtype
    htmltextwithsubtype --> warningtype
    htmltextwithsubtype -.->|extends| warningtype
    itassettype --> websitetype
    itassettype -.->|extends| websitetype
    checkcontenttype --> xccdf12checktype
    versiontype --> xccdf12itemtype
    itemtype --> xccdf12valuetype
    itemtype -.->|extends| xccdf12valuetype
```

## Table Index

| Table | Columns | Foreign Keys | Parent |
|-------|---------|--------------|--------|
| AddressDetails | 11 | 0 | - |
| AddressElementType | 2 | 0 | - |
| AddressIdentifierElementType | 3 | 0 | - |
| AddressLatitudeDirectionElementType | 2 | 0 | - |
| AddressLatitudeElementType | 2 | 0 | - |
| AddressLine | 2 | 0 | - |
| AddressLineElementType | 2 | 0 | - |
| AddressLinesType | 1 | 0 | - |
| AddressLongitudeDirectionElementType | 2 | 0 | - |
| AddressLongitudeElementType | 2 | 0 | - |
| AddresseeIndicatorElementType | 2 | 0 | - |
| AdministrativeArea | 6 | 0 | - |
| AdministrativeAreaElementType | 6 | 0 | - |
| AdministrativeAreaNameElementType | 2 | 0 | - |
| AffectedType | 4 | 0 | - |
| AliasElementType | 4 | 0 | - |
| ArcType | 1 | 0 | - |
| ArithmeticFunctionType | 2 | 0 | - |
| AssetIdentificationAssetElementType | 2 | 0 | - |
| RelationshipsContainerType | 2 | 0 | - |
| AssetsType | 3 | 1 | RelationshipsContainerType |
| AssetIdentificationType | 3 | 1 | AssetsType |
| AssetReportCollection | 7 | 1 | RelationshipsContainerType |
| AssetReportCollectionElementType | 7 | 1 | RelationshipsContainerType |
| AssetReportingFormAssetElementType | 2 | 0 | - |
| AssetType | 2 | 0 | - |
| AssetsElementType | 2 | 0 | - |
| BarcodeElementType | 2 | 0 | - |
| BeginFunctionType | 2 | 0 | - |
| Benchmark | 7 | 0 | - |
| VersionType | 3 | 0 | - |
| BenchmarkElementType | 7 | 1 | - |
| BenchmarkReferenceType | 3 | 0 | - |
| BirthdateElementType | 1 | 0 | - |
| BuildingNameType | 3 | 0 | - |
| CPE2idrefType | 2 | 0 | - |
| CanonicalizationMethodType | 2 | 0 | - |
| CheckContentRefType | 3 | 0 | - |
| CheckContentType | 1 | 0 | - |
| CheckExportType | 3 | 0 | - |
| CheckImportType | 3 | 0 | - |
| CidrElementType | 2 | 1 | CidrType |
| CircuitNameElementType | 1 | 0 | - |
| ItAssetType | 2 | 1 | AssetType |
| CircuitType | 3 | 1 | ItAssetType |
| ComplexCheckType | 5 | 1 | - |
| ComplexValueType | 2 | 0 | - |
| ComputingDeviceType | 6 | 1 | ItAssetType |
| ConcatFunctionType | 1 | 0 | - |
| ConnectionsElementType | 1 | 0 | - |
| VariableType | 6 | 0 | - |
| ConstantVariable | 2 | 1 | VariableType |
| ConstantVariableElementType | 2 | 1 | VariableType |
| ContentElementType | 3 | 0 | - |
| ContentElementType1 | 1 | 0 | - |
| CountFunctionType | 1 | 0 | - |
| CountryElementType | 2 | 0 | - |
| CountryName | 2 | 0 | - |
| CountryNameCodeElementType | 2 | 0 | - |
| CountryNameElementType | 2 | 0 | - |
| Cpe | 2 | 1 | CpeType |
| CpeElementType | 2 | 1 | CpeType |
| CriterionType | 5 | 0 | - |
| ExtendDefinitionType | 5 | 0 | - |
| CriteriaType | 8 | 3 | - |
| DSAKeyValueType | 8 | 0 | - |
| DataType | 2 | 1 | AssetType |
| DatabaseType | 3 | 1 | ItAssetType |
| DcStatusType | 1 | 0 | - |
| DefinitionType | 7 | 1 | - |
| DefinitionsType | 1 | 0 | - |
| Department | 4 | 0 | - |
| MailStopType | 4 | 0 | - |
| DepartmentElementType | 4 | 1 | - |
| DepartmentNameElementType | 2 | 0 | - |
| DependencyNameElementType | 3 | 1 | NameDetails |
| DependentLocalityNameElementType | 2 | 0 | - |
| DependentLocalityNumberElementType | 2 | 0 | - |
| DependentLocalityType | 10 | 1 | - |
| ThoroughfarePreDirectionType | 2 | 0 | - |
| ThoroughfareLeadingTypeType | 2 | 0 | - |
| ThoroughfareTrailingTypeType | 2 | 0 | - |
| ThoroughfarePostDirectionType | 2 | 0 | - |
| DependentThoroughfareElementType | 6 | 4 | - |
| DeprecatedInfoType | 4 | 0 | - |
| Dictionary20CheckType | 3 | 0 | - |
| SchemaVersionType | 3 | 1 | SchemaVersionPattern |
| Dictionary20GeneratorType | 5 | 1 | - |
| ReferencesType | 2 | 0 | - |
| Dictionary20ItemType | 6 | 1 | - |
| Dictionary20NotesType | 2 | 0 | - |
| Dictionary20TextType | 1 | 0 | - |
| DigestMethodType | 2 | 0 | - |
| DistinguishedNameElementType | 1 | 0 | - |
| DocumentRootElementType | 1 | 0 | - |
| ElementMapItemType | 2 | 0 | - |
| ElementMapType | 5 | 4 | - |
| EmailAddress | 1 | 0 | - |
| EmailAddressElementType | 1 | 0 | - |
| EndFunctionType | 2 | 0 | - |
| EndorsementLineCodeElementType | 2 | 0 | - |
| EntityComplexBaseType | 1 | 0 | - |
| EntitySimpleBaseType | 1 | 0 | - |
| EntityObjectAnySimpleType | 3 | 1 | EntitySimpleBaseType |
| EntityObjectBinaryType | 3 | 1 | EntitySimpleBaseType |
| EntityObjectBoolType | 3 | 1 | EntitySimpleBaseType |
| EntityObjectFieldType | 3 | 0 | - |
| EntityObjectFloatType | 3 | 1 | EntitySimpleBaseType |
| EntityObjectIPAddressStringType | 3 | 1 | EntitySimpleBaseType |
| EntityObjectIPAddressType | 3 | 1 | EntitySimpleBaseType |
| EntityObjectIntType | 3 | 1 | EntitySimpleBaseType |
| EntityObjectRecordType | 2 | 1 | EntityComplexBaseType |
| EntityObjectStringType | 3 | 1 | EntitySimpleBaseType |
| EntityObjectVersionType | 3 | 1 | EntitySimpleBaseType |
| EntityStateSimpleBaseType | 4 | 1 | EntitySimpleBaseType |
| EntityStateAnySimpleType | 3 | 1 | EntityStateSimpleBaseType |
| EntityStateBinaryType | 3 | 1 | EntityStateSimpleBaseType |
| EntityStateBoolType | 3 | 1 | EntityStateSimpleBaseType |
| EntityStateComplexBaseType | 4 | 1 | EntityComplexBaseType |
| EntityStateDebianEVRStringType | 3 | 1 | EntityStateSimpleBaseType |
| EntityStateEVRStringType | 3 | 1 | EntityStateSimpleBaseType |
| EntityStateFieldType | 3 | 0 | - |
| EntityStateFileSetRevisionType | 3 | 1 | EntityStateSimpleBaseType |
| EntityStateFloatType | 3 | 1 | EntityStateSimpleBaseType |
| EntityStateIOSVersionType | 3 | 1 | EntityStateSimpleBaseType |
| EntityStateIPAddressStringType | 3 | 1 | EntityStateSimpleBaseType |
| EntityStateIPAddressType | 3 | 1 | EntityStateSimpleBaseType |
| EntityStateIntType | 3 | 1 | EntityStateSimpleBaseType |
| EntityStateRecordType | 2 | 1 | EntityStateComplexBaseType |
| EntityStateStringType | 3 | 1 | EntityStateSimpleBaseType |
| EntityStateVersionType | 3 | 1 | EntityStateSimpleBaseType |
| EscapeRegexFunctionType | 1 | 0 | - |
| Extended | 1 | 0 | - |
| ExtendedInfoElementType | 2 | 0 | - |
| ExtendedInformationElementType | 1 | 0 | - |
| ExtendedInfosElementType | 2 | 0 | - |
| ExternalVariable | 4 | 1 | VariableType |
| PossibleValueType | 2 | 0 | - |
| PossibleRestrictionType | 3 | 0 | - |
| ExternalVariableElementType | 4 | 3 | VariableType |
| FactRefType | 2 | 0 | - |
| FactType | 3 | 0 | - |
| Filter | 3 | 1 | StateIDPattern |
| FilterElementType | 3 | 1 | StateIDPattern |
| FirmNameElementType | 2 | 0 | - |
| FirmType | 4 | 1 | - |
| FirstNameElementType | 4 | 0 | - |
| IdrefType | 2 | 0 | - |
| SubType | 3 | 1 | IdrefType |
| HtmlTextWithSubType | 3 | 1 | - |
| FixTextType | 7 | 1 | HtmlTextWithSubType |
| InstanceFixType | 2 | 0 | - |
| FixType | 10 | 2 | - |
| FormerNameElementType | 4 | 1 | PersonName |
| Fqdn | 1 | 0 | - |
| FqdnElementType | 1 | 0 | - |
| Function | 2 | 0 | - |
| GeneralSuffixElementType | 3 | 0 | - |
| GenerationIdentifierElementType | 3 | 0 | - |
| GlobToRegexFunctionType | 2 | 0 | - |
| SelectableItemType | 4 | 1 | ItemType |
| GroupType | 4 | 1 | SelectableItemType |
| HostElementType | 1 | 0 | - |
| HostnameElementType | 2 | 1 | HostnameType |
| HtmlTextType | 2 | 0 | - |
| IdentType | 2 | 0 | - |
| IdentityType | 3 | 0 | - |
| IdrefListType | 2 | 0 | - |
| InstallationIdElementType | 1 | 0 | - |
| InstanceNameElementType | 1 | 0 | - |
| InstanceResultType | 3 | 0 | - |
| IpAddressType | 3 | 0 | - |
| IpNetRangeElementType | 3 | 2 | - |
| IpV4ElementType | 2 | 1 | Ipv4Type |
| IpV6ElementType | 2 | 1 | Ipv6Type |
| JointPersonName | 3 | 0 | - |
| JointPersonNameElementType | 3 | 0 | - |
| KeyInfoType | 2 | 0 | - |
| KeyLineCodeElementType | 2 | 0 | - |
| KeyValueType | 1 | 0 | - |
| KnownAsElementType | 4 | 1 | PersonName |
| Language20TextType | 1 | 0 | - |
| LargeMailUserIdentifierElementType | 3 | 0 | - |
| LargeMailUserNameElementType | 3 | 0 | - |
| LargeMailUserType | 4 | 0 | - |
| LastNameElementType | 4 | 0 | - |
| LicenseElementType | 1 | 0 | - |
| ListType | 2 | 0 | - |
| LiteralComponentType | 2 | 0 | - |
| LocalVariable | 2 | 1 | VariableType |
| LocalVariableElementType | 2 | 1 | VariableType |
| LocaleElementType | 2 | 1 | LocaleType |
| Locality | 8 | 0 | - |
| LocalityElementType | 8 | 1 | - |
| LocalityNameElementType | 2 | 0 | - |
| LocationPoint | 5 | 0 | - |
| LocationPointElementType | 5 | 0 | - |
| LocationRegion | 1 | 0 | - |
| LocationRegionElementType | 1 | 0 | - |
| Locations | 1 | 0 | - |
| LocationsElementType | 1 | 0 | - |
| LocatorType | 1 | 0 | - |
| LogicalTestType | 3 | 0 | - |
| MacAddressElementType | 2 | 1 | MacAddressType |
| MailStopNameElementType | 2 | 0 | - |
| MailStopNumberElementType | 2 | 0 | - |
| ManifestType | 2 | 0 | - |
| MiddleNameElementType | 4 | 0 | - |
| Model | 2 | 0 | - |
| ModelElementType | 2 | 0 | - |
| MotherboardGuidElementType | 1 | 0 | - |
| NameDetailsElementType | 5 | 1 | NameDetails |
| NameLineType | 4 | 0 | - |
| NamePrefixElementType | 4 | 0 | - |
| NetworkInterfaceType | 5 | 2 | - |
| NetworkNameElementType | 1 | 0 | - |
| NetworkType | 5 | 1 | ItAssetType |
| Notes | 3 | 1 | NotesType |
| NotesElementType | 3 | 1 | NotesType |
| NoticeType | 2 | 0 | - |
| Ns09XmldsigObjectType | 4 | 0 | - |
| Ns09XmldsigReferenceType | 4 | 0 | - |
| Ns09XmldsigSignatureType | 2 | 0 | - |
| OasisNamesTcCiqXNameDetails | 5 | 1 | NameDetails |
| OasisNamesTcCiqXOrganisationNameDetails | 4 | 1 | OrganisationNameDetails |
| OasisNamesTcCiqXPersonName | 4 | 1 | PersonName |
| ObjectComponentType | 4 | 0 | - |
| ObjectRef | 2 | 0 | - |
| ObjectRefElementType | 2 | 0 | - |
| ObjectRefType | 2 | 0 | - |
| ObjectsType | 1 | 0 | - |
| OrganisationFormerNameElementType | 4 | 1 | OrganisationNameDetails |
| OrganisationKnownAsElementType | 4 | 1 | OrganisationNameDetails |
| OrganisationNameDetailsElementType | 4 | 1 | OrganisationNameDetails |
| OrganisationNameElementType | 4 | 0 | - |
| OrganisationTypeElementType | 4 | 0 | - |
| OrganizationType | 2 | 1 | AssetType |
| OtherNameElementType | 4 | 0 | - |
| OvalDefinitions | 7 | 0 | - |
| TestsType | 1 | 0 | - |
| StatesType | 1 | 0 | - |
| VariablesType | 1 | 0 | - |
| OvalDefinitionsElementType | 7 | 5 | - |
| OvalMitreOrgOvalGeneratorType | 4 | 0 | - |
| OvalMitreOrgOvalMessageType | 2 | 0 | - |
| OvalMitreOrgOvalMetadataType | 3 | 0 | - |
| OvalMitreOrgOvalNotesType | 2 | 0 | - |
| OvalMitreOrgOvalObjectType | 5 | 0 | - |
| OvalMitreOrgOvalReferenceType | 4 | 0 | - |
| OvalMitreOrgOvalValueType | 1 | 0 | - |
| OverrideType | 6 | 0 | - |
| OverrideableCPE2idrefType | 3 | 1 | CPE2idrefType |
| PGPDataType | 3 | 0 | - |
| ParamType | 2 | 0 | - |
| PersonNameElementType | 4 | 1 | PersonName |
| PersonType | 3 | 1 | AssetType |
| PlainTextType | 2 | 0 | - |
| PlatformSpecification | 1 | 0 | - |
| PlatformSpecificationElementType | 1 | 0 | - |
| PlatformType | 3 | 1 | - |
| PortElementType | 2 | 1 | PortType |
| PortRangeElementType | 3 | 0 | - |
| PostBox | 8 | 0 | - |
| PostBoxElementType | 8 | 1 | - |
| PostBoxNumberElementType | 1 | 0 | - |
| PostBoxNumberExtensionElementType | 2 | 0 | - |
| PostBoxNumberPrefixElementType | 2 | 0 | - |
| PostBoxNumberSuffixElementType | 2 | 0 | - |
| PostOffice | 6 | 0 | - |
| PostalRouteType | 4 | 0 | - |
| PostOfficeElementType | 6 | 1 | - |
| PostOfficeNameElementType | 2 | 0 | - |
| PostOfficeNumberElementType | 3 | 0 | - |
| PostTownElementType | 4 | 0 | - |
| PostTownNameElementType | 2 | 0 | - |
| PostTownSuffixElementType | 1 | 0 | - |
| PostalCode | 5 | 0 | - |
| PostalCodeElementType | 5 | 0 | - |
| PostalCodeNumberElementType | 2 | 0 | - |
| PostalCodeNumberExtensionElementType | 3 | 0 | - |
| PostalRouteNameElementType | 2 | 0 | - |
| PostalRouteNumberElementType | 1 | 0 | - |
| PostalServiceElementsElementType | 12 | 0 | - |
| PrecedingTitleElementType | 3 | 0 | - |
| Premise | 10 | 0 | - |
| PremiseElementType | 10 | 1 | - |
| PremiseLocationElementType | 1 | 0 | - |
| PremiseNameElementType | 3 | 0 | - |
| PremiseNumber | 6 | 0 | - |
| PremiseNumberElementType | 6 | 0 | - |
| PremiseNumberPrefix | 3 | 0 | - |
| PremiseNumberPrefixElementType | 3 | 0 | - |
| PremiseNumberRangeElementType | 9 | 0 | - |
| PremiseNumberRangeFromElementType | 1 | 0 | - |
| PremiseNumberRangeToElementType | 1 | 0 | - |
| PremiseNumberSuffix | 3 | 0 | - |
| PremiseNumberSuffixElementType | 3 | 0 | - |
| ProfileNoteType | 3 | 1 | - |
| ProfileRefineRuleType | 6 | 0 | - |
| ProfileRefineValueType | 4 | 0 | - |
| ProfileSelectType | 3 | 0 | - |
| ProfileSetComplexValueType | 3 | 1 | ComplexValueType |
| ProfileSetValueType | 2 | 0 | - |
| ProfileType | 13 | 1 | - |
| ProtocolElementType | 1 | 0 | - |
| RSAKeyValueType | 3 | 0 | - |
| ReferenceElementType | 2 | 0 | - |
| RegexCaptureFunctionType | 2 | 0 | - |
| RelationshipType | 5 | 0 | - |
| RelationshipsElementType | 1 | 0 | - |
| RemoteResource | 1 | 0 | - |
| RemoteResourceElementType | 1 | 0 | - |
| ReportRequestType | 3 | 0 | - |
| ReportRequestsElementType | 1 | 0 | - |
| ReportType | 3 | 0 | - |
| ReportsElementType | 1 | 0 | - |
| ResourceType | 1 | 0 | - |
| RestrictionType | 2 | 0 | - |
| RetrievalMethodType | 3 | 0 | - |
| RuleResultType | 9 | 0 | - |
| RuleType | 9 | 1 | SelectableItemType |
| SPKIDataType | 2 | 0 | - |
| ScoreType | 3 | 0 | - |
| SelChoicesType | 5 | 1 | - |
| SelComplexValueType | 3 | 1 | ComplexValueType |
| SelNumType | 2 | 0 | - |
| SelStringType | 2 | 0 | - |
| ServiceType | 6 | 1 | ItAssetType |
| Set | 2 | 0 | - |
| SetElementType | 2 | 0 | - |
| SignatureMethodType | 3 | 0 | - |
| SignaturePropertiesType | 2 | 0 | - |
| SignaturePropertyType | 3 | 0 | - |
| SignatureValueType | 2 | 0 | - |
| SignedInfoType | 2 | 0 | - |
| Simple | 1 | 0 | - |
| SoftwareType | 4 | 1 | ItAssetType |
| SortingCodeElementType | 2 | 0 | - |
| SplitFunctionType | 2 | 0 | - |
| StateRefType | 2 | 0 | - |
| StateType | 6 | 0 | - |
| Status | 3 | 1 | StatusType |
| StatusElementType | 3 | 1 | StatusType |
| SubAdministrativeAreaElementType | 5 | 0 | - |
| SubAdministrativeAreaNameElementType | 2 | 0 | - |
| SubPremiseLocationElementType | 1 | 0 | - |
| SubPremiseNameElementType | 3 | 0 | - |
| SubPremiseNumberElementType | 6 | 0 | - |
| SubPremiseNumberPrefixElementType | 3 | 0 | - |
| SubPremiseNumberSuffixElementType | 3 | 0 | - |
| SubPremiseType | 10 | 3 | - |
| SubstringFunctionType | 3 | 0 | - |
| SuffixElementType | 3 | 0 | - |
| SupplementaryPostalServiceDataElementType | 2 | 0 | - |
| SyntheticId | 3 | 0 | - |
| SyntheticIdElementType | 3 | 0 | - |
| SystemNameElementType | 1 | 0 | - |
| SystemType | 4 | 1 | ItAssetType |
| TailoringBenchmarkReferenceType | 3 | 1 | BenchmarkReferenceType |
| TailoringReferenceType | 5 | 0 | - |
| TailoringVersionType | 2 | 0 | - |
| TailoringType | 5 | 2 | - |
| TargetFactsType | 1 | 0 | - |
| TargetIdRefType | 4 | 0 | - |
| TelephoneNumber | 2 | 1 | TelephoneNumberType |
| TelephoneNumberElementType | 2 | 1 | TelephoneNumberType |
| TestResultType | 18 | 5 | - |
| TestType | 8 | 0 | - |
| TextWithSubType | 2 | 0 | - |
| Thoroughfare | 14 | 0 | - |
| ThoroughfareElementType | 14 | 4 | - |
| ThoroughfareNameType | 2 | 0 | - |
| ThoroughfareNumber | 6 | 0 | - |
| ThoroughfareNumberElementType | 6 | 0 | - |
| ThoroughfareNumberFromElementType | 1 | 0 | - |
| ThoroughfareNumberPrefix | 3 | 0 | - |
| ThoroughfareNumberPrefixElementType | 3 | 0 | - |
| ThoroughfareNumberRangeElementType | 9 | 0 | - |
| ThoroughfareNumberSuffix | 3 | 0 | - |
| ThoroughfareNumberSuffixElementType | 3 | 0 | - |
| ThoroughfareNumberToElementType | 1 | 0 | - |
| TimeDifferenceFunctionType | 3 | 0 | - |
| TitleElementType | 3 | 0 | - |
| TitleEltType | 1 | 0 | - |
| TransformType | 3 | 0 | - |
| TransformsType | 1 | 0 | - |
| UniqueFunctionType | 1 | 0 | - |
| UriRefType | 2 | 0 | - |
| UrlElementType | 1 | 0 | - |
| VariableComponentType | 2 | 0 | - |
| VersionElementType | 1 | 0 | - |
| WarningType | 3 | 1 | HtmlTextWithSubType |
| WebsiteType | 4 | 1 | ItAssetType |
| WebsiteUrl | 1 | 0 | - |
| WebsiteUrlElementType | 1 | 0 | - |
| X509DataType | 6 | 0 | - |
| X509IssuerSerialType | 3 | 0 | - |
| XAL | 2 | 0 | - |
| XALElementType | 2 | 0 | - |
| XNL | 2 | 0 | - |
| XNLElementType | 2 | 0 | - |
| Xccdf12CheckType | 7 | 1 | - |
| Xccdf12ItemType | 8 | 1 | - |
| Xccdf12MessageType | 2 | 0 | - |
| Xccdf12MetadataType | 1 | 0 | - |
| Xccdf12ReferenceType | 3 | 0 | - |
| Xccdf12SignatureType | 1 | 0 | - |
| Xccdf12TextType | 2 | 0 | - |
| Xccdf12ValueType | 12 | 1 | ItemType |
