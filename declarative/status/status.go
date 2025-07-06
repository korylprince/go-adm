// DO NOT EDIT
// generated from https://github.com/apple/device-management.git:0a4527c5ea21825fd23e08273ccdb9e2302458ce/declarative/status

package status

const DeviceManagementGenerateHash = "0a4527c5ea21825fd23e08273ccdb9e2302458ce"

var DeclarationMap = map[string]any{"status-reason": StatusReason{}}
var StatusItemType = map[string]any{
	"account.list.caldav":                                  StatusAccountListCalDAV{},
	"account.list.carddav":                                 StatusAccountListCardDAV{},
	"account.list.exchange":                                StatusAccountListExchange{},
	"account.list.google":                                  StatusAccountListGoogle{},
	"account.list.ldap":                                    StatusAccountListLDAP{},
	"account.list.mail.incoming":                           StatusAccountListMailIncoming{},
	"account.list.mail.outgoing":                           StatusAccountListMailOutgoing{},
	"account.list.subscribed-calendar":                     StatusAccountListSubscribedCalendar{},
	"app.managed.list":                                     StatusAppManagedList{},
	"device.identifier.serial-number":                      StatusDeviceSerialNumber{},
	"device.identifier.udid":                               StatusDeviceUDID{},
	"device.model.family":                                  StatusDeviceModelFamily{},
	"device.model.identifier":                              StatusDeviceModelIdentifier{},
	"device.model.marketing-name":                          StatusDeviceModelMarketingName{},
	"device.model.number":                                  StatusDeviceModelNumber{},
	"device.operating-system.build-version":                StatusDeviceOperatingSystemBuildVersion{},
	"device.operating-system.family":                       StatusDeviceOperatingSystemFamily{},
	"device.operating-system.marketing-name":               StatusDeviceOperatingSystemMarketingName{},
	"device.operating-system.supplemental.build-version":   StatusDeviceOperatingSystemSupplementalBuildVersion{},
	"device.operating-system.supplemental.extra-version":   StatusDeviceOperatingSystemSupplementalExtraVersion{},
	"device.operating-system.version":                      StatusDeviceOperatingSystemVersion{},
	"device.power.battery-health":                          StatusDeviceBatteryHealth{},
	"diskmanagement.filevault.enabled":                     StatusDiskManagementFileVaultEnabled{},
	"management.client-capabilities":                       StatusManagementClientCapabilities{},
	"management.declarations":                              StatusManagementDeclarations{},
	"mdm.app":                                              StatusMDMApp{},
	"package.list":                                         StatusPackageList{},
	"passcode.is-compliant":                                StatusPasscodeCompliance{},
	"passcode.is-present":                                  StatusPasscodeIsPresent{},
	"screensharing.connection.group.unresolved-connection": StatusScreenSharingConnectionGroupUnresolvedConnections{},
	"security.certificate.list":                            StatusSecurityCertificateList{},
	"services.background-task":                             StatusServicesBackgroundTask{},
	"softwareupdate.beta-enrollment":                       StatusSoftwareUpdateBetaEnrollment{},
	"softwareupdate.device-id":                             StatusSoftwareUpdateDeviceID{},
	"softwareupdate.failure-reason":                        StatusSoftwareUpdateFailureReason{},
	"softwareupdate.install-reason":                        StatusSoftwareUpdateInstallReason{},
	"softwareupdate.install-state":                         StatusSoftwareUpdateInstallState{},
	"softwareupdate.pending-version":                       StatusSoftwareUpdatePendingVersion{},
	"test.array-value":                                     StatusTestArrayValue{},
	"test.boolean-value":                                   StatusTestBooleanValue{},
	"test.dictionary-value":                                StatusTestDictionaryValue{},
	"test.error-value":                                     StatusTestErrorValue{},
	"test.integer-value":                                   StatusTestIntegerValue{},
	"test.real-value":                                      StatusTestRealValue{},
	"test.string-value":                                    StatusTestStringValue{},
}

// A status report of the client's Calendar accounts.
type StatusAccountListCalDAV struct {
	// A list of status values for the Calendar accounts.
	Accountlistcaldav []*AccountlistcaldavStatusValue `json:"account.list.caldav" plist:"account.list.caldav" required:"true"`
}

func (p *StatusAccountListCalDAV) StatusItemType() string {
	return "account.list.caldav"
}

// A status report of the client's Calendar account details.
type AccountlistcaldavStatusValue struct {
	// The unique identifier for the account.
	Identifier string `json:"identifier" plist:"identifier" required:"true"`
	// If `true`, the account is removed and the status item object only contains this key and the `identifier` key.
	Removed *bool `json:"_removed,omitempty" plist:"_removed,omitempty"`
	// The identifier of the declaration that installed the account. Only present if a declaration installed the account.
	DeclarationIdentifier *string `json:"declaration-identifier,omitempty" plist:"declaration-identifier,omitempty"`
	// The name of the account.
	VisibleName *string `json:"visible-name,omitempty" plist:"visible-name,omitempty"`
	// The server host name for the account.
	Hostname *string `json:"hostname,omitempty" plist:"hostname,omitempty"`
	// The server port for the account.
	Port *int64 `json:"port,omitempty" plist:"port,omitempty"`
	// The user name for the account.
	Username *string `json:"username,omitempty" plist:"username,omitempty"`
	// If `true`, the Calendar app is displaying calendars and events for the account.
	AreCalendarsEnabled *bool `json:"are-calendars-enabled,omitempty" plist:"are-calendars-enabled,omitempty"`
	// If `true`, the Reminders app is displaying reminders for the account.
	AreRemindersEnabled *bool `json:"are-reminders-enabled,omitempty" plist:"are-reminders-enabled,omitempty"`
}

// A status report of the client's Contacts accounts.
type StatusAccountListCardDAV struct {
	// A list of status values for the Contacts accounts.
	Accountlistcarddav []*AccountlistcarddavStatusValue `json:"account.list.carddav" plist:"account.list.carddav" required:"true"`
}

func (p *StatusAccountListCardDAV) StatusItemType() string {
	return "account.list.carddav"
}

// A status report of the client's Contacts account details.
type AccountlistcarddavStatusValue struct {
	// The unique identifier for the account.
	Identifier string `json:"identifier" plist:"identifier" required:"true"`
	// If `true`, the account is removed and the status item object only contains this key and the `identifier` key.
	Removed *bool `json:"_removed,omitempty" plist:"_removed,omitempty"`
	// The identifier of the declaration that installed the account. Only present if a declaration installed the account.
	DeclarationIdentifier *string `json:"declaration-identifier,omitempty" plist:"declaration-identifier,omitempty"`
	// The name of the account.
	VisibleName *string `json:"visible-name,omitempty" plist:"visible-name,omitempty"`
	// The server host name for the account.
	Hostname *string `json:"hostname,omitempty" plist:"hostname,omitempty"`
	// The server port for the account.
	Port *int64 `json:"port,omitempty" plist:"port,omitempty"`
	// The user name for the account.
	Username *string `json:"username,omitempty" plist:"username,omitempty"`
}

// A status report of the client's Exchange accounts.
type StatusAccountListExchange struct {
	// A list of status values for the Exchange accounts.
	Accountlistexchange []*AccountlistexchangeStatusValue `json:"account.list.exchange" plist:"account.list.exchange" required:"true"`
}

func (p *StatusAccountListExchange) StatusItemType() string {
	return "account.list.exchange"
}

