// DO NOT EDIT
// generated from https://github.com/apple/device-management.git:0a4527c5ea21825fd23e08273ccdb9e2302458ce/declarative/declarations/configurations

package configurations

const DeviceManagementGenerateHash = "0a4527c5ea21825fd23e08273ccdb9e2302458ce"

var DeclarationMap = map[string]any{
	"com.apple.configuration.account.caldav":                      AccountCalDAV{},
	"com.apple.configuration.account.carddav":                     AccountCardDAV{},
	"com.apple.configuration.account.exchange":                    AccountExchange{},
	"com.apple.configuration.account.google":                      AccountGoogle{},
	"com.apple.configuration.account.ldap":                        AccountLDAP{},
	"com.apple.configuration.account.mail":                        AccountMail{},
	"com.apple.configuration.account.subscribed-calendar":         AccountSubscribedCalendar{},
	"com.apple.configuration.app.managed":                         AppManaged{},
	"com.apple.configuration.audio-accessory.settings":            AudioAccessorySettings{},
	"com.apple.configuration.diskmanagement.settings":             DiskManagementSettings{},
	"com.apple.configuration.legacy":                              LegacyProfile{},
	"com.apple.configuration.legacy.interactive":                  LegacyInteractiveProfile{},
	"com.apple.configuration.management.status-subscriptions":     ManagementStatusSubscriptions{},
	"com.apple.configuration.management.test":                     ManagementTest{},
	"com.apple.configuration.math.settings":                       MathSettings{},
	"com.apple.configuration.package":                             Package{},
	"com.apple.configuration.passcode.settings":                   PasscodeSettings{},
	"com.apple.configuration.safari.bookmarks":                    SafariBookmarks{},
	"com.apple.configuration.safari.extensions.settings":          SafariExtensionSettings{},
	"com.apple.configuration.safari.settings":                     SafariSettings{},
	"com.apple.configuration.screensharing.connection":            ScreenSharingConnection{},
	"com.apple.configuration.screensharing.connection.group":      ScreenSharingConnectionGroup{},
	"com.apple.configuration.screensharing.host.settings":         ScreenSharingHostSettings{},
	"com.apple.configuration.security.certificate":                SecurityCertificate{},
	"com.apple.configuration.security.identity":                   SecurityIdentity{},
	"com.apple.configuration.security.passkey.attestation":        SecurityPasskeyAttestation{},
	"com.apple.configuration.services.background-tasks":           ServicesBackgroundTasks{},
	"com.apple.configuration.services.configuration-files":        ServicesConfigurationFiles{},
	"com.apple.configuration.softwareupdate.enforcement.specific": SoftwareUpdateEnforcementSpecific{},
	"com.apple.configuration.softwareupdate.settings":             SoftwareUpdateSettings{},
	"com.apple.configuration.watch.enrollment":                    WatchEnrollment{},
}

// The declaration to configure a Calendar account.
// A CalDAV configuration defines a CalDAV calendar and reminders account for a user.
type AccountCalDAV struct {
	// The name that apps show to the user for this calendar account. If not present, the system generates a suitable default.
	VisibleName *string `json:"VisibleName,omitempty" plist:"VisibleName,omitempty"`
	// The hostname or IP address of the CalDAV server.
	HostName string `json:"HostName" plist:"HostName" required:"true"`
	// The port number for the CalDAV server.
	Port *int64 `json:"Port,omitempty" plist:"Port,omitempty"`
	// The path for the CalDAV server.
	Path *string `json:"Path,omitempty" plist:"Path,omitempty"`
	// The identifier of an asset declaration that contains the credentials for this account. Set the corresponding asset type to `CredentialUserNameAndPassword`.
	AuthenticationCredentialsAssetReference *string `json:"AuthenticationCredentialsAssetReference,omitempty" plist:"AuthenticationCredentialsAssetReference,omitempty"`
}

func (p *AccountCalDAV) DeclarationType() string {
	return "com.apple.configuration.account.caldav"
}

// The declaration to configure a Contacts account.
// A CardDAV configuration defines a CardDAV contacts account for a user.
type AccountCardDAV struct {
	// The name that apps show to the user for this address book account. If not present, the system generates a suitable default.
	VisibleName *string `json:"VisibleName,omitempty" plist:"VisibleName,omitempty"`
	// The hostname or IP address of the CardDAV server.
	HostName string `json:"HostName" plist:"HostName" required:"true"`
	// The port number for the CardDAV server.
	Port *int64 `json:"Port,omitempty" plist:"Port,omitempty"`
	// The path for the CardDAV server.
	Path *string `json:"Path,omitempty" plist:"Path,omitempty"`
	// The identifier of an asset declaration that contains the credentials for this account. Set the corresponding asset type to `CredentialUserNameAndPassword`.
	AuthenticationCredentialsAssetReference *string `json:"AuthenticationCredentialsAssetReference,omitempty" plist:"AuthenticationCredentialsAssetReference,omitempty"`
}

func (p *AccountCardDAV) DeclarationType() string {
	return "com.apple.configuration.account.carddav"
}

// The declaration to configure an Exchange account.
// This payload configures an Exchange ActiveSync account on an iOS device.
type AccountExchange struct {
	// The name that apps show to the user for this Exchange account. If not present, the system generates a suitable default.
	VisibleName *string `json:"VisibleName,omitempty" plist:"VisibleName,omitempty"`
	// The set of protocol types to enable on the Exchange server, in order of preference. This is an array of unique strings with possible values:
	// - `EAS:` Exchange ActiveSync
	// - `EWS:` Exchange Web Services (EWS)
	// If the device supports one or more of the listed protocol types, it sets up an account for the first supported type.
	// If the device doesn't support any of the listed protocol types, it doesn't set up an account and the system reports an error.
	EnabledProtocolTypes []EnabledProtocolTypes `json:"EnabledProtocolTypes" plist:"EnabledProtocolTypes" required:"true"`
	// The identifier of an asset declaration that contains the user identity for this account. The corresponding asset must be of type `UserIdentity`.
	UserIdentityAssetReference *string `json:"UserIdentityAssetReference,omitempty" plist:"UserIdentityAssetReference,omitempty"`
	// The IP address or fully qualified domain name (FQDN) of the Exchange host.
	HostName *string `json:"HostName,omitempty" plist:"HostName,omitempty"`
	// The port number of the EWS server. The system uses this only when this declaration has a `HostName` value.
	Port *int64 `json:"Port,omitempty" plist:"Port,omitempty"`
	// The path of the EWS server. The system uses this only when this declaration has a `HostName` value.
	Path *string `json:"Path,omitempty" plist:"Path,omitempty"`
	// The external hostname of the EWS server (or IP address).
	ExternalHostName *string `json:"ExternalHostName,omitempty" plist:"ExternalHostName,omitempty"`
	// The external port number of the EWS server. The system uses this only when this declaration has a `ExternalHostName` value.
	ExternalPort *int64 `json:"ExternalPort,omitempty" plist:"ExternalPort,omitempty"`
	// The external path of the EWS server. The system uses this only when this declaration has a `ExternalHostName` value.
	ExternalPath *string `json:"External Path,omitempty" plist:"External Path,omitempty"`
	// The configuration settings for OAuth for this account.
	OAuth *OAuth `json:"OAuth,omitempty" plist:"OAuth,omitempty"`
	// The identifier of an asset declaration that contains the credentials for this account to authenticate with an Exchange server. Set the corresponding asset type to `CredentialUserNameAndPassword`.
	AuthenticationCredentialsAssetReference *string `json:"AuthenticationCredentialsAssetReference,omitempty" plist:"AuthenticationCredentialsAssetReference,omitempty"`
	// The identifier of a credential asset declaration that contains the identity that this account requires to authenticate with the Exchange server.
	AuthenticationIdentityAssetReference *string `json:"AuthenticationIdentityAssetReference,omitempty" plist:"AuthenticationIdentityAssetReference,omitempty"`
	// Settings for S/MIME.
	SMIME *AccountExchangeSMIME `json:"SMIME,omitempty" plist:"SMIME,omitempty"`
	// If `true`, the system activates the mail service for this account.
	MailServiceActive *bool `json:"MailServiceActive,omitempty" plist:"MailServiceActive,omitempty"`
	// If `true`, the system prevents the user from changing the status of the mail service for this account.
	LockMailService *bool `json:"LockMailService,omitempty" plist:"LockMailService,omitempty"`
	// If `true`, activates the address book service for this account.
	ContactsServiceActive *bool `json:"ContactsServiceActive,omitempty" plist:"ContactsServiceActive,omitempty"`
	// If `true`, the system prevents the user from changing the status of the address book service for this account.
	LockContactsService *bool `json:"LockContactsService,omitempty" plist:"LockContactsService,omitempty"`
	// If `true`, activates the calendar service for this account.
	CalendarServiceActive *bool `json:"CalendarServiceActive,omitempty" plist:"CalendarServiceActive,omitempty"`
	// If `true`, the system prevents the user from changing the status of the calendar service for this account.
	LockCalendarService *bool `json:"LockCalendarService,omitempty" plist:"LockCalendarService,omitempty"`
	// If `true`, the system activates the reminders service for this account.
	RemindersServiceActive *bool `json:"RemindersServiceActive,omitempty" plist:"RemindersServiceActive,omitempty"`
	// If `true`, the system prevents the user from changing the status of the reminders service for this account.
	LockRemindersService *bool `json:"LockRemindersService,omitempty" plist:"LockRemindersService,omitempty"`
	// If `true`, the system activates the notes service for this account.
	NotesServiceActive *bool `json:"NotesServiceActive,omitempty" plist:"NotesServiceActive,omitempty"`
	// If `true`, the system prevents the user from changing the status of the notes service for this account.
	LockNotesService *bool `json:"LockNotesService,omitempty" plist:"LockNotesService,omitempty"`
}

func (p *AccountExchange) DeclarationType() string {
	return "com.apple.configuration.account.exchange"
}

// The set of protocol types to enable on the Exchange server, in order of preference. This is an array of unique strings with possible values:
// - `EAS:` Exchange ActiveSync
// - `EWS:` Exchange Web Services (EWS)
// If the device supports one or more of the listed protocol types, it sets up an account for the first supported type.
// If the device doesn't support any of the listed protocol types, it doesn't set up an account and the system reports an error.
type EnabledProtocolTypes string

const (
	EnabledProtocolTypesEAS EnabledProtocolTypes = "EAS"
	EnabledProtocolTypesEWS EnabledProtocolTypes = "EWS"
)

// The configuration settings for OAuth for this account.
type OAuth struct {
	// If `true`, enables OAuth for this account.
	Enabled bool `json:"Enabled" plist:"Enabled" required:"true"`
	// The URL that this account uses for signing in with OAuth. The system ignores this value unless `Enabled` is `true`. The system doesn't use autodiscovery when a declaration contains this URL, so the declaration must also contain a `HostName`.
	SignInURL *string `json:"SignInURL,omitempty" plist:"SignInURL,omitempty"`
	// The URL that this account uses for token requests with OAuth. The system ignores this value unless `Enabled` is `true`.
	TokenRequestURL *string `json:"TokenRequestURL,omitempty" plist:"TokenRequestURL,omitempty"`
}

// Settings for S/MIME.
type AccountExchangeSMIME struct {
	// Settings for S/MIME signing.
	Signing *AccountExchangeSMIMESigning `json:"Signing,omitempty" plist:"Signing,omitempty"`
	// Settings for S/MIME encryption.
	Encryption *AccountExchangeSMIMEEncryption `json:"Encryption,omitempty" plist:"Encryption,omitempty"`
}

