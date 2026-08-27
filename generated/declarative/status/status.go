// DO NOT EDIT
// generated from https://github.com/apple/device-management.git:f878dea98fb88293a3686e44bcfb891f8e78f98f/declarative/status

package status

const DeviceManagementGenerateHash = "f878dea98fb88293a3686e44bcfb891f8e78f98f"

// StatusItemTypes lists every DDM status item these types can represent, by
// Apple's dotted status item identifier.
//
// Generated from the statusitemtype of each schema under declarative/status/.
// Each identifier corresponds to a field reachable from StatusItems -- the field's
// doc comment names the item it carries.
//
// These are the identifiers the protocol uses as string values: the items named by
// a status subscription, a status report's Errors[].StatusItem, and the client's
// management.client-capabilities.supported-payloads.status-items.
var StatusItemTypes = []string{
	"account.list.caldav",
	"account.list.carddav",
	"account.list.exchange",
	"account.list.google",
	"account.list.ldap",
	"account.list.mail.incoming",
	"account.list.mail.outgoing",
	"account.list.subscribed-calendar",
	"app.managed.list",
	"device.identifier.serial-number",
	"device.identifier.udid",
	"device.model.family",
	"device.model.identifier",
	"device.model.marketing-name",
	"device.model.number",
	"device.operating-system.build-version",
	"device.operating-system.family",
	"device.operating-system.marketing-name",
	"device.operating-system.supplemental.build-version",
	"device.operating-system.supplemental.extra-version",
	"device.operating-system.version",
	"device.power.battery-health",
	"diskmanagement.filevault.enabled",
	"management.client-capabilities",
	"management.declarations",
	"mdm.app",
	"package.list",
	"passcode.is-compliant",
	"passcode.is-present",
	"screensharing.connection.group.unresolved-connection",
	"security.certificate.list",
	"services.background-task",
	"softwareupdate.beta-enrollment",
	"softwareupdate.device-id",
	"softwareupdate.failure-reason",
	"softwareupdate.install-reason",
	"softwareupdate.install-state",
	"softwareupdate.pending-version",
	"test.array-value",
	"test.boolean-value",
	"test.dictionary-value",
	"test.error-value",
	"test.integer-value",
	"test.real-value",
	"test.string-value",
}
var DeclarationMap = map[string]any{"status-reason": StatusReason{}}

// Provides details about an error for an item in a status report.
type StatusReason struct {
	// The error code for this error.
	Code string `json:"code" required:"true"`
	// A description of this error.
	Description *string `json:"description,omitempty"`
	// A dictionary that contains additional details about the error.
	Details *map[string]any `json:"details,omitempty"`
}

func (p *StatusReason) DeclarationType() string {
	return "status-reason"
}

// Status sent by the client.
type StatusReport struct {
	// The status items for this report.
	StatusItems StatusItems `json:"StatusItems" required:"true"`
	// An array of errors for this status report.
	Errors []*StatusReportErrors `json:"Errors" required:"true"`
	// The system sets this to `true` to indicate that the status report contains the full set of current status, and is not an incremental report. A full status report includes the full set of items in any status array item, not just the changes. Servers use this to replace their entire status for the device, rather than do an incremental update to the existing status. The system sets this to `true` when sending a "safety sync" status report, which is typically sent every 24 hours or so.
	FullReport *bool `json:"FullReport,omitempty"`
}

// The status items for this report.
type StatusItems struct {
	// Status items under the `account` namespace.
	Account *StatusItemsAccount `json:"account,omitempty"`
	// Status items under the `app` namespace.
	App *StatusItemsApp `json:"app,omitempty"`
	// Status items under the `device` namespace.
	Device *StatusItemsDevice `json:"device,omitempty"`
	// Status items under the `diskmanagement` namespace.
	DiskManagement *StatusItemsDiskManagement `json:"diskmanagement,omitempty"`
	// Status items under the `management` namespace.
	Management *StatusItemsManagement `json:"management,omitempty"`
	// Status items under the `mdm` namespace.
	MDM *StatusItemsMDM `json:"mdm,omitempty"`
	// Status items under the `package` namespace.
	Package *StatusItemsPackage `json:"package,omitempty"`
	// Status items under the `passcode` namespace.
	Passcode *StatusItemsPasscode `json:"passcode,omitempty"`
	// Status items under the `screensharing` namespace.
	ScreenSharing *StatusItemsScreenSharing `json:"screensharing,omitempty"`
	// Status items under the `security` namespace.
	Security *StatusItemsSecurity `json:"security,omitempty"`
	// Status items under the `services` namespace.
	Services *StatusItemsServices `json:"services,omitempty"`
	// Status items under the `softwareupdate` namespace.
	SoftwareUpdate *StatusItemsSoftwareUpdate `json:"softwareupdate,omitempty"`
	// Status items under the `test` namespace.
	Test *StatusItemsTest `json:"test,omitempty"`
}

// Status items under the `account` namespace.
type StatusItemsAccount struct {
	// Status items under the `account.list` namespace.
	List *StatusItemsAccountList `json:"list,omitempty"`
}

// Status items under the `account.list` namespace.
type StatusItemsAccountList struct {
	// A list of status values for the Calendar accounts.
	// Status item: `account.list.caldav`.
	CalDAV *[]*AccountListCalDAV `json:"caldav,omitempty"`
	// A list of status values for the Contacts accounts.
	// Status item: `account.list.carddav`.
	CardDAV *[]*AccountListCardDAV `json:"carddav,omitempty"`
	// A list of status values for the Exchange accounts.
	// Status item: `account.list.exchange`.
	Exchange *[]*AccountListExchange `json:"exchange,omitempty"`
	// A list of status values for the Google accounts.
	// Status item: `account.list.google`.
	Google *[]*AccountListGoogle `json:"google,omitempty"`
	// A list of status values for the LDAP accounts.
	// Status item: `account.list.ldap`.
	LDAP *[]*AccountListLDAP `json:"ldap,omitempty"`
	// Status items under the `account.list.mail` namespace.
	Mail *StatusItemsAccountListMail `json:"mail,omitempty"`
	// A list of status values for the subscribed calendars.
	// Status item: `account.list.subscribed-calendar`.
	SubscribedCalendar *[]*AccountListSubscribedCalendar `json:"subscribed-calendar,omitempty"`
}

// A status report of the client's Calendar account details.
// Status item: `account.list.caldav`.
type AccountListCalDAV struct {
	// The unique identifier for the account.
	Identifier string `json:"identifier" required:"true"`
	// If `true`, the account is removed and the status item object only contains this key and the `identifier` key.
	Removed *bool `json:"_removed,omitempty"`
	// The identifier of the declaration that installed the account. Only present if a declaration installed the account.
	DeclarationIdentifier *string `json:"declaration-identifier,omitempty"`
	// The name of the account.
	VisibleName *string `json:"visible-name,omitempty"`
	// The server host name for the account.
	Hostname *string `json:"hostname,omitempty"`
	// The server port for the account.
	Port *int64 `json:"port,omitempty"`
	// The user name for the account.
	Username *string `json:"username,omitempty"`
	// If `true`, the Calendar app is displaying calendars and events for the account.
	AreCalendarsEnabled *bool `json:"are-calendars-enabled,omitempty"`
	// If `true`, the Reminders app is displaying reminders for the account.
	AreRemindersEnabled *bool `json:"are-reminders-enabled,omitempty"`
}

