// DO NOT EDIT
// generated from https://github.com/apple/device-management.git:b838baacf2e790db729b6ca3f52724adc8bfb96d/declarative/declarations/configurations

package configurations

const DeviceManagementGenerateHash = "b838baacf2e790db729b6ca3f52724adc8bfb96d"

// The set of protocol types to enable on the Exchange server, in order of preference. This is an array of unique strings with possible values:
// * 'EAS:' Exchange ActiveSync
// * 'EWS:' Exchange Web Services (EWS)
// If the device supports one or more of the listed protocol types, it sets up an account for the first supported type.
// If the device doesn't support any of the listed protocol types, it doesn't set up an account and the system reports an error.
type AccountExchangeEnabledProtocolTypes string

const (
	AccountExchangeEnabledProtocolTypesEAS AccountExchangeEnabledProtocolTypes = "EAS"
	AccountExchangeEnabledProtocolTypesEWS AccountExchangeEnabledProtocolTypes = "EWS"
)

// The type of recursion to use in the search.
// * 'Base': Only the 'SearchBase' node.
// * 'OneLevel': The 'SearchBase' node and its immediate children.
// * 'Subtree': The 'SearchBase' node and all its children, regardless of depth.
type SearchSettingsItemScope string

const (
	SearchSettingsItemScopeBase     SearchSettingsItemScope = "Base"
	SearchSettingsItemScopeOneLevel SearchSettingsItemScope = "OneLevel"
	SearchSettingsItemScopeSubtree  SearchSettingsItemScope = "Subtree"
)

// The mail protocol this account uses.
type IncomingServerServerType string

const (
	IncomingServerServerTypeIMAP IncomingServerServerType = "IMAP"
	IncomingServerServerTypePOP  IncomingServerServerType = "POP"
)

// The authentication method for the incoming mail server.
type IncomingServerAuthenticationMethod string

const (
	IncomingServerAuthenticationMethodNone     IncomingServerAuthenticationMethod = "None"
	IncomingServerAuthenticationMethodPassword IncomingServerAuthenticationMethod = "Password"
	IncomingServerAuthenticationMethodCRAMMD5  IncomingServerAuthenticationMethod = "CRAMMD5"
	IncomingServerAuthenticationMethodNTLM     IncomingServerAuthenticationMethod = "NTLM"
	IncomingServerAuthenticationMethodHTTPMD5  IncomingServerAuthenticationMethod = "HTTPMD5"
)

// The authentication method for the outgoing mail server.
type OutgoingServerAuthenticationMethod string

const (
	OutgoingServerAuthenticationMethodNone     OutgoingServerAuthenticationMethod = "None"
	OutgoingServerAuthenticationMethodPassword OutgoingServerAuthenticationMethod = "Password"
	OutgoingServerAuthenticationMethodCRAMMD5  OutgoingServerAuthenticationMethod = "CRAMMD5"
	OutgoingServerAuthenticationMethodNTLM     OutgoingServerAuthenticationMethod = "NTLM"
	OutgoingServerAuthenticationMethodHTTPMD5  OutgoingServerAuthenticationMethod = "HTTPMD5"
)

// The status the system reports back when the device implements the configuration. Use this to override the normal 'success' result.
type ManagementTestReturnStatus string

const (
	ManagementTestReturnStatusInstalled ManagementTestReturnStatus = "Installed"
	ManagementTestReturnStatusFailed    ManagementTestReturnStatus = "Failed"
	ManagementTestReturnStatusUnlocked  ManagementTestReturnStatus = "Unlocked"
)

// The type of display type to use for the connection.
// * Virtual1 — create one virtual display.
// * Virtual2 — create two virtual displays.
type DisplayConfigurationDisplayType string

const (
	DisplayConfigurationDisplayTypeVirtual1 DisplayConfigurationDisplayType = "Virtual1"
	DisplayConfigurationDisplayTypeVirtual2 DisplayConfigurationDisplayType = "Virtual2"
)

// Use this configuration to define settings for access to CalDAV servers.
// A CalDAV configuration defines a CalDAV calendar and reminders account for a user.
type AccountCalDAV struct {
	// The name that apps show to the user for this calendar account. If not present, the system generates a suitable default.
	VisibleName *string `json:"VisibleName,omitempty"`
	// The hostname of the CalDAV server (or IP address).
	HostName string `json:"HostName"`
	// The port number for the CalDAV server.
	Port *int64 `json:"Port,omitempty"`
	// The path for the CalDAV server.
	Path *string `json:"Path,omitempty"`
	// The identifier of an asset declaration that contains the credentials for this account. The corresponding asset must be of type 'CredentialUserNameAndPassword'.
	AuthenticationCredentialsAssetReference *string `json:"AuthenticationCredentialsAssetReference,omitempty"`
}

