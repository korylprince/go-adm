// DO NOT EDIT
// generated from https://github.com/apple/device-management.git:0a4527c5ea21825fd23e08273ccdb9e2302458ce/mdm/commands/commands

package commands

import "time"

const DeviceManagementGenerateHash = "0a4527c5ea21825fd23e08273ccdb9e2302458ce"

var CommandMap = map[string]any{
	"AccountConfiguration":            AccountConfigurationCommand{},
	"ActivationLockBypassCode":        ActivationLockBypassCodeCommand{},
	"ActiveNSExtensions":              ActiveNSExtensionsCommand{},
	"ApplyRedemptionCode":             ApplyRedemptionCodeCommand{},
	"AvailableOSUpdates":              AvailableOSUpdatesCommand{},
	"CertificateList":                 CertificateListCommand{},
	"ClearActivationLockBypassCode":   ClearActivationLockBypassCodeCommand{},
	"ClearPasscode":                   ClearPasscodeCommand{},
	"ClearRestrictionsPassword":       ClearRestrictionsPasswordCommand{},
	"ContentCachingInformation":       ContentCachingInformationCommand{},
	"DeclarativeManagement":           DeclarativeManagementCommand{},
	"DeleteUser":                      DeleteUserCommand{},
	"DeviceConfigured":                DeviceConfiguredCommand{},
	"DeviceInformation":               DeviceInformationCommand{},
	"DeviceLocation":                  DeviceLocationCommand{},
	"DeviceLock":                      DeviceLockCommand{},
	"DisableLostMode":                 DisableLostModeCommand{},
	"DisableRemoteDesktop":            DisableRemoteDesktopCommand{},
	"EnableLostMode":                  EnableLostModeCommand{},
	"EnableRemoteDesktop":             EnableRemoteDesktopCommand{},
	"EraseDevice":                     EraseDeviceCommand{},
	"InstallApplication":              InstallApplicationCommand{},
	"InstallEnterpriseApplication":    InstallEnterpriseApplicationCommand{},
	"InstallMedia":                    InstallMediaCommand{},
	"InstallProfile":                  InstallProfileCommand{},
	"InstallProvisioningProfile":      InstallProvisioningProfileCommand{},
	"InstalledApplicationList":        InstalledApplicationListCommand{},
	"InviteToProgram":                 InviteToProgramCommand{},
	"LOMDeviceRequest":                LOMDeviceRequestCommand{},
	"LOMSetupRequest":                 LOMSetupRequestCommand{},
	"LogOutUser":                      LogOutUserCommand{},
	"ManagedApplicationAttributes":    ManagedApplicationAttributesCommand{},
	"ManagedApplicationConfiguration": ManagedApplicationConfigurationCommand{},
	"ManagedApplicationFeedback":      ManagedApplicationFeedbackCommand{},
	"ManagedApplicationList":          ManagedApplicationListCommand{},
	"ManagedMediaList":                ManagedMediaListCommand{},
	"NSExtensionMappings":             NSExtensionMappingsCommand{},
	"OSUpdateStatus":                  OSUpdateStatusCommand{},
	"PlayLostModeSound":               PlayLostModeSoundCommand{},
	"ProfileList":                     ProfileListCommand{},
	"ProvisioningProfileList":         ProvisioningProfileListCommand{},
	"RefreshCellularPlans":            RefreshCellularPlansCommand{},
	"RemoveApplication":               RemoveApplicationCommand{},
	"RemoveMedia":                     RemoveMediaCommand{},
	"RemoveProfile":                   RemoveProfileCommand{},
	"RemoveProvisioningProfile":       RemoveProvisioningProfileCommand{},
	"RequestMirroring":                RequestMirroringCommand{},
	"RequestUnlockToken":              RequestUnlockTokenCommand{},
	"RestartDevice":                   RestartDeviceCommand{},
	"Restrictions":                    RestrictionsCommand{},
	"RotateFileVaultKey":              RotateFileVaultKeyCommand{},
	"ScheduleOSUpdate":                ScheduleOSUpdateCommand{},
	"ScheduleOSUpdateScan":            ScheduleOSUpdateScanCommand{},
	"SecurityInfo":                    SecurityInfoCommand{},
	"SetAutoAdminPassword":            SetAutoAdminPasswordCommand{},
	"SetFirmwarePassword":             SetFirmwarePasswordCommand{},
	"SetRecoveryLock":                 SetRecoveryLockCommand{},
	"Settings":                        SettingsCommand{},
	"ShutDownDevice":                  ShutDownDeviceCommand{},
	"StopMirroring":                   StopMirroringCommand{},
	"UnlockUserAccount":               UnlockUserAccountCommand{},
	"UserConfigured":                  UserConfiguredCommand{},
	"UserList":                        UserListCommand{},
	"ValidateApplications":            ValidateApplicationsCommand{},
	"VerifyFirmwarePassword":          VerifyFirmwarePasswordCommand{},
	"VerifyRecoveryLock":              VerifyRecoveryLockCommand{},
}

// Create and configure a local administrator account on a device.
// When a macOS (v10.11 and later) device is configured via DEP to enroll in an MDM server and the DEP profile has the await_device_configuration flag set to true, the AccountConfiguration command can be sent to the device to have it create the local administrator account (thereby skipping the page to create this account in Setup Assistant). This command can only be sent to a macOS device that is in the AwaitingConfiguration state.
type AccountConfigurationCommand struct {
	// If `true`, Setup Assistant skips the user interface for setting up primary accounts and disables autologin. If `true`, you must specify a value for `AutoSetupAdminAccounts`.
	SkipPrimarySetupAccountCreation *bool `json:"SkipPrimarySetupAccountCreation,omitempty" plist:"SkipPrimarySetupAccountCreation,omitempty"`
	// If `true`, Setup Assistant creates the primary accounts as regular users, and you must specify a value for `AutoSetupAdminAccounts`.
	SetPrimarySetupAccountAsRegularUser *bool `json:"SetPrimarySetupAccountAsRegularUser,omitempty" plist:"SetPrimarySetupAccountAsRegularUser,omitempty"`
	// The full name for the primary account. If present, Setup Assistant uses this value to prefill the Full Name field. However, Setup Assistant ignores this value if `DontAutoPopulatePrimaryAccountInfo` is `true`. This value is available in macOS 10.15 and later.
	PrimaryAccountFullName *string `json:"PrimaryAccountFullName,omitempty" plist:"PrimaryAccountFullName,omitempty"`
	// The account name for the primary account. If present, Setup Assistant uses this value to prefill the User Name field. However, Setup Assistant ignores this value if `DontAutoPopulatePrimaryAccountInfo` is `true`. This value is available in macOS 10.15 and later.
	PrimaryAccountUserName *string `json:"PrimaryAccountUserName,omitempty" plist:"PrimaryAccountUserName,omitempty"`
	// If `true`, Setup Assistant ignores the primary account information and requires the user to enter that information. If `false`, Setup Assistant prefills the Full Name field with `PrimaryAccountFullName` and the User Name field with `PrimaryAccountUserName`. This value is available in macOS 10.15 and later.
	DontAutoPopulatePrimaryAccountInfo *bool `json:"DontAutoPopulatePrimaryAccountInfo,omitempty" plist:"DontAutoPopulatePrimaryAccountInfo,omitempty"`
	// If `true`, and you provide values for `PrimaryAccountFullName` or `PrimaryAccountUserName`, Setup Assistant disables editing for the corresponding fields. `DontAutoPopulatePrimaryAccountInfo` must also be 0 (or missing).
	// If the user's password is also available from authentication through ConfigurationURL, Setup Assistant automatically creates the primary account with that information and skips showing the user interface to view or edit these fields.
	// This value is available in macOS 10.15 and later.
	LockPrimaryAccountInfo *bool `json:"LockPrimaryAccountInfo,omitempty" plist:"LockPrimaryAccountInfo,omitempty"`
	// A dictionary that describes the administrator account to create with Setup Assistant, which uses the first element and ignores additional elements.
	AutoSetupAdminAccounts *[]*AutoSetupAdminAccountItem `json:"AutoSetupAdminAccounts,omitempty" plist:"AutoSetupAdminAccounts,omitempty"`
	// If present, this is the short name of the local account to manage, which can also be the account that results from setting `AutoSetupAdminAccounts` to `true`. Otherwise, only the local account that Setup Assistant creates is a managed account. This value is available in macOS 11 and later.
	ManagedLocalUserShortName *string `json:"ManagedLocalUserShortName,omitempty" plist:"ManagedLocalUserShortName,omitempty"`
}

func (p *AccountConfigurationCommand) RequestType() string {
	return "AccountConfiguration"
}

type AccountConfigurationCommandResponse struct{}

// A dictionary that describes the administrator account to create with Setup Assistant, which uses the first element and ignores additional elements.
type AutoSetupAdminAccountItem struct {
	// The short name of the user.
	ShortName string `json:"shortName" plist:"shortName" required:"true"`
	// The full name of the user, which defaults to `shortName` if not specified.
	FullName *string `json:"fullName,omitempty" plist:"fullName,omitempty"`
	// Data that contains the pre-created salted PBKDF2 SHA512 password hash for the account.
	PasswordHash *[]byte `json:"passwordHash,omitempty" plist:"passwordHash,omitempty"`
	// If `true`, this sets the account attribute to make the account hidden in the Login Window and Users & Groups.
	Hidden *bool `json:"hidden,omitempty" plist:"hidden,omitempty"`
}

// Get a list of active extensions for a user on a device.
// Returns information about the active NSExtensions for a particular user.
// NSExtensions are installed and enabled at the user level. There is no concept of "device" NSExtensions.
// Requires "Query Installed Apps" right; supported on user channel only.
type ActiveNSExtensionsCommand struct {
	// An array of extension points. If you choose to provide this value, the response only includes the app extensions for the extension points you specify.
	FilterExtensionPoints *[]string `json:"FilterExtensionPoints,omitempty" plist:"FilterExtensionPoints,omitempty"`
}

func (p *ActiveNSExtensionsCommand) RequestType() string {
	return "ActiveNSExtensions"
}

type ActiveNSExtensionsCommandResponse struct {
	// An array of dictionaries that contains information about active extensions on the device.
	Extensions []*ActiveNSExtensionsCommandExtensionsExtensionsItem `json:"Extensions" plist:"Extensions" required:"true"`
}

// A dictionary that contains information about an extension.
type ActiveNSExtensionsCommandExtensionsExtensionsItem struct {
	// The identifier of the extension.
	Identifier string `json:"Identifier" plist:"Identifier" required:"true"`
	// The `NSExtensionPointIdentifier` for the extension.
	ExtensionPoint string `json:"ExtensionPoint" plist:"ExtensionPoint" required:"true"`
	// The extension's display name.
	DisplayName string `json:"DisplayName" plist:"DisplayName" required:"true"`
	// The display name of the container.
	ContainerDisplayName *string `json:"ContainerDisplayName,omitempty" plist:"ContainerDisplayName,omitempty"`
	// The identifier of the container.
	ContainerIdentifier *string `json:"ContainerIdentifier,omitempty" plist:"ContainerIdentifier,omitempty"`
	// The path to the extension.
	Path string `json:"Path" plist:"Path" required:"true"`
	// The version of the extension.
	Version string `json:"Version" plist:"Version" required:"true"`
	// The user-selected state of the extension, which a user sets in the Extensions preference pane in System Preferences.
	UserElection UserElection `json:"UserElection" plist:"UserElection" required:"true"`
}

// The user-selected state of the extension, which a user sets in the Extensions preference pane in System Preferences.
type UserElection string

const (
	UserElectionDefault UserElection = "Default"
	UserElectionUse     UserElection = "Use"
	UserElectionIgnore  UserElection = "Ignore"
)

// Get a list of the installed extensions for a user on a device.
// This command returns information about installed extensions for a user.
// The purpose of this command is to allow the server to build a mapping of
// extension identifiers to extension points to provide a UI for generating
// "com.apple.NSExtension" payloads.
// Requires "Query Installed Apps" right; supported on user channel only
type NSExtensionMappingsCommand struct{}

func (p *NSExtensionMappingsCommand) RequestType() string {
	return "NSExtensionMappings"
}

type NSExtensionMappingsCommandResponse struct {
	// An array of dictionaries that contains information about extensions on the device.
	Extensions []*NSExtensionMappingsCommandExtensionsExtensionsItem `json:"Extensions" plist:"Extensions" required:"true"`
}

// A dictionary that contains information about an extension.
type NSExtensionMappingsCommandExtensionsExtensionsItem struct {
	// The identifier of the extension.
	Identifier string `json:"Identifier" plist:"Identifier" required:"true"`
	// The `NSExtensionPointIdentifier` for the extension.
	ExtensionPoint string `json:"ExtensionPoint" plist:"ExtensionPoint" required:"true"`
	// The display name of the extension.
	DisplayName string `json:"DisplayName" plist:"DisplayName" required:"true"`
}

// Install an enterprise app on a device.
// This command allows the server to install an application on a device. It provides a more secure version of 'InstallApplication' that specifies a 'ManifestURL'.
type InstallEnterpriseApplicationCommand struct {
	// A dictionary that specifies where to download the app. This value is backward-compatible with the manifest for the `InstallApplicationCommand`; however, it also allows you to specify `sha256s` and `sha256-size` for SHA-256 hashes.
	Manifest *map[string]any `json:"Manifest,omitempty" plist:"Manifest,omitempty"`
	// The URL of the app manifest, which needs to begin with `https:`. The manifest is returned as a property list.
	ManifestURL *string `json:"ManifestURL,omitempty" plist:"ManifestURL,omitempty"`
	// An array of DER-encoded certificates to pin the connection when fetching the `ManifestURL`.
	ManifestURLPinningCerts *[][]byte `json:"ManifestURLPinningCerts,omitempty" plist:"ManifestURLPinningCerts,omitempty"`
	// If `true`, certificate revocation checks require a positive response when using certificate pinning with `ManifestURLPinningCerts`.
	PinningRevocationCheckRequired *bool `json:"PinningRevocationCheckRequired,omitempty" plist:"PinningRevocationCheckRequired,omitempty"`
	// If `true`, install the app as a managed app. Otherwise, the system installs the app as unmanaged. If you reinstall a manged app and omit this value or set it to `false`, the app becomes unmanaged.
	// For manifest-based installs, if `true`, the system only considers apps installed in `/Applications` as managed. In macOS 11 through 13, the system requires that the `pkg` only contains a single signed app.
	// Available in macOS 11 and later.
	InstallAsManaged *bool `json:"InstallAsManaged,omitempty" plist:"InstallAsManaged,omitempty"`
	// The management flags. The possible values are:
	// - `1`: If `InstallAsManaged` is `true`, remove the app upon removal of the MDM profile.
	// Available in macOS 11 and later.
	ManagementFlags *InstallEnterpriseApplicationCommandManagementFlags `json:"ManagementFlags,omitempty" plist:"ManagementFlags,omitempty"`
	// A dictionary that contains the initial configuration of the app, if you choose to provide it. Available in macOS 11 and later.
	Configuration *map[string]any `json:"Configuration,omitempty" plist:"Configuration,omitempty"`
	// The change management state. This value doesn't work with the user enrollments. The only possible value is:
	// - `Managed`: Take management of the app if the user installed it already and `InstallAsManaged` is `true`.
	// Available in macOS 11 and later.
	ChangeManagementState *InstallEnterpriseApplicationCommandChangeManagementState `json:"ChangeManagementState,omitempty" plist:"ChangeManagementState,omitempty"`
	// If `true`, the app is an iOS app that can run on a Mac computer with Apple silicon in macOS 11 and later.
	IOSApp *bool `json:"iOSApp,omitempty" plist:"iOSApp,omitempty"`
}

func (p *InstallEnterpriseApplicationCommand) RequestType() string {
	return "InstallEnterpriseApplication"
}

type InstallEnterpriseApplicationCommandResponse struct{}

// The management flags. The possible values are:
// - `1`: If `InstallAsManaged` is `true`, remove the app upon removal of the MDM profile.
// Available in macOS 11 and later.
type InstallEnterpriseApplicationCommandManagementFlags int64

const (
	InstallEnterpriseApplicationCommandManagementFlags1 InstallEnterpriseApplicationCommandManagementFlags = 1
)

// The change management state. This value doesn't work with the user enrollments. The only possible value is:
// - `Managed`: Take management of the app if the user installed it already and `InstallAsManaged` is `true`.
// Available in macOS 11 and later.
type InstallEnterpriseApplicationCommandChangeManagementState string

const (
	InstallEnterpriseApplicationCommandChangeManagementStateManaged InstallEnterpriseApplicationCommandChangeManagementState = "Managed"
)

// Install a third-party app on a device.
// This command allows the server to install an application on a device. If the app is already being managed, this command will update the app. This command will fail for apps that are managed by Declarative Device Management. macOS change - 10.9 user channel for VPP, 10.10 device channel, 10.11 both.
type InstallApplicationCommand struct {
	// The app's iTunes Store identifier.
	ITunesStoreID *int64 `json:"iTunesStoreID,omitempty" plist:"iTunesStoreID,omitempty"`
	// The app's bundle identifier.
	// > Important:
	// > For a watchOS app, the identifier needs to be the watch's bundle identifier, which differs from the main bundle identifier for the iPhone to which the watch is paired. Obtain the watch's bundle identifier for an app with a watch bundle, in the `watchBundleId` key that's part of the Content Metadata query. For more information on this query, see `Getting App and Book Information`.
	Identifier *string `json:"Identifier,omitempty" plist:"Identifier,omitempty"`
	// A dictionary that contains the app installation options.
	Options *Options `json:"Options,omitempty" plist:"Options,omitempty"`
	// The URL of the app manifest, which needs to begin with `https:`. The manifest is returned as a property list.
	ManifestURL *string `json:"ManifestURL,omitempty" plist:"ManifestURL,omitempty"`
	// A bitwise OR of the management flags. The possible values are:
	// - `1`: If `InstallAsManaged` is `true`, remove the app upon removal of the MDM profile.
	// - `4`: Prevent backup of app data.
	// - `5`: Both `1` and `4`.
	// Available in iOS 5 and later, macOS 11 and later, and tvOS 10.2 and later.
	ManagementFlags *InstallApplicationCommandManagementFlags `json:"ManagementFlags,omitempty" plist:"ManagementFlags,omitempty"`
	// A dictionary that contains the initial configuration of the app, if you choose to provide it. Available in iOS 7 and later, macOS 11 and later, and tvOS 10.2 and later.
	Configuration *map[string]any `json:"Configuration,omitempty" plist:"Configuration,omitempty"`
	// A dictionary that contains the initial attributes of the app, if you choose to provide it. Available in iOS 7 and later, and tvOS 10.2 and later.
	Attributes *InstallApplicationCommandAttributes `json:"Attributes,omitempty" plist:"Attributes,omitempty"`
	// The change management state. This value doesn't work with the user enrollment feature introduced in iOS 13. Available in iOS 9 and later, macOS 11 and later, and tvOS 10.2 and later. The only possible value is:
	// - `Managed`: Take management of the app if the user installed it already and `InstallAsManaged` is `true`.
	ChangeManagementState *InstallApplicationCommandChangeManagementState `json:"ChangeManagementState,omitempty" plist:"ChangeManagementState,omitempty"`
	// If `true`, install the app as a managed app. Otherwise, the system installs the app as unmanaged. If you reinstall a manged app and omit this value or set it to `false`, the app becomes unmanaged.
	// For manifest-based installs, if `true`, the system only considers apps installed in `/Applications` as managed. In macOS 11 through 13, the system requires that the `pkg` only contains a single signed app.
	// Available in macOS 11 and later.
	InstallAsManaged *bool `json:"InstallAsManaged,omitempty" plist:"InstallAsManaged,omitempty"`
	// If `true`, the app is an iOS app that can run on a Mac computer with Apple silicon in macOS 11 and later.
	IOSApp *bool `json:"iOSApp,omitempty" plist:"iOSApp,omitempty"`
}

func (p *InstallApplicationCommand) RequestType() string {
	return "InstallApplication"
}

type InstallApplicationCommandResponse struct {
	// The app's bundle identifier, if the user accepted the request.
	// > Note:
	// > For a watchOS app, the identifier is the watch's bundle identifier, which differs from the main bundle identifier for the iPhone that the watch is paired to.
	Identifier *string `json:"Identifier,omitempty" plist:"Identifier,omitempty"`
	// The app's installation state, if the user accepted the request. If this value is `NeedsRedemption`, the server needs to send a redemption code to complete the app installation.
	State *InstallApplicationCommandState `json:"State,omitempty" plist:"State,omitempty"`
	// The reason, if installation fails. macOS always returns "Other".
	RejectionReason *InstallApplicationCommandRejectionReason `json:"RejectionReason,omitempty" plist:"RejectionReason,omitempty"`
}

// A dictionary that contains the app installation options.
type Options struct {
	// The app's purchase type, which must be one of the following values:
	// - `0`: Free apps and Legacy Volume Purchase Program (VPP) with a redemption code. This option is only available in iOS.
	// - `1`: Volume Purchase Program (VPP) app assignment.
	// Set this value to `1` to install first-party apps without user login to the iTunes Store, such as Mail or Safari, or to install an iOS app with user enrollment.
	PurchaseMethod *PurchaseMethod `default:"0" json:"PurchaseMethod,omitempty" plist:"PurchaseMethod,omitempty"`
}

// The app's purchase type, which must be one of the following values:
// - `0`: Free apps and Legacy Volume Purchase Program (VPP) with a redemption code. This option is only available in iOS.
// - `1`: Volume Purchase Program (VPP) app assignment.
// Set this value to `1` to install first-party apps without user login to the iTunes Store, such as Mail or Safari, or to install an iOS app with user enrollment.
type PurchaseMethod int64

const (
	PurchaseMethod0 PurchaseMethod = 0
	PurchaseMethod1 PurchaseMethod = 1
)

// A bitwise OR of the management flags. The possible values are:
// - `1`: If `InstallAsManaged` is `true`, remove the app upon removal of the MDM profile.
// - `4`: Prevent backup of app data.
// - `5`: Both `1` and `4`.
// Available in iOS 5 and later, macOS 11 and later, and tvOS 10.2 and later.
type InstallApplicationCommandManagementFlags int64

const (
	InstallApplicationCommandManagementFlags1 InstallApplicationCommandManagementFlags = 1
	InstallApplicationCommandManagementFlags4 InstallApplicationCommandManagementFlags = 4
	InstallApplicationCommandManagementFlags5 InstallApplicationCommandManagementFlags = 5
)

// A dictionary that contains the initial attributes of the app, if you choose to provide it. Available in iOS 7 and later, and tvOS 10.2 and later.
type InstallApplicationCommandAttributes struct {
	// A per-app VPN unique identifier for this app. Available in iOS 7 and later.
	VPNUUID *string `json:"VPNUUID,omitempty" plist:"VPNUUID,omitempty"`
	// The content filter UUID for this app. Available in iOS 16 and later.
	ContentFilterUUID *string `json:"ContentFilterUUID,omitempty" plist:"ContentFilterUUID,omitempty"`
	// The DNS proxy UUID for this app. Available in iOS 16 and later.
	DNSProxyUUID *string `json:"DNSProxyUUID,omitempty" plist:"DNSProxyUUID,omitempty"`
	// The relay UUID for this app. Available in iOS 17 and later.
	RelayUUID *string `json:"RelayUUID,omitempty" plist:"RelayUUID,omitempty"`
	// An array that contains the associated domains to add to this app. Available in iOS 13 and later.
	AssociatedDomains *[]string `json:"AssociatedDomains,omitempty" plist:"AssociatedDomains,omitempty"`
	// If `true`, perform claimed site association verification directly at the domain instead of on Apple's servers. Only set this to `true` for domains that can't access the internet. Available in iOS 14 and later.
	AssociatedDomainsEnableDirectDownloads *bool `json:"AssociatedDomainsEnableDirectDownloads,omitempty" plist:"AssociatedDomainsEnableDirectDownloads,omitempty"`
	// If `false`, this app isn't removable while it's a managed app. Available in iOS 14 and later, and tvOS 14 and later.
	Removable *bool `json:"Removable,omitempty" plist:"Removable,omitempty"`
	// If `true`, Tap to Pay on iPhone requires users to use Face ID or a passcode to unlock their device after every transaction that requires a customer's card PIN. If `false`, the user can configure this setting on their device.
	// Available in iOS 16.4 and later.
	TapToPayScreenLock *bool `json:"TapToPayScreenLock,omitempty" plist:"TapToPayScreenLock,omitempty"`
	// The data network name (DNN) or app category. For DNN, the value is `DNN:name`, where `name` is the carrier-provided DNN name. For app category, the value is `AppCategory:category`, where `category` is a carrier-provided string like "Enterprise1".
	// Available in iOS 17 and later.
	CellularSliceUUID *string `json:"CellularSliceUUID,omitempty" plist:"CellularSliceUUID,omitempty"`
	// If `false`, the system prevents the user from hiding the app. It doesn't affect the user's ability to leave it in the App Library, while removing it from the Home Screen.
	Hideable *bool `json:"Hideable,omitempty" plist:"Hideable,omitempty"`
	// If `false`, the system prevents the user from locking the app. This also prevents the user from hiding the app.
	Lockable *bool `json:"Lockable,omitempty" plist:"Lockable,omitempty"`
}

// The change management state. This value doesn't work with the user enrollment feature introduced in iOS 13. Available in iOS 9 and later, macOS 11 and later, and tvOS 10.2 and later. The only possible value is:
// - `Managed`: Take management of the app if the user installed it already and `InstallAsManaged` is `true`.
type InstallApplicationCommandChangeManagementState string

const (
	InstallApplicationCommandChangeManagementStateManaged InstallApplicationCommandChangeManagementState = "Managed"
)

// The app's installation state, if the user accepted the request. If this value is `NeedsRedemption`, the server needs to send a redemption code to complete the app installation.
type InstallApplicationCommandState string

const (
	InstallApplicationCommandStateQueued                  InstallApplicationCommandState = "Queued"
	InstallApplicationCommandStateNeedsRedemption         InstallApplicationCommandState = "NeedsRedemption"
	InstallApplicationCommandStateRedeeming               InstallApplicationCommandState = "Redeeming"
	InstallApplicationCommandStatePrompting               InstallApplicationCommandState = "Prompting"
	InstallApplicationCommandStatePromptingForLogin       InstallApplicationCommandState = "PromptingForLogin"
	InstallApplicationCommandStateValidatingPurchase      InstallApplicationCommandState = "ValidatingPurchase"
	InstallApplicationCommandStateInstalling              InstallApplicationCommandState = "Installing"
	InstallApplicationCommandStateManaged                 InstallApplicationCommandState = "Managed"
	InstallApplicationCommandStateManagedButUninstalled   InstallApplicationCommandState = "ManagedButUninstalled"
	InstallApplicationCommandStateUserInstalledApp        InstallApplicationCommandState = "UserInstalledApp"
	InstallApplicationCommandStateUserRejected            InstallApplicationCommandState = "UserRejected"
	InstallApplicationCommandStatePromptingForUpdate      InstallApplicationCommandState = "PromptingForUpdate"
	InstallApplicationCommandStatePromptingForUpdateLogin InstallApplicationCommandState = "PromptingForUpdateLogin"
	InstallApplicationCommandStateValidatingUpdate        InstallApplicationCommandState = "ValidatingUpdate"
	InstallApplicationCommandStateUpdating                InstallApplicationCommandState = "Updating"
	InstallApplicationCommandStateUpdateRejected          InstallApplicationCommandState = "UpdateRejected"
	InstallApplicationCommandStatePromptingForManagement  InstallApplicationCommandState = "PromptingForManagement"
	InstallApplicationCommandStateManagementRejected      InstallApplicationCommandState = "ManagementRejected"
	InstallApplicationCommandStateFailed                  InstallApplicationCommandState = "Failed"
	InstallApplicationCommandStateUnknown                 InstallApplicationCommandState = "Unknown"
)

// The reason, if installation fails. macOS always returns "Other".
type InstallApplicationCommandRejectionReason string

const (
	InstallApplicationCommandRejectionReasonAppAlreadyInstalled                   InstallApplicationCommandRejectionReason = "AppAlreadyInstalled"
	InstallApplicationCommandRejectionReasonAppAlreadyQueued                      InstallApplicationCommandRejectionReason = "AppAlreadyQueued"
	InstallApplicationCommandRejectionReasonAppStoreDisabled                      InstallApplicationCommandRejectionReason = "AppStoreDisabled"
	InstallApplicationCommandRejectionReasonCouldNotVerifyAppID                   InstallApplicationCommandRejectionReason = "CouldNotVerifyAppID"
	InstallApplicationCommandRejectionReasonManagementChangeNotSupported          InstallApplicationCommandRejectionReason = "ManagementChangeNotSupported"
	InstallApplicationCommandRejectionReasonNotAnApp                              InstallApplicationCommandRejectionReason = "NotAnApp"
	InstallApplicationCommandRejectionReasonNotSupported                          InstallApplicationCommandRejectionReason = "NotSupported"
	InstallApplicationCommandRejectionReasonOther                                 InstallApplicationCommandRejectionReason = "Other"
	InstallApplicationCommandRejectionReasonPurchaseMethodNotSupported            InstallApplicationCommandRejectionReason = "PurchaseMethodNotSupported"
	InstallApplicationCommandRejectionReasonPurchaseMethodNotSupportedInMultiUser InstallApplicationCommandRejectionReason = "PurchaseMethodNotSupportedInMultiUser"
)

// Get a list of the installed apps on a device.
// This command allows the server to query for installed 3rd party applications. Starting in iOS 26, this query will also return system apps.
type InstalledApplicationListCommand struct {
	// An array of app identifiers. Provide this value to limit the response to only include these apps. This value is available in iOS 7 and later, macOS 10.15 and later, tvOS 10.2 and later, visionOS 1.1 and later, and watchOS 10 and later.
	// > Important:
	// > For a watchOS app, the identifier needs to be the watch's bundle identifier, which differs from the main bundle identifier for the iPhone to which the watch is paired. Obtain the watch's bundle identifier for an app with a watch bundle, in the `watchBundleId` key that's part of the Content Metadata query. For more information on this query, see `Getting App and Book Information`.
	Identifiers *[]string `json:"Identifiers,omitempty" plist:"Identifiers,omitempty"`
	// If `true`, only get a list of managed apps, excluding ones that are managed by Declarative Device Management. This value is available in iOS 7 and later, macOS 10.15 and later, and tvOS 10.2 and later.
	ManagedAppsOnly *bool `json:"ManagedAppsOnly,omitempty" plist:"ManagedAppsOnly,omitempty"`
	// An array of strings that represent keys in `InstalledApplicationListItem`. If present, the response only contains the keys listed here, except `Identifier` is always included. If not present, the response contains all keys.
	// > Tip:
	// > Only request the keys that you need, because some key values can take significant time and power to calculate on the device.
	Items *[]Items `json:"Items,omitempty" plist:"Items,omitempty"`
}

