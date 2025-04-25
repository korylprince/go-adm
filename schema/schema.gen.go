// DO NOT EDIT
// generated from https://github.com/apple/device-management.git:7d4ba1a2bde50a4053fa5a5e0ed6c17388d82ab2/docs/schema.yaml

package schema

import yamlschema "github.com/korylprince/go-adm/yamlschema"

const DeviceManagementGenerateHash = "7d4ba1a2bde50a4053fa5a5e0ed6c17388d82ab2"

// Indicates how multiple configurations of the same type are applied. If set to 'single', then only one configuration will be applied. If set to 'multiple', then each configuration is applied separately. If set to 'combined', then all configurations are combined into a single effective configuration.
type PayloadApply string

const (
	PayloadApplySingle   PayloadApply = "single"
	PayloadApplyMultiple PayloadApply = "multiple"
	PayloadApplyCombined PayloadApply = "combined"
)

// Indicates whether a payload or payload key can used with or without shared iPad in effect. If set to 'allowed', then the payload or payload key can be used both with or without shared iPad in effect. If set to 'required', then the payload or payload key can only be used if shared iPad is in effect. If set to 'forbidden', then the payload or payload key cannot be used if shared iPad is in effect. If set to 'ignored', then the payload or payload key can be used, but is ignored if shared iPad is in effect.
type SharediPadMode string

const (
	SharediPadModeAllowed   SharediPadMode = "allowed"
	SharediPadModeRequired  SharediPadMode = "required"
	SharediPadModeForbidden SharediPadMode = "forbidden"
	SharediPadModeIgnored   SharediPadMode = "ignored"
)

// Indicates how a payload or payload key can only be used if user enrollment is in effect. If set to 'allowed', then the payload or payload key can be used both with or without user enrollment in effect. If set to 'required', then the payload or payload key can only be used if user enrollment is in effect. If set to 'forbidden', then the payload or payload key cannot be used if user enrollment is in effect. If set to 'ignored', then the payload or payload key can be used, but is ignored if user enrollment is in effect.
type UserEnrollmentMode string

const (
	UserEnrollmentModeAllowed   UserEnrollmentMode = "allowed"
	UserEnrollmentModeRequired  UserEnrollmentMode = "required"
	UserEnrollmentModeForbidden UserEnrollmentMode = "forbidden"
	UserEnrollmentModeIgnored   UserEnrollmentMode = "ignored"
)

// The type of key. The value `<any>` may be used to indicate that any of the standard values can be used without any expectation that the value will be validated.
type PayloadKeyType string

const (
	PayloadKeyTypeString     PayloadKeyType = "<string>"
	PayloadKeyTypeInteger    PayloadKeyType = "<integer>"
	PayloadKeyTypeReal       PayloadKeyType = "<real>"
	PayloadKeyTypeBoolean    PayloadKeyType = "<boolean>"
	PayloadKeyTypeDate       PayloadKeyType = "<date>"
	PayloadKeyTypeData       PayloadKeyType = "<data>"
	PayloadKeyTypeArray      PayloadKeyType = "<array>"
	PayloadKeyTypeDictionary PayloadKeyType = "<dictionary>"
	PayloadKeyTypeAny        PayloadKeyType = "<any>"
)

// Indicates the expected format of the string value of the key, supporting additional validation of the value.
type PayloadKeySubType string

const (
	PayloadKeySubTypeURL      PayloadKeySubType = "<url>"
	PayloadKeySubTypeHostname PayloadKeySubType = "<hostname>"
	PayloadKeySubTypeEmail    PayloadKeySubType = "<email>"
)

// Whether the key is required or optional.
type PayloadKeyPresence string

const (
	PayloadKeyPresenceRequired PayloadKeyPresence = "required"
	PayloadKeyPresenceOptional PayloadKeyPresence = "optional"
)

// For a configuration that will be combined, indicates how this key is combined with ones from other configurations.
// * boolean-or - multiple <boolean> values are combined using a logical OR operation
// * boolean-and - multiple <boolean> values are combined using a logical AND operation
// * number-min - multiple <integer> or <real> values are combined by using the smallest value
// * number-max - multiple <integer> or <real> values are combined by using the largest value
// * enum-first - multiple <string> values with a rangelist are combined by using the value whose position is first in the range list
// * enum-last - multiple <string> values with a rangelist are combined by using the value whose position is last in the range list
// * first - multiple values are combined by using the first value that is processed
// * array-append - multiple <array> values are combined by concatenating the values in each array into a new array
// * set-union - multiple <array> values are combined by returning the unique union of all values in each array
// * set-intersection - multiple <array> values are combined by returning the unique intersection of all values in each array
type PayloadKeyCombineType string

