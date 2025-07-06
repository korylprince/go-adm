// DO NOT EDIT
// generated from https://github.com/apple/device-management.git:0a4527c5ea21825fd23e08273ccdb9e2302458ce/declarative/declarations/activations

package activations

const DeviceManagementGenerateHash = "0a4527c5ea21825fd23e08273ccdb9e2302458ce"

var DeclarationMap = map[string]any{"com.apple.activation.simple": ActivationSimple{}}

// The declaration to activate a set of configurations.
type ActivationSimple struct {
	// An array of strings that specify the identifiers of configurations to install. A failure to install one of the configurations doesn't prevent other configurations from installing.
	StandardConfigurations []string `json:"StandardConfigurations" plist:"StandardConfigurations" required:"true"`
	// A predicate format string as [Apple's Predicate Programming](https://developer.apple.com/library/archive/documentation/Cocoa/Conceptual/Predicates/AdditionalChapters/Introduction.html) describes. The activation only installs when the predicate evaluates to `true` or isn't present.
	Predicate *string `json:"Predicate,omitempty" plist:"Predicate,omitempty"`
}

func (p *ActivationSimple) DeclarationType() string {
	return "com.apple.activation.simple"
}