func (p *InstalledApplicationListCommand) RequestType() string {
	return "InstalledApplicationList"
}

type InstalledApplicationListCommandResponse struct {
	// An array of dictionaries that describes each installed app.
	InstalledApplicationList []*InstalledApplicationListItem `json:"InstalledApplicationList" plist:"InstalledApplicationList" required:"true"`
}

// An array of strings that represent keys in `InstalledApplicationListItem`. If present, the response only contains the keys listed here, except `Identifier` is always included. If not present, the response contains all keys.
// > Tip:
// > Only request the keys that you need, because some key values can take significant time and power to calculate on the device.
type Items string

const (
	ItemsAdHocCodeSigned           Items = "AdHocCodeSigned"
	ItemsAppStoreVendable          Items = "AppStoreVendable"
	ItemsBetaApp                   Items = "BetaApp"
	ItemsBundleSize                Items = "BundleSize"
	ItemsDeviceBasedVPP            Items = "DeviceBasedVPP"
	ItemsDistributorIdentifier     Items = "DistributorIdentifier"
	ItemsDynamicSize               Items = "DynamicSize"
	ItemsExternalVersionIdentifier Items = "ExternalVersionIdentifier"
	ItemsHasUpdateAvailable        Items = "HasUpdateAvailable"
	ItemsIdentifier                Items = "Identifier"
	ItemsInstalling                Items = "Installing"
	ItemsIsAppClip                 Items = "IsAppClip"
	ItemsIsValidated               Items = "IsValidated"
	ItemsName                      Items = "Name"
	ItemsShortVersion              Items = "ShortVersion"
	ItemsVersion                   Items = "Version"
)

// A dictionary that describes an app list item.
type InstalledApplicationListItem struct {
	// The app's identifier. This key is always be present on iOS and tvOS, but may be missing on macOS.
	// > Note:
	// > For a watchOS app, the identifier is the watch's bundle identifier, which differs from the main bundle identifier for the iPhone to which the watch is paired.
	Identifier *string `json:"Identifier,omitempty" plist:"Identifier,omitempty"`
	// The app's external version identifier, which you can use in the iTunes Search API to determine if an updated version of the app is available. Compare this value to the `externalId` value in the `contentMetadataLookupUrl` response from the `VPPServiceConfigSrv` endpoint. If these values don't match, an updated version of the app may be available.
	// > Note:
	// > A newer version of an app might not be available for installation on the device for a variety of reasons. A common reason is that the device's operating system version or hardware is incompatible with the available version of the app.
	ExternalVersionIdentifier *int64 `json:"ExternalVersionIdentifier,omitempty" plist:"ExternalVersionIdentifier,omitempty"`
	// The marketplace hosted application's distributor ID. This value is available in iOS 17.4 and later.
	DistributorIdentifier *string `json:"DistributorIdentifier,omitempty" plist:"DistributorIdentifier,omitempty"`
	// The app's version.
	Version *string `json:"Version,omitempty" plist:"Version,omitempty"`
	// The app's short version.
	ShortVersion *string `json:"ShortVersion,omitempty" plist:"ShortVersion,omitempty"`
	// The app's name.
	Name *string `json:"Name,omitempty" plist:"Name,omitempty"`
	// The app's static bundle size, in bytes. This value is available in iOS 5 and later, and macOS 10.7 and later, and tvOS 10.2 and later.
	BundleSize *int64 `json:"BundleSize,omitempty" plist:"BundleSize,omitempty"`
	// The size of the app's file system in bytes, including the Documents, Library, and other directories. This value is available in iOS 5 and later, and tvOS 10.2 and later.
	DynamicSize *int64 `json:"DynamicSize,omitempty" plist:"DynamicSize,omitempty"`
	// If `true`, the app is valid and can run on the device. If the app is enterprise-distributed and unvalidated, it won't be able to run until validation has occurred. This value is available in iOS 9.2 and later, and tvOS 10.2 and later.
	IsValidated *bool `json:"IsValidated,omitempty" plist:"IsValidated,omitempty"`
	// If `true`, the app is downloading. If `false`, it's already installed.
	Installing *bool `json:"Installing,omitempty" plist:"Installing,omitempty"`
	// If `true`, the app came from the App Store and can participate in store features. For device-based Volume Purchase Program (VPP) apps, this value is `false`. This value is available in iOS 11.3 and later, and tvOS 11.3 and later.
	AppStoreVendable *bool `json:"AppStoreVendable,omitempty" plist:"AppStoreVendable,omitempty"`
	// If `true`, installing the app didn't require an Apple Account. This value is available in iOS 11.3 and later, and tvOS 11.3 and later.
	DeviceBasedVPP *bool `json:"DeviceBasedVPP,omitempty" plist:"DeviceBasedVPP,omitempty"`
	// If `true`, the app is part of the Apple Beta Software Program. This value is available in iOS 11.3 and later, and tvOS 11.3 and later.
	BetaApp *bool `json:"BetaApp,omitempty" plist:"BetaApp,omitempty"`
	// If `true`, the app is ad-hoc code signed. This query is available in iOS 11.3 and later, and tvOS 11.3 and later.
	AdHocCodeSigned *bool `json:"AdHocCodeSigned,omitempty" plist:"AdHocCodeSigned,omitempty"`
	// If `true`, the app has an update available. This key is present only for App Store apps. In macOS, this key is present only for Volume Purchase Program (VPP) apps. This status updates daily and isn't always up-to-date when installing an app.
	HasUpdateAvailable *bool `json:"HasUpdateAvailable,omitempty" plist:"HasUpdateAvailable,omitempty"`
	// If `true`, the download failed.
	DownloadFailed *bool `json:"DownloadFailed,omitempty" plist:"DownloadFailed,omitempty"`
	// If `true`, the app is in the initial state, which is waiting to download.
	DownloadWaiting *bool `json:"DownloadWaiting,omitempty" plist:"DownloadWaiting,omitempty"`
	// If `true`, the user paused the download.
	DownloadPaused *bool `json:"DownloadPaused,omitempty" plist:"DownloadPaused,omitempty"`
	// If `true`, the user canceled the download.
	DownloadCancelled *bool `json:"DownloadCancelled,omitempty" plist:"DownloadCancelled,omitempty"`
	// If `true`, the app is an App Clip. Available in iOS 16 and later.
	IsAppClip *bool `json:"IsAppClip,omitempty" plist:"IsAppClip,omitempty"`
	// The source of the application. When the app is managed by Declarative Device Management this value is `Declarative Device Management`.
	Source *string `json:"Source,omitempty" plist:"Source,omitempty"`
}

// Invite a user to join the Volume Purchase Program (VPP).
// This command allows a server to invite a user to join a program. This command issues the invitation, but does not allow the server to monitor whether the user has joined the program. This command is supported in the user channel. This command will yield a NotNow status until the user exits Setup Assistant. This command does not work with Account Driven Device Enrollment.
type InviteToProgramCommand struct {
	// The program's identifier, which can only be `com.apple.cloudvpp`.
	ProgramID ProgramID `json:"ProgramID" plist:"ProgramID" required:"true"`
	// The Volume Purchase Program (VPP) invitation URL.
	InvitationURL string `json:"InvitationURL" plist:"InvitationURL" required:"true"`
}

func (p *InviteToProgramCommand) RequestType() string {
	return "InviteToProgram"
}

type InviteToProgramCommandResponse struct {
	// The result of the command.
	InvitationResult InvitationResult `json:"InvitationResult" plist:"InvitationResult" required:"true"`
}

// The program's identifier, which can only be `com.apple.cloudvpp`.
type ProgramID string

const (
	ProgramIDComapplecloudvpp ProgramID = "com.apple.cloudvpp"
)

// The result of the command.
type InvitationResult string

const (
	InvitationResultAcknowledged         InvitationResult = "Acknowledged"
	InvitationResultInvalidProgramID     InvitationResult = "InvalidProgramID"
	InvitationResultInvalidInvitationURL InvitationResult = "InvalidInvitationURL"
)

// Get the status of all managed apps on a device.
// This command allows the server to query the status of managed apps.
type ManagedApplicationListCommand struct {
	// The bundle identifiers of the managed apps to include in the response.
	// > Important:
	// > For a watchOS app, the identifier needs to be the watch's bundle identifier, which differs from the main bundle identifier for the iPhone to which the watch is paired. Obtain the watch's bundle identifier for an app with a watch bundle, in the `watchBundleId` key that's part of the Content Metadata query. For more information on this query, see `Getting App and Book Information`.
	Identifiers *[]string `json:"Identifiers,omitempty" plist:"Identifiers,omitempty"`
}

func (p *ManagedApplicationListCommand) RequestType() string {
	return "ManagedApplicationList"
}

type ManagedApplicationListCommandResponse struct {
	// A dictionary that contains status information about each managed app. The response doesn't include apps that Declarative Device Management manages.
	ManagedApplicationList ManagedApplicationList `json:"ManagedApplicationList" plist:"ManagedApplicationList" required:"true"`
}

// A dictionary that contains status information about each managed app. The response doesn't include apps that Declarative Device Management manages.
type ManagedApplicationList struct {
	// The bundle identifier of the managed app.
	ANYappidentifier ANYappidentifier `json:"ANY app identifier" plist:"ANY app identifier" required:"true"`
}

// The bundle identifier of the managed app.
type ANYappidentifier struct {
	// The status of the managed app, which is one of the following values:
	// - `Queued`: The app is scheduled for installation.
	// - `NeedsRedemption`: The app needs a redemption code to complete installation.
	// - `Redeeming`: The device is redeeming the redemption code for the app.
	// - `Prompting`: The app installation is prompting the user.
	// - `PromptingForLogin' - The app installation is prompting the user for App Store credentials.
	// - `ValidatingPurchase`: Validation of the app purchase is occurring.
	// - `PromptingForUpdate`: An app update is prompting the user.
	// - `PromptingForUpdateLogin`: An app update is prompting the user for App Store credentials.
	// - `PromptingForManagement`: Changing the app to a managed app is prompting the user.
	// - `ValidatingUpdate`: Validation of an app update is occurring.
	// - `Updating`: The app is updating.
	// - `Installing`: The app is installing.
	// - `Managed`: The installed app is a managed app.
	// - `ManagedButUninstalled`: The app is a managed app and the user removed it. Reinstalling the app reinstates it as a managed app.
	// - `Unknown`: The app state is unknown.
	// The following statuses are transient and report only once:
	// - `UserInstalledApp`: The user installed the app before managed app installation could occur.
	// - `UserRejected`: The user rejected the offer to install the app.
	// - `UpdateRejected`: The user rejected the offer to update the app.
	// - `ManagementRejected`:The user rejected management of an installed app.
	// - `Failed`: The app installation failed.
	Status ANYappidentifierStatus `json:"Status" plist:"Status" required:"true"`
	// The bitwise OR of the following management flags:
	// * '1': Remove app upon removal of MDM profile.
	// * '4': Prevent backup of app data.
	ManagementFlags int64 `json:"ManagementFlags" plist:"ManagementFlags" required:"true"`
	// If the user already purchased a paid app, this code is available for use by another user. This code reports only once. This value is available in iOS 5 and later.
	UnusedRedemptionCode string `json:"UnusedRedemptionCode" plist:"UnusedRedemptionCode" required:"true"`
	// If 'true', the app has an update available. This key is present only for App Store apps. In macOS, this key is present only for Volume Purchase Program (VPP) apps. This status updates daily and isn't always up-to-date when installing an app.
	HasConfiguration bool `json:"HasConfiguration" plist:"HasConfiguration" required:"true"`
	// If 'true', the app has feedback for the server. This value is available in iOS 7 and later, and tvOS 10.2 and later. On macOS 11.3 and later, this value is available if the request was sent on the user channel.
	HasFeedback bool `json:"HasFeedback" plist:"HasFeedback" required:"true"`
	// If 'true', the app is valid and can run on the device. If the app is enterprise-distributed and unvalidated, it won't be able to run until validation has occurred. This value is available in iOS 9.2 and later, and tvOS 10.2 and later.
	IsValidated bool `json:"IsValidated" plist:"IsValidated" required:"true"`
	// The app's external version identifier, which you can use in the iTunes Search API to determine if an updated version of the app is available. Compare this value to the 'externalId' value in the 'contentMetadataLookupUrl' response from the 'VPPServiceConfigSrv' endpoint. If these values don't match, an updated version of the app may be available. This value is available in iOS 10.3 and later, macOS 11.3 and later, and tvOS 10.2 and later.
	// A newer version of an app may not be available for installation on the device for a variety of reasons, including that the device's operating system version or hardware is incompatible with the available version of the app.
	ExternalVersionIdentifier int64 `json:"ExternalVersionIdentifier" plist:"ExternalVersionIdentifier" required:"true"`
}

// The status of the managed app, which is one of the following values:
// - `Queued`: The app is scheduled for installation.
// - `NeedsRedemption`: The app needs a redemption code to complete installation.
// - `Redeeming`: The device is redeeming the redemption code for the app.
// - `Prompting`: The app installation is prompting the user.
// - `PromptingForLogin' - The app installation is prompting the user for App Store credentials.
// - `ValidatingPurchase`: Validation of the app purchase is occurring.
// - `PromptingForUpdate`: An app update is prompting the user.
// - `PromptingForUpdateLogin`: An app update is prompting the user for App Store credentials.
// - `PromptingForManagement`: Changing the app to a managed app is prompting the user.
// - `ValidatingUpdate`: Validation of an app update is occurring.
// - `Updating`: The app is updating.
// - `Installing`: The app is installing.
// - `Managed`: The installed app is a managed app.
// - `ManagedButUninstalled`: The app is a managed app and the user removed it. Reinstalling the app reinstates it as a managed app.
// - `Unknown`: The app state is unknown.
// The following statuses are transient and report only once:
// - `UserInstalledApp`: The user installed the app before managed app installation could occur.
// - `UserRejected`: The user rejected the offer to install the app.
// - `UpdateRejected`: The user rejected the offer to update the app.
// - `ManagementRejected`:The user rejected management of an installed app.
// - `Failed`: The app installation failed.
type ANYappidentifierStatus string

const (
	ANYappidentifierStatusQueued                  ANYappidentifierStatus = "Queued"
	ANYappidentifierStatusNeedsRedemption         ANYappidentifierStatus = "NeedsRedemption"
	ANYappidentifierStatusRedeeming               ANYappidentifierStatus = "Redeeming"
	ANYappidentifierStatusPrompting               ANYappidentifierStatus = "Prompting"
	ANYappidentifierStatusPromptingForLogin       ANYappidentifierStatus = "PromptingForLogin"
	ANYappidentifierStatusValidatingPurchase      ANYappidentifierStatus = "ValidatingPurchase"
	ANYappidentifierStatusPromptingForUpdate      ANYappidentifierStatus = "PromptingForUpdate"
	ANYappidentifierStatusPromptingForUpdateLogin ANYappidentifierStatus = "PromptingForUpdateLogin"
	ANYappidentifierStatusPromptingForManagement  ANYappidentifierStatus = "PromptingForManagement"
	ANYappidentifierStatusValidatingUpdate        ANYappidentifierStatus = "ValidatingUpdate"
	ANYappidentifierStatusUpdating                ANYappidentifierStatus = "Updating"
	ANYappidentifierStatusInstalling              ANYappidentifierStatus = "Installing"
	ANYappidentifierStatusManaged                 ANYappidentifierStatus = "Managed"
	ANYappidentifierStatusManagedButUninstalled   ANYappidentifierStatus = "ManagedButUninstalled"
	ANYappidentifierStatusUnknown                 ANYappidentifierStatus = "Unknown"
	ANYappidentifierStatusUserInstalledApp        ANYappidentifierStatus = "UserInstalledApp"
	ANYappidentifierStatusUserRejected            ANYappidentifierStatus = "UserRejected"
	ANYappidentifierStatusUpdateRejected          ANYappidentifierStatus = "UpdateRejected"
	ANYappidentifierStatusManagementRejected      ANYappidentifierStatus = "ManagementRejected"
	ANYappidentifierStatusFailed                  ANYappidentifierStatus = "Failed"
)

// Complete the installation of an app using a redemption code.
// If a redemption code is needed during app installation, the server can use this command to complete the app installation.
type ApplyRedemptionCodeCommand struct {
	// The bundle identifier of the app.
	Identifier string `json:"Identifier" plist:"Identifier" required:"true"`
	// The redemption code that applies to the app pending installation.
	RedemptionCode string `json:"RedemptionCode" plist:"RedemptionCode" required:"true"`
}

func (p *ApplyRedemptionCodeCommand) RequestType() string {
	return "ApplyRedemptionCode"
}

type ApplyRedemptionCodeCommandResponse struct{}

// Remove an installed managed app.
// This command allows a server to remove managed apps. Starting in iOS 26 on supervised devices, this command will also allow a server to remove unmanaged and system apps. This command will fail for apps that are managed by Declarative Device Management.
type RemoveApplicationCommand struct {
	// The bundle identifier of the managed app.
	// > Important:
	// > For a watchOS app, the identifier needs to be the watch's bundle identifier, which differs from the main bundle identifier for the iPhone to which the watch is paired. Obtain the watch's bundle identifier for an app with a watch bundle, in the `watchBundleId` key that's part of the Content Metadata query. For more information on this query, see `Getting App and Book Information`.
	Identifier string `json:"Identifier" plist:"Identifier" required:"true"`
}

func (p *RemoveApplicationCommand) RequestType() string {
	return "RemoveApplication"
}

type RemoveApplicationCommandResponse struct{}

// Force validation of developer and universal provisioning profiles for enterprise apps.
// This command allows the server to query for installed 3rd party applications.
type ValidateApplicationsCommand struct {
	// The bundle identifiers of the enterprise apps to include for validation of associated provisioning profiles, if you choose to provide them. Otherwise, validation occurs for the provisioning profiles for the installed managed apps.
	Identifiers *[]string `json:"Identifiers,omitempty" plist:"Identifiers,omitempty"`
}

func (p *ValidateApplicationsCommand) RequestType() string {
	return "ValidateApplications"
}

type ValidateApplicationsCommandResponse struct{}

// Get a list of installed certificates on a device.
// This command allows the server to retrieve the list of installed certificates on the device. The command requires that the server has the Inspect Profile Manifest privilege.
// For userenrollment, this request will limit to certificates pushed via MDM.
type CertificateListCommand struct {
	// If `true`, only include certificates that MDM installed or that are in the same profile as the MDM payload. User-enrolled devices ignore this value and always only include managed certificates. This value is available in iOS 13 and later, macOS 10.15 and later, and tvOS 13 and later.
	ManagedOnly *bool `json:"ManagedOnly,omitempty" plist:"ManagedOnly,omitempty"`
}

func (p *CertificateListCommand) RequestType() string {
	return "CertificateList"
}

type CertificateListCommandResponse struct {
	// An array of certificate list items that describes each certificate.
	CertificateList []*CertificateListItem `json:"CertificateList" plist:"CertificateList" required:"true"`
}

// A dictionary that contains information about a certificate list item.
type CertificateListItem struct {
	// The certificate's common name.
	CommonName string `json:"CommonName" plist:"CommonName" required:"true"`
	// If `true`, this is an identity certificate.
	IsIdentity bool `json:"IsIdentity" plist:"IsIdentity" required:"true"`
	// The certificate in DER-encoded X.509 format.
	Data []byte `json:"Data" plist:"Data" required:"true"`
}

// Enable your server to support declarative management or trigger a declarative management synchronization operation on the device.
// This command allows the server to turn on the declarative management engine on the device (the first time it is used), or to trigger a declarative management synchronization operation.
type DeclarativeManagementCommand struct {
	// The base64-encoded declarative management JSON request using a `TokensResponse`.
	Data *[]byte `json:"Data,omitempty" plist:"Data,omitempty"`
}

func (p *DeclarativeManagementCommand) RequestType() string {
	return "DeclarativeManagement"
}

type DeclarativeManagementCommandResponse struct{}

// Get the code to bypass Activation Lock on a device.
// Retrieves the Activation Lock bypass code from the device. This bypass code is only available for 15 days after supervision.
type ActivationLockBypassCodeCommand struct{}

func (p *ActivationLockBypassCodeCommand) RequestType() string {
	return "ActivationLockBypassCode"
}

type ActivationLockBypassCodeCommandResponse struct {
	// The Activation Lock bypass code if it's available.
	ActivationLockBypassCode string `json:"ActivationLockBypassCode" plist:"ActivationLockBypassCode" required:"true"`
}

// Clear the Activation Lock bypass code on a device.
// Clears the Activation Lock bypass code from the device.
type ClearActivationLockBypassCodeCommand struct{}

func (p *ClearActivationLockBypassCodeCommand) RequestType() string {
	return "ClearActivationLockBypassCode"
}

type ClearActivationLockBypassCodeCommandResponse struct{}

// Inform the device that it can allow the user to continue in Setup Assistant.
// Informs the device that it can continue past DEP enrollment. Only works on devices in DEP that have their cloud configuration set to await configuration.
type DeviceConfiguredCommand struct{}

func (p *DeviceConfiguredCommand) RequestType() string {
	return "DeviceConfigured"
}

type DeviceConfiguredCommandResponse struct{}

// Remotely and immediately erase a device.
// This command allows the server to remotely erase the device. This command requires the Device Erase right.
type EraseDeviceCommand struct {
	// If `true`, preserve the data plan on an iPhone or iPad with eSIM functionality, if one exists. This value is available in iOS 11 and later.
	PreserveDataPlan *bool `json:"PreserveDataPlan,omitempty" plist:"PreserveDataPlan,omitempty"`
	// If `true`, disable Proximity Setup on the next reboot and skip the pane in Setup Assistant. This value is available in iOS 11 and later. Prior to iOS 14, don't use this option with any other option.
	DisallowProximitySetup *bool `json:"DisallowProximitySetup,omitempty" plist:"DisallowProximitySetup,omitempty"`
	// The six-character PIN for Find My. This value is available in macOS 10.8 and later.
	PIN *string `json:"PIN,omitempty" plist:"PIN,omitempty"`
	// This key defines the fallback behavior for erasing a device.
	// In macOS 12 and later, this command uses Erase All Content and Settings (EACS) on Mac computers with the Apple M1 chip or the Apple T2 Security Chip. On those devices, if EACS can't run, the device can use obliteration (macOS 11.x behavior). This key has no effect on machines prior to the T2 chip. For a list of supported macs, see [Mac models with the Apple T2 Security Chip](https://support.apple.com/en-us/HT208862).
	// Upon receiving this command, the device performs preflight checks to determine if the device is in a state that allows EACS. The `status` of the `EraseDeviceResponse` is either `Acknowledged` or `Error`.
	// The following values define the device's fallback behavior:
	// - `DoNotObliterate`: If EACS preflight fails, the device responds to the server with an `Error` status and doesn't attempt to erase itself.
	// If EACS preflight succeeds, but EACS fails, the device doesn't attempt to erase itself.
	// - `ObliterateWithWarning`: If EACS preflight fails, the device responds with an `Acknowledged` status and then attempts to erase itself.
	// If EACS preflight succeeds, but EACS fails, the device attempts to erase itself.
	// - `Always`: The system doesn't attempt EACS. T2 and later devices always obliterate.
	// - `Default`: If EACS preflight fails, the device responds to the server with an `Error` status and then attempts to erase itself.
	// If EACS preflight succeeds, but EACS fails, the device attempts to erase itself.
	ObliterationBehavior *ObliterationBehavior `json:"ObliterationBehavior,omitempty" plist:"ObliterationBehavior,omitempty"`
	// The configuration settings for return to service. This value is available in iOS 17 and later, with Shared iPad, in tvOS 18 and later, and in visionOS 26 and later.
	ReturnToService *ReturnToService `json:"ReturnToService,omitempty" plist:"ReturnToService,omitempty"`
}

func (p *EraseDeviceCommand) RequestType() string {
	return "EraseDevice"
}

type EraseDeviceCommandResponse struct{}

// This key defines the fallback behavior for erasing a device.
// In macOS 12 and later, this command uses Erase All Content and Settings (EACS) on Mac computers with the Apple M1 chip or the Apple T2 Security Chip. On those devices, if EACS can't run, the device can use obliteration (macOS 11.x behavior). This key has no effect on machines prior to the T2 chip. For a list of supported macs, see [Mac models with the Apple T2 Security Chip](https://support.apple.com/en-us/HT208862).
// Upon receiving this command, the device performs preflight checks to determine if the device is in a state that allows EACS. The `status` of the `EraseDeviceResponse` is either `Acknowledged` or `Error`.
// The following values define the device's fallback behavior:
// - `DoNotObliterate`: If EACS preflight fails, the device responds to the server with an `Error` status and doesn't attempt to erase itself.
// If EACS preflight succeeds, but EACS fails, the device doesn't attempt to erase itself.
// - `ObliterateWithWarning`: If EACS preflight fails, the device responds with an `Acknowledged` status and then attempts to erase itself.
// If EACS preflight succeeds, but EACS fails, the device attempts to erase itself.
// - `Always`: The system doesn't attempt EACS. T2 and later devices always obliterate.
// - `Default`: If EACS preflight fails, the device responds to the server with an `Error` status and then attempts to erase itself.
// If EACS preflight succeeds, but EACS fails, the device attempts to erase itself.
type ObliterationBehavior string

const (
	ObliterationBehaviorDefault               ObliterationBehavior = "Default"
	ObliterationBehaviorDoNotObliterate       ObliterationBehavior = "DoNotObliterate"
	ObliterationBehaviorObliterateWithWarning ObliterationBehavior = "ObliterateWithWarning"
	ObliterationBehaviorAlways                ObliterationBehavior = "Always"
)

// The configuration settings for return to service. This value is available in iOS 17 and later, with Shared iPad, in tvOS 18 and later, and in visionOS 26 and later.
type ReturnToService struct {
	// If `true`, the device tries to reenroll itself automatically after erasure. The user needs to deactivate all activation locks for this feature to work correctly.
	Enabled bool `json:"Enabled" plist:"Enabled" required:"true"`
	// The Wi-Fi profile that installs after erasure when using return to service. This is required when the device doesn't have an Ethernet or cellular connection.
	WiFiProfileData *[]byte `json:"WiFiProfileData,omitempty" plist:"WiFiProfileData,omitempty"`
	// The MDM profile that installs after erasure when using return to service. This key is required for all unsupervised devices, as well as supervised devices that don't enroll with Automated Device Enrollment. If provided, the device uses this profile directly instead of fetching it from the server. For devices that enroll with Automated Device Enrollment, this key isn't necessary unless the cloud configuration profile of the device contains the `configuration-web-url` key.
	// The cloud configuration still downloads from Apple's servers when the profile contains this key, so the supervision identity, MDM removability, and other settings from the cloud configuration still apply. However, the device doesn't use the specified URL in the cloud configuration to fetch the MDM profile.
	MDMProfileData *[]byte `json:"MDMProfileData,omitempty" plist:"MDMProfileData,omitempty"`
	// The bootstrap token the system uses to implement return to service with app preservation. Required when enabling return to service through the cloud configuration.
	BootstrapToken *[]byte `json:"BootstrapToken,omitempty" plist:"BootstrapToken,omitempty"`
}

// Query a carrier URL for active eSIM cellular-plan profiles on a device.
// Instructs the device to query for active cellular plan eSIM "profiles" (not a profile in the MDM sense)
// at the designated carrier eSIM server URL. This command is only supported on cellular devices, and only
// a subset of those devices support eSIM configuration management. (Need details from CoreTelephony.)
type RefreshCellularPlansCommand struct {
	// The carrier's eSIM server URL to query. Obtain this URL from each carrier separately.
	ESIMServerURL string `json:"eSIMServerURL" plist:"eSIMServerURL" required:"true"`
}

func (p *RefreshCellularPlansCommand) RequestType() string {
	return "RefreshCellularPlans"
}

type RefreshCellularPlansCommandResponse struct{}

// Remotely and immediately lock a device.
// This command allows the server to immediately lock the device. This command requires the Device Lock and Passcode Removal right.
type DeviceLockCommand struct {
	// The message to display on the Lock Screen of the device. This value doesn't apply to a Shared iPad device. This value is available in iOS 4 and later, and macOS 10.14 and later.
	Message *string `json:"Message,omitempty" plist:"Message,omitempty"`
	// The phone number to display on the Lock Screen. This value doesn't apply to a Shared iPad device. This value is available in iOS 7 and later and macOS 11.5 and later (for Mac computers with Apple silicon only).
	PhoneNumber *string `json:"PhoneNumber,omitempty" plist:"PhoneNumber,omitempty"`
	// The six-character PIN for Find My. This value is available in macOS 10.8 and later.
	PIN *string `json:"PIN,omitempty" plist:"PIN,omitempty"`
}

func (p *DeviceLockCommand) RequestType() string {
	return "DeviceLock"
}

type DeviceLockCommandResponse struct {
	// The message result if the command includes a message or phone number, which is one of the following values:
	// - `Success`: The message displayed successfully.
	// - `DeviceInLostMode`: The device is in Lost Mode.
	// - `NoPasscodeSet`: The message didn't display because there isn't a set passcode.
	// - `Unknown`: An unknown error occurred.
	MessageResult *string `json:"MessageResult,omitempty" plist:"MessageResult,omitempty"`
}

// Take the device out of Lost Mode.
// This command allows the server to take the device out of MDM lost mode.
type DisableLostModeCommand struct{}

func (p *DisableLostModeCommand) RequestType() string {
	return "DisableLostMode"
}

type DisableLostModeCommandResponse struct{}

