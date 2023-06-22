// DO NOT EDIT
// generated from https://github.com/apple/device-management.git:b838baacf2e790db729b6ca3f52724adc8bfb96d/declarative/declarations/activations

package activations

const DeviceManagementGenerateHash = "b838baacf2e790db729b6ca3f52724adc8bfb96d"

var DeclarationMap = map[string]any{"com.apple.activation.simple": ActivationSimple{}}

// An activation used to install a set of configurations.
type ActivationSimple struct {
	// An array of strings that specify the identifiers of configurations to install. A failure to install one of the configurations doesn't prevent other configurations from installing.
	StandardConfigurations []string `json:"StandardConfigurations" required:"true"`
	// A predicate format string as Apple's Predicate Programming <https://developer.apple.com/library/archive/documentation/Cocoa/Conceptual/Predicates/AdditionalChapters/Introduction.html> describes. The activation only installs when the predicate evaluates to 'true' or isn't present.
	Predicate *string `json:"Predicate,omitempty"`
}

func (p *ActivationSimple) DeclarationType() string {
	return "com.apple.activation.simple"
}
