// DO NOT EDIT
// generated from https://github.com/apple/device-management.git:f878dea98fb88293a3686e44bcfb891f8e78f98f/other

package other

// Enrollment SSO streamlines the MDM enrollment process, reduces sign-ins, and improves security.
type EnrollmentSSODocument struct {
	// The iTunes Store ID of the app to download prior to enrollment, to support Enrollment SSO during enrollment. Using developer mode ignores this key.
	ITunesStoreID *int64 `json:"iTunesStoreID,omitempty" plist:"iTunesStoreID,omitempty"`
	// An array of App IDs that specify apps that Enrollment SSO developer mode can use. In Enrollment SSO documents delivered through the developer endpoint, this key must be present and contain at least one value. In Enrollment SSO documents delivered by the standard Enrollment SSO endpoint, this key must not be present.
	AppIDs *[]string `json:"AppIDs,omitempty" plist:"AppIDs,omitempty"`
	// An array of associated domains that the device uses with the Enrollment SSO extension.
	AssociatedDomains *[]string `json:"AssociatedDomains,omitempty" plist:"AssociatedDomains,omitempty"`
	// If `true,` allows the domain to directly verify site association, instead of at Apple's servers. Use this verification only with domains that are inaccessible on the public Internet.
	AssociatedDomainsEnableDirectDownloads *bool `json:"AssociatedDomainsEnableDirectDownloads,omitempty" plist:"AssociatedDomainsEnableDirectDownloads,omitempty"`
	// The profile containing an `ExtensibleSingleSignOn` payload that specifies the SSO extension in the downloaded app prior to enrollment. This profile may contain certificate payloads.
	// One of `ConfigurationProfile` and `Declarations` must be present.
	ConfigurationProfile *[]byte `json:"ConfigurationProfile,omitempty" plist:"ConfigurationProfile,omitempty"`
	// An array of base64-encoded JSON formatted Declarative Device Management declarations that specify the managed app and its configuration, including any certificates or identities.
	// The set of declarations must include one `com.apple.configuration.app.managed` configuration, and one activation declaration that references the configuration. Asset declarations may be present if required by the app config.
	// The app configuration must include `AppStoreID` when developer mode is not being used, or it must include `BundleID` when developer mode is used.
	// One of `ConfigurationProfile` and `Declarations` must be present.
	Declarations *[][]byte `json:"Declarations,omitempty" plist:"Declarations,omitempty"`
}

