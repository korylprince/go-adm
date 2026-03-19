// DO NOT EDIT
// generated from https://github.com/apple/device-management.git:f878dea98fb88293a3686e44bcfb891f8e78f98f/mdm/checkin

package checkin

// Authenticates a user during MDM payload installation.
// Check-in protocol authenticate request and response.
type Authenticate struct {
	// The device's name.
	DeviceName string `json:"DeviceName" plist:"DeviceName" required:"true"`
	// The device's model name.
	ModelName string `json:"ModelName" plist:"ModelName" required:"true"`
	// The device's model.
	Model string `json:"Model" plist:"Model" required:"true"`
	// The message type, which requires a value of `Authenticate`.
	MessageType AuthenticateMessageType `json:"MessageType" plist:"MessageType" required:"true"`
	// The topic that the device subscribes to.
	Topic string `json:"Topic" plist:"Topic" required:"true"`
	// The device's UDID (unique device identifier). The system requires this value if the enrollment type is a device enrollment.
	UDID *string `json:"UDID,omitempty" plist:"UDID,omitempty"`
	// The per-enrollment identifier for the device. The system requires this value if the enrollment type is a user enrollment.
	// Available in iOS 13 and later, macOS 10.15 and later, and visionOS 2 and later.
	EnrollmentID *string `json:"EnrollmentID,omitempty" plist:"EnrollmentID,omitempty"`
	// The device's OS version.
	OSVersion *string `json:"OSVersion,omitempty" plist:"OSVersion,omitempty"`
	// The device's build version.
	BuildVersion *string `json:"BuildVersion,omitempty" plist:"BuildVersion,omitempty"`
	// The device's product name (such as `iPhone17,2`).
	ProductName *string `json:"ProductName,omitempty" plist:"ProductName,omitempty"`
	// The device's serial number.
	SerialNumber *string `json:"SerialNumber,omitempty" plist:"SerialNumber,omitempty"`
	// The device's IMEI (International Mobile Equipment Identity).
	IMEI *string `json:"IMEI,omitempty" plist:"IMEI,omitempty"`
	// The device's MEID (Mobile Equipment Identifier).
	MEID *string `json:"MEID,omitempty" plist:"MEID,omitempty"`
}

// The message type, which requires a value of `Authenticate`.
type AuthenticateMessageType string

const (
	AuthenticateMessageTypeAuthenticate AuthenticateMessageType = "Authenticate"
)

// Responds to the removal of the MDM enrollment profile from a device.
// Check-in protocol check out request and response.
type CheckOut struct {
	// The message type, which requires a value of `CheckOut`.
	MessageType CheckOutMessageType `json:"MessageType" plist:"MessageType" required:"true"`
	// The topic the device subscribes to.
	Topic string `json:"Topic" plist:"Topic" required:"true"`
	// The device's UDID (unique device identifier). The system requires this value if the enrollment type is a device enrollment.
	UDID string `json:"UDID" plist:"UDID" required:"true"`
	// The per-enrollment identifier for the device. The system requires this value if the enrollment type is a user enrollment.
	// Available in iOS 13 and later, macOS 10.15 and later, and visionOS 2 and later.
	EnrollmentID string `json:"EnrollmentID" plist:"EnrollmentID" required:"true"`
}

// The message type, which requires a value of `CheckOut`.
type CheckOutMessageType string

const (
	CheckOutMessageTypeCheckOut CheckOutMessageType = "CheckOut"
)

