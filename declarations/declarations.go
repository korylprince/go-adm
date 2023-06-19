//go:generate declgen -repo "https://github.com/apple/device-management.git" -commit "b838baacf2e790db729b6ca3f52724adc8bfb96d"

package declarations

import "encoding/json"

type DeclarationPayload interface {
	DeclarationType() string
}

type Declaration struct {
	// A string uniquely identifying this declaration.
	Identifier string `json:"Identifier"`

	// A unique token generated by the server specifying a particular revision
	ServerToken string `json:"ServerToken"`

	// The payload describing this declaration.
	Payload DeclarationPayload `json:"Payload"`
}

func (d *Declaration) Type() string {
	if d.Payload == nil {
		return ""
	}
	return d.Payload.DeclarationType()
}

func (d *Declaration) MarshalJSON() ([]byte, error) {
	StructDefaults(d.Payload)

	decl := map[string]any{
		"Type":        d.Type(),
		"Identifier":  d.Identifier,
		"ServerToken": d.ServerToken,
		"Payload":     d.Payload,
	}

	return json.Marshal(decl)
}