func (p *AccountCalDAV) DeclarationType() string {
	return "com.apple.configuration.account.caldav"
}

// Use this configuration to define settings for access to CardDAV servers.
// A CardDAV configuration defines a CardDAV contacts account for a user.
type AccountCardDAV struct {
	// The name that apps show to the user for this address book account. If not present, the system generates a suitable default.
	VisibleName *string `json:"VisibleName,omitempty"`
	// The hostname of the CardDAV server (or IP address).
	HostName string `json:"HostName"`
	// The port number for the CardDAV server.
	Port *int64 `json:"Port,omitempty"`
	// The path for the CardDAV server.
	Path *string `json:"Path,omitempty"`
	// The identifier of an asset declaration that contains the credentials for this account. The corresponding asset must be of type 'CredentialUserNameAndPassword'.
	AuthenticationCredentialsAssetReference *string `json:"AuthenticationCredentialsAssetReference,omitempty"`
}

func (p *AccountCardDAV) DeclarationType() string {
	return "com.apple.configuration.account.carddav"
}

// The configuration settings for OAuth for this account.
type OAuth struct {
	// If 'true', enables OAuth for this account.
	Enabled bool `json:"Enabled"`
	// The URL that this account uses for signing in with OAuth. The system ignores this value unless 'Enabled' is 'true'. The system doesn't use autodiscovery when a declaration contains this URL, so the declaration must also contain a 'HostName'.
	SignInURL *string `json:"SignInURL,omitempty"`
	// The URL that this account uses for token requests with OAuth. The system ignores this value unless 'Enabled' is 'true'.
	TokenRequestURL *string `json:"TokenRequestURL,omitempty"`
}

// Settings for S/MIME signing.
type AccountExchangeSMIMESigning struct {
	// It true, S/MIME signing is enabled.
	Enabled bool `json:"Enabled"`
	// Specifies the identifier of an asset declaration containing the identity required for S/MIME signing of messages sent from this account.
	IdentityAssetReference *string `json:"IdentityAssetReference,omitempty"`
	// If true, the user can turn S/MIME signing on or off in Settings.
	UserOverrideable *bool `default:"false" json:"UserOverrideable,omitempty"`
	// If true, the user can turn select an S/MIME signing identity in Settings.
	IdentityUserOverrideable *bool `default:"false" json:"IdentityUserOverrideable,omitempty"`
}

// Settings for S/MIME encryption.
type AccountExchangeSMIMEEncryption struct {
	// It true, S/MIME encryption by default is enabled. Cannot be overridden by the user if the "PerMessageSwitchEnabled" key is set to false.
	Enabled bool `json:"Enabled"`
	// Specifies the identifier of an asset declaration containing the identity required for S/MIME encryption. The public certificate is attached to outgoing mail to allow encrypted mail to be sent to this user. When the user sends encrypted mail, the public certificate is used to encrypt the copy of the mail in their Sent mailbox.
	IdentityAssetReference *string `json:"IdentityAssetReference,omitempty"`
	// If true, the user can turn S/MIME encryption by default on or off in Settings.
	UserOverrideable *bool `default:"false" json:"UserOverrideable,omitempty"`
	// If true, the user can turn select an S/MIME signing identity in Settings.
	IdentityUserOverrideable *bool `default:"false" json:"IdentityUserOverrideable,omitempty"`
	// If true, enables the per-message encryption switch in the compose view.
	PerMessageSwitchEnabled *bool `default:"false" json:"PerMessageSwitchEnabled,omitempty"`
}

// Settings for S/MIME.
type AccountExchangeSMIME struct {
	// Settings for S/MIME signing.
	Signing *AccountExchangeSMIMESigning `json:"Signing,omitempty"`
	// Settings for S/MIME encryption.
	Encryption *AccountExchangeSMIMEEncryption `json:"Encryption,omitempty"`
}