// Sends declarative management requests to the server.
// Check-in protocol declarative management request and response.
type DeclarativeManagement struct {
	// The message type, which requires a value of `DeclarativeManagement`.
	MessageType DeclarativeManagementMessageType `json:"MessageType" plist:"MessageType" required:"true"`
	// The type of operation the declaration is requesting. This key needs to be one of these values:
	// - `tokens`: For fetching synchronization tokens from the server
	// - `declaration-items`: For fetching the declaration manifest from the server
	// - `status`: For sending a status report to the server
	// - `declaration/…/…`: For fetching a specific declaration from the server. Include the declaration type and identifier separated by slash characters (`/`).
	Endpoint string `json:"Endpoint" plist:"Endpoint" required:"true"`
	// A Base64-encoded JSON object using the `SynchronizationTokens` schema.
	Data *[]byte `json:"Data,omitempty" plist:"Data,omitempty"`
	// The device's UDID (unique device identifier). The system requires this value if the enrollment type is a device enrollment.
	UDID string `json:"UDID" plist:"UDID" required:"true"`
	// The per-enrollment identifier for the device. The system requires this value if the enrollment type is a user enrollment.
	EnrollmentID string `json:"EnrollmentID" plist:"EnrollmentID" required:"true"`
	// The per-enrollment identifier for the user. The system requires this value if the enrollment type is a user enrollment on the user channel.
	EnrollmentUserID string `json:"EnrollmentUserID" plist:"EnrollmentUserID" required:"true"`
	// For macOS, this value is the short name of the user.
	// For Shared iPad, this value is the Managed Apple Account identifier of the user on Shared iPad. It indicates that the token is for the user channel.
	UserShortName *string `json:"UserShortName,omitempty" plist:"UserShortName,omitempty"`
	// For macOS, this value is the ID of the user.
	// For Shared iPad, this value is `FFFFFFFF-FFFF-FFFF-FFFF-FFFFFFFFFFFF` to indicate that authentication doesn't occur.
	UserID *string `json:"UserID,omitempty" plist:"UserID,omitempty"`
	// The full name of the user.
	UserLongName string `json:"UserLongName" plist:"UserLongName" required:"true"`
}

// The message type, which requires a value of `DeclarativeManagement`.
type DeclarativeManagementMessageType string

const (
	DeclarativeManagementMessageTypeDeclarativeManagement DeclarativeManagementMessageType = "DeclarativeManagement"
)

// Gets the bootstrap token from the server.
// Check-in protocol get bootstrap token data request and response.
type GetBootstrapToken struct {
	// The message type, which requires a value of `GetBootstrapToken`.
	MessageType GetBootstrapTokenMessageType `json:"MessageType" plist:"MessageType" required:"true"`
	// If `true`, the device is awaiting a `Device-Configured-Command` command before proceeding through Setup Assistant.
	AwaitingConfiguration *bool `json:"AwaitingConfiguration,omitempty" plist:"AwaitingConfiguration,omitempty"`
}
type GetBootstrapTokenResponse struct {
	// The current bootstrap token data for the device.
	BootstrapToken *[]byte `json:"BootstrapToken,omitempty" plist:"BootstrapToken,omitempty"`
}

// The message type, which requires a value of `GetBootstrapToken`.
type GetBootstrapTokenMessageType string

const (
	GetBootstrapTokenMessageTypeGetBootstrapToken GetBootstrapTokenMessageType = "GetBootstrapToken"
)

// Gets a token from the server.
// Check-in protocol get token data request and response.
type GetToken struct {
	// The message type, which requires a value of `GetToken`.
	MessageType GetTokenMessageType `json:"MessageType" plist:"MessageType" required:"true"`
	// A string that specifies the service for the requested token.
	TokenServiceType TokenServiceType `json:"TokenServiceType" plist:"TokenServiceType" required:"true"`
	// Parameters that the system uses to generate the token.
	TokenParameters *TokenParameters `json:"TokenParameters,omitempty" plist:"TokenParameters,omitempty"`
	// The device's UDID (unique device identifier). The system requires this value if the enrollment type is a device enrollment.
	UDID string `json:"UDID" plist:"UDID" required:"true"`
	// The per-enrollment identifier for the device. The system requires this value if the enrollment type is a user enrollment.
	EnrollmentID string `json:"EnrollmentID" plist:"EnrollmentID" required:"true"`
	// The per-enrollment identifier for the user. The system requires this value if the enrollment type is a user enrollment on the user channel.
	EnrollmentUserID string `json:"EnrollmentUserID" plist:"EnrollmentUserID" required:"true"`
	// For macOS, this value is the short name of the user.
	// For Shared iPad, this value is the Managed Apple Account identifier of the user. When present, it indicates that the token is for the user channel.
	UserShortName *string `json:"UserShortName,omitempty" plist:"UserShortName,omitempty"`
	// For macOS, this value is the ID of the user.
	// For Shared iPad, this value is `FFFFFFFF-FFFF-FFFF-FFFF-FFFFFFFFFFFF` to indicate that authentication doesn't occur.
	UserID *string `json:"UserID,omitempty" plist:"UserID,omitempty"`
	// The full name of the user.
	UserLongName string `json:"UserLongName" plist:"UserLongName" required:"true"`
}
type GetTokenResponse struct {
	// The token data. If the token is a string value, it needs to be a UTF-8-encoded string.
	TokenData []byte `json:"TokenData" plist:"TokenData" required:"true"`
}

