// DO NOT EDIT
// generated from https://github.com/apple/device-management.git:9e7fdbe23990b2551c4b29b39bc8b0de956805eb/declarative/declarations/configurations

package configurations

const DeviceManagementGenerateHash = "9e7fdbe23990b2551c4b29b39bc8b0de956805eb"

var DeclarationMap = map[string]any{
	"com.apple.configuration.accessibility.settings":              AccessibilitySettings{},
	"com.apple.configuration.account.caldav":                      AccountCalDAV{},
	"com.apple.configuration.account.carddav":                     AccountCardDAV{},
	"com.apple.configuration.account.exchange":                    AccountExchange{},
	"com.apple.configuration.account.google":                      AccountGoogle{},
	"com.apple.configuration.account.ldap":                        AccountLDAP{},
	"com.apple.configuration.account.mail":                        AccountMail{},
	"com.apple.configuration.account.settings":                    AccountSettings{},
	"com.apple.configuration.account.subscribed-calendar":         AccountSubscribedCalendar{},
	"com.apple.configuration.app.managed":                         AppManaged{},
	"com.apple.configuration.app.settings":                        AppSettings{},
	"com.apple.configuration.audio-accessory.settings":            AudioAccessorySettings{},
	"com.apple.configuration.content-cache.settings":              ContentCaching{},
	"com.apple.configuration.diskmanagement.settings":             DiskManagementSettings{},
	"com.apple.configuration.extensible-sso":                      ExtensibleSSO{},
	"com.apple.configuration.external-intelligence.settings":      ExternalIntelligenceSettings{},
	"com.apple.configuration.intelligence.settings":               IntelligenceSettings{},
	"com.apple.configuration.keyboard.settings":                   KeyboardSettings{},
	"com.apple.configuration.legacy":                              LegacyProfile{},
	"com.apple.configuration.legacy.interactive":                  LegacyInteractiveProfile{},
	"com.apple.configuration.management.status-subscriptions":     ManagementStatusSubscriptions{},
	"com.apple.configuration.management.test":                     ManagementTest{},
	"com.apple.configuration.math.settings":                       MathSettings{},
	"com.apple.configuration.migration-assistant.settings":        MigrationAssistantSettings{},
	"com.apple.configuration.network.dns-proxy":                   NetworkDNSProxy{},
	"com.apple.configuration.network.dns-settings":                NetworkDNSSettings{},
	"com.apple.configuration.network.relay":                       NetworkRelay{},
	"com.apple.configuration.network.vpn.always-on":               NetworkVPNAlwaysOn{},
	"com.apple.configuration.network.vpn.ikev2":                   NetworkVPNIKEv2{},
	"com.apple.configuration.network.vpn.ipsec":                   NetworkVPNIPSec{},
	"com.apple.configuration.network.vpn.vpn-plugin":              NetworkVPNVPNPlugin{},
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
	"com.apple.configuration.siri.settings":                       SiriSettings{},
	"com.apple.configuration.softwareupdate.enforcement.specific": SoftwareUpdateEnforcementSpecific{},
	"com.apple.configuration.softwareupdate.settings":             SoftwareUpdateSettings{},
	"com.apple.configuration.watch.enrollment":                    WatchEnrollment{},
	"com.apple.configuration.webcontent-filter.plugin":            WebContentFilterPlugin{},
}

// The declaration to configure accessibility settings.
type AccessibilitySettings struct {
	// If present, configures vision accessibility settings.
	Vision *Vision `json:"Vision,omitempty"`
}

func (p *AccessibilitySettings) DeclarationType() string {
	return "com.apple.configuration.accessibility.settings"
}

// If present, configures vision accessibility settings.
type Vision struct {
	// If `false`, disables the Accessibility Live Recognition Ask about Images and Surroundings features.
	AllowLiveRecognition *bool `json:"AllowLiveRecognition,omitempty"`
}

// The declaration to configure a Calendar account.
type AccountCalDAV struct {
	// The name that apps show to the user for this calendar account. If not present, the system generates a suitable default.
	VisibleName *string `json:"VisibleName,omitempty"`
	// The hostname or IP address of the CalDAV server.
	HostName string `json:"HostName" required:"true"`
	// The port number for the CalDAV server.
	Port *int64 `json:"Port,omitempty"`
	// The path for the CalDAV server.
	Path *string `json:"Path,omitempty"`
	// The identifier of an asset declaration that contains the credentials for this account. Set the corresponding asset type to `CredentialUserNameAndPassword`.
	AuthenticationCredentialsAssetReference *string `json:"AuthenticationCredentialsAssetReference,omitempty"`
}

func (p *AccountCalDAV) DeclarationType() string {
	return "com.apple.configuration.account.caldav"
}

// The declaration to configure a Contacts account.
type AccountCardDAV struct {
	// The name that apps show to the user for this address book account. If not present, the system generates a suitable default.
	VisibleName *string `json:"VisibleName,omitempty"`
	// The hostname or IP address of the CardDAV server.
	HostName string `json:"HostName" required:"true"`
	// The port number for the CardDAV server.
	Port *int64 `json:"Port,omitempty"`
	// The path for the CardDAV server.
	Path *string `json:"Path,omitempty"`
	// The identifier of an asset declaration that contains the credentials for this account. Set the corresponding asset type to `CredentialUserNameAndPassword`.
	AuthenticationCredentialsAssetReference *string `json:"AuthenticationCredentialsAssetReference,omitempty"`
}

func (p *AccountCardDAV) DeclarationType() string {
	return "com.apple.configuration.account.carddav"
}

// The declaration to configure an Exchange account.
type AccountExchange struct {
	// The name that apps show to the user for this Exchange account. If not present, the system generates a suitable default.
	VisibleName *string `json:"VisibleName,omitempty"`
	// The set of protocol types to enable on the Exchange server, in order of preference. This is an array of unique strings with possible values:
	// - `EAS:` Exchange ActiveSync
	// - `EWS:` Exchange Web Services
	// If the device supports one or more of the listed protocol types, it sets up an account for the first supported type.
	// If the device doesn't support any of the listed protocol types, it doesn't set up an account and the system reports an error.
	EnabledProtocolTypes []EnabledProtocolTypes `json:"EnabledProtocolTypes" required:"true"`
	// The identifier of an asset declaration that contains the user identity for this account. The corresponding asset must be of type `UserIdentity`.
	UserIdentityAssetReference *string `json:"UserIdentityAssetReference,omitempty"`
	// The IP address or fully qualified domain name (FQDN) of the Exchange host.
	HostName *string `json:"HostName,omitempty"`
	// The port number of the EWS server. The system uses this only when this declaration has a `HostName` value. Applicable for "EWS" only.
	Port *int64 `json:"Port,omitempty"`
	// The path of the EWS server. The system uses this only when this declaration has a `HostName` value. Applicable for "EWS" only.
	Path *string `json:"Path,omitempty"`
	// The external hostname of the EWS server (or IP address). Applicable for "EWS" only.
	ExternalHostName *string `json:"ExternalHostName,omitempty"`
	// The external port number of the EWS server. The system uses this only when this declaration has an `ExternalHostName` value. Applicable for "EWS" only.
	ExternalPort *int64 `json:"ExternalPort,omitempty"`
	// The external path of the EWS server. The system uses this only when this declaration has an `ExternalHostName` value. Applicable for "EWS" only.
	ExternalPath *string `json:"External Path,omitempty"`
	// The configuration settings for OAuth for this account.
	OAuth *OAuth `json:"OAuth,omitempty"`
	// The identifier of an asset declaration that contains the credentials for this account to authenticate with an Exchange server. Set the corresponding asset type to `CredentialUserNameAndPassword`.
	AuthenticationCredentialsAssetReference *string `json:"AuthenticationCredentialsAssetReference,omitempty"`
	// The identifier of a credential asset declaration that contains the identity that this account requires to authenticate with the Exchange server.
	AuthenticationIdentityAssetReference *string `json:"AuthenticationIdentityAssetReference,omitempty"`
	// Settings for S/MIME. Applicable for "EAS" only.
	SMIME *AccountExchangeSMIME `json:"SMIME,omitempty"`
	// If `true`, the system activates the mail service for this account.
	MailServiceActive *bool `json:"MailServiceActive,omitempty"`
	// If `true`, the system prevents the user from changing the status of the mail service for this account. Applicable for "EAS" only.
	LockMailService *bool `json:"LockMailService,omitempty"`
	// If `true`, activates the address book service for this account.
	ContactsServiceActive *bool `json:"ContactsServiceActive,omitempty"`
	// If `true`, the system prevents the user from changing the status of the address book service for this account. Applicable for "EAS" only.
	LockContactsService *bool `json:"LockContactsService,omitempty"`
	// If `true`, activates the calendar service for this account.
	CalendarServiceActive *bool `json:"CalendarServiceActive,omitempty"`
	// If `true`, the system prevents the user from changing the status of the calendar service for this account. Applicable for "EAS" only.
	LockCalendarService *bool `json:"LockCalendarService,omitempty"`
	// If `true`, the system activates the reminders service for this account.
	RemindersServiceActive *bool `json:"RemindersServiceActive,omitempty"`
	// If `true`, the system prevents the user from changing the status of the reminders service for this account. Applicable for "EAS" only.
	LockRemindersService *bool `json:"LockRemindersService,omitempty"`
	// If `true`, the system activates the notes service for this account.
	NotesServiceActive *bool `json:"NotesServiceActive,omitempty"`
	// If `true`, the system prevents the user from changing the status of the notes service for this account. Applicable for "EAS" only.
	LockNotesService *bool `json:"LockNotesService,omitempty"`
}

func (p *AccountExchange) DeclarationType() string {
	return "com.apple.configuration.account.exchange"
}

// The set of protocol types to enable on the Exchange server, in order of preference. This is an array of unique strings with possible values:
// - `EAS:` Exchange ActiveSync
// - `EWS:` Exchange Web Services
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
	Enabled bool `json:"Enabled" required:"true"`
	// The URL that this account uses for signing in with OAuth. The system ignores this value unless `Enabled` is `true`. The system doesn't use autodiscovery when a declaration contains this URL, so the declaration must also contain a `HostName`.
	SignInURL *string `json:"SignInURL,omitempty"`
	// The URL that this account uses for token requests with OAuth. The system ignores this value unless `Enabled` is `true`. Applicable for "EAS" only.
	TokenRequestURL *string `json:"TokenRequestURL,omitempty"`
}

// Settings for S/MIME. Applicable for "EAS" only.
type AccountExchangeSMIME struct {
	// Settings for S/MIME signing. Applicable for "EAS" only.
	Signing *AccountExchangeSMIMESigning `json:"Signing,omitempty"`
	// Settings for S/MIME encryption. Applicable for "EAS" only.
	Encryption *AccountExchangeSMIMEEncryption `json:"Encryption,omitempty"`
}

// Settings for S/MIME signing. Applicable for "EAS" only.
type AccountExchangeSMIMESigning struct {
	// If `true`, the system enables S/MIME signing. Applicable for "EAS" only.
	Enabled bool `json:"Enabled" required:"true"`
	// The identifier of an asset declaration containing the identity required for S/MIME signing of messages sent from this account. Applicable for "EAS" only.
	IdentityAssetReference *string `json:"IdentityAssetReference,omitempty"`
	// If `true`, the user can turn S/MIME signing on or off in Settings. Applicable for "EAS" only.
	UserOverrideable *bool `json:"UserOverrideable,omitempty"`
	// If `true`, the user can select an S/MIME signing identity in Settings. Applicable for "EAS" only.
	IdentityUserOverrideable *bool `json:"IdentityUserOverrideable,omitempty"`
}

// Settings for S/MIME encryption. Applicable for "EAS" only.
type AccountExchangeSMIMEEncryption struct {
	// If `true`, the system enables S/MIME encryption by default, which the user can't override if `PerMessageSwitchEnabled` is `false`. Applicable for "EAS" only.
	Enabled bool `json:"Enabled" required:"true"`
	// The identifier of an asset declaration containing the identity required for S/MIME encryption. The system attaches the public certificate to outgoing mail to allow the user to receive encrypted mail. When the user sends encrypted mail, the system uses the public certificate to encrypt the copy of the mail in their Sent mailbox. Applicable for "EAS" only.
	IdentityAssetReference *string `json:"IdentityAssetReference,omitempty"`
	// If `true`, the user can turn S/MIME encryption by default on or off in Settings. Applicable for "EAS" only.
	UserOverrideable *bool `json:"UserOverrideable,omitempty"`
	// If `true`, the user can select an S/MIME signing identity in Settings. Applicable for "EAS" only.
	IdentityUserOverrideable *bool `json:"IdentityUserOverrideable,omitempty"`
	// If `true`, the system enables the per-message encryption switch in the compose view. Applicable for "EAS" only.
	PerMessageSwitchEnabled *bool `json:"PerMessageSwitchEnabled,omitempty"`
}

// The declaration to configure a Google account.
type AccountGoogle struct {
	// The name that apps show to the user for this Google account. If not present, the system generates a suitable default.
	VisibleName *string `json:"VisibleName,omitempty"`
	// The identifier of an asset declaration that contains the user identity for this Google account. Set the corresponding asset type to `UserIdentity` and ensure that it contains an `EmailAddress` key that specifies the full Google email address for the account.
	UserIdentityAssetReference string `json:"UserIdentityAssetReference" required:"true"`
}

func (p *AccountGoogle) DeclarationType() string {
	return "com.apple.configuration.account.google"
}

// The declaration to configure a Lightweight Directory Access Protocol (LDAP) account.
type AccountLDAP struct {
	// The name that apps show to the user for this LDAP account. If not present, the system generates a suitable default.
	VisibleName *string `json:"VisibleName,omitempty"`
	// The hostname or IP address of the LDAP server.
	HostName string `json:"HostName" required:"true"`
	// The port number or IP address of the LDAP server.
	Port *int64 `json:"Port,omitempty"`
	// The identifier of an asset declaration that contains the credentials for this account. Set the corresponding asset type to `CredentialUserNameAndPassword`.
	AuthenticationCredentialsAssetReference *string `json:"AuthenticationCredentialsAssetReference,omitempty"`
	// The array of nodes to start LDAP searches from. There must be at least one node for this account to be useful. macOS only searches one node and ignores other items in the array.
	SearchSettings *[]*SearchSettingsItem `json:"SearchSettings,omitempty"`
}

func (p *AccountLDAP) DeclarationType() string {
	return "com.apple.configuration.account.ldap"
}