// Use this configuration to define settings for access to Exchange ActiveSync and Web Services servers.
// This payload configures an Exchange ActiveSync account on an iOS device.
type AccountExchange struct {
	// The name that apps show to the user for this Exchange account. If not present, the system generates a suitable default.
	VisibleName *string `json:"VisibleName,omitempty"`
	// The set of protocol types to enable on the Exchange server, in order of preference. This is an array of unique strings with possible values:
	// * 'EAS:' Exchange ActiveSync
	// * 'EWS:' Exchange Web Services (EWS)
	// If the device supports one or more of the listed protocol types, it sets up an account for the first supported type.
	// If the device doesn't support any of the listed protocol types, it doesn't set up an account and the system reports an error.
	EnabledProtocolTypes []string `json:"EnabledProtocolTypes"`
	// The identifier of an asset declaration that contains the user identity for this account. The corresponding asset must be of type 'UserIdentity'.
	UserIdentityAssetReference *string `json:"UserIdentityAssetReference,omitempty"`
	// The hostname of the EWS server (or IP address). This is a required field unless the declaration contains an 'OAuth' property, with a 'SignInURL' that has 'enabled' as 'true'.
	HostName *string `json:"HostName,omitempty"`
	// The port number of the EWS server. The system uses this only when this declaration has a 'HostName' value.
	Port *int64 `json:"Port,omitempty"`
	// The path of the EWS server. The system uses this only when this declaration has a 'HostName' value.
	Path *string `json:"Path,omitempty"`
	// The external hostname of the EWS server (or IP address). This is a required field unless the declaration contains an 'OAuth' property, with a 'SignInURL' that has 'enabled' as 'true'.
	ExternalHostName *string `json:"ExternalHostName,omitempty"`
	// The external port number of the EWS server. The system uses this only when this declaration has a 'HostName' value.
	ExternalPort *int64 `json:"ExternalPort,omitempty"`
	// The external path of the EWS server. The system uses this only when this declaration has a 'HostName' value.
	ExternalPath *string `json:"External Path,omitempty"`
	// The configuration settings for OAuth for this account.
	OAuth *OAuth `json:"OAuth,omitempty"`
	// The identifier of an asset declaration that contains the credentials for this account to authenticate with an Exchange server. The corresponding asset must be of type 'CredentialUserNameAndPassword'.
	AuthenticationCredentialsAssetReference *string `json:"AuthenticationCredentialsAssetReference,omitempty"`
	// The identifier of an asset declaration that contains the credentials for this account to authenticate with an Exchange server. The corresponding asset must be of type CredentialUserNameAndPassword.
	AuthenticationIdentityAssetReference *string `json:"AuthenticationIdentityAssetReference,omitempty"`
	// Settings for S/MIME.
	SMIME *AccountExchangeSMIME `json:"SMIME,omitempty"`
	// If 'true', activates the mail service for this account.
	MailServiceActive *bool `default:"true" json:"MailServiceActive,omitempty"`
	// If 'true', prevents the user from changing the status of the mail service for this account.
	LockMailService *bool `default:"false" json:"LockMailService,omitempty"`
	// If 'true', activates the address book service for this account.
	ContactsServiceActive *bool `default:"true" json:"ContactsServiceActive,omitempty"`
	// If 'true', prevents the user from changing the status of the address book service for this account.
	LockContactsService *bool `default:"false" json:"LockContactsService,omitempty"`
	// If 'true', activates the calendar service for this account.
	CalendarServiceActive *bool `default:"true" json:"CalendarServiceActive,omitempty"`
	// If 'true', prevents the user from changing the status of the calendar service for this account.
	LockCalendarService *bool `default:"false" json:"LockCalendarService,omitempty"`
	// If 'true', activates the reminders service for this account.
	RemindersServiceActive *bool `default:"true" json:"RemindersServiceActive,omitempty"`
	// If 'true', prevents the user from changing the status of the reminders service for this account.
	LockRemindersService *bool `default:"false" json:"LockRemindersService,omitempty"`
	// If 'true', activates the notes service for this account.
	NotesServiceActive *bool `default:"true" json:"NotesServiceActive,omitempty"`
	// If 'true', prevents the user from changing the status of the notes service for this account.
	LockNotesService *bool `default:"false" json:"LockNotesService,omitempty"`
}

func (p *AccountExchange) DeclarationType() string {
	return "com.apple.configuration.account.exchange"
}

// Use this configuration to define settings for access to Google services.
// A Google configuration defines a Google account for a user. The user will be prompted to enter their credentials shortly after the configuration successfully installs.
type AccountGoogle struct {
	// The name that apps show to the user for this Google account. If not present, the system generates a suitable default.
	VisibleName *string `json:"VisibleName,omitempty"`
	// The identifier of an asset declaration that contains the user identity for this Google account. The corresponding asset must be of type 'UserIdentity'. The asset must contain an 'EmailAddress' key that specifies the full Google email address for the account.
	UserIdentityAssetReference string `json:"UserIdentityAssetReference"`
}

func (p *AccountGoogle) DeclarationType() string {
	return "com.apple.configuration.account.google"
}