// Enable Lost Mode on a device, which provides a message and phone number on the Lock Screen.
// This command allows the server to put the device in MDM lost mode, with a message, phone number, and footnote text. A message or phone number must be provided.
type EnableLostModeCommand struct {
	// If present, the device displays this text on the Lock Screen. You must provide this value if you don't provide a value for `PhoneNumber`.
	Message *string `json:"Message,omitempty" plist:"Message,omitempty"`
	// If present, the device displays this phone number on the Lock Screen. You must provide this value if you don't provide a value for `Message`.
	PhoneNumber *string `json:"PhoneNumber,omitempty" plist:"PhoneNumber,omitempty"`
	// If present, the device displays this text at the bottom of the Lock Screen.
	Footnote *string `json:"Footnote,omitempty" plist:"Footnote,omitempty"`
}

func (p *EnableLostModeCommand) RequestType() string {
	return "EnableLostMode"
}

type EnableLostModeCommandResponse struct{}

// Request the location of a device when in Lost Mode.
type DeviceLocationCommand struct{}

func (p *DeviceLocationCommand) RequestType() string {
	return "DeviceLocation"
}

type DeviceLocationCommandResponse struct {
	// The latitude of the device's location.
	Latitude float64 `json:"Latitude" plist:"Latitude" required:"true"`
	// The longitude of the device's location.
	Longitude float64 `json:"Longitude" plist:"Longitude" required:"true"`
	// The radius of uncertainty for the location in meters, which is a negative value if the horizontal accuracy is unknown.
	HorizontalAccuracy float64 `json:"HorizontalAccuracy" plist:"HorizontalAccuracy" required:"true"`
	// The accuracy of the altitude value in meters, which is a negative value if the vertical accuracy is unknown.
	VerticalAccuracy float64 `json:"VerticalAccuracy" plist:"VerticalAccuracy" required:"true"`
	// The altitude of the device's location, which is a negative value if the altitude is unknown.
	Altitude float64 `json:"Altitude" plist:"Altitude" required:"true"`
	// The speed of the device in meters per second, which is a negative value if the speed is unknown.
	Speed float64 `json:"Speed" plist:"Speed" required:"true"`
	// The direction the device is traveling, which is a negative value if the course is unknown.
	Course float64 `json:"Course" plist:"Course" required:"true"`
	// The RFC 3339 timestamp of when the server determines the location of the device.
	Timestamp string `json:"Timestamp" plist:"Timestamp" required:"true"`
}

// Play the Lost Mode sound on a device that's in Lost Mode.
// This command allows the server to tell the device to play a sound if it is in MDM Lost Mode. The sound will play until the device is either removed from Lost Mode or a user disables the sound from the device.
type PlayLostModeSoundCommand struct{}

func (p *PlayLostModeSoundCommand) RequestType() string {
	return "PlayLostModeSound"
}

type PlayLostModeSoundCommandResponse struct{}

// Remotely and immediately restart a device.
// This command requires the Device Lock access right. The device will restart immediately.
type RestartDeviceCommand struct {
	// If `true`, the system rebuilds the kernel cache during a device restart. If `BootstrapTokenAllowedForAuthentication` is `true` in the `SecurityInfo` response, the device requests the bootstrap token from the MDM server prior to executing this command. This value is available in macOS 11 and later.
	RebuildKernelCache *bool `json:"RebuildKernelCache,omitempty" plist:"RebuildKernelCache,omitempty"`
	// If `RebuildKernelCache` is `true`, this value specifies the paths to kexts to add to the auxiliary kernel cache since the last kernel cache rebuild. If not present, the system only adds previously discovered kexts to the kernel cache. This value is available in macOS 11 and later.
	KextPaths *[]string `json:"KextPaths,omitempty" plist:"KextPaths,omitempty"`
	// If `true`, notifies the user to restart the device at their convenience. No forced restart occurs unless the device is at `loginwindow` with no logged-in users. The user can dismiss the notification and ignore the request. No further notifications display unless you resend the command.
	// This value is available in macOS 11.3 and later.
	NotifyUser *bool `json:"NotifyUser,omitempty" plist:"NotifyUser,omitempty"`
}

func (p *RestartDeviceCommand) RequestType() string {
	return "RestartDevice"
}

type RestartDeviceCommandResponse struct{}

// Clear the Screen Time password and the restrictions on a device.
type ClearRestrictionsPasswordCommand struct{}

func (p *ClearRestrictionsPasswordCommand) RequestType() string {
	return "ClearRestrictionsPassword"
}

type ClearRestrictionsPasswordCommandResponse struct{}

// Get a list of restrictions on the device.
// This command allows the server to determine what restrictions are being enforced on the device, and the total sum of all restrictions. This command requires the Restrictions Query access right. This technically does work on macOS but it returns a blank dictionary and there no plans to change this behavior.
type RestrictionsCommand struct {
	// If `true`, the device reports restrictions from each profile. This value is available in iOS 4 and later, and tvOS 6.1 and later.
	ProfileRestrictions *bool `json:"ProfileRestrictions,omitempty" plist:"ProfileRestrictions,omitempty"`
}

func (p *RestrictionsCommand) RequestType() string {
	return "Restrictions"
}

type RestrictionsCommandResponse struct {
	// A dictionary that contains the global restrictions in effect. This value is available in iOS 4 and later, and tvOS 6.1 and later.
	GlobalRestrictions GlobalRestrictions `json:"GlobalRestrictions" plist:"GlobalRestrictions" required:"true"`
	// A dictionary that contains dictionaries of restrictions from each profile. This value is only available when `ProfileRestrictions` is `true` in the command. The keys are the identifiers of the profiles. This value is available in iOS 4 and later, and tvOS 6.1 and later.
	ProfileRestrictions ProfileRestrictions `json:"ProfileRestrictions" plist:"ProfileRestrictions" required:"true"`
}

// A dictionary that contains the global restrictions in effect. This value is available in iOS 4 and later, and tvOS 6.1 and later.
type GlobalRestrictions struct {
	// A dictionary of Boolean profile restrictions.
	RestrictedBool *RestrictedBool `json:"restrictedBool,omitempty" plist:"restrictedBool,omitempty"`
	// A dictionary of numeric profile restrictions.
	RestrictedValue *RestrictedValue `json:"restrictedValue,omitempty" plist:"restrictedValue,omitempty"`
	// A dictionary of intersected profile restrictions. Intersected restrictions indicate that new restrictions can only reduce the number of strings in the set.
	Intersection *Intersection `json:"intersection,omitempty" plist:"intersection,omitempty"`
	// A dictionary of unioned profile restrictions. Unioned restrictions indicate that new restrictions can add to the set.
	Union *Union `json:"union,omitempty" plist:"union,omitempty"`
}

// A dictionary of Boolean profile restrictions.
type RestrictedBool struct {
	// The Boolean restriction parameters.
	ANYrestrictionname *RestrictedBoolANYrestrictionname `json:"ANY restriction name,omitempty" plist:"ANY restriction name,omitempty"`
}

// The Boolean restriction parameters.
type RestrictedBoolANYrestrictionname struct {
	// The value of the restriction.
	Value bool `json:"value" plist:"value" required:"true"`
}

// A dictionary of numeric profile restrictions.
type RestrictedValue struct {
	// The numeric restriction parameters.
	ANYrestrictionname *RestrictedValueANYrestrictionname `json:"ANY restriction name,omitempty" plist:"ANY restriction name,omitempty"`
}

// The numeric restriction parameters.
type RestrictedValueANYrestrictionname struct {
	// The value of the restriction.
	Value int64 `json:"value" plist:"value" required:"true"`
}

// A dictionary of intersected profile restrictions. Intersected restrictions indicate that new restrictions can only reduce the number of strings in the set.
type Intersection struct {
	// The intersected restriction parameters.
	ANYrestrictionname *IntersectionANYrestrictionname `json:"ANY restriction name,omitempty" plist:"ANY restriction name,omitempty"`
}

// The intersected restriction parameters.
type IntersectionANYrestrictionname struct {
	// The values of the restriction.
	Values []string `json:"values" plist:"values" required:"true"`
}

// A dictionary of unioned profile restrictions. Unioned restrictions indicate that new restrictions can add to the set.
type Union struct {
	// The unioned restriction parameters.
	ANYrestrictionname *UnionANYrestrictionname `json:"ANY restriction name,omitempty" plist:"ANY restriction name,omitempty"`
}

// The unioned restriction parameters.
type UnionANYrestrictionname struct {
	// The values of the restriction.
	Values []string `json:"values" plist:"values" required:"true"`
}

// A dictionary that contains dictionaries of restrictions from each profile. This value is only available when `ProfileRestrictions` is `true` in the command. The keys are the identifiers of the profiles. This value is available in iOS 4 and later, and tvOS 6.1 and later.
type ProfileRestrictions struct {
	// The profile identifiers. This dictionary is only available if `ProfileRestrictions` is `true` in the command.
	ANYprofileidentifier *ANYprofileidentifier `json:"ANY profile identifier,omitempty" plist:"ANY profile identifier,omitempty"`
}

// The profile identifiers. This dictionary is only available if `ProfileRestrictions` is `true` in the command.
type ANYprofileidentifier struct {
	// A dictionary of Boolean profile restrictions.
	RestrictedBool *RestrictedBool `json:"restrictedBool,omitempty" plist:"restrictedBool,omitempty"`
	// A dictionary of numeric profile restrictions.
	RestrictedValue *RestrictedValue `json:"restrictedValue,omitempty" plist:"restrictedValue,omitempty"`
	// A dictionary of intersected profile restrictions. Intersected restrictions indicate that new restrictions can only reduce the number of strings in the set.
	Intersection *Intersection `json:"intersection,omitempty" plist:"intersection,omitempty"`
	// A dictionary of unioned profile restrictions. Unioned restrictions indicate that new restrictions can add to the set.
	Union *Union `json:"union,omitempty" plist:"union,omitempty"`
}

// Remotely and immediately shut down a device.
// This command requires the Device Lock access right. The device will shut down immediately.
type ShutDownDeviceCommand struct{}

func (p *ShutDownDeviceCommand) RequestType() string {
	return "ShutDownDevice"
}

type ShutDownDeviceCommandResponse struct{}

// Get the status of the content caches on a device.
// This command allows the server to query for information about Content Caching.
type ContentCachingInformationCommand struct{}

func (p *ContentCachingInformationCommand) RequestType() string {
	return "ContentCachingInformation"
}

type ContentCachingInformationCommandResponse struct {
	// A dictionary that contains the status of content caching on a device.
	StatusResponse StatusResponse `json:"StatusResponse" plist:"StatusResponse" required:"true"`
}

// A dictionary that contains the status of content caching on a device.
type StatusResponse struct {
	// If `true`, the device has enabled content caching. Enabling content caching doesn't guarantee service. See the `Active` key for the readiness of content caching to serve requests.
	Activated *bool `json:"Activated,omitempty" plist:"Activated,omitempty"`
	// If `true`, content caching is ready to serve requests.
	Active *bool `json:"Active,omitempty" plist:"Active,omitempty"`
	// The actual amount of disk space, in bytes, that cached content uses. See related values `CacheUsed` and `PersonalCacheUsed`.
	ActualCacheUsed *int64 `json:"ActualCacheUsed,omitempty" plist:"ActualCacheUsed,omitempty"`
	// The error conditions the content cache detected in the `PeerFilterRanges` in the installed `com.apple.AssetCache.managed` payload.
	// To display these alerts on the device, set `DisplayAlerts` to `true` in the installed `ContentCaching` profile.
	AlertsForPeerFilterRanges *AlertsForPeerFilterRanges `json:"AlertsForPeerFilterRanges,omitempty" plist:"AlertsForPeerFilterRanges,omitempty"`
	// An array that contains the error conditions the content cache detected that aren't related to peer filter ranges, parent content caches, or peer content caches.
	// See `AlertsForPeerFilterRanges` for errors related to peer filter ranges.
	// See `Parents` and `Peers` for errors related to parent and peer content caches.
	// To display these alerts on the device, set `DisplayAlerts` to `true` in the installed `ContentCaching` profile.
	Alerts *[]AlertsItem `json:"Alerts,omitempty" plist:"Alerts,omitempty"`
	// The amount of disk space that various categories of cached content use. Apple defines these categories and they're subject to change.
	CacheDetails *CacheDetails `json:"CacheDetails,omitempty" plist:"CacheDetails,omitempty"`
	// The amount of disk space, in bytes, available to the content cache.
	CacheFree *int64 `json:"CacheFree,omitempty" plist:"CacheFree,omitempty"`
	// The maximum amount of disk space, in bytes, available to the content cache. A value of `0` indicates an unlimited amount. This value corresponds to `CacheLimit` in the installed `ContentCaching` profile.
	CacheLimit *int64 `json:"CacheLimit,omitempty" plist:"CacheLimit,omitempty"`
	// The level of cache pressure. `LowSpace` means cache pressure is high.
	CacheStatus *CacheStatus `json:"CacheStatus,omitempty" plist:"CacheStatus,omitempty"`
	// The amount of disk space, in bytes, cached content uses. Content caching allocates space in its cache for entire files even when it stores only part of those files in its cache.
	CacheUsed *int64 `json:"CacheUsed,omitempty" plist:"CacheUsed,omitempty"`
	// If `true`, the content cache finished moving from one volume to another.
	DataMigrationCompleted *bool `json:"DataMigrationCompleted,omitempty" plist:"DataMigrationCompleted,omitempty"`
	// The error that occurred while the content cache moved from one volume to another.
	DataMigrationError *DataMigrationError `json:"DataMigrationError,omitempty" plist:"DataMigrationError,omitempty"`
	// A floating-point number between `0.0` and `1.0` that indicates the percentage of progress in moving the content cache from one volume to another. A value of `1.0` indicates that the content cache has fully migrated.
	DataMigrationProgress *float64 `json:"DataMigrationProgress,omitempty" plist:"DataMigrationProgress,omitempty"`
	// A floating-point number between `0.0` and `1.0` that represents how often the cache needed more disk space over the last hour of operation. A lower value is better.
	MaxCachePressureLast1Hour *float64 `json:"MaxCachePressureLast1Hour,omitempty" plist:"MaxCachePressureLast1Hour,omitempty"`
	// An array of dictionaries that describes parent content caches.
	Parents *[]*ParentsItem `json:"Parents,omitempty" plist:"Parents,omitempty"`
	// An array of dictionaries that describes peer content caches.
	Peers *[]*PeersItem `json:"Peers,omitempty" plist:"Peers,omitempty"`
	// The amount of disk space, in bytes, available to the content cache for personal iCloud content.
	PersonalCacheFree *int64 `json:"PersonalCacheFree,omitempty" plist:"PersonalCacheFree,omitempty"`
	// The maximum amount of disk space, in bytes, available to the content cache for personal iCloud content. A value of `0` indicates an unlimited amount.
	PersonalCacheLimit *int64 `json:"PersonalCacheLimit,omitempty" plist:"PersonalCacheLimit,omitempty"`
	// The amount of disk space, in bytes, available to the content cache for personal iCloud content.
	PersonalCacheUsed *int64 `json:"PersonalCacheUsed,omitempty" plist:"PersonalCacheUsed,omitempty"`
	// The IP port number the content cache listens to for requests from clients, peers, and children.
	Port *int64 `json:"Port,omitempty" plist:"Port,omitempty"`
	// An array of the content cache's local IP addresses.
	PrivateAddresses *[]string `json:"PrivateAddresses,omitempty" plist:"PrivateAddresses,omitempty"`
	// The public IP address of the content cache.
	PublicAddress *string `json:"PublicAddress,omitempty" plist:"PublicAddress,omitempty"`
	// If present, the reason the content cache failed to register itself with Apple.
	RegistrationError *string `json:"RegistrationError,omitempty" plist:"RegistrationError,omitempty"`
	// If present, the HTTP response code the content cache received when it failed to register itself with Apple.
	RegistrationResponseCode *int64 `json:"RegistrationResponseCode,omitempty" plist:"RegistrationResponseCode,omitempty"`
	// The date when the content cache began registering itself with Apple. This value is only available during registration attempts.
	RegistrationStarted *time.Time `json:"RegistrationStarted,omitempty" plist:"RegistrationStarted,omitempty"`
	// The status of the content cache's registration with Apple, which is one of the following values:
	// - `-1:` Failed
	// -  `0:` Pending
	// -  `1:` Succeeded
	RegistrationStatus *RegistrationStatus `json:"RegistrationStatus,omitempty" plist:"RegistrationStatus,omitempty"`
	// If `true`, a restriction prevents caching of certain content types.
	RestrictedMedia *bool `json:"RestrictedMedia,omitempty" plist:"RestrictedMedia,omitempty"`
	// The unique identifier of the content cache.
	ServerGUID *string `json:"ServerGUID,omitempty" plist:"ServerGUID,omitempty"`
	// The status of the content cache's registration with Apple.
	StartupStatus *StartupStatus `json:"StartupStatus,omitempty" plist:"StartupStatus,omitempty"`
	// The status of tethered caching, which is content caching with a shared internet connection, which is one of the following values:
	// - `-1:` Unknown
	// -  `0:` Disabled
	// -  `1:` Enabled
	TetheratorStatus *TetheratorStatus `json:"TetheratorStatus,omitempty" plist:"TetheratorStatus,omitempty"`
	// The start date to use when collecting data for the other `TotalBytes` values.
	TotalBytesAreSince *time.Time `json:"TotalBytesAreSince,omitempty" plist:"TotalBytesAreSince,omitempty"`
	// The amount of data, in bytes, that the content cache downloaded, but couldn't add to its cache, since the `TotalBytesAreSince` date.
	TotalBytesDropped *int64 `json:"TotalBytesDropped,omitempty" plist:"TotalBytesDropped,omitempty"`
	// The amount of data, in bytes, that the content cache received since the `TotalBytesAreSince` date.
	TotalBytesImported *int64 `json:"TotalBytesImported,omitempty" plist:"TotalBytesImported,omitempty"`
	// The amount of data, in bytes, that the content cache served to its child content cache since the `TotalBytesAreSince` date.
	TotalBytesReturnedToChildren *int64 `json:"TotalBytesReturnedToChildren,omitempty" plist:"TotalBytesReturnedToChildren,omitempty"`
	// The amount of data, in bytes, that the content cache served to client iOS, macOS, and tvOS devices since the `TotalBytesAreSince` date.
	TotalBytesReturnedToClients *int64 `json:"TotalBytesReturnedToClients,omitempty" plist:"TotalBytesReturnedToClients,omitempty"`
	// The amount of data, in bytes, that the content cache served to peer content caches since the `TotalBytesAreSince` date.
	TotalBytesReturnedToPeers *int64 `json:"TotalBytesReturnedToPeers,omitempty" plist:"TotalBytesReturnedToPeers,omitempty"`
	// The amount of data, in bytes, that the content cache saved from the internet since the `TotalBytesAreSince` date.
	TotalBytesStoredFromOrigin *int64 `json:"TotalBytesStoredFromOrigin,omitempty" plist:"TotalBytesStoredFromOrigin,omitempty"`
	// The amount of data, in bytes, that the content cache saved from parent content caches since the `TotalBytesAreSince` date.
	TotalBytesStoredFromParents *int64 `json:"TotalBytesStoredFromParents,omitempty" plist:"TotalBytesStoredFromParents,omitempty"`
	// The amount of data, in bytes, that the content cache saved from peer content caches since the `TotalBytesAreSince` date.
	TotalBytesStoredFromPeers *int64 `json:"TotalBytesStoredFromPeers,omitempty" plist:"TotalBytesStoredFromPeers,omitempty"`
}

// The error conditions the content cache detected in the `PeerFilterRanges` in the installed `com.apple.AssetCache.managed` payload.
// To display these alerts on the device, set `DisplayAlerts` to `true` in the installed `ContentCaching` profile.
type AlertsForPeerFilterRanges struct {
	// A dictionary that describes the alerts for the peer filter ranges. The key name is the index into the `PeerFilterRanges` array in the installed `com.apple.AssetCache.managed` payload.
	ANYindex ANYindex `json:"ANY index" plist:"ANY index" required:"true"`
}

// A dictionary that describes the alerts for the peer filter ranges. The key name is the index into the `PeerFilterRanges` array in the installed `com.apple.AssetCache.managed` payload.
type ANYindex struct {
	// The type of the alert.
	ClassName ANYindexClassName `json:"className" plist:"className" required:"true"`
	// The date of the alert.
	PostDate time.Time `json:"postDate" plist:"postDate" required:"true"`
	// The index into the `PeerFilterRanges` in the installed ContentCaching payload.
	PeerFilterRangeIndex int64 `json:"peerFilterRangeIndex" plist:"peerFilterRangeIndex" required:"true"`
	// An array of local IP addresses of peer content caches that rejected requests from the content cache.
	Addresses []string `json:"addresses" plist:"addresses" required:"true"`
}

// The type of the alert.
type ANYindexClassName string

const (
	ANYindexClassNameAssetCacheUnfriendlyPeersInFilterRangeAlert ANYindexClassName = "AssetCacheUnfriendlyPeersInFilterRangeAlert"
)

// A dictionary that describes an alert from the content cache.
type AlertsItem struct {
	// The type of the alert.
	ClassName AlertsItemClassName `json:"className" plist:"className" required:"true"`
	// The date of the alert.
	PostDate time.Time `json:"postDate" plist:"postDate" required:"true"`
	// The limit, in bytes, for the content cache at the time of the alert. This value only applies to `AssetCacheLowSpaceAlert` and `AssetCacheNoSpaceAlert` types.
	CacheLimit *int64 `json:"cacheLimit,omitempty" plist:"cacheLimit,omitempty"`
	// The space, in bytes, that the system reserves at the time of the alert. This value only applies to the `AssetCacheLowSpaceAlert` and `AssetCacheNoSpaceAlert` types.
	ReservedVolumeSpace *int64 `json:"reservedVolumeSpace,omitempty" plist:"reservedVolumeSpace,omitempty"`
	// The resource that was missing or inaccessible at the time of the alert. This value only applies to the `AssetCacheResourceMissingAlert` type.
	Resource *string `json:"resource,omitempty" plist:"resource,omitempty"`
	// The subpath of the resource that was missing or inaccessible at the time of the alert. This value only applies to the `AssetCacheResourceMissingAlert` type.
	PathPreventingAccess *string `json:"pathPreventingAccess,omitempty" plist:"pathPreventingAccess,omitempty"`
}

// The type of the alert.
type AlertsItemClassName string

const (
	AlertsItemClassNameAssetCacheLowSpaceAlert                AlertsItemClassName = "AssetCacheLowSpaceAlert"
	AlertsItemClassNameAssetCacheNoSpaceAlert                 AlertsItemClassName = "AssetCacheNoSpaceAlert"
	AlertsItemClassNameAssetCacheRegistrationRejectedAlert    AlertsItemClassName = "AssetCacheRegistrationRejectedAlert"
	AlertsItemClassNameAssetCacheRegistrationUnavailableAlert AlertsItemClassName = "AssetCacheRegistrationUnavailableAlert"
	AlertsItemClassNameAssetCacheResourceMissingAlert         AlertsItemClassName = "AssetCacheResourceMissingAlert"
)

// The amount of disk space that various categories of cached content use. Apple defines these categories and they're subject to change.
type CacheDetails struct {
	// The amount of disk space, in bytes, that this category of cached content uses.
	CategoryName int64 `json:"Category Name" plist:"Category Name" required:"true"`
}

// The level of cache pressure. `LowSpace` means cache pressure is high.
type CacheStatus string

const (
	CacheStatusLOWSPACE CacheStatus = "LOWSPACE"
	CacheStatusOK       CacheStatus = "OK"
)

// The error that occurred while the content cache moved from one volume to another.
type DataMigrationError struct {
	// The error domain.
	Domain string `json:"domain" plist:"domain" required:"true"`
	// The error code.
	Code int64 `json:"code" plist:"code" required:"true"`
	// A dictionary that contains additional information about the error.
	UserInfo *map[string]any `json:"userInfo,omitempty" plist:"userInfo,omitempty"`
}

// A dictionary that describes a parent content cache.
type ParentsItem struct {
	// The local IP address of the parent content cache.
	Address string `json:"address" plist:"address" required:"true"`
	// A dictionary that describes an alert related to the parent content cache.
	Alert *ParentsItemAlert `json:"alert,omitempty" plist:"alert,omitempty"`
	// A dictionary that contains additional details about the parent content cache.
	Details ParentsItemDetails `json:"details" plist:"details" required:"true"`
	// The unique identifier of the parent content cache.
	Guid string `json:"guid" plist:"guid" required:"true"`
	// If `true,` the parent content cache is able to respond to requests from this content cache.
	Healthy bool `json:"healthy" plist:"healthy" required:"true"`
	// The IP port number the parent content cache listens to for requests.
	Port int64 `json:"port" plist:"port" required:"true"`
	// The version number of the parent content cache software.
	Version string `json:"version" plist:"version" required:"true"`
}

// A dictionary that describes an alert related to the parent content cache.
type ParentsItemAlert struct {
	// The type of the alert.
	ClassName ParentsItemAlertClassName `json:"className" plist:"className" required:"true"`
	// The date of the alert.
	PostDate time.Time `json:"postDate" plist:"postDate" required:"true"`
	// An array of local IP addresses of parent content caches.
	Addresses []string `json:"addresses" plist:"addresses" required:"true"`
}

// The type of the alert.
type ParentsItemAlertClassName string

const (
	ParentsItemAlertClassNameAssetCacheParentCycleAlert ParentsItemAlertClassName = "AssetCacheParentCycleAlert"
	ParentsItemAlertClassNameAssetCacheParentDepthAlert ParentsItemAlertClassName = "AssetCacheParentDepthAlert"
)

// A dictionary that contains additional details about the parent content cache.
type ParentsItemDetails struct {
	// If `true`, the parent content cache power source is AC; otherwise, an internal battery provides its power.
	AcPower *bool `json:"ac-power,omitempty" plist:"ac-power,omitempty"`
	// The maximum amount of disk space, in bytes, available to the parent content cache.
	CacheSize *int64 `json:"cache-size,omitempty" plist:"cache-size,omitempty"`
	// A dictionary that describes the capabilities of the parent content cache.
	Capabilities *ParentsItemDetailsCapabilities `json:"capabilities,omitempty" plist:"capabilities,omitempty"`
	// If `true`, the parent content cache computer is portable; for example, a laptop.
	IsPortable *bool `json:"is-portable,omitempty" plist:"is-portable,omitempty"`
	// A dictionary that describes the parent content cache's connection to its local network.
	LocalNetwork *ParentsItemDetailsLocalNetwork `json:"local-network,omitempty" plist:"local-network,omitempty"`
}

// A dictionary that describes the capabilities of the parent content cache.
type ParentsItemDetailsCapabilities struct {
	// If `true`, the parent content cache is capable of imports and uploads.
	Im *bool `json:"im,omitempty" plist:"im,omitempty"`
	// If `true`, the parent content cache is capable of handling namespaces, which is an aspect of personal caching.
	Ns *bool `json:"ns,omitempty" plist:"ns,omitempty"`
	// If `true`, the parent content cache is capable of caching personal iCloud content.
	Pc *bool `json:"pc,omitempty" plist:"pc,omitempty"`
	// If `true`, the parent content cache is capable of handling query parameters in URLs.
	QueryParameters *bool `json:"query-parameters,omitempty" plist:"query-parameters,omitempty"`
	// If `true`, the parent content cache is capable of caching shared non-iCloud content.
	Sc *bool `json:"sc,omitempty" plist:"sc,omitempty"`
	// If `true`, the parent content cache is capable of prioritizing imports and uploads.
	Ur *bool `json:"ur,omitempty" plist:"ur,omitempty"`
}

// A dictionary that describes the parent content cache's connection to its local network.
type ParentsItemDetailsLocalNetwork struct {
	// The transfer speed, in megabits per second, of the parent content cache's connection to its local network.
	Speed *int64 `json:"speed,omitempty" plist:"speed,omitempty"`
	// If `true`, the parent content cache has a wired connection to its local network. If `false`, it has a wireless connection; for example, Wi-Fi.
	Wired *bool `json:"wired,omitempty" plist:"wired,omitempty"`
}

// A dictionary that describes a peer content cache.
type PeersItem struct {
	// The local IP address of the peer content cache.
	Address string `json:"address" plist:"address" required:"true"`
	// A dictionary that describes an alert related to the peer content cache.
	Alert *PeersItemAlert `json:"alert,omitempty" plist:"alert,omitempty"`
	// A dictionary that contains additional details about the peer content cache.
	Details PeersItemDetails `json:"details" plist:"details" required:"true"`
	// If `true`, the peer content cache is able to respond to requests from the content cache.
	Friendly bool `json:"friendly" plist:"friendly" required:"true"`
	// The unique identifier of the peer content cache.
	Guid string `json:"guid" plist:"guid" required:"true"`
	// If `true`, the peer content cache is able to respond to requests from the content cache.
	Healthy bool `json:"healthy" plist:"healthy" required:"true"`
	// The IP port number the peer content cache listens to for requests.
	Port int64 `json:"port" plist:"port" required:"true"`
	// The version number of the peer content cache software.
	Version string `json:"version" plist:"version" required:"true"`
}

// A dictionary that describes an alert related to the peer content cache.
type PeersItemAlert struct {
	// The type of the alert.
	ClassName PeersItemAlertClassName `json:"className" plist:"className" required:"true"`
	// The date of the alert.
	PostDate time.Time `json:"postDate" plist:"postDate" required:"true"`
	// An array of local IP addresses of peer content caches.
	Addresses *[]string `json:"addresses,omitempty" plist:"addresses,omitempty"`
	// The local IP address of a peer content cache.
	PeerAddress *string `json:"peerAddress,omitempty" plist:"peerAddress,omitempty"`
}

// The type of the alert.
type PeersItemAlertClassName string

const (
	PeersItemAlertClassNameAssetCachePeerCycleAlert      PeersItemAlertClassName = "AssetCachePeerCycleAlert"
	PeersItemAlertClassNameAssetCacheUnfriendlyPeerAlert PeersItemAlertClassName = "AssetCacheUnfriendlyPeerAlert"
)

