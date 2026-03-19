// DO NOT EDIT
// generated from https://github.com/apple/device-management.git:f878dea98fb88293a3686e44bcfb891f8e78f98f/mdm/errors

package mdmerrors

// An error response that indicates Platform SSO is required.
type ErrorCodePlatformSSORequired struct {
	// Indicates that the device needs to do Platform SSO before enrollment and setup can proceed.
	Code ErrorCodePlatformSSORequiredCode `json:"code" plist:"code" required:"true"`
	// A description of the error. Only use this for logging purposes and don't display it to the user.
	Description *string `json:"description,omitempty" plist:"description,omitempty"`
	// A description of the error to display to the user.
	Message *string `json:"message,omitempty" plist:"message,omitempty"`
	// A dictionary that contains additional data about the error code.
	Details ErrorCodePlatformSSORequiredDetails `json:"details" plist:"details" required:"true"`
}

// Indicates that the device needs to do Platform SSO before enrollment and setup can proceed.
type ErrorCodePlatformSSORequiredCode string

const (
	ErrorCodePlatformSSORequiredCodeComApplePSSORequired ErrorCodePlatformSSORequiredCode = "com.apple.psso.required"
)

// A dictionary that contains additional data about the error code.
type ErrorCodePlatformSSORequiredDetails struct {
	// The URL of the profile containing an `ExtensibleSingleSignOn` profile payload that the device uses to configure the SSO extension for Platform SSO.
	ProfileURL string `json:"ProfileURL" plist:"ProfileURL" required:"true"`
	// A dictionary that specifies the package that the device uses to install an app with the SSO app extension used for Platform SSO.
	Package Package `json:"Package" plist:"Package" required:"true"`
	// The URL the device uses to create an `ASWebAuthenticationSession` to trigger Platform SSO authentication, once the profile and app are installed.
	AuthURL string `json:"AuthURL" plist:"AuthURL" required:"true"`
}

// A dictionary that specifies the package that the device uses to install an app with the SSO app extension used for Platform SSO.
type Package struct {
	// The URL of the app manifest, which needs to begin with `https:`.
	ManifestURL string `json:"ManifestURL" plist:"ManifestURL" required:"true"`
	// An array of DER-encoded certificates to pin the connection when fetching the `ManifestURL`.
	PinningCerts *[][]byte `json:"PinningCerts,omitempty" plist:"PinningCerts,omitempty"`
	// If `true`, certificate revocation checks require a positive response when using certificate pinning with `PinningCerts`.
	PinningRevocationCheckRequired *bool `json:"PinningRevocationCheckRequired,omitempty" plist:"PinningRevocationCheckRequired,omitempty"`
}

// An error response that indicates the system requires a software update.
type ErrorCodeSoftwareUpdateRequired struct {
	// Indicates that the device needs to perform a software update before enrollment and setup can proceed.
	Code ErrorCodeSoftwareUpdateRequiredCode `json:"code" plist:"code" required:"true"`
	// A description of the error. Only use this for logging purposes and don't display it to the user.
	Description *string `json:"description,omitempty" plist:"description,omitempty"`
	// A description of the error to display to the user.
	Message *string `json:"message,omitempty" plist:"message,omitempty"`
	// A dictionary that contains additional data about the error code.
	Details ErrorCodeSoftwareUpdateRequiredDetails `json:"details" plist:"details" required:"true"`
}

// Indicates that the device needs to perform a software update before enrollment and setup can proceed.
type ErrorCodeSoftwareUpdateRequiredCode string

const (
	ErrorCodeSoftwareUpdateRequiredCodeComAppleSoftwareUpdateRequired ErrorCodeSoftwareUpdateRequiredCode = "com.apple.softwareupdate.required"
)