type SearchSettingsItem struct {
	// The description of this search setting in the Contacts and Settings apps. If not present, the apps display no name.
	VisibleDescription *string `json:"VisibleDescription,omitempty"`
	// The path to the node where a search starts. For example, 'ou=people,o=example corp'.
	SearchBase string `json:"SearchBase"`
	// The type of recursion to use in the search.
	// * 'Base': Only the 'SearchBase' node.
	// * 'OneLevel': The 'SearchBase' node and its immediate children.
	// * 'Subtree': The 'SearchBase' node and all its children, regardless of depth.
	Scope *string `default:"Subtree" json:"Scope,omitempty"`
}

// Use this configuration to define settings for access to LDAP servers.
// An LDAP configuration defines an LDAP directory account for a user.
type AccountLDAP struct {
	// The name that apps show to the user for this LDAP account. If not present, the system generates a suitable default.
	VisibleName *string `json:"VisibleName,omitempty"`
	// The hostname of the LDAP server (or IP address).
	HostName string `json:"HostName"`
	// The port number of the LDAP server (or IP address).
	Port *int64 `json:"Port,omitempty"`
	// The identifier of an asset declaration that contains the credentials for this account. The corresponding asset must be of type 'CredentialUserNameAndPassword'.
	AuthenticationCredentialsAssetReference *string `json:"AuthenticationCredentialsAssetReference,omitempty"`
	// The array of nodes to start LDAP searches from. There must be at least one node for this account to be useful. macOS only searches one node and ignores other items in the array.
	SearchSettings *[]*SearchSettingsItem `json:"SearchSettings,omitempty"`
}

func (p *AccountLDAP) DeclarationType() string {
	return "com.apple.configuration.account.ldap"
}

// The settings for the incoming mail server for this account.
type IncomingServer struct {
	// The mail protocol this account uses.
	ServerType string `json:"ServerType"`
	// The host name for the incoming mail server.
	HostName string `json:"HostName"`
	// The port number for the incoming mail server.
	Port *int64 `json:"Port,omitempty"`
	// The authentication method for the incoming mail server.
	AuthenticationMethod string `json:"AuthenticationMethod"`
	// The identifier of an asset declaration that contains the credentials for this account to authenticate with an incoming mail server. The corresponding asset must be of type 'CredentialUserNameAndPassword'.
	// If the 'AuthenticationMethod' is 'None', this field must be blank. Otherwise, the declaration must contain this field.
	AuthenticationCredentialsAssetReference *string `json:"AuthenticationCredentialsAssetReference,omitempty"`
	// The path prefix for the IMAP server. The system uses this only when 'ServerType' is 'IMAP'.
	IMAPPathPrefix *string `json:"IMAPPathPrefix,omitempty"`
}

// The settings for the outgoing mail server for this account.
type OutgoingServer struct {
	// The host name for the outgoing mail server.
	HostName string `json:"HostName"`
	// The port number for the outgoing mail server.
	Port *int64 `json:"Port,omitempty"`
	// The authentication method for the outgoing mail server.
	AuthenticationMethod string `json:"AuthenticationMethod"`
	// The identifier of an asset declaration that contains the credentials for this account to authenticate with an outgoing mail server. The corresponding asset must be of type 'CredentialUserNameAndPassword'.
	// If the 'AuthenticationMethod' is 'None', this field must be blank. Otherwise, the declaration must contain this field.
	AuthenticationCredentialsAssetReference *string `json:"AuthenticationCredentialsAssetReference,omitempty"`
}

// Settings for S/MIME signing.
type AccountMailSMIMESigning struct {
	// It true, S/MIME signing is enabled.
	Enabled bool `json:"Enabled"`
	// Specifies the identifier of an asset declaration containing the identity required for S/MIME signing of messages sent from this account.
	IdentityAssetReference *string `json:"IdentityAssetReference,omitempty"`
	// If true, the user can turn S/MIME signing on or off in Settings.
	UserOverrideable *bool `default:"false" json:"UserOverrideable,omitempty"`
	// If true, the user can turn select an S/MIME signing identity in Settings.
	IdentityUserOverrideable *bool `default:"false" json:"IdentityUserOverrideable,omitempty"`
}

// Settings for S/MIME encryption.
type AccountMailSMIMEEncryption struct {
	// It true, S/MIME encryption by default is enabled. Cannot be overridden by the user if the "PerMessageSwitchEnabled" key is set to false.
	Enabled bool `json:"Enabled"`
	// Specifies the identifier of an asset declaration containing the identity required for S/MIME encryption. The public certificate is attached to outgoing mail to allow encrypted mail to be sent to this user. When the user sends encrypted mail, the public certificate is used to encrypt the copy of the mail in their Sent mailbox.
	IdentityAssetReference *string `json:"IdentityAssetReference,omitempty"`
	// If true, the user can turn S/MIME encryption by default on or off in Settings.
	UserOverrideable *bool `default:"false" json:"UserOverrideable,omitempty"`
	// If true, the user can turn select an S/MIME signing identity in Settings.
	IdentityUserOverrideable *bool `default:"false" json:"IdentityUserOverrideable,omitempty"`
	// If true, enables the per-message encryption switch in the compose view.
	PerMessageSwitchEnabled *bool `default:"false" json:"PerMessageSwitchEnabled,omitempty"`
}