// The message type, which requires a value of `GetToken`.
type GetTokenMessageType string

const (
	GetTokenMessageTypeGetToken GetTokenMessageType = "GetToken"
)

// A string that specifies the service for the requested token.
type TokenServiceType string

const (
	TokenServiceTypeComAppleMaid         TokenServiceType = "com.apple.maid"
	TokenServiceTypeComAppleWatchPairing TokenServiceType = "com.apple.watch.pairing"
)

// Parameters that the system uses to generate the token.
type TokenParameters struct {
	// A security token to generate the server token. Required by the `com.apple.watch.pairing` service type.
	SecurityToken *string `json:"SecurityToken,omitempty" plist:"SecurityToken,omitempty"`
	// The identifier of the phone paired to the watch. Required by the `com.apple.watch.pairing` service type.
	PhoneUDID *string `json:"PhoneUDID,omitempty" plist:"PhoneUDID,omitempty"`
	// The identifier of the watch paired to the phone. Required by the `com.apple.watch.pairing` service type.
	WatchUDID *string `json:"WatchUDID,omitempty" plist:"WatchUDID,omitempty"`
}

// Gets the return-to-service configuration from the server.
// Check-in protocol send ReturnToService request.
type ReturnToService1 struct {
	// The message type, which requires a value of `ReturnToService`.
	MessageType ReturnToServiceMessageType `json:"MessageType" plist:"MessageType" required:"true"`
	// The device's UDID (unique device identifier). The system requires this value if the enrollment type is a device enrollment.
	UDID string `json:"UDID" plist:"UDID" required:"true"`
}
type ReturnToServiceResponse struct {
	// A dictionary containing the configuration for return to service.
	ReturnToService ReturnToServiceReturnToService `json:"ReturnToService" plist:"ReturnToService" required:"true"`
}

// The message type, which requires a value of `ReturnToService`.
type ReturnToServiceMessageType string

const (
	ReturnToServiceMessageTypeReturnToService ReturnToServiceMessageType = "ReturnToService"
)

// A dictionary containing the configuration for return to service.
type ReturnToServiceReturnToService struct {
	// If `true`, the device automatically erases itself and then performs reenrollment.
	Enabled bool `json:"Enabled" plist:"Enabled" required:"true"`
	// The Wi-Fi profile that installs after erasure when using return to service. This is required when the device doesn't have Ethernet access.
	WiFiProfileData *[]byte `json:"WiFiProfileData,omitempty" plist:"WiFiProfileData,omitempty"`
	// The MDM profile that installs after erasure when using return to service. If provided, the device uses this profile directly instead of fetching it from the server. This key is required if the device's Automated Device Enrollment profile contains the `configuration-web-url` key.
	// The device always downloads the Automated Device Enrollment profile even when this key is present, so the supervision identity, MDM removability, and other settings still apply. However, the device doesn't use the specified URL in the Automated Device Enrollment profile to fetch the MDM profile.
	MDMProfileData *[]byte `json:"MDMProfileData,omitempty" plist:"MDMProfileData,omitempty"`
	// The system uses the bootstrap token for return to service with app preservation. Required when Automated Device Enrollment enables return to service for the device.
	// If the bootstrap token isn't present, the device performs a full erasure and a regular return to service, and can't preserve any data for app preservation.
	BootstrapToken *[]byte `json:"BootstrapToken,omitempty" plist:"BootstrapToken,omitempty"`
}

// Sends the bootstrap token to the server.
// Check-in protocol set bootstrap token data request and response.
type SetBootstrapToken struct {
	// The message type, which requires a value of `SetBootstrapToken`.
	MessageType SetBootstrapTokenMessageType `json:"MessageType" plist:"MessageType" required:"true"`
	// The device's bootstrap token data. If this field is missing or zero length, the server needs to remove the bootstrap token for this device.
	BootstrapToken *[]byte `json:"BootstrapToken,omitempty" plist:"BootstrapToken,omitempty"`
	// If `true`, the device is awaiting a `Device-Configured-Command` command before proceeding through Setup Assistant.
	AwaitingConfiguration *bool `json:"AwaitingConfiguration,omitempty" plist:"AwaitingConfiguration,omitempty"`
}