// A dictionary that contains additional data about the error code.
type ErrorCodeSoftwareUpdateRequiredDetails struct {
	// The OS version that the device needs to update to, for example, "16.1". This identifier can include a supplemental version identifier, for example, "16.1 (a)".
	OSVersion string `json:"OSVersion" plist:"OSVersion" required:"true"`
	// The build version that the device needs to update to, for example, "20A242. The systems uses the build version for testing during seeding periods. This identifier can include a supplemental version identifier, for example, "20A242a". If the `BuildVersion` isn't consistent with the `OSVersion`, `OSVersion` take precedence.
	BuildVersion *string `json:"BuildVersion,omitempty" plist:"BuildVersion,omitempty"`
	// The device enrolls in the beta program, allowing enforced software updates to beta program OS versions. The device remains in the beta program after the system completes the enforced software update.
	RequireBetaProgram *RequireBetaProgram `json:"RequireBetaProgram,omitempty" plist:"RequireBetaProgram,omitempty"`
}

// The device enrolls in the beta program, allowing enforced software updates to beta program OS versions. The device remains in the beta program after the system completes the enforced software update.
type RequireBetaProgram struct {
	// A human readable description of the beta program.
	Description string `json:"Description" plist:"Description" required:"true"`
	// The AxM seeding service token for the AxM organization the MDM server is part of. The system uses this token to enroll the device in the corresponding beta program.
	Token string `json:"Token" plist:"Token" required:"true"`
}

// An error response that indicates a device needs to unenroll.
type ErrorUnrecognizedDevice struct {
	// Indicates that the device is not recognized by the server. This causes the device to unenroll from MDM.
	Code ErrorUnrecognizedDeviceCode `json:"code" plist:"code" required:"true"`
	// A description of the error. Only use this for logging purposes and don't display it to the user.
	Description *string `json:"description,omitempty" plist:"description,omitempty"`
	// A description of the error to display to the user.
	Message *string `json:"message,omitempty" plist:"message,omitempty"`
}

// Indicates that the device is not recognized by the server. This causes the device to unenroll from MDM.
type ErrorUnrecognizedDeviceCode string

const (
	ErrorUnrecognizedDeviceCodeComAppleUnrecognizedDevice ErrorUnrecognizedDeviceCode = "com.apple.unrecognized.device"
)

// An error response that indicates a missing pairing token.
type ErrorCodePairingTokenMissing struct {
	// Indicates that the pairing token, which the system requires to enroll the watch, is missing.
	Code ErrorCodePairingTokenMissingCode `json:"code" plist:"code" required:"true"`
	// A description of the error. Only use this for logging purposes and don't display it to the user.
	Description *string `json:"description,omitempty" plist:"description,omitempty"`
	// A description of the error to display to the user.
	Message *string `json:"message,omitempty" plist:"message,omitempty"`
	// A dictionary that contains additional data about the error code.
	Details ErrorCodePairingTokenMissingDetails `json:"details" plist:"details" required:"true"`
}

// Indicates that the pairing token, which the system requires to enroll the watch, is missing.
type ErrorCodePairingTokenMissingCode string

const (
	ErrorCodePairingTokenMissingCodeComAppleWatchPairingTokenMissing ErrorCodePairingTokenMissingCode = "com.apple.watch.pairing.token.missing"
)

// A dictionary that contains additional data about the error code.
type ErrorCodePairingTokenMissingDetails struct {
	// The security token to pass to the phone's MDM server to create the pairing token. This token needs to be a random UUID string.
	SecurityToken string `json:"security-token" plist:"security-token" required:"true"`
}

// An error response that indicates a well-known service discovery request failed.
type ErrorWellKnownFailed struct {
	// Indicates that the well-known request has failed.
	Code ErrorWellKnownFailedCode `json:"code" plist:"code" required:"true"`
	// A description of the error. Only use this for logging purposes and don't display it to the user.
	Description *string `json:"description,omitempty" plist:"description,omitempty"`
	// A description of the error to display to the user.
	Message *string `json:"message,omitempty" plist:"message,omitempty"`
}

// Indicates that the well-known request has failed.
type ErrorWellKnownFailedCode string

const (
	ErrorWellKnownFailedCodeComAppleWellKnownFailed ErrorWellKnownFailedCode = "com.apple.well-known.failed"
)