type SearchSettingsItem struct {
	// The description of this search setting in the Contacts and Settings apps. If not present, the apps display no name.
	VisibleDescription *string `json:"VisibleDescription,omitempty"`
	// The path to the node where a search starts. For example, `ou=people,o=example corp`.
	SearchBase string `json:"SearchBase" required:"true"`
	// The type of recursion to use in the search:
	// - `Base`: The search uses only the `SearchBase` node.
	// - `OneLevel`: The search uses the `SearchBase` node and its immediate children.
	// - `Subtree`: The search uses the `SearchBase` node and all its children, regardless of depth.
	Scope *Scope `default:"Subtree" json:"Scope,omitempty"`
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
type AccountMail struct {
	// The name that apps show to the user for this mail account. If not present, the system generates a suitable default.
	VisibleName *string `json:"VisibleName,omitempty"`
	// The identifier of an asset declaration that contains the user identity for this account. Set the corresponding asset type to `UserIdentity`.
	UserIdentityAssetReference *string `json:"UserIdentityAssetReference,omitempty"`
	// The settings for the incoming mail server for this account.
	IncomingServer IncomingServer `json:"IncomingServer" required:"true"`
	// The settings for the outgoing mail server for this account.
	OutgoingServer OutgoingServer `json:"OutgoingServer" required:"true"`
	// Settings for S/MIME.
	SMIME *AccountMailSMIME `json:"SMIME,omitempty"`
}

func (p *AccountMail) DeclarationType() string {
	return "com.apple.configuration.account.mail"
}

// The settings for the incoming mail server for this account.
type IncomingServer struct {
	// The mail protocol this account uses.
	ServerType ServerType `json:"ServerType" required:"true"`
	// The host name for the incoming mail server.
	HostName string `json:"HostName" required:"true"`
	// The port number for the incoming mail server.
	Port *int64 `json:"Port,omitempty"`
	// The authentication method for the incoming mail server.
	AuthenticationMethod IncomingServerAuthenticationMethod `json:"AuthenticationMethod" required:"true"`
	// The identifier of an asset declaration that contains the credentials for this account to authenticate with an incoming mail server. The corresponding asset must be of type `CredentialUserNameAndPassword`.
	// If the `AuthenticationMethod` is `None`, this field must be blank. Otherwise, the declaration must contain this field.
	AuthenticationCredentialsAssetReference *string `json:"AuthenticationCredentialsAssetReference,omitempty"`
	// The path prefix for the IMAP server. The system uses this only when `ServerType` is `IMAP`.
	IMAPPathPrefix *string `json:"IMAPPathPrefix,omitempty"`
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
	HostName string `json:"HostName" required:"true"`
	// The port number for the outgoing mail server.
	Port *int64 `json:"Port,omitempty"`
	// The authentication method for the outgoing mail server.
	AuthenticationMethod OutgoingServerAuthenticationMethod `json:"AuthenticationMethod" required:"true"`
	// The identifier of an asset declaration that contains the credentials for this account to authenticate with an outgoing mail server. The corresponding asset must be of type `CredentialUserNameAndPassword`.
	// If the `AuthenticationMethod` is `None`, this field must be blank. Otherwise, the declaration must contain this field.
	AuthenticationCredentialsAssetReference *string `json:"AuthenticationCredentialsAssetReference,omitempty"`
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
	Signing *AccountMailSMIMESigning `json:"Signing,omitempty"`
	// Settings for S/MIME encryption.
	Encryption *AccountMailSMIMEEncryption `json:"Encryption,omitempty"`
}

// Settings for S/MIME signing.
type AccountMailSMIMESigning struct {
	// If `true`, the system enables S/MIME signing.
	Enabled bool `json:"Enabled" required:"true"`
	// Specifies the identifier of an asset declaration containing the identity required for S/MIME signing of messages sent from this account.
	IdentityAssetReference *string `json:"IdentityAssetReference,omitempty"`
	// If `true`, the user can turn S/MIME signing on or off in Settings.
	UserOverrideable *bool `json:"UserOverrideable,omitempty"`
	// If `true`, the user can select an S/MIME signing identity in Settings.
	IdentityUserOverrideable *bool `json:"IdentityUserOverrideable,omitempty"`
}

// Settings for S/MIME encryption.
type AccountMailSMIMEEncryption struct {
	// If `true`, the system enables S/MIME encryption by default, which the user can't override if `PerMessageSwitchEnabled` is `false`.
	Enabled bool `json:"Enabled" required:"true"`
	// Specifies the identifier of an asset declaration containing the identity required for S/MIME encryption. The system attaches the public certificate to outgoing mail to allow the user to receive encrypted mail. When the user sends encrypted mail, the system uses the public certificate to encrypt the copy of the mail in their Sent mailbox.
	IdentityAssetReference *string `json:"IdentityAssetReference,omitempty"`
	// If `true`, the user can set the default value for S/MIME encryption to on or off in Settings.
	UserOverrideable *bool `json:"UserOverrideable,omitempty"`
	// If `true`, the user can select an S/MIME signing identity in Settings.
	IdentityUserOverrideable *bool `json:"IdentityUserOverrideable,omitempty"`
	// If `true`, the system enables the per-message encryption switch in the compose view.
	PerMessageSwitchEnabled *bool `json:"PerMessageSwitchEnabled,omitempty"`
}

// The declaration to configure account settings.
type AccountSettings struct {
	// If present, configures settings for Exchange accounts.
	Exchange *Exchange `json:"Exchange,omitempty"`
}

func (p *AccountSettings) DeclarationType() string {
	return "com.apple.configuration.account.settings"
}

// If present, configures settings for Exchange accounts.
type Exchange struct {
	// If `false`, the system doesn't alert the user that an Exchange account has lost, or is about to lose, access because Microsoft is ending support for Exchange Web Services. Only applies to Exchange accounts installed by an organization.
	AllowEWSRetirementAlerts *bool `json:"AllowEWSRetirementAlerts,omitempty"`
}

// The declaration to configure a subscribed calendar.
type AccountSubscribedCalendar struct {
	// The name that apps show to the user for this calendar account. If not present, the system generates a suitable default.
	VisibleName *string `json:"VisibleName,omitempty"`
	// The URL of the subscribed calendar, which needs to start with `https://`.
	CalendarURL string `json:"CalendarURL" required:"true"`
	// The identifier of an asset declaration that contains the credentials for this account to authenticate with a calendar server. Set the corresponding asset type to `CredentialUserNameAndPassword`.
	AuthenticationCredentialsAssetReference *string `json:"AuthenticationCredentialsAssetReference,omitempty"`
}

func (p *AccountSubscribedCalendar) DeclarationType() string {
	return "com.apple.configuration.account.subscribed-calendar"
}

// The declaration to configure a managed app.
type AppManaged struct {
	// The App Store ID of the managed app that's downloaded from the App Store.
	// Only one of `AppStoreID`, `BundleID`, `ManifestURL`, or `AppComposedIdentifier` needs to be present.
	AppStoreID *string `json:"AppStoreID,omitempty"`
	// The bundle ID of the managed app that's downloaded from the App Store.
	// Only one of `AppStoreID`, `BundleID`, `ManifestURL`, or `AppComposedIdentifier` needs to be present.
	BundleID *string `json:"BundleID,omitempty"`
	// The URL of the manifest for the managed app that the device downloads from a web site. The manifest is a `ManifestURL` property list.
	// Only one of `AppStoreID`, `BundleID`, `ManifestURL`, or `AppComposedIdentifier` needs to be present.
	ManifestURL *string `json:"ManifestURL,omitempty"`
	// A string that specifies the composed identifier of an existing app that needs to be managed. The device uses this to take over management of an app installed by some other process, for example installed manually by the user, or via a package configuration. If the app isn't present when the device applies the configuration, the device takes over management of it when it does install. Management of the app occurs only if its code signature matches the composed identifier.
	// The following rules apply when the device takes over management:
	// - If the `InstallBehavior.Install` key is set to `Required`, the device takes over management of the app.
	// - If the `InstallBehavior.Install` key is set to `Optional`, the device takes over management of the app when the user "installs" it using an MDM management app.
	// The format of the composed identifier is either "Bundle-ID" or "Bundle-ID (Team-ID)". "Bundle-ID" is the bundle identifier string of the provider. "Team-ID" is the team identifier from the provider's code signature. For example, "com.example.app" for the bundle ID format, or "com.example.app (ABCD1234)" for the team ID format.
	// In macOS, only one of `AppStoreID`, `BundleID`, or `AppComposedIdentifier` needs to be present.
	AppComposedIdentifier *string `json:"AppComposedIdentifier,omitempty"`
	// If `true`, the device installs an iOS or iPadOS app that runs on a Mac with Apple silicon. This is only used when the app is an App Store app.
	IOSApp *bool `json:"iOSApp,omitempty"`
	// A dictionary that describes how and when to install the app.
	InstallBehavior *AppManagedInstallBehavior `json:"InstallBehavior,omitempty"`
	// A dictionary that specifies how the device updates apps.
	UpdateBehavior *UpdateBehavior `json:"UpdateBehavior,omitempty"`
	// If `true`, backups contain the app and its data.
	IncludeInBackup *bool `json:"IncludeInBackup,omitempty"`
	// A dictionary of values to associate with the app.
	Attributes *Attributes `json:"Attributes,omitempty"`
	// A dictionary of app config data and credentials.
	AppConfig *AppConfig `json:"AppConfig,omitempty"`
	// A dictionary of extension config data and credentials.
	ExtensionConfigs *map[string]AppConfigDictionary `json:"ExtensionConfigs,omitempty"`
	// The identifier of an asset declaration containing a reference to the app config data. The device provides the app config data to the app using the MDMv1 behavior. The corresponding asset needs to be of type `com.apple.asset.data`. The referenced data needs to be a property list file, and the asset's "ContentType"
	// value set to match the data type.
	LegacyAppConfigAssetReference *string `json:"LegacyAppConfigAssetReference,omitempty"`
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
	Install *AppManagedInstallBehaviorInstall `default:"Optional" json:"Install,omitempty"`
	// A dictionary that describes the app's license.
	License *License `json:"License,omitempty"`
	// The App Store external version identifier (EVID) of the version of the app the device installs. You can retrieve this value from the App Store. For more information, see `Apps and books metadata for organizations`. This key is ignored if the app isn't an App Store app.
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
	Version *int64 `json:"Version,omitempty"`
	// Indicates how the device uses a cellular network when it downloads the app for automatic install or update operations:
	// - `AlwaysOn`: The device downloads apps of any size using a cellular network.
	// - `AlwaysOff`: The device doesn't download apps using a cellular network. The device pauses the automatic install or update operation until a different network is active.
	// - `StoreSettings`: The device uses the settings for the corresponding store when downloading apps.
	// The device always uses the store settings to download apps when the install or update operation is user initiated.
	AllowDownloadsOverCellular *AllowDownloadsOverCellular `default:"StoreSettings" json:"AllowDownloadsOverCellular,omitempty"`
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
	Assignment *Assignment `json:"Assignment,omitempty"`
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

// Indicates how the device uses a cellular network when it downloads the app for automatic install or update operations:
// - `AlwaysOn`: The device downloads apps of any size using a cellular network.
// - `AlwaysOff`: The device doesn't download apps using a cellular network. The device pauses the automatic install or update operation until a different network is active.
// - `StoreSettings`: The device uses the settings for the corresponding store when downloading apps.
// The device always uses the store settings to download apps when the install or update operation is user initiated.
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
	// When you specify the `InstallBehavior.Version` key, the device ignores this key and Automatic App Updates are disabled.
	// In macOS, the device ignores this setting if the `AppComposedIdentifier` key is set in the configuration.
	AutomaticAppUpdates AutomaticAppUpdates `json:"AutomaticAppUpdates" required:"true"`
}

// Specifies whether the device automatically updates the app:
// - `AlwaysOn`: The device automatically updates the app to the latest version. For App Store apps, the device periodically checks the store for updates. For Enterprise apps, the device periodically downloads the manifest file and compares it to the previous manifest file. If the device detects a change to the bundle version in the manifest, it downloads and updates the app.
// - `AlwaysOff`: The device never automatically updates the app.
// - `StoreSettings`: The device uses the settings for the corresponding store to determine when to automatically update the app. For Enterprise apps, this setting behaves the same as `AlwaysOff`.
// When you specify the `InstallBehavior.Version` key, the device ignores this key and Automatic App Updates are disabled.
// In macOS, the device ignores this setting if the `AppComposedIdentifier` key is set in the configuration.
type AutomaticAppUpdates string

const (
	AutomaticAppUpdatesAlwaysOn      AutomaticAppUpdates = "AlwaysOn"
	AutomaticAppUpdatesAlwaysOff     AutomaticAppUpdates = "AlwaysOff"
	AutomaticAppUpdatesStoreSettings AutomaticAppUpdates = "StoreSettings"
)

// A dictionary of values to associate with the app.
type Attributes struct {
	// An array of domain names to associate with the app.
	AssociatedDomains *[]string `json:"AssociatedDomains,omitempty"`
	// If `true`, the system enables direct downloads for the `AssociatedDomains`.
	AssociatedDomainsEnableDirectDownloads *bool `json:"AssociatedDomainsEnableDirectDownloads,omitempty"`
	// The cellular slice identifier, which can be the data network name (DNN) or app category. For DNN, encode the value as "DNN:name", where "name" is the carrier-provided DNN name. For app category, encode the value as "AppCategory:category", where "category" is a carrier-provided string such as "Enterprise1".
	CellularSliceUUID *string `json:"CellularSliceUUID,omitempty"`
	// The UUID of the content filter to associate with the app.
	ContentFilterUUID *string `json:"ContentFilterUUID,omitempty"`
	// The UUID of the DNS proxy to associate with the app.
	DNSProxyUUID *string `json:"DNSProxyUUID,omitempty"`
	// If `false`, the system prevents the user from hiding the app. It doesn't affect the user's ability to leave it in the App Library, while removing it from the Home Screen.
	Hideable *bool `json:"Hideable,omitempty"`
	// If `false`, the system prevents the user from locking the app. This also prevents the user from hiding the app.
	Lockable *bool `json:"Lockable,omitempty"`
	// The UUID of the relay to associate with the app.
	RelayUUID *string `json:"RelayUUID,omitempty"`
	// If `true`, the device automatically locks after every transaction that requires a customer's card PIN. If `false`, the user can choose the behavior.
	TapToPayScreenLock *bool `json:"TapToPayScreenLock,omitempty"`
	// The UUID of the VPN to associate with the app.
	VPNUUID *string `json:"VPNUUID,omitempty"`
}

// A dictionary of app config data and credentials.
type AppConfig struct {
	// Specifies the identifier of an asset declaration containing a reference to the app or extension config data. The corresponding asset needs to be of type `com.apple.asset.data`. The referenced data needs to be a property list file, and the asset's "ContentType" value set to match the data type.
	DataAssetReference *string `json:"DataAssetReference,omitempty"`
	// Provides passwords to the managed app or extension. Each element in the array contains a password asset reference and an associated identifier which the app or extension uses to look up the password.
	Passwords *[]PasswordAppConfigItem `json:"Passwords,omitempty"`
	// Provides identities to the managed app or extension. Each element in the array contains an identity asset reference and an associated identifier which the app or extension uses to look up the identity.
	Identities *[]IdentityAppConfigItem `json:"Identities,omitempty"`
	// Provides certificates to the managed app or extension. Each element in the array contains a certificate asset reference and an associated identifier which the app or extension uses to look up the certificate.
	Certificates *[]CertificateAppConfigItem `json:"Certificates,omitempty"`
}

// A dictionary of values associated with a credential config.
type PasswordAppConfigItem struct {
	// The app or extension uses this identifier to fetch the corresponding password using the `ManagedApp` framework. App developers define the values for these identifiers.
	Identifier string `json:"Identifier" required:"true"`
	// Specifies the identifier of an asset declaration containing a username and password. The `ManagedApp` framework makes the password available to the app or extension. The `ManagedApp` framework ignores the username.
	AssetReference string `json:"AssetReference" required:"true"`
}

// A dictionary of values associated with a credential config.
type IdentityAppConfigItem struct {
	// The app or extension uses this identifier to fetch the corresponding identity using the `ManagedApp` framework. App developers define the values for these identifiers.
	Identifier string `json:"Identifier" required:"true"`
	// Specifies the identifier of an asset declaration containing an identity.
	AssetReference string `json:"AssetReference" required:"true"`
}

// A dictionary of values associated with a credential config.
type CertificateAppConfigItem struct {
	// The app or extension uses this identifier to fetch the corresponding certificate using the `ManagedApp` framework. App developers define the values for these identifiers.
	Identifier string `json:"Identifier" required:"true"`
	// Specifies the identifier of an asset declaration containing a certificate.
	AssetReference string `json:"AssetReference" required:"true"`
}

// A dictionary mapping extension composed identifiers to the extension config data and credentials.
// The format of the composed identifier is either "Bundle-ID" or "Bundle-ID (Team-ID)". "Bundle-ID" is the bundle identifier string of the provider. "Team-ID" is the team identifier from the provider's code signature. For example, "com.example.app" for the bundle ID format, or "com.example.app (ABCD1234)" for the team ID format.
type AppConfigDictionary struct {
	// Specifies the identifier of an asset declaration containing a reference to the app or extension config data. The corresponding asset needs to be of type `com.apple.asset.data`. The referenced data needs to be a property list file, and the asset's "ContentType" value set to match the data type.
	DataAssetReference *string `json:"DataAssetReference,omitempty"`
	// Provides passwords to the managed app or extension. Each element in the array contains a password asset reference and an associated identifier which the app or extension uses to look up the password.
	Passwords *[]PasswordAppConfigItem `json:"Passwords,omitempty"`
	// Provides identities to the managed app or extension. Each element in the array contains an identity asset reference and an associated identifier which the app or extension uses to look up the identity.
	Identities *[]IdentityAppConfigItem `json:"Identities,omitempty"`
	// Provides certificates to the managed app or extension. Each element in the array contains a certificate asset reference and an associated identifier which the app or extension uses to look up the certificate.
	Certificates *[]CertificateAppConfigItem `json:"Certificates,omitempty"`
}

// The declaration to configure app settings.
type AppSettings struct {
	// The dictionary of allowed app settings.
	Allowed *Allowed `json:"Allowed,omitempty"`
	// The dictionary of app settings.
	Privacy *AppSettingsPrivacy `json:"Privacy,omitempty"`
}

func (p *AppSettings) DeclarationType() string {
	return "com.apple.configuration.app.settings"
}

// The dictionary of allowed app settings.
type Allowed struct {
	// If present, the device only shows or launches apps with bundle IDs in the array. Include the value `com.apple.webapp` to allow all webclips. This applies to App Store apps, marketplace apps, and locally installed apps (using Configurator, Xcode, and so forth).
	AllowedApps *[]string `json:"AllowedApps,omitempty"`
	// If present, the device prevents showing or launching apps with bundle IDs in the
	// array. Include the value `com.apple.webapp` to restrict all webclips. This applies to
	// App Store apps, marketplace apps, and locally installed apps (using Configurator,
	// Xcode, and so forth).
	// > Note:
	// > Denying system apps may disable other functionality. For example, denying the App
	// Store app may prevent users from accepting the terms and conditions for the user-based
	// Volume Purchase Program (VPP).
	DeniedApps *[]string `json:"DeniedApps,omitempty"`
	// If present, the device only allows binaries that match the binary identifier properties to run. A binary only matches when all the binary identifiers match. The device always runs system critical processes. Use "codesign -dvvv <path_to_binary>" to show the information you need to generate these values.
	AllowedBinaries *[]BinaryIdentifier `json:"AllowedBinaries,omitempty"`
	// If present, the device doesn't allow binaries that match the binary identifier properties to run. A binary only matches when all the binary identifiers match.
	DeniedBinaries *[]BinaryIdentifier `json:"DeniedBinaries,omitempty"`
	// If `true`, the device implicitly includes managed apps in the effective allow list when `AllowedBinaries` is present.
	AlwaysAllowManagedApps *bool `json:"AlwaysAllowManagedApps,omitempty"`
}

// Dictionary containing one or more identifier fields to match a binary.
type BinaryIdentifier struct {
	// The 40 character code signature code directory hash of the binary.
	CDHash *string `json:"CDHash,omitempty"`
	// The code signature signing identifier of the binary.
	SigningID *string `json:"SigningID,omitempty"`
	// The code signature team identifier of the binary. Use the value "\*APPLE\*" instead of an empty string for Apple binaries with an empty team identifier.
	TeamID *string `json:"TeamID,omitempty"`
	// The file system path prefix to match binaries.
	PathPrefix *string `json:"PathPrefix,omitempty"`
	// The code signing state to match binaries.
	SigningState *SigningState `default:"All" json:"SigningState,omitempty"`
}

// The code signing state to match binaries.
type SigningState string

const (
	SigningStateAll         SigningState = "All"
	SigningStateTestFlight  SigningState = "TestFlight"
	SigningStateDeveloperID SigningState = "DeveloperID"
	SigningStateEnterprise  SigningState = "Enterprise"
	SigningStateAppStore    SigningState = "AppStore"
	SigningStateApple       SigningState = "Apple"
)

// The dictionary of app settings.
type AppSettingsPrivacy struct {
	// The dictionary of app privacy permission defaults. Each key in the dictionary is an app identifier. The dictionary values represent the permission defaults that the device applies for each matching app.
	// In iOS, the app identifier is a bundle ID, for example, "com.example.app".
	// In macOS, the app identifier is a composed identifier. The format of the composed identifier is "Bundle-ID {Designated-Requirement}". "Bundle-ID" is the bundle identifier string of the app. "Designated-Requirement" is the designated requirement string the device uses to match the code signature of the app. For example, "com.example.app {anchor apple generic}". The device only applies defaults for an app if its code signature matches the composed identifier.
	PermissionDefaults *map[string]AppDictionary `json:"PermissionDefaults,omitempty"`
}

// The dictionary that defines the app privacy permission defaults. Each key is an app identifier.
type AppDictionary struct {
	// Text you provide that clearly explains to the user the reason why the organization requires these app permission defaults. The device includes this text in the permission consent prompt it displays when it launches the app.
	OrganizationJustification string `json:"OrganizationJustification" required:"true"`
	// Controls whether an app privacy permission default is set.
	// * `None`: No app privacy permission default is set for use of accessibility.
	// * `Allow`: The app privacy permission default is set to allow use of accessibility.
	Accessibility *Accessibility `json:"Accessibility,omitempty"`
	// Controls whether an app privacy permission default is set.
	// * `None`: No app privacy permission default is set for use of Bluetooth.
	// * `Allow`: The app privacy permission default is set to allow use of Bluetooth.
	Bluetooth *Bluetooth `json:"Bluetooth,omitempty"`
	// Controls whether an app privacy permission default is set.
	// * `None`: No app privacy permission default is set for use of the camera.
	// * `Allow`: The app privacy permission default is set to allow use of the camera.
	Camera *AppDictionaryCamera `json:"Camera,omitempty"`
	// Controls whether an app privacy permission default is set.
	// * `None`: No app privacy permission default is set for use of dictation.
	// * `Allow`: The app privacy permission default is set to allow use of dictation.
	Dictation *Dictation `json:"Dictation,omitempty"`
	// Controls whether an app privacy permission default is set.
	// * `None`: No app privacy permission default is set for use of the local network.
	// * `Allow`: The app privacy permission default is set to allow use of the local network.
	LocalNetwork *LocalNetwork `json:"LocalNetwork,omitempty"`
	// Controls whether an app privacy permission default is set.
	// * `None`: No app privacy permission default is set for access to location.
	// * `WhileUsing`: The app privacy permission default is set to allow access to location only while the user is using the app In macOS, this is equivalent to `Always`.
	// * `Always`: The app privacy permission default is set to allow access to location always.
	Location *Location `json:"Location,omitempty"`
	// Controls whether an app privacy permission default is set.
	// * `None`: No app privacy permission default is set for access to precise location.
	// * `Approximate`: The app privacy permission default is set to allow approximate access to location.
	// * `Precise`: The app privacy permission default is set to allow precise access to location.
	LocationAccuracy *LocationAccuracy `json:"LocationAccuracy,omitempty"`
	// Controls whether an app privacy permission default is set.
	// * `None`: No app privacy permission default is set for use of the microphone.
	// * `Allow`: The app privacy permission default is set to allow use of the microphone.
	Microphone *AppDictionaryMicrophone `json:"Microphone,omitempty"`
}

// Controls whether an app privacy permission default is set.
// * `None`: No app privacy permission default is set for use of accessibility.
// * `Allow`: The app privacy permission default is set to allow use of accessibility.
type Accessibility string

const (
	AccessibilityNone  Accessibility = "None"
	AccessibilityAllow Accessibility = "Allow"
)

// Controls whether an app privacy permission default is set.
// * `None`: No app privacy permission default is set for use of Bluetooth.
// * `Allow`: The app privacy permission default is set to allow use of Bluetooth.
type Bluetooth string

const (
	BluetoothNone  Bluetooth = "None"
	BluetoothAllow Bluetooth = "Allow"
)

// Controls whether an app privacy permission default is set.
// * `None`: No app privacy permission default is set for use of the camera.
// * `Allow`: The app privacy permission default is set to allow use of the camera.
type AppDictionaryCamera string

const (
	AppDictionaryCameraNone  AppDictionaryCamera = "None"
	AppDictionaryCameraAllow AppDictionaryCamera = "Allow"
)

// Controls whether an app privacy permission default is set.
// * `None`: No app privacy permission default is set for use of dictation.
// * `Allow`: The app privacy permission default is set to allow use of dictation.
type Dictation string

const (
	DictationNone  Dictation = "None"
	DictationAllow Dictation = "Allow"
)

// Controls whether an app privacy permission default is set.
// * `None`: No app privacy permission default is set for use of the local network.
// * `Allow`: The app privacy permission default is set to allow use of the local network.
type LocalNetwork string

const (
	LocalNetworkNone  LocalNetwork = "None"
	LocalNetworkAllow LocalNetwork = "Allow"
)

// Controls whether an app privacy permission default is set.
// * `None`: No app privacy permission default is set for access to location.
// * `WhileUsing`: The app privacy permission default is set to allow access to location only while the user is using the app In macOS, this is equivalent to `Always`.
// * `Always`: The app privacy permission default is set to allow access to location always.
type Location string

const (
	LocationNone       Location = "None"
	LocationWhileUsing Location = "WhileUsing"
	LocationAlways     Location = "Always"
)

// Controls whether an app privacy permission default is set.
// * `None`: No app privacy permission default is set for access to precise location.
// * `Approximate`: The app privacy permission default is set to allow approximate access to location.
// * `Precise`: The app privacy permission default is set to allow precise access to location.
type LocationAccuracy string

const (
	LocationAccuracyNone        LocationAccuracy = "None"
	LocationAccuracyApproximate LocationAccuracy = "Approximate"
	LocationAccuracyPrecise     LocationAccuracy = "Precise"
)

// Controls whether an app privacy permission default is set.
// * `None`: No app privacy permission default is set for use of the microphone.
// * `Allow`: The app privacy permission default is set to allow use of the microphone.
type AppDictionaryMicrophone string

const (
	AppDictionaryMicrophoneNone  AppDictionaryMicrophone = "None"
	AppDictionaryMicrophoneAllow AppDictionaryMicrophone = "Allow"
)

// The declaration to configure audio accessory settings.
type AudioAccessorySettings struct {
	// A dictionary that describes audio accessory temporary pairing behavior. The device enables temporary pairing when this key is present and the `Disabled` key isn't `false`. The device doesn't synchronize pairing information with iCloud when temporary pairing is active.
	TemporaryPairing *TemporaryPairing `json:"TemporaryPairing,omitempty"`
}

func (p *AudioAccessorySettings) DeclarationType() string {
	return "com.apple.configuration.audio-accessory.settings"
}

// A dictionary that describes audio accessory temporary pairing behavior. The device enables temporary pairing when this key is present and the `Disabled` key isn't `false`. The device doesn't synchronize pairing information with iCloud when temporary pairing is active.
type TemporaryPairing struct {
	// If `true`, temporary pairing of audio accessories is disabled.
	Disabled *bool `json:"Disabled,omitempty"`
	// A dictionary providing configuration for temporary pairing. Required if `Disabled` isn't present or is `false`.
	Configuration *Configuration `json:"Configuration,omitempty"`
}

// A dictionary providing configuration for temporary pairing. Required if `Disabled` isn't present or is `false`.
type Configuration struct {
	// A dictionary that describes when the device automatically unpairs temporarily paired audio accessories.
	UnpairingTime UnpairingTime `json:"UnpairingTime" required:"true"`
}

// A dictionary that describes when the device automatically unpairs temporarily paired audio accessories.
type UnpairingTime struct {
	// A string that specifies the device's unpairing policy.
	// - `None`: The device doesn't automatically unpair. Use this only with a return to service device that you erase and reenroll when assigning it from one user to another.
	// - `Hour`: The device automatically unpairs temporarily paired audio accessories at the local time that the `Hour` key specifies.
	Policy Policy `json:"Policy" required:"true"`
	// The local time hour (24-hour clock) when the device automatically unpairs temporarily paired audio accessories. Required when setting the `Policy` key to `Hour`.
	Hour *int64 `json:"Hour,omitempty"`
}

// A string that specifies the device's unpairing policy.
// - `None`: The device doesn't automatically unpair. Use this only with a return to service device that you erase and reenroll when assigning it from one user to another.
// - `Hour`: The device automatically unpairs temporarily paired audio accessories at the local time that the `Hour` key specifies.
type Policy string

const (
	PolicyNone Policy = "None"
	PolicyHour Policy = "Hour"
)

// The declaration to configure the Content Caching service.
type ContentCaching struct {
	// If `true`, the system purges content from the cache automatically when it needs disk space for other apps when free disk space runs low on the computer. Set to `false` to maximize effectiveness of Content Caching.
	AllowCacheDelete *bool `json:"AllowCacheDelete,omitempty"`
	// If `true`, the system caches the user's iCloud data. Changes to this value don't have an immediate effect. Clients may take some time to react to changes.
	// > Note:
	// > At least one of the `AllowPersonalCaching` or `AllowSharedCaching` keys need to be `true`.
	AllowPersonalCaching *bool `json:"AllowPersonalCaching,omitempty"`
	// If `true`, the system caches non-iCloud content, such as apps and software updates. Changes to this value don't have an immediate effect. Clients may take some time to react to changes.
	// > Note:
	// > At least one of the `AllowPersonalCaching` or `AllowSharedCaching` keys need to be `true`.
	AllowSharedCaching *bool `json:"AllowSharedCaching,omitempty"`
	// If `true`, the system automatically activates the content cache when possible and prevents disabling it. If `allowContentCaching` is `false`, `AutoActivation` is also `false`.
	// Removing a profile that set `AutoActivation` to `true` doesn't deactivate the Content Cache.
	AutoActivation *bool `json:"AutoActivation,omitempty"`
	// If `true`, the system automatically enables Internet connection sharing when possible and prevent disabling Internet connection sharing. `DenyTetheredCaching` overrides `AutoEnableTetheredCaching`. Tethered caching requires Content Caching.
	AutoEnableTetheredCaching *bool `json:"AutoEnableTetheredCaching,omitempty"`
	// The maximum number of bytes of disk space to use for the content cache. Set to `0` for unlimited disk space. Also serves as the upper bound to the `PersonalCacheLimit`.
	CacheLimit *int64 `default:"0" json:"CacheLimit,omitempty"`
	// The path to the directory used to store cached content. Changing this setting manually doesn't automatically move cached content from the old location to the new one. To move content automatically, use the Sharing preference's Content Caching pane. The value must be (or end with) `/Library/Application Support/Apple/AssetCache/Data`.
	// The system creates a directory and its intermediates for the given data path if it doesn't already exist. The directory is owned by `_assetcache:_assetcache` and has mode 0750. Its immediate parent directory (`.../Library/Application Support/Apple/AssetCache`) is owned by `_assetcache:_assetcache` and has mode `0755`.
	DataPath *string `default:"/Library/Application Support/Apple/AssetCache/Data" json:"DataPath,omitempty"`
	// If `true`, the system disables Content Caching. This is the inverse of the `allowContentCaching` restriction in MDM. It overrides the `AutoActivation` key. Use this key to prevent launching the Content Caching service.
	DenyActivation *bool `json:"DenyActivation,omitempty"`
	// If `true`, the system disables tethered caching.
	DenyTetheredCaching *bool `json:"DenyTetheredCaching,omitempty"`
	// The time interval in seconds the system uses to update the `StatusContentCacheInfo` declarative status item. The reporting interval can't be less than 60 (1 per minute) or larger than 86400 (1 per day), defaults to 0  (off)
	DeclarativeStatusInterval *int64 `default:"0" json:"DeclarativeStatusInterval,omitempty"`
	// If `true`, Content Caching displays exceptional conditions (alerts) as system notifications in the upper corner of the screen.
	DisplayAlerts *bool `json:"DisplayAlerts,omitempty"`
	// If `true`, the system prevents the computer from sleeping as long as Content Caching is on (System Preferences > Sharing > Content Caching is on). Customers who want Content Caching to be as available as much as possible should turn this setting on.
	KeepAwake *bool `json:"KeepAwake,omitempty"`
	// An array of dictionaries that describe a range of client IP addresses to serve.
	ListenRanges *[]*RangesItem `json:"ListenRanges,omitempty"`
	// If `true`, the content cache provides content to the clients in the `ListenRanges`.
	ListenRangesOnly *bool `json:"ListenRangesOnly,omitempty"`
	// If `true`, the content cache provides content to the clients in the union of the `ListenRanges`, `PeerListenRanges` and `Parents`.
	ListenWithPeersAndParents *bool `json:"ListenWithPeersAndParents,omitempty"`
	// If `true`, the content cache offers content to clients only on the same immediate local network only. The content cache offers no content to clients on other networks reachable by the content cache. If `LocalSubnetsOnly` is `true`, the system ignores `ListenRanges`.
	LocalSubnetsOnly *bool `json:"LocalSubnetsOnly,omitempty"`
	// If `true`, the Content Cache logs the IP address and port number of the clients that request content.
	LogClientIdentity *bool `json:"LogClientIdentity,omitempty"`
	// An array of the local IP addresses of other content caches that this cache should download from or upload to, instead of downloading from or uploading to Apple directly. The system ignores invalid addresses and addresses of computers that aren't content caches. The system skips Parent caches that become unavailable. If all parent content caches become unavailable, the content cache downloads from or uploads to Apple directly, until a parent content cache becomes available again.
	Parents *[]string `json:"Parents,omitempty"`
	// The policy to implement when choosing among more than one configured parent content cache. With every policy, the system skips parent caches that are temporarily unavailable. Allowed values:
	// - `first-available`: Always use the first available parent in the Parents list. Use this policy to designate permanent primary, secondary, and subsequent parents.
	// - `url-path-hash`: Hash the path part of the requested URL so that the same parent is always used for the same URL. This is useful for maximizing the size of the combined caches of     the parents.
	// - `random`: Choose a parent at random. Use this policy for load balancing.
	// - `round-robin`: Rotate through the parents in order. Use this policy for load balancing.
	// - `sticky-available`: Use the first available parent in the Parents list until it becomes unavailable, then advance to the next one. Use this policy for designating floating primary, secondary, and subsequent parents.
	ParentSelectionPolicy *ParentSelectionPolicy `default:"round-robin" json:"ParentSelectionPolicy,omitempty"`
	// An array of dictionaries describing a range of peer IP addresses that the content cache uses to filter its list of peers to query for content. The content cache only queries peers in `PeerFilterRanges`. When `PeerFilterRanges` is an empty array, the content cache doesn't query any peers.
	PeerFilterRanges *[]*RangesItem `json:"PeerFilterRanges,omitempty"`
	// An array of dictionaries describing a range of peer IP addresses the content cache responds to. When `PeerListenRanges` is an empty array, the content cache responds with an error to all cache queries.
	PeerListenRanges *[]*RangesItem `json:"PeerListenRanges,omitempty"`
	// If `true`, the content cache only peers with other content caches on the same immediate local network, rather than with content caches that use the same public IP address as the device. When `PeerLocalSubnetsOnly` is `true`, it overrides the configuration of `PeerFilterRanges` and `PeerListenRanges`. If the network changes, the local network peering restrictions update appropriately. If `false`, the content cache defers to `PeerFilterRanges` and `PeerListenRanges` for configuring the peering restrictions.
	PeerLocalSubnetsOnly *bool `json:"PeerLocalSubnetsOnly,omitempty"`
	// The maximum number of bytes of disk space to use for the personal content cache. The content cache limits the maximum value to the `CacheLimit` value. Set to `0` to use the overall `CacheLimit`.
	PersonalCacheLimit *int64 `default:"0" json:"PersonalCacheLimit,omitempty"`
	// The TCP port number on which the content cache accepts requests for uploads or downloads. Set to `0` to pick a random, available port.
	Port *int64 `default:"0" json:"Port,omitempty"`
	// An array of dictionaries describing a range of public IP addresses that the cloud servers should use for matching clients to content caches.
	PublicRanges *[]*RangesItem `json:"PublicRanges,omitempty"`
	// The target URL the system uses to send management statistics using an HTTP PUT request with a json dictionary as the payload. If the URL is an https URL (recommended) the `ManagementSecurityConfig` defines how the system secures the connection. The system sends the management statistics report to the target URL at an interval set by the `ManagementReportingInterval` value.
	ManagementStatusTarget *string `json:"ManagementStatusTarget,omitempty"`
	// If the `ManagementStatusTarget` is an https URL, the system uses this field to determine how the connection is secured.
	// - `no-cert`: The system uses regular https processing using the built-in anchor certificates.
	// - `signedByCACert`: The system requires the https connection's certificate to be signed by the certificate provided in the `ManagementStatusCertificateReference` field. If your organization has a CA infrastructure this is the best choice.
	// - `specificServerCert`: The system requires the https connection's certificate to match the certificate provided in the `ManagementStatusCertificateReference` field.
	ManagementSecurityConfig *ManagementSecurityConfig `default:"no-cert" json:"ManagementSecurityConfig,omitempty"`
	// The reporting interval in seconds the system uses when `ManagementStatusTarget` is present. The reporting interval can't be less than 60 (1 per minute) or larger than 86400 (1 per day), and defaults to 300 (1 per 5 min).
	ManagementReportingInterval *int64 `default:"300" json:"ManagementReportingInterval,omitempty"`
	// The identifier of an asset declaration that contains the certificate the system uses to verify the security of the status reporting https connection. Required when `ManagementSecurityConfig` is set to `signedByCACert` or `specificServerCert`.
	ManagementStatusCertificateReference *string `json:"ManagementStatusCertificateReference,omitempty"`
}

func (p *ContentCaching) DeclarationType() string {
	return "com.apple.configuration.content-cache.settings"
}

// A range of IP addresses to cache.
type RangesItem struct {
	// The IP address type.
	Type *RangesItemType `default:"IPv4" json:"type,omitempty"`
	// The first IP address in the range.
	First string `json:"first" required:"true"`
	// The last IP address in the range.
	Last string `json:"last" required:"true"`
}

// The IP address type.
type RangesItemType string

const (
	RangesItemTypeIPv4 RangesItemType = "IPv4"
	RangesItemTypeIPv6 RangesItemType = "IPv6"
)

// The policy to implement when choosing among more than one configured parent content cache. With every policy, the system skips parent caches that are temporarily unavailable. Allowed values:
// - `first-available`: Always use the first available parent in the Parents list. Use this policy to designate permanent primary, secondary, and subsequent parents.
// - `url-path-hash`: Hash the path part of the requested URL so that the same parent is always used for the same URL. This is useful for maximizing the size of the combined caches of     the parents.
// - `random`: Choose a parent at random. Use this policy for load balancing.
// - `round-robin`: Rotate through the parents in order. Use this policy for load balancing.
// - `sticky-available`: Use the first available parent in the Parents list until it becomes unavailable, then advance to the next one. Use this policy for designating floating primary, secondary, and subsequent parents.
type ParentSelectionPolicy string

const (
	ParentSelectionPolicyFirstAvailable  ParentSelectionPolicy = "first-available"
	ParentSelectionPolicyURLPathHash     ParentSelectionPolicy = "url-path-hash"
	ParentSelectionPolicyRandom          ParentSelectionPolicy = "random"
	ParentSelectionPolicyRoundRobin      ParentSelectionPolicy = "round-robin"
	ParentSelectionPolicyStickyAvailable ParentSelectionPolicy = "sticky-available"
)

// If the `ManagementStatusTarget` is an https URL, the system uses this field to determine how the connection is secured.
// - `no-cert`: The system uses regular https processing using the built-in anchor certificates.
// - `signedByCACert`: The system requires the https connection's certificate to be signed by the certificate provided in the `ManagementStatusCertificateReference` field. If your organization has a CA infrastructure this is the best choice.
// - `specificServerCert`: The system requires the https connection's certificate to match the certificate provided in the `ManagementStatusCertificateReference` field.
type ManagementSecurityConfig string

const (
	ManagementSecurityConfigNoCert             ManagementSecurityConfig = "no-cert"
	ManagementSecurityConfigSignedByCACert     ManagementSecurityConfig = "signedByCACert"
	ManagementSecurityConfigSpecificServerCert ManagementSecurityConfig = "specificServerCert"
)

// The declaration to configure disk management settings on the device.
type DiskManagementSettings struct {
	// The restrictions for the disk.
	Restrictions *Restrictions `json:"Restrictions,omitempty"`
}

func (p *DiskManagementSettings) DeclarationType() string {
	return "com.apple.configuration.diskmanagement.settings"
}

// The restrictions for the disk.
type Restrictions struct {
	// Specifies the mount policy for external storage:
	// - `Allowed`: The system can mount external storage that's read-write or read-only.
	// - `ReadOnly`: The system can only mount read-only external storage. Note that external storage that's read-write won't be mounted read-only.
	// - `Disallowed`: The system can't mount any external storage.
	ExternalStorage *ExternalStorage `json:"ExternalStorage,omitempty"`
	// Specifies the mount policy for network storage:
	// - `Allowed`: The system can mount network storage that's read-write or read-only.
	// - `ReadOnly`: The system can only mount read-only network storage. Note that network storage that's read-write won't be mounted read-only.
	// - `Disallowed`: The system can't mount any network storage.
	NetworkStorage *NetworkStorage `json:"NetworkStorage,omitempty"`
}

// Specifies the mount policy for external storage:
// - `Allowed`: The system can mount external storage that's read-write or read-only.
// - `ReadOnly`: The system can only mount read-only external storage. Note that external storage that's read-write won't be mounted read-only.
// - `Disallowed`: The system can't mount any external storage.
type ExternalStorage string

const (
	ExternalStorageAllowed    ExternalStorage = "Allowed"
	ExternalStorageReadOnly   ExternalStorage = "ReadOnly"
	ExternalStorageDisallowed ExternalStorage = "Disallowed"
)

// Specifies the mount policy for network storage:
// - `Allowed`: The system can mount network storage that's read-write or read-only.
// - `ReadOnly`: The system can only mount read-only network storage. Note that network storage that's read-write won't be mounted read-only.
// - `Disallowed`: The system can't mount any network storage.
type NetworkStorage string

const (
	NetworkStorageAllowed    NetworkStorage = "Allowed"
	NetworkStorageReadOnly   NetworkStorage = "ReadOnly"
	NetworkStorageDisallowed NetworkStorage = "Disallowed"
)

// The declaration to configure Extensible Single Sign-On.
type ExtensibleSSO struct {
	// The identifier of the provider to use for this configuration. Useful for apps that contain more than one DNS proxy extension.
	// In iOS and visionOS, the identifier is a bundle ID, for example, "com.example.app.sso-extension".
	// In macOS, the identifier is a composed identifier. The format of the composed identifier is "Bundle-ID (Team-ID)". "Bundle-ID" is the bundle identifier string of the app extension. "Team-ID" is the team identifier from the app extension's code signature. For example, "com.example.app.sso-extension (ABCD1234)".
	ExtensionComposedIdentifier string `json:"ExtensionComposedIdentifier" required:"true"`
	// The type of SSO.
	Type ExtensibleSSOType `json:"Type" required:"true"`
	// The realm name for `Credential` payloads. Use proper capitalization for this value. Ignored for `Redirect` payloads.
	Realm *string `json:"Realm,omitempty"`
	// A dictionary of arbitrary data passed through to the app extension.
	ExtensionData *map[string]any `json:"ExtensionData,omitempty"`
	// An array of URL prefixes of identity providers where the app extension performs SSO.
	// Required for `Redirect` payloads. Ignored for `Credential` payloads.
	// The URLs need to begin with `http://` or `https://`.
	// The system:
	// - Matches scheme and host name case-insensitively
	// - Doesn't allow query parameters and URL fragments
	// - Requires that the URLs of all installed Extensible SSO payloads are unique
	URLs *[]string `json:"URLs,omitempty"`
	// An array of host or domain names that apps can authenticate through the app extension.
	// Required for `Credential` payloads. Ignored for `Redirect` payloads.
	// The system:
	// - Matches host or domain names case-insensitively
	// - Requires that all the host and domain names of all installed Extensible SSO payloads are unique
	// > Note:
	// > Host names that begin with a "." are wildcard suffixes that match all subdomains; otherwise the host name needs be an exact match.
	Hosts *[]string `json:"Hosts,omitempty"`
	// An array of bundle identifiers of apps that don't use SSO provided by this extension.
	DeniedBundleIdentifiers *[]string `json:"DeniedBundleIdentifiers,omitempty"`
	// If set to `Cancel`, the system cancels authentication requests when the screen is locked. If set to `DoNotHandle`, the request continues without SSO instead. This doesn't apply to requests where `userInterfaceEnabled` is `false`, or for background `URLSession` requests.
	ScreenLockedBehavior *ScreenLockedBehavior `default:"Cancel" json:"ScreenLockedBehavior,omitempty"`
	// The dictionary to configure Platform SSO.
	PlatformSSO *PlatformSSO `json:"PlatformSSO,omitempty"`
}

func (p *ExtensibleSSO) DeclarationType() string {
	return "com.apple.configuration.extensible-sso"
}

// The type of SSO.
type ExtensibleSSOType string

const (
	ExtensibleSSOTypeCredential ExtensibleSSOType = "Credential"
	ExtensibleSSOTypeRedirect   ExtensibleSSOType = "Redirect"
)

// If set to `Cancel`, the system cancels authentication requests when the screen is locked. If set to `DoNotHandle`, the request continues without SSO instead. This doesn't apply to requests where `userInterfaceEnabled` is `false`, or for background `URLSession` requests.
type ScreenLockedBehavior string

const (
	ScreenLockedBehaviorCancel      ScreenLockedBehavior = "Cancel"
	ScreenLockedBehaviorDoNotHandle ScreenLockedBehavior = "DoNotHandle"
)

// The dictionary to configure Platform SSO.
type PlatformSSO struct {
	// The Platform SSO authentication method to use with the extension. Requires that the SSO Extension also support the method.
	AuthenticationMethod *PlatformSSOAuthenticationMethod `json:"AuthenticationMethod,omitempty"`
	// The token this device uses for registration with Platform SSO. Use it for silent registration with the Identity Provider. Requires that `AuthenticationMethod` in `PlatformSSO` isn't empty.
	RegistrationToken *string `json:"RegistrationToken,omitempty"`
	// If `true`, the system uses the same signing and encryption keys for all users.
	UseSharedDeviceKeys *bool `json:"UseSharedDeviceKeys,omitempty"`
	// The duration, in seconds, until the system requires a full login instead of a refresh. The default value is 64,800 (18 hours). The minimum value is 3600 (1 hour).
	LoginFrequency *int64 `default:"64800" json:"LoginFrequency,omitempty"`
	// If `true`, the system includes the device UDID and serial number in Platform SSO attestations.
	AllowDeviceIdentifiersInAttestation *bool `json:"AllowDeviceIdentifiersInAttestation,omitempty"`
	// Account display and profile settings.
	Account *Account `json:"Account,omitempty"`
	// Settings for creating new users via Platform SSO.
	UserCreation *UserCreation `json:"UserCreation,omitempty"`
	// Settings for authorization prompts and group management.
	Authorization *Authorization `json:"Authorization,omitempty"`
	// Settings for Access Key authentication.
	AccessKey *AccessKey `json:"AccessKey,omitempty"`
	// Policies for login, unlock, and FileVault behavior.
	Policies *Policies `json:"Policies,omitempty"`
	// Settings for web authentication behavior.
	WebAuthentication *WebAuthentication `json:"WebAuthentication,omitempty"`
	// A dictionary with keys and values that are strings. The keys are HTTP header names, and the values are HTTP header values. The device adds these headers to nonce, pre-authentication, login, key and key exchange requests for Platform SSO. The device limits the number of headers to 20, the length of the header name to 128 characters, and the length of the header value to 2048 characters.
	AdditionalHTTPHeaders *map[string]string `json:"AdditionalHTTPHeaders,omitempty"`
}

// The Platform SSO authentication method to use with the extension. Requires that the SSO Extension also support the method.
type PlatformSSOAuthenticationMethod string

const (
	PlatformSSOAuthenticationMethodPassword             PlatformSSOAuthenticationMethod = "Password"
	PlatformSSOAuthenticationMethodUserSecureEnclaveKey PlatformSSOAuthenticationMethod = "UserSecureEnclaveKey"
	PlatformSSOAuthenticationMethodSmartCard            PlatformSSOAuthenticationMethod = "SmartCard"
	PlatformSSOAuthenticationMethodOpenID               PlatformSSOAuthenticationMethod = "OpenID"
)

// Account display and profile settings.
type Account struct {
	// The display name for the account in notifications and authentication requests.
	DisplayName *string `json:"DisplayName,omitempty"`
	// If `true`, the system requests the user's profile picture from the SSO extension.
	SynchronizeProfilePicture *bool `json:"SynchronizeProfilePicture,omitempty"`
}

// Settings for creating new users via Platform SSO.
type UserCreation struct {
	// Enables creating users at the Login Window with an `AuthenticationMethod` of either `Password` or `SmartCard`. Requires that `UseSharedDeviceKeys` is `true`.
	EnableAtLogin *bool `json:"EnableAtLogin,omitempty"`
	// If `true`, the device uses Platform SSO to create the first user account on the Mac during `Setup Assistant`.
	EnableFirstUserDuringSetup *bool `json:"EnableFirstUserDuringSetup,omitempty"`
	// If `true`, the system enables the PlatformSSO registration process during Setup Assistant on devices running macOS 26 and later. Set this key to `true` when configuring PlatformSSO before enrollment using the `com.apple.psso.required` error response.
	EnableRegistrationDuringSetup *bool `json:"EnableRegistrationDuringSetup,omitempty"`
	// The set of authentication methods to use for newly created accounts at login or during `Setup Assistant`. The system uses `Password` and `SmartCard` if this key isn't present.
	NewUserAuthenticationMethods *[]NewUserAuthenticationMethods `json:"NewUserAuthenticationMethods,omitempty"`
	// The permission to apply to newly created accounts at login. Allowed values:
	// * `Standard`: The account is a standard user.
	// * `Admin`: The system adds the account to the local administrators group.
	// * `Groups`: The system assigns groups to the account using `Authorization.AdministratorGroups`, `Authorization.AdditionalGroups`, or `Authorization.AuthorizationGroups`.
	// * `Temporary`: The system uses a temporary session configuration for newly created accounts at login.
	NewUserAuthorizationMode *NewUserAuthorizationMode `json:"NewUserAuthorizationMode,omitempty"`
	// The attribute mapping to use when creating users, or for authorization.
	TokenToUserMapping *TokenToUserMapping `json:"TokenToUserMapping,omitempty"`
	// If `true`, the system uses a quicker Authenticated Guest Mode login to Mac behavior. The system erases user data from only select locations in the user home directory after each session completes. Once every eight hours the system erases the full user home directory after a session completes. Turn this on for shared environments that have a high frequency of short sessions.
	TemporarySessionQuickLogin *bool `json:"TemporarySessionQuickLogin,omitempty"`
}

// The set of authentication methods to use for newly created accounts at login or during `Setup Assistant`. The system uses `Password` and `SmartCard` if this key isn't present.
// An authentication method to use for newly created accounts at login or during `Setup Assistant`. Allowed values:
// * `Password`: The account uses a password for authentication.
// * `SmartCard`: The account uses a smart card for authentication.
// * `AccessKey`: The account uses an access key for authentication.
// * `OpenID`: The account uses OpenID for authentication.
type NewUserAuthenticationMethods string

const (
	NewUserAuthenticationMethodsPassword  NewUserAuthenticationMethods = "Password"
	NewUserAuthenticationMethodsSmartCard NewUserAuthenticationMethods = "SmartCard"
	NewUserAuthenticationMethodsAccessKey NewUserAuthenticationMethods = "AccessKey"
	NewUserAuthenticationMethodsOpenID    NewUserAuthenticationMethods = "OpenID"
)

// The permission to apply to newly created accounts at login. Allowed values:
// * `Standard`: The account is a standard user.
// * `Admin`: The system adds the account to the local administrators group.
// * `Groups`: The system assigns groups to the account using `Authorization.AdministratorGroups`, `Authorization.AdditionalGroups`, or `Authorization.AuthorizationGroups`.
// * `Temporary`: The system uses a temporary session configuration for newly created accounts at login.
type NewUserAuthorizationMode string

const (
	NewUserAuthorizationModeStandard  NewUserAuthorizationMode = "Standard"
	NewUserAuthorizationModeAdmin     NewUserAuthorizationMode = "Admin"
	NewUserAuthorizationModeGroups    NewUserAuthorizationMode = "Groups"
	NewUserAuthorizationModeTemporary NewUserAuthorizationMode = "Temporary"
)

// The attribute mapping to use when creating users, or for authorization.
type TokenToUserMapping struct {
	// The claim name to use for the user's account name.
	AccountName *string `json:"AccountName,omitempty"`
	// The claim name to use for the user's full name.
	FullName *string `json:"FullName,omitempty"`
}

// Settings for authorization prompts and group management.
type Authorization struct {
	// Enables using identity provider accounts at authorization prompts. Requires that `UseSharedDeviceKeys` is `true`. The system assigns groups using `AdministratorGroups`, `AdditionalGroups`, or `AuthorizationGroups`.
	EnableIdentityProviderAccounts *bool `json:"EnableIdentityProviderAccounts,omitempty"`
	// The permission to apply to an account each time the user authenticates. Allowed values:
	// * `Standard`: The account is a standard user.
	// * `Admin`: The system adds the account to the local administrators group.
	// * `Groups`: The system assigns group to the account using `AdministratorGroups`, `AdditionalGroups`, or `AuthorizationGroups`.
	UserAuthorizationMode *UserAuthorizationMode `json:"UserAuthorizationMode,omitempty"`
	// The list of groups to use for administrator access. The system requests membership during authentication.
	AdministratorGroups *[]string `json:"AdministratorGroups,omitempty"`
	// The list of created groups that don't have administrator access.
	AdditionalGroups *[]string `json:"AdditionalGroups,omitempty"`
	// The pairing of Authorization Rights to group names. When using this, the system updates the Authorization Right to use the group.
	AuthorizationGroups *map[string]string `json:"AuthorizationGroups,omitempty"`
}

// The permission to apply to an account each time the user authenticates. Allowed values:
// * `Standard`: The account is a standard user.
// * `Admin`: The system adds the account to the local administrators group.
// * `Groups`: The system assigns group to the account using `AdministratorGroups`, `AdditionalGroups`, or `AuthorizationGroups`.
type UserAuthorizationMode string

const (
	UserAuthorizationModeStandard UserAuthorizationMode = "Standard"
	UserAuthorizationModeAdmin    UserAuthorizationMode = "Admin"
	UserAuthorizationModeGroups   UserAuthorizationMode = "Groups"
)

// Settings for Access Key authentication.
type AccessKey struct {
	// The reader group identifier for use with the `AccessKey`. The value needs to match the configured access key. Required if `UserCreation.AuthenticationMethods` contains `AccessKey`.
	ReaderGroupIdentifier *[]byte `json:"ReaderGroupIdentifier,omitempty"`
	// The identifier of an asset declaration that contains the identity to use as the Terminal identity of the Access Key. The Access Key needs to trust the identity. Required if `UserCreation.AuthenticationMethods` includes `AccessKey`.
	TerminalIdentityAssetReference *string `json:"TerminalIdentityAssetReference,omitempty"`
	// The identifier of an asset declaration that contains the certificate for the issuer certificate of the `Terminal` identity of the access key. Other specifications refer to the key as the "Reader CA Public Key". The key must be an elliptic curve key. Required if `UserCreation.AuthenticationMethods` includes `AccessKey`. The issuer of the Terminal identity of the access key needs to match this certificate, otherwise the device fails the authentication.
	ReaderIssuerCertificateAssetReference *string `json:"ReaderIssuerCertificateAssetReference,omitempty"`
	// If `true`, the system uses the access key in express mode, and doesn't require authentication before use.
	AllowExpressMode *bool `json:"AllowExpressMode,omitempty"`
}

// Policies for login, unlock, and FileVault behavior.
type Policies struct {
	// The policy to apply when using Platform SSO at FileVault unlock on a Mac with Apple silicon.
	// * `AttemptAuthentication`: The device attempts Platform SSO authentication before proceeding. If offline, unlock continues if the local account password matches. If online and the credential is incorrect, then the device requires a successful Platform SSO authentication is required, even if taken offline. Only use when `AuthenticationMethod` is `Password`.
	// * `RequireAuthentication`: The device requires Platform SSO authentication before proceeding. If the device is offline and `AllowOfflineGracePeriod` is enabled, then the device uses the offline `OfflineGracePeriod` to determine if the user can proceed or not. If online and the credential is incorrect, then the device requires a valid Platform SSO authentication to proceed, regardless of the `OfflineGracePeriod`. If the account isn't registered for Platform SSO and `AllowAuthenticationGracePeriod` is enabled, then the device uses `AuthenticationGracePeriod` to determine if the user can proceed or not. Only use when `AuthenticationMethod` is `Password`.
	// * `AllowOfflineGracePeriod`: The device allows the use of the `OfflineGracePeriod`. If `AllowOfflineGracePeriod` isn't set, then the device denies offline access. Only use when `AuthenticationMethod` is `Password` and `RequireAuthentication` is enabled, or `AuthenticationMethod` is `OpenID`.
	// * `AllowAuthenticationGracePeriod`: The device allows the use of the `AuthenticationGracePeriod` for other local accounts when `RequireAuthentication` is enabled. The `AuthenticationGracePeriod` starts when any of the policies are updated. If `AllowAuthenticationGracePeriod` isn't set, then the device denies unregistered account access. Only use when `AuthenticationMethod` is `Password`.
	// * `RequireTouchID`: The device requires the use of Touch ID (and not Apple Watch) for FileVault unlock. Only use when `AuthenticationMethod` is `Password` or `UserSecureEnclaveKey`.
	// * `RequireTouchIDOrWatch`: The device requires the use of Touch ID or Apple Watch for FileVault unlock. Only use when `AuthenticationMethod` is `Password` or `UserSecureEnclaveKey`.
	// * `AllowOpenIDForTouchIDFallback`: The device allows web login as a fallback if Touch ID fails or isn't available. Only use when `AuthenticationMethod` is `Password` or `UserSecureEnclaveKey`.
	FileVault *[]FileVault `json:"FileVault,omitempty"`
	// The policy to apply when using Platform SSO at the Login Window.
	// * `AttemptAuthentication`: The device attempts Platform SSO authentication before proceeding. If offline, login continues if the local account password matches. If online and the credential is incorrect, then the device requires a successful Platform SSO authentication to proceed, even if taken offline. Only use when `AuthenticationMethod` is `Password`.
	// * `RequireAuthentication`: The device requires Platform SSO authentication before proceeding. If the device is offline and `AllowOfflineGracePeriod` is enabled, then the device uses the offline `OfflineGracePeriod` to determine if the user can proceed or not. If online and the credential is incorrect, then the device requires a valid Platform SSO authentication to proceed, regardless of the `OfflineGracePeriod`. If the account isn't registered for Platform SSO and `AllowAuthenticationGracePeriod` is enabled, then the device uses the `AuthenticationGracePeriod` to determine if the user can proceed or not. Only use when `AuthenticationMethod` is `Password`.
	// * `AllowOfflineGracePeriod`: The device allows the use of the `OfflineGracePeriod`. If `AllowOfflineGracePeriod` isn't set, then the device denies offline access. Only use when `AuthenticationMethod` is `Password` and `RequireAuthentication` is enabled, or `AuthenticationMethod` is `OpenID`.
	// * `AllowAuthenticationGracePeriod`: The device allows the use of the `AuthenticationGracePeriod` for other local accounts when `RequireAuthentication` is enabled. The `AuthenticationGracePeriod` starts when any of the policies have been updated. If `AllowAuthenticationGracePeriod` isn't set, then the device denies unregistered account access. Only use when `AuthenticationMethod` is `Password`.
	// * `RequireTouchID`: The device requires the use of Touch ID (and not Apple Watch) for login. Only use when `AuthenticationMethod` is `Password` or `UserSecureEnclaveKey`.
	// * `RequireTouchIDOrWatch`: The device requires the use of Touch ID or Apple Watch for login. Only use when `AuthenticationMethod` is `Password` or `UserSecureEnclaveKey`.
	// * `AllowOpenIDForTouchIDFallback`: The device allows web login as fallback if Touch ID fails or isn't available. Only use when `AuthenticationMethod` is `Password` or `UserSecureEnclaveKey`.
	Login *[]Login `json:"Login,omitempty"`
	// The policy to apply when using Platform SSO at screensaver unlock.
	// * `AttemptAuthentication`: The device attempts Platform SSO authentication before proceeding. If offline, unlock will continue if the local account password matches. If online and the credential is incorrect, then the device requires a successful Platform SSO authentication to proceed, even if taken offline. Only use when `AuthenticationMethod` is `Password`.
	// * `RequireAuthentication`: The device requires Platform SSO authentication before proceeding. If the device is offline and `AllowOfflineGracePeriod` is enabled, then the offline `OfflineGracePeriod` is used to determine if the user can proceed or not. If online and the credential is incorrect, then the device requires a valid Platform SSO authentication to proceed regardless of the `OfflineGracePeriod`. If the account isn't registered for Platform SSO and `AllowAuthenticationGracePeriod` is enabled, then the device uses `AuthenticationGracePeriod` to determine if the user can proceed or not. Only use when `AuthenticationMethod` is `Password`.
	// * `AllowOfflineGracePeriod`: The device allows the use of the `OfflineGracePeriod`. If `AllowOfflineGracePeriod` isn't set, then the device denies offline access. Only use when `AuthenticationMethod` is `Password` and `RequireAuthentication` is enabled, or `AuthenticationMethod` is `OpenID`.
	// * `AllowAuthenticationGracePeriod`: The device allows the use of the `AuthenticationGracePeriod` for other local accounts when `RequireAuthentication` is enabled. The `AuthenticationGracePeriod` starts when any of the policies have been updated. If `AllowAuthenticationGracePeriod` isn't set, then the device denies the unregistered account access. Only use when `AuthenticationMethod` is `Password`.
	// * `AllowTouchIDOrWatchForUnlock`: The device allows Touch ID or Apple Watch to unlock the screensaver instead of Platform SSO authentication when `RequireAuthentication` is enabled. Only use when `AuthenticationMethod` is `Password`.
	// * `RequireTouchID`: The device requires the use of Touch ID (and not Apple Watch) for unlock. Only use when `AuthenticationMethod` is `Password` or `UserSecureEnclaveKey`.
	// * `RequireTouchIDOrWatch`: The device requires the use of Touch ID or Apple Watch for unlock. Only use when `AuthenticationMethod` is `Password` or `UserSecureEnclaveKey`.
	// * `AllowOpenIDForTouchIDFallback`: The device allows web login as fallback if Touch ID fails or isn't available. Only use when `AuthenticationMethod` is `Password` or `UserSecureEnclaveKey`.
	Unlock *[]Unlock `json:"Unlock,omitempty"`
	// If `true`, the system uses the Platform SSO login UI at the Login Window, screensaver unlock, and FileVault unlock for every configuration that supports it.
	// If `false`, the system uses the Platform SSO login UI only when `AuthenticationMethod` is set to `OpenID`, or when `LoginPolicy`, `UnlockPolicy` or `FileVaultPolicy` are set to `RequireTouchID` or `RequireTouchIDOrWatch`.
	AlwaysUseLoginUI *bool `json:"AlwaysUseLoginUI,omitempty"`
	// The amount of time (in seconds) after the last successful Platform SSO login for using a local account password offline. Required when setting `AllowOfflineGracePeriod`.
	OfflineGracePeriod *int64 `json:"OfflineGracePeriod,omitempty"`
	// The amount of time (in seconds) after receiving or updating a `Policies.FileVault`, `Policies.Login`, or `Policies.Unlock` that the system can use unregistered local accounts. Required when `AllowAuthenticationGracePeriod` is set.
	AuthenticationGracePeriod *int64 `json:"AuthenticationGracePeriod,omitempty"`
	// The list of local accounts that aren't subject to the `Policies.FileVault`, `Policies.Login`, or `Policies.Unlock` policies. The accounts don't receive a prompt to register for Platform SSO.
	NonPlatformSSOAccounts *[]string `json:"NonPlatformSSOAccounts,omitempty"`
}

// The policy to apply when using Platform SSO at FileVault unlock on a Mac with Apple silicon.
// * `AttemptAuthentication`: The device attempts Platform SSO authentication before proceeding. If offline, unlock continues if the local account password matches. If online and the credential is incorrect, then the device requires a successful Platform SSO authentication is required, even if taken offline. Only use when `AuthenticationMethod` is `Password`.
// * `RequireAuthentication`: The device requires Platform SSO authentication before proceeding. If the device is offline and `AllowOfflineGracePeriod` is enabled, then the device uses the offline `OfflineGracePeriod` to determine if the user can proceed or not. If online and the credential is incorrect, then the device requires a valid Platform SSO authentication to proceed, regardless of the `OfflineGracePeriod`. If the account isn't registered for Platform SSO and `AllowAuthenticationGracePeriod` is enabled, then the device uses `AuthenticationGracePeriod` to determine if the user can proceed or not. Only use when `AuthenticationMethod` is `Password`.
// * `AllowOfflineGracePeriod`: The device allows the use of the `OfflineGracePeriod`. If `AllowOfflineGracePeriod` isn't set, then the device denies offline access. Only use when `AuthenticationMethod` is `Password` and `RequireAuthentication` is enabled, or `AuthenticationMethod` is `OpenID`.
// * `AllowAuthenticationGracePeriod`: The device allows the use of the `AuthenticationGracePeriod` for other local accounts when `RequireAuthentication` is enabled. The `AuthenticationGracePeriod` starts when any of the policies are updated. If `AllowAuthenticationGracePeriod` isn't set, then the device denies unregistered account access. Only use when `AuthenticationMethod` is `Password`.
// * `RequireTouchID`: The device requires the use of Touch ID (and not Apple Watch) for FileVault unlock. Only use when `AuthenticationMethod` is `Password` or `UserSecureEnclaveKey`.
// * `RequireTouchIDOrWatch`: The device requires the use of Touch ID or Apple Watch for FileVault unlock. Only use when `AuthenticationMethod` is `Password` or `UserSecureEnclaveKey`.
// * `AllowOpenIDForTouchIDFallback`: The device allows web login as a fallback if Touch ID fails or isn't available. Only use when `AuthenticationMethod` is `Password` or `UserSecureEnclaveKey`.
// The policy to apply.
type FileVault string

const (
	FileVaultAttemptAuthentication          FileVault = "AttemptAuthentication"
	FileVaultRequireAuthentication          FileVault = "RequireAuthentication"
	FileVaultAllowOfflineGracePeriod        FileVault = "AllowOfflineGracePeriod"
	FileVaultAllowAuthenticationGracePeriod FileVault = "AllowAuthenticationGracePeriod"
	FileVaultRequireTouchID                 FileVault = "RequireTouchID"
	FileVaultRequireTouchIDOrWatch          FileVault = "RequireTouchIDOrWatch"
	FileVaultAllowOpenIDForTouchIDFallback  FileVault = "AllowOpenIDForTouchIDFallback"
)

// The policy to apply when using Platform SSO at the Login Window.
// * `AttemptAuthentication`: The device attempts Platform SSO authentication before proceeding. If offline, login continues if the local account password matches. If online and the credential is incorrect, then the device requires a successful Platform SSO authentication to proceed, even if taken offline. Only use when `AuthenticationMethod` is `Password`.
// * `RequireAuthentication`: The device requires Platform SSO authentication before proceeding. If the device is offline and `AllowOfflineGracePeriod` is enabled, then the device uses the offline `OfflineGracePeriod` to determine if the user can proceed or not. If online and the credential is incorrect, then the device requires a valid Platform SSO authentication to proceed, regardless of the `OfflineGracePeriod`. If the account isn't registered for Platform SSO and `AllowAuthenticationGracePeriod` is enabled, then the device uses the `AuthenticationGracePeriod` to determine if the user can proceed or not. Only use when `AuthenticationMethod` is `Password`.
// * `AllowOfflineGracePeriod`: The device allows the use of the `OfflineGracePeriod`. If `AllowOfflineGracePeriod` isn't set, then the device denies offline access. Only use when `AuthenticationMethod` is `Password` and `RequireAuthentication` is enabled, or `AuthenticationMethod` is `OpenID`.
// * `AllowAuthenticationGracePeriod`: The device allows the use of the `AuthenticationGracePeriod` for other local accounts when `RequireAuthentication` is enabled. The `AuthenticationGracePeriod` starts when any of the policies have been updated. If `AllowAuthenticationGracePeriod` isn't set, then the device denies unregistered account access. Only use when `AuthenticationMethod` is `Password`.
// * `RequireTouchID`: The device requires the use of Touch ID (and not Apple Watch) for login. Only use when `AuthenticationMethod` is `Password` or `UserSecureEnclaveKey`.
// * `RequireTouchIDOrWatch`: The device requires the use of Touch ID or Apple Watch for login. Only use when `AuthenticationMethod` is `Password` or `UserSecureEnclaveKey`.
// * `AllowOpenIDForTouchIDFallback`: The device allows web login as fallback if Touch ID fails or isn't available. Only use when `AuthenticationMethod` is `Password` or `UserSecureEnclaveKey`.
// The policy to apply.
type Login string

const (
	LoginAttemptAuthentication          Login = "AttemptAuthentication"
	LoginRequireAuthentication          Login = "RequireAuthentication"
	LoginAllowOfflineGracePeriod        Login = "AllowOfflineGracePeriod"
	LoginAllowAuthenticationGracePeriod Login = "AllowAuthenticationGracePeriod"
	LoginRequireTouchID                 Login = "RequireTouchID"
	LoginRequireTouchIDOrWatch          Login = "RequireTouchIDOrWatch"
	LoginAllowOpenIDForTouchIDFallback  Login = "AllowOpenIDForTouchIDFallback"
)

// The policy to apply when using Platform SSO at screensaver unlock.
// * `AttemptAuthentication`: The device attempts Platform SSO authentication before proceeding. If offline, unlock will continue if the local account password matches. If online and the credential is incorrect, then the device requires a successful Platform SSO authentication to proceed, even if taken offline. Only use when `AuthenticationMethod` is `Password`.
// * `RequireAuthentication`: The device requires Platform SSO authentication before proceeding. If the device is offline and `AllowOfflineGracePeriod` is enabled, then the offline `OfflineGracePeriod` is used to determine if the user can proceed or not. If online and the credential is incorrect, then the device requires a valid Platform SSO authentication to proceed regardless of the `OfflineGracePeriod`. If the account isn't registered for Platform SSO and `AllowAuthenticationGracePeriod` is enabled, then the device uses `AuthenticationGracePeriod` to determine if the user can proceed or not. Only use when `AuthenticationMethod` is `Password`.
// * `AllowOfflineGracePeriod`: The device allows the use of the `OfflineGracePeriod`. If `AllowOfflineGracePeriod` isn't set, then the device denies offline access. Only use when `AuthenticationMethod` is `Password` and `RequireAuthentication` is enabled, or `AuthenticationMethod` is `OpenID`.
// * `AllowAuthenticationGracePeriod`: The device allows the use of the `AuthenticationGracePeriod` for other local accounts when `RequireAuthentication` is enabled. The `AuthenticationGracePeriod` starts when any of the policies have been updated. If `AllowAuthenticationGracePeriod` isn't set, then the device denies the unregistered account access. Only use when `AuthenticationMethod` is `Password`.
// * `AllowTouchIDOrWatchForUnlock`: The device allows Touch ID or Apple Watch to unlock the screensaver instead of Platform SSO authentication when `RequireAuthentication` is enabled. Only use when `AuthenticationMethod` is `Password`.
// * `RequireTouchID`: The device requires the use of Touch ID (and not Apple Watch) for unlock. Only use when `AuthenticationMethod` is `Password` or `UserSecureEnclaveKey`.
// * `RequireTouchIDOrWatch`: The device requires the use of Touch ID or Apple Watch for unlock. Only use when `AuthenticationMethod` is `Password` or `UserSecureEnclaveKey`.
// * `AllowOpenIDForTouchIDFallback`: The device allows web login as fallback if Touch ID fails or isn't available. Only use when `AuthenticationMethod` is `Password` or `UserSecureEnclaveKey`.
// The policy to apply.
type Unlock string

const (
	UnlockAttemptAuthentication          Unlock = "AttemptAuthentication"
	UnlockRequireAuthentication          Unlock = "RequireAuthentication"
	UnlockAllowOfflineGracePeriod        Unlock = "AllowOfflineGracePeriod"
	UnlockAllowAuthenticationGracePeriod Unlock = "AllowAuthenticationGracePeriod"
	UnlockAllowTouchIDOrWatchForUnlock   Unlock = "AllowTouchIDOrWatchForUnlock"
	UnlockRequireTouchID                 Unlock = "RequireTouchID"
	UnlockRequireTouchIDOrWatch          Unlock = "RequireTouchIDOrWatch"
	UnlockAllowOpenIDForTouchIDFallback  Unlock = "AllowOpenIDForTouchIDFallback"
)

// Settings for web authentication behavior.
type WebAuthentication struct {
	// The set of allowed hosts that the system can load in the PSSO web view. Required if `AuthenticationMethod` is set to `OpenID`, or `UserCreation.AuthenticationMethods` contains `OpenID`.
	URLAllowList *[]string `json:"URLAllowList,omitempty"`
	// If `true`, the system detects the password during web authentication and synchronizes it to the local account password for the user.
	AllowPasswordSync *bool `json:"AllowPasswordSync,omitempty"`
}

// The declaration to configure External Intelligence Integrations settings.
type ExternalIntelligenceSettings struct {
	// If `false`, disables external intelligence integrations.
	Enabled *bool `json:"Enabled,omitempty"`
	// If `false`, disables sign-in for external intelligence integrations.
	AllowSignIn *bool `json:"AllowSignIn,omitempty"`
	// An array of strings, but currently restricted to a single element. If present, Apple Intelligence allows use of only the given external integration workspace ID, and requires a sign-in to make requests. The user is required to sign in to integrations that support signing in. Multiple values combine using an intersect operation.
	AllowedWorkspaceIDs *[]string `json:"AllowedWorkspaceIDs,omitempty"`
}

func (p *ExternalIntelligenceSettings) DeclarationType() string {
	return "com.apple.configuration.external-intelligence.settings"
}

// The declaration to configure Apple Intelligence settings.
type IntelligenceSettings struct {
	// If `false`, disables Apple Intelligence Report.
	AllowAppleIntelligenceReport *bool `json:"AllowAppleIntelligenceReport,omitempty"`
	// If `false`, disables Genmoji.
	AllowGenmoji *bool `json:"AllowGenmoji,omitempty"`
	// If `false`, disables Image Playground.
	AllowImagePlayground *bool `json:"AllowImagePlayground,omitempty"`
	// If `false`, disables Image Wand.
	AllowImageWand *bool `json:"AllowImageWand,omitempty"`
	// If `false`, disables Personalized Handwriting Results.
	AllowPersonalizedHandwritingResults *bool `json:"AllowPersonalizedHandwritingResults,omitempty"`
	// If `false`, disables Visual Intelligence.
	AllowVisualIntelligence *bool `json:"AllowVisualIntelligence,omitempty"`
	// If `false`, disables Visual Intelligence Summary.
	// Deprecated: use the `AllowVisualIntelligence` key.
	AllowVisualIntelligenceSummary *bool `json:"AllowVisualIntelligenceSummary,omitempty"`
	// If `false`, disables Writing Tools.
	AllowWritingTools *bool `json:"AllowWritingTools,omitempty"`
	// If present, configures app-specific Intelligence features.
	Apps *Apps `json:"Apps,omitempty"`
	// If `true`, forces On-Device Only Dictation.
	ForceOnDeviceOnlyDictation *bool `json:"ForceOnDeviceOnlyDictation,omitempty"`
	// If `true`, forces On-Device Only Translation.
	ForceOnDeviceOnlyTranslation *bool `json:"ForceOnDeviceOnlyTranslation,omitempty"`
}

func (p *IntelligenceSettings) DeclarationType() string {
	return "com.apple.configuration.intelligence.settings"
}

// If present, configures app-specific Intelligence features.
type Apps struct {
	// If present, configures Calendar and Reminders Intelligence features.
	Calendar *Calendar `json:"Calendar,omitempty"`
	// If present, configures Mail Intelligence features.
	Mail *Mail `json:"Mail,omitempty"`
	// If present, configures Notes Intelligence features.
	Notes *Notes `json:"Notes,omitempty"`
	// If present, configures Safari intelligence features.
	Safari *Safari `json:"Safari,omitempty"`
}

// If present, configures Calendar and Reminders Intelligence features.
type Calendar struct {
	// If `false`, disables Natural Language Editing in Calendar and Reminders.
	AllowNaturalLanguageEditing *bool `json:"AllowNaturalLanguageEditing,omitempty"`
}

// If present, configures Mail Intelligence features.
type Mail struct {
	// If `false`, disables Mail Smart Replies.
	AllowSmartReplies *bool `json:"AllowSmartReplies,omitempty"`
	// If `false`, disables Mail Summary.
	AllowSummary *bool `json:"AllowSummary,omitempty"`
}

// If present, configures Notes Intelligence features.
type Notes struct {
	// If `false`, disables Notes Transcription.
	AllowTranscription *bool `json:"AllowTranscription,omitempty"`
	// If `false`, disables Notes Transcription Summary.
	AllowTranscriptionSummary *bool `json:"AllowTranscriptionSummary,omitempty"`
}

// If present, configures Safari intelligence features.
type Safari struct {
	// If `false`, disables Safari Summary.
	AllowSummary *bool `json:"AllowSummary,omitempty"`
}

// The declaration to configure keyboard settings.
type KeyboardSettings struct {
	// If `false`, disables auto-correction.
	AllowAutoCorrection *bool `json:"AllowAutoCorrection,omitempty"`
	// If `false`, disables definition lookup.
	AllowDefinitionLookup *bool `json:"AllowDefinitionLookup,omitempty"`
	// If `false`, disables dictation.
	AllowDictation *bool `json:"AllowDictation,omitempty"`
	// If `false`, disables keyboard suggestions that include math solutions. This key is also supported by the math.settings configuration.
	AllowMathKeyboardSuggestions *bool `json:"AllowMathKeyboardSuggestions,omitempty"`
	// If `false`, disables predictive text.
	AllowPredictiveText *bool `json:"AllowPredictiveText,omitempty"`
	// If `false`, disables slide to type.
	AllowSlideToType *bool `json:"AllowSlideToType,omitempty"`
	// If `false`, disables spell check.
	AllowSpellCheck *bool `json:"AllowSpellCheck,omitempty"`
	// If `false`, disables text replacement.
	AllowTextReplacement *bool `json:"AllowTextReplacement,omitempty"`
}

func (p *KeyboardSettings) DeclarationType() string {
	return "com.apple.configuration.keyboard.settings"
}

// The declaration to configure an interactive legacy profile.
type LegacyInteractiveProfile struct {
	// The URL of the profile to download and install, which needs to start with `https://`. The request uses MDM semantics, which includes the device-identity certificate, and any user authentication. This is equivalent to an MDM request made to the `CheckInURL` or `ServerURL`.
	// One of `ProfileURL` or `ProfileAssetReference` needs to be present.
	ProfileURL *string `json:"ProfileURL,omitempty"`
	// The identifier of an asset declaration containing a reference to the profile data. The corresponding asset needs to be of type `com.apple.asset.data`. The referenced data needs to be a property list file, and the asset's "ContentType" value set to match the data type.
	// One of `ProfileURL` or `ProfileAssetReference` needs to be present.
	ProfileAssetReference *string `json:"ProfileAssetReference,omitempty"`
	// The visible name of the configuration. This name needs to indicate the nature of the profile.
	VisibleName string `json:"VisibleName" required:"true"`
}

func (p *LegacyInteractiveProfile) DeclarationType() string {
	return "com.apple.configuration.legacy.interactive"
}

// The declaration to configure a legacy profile.
type LegacyProfile struct {
	// The URL of the profile to download and install, which needs to start with `https://`. The request uses MDM semantics, which includes the device-identity certificate, and any user authentication. This is equivalent to an MDM request made to the `CheckInURL` or `ServerURL`.
	// One of `ProfileURL` or `ProfileAssetReference` needs to be present.
	ProfileURL *string `json:"ProfileURL,omitempty"`
	// The identifier of an asset declaration containing a reference to the profile data. The corresponding asset needs to be of type `com.apple.asset.data`. The referenced data needs to be a property list file, and the asset's "ContentType" value set to match the data type.
	// One of `ProfileURL` or `ProfileAssetReference` needs to be present.
	ProfileAssetReference *string `json:"ProfileAssetReference,omitempty"`
}

func (p *LegacyProfile) DeclarationType() string {
	return "com.apple.configuration.legacy"
}

// The declaration to configure status subscriptions.
type ManagementStatusSubscriptions struct {
	// An array of status items that the device notifies subscribers about.
	StatusItems []*StatusItem `json:"StatusItems" required:"true"`
}

func (p *ManagementStatusSubscriptions) DeclarationType() string {
	return "com.apple.configuration.management.status-subscriptions"
}

// The declaration for configuring a specific status subscription.
type StatusItem struct {
	// The name of the status item to send to subscribers.
	Name string `json:"Name" required:"true"`
}

// The declaration to configure a declarative device management test.
type ManagementTest struct {
	// The string to echo back in a status response reason.
	Echo string `json:"Echo" required:"true"`
	// The string to read from a data asset to echo back in status response reason description.
	EchoDataAssetReference *string `json:"EchoDataAssetReference,omitempty"`
	// The status the system reports back when the device implements the configuration. Use this to override the normal `success` result.
	ReturnStatus *ReturnStatus `default:"Installed" json:"ReturnStatus,omitempty"`
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
type MathSettings struct {
	// If present, configures the built-in Calculator app.
	Calculator *Calculator `json:"Calculator,omitempty"`
	// If present, configures math behavior in the system.
	SystemBehavior *SystemBehavior `json:"SystemBehavior,omitempty"`
}

func (p *MathSettings) DeclarationType() string {
	return "com.apple.configuration.math.settings"
}

// If present, configures the built-in Calculator app.
type Calculator struct {
	// If present, configures the basic mode of the calculator. Basic mode is always enabled.
	BasicMode *BasicMode `json:"BasicMode,omitempty"`
	// If present, configures the scientific mode of the calculator. If not present, scientific mode is enabled.
	ScientificMode *ScientificMode `json:"ScientificMode,omitempty"`
	// If present, configures the programmer mode of the calculator. If not present, programmer mode is enabled.
	ProgrammerMode *ProgrammerMode `json:"ProgrammerMode,omitempty"`
	// If present, configures the Math Notes mode of the calculator. If not present, Math Notes mode is enabled.
	MathNotesMode *MathNotesMode `json:"MathNotesMode,omitempty"`
	// If present, controls global input options of the calculator. If not present, all input modes are enabled.
	InputModes *InputModes `json:"InputModes,omitempty"`
}

// If present, configures the basic mode of the calculator. Basic mode is always enabled.
type BasicMode struct {
	// Add the square root button to the basic calculator by replacing the +/- button. Normally, the square root button is available in scientific mode, so this key can be used to make it available when the scientific mode is restricted.
	AddSquareRoot bool `json:"AddSquareRoot" required:"true"`
}

// If present, configures the scientific mode of the calculator. If not present, scientific mode is enabled.
type ScientificMode struct {
	// Controls whether the mode is enabled.
	Enabled bool `json:"Enabled" required:"true"`
}

// If present, configures the programmer mode of the calculator. If not present, programmer mode is enabled.
type ProgrammerMode struct {
	// Controls whether the mode is enabled.
	Enabled bool `json:"Enabled" required:"true"`
}

// If present, configures the Math Notes mode of the calculator. If not present, Math Notes mode is enabled.
type MathNotesMode struct {
	// Controls whether the mode is enabled.
	Enabled bool `json:"Enabled" required:"true"`
}

// If present, controls global input options of the calculator. If not present, all input modes are enabled.
type InputModes struct {
	// Configures whether unit conversions are enabled.
	UnitConversion bool `json:"UnitConversion" required:"true"`
	// Configures whether RPN input is enabled.
	RPN bool `json:"RPN" required:"true"`
}

// If present, configures math behavior in the system.
type SystemBehavior struct {
	// Controls whether keyboard suggestions include math solutions. This key is also supported by the keyboard.settings configuration.
	KeyboardSuggestions bool `json:"KeyboardSuggestions" required:"true"`
	// Controls whether Math Notes is allowed in other apps such as Notes.
	MathNotes bool `json:"MathNotes" required:"true"`
}

// The declaration to configure Migration Assistant settings.
type MigrationAssistantSettings struct {
	// If `true`, the device manages Migration Assistant.
	ShouldDoManagedMigration bool `json:"ShouldDoManagedMigration" required:"true"`
	// An array of strings that represent the user account short names the system excludes from migration.
	ExcludedAccounts *[]string `json:"ExcludedAccounts,omitempty"`
	// An array of strings that represent files and directories relative to the user's home directory that the system excludes from migration. Directory paths need to include a trailing "/". For example, to exclude the "Excluded" directory in the "Documents" folder of a user's home directory, use "Documents/Excluded/".
	ExcludedPaths *[]string `json:"ExcludedPaths,omitempty"`
	// An array of strings that represent files and directories relative to the user's home directory that the system needs to migrate. Directory paths need to include a trailing "/". For example, to require the "Required" directory in the "Documents" folder of a user's home directory, use "Documents/Required/".
	RequiredPaths *[]string `json:"RequiredPaths,omitempty"`
	// If `true`, the system migrates Security & Privacy settings.
	ShouldMigrateSecurityPrivacySettings bool `json:"ShouldMigrateSecurityPrivacySettings" required:"true"`
}

func (p *MigrationAssistantSettings) DeclarationType() string {
	return "com.apple.configuration.migration-assistant.settings"
}

// The declaration to configure DNS proxy settings.
type NetworkDNSProxy struct {
	// The name of the DNS proxy configuration that the system displays on the device.
	VisibleName string `json:"VisibleName" required:"true"`
	// The bundle identifier of the app containing the DNS proxy network extension.
	AppBundleIdentifier string `json:"AppBundleIdentifier" required:"true"`
	// The identifier of the provider to use for this configuration. Useful for apps that contain more than one DNS proxy extension.
	// In iOS and visionOS, the identifier is a bundle ID, for example, "com.example.app".
	// In macOS, the identifier is a composed identifier. The format of the composed identifier is either "Bundle-ID" or "Bundle-ID {Designated-Requirement}". "Bundle-ID" is the bundle identifier string of the provider. "Designated-Requirement" is the designated requirement string the device uses to match the code signature of the provider. For example, "com.example.app" for the bundle ID format, or "com.example.app {anchor apple generic}" for the designated requirement format.
	ProviderComposedIdentifier *string `json:"ProviderComposedIdentifier,omitempty"`
	// The dictionary of vendor-specific configuration items.
	ProviderConfiguration *map[string]any `json:"ProviderConfiguration,omitempty"`
	// A globally unique identifier for this DNS proxy configuration. The proxy processes DNS lookups traffic for managed apps with the same `DNSProxyUUID` in their app attributes. This key is required for user enrollment.
	DNSProxyUUID *string `json:"DNSProxyUUID,omitempty"`
}

func (p *NetworkDNSProxy) DeclarationType() string {
	return "com.apple.configuration.network.dns-proxy"
}

// The declaration to configure encrypted DNS settings.
type NetworkDNSSettings struct {
	// The name of the DNS settings that the system displays on the device.
	VisibleName string `json:"VisibleName" required:"true"`
	// A dictionary that defines a configuration for an encrypted DNS server.
	DNSSettings DNSSettings `json:"DNSSettings" required:"true"`
	// An array of rules that define the DNS settings. If not set, the system always applies the DNS settings. These rules are identical to the `OnDemandRules` array in VPN payloads.
	OnDemandRules *[]*OnDemandRulesElement `json:"OnDemandRules,omitempty"`
	// If `true`, the system prohibits users from disabling DNS settings.
	ProhibitDisablement *bool `json:"ProhibitDisablement,omitempty"`
}

func (p *NetworkDNSSettings) DeclarationType() string {
	return "com.apple.configuration.network.dns-settings"
}

// A dictionary that defines a configuration for an encrypted DNS server.
type DNSSettings struct {
	// The encrypted transport protocol used to communicate with the DNS server.
	DNSProtocol DNSProtocol `json:"DNSProtocol" required:"true"`
	// The URI template of a DNS-over-HTTPS server, as defined in RFC 8484. This URL needs to use the `https://` scheme, and the system uses the hostname or address in the URL to validate the server certificate. If no `ServerAddresses` are provided, the system uses the hostname or address in the URL to determine the server addresses. Required if `DNSProtocol` is `HTTPS`.
	ServerURL *string `json:"ServerURL,omitempty"`
	// The hostname of a DNS-over-TLS server used to validate the server certificate, as defined in RFC 7858. If no `ServerAddresses` are provided, the system uses the hostname to determine the server addresses. This key must be present only if the DNSProtocol is `TLS`.
	ServerName *string `json:"ServerName,omitempty"`
	// An unordered list of DNS server IP address strings. These IP addresses can be a mixture of IPv4 and IPv6 addresses.
	ServerAddresses *[]string `json:"ServerAddresses,omitempty"`
	// If `true`, the device allows failover to the default system DNS resolver.
	AllowFailover *bool `json:"AllowFailover,omitempty"`
	// Specifies the identifier of an asset declaration containing an identity that the system uses to authenticate the user to the DNS resolver.
	IdentityAssetReference *string `json:"IdentityAssetReference,omitempty"`
	// A list of domain strings used to determine which DNS queries use the DNS server. If not set, all domains use the DNS server.
	// The system supports a single wildcard (`\*`) prefix, but it's not required. For example, both `\*.example.com` and `example.com` match against `mydomain.example.com` and `your.domain.example.com`, but don't match against `mydomain-example.com`.
	SupplementalMatchDomains *[]string `json:"SupplementalMatchDomains,omitempty"`
}

// The encrypted transport protocol used to communicate with the DNS server.
type DNSProtocol string

const (
	DNSProtocolHTTPS DNSProtocol = "HTTPS"
	DNSProtocolTLS   DNSProtocol = "TLS"
)

type OnDemandRulesElement struct {
	// The action to take if this dictionary matches the current network. Allowed values:
	// - `Connect`: Apply DNS Settings when the dictionary matches.
	// - `Disconnect`: Don't apply DNS Settings when the dictionary matches.
	// - `EvaluateConnection`: Apply DNS Settings with per-domain exceptions when the dictionary matches.
	Action OnDemandRulesElementAction `json:"Action" required:"true"`
	// An array of dictionaries that provide per-connection rules. The system uses this array only for settings where the `Action` value is `EvaluateConnection`.
	ActionParameters *[]*OnDemandRulesElementActionParametersActionParameter `json:"ActionParameters,omitempty"`
	// An array of domain names. This rule matches if any of the domain names in the specified list matches any domain in the device's search domains list.
	// The system supports a single wildcard (`\*`) prefix, but it's not required. For example, both `\*.example.com` and `example.com` match against `mydomain.example.com` and `your.domain.example.com`, but don't match against `mydomain-example.com`.
	DNSDomainMatch *[]string `json:"DNSDomainMatch,omitempty"`
	// An array of IP addresses. This rule matches if any of the network's specified DNS servers match any entry in the array.
	// The system supports matching with a single wildcard. For example, `17.\*` matches any DNS server in the 17.0.0.0/8 subnet.
	DNSServerAddressMatch *[]string `json:"DNSServerAddressMatch,omitempty"`
	// An interface type. If specified, this rule matches only if the primary network interface hardware matches the specified type.
	InterfaceTypeMatch *OnDemandRulesElementInterfaceTypeMatch `json:"InterfaceTypeMatch,omitempty"`
	// An array of SSIDs to match against the current network. If the network isn't a Wi-Fi network or if the SSID doesn't appear in this array, the match fails. Omit this key and the corresponding array to match against any SSID.
	SSIDMatch *[]string `json:"SSIDMatch,omitempty"`
	// A URL to probe. This rule matches if this URL is successfully fetched and returns a 200 HTTP status code without redirection.
	URLStringProbe *string `json:"URLStringProbe,omitempty"`
}

// The action to take if this dictionary matches the current network. Allowed values:
// - `Connect`: Apply DNS Settings when the dictionary matches.
// - `Disconnect`: Don't apply DNS Settings when the dictionary matches.
// - `EvaluateConnection`: Apply DNS Settings with per-domain exceptions when the dictionary matches.
type OnDemandRulesElementAction string

const (
	OnDemandRulesElementActionConnect            OnDemandRulesElementAction = "Connect"
	OnDemandRulesElementActionDisconnect         OnDemandRulesElementAction = "Disconnect"
	OnDemandRulesElementActionEvaluateConnection OnDemandRulesElementAction = "EvaluateConnection"
)

// A dictionary that provides per-connection rules.
// The keys allowed in each dictionary are described below. Note: This array is only for dictionaries in which `EvaluateConnection` is the `Action` value.
type OnDemandRulesElementActionParametersActionParameter struct {
	// The domains for which this evaluation applies.
	Domains []string `json:"Domains" required:"true"`
	// The DNS settings behavior for the specified domains. Allowed values:
	// * 'NeverConnect': Don't use the DNS Settings for the specified domains.
	// * 'ConnectIfNeeded': Allow using the DNS Settings for the specified domains.
	DomainAction OnDemandRulesElementActionParametersActionParameterDomainAction `json:"DomainAction" required:"true"`
}

// The DNS settings behavior for the specified domains. Allowed values:
// * 'NeverConnect': Don't use the DNS Settings for the specified domains.
// * 'ConnectIfNeeded': Allow using the DNS Settings for the specified domains.
type OnDemandRulesElementActionParametersActionParameterDomainAction string

const (
	OnDemandRulesElementActionParametersActionParameterDomainActionNeverConnect    OnDemandRulesElementActionParametersActionParameterDomainAction = "NeverConnect"
	OnDemandRulesElementActionParametersActionParameterDomainActionConnectIfNeeded OnDemandRulesElementActionParametersActionParameterDomainAction = "ConnectIfNeeded"
)

// An interface type. If specified, this rule matches only if the primary network interface hardware matches the specified type.
type OnDemandRulesElementInterfaceTypeMatch string

const (
	OnDemandRulesElementInterfaceTypeMatchEthernet OnDemandRulesElementInterfaceTypeMatch = "Ethernet"
	OnDemandRulesElementInterfaceTypeMatchWiFi     OnDemandRulesElementInterfaceTypeMatch = "WiFi"
	OnDemandRulesElementInterfaceTypeMatchCellular OnDemandRulesElementInterfaceTypeMatch = "Cellular"
)

// The declaration to configure Network Relay settings.
type NetworkRelay struct {
	// The name of the network relays that the system displays on the device.
	VisibleName string `json:"VisibleName" required:"true"`
	// An array of dictionaries that describe one or more relay servers that the system can chain together.
	Relays []*Relay `json:"Relays" required:"true"`
	// A list of domain strings that the system uses to determine which connection to route through the servers in `Relays`.
	// Any connection that matches a domain in the list exactly or is a subdomain of the listed domain uses the relay servers, unless it matches a domain in `ExcludedDomains`.
	// If this list and `MatchFQDNs` are empty, the system routes traffic to all domains to the relay servers, except those that match an excluded domain or excluded FQDN.
	MatchDomains *[]string `json:"MatchDomains,omitempty"`
	// A list of domain strings to exclude from routing through the servers in `Relays`. Any connection that matches a domain in the list exactly or is a subdomain of the listed domain won't use the relay server.
	ExcludedDomains *[]string `json:"ExcludedDomains,omitempty"`
	// A list of Fully Qualified Domain Names (FQDNs) to route through the servers contained in `Relays`. Any connection that matches an FQDN in the list exactly uses the relay servers. If this list and `MatchDomains` are empty, the system routes traffic to all domains to the relay servers, except those that match an excluded domain or excluded FQDN.
	MatchFQDNs *[]string `json:"MatchFQDNs,omitempty"`
	// A list of Fully Qualified Domain Names (FQDNs) to exclude from routing through the servers contained in `Relays`. Any connection that matches an FQDN in the list exactly won't use the relay server. When `MatchDomains` is also present, any FQDN listed in the list should be a subdomain of at least one `MatchDomain` value, otherwise it won't have any effect.
	ExcludedFQDNs *[]string `json:"ExcludedFQDNs,omitempty"`
	// A globally unique identifier for this relay configuration. The system uses this UUID to route managed apps through the servers in `Relays`. This key is required for user enrollment.
	RelayUUID *string `json:"RelayUUID,omitempty"`
	// If `true`, the device allows the user to disable this network relay configuration.
	UIToggleEnabled *bool `json:"UIToggleEnabled,omitempty"`
	// If `true`, the device allows the relay to failover to the default system DNS resolver.
	AllowDNSFailover *bool `json:"AllowDNSFailover,omitempty"`
}

func (p *NetworkRelay) DeclarationType() string {
	return "com.apple.configuration.network.relay"
}

type Relay struct {
	// The URL or URI template, as defined in RFC 9298, of a relay server that's reachable using HTTP/3 and supports proxying TCP and UDP using the CONNECT method.
	// Each relay needs to include either `HTTP2RelayURL` or `HTTP3RelayURL`, or it can include both.
	HTTP3RelayURL *string `json:"HTTP3RelayURL,omitempty"`
	// The URL or URI template, as defined in RFC 9298, of a relay server that's reachable using HTTP/2 and supports proxying TCP and UDP using the CONNECT method.
	// Each relay needs to include either `HTTP2RelayURL` or `HTTP3RelayURL`, or it can include both.
	HTTP2RelayURL *string `json:"HTTP2RelayURL,omitempty"`
	// A dictionary that contains custom HTTP header keys and values to add to each request. The dictionary key name represents the HTTP header field name to use, and the dictionary value is the string to use as the HTTP header field value.
	AdditionalHTTPHeaderFields *map[string]string `json:"AdditionalHTTPHeaderFields,omitempty"`
	// The identifier of an asset declaration that contains the identity to install.
	IdentityAssetReference *string `json:"IdentityAssetReference,omitempty"`
	// An array of references to data assets containing DER-encoded public key data that the system uses to authenticate the server during a TLS handshake. The server needs to use one of the keys in the handshake to authenticate.
	// If this array is empty, the system uses the default TLS trust evaluation.
	PublicKeyData *[]string `json:"PublicKeyData,omitempty"`
}

// The declaration to configure a VPN using the Always On sub-type.
type NetworkVPNAlwaysOn struct {
	// The name of the VPN connection that the system displays on the device.
	VisibleName string `json:"VisibleName" required:"true"`
	// If `true`, allows the user to disable the VPN configuration.
	UIToggleEnabled *bool `json:"UIToggleEnabled,omitempty"`
	// An array that contains an arbitrary number of tunnel configurations.
	TunnelConfigurations []*TunnelConfigurationElement `json:"TunnelConfigurations" required:"true"`
	// An array that contains an arbitrary number of service exceptions.
	ServiceExceptions *[]*ServiceExceptionElement `json:"ServiceExceptions,omitempty"`
	// An array that contains an arbitrary number of apps whose connections occur outside the VPN.
	ApplicationExceptions *[]*ApplicationExceptionElement `json:"ApplicationExceptions,omitempty"`
	// If `true`, allows traffic from Captive Web Sheet outside the VPN tunnel.
	AllowCaptiveWebSheet *bool `json:"AllowCaptiveWebSheet,omitempty"`
	// If `true`, allows traffic from all captive networking apps outside the VPN tunnel to perform captive network handling.
	AllowAllCaptiveNetworkPlugins *bool `json:"AllowAllCaptiveNetworkPlugins,omitempty"`
	// The array of captive networking apps whose traffic is allowed outside the VPN tunnel, to perform captive network handling. Used only when `AllowAllCaptiveNetworkPlugins` is `false`.
	AllowedCaptiveNetworkPlugins *[]*AllowedCaptiveNetworkPluginElement `json:"AllowedCaptiveNetworkPlugins,omitempty"`
}

func (p *NetworkVPNAlwaysOn) DeclarationType() string {
	return "com.apple.configuration.network.vpn.always-on"
}

type TunnelConfigurationElement struct {
	// The type of connection, which needs to be `IKEv2`.
	ProtocolType ProtocolType `json:"ProtocolType" required:"true"`
	// The interfaces to apply this configuration to.
	Interfaces *[]Interfaces `json:"Interfaces,omitempty"`
	// The IKEv2 configuration for this tunnel.
	IKEv2 *IKEv2 `json:"IKEV2,omitempty"`
}

// The type of connection, which needs to be `IKEv2`.
type ProtocolType string

const (
	ProtocolTypeIKEv2 ProtocolType = "IKEv2"
)

// The interfaces to apply this configuration to.
type Interfaces string

const (
	InterfacesCellular Interfaces = "Cellular"
	InterfacesWiFi     Interfaces = "WiFi"
)

// The IKEv2 configuration for this tunnel.
type IKEv2 struct {
	// The IP address or hostname of the VPN server.
	HostName string `json:"HostName" required:"true"`
	// Identifier of the IKEv2 client.
	LocalIdentifier string `json:"LocalIdentifier" required:"true"`
	// The remote identifier.
	RemoteIdentifier string `json:"RemoteIdentifier" required:"true"`
	// Settings that control authentication.
	Authentication IKEv2Authentication `json:"Authentication" required:"true"`
	// Specifies details about how the system handles idle VPN connections.
	Idle *IKEv2Idle `json:"Idle,omitempty"`
	// If `true`, negotiations should use IKEv2 Configuration Attribute `INTERNAL_IP4_SUBNET` and `INTERNAL_IP6_SUBNET`.
	UseConfigurationAttributeInternalIPSubnet *bool `json:"UseConfigurationAttributeInternalIPSubnet,omitempty"`
	// If `true`, the system disables MOBIKE.
	DisableMOBIKE *bool `json:"DisableMOBIKE,omitempty"`
	// If `true`, the system disables IKEv2 redirect. If not set, the system redirects an IKEv2 connection when it receives a redirect request from the server.
	DisableRedirect *bool `json:"DisableRedirect,omitempty"`
	// If `true`, enables NAT keepalive offload for Always On VPN IKEv2 connections. The device sends keepalive packets to maintain NAT mappings for IKEv2 connections that have a NAT on the path. It sends keepalive packets at regular intervals when the device is awake. If `NATKeepAliveOffloadEnable` is `true`, the system offloads keepalive packets to hardware while the device is asleep.
	// NAT keepalive offload has an impact on the battery life due to the extra workload during sleep. The default interval for the keepalive offload packets is 20 seconds over Wi-Fi and 110 seconds over Cellular interface. The default NAT keepalive works well on networks with small NAT mapping timeouts but imposes a potential battery impact. If a network has larger NAT mapping timeouts, larger keepalive intervals may be safely used to minimize battery impact. Modify the keepalive interval through the `NATKeepAliveInterval` key.
	EnableNATKeepAliveOffload *bool `json:"EnableNATKeepAliveOffload,omitempty"`
	// The NAT Keepalive interval for Always On VPN IKEv2 connections. This value controls the interval that the device sends keepalive offload packets. The minimum value is 20 seconds. If no key is specified, the default is 20 seconds over Wi-Fi and 110 seconds over a cellular interface.
	NATKeepAliveInterval *int64 `default:"20" json:"NATKeepAliveInterval,omitempty"`
	// If `true`,  enables Perfect Forward Secrecy (PFS) for IKEv2 Connections.
	EnablePFS *bool `json:"EnablePFS,omitempty"`
	// If `true`, the system performs a certificate revocation check for IKEv2 connections. This is a best-effort revocation check and server response timeouts won't cause it to fail.
	EnableCertificateRevocationCheck *bool `json:"EnableCertificateRevocationCheck,omitempty"`
	// The Maximum Transmission Unit (MTU) specifies the maximum size in bytes of each packet that the system sends over the IKEv2 VPN interface.
	MTU *int64 `default:"1280" json:"MTU,omitempty"`
	// If set to `true`, the device doesn't allow DES, 3DES, and Diffie-Hellman groups less than 14. Also the device requires the encryption algorithm specified in `IKESecurityAssociationParameters` to be at least as cryptographically strong as the algorithm specified in `ChildSecurityAssociationParameters`. The device rejects this configuration if these requirements aren't met.
	EnforceStrictAlgorithmSelection *bool `json:"EnforceStrictAlgorithmSelection,omitempty"`
	// Post Quantum Key Exchange settings.
	PostQuantumKeyExchange *IKEv2PostQuantumKeyExchange `json:"PostQuantumKeyExchange,omitempty"`
	// These parameters apply to Child Security Association unless `ChildSecurityAssociationParameters` is specified.
	IKESecurityAssociationParameters *IKEv2IKESecurityAssociationParameters `json:"IKESecurityAssociationParameters,omitempty"`
	// The `ChildSecurityAssociationParameters` dictionaries.
	ChildSecurityAssociationParameters *IKEv2ChildSecurityAssociationParameters `json:"ChildSecurityAssociationParameters,omitempty"`
}

// Settings that control authentication.
type IKEv2Authentication struct {
	// The type of authentication method for the VPN.
	// To enable EAP-only authentication, set this to `None` and `ExtendedAuthEnabled` to `true`. If this is `None` and the `ExtendedAuthEnabled` key isn't set, the authentication configuration defaults to `SharedSecret`.
	Method IKEv2AuthenticationMethod `json:"Method" required:"true"`
	// The identifier of an asset declaration that contains the credentials (password) to authenticate with the VPN server. Required when `Authentication.Method` is set to `SharedSecret`.
	CredentialsAssetReference *string `json:"CredentialsAssetReference,omitempty"`
	// The type of key used by the identity set in the `IdentityAssetReference` to use for IKEv2 machine authentication. If this key is included, the system requires a value for `ServerCertificateIssuerCommonName`.
	IdentityCertificateType *IKEv2AuthenticationIdentityCertificateType `default:"RSA" json:"IdentityCertificateType,omitempty"`
	// The identifier of a credential asset declaration that contains the identity that this account requires to authenticate with the VPN server. If the value of `AuthenticationMethod` is `Certificate`, the system sends this certificate out for IKEv2 machine authentication. If extended authentication (EAP) is used, the system sends this certificate out for EAP-TLS authentication. Required when `Authentication.Method` is set to `Certificate`.
	IdentityAssetReference *string `json:"IdentityAssetReference,omitempty"`
	// Specifies details about extended authentication (EAP).
	ExtendedAuth *IKEv2AuthenticationExtendedAuth `json:"ExtendedAuth,omitempty"`
}

// The type of authentication method for the VPN.
// To enable EAP-only authentication, set this to `None` and `ExtendedAuthEnabled` to `true`. If this is `None` and the `ExtendedAuthEnabled` key isn't set, the authentication configuration defaults to `SharedSecret`.
type IKEv2AuthenticationMethod string

const (
	IKEv2AuthenticationMethodNone         IKEv2AuthenticationMethod = "None"
	IKEv2AuthenticationMethodSharedSecret IKEv2AuthenticationMethod = "SharedSecret"
	IKEv2AuthenticationMethodCertificate  IKEv2AuthenticationMethod = "Certificate"
)

// The type of key used by the identity set in the `IdentityAssetReference` to use for IKEv2 machine authentication. If this key is included, the system requires a value for `ServerCertificateIssuerCommonName`.
type IKEv2AuthenticationIdentityCertificateType string

const (
	IKEv2AuthenticationIdentityCertificateTypeRSA      IKEv2AuthenticationIdentityCertificateType = "RSA"
	IKEv2AuthenticationIdentityCertificateTypeECDSA256 IKEv2AuthenticationIdentityCertificateType = "ECDSA256"
	IKEv2AuthenticationIdentityCertificateTypeECDSA384 IKEv2AuthenticationIdentityCertificateType = "ECDSA384"
	IKEv2AuthenticationIdentityCertificateTypeECDSA521 IKEv2AuthenticationIdentityCertificateType = "ECDSA521"
	IKEv2AuthenticationIdentityCertificateTypeRSAPSS   IKEv2AuthenticationIdentityCertificateType = "RSA-PSS"
)

// Specifies details about extended authentication (EAP).
type IKEv2AuthenticationExtendedAuth struct {
	// If `true`, enables EAP-only authentication.
	Enabled *bool `json:"Enabled,omitempty"`
	// The identifier of an asset declaration that contains the credentials (user name and password) to authenticate with the VPN server. Required when `Enabled` is set to `true`. Implies the use of EAP-MSCHAPv2.
	CredentialsAssetReference *string `json:"CredentialsAssetReference,omitempty"`
	// Common Name of the server certificate issuer. If set, this field causes IKE to send a certificate request based on this certificate issuer to the server. This key is required if the `IdentityCertificateType` key is included and the `ExtendedAuth.Enabled` key is `true`.
	ServerCertificateIssuerCommonName *string `json:"ServerCertificateIssuerCommonName,omitempty"`
	// The common name of the server certificate. The system uses this name to validate the certificate sent by the IKE server. If not set, the system uses the remote identifier to validate the certificate.
	ServerCertificateCommonName *string `json:"ServerCertificateCommonName,omitempty"`
	// The minimum TLS version to use with EAP-TLS authentication.
	TLSMinimumVersion *IKEv2AuthenticationExtendedAuthTLSMinimumVersion `default:"1.0" json:"TLSMinimumVersion,omitempty"`
	// The maximum TLS version to use with EAP-TLS authentication.
	TLSMaximumVersion *IKEv2AuthenticationExtendedAuthTLSMaximumVersion `default:"1.2" json:"TLSMaximumVersion,omitempty"`
}

// The minimum TLS version to use with EAP-TLS authentication.
type IKEv2AuthenticationExtendedAuthTLSMinimumVersion string

const (
	IKEv2AuthenticationExtendedAuthTLSMinimumVersion10 IKEv2AuthenticationExtendedAuthTLSMinimumVersion = "1.0"
	IKEv2AuthenticationExtendedAuthTLSMinimumVersion11 IKEv2AuthenticationExtendedAuthTLSMinimumVersion = "1.1"
	IKEv2AuthenticationExtendedAuthTLSMinimumVersion12 IKEv2AuthenticationExtendedAuthTLSMinimumVersion = "1.2"
	IKEv2AuthenticationExtendedAuthTLSMinimumVersion13 IKEv2AuthenticationExtendedAuthTLSMinimumVersion = "1.3"
)

// The maximum TLS version to use with EAP-TLS authentication.
type IKEv2AuthenticationExtendedAuthTLSMaximumVersion string

const (
	IKEv2AuthenticationExtendedAuthTLSMaximumVersion10 IKEv2AuthenticationExtendedAuthTLSMaximumVersion = "1.0"
	IKEv2AuthenticationExtendedAuthTLSMaximumVersion11 IKEv2AuthenticationExtendedAuthTLSMaximumVersion = "1.1"
	IKEv2AuthenticationExtendedAuthTLSMaximumVersion12 IKEv2AuthenticationExtendedAuthTLSMaximumVersion = "1.2"
	IKEv2AuthenticationExtendedAuthTLSMaximumVersion13 IKEv2AuthenticationExtendedAuthTLSMaximumVersion = "1.3"
)

// Specifies details about how the system handles idle VPN connections.
type IKEv2Idle struct {
	// One of the following:
	// - `None`: No keepalive.
	// - `Low`: Send keepalive every 30 minutes.
	// - `Medium`: Send keepalive every 10 minutes.
	// - `High`: Send keepalive every 1 minute.
	DeadPeerDetectionRate *IKEv2IdleDeadPeerDetectionRate `default:"Medium" json:"DeadPeerDetectionRate,omitempty"`
}

// One of the following:
// - `None`: No keepalive.
// - `Low`: Send keepalive every 30 minutes.
// - `Medium`: Send keepalive every 10 minutes.
// - `High`: Send keepalive every 1 minute.
type IKEv2IdleDeadPeerDetectionRate string

const (
	IKEv2IdleDeadPeerDetectionRateNone   IKEv2IdleDeadPeerDetectionRate = "None"
	IKEv2IdleDeadPeerDetectionRateLow    IKEv2IdleDeadPeerDetectionRate = "Low"
	IKEv2IdleDeadPeerDetectionRateMedium IKEv2IdleDeadPeerDetectionRate = "Medium"
	IKEv2IdleDeadPeerDetectionRateHigh   IKEv2IdleDeadPeerDetectionRate = "High"
)

// Post Quantum Key Exchange settings.
type IKEv2PostQuantumKeyExchange struct {
	// The Post-quantum Pre-shared key (PPK) the device uses for this VPN. This key is is used with VPN servers that support RFC 8784. If this key is present `PPKIdentifier` must also be present.
	PPK *[]byte `json:"PPK,omitempty"`
	// The identifier for the Post-quantum Pre-shared key (PPK) the device uses for this VPN. This key is is used with VPN servers that support RFC 8784. If this key is present `PPK` must also be present.
	PPKIdentifier *string `json:"PPKIdentifier,omitempty"`
	// If set to `true`, the VPN doesn't establish a connection if the server doesn't support RFC 8784 or doesn't accept the PPK identifier specified in `PPKIdentifier`. The device ignores this key if `PPK` and `PPKIdentifier` aren't present.
	PPKMandatory *bool `json:"PPKMandatory,omitempty"`
	// If set to `false`, the VPN doesn't establish a connection if the server doesn't support or doesn't allow post-quantum key exchanges. Thd device ignores this key if `PostQuantumKeyExchangeMethods` isn't present in `IKESecurityAssociationParameters` or `ChildSecurityAssociationParameters`.
	AllowFallback *bool `json:"AllowFallback,omitempty"`
}

// These parameters apply to Child Security Association unless `ChildSecurityAssociationParameters` is specified.
type IKEv2IKESecurityAssociationParameters struct {
	// The encryption algorithm.
	// In tvOS the default value is `AES-256-GCM`.
	EncryptionAlgorithm *IKEv2IKESecurityAssociationParametersEncryptionAlgorithm `default:"AES-256" json:"EncryptionAlgorithm,omitempty"`
	// The integrity algorithm.
	IntegrityAlgorithm *IKEv2IKESecurityAssociationParametersIntegrityAlgorithm `default:"SHA2-256" json:"IntegrityAlgorithm,omitempty"`
	// The Diffie-Hellman group.
	// For `AlwaysOn` VPN, the minimum allowed value is `14`.
	DiffieHellmanGroup *IKEv2IKESecurityAssociationParametersDiffieHellmanGroup `default:"14" json:"DiffieHellmanGroup,omitempty"`
	// An array of integers representing postquantum key exchange methods the device uses during SA establishment and rekey. You can specify up to seven items, which correspond to ADDKE1 - ADDKE7 from RFC 9370.
	PostQuantumKeyExchangeMethods *[]IKEv2IKESecurityAssociationParametersPostQuantumKeyExchangeMethods `json:"PostQuantumKeyExchangeMethods,omitempty"`
	// The SA lifetime (rekey interval) in minutes.
	LifeTimeInMinutes *int64 `default:"1440" json:"LifeTimeInMinutes,omitempty"`
}

// The encryption algorithm.
// In tvOS the default value is `AES-256-GCM`.
type IKEv2IKESecurityAssociationParametersEncryptionAlgorithm string

const (
	IKEv2IKESecurityAssociationParametersEncryptionAlgorithmAES128           IKEv2IKESecurityAssociationParametersEncryptionAlgorithm = "AES-128"
	IKEv2IKESecurityAssociationParametersEncryptionAlgorithmAES256           IKEv2IKESecurityAssociationParametersEncryptionAlgorithm = "AES-256"
	IKEv2IKESecurityAssociationParametersEncryptionAlgorithmAES128GCM        IKEv2IKESecurityAssociationParametersEncryptionAlgorithm = "AES-128-GCM"
	IKEv2IKESecurityAssociationParametersEncryptionAlgorithmAES256GCM        IKEv2IKESecurityAssociationParametersEncryptionAlgorithm = "AES-256-GCM"
	IKEv2IKESecurityAssociationParametersEncryptionAlgorithmChaCha20Poly1305 IKEv2IKESecurityAssociationParametersEncryptionAlgorithm = "ChaCha20Poly1305"
)

// The integrity algorithm.
type IKEv2IKESecurityAssociationParametersIntegrityAlgorithm string

const (
	IKEv2IKESecurityAssociationParametersIntegrityAlgorithmSHA2256 IKEv2IKESecurityAssociationParametersIntegrityAlgorithm = "SHA2-256"
	IKEv2IKESecurityAssociationParametersIntegrityAlgorithmSHA2384 IKEv2IKESecurityAssociationParametersIntegrityAlgorithm = "SHA2-384"
	IKEv2IKESecurityAssociationParametersIntegrityAlgorithmSHA2512 IKEv2IKESecurityAssociationParametersIntegrityAlgorithm = "SHA2-512"
)

// The Diffie-Hellman group.
// For `AlwaysOn` VPN, the minimum allowed value is `14`.
type IKEv2IKESecurityAssociationParametersDiffieHellmanGroup int64

const (
	IKEv2IKESecurityAssociationParametersDiffieHellmanGroup14 IKEv2IKESecurityAssociationParametersDiffieHellmanGroup = 14
	IKEv2IKESecurityAssociationParametersDiffieHellmanGroup15 IKEv2IKESecurityAssociationParametersDiffieHellmanGroup = 15
	IKEv2IKESecurityAssociationParametersDiffieHellmanGroup16 IKEv2IKESecurityAssociationParametersDiffieHellmanGroup = 16
	IKEv2IKESecurityAssociationParametersDiffieHellmanGroup17 IKEv2IKESecurityAssociationParametersDiffieHellmanGroup = 17
	IKEv2IKESecurityAssociationParametersDiffieHellmanGroup18 IKEv2IKESecurityAssociationParametersDiffieHellmanGroup = 18
	IKEv2IKESecurityAssociationParametersDiffieHellmanGroup19 IKEv2IKESecurityAssociationParametersDiffieHellmanGroup = 19
	IKEv2IKESecurityAssociationParametersDiffieHellmanGroup20 IKEv2IKESecurityAssociationParametersDiffieHellmanGroup = 20
	IKEv2IKESecurityAssociationParametersDiffieHellmanGroup21 IKEv2IKESecurityAssociationParametersDiffieHellmanGroup = 21
	IKEv2IKESecurityAssociationParametersDiffieHellmanGroup31 IKEv2IKESecurityAssociationParametersDiffieHellmanGroup = 31
	IKEv2IKESecurityAssociationParametersDiffieHellmanGroup32 IKEv2IKESecurityAssociationParametersDiffieHellmanGroup = 32
)

// An array of integers representing postquantum key exchange methods the device uses during SA establishment and rekey. You can specify up to seven items, which correspond to ADDKE1 - ADDKE7 from RFC 9370.
type IKEv2IKESecurityAssociationParametersPostQuantumKeyExchangeMethods int64

const (
	IKEv2IKESecurityAssociationParametersPostQuantumKeyExchangeMethods0  IKEv2IKESecurityAssociationParametersPostQuantumKeyExchangeMethods = 0
	IKEv2IKESecurityAssociationParametersPostQuantumKeyExchangeMethods36 IKEv2IKESecurityAssociationParametersPostQuantumKeyExchangeMethods = 36
	IKEv2IKESecurityAssociationParametersPostQuantumKeyExchangeMethods37 IKEv2IKESecurityAssociationParametersPostQuantumKeyExchangeMethods = 37
)

// The `ChildSecurityAssociationParameters` dictionaries.
type IKEv2ChildSecurityAssociationParameters struct {
	// The encryption algorithm.
	// In tvOS the default value is `AES-256-GCM`.
	EncryptionAlgorithm *IKEv2IKESecurityAssociationParametersEncryptionAlgorithm `default:"AES-256" json:"EncryptionAlgorithm,omitempty"`
	// The integrity algorithm.
	IntegrityAlgorithm *IKEv2IKESecurityAssociationParametersIntegrityAlgorithm `default:"SHA2-256" json:"IntegrityAlgorithm,omitempty"`
	// The Diffie-Hellman group.
	// For `AlwaysOn` VPN, the minimum allowed value is `14`.
	DiffieHellmanGroup *IKEv2IKESecurityAssociationParametersDiffieHellmanGroup `default:"14" json:"DiffieHellmanGroup,omitempty"`
	// An array of integers representing postquantum key exchange methods the device uses during SA establishment and rekey. You can specify up to seven items, which correspond to ADDKE1 - ADDKE7 from RFC 9370.
	PostQuantumKeyExchangeMethods *[]IKEv2IKESecurityAssociationParametersPostQuantumKeyExchangeMethods `json:"PostQuantumKeyExchangeMethods,omitempty"`
	// The SA lifetime (rekey interval) in minutes.
	LifeTimeInMinutes *int64 `default:"1440" json:"LifeTimeInMinutes,omitempty"`
}
type ServiceExceptionElement struct {
	// The name of a service that's exempt from Always On VPN.
	// `CellularServices` exempts `VoLTE`, `IMS`, `MMS`, and Wi-Fi calling.
	// `DeviceCommunication` exempts network traffic used for communicating with devices connected via USB or Wi-Fi.
	ServiceName ServiceName `json:"ServiceName" required:"true"`
	// The action to take with network connections from the named service.
	Action ServiceExceptionElementAction `json:"Action" required:"true"`
}

// The name of a service that's exempt from Always On VPN.
// `CellularServices` exempts `VoLTE`, `IMS`, `MMS`, and Wi-Fi calling.
// `DeviceCommunication` exempts network traffic used for communicating with devices connected via USB or Wi-Fi.
type ServiceName string

const (
	ServiceNameVoiceMail           ServiceName = "VoiceMail"
	ServiceNameAirPrint            ServiceName = "AirPrint"
	ServiceNameCellularServices    ServiceName = "CellularServices"
	ServiceNameDeviceCommunication ServiceName = "DeviceCommunication"
)

// The action to take with network connections from the named service.
type ServiceExceptionElementAction string

const (
	ServiceExceptionElementActionAllow ServiceExceptionElementAction = "Allow"
	ServiceExceptionElementActionDrop  ServiceExceptionElementAction = "Drop"
)

type ApplicationExceptionElement struct {
	// The app's bundle identifier.
	BundleIdentifier string `json:"BundleIdentifier" required:"true"`
	// Limit the exception to only the specified list of protocols, with support for `UDP` only.
	LimitToProtocols *[]LimitToProtocols `json:"LimitToProtocols,omitempty"`
}

// Limit the exception to only the specified list of protocols, with support for `UDP` only.
type LimitToProtocols string

const (
	LimitToProtocolsUDP LimitToProtocols = "UDP"
)

type AllowedCaptiveNetworkPluginElement struct {
	// The bundle identifier for the app that's allowed on the captive network.
	BundleIdentifier string `json:"BundleIdentifier" required:"true"`
}

// The declaration to configure a VPN using the IKEv2 sub-type.
type NetworkVPNIKEv2 struct {
	// The name of the VPN connection that the system displays on the device.
	VisibleName string `json:"VisibleName" required:"true"`
	// The IP address or hostname of the VPN server.
	HostName string `json:"HostName" required:"true"`
	// Identifier of the IKEv2 client.
	LocalIdentifier string `json:"LocalIdentifier" required:"true"`
	// The remote identifier.
	RemoteIdentifier string `json:"RemoteIdentifier" required:"true"`
	// Settings that control authentication.
	Authentication NetworkVPNIKEv2Authentication `json:"Authentication" required:"true"`
	// Specifies details about how the VPN routes different types of network traffic.
	NetworkRouting *NetworkVPNIKEv2NetworkRouting `json:"NetworkRouting,omitempty"`
	// Specifies details about how the system handles idle VPN connections.
	Idle *NetworkVPNIKEv2Idle `json:"Idle,omitempty"`
	// Specifies details about how the system controls on-demand VPN.
	OnDemand *NetworkVPNIKEv2OnDemand `json:"OnDemand,omitempty"`
	// If `true`, negotiations should use IKEv2 Configuration Attribute `INTERNAL_IP4_SUBNET` and `INTERNAL_IP6_SUBNET`.
	UseConfigurationAttributeInternalIPSubnet *bool `json:"UseConfigurationAttributeInternalIPSubnet,omitempty"`
	// If `true`, the system disables MOBIKE.
	DisableMOBIKE *bool `json:"DisableMOBIKE,omitempty"`
	// If `true`, the system disables IKEv2 redirect. If not set, the system redirects an IKEv2 connection when it receives a redirect request from the server.
	DisableRedirect *bool `json:"DisableRedirect,omitempty"`
	// If `true`,  enables Perfect Forward Secrecy (PFS) for IKEv2 Connections.
	EnablePFS *bool `json:"EnablePFS,omitempty"`
	// If `true`, the system performs a certificate revocation check for IKEv2 connections. This is a best-effort revocation check and server response timeouts won't cause it to fail.
	EnableCertificateRevocationCheck *bool `json:"EnableCertificateRevocationCheck,omitempty"`
	// If `true`, the system enables a tunnel over cellular data to carry traffic that's eligible for Wi-Fi Assist and also requires VPN.
	// Enabling fallback requires that the server support multiple tunnels for a single user.
	EnableFallback *bool `json:"EnableFallback,omitempty"`
	// The Maximum Transmission Unit (MTU) specifies the maximum size in bytes of each packet that the system sends over the IKEv2 VPN interface.
	MTU *int64 `default:"1280" json:"MTU,omitempty"`
	// If set to `true`, the device doesn't allow DES, 3DES, and Diffie-Hellman groups less than 14. Also the device requires the encryption algorithm specified in `IKESecurityAssociationParameters` to be at least as cryptographically strong as the algorithm specified in `ChildSecurityAssociationParameters`. The device rejects this configuration if these requirements aren't met.
	EnforceStrictAlgorithmSelection *bool `json:"EnforceStrictAlgorithmSelection,omitempty"`
	// Post Quantum Key Exchange settings.
	PostQuantumKeyExchange *NetworkVPNIKEv2PostQuantumKeyExchange `json:"PostQuantumKeyExchange,omitempty"`
	// These parameters apply to Child Security Association unless `ChildSecurityAssociationParameters` is specified.
	IKESecurityAssociationParameters *NetworkVPNIKEv2IKESecurityAssociationParameters `json:"IKESecurityAssociationParameters,omitempty"`
	// The `ChildSecurityAssociationParameters` dictionaries.
	ChildSecurityAssociationParameters *NetworkVPNIKEv2ChildSecurityAssociationParameters `json:"ChildSecurityAssociationParameters,omitempty"`
	// A dictionary to use for all VPN types.
	DNS *NetworkVPNIKEv2DNS `json:"DNS,omitempty"`
	// The dictionary to use to configure `Proxies` for use with `VPN`.
	Proxies *NetworkVPNIKEv2Proxies `json:"Proxies,omitempty"`
}

func (p *NetworkVPNIKEv2) DeclarationType() string {
	return "com.apple.configuration.network.vpn.ikev2"
}

// Settings that control authentication.
type NetworkVPNIKEv2Authentication struct {
	// The type of authentication method for the VPN.
	// To enable EAP-only authentication, set this to `None` and `ExtendedAuthEnabled` to `true`. If this is `None` and the `ExtendedAuthEnabled` key isn't set, the authentication configuration defaults to `SharedSecret`.
	Method NetworkVPNIKEv2AuthenticationMethod `json:"Method" required:"true"`
	// The identifier of an asset declaration that contains the credentials (password) to authenticate with the VPN server. Required when `Authentication.Method` is set to `SharedSecret`.
	CredentialsAssetReference *string `json:"CredentialsAssetReference,omitempty"`
	// The type of key used by the identity set in the `IdentityAssetReference` to use for IKEv2 machine authentication. If this key is included, the system requires a value for `ServerCertificateIssuerCommonName`.
	IdentityCertificateType *NetworkVPNIKEv2AuthenticationIdentityCertificateType `default:"RSA" json:"IdentityCertificateType,omitempty"`
	// The identifier of a credential asset declaration that contains the identity that this account requires to authenticate with the VPN server. If the value of `AuthenticationMethod` is `Certificate`, the system sends this certificate out for IKEv2 machine authentication. If extended authentication (EAP) is used, the system sends this certificate out for EAP-TLS authentication. Required when `Authentication.Method` is set to `Certificate`.
	IdentityAssetReference *string `json:"IdentityAssetReference,omitempty"`
	// Specifies details about extended authentication (EAP).
	ExtendedAuth *NetworkVPNIKEv2AuthenticationExtendedAuth `json:"ExtendedAuth,omitempty"`
}

// The type of authentication method for the VPN.
// To enable EAP-only authentication, set this to `None` and `ExtendedAuthEnabled` to `true`. If this is `None` and the `ExtendedAuthEnabled` key isn't set, the authentication configuration defaults to `SharedSecret`.
type NetworkVPNIKEv2AuthenticationMethod string

const (
	NetworkVPNIKEv2AuthenticationMethodNone         NetworkVPNIKEv2AuthenticationMethod = "None"
	NetworkVPNIKEv2AuthenticationMethodSharedSecret NetworkVPNIKEv2AuthenticationMethod = "SharedSecret"
	NetworkVPNIKEv2AuthenticationMethodCertificate  NetworkVPNIKEv2AuthenticationMethod = "Certificate"
)

// The type of key used by the identity set in the `IdentityAssetReference` to use for IKEv2 machine authentication. If this key is included, the system requires a value for `ServerCertificateIssuerCommonName`.
type NetworkVPNIKEv2AuthenticationIdentityCertificateType string

const (
	NetworkVPNIKEv2AuthenticationIdentityCertificateTypeRSA      NetworkVPNIKEv2AuthenticationIdentityCertificateType = "RSA"
	NetworkVPNIKEv2AuthenticationIdentityCertificateTypeECDSA256 NetworkVPNIKEv2AuthenticationIdentityCertificateType = "ECDSA256"
	NetworkVPNIKEv2AuthenticationIdentityCertificateTypeECDSA384 NetworkVPNIKEv2AuthenticationIdentityCertificateType = "ECDSA384"
	NetworkVPNIKEv2AuthenticationIdentityCertificateTypeECDSA521 NetworkVPNIKEv2AuthenticationIdentityCertificateType = "ECDSA521"
	NetworkVPNIKEv2AuthenticationIdentityCertificateTypeRSAPSS   NetworkVPNIKEv2AuthenticationIdentityCertificateType = "RSA-PSS"
)

// Specifies details about extended authentication (EAP).
type NetworkVPNIKEv2AuthenticationExtendedAuth struct {
	// If `true`, enables EAP-only authentication.
	Enabled *bool `json:"Enabled,omitempty"`
	// The identifier of an asset declaration that contains the credentials (user name and password) to authenticate with the VPN server. Required when `Enabled` is set to `true`. Implies the use of EAP-MSCHAPv2.
	CredentialsAssetReference *string `json:"CredentialsAssetReference,omitempty"`
	// Common Name of the server certificate issuer. If set, this field causes IKE to send a certificate request based on this certificate issuer to the server. This key is required if the `IdentityCertificateType` key is included and the `ExtendedAuth.Enabled` key is `true`.
	ServerCertificateIssuerCommonName *string `json:"ServerCertificateIssuerCommonName,omitempty"`
	// The common name of the server certificate. The system uses this name to validate the certificate sent by the IKE server. If not set, the system uses the remote identifier to validate the certificate.
	ServerCertificateCommonName *string `json:"ServerCertificateCommonName,omitempty"`
	// The minimum TLS version to use with EAP-TLS authentication.
	TLSMinimumVersion *NetworkVPNIKEv2AuthenticationExtendedAuthTLSMinimumVersion `default:"1.0" json:"TLSMinimumVersion,omitempty"`
	// The maximum TLS version to use with EAP-TLS authentication.
	TLSMaximumVersion *NetworkVPNIKEv2AuthenticationExtendedAuthTLSMaximumVersion `default:"1.2" json:"TLSMaximumVersion,omitempty"`
}

// The minimum TLS version to use with EAP-TLS authentication.
type NetworkVPNIKEv2AuthenticationExtendedAuthTLSMinimumVersion string

const (
	NetworkVPNIKEv2AuthenticationExtendedAuthTLSMinimumVersion10 NetworkVPNIKEv2AuthenticationExtendedAuthTLSMinimumVersion = "1.0"
	NetworkVPNIKEv2AuthenticationExtendedAuthTLSMinimumVersion11 NetworkVPNIKEv2AuthenticationExtendedAuthTLSMinimumVersion = "1.1"
	NetworkVPNIKEv2AuthenticationExtendedAuthTLSMinimumVersion12 NetworkVPNIKEv2AuthenticationExtendedAuthTLSMinimumVersion = "1.2"
	NetworkVPNIKEv2AuthenticationExtendedAuthTLSMinimumVersion13 NetworkVPNIKEv2AuthenticationExtendedAuthTLSMinimumVersion = "1.3"
)

// The maximum TLS version to use with EAP-TLS authentication.
type NetworkVPNIKEv2AuthenticationExtendedAuthTLSMaximumVersion string

const (
	NetworkVPNIKEv2AuthenticationExtendedAuthTLSMaximumVersion10 NetworkVPNIKEv2AuthenticationExtendedAuthTLSMaximumVersion = "1.0"
	NetworkVPNIKEv2AuthenticationExtendedAuthTLSMaximumVersion11 NetworkVPNIKEv2AuthenticationExtendedAuthTLSMaximumVersion = "1.1"
	NetworkVPNIKEv2AuthenticationExtendedAuthTLSMaximumVersion12 NetworkVPNIKEv2AuthenticationExtendedAuthTLSMaximumVersion = "1.2"
	NetworkVPNIKEv2AuthenticationExtendedAuthTLSMaximumVersion13 NetworkVPNIKEv2AuthenticationExtendedAuthTLSMaximumVersion = "1.3"
)

// Specifies details about how the VPN routes different types of network traffic.
type NetworkVPNIKEv2NetworkRouting struct {
	// If `true`, all the VPN's non-default routes take precedence over any locally defined routes.
	// If `IncludeAllNetworks` is `true`, the system ignores the value of `EnforceRoutes`.
	EnforceRoutes *bool `json:"EnforceRoutes,omitempty"`
	// If `true`, routes all traffic through the VPN, with some exclusions. Several of the exclusions can be controlled with the `ExcludeLocalNetworks`, `ExcludeCellularServices`, `ExcludeAPNs` and `ExcludeDeviceCommunication` properties. The following traffic is always excluded from the tunnel:
	// - Traffic necessary for connecting and maintaining the device's network connection, such as DHCP.
	// - Traffic necessary for connecting to captive networks.
	// - Certain cellular services traffic that's not routable over the internet and is instead directly routed to the cellular network. See the ExcludeCellularServices property for more details.
	IncludeAllNetworks *bool `json:"IncludeAllNetworks,omitempty"`
	// If `true` and `IncludeAllNetworks` is `true`, routes all local network traffic outside the VPN.
	ExcludeLocalNetworks *bool `json:"ExcludeLocalNetworks,omitempty"`
	// If `true` and `IncludeAllNetworks` is `true`, then the system excludes internet-routable network traffic for cellular services (VoLTE, Wi-Fi Calling, IMS, MMS, Visual Voicemail, etc.) from the tunnel. Note that some cellular carriers route cellular services traffic directly to the carrier network, bypassing the internet. Such cellular services traffic is always excluded from the tunnel.
	ExcludeCellularServices *bool `json:"ExcludeCellularServices,omitempty"`
	// If `true` and `IncludeAllNetworks` is `true`, then the system excludes the network traffic for the Apple Push Notification service (APNs) from the tunnel.
	ExcludeAPNs *bool `json:"ExcludeAPNs,omitempty"`
	// If set to `true` and `IncludeAllNetworks` is set to `true`, the device excludes network traffic used for communicating with devices connected via USB or Wi-Fi from the tunnel.
	ExcludeDeviceCommunication *bool `json:"ExcludeDeviceCommunication,omitempty"`
}

// Specifies details about how the system handles idle VPN connections.
type NetworkVPNIKEv2Idle struct {
	// If `true`, disconnects after an on-demand connection idles.
	Disconnect *bool `json:"Disconnect,omitempty"`
	// The length of time to wait, in seconds, before disconnecting an on-demand connection.
	Timer *int64 `json:"Timer,omitempty"`
	// One of the following:
	// - `None`: No keepalive.
	// - `Low`: Send keepalive every 30 minutes.
	// - `Medium`: Send keepalive every 10 minutes.
	// - `High`: Send keepalive every 1 minute.
	DeadPeerDetectionRate *NetworkVPNIKEv2IdleDeadPeerDetectionRate `default:"Medium" json:"DeadPeerDetectionRate,omitempty"`
}

// One of the following:
// - `None`: No keepalive.
// - `Low`: Send keepalive every 30 minutes.
// - `Medium`: Send keepalive every 10 minutes.
// - `High`: Send keepalive every 1 minute.
type NetworkVPNIKEv2IdleDeadPeerDetectionRate string

const (
	NetworkVPNIKEv2IdleDeadPeerDetectionRateNone   NetworkVPNIKEv2IdleDeadPeerDetectionRate = "None"
	NetworkVPNIKEv2IdleDeadPeerDetectionRateLow    NetworkVPNIKEv2IdleDeadPeerDetectionRate = "Low"
	NetworkVPNIKEv2IdleDeadPeerDetectionRateMedium NetworkVPNIKEv2IdleDeadPeerDetectionRate = "Medium"
	NetworkVPNIKEv2IdleDeadPeerDetectionRateHigh   NetworkVPNIKEv2IdleDeadPeerDetectionRate = "High"
)

// Specifies details about how the system controls on-demand VPN.
type NetworkVPNIKEv2OnDemand struct {
	// If `true`, enables VPN On Demand.
	Enabled *bool `json:"Enabled,omitempty"`
	// If `true`, the device disables the Connect On Demand toggle in Settings for this configuration.
	DisableUserOverride *bool `json:"DisableUserOverride,omitempty"`
	// An array of dictionaries defining On Demand Rules.
	Rules *[]*NetworkVPNIKEv2OnDemandRulesRulesElement `json:"Rules,omitempty"`
}
type NetworkVPNIKEv2OnDemandRulesRulesElement struct {
	// The action to take if this dictionary matches the current network. Possible values are:
	// - `Allow`: Deprecated. Allow VPN On Demand to connect if triggered.
	// - `Connect`: Unconditionally initiate a VPN connection on the next network attempt.
	// - `Disconnect`: Tear down the VPN connection and don't reconnect on demand as long as this dictionary matches.
	// - `EvaluateConnection`: Evaluate the ActionParameters array for each connection attempt.
	// - `Ignore`: Leave any existing VPN connection up, but don't reconnect on demand as long as this dictionary matches.
	Action NetworkVPNIKEv2OnDemandRulesRulesElementAction `json:"Action" required:"true"`
	// An array of dictionaries that provides rules similar to the `OnDemandRules` dictionary, but evaluated on each connection instead of when the network changes. This value is only for use with dictionaries in which the `Action` value is `EvaluateConnection`. The system evaluates these dictionaries in order and the first dictionary that matches determines the behavior.
	ActionParameters *[]*NetworkVPNIKEv2OnDemandRulesRulesElementActionParametersActionParameter `json:"ActionParameters,omitempty"`
	// An array of domain names. This rule matches if any of the domain names in the specified list matches any domain in the device's search domains list.
	// The system supports a wildcard (`\*`) prefix. For example, `\*.example.com` matches against either `mydomain.example.com` or `yourdomain.example.com`.
	DNSDomainMatch *[]string `json:"DNSDomainMatch,omitempty"`
	// An array of IP addresses. This rule matches if any of the network's specified DNS servers match any entry in the array.
	// The system supports matching with a single wildcard. For example, `17.\*` matches any DNS server in the `17.0.0.0/8` subnet.
	DNSServerAddressMatch *[]string `json:"DNSServerAddressMatch,omitempty"`
	// An interface type. If specified, this rule matches only if the primary network interface hardware matches the specified type.
	InterfaceTypeMatch *NetworkVPNIKEv2OnDemandRulesRulesElementInterfaceTypeMatch `json:"InterfaceTypeMatch,omitempty"`
	// An array of SSIDs to match against the current network. If the network isn't a Wi-Fi network or if the SSID doesn't appear in this array, the match fails.
	// Omit this key and the corresponding array to match against any SSID.
	SSIDMatch *[]string `json:"SSIDMatch,omitempty"`
	// A URL to probe. This rule matches when this URL is successfully fetched (returns a `200` HTTP status code) without redirection.
	URLStringProbe *string `json:"URLStringProbe,omitempty"`
}

// The action to take if this dictionary matches the current network. Possible values are:
// - `Allow`: Deprecated. Allow VPN On Demand to connect if triggered.
// - `Connect`: Unconditionally initiate a VPN connection on the next network attempt.
// - `Disconnect`: Tear down the VPN connection and don't reconnect on demand as long as this dictionary matches.
// - `EvaluateConnection`: Evaluate the ActionParameters array for each connection attempt.
// - `Ignore`: Leave any existing VPN connection up, but don't reconnect on demand as long as this dictionary matches.
type NetworkVPNIKEv2OnDemandRulesRulesElementAction string

const (
	NetworkVPNIKEv2OnDemandRulesRulesElementActionAllow              NetworkVPNIKEv2OnDemandRulesRulesElementAction = "Allow"
	NetworkVPNIKEv2OnDemandRulesRulesElementActionConnect            NetworkVPNIKEv2OnDemandRulesRulesElementAction = "Connect"
	NetworkVPNIKEv2OnDemandRulesRulesElementActionDisconnect         NetworkVPNIKEv2OnDemandRulesRulesElementAction = "Disconnect"
	NetworkVPNIKEv2OnDemandRulesRulesElementActionEvaluateConnection NetworkVPNIKEv2OnDemandRulesRulesElementAction = "EvaluateConnection"
	NetworkVPNIKEv2OnDemandRulesRulesElementActionIgnore             NetworkVPNIKEv2OnDemandRulesRulesElementAction = "Ignore"
)

// A dictionary that provides rules similar to the OnDemandRules dictionary, but evaluated on each connection instead of when the network changes. These dictionaries are evaluated in order, and the behavior is determined by the first dictionary that matches.
// The keys allowed in each dictionary are described below. Note: This array is used only for dictionaries in which EvaluateConnection is the Action value.
type NetworkVPNIKEv2OnDemandRulesRulesElementActionParametersActionParameter struct {
	// The domains to apply this evaluation.
	Domains []string `json:"Domains" required:"true"`
	// Defines the VPN behavior for the specified domains. Allowed values are:
	// * 'ConnectIfNeeded': The specified domains should trigger a VPN connection attempt if domain name resolution fails, such as when the DNS server indicates that it can't resolve the domain, responds with a redirection to a different server, or fails to respond (timeout).
	// * 'NeverConnect': The specified domains should never trigger a VPN connection attempt.
	DomainAction NetworkVPNIKEv2OnDemandRulesRulesElementActionParametersActionParameterDomainAction `json:"DomainAction" required:"true"`
	// An array of IP addresses of DNS servers to use for resolving the specified domains. These servers don't need to be part of the device's current network configuration. If these DNS servers aren't reachable, the system establishes a VPN connection. These DNS servers need to be either internal DNS servers or trusted external DNS servers.
	// This key is valid only if the value of 'DomainAction' is 'ConnectIfNeeded'.
	RequiredDNSServers *[]string `json:"RequiredDNSServers,omitempty"`
	// An HTTP or HTTPS (preferred) URL to probe, using a GET request. If the URL's hostname can't be resolved, if the server is unreachable, or if the server doesn't respond with a 200 HTTP status code, a VPN connection is established in response.
	// This key is valid only if the value of 'DomainAction' is 'ConnectIfNeeded'.
	RequiredURLStringProbe *string `json:"RequiredURLStringProbe,omitempty"`
}

// Defines the VPN behavior for the specified domains. Allowed values are:
// * 'ConnectIfNeeded': The specified domains should trigger a VPN connection attempt if domain name resolution fails, such as when the DNS server indicates that it can't resolve the domain, responds with a redirection to a different server, or fails to respond (timeout).
// * 'NeverConnect': The specified domains should never trigger a VPN connection attempt.
type NetworkVPNIKEv2OnDemandRulesRulesElementActionParametersActionParameterDomainAction string

const (
	NetworkVPNIKEv2OnDemandRulesRulesElementActionParametersActionParameterDomainActionConnectIfNeeded NetworkVPNIKEv2OnDemandRulesRulesElementActionParametersActionParameterDomainAction = "ConnectIfNeeded"
	NetworkVPNIKEv2OnDemandRulesRulesElementActionParametersActionParameterDomainActionNeverConnect    NetworkVPNIKEv2OnDemandRulesRulesElementActionParametersActionParameterDomainAction = "NeverConnect"
)

// An interface type. If specified, this rule matches only if the primary network interface hardware matches the specified type.
type NetworkVPNIKEv2OnDemandRulesRulesElementInterfaceTypeMatch string

const (
	NetworkVPNIKEv2OnDemandRulesRulesElementInterfaceTypeMatchEthernet NetworkVPNIKEv2OnDemandRulesRulesElementInterfaceTypeMatch = "Ethernet"
	NetworkVPNIKEv2OnDemandRulesRulesElementInterfaceTypeMatchWiFi     NetworkVPNIKEv2OnDemandRulesRulesElementInterfaceTypeMatch = "WiFi"
	NetworkVPNIKEv2OnDemandRulesRulesElementInterfaceTypeMatchCellular NetworkVPNIKEv2OnDemandRulesRulesElementInterfaceTypeMatch = "Cellular"
)

// Post Quantum Key Exchange settings.
type NetworkVPNIKEv2PostQuantumKeyExchange struct {
	// The Post-quantum Pre-shared key (PPK) the device uses for this VPN. This key is is used with VPN servers that support RFC 8784. If this key is present `PPKIdentifier` must also be present.
	PPK *[]byte `json:"PPK,omitempty"`
	// The identifier for the Post-quantum Pre-shared key (PPK) the device uses for this VPN. This key is is used with VPN servers that support RFC 8784. If this key is present `PPK` must also be present.
	PPKIdentifier *string `json:"PPKIdentifier,omitempty"`
	// If set to `true`, the VPN doesn't establish a connection if the server doesn't support RFC 8784 or doesn't accept the PPK identifier specified in `PPKIdentifier`. The device ignores this key if `PPK` and `PPKIdentifier` aren't present.
	PPKMandatory *bool `json:"PPKMandatory,omitempty"`
	// If set to `false`, the VPN doesn't establish a connection if the server doesn't support or doesn't allow post-quantum key exchanges. Thd device ignores this key if `PostQuantumKeyExchangeMethods` isn't present in `IKESecurityAssociationParameters` or `ChildSecurityAssociationParameters`.
	AllowFallback *bool `json:"AllowFallback,omitempty"`
}

// These parameters apply to Child Security Association unless `ChildSecurityAssociationParameters` is specified.
type NetworkVPNIKEv2IKESecurityAssociationParameters struct {
	// The encryption algorithm.
	// On tvOS, the default value is `AES-256-GCM`.
	EncryptionAlgorithm *NetworkVPNIKEv2IKESecurityAssociationParametersEncryptionAlgorithm `default:"AES-256" json:"EncryptionAlgorithm,omitempty"`
	// The integrity algorithm.
	IntegrityAlgorithm *NetworkVPNIKEv2IKESecurityAssociationParametersIntegrityAlgorithm `default:"SHA2-256" json:"IntegrityAlgorithm,omitempty"`
	// The Diffie-Hellman group.
	// For `AlwaysOn` VPN, the minimum allowed value is `14`.
	DiffieHellmanGroup *NetworkVPNIKEv2IKESecurityAssociationParametersDiffieHellmanGroup `default:"14" json:"DiffieHellmanGroup,omitempty"`
	// An array of integers representing postquantum key exchange methods the device uses during SA establishment and rekey. You can specify up to seven items, which correspond to ADDKE1 - ADDKE7 from RFC 9370.
	PostQuantumKeyExchangeMethods *[]NetworkVPNIKEv2IKESecurityAssociationParametersPostQuantumKeyExchangeMethods `json:"PostQuantumKeyExchangeMethods,omitempty"`
	// The SA lifetime (rekey interval) in minutes.
	LifeTimeInMinutes *int64 `default:"1440" json:"LifeTimeInMinutes,omitempty"`
}

// The encryption algorithm.
// On tvOS, the default value is `AES-256-GCM`.
type NetworkVPNIKEv2IKESecurityAssociationParametersEncryptionAlgorithm string

const (
	NetworkVPNIKEv2IKESecurityAssociationParametersEncryptionAlgorithmAES128           NetworkVPNIKEv2IKESecurityAssociationParametersEncryptionAlgorithm = "AES-128"
	NetworkVPNIKEv2IKESecurityAssociationParametersEncryptionAlgorithmAES256           NetworkVPNIKEv2IKESecurityAssociationParametersEncryptionAlgorithm = "AES-256"
	NetworkVPNIKEv2IKESecurityAssociationParametersEncryptionAlgorithmAES128GCM        NetworkVPNIKEv2IKESecurityAssociationParametersEncryptionAlgorithm = "AES-128-GCM"
	NetworkVPNIKEv2IKESecurityAssociationParametersEncryptionAlgorithmAES256GCM        NetworkVPNIKEv2IKESecurityAssociationParametersEncryptionAlgorithm = "AES-256-GCM"
	NetworkVPNIKEv2IKESecurityAssociationParametersEncryptionAlgorithmChaCha20Poly1305 NetworkVPNIKEv2IKESecurityAssociationParametersEncryptionAlgorithm = "ChaCha20Poly1305"
)

// The integrity algorithm.
type NetworkVPNIKEv2IKESecurityAssociationParametersIntegrityAlgorithm string

const (
	NetworkVPNIKEv2IKESecurityAssociationParametersIntegrityAlgorithmSHA2256 NetworkVPNIKEv2IKESecurityAssociationParametersIntegrityAlgorithm = "SHA2-256"
	NetworkVPNIKEv2IKESecurityAssociationParametersIntegrityAlgorithmSHA2384 NetworkVPNIKEv2IKESecurityAssociationParametersIntegrityAlgorithm = "SHA2-384"
	NetworkVPNIKEv2IKESecurityAssociationParametersIntegrityAlgorithmSHA2512 NetworkVPNIKEv2IKESecurityAssociationParametersIntegrityAlgorithm = "SHA2-512"
)

// The Diffie-Hellman group.
// For `AlwaysOn` VPN, the minimum allowed value is `14`.
type NetworkVPNIKEv2IKESecurityAssociationParametersDiffieHellmanGroup int64

const (
	NetworkVPNIKEv2IKESecurityAssociationParametersDiffieHellmanGroup14 NetworkVPNIKEv2IKESecurityAssociationParametersDiffieHellmanGroup = 14
	NetworkVPNIKEv2IKESecurityAssociationParametersDiffieHellmanGroup15 NetworkVPNIKEv2IKESecurityAssociationParametersDiffieHellmanGroup = 15
	NetworkVPNIKEv2IKESecurityAssociationParametersDiffieHellmanGroup16 NetworkVPNIKEv2IKESecurityAssociationParametersDiffieHellmanGroup = 16
	NetworkVPNIKEv2IKESecurityAssociationParametersDiffieHellmanGroup17 NetworkVPNIKEv2IKESecurityAssociationParametersDiffieHellmanGroup = 17
	NetworkVPNIKEv2IKESecurityAssociationParametersDiffieHellmanGroup18 NetworkVPNIKEv2IKESecurityAssociationParametersDiffieHellmanGroup = 18
	NetworkVPNIKEv2IKESecurityAssociationParametersDiffieHellmanGroup19 NetworkVPNIKEv2IKESecurityAssociationParametersDiffieHellmanGroup = 19
	NetworkVPNIKEv2IKESecurityAssociationParametersDiffieHellmanGroup20 NetworkVPNIKEv2IKESecurityAssociationParametersDiffieHellmanGroup = 20
	NetworkVPNIKEv2IKESecurityAssociationParametersDiffieHellmanGroup21 NetworkVPNIKEv2IKESecurityAssociationParametersDiffieHellmanGroup = 21
	NetworkVPNIKEv2IKESecurityAssociationParametersDiffieHellmanGroup31 NetworkVPNIKEv2IKESecurityAssociationParametersDiffieHellmanGroup = 31
	NetworkVPNIKEv2IKESecurityAssociationParametersDiffieHellmanGroup32 NetworkVPNIKEv2IKESecurityAssociationParametersDiffieHellmanGroup = 32
)

// An array of integers representing postquantum key exchange methods the device uses during SA establishment and rekey. You can specify up to seven items, which correspond to ADDKE1 - ADDKE7 from RFC 9370.
type NetworkVPNIKEv2IKESecurityAssociationParametersPostQuantumKeyExchangeMethods int64

const (
	NetworkVPNIKEv2IKESecurityAssociationParametersPostQuantumKeyExchangeMethods0  NetworkVPNIKEv2IKESecurityAssociationParametersPostQuantumKeyExchangeMethods = 0
	NetworkVPNIKEv2IKESecurityAssociationParametersPostQuantumKeyExchangeMethods36 NetworkVPNIKEv2IKESecurityAssociationParametersPostQuantumKeyExchangeMethods = 36
	NetworkVPNIKEv2IKESecurityAssociationParametersPostQuantumKeyExchangeMethods37 NetworkVPNIKEv2IKESecurityAssociationParametersPostQuantumKeyExchangeMethods = 37
)

// The `ChildSecurityAssociationParameters` dictionaries.
type NetworkVPNIKEv2ChildSecurityAssociationParameters struct {
	// The encryption algorithm.
	// On tvOS, the default value is `AES-256-GCM`.
	EncryptionAlgorithm *NetworkVPNIKEv2IKESecurityAssociationParametersEncryptionAlgorithm `default:"AES-256" json:"EncryptionAlgorithm,omitempty"`
	// The integrity algorithm.
	IntegrityAlgorithm *NetworkVPNIKEv2IKESecurityAssociationParametersIntegrityAlgorithm `default:"SHA2-256" json:"IntegrityAlgorithm,omitempty"`
	// The Diffie-Hellman group.
	// For `AlwaysOn` VPN, the minimum allowed value is `14`.
	DiffieHellmanGroup *NetworkVPNIKEv2IKESecurityAssociationParametersDiffieHellmanGroup `default:"14" json:"DiffieHellmanGroup,omitempty"`
	// An array of integers representing postquantum key exchange methods the device uses during SA establishment and rekey. You can specify up to seven items, which correspond to ADDKE1 - ADDKE7 from RFC 9370.
	PostQuantumKeyExchangeMethods *[]NetworkVPNIKEv2IKESecurityAssociationParametersPostQuantumKeyExchangeMethods `json:"PostQuantumKeyExchangeMethods,omitempty"`
	// The SA lifetime (rekey interval) in minutes.
	LifeTimeInMinutes *int64 `default:"1440" json:"LifeTimeInMinutes,omitempty"`
}

// A dictionary to use for all VPN types.
type NetworkVPNIKEv2DNS struct {
	// The array of DNS server IP address strings. These IP addresses can be a mixture of IPv4 and IPv6 addresses.
	ServerAddresses []string `json:"ServerAddresses" required:"true"`
	// The list of domain strings used to fully qualify single-label host names.
	SearchDomains *[]string `json:"SearchDomains,omitempty"`
	// The primary domain of the tunnel.
	DomainName *string `json:"DomainName,omitempty"`
	// The list of domain strings used to determine which DNS queries use the DNS resolver settings in `ServerAddresses`. The system uses this key to create a split DNS configuration where it resolves only hosts in certain domains using the tunnel's DNS resolver. The system uses the default resolver for hosts that aren't in one of the domains in this list.
	// If `SupplementalMatchDomains` contains the empty string it becomes the default domain.
	// Split-tunnel configurations can direct all DNS queries to the VPN DNS servers before the primary DNS servers. If the VPN tunnel becomes the network's default route, the servers listed in `ServerAddresses` become the default resolver and the system ignores the `SupplementalMatchDomains` list.
	SupplementalMatchDomains *[]string `json:"SupplementalMatchDomains,omitempty"`
	// If `true`, don't append the domains in the `SupplementalMatchDomains` list to the resolver's list of search domains.
	SupplementalMatchDomainsNoSearch *bool `json:"SupplementalMatchDomainsNoSearch,omitempty"`
	// The identifier of a credential asset declaration that contains the identity that the system uses to authenticate the user to the DNS resolver.
	IdentityAssetReference *string `json:"IdentityAssetReference,omitempty"`
}

// The dictionary to use to configure `Proxies` for use with `VPN`.
type NetworkVPNIKEv2Proxies struct {
	// If `true`, enables automatic proxy configuration.
	AutoConfigEnable *bool `json:"AutoConfigEnable,omitempty"`
	// If `true`, enables proxy auto discovery.
	AutoDiscoveryEnable *bool `json:"AutoDiscoveryEnable,omitempty"`
	// The URL to the location of the proxy auto-configuration file. Used only when `ProxyAutoConfigEnable` is `true`.
	AutoConfigURLString *string `json:"AutoConfigURLString,omitempty"`
	// An array of domains that defines which hosts use proxy settings for hosts.
	SupplementalMatchDomains *[]string `json:"SupplementalMatchDomains,omitempty"`
	// The dictionary to use to configure HTTP servers  for `Proxies` for use with `VPN`.
	Protocol *NetworkVPNIKEv2ProxiesProtocol `json:"Protocol,omitempty"`
}

// The dictionary to use to configure HTTP servers  for `Proxies` for use with `VPN`.
type NetworkVPNIKEv2ProxiesProtocol struct {
	// The dictionary to use to configure the HTTP (non-TLS) server.
	HTTP *NetworkVPNIKEv2ProxiesProtocolHTTP `json:"HTTP,omitempty"`
	// The dictionary to use to configure the HTTPS (TLS) server.
	HTTPS *NetworkVPNIKEv2ProxiesProtocolHTTPS `json:"HTTPS,omitempty"`
	// The identifier of an asset declaration that contains the credentials (user name and password) to authenticate with the proxy server.
	CredentialsAssetReference *string `json:"CredentialsAssetReference,omitempty"`
}

// The dictionary to use to configure the HTTP (non-TLS) server.
type NetworkVPNIKEv2ProxiesProtocolHTTP struct {
	// If `true`, enables proxy for HTTP traffic.
	Enable *bool `json:"Enable,omitempty"`
	// The host name of the HTTP proxy.
	HostName *string `json:"HostName,omitempty"`
	// The port number of the HTTP proxy. This field is required if `HostName` is specified.
	Port *int64 `json:"Port,omitempty"`
}

// The dictionary to use to configure the HTTPS (TLS) server.
type NetworkVPNIKEv2ProxiesProtocolHTTPS struct {
	// If `true`, enables proxy for HTTPS traffic.
	Enable *bool `json:"Enable,omitempty"`
	// The host name of the HTTPS proxy.
	HostName *string `json:"HostName,omitempty"`
	// The port number of the HTTPS proxy. This field is required if `HostName` is specified.
	Port *int64 `json:"Port,omitempty"`
}

// The declaration to configure a VPN using the IPSec sub-type.
type NetworkVPNIPSec struct {
	// The name of the VPN connection that the system displays on the device.
	VisibleName string `json:"VisibleName" required:"true"`
	// The IP address or hostname of the VPN server.
	HostName string `json:"HostName" required:"true"`
	// Settings that control authentication.
	Authentication NetworkVPNIPSecAuthentication `json:"Authentication" required:"true"`
	// Specifies details about how the system handles idle VPN connections.
	Idle *NetworkVPNIPSecIdle `json:"Idle,omitempty"`
	// Specifies details about how the system controls on-demand VPN.
	OnDemand *NetworkVPNIPSecOnDemand `json:"OnDemand,omitempty"`
	// A dictionary to use for all VPN types.
	DNS *NetworkVPNIPSecDNS `json:"DNS,omitempty"`
	// The dictionary to use to configure `Proxies` for use with `VPN`.
	Proxies *NetworkVPNIPSecProxies `json:"Proxies,omitempty"`
}

func (p *NetworkVPNIPSec) DeclarationType() string {
	return "com.apple.configuration.network.vpn.ipsec"
}

// Settings that control authentication.
type NetworkVPNIPSecAuthentication struct {
	// The authentication method to use.
	Method NetworkVPNIPSecAuthenticationMethod `json:"Method" required:"true"`
	// The identifier of an asset declaration that contains the credentials
	// (password) to authenticate with the VPN servers.
	// Only use this with Cisco IPSec VPNs and if the `Authentication.Method` key is to `SharedSecret`.
	CredentialsAssetReference *string `json:"CredentialsAssetReference,omitempty"`
	// The identifier of a credential asset declaration that contains the identity
	// that this account requires to authenticate with the VPN servers.
	// Only use this with Cisco IPSec VPNs and if the `Authentication.Method` key is to `Certificate`.
	IdentityAssetReference *string `json:"IdentityAssetReference,omitempty"`
	// The name of the group. For hybrid authentication, the string needs to end with "hybrid".
	// Present only for Cisco IPSec if `Authentication.Method` is `SharedSecret`.
	LocalIdentifier *string `json:"LocalIdentifier,omitempty"`
	// Present only if `Authentication.Method` is `SharedSecret`. The value is `KeyID`. The system uses this value for Cisco IPSec VPNs.
	LocalIdentifierType *LocalIdentifierType `default:"KeyID" json:"LocalIdentifierType,omitempty"`
	// If `true`, prompts for a PIN when connecting to Cisco IPSec VPNs.
	PromptForVPNPIN *bool `json:"PromptForVPNPIN,omitempty"`
	// Settings that control XAuth.
	XAuth *XAuth `json:"XAuth,omitempty"`
}

// The authentication method to use.
type NetworkVPNIPSecAuthenticationMethod string

const (
	NetworkVPNIPSecAuthenticationMethodSharedSecret NetworkVPNIPSecAuthenticationMethod = "SharedSecret"
	NetworkVPNIPSecAuthenticationMethodCertificate  NetworkVPNIPSecAuthenticationMethod = "Certificate"
)

// Present only if `Authentication.Method` is `SharedSecret`. The value is `KeyID`. The system uses this value for Cisco IPSec VPNs.
type LocalIdentifierType string

const (
	LocalIdentifierTypeKeyID LocalIdentifierType = "KeyID"
)

// Settings that control XAuth.
type XAuth struct {
	// If `true`, enables Xauth for Cisco IPSec VPNs.
	Enabled bool `json:"Enabled" required:"true"`
	// The identifier of an asset declaration that contains the credentials (user name and password) required for XAuth. Required when `Enabled` key is set to `true`.
	CredentialsAssetReference *string `json:"CredentialsAssetReference,omitempty"`
	// A string that either has the value `Prompt` or isn't present.
	PasswordEncryption *PasswordEncryption `json:"PasswordEncryption,omitempty"`
}

// A string that either has the value `Prompt` or isn't present.
type PasswordEncryption string

const (
	PasswordEncryptionPrompt PasswordEncryption = "Prompt"
)

// Specifies details about how the system handles idle VPN connections.
type NetworkVPNIPSecIdle struct {
	// If `true`, disconnects after an on-demand connection idles.
	Disconnect *bool `json:"Disconnect,omitempty"`
	// The length of time to wait, in seconds, before disconnecting an on-demand connection.
	Timer *int64 `json:"Timer,omitempty"`
}

// Specifies details about how the system controls on-demand VPN.
type NetworkVPNIPSecOnDemand struct {
	// If `true`, enables VPN On Demand.
	Enabled *bool `json:"Enabled,omitempty"`
	// An array of dictionaries defining On Demand Rules.
	Rules *[]*NetworkVPNIPSecOnDemandRulesRulesElement `json:"Rules,omitempty"`
}
type NetworkVPNIPSecOnDemandRulesRulesElement struct {
	// The action to take if this dictionary matches the current network. Possible values are:
	// - `Allow`: Deprecated. Allow VPN On Demand to connect if triggered.
	// - `Connect`: Unconditionally initiate a VPN connection on the next network attempt.
	// - `Disconnect`: Tear down the VPN connection and don't reconnect on demand as long as this dictionary matches.
	// - `EvaluateConnection`: Evaluate the ActionParameters array for each connection attempt.
	// - `Ignore`: Leave any existing VPN connection up, but don't reconnect on demand as long as this dictionary matches.
	Action NetworkVPNIPSecOnDemandRulesRulesElementAction `json:"Action" required:"true"`
	// An array of dictionaries that provides rules similar to the `OnDemandRules` dictionary, but evaluated on each connection instead of when the network changes. This value is only for use with dictionaries in which the `Action` value is `EvaluateConnection`. The system evaluates these dictionaries in order and the first dictionary that matches determines the behavior.
	ActionParameters *[]*NetworkVPNIPSecOnDemandRulesRulesElementActionParametersActionParameter `json:"ActionParameters,omitempty"`
	// An array of domain names. This rule matches if any of the domain names in the specified list matches any domain in the device's search domains list.
	// The system supports a wildcard (`\*`) prefix. For example, `\*.example.com` matches against either `mydomain.example.com` or `yourdomain.example.com`.
	DNSDomainMatch *[]string `json:"DNSDomainMatch,omitempty"`
	// An array of IP addresses. This rule matches if any of the network's specified DNS servers match any entry in the array.
	// The system supports matching with a single wildcard. For example, `17.\*` matches any DNS server in the `17.0.0.0/8` subnet.
	DNSServerAddressMatch *[]string `json:"DNSServerAddressMatch,omitempty"`
	// An interface type. If specified, this rule matches only if the primary network interface hardware matches the specified type.
	InterfaceTypeMatch *NetworkVPNIPSecOnDemandRulesRulesElementInterfaceTypeMatch `json:"InterfaceTypeMatch,omitempty"`
	// An array of SSIDs to match against the current network. If the network isn't a Wi-Fi network or if the SSID doesn't appear in this array, the match fails.
	// Omit this key and the corresponding array to match against any SSID.
	SSIDMatch *[]string `json:"SSIDMatch,omitempty"`
	// A URL to probe. This rule matches when this URL is successfully fetched (returns a `200` HTTP status code) without redirection.
	URLStringProbe *string `json:"URLStringProbe,omitempty"`
}

// The action to take if this dictionary matches the current network. Possible values are:
// - `Allow`: Deprecated. Allow VPN On Demand to connect if triggered.
// - `Connect`: Unconditionally initiate a VPN connection on the next network attempt.
// - `Disconnect`: Tear down the VPN connection and don't reconnect on demand as long as this dictionary matches.
// - `EvaluateConnection`: Evaluate the ActionParameters array for each connection attempt.
// - `Ignore`: Leave any existing VPN connection up, but don't reconnect on demand as long as this dictionary matches.
type NetworkVPNIPSecOnDemandRulesRulesElementAction string

const (
	NetworkVPNIPSecOnDemandRulesRulesElementActionAllow              NetworkVPNIPSecOnDemandRulesRulesElementAction = "Allow"
	NetworkVPNIPSecOnDemandRulesRulesElementActionConnect            NetworkVPNIPSecOnDemandRulesRulesElementAction = "Connect"
	NetworkVPNIPSecOnDemandRulesRulesElementActionDisconnect         NetworkVPNIPSecOnDemandRulesRulesElementAction = "Disconnect"
	NetworkVPNIPSecOnDemandRulesRulesElementActionEvaluateConnection NetworkVPNIPSecOnDemandRulesRulesElementAction = "EvaluateConnection"
	NetworkVPNIPSecOnDemandRulesRulesElementActionIgnore             NetworkVPNIPSecOnDemandRulesRulesElementAction = "Ignore"
)

// A dictionary that provides rules similar to the OnDemandRules dictionary, but evaluated on each connection instead of when the network changes. These dictionaries are evaluated in order, and the behavior is determined by the first dictionary that matches.
// The keys allowed in each dictionary are described below. Note: This array is used only for dictionaries in which EvaluateConnection is the Action value.
type NetworkVPNIPSecOnDemandRulesRulesElementActionParametersActionParameter struct {
	// The domains to apply this evaluation.
	Domains []string `json:"Domains" required:"true"`
	// Defines the VPN behavior for the specified domains. Allowed values are:
	// * 'ConnectIfNeeded': The specified domains should trigger a VPN connection attempt if domain name resolution fails, such as when the DNS server indicates that it can't resolve the domain, responds with a redirection to a different server, or fails to respond (timeout).
	// * 'NeverConnect': The specified domains should never trigger a VPN connection attempt.
	DomainAction NetworkVPNIPSecOnDemandRulesRulesElementActionParametersActionParameterDomainAction `json:"DomainAction" required:"true"`
	// An array of IP addresses of DNS servers to use for resolving the specified domains. These servers don't need to be part of the device's current network configuration. If these DNS servers aren't reachable, the system establishes a VPN connection. These DNS servers need to be either internal DNS servers or trusted external DNS servers.
	// This key is valid only if the value of 'DomainAction' is 'ConnectIfNeeded'.
	RequiredDNSServers *[]string `json:"RequiredDNSServers,omitempty"`
	// An HTTP or HTTPS (preferred) URL to probe, using a GET request. If the URL's hostname can't be resolved, if the server is unreachable, or if the server doesn't respond with a 200 HTTP status code, a VPN connection is established in response.
	// This key is valid only if the value of 'DomainAction' is 'ConnectIfNeeded'.
	RequiredURLStringProbe *string `json:"RequiredURLStringProbe,omitempty"`
}

// Defines the VPN behavior for the specified domains. Allowed values are:
// * 'ConnectIfNeeded': The specified domains should trigger a VPN connection attempt if domain name resolution fails, such as when the DNS server indicates that it can't resolve the domain, responds with a redirection to a different server, or fails to respond (timeout).
// * 'NeverConnect': The specified domains should never trigger a VPN connection attempt.
type NetworkVPNIPSecOnDemandRulesRulesElementActionParametersActionParameterDomainAction string

const (
	NetworkVPNIPSecOnDemandRulesRulesElementActionParametersActionParameterDomainActionConnectIfNeeded NetworkVPNIPSecOnDemandRulesRulesElementActionParametersActionParameterDomainAction = "ConnectIfNeeded"
	NetworkVPNIPSecOnDemandRulesRulesElementActionParametersActionParameterDomainActionNeverConnect    NetworkVPNIPSecOnDemandRulesRulesElementActionParametersActionParameterDomainAction = "NeverConnect"
)

// An interface type. If specified, this rule matches only if the primary network interface hardware matches the specified type.
type NetworkVPNIPSecOnDemandRulesRulesElementInterfaceTypeMatch string

const (
	NetworkVPNIPSecOnDemandRulesRulesElementInterfaceTypeMatchEthernet NetworkVPNIPSecOnDemandRulesRulesElementInterfaceTypeMatch = "Ethernet"
	NetworkVPNIPSecOnDemandRulesRulesElementInterfaceTypeMatchWiFi     NetworkVPNIPSecOnDemandRulesRulesElementInterfaceTypeMatch = "WiFi"
	NetworkVPNIPSecOnDemandRulesRulesElementInterfaceTypeMatchCellular NetworkVPNIPSecOnDemandRulesRulesElementInterfaceTypeMatch = "Cellular"
)

// A dictionary to use for all VPN types.
type NetworkVPNIPSecDNS struct {
	// The array of DNS server IP address strings. These IP addresses can be a mixture of IPv4 and IPv6 addresses.
	ServerAddresses []string `json:"ServerAddresses" required:"true"`
	// The list of domain strings used to fully qualify single-label host names.
	SearchDomains *[]string `json:"SearchDomains,omitempty"`
	// The primary domain of the tunnel.
	DomainName *string `json:"DomainName,omitempty"`
	// The list of domain strings used to determine which DNS queries use the DNS resolver settings in `ServerAddresses`. The system uses this key to create a split DNS configuration where it resolves only hosts in certain domains using the tunnel's DNS resolver. The system uses the default resolver for hosts that aren't in one of the domains in this list.
	// If `SupplementalMatchDomains` contains the empty string it becomes the default domain.
	// Split-tunnel configurations can direct all DNS queries to the VPN DNS servers before the primary DNS servers. If the VPN tunnel becomes the network's default route, the servers listed in `ServerAddresses` become the default resolver and the system ignores the `SupplementalMatchDomains` list.
	SupplementalMatchDomains *[]string `json:"SupplementalMatchDomains,omitempty"`
	// If `true`, don't append the domains in the `SupplementalMatchDomains` list to the resolver's list of search domains.
	SupplementalMatchDomainsNoSearch *bool `json:"SupplementalMatchDomainsNoSearch,omitempty"`
	// The identifier of a credential asset declaration that contains the identity that the system uses to authenticate the user to the DNS resolver.
	IdentityAssetReference *string `json:"IdentityAssetReference,omitempty"`
}

// The dictionary to use to configure `Proxies` for use with `VPN`.
type NetworkVPNIPSecProxies struct {
	// If `true`, enables automatic proxy configuration.
	AutoConfigEnable *bool `json:"AutoConfigEnable,omitempty"`
	// If `true`, enables proxy auto discovery.
	AutoDiscoveryEnable *bool `json:"AutoDiscoveryEnable,omitempty"`
	// The URL to the location of the proxy auto-configuration file. Used only when `ProxyAutoConfigEnable` is `true`.
	AutoConfigURLString *string `json:"AutoConfigURLString,omitempty"`
	// An array of domains that defines which hosts use proxy settings for hosts.
	SupplementalMatchDomains *[]string `json:"SupplementalMatchDomains,omitempty"`
	// The dictionary to use to configure HTTP servers  for `Proxies` for use with `VPN`.
	Protocol *NetworkVPNIPSecProxiesProtocol `json:"Protocol,omitempty"`
}

// The dictionary to use to configure HTTP servers  for `Proxies` for use with `VPN`.
type NetworkVPNIPSecProxiesProtocol struct {
	// The dictionary to use to configure the HTTP (non-TLS) server.
	HTTP *NetworkVPNIPSecProxiesProtocolHTTP `json:"HTTP,omitempty"`
	// The dictionary to use to configure the HTTPS (TLS) server.
	HTTPS *NetworkVPNIPSecProxiesProtocolHTTPS `json:"HTTPS,omitempty"`
	// The identifier of an asset declaration that contains the credentials (user name and password) to authenticate with the proxy server.
	CredentialsAssetReference *string `json:"CredentialsAssetReference,omitempty"`
}

// The dictionary to use to configure the HTTP (non-TLS) server.
type NetworkVPNIPSecProxiesProtocolHTTP struct {
	// If `true`, enables proxy for HTTP traffic.
	Enable *bool `json:"Enable,omitempty"`
	// The host name of the HTTP proxy.
	HostName *string `json:"HostName,omitempty"`
	// The port number of the HTTP proxy. This field is required if `HostName` is specified.
	Port *int64 `json:"Port,omitempty"`
}

// The dictionary to use to configure the HTTPS (TLS) server.
type NetworkVPNIPSecProxiesProtocolHTTPS struct {
	// If `true`, enables proxy for HTTPS traffic.
	Enable *bool `json:"Enable,omitempty"`
	// The host name of the HTTPS proxy.
	HostName *string `json:"HostName,omitempty"`
	// The port number of the HTTPS proxy. This field is required if `HostName` is specified.
	Port *int64 `json:"Port,omitempty"`
}

// The declaration to configure a VPN using the VPN plugin sub-type.
type NetworkVPNVPNPlugin struct {
	// The name of the VPN connection that the system displays on the device.
	VisibleName string `json:"VisibleName" required:"true"`
	// The IP address or hostname of the VPN server.
	HostName string `json:"HostName" required:"true"`
	// An identifier for a vendor-specified configuration dictionary.
	// If the configuration targets a VPN solution that uses a VPN plugin, then this field contains the bundle identifier of the plugin. Here are some examples:
	// - Cisco AnyConnect: `com.cisco.anyconnect.applevpn.plugin`
	// - Juniper SSL: `net.juniper.sslvpn`
	// - F5 SSL: `com.f5.F5-Edge-Client.vpnplugin`
	// - SonicWALL Mobile Connect: `com.sonicwall.SonicWALL-SSLVPN.vpnplugin`
	// - ``Aruba VIA: `com.arubanetworks.aruba-via.vpnplugin`
	// If the configuration targets a VPN solution that uses a network extension provider, then this field contains the bundle identifier of the app that contains the provider. Contact the VPN solution vendor for the value of the identifier.
	SubType string `json:"SubType" required:"true"`
	// The vendor-specific configuration dictionary, which the system reads only when `SubType` has a value.
	VendorConfig *VendorConfig `json:"VendorConfig,omitempty"`
	// Settings that control authentication.
	Authentication NetworkVPNVPNPluginAuthentication `json:"Authentication" required:"true"`
	// Specifies details about the provider.
	Provider *Provider `json:"Provider,omitempty"`
	// Specifies details about how the VPN routes different types of network traffic.
	NetworkRouting *NetworkVPNVPNPluginNetworkRouting `json:"NetworkRouting,omitempty"`
	// Specifies details about how the system handles idle VPN connections.
	Idle *NetworkVPNVPNPluginIdle `json:"Idle,omitempty"`
	// Specifies details about how the system controls on-demand VPN.
	OnDemand *NetworkVPNVPNPluginOnDemand `json:"OnDemand,omitempty"`
	// A dictionary to use for all VPN types.
	DNS *NetworkVPNVPNPluginDNS `json:"DNS,omitempty"`
	// The dictionary to use to configure `Proxies` for use with `VPN`.
	Proxies *NetworkVPNVPNPluginProxies `json:"Proxies,omitempty"`
}

func (p *NetworkVPNVPNPlugin) DeclarationType() string {
	return "com.apple.configuration.network.vpn.vpn-plugin"
}

// The vendor-specific configuration dictionary, which the system reads only when `SubType` has a value.
type VendorConfig struct {
	// The Kerberos realm name, which needs to be properly capitalized. Valid only for Juniper SSL and Pulse Secure.
	Realm *string `json:"Realm,omitempty"`
	// The role to select when connecting to the server. Valid only for Juniper SSL and Pulse Secure.
	Role *string `json:"Role,omitempty"`
	// The group to connect to on the head end. Valid for Cisco AnyConnect and Cisco Legacy AnyConnect.
	Group *string `json:"Group,omitempty"`
	// The login group or domain. Valid only for SonicWALL Mobile Connect.
	LoginGroupOrDomain *string `json:"LoginGroupOrDomain,omitempty"`
}

// Settings that control authentication.
type NetworkVPNVPNPluginAuthentication struct {
	// The authentication method to use.
	Method NetworkVPNVPNPluginAuthenticationMethod `json:"Method" required:"true"`
	// The identifier of an asset declaration that contains the credentials (user name and password) to authenticate with the VPN server. Required when `Authentication.Method` is set to `Password`.
	CredentialsAssetReference *string `json:"CredentialsAssetReference,omitempty"`
	// The identifier of a credential asset declaration that contains the identity that this account requires to authenticate with the VPN server. Required when `Authentication.Method` is set to `Certificate`.
	IdentityAssetReference *string `json:"IdentityAssetReference,omitempty"`
}

// The authentication method to use.
type NetworkVPNVPNPluginAuthenticationMethod string

const (
	NetworkVPNVPNPluginAuthenticationMethodPassword            NetworkVPNVPNPluginAuthenticationMethod = "Password"
	NetworkVPNVPNPluginAuthenticationMethodCertificate         NetworkVPNVPNPluginAuthenticationMethod = "Certificate"
	NetworkVPNVPNPluginAuthenticationMethodPasswordCertificate NetworkVPNVPNPluginAuthenticationMethod = "Password+Certificate"
)

// Specifies details about the provider.
type Provider struct {
	// The type of VPN service. If the value is `app-proxy`, the service tunnels traffic at the app level. If the value is `packet-tunnel`, the service tunnels traffic at the IP layer.
	Type *ProviderType `default:"packet-tunnel" json:"Type,omitempty"`
	// If the `SubType` field contains the bundle identifier of an app that contains multiple VPN providers of the same type (app-proxy or packet-tunnel), then the system uses this field to choose which provider to use for this configuration. If the VPN provider is implemented as a System Extension, then this field is required.
	// In iOS, tvOS, and visionOS, the identifier is a bundle ID, for example, "com.example.app".
	// In macOS, the identifier is a composed identifier. The format of the composed identifier is either "Bundle-ID" or "Bundle-ID {Designated-Requirement}". "Bundle-ID" is the bundle identifier string of the provider. "Designated-Requirement" is the designated requirement string the device uses to match the code signature of the provider. For example, "com.example.app" for the bundle ID format, or "com.example.app {anchor apple generic}" for the designated requirement format.
	ComposedIdentifier *string `json:"ComposedIdentifier,omitempty"`
}

// The type of VPN service. If the value is `app-proxy`, the service tunnels traffic at the app level. If the value is `packet-tunnel`, the service tunnels traffic at the IP layer.
type ProviderType string

const (
	ProviderTypePacketTunnel ProviderType = "packet-tunnel"
	ProviderTypeAppProxy     ProviderType = "app-proxy"
)

// Specifies details about how the VPN routes different types of network traffic.
type NetworkVPNVPNPluginNetworkRouting struct {
	// If `true`, all the VPN's non-default routes take precedence over any locally defined routes.
	// If `IncludeAllNetworks` is `true`, the system ignores the value of `EnforceRoutes`.
	EnforceRoutes *bool `json:"EnforceRoutes,omitempty"`
	// If `true`, routes all traffic through the VPN, with some exclusions. Several of the exclusions can be controlled with the `ExcludeLocalNetworks`, `ExcludeCellularServices`, `ExcludeAPNs` and `ExcludeDeviceCommunication` properties. The following traffic is always excluded from the tunnel:
	// - Traffic necessary for connecting and maintaining the device's network connection, such as DHCP.
	// - Traffic necessary for connecting to captive networks.
	// - Certain cellular services traffic that's not routable over the internet and is instead directly routed to the cellular network. See the ExcludeCellularServices property for more details.
	IncludeAllNetworks *bool `json:"IncludeAllNetworks,omitempty"`
	// If `true` and `IncludeAllNetworks` is `true`, routes all local network traffic outside the VPN.
	ExcludeLocalNetworks *bool `json:"ExcludeLocalNetworks,omitempty"`
	// If `true` and `IncludeAllNetworks` is `true`, then the system excludes internet-routable network traffic for cellular services (VoLTE, Wi-Fi Calling, IMS, MMS, Visual Voicemail, etc.) from the tunnel. Note that some cellular carriers route cellular services traffic directly to the carrier network, bypassing the internet. Such cellular services traffic is always excluded from the tunnel.
	ExcludeCellularServices *bool `json:"ExcludeCellularServices,omitempty"`
	// If `true` and `IncludeAllNetworks` is `true`, then the system excludes the network traffic for the Apple Push Notification service (APNs) from the tunnel.
	ExcludeAPNs *bool `json:"ExcludeAPNs,omitempty"`
	// If set to `true` and `IncludeAllNetworks` is set to `true`, the device excludes network traffic used for communicating with devices connected via USB or Wi-Fi from the tunnel.
	ExcludeDeviceCommunication *bool `json:"ExcludeDeviceCommunication,omitempty"`
}

// Specifies details about how the system handles idle VPN connections.
type NetworkVPNVPNPluginIdle struct {
	// If `true`, disconnects after an on-demand connection idles.
	Disconnect *bool `json:"Disconnect,omitempty"`
	// The length of time to wait, in seconds, before disconnecting an on-demand connection.
	Timer *int64 `json:"Timer,omitempty"`
}

// Specifies details about how the system controls on-demand VPN.
type NetworkVPNVPNPluginOnDemand struct {
	// If `true`, enables VPN On Demand.
	Enabled *bool `json:"Enabled,omitempty"`
	// If `true`, the device disables the Connect On Demand toggle in Settings for this configuration.
	DisableUserOverride *bool `json:"DisableUserOverride,omitempty"`
	// An array of dictionaries defining On Demand Rules.
	Rules *[]*NetworkVPNVPNPluginOnDemandRulesRulesElement `json:"Rules,omitempty"`
}
type NetworkVPNVPNPluginOnDemandRulesRulesElement struct {
	// The action to take if this dictionary matches the current network. Possible values are:
	// - `Allow`: Deprecated. Allow VPN On Demand to connect if triggered.
	// - `Connect`: Unconditionally initiate a VPN connection on the next network attempt.
	// - `Disconnect`: Tear down the VPN connection and don't reconnect on demand as long as this dictionary matches.
	// - `EvaluateConnection`: Evaluate the ActionParameters array for each connection attempt.
	// - `Ignore`: Leave any existing VPN connection up, but don't reconnect on demand as long as this dictionary matches.
	Action NetworkVPNVPNPluginOnDemandRulesRulesElementAction `json:"Action" required:"true"`
	// An array of dictionaries that provides rules similar to the `OnDemandRules` dictionary, but evaluated on each connection instead of when the network changes. This value is only for use with dictionaries in which the `Action` value is `EvaluateConnection`. The system evaluates these dictionaries in order and the first dictionary that matches determines the behavior.
	ActionParameters *[]*NetworkVPNVPNPluginOnDemandRulesRulesElementActionParametersActionParameter `json:"ActionParameters,omitempty"`
	// An array of domain names. This rule matches if any of the domain names in the specified list matches any domain in the device's search domains list.
	// The system supports a wildcard (`\*`) prefix. For example, `\*.example.com` matches against either `mydomain.example.com` or `yourdomain.example.com`.
	DNSDomainMatch *[]string `json:"DNSDomainMatch,omitempty"`
	// An array of IP addresses. This rule matches if any of the network's specified DNS servers match any entry in the array.
	// The system supports matching with a single wildcard. For example, `17.\*` matches any DNS server in the `17.0.0.0/8` subnet.
	DNSServerAddressMatch *[]string `json:"DNSServerAddressMatch,omitempty"`
	// An interface type. If specified, this rule matches only if the primary network interface hardware matches the specified type.
	InterfaceTypeMatch *NetworkVPNVPNPluginOnDemandRulesRulesElementInterfaceTypeMatch `json:"InterfaceTypeMatch,omitempty"`
	// An array of SSIDs to match against the current network. If the network isn't a Wi-Fi network or if the SSID doesn't appear in this array, the match fails.
	// Omit this key and the corresponding array to match against any SSID.
	SSIDMatch *[]string `json:"SSIDMatch,omitempty"`
	// A URL to probe. This rule matches when this URL is successfully fetched (returns a `200` HTTP status code) without redirection.
	URLStringProbe *string `json:"URLStringProbe,omitempty"`
}

// The action to take if this dictionary matches the current network. Possible values are:
// - `Allow`: Deprecated. Allow VPN On Demand to connect if triggered.
// - `Connect`: Unconditionally initiate a VPN connection on the next network attempt.
// - `Disconnect`: Tear down the VPN connection and don't reconnect on demand as long as this dictionary matches.
// - `EvaluateConnection`: Evaluate the ActionParameters array for each connection attempt.
// - `Ignore`: Leave any existing VPN connection up, but don't reconnect on demand as long as this dictionary matches.
type NetworkVPNVPNPluginOnDemandRulesRulesElementAction string

const (
	NetworkVPNVPNPluginOnDemandRulesRulesElementActionAllow              NetworkVPNVPNPluginOnDemandRulesRulesElementAction = "Allow"
	NetworkVPNVPNPluginOnDemandRulesRulesElementActionConnect            NetworkVPNVPNPluginOnDemandRulesRulesElementAction = "Connect"
	NetworkVPNVPNPluginOnDemandRulesRulesElementActionDisconnect         NetworkVPNVPNPluginOnDemandRulesRulesElementAction = "Disconnect"
	NetworkVPNVPNPluginOnDemandRulesRulesElementActionEvaluateConnection NetworkVPNVPNPluginOnDemandRulesRulesElementAction = "EvaluateConnection"
	NetworkVPNVPNPluginOnDemandRulesRulesElementActionIgnore             NetworkVPNVPNPluginOnDemandRulesRulesElementAction = "Ignore"
)

// A dictionary that provides rules similar to the OnDemandRules dictionary, but evaluated on each connection instead of when the network changes. These dictionaries are evaluated in order, and the behavior is determined by the first dictionary that matches.
// The keys allowed in each dictionary are described below. Note: This array is used only for dictionaries in which EvaluateConnection is the Action value.
type NetworkVPNVPNPluginOnDemandRulesRulesElementActionParametersActionParameter struct {
	// The domains to apply this evaluation.
	Domains []string `json:"Domains" required:"true"`
	// Defines the VPN behavior for the specified domains. Allowed values are:
	// * 'ConnectIfNeeded': The specified domains should trigger a VPN connection attempt if domain name resolution fails, such as when the DNS server indicates that it can't resolve the domain, responds with a redirection to a different server, or fails to respond (timeout).
	// * 'NeverConnect': The specified domains should never trigger a VPN connection attempt.
	DomainAction NetworkVPNVPNPluginOnDemandRulesRulesElementActionParametersActionParameterDomainAction `json:"DomainAction" required:"true"`
	// An array of IP addresses of DNS servers to use for resolving the specified domains. These servers don't need to be part of the device's current network configuration. If these DNS servers aren't reachable, the system establishes a VPN connection. These DNS servers need to be either internal DNS servers or trusted external DNS servers.
	// This key is valid only if the value of 'DomainAction' is 'ConnectIfNeeded'.
	RequiredDNSServers *[]string `json:"RequiredDNSServers,omitempty"`
	// An HTTP or HTTPS (preferred) URL to probe, using a GET request. If the URL's hostname can't be resolved, if the server is unreachable, or if the server doesn't respond with a 200 HTTP status code, a VPN connection is established in response.
	// This key is valid only if the value of 'DomainAction' is 'ConnectIfNeeded'.
	RequiredURLStringProbe *string `json:"RequiredURLStringProbe,omitempty"`
}

// Defines the VPN behavior for the specified domains. Allowed values are:
// * 'ConnectIfNeeded': The specified domains should trigger a VPN connection attempt if domain name resolution fails, such as when the DNS server indicates that it can't resolve the domain, responds with a redirection to a different server, or fails to respond (timeout).
// * 'NeverConnect': The specified domains should never trigger a VPN connection attempt.
type NetworkVPNVPNPluginOnDemandRulesRulesElementActionParametersActionParameterDomainAction string

const (
	NetworkVPNVPNPluginOnDemandRulesRulesElementActionParametersActionParameterDomainActionConnectIfNeeded NetworkVPNVPNPluginOnDemandRulesRulesElementActionParametersActionParameterDomainAction = "ConnectIfNeeded"
	NetworkVPNVPNPluginOnDemandRulesRulesElementActionParametersActionParameterDomainActionNeverConnect    NetworkVPNVPNPluginOnDemandRulesRulesElementActionParametersActionParameterDomainAction = "NeverConnect"
)

// An interface type. If specified, this rule matches only if the primary network interface hardware matches the specified type.
type NetworkVPNVPNPluginOnDemandRulesRulesElementInterfaceTypeMatch string

const (
	NetworkVPNVPNPluginOnDemandRulesRulesElementInterfaceTypeMatchEthernet NetworkVPNVPNPluginOnDemandRulesRulesElementInterfaceTypeMatch = "Ethernet"
	NetworkVPNVPNPluginOnDemandRulesRulesElementInterfaceTypeMatchWiFi     NetworkVPNVPNPluginOnDemandRulesRulesElementInterfaceTypeMatch = "WiFi"
	NetworkVPNVPNPluginOnDemandRulesRulesElementInterfaceTypeMatchCellular NetworkVPNVPNPluginOnDemandRulesRulesElementInterfaceTypeMatch = "Cellular"
)

// A dictionary to use for all VPN types.
type NetworkVPNVPNPluginDNS struct {
	// The array of DNS server IP address strings. These IP addresses can be a mixture of IPv4 and IPv6 addresses.
	ServerAddresses []string `json:"ServerAddresses" required:"true"`
	// The list of domain strings used to fully qualify single-label host names.
	SearchDomains *[]string `json:"SearchDomains,omitempty"`
	// The primary domain of the tunnel.
	DomainName *string `json:"DomainName,omitempty"`
	// The list of domain strings used to determine which DNS queries use the DNS resolver settings in `ServerAddresses`. The system uses this key to create a split DNS configuration where it resolves only hosts in certain domains using the tunnel's DNS resolver. The system uses the default resolver for hosts that aren't in one of the domains in this list.
	// If `SupplementalMatchDomains` contains the empty string it becomes the default domain.
	// Split-tunnel configurations can direct all DNS queries to the VPN DNS servers before the primary DNS servers. If the VPN tunnel becomes the network's default route, the servers listed in `ServerAddresses` become the default resolver and the system ignores the `SupplementalMatchDomains` list.
	SupplementalMatchDomains *[]string `json:"SupplementalMatchDomains,omitempty"`
	// If `true`, don't append the domains in the `SupplementalMatchDomains` list to the resolver's list of search domains.
	SupplementalMatchDomainsNoSearch *bool `json:"SupplementalMatchDomainsNoSearch,omitempty"`
	// The identifier of a credential asset declaration that contains the identity that the system uses to authenticate the user to the DNS resolver.
	IdentityAssetReference *string `json:"IdentityAssetReference,omitempty"`
}

// The dictionary to use to configure `Proxies` for use with `VPN`.
type NetworkVPNVPNPluginProxies struct {
	// If `true`, enables automatic proxy configuration.
	AutoConfigEnable *bool `json:"AutoConfigEnable,omitempty"`
	// If `true`, enables proxy auto discovery.
	AutoDiscoveryEnable *bool `json:"AutoDiscoveryEnable,omitempty"`
	// The URL to the location of the proxy auto-configuration file. Used only when `ProxyAutoConfigEnable` is `true`.
	AutoConfigURLString *string `json:"AutoConfigURLString,omitempty"`
	// An array of domains that defines which hosts use proxy settings for hosts.
	SupplementalMatchDomains *[]string `json:"SupplementalMatchDomains,omitempty"`
	// The dictionary to use to configure HTTP servers  for `Proxies` for use with `VPN`.
	Protocol *NetworkVPNVPNPluginProxiesProtocol `json:"Protocol,omitempty"`
}

// The dictionary to use to configure HTTP servers  for `Proxies` for use with `VPN`.
type NetworkVPNVPNPluginProxiesProtocol struct {
	// The dictionary to use to configure the HTTP (non-TLS) server.
	HTTP *NetworkVPNVPNPluginProxiesProtocolHTTP `json:"HTTP,omitempty"`
	// The dictionary to use to configure the HTTPS (TLS) server.
	HTTPS *NetworkVPNVPNPluginProxiesProtocolHTTPS `json:"HTTPS,omitempty"`
	// The identifier of an asset declaration that contains the credentials (user name and password) to authenticate with the proxy server.
	CredentialsAssetReference *string `json:"CredentialsAssetReference,omitempty"`
}

// The dictionary to use to configure the HTTP (non-TLS) server.
type NetworkVPNVPNPluginProxiesProtocolHTTP struct {
	// If `true`, enables proxy for HTTP traffic.
	Enable *bool `json:"Enable,omitempty"`
	// The host name of the HTTP proxy.
	HostName *string `json:"HostName,omitempty"`
	// The port number of the HTTP proxy. This field is required if `HostName` is specified.
	Port *int64 `json:"Port,omitempty"`
}

// The dictionary to use to configure the HTTPS (TLS) server.
type NetworkVPNVPNPluginProxiesProtocolHTTPS struct {
	// If `true`, enables proxy for HTTPS traffic.
	Enable *bool `json:"Enable,omitempty"`
	// The host name of the HTTPS proxy.
	HostName *string `json:"HostName,omitempty"`
	// The port number of the HTTPS proxy. This field is required if `HostName` is specified.
	Port *int64 `json:"Port,omitempty"`
}

// The declaration to configure a package.
type Package struct {
	// The URL of the manifest document for the package that the device downloads. The manifest is returned as a `ManifestURL` property list. The `url` property of the manifest must point to the package (.pkg) file to install.
	ManifestURL string `json:"ManifestURL" required:"true"`
	// A dictionary that describes how and when to install the package.
	InstallBehavior *PackageInstallBehavior `json:"InstallBehavior,omitempty"`
	// A dictionary that describes how to uninstall the package.
	UninstallBehavior *UninstallBehavior `json:"UninstallBehavior,omitempty"`
}

func (p *Package) DeclarationType() string {
	return "com.apple.configuration.package"
}

// A dictionary that describes how and when to install the package.
type PackageInstallBehavior struct {
	// A string that specifies when the system installs the package:
	// - `Optional`: The user can install the package after the system activates the configuration.
	// - `Required`: The system installs the package after it activates the configuration.
	Install *PackageInstallBehaviorInstall `default:"Optional" json:"Install,omitempty"`
}

// A string that specifies when the system installs the package:
// - `Optional`: The user can install the package after the system activates the configuration.
// - `Required`: The system installs the package after it activates the configuration.
type PackageInstallBehaviorInstall string

const (
	PackageInstallBehaviorInstallOptional PackageInstallBehaviorInstall = "Optional"
	PackageInstallBehaviorInstallRequired PackageInstallBehaviorInstall = "Required"
)

// A dictionary that describes how to uninstall the package.
type UninstallBehavior struct {
	// If `true`, the system removes the files that the package installs when removing the configuration.
	Remove *bool `json:"Remove,omitempty"`
}

// The declaration to configure passcode policy settings.
type PasscodeSettings struct {
	// If `true`, the system requires the user to set a passcode without any requirements about the length or quality of the passcode. The presence of any other keys implicitly requires a passcode, and overrides this key's value.
	RequirePasscode *bool `json:"RequirePasscode,omitempty"`
	// If `true`, the passcode needs to consist of at least one alphabetic character and at least one number.
	RequireAlphanumericPasscode *bool `json:"RequireAlphanumericPasscode,omitempty"`
	// If `true`, the system requires a complex passcode. A complex passcode is one that doesn't contain repeated characters or increasing or decreasing characters (such as 123 or CBA).
	RequireComplexPasscode *bool `json:"RequireComplexPasscode,omitempty"`
	// The minimum number of characters a passcode can contain.
	MinimumLength *int64 `default:"0" json:"MinimumLength,omitempty"`
	// Specifies the minimum number of complex characters in the password. A complex character is a character other than a number or a letter, such as `&`, `%`, `$`, and `#`.
	MinimumComplexCharacters *int64 `default:"0" json:"MinimumComplexCharacters,omitempty"`
	// The number of failed passcode attempts that the system allows the user before it erases or locks the device. After six failed attempts, the device imposes a time delay before the user can enter a passcode again. The time delay increases with each failed attempt. On macOS, set `FailedAttemptsResetInMinutes` to define the time delay. The time delay begins after the sixth attempt, so if `MaximumFailedAttempts` is six or lower, the system has no time delay and triggers the erase or lock as soon as the user exceeds the limit.
	// After the final failed attempt, the system locks a macOS device, or securely erases all data and settings from an iOS, visionOS, or watchOS device.
	MaximumFailedAttempts *int64 `default:"11" json:"MaximumFailedAttempts,omitempty"`
	// The number of minutes before the login is reset after the maximum number of failed attempts. Also set the `MaximumFailedAttempts` key for this to take effect.
	FailedAttemptsResetInMinutes *int64 `json:"FailedAttemptsResetInMinutes,omitempty"`
	// The maximum period that a user can select, during which the user can unlock the device without a passcode. A value of `0` means no grace period, and the device requires a passcode immediately. In the absence of this key, the user can select any period. In macOS, the system translates this to screensaver settings.
	MaximumGracePeriodInMinutes *int64 `json:"MaximumGracePeriodInMinutes,omitempty"`
	// The maximum period that a user can select, during which the device can be idle before the system automatically locks it. When the device reaches this limit, the device locks and the user must enter the passcode to unlock it. In the absence of this key, the user can select any period. In macOS, the system translates this to screensaver settings.
	MaximumInactivityInMinutes *int64 `json:"MaximumInactivityInMinutes,omitempty"`
	// Specifies the maximum number of days that the passcode can remain unchanged. After this number of days, the system forces the user to change the passcode before it unlocks the device.
	MaximumPasscodeAgeInDays *int64 `json:"MaximumPasscodeAgeInDays,omitempty"`
	// The number of historical passcode entries the system checks when validating a new passcode. The device refuses a new passcode if it matches a previously used passcode within the specified passcode history range. In the absence of this key, the system performs no historical check.
	PasscodeReuseLimit *int64 `json:"PasscodeReuseLimit,omitempty"`
	// If `true`, the system forces a password reset the next time the user tries to authenticate. If you set this key in a configuration in the system scope (device channel), the setting takes effect for all users, and admin authentication may fail until the admin user password is also reset.
	ChangeAtNextAuth *bool `json:"ChangeAtNextAuth,omitempty"`
	// Specifies a regular expression, and its description, to enforce password compliance. Use the simpler passcode settings whenever possible, and rely on regular expression matching only when necessary. Mistakes in regular expressions can lead to frustrating user experiences, such as unsatisfiable passcode policies, or policy descriptions that don't match the enforced policy.
	CustomRegex *CustomRegex `json:"CustomRegex,omitempty"`
}

func (p *PasscodeSettings) DeclarationType() string {
	return "com.apple.configuration.passcode.settings"
}

// Specifies a regular expression, and its description, to enforce password compliance. Use the simpler passcode settings whenever possible, and rely on regular expression matching only when necessary. Mistakes in regular expressions can lead to frustrating user experiences, such as unsatisfiable passcode policies, or policy descriptions that don't match the enforced policy.
type CustomRegex struct {
	// A regular expression string to match against the password to determine whether it complies with a policy. The regular expression uses the ICU syntax. The string can't exceed 2048 characters in length.
	Regex string `json:"Regex" required:"true"`
	// A dictionary with supported OS language IDs for the keys (such as `en-US`), and values that represent a localized description of the policy that the regular expression enforces. Use the special `default` key for languages that the dictionary doesn't contain.
	Description *map[string]string `json:"Description,omitempty"`
}

// The declaration to configure managed bookmarks in Safari.
type SafariBookmarks struct {
	// A dictionary that specifies a set of managed bookmarks.
	ManagedBookmarks *[]BookmarkGroup `json:"ManagedBookmarks,omitempty"`
}

func (p *SafariBookmarks) DeclarationType() string {
	return "com.apple.configuration.safari.bookmarks"
}

// A group of managed bookmarks.
type BookmarkGroup struct {
	// A string that specifies the unique identifier for this group of managed bookmarks. Safari displays a separate managed bookmarks item for each set of unique managed bookmarks based on the value of this key. If multiple active configurations use the same value for this key, Safari displays a single group formed by merging the list of `Bookmarks` from each group.
	GroupIdentifier string `json:"GroupIdentifier" required:"true"`
	// The name of the bookmarks folder. Safari uses this as the title for the top-level managed bookmarks item.
	Title string `json:"Title" required:"true"`
	// A list of bookmarks. Either a `URL` or `Folder` key must be present in each item.
	Bookmarks []BookmarksItem `json:"Bookmarks" required:"true"`
}

// A bookmark that specifies a title, and either a URL for the bookmark, or a nested folder of bookmarks.
type BookmarksItem struct {
	// The title of the bookmark shown in Safari.
	Title string `json:"Title" required:"true"`
	// The URL for the bookmark item.
	// Only one of `URL` or `Folder` must be present.
	URL *string `json:"URL,omitempty"`
	// An array of bookmarks for each bookmark in the folder. Folders can include bookmark items and bookmark folders.
	// Only one of `URL` or `Folder` must be present.
	Folder *[]FolderItem `json:"Folder,omitempty"`
}

// A bookmark that specifies a title, and either a URL for the bookmark, or a nested folder of bookmarks.
type FolderItem struct {
	// The title of the bookmark shown in Safari.
	Title string `json:"Title" required:"true"`
	// The URL for the bookmark item.
	// Only one of `URL` or `Folder` must be present.
	URL *string `json:"URL,omitempty"`
	// An array of bookmarks for each bookmark in the folder. Folders can include bookmark items and bookmark folders.
	// Only one of `URL` or `Folder` must be present.
	Folder *[]FolderItem `json:"Folder,omitempty"`
}

// The declaration to configure Safari Extensions.
type SafariExtensionSettings struct {
	// The dictionary of managed extensions settings. Each key in the dictionary represents a composed identifier for a specific managed extension, or you can specify a single "\*" character to match any extension. The dictionary values represent the settings that Safari applies to each extension that matches the key. In order for the extension to be managed, its host app needs to be present on the device.
	// The composed identifier of a managed extension uses the format "Identifier (TeamIdentifier)", for example "com.example.app (ABCD1234)". Use `codesign -dv <path_to_appex>` to show the information you need to generate this string on macOS, using the path to the extension bundle located in the "PlugIns" folder inside the app bundle. For other platforms, request this information from the app developer.
	ManagedExtensions *map[string]ExtensionDictionary `json:"ManagedExtensions,omitempty"`
}

func (p *SafariExtensionSettings) DeclarationType() string {
	return "com.apple.configuration.safari.extensions.settings"
}

// The dictionary that defines the settings for a managed extension. Each key represents a specific managed extension, or you can specify a single "*" character to match any extension.
type ExtensionDictionary struct {
	// Controls whether an extension is allowed. The device uses this key when the extension identifier is a composed identifier or a single "\*" character.
	// * `Allowed` - The user is allowed to turn the extension on or off.
	// * `AlwaysOn` - The extension will always be on.
	// * `AlwaysOff` - The extension will always be off.
	State *State `json:"State,omitempty"`
	// Controls whether an extension is allowed in Private Browsing. The device uses this key when the extension identifier is a composed identifier or a single "\*" character.
	// * `Allowed` - The user is allowed to turn the extension on or off in Private Browsing.
	// * `AlwaysOn` - The extension will always be on in Private Browsing if the extension is on outside of Private Browsing.
	// * `AlwaysOff` - The extension will never be on in Private Browsing.
	PrivateBrowsing *PrivateBrowsing `json:"PrivateBrowsing,omitempty"`
	// Controls the domains and sub-domains the extension can access. The device ignores this key when the extension identifier is a single "*" character.
	AllowedDomains *[]string `json:"AllowedDomains,omitempty"`
	// Controls the domains and sub-domains the extension isn't allowed to access. The device uses this key when the extension identifier is a composed identifier or a single "*" character.
	DeniedDomains *[]string `json:"DeniedDomains,omitempty"`
}

// Controls whether an extension is allowed. The device uses this key when the extension identifier is a composed identifier or a single "\*" character.
// * `Allowed` - The user is allowed to turn the extension on or off.
// * `AlwaysOn` - The extension will always be on.
// * `AlwaysOff` - The extension will always be off.
type State string

const (
	StateAllowed   State = "Allowed"
	StateAlwaysOn  State = "AlwaysOn"
	StateAlwaysOff State = "AlwaysOff"
)

// Controls whether an extension is allowed in Private Browsing. The device uses this key when the extension identifier is a composed identifier or a single "\*" character.
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
	AcceptCookies *AcceptCookies `default:"Always" json:"AcceptCookies,omitempty"`
	// If `false`, the system forces fraud warnings on in Safari.
	AllowDisablingFraudWarning *bool `json:"AllowDisablingFraudWarning,omitempty"`
	// If `false`, the system disables clearing history in Safari.
	AllowHistoryClearing *bool `json:"AllowHistoryClearing,omitempty"`
	// If `false`, the system disables JavaScript in Safari.
	AllowJavaScript *bool `json:"AllowJavaScript,omitempty"`
	// If `false`, the system disables private browsing in Safari.
	AllowPrivateBrowsing *bool `json:"AllowPrivateBrowsing,omitempty"`
	// If `false`, the system disables popups in Safari.
	AllowPopups *bool `json:"AllowPopups,omitempty"`
	// If `false`, the system disables summarization of content in Safari.
	AllowSummary *bool `json:"AllowSummary,omitempty"`
	// Sets the start page for new tabs in Safari.
	NewTabStartPage *NewTabStartPage `json:"NewTabStartPage,omitempty"`
	// The dictionary of website privacy settings.
	Privacy *SafariSettingsPrivacy `json:"Privacy,omitempty"`
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
	PageType PageType `json:"PageType" required:"true"`
	// The URL of the homepage which needs to start with `https://` or `http://`. Required when setting `PageType` to `Home`.
	HomepageURL *string `json:"HomepageURL,omitempty"`
	// The composed identifier of the extension that provides the start page. The required format is "Identifier (TeamIdentifier)", for example "com.example.app (ABCD1234)". Required when setting `PageType` to `Extension`.
	ExtensionIdentifier *string `json:"ExtensionIdentifier,omitempty"`
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

// The dictionary of website privacy settings.
type SafariSettingsPrivacy struct {
	// The dictionary of website permission defaults. Each key in the dictionary represents a single website, or a website and its sub-domains. The dictionary values represent the permission defaults that Safari applies for each website that matches the key.
	// Safari supports the following patterns for the website key:
	// - A specific domain such as "example.com" or "www.example.com". The permission defaults apply to that website only.
	// - A wildcard domain that uses a single "\*" character as a prefix for the domain, such as "\*example.com". The permission defaults apply to both the exact domain "example.com", and any sub-domains such as "www.example.com". It won't match other domains with a similar string suffix such as "myexample.com".
	// When multiple patterns match a website, Safari uses the most precise pattern. For example, for the website "www.example.com", if rules "www.example.com" and "\*example.com" match, Safari uses the former.
	PermissionDefaults *map[string]WebsiteDictionary `json:"PermissionDefaults,omitempty"`
}

// The dictionary that defines the website privacy permission defaults. Each key represents a website.
type WebsiteDictionary struct {
	// Text that clearly explains to the Safari user the reason why the organization requires these website privacy permission defaults. Safari includes this text in the permission consent prompt it displays when it first displays the website.
	OrganizationJustification string `json:"OrganizationJustification" required:"true"`
	// Controls whether a website privacy permission default is set.
	// * `None`: Safari sets no website privacy permission default for use of the camera.
	// * `Allow`: Safari sets the website privacy permission default to allow use of the camera.
	Camera *WebsiteDictionaryCamera `json:"Camera,omitempty"`
	// Controls whether a website privacy permission default is set.
	// * `None`: Safari sets no website privacy permission default for use of the microphone.
	// * `Allow`: Safari sets the website privacy permission default to allow use of the microphone.
	Microphone *WebsiteDictionaryMicrophone `json:"Microphone,omitempty"`
}

// Controls whether a website privacy permission default is set.
// * `None`: Safari sets no website privacy permission default for use of the camera.
// * `Allow`: Safari sets the website privacy permission default to allow use of the camera.
type WebsiteDictionaryCamera string

const (
	WebsiteDictionaryCameraNone  WebsiteDictionaryCamera = "None"
	WebsiteDictionaryCameraAllow WebsiteDictionaryCamera = "Allow"
)

// Controls whether a website privacy permission default is set.
// * `None`: Safari sets no website privacy permission default for use of the microphone.
// * `Allow`: Safari sets the website privacy permission default to allow use of the microphone.
type WebsiteDictionaryMicrophone string

const (
	WebsiteDictionaryMicrophoneNone  WebsiteDictionaryMicrophone = "None"
	WebsiteDictionaryMicrophoneAllow WebsiteDictionaryMicrophone = "Allow"
)

// The declaration to configure a group of screen-sharing connections.
type ScreenSharingConnectionGroup struct {
	// A unique identifier for this connection group.
	ConnectionGroupUUID string `json:"ConnectionGroupUUID" required:"true"`
	// The name of the connection group.
	GroupName string `json:"GroupName" required:"true"`
	// An array of `ConnectionUUID`s that represent connections declared in `ScreenSharingConnection` configurations that are members of this group.
	Members []string `json:"Members" required:"true"`
}

func (p *ScreenSharingConnectionGroup) DeclarationType() string {
	return "com.apple.configuration.screensharing.connection.group"
}

// The declaration to configure a connection to a screen-sharing host.
type ScreenSharingConnection struct {
	// A unique identifier for this connection when it's in a connection group.
	ConnectionUUID string `json:"ConnectionUUID" required:"true"`
	// The name of the connection.
	DisplayName string `json:"DisplayName" required:"true"`
	// The host name or IP address of the Mac that hosts the screen-sharing connection.
	HostName string `json:"HostName" required:"true"`
	// The TCP port number on the host to initiate the connection.
	Port *int64 `json:"Port,omitempty"`
	// The display configuration for this connection.
	DisplayConfiguration DisplayConfiguration `json:"DisplayConfiguration" required:"true"`
	// The identifier of an asset declaration that contains the required credentials for this connection to authenticate with the screen-sharing server. Set the corresponding asset type to `com.apple.asset.credential.userpassword`.
	AuthenticationCredentialsAssetReference *string `json:"AuthenticationCredentialsAssetReference,omitempty"`
}

func (p *ScreenSharingConnection) DeclarationType() string {
	return "com.apple.configuration.screensharing.connection"
}

// The display configuration for this connection.
type DisplayConfiguration struct {
	// The type of display for the connection, which has these allowed values:
	// - `Virtual1`: Create one virtual display.
	// - `Virtual2`: Create two virtual displays.
	DisplayType DisplayType `json:"DisplayType" required:"true"`
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
	MaximumVirtualDisplays *int64 `json:"MaximumVirtualDisplays,omitempty"`
	// The initial UDP port number to connect to the host. Screen sharing requires multiple connections, so the system increments this value by 1 for each additional connection. This doesn't change the port number that the system uses to initially establish a connection with a host, which is always TCP port 5900.
	PortBase *int64 `json:"PortBase,omitempty"`
	// If `true`, the system prevents users from copying files from the screen-sharing host.
	PreventCopyFilesFromHost *bool `json:"PreventCopyFilesFromHost,omitempty"`
	// If `true`, the system prevents users from copying files to the screen-sharing host.
	PreventCopyFilesToHost *bool `json:"PreventCopyFilesToHost,omitempty"`
	// If `true`, the system prevents clients from establishing high-performance connections to the host.
	PreventHighPerformanceConnections *bool `json:"PreventHighPerformanceConnections,omitempty"`
}

func (p *ScreenSharingHostSettings) DeclarationType() string {
	return "com.apple.configuration.screensharing.host.settings"
}

// The declaration to configure a certificate.
type SecurityCertificate struct {
	// The identifier of an asset declaration that contains the certificate to install.
	CredentialAssetReference string `json:"CredentialAssetReference" required:"true"`
}

func (p *SecurityCertificate) DeclarationType() string {
	return "com.apple.configuration.security.certificate"
}

// The declaration to configure an identity.
type SecurityIdentity struct {
	// The identifier of an asset declaration that contains the identity to install.
	CredentialAssetReference string `json:"CredentialAssetReference" required:"true"`
	// If `true`, apps can access the private key.
	AllowAllAppsAccess *bool `json:"AllowAllAppsAccess,omitempty"`
	// If `true`, the private key is extractable in the keychain.
	KeyIsExtractable *bool `json:"KeyIsExtractable,omitempty"`
}

func (p *SecurityIdentity) DeclarationType() string {
	return "com.apple.configuration.security.identity"
}

// The declaration to configure the device to allow WebAuthn enterprise attestation for certain passkeys.
type SecurityPasskeyAttestation struct {
	// The identifier of an asset declaration that contains the identity to install and use for passkey attestation.
	AttestationIdentityAssetReference string `json:"AttestationIdentityAssetReference" required:"true"`
	// If `true`, the private key for the attestation identity is extractable in the keychain.
	AttestationIdentityKeyIsExtractable *bool `json:"AttestationIdentityKeyIsExtractable,omitempty"`
	// An array of the relying parties to allow enterprise attestation.
	RelyingParties []string `json:"RelyingParties" required:"true"`
}

func (p *SecurityPasskeyAttestation) DeclarationType() string {
	return "com.apple.configuration.security.passkey.attestation"
}

// The declaration to configure background tasks.
type ServicesBackgroundTasks struct {
	// The unique identifier of the set of background tasks managed with this configuration. This should be a reverse DNS style identifier. The device uses this identifier to differentiate between tasks in different configurations.
	TaskType string `json:"TaskType" required:"true"`
	// A description of the set of background tasks this configuration manages.
	TaskDescription *string `json:"TaskDescription,omitempty"`
	// Specifies the identifier of an asset declaration containing a reference to the files to be used for the background task configuration. The corresponding asset must be of type `com.apple.asset.data`.
	// The referenced data must be a zip archive of an entire directory that the device expands and stores in a well known location for the background task. The asset's "ContentType" and "Hash-SHA-256" keys in the "Reference" key are required.
	// This zip archive should contain background task executables, scripts, and configuration files, but not the `launchd` configuration files.
	// The device stores the expanded zip archive in the tamper-proof `/var/db/ManagedConfigurationFiles/BackgroundTaskServices/Services/<TaskType>` directory, where `<TaskType>` is the value of the `TaskType` key in the configuration.
	ExecutableAssetReference *string `json:"ExecutableAssetReference,omitempty"`
	// An array of `launchd` configuration files used to run the background tasks.
	// The device stores the configuration files in the tamper-proof `/var/db/ManagedConfigurationFiles/BackgroundTaskServices/LaunchDaemons` or `/var/db/ManagedConfigurationFiles/BackgroundTaskServices/LaunchAgents` directories, based on the `Context` value for each item.
	LaunchdConfigurations *[]LaunchdItem `json:"LaunchdConfigurations,omitempty"`
}

func (p *ServicesBackgroundTasks) DeclarationType() string {
	return "com.apple.configuration.services.background-tasks"
}

// A dictionary of launchd configurations.
type LaunchdItem struct {
	// Specifies the identifier of an asset declaration containing a reference to the launchd configuration file for the background task. The referenced data must be a property list file conforming to the launchd.plist format. The asset's "ContentType" and "Hash-SHA-256" keys in the "Reference" key are required.
	FileAssetReference string `json:"FileAssetReference" required:"true"`
	// Indicates whether the device applies the launchd configuration file to the system daemon or system agent domain.
	Context Context `json:"Context" required:"true"`
}

// Indicates whether the device applies the launchd configuration file to the system daemon or system agent domain.
type Context string

const (
	ContextDaemon Context = "daemon"
	ContextAgent  Context = "agent"
)

// The declaration to configure managed configuration files for services.
type ServicesConfigurationFiles struct {
	// The identifier of the system service with managed configuration files. Use a reverse DNS style for this identifier.
	ServiceType string `json:"ServiceType" required:"true"`
	// The identifier of an asset declaration that contains a reference to the files to use for system service configuration. Ensure that the corresponding asset:
	// - Is of type `com.apple.asset.data`
	// - Is a zip archive of an entire directory
	// - Has a `Reference` key that includes the `ContentType` and `Hash-SHA-256` keys, which the system requires
	// The system expands the zip archive and stores the data in a well-known location for the service.
	DataAssetReference string `json:"DataAssetReference" required:"true"`
}

func (p *ServicesConfigurationFiles) DeclarationType() string {
	return "com.apple.configuration.services.configuration-files"
}

// The declaration to configure Siri settings.
type SiriSettings struct {
	// If `false`, disables Siri.
	Enabled *bool `json:"Enabled,omitempty"`
	// If `false`, disables the Siri AI specific features of Siri.
	AllowSiriAI *bool `json:"AllowSiriAI,omitempty"`
	// If `false`, disables Siri user-generated content.
	AllowUserGeneratedContent *bool `json:"AllowUserGeneratedContent,omitempty"`
	// If `false`, disables Siri while locked.
	AllowWhileLocked *bool `json:"AllowWhileLocked,omitempty"`
	// If `true`, forces Siri profanity filter.
	ForceProfanityFilter *bool `json:"ForceProfanityFilter,omitempty"`
	// If `true`, forces Siri to reduce sensitive content.
	ForceReduceSensitiveContent *bool `json:"ForceReduceSensitiveContent,omitempty"`
}

func (p *SiriSettings) DeclarationType() string {
	return "com.apple.configuration.siri.settings"
}

// The declaration to configure a software update enforcement policy for a specific OS release.
type SoftwareUpdateEnforcementSpecific struct {
	// The target OS version to update the device to by the appropriate time. This is the OS version number, for example, `16.1`.
	TargetOSVersion string `json:"TargetOSVersion" required:"true"`
	// The target build version to update the device to by the appropriate time, for example, `20A242`. Use the build version for testing during seeding periods. The build version can include a supplemental version identifier, for example, `20A242a`.
	TargetBuildVersion *string `json:"TargetBuildVersion,omitempty"`
	// The local date time value that specifies when to force install the software update. Use the format `yyyy-mm-ddThh:mm:ss`, which is derived from RFC 3339 but doesn't include a time zone offset or fractional seconds. If the user doesn't trigger the software update before this time, the device force installs it.
	TargetLocalDateTime string `json:"TargetLocalDateTime" required:"true"`
	// The URL of a web page that shows details that the organization provides about the enforced software update.
	DetailsURL *string `json:"DetailsURL,omitempty"`
}

func (p *SoftwareUpdateEnforcementSpecific) DeclarationType() string {
	return "com.apple.configuration.softwareupdate.enforcement.specific"
}

// The declaration to configure software updates.
type SoftwareUpdateSettings struct {
	// If set to `true`, the device shows all software update enforcement notifications.
	// If set to `false`, the device only shows notifications triggered one hour before the enforcement deadline, and the restart countdown notification.
	Notifications *bool `json:"Notifications,omitempty"`
	// This object configures the deferral of software updates. Background Security Improvements aren't considered in `Major`, `Minor`, or `System` deferral mechanism.
	Deferrals *Deferrals `json:"Deferrals,omitempty"`
	// This string specifies how the device shows software updates to the user. When more than one update is available update, the device behaves as follows:
	// - `All` - Shows all software update versions.
	// - `Oldest` - Shows only the oldest (lower numbered) software update version.
	// - `Newest` - Shows only the newest (highest numbered) software update version.
	RecommendedCadence *RecommendedCadence `json:"RecommendedCadence,omitempty"`
	// This object configures various automatic Software Update functionality.
	AutomaticActions *AutomaticActions `json:"AutomaticActions,omitempty"`
	// These configurations set user access to interacting with Background Security Improvement.
	RapidSecurityResponse *RapidSecurityResponse `json:"RapidSecurityResponse,omitempty"`
	// If set to `true`, a standard user can perform Major and Minor Software Updates.
	// If set to `false`, only administrators can perform Major and Minor Software Updates.
	AllowStandardUserOSUpdates *bool `json:"AllowStandardUserOSUpdates,omitempty"`
	// This object configures the beta program settings for a device.
	Beta *Beta `json:"Beta,omitempty"`
}

func (p *SoftwareUpdateSettings) DeclarationType() string {
	return "com.apple.configuration.softwareupdate.settings"
}

// This object configures the deferral of software updates. Background Security Improvements aren't considered in `Major`, `Minor`, or `System` deferral mechanism.
type Deferrals struct {
	// Specifies the number of days to defer a major or minor OS software update on the device. When set, software updates only appear after the specified delay, following the release of the software update.
	CombinedPeriodInDays *int64 `json:"CombinedPeriodInDays,omitempty"`
	// Specifies the number of days to defer a major OS software update on the device. When set, software updates only appear after the specified delay, following the release of the software update.
	MajorPeriodInDays *int64 `json:"MajorPeriodInDays,omitempty"`
	// Specifies the number of days to defer a minor OS software update on the device. It also defers major updates for iOS. When set, software updates only appear after the specified delay, following the release of the software update.
	MinorPeriodInDays *int64 `json:"MinorPeriodInDays,omitempty"`
	// Specifies the number of days to defer system or non-OS updates. When set, updates only appear after the specified delay, following the release of the update.
	SystemPeriodInDays *int64 `json:"SystemPeriodInDays,omitempty"`
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
	Download *Download `default:"Allowed" json:"Download,omitempty"`
	// Specifies whether the user can control automatic installation of available updates:
	// - `Allowed` - the user can enable or disable automatic installation.
	// - `AlwaysOn` - automatic installations are always enabled.
	// - `AlwaysOff` - automatic installations are always disabled.
	// > Note:
	// > The device uses this only with automatic downloads enabled.
	InstallOSUpdates *InstallOSUpdates `default:"Allowed" json:"InstallOSUpdates,omitempty"`
	// Specifies whether the user can control automatic installation of available security updates:
	// - `Allowed` - the user can enable or disable automatic installation.
	// - `AlwaysOn` - automatic installations are always enabled.
	// - `AlwaysOff` - automatic installations are always disabled.
	// > Note:
	// > The device uses this only with automatic downloads enabled.
	InstallSecurityUpdate *InstallSecurityUpdate `default:"Allowed" json:"InstallSecurityUpdate,omitempty"`
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
// > Note:
// > The device uses this only with automatic downloads enabled.
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
// > Note:
// > The device uses this only with automatic downloads enabled.
type InstallSecurityUpdate string

const (
	InstallSecurityUpdateAllowed   InstallSecurityUpdate = "Allowed"
	InstallSecurityUpdateAlwaysOn  InstallSecurityUpdate = "AlwaysOn"
	InstallSecurityUpdateAlwaysOff InstallSecurityUpdate = "AlwaysOff"
)

// These configurations set user access to interacting with Background Security Improvement.
type RapidSecurityResponse struct {
	// If set to `false`, Background Security Improvements aren't offered for user installation. The system can still install Background Security Improvements with `com.apple.configuration.softwareupdate.enforcement.specific` configurations.
	// If set to `true`, the system offers Background Security Improvements to the user.
	Enable *bool `json:"Enable,omitempty"`
	// If set to `false`, the system doesn't offer Background Security Improvement rollbacks to the user.
	// If set to `true`, the system offers Background Security Improvement rollbacks to the user.
	EnableRollback *bool `json:"EnableRollback,omitempty"`
}

// This object configures the beta program settings for a device.
type Beta struct {
	// Specifies whether the user can control beta program enrollment in the software update settings UI:
	// - `Allowed` - the user can enroll in any applicable beta programs associated with their logged in Apple Account. If the `OfferPrograms` key is present, then the programs listed in that key are also presented to the user.
	// - `AlwaysOn` - the device uses the beta programs the organization specifies, and the user isn't able to enroll in a beta program using their logged in Apple Account. The device is automatically enrolled into the beta program specified by the `RequireProgram` key if it's present. Otherwise, the system presents the programs listed in the `OfferPrograms` key to the user to choose which to enroll with.
	// - `AlwaysOff` - The device isn't allowed to enroll in any beta programs. The system removes the device from any beta programs, if already enrolled.
	ProgramEnrollment *ProgramEnrollment `default:"Allowed" json:"ProgramEnrollment,omitempty"`
	// An array of beta programs allowed on the device. This key must only be present if the `ProgramEnrollment` key is set to `Allowed` or `AlwaysOn`. This key must not be present if the `RequireProgram` key is present. This key can be present on unsupervised devices where the `ProgramEnrollment` key isn't supported but is implicitly set to `Allowed`.
	OfferPrograms *[]Program `json:"OfferPrograms,omitempty"`
	// The device automatically enrolls in this beta program. This key must only be present if the `ProgramEnrollment` key is set to `AlwaysOn`. The `OfferPrograms` key must not be present if this key is present.
	RequireProgram *RequireProgram `json:"RequireProgram,omitempty"`
}

// Specifies whether the user can control beta program enrollment in the software update settings UI:
// - `Allowed` - the user can enroll in any applicable beta programs associated with their logged in Apple Account. If the `OfferPrograms` key is present, then the programs listed in that key are also presented to the user.
// - `AlwaysOn` - the device uses the beta programs the organization specifies, and the user isn't able to enroll in a beta program using their logged in Apple Account. The device is automatically enrolled into the beta program specified by the `RequireProgram` key if it's present. Otherwise, the system presents the programs listed in the `OfferPrograms` key to the user to choose which to enroll with.
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
	Description string `json:"Description" required:"true"`
	// The Apple School Manager or Apple Business seeding service token for the organization the MDM server is part of. The system uses this token to enroll the device in the corresponding beta program.
	Token string `json:"Token" required:"true"`
}

// The device automatically enrolls in this beta program. This key must only be present if the `ProgramEnrollment` key is set to `AlwaysOn`. The `OfferPrograms` key must not be present if this key is present.
type RequireProgram struct {
	// A human readable description of the beta program.
	Description string `json:"Description" required:"true"`
	// The Apple School Manager or Apple Business seeding service token for the organization the MDM server is part of. The system uses this token to enroll the device in the corresponding beta program.
	Token string `json:"Token" required:"true"`
}

// The declaration to configure an MDMv1 profile for Apple Watch enrollment.
type WatchEnrollment struct {
	// The URL of the profile that the Apple Watch downloads and installs if the user opts in to management during the pairing process, which needs to start with `https://`. Successful enrollment requires that the pairing iPhone is supervised and the profile contains an MDM payload. Apple Watch attempts to install each payload that the profile contains.
	EnrollmentProfileURL string `json:"EnrollmentProfileURL" required:"true"`
	// An array of identifiers of asset declarations that contain anchor certificates to use to evaluate the trust of the enrollment profile server. Set the type of the corresponding assets to `com.apple.asset.credential.certificate`.
	// These certificates are pinned, meaning that the server specified by the `EnrollmentProfileURL` must use a certificate that chains to one of the certs in this array.
	// If it chains to one of the built-in trusted root certificates but not one of the `AnchorCertificateAssetReferences` certs, the connection will fail.
	AnchorCertificateAssetReferences *[]string `json:"AnchorCertificateAssetReferences,omitempty"`
}

func (p *WatchEnrollment) DeclarationType() string {
	return "com.apple.configuration.watch.enrollment"
}

// The declaration to configure a WebContent Filter that uses a plugin.
type WebContentFilterPlugin struct {
	// The name of the web content filter that the system displays on the device.
	VisibleName string `json:"VisibleName" required:"true"`
	// The bundle ID of the plug-in that provides filtering service. Consult your filtering solution vendor to determine what to specify for this value.
	PluginBundleID string `json:"PluginBundleID" required:"true"`
	// A globally unique identifier for this content filter configuration. The content filter processes network traffic for managed apps with the same `ContentFilterUUID` in their app attributes. This key must be present for unsupervised devices and user enrollment.
	ContentFilterUUID *string `json:"ContentFilterUUID,omitempty"`
	// The server address, which may be the IP address, hostname, or URL.
	ServerAddress *string `json:"ServerAddress,omitempty"`
	// The organization string to pass to the third-party plug-in.
	Organization *string `json:"Organization,omitempty"`
	// The custom dictionary that the filtering service plug-in needs.
	VendorConfig *map[string]any `json:"VendorConfig,omitempty"`
	// Settings that control authentication.
	Authentication *WebContentFilterPluginAuthentication `json:"Authentication,omitempty"`
	// Settings that control authentication.
	Filter *Filter `json:"Filter,omitempty"`
}

func (p *WebContentFilterPlugin) DeclarationType() string {
	return "com.apple.configuration.webcontent-filter.plugin"
}

// Settings that control authentication.
type WebContentFilterPluginAuthentication struct {
	// The identifier of an asset declaration that contains the credentials (user name and password) to authenticate with the service.
	CredentialsAssetReference *string `json:"CredentialsAssetReference,omitempty"`
	// The identifier of a credential asset declaration that contains the identity that this account requires to authenticate with the service.
	IdentityAssetReference *string `json:"IdentityAssetReference,omitempty"`
}

// Settings that control authentication.
type Filter struct {
	// The system uses this value to derive the relative order of content filters. Filters with a grade of `firewall` see network traffic before filters with a grade of `inspector`. However, the system doesn't define the order of filters within a grade.
	Grade *Grade `default:"firewall" json:"Grade,omitempty"`
	// Settings that control the browser filter. If not present, the system doesn't use browser filtering.
	Browsers *Browsers `json:"Browsers,omitempty"`
	// Settings that control the socket filter. If not present, the system doesn't use socket filtering.
	Sockets *Sockets `json:"Sockets,omitempty"`
	// Settings that control the packet filter. If not present, the system doesn't use packet filtering.
	Packets *Packets `json:"Packets,omitempty"`
	// Settings that control the URL filter. If not present, the system doesn't use URL filtering.
	URLs *URLs `json:"URLs,omitempty"`
}

// The system uses this value to derive the relative order of content filters. Filters with a grade of `firewall` see network traffic before filters with a grade of `inspector`. However, the system doesn't define the order of filters within a grade.
type Grade string

const (
	GradeFirewall  Grade = "firewall"
	GradeInspector Grade = "inspector"
)

// Settings that control the browser filter. If not present, the system doesn't use browser filtering.
type Browsers struct {
	// If `true`, the system enables filtering WebKit traffic.
	Enabled bool `json:"Enabled" required:"true"`
}

// Settings that control the socket filter. If not present, the system doesn't use socket filtering.
type Sockets struct {
	// If `true`, enables the filtering of socket traffic.
	Enabled bool `json:"Enabled" required:"true"`
	// The data provider identifier. This string identifies the filter data provider when the filter starts running. Required when Enabled is true.
	// In iOS and visionOS, the identifier is a bundle ID, for example, "com.example.app".
	// In macOS, the identifier is a composed identifier. The format of the composed identifier is "Bundle-ID {Designated-Requirement}". "Bundle-ID" is the bundle identifier string of the provider. "Designated-Requirement" is the designated requirement string the device uses to match the code signature of the provider. For example, "com.example.app {anchor apple generic}".
	ProviderComposedIdentifier *string `json:"ProviderComposedIdentifier,omitempty"`
}

// Settings that control the packet filter. If not present, the system doesn't use packet filtering.
type Packets struct {
	// If `true`, the system enables filtering network packets.
	Enabled bool `json:"Enabled" required:"true"`
	// The packet provider identifier. This string identifies the filter data provider when the filter starts running. Required when Enabled is true.
	// The identifier is a composed identifier. The format of the composed identifier is "Bundle-ID {Designated-Requirement}". "Bundle-ID" is the bundle identifier string of the provider. "Designated-Requirement" is the designated requirement string the device uses to match the code signature of the provider. For example, "com.example.app {anchor apple generic}".
	ProviderComposedIdentifier *string `json:"ProviderComposedIdentifier,omitempty"`
}

// Settings that control the URL filter. If not present, the system doesn't use URL filtering.
type URLs struct {
	// If `true`, the system filters URL requests.
	Enabled bool `json:"Enabled" required:"true"`
	// A dictionary containing URL filter parameters. Required when `Enabled` is `true`.
	Parameters *Parameters `json:"Parameters,omitempty"`
}

// A dictionary containing URL filter parameters. Required when `Enabled` is `true`.
type Parameters struct {
	// The URL filter control provider identifier. This string identifies the filter data provider when the filter starts running. Required when Enabled is true.
	// In iOS, the identifier is a bundle ID, for example, "com.example.app".
	// In macOS, the identifier is a composed identifier. The format of the composed identifier is "Bundle-ID {Designated-Requirement}". "Bundle-ID" is the bundle identifier string of the provider. "Designated-Requirement" is the designated requirement string the device uses to match the code signature of the provider. For example, "com.example.app {anchor apple generic}".
	ProviderComposedIdentifier string `json:"ProviderComposedIdentifier" required:"true"`
	// A dictionary containing Private Information Retrieval server settings.
	PIR PIR `json:"PIR" required:"true"`
	// If `true`, the system blocks URLs if the filter is enabled, but it fails to make any filtering decision; for example, if there's a communication failure with the PIR server. If `false`, the system allows URLs if the filter is enabled, but it fails to make any filtering decision.
	FailClosed *bool `json:"FailClosed,omitempty"`
	// The time interval in seconds that the system uses to periodically run the `NEURLFilterControlProvider` app extension. The default value is 86400 seconds (1 day). The minimum allowed value is 2700 seconds (45 minutes). The system allows `NEURLFilterControlProvider` implementations to download prefilter Bloom filter data onto the device periodically at the specified interval. Implementations need to allow for a slight difference between the scheduled time and the actual runtime of the task, due to the scheduling mechanism on the system.
	PrefilterFetchFrequency *int64 `default:"86400" json:"PrefilterFetchFrequency,omitempty"`
}

// A dictionary containing Private Information Retrieval server settings.
type PIR struct {
	// The URL containing the domain name of the private information retrieval server.
	ServerURL string `json:"ServerURL" required:"true"`
	// The URL containing the domain name of Privacy Pass Issuer.
	PrivacyPassIssuerURL string `json:"PrivacyPassIssuerURL" required:"true"`
	// The identifier of an asset declaration containing the HTTP bearer token required to authenticate with the service. The bearer token is provided in the `Password` field of the asset data. The system uses this token to attest that it's a valid user when requesting anonymous authentication tokens for PIR exchanges.
	AuthenticationTokenAssetReference *string `json:"AuthenticationTokenAssetReference,omitempty"`
}