// A status report of the client's Contacts account details.
// Status item: `account.list.carddav`.
type AccountListCardDAV struct {
	// The unique identifier for the account.
	Identifier string `json:"identifier" required:"true"`
	// If `true`, the account is removed and the status item object only contains this key and the `identifier` key.
	Removed *bool `json:"_removed,omitempty"`
	// The identifier of the declaration that installed the account. Only present if a declaration installed the account.
	DeclarationIdentifier *string `json:"declaration-identifier,omitempty"`
	// The name of the account.
	VisibleName *string `json:"visible-name,omitempty"`
	// The server host name for the account.
	Hostname *string `json:"hostname,omitempty"`
	// The server port for the account.
	Port *int64 `json:"port,omitempty"`
	// The user name for the account.
	Username *string `json:"username,omitempty"`
}

// A status report of the client's Exchange account details.
// Status item: `account.list.exchange`.
type AccountListExchange struct {
	// The unique identifier for the account.
	Identifier string `json:"identifier" required:"true"`
	// If `true`, the account is removed and the status item object only contains this key and the `identifier` key.
	Removed *bool `json:"_removed,omitempty"`
	// The identifier of the declaration that installed the account. Only present if a declaration installed the account.
	DeclarationIdentifier *string `json:"declaration-identifier,omitempty"`
	// The name of the account.
	VisibleName *string `json:"visible-name,omitempty"`
	// The server host name for the account.
	Hostname *string `json:"hostname,omitempty"`
	// The server port for the account.
	Port *int64 `json:"port,omitempty"`
	// The user name for the account.
	Username *string `json:"username,omitempty"`
	// A Boolean value that indicates whether the Mail app displays mail for this account.
	IsMailEnabled *bool `json:"is-mail-enabled,omitempty"`
	// A Boolean value that indicates whether the Calendar app displays calendars and events for this account.
	AreCalendarsEnabled *bool `json:"are-calendars-enabled,omitempty"`
	// A Boolean value that indicates whether the Contacts app displays contacts for this account.
	AreContactsEnabled *bool `json:"are-contacts-enabled,omitempty"`
	// A Boolean value that indicates whether the Notes app displays notes for this account.
	AreNotesEnabled *bool `json:"are-notes-enabled,omitempty"`
	// A Boolean value that indicates whether the Reminders app displays reminders for this account.
	AreRemindersEnabled *bool `json:"are-reminders-enabled,omitempty"`
}

// A status report of the client's Google account details.
// Status item: `account.list.google`.
type AccountListGoogle struct {
	// The unique identifier for the account.
	Identifier string `json:"identifier" required:"true"`
	// If `true`, the account is removed and the status item object only contains this key and the `identifier` key.
	Removed *bool `json:"_removed,omitempty"`
	// The identifier of the declaration that installed the account. Only present if a declaration installed the account.
	DeclarationIdentifier *string `json:"declaration-identifier,omitempty"`
	// The name of the account.
	VisibleName *string `json:"visible-name,omitempty"`
	// The user name for the account.
	Username *string `json:"username,omitempty"`
	// A Boolean value that indicates whether the Mail app displays mail for this account.
	IsMailEnabled *bool `json:"is-mail-enabled,omitempty"`
	// A Boolean value that indicates whether the Calendar app displays calendars and events for this account.
	AreCalendarsEnabled *bool `json:"are-calendars-enabled,omitempty"`
	// A Boolean value that indicates whether the Contacts app displays contacts for this account.
	AreContactsEnabled *bool `json:"are-contacts-enabled,omitempty"`
	// A Boolean value that indicates whether the Notes app displays notes for this account.
	AreNotesEnabled *bool `json:"are-notes-enabled,omitempty"`
}

// A status report of the client's LDAP account details.
// Status item: `account.list.ldap`.
type AccountListLDAP struct {
	// The unique identifier for the account.
	Identifier string `json:"identifier" required:"true"`
	// If `true`, the account is removed and the status item object only contains this key and the `identifier` key.
	Removed *bool `json:"_removed,omitempty"`
	// The identifier of the declaration that installed the account. Only present if a declaration installed the account.
	DeclarationIdentifier *string `json:"declaration-identifier,omitempty"`
	// The name of the account.
	VisibleName *string `json:"visible-name,omitempty"`
	// The server host name for the account.
	Hostname *string `json:"hostname,omitempty"`
	// The server port for the account.
	Port *int64 `json:"port,omitempty"`
	// The user name for the account.
	Username *string `json:"username,omitempty"`
	// A Boolean value that indicates whether the account is enabled for use with the Contacts app.
	IsEnabled *bool `json:"is-enabled,omitempty"`
}

// Status items under the `account.list.mail` namespace.
type StatusItemsAccountListMail struct {
	// A list of status values for the incoming Mail accounts.
	// Status item: `account.list.mail.incoming`.
	Incoming *[]*AccountListMailIncoming `json:"incoming,omitempty"`
	// A list of status values for the outgoing Mail accounts.
	// Status item: `account.list.mail.outgoing`.
	Outgoing *[]*AccountListMailOutgoing `json:"outgoing,omitempty"`
}

// A status report of the client's incoming Mail account details.
// Status item: `account.list.mail.incoming`.
type AccountListMailIncoming struct {
	// The unique identifier for the account.
	Identifier string `json:"identifier" required:"true"`
	// If `true`, the account is removed and the status item object only contains this key and the `identifier` key.
	Removed *bool `json:"_removed,omitempty"`
	// The identifier of the declaration that installed the account. Only present if a declaration installed the account.
	DeclarationIdentifier *string `json:"declaration-identifier,omitempty"`
	// The name of the account.
	VisibleName *string `json:"visible-name,omitempty"`
	// The server host name for the account.
	Hostname *string `json:"hostname,omitempty"`
	// The server port for the account.
	Port *int64 `json:"port,omitempty"`
	// The user name for the account.
	Username *string `json:"username,omitempty"`
	// A Boolean value that indicates whether the Mail app displays mail for this account.
	IsMailEnabled *bool `json:"is-mail-enabled,omitempty"`
	// A Boolean value that indicates whether the Notes app displays notes for this account.
	AreNotesEnabled *bool `json:"are-notes-enabled,omitempty"`
}