// A status report of the client's Exchange account details.
type AccountlistexchangeStatusValue struct {
	// The unique identifier for the account.
	Identifier string `json:"identifier" plist:"identifier" required:"true"`
	// If `true`, the account is removed and the status item object only contains this key and the `identifier` key.
	Removed *bool `json:"_removed,omitempty" plist:"_removed,omitempty"`
	// The identifier of the declaration that installed the account. Only present if a declaration installed the account.
	DeclarationIdentifier *string `json:"declaration-identifier,omitempty" plist:"declaration-identifier,omitempty"`
	// The name of the account.
	VisibleName *string `json:"visible-name,omitempty" plist:"visible-name,omitempty"`
	// The server host name for the account.
	Hostname *string `json:"hostname,omitempty" plist:"hostname,omitempty"`
	// The server port for the account.
	Port *int64 `json:"port,omitempty" plist:"port,omitempty"`
	// The user name for the account.
	Username *string `json:"username,omitempty" plist:"username,omitempty"`
	// A Boolean value that indicates whether the Mail app displays mail for this account.
	IsMailEnabled *bool `json:"is-mail-enabled,omitempty" plist:"is-mail-enabled,omitempty"`
	// A Boolean value that indicates whether the Calendar app displays calendars and events for this account.
	AreCalendarsEnabled *bool `json:"are-calendars-enabled,omitempty" plist:"are-calendars-enabled,omitempty"`
	// A Boolean value that indicates whether the Contacts app displays contacts for this account.
	AreContactsEnabled *bool `json:"are-contacts-enabled,omitempty" plist:"are-contacts-enabled,omitempty"`
	// A Boolean value that indicates whether the Notes app displays notes for this account.
	AreNotesEnabled *bool `json:"are-notes-enabled,omitempty" plist:"are-notes-enabled,omitempty"`
	// A Boolean value that indicates whether the Reminders app displays reminders for this account.
	AreRemindersEnabled *bool `json:"are-reminders-enabled,omitempty" plist:"are-reminders-enabled,omitempty"`
}

// A status report of the client's Google accounts.
type StatusAccountListGoogle struct {
	// A list of status values for the Google accounts.
	Accountlistgoogle []*AccountlistgoogleStatusValue `json:"account.list.google" plist:"account.list.google" required:"true"`
}

func (p *StatusAccountListGoogle) StatusItemType() string {
	return "account.list.google"
}

// A status report of the client's Google account details.
type AccountlistgoogleStatusValue struct {
	// The unique identifier for the account.
	Identifier string `json:"identifier" plist:"identifier" required:"true"`
	// If `true`, the account is removed and the status item object only contains this key and the `identifier` key.
	Removed *bool `json:"_removed,omitempty" plist:"_removed,omitempty"`
	// The identifier of the declaration that installed the account. Only present if a declaration installed the account.
	DeclarationIdentifier *string `json:"declaration-identifier,omitempty" plist:"declaration-identifier,omitempty"`
	// The name of the account.
	VisibleName *string `json:"visible-name,omitempty" plist:"visible-name,omitempty"`
	// The user name for the account.
	Username *string `json:"username,omitempty" plist:"username,omitempty"`
	// A Boolean value that indicates whether the Mail app displays mail for this account.
	IsMailEnabled *bool `json:"is-mail-enabled,omitempty" plist:"is-mail-enabled,omitempty"`
	// A Boolean value that indicates whether the Calendar app displays calendars and events for this account.
	AreCalendarsEnabled *bool `json:"are-calendars-enabled,omitempty" plist:"are-calendars-enabled,omitempty"`
	// A Boolean value that indicates whether the Contacts app displays contacts for this account.
	AreContactsEnabled *bool `json:"are-contacts-enabled,omitempty" plist:"are-contacts-enabled,omitempty"`
	// A Boolean value that indicates whether the Notes app displays notes for this account.
	AreNotesEnabled *bool `json:"are-notes-enabled,omitempty" plist:"are-notes-enabled,omitempty"`
}

// A status report of the client's Lightweight Directory Access Protocol (LDAP) accounts.
type StatusAccountListLDAP struct {
	// A list of status values for the LDAP accounts.
	Accountlistldap []*AccountlistldapStatusValue `json:"account.list.ldap" plist:"account.list.ldap" required:"true"`
}

func (p *StatusAccountListLDAP) StatusItemType() string {
	return "account.list.ldap"
}

// A status report of the client's LDAP account details.
type AccountlistldapStatusValue struct {
	// The unique identifier for the account.
	Identifier string `json:"identifier" plist:"identifier" required:"true"`
	// If `true`, the account is removed and the status item object only contains this key and the `identifier` key.
	Removed *bool `json:"_removed,omitempty" plist:"_removed,omitempty"`
	// The identifier of the declaration that installed the account. Only present if a declaration installed the account.
	DeclarationIdentifier *string `json:"declaration-identifier,omitempty" plist:"declaration-identifier,omitempty"`
	// The name of the account.
	VisibleName *string `json:"visible-name,omitempty" plist:"visible-name,omitempty"`
	// The server host name for the account.
	Hostname *string `json:"hostname,omitempty" plist:"hostname,omitempty"`
	// The server port for the account.
	Port *int64 `json:"port,omitempty" plist:"port,omitempty"`
	// The user name for the account.
	Username *string `json:"username,omitempty" plist:"username,omitempty"`
	// A Boolean value that indicates whether the account is enabled for use with the Contacts app.
	IsEnabled *bool `json:"is-enabled,omitempty" plist:"is-enabled,omitempty"`
}

// A status report of the client's incoming Mail accounts.
type StatusAccountListMailIncoming struct {
	// A list of status values for the incoming Mail accounts.
	Accountlistmailincoming []*AccountlistmailincomingStatusValue `json:"account.list.mail.incoming" plist:"account.list.mail.incoming" required:"true"`
}

func (p *StatusAccountListMailIncoming) StatusItemType() string {
	return "account.list.mail.incoming"
}

// A status report of the client's incoming Mail account details.
type AccountlistmailincomingStatusValue struct {
	// The unique identifier for the account.
	Identifier string `json:"identifier" plist:"identifier" required:"true"`
	// If `true`, the account is removed and the status item object only contains this key and the `identifier` key.
	Removed *bool `json:"_removed,omitempty" plist:"_removed,omitempty"`
	// The identifier of the declaration that installed the account. Only present if a declaration installed the account.
	DeclarationIdentifier *string `json:"declaration-identifier,omitempty" plist:"declaration-identifier,omitempty"`
	// The name of the account.
	VisibleName *string `json:"visible-name,omitempty" plist:"visible-name,omitempty"`
	// The server host name for the account.
	Hostname *string `json:"hostname,omitempty" plist:"hostname,omitempty"`
	// The server port for the account.
	Port *int64 `json:"port,omitempty" plist:"port,omitempty"`
	// The user name for the account.
	Username *string `json:"username,omitempty" plist:"username,omitempty"`
	// A Boolean value that indicates whether the Mail app displays mail for this account.
	IsMailEnabled *bool `json:"is-mail-enabled,omitempty" plist:"is-mail-enabled,omitempty"`
	// A Boolean value that indicates whether the Notes app displays notes for this account.
	AreNotesEnabled *bool `json:"are-notes-enabled,omitempty" plist:"are-notes-enabled,omitempty"`
}

// A status report of the client's outgoing Mail accounts.
type StatusAccountListMailOutgoing struct {
	// A list of status values for the outgoing Mail accounts.
	Accountlistmailoutgoing []*AccountlistmailoutgoingStatusValue `json:"account.list.mail.outgoing" plist:"account.list.mail.outgoing" required:"true"`
}

func (p *StatusAccountListMailOutgoing) StatusItemType() string {
	return "account.list.mail.outgoing"
}

// A status report of the client's outgoing Mail account details.
type AccountlistmailoutgoingStatusValue struct {
	// The unique identifier for the account.
	Identifier string `json:"identifier" plist:"identifier" required:"true"`
	// If `true`, the account is removed and the status item object only contains this key and the `identifier` key.
	Removed *bool `json:"_removed,omitempty" plist:"_removed,omitempty"`
	// The identifier of the declaration that installed the account. Only present if a declaration installed the account.
	DeclarationIdentifier *string `json:"declaration-identifier,omitempty" plist:"declaration-identifier,omitempty"`
	// The name of the account.
	VisibleName *string `json:"visible-name,omitempty" plist:"visible-name,omitempty"`
	// The server host name for the account.
	Hostname *string `json:"hostname,omitempty" plist:"hostname,omitempty"`
	// The server port for the account.
	Port *int64 `json:"port,omitempty" plist:"port,omitempty"`
	// The user name for the account.
	Username *string `json:"username,omitempty" plist:"username,omitempty"`
}