// Settings for S/MIME signing.
type AccountExchangeSMIMESigning struct {
	// If `true`, the system enables S/MIME signing.
	Enabled bool `json:"Enabled" plist:"Enabled" required:"true"`
	// Specifies the identifier of an asset declaration containing the identity required for S/MIME signing of messages sent from this account.
	IdentityAssetReference *string `json:"IdentityAssetReference,omitempty" plist:"IdentityAssetReference,omitempty"`
	// If `true`, the user can turn S/MIME signing on or off in Settings.
	UserOverrideable *bool `json:"UserOverrideable,omitempty" plist:"UserOverrideable,omitempty"`
	// If `true`, the user can select an S/MIME signing identity in Settings.
	IdentityUserOverrideable *bool `json:"IdentityUserOverrideable,omitempty" plist:"IdentityUserOverrideable,omitempty"`
}

// Settings for S/MIME encryption.
type AccountExchangeSMIMEEncryption struct {
	// If `true`, the system enables S/MIME encryption by default, which the user can't override if `PerMessageSwitchEnabled` is `false`.
	Enabled bool `json:"Enabled" plist:"Enabled" required:"true"`
	// Specifies the identifier of an asset declaration containing the identity required for S/MIME encryption. The system attaches the public certificate to outgoing mail to allow the user to receive encrypted mail. When the user sends encrypted mail, the system uses the public certificate to encrypt the copy of the mail in their Sent mailbox.
	IdentityAssetReference *string `json:"IdentityAssetReference,omitempty" plist:"IdentityAssetReference,omitempty"`
	// If `true`, the user can turn S/MIME encryption by default on or off in Settings.
	UserOverrideable *bool `json:"UserOverrideable,omitempty" plist:"UserOverrideable,omitempty"`
	// If `true`, the user can select an S/MIME signing identity in Settings.
	IdentityUserOverrideable *bool `json:"IdentityUserOverrideable,omitempty" plist:"IdentityUserOverrideable,omitempty"`
	// If `true`, the system enables the per-message encryption switch in the compose view.
	PerMessageSwitchEnabled *bool `json:"PerMessageSwitchEnabled,omitempty" plist:"PerMessageSwitchEnabled,omitempty"`
}

// The declaration to configure a Google account.
// A Google configuration defines a Google account for a user. The user will be prompted to enter their credentials shortly after the configuration successfully installs.
type AccountGoogle struct {
	// The name that apps show to the user for this Google account. If not present, the system generates a suitable default.
	VisibleName *string `json:"VisibleName,omitempty" plist:"VisibleName,omitempty"`
	// The identifier of an asset declaration that contains the user identity for this Google account. Set the corresponding asset type to `UserIdentity` and ensure that it contains an `EmailAddress` key that specifies the full Google email address for the account.
	UserIdentityAssetReference string `json:"UserIdentityAssetReference" plist:"UserIdentityAssetReference" required:"true"`
}

func (p *AccountGoogle) DeclarationType() string {
	return "com.apple.configuration.account.google"
}

// The declaration to configure a Lightweight Directory Access Protocol (LDAP) account.
// An LDAP configuration defines an LDAP directory account for a user.
type AccountLDAP struct {
	// The name that apps show to the user for this LDAP account. If not present, the system generates a suitable default.
	VisibleName *string `json:"VisibleName,omitempty" plist:"VisibleName,omitempty"`
	// The hostname or IP address of the LDAP server.
	HostName string `json:"HostName" plist:"HostName" required:"true"`
	// The port number or IP address of the LDAP server.
	Port *int64 `json:"Port,omitempty" plist:"Port,omitempty"`
	// The identifier of an asset declaration that contains the credentials for this account. Set the corresponding asset type to `CredentialUserNameAndPassword`.
	AuthenticationCredentialsAssetReference *string `json:"AuthenticationCredentialsAssetReference,omitempty" plist:"AuthenticationCredentialsAssetReference,omitempty"`
	// The array of nodes to start LDAP searches from. There must be at least one node for this account to be useful. macOS only searches one node and ignores other items in the array.
	SearchSettings *[]*SearchSettingsItem `json:"SearchSettings,omitempty" plist:"SearchSettings,omitempty"`
}

func (p *AccountLDAP) DeclarationType() string {
	return "com.apple.configuration.account.ldap"
}

type SearchSettingsItem struct {
	// The description of this search setting in the Contacts and Settings apps. If not present, the apps display no name.
	VisibleDescription *string `json:"VisibleDescription,omitempty" plist:"VisibleDescription,omitempty"`
	// The path to the node where a search starts. For example, `ou=people,o=example corp`.
	SearchBase string `json:"SearchBase" plist:"SearchBase" required:"true"`
	// The type of recursion to use in the search:
	// - `Base`: The search uses only the `SearchBase` node.
	// - `OneLevel`: The search uses the `SearchBase` node and its immediate children.
	// - `Subtree`: The search uses the `SearchBase` node and all its children, regardless of depth.
	Scope *Scope `default:"Subtree" json:"Scope,omitempty" plist:"Scope,omitempty"`
}

// The type of recursion to use in the search:
// - `Base`: The search uses only the `SearchBase` node.
// - `OneLevel`: The search uses the `SearchBase` node and its immediate children.
// - `Subtree`: The search uses the `SearchBase` node and all its children, regardless of depth.
type Scope string

const (
	ScopeBase     Scope = "Base"
	ScopeOneLevel Scope = "OneLevel"
	ScopeSubtree  Scope = "Subtree"
)

// The declaration to configure a Mail account.
// An email configuration defines an email account for a user.
type AccountMail struct {
	// The name that apps show to the user for this mail account. If not present, the system generates a suitable default.
	VisibleName *string `json:"VisibleName,omitempty" plist:"VisibleName,omitempty"`
	// The identifier of an asset declaration that contains the user identity for this account. Set the corresponding asset type to `UserIdentity`.
	UserIdentityAssetReference *string `json:"UserIdentityAssetReference,omitempty" plist:"UserIdentityAssetReference,omitempty"`
	// The settings for the incoming mail server for this account.
	IncomingServer IncomingServer `json:"IncomingServer" plist:"IncomingServer" required:"true"`
	// The settings for the outgoing mail server for this account.
	OutgoingServer OutgoingServer `json:"OutgoingServer" plist:"OutgoingServer" required:"true"`
	// Settings for S/MIME.
	SMIME *AccountMailSMIME `json:"SMIME,omitempty" plist:"SMIME,omitempty"`
}

func (p *AccountMail) DeclarationType() string {
	return "com.apple.configuration.account.mail"
}

// The settings for the incoming mail server for this account.
type IncomingServer struct {
	// The mail protocol this account uses.
	ServerType ServerType `json:"ServerType" plist:"ServerType" required:"true"`
	// The host name for the incoming mail server.
	HostName string `json:"HostName" plist:"HostName" required:"true"`
	// The port number for the incoming mail server.
	Port *int64 `json:"Port,omitempty" plist:"Port,omitempty"`
	// The authentication method for the incoming mail server.
	AuthenticationMethod IncomingServerAuthenticationMethod `json:"AuthenticationMethod" plist:"AuthenticationMethod" required:"true"`
	// The identifier of an asset declaration that contains the credentials for this account to authenticate with an incoming mail server. The corresponding asset must be of type `CredentialUserNameAndPassword`.
	// If the `AuthenticationMethod` is `None`, this field must be blank. Otherwise, the declaration must contain this field.
	AuthenticationCredentialsAssetReference *string `json:"AuthenticationCredentialsAssetReference,omitempty" plist:"AuthenticationCredentialsAssetReference,omitempty"`
	// The path prefix for the IMAP server. The system uses this only when `ServerType` is `IMAP`.
	IMAPPathPrefix *string `json:"IMAPPathPrefix,omitempty" plist:"IMAPPathPrefix,omitempty"`
}

// The mail protocol this account uses.
type ServerType string

