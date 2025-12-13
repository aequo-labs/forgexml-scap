package schemas

import (
	"encoding/xml"
)

type FileBehaviors struct {
	XMLName xml.Name `xml:"http://oval.mitre.org/XMLSchema/oval-definitions-5#windows FileBehaviors"`
}

type FileAuditPermissionsBehaviors struct {
	XMLName xml.Name `xml:"http://oval.mitre.org/XMLSchema/oval-definitions-5#windows FileAuditPermissionsBehaviors"`
}

type FileEffectiveRightsBehaviors struct {
	XMLName xml.Name `xml:"http://oval.mitre.org/XMLSchema/oval-definitions-5#windows FileEffectiveRightsBehaviors"`
}

type RegistryBehaviors struct {
	XMLName xml.Name `xml:"http://oval.mitre.org/XMLSchema/oval-definitions-5#windows RegistryBehaviors"`
}

type RegkeyAuditPermissionsBehaviors struct {
	XMLName xml.Name `xml:"http://oval.mitre.org/XMLSchema/oval-definitions-5#windows RegkeyAuditPermissionsBehaviors"`
}

type RegkeyEffectiveRightsBehaviors struct {
	XMLName xml.Name `xml:"http://oval.mitre.org/XMLSchema/oval-definitions-5#windows RegkeyEffectiveRightsBehaviors"`
}

type SidBehaviors struct {
	XMLName xml.Name `xml:"http://oval.mitre.org/XMLSchema/oval-definitions-5#windows SidBehaviors"`
}

type EntityStateAddrTypeType struct {
	XMLName xml.Name `xml:"http://oval.mitre.org/XMLSchema/oval-definitions-5#windows EntityStateAddrTypeType"`
}

type EntityStateAdstypeType struct {
	XMLName xml.Name `xml:"http://oval.mitre.org/XMLSchema/oval-definitions-5#windows EntityStateAdstypeType"`
}

type EntityStateAuditType struct {
	XMLName xml.Name `xml:"http://oval.mitre.org/XMLSchema/oval-definitions-5#windows EntityStateAuditType"`
}

type EntityStateInterfaceTypeType struct {
	XMLName xml.Name `xml:"http://oval.mitre.org/XMLSchema/oval-definitions-5#windows EntityStateInterfaceTypeType"`
}

type EntityStateFileTypeType struct {
	XMLName xml.Name `xml:"http://oval.mitre.org/XMLSchema/oval-definitions-5#windows EntityStateFileTypeType"`
}

type EntityObjectNamingContextType struct {
	XMLName xml.Name `xml:"http://oval.mitre.org/XMLSchema/oval-definitions-5#windows EntityObjectNamingContextType"`
}

type EntityStateNamingContextType struct {
	XMLName xml.Name `xml:"http://oval.mitre.org/XMLSchema/oval-definitions-5#windows EntityStateNamingContextType"`
}

type EntityObjectProtocolType struct {
	XMLName xml.Name `xml:"http://oval.mitre.org/XMLSchema/oval-definitions-5#windows EntityObjectProtocolType"`
}

type EntityStateProtocolType struct {
	XMLName xml.Name `xml:"http://oval.mitre.org/XMLSchema/oval-definitions-5#windows EntityStateProtocolType"`
}

type EntityObjectRegistryHiveType struct {
	XMLName xml.Name `xml:"http://oval.mitre.org/XMLSchema/oval-definitions-5#windows EntityObjectRegistryHiveType"`
}

type EntityStateRegistryHiveType struct {
	XMLName xml.Name `xml:"http://oval.mitre.org/XMLSchema/oval-definitions-5#windows EntityStateRegistryHiveType"`
}

type EntityStateRegistryTypeType struct {
	XMLName xml.Name `xml:"http://oval.mitre.org/XMLSchema/oval-definitions-5#windows EntityStateRegistryTypeType"`
}