// A status report of the client's subscribed calendars.
type StatusAccountListSubscribedCalendar struct {
	// A list of status values for the subscribed calendars.
	AccountlistsubscribedCalendar []*AccountlistsubscribedCalendarStatusValue `json:"account.list.subscribed-calendar" plist:"account.list.subscribed-calendar" required:"true"`
}

func (p *StatusAccountListSubscribedCalendar) StatusItemType() string {
	return "account.list.subscribed-calendar"
}

// A status report of the client's subscribed calendar details.
type AccountlistsubscribedCalendarStatusValue struct {
	// The unique identifier for the subscribed calendar.
	Identifier string `json:"identifier" plist:"identifier" required:"true"`
	// If `true`, the subscribed calendar is removed and the status item object only contains this key and the `identifier` key.
	Removed *bool `json:"_removed,omitempty" plist:"_removed,omitempty"`
	// The identifier of the declaration that installed the subscribed calendar. Only present if a declaration installed the subscribed calendar.
	DeclarationIdentifier *string `json:"declaration-identifier,omitempty" plist:"declaration-identifier,omitempty"`
	// The name of the subscribed calendar.
	VisibleName *string `json:"visible-name,omitempty" plist:"visible-name,omitempty"`
	// The URL of the subscribed calendar.
	CalendarUrl *string `json:"calendar-url,omitempty" plist:"calendar-url,omitempty"`
	// The user name for authenticating with the subscribed calendar.
	Username *string `json:"username,omitempty" plist:"username,omitempty"`
	// A Boolean value that indicates whether the Calendar app displays this subscribed calendar.
	IsEnabled *bool `json:"is-enabled,omitempty" plist:"is-enabled,omitempty"`
}

// The device's declarative managed apps.
type StatusAppManagedList struct {
	// An array of dictionaries that describe the device's declarative managed apps.
	Appmanagedlist []*AppmanagedlistStatusValue `json:"app.managed.list" plist:"app.managed.list" required:"true"`
}

func (p *StatusAppManagedList) StatusItemType() string {
	return "app.managed.list"
}

// A dictionary that describes a declarative managed app.
type AppmanagedlistStatusValue struct {
	// The app's bundle id, which is unique.
	Identifier string `json:"identifier" plist:"identifier" required:"true"`
	// If `true`, the system removed the app and only this key and the `identifier` key are present in the status item object.
	Removed *bool `json:"_removed,omitempty" plist:"_removed,omitempty"`
	// The identifier of the declaration that controls the app.
	DeclarationIdentifier *string `json:"declaration-identifier,omitempty" plist:"declaration-identifier,omitempty"`
	// The name of the app.
	Name *string `json:"name,omitempty" plist:"name,omitempty"`
	// The app's external version ID. You can also retrieve this value from the store through the `contentMetadataLookupUrl` of the `VPPServiceConfigSrv`.
	// In the response from `uclient-api.itunes.apple.com` URL, there's an `externalId` at the path `results.<adamId>.offers[0].version.externalId`. If the current external version identifier of an app on the store doesn't match the external version identifier reported by the device, there may be an app update available for the device.
	ExternalVersionId *int64 `json:"external-version-id,omitempty" plist:"external-version-id,omitempty"`
	// The version of the app.
	Version *string `json:"version,omitempty" plist:"version,omitempty"`
	// The short version of the app.
	ShortVersion *string `json:"short-version,omitempty" plist:"short-version,omitempty"`
	// The status of the app, which has the following possible values:
	// - `optional`: The app is optional and the user has to trigger its installation.
	// - `queued`: The system has started installation of the app.
	// - `not-present`: Management of the app occurs after it is installed.
	// - `prompting-for-consent`: The system is displaying a prompt to the user to proceed with app installation.
	// - `prompting-for-login`: The system is displaying an App Store sign-in prompt to the user to allow app installation.
	// - `prompting-for-management`: The system is displaying a prompt to the user to allow changing the installed app to a managed app.
	// - `downloading`: The system is downloading the app.
	// - `installing`: The system is installing the app.
	// - `managed`: The app is installed and managed.
	// - `managed-but-uninstalled`: The app is required, but the system hasn't installed it. The app becomes managed if the system installs it again. If the user removes an optional app, its state is `optional`, not `managed-but-uninstalled`.
	// - `failed`: The app install failed.
	State *AppmanagedlistStatusValueState `json:"state,omitempty" plist:"state,omitempty"`
	// The update status of the app, which has the following possible values:
	// - `available`: An update is available for the app.
	// - `prompting-for-update`: The system is displaying a prompt to the user to proceed with app installation.
	// - `prompting-for-update-login`: The system is displaying an App Store sign-in prompt to the user to allow app installation.
	// - `updating`: The app is updating.
	// - `failed`: The app update failed.
	// > Note:
	// > This key is only present if `state` is `managed` and an update is available.
	UpdateState *UpdateState `json:"update-state,omitempty" plist:"update-state,omitempty"`
	// The status of app or extension managed configurations. This key is only present when managed configurations are available for the managed app or any of its extensions.
	ConfigState *ConfigState `json:"config-state,omitempty" plist:"config-state,omitempty"`
	// An array that contains additional details about the app state, including errors.
	Reasons *[]*AppmanagedlistStatusValueReasonsReasons `json:"reasons,omitempty" plist:"reasons,omitempty"`
}

// The status of the app, which has the following possible values:
// - `optional`: The app is optional and the user has to trigger its installation.
// - `queued`: The system has started installation of the app.
// - `not-present`: Management of the app occurs after it is installed.
// - `prompting-for-consent`: The system is displaying a prompt to the user to proceed with app installation.
// - `prompting-for-login`: The system is displaying an App Store sign-in prompt to the user to allow app installation.
// - `prompting-for-management`: The system is displaying a prompt to the user to allow changing the installed app to a managed app.
// - `downloading`: The system is downloading the app.
// - `installing`: The system is installing the app.
// - `managed`: The app is installed and managed.
// - `managed-but-uninstalled`: The app is required, but the system hasn't installed it. The app becomes managed if the system installs it again. If the user removes an optional app, its state is `optional`, not `managed-but-uninstalled`.
// - `failed`: The app install failed.
type AppmanagedlistStatusValueState string

const (
	AppmanagedlistStatusValueStateOptional               AppmanagedlistStatusValueState = "optional"
	AppmanagedlistStatusValueStateQueued                 AppmanagedlistStatusValueState = "queued"
	AppmanagedlistStatusValueStateNotPresent             AppmanagedlistStatusValueState = "not-present"
	AppmanagedlistStatusValueStatePromptingForConsent    AppmanagedlistStatusValueState = "prompting-for-consent"
	AppmanagedlistStatusValueStatePromptingForLogin      AppmanagedlistStatusValueState = "prompting-for-login"
	AppmanagedlistStatusValueStatePromptingForManagement AppmanagedlistStatusValueState = "prompting-for-management"
	AppmanagedlistStatusValueStateDownloading            AppmanagedlistStatusValueState = "downloading"
	AppmanagedlistStatusValueStateInstalling             AppmanagedlistStatusValueState = "installing"
	AppmanagedlistStatusValueStateManaged                AppmanagedlistStatusValueState = "managed"
	AppmanagedlistStatusValueStateManagedButUninstalled  AppmanagedlistStatusValueState = "managed-but-uninstalled"
	AppmanagedlistStatusValueStateFailed                 AppmanagedlistStatusValueState = "failed"
)

// The update status of the app, which has the following possible values:
// - `available`: An update is available for the app.
// - `prompting-for-update`: The system is displaying a prompt to the user to proceed with app installation.
// - `prompting-for-update-login`: The system is displaying an App Store sign-in prompt to the user to allow app installation.
// - `updating`: The app is updating.
// - `failed`: The app update failed.
// > Note:
// > This key is only present if `state` is `managed` and an update is available.
type UpdateState string

const (
	UpdateStateAvailable               UpdateState = "available"
	UpdateStatePromptingForUpdate      UpdateState = "prompting-for-update"
	UpdateStatePromptingForUpdateLogin UpdateState = "prompting-for-update-login"
	UpdateStateUpdating                UpdateState = "updating"
	UpdateStateFailed                  UpdateState = "failed"
)