const (
	ServerTypeIMAP ServerType = "IMAP"
	ServerTypePOP  ServerType = "POP"
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

// The settings for the outgoing mail server for this account.
type OutgoingServer struct {
	// The host name for the outgoing mail server.
	HostName string `json:"HostName" plist:"HostName" required:"true"`
	// The port number for the outgoing mail server.
	Port *int64 `json:"Port,omitempty" plist:"Port,omitempty"`
	// The authentication method for the outgoing mail server.
	AuthenticationMethod OutgoingServerAuthenticationMethod `json:"AuthenticationMethod" plist:"AuthenticationMethod" required:"true"`
	// The identifier of an asset declaration that contains the credentials for this account to authenticate with an outgoing mail server. The corresponding asset must be of type `CredentialUserNameAndPassword`.
	// If the `AuthenticationMethod` is `None`, this field must be blank. Otherwise, the declaration must contain this field.
	AuthenticationCredentialsAssetReference *string `json:"AuthenticationCredentialsAssetReference,omitempty" plist:"AuthenticationCredentialsAssetReference,omitempty"`
}

// The authentication method for the outgoing mail server.
type OutgoingServerAuthenticationMethod string

const (
	OutgoingServerAuthenticationMethodNone     OutgoingServerAuthenticationMethod = "None"
	OutgoingServerAuthenticationMethodPassword OutgoingServerAuthenticationMethod = "Password"
	OutgoingServerAuthenticationMethodCRAMMD5  OutgoingServerAuthenticationMethod = "CRAMMD5"
	OutgoingServerAuthenticationMethodNTLM     OutgoingServerAuthenticationMethod = "NTLM"
	OutgoingServerAuthenticationMethodHTTPMD5  OutgoingServerAuthenticationMethod = "HTTPMD5"
)

// Settings for S/MIME.
type AccountMailSMIME struct {
	// Settings for S/MIME signing.
	Signing *AccountMailSMIMESigning `json:"Signing,omitempty" plist:"Signing,omitempty"`
	// Settings for S/MIME encryption.
	Encryption *AccountMailSMIMEEncryption `json:"Encryption,omitempty" plist:"Encryption,omitempty"`
}

// Settings for S/MIME signing.
type AccountMailSMIMESigning struct {
	// If `true`, the system enables S/MIME signing.
	Enabled bool `json:"Enabled" plist:"Enabled" required:"true"`
	// Specifies the identifier of an asset declaration containing the identity required for S/MIME signing of messages sent from this account.
	IdentityAssetReference *string `json:"IdentityAssetReference,omitempty" plist:"IdentityAssetReference,omitempty"`
	// If `true`, the user can turn S/MIME signing on or off in Settings.
	UserOverrideable *bool `json:"UserOverrideable,omitempty" plist:"UserOverrideable,omitempty"`
	// If `true`, the user can select an S/MIME signing identity in Settings.
	IdentityUserOverrideable *bool `json:"IdentityUserOverrideable,omitempty" plist:"IdentityUserOverrideable,omitempty"`
}

// Settings for S/MIME encryption.
type AccountMailSMIMEEncryption struct {
	// If `true`, the system enables S/MIME encryption by default, which the user can't override if `PerMessageSwitchEnabled` is `false`.
	Enabled bool `json:"Enabled" plist:"Enabled" required:"true"`
	// Specifies the identifier of an asset declaration containing the identity required for S/MIME encryption. The system attaches the public certificate to outgoing mail to allow the user to receive encrypted mail. When the user sends encrypted mail, the system uses the public certificate to encrypt the copy of the mail in their Sent mailbox.
	IdentityAssetReference *string `json:"IdentityAssetReference,omitempty" plist:"IdentityAssetReference,omitempty"`
	// If `true`, the user can set the default value for S/MIME encryption to on or off in Settings.
	UserOverrideable *bool `json:"UserOverrideable,omitempty" plist:"UserOverrideable,omitempty"`
	// If `true`, the user can select an S/MIME signing identity in Settings.
	IdentityUserOverrideable *bool `json:"IdentityUserOverrideable,omitempty" plist:"IdentityUserOverrideable,omitempty"`
	// If `true`, the system enables the per-message encryption switch in the compose view.
	PerMessageSwitchEnabled *bool `json:"PerMessageSwitchEnabled,omitempty" plist:"PerMessageSwitchEnabled,omitempty"`
}

// The declaration to configure a subscribed calendar.
// A subscribed calendar configuration defines a subscribed calendar for a user.
type AccountSubscribedCalendar struct {
	// The name that apps show to the user for this calendar account. If not present, the system generates a suitable default.
	VisibleName *string `json:"VisibleName,omitempty" plist:"VisibleName,omitempty"`
	// The URL of the subscribed calendar, which needs to start with `https://`.
	CalendarURL string `json:"CalendarURL" plist:"CalendarURL" required:"true"`
	// The identifier of an asset declaration that contains the credentials for this account to authenticate with a calendar server. Set the corresponding asset type to `CredentialUserNameAndPassword`.
	AuthenticationCredentialsAssetReference *string `json:"AuthenticationCredentialsAssetReference,omitempty" plist:"AuthenticationCredentialsAssetReference,omitempty"`
}

func (p *AccountSubscribedCalendar) DeclarationType() string {
	return "com.apple.configuration.account.subscribed-calendar"
}

// The declaration to configure a managed app.
type AppManaged struct {
	// The App Store ID of the managed app that is downloaded from the App Store.
	// Only one of `AppStoreID`, `BundleID`, `ManifestURL`, or `AppComposedIdentifier` needs to be present.
	AppStoreID *string `json:"AppStoreID,omitempty" plist:"AppStoreID,omitempty"`
	// The bundle ID of the managed app that is downloaded from the App Store.
	// Only one of `AppStoreID`, `BundleID`, `ManifestURL`, or `AppComposedIdentifier` needs to be present.
	BundleID *string `json:"BundleID,omitempty" plist:"BundleID,omitempty"`
	// The URL of the manifest for the managed app that the device downloads from a web site. The manifest is returned as a `ManifestURL` property list.
	// Only one of `AppStoreID`, `BundleID`, `ManifestURL`, or `AppComposedIdentifier` needs to be present.
	// Available only in iOS and visionOS.
	ManifestURL *string `json:"ManifestURL,omitempty" plist:"ManifestURL,omitempty"`
	// A string that specifies the composed identifier of an existing app that needs to be managed. The device uses this to take over management of an app installed by some other process, for example installed manually by the user, or via a package configuration. If the app isn't present when the device applies the configuration, the device takes over management of it when it does install.
	// The following rules apply when the device takes over management:
	// - If the `InstallBehavior.Install` key is set to `Required`, the device takes over management of the app.
	// - If the `InstallBehavior.Install` key is set to `Optional`, the device takes over management of the app when the user "installs" it using an MDM management app.
	// The format of the composed identifier is either "Bundle-ID (Team-ID)" or "Bundle-ID {Designated-Requirement}". For example, `com.example.app (ABCD1234)` for the team ID format, or `com.example.app {anchor apple generic}` for the designated requirement format. Management of the app occurs only if its code signature matches the composed identifier.
	// In macOS, only one of `AppStoreID`, `BundleID`, or `AppComposedIdentifier` needs to be present.
	// Available only in macOS.
	AppComposedIdentifier *string `json:"AppComposedIdentifier,omitempty" plist:"AppComposedIdentifier,omitempty"`
	// If `true`, the device install an iOS or iPadOS app that runs on an Apple Silicon Mac. This is only used when the app is an App Store app.
	// Available only in macOS.
	IOSApp *bool `json:"iOSApp,omitempty" plist:"iOSApp,omitempty"`
	// A dictionary that describes how and when to install the app.
	InstallBehavior *AppManagedInstallBehavior `json:"InstallBehavior,omitempty" plist:"InstallBehavior,omitempty"`
	// A dictionary that specifies how the device updates apps.
	UpdateBehavior *UpdateBehavior `json:"UpdateBehavior,omitempty" plist:"UpdateBehavior,omitempty"`
	// If `true`, backups contain the app and its data.
	// Available only in iOS and visionOS.
	IncludeInBackup *bool `json:"IncludeInBackup,omitempty" plist:"IncludeInBackup,omitempty"`
	// A dictionary of values to associate with the app.
	// Available only in iOS and visionOS.
	Attributes *Attributes `json:"Attributes,omitempty" plist:"Attributes,omitempty"`
	// A dictionary of app config data and credentials.
	// Available only in iOS and visionOS.
	AppConfig *AppConfig `json:"AppConfig,omitempty" plist:"AppConfig,omitempty"`
	// A dictionary of extension config data and credentials.
	// Available only in iOS and visionOS.
	ExtensionConfigs *ExtensionConfigs `json:"ExtensionConfigs,omitempty" plist:"ExtensionConfigs,omitempty"`
	// The identifier of an asset declaration containing a reference to the app config data. The device provides the app config data to the app using the MDMv1 behavior. The corresponding asset needs to be of type `com.apple.asset.data`. The referenced data needs to be a property list file, and the asset's "ContentType"
	// value set to match the data type.
	// Available only in iOS and visionOS.
	LegacyAppConfigAssetReference *string `json:"LegacyAppConfigAssetReference,omitempty" plist:"LegacyAppConfigAssetReference,omitempty"`
}

func (p *AppManaged) DeclarationType() string {
	return "com.apple.configuration.app.managed"
}

// A dictionary that describes how and when to install the app.
type AppManagedInstallBehavior struct {
	// A string that specifies if the app needs to remain on the device at all times or if the user can freely install and remove it, which is one of the following values:
	// - `Optional`: The user can install and remove the app after the system activates the configuration.
	// - `Required`: The system installs the app after it activates the configuration. The user can't remove the app.
	// The system automatically installs apps on supervised devices. Otherwise, the device prompts the user to approve installation of the app.
	Install *AppManagedInstallBehaviorInstall `default:"Optional" json:"Install,omitempty" plist:"Install,omitempty"`
	// A dictionary that describes the app's license.
	License *License `json:"License,omitempty" plist:"License,omitempty"`
	// The App Store External Version Identifier (EVID) of the version of the app the device installs. This key is ignored if the app isn't an App Store app.
	// The following rules apply when the device applies or updates the configuration:
	// - If this key isn't present:
	// - If the app isn't present, the device installs the latest version.
	// - If the app is present, if allowed the device takes over management of the current version of the app.
	// - If this key is present:
	// - If the app isn't present, the device installs the app with the specified version.
	// - If an app with the same version is present, if allowed the device takes over management of that app.
	// - If an app with an older version is present, if allowed the device updates the app to the specified version and takes over management of it.
	// - If an app with a newer version is present, the device doesn't take over management of the app. The device reports an app status failure.
	// > Note:
	// > The device never installs an older version of the app over a newer version.
	// You can retrieve this value from the store through the `contentMetadataLookupUrl` of the `VPPServiceConfigSrv`.
	Version *int64 `json:"Version,omitempty" plist:"Version,omitempty"`
	// Indicates how the device uses a cellular network when it downloads the app for automatic install or update operations:
	// - `AlwaysOn`: The device downloads apps of any size using a cellular network.
	// - `AlwaysOff`: The device doesn't download apps using a cellular network. The device pauses the automatic install or update operation until a different network is active.
	// - `StoreSettings`: The device uses the settings for the corresponding store when downloading apps.
	// The device always uses the store settings to download apps when the install or update operation is user initiated.
	// Available only in iOS.
	AllowDownloadsOverCellular *AllowDownloadsOverCellular `default:"StoreSettings" json:"AllowDownloadsOverCellular,omitempty" plist:"AllowDownloadsOverCellular,omitempty"`
}

// A string that specifies if the app needs to remain on the device at all times or if the user can freely install and remove it, which is one of the following values:
// - `Optional`: The user can install and remove the app after the system activates the configuration.
// - `Required`: The system installs the app after it activates the configuration. The user can't remove the app.
// The system automatically installs apps on supervised devices. Otherwise, the device prompts the user to approve installation of the app.
type AppManagedInstallBehaviorInstall string

const (
	AppManagedInstallBehaviorInstallOptional AppManagedInstallBehaviorInstall = "Optional"
	AppManagedInstallBehaviorInstallRequired AppManagedInstallBehaviorInstall = "Required"
)

// A dictionary that describes the app's license.
type License struct {
	// The type of license that the app uses for installation through the App Store, which is one of the following values:
	// - `Device`: The app has a device license.
	// - `User`: The app has a user license.
	// This key needs to be present for App Store apps, when either `AppStoreID` or `BundleID` are present in the configuration.
	Assignment *Assignment `json:"Assignment,omitempty" plist:"Assignment,omitempty"`
	// The type of VPP license that the app uses for installation through the App Store, which is one of the following values:
	// - `Device`: The app has a VPP device license.
	// - `User`: The app has a VPP user license.
	// This key needs to be present to install an app through the App Store.
	VPPType *VPPType `json:"VPPType,omitempty" plist:"VPPType,omitempty"`
}

// The type of license that the app uses for installation through the App Store, which is one of the following values:
// - `Device`: The app has a device license.
// - `User`: The app has a user license.
// This key needs to be present for App Store apps, when either `AppStoreID` or `BundleID` are present in the configuration.
type Assignment string

const (
	AssignmentDevice Assignment = "Device"
	AssignmentUser   Assignment = "User"
)

// The type of VPP license that the app uses for installation through the App Store, which is one of the following values:
// - `Device`: The app has a VPP device license.
// - `User`: The app has a VPP user license.
// This key needs to be present to install an app through the App Store.
type VPPType string

const (
	VPPTypeDevice VPPType = "Device"
	VPPTypeUser   VPPType = "User"
)

// Indicates how the device uses a cellular network when it downloads the app for automatic install or update operations:
// - `AlwaysOn`: The device downloads apps of any size using a cellular network.
// - `AlwaysOff`: The device doesn't download apps using a cellular network. The device pauses the automatic install or update operation until a different network is active.
// - `StoreSettings`: The device uses the settings for the corresponding store when downloading apps.
// The device always uses the store settings to download apps when the install or update operation is user initiated.
// Available only in iOS.
type AllowDownloadsOverCellular string

const (
	AllowDownloadsOverCellularAlwaysOn      AllowDownloadsOverCellular = "AlwaysOn"
	AllowDownloadsOverCellularAlwaysOff     AllowDownloadsOverCellular = "AlwaysOff"
	AllowDownloadsOverCellularStoreSettings AllowDownloadsOverCellular = "StoreSettings"
)

// A dictionary that specifies how the device updates apps.
type UpdateBehavior struct {
	// Specifies whether the device automatically updates the app:
	// - `AlwaysOn`: The device automatically updates the app to the latest version. For App Store apps, the device periodically checks the store for updates. For Enterprise apps, the device periodically downloads the manifest file and compares it to the previous manifest file. If the device detects a change to the bundle version in the manifest, it downloads and updates the app.
	// - `AlwaysOff`: The device never automatically updates the app.
	// - `StoreSettings`: The device uses the settings for the corresponding store to determine when to automatically update the app. For Enterprise apps, this setting behaves the same as `AlwaysOff`.
	// When the `InstallBehavior.Version` key is specified, the device ignores this key and Automatic App Updates are disabled.
	// In macOS, the device ignores this setting if the `AppComposedIdentifier` key is set in the configuration.
	AutomaticAppUpdates AutomaticAppUpdates `json:"AutomaticAppUpdates" plist:"AutomaticAppUpdates" required:"true"`
}

// Specifies whether the device automatically updates the app:
// - `AlwaysOn`: The device automatically updates the app to the latest version. For App Store apps, the device periodically checks the store for updates. For Enterprise apps, the device periodically downloads the manifest file and compares it to the previous manifest file. If the device detects a change to the bundle version in the manifest, it downloads and updates the app.
// - `AlwaysOff`: The device never automatically updates the app.
// - `StoreSettings`: The device uses the settings for the corresponding store to determine when to automatically update the app. For Enterprise apps, this setting behaves the same as `AlwaysOff`.
// When the `InstallBehavior.Version` key is specified, the device ignores this key and Automatic App Updates are disabled.
// In macOS, the device ignores this setting if the `AppComposedIdentifier` key is set in the configuration.
type AutomaticAppUpdates string

const (
	AutomaticAppUpdatesAlwaysOn      AutomaticAppUpdates = "AlwaysOn"
	AutomaticAppUpdatesAlwaysOff     AutomaticAppUpdates = "AlwaysOff"
	AutomaticAppUpdatesStoreSettings AutomaticAppUpdates = "StoreSettings"
)

// A dictionary of values to associate with the app.
// Available only in iOS and visionOS.
type Attributes struct {
	// An array of domain names to associate with the app.
	AssociatedDomains *[]string `json:"AssociatedDomains,omitempty" plist:"AssociatedDomains,omitempty"`
	// If `true`, the system enables direct downloads for the `AssociatedDomains`.
	AssociatedDomainsEnableDirectDownloads *bool `json:"AssociatedDomainsEnableDirectDownloads,omitempty" plist:"AssociatedDomainsEnableDirectDownloads,omitempty"`
	// The cellular slice identifier, which can be the data network name (DNN) or app category. For DNN, encode the value as "DNN:name", where "name" is the carrier-provided DNN name. For app category, encode the value as "AppCategory:category", where "category" is a carrier-provided string such as "Enterprise1".
	CellularSliceUUID *string `json:"CellularSliceUUID,omitempty" plist:"CellularSliceUUID,omitempty"`
	// The UUID of the content filter to associate with the app.
	ContentFilterUUID *string `json:"ContentFilterUUID,omitempty" plist:"ContentFilterUUID,omitempty"`
	// The UUID of the DNS proxy to associate with the app.
	DNSProxyUUID *string `json:"DNSProxyUUID,omitempty" plist:"DNSProxyUUID,omitempty"`
	// If `false`, the system prevents the user from hiding the app. It doesn't affect the user's ability to leave it in the App Library, while removing it from the Home Screen.
	Hideable *bool `json:"Hideable,omitempty" plist:"Hideable,omitempty"`
	// If `false`, the system prevents the user from locking the app. This also prevents the user from hiding the app.
	Lockable *bool `json:"Lockable,omitempty" plist:"Lockable,omitempty"`
	// The UUID of the relay to associate with the app.
	RelayUUID *string `json:"RelayUUID,omitempty" plist:"RelayUUID,omitempty"`
	// If `true`, the device automatically locks after every transaction that requires a customer's card PIN. If `false`, the user can choose the behavior.
	TapToPayScreenLock *bool `json:"TapToPayScreenLock,omitempty" plist:"TapToPayScreenLock,omitempty"`
	// The UUID of the VPN to associate with the app.
	VPNUUID *string `json:"VPNUUID,omitempty" plist:"VPNUUID,omitempty"`
}

// A dictionary of app config data and credentials.
// Available only in iOS and visionOS.
type AppConfig struct {
	// Specifies the identifier of an asset declaration containing a reference to the app or extension config data. The corresponding asset needs to be of type `com.apple.asset.data`. The referenced data needs to be a property list file, and the asset's "ContentType" value set to match the data type.
	DataAssetReference *string `json:"DataAssetReference,omitempty" plist:"DataAssetReference,omitempty"`
	// Provides passwords to the managed app or extension. Each element in the array contains a password asset reference and an associated identifier which the app or extension uses to look up the password.
	Passwords *[]PasswordAppConfigItem `json:"Passwords,omitempty" plist:"Passwords,omitempty"`
	// Provides identities to the managed app or extension. Each element in the array contains an identity asset reference and an associated identifier which the app or extension uses to look up the identity.
	Identities *[]IdentityAppConfigItem `json:"Identities,omitempty" plist:"Identities,omitempty"`
	// Provides certificates to the managed app or extension. Each element in the array contains a certificate asset reference and an associated identifier which the app or extension uses to look up the certificate.
	Certificates *[]CertificateAppConfigItem `json:"Certificates,omitempty" plist:"Certificates,omitempty"`
}

// A dictionary of values associated with a credential config.
type PasswordAppConfigItem struct {
	// The app or extension uses this identifier to fetch the corresponding password using the `ManagedApp` framework. App developers define the values for these identifiers.
	Identifier string `json:"Identifier" plist:"Identifier" required:"true"`
	// Specifies the identifier of an asset declaration containing a username and password. The `ManagedApp` framework makes the password available to the app or extension. The `ManagedApp` framework ignores the username.
	AssetReference string `json:"AssetReference" plist:"AssetReference" required:"true"`
}

// A dictionary of values associated with a credential config.
type IdentityAppConfigItem struct {
	// The app or extension uses this identifier to fetch the corresponding identity using the `ManagedApp` framework. App developers define the values for these identifiers.
	Identifier string `json:"Identifier" plist:"Identifier" required:"true"`
	// Specifies the identifier of an asset declaration containing an identity.
	AssetReference string `json:"AssetReference" plist:"AssetReference" required:"true"`
}

// A dictionary of values associated with a credential config.
type CertificateAppConfigItem struct {
	// The app or extension uses this identifier to fetch the corresponding certificate using the `ManagedApp` framework. App developers define the values for these identifiers.
	Identifier string `json:"Identifier" plist:"Identifier" required:"true"`
	// Specifies the identifier of an asset declaration containing a certificate.
	AssetReference string `json:"AssetReference" plist:"AssetReference" required:"true"`
}

// A dictionary of extension config data and credentials.
// Available only in iOS and visionOS.
type ExtensionConfigs struct {
	// A dictionary mapping extension composed identifiers to the extension config data and credentials. The expected format is "Identifier (TeamIdentifier)".
	ANY *ExtensionConfigsANY `json:"ANY,omitempty" plist:"ANY,omitempty"`
}

// A dictionary mapping extension composed identifiers to the extension config data and credentials. The expected format is "Identifier (TeamIdentifier)".
type ExtensionConfigsANY struct {
	// Specifies the identifier of an asset declaration containing a reference to the app or extension config data. The corresponding asset needs to be of type `com.apple.asset.data`. The referenced data needs to be a property list file, and the asset's "ContentType" value set to match the data type.
	DataAssetReference *string `json:"DataAssetReference,omitempty" plist:"DataAssetReference,omitempty"`
	// Provides passwords to the managed app or extension. Each element in the array contains a password asset reference and an associated identifier which the app or extension uses to look up the password.
	Passwords *[]PasswordAppConfigItem `json:"Passwords,omitempty" plist:"Passwords,omitempty"`
	// Provides identities to the managed app or extension. Each element in the array contains an identity asset reference and an associated identifier which the app or extension uses to look up the identity.
	Identities *[]IdentityAppConfigItem `json:"Identities,omitempty" plist:"Identities,omitempty"`
	// Provides certificates to the managed app or extension. Each element in the array contains a certificate asset reference and an associated identifier which the app or extension uses to look up the certificate.
	Certificates *[]CertificateAppConfigItem `json:"Certificates,omitempty" plist:"Certificates,omitempty"`
}

// The declaration to configure audio accessory settings.
type AudioAccessorySettings struct {
	// A dictionary that describes audio accessory temporary pairing behavior. The device enables temporary pairing when this key is present and the `Disabled` key isn't `false`. The device doesn't synchronize pairing information with iCloud when temporary pairing is active.
	TemporaryPairing *TemporaryPairing `json:"TemporaryPairing,omitempty" plist:"TemporaryPairing,omitempty"`
}

func (p *AudioAccessorySettings) DeclarationType() string {
	return "com.apple.configuration.audio-accessory.settings"
}

// A dictionary that describes audio accessory temporary pairing behavior. The device enables temporary pairing when this key is present and the `Disabled` key isn't `false`. The device doesn't synchronize pairing information with iCloud when temporary pairing is active.
type TemporaryPairing struct {
	// If `true`, temporary pairing of audio accessories is disabled.
	Disabled *bool `json:"Disabled,omitempty" plist:"Disabled,omitempty"`
	// A dictionary providing configuration for temporary pairing. Required if `Disabled` isn't present or is `false`.
	Configuration *Configuration `json:"Configuration,omitempty" plist:"Configuration,omitempty"`
}

// A dictionary providing configuration for temporary pairing. Required if `Disabled` isn't present or is `false`.
type Configuration struct {
	// A dictionary that describes when the device automatically unpairs temporarily paired audio accessories.
	UnpairingTime UnpairingTime `json:"UnpairingTime" plist:"UnpairingTime" required:"true"`
}

// A dictionary that describes when the device automatically unpairs temporarily paired audio accessories.
type UnpairingTime struct {
	// A string that specifies the device's unpairing policy.
	// - `None`: The device doesn't automatically unpair. Use this only with a return to service device that you erase and reenroll when assigning it from one user to another.
	// - `Hour`: The device automatically unpairs temporarily paired audio accessories at the local time that the `Hour` key specifies.
	Policy Policy `json:"Policy" plist:"Policy" required:"true"`
	// The local time hour (24-hour clock) when the device automatically unpairs temporarily paired audio accessories. Required when setting the `Policy` key to `Hour`.
	Hour *int64 `json:"Hour,omitempty" plist:"Hour,omitempty"`
}

// A string that specifies the device's unpairing policy.
// - `None`: The device doesn't automatically unpair. Use this only with a return to service device that you erase and reenroll when assigning it from one user to another.
// - `Hour`: The device automatically unpairs temporarily paired audio accessories at the local time that the `Hour` key specifies.
type Policy string

const (
	PolicyNone Policy = "None"
	PolicyHour Policy = "Hour"
)

// The declaration to configure disk management settings on the device.
type DiskManagementSettings struct {
	// The restrictions for the disk.
	Restrictions *Restrictions `json:"Restrictions,omitempty" plist:"Restrictions,omitempty"`
}

func (p *DiskManagementSettings) DeclarationType() string {
	return "com.apple.configuration.diskmanagement.settings"
}

// The restrictions for the disk.
type Restrictions struct {
	// Specifies the mount policy for external storage:
	// - `Allowed`: The system can mount external storage that is read-write or read-only.
	// - `ReadOnly`: The system can only mount read-only external storage. Note that external storage that is read-write will not be mounted read-only.
	// - `Disallowed`: The system can't mount any external storage.
	ExternalStorage *ExternalStorage `json:"ExternalStorage,omitempty" plist:"ExternalStorage,omitempty"`
	// Specifies the mount policy for network storage:
	// - `Allowed`: The system can mount network storage that is read-write or read-only.
	// - `ReadOnly`: The system can only mount read-only network storage. Note that network storage that is read-write will not be mounted read-only.
	// - `Disallowed`: The system can't mount any network storage.
	NetworkStorage *NetworkStorage `json:"NetworkStorage,omitempty" plist:"NetworkStorage,omitempty"`
}

// Specifies the mount policy for external storage:
// - `Allowed`: The system can mount external storage that is read-write or read-only.
// - `ReadOnly`: The system can only mount read-only external storage. Note that external storage that is read-write will not be mounted read-only.
// - `Disallowed`: The system can't mount any external storage.
type ExternalStorage string

const (
	ExternalStorageAllowed    ExternalStorage = "Allowed"
	ExternalStorageReadOnly   ExternalStorage = "ReadOnly"
	ExternalStorageDisallowed ExternalStorage = "Disallowed"
)

// Specifies the mount policy for network storage:
// - `Allowed`: The system can mount network storage that is read-write or read-only.
// - `ReadOnly`: The system can only mount read-only network storage. Note that network storage that is read-write will not be mounted read-only.
// - `Disallowed`: The system can't mount any network storage.
type NetworkStorage string

const (
	NetworkStorageAllowed    NetworkStorage = "Allowed"
	NetworkStorageReadOnly   NetworkStorage = "ReadOnly"
	NetworkStorageDisallowed NetworkStorage = "Disallowed"
)

// The declaration to configure an interactive legacy profile.
type LegacyInteractiveProfile struct {
	// The URL of the profile to download and install, which needs to start with `https://`, and must be hosted by the MDM server.
	// If a user enrollment triggers this configuration, the system silently ignores any MDMv1 payloads in macOS that are forbidden with user enrollment. In iOS, the system rejects the entire profile.
	ProfileURL string `json:"ProfileURL" plist:"ProfileURL" required:"true"`
	// The visible name of the configuration. This name needs to indicate the nature of the profile.
	VisibleName string `json:"VisibleName" plist:"VisibleName" required:"true"`
}

func (p *LegacyInteractiveProfile) DeclarationType() string {
	return "com.apple.configuration.legacy.interactive"
}

// The declaration to configure a legacy profile.
type LegacyProfile struct {
	// The URL of the profile to download and install, which needs to start with `https://`, and must be hosted by the MDM server.
	// If a user enrollment triggers this configuration, the system silently ignores any MDMv1 payloads in macOS where the User Enrollment Mode setting is `forbidden`. In iOS, the system rejects the entire profile.
	ProfileURL string `json:"ProfileURL" plist:"ProfileURL" required:"true"`
}

func (p *LegacyProfile) DeclarationType() string {
	return "com.apple.configuration.legacy"
}

// The declaration to configure status subscriptions.
type ManagementStatusSubscriptions struct {
	// An array of status items that the device notifies subscribers about.
	StatusItems []*StatusItem `json:"StatusItems" plist:"StatusItems" required:"true"`
}

func (p *ManagementStatusSubscriptions) DeclarationType() string {
	return "com.apple.configuration.management.status-subscriptions"
}

// The declaration for configuring a specific status subscription.
type StatusItem struct {
	// The name of the status item to send to subscribers.
	Name string `json:"Name" plist:"Name" required:"true"`
}

// The declaration to test declarative device management.
type ManagementTest struct {
	// The string to echo back in a status response reason.
	Echo string `json:"Echo" plist:"Echo" required:"true"`
	// The string to read from a data asset to echo back in status response reason description.
	EchoDataAssetReference *string `json:"EchoDataAssetReference,omitempty" plist:"EchoDataAssetReference,omitempty"`
	// The status the system reports back when the device implements the configuration. Use this to override the normal `success` result.
	ReturnStatus *ReturnStatus `default:"Installed" json:"ReturnStatus,omitempty" plist:"ReturnStatus,omitempty"`
}

func (p *ManagementTest) DeclarationType() string {
	return "com.apple.configuration.management.test"
}

// The status the system reports back when the device implements the configuration. Use this to override the normal `success` result.
type ReturnStatus string

const (
	ReturnStatusInstalled ReturnStatus = "Installed"
	ReturnStatusFailed    ReturnStatus = "Failed"
	ReturnStatusUnlocked  ReturnStatus = "Unlocked"
)

// The declaration to configure the math and calculator apps.
// Configures the built-in math and calculator app settings.
type MathSettings struct {
	// If present, configures the built-in Calculator app.
	Calculator *Calculator `json:"Calculator,omitempty" plist:"Calculator,omitempty"`
	// If present, configures math behavior in the system.
	SystemBehavior *SystemBehavior `json:"SystemBehavior,omitempty" plist:"SystemBehavior,omitempty"`
}

func (p *MathSettings) DeclarationType() string {
	return "com.apple.configuration.math.settings"
}

// If present, configures the built-in Calculator app.
type Calculator struct {
	// If present, configures the basic mode of the calculator. Basic mode is always enabled.
	BasicMode *BasicMode `json:"BasicMode,omitempty" plist:"BasicMode,omitempty"`
	// If present, configures the scientific mode of the calculator. If not present, scientific mode is enabled.
	ScientificMode *ScientificMode `json:"ScientificMode,omitempty" plist:"ScientificMode,omitempty"`
	// If present, configures the programmer mode of the calculator. If not present, programmer mode is enabled.
	ProgrammerMode *ProgrammerMode `json:"ProgrammerMode,omitempty" plist:"ProgrammerMode,omitempty"`
	// If present, configures the Math Notes mode of the calculator. If not present, Math Notes mode is enabled.
	MathNotesMode *MathNotesMode `json:"MathNotesMode,omitempty" plist:"MathNotesMode,omitempty"`
	// If present, controls global input options of the calculator. If not present, all input modes are enabled.
	InputModes *InputModes `json:"InputModes,omitempty" plist:"InputModes,omitempty"`
}

// If present, configures the basic mode of the calculator. Basic mode is always enabled.
type BasicMode struct {
	// Add the square root button to the basic calculator by replacing the +/- button. Normally, the square root button is available in scientific mode, so this key can be used to make it available when the scientific mode is restricted.
	AddSquareRoot bool `json:"AddSquareRoot" plist:"AddSquareRoot" required:"true"`
}

// If present, configures the scientific mode of the calculator. If not present, scientific mode is enabled.
type ScientificMode struct {
	// Controls whether the mode is enabled.
	Enabled bool `json:"Enabled" plist:"Enabled" required:"true"`
}

// If present, configures the programmer mode of the calculator. If not present, programmer mode is enabled.
type ProgrammerMode struct {
	// Controls whether the mode is enabled.
	Enabled bool `json:"Enabled" plist:"Enabled" required:"true"`
}

// If present, configures the Math Notes mode of the calculator. If not present, Math Notes mode is enabled.
type MathNotesMode struct {
	// Controls whether the mode is enabled.
	Enabled bool `json:"Enabled" plist:"Enabled" required:"true"`
}

// If present, controls global input options of the calculator. If not present, all input modes are enabled.
type InputModes struct {
	// Configures whether unit conversions are enabled.
	UnitConversion bool `json:"UnitConversion" plist:"UnitConversion" required:"true"`
	// Configures whether RPN input is enabled.
	RPN bool `json:"RPN" plist:"RPN" required:"true"`
}

// If present, configures math behavior in the system.
type SystemBehavior struct {
	// Controls whether keyboard suggestions include math solutions.
	KeyboardSuggestions bool `json:"KeyboardSuggestions" plist:"KeyboardSuggestions" required:"true"`
	// Controls whether Math Notes is allowed in other apps such as Notes.
	MathNotes bool `json:"MathNotes" plist:"MathNotes" required:"true"`
}

// The declaration to install a package.
type Package struct {
	// The URL of the manifest document for the package that the device downloads. The manifest is returned as a `ManifestURL` property list. The `url` property of the manifest must point to the package (.pkg) file to install.
	ManifestURL string `json:"ManifestURL" plist:"ManifestURL" required:"true"`
	// A dictionary that describes how and when to install the package.
	InstallBehavior *PackageInstallBehavior `json:"InstallBehavior,omitempty" plist:"InstallBehavior,omitempty"`
}

func (p *Package) DeclarationType() string {
	return "com.apple.configuration.package"
}

// A dictionary that describes how and when to install the package.
type PackageInstallBehavior struct {
	// A string that specifies when the system installs the package:
	// - `Optional`: The user can install the package after the system activates the configuration.
	// - `Required`: The system installs the package after it activates the configuration.
	Install *PackageInstallBehaviorInstall `default:"Optional" json:"Install,omitempty" plist:"Install,omitempty"`
}

// A string that specifies when the system installs the package:
// - `Optional`: The user can install the package after the system activates the configuration.
// - `Required`: The system installs the package after it activates the configuration.
type PackageInstallBehaviorInstall string

const (
	PackageInstallBehaviorInstallOptional PackageInstallBehaviorInstall = "Optional"
	PackageInstallBehaviorInstallRequired PackageInstallBehaviorInstall = "Required"
)

// The declaration to configure passcode policy settings.
type PasscodeSettings struct {
	// If `true`, the system requires the user to set a passcode without any requirements about the length or quality of the passcode. The presence of any other keys implicitly requires a passcode, and overrides this key's value.
	RequirePasscode *bool `json:"RequirePasscode,omitempty" plist:"RequirePasscode,omitempty"`
	// If `true`, the passcode needs to consist of at least one alphabetic character and at least one number.
	RequireAlphanumericPasscode *bool `json:"RequireAlphanumericPasscode,omitempty" plist:"RequireAlphanumericPasscode,omitempty"`
	// If `true`, the system requires a complex passcode. A complex passcode is one that doesn't contain repeated characters or increasing or decreasing characters (such as 123 or CBA).
	RequireComplexPasscode *bool `json:"RequireComplexPasscode,omitempty" plist:"RequireComplexPasscode,omitempty"`
	// The minimum number of characters a passcode can contain.
	MinimumLength *int64 `default:"0" json:"MinimumLength,omitempty" plist:"MinimumLength,omitempty"`
	// Specifies the minimum number of complex characters in the password. A complex character is a character other than a number or a letter, such as `&`, `%`, `$`, and `#`.
	MinimumComplexCharacters *int64 `default:"0" json:"MinimumComplexCharacters,omitempty" plist:"MinimumComplexCharacters,omitempty"`
	// The number of failed passcode attempts that the system allows the user before iOS erases the device or macOS locks the device. If you don't change this setting, after six failed attempts, the device imposes a time delay before the user can enter a passcode again. The time delay increases with each failed attempt.
	// After the final failed attempt, the system securely erases all data and settings from the iOS device. A macOS device locks after the final attempt. The passcode time delay begins after the sixth attempt, so if this value is six or lower, the system has no time delay and triggers the erase or lock as soon as the user exceeds the limit.
	MaximumFailedAttempts *int64 `default:"11" json:"MaximumFailedAttempts,omitempty" plist:"MaximumFailedAttempts,omitempty"`
	// The number of minutes before the login is reset after the maximum number of failed attempts. Also set the `MaximumFailedAttempts` key for this to take effect.
	FailedAttemptsResetInMinutes *int64 `json:"FailedAttemptsResetInMinutes,omitempty" plist:"FailedAttemptsResetInMinutes,omitempty"`
	// The maximum period that a user can select, during which the user can unlock the device without a passcode. A value of `0` means no grace period, and the device requires a passcode immediately. In the absence of this key, the user can select any period. In macOS, the system translates this to screensaver settings.
	MaximumGracePeriodInMinutes *int64 `json:"MaximumGracePeriodInMinutes,omitempty" plist:"MaximumGracePeriodInMinutes,omitempty"`
	// The maximum period that a user can select, during which the device can be idle before the system automatically locks it. When the device reaches this limit, the device locks and the user must enter the passcode to unlock it. In the absence of this key, the user can select any period. In macOS, the system translates this to screensaver settings.
	MaximumInactivityInMinutes *int64 `json:"MaximumInactivityInMinutes,omitempty" plist:"MaximumInactivityInMinutes,omitempty"`
	// Specifies the maximum number of days that the passcode can remain unchanged. After this number of days, the system forces the user to change the passcode before it unlocks the device.
	MaximumPasscodeAgeInDays *int64 `json:"MaximumPasscodeAgeInDays,omitempty" plist:"MaximumPasscodeAgeInDays,omitempty"`
	// The number of historical passcode entries the system checks when validating a new passcode. The device refuses a new passcode if it matches a previously used passcode within the specified passcode history range. In the absence of this key, the system performs no historical check.
	PasscodeReuseLimit *int64 `json:"PasscodeReuseLimit,omitempty" plist:"PasscodeReuseLimit,omitempty"`
	// If `true`, the system forces a password reset the next time the user tries to authenticate. If you set this key in a configuration in the system scope (device channel), the setting takes effect for all users, and admin authentication may fail until the admin user password is also reset.
	ChangeAtNextAuth *bool `json:"ChangeAtNextAuth,omitempty" plist:"ChangeAtNextAuth,omitempty"`
	// Specifies a regular expression, and its description, to enforce password compliance. Use the simpler passcode settings whenever possible, and rely on regular expression matching only when necessary. Mistakes in regular expressions can lead to frustrating user experiences, such as unsatisfiable passcode policies, or policy descriptions that don't match the enforced policy.
	CustomRegex *CustomRegex `json:"CustomRegex,omitempty" plist:"CustomRegex,omitempty"`
}

func (p *PasscodeSettings) DeclarationType() string {
	return "com.apple.configuration.passcode.settings"
}

// Specifies a regular expression, and its description, to enforce password compliance. Use the simpler passcode settings whenever possible, and rely on regular expression matching only when necessary. Mistakes in regular expressions can lead to frustrating user experiences, such as unsatisfiable passcode policies, or policy descriptions that don't match the enforced policy.
type CustomRegex struct {
	// A regular expression string to match against the password to determine whether it complies with a policy. The regular expression uses the ICU syntax. The string can't exceed 2048 characters in length.
	Regex string `json:"Regex" plist:"Regex" required:"true"`
	// A dictionary with supported OS language IDs for the keys (such as `en-US`), and values that represent a localized description of the policy that the regular expression enforces. Use the special `default` key for languages that the dictionary doesn't contain.
	Description *Description `json:"Description,omitempty" plist:"Description,omitempty"`
}

// A dictionary with supported OS language IDs for the keys (such as `en-US`), and values that represent a localized description of the policy that the regular expression enforces. Use the special `default` key for languages that the dictionary doesn't contain.
type Description struct {
	// A localized description.
	ANY *string `json:"ANY,omitempty" plist:"ANY,omitempty"`
}

// The declaration to configure managed bookmarks in Safari.
type SafariBookmarks struct {
	// A dictionary that specifies a set of managed bookmarks.
	ManagedBookmarks *[]BookmarkGroup `json:"ManagedBookmarks,omitempty" plist:"ManagedBookmarks,omitempty"`
}

func (p *SafariBookmarks) DeclarationType() string {
	return "com.apple.configuration.safari.bookmarks"
}

// A group of managed bookmarks.
type BookmarkGroup struct {
	// A string that specifies the unique identifier for this group of managed bookmarks. Safari displays a separate managed bookmarks item for each set of unique managed bookmarks based on the value of this key. If multiple active configurations use the same value for this key, Safari displays a single group formed by merging the list of `Bookmarks` from each group.
	GroupIdentifier string `json:"GroupIdentifier" plist:"GroupIdentifier" required:"true"`
	// The name of the bookmarks folder. Safari uses this as the title for the top-level managed bookmarks item.
	Title string `json:"Title" plist:"Title" required:"true"`
	// A list of bookmarks. Either a `URL` or `Folder` key must be present in each item.
	Bookmarks []BookmarksItem `json:"Bookmarks" plist:"Bookmarks" required:"true"`
}

// A bookmark that specifies a title, and either a URL for the bookmark, or a nested folder of bookmarks.
type BookmarksItem struct {
	// The title of the bookmark shown in Safari.
	Title string `json:"Title" plist:"Title" required:"true"`
	// The URL for the bookmark item.
	// Only one of `URL` or `Folder` must be present.
	URL *string `json:"URL,omitempty" plist:"URL,omitempty"`
	// An array of bookmarks for each bookmark in the folder. Folders can include bookmark items and bookmark folders.
	// Only one of `URL` or `Folder` must be present.
	Folder *[]FolderItem `json:"Folder,omitempty" plist:"Folder,omitempty"`
}

// A bookmark that specifies a title, and either a URL for the bookmark, or a nested folder of bookmarks.
type FolderItem struct{}

// The declaration to configure Safari Extensions.
type SafariExtensionSettings struct {
	// The dictionary of managed extensions settings.
	ManagedExtensions *ManagedExtensions `json:"ManagedExtensions,omitempty" plist:"ManagedExtensions,omitempty"`
}

func (p *SafariExtensionSettings) DeclarationType() string {
	return "com.apple.configuration.safari.extensions.settings"
}

// The dictionary of managed extensions settings.
type ManagedExtensions struct {
	// The composed identifier of the managed extension, or "*" for all extensions. In order for the extension to be managed, its host app must be present on the device.
	// To generate this string use `codesign -dv <path_to_appex>`. The browser extension is located in the PlugIns folder inside the app bundle. The expected format is "Identifier (TeamIdentifier)", for example "com.example.app (ABCD1234)".
	// For extensions that aren't also available on macOS the app developer needs to provide this information.
	ANY *ManagedExtensionsANY `json:"ANY,omitempty" plist:"ANY,omitempty"`
}

// The composed identifier of the managed extension, or "*" for all extensions. In order for the extension to be managed, its host app must be present on the device.
// To generate this string use `codesign -dv <path_to_appex>`. The browser extension is located in the PlugIns folder inside the app bundle. The expected format is "Identifier (TeamIdentifier)", for example "com.example.app (ABCD1234)".
// For extensions that aren't also available on macOS the app developer needs to provide this information.
type ManagedExtensionsANY struct {
	// Controls whether an extension is allowed.
	// * `Allowed` - The user is allowed to turn the extension on or off.
	// * `AlwaysOn` - The extension will always be on.
	// * `AlwaysOff` - The extension will always be off.
	State *State `json:"State,omitempty" plist:"State,omitempty"`
	// Controls whether an extension is allowed in Private Browsing.
	// * `Allowed` - The user is allowed to turn the extension on or off in Private Browsing.
	// * `AlwaysOn` - The extension will always be on in Private Browsing if the extension is on outside of Private Browsing.
	// * `AlwaysOff` - The extension will never be on in Private Browsing.
	PrivateBrowsing *PrivateBrowsing `json:"PrivateBrowsing,omitempty" plist:"PrivateBrowsing,omitempty"`
	// Controls the domains and sub-domains the extension is granted access to. Any non-prefixed domains take precedence over prefixed domains, and `DeniedDomains` takes precedence over `AllowedDomains`. The user can configure any domains not matched by the values in `AllowedDomains` or `DeniedDomains`.
	AllowedDomains *[]string `json:"AllowedDomains,omitempty" plist:"AllowedDomains,omitempty"`
	// Controls the domains and sub-domains the extension isn't allowed to access. Any non-prefixed domains take precedence over prefixed domains, and `DeniedDomains` takes precedence over `AllowedDomains`. The user can configure any domains not matched by the values in `AllowedDomains` or `DeniedDomains`.
	DeniedDomains *[]string `json:"DeniedDomains,omitempty" plist:"DeniedDomains,omitempty"`
}

// Controls whether an extension is allowed.
// * `Allowed` - The user is allowed to turn the extension on or off.
// * `AlwaysOn` - The extension will always be on.
// * `AlwaysOff` - The extension will always be off.
type State string

const (
	StateAllowed   State = "Allowed"
	StateAlwaysOn  State = "AlwaysOn"
	StateAlwaysOff State = "AlwaysOff"
)

// Controls whether an extension is allowed in Private Browsing.
// * `Allowed` - The user is allowed to turn the extension on or off in Private Browsing.
// * `AlwaysOn` - The extension will always be on in Private Browsing if the extension is on outside of Private Browsing.
// * `AlwaysOff` - The extension will never be on in Private Browsing.
type PrivateBrowsing string

const (
	PrivateBrowsingAllowed   PrivateBrowsing = "Allowed"
	PrivateBrowsingAlwaysOn  PrivateBrowsing = "AlwaysOn"
	PrivateBrowsingAlwaysOff PrivateBrowsing = "AlwaysOff"
)

// The declaration to configure Safari settings.
type SafariSettings struct {
	// The policy Safari uses for managing cookies:
	// - `Never`: Safari always blocks cookies.
	// - `CurrentWebsite`: Safari allows cookies only from the current website.
	// - `VisitedWebsites`: Safari allows cookies only from visited websites.
	// - `Always`: Safari always allows cookies.
	AcceptCookies *AcceptCookies `default:"Always" json:"AcceptCookies,omitempty" plist:"AcceptCookies,omitempty"`
	// If `false`, the system forces fraud warnings on in Safari.
	AllowDisablingFraudWarning *bool `json:"AllowDisablingFraudWarning,omitempty" plist:"AllowDisablingFraudWarning,omitempty"`
	// If `false`, the system disables clearing history in Safari.
	AllowHistoryClearing *bool `json:"AllowHistoryClearing,omitempty" plist:"AllowHistoryClearing,omitempty"`
	// If `false`, the system disables JavaScript in Safari.
	AllowJavaScript *bool `json:"AllowJavaScript,omitempty" plist:"AllowJavaScript,omitempty"`
	// If `false`, the system disables private browsing in Safari.
	AllowPrivateBrowsing *bool `json:"AllowPrivateBrowsing,omitempty" plist:"AllowPrivateBrowsing,omitempty"`
	// If `false`, the system disables popups in Safari.
	AllowPopups *bool `json:"AllowPopups,omitempty" plist:"AllowPopups,omitempty"`
	// If `false`, the system disables summarization of content in Safari.
	AllowSummary *bool `json:"AllowSummary,omitempty" plist:"AllowSummary,omitempty"`
	// Sets the start page for new tabs in Safari.
	NewTabStartPage *NewTabStartPage `json:"NewTabStartPage,omitempty" plist:"NewTabStartPage,omitempty"`
}

func (p *SafariSettings) DeclarationType() string {
	return "com.apple.configuration.safari.settings"
}

// The policy Safari uses for managing cookies:
// - `Never`: Safari always blocks cookies.
// - `CurrentWebsite`: Safari allows cookies only from the current website.
// - `VisitedWebsites`: Safari allows cookies only from visited websites.
// - `Always`: Safari always allows cookies.
type AcceptCookies string

const (
	AcceptCookiesNever           AcceptCookies = "Never"
	AcceptCookiesCurrentWebsite  AcceptCookies = "CurrentWebsite"
	AcceptCookiesVisitedWebsites AcceptCookies = "VisitedWebsites"
	AcceptCookiesAlways          AcceptCookies = "Always"
)

// Sets the start page for new tabs in Safari.
type NewTabStartPage struct {
	// Sets the start page type in Safari:
	// - `Start` - Safari uses the default start page. Safari disables the Homepage.
	// - `Home` - Safari uses the page specified by `HomepageURL`, and Safari also sets that as the Homepage.
	// - `Extension` - Safari uses the page specified by the Safari extension whose identifier is `ExtensionIdentifier`. Safari disables the Homepage.
	PageType PageType `json:"PageType" plist:"PageType" required:"true"`
	// The URL of the homepage which needs to start with `https://` or `http://`. Required when setting `PageType` to `Home`.
	HomepageURL *string `json:"HomepageURL,omitempty" plist:"HomepageURL,omitempty"`
	// The composed identifier of the extension that provides the start page. The required format is "Identifier (TeamIdentifier)", for example "com.example.app (ABCD1234)". Required when setting `PageType` to `Extension`.
	ExtensionIdentifier *string `json:"ExtensionIdentifier,omitempty" plist:"ExtensionIdentifier,omitempty"`
}

// Sets the start page type in Safari:
// - `Start` - Safari uses the default start page. Safari disables the Homepage.
// - `Home` - Safari uses the page specified by `HomepageURL`, and Safari also sets that as the Homepage.
// - `Extension` - Safari uses the page specified by the Safari extension whose identifier is `ExtensionIdentifier`. Safari disables the Homepage.
type PageType string

const (
	PageTypeStart     PageType = "Start"
	PageTypeHome      PageType = "Home"
	PageTypeExtension PageType = "Extension"
)

// The declaration to configure a group of screen-sharing connections.
type ScreenSharingConnectionGroup struct {
	// A unique identifier for this connection group.
	ConnectionGroupUUID string `json:"ConnectionGroupUUID" plist:"ConnectionGroupUUID" required:"true"`
	// The name of the connection group.
	GroupName string `json:"GroupName" plist:"GroupName" required:"true"`
	// An array of `ConnectionUUID`s that represent connections declared in `ScreenSharingConnection` configurations that are members of this group.
	Members []string `json:"Members" plist:"Members" required:"true"`
}

func (p *ScreenSharingConnectionGroup) DeclarationType() string {
	return "com.apple.configuration.screensharing.connection.group"
}

// The declaration to configure a connection to a screen-sharing host.
type ScreenSharingConnection struct {
	// A unique identifier for this connection when it's in a connection group.
	ConnectionUUID string `json:"ConnectionUUID" plist:"ConnectionUUID" required:"true"`
	// The name of the connection.
	DisplayName string `json:"DisplayName" plist:"DisplayName" required:"true"`
	// The host name or IP address of the Mac that hosts the screen-sharing connection.
	HostName string `json:"HostName" plist:"HostName" required:"true"`
	// The TCP port number on the host to initiate the connection.
	Port *int64 `json:"Port,omitempty" plist:"Port,omitempty"`
	// The display configuration for this connection.
	DisplayConfiguration DisplayConfiguration `json:"DisplayConfiguration" plist:"DisplayConfiguration" required:"true"`
	// The identifier of an asset declaration that contains the required credentials for this connection to authenticate with the screen-sharing server. Set the corresponding asset type to `com.apple.asset.credential.userpassword`.
	AuthenticationCredentialsAssetReference *string `json:"AuthenticationCredentialsAssetReference,omitempty" plist:"AuthenticationCredentialsAssetReference,omitempty"`
}

func (p *ScreenSharingConnection) DeclarationType() string {
	return "com.apple.configuration.screensharing.connection"
}

// The display configuration for this connection.
type DisplayConfiguration struct {
	// The type of display for the connection, which has these allowed values:
	// - `Virtual1`: Create one virtual display.
	// - `Virtual2`: Create two virtual displays.
	DisplayType DisplayType `json:"DisplayType" plist:"DisplayType" required:"true"`
}

// The type of display for the connection, which has these allowed values:
// - `Virtual1`: Create one virtual display.
// - `Virtual2`: Create two virtual displays.
type DisplayType string

const (
	DisplayTypeVirtual1 DisplayType = "Virtual1"
	DisplayTypeVirtual2 DisplayType = "Virtual2"
)

// The declaration to configure screen-sharing host settings and restrictions.
type ScreenSharingHostSettings struct {
	// The maximum number of virtual displays to make available to clients.
	MaximumVirtualDisplays *int64 `json:"MaximumVirtualDisplays,omitempty" plist:"MaximumVirtualDisplays,omitempty"`
	// The initial UDP port number to connect to the host. Screen sharing requires multiple connections, so the system increments this value by 1 for each additional connection. This doesn't change the port number that the system uses to initially establish a connection with a host, which is always TCP port 5900.
	PortBase *int64 `json:"PortBase,omitempty" plist:"PortBase,omitempty"`
	// If `true`, the system prevents users from copying files from the screen-sharing host.
	PreventCopyFilesFromHost *bool `json:"PreventCopyFilesFromHost,omitempty" plist:"PreventCopyFilesFromHost,omitempty"`
	// If `true`, the system prevents users from copying files to the screen-sharing host.
	PreventCopyFilesToHost *bool `json:"PreventCopyFilesToHost,omitempty" plist:"PreventCopyFilesToHost,omitempty"`
	// If `true`, the system prevents clients from establishing high-performance connections to the host.
	PreventHighPerformanceConnections *bool `json:"PreventHighPerformanceConnections,omitempty" plist:"PreventHighPerformanceConnections,omitempty"`
}

func (p *ScreenSharingHostSettings) DeclarationType() string {
	return "com.apple.configuration.screensharing.host.settings"
}

// The declaration to add a certificate to the device.
type SecurityCertificate struct {
	// The identifier of an asset declaration that contains the certificate to install.
	CredentialAssetReference string `json:"CredentialAssetReference" plist:"CredentialAssetReference" required:"true"`
}

func (p *SecurityCertificate) DeclarationType() string {
	return "com.apple.configuration.security.certificate"
}

// The declaration to install an identity on the device.
type SecurityIdentity struct {
	// The identifier of an asset declaration that contains the identity to install.
	CredentialAssetReference string `json:"CredentialAssetReference" plist:"CredentialAssetReference" required:"true"`
	// If `true`, apps can access the private key.
	AllowAllAppsAccess *bool `json:"AllowAllAppsAccess,omitempty" plist:"AllowAllAppsAccess,omitempty"`
	// If `true`, the private key is extractable in the keychain.
	KeyIsExtractable *bool `json:"KeyIsExtractable,omitempty" plist:"KeyIsExtractable,omitempty"`
}

func (p *SecurityIdentity) DeclarationType() string {
	return "com.apple.configuration.security.identity"
}

// The declaration to configure the device to allow WebAuthn enterprise attestation for certain passkeys.
type SecurityPasskeyAttestation struct {
	// The identifier of an asset declaration that contains the identity to install and use for passkey attestation.
	AttestationIdentityAssetReference string `json:"AttestationIdentityAssetReference" plist:"AttestationIdentityAssetReference" required:"true"`
	// If `true`, the private key for the attestation identity is extractable in the keychain.
	AttestationIdentityKeyIsExtractable *bool `json:"AttestationIdentityKeyIsExtractable,omitempty" plist:"AttestationIdentityKeyIsExtractable,omitempty"`
	// An array of the relying parties to allow enterprise attestation.
	RelyingParties []string `json:"RelyingParties" plist:"RelyingParties" required:"true"`
}

func (p *SecurityPasskeyAttestation) DeclarationType() string {
	return "com.apple.configuration.security.passkey.attestation"
}

// The declaration to configure background tasks.
type ServicesBackgroundTasks struct {
	// The unique identifier of the set of background tasks managed with this configuration. This should be a reverse DNS style identifier. The system uses this identifier to differentiate between tasks in different configurations.
	TaskType string `json:"TaskType" plist:"TaskType" required:"true"`
	// A description of the set of background tasks managed by this configuration.
	TaskDescription *string `json:"TaskDescription,omitempty" plist:"TaskDescription,omitempty"`
	// Specifies the identifier of an asset declaration containing a reference to the files to be used for the background task configuration. The corresponding asset must be of type `com.apple.asset.data`.
	// The referenced data must be a zip archive of an entire directory, that will be expanded and stored in a well known location for the background task. The asset's "ContentType" and "Hash-SHA-256" keys in the "Reference" key are required.
	// This file should contain background task executables, scripts, and configuration files, but not the `launchd` configuration files.
	ExecutableAssetReference *string `json:"ExecutableAssetReference,omitempty" plist:"ExecutableAssetReference,omitempty"`
	// An array of `launchd` configuration files used to run the background tasks.
	LaunchdConfigurations *[]LaunchdItem `json:"LaunchdConfigurations,omitempty" plist:"LaunchdConfigurations,omitempty"`
}

func (p *ServicesBackgroundTasks) DeclarationType() string {
	return "com.apple.configuration.services.background-tasks"
}

// A dictionary of launchd configurations.
type LaunchdItem struct {
	// Specifies the identifier of an asset declaration containing a reference to the launchd configuration file for the background task. The referenced data must be a property list file conforming to the launchd.plist format. The asset's "ContentType" and "Hash-SHA-256" keys in the "Reference" key are required.
	FileAssetReference string `json:"FileAssetReference" plist:"FileAssetReference" required:"true"`
	// Indicates whether the launchd configuration file is applied to the system daemon, or system agent domain.
	Context Context `json:"Context" plist:"Context" required:"true"`
}

// Indicates whether the launchd configuration file is applied to the system daemon, or system agent domain.
type Context string

const (
	ContextDaemon Context = "daemon"
	ContextAgent  Context = "agent"
)

// The managed configuration files for services.
type ServicesConfigurationFiles struct {
	// The identifier of the system service with managed configuration files. Use a reverse DNS style for this identifier. However, the system reserves `com.apple.` prefix for built-in services. The available built-in services are:
	// - `com.apple.sshd` configures sshd
	// - `com.apple.sudo` configures sudo
	// - `com.apple.pam` configures PAM
	// - `com.apple.cups` configures CUPS
	// - `com.apple.apache.httpd` configures Apache httpd
	// - `com.apple.bash` configures bash
	// - `com.apple.zsh` configures zsh
	ServiceType string `json:"ServiceType" plist:"ServiceType" required:"true"`
	// The identifier of an asset declaration that contains a reference to the files to use for system service configuration. Ensure that the corresponding asset:
	// - Is of type `com.apple.asset.data`
	// - Is a zip archive of an entire directory
	// - Has a `Reference` key that includes the `ContentType` and `Hash-SHA-256` keys, which the system requires
	// The system expands the zip archive and stores the data in a well-known location for the service.
	DataAssetReference string `json:"DataAssetReference" plist:"DataAssetReference" required:"true"`
}

func (p *ServicesConfigurationFiles) DeclarationType() string {
	return "com.apple.configuration.services.configuration-files"
}

// A software update enforcement policy for a specific OS release.
type SoftwareUpdateEnforcementSpecific struct {
	// The target OS version to update the device to by the appropriate time. This is the OS version number, for example, `16.1`.
	TargetOSVersion string `json:"TargetOSVersion" plist:"TargetOSVersion" required:"true"`
	// The target build version to update the device to by the appropriate time, for example, `20A242`. The system uses the build version for testing during seeding periods. The build version can include a supplemental version identifier, for example, `20A242a`. If the build version isn't consistent with the target OS version specified in the `TargetOSVersion` key, the target OS version takes precedence.
	TargetBuildVersion *string `json:"TargetBuildVersion,omitempty" plist:"TargetBuildVersion,omitempty"`
	// The local date time value that specifies when to force install the software update. Use the format `yyyy-mm-ddThh:mm:ss`, which is derived from RFC3339 but doesn't include a time zone offset. If the user doesn't trigger the software update before this time, the device force installs it.
	TargetLocalDateTime string `json:"TargetLocalDateTime" plist:"TargetLocalDateTime" required:"true"`
	// The URL of a web page that shows details that the organization provides about the enforced update.
	DetailsURL *string `json:"DetailsURL,omitempty" plist:"DetailsURL,omitempty"`
}

func (p *SoftwareUpdateEnforcementSpecific) DeclarationType() string {
	return "com.apple.configuration.softwareupdate.enforcement.specific"
}

// The declaration to configure software updates.
type SoftwareUpdateSettings struct {
	// If set to `true`, the device shows all software update enforcement notifications.
	// If set to `false`, the device only shows notifications triggered one hour before the enforcement deadline, and the restart countdown notification.
	Notifications *bool `json:"Notifications,omitempty" plist:"Notifications,omitempty"`
	// This object configures the deferral of software updates. Rapid Security Responses aren't considered in `Major`, `Minor`, or `System` deferral mechanism.
	Deferrals *Deferrals `json:"Deferrals,omitempty" plist:"Deferrals,omitempty"`
	// This string specifies how the device shows software updates to the user. When more than one update is available update, the device behaves as follows:
	// - `All` - Shows all software update versions.
	// - `Oldest` - Shows only the oldest (lower numbered) software update version.
	// - `Newest` - Shows only the newest (highest numbered) software update version.
	RecommendedCadence *RecommendedCadence `json:"RecommendedCadence,omitempty" plist:"RecommendedCadence,omitempty"`
	// This object configures various automatic Software Update functionality.
	AutomaticActions *AutomaticActions `json:"AutomaticActions,omitempty" plist:"AutomaticActions,omitempty"`
	// These configurations set user access to interacting with Rapid Security Responses (RSRs).
	RapidSecurityResponse *RapidSecurityResponse `json:"RapidSecurityResponse,omitempty" plist:"RapidSecurityResponse,omitempty"`
	// If set to `true`, a standard user can perform Major and Minor Software Updates.
	// If set to `false`, only administrators can perform Major and Minor Software Updates.
	AllowStandardUserOSUpdates *bool `json:"AllowStandardUserOSUpdates,omitempty" plist:"AllowStandardUserOSUpdates,omitempty"`
	// This object configures the beta program settings for a device.
	Beta *Beta `json:"Beta,omitempty" plist:"Beta,omitempty"`
}

func (p *SoftwareUpdateSettings) DeclarationType() string {
	return "com.apple.configuration.softwareupdate.settings"
}

// This object configures the deferral of software updates. Rapid Security Responses aren't considered in `Major`, `Minor`, or `System` deferral mechanism.
type Deferrals struct {
	// Specifies the number of days to defer a major or minor OS software update on the device. When set, software updates only appear after the specified delay, following the release of the software update. Available in iOS 18 and later.
	CombinedPeriodInDays *int64 `json:"CombinedPeriodInDays,omitempty" plist:"CombinedPeriodInDays,omitempty"`
	// Specifies the number of days to defer a major OS software update on the device. When set, software updates only appear after the specified delay, following the release of the software update. Available in macOS 15 and later.
	MajorPeriodInDays *int64 `json:"MajorPeriodInDays,omitempty" plist:"MajorPeriodInDays,omitempty"`
	// Specifies the number of days to defer a minor OS software update on the device. It also defers major updates for iOS. When set, software updates only appear after the specified delay, following the release of the software update. Available in macOS 15 and later.
	MinorPeriodInDays *int64 `json:"MinorPeriodInDays,omitempty" plist:"MinorPeriodInDays,omitempty"`
	// Specifies the number of days to defer system or non-OS updates. When set, updates only appear after the specified delay, following the release of the update. Available in macOS 15 and later.
	SystemPeriodInDays *int64 `json:"SystemPeriodInDays,omitempty" plist:"SystemPeriodInDays,omitempty"`
}

// This string specifies how the device shows software updates to the user. When more than one update is available update, the device behaves as follows:
// - `All` - Shows all software update versions.
// - `Oldest` - Shows only the oldest (lower numbered) software update version.
// - `Newest` - Shows only the newest (highest numbered) software update version.
type RecommendedCadence string

const (
	RecommendedCadenceAll    RecommendedCadence = "All"
	RecommendedCadenceOldest RecommendedCadence = "Oldest"
	RecommendedCadenceNewest RecommendedCadence = "Newest"
)

// This object configures various automatic Software Update functionality.
type AutomaticActions struct {
	// Specifies whether the user can control automatic downloads of available updates:
	// - `Allowed` - the user can enable or disable automatic downloads.
	// - `AlwaysOn` - automatic downloads are always enabled.
	// - `AlwaysOff` - automatic downloads are always disabled.
	Download *Download `default:"Allowed" json:"Download,omitempty" plist:"Download,omitempty"`
	// Specifies whether the user can control automatic installation of available updates:
	// - `Allowed` - the user can enable or disable automatic installation.
	// - `AlwaysOn` - automatic installations are always enabled.
	// - `AlwaysOff` - automatic installations are always disabled.
	InstallOSUpdates *InstallOSUpdates `default:"Allowed" json:"InstallOSUpdates,omitempty" plist:"InstallOSUpdates,omitempty"`
	// Specifies whether the user can control automatic installation of available security updates:
	// - `Allowed` - the user can enable or disable automatic installation.
	// - `AlwaysOn` - automatic installations are always enabled.
	// - `AlwaysOff` - automatic installations are always disabled.
	InstallSecurityUpdate *InstallSecurityUpdate `default:"Allowed" json:"InstallSecurityUpdate,omitempty" plist:"InstallSecurityUpdate,omitempty"`
}

// Specifies whether the user can control automatic downloads of available updates:
// - `Allowed` - the user can enable or disable automatic downloads.
// - `AlwaysOn` - automatic downloads are always enabled.
// - `AlwaysOff` - automatic downloads are always disabled.
type Download string

const (
	DownloadAllowed   Download = "Allowed"
	DownloadAlwaysOn  Download = "AlwaysOn"
	DownloadAlwaysOff Download = "AlwaysOff"
)

// Specifies whether the user can control automatic installation of available updates:
// - `Allowed` - the user can enable or disable automatic installation.
// - `AlwaysOn` - automatic installations are always enabled.
// - `AlwaysOff` - automatic installations are always disabled.
type InstallOSUpdates string

const (
	InstallOSUpdatesAllowed   InstallOSUpdates = "Allowed"
	InstallOSUpdatesAlwaysOn  InstallOSUpdates = "AlwaysOn"
	InstallOSUpdatesAlwaysOff InstallOSUpdates = "AlwaysOff"
)

// Specifies whether the user can control automatic installation of available security updates:
// - `Allowed` - the user can enable or disable automatic installation.
// - `AlwaysOn` - automatic installations are always enabled.
// - `AlwaysOff` - automatic installations are always disabled.
type InstallSecurityUpdate string

const (
	InstallSecurityUpdateAllowed   InstallSecurityUpdate = "Allowed"
	InstallSecurityUpdateAlwaysOn  InstallSecurityUpdate = "AlwaysOn"
	InstallSecurityUpdateAlwaysOff InstallSecurityUpdate = "AlwaysOff"
)

// These configurations set user access to interacting with Rapid Security Responses (RSRs).
type RapidSecurityResponse struct {
	// If set to `false`, Rapid Security Responses aren't offered for user installation. The system can still install Rapid Security Responses with `com.apple.configuration.softwareupdate.enforcement.specific` configurations.
	// If set to `true`, the system offers Rapid Security Responses to the user.
	Enable *bool `json:"Enable,omitempty" plist:"Enable,omitempty"`
	// If set to `false`, the system doesn't offer Rapid Security Response rollbacks to the user.
	// If set to `true`, the system offers Rapid Security Response rollbacks to the user.
	EnableRollback *bool `json:"EnableRollback,omitempty" plist:"EnableRollback,omitempty"`
}

// This object configures the beta program settings for a device.
type Beta struct {
	// Specifies whether the user can control beta program enrollment in the software update settings UI:
	// - `Allowed` - the user can enroll in any applicable beta programs associated with their logged in Apple Account. If the `OfferPrograms` key is present, then the programs listed in that key are also presented to the user.
	// - `AlwaysOn` - the beta programs specified by the organization are used, and the user isn't able to enroll in a beta program using their logged in Apple Account. The device is automatically enrolled into the beta program specified by the `RequireProgram` key if it's present. Otherwise, the system presents the programs listed in the `OfferPrograms` key to the user to choose which to enroll with.
	// - `AlwaysOff` - The device isn't allowed to enroll in any beta programs. The system removes the device from any beta programs, if already enrolled.
	ProgramEnrollment *ProgramEnrollment `default:"Allowed" json:"ProgramEnrollment,omitempty" plist:"ProgramEnrollment,omitempty"`
	// An array of beta programs allowed on the device. This key must only be present if the `ProgramEnrollment` key is set to `Allowed` or `AlwaysOn`. This key must not be present if the `RequireProgram` key is present. This key can be present on unsupervised devices where the `ProgramEnrollment` key isn't supported but is implicitly set to `Allowed`.
	OfferPrograms *[]Program `json:"OfferPrograms,omitempty" plist:"OfferPrograms,omitempty"`
	// The device automatically enrolls in this beta program. This key must only be present if the `ProgramEnrollment` key is set to `AlwaysOn`. The `OfferPrograms` key must not be present if this key is present.
	RequireProgram *RequireProgram `json:"RequireProgram,omitempty" plist:"RequireProgram,omitempty"`
}

// Specifies whether the user can control beta program enrollment in the software update settings UI:
// - `Allowed` - the user can enroll in any applicable beta programs associated with their logged in Apple Account. If the `OfferPrograms` key is present, then the programs listed in that key are also presented to the user.
// - `AlwaysOn` - the beta programs specified by the organization are used, and the user isn't able to enroll in a beta program using their logged in Apple Account. The device is automatically enrolled into the beta program specified by the `RequireProgram` key if it's present. Otherwise, the system presents the programs listed in the `OfferPrograms` key to the user to choose which to enroll with.
// - `AlwaysOff` - The device isn't allowed to enroll in any beta programs. The system removes the device from any beta programs, if already enrolled.
type ProgramEnrollment string

const (
	ProgramEnrollmentAllowed   ProgramEnrollment = "Allowed"
	ProgramEnrollmentAlwaysOn  ProgramEnrollment = "AlwaysOn"
	ProgramEnrollmentAlwaysOff ProgramEnrollment = "AlwaysOff"
)

// The name and token associated with a specific beta program to be allowed.
type Program struct {
	// A human readable description of the beta program.
	Description string `json:"Description" plist:"Description" required:"true"`
	// The Apple Business Manager or Apple School Manager seeding service token for the organization the MDM server is part of. The system uses this token to enroll the device in the corresponding beta program.
	Token string `json:"Token" plist:"Token" required:"true"`
}

// The device automatically enrolls in this beta program. This key must only be present if the `ProgramEnrollment` key is set to `AlwaysOn`. The `OfferPrograms` key must not be present if this key is present.
type RequireProgram struct {
	// A human readable description of the beta program.
	Description string `json:"Description" plist:"Description" required:"true"`
	// The Apple Business Manager or Apple School Manager seeding service token for the organization the MDM server is part of. The system uses this token to enroll the device in the corresponding beta program.
	Token string `json:"Token" plist:"Token" required:"true"`
}

// The declaration to configure an MDMv1 profile for Apple Watch enrollment.
type WatchEnrollment struct {
	// The URL of the profile that the Apple Watch downloads and installs if the user opts in to management during the pairing process, which needs to start with `https://`. Successful enrollment requires that the pairing iPhone is supervised and the profile contains an MDM payload. Apple Watch attempts to install each payload that the profile contains.
	EnrollmentProfileURL string `json:"EnrollmentProfileURL" plist:"EnrollmentProfileURL" required:"true"`
	// An array of identifiers of asset declarations that contain anchor certificates to use to evaluate the trust of the enrollment profile server. Set the type of the corresponding assets to `com.apple.asset.credential.certificate`.
	// These certificates are pinned, meaning that the server specified by the `EnrollmentProfileURL` must use a certificate that chains to one of the certs in this array.
	// If it chains to one of the built-in trusted root certificates but not one of the `AnchorCertificateAssetReferences` certs, the connection will fail.
	AnchorCertificateAssetReferences *[]string `json:"AnchorCertificateAssetReferences,omitempty" plist:"AnchorCertificateAssetReferences,omitempty"`
}

func (p *WatchEnrollment) DeclarationType() string {
	return "com.apple.configuration.watch.enrollment"
}
