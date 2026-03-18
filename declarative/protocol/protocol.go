// DO NOT EDIT
// generated from https://github.com/apple/device-management.git:f878dea98fb88293a3686e44bcfb891f8e78f98f/declarative/protocol

package protocol

const DeviceManagementGenerateHash = "f878dea98fb88293a3686e44bcfb891f8e78f98f"

// The set of available declarations on the server.
type DeclarationItemsResponse struct {
	// The set of available declarations on the server.
	Declarations Declarations `json:"Declarations" required:"true"`
	// The current value of the declarations token. Clients use this to detect when declarations change so they can refetch the token.
	DeclarationsToken string `json:"DeclarationsToken" required:"true"`
}

// The set of available declarations on the server.
type Declarations struct {
	// The list of available activation declarations on the server.
	Activations []*Activations `json:"Activations" required:"true"`
	// The list of available configuration declarations on the server.
	Configurations []*Configurations `json:"Configurations" required:"true"`
	// The list of available asset declarations on the server.
	Assets []*Assets `json:"Assets" required:"true"`
	// The list of available management declarations on the server.
	Management []*Management `json:"Management" required:"true"`
}

// Information about an available declaration on the server.
type Activations struct {
	// The declaration's identifier.
	Identifier string `json:"Identifier" required:"true"`
	// The 'ServerToken' value of the declaration.
	// The client uses this to determine if the actual payload is different from the one on the client. Servers must compute the token over the entire declaration content to ensure the value always changes whenever there's any change to the content.
	ServerToken string `json:"ServerToken" required:"true"`
}

// Information about an available declaration on the server.
type Configurations struct {
	// The declaration's identifier.
	Identifier string `json:"Identifier" required:"true"`
	// The 'ServerToken' value of the declaration.
	// The client uses this to determine if the actual payload is different from the one on the client. Servers must compute the token over the entire declaration content to ensure the value always changes whenever there's any change to the content.
	ServerToken string `json:"ServerToken" required:"true"`
}

// Information about an available declaration on the server.
type Assets struct {
	// The declaration's identifier.
	Identifier string `json:"Identifier" required:"true"`
	// The 'ServerToken' value of the declaration.
	// The client uses this to determine if the actual payload is different from the one on the client. Servers must compute the token over the entire declaration content to ensure the value always changes whenever there's any change to the content.
	ServerToken string `json:"ServerToken" required:"true"`
}

// Information about an available declaration on the server.
type Management struct {
	// The declaration's identifier.
	Identifier string `json:"Identifier" required:"true"`
	// The 'ServerToken' value of the declaration.
	// The client uses this to determine if the actual payload is different from the one on the client. Servers must compute the token over the entire declaration content to ensure the value always changes whenever there's any change to the content.
	ServerToken string `json:"ServerToken" required:"true"`
}

// Status sent by the client.
type StatusReport struct {
	// The status items for this report.
	StatusItems StatusItems `json:"StatusItems" required:"true"`
	// An array of errors for this status report.
	Errors []*Errors `json:"Errors" required:"true"`
	// The system sets this to `true` to indicate that the status report contains the full set of current status, and is not an incremental report. A full status report includes the full set of items in any status array item, not just the changes. Servers use this to replace their entire status for the device, rather than do an incremental update to the existing status. The system sets this to `true` when sending a "safety sync" status report, which is typically sent every 24 hours or so.
	FullReport *bool `json:"FullReport,omitempty"`
}

// The status items for this report.
type StatusItems struct{}

// Error information for a status item that cannot be returned.
type Errors struct {
	// The status item that this error pertains to.
	StatusItem string `json:"StatusItem" required:"true"`
	// An array of reasons for the error.
	Reasons *[]*Reasons `json:"Reasons,omitempty"`
}

// Information about a status error.
type Reasons struct {
	// The error code for this error.
	Code string `json:"Code" required:"true"`
	// The description for this error.
	Description *string `json:"Description,omitempty"`
	// A dictionary that contains further details about this error.
	Details *Details `json:"Details,omitempty"`
}

// A dictionary that contains further details about this error.
type Details struct{}

// The server's synchronization tokens.
type TokensResponse struct {
	// A dictionary of synchronization tokens that describes the state of different types of data on the server. The client uses these tokens to determine which endpoints it needs to use to fetch new or updated data on the server.
	SyncTokens SyncTokens `json:"SyncTokens" required:"true"`
}

// A dictionary of synchronization tokens that describes the state of different types of data on the server. The client uses these tokens to determine which endpoints it needs to use to fetch new or updated data on the server.
type SyncTokens struct{}