// The status of app or extension managed configurations. This key is only present when managed configurations are available for the managed app or any of its extensions.
type ConfigState struct {
	// The status of any app managed configuration. This key is only present when the managed app has a managed configuration.
	AppConfigState *AppConfigState `json:"app-config-state,omitempty" plist:"app-config-state,omitempty"`
	// The status of any app extension managed configuration. This key's value is a dictionary whose keys are the bundle identifiers of app extensions that have a managed configuration. The values of each key represent the status of the corresponding app extension's managed configuration.
	ExtensionConfigState *ExtensionConfigState `json:"extension-config-state,omitempty" plist:"extension-config-state,omitempty"`
}

// The status of any app managed configuration. This key is only present when the managed app has a managed configuration.
type AppConfigState struct {
	// The managed configuration status.
	// - `unknown`: The managed configuration has not been read
	// - `invalid`: The managed configuration was read and deemed to be invalid
	// - `valid`: The managed configuration was read and deemed to be valid
	State AppConfigStateState `json:"state" plist:"state" required:"true"`
}

// The managed configuration status.
// - `unknown`: The managed configuration has not been read
// - `invalid`: The managed configuration was read and deemed to be invalid
// - `valid`: The managed configuration was read and deemed to be valid
type AppConfigStateState string

const (
	AppConfigStateStateUnknown AppConfigStateState = "unknown"
	AppConfigStateStateInvalid AppConfigStateState = "invalid"
	AppConfigStateStateValid   AppConfigStateState = "valid"
)

// The status of any app extension managed configuration. This key's value is a dictionary whose keys are the bundle identifiers of app extensions that have a managed configuration. The values of each key represent the status of the corresponding app extension's managed configuration.
type ExtensionConfigState struct {
	// The bundle identifier of the managed app extension.
	ANY *ANY `json:"ANY,omitempty" plist:"ANY,omitempty"`
}

// The bundle identifier of the managed app extension.
type ANY struct {
	// The managed configuration status.
	// - `unknown`: The managed configuration has not been read
	// - `invalid`: The managed configuration was read and deemed to be invalid
	// - `valid`: The managed configuration was read and deemed to be valid
	State AppConfigStateState `json:"state" plist:"state" required:"true"`
}

// Information about a status error.
type AppmanagedlistStatusValueReasonsReasons struct {
	// A code for the state.
	Code string `json:"code" plist:"code" required:"true"`
	// A description of the state.
	Description *string `json:"description,omitempty" plist:"description,omitempty"`
	// A dictionary that contains additional details about the state.
	Details *map[string]any `json:"details,omitempty" plist:"details,omitempty"`
}

// A status report of the device's serial number.
type StatusDeviceSerialNumber struct {
	// The device's serial number.
	DeviceidentifierserialNumber string `json:"device.identifier.serial-number" plist:"device.identifier.serial-number" required:"true"`
}

func (p *StatusDeviceSerialNumber) StatusItemType() string {
	return "device.identifier.serial-number"
}

// A status report of the device's UDID.
type StatusDeviceUDID struct {
	// The device's UDID. This value is always available on the device channel. This value is only available on user channels whose organization matches that of the device channel.
	Deviceidentifierudid string `json:"device.identifier.udid" plist:"device.identifier.udid" required:"true"`
}

func (p *StatusDeviceUDID) StatusItemType() string {
	return "device.identifier.udid"
}

// A status report of the device's hardware family.
type StatusDeviceModelFamily struct {
	// The hardware family of the device, such as `Mac`, `iPhone`, or `iPad`.
	Devicemodelfamily string `json:"device.model.family" plist:"device.model.family" required:"true"`
}

func (p *StatusDeviceModelFamily) StatusItemType() string {
	return "device.model.family"
}

// A status report of the device's hardware identifier.
type StatusDeviceModelIdentifier struct {
	// A two-part string that specifies the device's model. The first part specifies device's model family, and the second part specifies the model's version. The model's version is a comma-separated number where the first part of the number is the version, and the second part is a variant, such as `MacBookPro15,1` or `iPhone13,2`.
	Devicemodelidentifier string `json:"device.model.identifier" plist:"device.model.identifier" required:"true"`
}

func (p *StatusDeviceModelIdentifier) StatusItemType() string {
	return "device.model.identifier"
}

// A status report of the device's marketing name.
type StatusDeviceModelMarketingName struct {
	// The device's marketing name, such as `iPhone 12`. This value may not always be available.
	DevicemodelmarketingName string `json:"device.model.marketing-name" plist:"device.model.marketing-name" required:"true"`
}

func (p *StatusDeviceModelMarketingName) StatusItemType() string {
	return "device.model.marketing-name"
}

// A status report of the device's hardware number.
type StatusDeviceModelNumber struct {
	// The device's model number.
	Devicemodelnumber string `json:"device.model.number" plist:"device.model.number" required:"true"`
}

func (p *StatusDeviceModelNumber) StatusItemType() string {
	return "device.model.number"
}

// A status report of the device's software build identifier.
type StatusDeviceOperatingSystemBuildVersion struct {
	// The operating system's build version on the device, such as `18F132`.
	DeviceoperatingSystembuildVersion string `json:"device.operating-system.build-version" plist:"device.operating-system.build-version" required:"true"`
}

func (p *StatusDeviceOperatingSystemBuildVersion) StatusItemType() string {
	return "device.operating-system.build-version"
}

// A status report of the device's operating system family.
type StatusDeviceOperatingSystemFamily struct {
	// The operating system family in use on the device, such as `macOS` or `iOS`.
	DeviceoperatingSystemfamily string `json:"device.operating-system.family" plist:"device.operating-system.family" required:"true"`
}

func (p *StatusDeviceOperatingSystemFamily) StatusItemType() string {
	return "device.operating-system.family"
}

// A status report of the device's operating system marketing name.
type StatusDeviceOperatingSystemMarketingName struct {
	// The operating system's marketing name in use on the device, such as `Catalina`.
	DeviceoperatingSystemmarketingName string `json:"device.operating-system.marketing-name" plist:"device.operating-system.marketing-name" required:"true"`
}

func (p *StatusDeviceOperatingSystemMarketingName) StatusItemType() string {
	return "device.operating-system.marketing-name"
}

// A status report of the device's operating system supplemental build identifier.
type StatusDeviceOperatingSystemSupplementalBuildVersion struct {
	// The operating system's build and rapid security response versions in use on the device, for example, `20A123a` or `20B27c`.
	DeviceoperatingSystemsupplementalbuildVersion string `json:"device.operating-system.supplemental.build-version" plist:"device.operating-system.supplemental.build-version" required:"true"`
}

func (p *StatusDeviceOperatingSystemSupplementalBuildVersion) StatusItemType() string {
	return "device.operating-system.supplemental.build-version"
}

// A status report of the device's operating system's rapid security response identifier.
type StatusDeviceOperatingSystemSupplementalExtraVersion struct {
	// The operating system's rapid security response version in use on the device, for example, `a`.
	DeviceoperatingSystemsupplementalextraVersion string `json:"device.operating-system.supplemental.extra-version" plist:"device.operating-system.supplemental.extra-version" required:"true"`
}

func (p *StatusDeviceOperatingSystemSupplementalExtraVersion) StatusItemType() string {
	return "device.operating-system.supplemental.extra-version"
}

// A status report of the device's operating system version.
type StatusDeviceOperatingSystemVersion struct {
	// The operating system's version in use on the device, such as `15.0`.
	DeviceoperatingSystemversion string `json:"device.operating-system.version" plist:"device.operating-system.version" required:"true"`
}

func (p *StatusDeviceOperatingSystemVersion) StatusItemType() string {
	return "device.operating-system.version"
}

// The device's battery health.
type StatusDeviceBatteryHealth struct {
	// The battery health status, which has the following values:
	// - `non-genuine`: The battery isn't a genuine Apple battery.
	// - `normal`: The battery is operating normally.
	// - `service-recommended`: The system recommends battery service.
	// - `unknown`: The system couldn't determine battery health information.
	// - `unsupported`: The device doesn't support battery health reporting.
	// Available in iOS 17 and later on iPhone, iPadOS 18.4 and later on supported iPad models, and macOS 14.4 and later on Mac computers with Apple silicon.
	DevicepowerbatteryHealth DevicepowerbatteryHealth `json:"device.power.battery-health" plist:"device.power.battery-health" required:"true"`
}