const (
	PayloadKeyCombineTypeBooleanOr       PayloadKeyCombineType = "boolean-or"
	PayloadKeyCombineTypeBooleanAnd      PayloadKeyCombineType = "boolean-and"
	PayloadKeyCombineTypeNumberMin       PayloadKeyCombineType = "number-min"
	PayloadKeyCombineTypeNumberMax       PayloadKeyCombineType = "number-max"
	PayloadKeyCombineTypeEnumFirst       PayloadKeyCombineType = "enum-first"
	PayloadKeyCombineTypeEnumLast        PayloadKeyCombineType = "enum-last"
	PayloadKeyCombineTypeFirst           PayloadKeyCombineType = "first"
	PayloadKeyCombineTypeArrayAppend     PayloadKeyCombineType = "array-append"
	PayloadKeyCombineTypeSetUnion        PayloadKeyCombineType = "set-union"
	PayloadKeyCombineTypeSetIntersection PayloadKeyCombineType = "set-intersection"
)

// The type of the dictionary value.
type DetailsType string

const (
	DetailsTypeString     DetailsType = "<string>"
	DetailsTypeInteger    DetailsType = "<integer>"
	DetailsTypeReal       DetailsType = "<real>"
	DetailsTypeBoolean    DetailsType = "<boolean>"
	DetailsTypeDate       DetailsType = "<date>"
	DetailsTypeData       DetailsType = "<data>"
	DetailsTypeArray      DetailsType = "<array>"
	DetailsTypeDictionary DetailsType = "<dictionary>"
	DetailsTypeAny        DetailsType = "<any>"
)

type Schema struct {
	// Title for this schema object.
	Title string `yaml:"title"`
	// Description of this schema object.
	Description string `yaml:"description,omitempty"`
	// Overall properties of the payload.
	Payload *Payload `yaml:"payload,omitempty"`
	// An array of payload keys.
	PayloadKeys []*PayloadKey `yaml:"payloadkeys,omitempty"`
	// An array of payload keys.
	ResponseKeys []*PayloadKey `yaml:"responsekeys,omitempty"`
	// An array of Remote Management status reason codes.
	Reasons []*Reasons `yaml:"reasons,omitempty"`
	// An array describing any status items that might be related to a configuration.
	RelatedStatusItems []*RelatedStatusItems `yaml:"related-status-items,omitempty"`
	// An array of additional notes about a payload. These are published to the open source repository.
	Notes []*Notes `yaml:"notes,omitempty"`
}

// Overall properties of the payload.
type Payload struct {
	// Type of the profile payload.
	PayloadType string `yaml:"payloadtype,omitempty"`
	// Type of the MDM command.
	RequestType string `yaml:"requesttype,omitempty"`
	// Type of the declaration payload.
	DeclarationType string `yaml:"declarationtype,omitempty"`
	// Type of the status payload.
	StatusItemType string `yaml:"statusitemtype,omitempty"`
	// Type of the credential asset data.
	CredentialType string `yaml:"credentialtype,omitempty"`
	// Identifies the range of supported OS versions that support the entire payload.
	SupportedOS *SupportedOS `yaml:"supportedOS,omitempty"`
	// Indicates how multiple configurations of the same type are applied. If set to 'single', then only one configuration will be applied. If set to 'multiple', then each configuration is applied separately. If set to 'combined', then all configurations are combined into a single effective configuration.
	Apply PayloadApply `yaml:"apply,omitempty"`
	// Indicates that this entire payload should be considered a beta release. It may change in an incompatible way prior to final release.
	Beta bool `yaml:"beta,omitempty"`
	// Description of the payload.
	Content string `yaml:"content,omitempty"`
}

// Identifies the range of supported OS versions that support the entire payload.
type SupportedOS struct {
	// Supported range on this OS.
	IOS *OS `yaml:"iOS,omitempty"`
	// Supported range on this OS.
	MacOS *OS `yaml:"macOS,omitempty"`
	// Supported range on this OS.
	TVOS *OS `yaml:"tvOS,omitempty"`
	// Supported range on this OS.
	VisionOS *OS `yaml:"visionOS,omitempty"`
	// Supported range on this OS.
	WatchOS *OS `yaml:"watchOS,omitempty"`
}