// A dictionary that contains additional details about the peer content cache.
type PeersItemDetails struct {
	// If `true`, the peer content cache power source is AC; otherwise, an internal battery provides its power.
	AcPower *bool `json:"ac-power,omitempty" plist:"ac-power,omitempty"`
	// The maximum amount of disk space, in bytes, available to the peer content cache.
	CacheSize *int64 `json:"cache-size,omitempty" plist:"cache-size,omitempty"`
	// A dictionary that describes the capabilities of the peer content cache.
	Capabilities *PeersItemDetailsCapabilities `json:"capabilities,omitempty" plist:"capabilities,omitempty"`
	// If `true`, the peer content cache computer is portable; for example, a laptop.
	IsPortable *bool `json:"is-portable,omitempty" plist:"is-portable,omitempty"`
	// A dictionary that describes the peer content cache's connection to its local network.
	LocalNetwork *PeersItemDetailsLocalNetwork `json:"local-network,omitempty" plist:"local-network,omitempty"`
}

// A dictionary that describes the capabilities of the peer content cache.
type PeersItemDetailsCapabilities struct {
	// If `true`, the peer content cache is capable of imports and uploads.
	Im *bool `json:"im,omitempty" plist:"im,omitempty"`
	// If `true`, the peer content cache is capable of handling namespaces, which is an aspect of personal caching.
	Ns *bool `json:"ns,omitempty" plist:"ns,omitempty"`
	// If `true`, the peer content cache is capable of caching personal iCloud content.
	Pc *bool `json:"pc,omitempty" plist:"pc,omitempty"`
	// If `true`, the peer content cache is capable of handling query parameters in URLs.
	QueryParameters *bool `json:"query-parameters,omitempty" plist:"query-parameters,omitempty"`
	// If `true`, the peer content cache is capable of caching shared non-iCloud content.
	Sc *bool `json:"sc,omitempty" plist:"sc,omitempty"`
	// If `true`, the peer content cache is capable of prioritizing imports and uploads.
	Ur *bool `json:"ur,omitempty" plist:"ur,omitempty"`
}

// A dictionary that describes the peer content cache's connection to its local network.
type PeersItemDetailsLocalNetwork struct {
	// The transfer speed, in megabits per second, of the peer content cache's connection to its local network.
	Speed *int64 `json:"speed,omitempty" plist:"speed,omitempty"`
	// If `true`, the peer content cache has a wired connection to its local network. If `false`, it has a wireless connection; for example, Wi-Fi.
	Wired *bool `json:"wired,omitempty" plist:"wired,omitempty"`
}

// The status of the content cache's registration with Apple, which is one of the following values:
// - `-1:` Failed
// -  `0:` Pending
// -  `1:` Succeeded
type RegistrationStatus int64

const (
	RegistrationStatusNeg1 RegistrationStatus = -1
	RegistrationStatus0    RegistrationStatus = 0
	RegistrationStatus1    RegistrationStatus = 1
)

// The status of the content cache's registration with Apple.
type StartupStatus string

const (
	StartupStatusFAILED        StartupStatus = "FAILED"
	StartupStatusMIGRATINGDATA StartupStatus = "MIGRATING_DATA"
	StartupStatusOK            StartupStatus = "OK"
	StartupStatusPENDING       StartupStatus = "PENDING"
)

// The status of tethered caching, which is content caching with a shared internet connection, which is one of the following values:
// - `-1:` Unknown
// -  `0:` Disabled
// -  `1:` Enabled
type TetheratorStatus int64

const (
	TetheratorStatusNeg1 TetheratorStatus = -1
	TetheratorStatus0    TetheratorStatus = 0
	TetheratorStatus1    TetheratorStatus = 1
)

// Get detailed information about a device.
// This command allows the server to query for specific device information. It's supported in the user channel.
type DeviceInformationCommand struct {
	// An array of query dictionaries to get information about a device.
	Queries []string `json:"Queries" plist:"Queries" required:"true"`
	// A freshness code that appears in the resulting attestation. This value can contain up to 32 bytes of data. If specified, queries need to contain `DevicePropertiesAttestation`.
	// The MDM server uses this value to prove that an attestation was recently generated. The system caches the most recently generated attestation on the device. If omitted or if the value matches the cached attestation, the system returns the cached attestation. To request a new attestation, provide a new freshness code. Requests for new attestations are rate limited. If it is fewer than 7 days since the system generated an attestation, the device returns the cached attestation rather than generating a new one.
	// Available in iOS 16 and later, macOS 14 and later, tvOS 16 and later, and watchOS 10 and later. The hardware requirements for attestation are described below.
	DeviceAttestationNonce *[]byte `json:"DeviceAttestationNonce,omitempty" plist:"DeviceAttestationNonce,omitempty"`
}

func (p *DeviceInformationCommand) RequestType() string {
	return "DeviceInformation"
}

type DeviceInformationCommandResponse struct {
	// A dictionary that contains information about the device.
	QueryResponses QueryResponses `json:"QueryResponses" plist:"QueryResponses" required:"true"`
}

// A dictionary that contains information about the device.
type QueryResponses struct {
	// The unique identifier of the device.
	UDID *string `json:"UDID" plist:"UDID"`
	// The device identifier to use in provisioning profiles. This value differs from the UDID on Mac computers with Apple silicon. Available in macOS 11.3 and later.
	ProvisioningUDID *string `json:"ProvisioningUDID" plist:"ProvisioningUDID"`
	// The contents of `OrganizationInfo`.
	OrganizationInfo *QueryResponsesOrganizationInfo `json:"OrganizationInfo" plist:"OrganizationInfo"`
	// The contents of `MDMOptions`.
	MDMOptions *QueryResponsesMDMOptions `json:"MDMOptions" plist:"MDMOptions"`
	// The date of the last iCloud backup. Available in iOS 8 and later.
	LastCloudBackupDate *time.Time `json:"LastCloudBackupDate" plist:"LastCloudBackupDate"`
	// If `true` on the device channel, the device is still waiting for a `Device-Configured-Command` command to continue through Setup Assistant.
	// If `true` on the user channel (Shared iPad only), the device is still waiting for a `User-Configured-Command` command to continue through Setup Assistant and finish login.
	AwaitingConfiguration *bool `json:"AwaitingConfiguration" plist:"AwaitingConfiguration"`
	// If `true`, the device has an active iTunes Store account. Requires the App Installation access right.
	ITunesStoreAccountIsActive *bool `json:"iTunesStoreAccountIsActive" plist:"iTunesStoreAccountIsActive"`
	// A hash of the logged-in iTunes Store account. Also see `GetVppUserRequest`. Requires the App Installation access right.
	ITunesStoreAccountHash *string `json:"iTunesStoreAccountHash" plist:"iTunesStoreAccountHash"`
	// The device name. Requires the Device Information access right.
	DeviceName *string `json:"DeviceName" plist:"DeviceName"`
	// The operating system version. Requires the Device Information access right.
	OSVersion *string `json:"OSVersion" plist:"OSVersion"`
	// The OS update rapid security response version letter.
	SupplementalOSVersionExtra *string `json:"SupplementalOSVersionExtra" plist:"SupplementalOSVersionExtra"`
	// The operating system version. Requires the Device Information access right.
	BuildVersion *string `json:"BuildVersion" plist:"BuildVersion"`
	// The supplemental OS build version.
	SupplementalBuildVersion *string `json:"SupplementalBuildVersion" plist:"SupplementalBuildVersion"`
	// The model name, such as _iPhone_. Requires the Device Information access right.
	ModelName *string `json:"ModelName" plist:"ModelName"`
	// The model. Requires the Device Information access right.
	Model *string `json:"Model" plist:"Model"`
	// The device's hardware model number including region info, for example, `MK1A3LL/A`. Requires the Device Information access right. Requires a Mac computer with Apple silicon on macOS.
	ModelNumber *string `json:"ModelNumber" plist:"ModelNumber"`
	// If `true`, the macOS device uses an Apple silicon chip.
	IsAppleSilicon *bool `json:"IsAppleSilicon" plist:"IsAppleSilicon"`
	// The product name, such as _iPad8,12_. Requires the Device Information access right.
	ProductName *string `json:"ProductName" plist:"ProductName"`
	// The serial number. Requires the Device Information access right.
	SerialNumber *string `json:"SerialNumber" plist:"SerialNumber"`
	// The total capacity in floating-point base-10 gigabytes (GB) on iOS and macOS 12 or later. The capacity is in base-2 gibibytes (GiB) on macOS 11 and earlier. Requires the Device Information access right. Available in iOS 4 and later, and macOS 10.7 and later.
	DeviceCapacity *float64 `json:"DeviceCapacity" plist:"DeviceCapacity"`
	// The available capacity in floating-point base-10 gigabytes (GB) in iOS and macOS 12 or later. The capacity is in base-2 gibibytes (GiB) in macOS 11 and earlier. Requires the Device Information access right. Available in iOS 4 and later, and macOS 10.7 and later.
	AvailableDeviceCapacity *float64 `json:"AvailableDeviceCapacity" plist:"AvailableDeviceCapacity"`
	// The International Mobile Equipment Identity (IMEI) number. Requires the Device Information access right. Available as of iOS 4 and deprecated in iOS 16.
	IMEI *string `json:"IMEI" plist:"IMEI"`
	// The mobile equipment identifier (MEID) number. Requires the Device Information access right. Available as of iOS 4 and deprecated in iOS 16.
	MEID *string `json:"MEID" plist:"MEID"`
	// The modem firmware version. Requires the Device Information access right. Available in iOS 4 and later.
	ModemFirmwareVersion *string `json:"ModemFirmwareVersion" plist:"ModemFirmwareVersion"`
	// The cellular technology type, which is one of the following values:
	// - `0`: None
	// - `1`: GSM
	// - `2`: CDMA
	// - `3`: GSM and CDMA
	// Requires the Device Information access right. Available in iOS 4.2.6 and later.
	CellularTechnology *CellularTechnology `json:"CellularTechnology" plist:"CellularTechnology"`
	// The battery level, between `0.0` and `1.0`, or `-1.0` if MDM can't determine the battery level. Requires the Device Information access right. Available in iOS 5 and later, and macOS 13.3 and later.
	BatteryLevel *float64 `json:"BatteryLevel" plist:"BatteryLevel"`
	// If `true`, the device has an internal battery.
	HasBattery *bool `json:"HasBattery" plist:"HasBattery"`
	// If `true`, it's a supervised device. Requires the Device Information access right. Available in iOS 6 and later, macOS 10.15 and later, and tvOS 9 and later.
	IsSupervised *bool `json:"IsSupervised" plist:"IsSupervised"`
	// If `true`, the device is a Shared iPad. Requires the Device Information access right. Available in iOS 9.3 and later.
	IsMultiUser *bool `json:"IsMultiUser" plist:"IsMultiUser"`
	// If `true`, the device has enabled a device locator service, such as Find My. Requires the Device Information access right. Available in iOS 7 and later.
	IsDeviceLocatorServiceEnabled *bool `json:"IsDeviceLocatorServiceEnabled" plist:"IsDeviceLocatorServiceEnabled"`
	// If `true`, the device has enabled Activation Lock. Requires the Device Information access right. Available as of iOS 7 and macOS 10.9, and deprecated in iOS 16 and macOS 13.
	IsActivationLockEnabled *bool `json:"IsActivationLockEnabled" plist:"IsActivationLockEnabled"`
	// If `true`, the device supports Activation Lock. Also see `IsActivationLockManageable` in `ManagementStatus`. Available in macOS 10.9 and later.
	IsActivationLockSupported *bool `json:"IsActivationLockSupported" plist:"IsActivationLockSupported"`
	// If `true`, the device is in Do Not Disturb (DND) mode. This value is `true` even if DND is only in effect for a locked device. Requires the Device Information access right. Available in iOS 7 and later.
	IsDoNotDisturbInEffect *bool `json:"IsDoNotDisturbInEffect" plist:"IsDoNotDisturbInEffect"`
	// If `true`, the device can receive `PowerON`, `PowerOFF`, and `Reset` commands from a lights-out management (LOM) controller. Available in macOS 11 and later.
	SupportsLOMDevice *bool `json:"SupportsLOMDevice" plist:"SupportsLOMDevice"`
	// The device identifier. Requires the Device Information access right. Available in tvOS 6 and later.
	DeviceID *string `json:"DeviceID" plist:"DeviceID"`
	// The device identifier for Exchange Active Sync (EAS). Requires the Device Information access right. Available in iOS 7 and later.
	EASDeviceIdentifier *string `json:"EASDeviceIdentifier" plist:"EASDeviceIdentifier"`
	// If `true`, the device has enabled iCloud backup. Requires the Device Information access right. Available in iOS 7.1 and later.
	IsCloudBackupEnabled *bool `json:"IsCloudBackupEnabled" plist:"IsCloudBackupEnabled"`
	// An array of the directory GUIDs of the logged-in managed users. If one of these users is currently logged in to the console, the `CurrentConsoleManagedUser` key returns the GUID of that user. Requires the Device Information access right. Available in macOS 10.11 and later.
	ActiveManagedUsers *[]string `json:"ActiveManagedUsers" plist:"ActiveManagedUsers"`
	// The contents of “OSUpdateSettings-dictionary“. Requires the Device Information access right. Available in macOS 10.11 and later.
	OSUpdateSettings *OSUpdateSettings `json:"OSUpdateSettings" plist:"OSUpdateSettings"`
	// The local host name from Bonjour. Available in macOS 10.11 and later.
	LocalHostName *string `json:"LocalHostName" plist:"LocalHostName"`
	// The host name. Available in macOS 10.11 and later.
	HostName *string `json:"HostName" plist:"HostName"`
	// The contents of “AutoSetupAdminAccountsItem“, which Setup Assistant automatically creates during DEP enrollment. Requires the Device Information access right. Available in macOS 10.11 and later.
	AutoSetupAdminAccounts *[]*AutoSetupAdminAccountsItem `json:"AutoSetupAdminAccounts" plist:"AutoSetupAdminAccounts"`
	// If `true`, the device has enabled System Integrity Protection. Requires the Device Information access right. Available in macOS 10.12 and later.
	SystemIntegrityProtectionEnabled *bool `json:"SystemIntegrityProtectionEnabled" plist:"SystemIntegrityProtectionEnabled"`
	// If `true`, the device has enabled Managed Lost Mode. Requires the Device Information access right. Available in iOS 9.3 and later.
	IsMDMLostModeEnabled *bool `json:"IsMDMLostModeEnabled" plist:"IsMDMLostModeEnabled"`
	// The maximum number of users that can use this Shared iPad device. Starting with iOS 13.4, the value that returns is always `32`. Requires the Device Information access right. Available in iOS 9.3 and later.
	MaximumResidentUsers *int64 `json:"MaximumResidentUsers" plist:"MaximumResidentUsers"`
	// The estimated number of users that can use this Shared iPad device, according to the space available on the device and each user's quota. Requires the Device Information access right. Available in iOS 14 and later.
	EstimatedResidentUsers *int64 `json:"EstimatedResidentUsers" plist:"EstimatedResidentUsers"`
	// The quota size in megabytes for each user on this Shared iPad device. Requires the Device Information access right. Available in iOS 13.4 and later.
	QuotaSize *int64 `json:"QuotaSize" plist:"QuotaSize"`
	// The number of users currently on this Shared iPad device. Requires the Device Information access right. Available in iOS 13.4 and later.
	ResidentUsers *int64 `json:"ResidentUsers" plist:"ResidentUsers"`
	// The timeout interval for the user session. A value of `0` indicates that there's no timeout.
	UserSessionTimeout *int64 `json:"UserSessionTimeout" plist:"UserSessionTimeout"`
	// The timeout interval for the temporary session. A value of `0` indicates that there's no timeout.
	TemporarySessionTimeout *int64 `json:"TemporarySessionTimeout" plist:"TemporarySessionTimeout"`
	// If `true`, the device allows only temporary sessions.
	TemporarySessionOnly *bool `json:"TemporarySessionOnly" plist:"TemporarySessionOnly"`
	// The list of domains that the device suggests on the Shared iPad login screen. Available in iOS 16 and later.
	ManagedAppleIDDefaultDomains *[]string `json:"ManagedAppleIDDefaultDomains" plist:"ManagedAppleIDDefaultDomains"`
	// The grace period for Shared iPad online authentication (in days). A value of `0` indicates that the device requires online authentication for every login. Available in iOS 16 and later.
	OnlineAuthenticationGracePeriod *int64 `json:"OnlineAuthenticationGracePeriod" plist:"OnlineAuthenticationGracePeriod"`
	// If `true`, skip the language and country/region panes for new users on Shared iPad.
	SkipLanguageAndLocaleSetupForNewUsers *bool `json:"SkipLanguageAndLocaleSetupForNewUsers" plist:"SkipLanguageAndLocaleSetupForNewUsers"`
	// The push token for the user-channel connection, in the same format as in `TokenUpdateRequest`. MDM ignores this query for the device channel. Requires the Device Information access right. Available in iOS 9.3 and later, and macOS 10.12 and later.
	PushToken *[]byte `json:"PushToken" plist:"PushToken"`
	// If `true`, the device has enabled diagnostic submission. Requires the Device Information access right. Available in iOS 9.3 and later.
	DiagnosticSubmissionEnabled *bool `json:"DiagnosticSubmissionEnabled" plist:"DiagnosticSubmissionEnabled"`
	// If `true`, the device is sharing app analytics. Requires the Device Information access right. Available in iOS 9.3 and later.
	AppAnalyticsEnabled *bool `json:"AppAnalyticsEnabled" plist:"AppAnalyticsEnabled"`
	// The current Internet Assigned Numbers Authority (IANA) time zone database name. Requires the Device Information access right. Available in iOS 14 and later, and tvOS 14 and later.
	TimeZone *string `json:"TimeZone" plist:"TimeZone"`
	// The integrated circuit card (ICC) identifier for the installed SIM card. Requires the Network Information access right. Available as of iOS 4 and deprecated in iOS 16.
	ICCID *string `json:"ICCID" plist:"ICCID"`
	// The Bluetooth media access control (MAC) address. Requires the Network Information access right.
	BluetoothMAC *string `json:"BluetoothMAC" plist:"BluetoothMAC"`
	// The Wi-Fi MAC address. Requires the Network Information access right.
	WiFiMAC *string `json:"WiFiMAC" plist:"WiFiMAC"`
	// The primary Ethernet MAC address. Requires the Network Information access right. Available in macOS 10.7 and later.
	EthernetMAC *string `json:"EthernetMAC" plist:"EthernetMAC"`
	// The name of the current carrier network. Requires the Network Information access right. Available as of iOS 4 and deprecated in iOS 16.
	CurrentCarrierNetwork *string `json:"CurrentCarrierNetwork" plist:"CurrentCarrierNetwork"`
	// Apple no longer supports this query. Use `SubscriberCarrierNetwork` instead.
	SIMCarrierNetwork *string `json:"SIMCarrierNetwork" plist:"SIMCarrierNetwork"`
	// The name of the home carrier network. Requires the Network Information access right. Available as of iOS 5 and deprecated in iOS 16.
	SubscriberCarrierNetwork *string `json:"SubscriberCarrierNetwork" plist:"SubscriberCarrierNetwork"`
	// The version of the carrier settings. Requires the Network Information access right. Available as of iOS 4 and deprecated in iOS 16.
	CarrierSettingsVersion *string `json:"CarrierSettingsVersion" plist:"CarrierSettingsVersion"`
	// The raw phone number without punctuation and including the country code. Requires the Network Information access right. Available as of iOS 4 and deprecated in iOS 16.
	PhoneNumber *string `json:"PhoneNumber" plist:"PhoneNumber"`
	// If `true`, the device has enabled data roaming. Requires the Network Information access right. Available in iOS 5 and later.
	DataRoamingEnabled *bool `json:"DataRoamingEnabled" plist:"DataRoamingEnabled"`
	// If `true`, the device has enabled voice roaming, which isn't available for all carriers. Requires the Network Information access right. Requires the Device Information access right. Available as of iOS 5 and deprecated in iOS 16.
	VoiceRoamingEnabled *bool `json:"VoiceRoamingEnabled" plist:"VoiceRoamingEnabled"`
	// If `true,` the device has enabled Personal Hotspot, which isn't available for all carriers. Requires the Network Information access right. Available in iOS 7 and later.
	PersonalHotspotEnabled *bool `json:"PersonalHotspotEnabled" plist:"PersonalHotspotEnabled"`
	// If `true`, the device is network-tethered. Requires the Network Information access right. Available in iOS 10.3 and later.
	IsNetworkTethered *bool `json:"IsNetworkTethered" plist:"IsNetworkTethered"`
	// If `true`, the device is roaming. Requires the Network Information access right. IAvailable as of iOS 4.2 and deprecated in iOS 16.
	IsRoaming *bool `json:"IsRoaming" plist:"IsRoaming"`
	// Apple no longer supports this query. Use `SubscriberMCC` instead.
	SIMMCC *string `json:"SIMMCC" plist:"SIMMCC"`
	// Apple no longer supports this query. Use `SubscriberMNC` instead.
	SIMMNC *string `json:"SIMMNC" plist:"SIMMNC"`
	// The home Mobile Country Code (MCC). Requires the Network Information access right. Available as of iOS 4.2.6 and deprecated in iOS 16.
	SubscriberMCC *string `json:"SubscriberMCC" plist:"SubscriberMCC"`
	// The key to get the home Mobile Network Code (MNC). Requires the Network Information access right. Available as of iOS 4.2.6 and deprecated in iOS 16.
	SubscriberMNC *string `json:"SubscriberMNC" plist:"SubscriberMNC"`
	// The current mobile country code (MCC). Requires the Network Information access right. Available as of iOS 4 and deprecated in iOS 16.
	CurrentMCC *string `json:"CurrentMCC" plist:"CurrentMCC"`
	// The current mobile network code (MNC). Requires the Network Information access right. Available as of iOS 4 and deprecated in iOS 16.
	CurrentMNC *string `json:"CurrentMNC" plist:"CurrentMNC"`
	// The contents of “ServiceSubscriptionProperty“. Requires the Network Information access right.
	ServiceSubscriptions *[]*ServiceSubscriptionProperty `json:"ServiceSubscriptions" plist:"ServiceSubscriptions"`
	// If `true`, the `EraseDeviceCommand` requires a PIN. Available in macOS 11 and later.
	PINRequiredForEraseDevice *bool `json:"PINRequiredForEraseDevice" plist:"PINRequiredForEraseDevice"`
	// If `true`, the `DeviceLockCommand` requires a PIN. Available in macOS 11 and later.
	PINRequiredForDeviceLock *bool `json:"PINRequiredForDeviceLock" plist:"PINRequiredForDeviceLock"`
	// If `true`, the device supports iOS or iPadOS app installs through MDM. Available in macOS 11 and later.
	SupportsiOSAppInstalls *bool `json:"SupportsiOSAppInstalls" plist:"SupportsiOSAppInstalls"`
	// The device identifier to look up available OS updates through [https://gdmf.apple.com/v2/pmv](https://gdmf.apple.com/v2/pmv). Available in iOS 15 and later, and macOS 12 and later.
	SoftwareUpdateDeviceID *string `json:"SoftwareUpdateDeviceID" plist:"SoftwareUpdateDeviceID"`
	// The device settings that control which updates appear in the Software Update pane in Settings. Available in iOS 14.5 and later.
	SoftwareUpdateSettings *QueryResponsesSoftwareUpdateSettings `json:"SoftwareUpdateSettings" plist:"SoftwareUpdateSettings"`
	// The current state of settable accessibility settings. Available in iOS 16 and later.
	AccessibilitySettings *QueryResponsesAccessibilitySettings `json:"AccessibilitySettings" plist:"AccessibilitySettings"`
	// The key to get an attestation of the device's properties. Available in iOS 16 and later, macOS 14 and later, tvOS 16 and later, and watchOS 10 and later. The hardware requirements for attestation are described below.
	// The value is an array of certificates in DER form that forms a certificate chain. The chain is rooted with the Apple CA `Apple Enterprise Attestation Root CA`. The first array item is the leaf certificate. The leaf certificate contains custom OIDs describing a device. The OS version of the device, and the type of enrollment, determine which OIDs are present in the certificate. If Apple's attestation servers are unable to verify a device property they generate a blank value, omit the OID entirely, or refuse to issue an attestation certificate.
	// The following OIDs were introduced in iOS 16, iPadOS 16, tvOS 16, watchOS 10, visionOS 1 and macOS 14:
	// - `1.2.840.113635.100.8.9.1` serial number: This is the serial number of the device. It is omitted if the enrollment is a user enrollment.
	// - `1.2.840.113635.100.8.9.2` UDID: For a Mac this has the same value as the `ProvisioningUDID` key, and does not match the UDID used elsewhere in the MDM protocol. It is omitted if the enrollment is a user enrollment.
	// - `1.2.840.113635.100.8.10.2` sepOS version: This is the version of the operating system running on the Secure Enclave when the attestation is generated. Typically this matches the version of the main operating system.
	// - `1.2.840.113635.100.8.11.1` Freshness code: This is the freshness code. See the `DeviceAttestationNonce`. This may not match the requested freshness code if a cached attestation was returned.
	// The following OIDs were introduced in iOS 17.2, iPadOS 17.2, tvOS 17.2, watchOS 10.2, visionOS 1.l0, and macOS 14.2:
	// - `1.2.840.113635.100.8.9.4` Software Update Device ID: This is an identifier of the device model. It is expected to match the `SoftwareUpdateDeviceID` in the `DeviceInformation“ response. This is the device identifier to use when looking up available OS updates through <https://gdmf.apple.com/v2/pmv>.
	// - `1.2.840.113635.100.8.10.1` OS Version: This is the version of iOS, iPadOS or tvOS running on the device when the attestation is generated.
	// - `1.2.840.113635.100.8.10.3` LLB Version: This is the version of the Low Level Bootloader firmware running on the device when the attestation is generated. For more information about the boot process, see the documentation of the boot process in the Apple Platform Security guide.
	// The following OIDs were introduced in macOS 14.2:
	// - `1.2.840.113635.100.8.13.1` System Integrity Protection (SIP) status: This indicates whether SIP is enabled or disabled when the attestation is generated. `0` indicates enabled, `1` indicates disabled.
	// - `1.2.840.113635.100.8.13.2` Secure boot status: This describes part of the configuration of the LocalPolicy when the attestation is generated. The values are `Full Security`, `Reduced Security`, or `Permissive Security`. For a description of these values see the Apple Platform Security guide.
	// - `1.2.840.113635.100.8.13.3` Third party kernel extensions allowed: This indicates whether third party kernel extensions are allowed. A value of `0` indicates third party kernel extensions are not allowed. Any other value means that some kinds of third party kernel extensions are allowed.
	DevicePropertiesAttestation *[][]byte `json:"DevicePropertiesAttestation" plist:"DevicePropertiesAttestation"`
	// Specifies whether the device can perform an `EraseDeviceCommand` using Erase All Content and Settings (EACS), which is one of the following values:
	// - `success`: The device supports EACS.
	// - `not supported`: The device is too old to support EACS.
	// - `unknown failure`: A problem occurred for which there isn't a more specific error message.
	// - `(other string)`: A reason why the device can't perform EACS, such as "System is not sealed"
	EACSPreflight *string `json:"EACSPreflight" plist:"EACSPreflight"`
}

// The contents of `OrganizationInfo`.
type QueryResponsesOrganizationInfo struct {
	// A string that describes the organization operating the MDM server. This value is available in iOS 7 and later, macOS 10.11 and later, and tvOS 9 and later.
	OrganizationName string `json:"OrganizationName" plist:"OrganizationName" required:"true"`
	// The organization's address. Use the LF character (`&#10`) to insert line breaks. This value is available in iOS 7 and later, macOS 10.11 and later, and tvOS 9 and later.
	OrganizationAddress *string `json:"OrganizationAddress,omitempty" plist:"OrganizationAddress,omitempty"`
	// The organization's phone number. This value is available in iOS 7 and later, macOS 10.11 and later, and tvOS 9 and later.
	OrganizationPhone *string `json:"OrganizationPhone,omitempty" plist:"OrganizationPhone,omitempty"`
	// The organization's support email address. This value is available in iOS 7 and later, macOS 10.11 and later, and tvOS 9 and later.
	OrganizationEmail *string `json:"OrganizationEmail,omitempty" plist:"OrganizationEmail,omitempty"`
	// A unique identifier for the various services a single organization manages. This value is available in iOS 7 and later, macOS 10.11 and later, and tvOS 9 and later.
	OrganizationMagic *string `json:"OrganizationMagic,omitempty" plist:"OrganizationMagic,omitempty"`
}

// The contents of `MDMOptions`.
type QueryResponsesMDMOptions struct {
	// If `true`, a supervised device registers itself with Activation Lock when the user enables Find My. Unsupervised devices ignore this value. This value is available in iOS 7 and later, macOS 11 and later, and tvOS 9 and later.
	ActivationLockAllowedWhileSupervised *bool `json:"ActivationLockAllowedWhileSupervised,omitempty" plist:"ActivationLockAllowedWhileSupervised,omitempty"`
	// If `true`, the server supports Bootstrap Token commands. This value is available in macOS 11 and later.
	BootstrapTokenAllowed *bool `json:"BootstrapTokenAllowed,omitempty" plist:"BootstrapTokenAllowed,omitempty"`
	// If `true`, the device can accept a Bootstrap Token from the MDM server instead of prompting for user authentication prior to installation. This only applies when `BootstrapTokenAllowedForAuthentication` is `true` in the `SecurityInfo` response. This value is available for Mac computers with Apple silicon in macOS 11 and later.
	PromptUserToAllowBootstrapTokenForAuthentication *bool `json:"PromptUserToAllowBootstrapTokenForAuthentication,omitempty" plist:"PromptUserToAllowBootstrapTokenForAuthentication,omitempty"`
}

// The cellular technology type, which is one of the following values:
// - `0`: None
// - `1`: GSM
// - `2`: CDMA
// - `3`: GSM and CDMA
// Requires the Device Information access right. Available in iOS 4.2.6 and later.
type CellularTechnology int64

const (
	CellularTechnology0 CellularTechnology = 0
	CellularTechnology1 CellularTechnology = 1
	CellularTechnology2 CellularTechnology = 2
	CellularTechnology3 CellularTechnology = 3
)