func (p *StatusDeviceBatteryHealth) StatusItemType() string {
	return "device.power.battery-health"
}

// The battery health status, which has the following values:
// - `non-genuine`: The battery isn't a genuine Apple battery.
// - `normal`: The battery is operating normally.
// - `service-recommended`: The system recommends battery service.
// - `unknown`: The system couldn't determine battery health information.
// - `unsupported`: The device doesn't support battery health reporting.
// Available in iOS 17 and later on iPhone, iPadOS 18.4 and later on supported iPad models, and macOS 14.4 and later on Mac computers with Apple silicon.
type DevicepowerbatteryHealth string

const (
	DevicepowerbatteryHealthNonGenuine         DevicepowerbatteryHealth = "non-genuine"
	DevicepowerbatteryHealthNormal             DevicepowerbatteryHealth = "normal"
	DevicepowerbatteryHealthServiceRecommended DevicepowerbatteryHealth = "service-recommended"
	DevicepowerbatteryHealthUnknown            DevicepowerbatteryHealth = "unknown"
	DevicepowerbatteryHealthUnsupported        DevicepowerbatteryHealth = "unsupported"
)

// The enabled status of the File Vault.
type StatusDiskManagementFileVaultEnabled struct {
	// A Boolean value that specifies the File Vault enabled status on the device.
	Diskmanagementfilevaultenabled bool `json:"diskmanagement.filevault.enabled" plist:"diskmanagement.filevault.enabled" required:"true"`
}

func (p *StatusDiskManagementFileVaultEnabled) StatusItemType() string {
	return "diskmanagement.filevault.enabled"
}

// A status report of the client's protocol capabilities.
type StatusManagementClientCapabilities struct {
	// An object that contains the client's protocol capabilities. These typically only change when the device upgrades its software. An implicit status subscription for this status item is always present, so the client always reports changes to the server.
	ManagementclientCapabilities ManagementclientCapabilities `json:"management.client-capabilities" plist:"management.client-capabilities" required:"true"`
}

func (p *StatusManagementClientCapabilities) StatusItemType() string {
	return "management.client-capabilities"
}

// An object that contains the client's protocol capabilities. These typically only change when the device upgrades its software. An implicit status subscription for this status item is always present, so the client always reports changes to the server.
type ManagementclientCapabilities struct {
	// A list of protocol versions that the client supports.
	SupportedVersions []string `json:"supported-versions" plist:"supported-versions" required:"true"`
	// A set of optional protocol features that the client supports. Each object's key represents a feature, and the property value represents the feature's associated parameters.
	SupportedFeatures map[string]any `json:"supported-features" plist:"supported-features" required:"true"`
	// A set of declaration and status items that the client supports.
	SupportedPayloads SupportedPayloads `json:"supported-payloads" plist:"supported-payloads" required:"true"`
}

// A set of declaration and status items that the client supports.
type SupportedPayloads struct {
	// A set of declarations that the client supports.
	Declarations Declarations `json:"declarations" plist:"declarations" required:"true"`
	// A list of status items that the client supports.
	StatusItems []string `json:"status-items" plist:"status-items" required:"true"`
}

// A set of declarations that the client supports.
type Declarations struct {
	// An array of strings that represents the activation types that the client supports.
	Activations *[]string `json:"activations,omitempty" plist:"activations,omitempty"`
	// An array of strings that represents the assets that the client supports.
	Assets *[]string `json:"assets,omitempty" plist:"assets,omitempty"`
	// An array of strings that represents the configuration types that the client supports.
	Configurations *[]string `json:"configurations,omitempty" plist:"configurations,omitempty"`
	// An array of strings that represents the declaration types that the client supports.
	Management *[]string `json:"management,omitempty" plist:"management,omitempty"`
}

// A status report of the client's processed declarations.
type StatusManagementDeclarations struct {
	// A collection of the client's processed declarations.
	Managementdeclarations Managementdeclarations `json:"management.declarations" plist:"management.declarations" required:"true"`
}

func (p *StatusManagementDeclarations) StatusItemType() string {
	return "management.declarations"
}

// A collection of the client's processed declarations.
type Managementdeclarations struct {
	// An array of declarations that represent the client's processed activation types.
	Activations []*Activations `json:"activations" plist:"activations" required:"true"`
	// An array of declarations that represent the client's processed configuration types.
	Configurations []*Configurations `json:"configurations" plist:"configurations" required:"true"`
	// An array of declarations that represent the client's processed assets.
	Assets []*Assets `json:"assets" plist:"assets" required:"true"`
	// An array of declarations that represent the client's processed declaration types.
	Management []*Management `json:"management" plist:"management" required:"true"`
}

// Status for a declaration processed by the client.
type Activations struct {
	// The `identifier` of the declaration this status report refers to.
	Identifier string `json:"identifier" plist:"identifier" required:"true"`
	// The `ServerToken` of the declaration this status report refers to.
	ServerToken string `json:"server-token" plist:"server-token" required:"true"`
	// If `true`, the declaration is active on the device.
	Active bool `json:"active" plist:"active" required:"true"`
	// This string defines the validity of the declaration. If it's `invalid`, the `reasons` property contains more details.
	Valid Valid `json:"valid" plist:"valid" required:"true"`
	// The details of any client errors.
	Reasons *[]*ActivationsReasonsReasons `json:"reasons,omitempty" plist:"reasons,omitempty"`
}

// This string defines the validity of the declaration. If it's `invalid`, the `reasons` property contains more details.
type Valid string

const (
	ValidUnknown Valid = "unknown"
	ValidInvalid Valid = "invalid"
	ValidValid   Valid = "valid"
)

// Information about a status error.
type ActivationsReasonsReasons struct {
	// The error code for this error.
	Code string `json:"code" plist:"code" required:"true"`
	// The description for this error.
	Description *string `json:"description,omitempty" plist:"description,omitempty"`
	// A dictionary that contains further details about this error.
	Details *map[string]any `json:"details,omitempty" plist:"details,omitempty"`
}

// Status for a declaration processed by the client.
type Configurations struct {
	// The `identifier` of the declaration this status report refers to.
	Identifier string `json:"identifier" plist:"identifier" required:"true"`
	// The `ServerToken` of the declaration this status report refers to.
	ServerToken string `json:"server-token" plist:"server-token" required:"true"`
	// If `true`, the declaration is active on the device.
	Active bool `json:"active" plist:"active" required:"true"`
	// This string defines the validity of the declaration. If it's `invalid`, the `reasons` property contains more details.
	Valid Valid `json:"valid" plist:"valid" required:"true"`
	// The details of any client errors.
	Reasons *[]*ActivationsReasonsReasons `json:"reasons,omitempty" plist:"reasons,omitempty"`
}

// Status for a declaration processed by the client.
type Assets struct {
	// The `identifier` of the declaration this status report refers to.
	Identifier string `json:"identifier" plist:"identifier" required:"true"`
	// The `ServerToken` of the declaration this status report refers to.
	ServerToken string `json:"server-token" plist:"server-token" required:"true"`
	// If `true`, the declaration is active on the device.
	Active bool `json:"active" plist:"active" required:"true"`
	// This string defines the validity of the declaration. If it's `invalid`, the `reasons` property contains more details.
	Valid Valid `json:"valid" plist:"valid" required:"true"`
	// The details of any client errors.
	Reasons *[]*ActivationsReasonsReasons `json:"reasons,omitempty" plist:"reasons,omitempty"`
}

// Status for a declaration processed by the client.
type Management struct {
	// The `identifier` of the declaration this status report refers to.
	Identifier string `json:"identifier" plist:"identifier" required:"true"`
	// The `ServerToken` of the declaration this status report refers to.
	ServerToken string `json:"server-token" plist:"server-token" required:"true"`
	// If `true`, the declaration is active on the device.
	Active bool `json:"active" plist:"active" required:"true"`
	// This string defines the validity of the declaration. If it's `invalid`, the `reasons` property contains more details.
	Valid Valid `json:"valid" plist:"valid" required:"true"`
	// The details of any client errors.
	Reasons *[]*ActivationsReasonsReasons `json:"reasons,omitempty" plist:"reasons,omitempty"`
}