// Supported range on this OS.
type OS struct {
	// OS version where feature was introduced.
	Introduced string `yaml:"introduced,omitempty"`
	// OS version where feature was deprecated.
	Deprecated string `yaml:"deprecated,omitempty"`
	// OS version where feature was removed.
	Removed string `yaml:"removed,omitempty"`
	// The MDM protocol access rights required on the device to execute the command.
	AccessRights string `yaml:"accessrights,omitempty"`
	// Indicates whether multiple copies of the payload can be installed
	Multiple bool `yaml:"multiple,omitempty"`
	// Indicates whether the command or profile is supported on the device channel. If this key is present it overrides the the `devicechannel` key in the top-level payload !!(payload) key.
	DeviceChannel bool `yaml:"devicechannel,omitempty"`
	// indicates whether the command or profile is supported on the user channel. If this key is present it overrides the the `userchannel` key in the top-level payload !!(payload) key.
	UserChannel bool `yaml:"userchannel,omitempty"`
	// Indicates whether the command or profile can only be executed on supervised devices. If this key is present it overrides the the `supervised` key in the top-level payload !!(payload) key.
	Supervised bool `yaml:"supervised,omitempty"`
	// If True, the command can only be executed on devices provisioned in DEP.
	RequiresDEP bool `yaml:"requiresdep,omitempty"`
	// If True, the command can only be executed on devices with user-approved MDM enrollment.
	UserApprovedMDM bool `yaml:"userapprovedmdm,omitempty"`
	// If True, the profile can be installed manually by a user on the device.
	AllowManualInstall bool `yaml:"allowmanualinstall,omitempty"`
	// Array of allowed enrollment types.
	AllowedEnrollments []string `yaml:"allowed-enrollments,omitempty"`
	// Array of allowed scopes.
	AllowedScopes []string `yaml:"allowed-scopes,omitempty"`
	// Additional behavior specific to shared iPad devices.
	SharediPad *SharediPad `yaml:"sharedipad,omitempty"`
	// Additional behavior when user enrollment is in effect. If this key is not present, then the corresponding payload or payload key can be used both with or without user enrollment in effect, without any changes to normal behavior.
	UserEnrollment *UserEnrollment `yaml:"userenrollment,omitempty"`
	// If true, indicates that the skip key's corresponding Setup pane is always skipped. If false, indicates that the skip key's corresponding Setup pane may be shown, depending on exactly when during the setup flow it occurs.
	AlwaysSkippable bool `yaml:"always-skippable,omitempty"`
	// Indicates that this payload should be considered a beta release for this OS. It may change in an incompatible way prior to final release.
	Beta bool `yaml:"beta,omitempty"`
}

// Additional behavior specific to shared iPad devices.
type SharediPad struct {
	// Indicates whether a payload or payload key can used with or without shared iPad in effect. If set to 'allowed', then the payload or payload key can be used both with or without shared iPad in effect. If set to 'required', then the payload or payload key can only be used if shared iPad is in effect. If set to 'forbidden', then the payload or payload key cannot be used if shared iPad is in effect. If set to 'ignored', then the payload or payload key can be used, but is ignored if shared iPad is in effect.
	Mode SharediPadMode `yaml:"mode,omitempty"`
	// Defines if the payload can be installed on the device MDM channel.
	DeviceChannel bool `yaml:"devicechannel,omitempty"`
	// Defines if the payload can be installed on the user MDM channel.
	UserChannel bool `yaml:"userchannel,omitempty"`
	// Array of allowed scopes.
	AllowedScopes []string `yaml:"allowed-scopes,omitempty"`
}

// Additional behavior when user enrollment is in effect. If this key is not present, then the corresponding payload or payload key can be used both with or without user enrollment in effect, without any changes to normal behavior.
type UserEnrollment struct {
	// Indicates how a payload or payload key can only be used if user enrollment is in effect. If set to 'allowed', then the payload or payload key can be used both with or without user enrollment in effect. If set to 'required', then the payload or payload key can only be used if user enrollment is in effect. If set to 'forbidden', then the payload or payload key cannot be used if user enrollment is in effect. If set to 'ignored', then the payload or payload key can be used, but is ignored if user enrollment is in effect.
	Mode UserEnrollmentMode `yaml:"mode,omitempty"`
	// Describes any special behavior for the payload or payload key if user enrollment is in effect.
	Behavior string `yaml:"behavior,omitempty"`
}