// The contents of “OSUpdateSettings-dictionary“. Requires the Device Information access right. Available in macOS 10.11 and later.
type OSUpdateSettings struct {
	// The URL to the software update catalog the client is using. This value is available in macOS 10.11 and later.
	CatalogURL *string `json:"CatalogURL,omitempty" plist:"CatalogURL,omitempty"`
	// If `true`, `CatalogURL` is the default catalog. This value is available in macOS 10.11 and later.
	IsDefaultCatalog *bool `json:"IsDefaultCatalog" plist:"IsDefaultCatalog"`
	// The date of the last software update scan. This value is available in macOS 10.11 and later.
	PreviousScanDate *time.Time `json:"PreviousScanDate" plist:"PreviousScanDate"`
	// The result code of last software update scan; `0` = success. This value is available in macOS 10.11 and no longer available in macOS 15 and later.
	PreviousScanResult *string `json:"PreviousScanResult,omitempty" plist:"PreviousScanResult,omitempty"`
	// If `true`, start a new scan. This value is available in macOS 10.11 and later.
	PerformPeriodicCheck *bool `json:"PerformPeriodicCheck" plist:"PerformPeriodicCheck"`
	// The preference to automatically check for app updates. This value is available in macOS 10.11 and later.
	AutoCheckEnabled *bool `json:"AutoCheckEnabled" plist:"AutoCheckEnabled"`
	// The preference to download app updates in the background. This value is available in macOS 10.11 and later.
	BackgroundDownloadEnabled *bool `json:"BackgroundDownloadEnabled" plist:"BackgroundDownloadEnabled"`
	// The preference to automatically install app updates. This value is available in macOS 10.11 and later.
	AutomaticAppInstallationEnabled *bool `json:"AutomaticAppInstallationEnabled" plist:"AutomaticAppInstallationEnabled"`
	// The preference to automatically install operating system updates. This value is available in macOS 10.11 and later.
	AutomaticOSInstallationEnabled *bool `json:"AutomaticOSInstallationEnabled" plist:"AutomaticOSInstallationEnabled"`
	// The preference to automatically install system data files and security updates. This value is available in macOS 10.11 and later.
	AutomaticSecurityUpdatesEnabled *bool `json:"AutomaticSecurityUpdatesEnabled" plist:"AutomaticSecurityUpdatesEnabled"`
}

// The response dictionary that contains the administrator setup information.
type AutoSetupAdminAccountsItem struct {
	// The `GeneratedUID` of the administrator account. This value is available in macOS 10.11 and later.
	GUID *string `json:"GUID" plist:"GUID"`
	// The short name of the administrator account. This value is available in macOS 10.11 and later.
	ShortName *string `json:"shortName" plist:"shortName"`
}

// Properties of this Service Subscription. See below.
type ServiceSubscriptionProperty struct {
	// The version of the carrier settings. This value is available in iOS 12 and later.
	CarrierSettingsVersion *string `json:"CarrierSettingsVersion" plist:"CarrierSettingsVersion"`
	// The name of the current carrier network. This value is available in iOS 12 and later.
	CurrentCarrierNetwork *string `json:"CurrentCarrierNetwork" plist:"CurrentCarrierNetwork"`
	// The current mobile country code (MCC). This value is available in iOS 12 and later.
	CurrentMCC *string `json:"CurrentMCC" plist:"CurrentMCC"`
	// The current mobile network code (MNC). This value is available in iOS 12 and later.
	CurrentMNC *string `json:"CurrentMNC" plist:"CurrentMNC"`
	// The integrated circuit card identifier (ICCID) value. This value is available in iOS 12 and later.
	ICCID *string `json:"ICCID" plist:"ICCID"`
	// The eSIM identifier. This value is available in iOS 14 and later.
	EID *string `json:"EID" plist:"EID"`
	// The device International Mobile Equipment Identity (IMEI) number. This value is available in iOS 12 and later.
	IMEI *string `json:"IMEI" plist:"IMEI"`
	// If `true`, this subscription is the preference for data. This value is available in iOS 12 and later.
	IsDataPreferred *bool `json:"IsDataPreferred" plist:"IsDataPreferred"`
	// If `true`, the phone is roaming. This value is available in iOS 12 and later.
	IsRoaming *bool `json:"IsRoaming" plist:"IsRoaming"`
	// If `true`, this subscription is the preference for voice. This value is available in iOS 12 and later.
	IsVoicePreferred *bool `json:"IsVoicePreferred" plist:"IsVoicePreferred"`
	// The label of this subscription. This value is available in iOS 12 and later.
	Label *string `json:"Label" plist:"Label"`
	// The unique identifier for this subscription. This value is available in iOS 12 and later.
	LabelID *string `json:"LabelID" plist:"LabelID"`
	// The device Mobile Equipment Identifier (MEID) number. This query is available in iOS 12 and later.
	MEID *string `json:"MEID" plist:"MEID"`
	// The raw phone number without punctuation and including country code. This value is available in iOS 12 and later.
	PhoneNumber *string `json:"PhoneNumber" plist:"PhoneNumber"`
	// The description of the slot that contains the SIM representing this subscription. This value is available in iOS 12 and later.
	Slot *string `json:"Slot" plist:"Slot"`
	// The name of the home carrier network. This value is available in iOS 16 and later.
	SubscriberCarrierNetwork *string `json:"SubscriberCarrierNetwork" plist:"SubscriberCarrierNetwork"`
}

// The device settings that control which updates appear in the Software Update pane in Settings. Available in iOS 14.5 and later.
type QueryResponsesSoftwareUpdateSettings struct {
	// Which software updates to present to the user.
	// - `0`: Allows all updates (the default value).
	// - `1`: Allows only older updates.
	// - `2`: Allows only newer updates.
	// No effect if the device qualifies for only a single update.
	RecommendationsCadence *int64 `json:"RecommendationsCadence" plist:"RecommendationsCadence"`
}

// The current state of settable accessibility settings. Available in iOS 16 and later.
type QueryResponsesAccessibilitySettings struct {
	// If `true`, the device has enabled bold text.
	BoldTextEnabled *bool `json:"BoldTextEnabled" plist:"BoldTextEnabled"`
	// If `true`, the device has enabled increase contrast.
	IncreaseContrastEnabled *bool `json:"IncreaseContrastEnabled" plist:"IncreaseContrastEnabled"`
	// If `true`, the device has enabled reduced motion.
	ReduceMotionEnabled *bool `json:"ReduceMotionEnabled" plist:"ReduceMotionEnabled"`
	// If `true`, the device has enabled reduced transparency.
	ReduceTransparencyEnabled *bool `json:"ReduceTransparencyEnabled" plist:"ReduceTransparencyEnabled"`
	// The accessibility text size apps that support dynamic text use. 0 is the smallest value, and 11 is the largest available.
	// `-1` indicates that the current size is unknown or hasn't been explicitly set.
	TextSize *QueryResponsesAccessibilitySettingsTextSize `json:"TextSize" plist:"TextSize"`
	// If `true`, the device has enabled touch accommodations.
	TouchAccommodationsEnabled *bool `json:"TouchAccommodationsEnabled" plist:"TouchAccommodationsEnabled"`
	// If `true`, the device has enabled voiceover.
	VoiceOverEnabled *bool `json:"VoiceOverEnabled" plist:"VoiceOverEnabled"`
	// If `true`, the device has enabled zoom.
	ZoomEnabled *bool `json:"ZoomEnabled" plist:"ZoomEnabled"`
	// If `true`, the device has enabled grayscale display.
	GrayscaleEnabled *bool `json:"GrayscaleEnabled" plist:"GrayscaleEnabled"`
}

// The accessibility text size apps that support dynamic text use. 0 is the smallest value, and 11 is the largest available.
// `-1` indicates that the current size is unknown or hasn't been explicitly set.
type QueryResponsesAccessibilitySettingsTextSize int64

const (
	QueryResponsesAccessibilitySettingsTextSizeNeg1 QueryResponsesAccessibilitySettingsTextSize = -1
	QueryResponsesAccessibilitySettingsTextSize0    QueryResponsesAccessibilitySettingsTextSize = 0
	QueryResponsesAccessibilitySettingsTextSize1    QueryResponsesAccessibilitySettingsTextSize = 1
	QueryResponsesAccessibilitySettingsTextSize2    QueryResponsesAccessibilitySettingsTextSize = 2
	QueryResponsesAccessibilitySettingsTextSize3    QueryResponsesAccessibilitySettingsTextSize = 3
	QueryResponsesAccessibilitySettingsTextSize4    QueryResponsesAccessibilitySettingsTextSize = 4
	QueryResponsesAccessibilitySettingsTextSize5    QueryResponsesAccessibilitySettingsTextSize = 5
	QueryResponsesAccessibilitySettingsTextSize6    QueryResponsesAccessibilitySettingsTextSize = 6
	QueryResponsesAccessibilitySettingsTextSize7    QueryResponsesAccessibilitySettingsTextSize = 7
	QueryResponsesAccessibilitySettingsTextSize8    QueryResponsesAccessibilitySettingsTextSize = 8
	QueryResponsesAccessibilitySettingsTextSize9    QueryResponsesAccessibilitySettingsTextSize = 9
	QueryResponsesAccessibilitySettingsTextSize10   QueryResponsesAccessibilitySettingsTextSize = 10
	QueryResponsesAccessibilitySettingsTextSize11   QueryResponsesAccessibilitySettingsTextSize = 11
)

// Get security-related information about a device.
// This command queries the device for security-related information. Queries are available if the MDM host has the Security Query right.
type SecurityInfoCommand struct{}

func (p *SecurityInfoCommand) RequestType() string {
	return "SecurityInfo"
}

type SecurityInfoCommandResponse struct {
	// A dictionary that contains security-related information.
	SecurityInfo SecurityInfo `json:"SecurityInfo" plist:"SecurityInfo" required:"true"`
}

// A dictionary that contains security-related information.
type SecurityInfo struct {
	// An integer that indicates the underlying hardware encryption capabilities of the device, which is one of the following values:
	// - `1`: Block-level encryption
	// - `2`: File-level encryption
	// - `3`: Both block-level and file-level encryption
	// > Important:
	// > For a device to have data protection, `HardwareEncryptionCaps` must be `3` and `PasscodePresent` must `true`.
	// This value is available in iOS 4 and later, and tvOS 6 and later.
	HardwareEncryptionCaps *int64 `json:"HardwareEncryptionCaps" plist:"HardwareEncryptionCaps"`
	// If `true`, the device has a passcode. This key doesn't apply to User-Enrolled devices. This value is available in iOS 4 and later, and tvOS 6 and later.
	PasscodePresent *bool `json:"PasscodePresent" plist:"PasscodePresent"`
	// If `true`, the user's passcode is compliant with all requirements on the device, including Exchange and other accounts. This value is available in iOS 4 and later, and tvOS 6 and later.
	PasscodeCompliant *bool `json:"PasscodeCompliant" plist:"PasscodeCompliant"`
	// If `true`, the user's passcode is compliant with requirements from profiles. This key doesn't apply to User-Enrolled devices. This value is available in iOS 4 and later, and tvOS 6 and later.
	PasscodeCompliantWithProfiles *bool `json:"PasscodeCompliantWithProfiles" plist:"PasscodeCompliantWithProfiles"`
	// The user preference for the number of seconds before a locked screen requires the device passcode to unlock it. This value is only available for Shared iPad.
	PasscodeLockGracePeriod *int64 `json:"PasscodeLockGracePeriod" plist:"PasscodeLockGracePeriod"`
	// The enforced value for the number of seconds before a locked screen requires the device passcode to unlock it. If a device has a passcode, changing `PasscodeLockGracePeriod` to a larger value doesn't take effect until the user logs out or removes the passcode. This value is only available for Shared iPad.
	PasscodeLockGracePeriodEnforced *int64 `json:"PasscodeLockGracePeriodEnforced" plist:"PasscodeLockGracePeriodEnforced"`
	// The number of seconds before a device goes to sleep after being idle. This value is only available on Shared iPad in iOS 17 and later.
	AutoLockTime *int64 `json:"AutoLockTime" plist:"AutoLockTime"`
	// If `true`, the device has enabled FileVault full disk encryption (FDE). This value is available in macOS 10.9 and later.
	FDEEnabled *bool `json:"FDE_Enabled" plist:"FDE_Enabled"`
	// If `true`, FileVault FDE has a personal recovery key. This value is available in macOS 10.9 and later.
	FDEHasPersonalRecoveryKey *bool `json:"FDE_HasPersonalRecoveryKey" plist:"FDE_HasPersonalRecoveryKey"`
	// If `true`, FileVault FDE has an institutional recovery key. This value is available in macOS 10.9 and later.
	FDEHasInstitutionalRecoveryKey *bool `json:"FDE_HasInstitutionalRecoveryKey" plist:"FDE_HasInstitutionalRecoveryKey"`
	// If the FileVault personal recovery key has enabled escrow with a recovery key, this value contains the key. The certificate from the `FDERecoveryKeyEscrow` profile encrypts the key and wraps it as CMS data. This value is available in macOS 10.13 and later.
	FDEPersonalRecoveryKeyCMS *[]byte `json:"FDE_PersonalRecoveryKeyCMS" plist:"FDE_PersonalRecoveryKeyCMS"`
	// If the FileVault personal recovery key has enabled escrow with a recovery key, this value is the device serial number. This is the value that displays to the user at the EFI Login Window as part of the help message if they enter their password incorrectly three times. The server also uses this value as an index when saving the device personal recovery key. This replaces the `recordNumber` that the server returned in the previous escrow mechanism. This value is available in macOS 10.13 and later.
	FDEPersonalRecoveryKeyDeviceKey *string `json:"FDE_PersonalRecoveryKeyDeviceKey" plist:"FDE_PersonalRecoveryKeyDeviceKey"`
	// If `true`, System Integrity Protection (SIP) is active on the device. This value is available in macOS 10.12 and later.
	SystemIntegrityProtectionEnabled *bool `json:"SystemIntegrityProtectionEnabled" plist:"SystemIntegrityProtectionEnabled"`
	// A dictionary that contains the firewall settings. This value is available in macOS 10.12 and later.
	FirewallSettings *FirewallSettings `json:"FirewallSettings" plist:"FirewallSettings"`
	// A dictionary that contains the status of the EFI firmware password. This value is available in macOS 10.13 and later.
	FirmwarePasswordStatus *FirmwarePasswordStatus `json:"FirmwarePasswordStatus" plist:"FirmwarePasswordStatus"`
	// A dictionary that contains the status of the device's MDM enrollment.
	ManagementStatus *ManagementStatus `json:"ManagementStatus" plist:"ManagementStatus"`
	// A dictionary that contains the device's Secure Boot settings. This value is available in macOS 10.15 and later.
	SecureBoot *SecureBoot `json:"SecureBoot" plist:"SecureBoot"`
	// If `true`, Remote Desktop is active on the device. This value is available in macOS 10.14.4 and later.
	RemoteDesktopEnabled *bool `json:"RemoteDesktopEnabled" plist:"RemoteDesktopEnabled"`
	// If `true`, the system booted using an Authenticated Root Volume. This value is available in macOS 11 and later.
	AuthenticatedRootVolumeEnabled *bool `json:"AuthenticatedRootVolumeEnabled" plist:"AuthenticatedRootVolumeEnabled"`
	// This value specifies whether the Secure Enclave Processor (SEP) supports and allows secure operations to use the Bootstrap Token. The value is automatically set for devices enrolled through the Device Enrollment Program (DEP). The user can also manually set this value in the RecoveryOS.
	// This value is available for Mac computers with Apple silicon in macOS 11 and later. Not available for user enrollment.
	BootstrapTokenAllowedForAuthentication *BootstrapTokenAllowedForAuthentication `json:"BootstrapTokenAllowedForAuthentication" plist:"BootstrapTokenAllowedForAuthentication"`
	// If `true`, the device can accept a Bootstrap Token from the MDM server instead of prompting for user authentication prior to installation. This only applies when `BootstrapTokenAllowedForAuthentication` is `true` in the `SecurityInfo` response.
	// This value is available for Mac computers with Apple silicon in macOS 11 and later. Not available for user enrollment.
	BootstrapTokenRequiredForSoftwareUpdate *bool `json:"BootstrapTokenRequiredForSoftwareUpdate" plist:"BootstrapTokenRequiredForSoftwareUpdate"`
	// If `true`, the device can accept a Bootstrap Token from the MDM server instead of prompting for user authentication prior to enabling kernel extensions. This includes enabling kexts through the `com.apple.syspolicy.kernel-extension-policy` payload or triggering the `RestartDevice` command with `RebuildKernelCache` set to `true`. This only applies when `BootstrapTokenAllowedForAuthentication` is `true` in the `SecurityInfo` response.
	// This value is available for Mac computers with Apple silicon in macOS 11 and later. Not available for user enrollment.
	BootstrapTokenRequiredForKernelExtensionApproval *bool `json:"BootstrapTokenRequiredForKernelExtensionApproval" plist:"BootstrapTokenRequiredForKernelExtensionApproval"`
	// If `true`, a password is required to enter recovery (see `SetRecoveryLockCommand`). Available in macOS 11.5 and later and only on Mac computers with Apple silicon.
	IsRecoveryLockEnabled *bool `json:"IsRecoveryLockEnabled" plist:"IsRecoveryLockEnabled"`
}

// A dictionary that contains the firewall settings. This value is available in macOS 10.12 and later.
type FirewallSettings struct {
	// If `true`, the firewall is on.
	FirewallEnabled *bool `json:"FirewallEnabled" plist:"FirewallEnabled"`
	// If `true`, the firewall blocks all incoming connections.
	BlockAllIncoming *bool `json:"BlockAllIncoming" plist:"BlockAllIncoming"`
	// If true, stealth mode is active for the firewall.
	StealthMode *bool `json:"StealthMode" plist:"StealthMode"`
	// An array of dictionaries that describes the allowed applications.
	Applications *[]*ApplicationsItem `json:"Applications" plist:"Applications"`
	// If `true`, logging is enabled.
	LoggingEnabled *bool `json:"LoggingEnabled" plist:"LoggingEnabled"`
	// The type of logging emitted by the firewall.
	LoggingOption *LoggingOption `json:"LoggingOption" plist:"LoggingOption"`
}

// A dictionary that describes the allowed apps.
type ApplicationsItem struct {
	// If `true`, the app is an allowed app.
	Allowed *bool `json:"Allowed" plist:"Allowed"`
	// The app's bundle identifier.
	BundleID *string `json:"BundleID" plist:"BundleID"`
	// The app's display name if it's determinable from the `BundleID`.
	Name *string `json:"Name" plist:"Name"`
}

// The type of logging emitted by the firewall.
type LoggingOption string

const (
	LoggingOptionThrottled LoggingOption = "throttled"
	LoggingOptionBrief     LoggingOption = "brief"
	LoggingOptionDetail    LoggingOption = "detail"
)

// A dictionary that contains the status of the EFI firmware password. This value is available in macOS 10.13 and later.
type FirmwarePasswordStatus struct {
	// If `true`, the device has an EFI firmware password.
	PasswordExists *bool `json:"PasswordExists" plist:"PasswordExists"`
	// If `true`, a firmware password change is pending. A device restart is necessary for this change to take effect. Until then, additional attempts to change the password fail.
	// > Note:
	// > If `true`, the other values show the current state of the device, not the state after a restart.
	ChangePending *bool `json:"ChangePending" plist:"ChangePending"`
	// If `true`, enable ROMs.
	AllowOroms *bool `json:"AllowOroms" plist:"AllowOroms"`
}

// A dictionary that contains the status of the device's MDM enrollment.
type ManagementStatus struct {
	// If `true`, the device enrolled in MDM through the Device Enrollment Program (DEP). This value is available in macOS 10.13.2 and later.
	EnrolledViaDEP *bool `json:"EnrolledViaDEP" plist:"EnrolledViaDEP"`
	// If `true`, the enrollment was user-approved. If `false`, the device may reject certain security-sensitive payloads or commands. This value is available in macOS 10.13.2 and later.
	UserApprovedEnrollment *bool `json:"UserApprovedEnrollment" plist:"UserApprovedEnrollment"`
	// If `true`, the device is user-enrolled. This value is available in iOS 13 and later, and macOS 10.15 and later.
	IsUserEnrollment *bool `json:"IsUserEnrollment" plist:"IsUserEnrollment"`
	// If `true`, the type of enrollment allows the MDM to manage Activation Lock for this device. This value is available in macOS 10.15 and later.
	IsActivationLockManageable *bool `json:"IsActivationLockManageable" plist:"IsActivationLockManageable"`
}

// A dictionary that contains the device's Secure Boot settings. This value is available in macOS 10.15 and later.
type SecureBoot struct {
	// The security level for the bootable operating system versions.
	SecureBootLevel *SecureBootLevel `json:"SecureBootLevel" plist:"SecureBootLevel"`
	// The device's external boot level, which indicates whether it allows booting from an external device, disallows it, or doesn't support it.
	ExternalBootLevel *ExternalBootLevel `json:"ExternalBootLevel" plist:"ExternalBootLevel"`
	// Reports which security features the user disables in `recoveryOS`. This property is only present for Mac computers with Apple silicon when `SecureBootLevel` is `medium`.
	// Available in iOS 11 and later.
	ReducedSecurity *[]string `json:"ReducedSecurity" plist:"ReducedSecurity"`
}

// The security level for the bootable operating system versions.
type SecureBootLevel string

const (
	SecureBootLevelOff          SecureBootLevel = "off"
	SecureBootLevelMedium       SecureBootLevel = "medium"
	SecureBootLevelFull         SecureBootLevel = "full"
	SecureBootLevelNotsupported SecureBootLevel = "not supported"
)

// The device's external boot level, which indicates whether it allows booting from an external device, disallows it, or doesn't support it.
type ExternalBootLevel string

const (
	ExternalBootLevelAllowed      ExternalBootLevel = "allowed"
	ExternalBootLevelDisallowed   ExternalBootLevel = "disallowed"
	ExternalBootLevelNotsupported ExternalBootLevel = "not supported"
)

// This value specifies whether the Secure Enclave Processor (SEP) supports and allows secure operations to use the Bootstrap Token. The value is automatically set for devices enrolled through the Device Enrollment Program (DEP). The user can also manually set this value in the RecoveryOS.
// This value is available for Mac computers with Apple silicon in macOS 11 and later. Not available for user enrollment.
type BootstrapTokenAllowedForAuthentication string

const (
	BootstrapTokenAllowedForAuthenticationAllowed      BootstrapTokenAllowedForAuthentication = "allowed"
	BootstrapTokenAllowedForAuthenticationDisallowed   BootstrapTokenAllowedForAuthentication = "disallowed"
	BootstrapTokenAllowedForAuthenticationNotsupported BootstrapTokenAllowedForAuthentication = "not supported"
)

// Send requests to a device using lights-out management (LOM).
// Used to send LOM requests ("PowerON", "PowerOFF", "Reset") to LOM Controller which then forwards the request to LOM Devices.
type LOMDeviceRequestCommand struct {
	// An array of requests to perform.
	RequestList []RequestListItem `json:"RequestList" plist:"RequestList" required:"true"`
}

func (p *LOMDeviceRequestCommand) RequestType() string {
	return "LOMDeviceRequest"
}

type LOMDeviceRequestCommandResponse struct {
	// An array of dictionaries that describes the status of each request.
	ResponseList []ResponseListItem `json:"ResponseList" plist:"ResponseList" required:"true"`
}

// A dictionary that contains a requested action to perform on a device using lights-out management (LOM).
type RequestListItem struct {
	// The requested action to perform on the device.
	DeviceRequestType DeviceRequestType `json:"DeviceRequestType" plist:"DeviceRequestType" required:"true"`
	// The unique identifier of the request.
	DeviceRequestUUID string `json:"DeviceRequestUUID" plist:"DeviceRequestUUID" required:"true"`
	// The DNS name of the device. This should match the `dNSName` in `SubjectAltName` or an equivalent in a PKCS12 identity.
	DeviceDNSName string `json:"DeviceDNSName" plist:"DeviceDNSName" required:"true"`
	// An array that contains the IPv6 addresses for primary LOM-compatible Ethernet interfaces for the device.
	PrimaryIPv6AddressList []string `json:"PrimaryIPv6AddressList" plist:"PrimaryIPv6AddressList" required:"true"`
	// An array that contains the IPv6 addresses for secondary LOM-compatible Ethernet interfaces for the device.
	SecondaryIPv6AddressList []string `json:"SecondaryIPv6AddressList" plist:"SecondaryIPv6AddressList" required:"true"`
	// The LOM protocol version that the device supports. Provide the same value that `LOMProtocolVersion` receives in the `LOMSetupRequestResponse`.
	LOMProtocolVersion int64 `json:"LOMProtocolVersion" plist:"LOMProtocolVersion" required:"true"`
}

// The requested action to perform on the device.
type DeviceRequestType string

const (
	DeviceRequestTypePowerON  DeviceRequestType = "PowerON"
	DeviceRequestTypePowerOFF DeviceRequestType = "PowerOFF"
	DeviceRequestTypeReset    DeviceRequestType = "Reset"
)

// A dictionary that describes a response list item.
type ResponseListItem struct {
	// If `true`, the request was successful.
	DeviceRequestSuccess bool `json:"DeviceRequestSuccess" plist:"DeviceRequestSuccess" required:"true"`
	// The unique identifier of the request for this response list item.
	DeviceRequestUUID string `json:"DeviceRequestUUID" plist:"DeviceRequestUUID" required:"true"`
	// If present, a description of the error for a failed request.
	DeviceRequestReturnError *string `json:"DeviceRequestReturnError,omitempty" plist:"DeviceRequestReturnError,omitempty"`
}

// Get information from a device to set up lights-out management (LOM).
// Queries the device for LOM setup information such as IP addresses, protocol version, etc. The MDM server must send this command prior to sending the LOMDeviceRequest command.
type LOMSetupRequestCommand struct{}

func (p *LOMSetupRequestCommand) RequestType() string {
	return "LOMSetupRequest"
}

type LOMSetupRequestCommandResponse struct {
	// An array that contains the IPv6 addresses for primary LOM-compatible Ethernet interfaces for the device.
	PrimaryIPv6AddressList []string `json:"PrimaryIPv6AddressList" plist:"PrimaryIPv6AddressList" required:"true"`
	// An array that contains the IPv6 addresses for secondary LOM-compatible Ethernet interfaces for the device.
	SecondaryIPv6AddressList []string `json:"SecondaryIPv6AddressList" plist:"SecondaryIPv6AddressList" required:"true"`
	// The LOM protocol version that the device supports.
	LOMProtocolVersion int64 `json:"LOMProtocolVersion" plist:"LOMProtocolVersion" required:"true"`
}

// Query attributes in managed apps on a device.
// Queries managed application attributes. Attributes can be set on managed apps. These attributes can be changed over time. The response will not include apps that are managed by Declarative Device Management.
type ManagedApplicationAttributesCommand struct {
	// The bundle identifiers of the managed apps.
	// > Important:
	// > For a watchOS app, the identifier needs to be the watch's bundle identifier, which differs from the main bundle identifier for the iPhone to which the watch is paired. Obtain the watch's bundle identifier for an app with a watch bundle, in the `watchBundleId` key that's part of the Content Metadata query. For more information on this query, see `Getting App and Book Information`.
	Identifiers []string `json:"Identifiers" plist:"Identifiers" required:"true"`
}

func (p *ManagedApplicationAttributesCommand) RequestType() string {
	return "ManagedApplicationAttributes"
}

type ManagedApplicationAttributesCommandResponse struct {
	// An array of app attribute items.
	ApplicationAttributes []*ApplicationAttributesApplicationAttributesItem `json:"ApplicationAttributes" plist:"ApplicationAttributes" required:"true"`
}

// A dictionary that contains a managed app attributes item.
type ApplicationAttributesApplicationAttributesItem struct {
	// The app's bundle identifier.
	// > Note:
	// > For a watchOS app, the identifier is the watch's bundle identifier, which differs from the main bundle identifier for the iPhone to which the watch is paired.
	Identifier string `json:"Identifier" plist:"Identifier" required:"true"`
	// The app's attributes.
	Attributes *ApplicationAttributesItemAttributes `json:"Attributes,omitempty" plist:"Attributes,omitempty"`
}

// The app's attributes.
type ApplicationAttributesItemAttributes struct {
	// A per-app VPN unique identifier for this app.
	VPNUUID *string `json:"VPNUUID,omitempty" plist:"VPNUUID,omitempty"`
	// The content Filter UUID assigned to this app.
	// Available in iOS 16 and later.
	ContentFilterUUID *string `json:"ContentFilterUUID,omitempty" plist:"ContentFilterUUID,omitempty"`
	// The DNS Proxy UUID assigned to this app.
	// Available in iOS 16 and later.
	DNSProxyUUID *string `json:"DNSProxyUUID,omitempty" plist:"DNSProxyUUID,omitempty"`
	// The relay UUID for this app.
	RelayUUID *string `json:"RelayUUID,omitempty" plist:"RelayUUID,omitempty"`
	// This app's associated domains. This value is available in iOS 13 and later.
	AssociatedDomains *[]string `json:"AssociatedDomains,omitempty" plist:"AssociatedDomains,omitempty"`
	// If `true`, perform claimed site association verification directly at the domain instead of on Apple's servers. Only set this to `true` for domains that can't access the internet. This value is available in iOS 14 and later.
	AssociatedDomainsEnableDirectDownloads *bool `json:"AssociatedDomainsEnableDirectDownloads,omitempty" plist:"AssociatedDomainsEnableDirectDownloads,omitempty"`
	// If `false`, this app isn't removable while it's a managed app. This value is available in iOS 14 and later.
	Removable *bool `json:"Removable,omitempty" plist:"Removable,omitempty"`
	// If `true`, Tap to Pay on iPhone requires users to use Face ID or a passcode to unlock their device after every transaction that requires a customer's card PIN. If `false`, the user can configure this setting on their device.
	// Available in iOS 16.4 and later.
	TapToPayScreenLock *bool `json:"TapToPayScreenLock,omitempty" plist:"TapToPayScreenLock,omitempty"`
	// The data network name (DNN) or app category. For DNN, the value is `DNN:name`, where `name` is the carrier-provided DNN name. For app category, the value is `AppCategory:category`, where `category` is a carrier-provided string like "Enterprise1".
	// Available in iOS 17 and later.
	CellularSliceUUID *string `json:"CellularSliceUUID,omitempty" plist:"CellularSliceUUID,omitempty"`
	// If `false`, the system prevents the user from hiding the app. It doesn't affect the user's ability to leave it in the App Library, while removing it from the Home Screen.
	Hideable *bool `json:"Hideable,omitempty" plist:"Hideable,omitempty"`
	// If `false`, the system prevents the user from locking the app. This also prevents the user from hiding the app.
	Lockable *bool `json:"Lockable,omitempty" plist:"Lockable,omitempty"`
}