// A status report of the client's MDM-installed apps.
type StatusMDMApp struct {
	// The list of apps. The response doesn't include apps that are managed by Declarative Device Management.
	Mdmapp []*MdmappStatusValue `json:"mdm.app" plist:"mdm.app" required:"true"`
}

func (p *StatusMDMApp) StatusItemType() string {
	return "mdm.app"
}

// A status report that contains details about an MDM-installed app.
type MdmappStatusValue struct {
	// The app's bundle id, which is unique.
	Identifier string `json:"identifier" plist:"identifier" required:"true"`
	// If `true`, the system removed the app and only this key and the `identifier` key are present in the status item object. The device reports an MDM-installed app as removed if management of the app has been transferred to Declarative Device Management.
	Removed *bool `json:"_removed,omitempty" plist:"_removed,omitempty"`
	// The name of the app.
	Name *string `json:"name,omitempty" plist:"name,omitempty"`
	// The application's external version ID. Use `Service-Config` to get the `contentMetadataLookupUrl` endpoint. In the response from that URL, find a key named `externalId` at the path `results.<adamId>.offers[0].version.externalId`. If the current external version identifier of an app on the store doesn't match the external version identifier reported by the device, there may be an app update available for the device.
	ExternalVersionId *string `json:"external-version-id,omitempty" plist:"external-version-id,omitempty"`
	// The version of the app.
	Version *string `json:"version,omitempty" plist:"version,omitempty"`
	// The short version of the app.
	ShortVersion *string `json:"short-version,omitempty" plist:"short-version,omitempty"`
	// The status of the app that `ManagedApplicationListCommand` reports.
	State *MdmappStatusValueState `json:"state,omitempty" plist:"state,omitempty"`
}

// The status of the app that `ManagedApplicationListCommand` reports.
type MdmappStatusValueState string

const (
	MdmappStatusValueStateQueued                  MdmappStatusValueState = "queued"
	MdmappStatusValueStateNeedsRedemption         MdmappStatusValueState = "needs-redemption"
	MdmappStatusValueStateRedeeming               MdmappStatusValueState = "redeeming"
	MdmappStatusValueStatePrompting               MdmappStatusValueState = "prompting"
	MdmappStatusValueStatePromptingForLogin       MdmappStatusValueState = "prompting-for-login"
	MdmappStatusValueStateValidatingPurchase      MdmappStatusValueState = "validating-purchase"
	MdmappStatusValueStatePromptingForUpdate      MdmappStatusValueState = "prompting-for-update"
	MdmappStatusValueStatePromptingForUpdateLogin MdmappStatusValueState = "prompting-for-update-login"
	MdmappStatusValueStatePromptingForManagement  MdmappStatusValueState = "prompting-for-management"
	MdmappStatusValueStateValidatingUpdate        MdmappStatusValueState = "validating-update"
	MdmappStatusValueStateUpdating                MdmappStatusValueState = "updating"
	MdmappStatusValueStateInstalling              MdmappStatusValueState = "installing"
	MdmappStatusValueStateManaged                 MdmappStatusValueState = "managed"
	MdmappStatusValueStateManagedButUninstalled   MdmappStatusValueState = "managed-but-uninstalled"
	MdmappStatusValueStateUnknown                 MdmappStatusValueState = "unknown"
	MdmappStatusValueStateUserInstalledApp        MdmappStatusValueState = "user-installed-app"
	MdmappStatusValueStateUserRejected            MdmappStatusValueState = "user-rejected"
	MdmappStatusValueStateUpdateRejected          MdmappStatusValueState = "update-rejected"
	MdmappStatusValueStateManagementRejected      MdmappStatusValueState = "management-rejected"
	MdmappStatusValueStateFailed                  MdmappStatusValueState = "failed"
)

// The client's declarative packages.
type StatusPackageList struct {
	// An array of dictionaries that describe the device's declarative packages.
	Packagelist []*PackagelistStatusValue `json:"package.list" plist:"package.list" required:"true"`
}

func (p *StatusPackageList) StatusItemType() string {
	return "package.list"
}

// A dictionary that describes a declarative package.
type PackagelistStatusValue struct {
	// The package's unique identifier. This is the package identifier value of the package file.
	Identifier string `json:"identifier" plist:"identifier" required:"true"`
	// If `true`, the system removed the package and only this key and the `identifier` key are present in the status item object.
	Removed *bool `json:"_removed,omitempty" plist:"_removed,omitempty"`
	// The identifier of the declaration that controls the package.
	DeclarationIdentifier *string `json:"declaration-identifier,omitempty" plist:"declaration-identifier,omitempty"`
	// The name of the package.
	Name *string `json:"name,omitempty" plist:"name,omitempty"`
	// The version of the package. This will be the package version value of the package file.
	Version *string `json:"version,omitempty" plist:"version,omitempty"`
	// The status of the package, which has the following possible values:
	// - `optional`: The package is optional and the user has to trigger its installation.
	// - `queued`: Installation of the package has started.
	// - `prompting-for-consent`: The system is displaying a prompt to the user to proceed with package installation.
	// - `downloading`: The system is downloading the package.
	// - `installing`: The system is installing the package.
	// - `installed`: The package is installed.
	// - `failed`: The package install failed.
	State *PackagelistStatusValueState `json:"state,omitempty" plist:"state,omitempty"`
	// An array that contains additional details about the package state, including errors.
	Reasons *[]*PackagelistStatusValueReasonsReasons `json:"reasons,omitempty" plist:"reasons,omitempty"`
}

// The status of the package, which has the following possible values:
// - `optional`: The package is optional and the user has to trigger its installation.
// - `queued`: Installation of the package has started.
// - `prompting-for-consent`: The system is displaying a prompt to the user to proceed with package installation.
// - `downloading`: The system is downloading the package.
// - `installing`: The system is installing the package.
// - `installed`: The package is installed.
// - `failed`: The package install failed.
type PackagelistStatusValueState string

const (
	PackagelistStatusValueStateOptional            PackagelistStatusValueState = "optional"
	PackagelistStatusValueStateQueued              PackagelistStatusValueState = "queued"
	PackagelistStatusValueStatePromptingForConsent PackagelistStatusValueState = "prompting-for-consent"
	PackagelistStatusValueStateDownloading         PackagelistStatusValueState = "downloading"
	PackagelistStatusValueStateInstalling          PackagelistStatusValueState = "installing"
	PackagelistStatusValueStateInstalled           PackagelistStatusValueState = "installed"
	PackagelistStatusValueStateFailed              PackagelistStatusValueState = "failed"
)

// Information about a status error.
type PackagelistStatusValueReasonsReasons struct {
	// A code for the state.
	Code string `json:"code" plist:"code" required:"true"`
	// A description of the state.
	Description *string `json:"description,omitempty" plist:"description,omitempty"`
	// A dictionary that contains additional details about the state.
	Details *map[string]any `json:"details,omitempty" plist:"details,omitempty"`
}

// A status report of passcode compliance.
type StatusPasscodeCompliance struct {
	// If `true`, the passcode is in compliance with all passcode policies set on the device. If `false`, the passcode isn't in compliance with one or more passcode policies set on the device. When there are no passcode policies on the device, this value `true`.
	PasscodeisCompliant bool `json:"passcode.is-compliant" plist:"passcode.is-compliant" required:"true"`
}

func (p *StatusPasscodeCompliance) StatusItemType() string {
	return "passcode.is-compliant"
}

// A status report of the passcode on the device.
type StatusPasscodeIsPresent struct {
	// If `true`, a passcode is present on the device. If `false`, a passcode isn't present on the device. When a passcode is present, the specific attributes of the passcode, such as length or number of complex characters, aren't reported. Instead, use the `passcode.is-compliant` status item to verify that the passcode complies with all passcode policies set on the device.
	PasscodeisPresent bool `json:"passcode.is-present" plist:"passcode.is-present" required:"true"`
}

func (p *StatusPasscodeIsPresent) StatusItemType() string {
	return "passcode.is-present"
}