// A device's information in response to a MDM enrollment profile request.
type MachineInfo struct {
	// The device's UDID.
	UDID string `json:"UDID" plist:"UDID" required:"true"`
	// The device's serial number.
	SERIAL string `json:"SERIAL" plist:"SERIAL" required:"true"`
	// The device's product type, for example, `iPhone5,1`.
	PRODUCT string `json:"PRODUCT" plist:"PRODUCT" required:"true"`
	// The build version installed on the device, for example, `7A182`.
	VERSION string `json:"VERSION" plist:"VERSION" required:"true"`
	// The OS version installed on the device, for example, 17.0.
	// Available on iOS 17 and later, macOS 14 and later, tvOS 17 and later, and watchOS 10 and later.
	OSVERSION string `json:"OS_VERSION" plist:"OS_VERSION" required:"true"`
	// The device's operating system supplemental build version (if available).
	// Available on iOS 17 and later, macOS 14 and later, tvOS 17 and later, and watchOS 10 and later.
	SUPPLEMENTALBUILDVERSION *string `json:"SUPPLEMENTAL_BUILD_VERSION,omitempty" plist:"SUPPLEMENTAL_BUILD_VERSION,omitempty"`
	// The device's operating system supplemental OS version extra (if available).
	// Available on iOS 17 and later, macOS 14 and later, tvOS 17 and later, and watchOS 10 and later.
	SUPPLEMENTALOSVERSIONEXTRA *string `json:"SUPPLEMENTAL_OS_VERSION_EXTRA,omitempty" plist:"SUPPLEMENTAL_OS_VERSION_EXTRA,omitempty"`
	// The device's IMEI (if available).
	IMEI *string `json:"IMEI,omitempty" plist:"IMEI,omitempty"`
	// The device's MEID (if available).
	MEID *string `json:"MEID,omitempty" plist:"MEID,omitempty"`
	// The user's currently-selected language, for example, `en`.
	LANGUAGE *string `json:"LANGUAGE,omitempty" plist:"LANGUAGE,omitempty"`
	// If `true`, indicates that the server can trigger the device to do a required software update.
	// Available on iOS 17 and later, and macOS 14 and later.
	MDMCANREQUESTSOFTWAREUPDATE *bool `json:"MDM_CAN_REQUEST_SOFTWARE_UPDATE,omitempty" plist:"MDM_CAN_REQUEST_SOFTWARE_UPDATE,omitempty"`
	// The pairing token to validate when a watch is enrolling.
	// Available on watchOS 10 and later.
	PAIRINGTOKEN *[]byte `json:"PAIRING_TOKEN,omitempty" plist:"PAIRING_TOKEN,omitempty"`
	// The device model identifier used to lookup available OS updates through https://gdmf.apple.com/v2/pmv. Available on iOS 17.4 and later, macOS 14.4 and later, and visionOS 1.1 and later.
	SOFTWAREUPDATEDEVICEID *string `json:"SOFTWARE_UPDATE_DEVICE_ID,omitempty" plist:"SOFTWARE_UPDATE_DEVICE_ID,omitempty"`
	// If `true`, indicates that the server can trigger the device to do a required Platform SSO authentication before enrolling.
	// Available on macOS 26 and later.
	MDMCANREQUESTPSSOCONFIG *bool `json:"MDM_CAN_REQUEST_PSSO_CONFIG,omitempty" plist:"MDM_CAN_REQUEST_PSSO_CONFIG,omitempty"`
	// If `true`, indicates that the device requires a mandatory software update during Setup Assistant. The MDM server can return a 403 with a `ErrorCodeSoftwareUpdateRequired` error to force the device to update to a specific version instead of the device choosing a version.
	// Available on macOS 26.1 and later.
	MANDATORYSOFTWAREUPDATEREQUIRED *bool `json:"MANDATORY_SOFTWARE_UPDATE_REQUIRED,omitempty" plist:"MANDATORY_SOFTWARE_UPDATE_REQUIRED,omitempty"`
}

// The URL to the app manifest.
type ManifestURL struct {
	// An array of dictionaries representing what the manifest installs.
	Items []ItemsItem `json:"items" plist:"items" required:"true"`
}
type ItemsItem struct {
	// An array of dictionaries that describe an item to install.
	Assets []AssetsItem `json:"assets" plist:"assets" required:"true"`
	// The metadata for an application or package manifest item.
	Metadata Metadata `json:"metadata" plist:"metadata" required:"true"`
}
type AssetsItem struct {
	// The kind of manifest item to install. Use `software-package` for apps and macOS packages.
	Kind AssetsItemKind `json:"kind" plist:"kind" required:"true"`
	// The MD5 hash value the device uses when verifying the hash of the manifest item data. When this key is present, the device ignores the `md5-size` and `md5s` keys.
	Md5 *string `json:"md5,omitempty" plist:"md5,omitempty"`
	// The data *chunk* size the device uses when verifying the hash of the manifest item data. Required when the `md5s` key is present.
	Md5Size *int64 `json:"md5-size,omitempty" plist:"md5-size,omitempty"`
	// An array of strings representing a set of MD5 hash values. The device uses these values to verify the integrity of the downloaded manifest item data. Required when the `md5-size` key is present.
	Md5s *[]string `json:"md5s,omitempty" plist:"md5s,omitempty"`
	// The SHA-256 hash value the device uses when verifying the hash of the manifest item data. When this key is present, the device ignores the `sha256-size` and `sha256` keys.
	SHA256 *string `json:"sha256,omitempty" plist:"sha256,omitempty"`
	// The data *chunk* size the device uses when verifying the hash of the manifest item data. Required when the `sha256s` key is present.
	SHA256Size *int64 `json:"sha256-size,omitempty" plist:"sha256-size,omitempty"`
	// An array of strings representing a set of SHA-256 hash values. The device uses these values to verify the integrity of the downloaded manifest item data. Required when the `sha256-size` key is present.
	SHA256s *[]string `json:"sha256s,omitempty" plist:"sha256s,omitempty"`
	// The URL that hosts the manifest item data. The URL needs to start with `https://`.
	URL string `json:"url" plist:"url" required:"true"`
	// removed
	NeedsShine *bool `json:"needs-shine,omitempty" plist:"needs-shine,omitempty"`
}