// Settings for S/MIME.
type AccountMailSMIME struct {
	// Settings for S/MIME signing.
	Signing *AccountMailSMIMESigning `json:"Signing,omitempty"`
	// Settings for S/MIME encryption.
	Encryption *AccountMailSMIMEEncryption `json:"Encryption,omitempty"`
}

// Use this configuration to define settings for access to email servers.
// An email configuration defines an email account for a user.
type AccountMail struct {
	// The name that apps show to the user for this mail account. If not present, the system generates a suitable default.
	VisibleName *string `json:"VisibleName,omitempty"`
	// The identifier of an asset declaration that contains the user identity for this account. The corresponding asset must be of type 'UserIdentity'.
	UserIdentityAssetReference *string `json:"UserIdentityAssetReference,omitempty"`
	// The settings for the incoming mail server for this account.
	IncomingServer *IncomingServer `json:"IncomingServer"`
	// The settings for the outgoing mail server for this account.
	OutgoingServer *OutgoingServer `json:"OutgoingServer"`
	// Settings for S/MIME.
	SMIME *AccountMailSMIME `json:"SMIME,omitempty"`
}

func (p *AccountMail) DeclarationType() string {
	return "com.apple.configuration.account.mail"
}

// Use this configuration to define settings for a subscribed calendar.
// A subscribed calendar configuration defines a subscribed calendar for a user.
type AccountSubscribedCalendar struct {
	// The name that apps show to the user for this calendar account. If not present, the system generates a suitable default.
	VisibleName *string `json:"VisibleName,omitempty"`
	// The URL of the subscribed calendar. The URL must start with 'https://'.
	CalendarURL string `json:"CalendarURL"`
	// The identifier of an asset declaration that contains the credentials for this account to authenticate with a calendar server. The corresponding asset must be of type 'CredentialUserNameAndPassword'.
	AuthenticationCredentialsAssetReference *string `json:"AuthenticationCredentialsAssetReference,omitempty"`
}

func (p *AccountSubscribedCalendar) DeclarationType() string {
	return "com.apple.configuration.account.subscribed-calendar"
}

// Specifies an MDMv1 profile to present to the user who may choose to download and install it
type LegacyInteractiveProfile struct {
	// The URL of the profile to download and install. This must be an 'https://' URL that is hosted by the MDM server. When fetching the profile data, the device will use standard MDM authentication for the HTTP request.
	// If a user enrollment triggers this configuration, the system silently ignores any MDM 1 payloads in macOS where the User Enrollment Mode setting is 'forbidden'. In iOS and tvOS, the system rejects the entire profile.
	ProfileURL string `json:"ProfileURL"`
	// The visible name of the configuration. This name needs to indicate the nature of the profile.
	VisibleName string `json:"VisibleName"`
}

func (p *LegacyInteractiveProfile) DeclarationType() string {
	return "com.apple.configuration.legacy.interactive"
}

// Specifies an MDMv1 profile to download and install
type LegacyProfile struct {
	// The URL of the profile to download and install. This must be an 'https://' URL that is hosted by the MDM server.When fetching the profile data, the device will use standard MDM authentication for the HTTP request.
	// If a user enrollment triggers this configuration, the system silently ignores any MDM 1 payloads in macOS where the User Enrollment Mode setting is 'forbidden'. In iOS and tvOS, the system rejects the entire profile.
	ProfileURL string `json:"ProfileURL"`
}

func (p *LegacyProfile) DeclarationType() string {
	return "com.apple.configuration.legacy"
}

type StatusItem struct {
	// The name of the status item to send to subscribers.
	Name string `json:"Name"`
}

// Use this configuration to define the status subscriptions that cause status to be reported by the client.
type ManagementStatusSubscriptions struct {
	// An array of status items that the device notifies subscribers about.
	StatusItems []*StatusItem `json:"StatusItems"`
}

func (p *ManagementStatusSubscriptions) DeclarationType() string {
	return "com.apple.configuration.management.status-subscriptions"
}

// A configuration used for testing only
type ManagementTest struct {
	// The string to echo back in a status response reason.
	Echo string `json:"Echo"`
	// String read from a data asset to echo back in status response reason description.
	EchoDataAssetReference *string `json:"EchoDataAssetReference,omitempty"`
	// The status the system reports back when the device implements the configuration. Use this to override the normal 'success' result.
	ReturnStatus *string `default:"Installed" json:"ReturnStatus,omitempty"`
}