// Information about connection groups with member connection references that the system couldn't resolve.
type StatusScreenSharingConnectionGroupUnresolvedConnections struct {
	// A status item that contains an array of unresolved connection groups.
	ScreensharingconnectiongroupunresolvedConnection []*UnresolvedGroup `json:"screensharing.connection.group.unresolved-connection" plist:"screensharing.connection.group.unresolved-connection" required:"true"`
}

func (p *StatusScreenSharingConnectionGroupUnresolvedConnections) StatusItemType() string {
	return "screensharing.connection.group.unresolved-connection"
}

// A status item that contains an unresolved connection group.
type UnresolvedGroup struct {
	// The unique `ConnectionGroupUUID` identifier of the connection group.
	Identifier string `json:"identifier" plist:"identifier" required:"true"`
	// If `true`, the system removed the unresolved connection group and only this key and the `identifier` key are present in the status item object.
	Removed *bool `json:"_removed,omitempty" plist:"_removed,omitempty"`
	// An array of `ConnectionUUID` values specified in the `Members` key in the group's declaration for the unresolved connections.
	UnresolvedConnections *[]string `json:"unresolved_connections,omitempty" plist:"unresolved_connections,omitempty"`
}

// A status report of the client's managed certificates.
type StatusSecurityCertificateList struct {
	// A list of the device's managed certificates.
	Securitycertificatelist []*SecuritycertificatelistStatusValue `json:"security.certificate.list" plist:"security.certificate.list" required:"true"`
}

func (p *StatusSecurityCertificateList) StatusItemType() string {
	return "security.certificate.list"
}

// A status report of a security certificate.
type SecuritycertificatelistStatusValue struct {
	// The unique identifier of the certificate which the system uses as the primary key.
	Identifier string `json:"identifier" plist:"identifier" required:"true"`
	// If `true`, the system removed the app and only this key and the `identifier` key are present in the status item object.
	Removed *bool `json:"_removed,omitempty" plist:"_removed,omitempty"`
	// The identifier of the asset declaration that installed the certificate, which is only present if a declaration installed the certificate.
	DeclarationIdentifier *string `json:"declaration-identifier,omitempty" plist:"declaration-identifier,omitempty"`
	// The summary of the certificate's subject.
	SubjectSummary string `json:"subject-summary" plist:"subject-summary" required:"true"`
	// If `true`, the certificate is an identity certificate.
	IsIdentity bool `json:"is-identity" plist:"is-identity" required:"true"`
	// The certificate data in DER-encoded X.509 format.
	Data []byte `json:"data" plist:"data" required:"true"`
}

// A status report of the device's background task details.
type StatusServicesBackgroundTask struct {
	// The background task.
	ServicesbackgroundTask []*ServicesbackgroundTaskStatusValue `json:"services.background-task" plist:"services.background-task" required:"true"`
}

func (p *StatusServicesBackgroundTask) StatusItemType() string {
	return "services.background-task"
}

// A status report of a background task.
type ServicesbackgroundTaskStatusValue struct {
	// The background task UUID which the system uses as the primary key.
	Identifier string `json:"identifier" plist:"identifier" required:"true"`
	// If `true`, the background task is removed and the status item object only contains this key and the `identifier` key.
	Removed *bool `json:"_removed,omitempty" plist:"_removed,omitempty"`
	// For types other than `agent` or `daemon`, this is the code signature designated requirement of the item, if available.
	CodeSignature *string `json:"code-signature,omitempty" plist:"code-signature,omitempty"`
	// The numeric user identifier of the owner of the background task.
	Uid int64 `json:"uid" plist:"uid" required:"true"`
	// For an `agent` or `daemon`, the path to the `launchd` `plist` file. For other types, the path to the app or the document.
	Path string `json:"path" plist:"path" required:"true"`
	// The `SMAppService.Status` enumeration.
	State ServicesbackgroundTaskStatusValueState `json:"state" plist:"state" required:"true"`
	// The daemon, agent, or SFL login item type.
	Type Type `json:"type" plist:"type" required:"true"`
	// Details about a `launchd`-based background task, which is only present when the type is `daemon` or `agent`.
	Launchd *Launchd `json:"launchd,omitempty" plist:"launchd,omitempty"`
}

// The `SMAppService.Status` enumeration.
type ServicesbackgroundTaskStatusValueState string

const (
	ServicesbackgroundTaskStatusValueStateNotRegistered    ServicesbackgroundTaskStatusValueState = "not-registered"
	ServicesbackgroundTaskStatusValueStateEnabled          ServicesbackgroundTaskStatusValueState = "enabled"
	ServicesbackgroundTaskStatusValueStateRequiresApproval ServicesbackgroundTaskStatusValueState = "requires-approval"
	ServicesbackgroundTaskStatusValueStateNotFound         ServicesbackgroundTaskStatusValueState = "not-found"
)

// The daemon, agent, or SFL login item type.
type Type string

const (
	TypeDaemon    Type = "daemon"
	TypeAgent     Type = "agent"
	TypeLoginItem Type = "login-item"
	TypeApp       Type = "app"
	TypeUserItem  Type = "user-item"
)

// Details about a `launchd`-based background task, which is only present when the type is `daemon` or `agent`.
type Launchd struct {
	// The label of the `launchd`-based background task.
	Label string `json:"label" plist:"label" required:"true"`
	// The program that the `launchd` `plist` file specifies.
	Program string `json:"program" plist:"program" required:"true"`
	// The program arguments that the `launchd` `plist` file specifies.
	ProgramArguments *[]string `json:"program-arguments,omitempty" plist:"program-arguments,omitempty"`
	// The hash value of the `launchd` `plist` file.
	Checksum string `json:"checksum" plist:"checksum" required:"true"`
	// A dictionary that indicates a `ServicesBackgroundTasks` configuration created this background task. The dictionary contains properties that identify the configuration and the declaration asset that provided the launchd plist for the task.
	DeviceManagement *DeviceManagement `json:"device-management,omitempty" plist:"device-management,omitempty"`
}

// A dictionary that indicates a `ServicesBackgroundTasks` configuration created this background task. The dictionary contains properties that identify the configuration and the declaration asset that provided the launchd plist for the task.
type DeviceManagement struct {
	// The identifier of the `ServicesBackgroundTasks` configuration that created this task.
	ConfigurationIdentifier string `json:"configuration-identifier" plist:"configuration-identifier" required:"true"`
	// The `Identifier` of the declaration asset that provided the launchd plist for this task.
	AssetIdentifier string `json:"asset-identifier" plist:"asset-identifier" required:"true"`
	// The `ServerToken` of the declaration asset that provided the launchd plist for this task.
	AssetServerToken string `json:"asset-server-token" plist:"asset-server-token" required:"true"`
}

// A status report of the device's enrolled beta program.
type StatusSoftwareUpdateBetaEnrollment struct {
	// The device's enrolled beta program name, or an empty string if there's no enrolled beta program.
	SoftwareupdatebetaEnrollment string `json:"softwareupdate.beta-enrollment" plist:"softwareupdate.beta-enrollment" required:"true"`
}

func (p *StatusSoftwareUpdateBetaEnrollment) StatusItemType() string {
	return "softwareupdate.beta-enrollment"
}

// A status report of the device's update device ID.
type StatusSoftwareUpdateDeviceID struct {
	// The device identifier to use when looking up available software updates via `https://gdmf.apple.com/v2/pmv`.
	SoftwareupdatedeviceId string `json:"softwareupdate.device-id" plist:"softwareupdate.device-id" required:"true"`
}

func (p *StatusSoftwareUpdateDeviceID) StatusItemType() string {
	return "softwareupdate.device-id"
}

// A status report of a software update failure reason.
type StatusSoftwareUpdateFailureReason struct {
	// Details about a software update failure.
	SoftwareupdatefailureReason SoftwareupdatefailureReason `json:"softwareupdate.failure-reason" plist:"softwareupdate.failure-reason" required:"true"`
}

func (p *StatusSoftwareUpdateFailureReason) StatusItemType() string {
	return "softwareupdate.failure-reason"
}