// Get app configurations from managed apps on a device.
// This command queries the device for the current configuration of managed applications. This command requires the App Management right. The response will not include apps that are managed by Declarative Device Management.
type ManagedApplicationConfigurationCommand struct {
	// The bundle identifiers of the managed apps.
	// > Important:
	// > For a watchOS app, the identifier needs to be the watch's bundle identifier, which differs from the main bundle identifier for the iPhone to which the watch is paired. Obtain the watch's bundle identifier for an app with a watch bundle, in the `watchBundleId` key that's part of the Content Metadata query. For more information on this query, see `Getting App and Book Information`.
	Identifiers []string `json:"Identifiers" plist:"Identifiers" required:"true"`
}

func (p *ManagedApplicationConfigurationCommand) RequestType() string {
	return "ManagedApplicationConfiguration"
}

type ManagedApplicationConfigurationCommandResponse struct {
	// An array of app configuration items.
	ApplicationConfigurations []*ApplicationConfigurationsItem `json:"ApplicationConfigurations" plist:"ApplicationConfigurations" required:"true"`
}

// A dictionary that contains a managed app's configurations item.
type ApplicationConfigurationsItem struct {
	// The app's bundle identifier.
	// > Note:
	// > For a watchOS app, the identifier is the watch's bundle identifier, which differs from the main bundle identifier for the iPhone to which the watch is paired.
	Identifier string `json:"Identifier" plist:"Identifier" required:"true"`
	// The app's configurations.
	Configuration *map[string]any `json:"Configuration,omitempty" plist:"Configuration,omitempty"`
}

// Get app feedback from a managed app on the device.
// This command queries the device for application feedback information. This command requires the App Management right. The response will not include apps that are managed by Declarative Device Management.
type ManagedApplicationFeedbackCommand struct {
	// The bundle identifiers of the managed apps.
	Identifiers []string `json:"Identifiers" plist:"Identifiers" required:"true"`
	// If `true`, delete the app's feedback dictionary after the server reads it. Apps that are managed by Declarative Device Management are ignored.
	DeleteFeedback *bool `json:"DeleteFeedback,omitempty" plist:"DeleteFeedback,omitempty"`
}

func (p *ManagedApplicationFeedbackCommand) RequestType() string {
	return "ManagedApplicationFeedback"
}

type ManagedApplicationFeedbackCommandResponse struct {
	// An array of managed app feedback items.
	ManagedApplicationFeedback []*ManagedApplicationFeedbackItem `json:"ManagedApplicationFeedback" plist:"ManagedApplicationFeedback" required:"true"`
}

// A dictionary that contains a managed app's feedback item.
type ManagedApplicationFeedbackItem struct {
	// The app's bundle identifier.
	Identifier string `json:"Identifier" plist:"Identifier" required:"true"`
	// The app's feedback.
	Feedback *map[string]any `json:"Feedback,omitempty" plist:"Feedback,omitempty"`
}

// Install a book on a device.
// This command allows the server to install a book on a device. If the book is already being managed, this command will update the book.
type InstallMediaCommand struct {
	// The book's iTunes Store identifier.
	ITunesStoreID *int64 `json:"iTunesStoreID,omitempty" plist:"iTunesStoreID,omitempty"`
	// The URL to retrieve the book. This value is available in iOS 8 and later.
	MediaURL *string `json:"MediaURL,omitempty" plist:"MediaURL,omitempty"`
	// The media type, which can only be `Book`.
	MediaType InstallMediaCommandMediaType `json:"MediaType" plist:"MediaType" required:"true"`
	// The book's persistent identifier in reverse-DNS form; for example, `com.acme.manuals.training`. This value is available in iOS 8 and later.
	PersistentID *string `json:"PersistentID,omitempty" plist:"PersistentID,omitempty"`
	// The kind of the media, which can be one of the following values:
	// - `pdf`: A PDF file
	// - `epub`: An EPUB file in `gzip` format.
	// - `ibooks`: An iBooks Author file in `gzip` format.
	// If you omit this value, its value is the file extension in the URL. This value is available in iOS 8 and later.
	Kind *Kind `json:"Kind,omitempty" plist:"Kind,omitempty"`
	// The book's version number. This value is available in iOS 8 and later.
	Version *string `json:"Version,omitempty" plist:"Version,omitempty"`
	// The name of the book's author. This value is available in iOS 8 and later.
	Author *string `json:"Author,omitempty" plist:"Author,omitempty"`
	// The book's title. This value is available in iOS 8 and later.
	Title *string `json:"Title,omitempty" plist:"Title,omitempty"`
}

func (p *InstallMediaCommand) RequestType() string {
	return "InstallMedia"
}

type InstallMediaCommandResponse struct {
	// The book's iTunes Store identifier, if present in the command.
	ITunesStoreID *int64 `json:"iTunesStoreID,omitempty" plist:"iTunesStoreID,omitempty"`
	// The URL to retrieve the book, if present in the command. This value is available in iOS 8 and later.
	MediaURL *string `json:"MediaURL,omitempty" plist:"MediaURL,omitempty"`
	// The book's persistent identifier, if present in the command. This value is available in iOS 8 and later.
	PersistentID *string `json:"PersistentID,omitempty" plist:"PersistentID,omitempty"`
	// The media type, which can only be `Book`.
	MediaType *string `json:"MediaType,omitempty" plist:"MediaType,omitempty"`
	// The installation state of this book. The `Failed` and `Unknown` states are transient and the device only reports them once. Books from the Book Store report their state as `Installed` instead of `Managed`.
	State *InstallMediaCommandState `json:"State,omitempty" plist:"State,omitempty"`
	// The reason, if installation fails, which is one of the following values:
	// - `CouldNotVerifyITunesStoreID`: The `iTunesStoreID` is invalid.
	// - `PurchaseNotFound`: The Volume Purchase Program (VPP) license isn't in the user's history.
	// - `AppStoreDisabled`: App Store isn't available on the device.
	// - `WrongMediaType`: The media type is invalid. The only valid type is `Book`.
	// - `DownloadInvalid`: The URL doesn't lead to a valid book.
	// - `EnterpriseBooksNotSupportedInMultiUser`: Multiuser mode doesn't support enterprise books.
	RejectionReason *InstallMediaCommandRejectionReason `json:"RejectionReason,omitempty" plist:"RejectionReason,omitempty"`
}

// The media type, which can only be `Book`.
type InstallMediaCommandMediaType string

const (
	InstallMediaCommandMediaTypeBook InstallMediaCommandMediaType = "Book"
)

// The kind of the media, which can be one of the following values:
// - `pdf`: A PDF file
// - `epub`: An EPUB file in `gzip` format.
// - `ibooks`: An iBooks Author file in `gzip` format.
// If you omit this value, its value is the file extension in the URL. This value is available in iOS 8 and later.
type Kind string

const (
	KindPdf    Kind = "pdf"
	KindEpub   Kind = "epub"
	KindIbooks Kind = "ibooks"
)

// The installation state of this book. The `Failed` and `Unknown` states are transient and the device only reports them once. Books from the Book Store report their state as `Installed` instead of `Managed`.
type InstallMediaCommandState string

const (
	InstallMediaCommandStateQueued                InstallMediaCommandState = "Queued"
	InstallMediaCommandStatePromptingForLogin     InstallMediaCommandState = "PromptingForLogin"
	InstallMediaCommandStateUpdating              InstallMediaCommandState = "Updating"
	InstallMediaCommandStateInstalling            InstallMediaCommandState = "Installing"
	InstallMediaCommandStateManaged               InstallMediaCommandState = "Managed"
	InstallMediaCommandStateManagedButUninstalled InstallMediaCommandState = "ManagedButUninstalled"
	InstallMediaCommandStateInstalled             InstallMediaCommandState = "Installed"
	InstallMediaCommandStateUninstalled           InstallMediaCommandState = "Uninstalled"
	InstallMediaCommandStateFailed                InstallMediaCommandState = "Failed"
	InstallMediaCommandStateUnknown               InstallMediaCommandState = "Unknown"
)

// The reason, if installation fails, which is one of the following values:
// - `CouldNotVerifyITunesStoreID`: The `iTunesStoreID` is invalid.
// - `PurchaseNotFound`: The Volume Purchase Program (VPP) license isn't in the user's history.
// - `AppStoreDisabled`: App Store isn't available on the device.
// - `WrongMediaType`: The media type is invalid. The only valid type is `Book`.
// - `DownloadInvalid`: The URL doesn't lead to a valid book.
// - `EnterpriseBooksNotSupportedInMultiUser`: Multiuser mode doesn't support enterprise books.
type InstallMediaCommandRejectionReason string

const (
	InstallMediaCommandRejectionReasonCouldNotVerifyITunesStoreID            InstallMediaCommandRejectionReason = "CouldNotVerifyITunesStoreID"
	InstallMediaCommandRejectionReasonPurchaseNotFound                       InstallMediaCommandRejectionReason = "PurchaseNotFound"
	InstallMediaCommandRejectionReasonAppStoreDisabled                       InstallMediaCommandRejectionReason = "AppStoreDisabled"
	InstallMediaCommandRejectionReasonWrongMediaType                         InstallMediaCommandRejectionReason = "WrongMediaType"
	InstallMediaCommandRejectionReasonDownloadInvalid                        InstallMediaCommandRejectionReason = "DownloadInvalid"
	InstallMediaCommandRejectionReasonEnterpriseBooksNotSupportedInMultiUser InstallMediaCommandRejectionReason = "EnterpriseBooksNotSupportedInMultiUser"
)

// Get a list of the managed books on a device.
// This command allows the server to query for installed 3rd party applications.
type ManagedMediaListCommand struct{}

func (p *ManagedMediaListCommand) RequestType() string {
	return "ManagedMediaList"
}

type ManagedMediaListCommandResponse struct {
	// An array of dictionaries that describes managed books.
	Books []*BooksItem `json:"Books" plist:"Books" required:"true"`
}

// A dictionary that describes a managed book.
type BooksItem struct {
	// The book's iTunes Store identifier.
	ITunesStoreID int64 `json:"iTunesStoreID" plist:"iTunesStoreID" required:"true"`
	// The installation state of this book, which is one of the following values:
	// - `Queued`
	// - `PromptingForLogin`
	// - `Updating`
	// - `Installing`
	// - `Managed`
	// - `ManagedButUninstalled`
	// - `Installed`
	// - `Uninstalled`
	// - `Failed`
	// - `Unknown`
	// The `Failed` and `Unknown` states are transient and the device only reports them once. Books from the Book Store report their state as `Installed` instead of `Managed`.
	State *string `json:"State,omitempty" plist:"State,omitempty"`
	// The book's persistent identifier in reverse-DNS form; for example, `com.acme.manuals.training`.
	PersistentID *string `json:"PersistentID,omitempty" plist:"PersistentID,omitempty"`
	// The kind of the media, which is one of the following values:
	// - `pdf`: A PDF file
	// - `epub`: An EPUB file in `gzip` format
	// - `ibooks`: An iBooks Author file in `gzip` format
	// - The file extension in the URL
	// This value is available in iOS 8 and later.
	Kind *string `json:"Kind,omitempty" plist:"Kind,omitempty"`
	// The book's version number.
	Version *string `json:"Version,omitempty" plist:"Version,omitempty"`
	// The name of the book's author.
	Author *string `json:"Author,omitempty" plist:"Author,omitempty"`
	// The book's title.
	Title *string `json:"Title,omitempty" plist:"Title,omitempty"`
}

// Remove a previously installed book from a device.
// This command allows an MDM server to remove managed media. This command returns Acknowledged even if the item is not found.
type RemoveMediaCommand struct {
	// The media type, which can only be `Book`.
	MediaType RemoveMediaCommandMediaType `json:"MediaType" plist:"MediaType" required:"true"`
	// The book's iTunes Store identifier.
	ITunesStoreID *string `json:"iTunesStoreID,omitempty" plist:"iTunesStoreID,omitempty"`
	// The book's persistent identifier in reverse-DNS form; for example, `com.acme.manuals.training`.
	PersistentID *string `json:"PersistentID,omitempty" plist:"PersistentID,omitempty"`
}

func (p *RemoveMediaCommand) RequestType() string {
	return "RemoveMedia"
}

type RemoveMediaCommandResponse struct{}

// The media type, which can only be `Book`.
type RemoveMediaCommandMediaType string

const (
	RemoveMediaCommandMediaTypeBook RemoveMediaCommandMediaType = "Book"
)

// Prompt the user to share their screen using AirPlay Mirroring.
// This command prompts the user to share their screen using AirPlay Mirroring.
type RequestMirroringCommand struct {
	// The name of the AirPlay Mirroring destination.
	DestinationName *string `json:"DestinationName,omitempty" plist:"DestinationName,omitempty"`
	// The hardware address of the AirPlay Mirroring destination that identifies the device, in the format `xx:xx:xx:xx:xx`. This value isn't case-sensitive. Not available for Apple TV devices running tvOS 18 or later, use `DestinationName` instead.
	DestinationDeviceID *string `json:"DestinationDeviceID,omitempty" plist:"DestinationDeviceID,omitempty"`
	// The number of seconds, from `10` to `300`, for the device to spend searching for the destination. The default value is `30`.
	ScanTime *int64 `json:"ScanTime,omitempty" plist:"ScanTime,omitempty"`
	// The screen-sharing password that the device uses when connecting to the destination.
	Password *string `json:"Password,omitempty" plist:"Password,omitempty"`
}

func (p *RequestMirroringCommand) RequestType() string {
	return "RequestMirroring"
}

type RequestMirroringCommandResponse struct {
	// The result of the request. One of these values:
	// - `Prompting`: The user is receiving a prompt to share their screen.
	// - `DestinationNotFound`: The device is unable to reach the destination.
	// - `Cancelled`: The user canceled the request.
	// - `Unknown`: An unknown error occurred.
	MirroringResult *string `json:"MirroringResult,omitempty" plist:"MirroringResult,omitempty"`
}

// Stop mirroring the display to another device.
// This command stops AirPlay mirroring.
type StopMirroringCommand struct{}

func (p *StopMirroringCommand) RequestType() string {
	return "StopMirroring"
}

type StopMirroringCommandResponse struct{}

// Remove the passcode from a device.
// This command allows the server to clear the passcode on the device. This command requires the Device Lock and Passcode Removal right.
type ClearPasscodeCommand struct {
	// The unlock token value that the device provides in its `TokenUpdateMessage` check-in message.
	UnlockToken []byte `json:"UnlockToken" plist:"UnlockToken" required:"true"`
}

func (p *ClearPasscodeCommand) RequestType() string {
	return "ClearPasscode"
}

type ClearPasscodeCommandResponse struct{}

// Change or clear the firmware password on a device.
// Changes or clears the firmware password for the device. Requires the "Device lock and passcode removal right". This command is not available on Mac computers with Apple silicon.
type SetFirmwarePasswordCommand struct {
	// The current password, which you must set if the device has a firmware password.
	CurrentPassword *string `json:"CurrentPassword,omitempty" plist:"CurrentPassword,omitempty"`
	// The new firmware password. Set to an empty string to clear the password. The characters in this value must consist of low-ASCII, printable characters (`0x20` through `0x7E`) to ensure that all characters are enterable on the EFI login screen.
	NewPassword string `json:"NewPassword" plist:"NewPassword" required:"true"`
	// If `true`, enable ROMs.
	AllowOroms *bool `json:"AllowOroms,omitempty" plist:"AllowOroms,omitempty"`
}

func (p *SetFirmwarePasswordCommand) RequestType() string {
	return "SetFirmwarePassword"
}

type SetFirmwarePasswordCommandResponse struct {
	// A dictionary containing the results of the command.
	SetFirmwarePassword SetFirmwarePassword `json:"SetFirmwarePassword" plist:"SetFirmwarePassword" required:"true"`
}

// A dictionary containing the results of the command.
type SetFirmwarePassword struct {
	// If `true`, the password change succeeded.
	PasswordChanged bool `json:"PasswordChanged" plist:"PasswordChanged" required:"true"`
}

// Verify the firmware password on a device.
// Verifies the device's firmware password. This command is not available on Mac computers with Apple silicon.
type VerifyFirmwarePasswordCommand struct {
	// The password to verify.
	Password string `json:"Password" plist:"Password" required:"true"`
}

func (p *VerifyFirmwarePasswordCommand) RequestType() string {
	return "VerifyFirmwarePassword"
}

type VerifyFirmwarePasswordCommandResponse struct {
	// A dictionary containing the results of the command.
	VerifyFirmwarePassword VerifyFirmwarePassword `json:"VerifyFirmwarePassword" plist:"VerifyFirmwarePassword" required:"true"`
}

// A dictionary containing the results of the command.
type VerifyFirmwarePassword struct {
	// If 'true', the provided password matched the firmware password set for the device.
	PasswordVerified bool `json:"PasswordVerified" plist:"PasswordVerified" required:"true"`
}

// Set or clear the Recovery Lock password.
// Sets or clears the recovery lock password (Apple Silicon devices only). Requires the "Device lock and passcode removal right".
type SetRecoveryLockCommand struct {
	// If the device has a Recovery Lock password set, the system requires the current password.
	CurrentPassword *string `json:"CurrentPassword,omitempty" plist:"CurrentPassword,omitempty"`
	// The new password for Recovery Lock. Set as an empty string to clear the Recovery Lock password.
	NewPassword string `json:"NewPassword" plist:"NewPassword" required:"true"`
}

func (p *SetRecoveryLockCommand) RequestType() string {
	return "SetRecoveryLock"
}

type SetRecoveryLockCommandResponse struct{}

// Verify the device's Recovery Lock password.
// Verifies the device's recovery lock password. (AppleSilicon devices only)
type VerifyRecoveryLockCommand struct {
	// The password to verify.
	Password string `json:"Password" plist:"Password" required:"true"`
}

func (p *VerifyRecoveryLockCommand) RequestType() string {
	return "VerifyRecoveryLock"
}

type VerifyRecoveryLockCommandResponse struct {
	// If `true`, the device verified the password.
	PasswordVerified bool `json:"PasswordVerified" plist:"PasswordVerified" required:"true"`
}

// Request an unlock token from a device.
// This command requests an UnlockToken from the device. Pass this token to the ClearPasscode command to unlock the device.
type RequestUnlockTokenCommand struct{}

func (p *RequestUnlockTokenCommand) RequestType() string {
	return "RequestUnlockToken"
}

type RequestUnlockTokenCommandResponse struct {
	// The unlock token. Erasing the user partition invalidates this token.
	UnlockToken []byte `json:"UnlockToken" plist:"UnlockToken" required:"true"`
}

// Install a configuration profile on a device.
// This command allows the host to install a configuration profile. The profile may be encrypted using any installed identity certificate. The profile may also be signed. This command requires the Profile Installation and Removal right. It's supported in the user channel.
type InstallProfileCommand struct {
	// The profile to install, which you can encrypt using any identity certificate installed on the device. You can also sign the profile.
	Payload []byte `json:"Payload" plist:"Payload" required:"true"`
}

func (p *InstallProfileCommand) RequestType() string {
	return "InstallProfile"
}

type InstallProfileCommandResponse struct{}

// Get a list of installed profiles on a device.
// This command allows the MDM server to query for the profiles installed on the device. This command requires the Inspect Profile Manifest right. It's supported on the user channel.
type ProfileListCommand struct {
	// If `true`, only include profiles that MDM has installed. For user enrollments, the device ignores this key and always limits the results to managed profiles. This value is available in iOS 13 and later, macOS 10.5 and later, and tvOS 13 and later.
	ManagedOnly *bool `json:"ManagedOnly,omitempty" plist:"ManagedOnly,omitempty"`
}

func (p *ProfileListCommand) RequestType() string {
	return "ProfileList"
}

type ProfileListCommandResponse struct {
	// An array of dictionaries that describes each installed profile.
	ProfileList []*ProfileListItem `json:"ProfileList" plist:"ProfileList" required:"true"`
}

// A dictionary that describes a profile list item.
type ProfileListItem struct {
	// The unique identifier for the profile.
	PayloadUUID string `json:"PayloadUUID" plist:"PayloadUUID" required:"true"`
	// The reverse-DNS-style identifier of the profile; for example, `com.example.myprofile`.
	PayloadIdentifier string `json:"PayloadIdentifier" plist:"PayloadIdentifier" required:"true"`
	// The version of the configuration profile as a whole, not of the individual profiles within it. The value should be `1`.
	PayloadVersion *int64 `json:"PayloadVersion,omitempty" plist:"PayloadVersion,omitempty"`
	// The human-readable name of the profile.
	PayloadDisplayName *string `json:"PayloadDisplayName,omitempty" plist:"PayloadDisplayName,omitempty"`
	// The human-readable name of the organization that provided the profile.
	PayloadOrganization *string `json:"PayloadOrganization,omitempty" plist:"PayloadOrganization,omitempty"`
	// The description of the profile.
	PayloadDescription *string `json:"PayloadDescription,omitempty" plist:"PayloadDescription,omitempty"`
	// If `true`, the user can't delete the profile unless it has a removal password and the user provides it. The framework ignores this field on unsupervised devices.
	PayloadRemovalDisallowed *bool `json:"PayloadRemovalDisallowed,omitempty" plist:"PayloadRemovalDisallowed,omitempty"`
	// If `true`, the profile has a passcode for removal.
	HasRemovalPasscode *bool `json:"HasRemovalPasscode,omitempty" plist:"HasRemovalPasscode,omitempty"`
	// If `true`, it's an encrypted profile.
	IsEncrypted *bool `json:"IsEncrypted,omitempty" plist:"IsEncrypted,omitempty"`
	// An array that contains the certificate for signing the profile, followed by any intermediate certificates, in DER-encoded X.509 format.
	SignerCertificates *[][]byte `json:"SignerCertificates,omitempty" plist:"SignerCertificates,omitempty"`
	// If `true`, the current MDM service installed the profile. MDM doesn't return this value for supervised devices, and can remove or replace all profiles on supervised devices.
	IsManaged *bool `json:"IsManaged,omitempty" plist:"IsManaged,omitempty"`
	// A string set to `Declarative Device Management` when the profile is managed by Declarative Device Management.
	Source *string `json:"Source,omitempty" plist:"Source,omitempty"`
	// An array of payload content items. This value isn't present if `IsEncrypted` is `true`.
	PayloadContent *[]*PayloadContentItem `json:"PayloadContent,omitempty" plist:"PayloadContent,omitempty"`
}

// A dictionary that describes a profile payload content item.
type PayloadContentItem struct {
	// The type of payload, such as `com.apple.wifi.managed`.
	PayloadType string `json:"PayloadType" plist:"PayloadType" required:"true"`
	// The version of the payload. The value is `1`.
	PayloadVersion PayloadVersion `json:"PayloadVersion" plist:"PayloadVersion" required:"true"`
	// The reverse-DNS-style identifier of the payload, such as `com.example.mypayload`.
	PayloadIdentifier string `json:"PayloadIdentifier" plist:"PayloadIdentifier" required:"true"`
	// The unique identifier of the payload.
	PayloadUUID string `json:"PayloadUUID" plist:"PayloadUUID" required:"true"`
	// The human-readable name of the payload.
	PayloadDisplayName *string `json:"PayloadDisplayName,omitempty" plist:"PayloadDisplayName,omitempty"`
	// A description of the payload.
	PayloadDescription *string `json:"PayloadDescription,omitempty" plist:"PayloadDescription,omitempty"`
	// The human-readable name of the organization that provided the payload.
	PayloadOrganization *string `json:"PayloadOrganization,omitempty" plist:"PayloadOrganization,omitempty"`
}

// The version of the payload. The value is `1`.
type PayloadVersion int64

const (
	PayloadVersion1 PayloadVersion = 1
)

// Install a provisioning profile on a device.
// This command allows the server to install a provisioning profile. No error occurs if the provisioning profile is already installed. This command requires the Provisioning Profile Installation and Removal right. On macOS, this command is for iOS and iPadOS style provisioning profiles only.
type InstallProvisioningProfileCommand struct {
	// The provisioning profile.
	ProvisioningProfile []byte `json:"ProvisioningProfile" plist:"ProvisioningProfile" required:"true"`
}

func (p *InstallProvisioningProfileCommand) RequestType() string {
	return "InstallProvisioningProfile"
}

type InstallProvisioningProfileCommandResponse struct{}

// Get a list of installed provisioning profiles on a device.
// This command allows the server to retrieve the list of installed provisioning profiles on the device. This command requires the Inspect Provisioning Profiles right. On macOS, this command is for iOS and iPadOS style provisioning profiles only.
type ProvisioningProfileListCommand struct {
	// If `true`, only include profiles that MDM has installed. For user enrollments, the device ignores this key and always limits the results to managed profiles. This value is available in iOS 13 and later, and tvOS 13 and later.
	ManagedOnly *bool `json:"ManagedOnly,omitempty" plist:"ManagedOnly,omitempty"`
}

func (p *ProvisioningProfileListCommand) RequestType() string {
	return "ProvisioningProfileList"
}

type ProvisioningProfileListCommandResponse struct {
	// An array of dictionaries that describes each installed profile.
	ProvisioningProfileList []*ProvisioningProfileListItem `json:"ProvisioningProfileList" plist:"ProvisioningProfileList" required:"true"`
}

// A dictionary that describes a provisioning profile list item.
type ProvisioningProfileListItem struct {
	// The display name of the provisioning profile.
	Name string `json:"Name" plist:"Name" required:"true"`
	// The unique identifier for the provisioning profile.
	UUID string `json:"UUID" plist:"UUID" required:"true"`
	// The expiry date of the provisioning profile.
	ExpiryDate *time.Time `json:"ExpiryDate,omitempty" plist:"ExpiryDate,omitempty"`
}

// Remove a previously installed provisioning profile from a device.
// This command allows the server to remove a provisioning profile. This command requires the Provisioning Profile Installation and Removal right. On macOS, this command is for iOS and iPadOS style provisioning profiles only.
type RemoveProvisioningProfileCommand struct {
	// The unique identifier of the provisioning profile to remove.
	UUID string `json:"UUID" plist:"UUID" required:"true"`
}

func (p *RemoveProvisioningProfileCommand) RequestType() string {
	return "RemoveProvisioningProfile"
}

type RemoveProvisioningProfileCommandResponse struct{}

// Remove a previously installed profile from the device.
// This command allows the server to remove a profile. This command requires the Profile Installation and Removal Right. It's supported in the user channel.
type RemoveProfileCommand struct {
	// The identifier of the profile to remove.
	Identifier string `json:"Identifier" plist:"Identifier" required:"true"`
}

func (p *RemoveProfileCommand) RequestType() string {
	return "RemoveProfile"
}

type RemoveProfileCommandResponse struct{}

// Disable Remote Desktop on a device.
// Disable Remote Desktop.
type DisableRemoteDesktopCommand struct{}

func (p *DisableRemoteDesktopCommand) RequestType() string {
	return "DisableRemoteDesktop"
}

type DisableRemoteDesktopCommandResponse struct{}

// Enable Remote Desktop on a device.
// Enable Remote Desktop.
type EnableRemoteDesktopCommand struct{}

func (p *EnableRemoteDesktopCommand) RequestType() string {
	return "EnableRemoteDesktop"
}

type EnableRemoteDesktopCommandResponse struct{}

// Change the FileVault primary password on a device.
// This command allows for changing a device's FileVaultMaster password.
type RotateFileVaultKeyCommand struct {
	// The type of FileVault key you want to change the password for. Set this value to `personal` and set a value for `Password` in the `FileVaultUnlock` dictionary to enable unlocking a device with a password. Set this value to `institutional` and set values for `PrivateKeyExport` and `PrivateKeyExportPassword` in the `FileVaultUnlock` dictionary.
	KeyType KeyType `json:"KeyType" plist:"KeyType" required:"true"`
	// A dictionary that contains FileVault unlock options.
	FileVaultUnlock FileVaultUnlock `json:"FileVaultUnlock" plist:"FileVaultUnlock" required:"true"`
	// A DER-encoded certificate for creating a new institutional recovery key, which the system requires if `KeyType` is `institutional`.
	NewCertificate *[]byte `json:"NewCertificate,omitempty" plist:"NewCertificate,omitempty"`
	// A DER-encoded certificate for encrypting the new personal recovery key in a wrapper conforming to the IETF Cryptographic Message Syntax (CMS) standard.
	ReplyEncryptionCertificate *[]byte `json:"ReplyEncryptionCertificate,omitempty" plist:"ReplyEncryptionCertificate,omitempty"`
}

func (p *RotateFileVaultKeyCommand) RequestType() string {
	return "RotateFileVaultKey"
}

type RotateFileVaultKeyCommandResponse struct {
	// The result of rotating the personal recovery key.
	RotateResult *RotateResult `json:"RotateResult,omitempty" plist:"RotateResult,omitempty"`
}

// The type of FileVault key you want to change the password for. Set this value to `personal` and set a value for `Password` in the `FileVaultUnlock` dictionary to enable unlocking a device with a password. Set this value to `institutional` and set values for `PrivateKeyExport` and `PrivateKeyExportPassword` in the `FileVaultUnlock` dictionary.
type KeyType string

const (
	KeyTypePersonal      KeyType = "personal"
	KeyTypeInstitutional KeyType = "institutional"
)

// A dictionary that contains FileVault unlock options.
type FileVaultUnlock struct {
	// A FileVault user's password, or if using a CoreStorage volume, the personal recovery key.
	Password *string `json:"Password,omitempty" plist:"Password,omitempty"`
	// The data for a .p12 export of the private key for the current institutional recovery key, which requires that `KeyType` is `institutional`. The system ignores this key on APFS volumes.
	PrivateKeyExport *[]byte `json:"PrivateKeyExport,omitempty" plist:"PrivateKeyExport,omitempty"`
	// The password for `PrivateKeyExport`. Either `Password` or both `PrivateKeyExport` and `PrivateKeyExportPassword` must be present. The system ignores this key on APFS volumes.
	PrivateKeyExportPassword *string `json:"PrivateKeyExportPassword,omitempty" plist:"PrivateKeyExportPassword,omitempty"`
}