func (p *ManagementTest) DeclarationType() string {
	return "com.apple.configuration.management.test"
}

// Contains a dictionary whose keys are supported OS language IDs (e.g. "en-US"), and whose values represent a localized description of the policy the regular expression enforces. The special "default" key can be used for languages that are not contained in the dictionary.
type Description struct {
	// A localized description.
	ANY *string `json:"ANY,omitempty"`
}

// Specifies a regular expression, and its description, to be used to enforce password compliance. Note that we strongly recommend using the simpler passcode settings whenever possible, and rely on regular expression matching only when necessary. Mistakes in regular expressions can lead to frustrating user experiences, such as unsatisfiable passcode policies, or policy descriptions that don't match the enforced policy.
type CustomRegex struct {
	// A regular expression string that is matched against the password to determine whether it complies with a policy. The regular expression uses the ICU syntax (https://unicode-org.github.io/icu/userguide/strings/regexp.html). The string must not exceed 2048 characters in length.
	Regex string `json:"Regex"`
	// Contains a dictionary whose keys are supported OS language IDs (e.g. "en-US"), and whose values represent a localized description of the policy the regular expression enforces. The special "default" key can be used for languages that are not contained in the dictionary.
	Description *Description `json:"Description,omitempty"`
}

// Use this configuration to define passcode policy settings
type PasscodeSettings struct {
	// If 'true', requires the user to set a passcode without any requirements about the length or quality of the passcode. The presence of any other keys implicitly requires a passcode, and overrides this key's value.
	RequirePasscode *bool `default:"false" json:"RequirePasscode,omitempty"`
	// If set to true, the passcode must consist of at least one alphabetic characters ("abcd"), and at least one number.
	RequireAlphanumericPasscode *bool `default:"false" json:"RequireAlphanumericPasscode,omitempty"`
	// If 'true', requires a complex passcode. A complex passcode is one that doesn't contain repeated characters or increasing/decreasing characters (such as 123 or CBA).
	RequireComplexPasscode *bool `default:"false" json:"RequireComplexPasscode,omitempty"`
	// The minimum number of characters a passcode can contain.
	MinimumLength *int64 `default:"0" json:"MinimumLength,omitempty"`
	// Specifies the minimum number of complex characters that must be present. A "complex" character is a character other than a number or a letter, such as &%$#.
	MinimumComplexCharacters *int64 `default:"0" json:"MinimumComplexCharacters,omitempty"`
	// The number of failed passcode attempts that the system allows the user before iOS erases the device or macOS locks the device. If you don't change this setting, after six failed attempts, the device imposes a time delay before the user can enter a passcode again. The time delay increases with each failed attempt.
	// After the final failed attempt, the system securely erases all data and settings from the iOS device. A macOS device locks after the final attempt. The passcode time delay begins after the sixth attempt, so if this value is six or lower, the system has no time delay and triggers the erase or lock as soon as the user exceeds the limit.
	MaximumFailedAttempts *int64 `default:"11" json:"MaximumFailedAttempts,omitempty"`
	// The number of minutes before the login will be reset after the maximum number of failed attempts has been reached. The MaximumFailedAttempts key must be set for this to take effect.
	FailedAttemptsResetInMinutes *int64 `json:"FailedAttemptsResetInMinutes,omitempty"`
	// The maximum period that a user can select, during which the user can unlock the device without a passcode. A value of '0' means no grace period, and the device requires a passcode immediately. In the absence of this key, the user can select any period.
	// macOS translates this to screensaver settings.
	MaximumGracePeriodInMinutes *int64 `json:"MaximumGracePeriodInMinutes,omitempty"`
	// The maximum period that a user can select, during which the device can be idle before the system automatically locks it. When the device reaches this limit, the device locks and the user must enter the passcode to unlock it. In the absence of this key, the user can select any period.
	// macOS translates this to screensaver settings.
	MaximumInactivityInMinutes *int64 `json:"MaximumInactivityInMinutes,omitempty"`
	// Specifies the maximum number of days for which the passcode can remain unchanged. After this number of days, the user is forced to change the passcode before the device is unlocked.
	MaximumPasscodeAgeInDays *int64 `json:"MaximumPasscodeAgeInDays,omitempty"`
	// The number of historical passcode entries the system checks when validating a new passcode. The device refuses a new passcode if it matches a previously used passcode within the specified passcode history range. In the absence of this key, the system performs no historical check.
	PasscodeReuseLimit *int64 `json:"PasscodeReuseLimit,omitempty"`
	// If set to true, forces a password reset to occur the next time the user tries to authenticate. If this key is set in a configuration in the system scope (device channel), the setting takes effect for all users, and admin authentications may fail until the admin user password is also reset.
	ChangeAtNextAuth *bool `default:"false" json:"ChangeAtNextAuth,omitempty"`
	// Specifies a regular expression, and its description, to be used to enforce password compliance. Note that we strongly recommend using the simpler passcode settings whenever possible, and rely on regular expression matching only when necessary. Mistakes in regular expressions can lead to frustrating user experiences, such as unsatisfiable passcode policies, or policy descriptions that don't match the enforced policy.
	CustomRegex *CustomRegex `json:"CustomRegex,omitempty"`
}