// Details about a software update failure.
type SoftwareupdatefailureReason struct {
	// The number of times the current software update failed. If there are no failures, or no pending software update, this is `0`.
	Count int64 `json:"count" plist:"count" required:"true"`
	// If present, this describes the reason for last software update failure. This key isn't present if there are no failures or no pending software update.
	Reason *string `json:"reason,omitempty" plist:"reason,omitempty"`
	// If present, this is the RFC 3339 timestamp of the last software update failure. This key isn't present if there are no failures or no pending software update.
	Timestamp *string `json:"timestamp,omitempty" plist:"timestamp,omitempty"`
}

// A status report of the reason for a pending software update on the device.
type StatusSoftwareUpdateInstallReason struct {
	// Details about the reason for a pending software update.
	SoftwareupdateinstallReason SoftwareupdateinstallReason `json:"softwareupdate.install-reason" plist:"softwareupdate.install-reason" required:"true"`
}

func (p *StatusSoftwareUpdateInstallReason) StatusItemType() string {
	return "softwareupdate.install-reason"
}

// Details about the reason for a pending software update.
type SoftwareupdateinstallReason struct {
	// A list of reasons for the pending software update. An empty list indicates that no software update is pending.
	Reason []Reason `json:"reason" plist:"reason" required:"true"`
	// The identifier of the declaration that caused the software update to occur. This key is present only if the `reason` array contains the `declaration` value.
	DeclarationId *string `json:"declaration-id,omitempty" plist:"declaration-id,omitempty"`
}

// A list of reasons for the pending software update. An empty list indicates that no software update is pending.
// The software update install reason state:
// * system-settings - software update was triggered via Settings.app
// * install-tonight - software update was triggered via install tonight action
// * auto-update - software update was triggered via an automatic update
// * notification - software update was triggered via user notification action
// * setup-assistant - software update was triggered via Setup Assistant
// * command-line - software update was triggered via `softwareupdate` command line tool
// * mdm - software update was triggered via an MDM command
// * declaration - software update was triggered via a declarative device management configuration
type Reason string

const (
	ReasonSystemSettings Reason = "system-settings"
	ReasonInstallTonight Reason = "install-tonight"
	ReasonAutoUpdate     Reason = "auto-update"
	ReasonNotification   Reason = "notification"
	ReasonSetupAssistant Reason = "setup-assistant"
	ReasonCommandLine    Reason = "command-line"
	ReasonMdm            Reason = "mdm"
	ReasonDeclaration    Reason = "declaration"
)

// A status report of the software update install state.
type StatusSoftwareUpdateInstallState struct {
	// The software update install status, which has the following values:
	// - `none`: There's no software update pending, and any previous software update succeeded.
	// - `waiting': A software update is waiting to start.
	// - `downloading`: The system is downloading data for a software update.
	// - `prepared`: The system prepared the software update and it's ready for installation.
	// - `installing`: The system is installing the software update.
	// - `failed`: The software update failed.
	SoftwareupdateinstallState SoftwareupdateinstallState `json:"softwareupdate.install-state" plist:"softwareupdate.install-state" required:"true"`
}

func (p *StatusSoftwareUpdateInstallState) StatusItemType() string {
	return "softwareupdate.install-state"
}

// The software update install status, which has the following values:
// - `none`: There's no software update pending, and any previous software update succeeded.
// - `waiting': A software update is waiting to start.
// - `downloading`: The system is downloading data for a software update.
// - `prepared`: The system prepared the software update and it's ready for installation.
// - `installing`: The system is installing the software update.
// - `failed`: The software update failed.
type SoftwareupdateinstallState string

const (
	SoftwareupdateinstallStateNone        SoftwareupdateinstallState = "none"
	SoftwareupdateinstallStateDownloading SoftwareupdateinstallState = "downloading"
	SoftwareupdateinstallStatePrepared    SoftwareupdateinstallState = "prepared"
	SoftwareupdateinstallStateInstalling  SoftwareupdateinstallState = "installing"
	SoftwareupdateinstallStateFailed      SoftwareupdateinstallState = "failed"
)

// A status report of the pending software update version.
type StatusSoftwareUpdatePendingVersion struct {
	// A dictionary that contains the build and OS versions of the software update that's pending on the device.
	SoftwareupdatependingVersion SoftwareupdatependingVersion `json:"softwareupdate.pending-version" plist:"softwareupdate.pending-version" required:"true"`
}

func (p *StatusSoftwareUpdatePendingVersion) StatusItemType() string {
	return "softwareupdate.pending-version"
}

// A dictionary that contains the build and OS versions of the software update that's pending on the device.
type SoftwareupdatependingVersion struct {
	// The OS version of the pending software update, including any rapid security response version. This string is empty if no update is pending.
	OsVersion string `json:"os-version" plist:"os-version" required:"true"`
	// The build version of the pending software update, including any rapid security response version. This string is empty if no update is pending.
	BuildVersion string `json:"build-version" plist:"build-version" required:"true"`
	// The local date time value that indicates when the pending software update will be installed. This key is only present when the pending software update is being enforced.
	TargetLocalDateTime *string `json:"target-local-date-time,omitempty" plist:"target-local-date-time,omitempty"`
}

// Provides details about an error for an item in a status report.
type StatusReason struct {
	// The error code for this error.
	Code string `json:"code" plist:"code" required:"true"`
	// A description of this error.
	Description *string `json:"description,omitempty" plist:"description,omitempty"`
	// A dictionary that contains additional details about the error.
	Details *Details `json:"details,omitempty" plist:"details,omitempty"`
}

func (p *StatusReason) DeclarationType() string {
	return "status-reason"
}

// A dictionary that contains additional details about the error.
type Details struct{}

// A test status item for an array.
type StatusTestArrayValue struct {
	// The test status item array value.
	TestarrayValue []*TestarrayValueStatusValue `json:"test.array-value" plist:"test.array-value" required:"true"`
}

func (p *StatusTestArrayValue) StatusItemType() string {
	return "test.array-value"
}

// A status value for the test status item array.
type TestarrayValueStatusValue struct {
	// The value of the first sub-key.
	Key1 string `json:"key1" plist:"key1" required:"true"`
	// The value of the second sub-key.
	Key2 *string `json:"key2,omitempty" plist:"key2,omitempty"`
}

// A test status item for a Boolean value.
type StatusTestBooleanValue struct {
	// The test status Boolean value.
	TestbooleanValue bool `json:"test.boolean-value" plist:"test.boolean-value" required:"true"`
}

func (p *StatusTestBooleanValue) StatusItemType() string {
	return "test.boolean-value"
}

// A test status item for a dictionary.
type StatusTestDictionaryValue struct {
	// The test status dictionary value.
	TestdictionaryValue TestdictionaryValue `json:"test.dictionary-value" plist:"test.dictionary-value" required:"true"`
}

func (p *StatusTestDictionaryValue) StatusItemType() string {
	return "test.dictionary-value"
}

// The test status dictionary value.
type TestdictionaryValue struct {
	// The value of the first sub-key.
	Key1 string `json:"key1" plist:"key1" required:"true"`
	// The value of the second sub-key.
	Key2 *string `json:"key2,omitempty" plist:"key2,omitempty"`
}

// A test status item for an error.
type StatusTestErrorValue struct {
	// The test status error value.
	TesterrorValue string `json:"test.error-value" plist:"test.error-value" required:"true"`
}

func (p *StatusTestErrorValue) StatusItemType() string {
	return "test.error-value"
}

// A test status item for an integer.
type StatusTestIntegerValue struct {
	// The test status integer value.
	TestintegerValue int64 `json:"test.integer-value" plist:"test.integer-value" required:"true"`
}

func (p *StatusTestIntegerValue) StatusItemType() string {
	return "test.integer-value"
}

// A test status item for a real value.
type StatusTestRealValue struct {
	// The test status real value.
	TestrealValue float64 `json:"test.real-value" plist:"test.real-value" required:"true"`
}

func (p *StatusTestRealValue) StatusItemType() string {
	return "test.real-value"
}

// A test status item for a string.
type StatusTestStringValue struct {
	// The test status string value.
	TeststringValue string `json:"test.string-value" plist:"test.string-value" required:"true"`
}

func (p *StatusTestStringValue) StatusItemType() string {
	return "test.string-value"
}
