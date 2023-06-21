// DO NOT EDIT
// generated from https://github.com/apple/device-management.git:b838baacf2e790db729b6ca3f52724adc8bfb96d/declarative/declarations/management

package management

const DeviceManagementGenerateHash = "b838baacf2e790db729b6ca3f52724adc8bfb96d"

var DeclarationMap = map[string]any{
	"com.apple.management.organization-info":   ManagementOrganizationInformation{},
	"com.apple.management.properties":          ManagementProperties{},
	"com.apple.management.server-capabilities": ManagementServerCapabilities{},
}

// The additional properties that verify the identity and authenticity of the organization.
type Proof struct {
	// A token that verifies the identity of the organization when using this service.
	IdentityToken *string `json:"IdentityToken,omitempty"`
}

// Use this declaration to tell the client about the server's organization information.
type ManagementOrganizationInformation struct {
	// The name of the organization.
	Name string `json:"Name"`
	// The email address of the contact person for the organization.
	Email *string `json:"Email,omitempty"`
	// The website of the organization to contact for support.
	URL *string `json:"URL,omitempty"`
	// The additional properties that verify the identity and authenticity of the organization.
	Proof *Proof `json:"Proof,omitempty"`
}

func (p *ManagementOrganizationInformation) DeclarationType() string {
	return "com.apple.management.organization-info"
}

// Use this declaration to set properties on the device.
type ManagementProperties struct {
	// Each entry represents a property key/value.
	ANY *any `json:"ANY,omitempty"`
}

func (p *ManagementProperties) DeclarationType() string {
	return "com.apple.management.properties"
}

// A dictionary that contains the server's optional protocol features.
// Each dictionary item uses the key name to represent a feature, and the value to hold the feature's associated parameters. This protocol reserves keys with a prefix of “'com.apple.'”, which appear as subkeys in this dictionary.
type SupportedFeatures struct {
	// Additional keys may be present.
	ANY *any `json:"ANY,omitempty"`
}

// Use this declaration to tell the client about the server's capabilities.
type ManagementServerCapabilities struct {
	// The server's protocol version.
	Version string `json:"Version"`
	// A dictionary that contains the server's optional protocol features.
	// Each dictionary item uses the key name to represent a feature, and the value to hold the feature's associated parameters. This protocol reserves keys with a prefix of “'com.apple.'”, which appear as subkeys in this dictionary.
	SupportedFeatures SupportedFeatures `json:"SupportedFeatures"`
}

func (p *ManagementServerCapabilities) DeclarationType() string {
	return "com.apple.management.server-capabilities"
}