// The kind of manifest item to install. Use `software-package` for apps and macOS packages.
type AssetsItemKind string

const (
	AssetsItemKindAssetPackManifest AssetsItemKind = "asset-pack-manifest"
	AssetsItemKindDisplayImage      AssetsItemKind = "display-image"
	AssetsItemKindFullSizeImage     AssetsItemKind = "full-size-image"
	AssetsItemKindSoftwarePackage   AssetsItemKind = "software-package"
)

// The metadata for an application or package manifest item.
type Metadata struct {
	// The bundle identifier of the app or package manifest item.
	BundleIdentifier string `json:"bundle-identifier" plist:"bundle-identifier" required:"true"`
	// The bundle version of the app or package manifest item.
	BundleVersion *string `json:"bundle-version,omitempty" plist:"bundle-version,omitempty"`
	// The type of manifest item. For an app or package, this needs to be `software`.
	Kind MetadataKind `json:"kind" plist:"kind" required:"true"`
	// removed
	SizeInBytes *int64 `json:"sizeInBytes,omitempty" plist:"sizeInBytes,omitempty"`
	// The name of the app or package developer.
	Subtitle *string `json:"subtitle,omitempty" plist:"subtitle,omitempty"`
	// The title of the app or package being installed.
	Title string `json:"title" plist:"title" required:"true"`
	// Removed
	Items *[]*MetadataItemsItem `json:"items,omitempty" plist:"items,omitempty"`
}

// The type of manifest item. For an app or package, this needs to be `software`.
type MetadataKind string

const (
	MetadataKindSoftware MetadataKind = "software"
)

type MetadataItemsItem struct {
	// removed
	BundleIdentifier *string `json:"bundle-identifier,omitempty" plist:"bundle-identifier,omitempty"`
	// removed
	BundleVersion *string `json:"bundle-version,omitempty" plist:"bundle-version,omitempty"`
}

// A dictionary that contains the password hash for the account.
type PasswordHash struct {
	// A dictionary that contains the `entropy`, `iterations`, and `salt` elements to create the password hash using the CommonCrypto libraries, or equivalent. Convert this dictionary to binary data before setting it as the value for the password hash.
	SALTEDSHA512PBKDF2 SALTEDSHA512PBKDF2 `json:"SALTED-SHA512-PBKDF2" plist:"SALTED-SHA512-PBKDF2" required:"true"`
}

// A dictionary that contains the `entropy`, `iterations`, and `salt` elements to create the password hash using the CommonCrypto libraries, or equivalent. Convert this dictionary to binary data before setting it as the value for the password hash.
type SALTEDSHA512PBKDF2 struct {
	// The derived key from the password hash; for example, from `CCKeyDerivationPBKDF()`.
	Entropy []byte `json:"entropy" plist:"entropy" required:"true"`
	// The number of iterations; for example, from `CCCalibratePBKDF()` using a minimum hash time of 100 milliseconds, or if unknown, a number in the range of 20,000 to 40,000 iterations.
	Iterations int64 `json:"iterations" plist:"iterations" required:"true"`
	// The 32-byte randomized data; for example, from `CCRandomCopyBytes()`.
	Salt []byte `json:"salt" plist:"salt" required:"true"`
}