func (p *PasscodeSettings) DeclarationType() string {
	return "com.apple.configuration.passcode.settings"
}

// Use this configuration to define a group of Screen Sharing connections.
type ScreenSharingConnectionGroup struct {
	// A string which uniquely identifies this connection group.
	ConnectionGroupUUID string `json:"ConnectionGroupUUID"`
	// The name of the Connection Group.
	GroupName string `json:"GroupName"`
	// Array of ConnectionUUIDs (matching a connection declared in a
	// com.apple.configuration.screensharing.connection configuration) of the Connections
	// that should be members of this group.
	Members []string `json:"Members"`
}

func (p *ScreenSharingConnectionGroup) DeclarationType() string {
	return "com.apple.configuration.screensharing.connection.group"
}

// The display configuration to use for this connection.
type DisplayConfiguration struct {
	// The type of display type to use for the connection.
	// * Virtual1 — create one virtual display.
	// * Virtual2 — create two virtual displays.
	DisplayType string `json:"DisplayType"`
}

// Use this configuration to define a connection to a Screen Sharing host.
type ScreenSharingConnection struct {
	// A string which uniquely identifies this connection. This is used to include a connection in a connection group.
	ConnectionUUID string `json:"ConnectionUUID"`
	// The name of the connection.
	DisplayName string `json:"DisplayName"`
	// The host name or IP address of the Mac that will host the screen sharing connection.
	HostName string `json:"HostName"`
	// Specifies the TCP port number on the host used to initiate the connection.
	Port *int64 `json:"Port,omitempty"`
	// The display configuration to use for this connection.
	DisplayConfiguration *DisplayConfiguration `json:"DisplayConfiguration"`
	// Specifies the identifier of an asset declaration containing the credentials required for this connection to authenticate with the Screen Sharing server. The corresponding asset must be of type "com.apple.asset.credential.userpassword".
	AuthenticationCredentialsAssetReference *string `json:"AuthenticationCredentialsAssetReference,omitempty"`
}

func (p *ScreenSharingConnection) DeclarationType() string {
	return "com.apple.configuration.screensharing.connection"
}

// Use this configuration to define Screen Sharing host settings and restrictions.
type ScreenSharingHostSettings struct {
	// Sets the maximum number of Virtual Displays to make available to clients.
	MaximumVirtualDisplays *int64 `json:"MaximumVirtualDisplays,omitempty"`
	// Specifies the initial UDP port number for connecting to the host. Screen Sharing needs multiple connections
	// so additional connections will increment this base port number by 1 for each needed connection. This does not
	// change the port number used to initially establish a connection with a host, which is always TCP port 5900.
	PortBase *int64 `json:"PortBase,omitempty"`
	// Set to true to prevent users from copying files from the Screen Sharing host.
	PreventCopyFilesFromHost *bool `default:"false" json:"PreventCopyFilesFromHost,omitempty"`
	// Set to true to prevent users from copying files to the Screen Sharing host.
	PreventCopyFilesToHost *bool `default:"false" json:"PreventCopyFilesToHost,omitempty"`
	// Set to true to prevent clients from establishing High Performance connections to the host.
	PreventHighPerformanceConnections *bool `default:"false" json:"PreventHighPerformanceConnections,omitempty"`
}

func (p *ScreenSharingHostSettings) DeclarationType() string {
	return "com.apple.configuration.screensharing.host.settings"
}

// Use this configuration to add a certificate to the device.
type SecurityCertificate struct {
	// Specifies the identifier of an asset declaration containing the certificate to be installed.
	CredentialAssetReference string `json:"CredentialAssetReference"`
}

func (p *SecurityCertificate) DeclarationType() string {
	return "com.apple.configuration.security.certificate"
}