// A single payload key.
type PayloadKey struct {
	// The name of the key.
	Key string `yaml:"key"`
	// The title of the key.
	Title string `yaml:"title,omitempty"`
	// Identifies the range of supported OS versions that support the entire payload.
	SupportedOS *SupportedOS `yaml:"supportedOS,omitempty"`
	// The type of key. The value `<any>` may be used to indicate that any of the standard values can be used without any expectation that the value will be validated.
	Type PayloadKeyType `yaml:"type"`
	// Indicates the expected format of the string value of the key, supporting additional validation of the value.
	SubType PayloadKeySubType `yaml:"subtype,omitempty"`
	// Indicates the set of allowed asset types.
	AssetTypes []string `yaml:"assettypes,omitempty"`
	// Whether the key is required or optional.
	Presence PayloadKeyPresence `yaml:"presence,omitempty"`
	// List of allowed values for this key.
	Rangelist []yamlschema.IntegerNumberString `yaml:"rangelist,omitempty"`
	// Bounds for the value of this key.
	Range *Range `yaml:"range,omitempty"`
	// The default value (if any) for the key.
	Default yamlschema.BooleanIntegerNumberString `yaml:"default,omitempty"`
	// The format for the value expressed as a regular expression.
	Format string `yaml:"format,omitempty"`
	// Cardinality for this value.
	Repetition *Repetition `yaml:"repetition,omitempty"`
	// For a configuration that will be combined, indicates how this key is combined with ones from other configurations.
	// * boolean-or - multiple <boolean> values are combined using a logical OR operation
	// * boolean-and - multiple <boolean> values are combined using a logical AND operation
	// * number-min - multiple <integer> or <real> values are combined by using the smallest value
	// * number-max - multiple <integer> or <real> values are combined by using the largest value
	// * enum-first - multiple <string> values with a rangelist are combined by using the value whose position is first in the range list
	// * enum-last - multiple <string> values with a rangelist are combined by using the value whose position is last in the range list
	// * first - multiple values are combined by using the first value that is processed
	// * array-append - multiple <array> values are combined by concatenating the values in each array into a new array
	// * set-union - multiple <array> values are combined by returning the unique union of all values in each array
	// * set-intersection - multiple <array> values are combined by returning the unique intersection of all values in each array
	CombineType PayloadKeyCombineType `yaml:"combinetype,omitempty"`
	// Description of the payload key.
	Content string `yaml:"content,omitempty"`
	// A name that uniquely represents the structured subkey object. This is used when structured subkeys are referenced multiple times.
	SubKeyType string `yaml:"subkeytype,omitempty"`
	// An array of payload keys.
	SubKeys []*PayloadKey `yaml:"subkeys,omitempty"`
}

// Bounds for the value of this key.
type Range struct {
	// Lower bound.
	Min yamlschema.IntegerNumber `yaml:"min,omitempty"`
	// Upper bound.
	Max yamlschema.IntegerNumber `yaml:"max,omitempty"`
}

// Cardinality for this value.
type Repetition struct {
	// Lower bound.
	Min int64 `yaml:"min"`
	// Upper bound.
	Max int64 `yaml:"max"`
}

// An Remote Management reason code.
type Reasons struct {
	// The Remote Management reason code.
	Value string `yaml:"value,omitempty"`
	// Description of the Remote Management reason code.
	Description string `yaml:"description,omitempty"`
	// Keys defined in the Details dictionary
	Details []*Details `yaml:"details,omitempty"`
}

// Details dictionary keys
type Details struct {
	// The name of the dictionary key.
	Key string `yaml:"key,omitempty"`
	// Description of the dictionary item.
	Description string `yaml:"description,omitempty"`
	// The type of the dictionary value.
	Type DetailsType `yaml:"type,omitempty"`
}

// An additional note about a payload. A note is written in "markdown" and can be transformed to HTML if needed.
type RelatedStatusItems struct {
	// List of status-item types.
	StatusItems []string `yaml:"status-items"`
	// A description of the relationship.
	Note string `yaml:"note,omitempty"`
}

// An additional note about a payload. A note is written in "markdown" and can be transformed to HTML if needed.
type Notes struct {
	// Title for the note.
	Title string `yaml:"title"`
	// The note content in "markdown" format.
	Content string `yaml:"content"`
}