// The list of skip keys for setup panes.
type SkipKeys struct {
	// The key to skip the Accessibility pane, when creating additional users.
	// > Note:
	// > This key doesn't skip the Accessibility pane in Setup Assistant during initial device set up. It does skip the Accessibility pane when Setup Assistant runs due to a new user log in.
	Accessibility *string `json:"Accessibility,omitempty" plist:"Accessibility,omitempty"`
	// The key to skip the Action Button configuration pane.
	ActionButton *string `json:"ActionButton,omitempty" plist:"ActionButton,omitempty"`
	// The key to skip the Age Assurance pane.
	AgeAssurance *string `json:"AgeAssurance,omitempty" plist:"AgeAssurance,omitempty"`
	// The key to skip the Age Based Safety Settings pane.
	AgeBasedSafetySettings *string `json:"AgeBasedSafetySettings,omitempty" plist:"AgeBasedSafetySettings,omitempty"`
	// If the Restore pane isn't skipped, this is the key to remove the Move from Android option in the Restore pane.
	Android *string `json:"Android,omitempty" plist:"Android,omitempty"`
	// The key to skip the Choose Your Look screen.
	Appearance *string `json:"Appearance,omitempty" plist:"Appearance,omitempty"`
	// The key to skip Apple Account setup.
	AppleID *string `json:"AppleID,omitempty" plist:"AppleID,omitempty"`
	// The key to skip the App Store pane.
	AppStore *string `json:"AppStore,omitempty" plist:"AppStore,omitempty"`
	// The key to skip biometric setup.
	Biometric *string `json:"Biometric,omitempty" plist:"Biometric,omitempty"`
	// The key to skip the Camera Button pane.
	CameraButton *string `json:"CameraButton,omitempty" plist:"CameraButton,omitempty"`
	// The key to skip Device to Device Migration pane.
	DeviceToDeviceMigration *string `json:"DeviceToDeviceMigration,omitempty" plist:"DeviceToDeviceMigration,omitempty"`
	// The key to skip the App Analytics pane.
	Diagnostics *string `json:"Diagnostics,omitempty" plist:"Diagnostics,omitempty"`
	// The key to skip DisplayTone setup.
	DisplayTone *string `json:"DisplayTone,omitempty" plist:"DisplayTone,omitempty"`
	// The key to skip the Lockdown Mode pane if an Apple Account is set up.
	EnableLockdownMode *string `json:"EnableLockdownMode,omitempty" plist:"EnableLockdownMode,omitempty"`
	// The key to disable the FileVault Setup Assistant screen.
	FileVault *string `json:"FileVault,omitempty" plist:"FileVault,omitempty"`
	// The key to skip the Meet the New Home Button screen on iPhone 7, iPhone 7 Plus, iPhone 8, iPhone 8 Plus, and iPhone SE.
	HomeButtonSensitivity *string `json:"HomeButtonSensitivity,omitempty" plist:"HomeButtonSensitivity,omitempty"`
	// The key to skip the iCloud Analytics screen.
	ICloudDiagnostics *string `json:"iCloudDiagnostics,omitempty" plist:"iCloudDiagnostics,omitempty"`
	// The key to skip the iCloud Documents and Desktop screen.
	ICloudStorage *string `json:"iCloudStorage,omitempty" plist:"iCloudStorage,omitempty"`
	// The key to skip the iMessage and FaceTime screen.
	IMessageAndFaceTime *string `json:"iMessageAndFaceTime,omitempty" plist:"iMessageAndFaceTime,omitempty"`
	// The key to skip the Intelligence pane.
	Intelligence *string `json:"Intelligence,omitempty" plist:"Intelligence,omitempty"`
	// The key to skip the Keyboard pane. This pane isn't always skippable because it appears before the device retrieves the Cloud Configuration from the server.
	Keyboard *string `json:"Keyboard,omitempty" plist:"Keyboard,omitempty"`
	// The key to disable Location Services.
	Location *string `json:"Location,omitempty" plist:"Location,omitempty"`
	// The key to skip the iMessage pane.
	MessagingActivationUsingPhoneNumber *string `json:"MessagingActivationUsingPhoneNumber,omitempty" plist:"MessagingActivationUsingPhoneNumber,omitempty"`
	// The key to skip the Multitasking pane.
	Multitasking *string `json:"Multitasking,omitempty" plist:"Multitasking,omitempty"`
	// The key to skip the on-boarding informational screens for user education (Go Home, Cover Sheet, Multitasking & Control Center, for example).
	OnBoarding *string `json:"OnBoarding,omitempty" plist:"OnBoarding,omitempty"`
	// The key to skip the OS Showcase pane.
	OSShowcase *string `json:"OSShowcase,omitempty" plist:"OSShowcase,omitempty"`
	// The key to hide and disable the passcode pane.
	Passcode *string `json:"Passcode,omitempty" plist:"Passcode,omitempty"`
	// The key to skip Apple Pay setup.
	Payment *string `json:"Payment,omitempty" plist:"Payment,omitempty"`
	// The key to skip the privacy pane.
	Privacy *string `json:"Privacy,omitempty" plist:"Privacy,omitempty"`
	// The key to disable restoring from backup.
	Restore *string `json:"Restore,omitempty" plist:"Restore,omitempty"`
	// The key to skip the Restore Completed pane.
	RestoreCompleted *string `json:"RestoreCompleted,omitempty" plist:"RestoreCompleted,omitempty"`
	// The key to skip the tvOS screen about using aerial screensavers on an Apple TV.
	ScreenSaver *string `json:"ScreenSaver,omitempty" plist:"ScreenSaver,omitempty"`
	// The key to skip the Safety pane.
	Safety *string `json:"Safety,omitempty" plist:"Safety,omitempty"`
	// The key to skip the Safety and Handling pane. This pane isn't always skippable because it appears before the device retrieves the Cloud Configuration from the server.
	SafetyAndHandling *string `json:"SafetyAndHandling,omitempty" plist:"SafetyAndHandling,omitempty"`
	// The key to skip the Screen Time pane.
	ScreenTime *string `json:"ScreenTime,omitempty" plist:"ScreenTime,omitempty"`
	// The key to skip the add cellular plan pane. Skipping this pane prevents automatic eSIM setup during Setup Assistant.
	SIMSetup *string `json:"SIMSetup,omitempty" plist:"SIMSetup,omitempty"`
	// The key to disable Siri.
	Siri *string `json:"Siri,omitempty" plist:"Siri,omitempty"`
	// The key to skip the mandatory software update screen.
	SoftwareUpdate *string `json:"SoftwareUpdate,omitempty" plist:"SoftwareUpdate,omitempty"`
	// The key to skip the Dictation pane. This pane isn't always skippable because it appears before the device retrieves the Cloud Configuration from the server.
	SpokenLanguage *string `json:"SpokenLanguage,omitempty" plist:"SpokenLanguage,omitempty"`
	// The key to skip the Tap To Set Up option in Apple TV related to using an iOS device to set up your Apple TV.
	TapToSetup *string `json:"TapToSetup,omitempty" plist:"TapToSetup,omitempty"`
	// The key to skip the Terms of Address pane. This key isn't always skippable because this pane appears before the device retrieves the Cloud Configuration from the server.
	TermsOfAddress *string `json:"TermsOfAddress,omitempty" plist:"TermsOfAddress,omitempty"`
	// The key to skip the Tips pane.
	Tips *string `json:"Tips,omitempty" plist:"Tips,omitempty"`
	// The key to skip Terms and Conditions.
	TOS *string `json:"TOS,omitempty" plist:"TOS,omitempty"`
	// The key to skip TV Home Screen layout sync screen.
	TVHomeScreenSync *string `json:"TVHomeScreenSync,omitempty" plist:"TVHomeScreenSync,omitempty"`
	// The key to skip the TV provider sign in screen.
	TVProviderSignIn *string `json:"TVProviderSignIn,omitempty" plist:"TVProviderSignIn,omitempty"`
	// The key to skip the "Where is this Apple TV?" screen.
	TVRoom *string `json:"TVRoom,omitempty" plist:"TVRoom,omitempty"`
	// The key to skip the "Unlock with Apple Watch" screen.
	UnlockWithWatch *string `json:"UnlockWithWatch,omitempty" plist:"UnlockWithWatch,omitempty"`
	// The key to skip the Software Update Complete pane.
	UpdateCompleted *string `json:"UpdateCompleted,omitempty" plist:"UpdateCompleted,omitempty"`
	// The key to skip Wallpaper setup.
	Wallpaper *string `json:"Wallpaper,omitempty" plist:"Wallpaper,omitempty"`
	// The key to skip the screen for watch migration.
	WatchMigration *string `json:"WatchMigration,omitempty" plist:"WatchMigration,omitempty"`
	// The key to skip the Web Content Filtering pane.
	WebContentFiltering *string `json:"WebContentFiltering,omitempty" plist:"WebContentFiltering,omitempty"`
	// The key to skip the Get Started pane.
	Welcome *string `json:"Welcome,omitempty" plist:"Welcome,omitempty"`
	// The key to skip zoom setup.
	Zoom *string `json:"Zoom,omitempty" plist:"Zoom,omitempty"`
}