// The result of rotating the personal recovery key.
type RotateResult struct {
	// A new personal recovery key that is encrypted using a `ReplyEncryptionCertificate` as a CMS-compliant envelope.
	EncryptedNewRecoveryKey *[]byte `json:"EncryptedNewRecoveryKey,omitempty" plist:"EncryptedNewRecoveryKey,omitempty"`
}

// Update the local administrator account password.
// Allows changing the password of a local admin account that was created by Setup Assistant during DEP enrollment via the AccountConfiguration command.
type SetAutoAdminPasswordCommand struct {
	// The unique identifier of the local administrator account. If this value doesn't match the GUID of an administrator account that MDM created during Device Enrollment Program (DEP) enrollment, the command returns an error.
	GUID string `json:"GUID" plist:"GUID" required:"true"`
	// The precreated salted SHA-512 PBKDF2 password hash for the account.
	// Create this hash on the server using the CommonCrypto libraries, or equivalent, as a salted SHA-512 PBKDF2 dictionary that contains these elements:
	// - `entropy`: The derived key from the password hash; for example, from `CCKeyDerivationPBKDF()`
	// - `salt`: The 32-byte randomized salt; for example, from `CCRandomCopyBytes()`
	// - `iterations:` The number of iterations; for example, from `CCCalibratePBKDF()` using a minimum hash time of 100 milliseconds, or if unknown, a number in the range of 20,000 to 40,000 iterations
	// Place the dictionary that contains these elements into an outer dictionary with the key `SALTED-SHA512-PBKDF2`. Convert this dictionary to binary data before setting it as the value for `passwordHash`.
	PasswordHash []byte `json:"passwordHash" plist:"passwordHash" required:"true"`
}

func (p *SetAutoAdminPasswordCommand) RequestType() string {
	return "SetAutoAdminPassword"
}

type SetAutoAdminPasswordCommandResponse struct{}

// Configure settings on a device.
// This command allows the server to set settings on the device. These settings take effect on a one-time basis. The user may still be able to change the settings at a later time. This command requires the ApplySettings right.
type SettingsCommand struct {
	// An array of dictionaries that contains the settings.
	Settings []*Wallpaper `json:"Settings" plist:"Settings" required:"true"`
}

func (p *SettingsCommand) RequestType() string {
	return "Settings"
}

type SettingsCommandResponse struct {
	// A dictionary that describes the results of configuring settings.
	Settings *Settings `json:"Settings,omitempty" plist:"Settings,omitempty"`
}

// A dictionary that contains wallpaper settings. This setting doesn't support user enrollment. Available in iOS 8 and later. Starting in iOS 16 and iPadOS 17, when setting the wallpaper for the first time, both locations update. After that, you can set either location separately.
type Wallpaper struct {
	// A string that identifies this setting.
	Item WallpaperItem `json:"Item" plist:"Item" required:"true"`
	// A Base64-encoded image in either PNG or JPG format to use for wallpaper.
	Image []byte `json:"Image" plist:"Image" required:"true"`
	// A number that indicates where to use the wallpaper, which is one of the following values:
	// - `1`: Lock Screen
	// - `2`: Home Screen
	// - `3`: Both Lock and Home Screens.
	// In iOS 16 and later, and iPadOS 17 and later, when you set the wallpaper for the first time, the system sets both the Lock Screen and Home Screen. After that, you can separately set each location.
	Where Where `json:"Where" plist:"Where" required:"true"`
}

// A string that identifies this setting.
type WallpaperItem string

const (
	WallpaperItemWallpaper WallpaperItem = "Wallpaper"
)

// A number that indicates where to use the wallpaper, which is one of the following values:
// - `1`: Lock Screen
// - `2`: Home Screen
// - `3`: Both Lock and Home Screens.
// In iOS 16 and later, and iPadOS 17 and later, when you set the wallpaper for the first time, the system sets both the Lock Screen and Home Screen. After that, you can separately set each location.
type Where int64

const (
	Where1 Where = 1
	Where2 Where = 2
	Where3 Where = 3
)

// A dictionary that contains data roaming settings. This setting requires the Network Information access right, and doesn't support user enrollment. Available in iOS 5 and later.
type DataRoaming struct {
	// A string that identifies this setting.
	Item DataRoamingItem `json:"Item" plist:"Item" required:"true"`
	// If `true`, enable data roaming, which also enables voice roaming. If `false`, disable data roaming.
	Enabled bool `json:"Enabled" plist:"Enabled" required:"true"`
}

// A string that identifies this setting.
type DataRoamingItem string

const (
	DataRoamingItemDataRoaming DataRoamingItem = "DataRoaming"
)

// A dictionary that contains voice roaming settings. This setting requires the Network Information access right, and doesn't support user enrollment. Available in iOS 5 and later.
type VoiceRoaming struct {
	// A string that identifies this setting.
	Item VoiceRoamingItem `json:"Item" plist:"Item" required:"true"`
	// If `true`, enable voice roaming. If `false`, disable voice roaming, which also disables data roaming. The setting is only available for certain carriers.
	Enabled bool `json:"Enabled" plist:"Enabled" required:"true"`
}

// A string that identifies this setting.
type VoiceRoamingItem string

const (
	VoiceRoamingItemVoiceRoaming VoiceRoamingItem = "VoiceRoaming"
)

// A dictionary that contains Personal Hotspot settings. This setting requires the Network Information access right, and doesn't support user enrollment. Available in iOS 5 and later.
type PersonalHotspot struct {
	// A string that identifies this setting.
	Item PersonalHotspotItem `json:"Item" plist:"Item" required:"true"`
	// If `true`, enable Personal Hotspot. If `false`, disable Personal Hotspot.
	Enabled bool `json:"Enabled" plist:"Enabled" required:"true"`
}

// A string that identifies this setting.
type PersonalHotspotItem string

const (
	PersonalHotspotItemPersonalHotspot PersonalHotspotItem = "PersonalHotspot"
)

// A dictionary that contains Bluetooth settings. This setting requires the Network Information access right, doesn't support user enrollment, and is available only on supervised devices. Available in iOS 11.3 and later, and macOS 10.13.4 and later.
type Bluetooth struct {
	// A string that identifies this setting.
	Item BluetoothItem `json:"Item" plist:"Item" required:"true"`
	// If `true`, enable the Bluetooth setting. If `false`, disable the Bluetooth setting.
	Enabled bool `json:"Enabled" plist:"Enabled" required:"true"`
}

// A string that identifies this setting.
type BluetoothItem string

const (
	BluetoothItemBluetooth BluetoothItem = "Bluetooth"
)

// A dictionary that contains the configurations to apply to the app. Omit this setting to remove existing configurations. This setting requires the App Management access right, supports user enrollment, and is available in iOS 7 and later, macOS 10.15 and later, and tvOS 10.2 and later. This setting fails for apps that Declarative Device Management manages.
type ApplicationConfiguration struct {
	// A string that identifies this setting.
	Item ApplicationConfigurationItem `json:"Item" plist:"Item" required:"true"`
	// The bundle identifier of the managed app.
	// > Important:
	// > For a watchOS app, the identifier needs to be the watch's bundle identifier, which differs from the main bundle identifier for the iPhone to which the watch is paired. Obtain the watch's bundle identifier for an app with a watch bundle, in the `watchBundleId` key that's part of the Content Metadata query. For more information on this query, see `Getting App and Book Information`.
	Identifier string `json:"Identifier" plist:"Identifier" required:"true"`
	// A dictionary that contains the configurations to apply to the app. Omit this setting to remove existing configurations.
	Configuration *map[string]any `json:"Configuration,omitempty" plist:"Configuration,omitempty"`
}

// A string that identifies this setting.
type ApplicationConfigurationItem string

const (
	ApplicationConfigurationItemApplicationConfiguration ApplicationConfigurationItem = "ApplicationConfiguration"
)

// A dictionary that contains the attributes to apply to the app. Omit this setting to remove existing attributes. This setting supports user enrollment, is available in iOS 7 and later, and tvOS 10.2 and later. This setting fails for apps that Declarative Device Management manages.
type ApplicationAttributes struct {
	// A string that identifies this setting.
	Item SettingsApplicationAttributesItem `json:"Item" plist:"Item" required:"true"`
	// The bundle identifier of the app.
	// > Important:
	// > For a watchOS app, the identifier needs to be the watch's bundle identifier, which differs from the main bundle identifier for the iPhone to which the watch is paired. Obtain the watch's bundle identifier for an app with a watch bundle, in the `watchBundleId` key that's part of the Content Metadata query. For more information on this query, see `Getting App and Book Information`.
	Identifier string `json:"Identifier" plist:"Identifier" required:"true"`
	// A dictionary that contains the attributes to apply to the app. Omit this setting to remove existing attributes. This setting is available in iOS 7 and later, and tvOS 10.2 and later.
	Attributes *ApplicationAttributesAttributes `json:"Attributes,omitempty" plist:"Attributes,omitempty"`
}

// A string that identifies this setting.
type SettingsApplicationAttributesItem string

const (
	SettingsApplicationAttributesItemApplicationAttributes SettingsApplicationAttributesItem = "ApplicationAttributes"
)

// A dictionary that contains the attributes to apply to the app. Omit this setting to remove existing attributes. This setting is available in iOS 7 and later, and tvOS 10.2 and later.
type ApplicationAttributesAttributes struct {
	// A per-app VPN unique identifier for this app. Available in iOS 7 and later.
	VPNUUID *string `json:"VPNUUID,omitempty" plist:"VPNUUID,omitempty"`
	// The content filter UUID for this app. Available in iOS 16 and later.
	ContentFilterUUID *string `json:"ContentFilterUUID,omitempty" plist:"ContentFilterUUID,omitempty"`
	// The DNS proxy UUID for this app. Available in iOS 16 and later.
	DNSProxyUUID *string `json:"DNSProxyUUID,omitempty" plist:"DNSProxyUUID,omitempty"`
	// The relay UUID for this app. Available in iOS 17 and later.
	RelayUUID *string `json:"RelayUUID,omitempty" plist:"RelayUUID,omitempty"`
	// An array that contains the associated domains to add to this app. Available in iOS 13 and later.
	AssociatedDomains *[]string `json:"AssociatedDomains,omitempty" plist:"AssociatedDomains,omitempty"`
	// If `true`, perform claimed site association verification directly at the domain, instead of on Apple's servers. Only set this to `true` for domains that can't access the internet. Available in iOS 14 and later.
	AssociatedDomainsEnableDirectDownloads *bool `json:"AssociatedDomainsEnableDirectDownloads,omitempty" plist:"AssociatedDomainsEnableDirectDownloads,omitempty"`
	// If `false`, this app isn't removable while it's managed. Available in iOS 14 and later, and tvOS 14 and later.
	Removable *bool `json:"Removable,omitempty" plist:"Removable,omitempty"`
	// If true, the system require Tap to Pay on iPhone users to use Face ID or a passcode to unlock their device after every transaction that requires a customer's card PIN. If `false`, the user can configure this setting on their device.
	// Available in iOS 16.4 and later.
	TapToPayScreenLock *bool `json:"TapToPayScreenLock,omitempty" plist:"TapToPayScreenLock,omitempty"`
	// The data network name (DNN) or app category. For DNN, the value is `DNN:name`, where `name` is the carrier-provided DNN name. For app category, the value is `AppCategory:category`, where `category` is a carrier-provided string like "Enterprise1"`.`
	// Available in iOS 17 and later.
	CellularSliceUUID *string `json:"CellularSliceUUID,omitempty" plist:"CellularSliceUUID,omitempty"`
	// If `false`, the system prevents the user from hiding the app. It doesn't affect the user's ability to leave it in the App Library, while removing it from the Home Screen.
	Hideable *bool `json:"Hideable,omitempty" plist:"Hideable,omitempty"`
	// If `false`, the system prevents the user from locking the app. This also prevents the user from hiding the app.
	Lockable *bool `json:"Lockable,omitempty" plist:"Lockable,omitempty"`
}

// A dictionary that contains device name settings. This setting doesn't support user enrollment, and is available only on supervised devices. Available in iOS 5 and later, macOS 10.10 and later, and visionOS 2 and later.
type DeviceName struct {
	// A string that identifies this setting.
	Item DeviceNameItem `json:"Item" plist:"Item" required:"true"`
	// The device's name.
	DeviceName string `json:"DeviceName" plist:"DeviceName" required:"true"`
}

// A string that identifies this setting.
type DeviceNameItem string

const (
	DeviceNameItemDeviceName DeviceNameItem = "DeviceName"
)

// A dictionary that contains hostname settings. This setting doesn't support user enrollment, and is available in macOS 10.11 and later.
type HostName struct {
	// The string that defines this setting type.
	Item HostNameItem `json:"Item" plist:"Item" required:"true"`
	// The hostname for the device.
	HostName string `json:"HostName" plist:"HostName" required:"true"`
}

// The string that defines this setting type.
type HostNameItem string

const (
	HostNameItemHostName HostNameItem = "HostName"
)

// A dictionary that contains settings about the organization operating the MDM server. This setting supports user enrollment. Available in iOS 5 and later.
type SettingsOrganizationInfo struct {
	// The string that defines this setting type.
	Item OrganizationInfoItem `json:"Item" plist:"Item" required:"true"`
	// A dictionary that contains information about the organization operating the MDM server. Omit this setting to remove existing information.
	OrganizationInfo *OrganizationInfoOrganizationInfo `json:"OrganizationInfo,omitempty" plist:"OrganizationInfo,omitempty"`
}

// The string that defines this setting type.
type OrganizationInfoItem string

const (
	OrganizationInfoItemOrganizationInfo OrganizationInfoItem = "OrganizationInfo"
)

// A dictionary that contains information about the organization operating the MDM server. Omit this setting to remove existing information.
type OrganizationInfoOrganizationInfo struct {
	// A string that describes the organization operating the MDM server for display to the user during certain operations, such as purchasing or installing apps.
	OrganizationName string `json:"OrganizationName" plist:"OrganizationName" required:"true"`
	// A shorter version of `OrganizationName`, preferably a single word or abbreviation, suitable for display to the user in places where a very short name is necessary.
	OrganizationShortName *string `json:"OrganizationShortName,omitempty" plist:"OrganizationShortName,omitempty"`
	// The organization's address. Use the LF character (`&#10`) to insert line breaks.
	OrganizationAddress *string `json:"OrganizationAddress,omitempty" plist:"OrganizationAddress,omitempty"`
	// The organization's phone number.
	OrganizationPhone *string `json:"OrganizationPhone,omitempty" plist:"OrganizationPhone,omitempty"`
	// The organization's support email address.
	OrganizationEmail *string `json:"OrganizationEmail,omitempty" plist:"OrganizationEmail,omitempty"`
	// A unique identifier for the various services a single organization manages.
	OrganizationMagic *string `json:"OrganizationMagic,omitempty" plist:"OrganizationMagic,omitempty"`
}

// A dictionary that contains default application bundle identifiers for each default application type that can be set.
type DefaultApplications struct {
	// The bundle identifier of the app the system sets as the default web browser. This app must be an eligible web browser for the region of the device.
	WebBrowser *string `json:"WebBrowser,omitempty" plist:"WebBrowser,omitempty"`
}

// A dictionary that contains settings related to the MDM protocol. This setting doesn't support user enrollment. Available in iOS 7 and later, macOS 10.15 and later, and visionOS 2 and later.
type SettingsMDMOptions struct {
	// The string that defines this setting type.
	Item MDMOptionsItem `json:"Item" plist:"Item" required:"true"`
	// A dictionary of MDM options.
	MDMOptions MDMOptionsMDMOptions `json:"MDMOptions" plist:"MDMOptions" required:"true"`
}

// The string that defines this setting type.
type MDMOptionsItem string

const (
	MDMOptionsItemMDMOptions MDMOptionsItem = "MDMOptions"
)

// A dictionary of MDM options.
type MDMOptionsMDMOptions struct {
	// If `true`, a supervised device registers itself with Activation Lock when the user enables Find My. This setting is available for supervised devices in iOS 7 and later, and macOS 10.15 and later.
	ActivationLockAllowedWhileSupervised *bool `json:"ActivationLockAllowedWhileSupervised,omitempty" plist:"ActivationLockAllowedWhileSupervised,omitempty"`
	// If `true`, the server supports the Bootstrap Token commands.
	BootstrapTokenAllowed *bool `json:"BootstrapTokenAllowed,omitempty" plist:"BootstrapTokenAllowed,omitempty"`
	// If `true`, warn the user that they need to reboot into RecoveryOS and allow the MDM server to use the Bootstrap Token for authentication for certain sensitive operations; for example, enabling kernel extensions or installing certain types of software updates. Set this value to `false` if your MDM server doesn't need to perform these operations. The value provided here overrides the value specified in MDM, and only applies when `BootstrapTokenAllowedForAuthentication` is `true` in the `SecurityInfo` response. This value is available for Mac computers with Apple silicon in macOS 11 and later.
	PromptUserToAllowBootstrapTokenForAuthentication *bool `json:"PromptUserToAllowBootstrapTokenForAuthentication,omitempty" plist:"PromptUserToAllowBootstrapTokenForAuthentication,omitempty"`
	// If `true`, the device automatically reboots while locked after several days of inactivity. This is set to `false` by default when a supervised device enrolls.
	IdleRebootAllowed *bool `json:"IdleRebootAllowed,omitempty" plist:"IdleRebootAllowed,omitempty"`
}

// A dictionary that contains settings for maximum resident users. Apple deprecated this setting in iOS 13.4. Use 'SharedDeviceConfiguration` instead. This setting is available only for Shared iPad.
type MaximumResidentUsers struct {
	// A string that identifies this setting.
	Item MaximumResidentUsersItem `json:"Item" plist:"Item" required:"true"`
	// The maximum number of users that can use the device. If this value is greater than the value for the maximum possible number of users that the device supports, the MDM server uses that value instead.
	// This setting requires that the device is in the `AwaitingConfiguration` phase before it receives the [DeviceConfigured](https://developer.apple.com/library/archive/documentation/Miscellaneous/Reference/MobileDeviceManagementProtocolRef/3-MDM_Protocol/MDM_Protocol.html#//apple_ref/doc/uid/TP40017387-CH3-SW301) message.
	// When a device reaches the maximum number of resident users and a new user tries to sign in, the MDM server removes a synchronized user to make space for the new user. If there are no synchronized users, the new user sign-in fails. A synchronized user is a user that has completed syncing their data.
	MaximumResidentUsers int64 `json:"MaximumResidentUsers" plist:"MaximumResidentUsers" required:"true"`
}

// A string that identifies this setting.
type MaximumResidentUsersItem string

const (
	MaximumResidentUsersItemMaximumResidentUsers MaximumResidentUsersItem = "MaximumResidentUsers"
)

// A dictionary that contains shared device configuration settings. This setting is available only for Shared iPad in iOS 13.4 and later.
type SharedDeviceConfiguration struct {
	// A string that identifies this setting.
	Item SharedDeviceConfigurationItem `json:"Item" plist:"Item" required:"true"`
	// The quota size, in megabytes (MB), for each user on the shared device, or if the quota size is too small, the minimum quota size. Available to Temporary Sessions Only guest users on iOS 17+.
	QuotaSize *int64 `json:"QuotaSize,omitempty" plist:"QuotaSize,omitempty"`
	// The expected number of users. If this value is greater than the value for the maximum possible number of users that the device supports, the MDM server uses that value instead.
	ResidentUsers *int64 `json:"ResidentUsers,omitempty" plist:"ResidentUsers,omitempty"`
	// The timeout, in seconds, for the user session. The user session logs out automatically after the specified period of inactivity. The minimum value is 30 seconds. Setting this value to `0` removes the timeout.
	// Available in iOS 14.5 and later.
	UserSessionTimeout *int64 `json:"UserSessionTimeout,omitempty" plist:"UserSessionTimeout,omitempty"`
	// The timeout, in seconds, for the temporary session. The temporary session logs out automatically after the specified period of inactivity. The minimum value is 30 seconds. Setting this value to `0` removes the timeout.
	// Available in iOS 14.5 and later.
	TemporarySessionTimeout *int64 `json:"TemporarySessionTimeout,omitempty" plist:"TemporarySessionTimeout,omitempty"`
	// If `true`, the user only sees the Guest Welcome pane and can only log in as a guest user.
	// If `false`, the user can sign in with a Managed Apple Account (the existing behavior).
	// Available in iOS 14.5 and later.
	TemporarySessionOnly *bool `json:"TemporarySessionOnly,omitempty" plist:"TemporarySessionOnly,omitempty"`
	// A list of domains that the Shared iPad login screen displays. The user can pick a domain from the list to complete their Managed Apple Account.
	// If this list contains more than 3 domains, the system picks 3 at random for display. Available in iOS 16 and later.
	ManagedAppleIDDefaultDomains *[]string `json:"ManagedAppleIDDefaultDomains,omitempty" plist:"ManagedAppleIDDefaultDomains,omitempty"`
	// A grace period (in days) for Shared iPad online authentication. The Shared iPad only verifies the user's passcode locally during login for users that already exist on the device. However, the system requires an online authentication (against Apple's identity server) after the number of days specified by this setting.
	// Setting this value to 0 enforces online authentication every time.
	// Available in iOS 16 and later.
	OnlineAuthenticationGracePeriod *int64 `json:"OnlineAuthenticationGracePeriod,omitempty" plist:"OnlineAuthenticationGracePeriod,omitempty"`
	// If `true`, the system picks the system language and locale automatically for the new Shared iPad user.
	// Available in iOS 16.2 and later.
	SkipLanguageAndLocaleSetupForNewUsers *bool `json:"SkipLanguageAndLocaleSetupForNewUsers,omitempty" plist:"SkipLanguageAndLocaleSetupForNewUsers,omitempty"`
	// If enabled, the Shared iPad device enters Setup Assistant after the user triggers a login. The MDM server has a chance to configure the device and user. After configuration, the server needs to send a `User-Configured-Command` command to the user channel to unblock the login. This feature requires the device to have network access during the login process.
	// Available in iOS 17 and later.
	AwaitUserConfiguration *AwaitUserConfiguration `json:"AwaitUserConfiguration,omitempty" plist:"AwaitUserConfiguration,omitempty"`
	// A dictionary that contains passcode policies.
	PasscodePolicy *PasscodePolicy `json:"PasscodePolicy,omitempty" plist:"PasscodePolicy,omitempty"`
}

// A string that identifies this setting.
type SharedDeviceConfigurationItem string

const (
	SharedDeviceConfigurationItemSharedDeviceConfiguration SharedDeviceConfigurationItem = "SharedDeviceConfiguration"
)

// If enabled, the Shared iPad device enters Setup Assistant after the user triggers a login. The MDM server has a chance to configure the device and user. After configuration, the server needs to send a `User-Configured-Command` command to the user channel to unblock the login. This feature requires the device to have network access during the login process.
// Available in iOS 17 and later.
type AwaitUserConfiguration struct {
	// If `true`, the device stops at the Setup Assistant pane after user login. The user can't use the device until it receives a `User-Configured-Command` command.
	Enabled bool `json:"Enabled" plist:"Enabled" required:"true"`
}

// A dictionary that contains passcode policies.
type PasscodePolicy struct {
	// The number of seconds before a locked screen requires the user to enter the device passcode to unlock it. The minimum value is `0` seconds and the maximum value is `14400` seconds.
	// If a device has a passcode, a change to a larger value doesn't take effect until the user logs out or removes the passcode. For this reason, it's better to set this value before the user sets a passcode.
	// If the value is less than one of the known values, the device uses the next lowest value. For example a value of 299 results in an effective setting of 60.
	// This setting won't take effect if `TemporarySessionOnly` is `true` because there's no passcode for a temporary session.
	PasscodeLockGracePeriod *PasscodePolicyPasscodeLockGracePeriod `json:"PasscodeLockGracePeriod,omitempty" plist:"PasscodeLockGracePeriod,omitempty"`
	// The number of seconds before a device goes to sleep after being idle. The minimum value for this setting is `120` seconds.
	AutoLockTime *int64 `json:"AutoLockTime,omitempty" plist:"AutoLockTime,omitempty"`
}

// The number of seconds before a locked screen requires the user to enter the device passcode to unlock it. The minimum value is `0` seconds and the maximum value is `14400` seconds.
// If a device has a passcode, a change to a larger value doesn't take effect until the user logs out or removes the passcode. For this reason, it's better to set this value before the user sets a passcode.
// If the value is less than one of the known values, the device uses the next lowest value. For example a value of 299 results in an effective setting of 60.
// This setting won't take effect if `TemporarySessionOnly` is `true` because there's no passcode for a temporary session.
type PasscodePolicyPasscodeLockGracePeriod int64

const (
	PasscodePolicyPasscodeLockGracePeriod0     PasscodePolicyPasscodeLockGracePeriod = 0
	PasscodePolicyPasscodeLockGracePeriod60    PasscodePolicyPasscodeLockGracePeriod = 60
	PasscodePolicyPasscodeLockGracePeriod300   PasscodePolicyPasscodeLockGracePeriod = 300
	PasscodePolicyPasscodeLockGracePeriod900   PasscodePolicyPasscodeLockGracePeriod = 900
	PasscodePolicyPasscodeLockGracePeriod3600  PasscodePolicyPasscodeLockGracePeriod = 3600
	PasscodePolicyPasscodeLockGracePeriod14400 PasscodePolicyPasscodeLockGracePeriod = 14400
)

// A dictionary that contains diagnostic submission settings. This setting is available only for Shared iPad in iOS 9.3 and later.
type DiagnosticSubmission struct {
	// The string that defines this setting type.
	Item DiagnosticSubmissionItem `json:"Item" plist:"Item" required:"true"`
	// If `true`, enables diagnostic submission. If `false`, disables diagnostic submission.
	Enabled bool `json:"Enabled" plist:"Enabled" required:"true"`
}

// The string that defines this setting type.
type DiagnosticSubmissionItem string

const (
	DiagnosticSubmissionItemDiagnosticSubmission DiagnosticSubmissionItem = "DiagnosticSubmission"
)

// A dictionary that contains settings for sharing app analytics. This setting is available only for Shared iPad in iOS 9.3.2 and later.
type AppAnalytics struct {
	// A string that identifies this setting.
	Item AppAnalyticsItem `json:"Item" plist:"Item" required:"true"`
	// If `true`, enable sharing app analytics with app developers. If `false`, disable sharing app analytics.
	Enabled bool `json:"Enabled" plist:"Enabled" required:"true"`
}

// A string that identifies this setting.
type AppAnalyticsItem string

const (
	AppAnalyticsItemAppAnalytics AppAnalyticsItem = "AppAnalytics"
)

// A dictionary that contains password lock grace period settings. This setting is available only for Shared iPad in iOS 9.3.2 and later.
// This key is deprecated. Use 'PasscodeLockGracePeriod' in SettingsCommand.Command.Settings.SharedDeviceConfiguration.PasscodePolicy instead.
type SettingsPasscodeLockGracePeriod struct {
	// A string that identifies this setting.
	Item PasscodeLockGracePeriodItem `json:"Item" plist:"Item" required:"true"`
	// The number of seconds before a locked screen requires the user to enter the device passcode to unlock it. The minimum value is `0` seconds and the maximum value is `14400` seconds.
	// If a device has a passcode, a change to a larger value doesn't take effect until the user logs out or removes the passcode. For this reason, it's better to set this value before the user sets a passcode.
	// If the value is less than one of the known values, the device uses the next lowest value. For example a value of 299 results in an effective setting of 60.
	// This setting won't take effect if `TemporarySessionOnly` is `true` because there's no passcode for a temporary session.
	PasscodeLockGracePeriod PasscodeLockGracePeriodPasscodeLockGracePeriod `json:"PasscodeLockGracePeriod" plist:"PasscodeLockGracePeriod" required:"true"`
}

// A string that identifies this setting.
type PasscodeLockGracePeriodItem string

const (
	PasscodeLockGracePeriodItemPasscodeLockGracePeriod PasscodeLockGracePeriodItem = "PasscodeLockGracePeriod"
)

// The number of seconds before a locked screen requires the user to enter the device passcode to unlock it. The minimum value is `0` seconds and the maximum value is `14400` seconds.
// If a device has a passcode, a change to a larger value doesn't take effect until the user logs out or removes the passcode. For this reason, it's better to set this value before the user sets a passcode.
// If the value is less than one of the known values, the device uses the next lowest value. For example a value of 299 results in an effective setting of 60.
// This setting won't take effect if `TemporarySessionOnly` is `true` because there's no passcode for a temporary session.
type PasscodeLockGracePeriodPasscodeLockGracePeriod int64

const (
	PasscodeLockGracePeriodPasscodeLockGracePeriod0     PasscodeLockGracePeriodPasscodeLockGracePeriod = 0
	PasscodeLockGracePeriodPasscodeLockGracePeriod60    PasscodeLockGracePeriodPasscodeLockGracePeriod = 60
	PasscodeLockGracePeriodPasscodeLockGracePeriod300   PasscodeLockGracePeriodPasscodeLockGracePeriod = 300
	PasscodeLockGracePeriodPasscodeLockGracePeriod900   PasscodeLockGracePeriodPasscodeLockGracePeriod = 900
	PasscodeLockGracePeriodPasscodeLockGracePeriod3600  PasscodeLockGracePeriodPasscodeLockGracePeriod = 3600
	PasscodeLockGracePeriodPasscodeLockGracePeriod14400 PasscodeLockGracePeriodPasscodeLockGracePeriod = 14400
)

// A dictionary that contains time zone settings. This setting is available only on supervised devices and doesn't support user enrollment. Available in iOS 14 and later, tvOS 14 and later, and visionOS 2 and later.
type TimeZone struct {
	// A string that identifies this setting.
	Item TimeZoneItem `json:"Item" plist:"Item" required:"true"`
	// The Internet Assigned Numbers Authority (IANA) time zone database name.
	// If the `forceAutomaticDateAndTime` restriction is set in `Restrictions`, this setting fails with an error. Otherwise, setting this value disables automatic time zone logic. The user is still be able to change the time zone; for example, by turning automatic date and time back on. The intention is to allow setting the time zone when automatic determination isn't be available, such as when Location Services are off.
	TimeZone string `json:"TimeZone" plist:"TimeZone" required:"true"`
}

