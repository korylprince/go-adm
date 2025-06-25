// DO NOT EDIT
// generated from https://github.com/apple/device-management.git:0a4527c5ea21825fd23e08273ccdb9e2302458ce/declarative/declarations/management

package management

const DeviceManagementGenerateHash = "0a4527c5ea21825fd23e08273ccdb9e2302458ce"

var DeclarationMap = map[string]any{
	"com.apple.management.organization-info":   ManagementOrganizationInformation{},
	"com.apple.management.properties":          ManagementProperties{},
	"com.apple.management.server-capabilities": ManagementServerCapabilities{},
}

// The declaration to configure the managing organization's contact information.
type ManagementOrganizationInformation struct {
	// The name of the organization.
	Name string `json:"Name" plist:"Name" required:"true"`
	// The email address of the contact person for the organization.
	Email *string `json:"Email,omitempty" plist:"Email,omitempty"`
	// The website of the organization to contact for support.
	URL *string `json:"URL,omitempty" plist:"URL,omitempty"`
	// The additional properties that verify the identity and authenticity of the organization.
	Proof *Proof `json:"Proof,omitempty" plist:"Proof,omitempty"`
}

func (p *ManagementOrganizationInformation) DeclarationType() string {
	return "com.apple.management.organization-info"
}

// The additional properties that verify the identity and authenticity of the organization.
type Proof struct {
	// A token that verifies the identity of the organization when using this service.
	IdentityToken *string `json:"IdentityToken,omitempty" plist:"IdentityToken,omitempty"`
}

// The declaration to configure the properties on the device.
type ManagementProperties map[string]any

func (p ManagementProperties) DeclarationType() string {
	return "com.apple.management.properties"
}

// The declaration to configure the server's feature set.
type ManagementServerCapabilities struct {
	// The server's protocol version.
	Version string `json:"Version" plist:"Version" required:"true"`
	// A dictionary that contains the server's optional protocol features.
	// Each dictionary item uses the key name to represent a feature, and the value to hold the feature's associated parameters. This protocol reserves keys with a prefix of `com.apple.`, which appear as subkeys in this dictionary.
	SupportedFeatures map[string]any `json:"SupportedFeatures" plist:"SupportedFeatures" required:"true"`
}

func (p *ManagementServerCapabilities) DeclarationType() string {
	return "com.apple.management.server-capabilities"
}