// The message type, which requires a value of `SetBootstrapToken`.
type SetBootstrapTokenMessageType string

const (
	SetBootstrapTokenMessageTypeSetBootstrapToken SetBootstrapTokenMessageType = "SetBootstrapToken"
)

// Updates the token for a device on the server.
// Check-in protocol token update request and response.
type TokenUpdate struct {
	// If `true`, the device isn't on-console.
	NotOnConsole bool `json:"NotOnConsole" plist:"NotOnConsole" required:"true"`
	// The message type, which requires a value of `TokenUpdate`.
	MessageType TokenUpdateMessageType `json:"MessageType" plist:"MessageType" required:"true"`
	// The topic the device subscribes to.
	Topic string `json:"Topic" plist:"Topic" required:"true"`
	// The device's UDID (unique device identifier). The system requires this value if the enrollment type is a device enrollment.
	UDID string `json:"UDID" plist:"UDID" required:"true"`
	// The per-enrollment identifier for the device. The system requires this value if the enrollment type is a user enrollment.
	// Available in iOS 13 and later, macOS 10.15 and later, and visionOS 2 and later.
	EnrollmentID string `json:"EnrollmentID" plist:"EnrollmentID" required:"true"`
	// The per-enrollment identifier for the user. The system requires this value if the enrollment type is a user enrollment on the user channel.
	// Available in macOS 10.15 and later.
	EnrollmentUserID string `json:"EnrollmentUserID" plist:"EnrollmentUserID" required:"true"`
	// For macOS, this value is the short name of the user.
	// For Shared iPad, this value is the Managed Apple Account identifier of the user on Shared iPad. It indicates that the token is for the user channel.
	UserShortName *string `json:"UserShortName,omitempty" plist:"UserShortName,omitempty"`
	// For macOS, this value is the ID of the user.
	// For Shared iPad, this value is `FFFFFFFF-FFFF-FFFF-FFFF-FFFFFFFFFFFF` to indicate that authentication doesn't occur.
	UserID *string `json:"UserID,omitempty" plist:"UserID,omitempty"`
	// The full name of the user.
	UserLongName string `json:"UserLongName" plist:"UserLongName" required:"true"`
	// The push token for the device.
	Token []byte `json:"Token" plist:"Token" required:"true"`
	// The magic string to include in the push notification message.
	PushMagic string `json:"PushMagic" plist:"PushMagic" required:"true"`
	// The data to use to unlock the device. If provided, the server needs to retain this data and send it when trying to implement `Clear-Passcode-Command`.
	UnlockToken *[]byte `json:"UnlockToken,omitempty" plist:"UnlockToken,omitempty"`
	// If `true` from the device channel, the device is awaiting a `Device-Configured-Command` command before proceeding through Setup Assistant.
	// If `true` from the user channel (Shared iPad only), the device is awaiting a `User-Configured-Command` command before proceeding through Setup Assistant.
	AwaitingConfiguration *bool `json:"AwaitingConfiguration,omitempty" plist:"AwaitingConfiguration,omitempty"`
}

// The message type, which requires a value of `TokenUpdate`.
type TokenUpdateMessageType string

const (
	TokenUpdateMessageTypeTokenUpdate TokenUpdateMessageType = "TokenUpdate"
)

// Authenticates a user with a two-step authentication protocol.
// Authenticate network or mobile users with MDM.
type UserAuthenticate struct {
	// The message type, which requires a value of `UserAuthenticate`.
	MessageType UserAuthenticateMessageType `json:"MessageType" plist:"MessageType" required:"true"`
	// The device's UDID (unique device identifier). The system requires this value if the enrollment type is a device enrollment.
	UDID string `json:"UDID" plist:"UDID" required:"true"`
	// The local mobile user's GUID or the network user's GUID from an Open Directory record.
	UserID string `json:"UserID" plist:"UserID" required:"true"`
	// A string that the client provides in the second `User-Authenticate` request after receiving `DigestChallenge` from the server on the first `User-Authenticate` request.
	DigestResponse string `json:"DigestResponse" plist:"DigestResponse" required:"true"`
}

// The message type, which requires a value of `UserAuthenticate`.
type UserAuthenticateMessageType string

const (
	UserAuthenticateMessageTypeUserAuthenticate UserAuthenticateMessageType = "UserAuthenticate"
)