// Use this configuration to install an identity on the device.
type SecurityIdentity struct {
	// Specifies the identifier of an asset declaration containing the identity to be installed.
	CredentialAssetReference string `json:"CredentialAssetReference"`
	// If set to "true", apps have access to the private key.
	AllowAllAppsAccess *bool `default:"false" json:"AllowAllAppsAccess,omitempty"`
	// If set to "true", the private key will be marked as extractable in the keychain.
	KeyIsExtractable *bool `default:"true" json:"KeyIsExtractable,omitempty"`
}

func (p *SecurityIdentity) DeclarationType() string {
	return "com.apple.configuration.security.identity"
}

// Configures the device to allow WebAuthn enterprise attestation for certain passkeys.
type SecurityPasskeyAttestation struct {
	// Specifies the identifier of an asset declaration containing the identity to be installed and used for passkey attestation.
	AttestationIdentityAssetReference string `json:"AttestationIdentityAssetReference"`
	// If set to "true", the private key for the attestation identity will be marked as extractable in the keychain.
	AttestationIdentityKeyIsExtractable *bool `default:"true" json:"AttestationIdentityKeyIsExtractable,omitempty"`
	// Relying parties to allow enterprise attestation.
	RelyingParties []string `json:"RelyingParties"`
}

func (p *SecurityPasskeyAttestation) DeclarationType() string {
	return "com.apple.configuration.security.passkey.attestation"
}

// Specifies managed configuration files for services
type ServicesConfigurationFiles struct {
	// The identifier of the system service whose managed configuration files are being specified.
	// This should be a reverse DNS style identifier. The "com.apple." prefix is reserved for built-in services.
	// The available built-in services are:
	// * com.apple.sshd - configures sshd
	// * com.apple.sudo - configures sudo
	// * com.apple.pam - configures PAM
	// * com.apple.cups - configures CUPS
	// * com.apple.apache.httpd - configures Apache httpd
	// * com.apple.bash - configures bash
	// * com.apple.zsh - configures zsh
	ServiceType string `json:"ServiceType"`
	// Specifies the identifier of an asset declaration containing a reference to the files to be used for the system service configuration. The corresponding asset must be of type "com.apple.asset.data". The referenced data must be a zip archive of an entire directory, that will be expanded and stored in a well known location for the system service. The asset's "ContentType" and "Hash-SHA-256" keys in the "Reference" key are required.
	DataAssetReference string `json:"DataAssetReference"`
}

func (p *ServicesConfigurationFiles) DeclarationType() string {
	return "com.apple.configuration.services.configuration-files"
}

// A software update enforcement policy for a specific OS release
type SoftwareUpdateEnforcementSpecific struct {
	// The target OS version that the device is required to update to by the appropriate time. This is the OS version number (e.g., "16.1"). A supplemental version identifier can be included (e.g., "16.1 (a)").
	TargetOSVersion string `json:"TargetOSVersion"`
	// The target build version that the device is required to update to by the appropriate time. This is the build version (e.g., "20A242). The build version is used for testing during seeding periods. A supplemental version identifier can be included (e.g., "20A242a"). If the build version is not consistent with the target OS version specified in the 'TargetOSVersion' key, the target OS version will take precedence.
	TargetBuildVersion *string `json:"TargetBuildVersion,omitempty"`
	// The local date time value when the software update will be force installed. This value must use the `yyyy-mm-ddThh:mm:ss` format derived from RFC3339 (https://www.rfc-editor.org/rfc/rfc3339.txt), but must not include a time zone offset. If the user does not trigger the software update before this time, the device will force install it.
	TargetLocalDateTime string `json:"TargetLocalDateTime"`
	// The URL of a web page that shows details, provided by the organization, about the enforced update.
	DetailsURL *string `json:"DetailsURL,omitempty"`
}

func (p *SoftwareUpdateEnforcementSpecific) DeclarationType() string {
	return "com.apple.configuration.softwareupdate.enforcement.specific"
}

// Specifies an MDMv1 Apple Watch enrollment profile
type WatchEnrollment struct {
	// The URL of the profile that the Apple Watch will download and install if the user opts in to management during the pairing process. This must be an "https" URL. The pairing iPhone must be supervised for this device enrollment to succeed. The profile must contain an MDM payload or it will be ignored and Apple Watch will not be enrolled. Apple Watch will attempt to install each payload contained in the profile.
	EnrollmentProfileURL string `json:"EnrollmentProfileURL"`
	// Specifies an array of identifiers of asset declarations containing anchor certificates to be used when evaluating the trust of the enrollment profile server. The corresponding assets must be of type "com.apple.asset.credential.certificate".
	AnchorCertificateAssetReferences *[]string `json:"AnchorCertificateAssetReferences,omitempty"`
}

func (p *WatchEnrollment) DeclarationType() string {
	return "com.apple.configuration.watch.enrollment"
}