// A status report of the client's outgoing Mail account details.
// Status item: `account.list.mail.outgoing`.
type AccountListMailOutgoing struct {
	// The unique identifier for the account.
	Identifier string `json:"identifier" required:"true"`
	// If `true`, the account is removed and the status item object only contains this key and the `identifier` key.
	Removed *bool `json:"_removed,omitempty"`
	// The identifier of the declaration that installed the account. Only present if a declaration installed the account.
	DeclarationIdentifier *string `json:"declaration-identifier,omitempty"`
	// The name of the account.
	VisibleName *string `json:"visible-name,omitempty"`
	// The server host name for the account.
	Hostname *string `json:"hostname,omitempty"`
	// The server port for the account.
	Port *int64 `json:"port,omitempty"`
	// The user name for the account.
	Username *string `json:"username,omitempty"`
}

// A status report of the client's subscribed calendar details.
// Status item: `account.list.subscribed-calendar`.
type AccountListSubscribedCalendar struct {
	// The unique identifier for the subscribed calendar.
	Identifier string `json:"identifier" required:"true"`
	// If `true`, the subscribed calendar is removed and the status item object only contains this key and the `identifier` key.
	Removed *bool `json:"_removed,omitempty"`
	// The identifier of the declaration that installed the subscribed calendar. Only present if a declaration installed the subscribed calendar.
	DeclarationIdentifier *string `json:"declaration-identifier,omitempty"`
	// The name of the subscribed calendar.
	VisibleName *string `json:"visible-name,omitempty"`
	// The URL of the subscribed calendar.
	CalendarURL *string `json:"calendar-url,omitempty"`
	// The user name for authenticating with the subscribed calendar.
	Username *string `json:"username,omitempty"`
	// A Boolean value that indicates whether the Calendar app displays this subscribed calendar.
	IsEnabled *bool `json:"is-enabled,omitempty"`
}

// Status items under the `app` namespace.
type StatusItemsApp struct {
	// Status items under the `app.managed` namespace.
	Managed *StatusItemsAppManaged `json:"managed,omitempty"`
}

// Status items under the `app.managed` namespace.
type StatusItemsAppManaged struct {
	// An array of dictionaries that describe the device's declarative managed apps.
	// Status item: `app.managed.list`.
	List *[]*AppManagedList `json:"list,omitempty"`
}