// A string that identifies this setting.
type TimeZoneItem string

const (
	TimeZoneItemTimeZone TimeZoneItem = "TimeZone"
)

// A dictionary that contains software update settings. This setting doesn't support user enrollment. Available in iOS 14.5 and later.
type SettingsSoftwareUpdateSettings struct {
	// A string that represents the type of updates that should appear in the Software Update pane in Settings. Supervised only.
	Item SoftwareUpdateSettingsItem `json:"Item" plist:"Item" required:"true"`
	// This value defines how the system presents software updates to the user. When there's more than one available update for the user, the system behaves as follows:
	// - `0`: Presents both options to the user.
	// - `1`: Presents the lower numbered (oldest) software update version.
	// - `2`: Presents only the highest numbered (most recent) release available for the device.
	// This value has no effect when there's only one available update; the system shows the single available update to the user regardless of the value of this setting.
	// Available in iOS 14.5 and later.
	RecommendationCadence RecommendationCadence `json:"RecommendationCadence" plist:"RecommendationCadence" required:"true"`
}

// A string that represents the type of updates that should appear in the Software Update pane in Settings. Supervised only.
type SoftwareUpdateSettingsItem string

const (
	SoftwareUpdateSettingsItemSoftwareUpdateSettings SoftwareUpdateSettingsItem = "SoftwareUpdateSettings"
)

// This value defines how the system presents software updates to the user. When there's more than one available update for the user, the system behaves as follows:
// - `0`: Presents both options to the user.
// - `1`: Presents the lower numbered (oldest) software update version.
// - `2`: Presents only the highest numbered (most recent) release available for the device.
// This value has no effect when there's only one available update; the system shows the single available update to the user regardless of the value of this setting.
// Available in iOS 14.5 and later.
type RecommendationCadence int64

const (
	RecommendationCadence0 RecommendationCadence = 0
	RecommendationCadence1 RecommendationCadence = 1
	RecommendationCadence2 RecommendationCadence = 2
)

// A dictionary that contains accessibility settings. Available in iOS 16 and later.
type SettingsAccessibilitySettings struct {
	// Sets various accessibility settings. The system allows only keys with explicitly provided values.
	Item AccessibilitySettingsItem `json:"Item" plist:"Item" required:"true"`
	// If `true`, the system enables bold text.
	BoldTextEnabled *bool `json:"BoldTextEnabled,omitempty" plist:"BoldTextEnabled,omitempty"`
	// If `true`, the system enables increase contrast.
	IncreaseContrastEnabled *bool `json:"IncreaseContrastEnabled,omitempty" plist:"IncreaseContrastEnabled,omitempty"`
	// If `true`, the system enables reduced motion.
	ReduceMotionEnabled *bool `json:"ReduceMotionEnabled,omitempty" plist:"ReduceMotionEnabled,omitempty"`
	// If `true`, the system enables reduced transparency.
	ReduceTransparencyEnabled *bool `json:"ReduceTransparencyEnabled,omitempty" plist:"ReduceTransparencyEnabled,omitempty"`
	// The accessibility text size apps that support dynamic text use. `0` is the smallest value, and `11` is the largest available.
	TextSize *SettingsAccessibilitySettingsTextSize `default:"4" json:"TextSize,omitempty" plist:"TextSize,omitempty"`
	// If `true`, the system enables touch accommodations.
	TouchAccommodationsEnabled *bool `json:"TouchAccommodationsEnabled,omitempty" plist:"TouchAccommodationsEnabled,omitempty"`
	// If `true`, the system enables voiceover.
	VoiceOverEnabled *bool `json:"VoiceOverEnabled,omitempty" plist:"VoiceOverEnabled,omitempty"`
	// If `true`, the system enables zoom.
	ZoomEnabled *bool `json:"ZoomEnabled,omitempty" plist:"ZoomEnabled,omitempty"`
	// If `true`, the system enables grayscale display.
	GrayscaleEnabled *bool `json:"GrayscaleEnabled,omitempty" plist:"GrayscaleEnabled,omitempty"`
}

// Sets various accessibility settings. The system allows only keys with explicitly provided values.
type AccessibilitySettingsItem string

const (
	AccessibilitySettingsItemAccessibilitySettings AccessibilitySettingsItem = "AccessibilitySettings"
)

// The accessibility text size apps that support dynamic text use. `0` is the smallest value, and `11` is the largest available.
type SettingsAccessibilitySettingsTextSize int64

const (
	SettingsAccessibilitySettingsTextSize0  SettingsAccessibilitySettingsTextSize = 0
	SettingsAccessibilitySettingsTextSize1  SettingsAccessibilitySettingsTextSize = 1
	SettingsAccessibilitySettingsTextSize2  SettingsAccessibilitySettingsTextSize = 2
	SettingsAccessibilitySettingsTextSize3  SettingsAccessibilitySettingsTextSize = 3
	SettingsAccessibilitySettingsTextSize4  SettingsAccessibilitySettingsTextSize = 4
	SettingsAccessibilitySettingsTextSize5  SettingsAccessibilitySettingsTextSize = 5
	SettingsAccessibilitySettingsTextSize6  SettingsAccessibilitySettingsTextSize = 6
	SettingsAccessibilitySettingsTextSize7  SettingsAccessibilitySettingsTextSize = 7
	SettingsAccessibilitySettingsTextSize8  SettingsAccessibilitySettingsTextSize = 8
	SettingsAccessibilitySettingsTextSize9  SettingsAccessibilitySettingsTextSize = 9
	SettingsAccessibilitySettingsTextSize10 SettingsAccessibilitySettingsTextSize = 10
	SettingsAccessibilitySettingsTextSize11 SettingsAccessibilitySettingsTextSize = 11
)

// A dictionary that describes the results of configuring settings.
type Settings struct {
	// The status of the setting, which is one of the following values:
	// - `Acknowledged`: The device processed the command successfully.
	// - `Error`: An error occurred. See the `ErrorChain` for more details.
	Status string `json:"Status" plist:"Status" required:"true"`
	// An array of dictionaries that describes any errors that occurred.
	ErrorChain *[]*map[string]any `json:"ErrorChain,omitempty" plist:"ErrorChain,omitempty"`
	// The app identifier to which this error applies.
	// > Note:
	// > For a watchOS app, the identifier is the watch's bundle identifier, which differs from the main bundle identifier for the iPhone to which the watch is paired.
	Identifier *string `json:"Identifier,omitempty" plist:"Identifier,omitempty"`
}

// Get a list of available operating-system updates for a device.
// Queries the device for a list of available OS updates. On OS X, a ScheduleOSUpdateScan must be performed to update the results returned by this query.
type AvailableOSUpdatesCommand struct{}

func (p *AvailableOSUpdatesCommand) RequestType() string {
	return "AvailableOSUpdates"
}

type AvailableOSUpdatesCommandResponse struct {
	// An array of dictionaries that contains only the most recent available updates in iOS and tvOS, and possibly multiple available updates in macOS. Follow the instructions in the Managed Apps and Updates section of the Apple Software Lookup Service to find a complete catalog of iOS and tvOS updates.
	// In macOS 14 and later, `AvailableOSUpdates` doesn't include InstallAssistant-based, full-replacement installers. It only contains over-the-air (OTA) updates. OTA updates can update or upgrade the OS and support all `InstallAction` options.
	// If a Software Update is actively managed by a Declarative Device Management Specific Enforcement configuration, the device ignores this command as it applies to the actively managed update. This command can return information for unmanaged updates, such as System Applications and Configuration Data. For information about available updates when using Declarative Device Management, see [Using the Apple Software Lookup Service](https://support.apple.com/guide/deployment/depafd2fad80/web).
	AvailableOSUpdates []AvailableOSUpdatesItem `json:"AvailableOSUpdates" plist:"AvailableOSUpdates" required:"true"`
}

// The response dictionary that describes the available operating-system updates item.
type AvailableOSUpdatesItem struct {
	// The product key that represents the update.
	ProductKey string `json:"ProductKey" plist:"ProductKey" required:"true"`
	// The human-readable name of the update in the current user's current locale.
	HumanReadableName string `json:"HumanReadableName" plist:"HumanReadableName" required:"true"`
	// The locale, in IOS639-1 Alpha-2 code format, of the `HumanReadableName` value. This value is available in macOS 10.11 and later.
	HumanReadableNameLocale string `json:"HumanReadableNameLocale" plist:"HumanReadableNameLocale" required:"true"`
	// A URL where the MDM server can request additional localized names for this update. This key isn't present for certain updates, such as mobile software updates (MSUs) or major OS updates. This value is available in macOS 10.11 and later.
	MetadataURL string `json:"MetadataURL" plist:"MetadataURL" required:"true"`
	// The product name; for example, _iOS_. This value is available in iOS 9.0 and later, and tvOS 12.0 and later.
	ProductName string `json:"ProductName" plist:"ProductName" required:"true"`
	// The version of the update.
	Version string `json:"Version" plist:"Version" required:"true"`
	// The build number of the update.
	Build string `json:"Build" plist:"Build" required:"true"`
	// The storage size necessary to download the software update. Prior to macOS 10.14, this only includes major operating-system updates. In macOS 10.14 and later, this also includes minor updates.
	DownloadSize int64 `json:"DownloadSize" plist:"DownloadSize" required:"true"`
	// The storage size necessary to install the update. This value is available in iOS 9.0 and later, and tvOS 12.0 and later.
	InstallSize int64 `json:"InstallSize" plist:"InstallSize" required:"true"`
	// An array that contains app identifiers of apps to close so you can install the update. This value is available in macOS 10.11 and later.
	AppIdentifiersToClose []string `json:"AppIdentifiersToClose" plist:"AppIdentifiersToClose" required:"true"`
	// If `true`, this is a critical update.
	IsCritical *bool `json:"IsCritical,omitempty" plist:"IsCritical,omitempty"`
	// If `true`, this is an update to a configuration file. This value is available in macOS 10.11 and later.
	IsConfigDataUpdate *bool `json:"IsConfigDataUpdate,omitempty" plist:"IsConfigDataUpdate,omitempty"`
	// If `true`, this is an update to firmware. This value is available in macOS 10.11 and later.
	IsFirmwareUpdate *bool `json:"IsFirmwareUpdate,omitempty" plist:"IsFirmwareUpdate,omitempty"`
	// If `true`, this is a major update; for example, 10.15.x to 11. This value is available in macOS 10.11 and later.
	IsMajorOSUpdate *bool `json:"IsMajorOSUpdate,omitempty" plist:"IsMajorOSUpdate,omitempty"`
	// If `true`, the device restarts after installing the update.
	RestartRequired *bool `json:"RestartRequired,omitempty" plist:"RestartRequired,omitempty"`
	// If `true`, download the software update and install it later.
	AllowsInstallLater *bool `json:"AllowsInstallLater,omitempty" plist:"AllowsInstallLater,omitempty"`
	// If present, the date when you want the update to install. This value is available in macOS 10.12.4 and later.
	DeferredUntil *time.Time `json:"DeferredUntil,omitempty" plist:"DeferredUntil,omitempty"`
	// If `true`, the device can accept a Bootstrap Token from the MDM server instead of prompting for user authentication prior to installation. This only applies when `BootstrapTokenAllowedForAuthentication` is `true` in the `SecurityInfo` response. This value is available for Mac computers with Apple silicon in macOS 11 and later.
	RequiresBootstrapToken *bool `json:"RequiresBootstrapToken,omitempty" plist:"RequiresBootstrapToken,omitempty"`
	// If `true`, this update is a Rapid Security Response.
	IsSecurityResponse bool `json:"IsSecurityResponse" plist:"IsSecurityResponse" required:"true"`
	// The build version for the Rapid Security Response update, for example, `13A999`, which is the same as `Build`.
	SupplementalBuildVersion *string `json:"SupplementalBuildVersion,omitempty" plist:"SupplementalBuildVersion,omitempty"`
	// The Rapid Security Response OS version suffix, for example, `(a)`. Only present if this is a Rapid Security Response update.
	SupplementalOSVersionExtra *string `json:"SupplementalOSVersionExtra,omitempty" plist:"SupplementalOSVersionExtra,omitempty"`
}

// Schedule a background scan for operating-system updates on a device.
// Requests that the device perform a background scan for OS updates.
type ScheduleOSUpdateScanCommand struct {
	// If `true`, force a scan to start immediately. Otherwise, the scan starts at a system-determined time.
	Force *bool `json:"Force,omitempty" plist:"Force,omitempty"`
}

func (p *ScheduleOSUpdateScanCommand) RequestType() string {
	return "ScheduleOSUpdateScan"
}

type ScheduleOSUpdateScanCommandResponse struct {
	// If `true`, the scan started successfully.
	ScanInitiated bool `json:"ScanInitiated" plist:"ScanInitiated" required:"true"`
}

// Schedule an update of the operating system on a device.
// This command allows the server to schedule an OS update.
type ScheduleOSUpdateCommand struct {
	// An array of dictionaries specifying the updates to download or install. If this value is missing, the device applies the default behavior for handling updates.
	// The device ignores this command and an informational error is returned, if a software update is managed by a Declarative Device Management `SoftwareUpdateEnforcementSpecific` configuration, as the configuration takes precedence.
	Updates []UpdatesItem `json:"Updates" plist:"Updates" required:"true"`
}

func (p *ScheduleOSUpdateCommand) RequestType() string {
	return "ScheduleOSUpdate"
}

type ScheduleOSUpdateCommandResponse struct {
	// An array of dictionaries that describes the results of processing operating-system updates.
	UpdateResults []UpdateResultsItem `json:"UpdateResults" plist:"UpdateResults" required:"true"`
}

// A dictionary that describes the available operating-system updates item.
type UpdatesItem struct {
	// The product key that represents the update.
	ProductKey *string `json:"ProductKey,omitempty" plist:"ProductKey,omitempty"`
	// The version of the update, which the system requires if `ProductKey` isn't present. This value is available in iOS 11.3 and later, macOS 12 and later, and tvOS 12.2 and later.
	// > Note:
	// > This value isn't available for use with Rapid Security Response (RSR) updates.
	ProductVersion *string `json:"ProductVersion,omitempty" plist:"ProductVersion,omitempty"`
	// The install action, which is one of the following values:
	// - `Default`: Download or install the update, depending on the current state. You can check the `UpdateResults` dictionary to review scheduled updates. This value is available in iOS 9 and later, macOS 10.11 and later, and tvOS 12 and later.
	// - `DownloadOnly`: Download the software update without installing it. This value is available in iOS 9 and later, macOS 11 and later, and tvOS 12 and later.
	// - `InstallASAP`: In iOS and tvOS, install a previously downloaded software update. In macOS, download the software update and trigger the restart countdown notification. This value is available in iOS 9 and later, macOS 10.11 and later, and tvOS 12 and later.
	// - `NotifyOnly`: Download the software update and notify the user through the App Store. This value is available in macOS 10.11 and later.
	// - `InstallLater`: Download the software update and install it at a later time. This value is available in macOS 10.11 and later.
	// - `InstallForceRestart`: Perform the `Default` action, and then force a restart if the update requires it. This value is available in macOS 11 and later.
	// > Warning:
	// > `InstallForceRestart` may result in data loss.
	InstallAction UpdatesItemInstallAction `json:"InstallAction" plist:"InstallAction" required:"true"`
	// The maximum number of times the system allows the user to postpone an update before it's installed. The system prompts the user once a day.
	// This key is only supported when `InstallAction` is `InstallLater` and only supported for minor OS updates (for example, macOS 12.x to 12.y).
	MaxUserDeferrals *int64 `json:"MaxUserDeferrals,omitempty" plist:"MaxUserDeferrals,omitempty"`
	// The scheduling priority for downloading and preparing the requested update. This is only supported for minor OS updates (macOS 12.x to 12.y).
	// Available in macOS 12.3 and later. Prior versions of macOS used a priority of `Low`.
	Priority *Priority `default:"Low" json:"Priority,omitempty" plist:"Priority,omitempty"`
}

// The install action, which is one of the following values:
// - `Default`: Download or install the update, depending on the current state. You can check the `UpdateResults` dictionary to review scheduled updates. This value is available in iOS 9 and later, macOS 10.11 and later, and tvOS 12 and later.
// - `DownloadOnly`: Download the software update without installing it. This value is available in iOS 9 and later, macOS 11 and later, and tvOS 12 and later.
// - `InstallASAP`: In iOS and tvOS, install a previously downloaded software update. In macOS, download the software update and trigger the restart countdown notification. This value is available in iOS 9 and later, macOS 10.11 and later, and tvOS 12 and later.
// - `NotifyOnly`: Download the software update and notify the user through the App Store. This value is available in macOS 10.11 and later.
// - `InstallLater`: Download the software update and install it at a later time. This value is available in macOS 10.11 and later.
// - `InstallForceRestart`: Perform the `Default` action, and then force a restart if the update requires it. This value is available in macOS 11 and later.
// > Warning:
// > `InstallForceRestart` may result in data loss.
type UpdatesItemInstallAction string

const (
	UpdatesItemInstallActionDefault             UpdatesItemInstallAction = "Default"
	UpdatesItemInstallActionDownloadOnly        UpdatesItemInstallAction = "DownloadOnly"
	UpdatesItemInstallActionInstallASAP         UpdatesItemInstallAction = "InstallASAP"
	UpdatesItemInstallActionNotifyOnly          UpdatesItemInstallAction = "NotifyOnly"
	UpdatesItemInstallActionInstallLater        UpdatesItemInstallAction = "InstallLater"
	UpdatesItemInstallActionInstallForceRestart UpdatesItemInstallAction = "InstallForceRestart"
)

// The scheduling priority for downloading and preparing the requested update. This is only supported for minor OS updates (macOS 12.x to 12.y).
// Available in macOS 12.3 and later. Prior versions of macOS used a priority of `Low`.
type Priority string

const (
	PriorityLow  Priority = "Low"
	PriorityHigh Priority = "High"
)

// The response dictionary that describes the result of processing an operating-system update.
type UpdateResultsItem struct {
	// The product key that represents the update.
	ProductKey string `json:"ProductKey" plist:"ProductKey" required:"true"`
	// The install action that the device scheduled, which is one of the following values:
	// - `Error`: An error occurred during scheduling.
	// - `DownloadOnly`: Download the software update without installing it.
	// - `InstallASAP`: Install a previously downloaded software update.
	// - `NotifyOnly`: Download the software update and notify the user through the App Store. This value is available in macOS 10.11 and later.
	// - `InstallLater`: Download the software update and install it at a later time. This value is available in macOS 10.11 and later.
	// - `InstallForceRestart`: Perform the `Default` action, and then force a restart if the update requires it. This value is available in macOS 11 and later.
	InstallAction UpdateResultsItemInstallAction `json:"InstallAction" plist:"InstallAction" required:"true"`
	// The status of the update, which is one of the following values:
	// - `Idle`: The update is idle.
	// - `Downloading`: The software update is downloading.
	// - `DownloadFailed`: The download failed.
	// - `DownloadRequiresComputer`: Tether the device to download this update. This value is only available in iOS.
	// - `DownloadInsufficientSpace`: There isn't enough space to download the update.
	// - `DownloadInsufficientPower`: There isn't enough power to download the update.
	// - `DownloadInsufficientNetwork`: The network capacity is insufficient to download the update.
	// - `Installing`: The software update is installing.
	// - `InstallInsufficientSpace`: There isn't enough space to install the update.
	// - `InstallInsufficientPower`: There isn't enough power to install the update.
	// - `InstallPhoneCallInProgress`: Installation couldn't occur because a phone call is in progress.
	// - `InstallFailed`: Installation failed due to an unspecified reason.
	Status UpdateResultsItemStatus `json:"Status" plist:"Status" required:"true"`
	// A dictionary that describes an error chain.
	ErrorChain *[]*map[string]any `json:"ErrorChain,omitempty" plist:"ErrorChain,omitempty"`
}

// The install action that the device scheduled, which is one of the following values:
// - `Error`: An error occurred during scheduling.
// - `DownloadOnly`: Download the software update without installing it.
// - `InstallASAP`: Install a previously downloaded software update.
// - `NotifyOnly`: Download the software update and notify the user through the App Store. This value is available in macOS 10.11 and later.
// - `InstallLater`: Download the software update and install it at a later time. This value is available in macOS 10.11 and later.
// - `InstallForceRestart`: Perform the `Default` action, and then force a restart if the update requires it. This value is available in macOS 11 and later.
type UpdateResultsItemInstallAction string

const (
	UpdateResultsItemInstallActionError               UpdateResultsItemInstallAction = "Error"
	UpdateResultsItemInstallActionDownloadOnly        UpdateResultsItemInstallAction = "DownloadOnly"
	UpdateResultsItemInstallActionInstallASAP         UpdateResultsItemInstallAction = "InstallASAP"
	UpdateResultsItemInstallActionNotifyOnly          UpdateResultsItemInstallAction = "NotifyOnly"
	UpdateResultsItemInstallActionInstallLater        UpdateResultsItemInstallAction = "InstallLater"
	UpdateResultsItemInstallActionInstallForceRestart UpdateResultsItemInstallAction = "InstallForceRestart"
)

// The status of the update, which is one of the following values:
// - `Idle`: The update is idle.
// - `Downloading`: The software update is downloading.
// - `DownloadFailed`: The download failed.
// - `DownloadRequiresComputer`: Tether the device to download this update. This value is only available in iOS.
// - `DownloadInsufficientSpace`: There isn't enough space to download the update.
// - `DownloadInsufficientPower`: There isn't enough power to download the update.
// - `DownloadInsufficientNetwork`: The network capacity is insufficient to download the update.
// - `Installing`: The software update is installing.
// - `InstallInsufficientSpace`: There isn't enough space to install the update.
// - `InstallInsufficientPower`: There isn't enough power to install the update.
// - `InstallPhoneCallInProgress`: Installation couldn't occur because a phone call is in progress.
// - `InstallFailed`: Installation failed due to an unspecified reason.
type UpdateResultsItemStatus string

const (
	UpdateResultsItemStatusIdle                        UpdateResultsItemStatus = "Idle"
	UpdateResultsItemStatusDownloading                 UpdateResultsItemStatus = "Downloading"
	UpdateResultsItemStatusDownloadFailed              UpdateResultsItemStatus = "DownloadFailed"
	UpdateResultsItemStatusDownloadRequiresComputer    UpdateResultsItemStatus = "DownloadRequiresComputer"
	UpdateResultsItemStatusDownloadInsufficientSpace   UpdateResultsItemStatus = "DownloadInsufficientSpace"
	UpdateResultsItemStatusDownloadInsufficientPower   UpdateResultsItemStatus = "DownloadInsufficientPower"
	UpdateResultsItemStatusDownloadInsufficientNetwork UpdateResultsItemStatus = "DownloadInsufficientNetwork"
	UpdateResultsItemStatusInstalling                  UpdateResultsItemStatus = "Installing"
	UpdateResultsItemStatusInstallInsufficientSpace    UpdateResultsItemStatus = "InstallInsufficientSpace"
	UpdateResultsItemStatusInstallInsufficientPower    UpdateResultsItemStatus = "InstallInsufficientPower"
	UpdateResultsItemStatusInstallPhoneCallInProgress  UpdateResultsItemStatus = "InstallPhoneCallInProgress"
	UpdateResultsItemStatusInstallFailed               UpdateResultsItemStatus = "InstallFailed"
)

// Get the status of operating-system updates on a device.
// Queries the device for the status of software updates.
type OSUpdateStatusCommand struct{}

func (p *OSUpdateStatusCommand) RequestType() string {
	return "OSUpdateStatus"
}

type OSUpdateStatusCommandResponse struct {
	// An array of dictionaries that describes the statuses of software updates. The array is empty if there are no software updates currently in progress.
	// This command only returns the status for System Applications and Configuration Data updates when a software update is managed by a Declarative Device Management `SoftwareUpdateEnforcementSpecific` configuration.
	OSUpdateStatus []OSUpdateStatusItem `json:"OSUpdateStatus" plist:"OSUpdateStatus" required:"true"`
}

// A dictionary that describes the status of a software update.
type OSUpdateStatusItem struct {
	// The product key that represents the update.
	ProductKey string `json:"ProductKey" plist:"ProductKey" required:"true"`
	// If `true`, the update has finished downloading.
	IsDownloaded bool `json:"IsDownloaded" plist:"IsDownloaded" required:"true"`
	// A floating-point number between `0.0` and `1.0` that indicates the download progress as a percentage.
	DownloadPercentComplete float64 `json:"DownloadPercentComplete" plist:"DownloadPercentComplete" required:"true"`
	// The status of the update, which is one of the following values:
	// - `Idle`: The update is idle.
	// - `Downloading`: The software update is downloading and subsequently preparing.
	// - `Installing`: The software update is installing.
	Status string `json:"Status" plist:"Status" required:"true"`
	// The number of times a user can defer this OS update.
	// Available in macOS 12.3 and later.
	MaxDeferrals *int64 `json:"MaxDeferrals,omitempty" plist:"MaxDeferrals,omitempty"`
	// The number of remaining user deferrals for this OS update.
	// Available in macOS 12.3 and later.
	DeferralsRemaining *int64 `json:"DeferralsRemaining,omitempty" plist:"DeferralsRemaining,omitempty"`
	// The date of the next attempt at installing this OS update.
	// Available in macOS 12.3 and later.
	NextScheduledInstall *time.Time `json:"NextScheduledInstall,omitempty" plist:"NextScheduledInstall,omitempty"`
	// The dates/times when the OS notified the user about installing this OS update.
	// Available in macOS 12.3 and later.
	PastNotifications *[]time.Time `json:"PastNotifications,omitempty" plist:"PastNotifications,omitempty"`
}

// Inform the device that it can continue past Setup Assistant and finish login.
// Inform the device that it can continue past Setup Assistant and finish login. Only works on Shared iPads that have the AwaitUserConfiguration feature enabled.
type UserConfiguredCommand struct{}

func (p *UserConfiguredCommand) RequestType() string {
	return "UserConfigured"
}

type UserConfiguredCommandResponse struct{}

// Delete a user's account from a device.
// This command allows the server to delete a user that has an active account on the device.
type DeleteUserCommand struct {
	// The user name of the account to delete. This key is required when the value for `DeleteAllUsers` is absent or `false`.
	UserName *string `json:"UserName,omitempty" plist:"UserName,omitempty"`
	// If `true`, the system deletes the account even if the user has data that's pending sync to the cloud. This value is available on iOS 9.3 and later.
	ForceDeletion *bool `json:"ForceDeletion,omitempty" plist:"ForceDeletion,omitempty"`
	// If `true`, the system attempts to delete all users from the device. If `ForceDeletion` is `false`, the system generates an error instead and doesn't delete users who have data that's pending sync. This value is available in iOS 14 and later.
	DeleteAllUsers *bool `json:"DeleteAllUsers,omitempty" plist:"DeleteAllUsers,omitempty"`
}

func (p *DeleteUserCommand) RequestType() string {
	return "DeleteUser"
}

type DeleteUserCommandResponse struct{}

// Get a list of users with active accounts on a device.
// This command allows the server to query for a list of users that have an active account on the device.
type UserListCommand struct{}

func (p *UserListCommand) RequestType() string {
	return "UserList"
}

type UserListCommandResponse struct {
	// An array of user dictionaries that contains information about the active accounts.
	Users []UsersItem `json:"Users" plist:"Users" required:"true"`
}

// A dictionary that contains information about an active account on a device.
type UsersItem struct {
	// The user name for the account. In macOS, this is the short name of the user account. This value is available in iOS 9.3 and later, and macOS 10.13 and later.
	UserName string `json:"UserName" plist:"UserName" required:"true"`
	// The user's full name. This value is available in macOS 10.13 and later.
	FullName string `json:"FullName" plist:"FullName" required:"true"`
	// The user's unique identifier. This value is available in macOS 10.13 and later.
	UID int64 `json:"UID" plist:"UID" required:"true"`
	// The user's `GeneratedUID`. This value is available in macOS 10.13 and later.
	UserGUID string `json:"UserGUID" plist:"UserGUID" required:"true"`
	// If `true`, the user is currently logged in on the device. This value is available in iOS 9.3 and later, and macOS 10.13 and later.
	IsLoggedIn bool `json:"IsLoggedIn" plist:"IsLoggedIn" required:"true"`
	// If `true`, the user has data to sync to the cloud. This value is available in iOS 9.3 and later.
	HasDataToSync bool `json:"HasDataToSync" plist:"HasDataToSync" required:"true"`
	// If present, the user's data quota in bytes. This isn't present if the account doesn't enforce a quota. This value is available in iOS 9.3 and later.
	DataQuota int64 `json:"DataQuota" plist:"DataQuota" required:"true"`
	// The amount of data, in bytes, that the user has used. This value is available in iOS 9.3 and later.
	DataUsed int64 `json:"DataUsed" plist:"DataUsed" required:"true"`
	// If `true`, the account is a mobile account. This value is available in macOS 10.13 and later.
	MobileAccount bool `json:"MobileAccount" plist:"MobileAccount" required:"true"`
	// If `true`, the user currently has a secure token set. This value is available in macOS 11 and later.
	HasSecureToken bool `json:"HasSecureToken" plist:"HasSecureToken" required:"true"`
}

// Force the current user to log out of a device.
// This command allows the server to force the current user to logout.
type LogOutUserCommand struct{}

func (p *LogOutUserCommand) RequestType() string {
	return "LogOutUser"
}

type LogOutUserCommandResponse struct{}

// Unlock a user account that the system locked because of too many failed password attempts.
// This command allows the server to unlock a local user account that has been locked due to too many failed password attempts. Requires "Device lock and passcode removal right".
type UnlockUserAccountCommand struct {
	// The user name of the local account, which can be any local account on the system, not just a managed user account.
	UserName string `json:"UserName" plist:"UserName" required:"true"`
}

func (p *UnlockUserAccountCommand) RequestType() string {
	return "UnlockUserAccount"
}

type UnlockUserAccountCommandResponse struct{}