// A dictionary that describes a declarative managed app.
// Status item: `app.managed.list`.
type AppManagedList struct {
	// The app's bundle id, which is unique.
	Identifier string `json:"identifier" required:"true"`
	// If `true`, the system removed the app and only this key and the `identifier` key are present in the status item object.
	Removed *bool `json:"_removed,omitempty"`
	// The identifier of the declaration that controls the app.
	DeclarationIdentifier *string `json:"declaration-identifier,omitempty"`
	// The name of the app.
	Name *string `json:"name,omitempty"`
	// The app's external version identifier. You can also retrieve this value from the App Store. For more information, see `Apps and Books for Organizations`.
	// If the current external version identifier of an app on the App Store doesn't match the external version identifier reported by the device, there may be an app update available for the device.
	ExternalVersionID *int64 `json:"external-version-id,omitempty"`
	// The version of the app.
	Version *string `json:"version,omitempty"`
	// The short version of the app.
	ShortVersion *string `json:"short-version,omitempty"`
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
	State *AppManagedListState `json:"state,omitempty"`
	// The update status of the app, which has the following possible values:
	// - `available`: An update is available for the app.
	// - `prompting-for-update`: The system is displaying a prompt to the user to proceed with app installation.
	// - `prompting-for-update-login`: The system is displaying an App Store sign-in prompt to the user to allow app installation.
	// - `updating`: The app is updating.
	// - `failed`: The app update failed.
	// > Note:
	// > This key is only present if `state` is `managed` and an update is available.
	UpdateState *AppManagedListUpdateState `json:"update-state,omitempty"`
	// The status of app or extension managed configurations. This key is only present when managed configurations are available for the managed app or any of its extensions.
	ConfigState *AppManagedListConfigState `json:"config-state,omitempty"`
	// An array that contains additional details about the app state, including errors.
	Reasons *[]*AppManagedListReasons `json:"reasons,omitempty"`
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
type AppManagedListState string

const (
	AppManagedListStateOptional               AppManagedListState = "optional"
	AppManagedListStateQueued                 AppManagedListState = "queued"
	AppManagedListStateNotPresent             AppManagedListState = "not-present"
	AppManagedListStatePromptingForConsent    AppManagedListState = "prompting-for-consent"
	AppManagedListStatePromptingForLogin      AppManagedListState = "prompting-for-login"
	AppManagedListStatePromptingForManagement AppManagedListState = "prompting-for-management"
	AppManagedListStateDownloading            AppManagedListState = "downloading"
	AppManagedListStateInstalling             AppManagedListState = "installing"
	AppManagedListStateManaged                AppManagedListState = "managed"
	AppManagedListStateManagedButUninstalled  AppManagedListState = "managed-but-uninstalled"
	AppManagedListStateFailed                 AppManagedListState = "failed"
)

// The update status of the app, which has the following possible values:
// - `available`: An update is available for the app.
// - `prompting-for-update`: The system is displaying a prompt to the user to proceed with app installation.
// - `prompting-for-update-login`: The system is displaying an App Store sign-in prompt to the user to allow app installation.
// - `updating`: The app is updating.
// - `failed`: The app update failed.
// > Note:
// > This key is only present if `state` is `managed` and an update is available.
type AppManagedListUpdateState string

const (
	AppManagedListUpdateStateAvailable               AppManagedListUpdateState = "available"
	AppManagedListUpdateStatePromptingForUpdate      AppManagedListUpdateState = "prompting-for-update"
	AppManagedListUpdateStatePromptingForUpdateLogin AppManagedListUpdateState = "prompting-for-update-login"
	AppManagedListUpdateStateUpdating                AppManagedListUpdateState = "updating"
	AppManagedListUpdateStateFailed                  AppManagedListUpdateState = "failed"
)

// The status of app or extension managed configurations. This key is only present when managed configurations are available for the managed app or any of its extensions.
type AppManagedListConfigState struct {
	// The status of any app managed configuration. This key is only present when the managed app has a managed configuration.
	AppConfigState *AppManagedListConfigStateAppConfigState `json:"app-config-state,omitempty"`
	// The status of any app extension managed configuration. This key's value is a dictionary whose keys are the bundle identifiers of app extensions that have a managed configuration. The values of each key represent the status of the corresponding app extension's managed configuration.
	ExtensionConfigState *map[string]AppManagedListConfigStateExtensionConfigStateManagedConfigurationState `json:"extension-config-state,omitempty"`
}

// The status of any app managed configuration. This key is only present when the managed app has a managed configuration.
type AppManagedListConfigStateAppConfigState struct {
	// The managed configuration status.
	// - `unknown`: The managed configuration has not been read
	// - `invalid`: The managed configuration was read and deemed to be invalid
	// - `valid`: The managed configuration was read and deemed to be valid
	State AppManagedListConfigStateAppConfigStateState `json:"state" required:"true"`
}

// The managed configuration status.
// - `unknown`: The managed configuration has not been read
// - `invalid`: The managed configuration was read and deemed to be invalid
// - `valid`: The managed configuration was read and deemed to be valid
type AppManagedListConfigStateAppConfigStateState string

const (
	AppManagedListConfigStateAppConfigStateStateUnknown AppManagedListConfigStateAppConfigStateState = "unknown"
	AppManagedListConfigStateAppConfigStateStateInvalid AppManagedListConfigStateAppConfigStateState = "invalid"
	AppManagedListConfigStateAppConfigStateStateValid   AppManagedListConfigStateAppConfigStateState = "valid"
)

// The bundle identifier of the managed app extension.
type AppManagedListConfigStateExtensionConfigStateManagedConfigurationState struct {
	// The managed configuration status.
	// - `unknown`: The managed configuration has not been read
	// - `invalid`: The managed configuration was read and deemed to be invalid
	// - `valid`: The managed configuration was read and deemed to be valid
	State AppManagedListConfigStateAppConfigStateState `json:"state" required:"true"`
}

// Information about a status error.
type AppManagedListReasons struct {
	// A code for the state.
	Code string `json:"code" required:"true"`
	// A description of the state.
	Description *string `json:"description,omitempty"`
	// A dictionary that contains additional details about the state.
	Details *map[string]any `json:"details,omitempty"`
}

// Status items under the `device` namespace.
type StatusItemsDevice struct {
	// Status items under the `device.identifier` namespace.
	Identifier *StatusItemsDeviceIdentifier `json:"identifier,omitempty"`
	// Status items under the `device.model` namespace.
	Model *StatusItemsDeviceModel `json:"model,omitempty"`
	// Status items under the `device.operating-system` namespace.
	OperatingSystem *StatusItemsDeviceOperatingSystem `json:"operating-system,omitempty"`
	// Status items under the `device.power` namespace.
	Power *StatusItemsDevicePower `json:"power,omitempty"`
}

// Status items under the `device.identifier` namespace.
type StatusItemsDeviceIdentifier struct {
	// The device's serial number.
	// Status item: `device.identifier.serial-number`.
	SerialNumber *string `json:"serial-number,omitempty"`
	// The device's UDID. This value is always available on the device channel. This value is only available on user channels whose organization matches that of the device channel.
	// Status item: `device.identifier.udid`.
	UDID *string `json:"udid,omitempty"`
}

// Status items under the `device.model` namespace.
type StatusItemsDeviceModel struct {
	// The hardware family of the device, such as `Mac`, `iPhone`, or `iPad`.
	// Status item: `device.model.family`.
	Family *string `json:"family,omitempty"`
	// A two-part string that specifies the device's model. The first part specifies device's model family, and the second part specifies the model's version. The model's version is a comma-separated number where the first part of the number is the version, and the second part is a variant, such as `MacBookPro15,1` or `iPhone13,2`.
	// Status item: `device.model.identifier`.
	Identifier *string `json:"identifier,omitempty"`
	// The device's marketing name, such as `iPhone 12`. This value may not always be available.
	// Status item: `device.model.marketing-name`.
	MarketingName *string `json:"marketing-name,omitempty"`
	// The device's model number.
	// Status item: `device.model.number`.
	Number *string `json:"number,omitempty"`
}

// Status items under the `device.operating-system` namespace.
type StatusItemsDeviceOperatingSystem struct {
	// The operating system's build version on the device, such as `18F132`.
	// Status item: `device.operating-system.build-version`.
	BuildVersion *string `json:"build-version,omitempty"`
	// The operating system family in use on the device, such as `macOS` or `iOS`.
	// Status item: `device.operating-system.family`.
	Family *string `json:"family,omitempty"`
	// The operating system's marketing name in use on the device, such as `Catalina`.
	// Status item: `device.operating-system.marketing-name`.
	MarketingName *string `json:"marketing-name,omitempty"`
	// Status items under the `device.operating-system.supplemental` namespace.
	Supplemental *StatusItemsDeviceOperatingSystemSupplemental `json:"supplemental,omitempty"`
	// The operating system's version in use on the device, such as `15.0`.
	// Status item: `device.operating-system.version`.
	Version *string `json:"version,omitempty"`
}

// Status items under the `device.operating-system.supplemental` namespace.
type StatusItemsDeviceOperatingSystemSupplemental struct {
	// The operating system's build and Background Security Improvement versions in use on the device, for example, `20A123a` or `20B27c`.
	// Status item: `device.operating-system.supplemental.build-version`.
	BuildVersion *string `json:"build-version,omitempty"`
	// The operating system's Background Security Improvement version in use on the device, for example, `a`.
	// Status item: `device.operating-system.supplemental.extra-version`.
	ExtraVersion *string `json:"extra-version,omitempty"`
}

// Status items under the `device.power` namespace.
type StatusItemsDevicePower struct {
	// The battery health status, which has the following values:
	// - `non-genuine`: The battery isn't a genuine Apple battery.
	// - `normal`: The battery is operating normally.
	// - `service-recommended`: The system recommends battery service.
	// - `unknown`: The system couldn't determine battery health information.
	// - `unsupported`: The device doesn't support battery health reporting.
	// Available in iOS 17 and later on iPhone, iPadOS 18.4 and later on supported iPad models, and macOS 14.4 and later on a Mac with Apple silicon.
	// Status item: `device.power.battery-health`.
	BatteryHealth *DevicePowerBatteryHealth `json:"battery-health,omitempty"`
}

// The battery health status, which has the following values:
// - `non-genuine`: The battery isn't a genuine Apple battery.
// - `normal`: The battery is operating normally.
// - `service-recommended`: The system recommends battery service.
// - `unknown`: The system couldn't determine battery health information.
// - `unsupported`: The device doesn't support battery health reporting.
// Available in iOS 17 and later on iPhone, iPadOS 18.4 and later on supported iPad models, and macOS 14.4 and later on a Mac with Apple silicon.
// Status item: `device.power.battery-health`.
type DevicePowerBatteryHealth string

const (
	DevicePowerBatteryHealthNonGenuine         DevicePowerBatteryHealth = "non-genuine"
	DevicePowerBatteryHealthNormal             DevicePowerBatteryHealth = "normal"
	DevicePowerBatteryHealthServiceRecommended DevicePowerBatteryHealth = "service-recommended"
	DevicePowerBatteryHealthUnknown            DevicePowerBatteryHealth = "unknown"
	DevicePowerBatteryHealthUnsupported        DevicePowerBatteryHealth = "unsupported"
)

// Status items under the `diskmanagement` namespace.
type StatusItemsDiskManagement struct {
	// Status items under the `diskmanagement.filevault` namespace.
	FileVault *StatusItemsDiskManagementFileVault `json:"filevault,omitempty"`
}

// Status items under the `diskmanagement.filevault` namespace.
type StatusItemsDiskManagementFileVault struct {
	// A Boolean value that specifies the File Vault enabled status on the device.
	// Status item: `diskmanagement.filevault.enabled`.
	Enabled *bool `json:"enabled,omitempty"`
}

// Status items under the `management` namespace.
type StatusItemsManagement struct {
	// An object that contains the client's protocol capabilities. These typically only change when the device upgrades its software. An implicit status subscription for this status item is always present, so the client always reports changes to the server.
	// Status item: `management.client-capabilities`.
	ClientCapabilities *ManagementClientCapabilities `json:"client-capabilities,omitempty"`
	// A collection of the client's processed declarations.
	// Status item: `management.declarations`.
	Declarations *ManagementDeclarations `json:"declarations,omitempty"`
}

// An object that contains the client's protocol capabilities. These typically only change when the device upgrades its software. An implicit status subscription for this status item is always present, so the client always reports changes to the server.
// Status item: `management.client-capabilities`.
type ManagementClientCapabilities struct {
	// A list of protocol versions that the client supports.
	SupportedVersions []string `json:"supported-versions" required:"true"`
	// A set of optional protocol features that the client supports. Each object's key represents a feature, and the property value represents the feature's associated parameters.
	SupportedFeatures map[string]any `json:"supported-features" required:"true"`
	// A set of declaration and status items that the client supports.
	SupportedPayloads ManagementClientCapabilitiesSupportedPayloads `json:"supported-payloads" required:"true"`
}

// A set of declaration and status items that the client supports.
type ManagementClientCapabilitiesSupportedPayloads struct {
	// A set of declarations that the client supports.
	Declarations ManagementClientCapabilitiesSupportedPayloadsDeclarations `json:"declarations" required:"true"`
	// A list of status items that the client supports.
	StatusItems []string `json:"status-items" required:"true"`
}

// A set of declarations that the client supports.
type ManagementClientCapabilitiesSupportedPayloadsDeclarations struct {
	// An array of strings that represents the activation types that the client supports.
	Activations *[]string `json:"activations,omitempty"`
	// An array of strings that represents the assets that the client supports.
	Assets *[]string `json:"assets,omitempty"`
	// An array of strings that represents the configuration types that the client supports.
	Configurations *[]string `json:"configurations,omitempty"`
	// An array of strings that represents the declaration types that the client supports.
	Management *[]string `json:"management,omitempty"`
}

// A collection of the client's processed declarations.
// Status item: `management.declarations`.
type ManagementDeclarations struct {
	// An array of declarations that represent the client's processed activation types.
	Activations []*ManagementDeclarationsActivations `json:"activations" required:"true"`
	// An array of declarations that represent the client's processed configuration types.
	Configurations []*ManagementDeclarationsConfigurations `json:"configurations" required:"true"`
	// An array of declarations that represent the client's processed assets.
	Assets []*ManagementDeclarationsAssets `json:"assets" required:"true"`
	// An array of declarations that represent the client's processed declaration types.
	Management []*ManagementDeclarationsManagement `json:"management" required:"true"`
}

// Status for a declaration processed by the client.
type ManagementDeclarationsActivations struct {
	// The `identifier` of the declaration this status report refers to.
	Identifier string `json:"identifier" required:"true"`
	// The `ServerToken` of the declaration this status report refers to.
	ServerToken string `json:"server-token" required:"true"`
	// If `true`, the declaration is active on the device.
	Active bool `json:"active" required:"true"`
	// This string defines the validity of the declaration. If it's `invalid`, the `reasons` property contains more details.
	Valid ManagementDeclarationsActivationsValid `json:"valid" required:"true"`
	// The details of any client errors.
	Reasons *[]*ManagementDeclarationsActivationsReasons `json:"reasons,omitempty"`
}

// This string defines the validity of the declaration. If it's `invalid`, the `reasons` property contains more details.
type ManagementDeclarationsActivationsValid string

const (
	ManagementDeclarationsActivationsValidUnknown ManagementDeclarationsActivationsValid = "unknown"
	ManagementDeclarationsActivationsValidInvalid ManagementDeclarationsActivationsValid = "invalid"
	ManagementDeclarationsActivationsValidValid   ManagementDeclarationsActivationsValid = "valid"
)

// Information about a status error.
type ManagementDeclarationsActivationsReasons struct {
	// The error code for this error.
	Code string `json:"code" required:"true"`
	// The description for this error.
	Description *string `json:"description,omitempty"`
	// A dictionary that contains further details about this error.
	Details *map[string]any `json:"details,omitempty"`
}

// Status for a declaration processed by the client.
type ManagementDeclarationsConfigurations struct {
	// The `identifier` of the declaration this status report refers to.
	Identifier string `json:"identifier" required:"true"`
	// The `ServerToken` of the declaration this status report refers to.
	ServerToken string `json:"server-token" required:"true"`
	// If `true`, the declaration is active on the device.
	Active bool `json:"active" required:"true"`
	// This string defines the validity of the declaration. If it's `invalid`, the `reasons` property contains more details.
	Valid ManagementDeclarationsActivationsValid `json:"valid" required:"true"`
	// The details of any client errors.
	Reasons *[]*ManagementDeclarationsActivationsReasons `json:"reasons,omitempty"`
}

// Status for a declaration processed by the client.
type ManagementDeclarationsAssets struct {
	// The `identifier` of the declaration this status report refers to.
	Identifier string `json:"identifier" required:"true"`
	// The `ServerToken` of the declaration this status report refers to.
	ServerToken string `json:"server-token" required:"true"`
	// If `true`, the declaration is active on the device.
	Active bool `json:"active" required:"true"`
	// This string defines the validity of the declaration. If it's `invalid`, the `reasons` property contains more details.
	Valid ManagementDeclarationsActivationsValid `json:"valid" required:"true"`
	// The details of any client errors.
	Reasons *[]*ManagementDeclarationsActivationsReasons `json:"reasons,omitempty"`
}

// Status for a declaration processed by the client.
type ManagementDeclarationsManagement struct {
	// The `identifier` of the declaration this status report refers to.
	Identifier string `json:"identifier" required:"true"`
	// The `ServerToken` of the declaration this status report refers to.
	ServerToken string `json:"server-token" required:"true"`
	// If `true`, the declaration is active on the device.
	Active bool `json:"active" required:"true"`
	// This string defines the validity of the declaration. If it's `invalid`, the `reasons` property contains more details.
	Valid ManagementDeclarationsActivationsValid `json:"valid" required:"true"`
	// The details of any client errors.
	Reasons *[]*ManagementDeclarationsActivationsReasons `json:"reasons,omitempty"`
}

// Status items under the `mdm` namespace.
type StatusItemsMDM struct {
	// The list of apps. The response doesn't include apps that are managed by Declarative Device Management.
	// Status item: `mdm.app`.
	App *[]*MDMApp `json:"app,omitempty"`
}

// A status report that contains details about an MDM-installed app.
// Status item: `mdm.app`.
type MDMApp struct {
	// The app's bundle id, which is unique.
	Identifier string `json:"identifier" required:"true"`
	// If `true`, the system removed the app and only this key and the `identifier` key are present in the status item object. The device reports an MDM-installed app as removed if management of the app has been transferred to Declarative Device Management.
	Removed *bool `json:"_removed,omitempty"`
	// The name of the app.
	Name *string `json:"name,omitempty"`
	// The app's external version identifier. You can also retrieve this value from the App Store. For more information, see `Apps and Books for Organizations`.
	// If the current external version identifier of an app on the App Store doesn't match the external version identifier reported by the device, there may be an app update available for the device.
	ExternalVersionID *string `json:"external-version-id,omitempty"`
	// The version of the app.
	Version *string `json:"version,omitempty"`
	// The short version of the app.
	ShortVersion *string `json:"short-version,omitempty"`
	// The status of the app that `ManagedApplicationListCommand` reports.
	State *MDMAppState `json:"state,omitempty"`
}

// The status of the app that `ManagedApplicationListCommand` reports.
type MDMAppState string

const (
	MDMAppStateQueued                  MDMAppState = "queued"
	MDMAppStateNeedsRedemption         MDMAppState = "needs-redemption"
	MDMAppStateRedeeming               MDMAppState = "redeeming"
	MDMAppStatePrompting               MDMAppState = "prompting"
	MDMAppStatePromptingForLogin       MDMAppState = "prompting-for-login"
	MDMAppStateValidatingPurchase      MDMAppState = "validating-purchase"
	MDMAppStatePromptingForUpdate      MDMAppState = "prompting-for-update"
	MDMAppStatePromptingForUpdateLogin MDMAppState = "prompting-for-update-login"
	MDMAppStatePromptingForManagement  MDMAppState = "prompting-for-management"
	MDMAppStateValidatingUpdate        MDMAppState = "validating-update"
	MDMAppStateUpdating                MDMAppState = "updating"
	MDMAppStateInstalling              MDMAppState = "installing"
	MDMAppStateManaged                 MDMAppState = "managed"
	MDMAppStateManagedButUninstalled   MDMAppState = "managed-but-uninstalled"
	MDMAppStateUnknown                 MDMAppState = "unknown"
	MDMAppStateUserInstalledApp        MDMAppState = "user-installed-app"
	MDMAppStateUserRejected            MDMAppState = "user-rejected"
	MDMAppStateUpdateRejected          MDMAppState = "update-rejected"
	MDMAppStateManagementRejected      MDMAppState = "management-rejected"
	MDMAppStateFailed                  MDMAppState = "failed"
)

// Status items under the `package` namespace.
type StatusItemsPackage struct {
	// An array of dictionaries that describe the device's declarative packages.
	// Status item: `package.list`.
	List *[]*PackageList `json:"list,omitempty"`
}

// A dictionary that describes a declarative package.
// Status item: `package.list`.
type PackageList struct {
	// The package's unique identifier. This is the package identifier value of the package file.
	Identifier string `json:"identifier" required:"true"`
	// If `true`, the system removed the package and only this key and the `identifier` key are present in the status item object.
	Removed *bool `json:"_removed,omitempty"`
	// The identifier of the declaration that controls the package.
	DeclarationIdentifier *string `json:"declaration-identifier,omitempty"`
	// The name of the package.
	Name *string `json:"name,omitempty"`
	// The version of the package. This will be the package version value of the package file.
	Version *string `json:"version,omitempty"`
	// The status of the package, which has the following possible values:
	// - `optional`: The package is optional and the user has to trigger its installation.
	// - `queued`: Installation of the package has started.
	// - `prompting-for-consent`: The system is displaying a prompt to the user to proceed with package installation.
	// - `downloading`: The system is downloading the package.
	// - `installing`: The system is installing the package.
	// - `installed`: The package is installed.
	// - `failed`: The package install failed.
	State *PackageListState `json:"state,omitempty"`
	// An array that contains additional details about the package state, including errors.
	Reasons *[]*PackageListReasons `json:"reasons,omitempty"`
}

// The status of the package, which has the following possible values:
// - `optional`: The package is optional and the user has to trigger its installation.
// - `queued`: Installation of the package has started.
// - `prompting-for-consent`: The system is displaying a prompt to the user to proceed with package installation.
// - `downloading`: The system is downloading the package.
// - `installing`: The system is installing the package.
// - `installed`: The package is installed.
// - `failed`: The package install failed.
type PackageListState string

const (
	PackageListStateOptional            PackageListState = "optional"
	PackageListStateQueued              PackageListState = "queued"
	PackageListStatePromptingForConsent PackageListState = "prompting-for-consent"
	PackageListStateDownloading         PackageListState = "downloading"
	PackageListStateInstalling          PackageListState = "installing"
	PackageListStateInstalled           PackageListState = "installed"
	PackageListStateFailed              PackageListState = "failed"
)

// Information about a status error.
type PackageListReasons struct {
	// A code for the state.
	Code string `json:"code" required:"true"`
	// A description of the state.
	Description *string `json:"description,omitempty"`
	// A dictionary that contains additional details about the state.
	Details *map[string]any `json:"details,omitempty"`
}

// Status items under the `passcode` namespace.
type StatusItemsPasscode struct {
	// If `true`, the passcode is in compliance with all passcode policies set on the device. If `false`, the passcode isn't in compliance with one or more passcode policies set on the device. When there are no passcode policies on the device, this value `true`.
	// Status item: `passcode.is-compliant`.
	IsCompliant *bool `json:"is-compliant,omitempty"`
	// If `true`, a passcode is present on the device. If `false`, a passcode isn't present on the device. When a passcode is present, the specific attributes of the passcode, such as length or number of complex characters, aren't reported. Instead, use the `passcode.is-compliant` status item to verify that the passcode complies with all passcode policies set on the device.
	// Status item: `passcode.is-present`.
	IsPresent *bool `json:"is-present,omitempty"`
}

// Status items under the `screensharing` namespace.
type StatusItemsScreenSharing struct {
	// Status items under the `screensharing.connection` namespace.
	Connection *StatusItemsScreenSharingConnection `json:"connection,omitempty"`
}

// Status items under the `screensharing.connection` namespace.
type StatusItemsScreenSharingConnection struct {
	// Status items under the `screensharing.connection.group` namespace.
	Group *StatusItemsScreenSharingConnectionGroup `json:"group,omitempty"`
}

// Status items under the `screensharing.connection.group` namespace.
type StatusItemsScreenSharingConnectionGroup struct {
	// A status item that contains an array of unresolved connection groups.
	// Status item: `screensharing.connection.group.unresolved-connection`.
	UnresolvedConnection *[]*ScreenSharingConnectionGroupUnresolvedConnection `json:"unresolved-connection,omitempty"`
}

// A status item that contains an unresolved connection group.
// Status item: `screensharing.connection.group.unresolved-connection`.
type ScreenSharingConnectionGroupUnresolvedConnection struct {
	// The unique `ConnectionGroupUUID` identifier of the connection group.
	Identifier string `json:"identifier" required:"true"`
	// If `true`, the system removed the unresolved connection group and only this key and the `identifier` key are present in the status item object.
	Removed *bool `json:"_removed,omitempty"`
	// An array of `ConnectionUUID` values specified in the `Members` key in the group's declaration for the unresolved connections.
	UnresolvedConnections *[]string `json:"unresolved_connections,omitempty"`
}

// Status items under the `security` namespace.
type StatusItemsSecurity struct {
	// Status items under the `security.certificate` namespace.
	Certificate *StatusItemsSecurityCertificate `json:"certificate,omitempty"`
}

// Status items under the `security.certificate` namespace.
type StatusItemsSecurityCertificate struct {
	// A list of the device's managed certificates.
	// Status item: `security.certificate.list`.
	List *[]*SecurityCertificateList `json:"list,omitempty"`
}

// A status report of a security certificate.
// Status item: `security.certificate.list`.
type SecurityCertificateList struct {
	// The unique identifier of the certificate which the system uses as the primary key.
	Identifier string `json:"identifier" required:"true"`
	// If `true`, the system removed the app and only this key and the `identifier` key are present in the status item object.
	Removed *bool `json:"_removed,omitempty"`
	// The identifier of the asset declaration that installed the certificate, which is only present if a declaration installed the certificate.
	DeclarationIdentifier *string `json:"declaration-identifier,omitempty"`
	// The summary of the certificate's subject.
	SubjectSummary string `json:"subject-summary" required:"true"`
	// If `true`, the certificate is an identity certificate.
	IsIdentity bool `json:"is-identity" required:"true"`
	// The certificate data in DER-encoded X.509 format.
	Data []byte `json:"data" required:"true"`
}

// Status items under the `services` namespace.
type StatusItemsServices struct {
	// The background task.
	// Status item: `services.background-task`.
	BackgroundTask *[]*ServicesBackgroundTask `json:"background-task,omitempty"`
}

// A status report of a background task.
// Status item: `services.background-task`.
type ServicesBackgroundTask struct {
	// The background task UUID which the system uses as the primary key.
	Identifier string `json:"identifier" required:"true"`
	// If `true`, the background task is removed and the status item object only contains this key and the `identifier` key.
	Removed *bool `json:"_removed,omitempty"`
	// For types other than `agent` or `daemon`, this is the code signature designated requirement of the item, if available.
	CodeSignature *string `json:"code-signature,omitempty"`
	// The numeric user identifier of the owner of the background task.
	UID int64 `json:"uid" required:"true"`
	// For an `agent` or `daemon`, the path to the `launchd` `plist` file. For other types, the path to the app or the document.
	Path string `json:"path" required:"true"`
	// The `SMAppService.Status` enumeration.
	State ServicesBackgroundTaskState `json:"state" required:"true"`
	// The daemon, agent, or SFL login item type.
	Type ServicesBackgroundTaskType `json:"type" required:"true"`
	// Details about a `launchd`-based background task, which is only present when the type is `daemon` or `agent`.
	Launchd *ServicesBackgroundTaskLaunchd `json:"launchd,omitempty"`
}

// The `SMAppService.Status` enumeration.
type ServicesBackgroundTaskState string

const (
	ServicesBackgroundTaskStateNotRegistered    ServicesBackgroundTaskState = "not-registered"
	ServicesBackgroundTaskStateEnabled          ServicesBackgroundTaskState = "enabled"
	ServicesBackgroundTaskStateRequiresApproval ServicesBackgroundTaskState = "requires-approval"
	ServicesBackgroundTaskStateNotFound         ServicesBackgroundTaskState = "not-found"
)

// The daemon, agent, or SFL login item type.
type ServicesBackgroundTaskType string

const (
	ServicesBackgroundTaskTypeDaemon    ServicesBackgroundTaskType = "daemon"
	ServicesBackgroundTaskTypeAgent     ServicesBackgroundTaskType = "agent"
	ServicesBackgroundTaskTypeLoginItem ServicesBackgroundTaskType = "login-item"
	ServicesBackgroundTaskTypeApp       ServicesBackgroundTaskType = "app"
	ServicesBackgroundTaskTypeUserItem  ServicesBackgroundTaskType = "user-item"
)

// Details about a `launchd`-based background task, which is only present when the type is `daemon` or `agent`.
type ServicesBackgroundTaskLaunchd struct {
	// The label of the `launchd`-based background task.
	Label string `json:"label" required:"true"`
	// The program that the `launchd` `plist` file specifies.
	Program string `json:"program" required:"true"`
	// The program arguments that the `launchd` `plist` file specifies.
	ProgramArguments *[]string `json:"program-arguments,omitempty"`
	// The hash value of the `launchd` `plist` file.
	Checksum string `json:"checksum" required:"true"`
	// A dictionary that indicates a `ServicesBackgroundTasks` configuration created this background task. The dictionary contains properties that identify the configuration and the declaration asset that provided the launchd plist for the task.
	DeviceManagement *ServicesBackgroundTaskLaunchdDeviceManagement `json:"device-management,omitempty"`
}

// A dictionary that indicates a `ServicesBackgroundTasks` configuration created this background task. The dictionary contains properties that identify the configuration and the declaration asset that provided the launchd plist for the task.
type ServicesBackgroundTaskLaunchdDeviceManagement struct {
	// The identifier of the `ServicesBackgroundTasks` configuration that created this task.
	ConfigurationIdentifier string `json:"configuration-identifier" required:"true"`
	// The `Identifier` of the declaration asset that provided the launchd plist for this task.
	AssetIdentifier string `json:"asset-identifier" required:"true"`
	// The `ServerToken` of the declaration asset that provided the launchd plist for this task.
	AssetServerToken string `json:"asset-server-token" required:"true"`
}

// Status items under the `softwareupdate` namespace.
type StatusItemsSoftwareUpdate struct {
	// The device's enrolled beta program name, or an empty string if there's no enrolled beta program.
	// Status item: `softwareupdate.beta-enrollment`.
	BetaEnrollment *string `json:"beta-enrollment,omitempty"`
	// The device identifier to use when looking up available software updates via `https://gdmf.apple.com/v2/pmv`.
	// Status item: `softwareupdate.device-id`.
	DeviceID *string `json:"device-id,omitempty"`
	// Details about a software update failure.
	// Status item: `softwareupdate.failure-reason`.
	FailureReason *SoftwareUpdateFailureReason `json:"failure-reason,omitempty"`
	// Details about the reason for a pending software update.
	// Status item: `softwareupdate.install-reason`.
	InstallReason *SoftwareUpdateInstallReason `json:"install-reason,omitempty"`
	// The software update install status, which has the following values:
	// - `none`: There's no software update pending, and any previous software update succeeded.
	// - `waiting': A software update is waiting to start.
	// - `downloading`: The system is downloading data for a software update.
	// - `prepared`: The system prepared the software update and it's ready for installation.
	// - `installing`: The system is installing the software update.
	// - `failed`: The software update failed.
	// Status item: `softwareupdate.install-state`.
	InstallState *SoftwareUpdateInstallState `json:"install-state,omitempty"`
	// A dictionary that contains the build and OS versions of the software update that's pending on the device.
	// Status item: `softwareupdate.pending-version`.
	PendingVersion *SoftwareUpdatePendingVersion `json:"pending-version,omitempty"`
}

// Details about a software update failure.
// Status item: `softwareupdate.failure-reason`.
type SoftwareUpdateFailureReason struct {
	// The number of times the current software update failed. If there are no failures, or no pending software update, this is `0`.
	Count int64 `json:"count" required:"true"`
	// If present, this describes the reason for last software update failure. This key isn't present if there are no failures or no pending software update.
	Reason *string `json:"reason,omitempty"`
	// If present, this is the RFC 3339 timestamp of the last software update failure. This key isn't present if there are no failures or no pending software update.
	Timestamp *string `json:"timestamp,omitempty"`
}

// Details about the reason for a pending software update.
// Status item: `softwareupdate.install-reason`.
type SoftwareUpdateInstallReason struct {
	// A list of reasons for the pending software update. An empty list indicates that no software update is pending.
	Reason []SoftwareUpdateInstallReasonReason `json:"reason" required:"true"`
	// The identifier of the declaration that caused the software update to occur. This key is present only if the `reason` array contains the `declaration` value.
	DeclarationID *string `json:"declaration-id,omitempty"`
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
type SoftwareUpdateInstallReasonReason string

const (
	SoftwareUpdateInstallReasonReasonSystemSettings SoftwareUpdateInstallReasonReason = "system-settings"
	SoftwareUpdateInstallReasonReasonInstallTonight SoftwareUpdateInstallReasonReason = "install-tonight"
	SoftwareUpdateInstallReasonReasonAutoUpdate     SoftwareUpdateInstallReasonReason = "auto-update"
	SoftwareUpdateInstallReasonReasonNotification   SoftwareUpdateInstallReasonReason = "notification"
	SoftwareUpdateInstallReasonReasonSetupAssistant SoftwareUpdateInstallReasonReason = "setup-assistant"
	SoftwareUpdateInstallReasonReasonCommandLine    SoftwareUpdateInstallReasonReason = "command-line"
	SoftwareUpdateInstallReasonReasonMDM            SoftwareUpdateInstallReasonReason = "mdm"
	SoftwareUpdateInstallReasonReasonDeclaration    SoftwareUpdateInstallReasonReason = "declaration"
)

// The software update install status, which has the following values:
// - `none`: There's no software update pending, and any previous software update succeeded.
// - `waiting': A software update is waiting to start.
// - `downloading`: The system is downloading data for a software update.
// - `prepared`: The system prepared the software update and it's ready for installation.
// - `installing`: The system is installing the software update.
// - `failed`: The software update failed.
// Status item: `softwareupdate.install-state`.
type SoftwareUpdateInstallState string

const (
	SoftwareUpdateInstallStateNone        SoftwareUpdateInstallState = "none"
	SoftwareUpdateInstallStateDownloading SoftwareUpdateInstallState = "downloading"
	SoftwareUpdateInstallStatePrepared    SoftwareUpdateInstallState = "prepared"
	SoftwareUpdateInstallStateInstalling  SoftwareUpdateInstallState = "installing"
	SoftwareUpdateInstallStateFailed      SoftwareUpdateInstallState = "failed"
)

// A dictionary that contains the build and OS versions of the software update that's pending on the device.
// Status item: `softwareupdate.pending-version`.
type SoftwareUpdatePendingVersion struct {
	// The OS version of the pending software update, including any Background Security Improvement version. This string is empty if no update is pending.
	OSVersion string `json:"os-version" required:"true"`
	// The build version of the pending software update, including any Background Security Improvement version. This string is empty if no update is pending.
	BuildVersion string `json:"build-version" required:"true"`
	// The local date time value that indicates when the pending software update will be installed. This key is only present when the pending software update is being enforced.
	TargetLocalDateTime *string `json:"target-local-date-time,omitempty"`
}

// Status items under the `test` namespace.
type StatusItemsTest struct {
	// The test status item array value.
	// Status item: `test.array-value`.
	ArrayValue *[]*TestArrayValue `json:"array-value,omitempty"`
	// The test status Boolean value.
	// Status item: `test.boolean-value`.
	BooleanValue *bool `json:"boolean-value,omitempty"`
	// The test status dictionary value.
	// Status item: `test.dictionary-value`.
	DictionaryValue *TestDictionaryValue `json:"dictionary-value,omitempty"`
	// The test status error value.
	// Status item: `test.error-value`.
	ErrorValue *string `json:"error-value,omitempty"`
	// The test status integer value.
	// Status item: `test.integer-value`.
	IntegerValue *int64 `json:"integer-value,omitempty"`
	// The test status real value.
	// Status item: `test.real-value`.
	RealValue *float64 `json:"real-value,omitempty"`
	// The test status string value.
	// Status item: `test.string-value`.
	StringValue *string `json:"string-value,omitempty"`
}

// A status value for the test status item array.
// Status item: `test.array-value`.
type TestArrayValue struct {
	// The value of the first sub-key.
	Key1 string `json:"key1" required:"true"`
	// The value of the second sub-key.
	Key2 *string `json:"key2,omitempty"`
}

// The test status dictionary value.
// Status item: `test.dictionary-value`.
type TestDictionaryValue struct {
	// The value of the first sub-key.
	Key1 string `json:"key1" required:"true"`
	// The value of the second sub-key.
	Key2 *string `json:"key2,omitempty"`
}

// Error information for a status item that cannot be returned.
type StatusReportErrors struct {
	// The status item that this error pertains to.
	StatusItem string `json:"StatusItem" required:"true"`
	// An array of reasons for the error.
	Reasons *[]*StatusReportErrorsReasons `json:"Reasons,omitempty"`
}

// Information about a status error.
type StatusReportErrorsReasons struct {
	// The error code for this error.
	Code string `json:"Code" required:"true"`
	// The description for this error.
	Description *string `json:"Description,omitempty"`
	// A dictionary that contains further details about this error.
	Details *map[string]any `json:"Details,omitempty"`
}
