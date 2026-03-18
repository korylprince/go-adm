// DO NOT EDIT
// generated from https://github.com/apple/device-management.git:f878dea98fb88293a3686e44bcfb891f8e78f98f/mdm/profiles/profiles

package profiles

import "time"

const DeviceManagementGenerateHash = "f878dea98fb88293a3686e44bcfb891f8e78f98f"

var PayloadMap = map[string]any{
	".GlobalPreferences":                                     GlobalPreferences{},
	"com.apple.ADCertificate.managed":                        ActiveDirectoryCertificate{},
	"com.apple.AIM.account":                                  AIMAccount{},
	"com.apple.AssetCache.managed":                           ContentCaching{},
	"com.apple.Dictionary":                                   ParentalControlsDictionary{},
	"com.apple.DirectoryService.managed":                     DirectoryService{},
	"com.apple.DiscRecording":                                MediaManagementDiscBurning{},
	"com.apple.MCX.Accounts":                                 Accounts{},
	"com.apple.MCX.EnergySaver":                              EnergySaver{},
	"com.apple.MCX.FDEFileVaultOptions":                      FDEFileVaultOptions{},
	"com.apple.MCX.FileVault2":                               FDEFileVault{},
	"com.apple.MCX.MobileAccounts":                           MobileAccounts{},
	"com.apple.MCX.TimeMachine":                              TimeMachine{},
	"com.apple.MCX.TimeServer":                               TimeServer{},
	"com.apple.MCX.WiFiManagedSettings":                      WiFiManagedSettings{},
	"com.apple.ManagedClient.preferences":                    ManagedPreferences{},
	"com.apple.NSExtension":                                  NSExtensionManagement{},
	"com.apple.SetupAssistant.managed":                       SetupAssistant{},
	"com.apple.ShareKitHelper":                               ShareKit{},
	"com.apple.SoftwareUpdate":                               SoftwareUpdate{},
	"com.apple.SystemConfiguration":                          NetworkProxyConfiguration{},
	"com.apple.TCC.configuration-profile-policy":             PrivacyPreferencesPolicyControl{},
	"com.apple.airplay":                                      AirPlay{},
	"com.apple.airplay.security":                             AirPlaySecurity{},
	"com.apple.airprint":                                     AirPrint{},
	"com.apple.apn.managed":                                  APN{},
	"com.apple.app.lock":                                     AppLock{},
	"com.apple.applicationaccess":                            Restrictions{},
	"com.apple.applicationaccess.new":                        ParentalControlsApplicationRestrictions{},
	"com.apple.appstore":                                     AppStore{},
	"com.apple.asam":                                         AutonomousSingleAppMode{},
	"com.apple.associated-domains":                           AssociatedDomains{},
	"com.apple.caldav.account":                               CalDAV{},
	"com.apple.carddav.account":                              CardDAV{},
	"com.apple.cellular":                                     Cellular{},
	"com.apple.cellularprivatenetwork.managed":               CellularPrivateNetwork{},
	"com.apple.conferenceroomdisplay":                        ConferenceRoomDisplay{},
	"com.apple.configurationprofile.identification":          Identification{},
	"com.apple.dashboard":                                    ParentalControlsDashboardWidgetRestrictions{},
	"com.apple.declarations":                                 Declarations{},
	"com.apple.desktop":                                      Desktop{},
	"com.apple.dnsProxy.managed":                             DNSProxy{},
	"com.apple.dnsSettings.managed":                          DNSSettings1{},
	"com.apple.dock":                                         Dock{},
	"com.apple.domains":                                      Domains{},
	"com.apple.eas.account":                                  ExchangeActiveSync{},
	"com.apple.education":                                    EducationConfiguration{},
	"com.apple.ews.account":                                  ExchangeWebServices{},
	"com.apple.extensiblesso.ExtensibleSingleSignOn":         ExtensibleSingleSignOn{},
	"com.apple.extensiblesso.ExtensibleSingleSignOnKerberos": ExtensibleSingleSignOnKerberos{},
	"com.apple.familycontrols.contentfilter":                 ParentalControlsContentFilter{},
	"com.apple.familycontrols.timelimits.v2":                 ParentalControlsTimeLimits{},
	"com.apple.fileproviderd":                                FileProvider{},
	"com.apple.finder":                                       Finder{},
	"com.apple.firstactiveethernet.managed":                  Net8021XFirstActiveEthernet{},
	"com.apple.firstethernet.managed":                        Net8021XFirstEthernet{},
	"com.apple.font":                                         Font{},
	"com.apple.gamed":                                        ParentalControlsGameCenter{},
	"com.apple.globalethernet.managed":                       Net8021XGlobalEthernet{},
	"com.apple.google-oauth":                                 GoogleAccount{},
	"com.apple.homescreenlayout":                             HomeScreenLayout{},
	"com.apple.ironwood.support":                             ParentalControlDictationandProfanity{},
	"com.apple.jabber.account":                               JabberAccount{},
	"com.apple.ldap.account":                                 LDAP{},
	"com.apple.loginitems.managed":                           LoginItemsManagedItems{},
	"com.apple.loginwindow":                                  LoginWindow{},
	"com.apple.lom":                                          LightsOutManagementLOM{},
	"com.apple.mail.managed":                                 Mail{},
	"com.apple.mcxMenuExtras":                                ManagedMenuExtras{},
	"com.apple.mcxloginscripts":                              LoginWindowScripts{},
	"com.apple.mcxprinting":                                  Printing{},
	"com.apple.mdm":                                          MDM{},
	"com.apple.mobiledevice.passwordpolicy":                  Passcode{},
	"com.apple.networkusagerules":                            NetworkUsageRules{},
	"com.apple.notificationsettings":                         Notifications{},
	"com.apple.osxserver.account":                            MacOSServerAccount{},
	"com.apple.preference.security":                          SecurityPreferences{},
	"com.apple.preference.users":                             UserPreferences{},
	"com.apple.profileRemovalPassword":                       ProfileRemovalPassword{},
	"com.apple.proxy.http.global":                            GlobalHTTPProxy{},
	"com.apple.relay.managed":                                Relay1{},
	"com.apple.screensaver":                                  Screensaver{},
	"com.apple.screensaver.user":                             ScreensaverUser{},
	"com.apple.secondactiveethernet.managed":                 Net8021XSecondActiveEthernet{},
	"com.apple.secondethernet.managed":                       Net8021XSecondEthernet{},
	"com.apple.security.FDERecoveryKeyEscrow":                FDERecoveryKeyEscrow{},
	"com.apple.security.FDERecoveryRedirect":                 FDERecoveryKeyRedirection{},
	"com.apple.security.acme":                                ACMECertificate{},
	"com.apple.security.certificatepreference":               CertificatePreference{},
	"com.apple.security.certificaterevocation":               CertificateRevocation{},
	"com.apple.security.certificatetransparency":             CertificateTransparency{},
	"com.apple.security.firewall":                            Firewall{},
	"com.apple.security.identitypreference":                  IdentityPreference{},
	"com.apple.security.pem":                                 CertificatePEM{},
	"com.apple.security.pkcs1":                               CertificatePKCS1{},
	"com.apple.security.pkcs12":                              CertificatePKCS12{},
	"com.apple.security.root":                                CertificateRoot{},
	"com.apple.security.scep":                                SCEP{},
	"com.apple.security.smartcard":                           SmartCard{},
	"com.apple.servicemanagement":                            ServiceManagementManagedLoginItems{},
	"com.apple.shareddeviceconfiguration":                    LockScreenMessage{},
	"com.apple.sso":                                          SingleSignOn{},
	"com.apple.subscribedcalendar.account":                   SubscribedCalendars{},
	"com.apple.syspolicy.kernel-extension-policy":            SystemPolicyKernelExtensions{},
	"com.apple.system-extension-policy":                      SystemExtensions{},
	"com.apple.system.logging":                               SystemLogging{},
	"com.apple.systemmigration":                              SystemMigration{},
	"com.apple.systempolicy.control":                         SystemPolicyControl{},
	"com.apple.systempolicy.managed":                         SystemPolicyManaged{},
	"com.apple.systempolicy.rule":                            SystemPolicyRule{},
	"com.apple.systempreferences":                            SystemPreferences{},
	"com.apple.systemuiserver":                               MediaManagementAllowedMedia{},
	"com.apple.thirdactiveethernet.managed":                  Net8021XThirdActiveEthernet{},
	"com.apple.thirdethernet.managed":                        Net8021XThirdEthernet{},
	"com.apple.tvremote":                                     TVRemote{},
	"com.apple.universalaccess":                              Accessibility{},
	"com.apple.vpn.managed":                                  VPN1{},
	"com.apple.vpn.managed.applayer":                         AppLayerVPN{},
	"com.apple.vpn.managed.appmapping":                       AppToAppLayerVPNMapping{},
	"com.apple.webClip.managed":                              WebClip{},
	"com.apple.webcontent-filter":                            WebContentFilter{},
	"com.apple.wifi.managed":                                 WiFi{},
	"com.apple.xsan":                                         Xsan{},
	"com.apple.xsan.preferences":                             XsanPreferences{},
	"loginwindow":                                            LoginWindowLoginItems{},
}

// The properties common to all payloads.
type CommonPayloadKeys struct {
	// The reverse-DNS-style identifier for the payload. This identifier is usually the same as the `TopLevel` value, with an additional appended component. This string must be unique within the profile.
	// During a profile replacement, the system updates payloads with the same `PayloadIdentifier` and `PayloadUUID` in the old and new profiles.
	PayloadIdentifier string `json:"PayloadIdentifier" plist:"PayloadIdentifier" required:"true"`
	// The globally unique identifier for the payload. The actual content is unimportant, but must be globally unique. In macOS, use `uuidgen` to generate UUIDs.
	// During a profile replacement, the system updates payloads with the same `PayloadIdentifier` and `PayloadUUID` in the old and new profiles.
	PayloadUUID string `json:"PayloadUUID" plist:"PayloadUUID" required:"true"`
	// The payload type, which each payload domain's reference page specifies.
	PayloadType string `json:"PayloadType" plist:"PayloadType" required:"true"`
	// The version of this specific payload.
	PayloadVersion CommonPayloadKeysPayloadVersion `json:"PayloadVersion" plist:"PayloadVersion" required:"true"`
	// The human-readable description of this payload. This description appears on the Detail screen.
	PayloadDescription *string `json:"PayloadDescription,omitempty" plist:"PayloadDescription,omitempty"`
	// The human-readable name for the profile payload. The name appears on the Detail screen and doesn't need to be unique.
	PayloadDisplayName *string `json:"PayloadDisplayName,omitempty" plist:"PayloadDisplayName,omitempty"`
	// The human-readable string containing the name of the organization that provides the profile. This value doesn't need to match the organization payload value in the enclosing dictionary.
	PayloadOrganization *string `json:"PayloadOrganization,omitempty" plist:"PayloadOrganization,omitempty"`
}

// The version of this specific payload.
type CommonPayloadKeysPayloadVersion int64

const (
	CommonPayloadKeysPayloadVersion1 CommonPayloadKeysPayloadVersion = 1
)

// The payload to configure global preferences.
// Global preferences on macOS
type GlobalPreferences struct {
	*CommonPayloadKeys
	// If `false`, disables fast user switching.
	MultipleSessionEnabled *bool `json:"MultipleSessionEnabled,omitempty" plist:"MultipleSessionEnabled,omitempty"`
	// The `autologout` delay, in seconds. A value of `0` means `autologout` is off. In some cases, this delay may be restricted to values between 5 minutes and 24 hours.
	ComappleautologoutAutoLogOutDelay *float64 `json:"com.apple.autologout.AutoLogOutDelay,omitempty" plist:"com.apple.autologout.AutoLogOutDelay,omitempty"`
}

func (p *GlobalPreferences) PayloadType() string {
	return ".GlobalPreferences"
}

// The top-level payload properties for all profiles.
type TopLevel struct {
	// The reverse-DNS style identifier (`com.example.myprofile`, for example) that identifies the profile. The system uses this string to determine whether to replace an existing profile or add it as a new profile.
	PayloadIdentifier string `json:"PayloadIdentifier" plist:"PayloadIdentifier" required:"true"`
	// The globally unique identifier for the profile. The actual content is unimportant. In macOS, you can use `uuidgen` to generate reasonable UUIDs.
	PayloadUUID string `json:"PayloadUUID" plist:"PayloadUUID" required:"true"`
	// The type of payload. The only supported value is `Configuration`.
	PayloadType PayloadType `json:"PayloadType" plist:"PayloadType" required:"true"`
	// The version number of the profile format, which needs to be `1`. This number represents the version of the configuration profile as a whole, not of the individual profiles within it.
	PayloadVersion TopLevelPayloadVersion `json:"PayloadVersion" plist:"PayloadVersion" required:"true"`
	// The array of payload dictionaries. If `IsEncrypted` is `true`, this array isn't needed.
	PayloadContent []ProfilePayload `json:"PayloadContent" plist:"PayloadContent" required:"true"`
	// Enabled if `IsEncrypted` is `true`.
	EncryptedPayloadContent *[]byte `json:"EncryptedPayloadContent,omitempty" plist:"EncryptedPayloadContent,omitempty"`
	// The description of the profile, shown on the Detail screen for the profile. Make this description detailed enough to help the user decide whether to install the profile.
	PayloadDescription *string `json:"PayloadDescription,omitempty" plist:"PayloadDescription,omitempty"`
	// The human-readable name for the profile, which doesn't need to be unique. The system displays this value on the Detail screen.
	PayloadDisplayName *string `json:"PayloadDisplayName,omitempty" plist:"PayloadDisplayName,omitempty"`
	// The human-readable string that contains the name of the organization that provided the profile.
	PayloadOrganization *string `json:"PayloadOrganization,omitempty" plist:"PayloadOrganization,omitempty"`
	// If present and set to `true`, the user can't delete the profile unless the profile has a removal password and the user provides it.
	// On macOS 10.15 and later, this key only affects removal of _manually_ installed profiles. If set to `true` and no profile removal payload is present, removing the profile requires admin auth.
	// On macOS versions prior to 10.15, this key prevents admins from removing MDM installed profiles. However, as of macOS 10.15, users can never remove MDM profiles, not even the admin.
	// On iOS users can't remove a MDM profile.
	// Requires a supervised device.
	PayloadRemovalDisallowed *bool `json:"PayloadRemovalDisallowed,omitempty" plist:"PayloadRemovalDisallowed,omitempty"`
	// A string that defines whether to install the profile for the system or the user. In many cases, it determines the location of certificate items, such as keychains. Though it's not possible to declare different payload scopes, payloads like VPN can automatically install their items in both scopes, if needed.
	PayloadScope *PayloadScope `json:"PayloadScope,omitempty" plist:"PayloadScope,omitempty"`
	// The date when the system automatically removes the profile.
	RemovalDate *time.Time `json:"RemovalDate,omitempty" plist:"RemovalDate,omitempty"`
	// The number of seconds until the profile is automatically removed. If the `RemovalDate` key is present, the system uses whichever field yields the earliest date.
	DurationUntilRemoval *float64 `json:"DurationUntilRemoval,omitempty" plist:"DurationUntilRemoval,omitempty"`
	// The date when a profile is no longer valid and the system presents an update button to the user.
	PayloadExpirationDate *time.Time `json:"PayloadExpirationDate,omitempty" plist:"PayloadExpirationDate,omitempty"`
	// The type of platform of the target device. Specifying the platform type helps prevent unintended installations.
	// For interactive installations on iOS devices, specifying a target platform avoids interstitial alerts that prompt the user to choose a profile target when multiple targets are eligible.
	// Allowed values:
	// - `0`: Any/unspecified
	// - `1`: iPhone/iPad/iPod Touch
	// - `2`: Apple Watch
	// - `3`: HomePod
	// - `4`: Apple TV
	// - `5`: Mac
	// - `6`: Vision Pro
	TargetDeviceType *TargetDeviceType `default:"0" json:"TargetDeviceType,omitempty" plist:"TargetDeviceType,omitempty"`
	// A dictionary that includes:
	// - A key that contains the IETF BCP 47 identifier for a language, such as _en_ or _jp_
	// - A value that contains the agreement localized to language specified by the key
	// The dictionary can also contain an optional key, `default`, with its value consisting of the unlocalized (usually in _en_) agreement.
	// The system always displays the agreement in a dialog, and the user needs to agree before the system can install the profile.
	// The system chooses a localized version in the order of preference that the user specifies in macOS, or based on the user's current language setting in iOS. If there's no exact match, the system uses the default localization. If there's no default localization, the system uses the _en_ localization. If there's no _en_ localization, the system uses the first available localization.
	// > Tip:
	// > Provide a default value, if possible. The system won't display a warning if the user's locale doesn't match any localization in the `ConsentText` dictionary.
	ConsentText *ConsentText `json:"ConsentText,omitempty" plist:"ConsentText,omitempty"`
}

// The type of payload. The only supported value is `Configuration`.
type PayloadType string

const (
	PayloadTypeConfiguration PayloadType = "Configuration"
)

// The version number of the profile format, which needs to be `1`. This number represents the version of the configuration profile as a whole, not of the individual profiles within it.
type TopLevelPayloadVersion int64

const (
	TopLevelPayloadVersion1 TopLevelPayloadVersion = 1
)

// A string that defines whether to install the profile for the system or the user. In many cases, it determines the location of certificate items, such as keychains. Though it's not possible to declare different payload scopes, payloads like VPN can automatically install their items in both scopes, if needed.
type PayloadScope string

const (
	PayloadScopeSystem PayloadScope = "System"
	PayloadScopeUser   PayloadScope = "User"
)

// The type of platform of the target device. Specifying the platform type helps prevent unintended installations.
// For interactive installations on iOS devices, specifying a target platform avoids interstitial alerts that prompt the user to choose a profile target when multiple targets are eligible.
// Allowed values:
// - `0`: Any/unspecified
// - `1`: iPhone/iPad/iPod Touch
// - `2`: Apple Watch
// - `3`: HomePod
// - `4`: Apple TV
// - `5`: Mac
// - `6`: Vision Pro
type TargetDeviceType int64

const (
	TargetDeviceType0 TargetDeviceType = 0
	TargetDeviceType1 TargetDeviceType = 1
	TargetDeviceType2 TargetDeviceType = 2
	TargetDeviceType3 TargetDeviceType = 3
	TargetDeviceType4 TargetDeviceType = 4
	TargetDeviceType5 TargetDeviceType = 5
	TargetDeviceType6 TargetDeviceType = 6
)

// A dictionary that includes:
// - A key that contains the IETF BCP 47 identifier for a language, such as _en_ or _jp_
// - A value that contains the agreement localized to language specified by the key
// The dictionary can also contain an optional key, `default`, with its value consisting of the unlocalized (usually in _en_) agreement.
// The system always displays the agreement in a dialog, and the user needs to agree before the system can install the profile.
// The system chooses a localized version in the order of preference that the user specifies in macOS, or based on the user's current language setting in iOS. If there's no exact match, the system uses the default localization. If there's no default localization, the system uses the _en_ localization. If there's no _en_ localization, the system uses the first available localization.
// > Tip:
// > Provide a default value, if possible. The system won't display a warning if the user's locale doesn't match any localization in the `ConsentText` dictionary.
type ConsentText struct {
	// The dictionary containing a key that consists of the IETF BCP 47 identifier for a language (for example, en or jp) and a value that consists of the agreement localized to that language.
	ConsentTextItem map[string]string `json:"ConsentTextItem" plist:"ConsentTextItem" required:"true"`
}

// The payload that configures Active Directory Certificate settings.
// A certificate can be requested from a Microsoft Certificate Authority (CA) using DCE/RPC and the Active Directory Certificate profile payload instructions detailed at support.apple.com/kb/HT5357.
type ActiveDirectoryCertificate struct {
	*CommonPayloadKeys
	// The fully qualified host name of the CA.
	CertServer string `json:"CertServer" plist:"CertServer" required:"true"`
	// The certificate template for your environment. The default user certificate value is \`User\`. The default computer certificate value is \`Machine\`.
	CertTemplate string `json:"CertTemplate" plist:"CertTemplate" required:"true"`
	// A user-friendly description of the certification identity.
	Description *string `json:"Description,omitempty" plist:"Description,omitempty"`
	// The number of days in advance of certificate expiration that the notification center notifies the user.
	CertificateRenewalTimeInterval *int64 `json:"CertificateRenewalTimeInterval,omitempty" plist:"CertificateRenewalTimeInterval,omitempty"`
	// The name of the certificate authority (CA), which is determined from the common name (CN) of the Active Directory entry. Available in macOS 10.8 and later. Valid values:
	// - CN=<your CA Name>
	// - CN=`Certification Authorities`
	// - CN=`Public Key Services`
	// - CN=`Services`
	// - CN=`Configuration`
	// - CN=<your base Domain Name>
	CertificateAuthority *string `json:"CertificateAuthority,omitempty" plist:"CertificateAuthority,omitempty"`
	// This value is most commonly `RPC`; if using web enrollment, use `HTTP`. Available in macOS 10.8 and later.
	CertificateAcquisitionMechanism *string `json:"CertificateAcquisitionMechanism,omitempty" plist:"CertificateAcquisitionMechanism,omitempty"`
	// If `true`, gives apps access to the private key. Available in macOS 10.10 and later.
	AllowAllAppsAccess *bool `json:"AllowAllAppsAccess,omitempty" plist:"AllowAllAppsAccess,omitempty"`
	// If `true`, the system prompts the user for credentials when is installs the profile. This key applies only to user certificates with the Manual Download profile delivery method. Omit this key for computer certificates. Available in macOS 10.8 and later.
	PromptForCredentials *bool `json:"PromptForCredentials,omitempty" plist:"PromptForCredentials,omitempty"`
	// If `true`, the system allows exporting the private key. Available in macOS 10.10 and later.
	KeyIsExtractable *bool `json:"KeyIsExtractable,omitempty" plist:"KeyIsExtractable,omitempty"`
	// The RSA key size for the certificate signing request (CSR). Available in macOS 10.11 and later.
	Keysize *int64 `default:"2048" json:"Keysize,omitempty" plist:"Keysize,omitempty"`
	// If `true`, the certificate obtained with this payload attempts auto-renewal. Auto-renewal can only be used with device Active Directory certificate payloads. Available in macOS 10.13.4 and later.
	EnableAutoRenewal *bool `json:"EnableAutoRenewal,omitempty" plist:"EnableAutoRenewal,omitempty"`
}

func (p *ActiveDirectoryCertificate) PayloadType() string {
	return "com.apple.ADCertificate.managed"
}

// The payload that configures an AIM account on the device.
// An AIM payload creates an AIM account on the device.
type AIMAccount struct {
	*CommonPayloadKeys
	// The description of the account.
	AIMAccountDescription *string `json:"AIMAccountDescription,omitempty" plist:"AIMAccountDescription,omitempty"`
	// The server address.
	AIMHostName AIMHostName `json:"AIMHostName" plist:"AIMHostName" required:"true"`
	// The user's login name.
	AIMUserName *string `json:"AIMUserName,omitempty" plist:"AIMUserName,omitempty"`
	// The user's password.
	AIMPassword *string `json:"AIMPassword,omitempty" plist:"AIMPassword,omitempty"`
	// If `true`, enables SSL.
	AIMUseSSL *bool `json:"AIMUseSSL,omitempty" plist:"AIMUseSSL,omitempty"`
	// The connection port for the server.
	AIMPort *int64 `default:"5190" json:"AIMPort,omitempty" plist:"AIMPort,omitempty"`
	// The authentication method for the account.
	AIMAuthentication AIMAuthentication `json:"AIMAuthentication" plist:"AIMAuthentication" required:"true"`
}

func (p *AIMAccount) PayloadType() string {
	return "com.apple.AIM.account"
}

// The server address.
type AIMHostName string

const (
	AIMHostNameSloginoscaraolcom AIMHostName = "slogin.oscar.aol.com"
)

// The authentication method for the account.
type AIMAuthentication string

const (
	AIMAuthenticationAIMAuthPassword AIMAuthentication = "AIMAuthPassword"
)

// The payload that configures the Content Caching service.
// Configures the Content Caching service.
type ContentCaching struct {
	*CommonPayloadKeys
	// If true, the system purges content from the cache automatically when it needs disk space for other apps when free disk space runs low on the computer. Set to `false` to maximize effectiveness of Content Caching. Available in macOS 10.15 and later.
	AllowCacheDelete *bool `json:"AllowCacheDelete,omitempty" plist:"AllowCacheDelete,omitempty"`
	// If `true`, the system caches the user's iCloud data. Changes to this value don't have an immediate effect. Clients may take some time, such as hours or days, to react to changes.
	// > Note:
	// > At least one of the `AllowPersonalCaching` or `AllowSharedCaching` keys need to be `true`.
	AllowPersonalCaching *bool `json:"AllowPersonalCaching,omitempty" plist:"AllowPersonalCaching,omitempty"`
	// If `true`, the system caches non-iCloud content, such as apps and software updates. Changes to this value don't have an immediate effect. Clients may take some time, such as hours or days, to react to changes.
	// > Note:
	// > At least one of the `AllowPersonalCaching` or `AllowSharedCaching` keys need to be `true`.
	AllowSharedCaching *bool `json:"AllowSharedCaching,omitempty" plist:"AllowSharedCaching,omitempty"`
	// If `true`, the system automatically activates the content cache when possible and prevents disabling it. If `allowContentCaching` is `false`, `AutoActivation` is also `false`.
	// Removing a profile that set `AutoActivation` to `true` doesn't deactivate the Content Cache.
	AutoActivation *bool `json:"AutoActivation,omitempty" plist:"AutoActivation,omitempty"`
	// If `true`, the system automatically enables Internet connection sharing when possible and prevent disabling Internet connection sharing. `DenyTetheredCaching` overrides `AutoEnableTetheredCaching`. Tethered caching requires Content Caching.
	// Available in macOS 10.15.4 and later.
	AutoEnableTetheredCaching *bool `json:"AutoEnableTetheredCaching,omitempty" plist:"AutoEnableTetheredCaching,omitempty"`
	// The maximum number of bytes of disk space to use for the content cache. Set to `0` for unlimited disk space.
	CacheLimit *int64 `default:"0" json:"CacheLimit,omitempty" plist:"CacheLimit,omitempty"`
	// The path to the directory used to store cached content. Changing this setting manually doesn't automatically move cached content from the old location to the new one. To move content automatically, use the Sharing preference's Content Caching pane. The value must be (or end with) `/Library/Application Support/Apple/AssetCache/Data`.
	// The system creates a directory and its intermediates for the given data path if it doesn't already exist. The directory is owned by `_assetcache:_assetcache` and has mode 0750. Its immediate parent directory (`.../Library/Application Support/Apple/AssetCache`) is owned by `_assetcache:_assetcache` and has mode `0755`.
	DataPath *string `default:"/Library/Application Support/Apple/AssetCache/Data" json:"DataPath,omitempty" plist:"DataPath,omitempty"`
	// If `true`, the system disables tethered caching.
	DenyTetheredCaching *bool `json:"DenyTetheredCaching,omitempty" plist:"DenyTetheredCaching,omitempty"`
	// If `true`, Content Caching displays exceptional conditions (alerts) as system notifications in the upper corner of the screen. Alerts were automatically displayed starting in macOS 10.13. In macOS 10.15 the alerts are off by default, but still available through this setting. Available in macOS 10.15 and later.
	DisplayAlerts *bool `json:"DisplayAlerts,omitempty" plist:"DisplayAlerts,omitempty"`
	// If `true`, the system prevents the computer from sleeping as long as Content Caching is on (System Preferences > Sharing > Content Caching is on). Customers who want Content Caching to be as available as much as possible should turn this setting on. Available in macOS 10.15 and later.
	KeepAwake *bool `json:"KeepAwake,omitempty" plist:"KeepAwake,omitempty"`
	// An array of dictionaries that describe a range of client IP addresses to serve.
	ListenRanges *[]*RangesItem `json:"ListenRanges,omitempty" plist:"ListenRanges,omitempty"`
	// If `true`, the content cache provides content to the clients in the `ListenRanges`.
	ListenRangesOnly *bool `json:"ListenRangesOnly,omitempty" plist:"ListenRangesOnly,omitempty"`
	// If `true`, the content cache provides content to the clients in the union of the `ListenRanges`, `PeerListenRanges` and `Parents`.
	ListenWithPeersAndParents *bool `json:"ListenWithPeersAndParents,omitempty" plist:"ListenWithPeersAndParents,omitempty"`
	// If `true`, the content cache offers content to clients only on the same immediate local network only. No content is offered to clients on other networks reachable by the content cache. If `LocalSubnetsOnly` is `true`, the system ignores `ListenRanges`.
	LocalSubnetsOnly *bool `json:"LocalSubnetsOnly,omitempty" plist:"LocalSubnetsOnly,omitempty"`
	// If `true`, the Content Cache logs the IP address and port number of the clients that request content.
	LogClientIdentity *bool `json:"LogClientIdentity,omitempty" plist:"LogClientIdentity,omitempty"`
	// An array of the local IP addresses of other content caches that this cache should download from or upload to, instead of downloading from or uploading to Apple directly. The system ignores invalid addresses and addresses of computers that aren't content caches. The system skips Parent caches that become unavailable. If all parent content caches become unavailable, the content cache downloads from or uploads to Apple directly, until a parent content cache becomes available again.
	Parents *[]string `json:"Parents,omitempty" plist:"Parents,omitempty"`
	// The policy to implement when choosing among more than one configured parent content cache. With every policy, the system skips parent caches that are temporarily unavailable. Allowed values:
	// - `first-available`: Always use the first available parent in the Parents list. Use this policy to designate permanent primary, secondary, and subsequent parents.
	// - `url-path-hash`: Hash the path part of the requested URL so that the same parent is always used for the same URL. This is useful for maximizing the size of the combined caches of     the parents.
	// - `random`: Choose a parent at random. Use this policy for load balancing.
	// - `round-robin`: Rotate through the parents in order. Use this policy for load balancing.
	// - `sticky-available`: Use the first available parent in the Parents list until it becomes unavailable, then advance to the next one. Use this policy for designating floating primary, secondary, and subsequent parents.
	ParentSelectionPolicy *ParentSelectionPolicy `default:"round-robin" json:"ParentSelectionPolicy,omitempty" plist:"ParentSelectionPolicy,omitempty"`
	// An array of dictionaries describing a range of peer IP addresses that the content cache uses to filter its list of peers to query for content. The content cache only queries peers in `PeerFilterRanges`. When `PeerFilterRanges` is an empty array, the content cache doesn't query any peers.
	PeerFilterRanges *[]*RangesItem `json:"PeerFilterRanges,omitempty" plist:"PeerFilterRanges,omitempty"`
	// An array of dictionaries describing a range of peer IP addresses the content cache responds to. When `PeerListenRanges` is an empty array, the content cache responds with an error to all cache queries.
	PeerListenRanges *[]*RangesItem `json:"PeerListenRanges,omitempty" plist:"PeerListenRanges,omitempty"`
	// If `true`, the content cache only peers with other content caches on the same immediate local network, rather than with content caches that use the same public IP address as the device. When `PeerLocalSubnetsOnly` is `true`, it overrides the configuration of `PeerFilterRanges` and `PeerListenRanges`. If the network changes, the local network peering restrictions update appropriately. If `false`, the content cache defers to `PeerFilterRanges` and `PeerListenRanges` for configuring the peering restrictions.
	PeerLocalSubnetsOnly *bool `json:"PeerLocalSubnetsOnly,omitempty" plist:"PeerLocalSubnetsOnly,omitempty"`
	// The TCP port number on which the content cache accepts requests for uploads or downloads. Set to `0` to pick a random, available port.
	Port *int64 `default:"0" json:"Port,omitempty" plist:"Port,omitempty"`
	// An array of dictionaries describing a range of public IP addresses that the cloud servers should use for matching clients to content caches.
	PublicRanges *[]*RangesItem `json:"PublicRanges,omitempty" plist:"PublicRanges,omitempty"`
}

func (p *ContentCaching) PayloadType() string {
	return "com.apple.AssetCache.managed"
}

// A range of IP addresses to cache.
type RangesItem struct {
	// The IP address type.
	Type *RangesItemType `default:"IPv4" json:"type,omitempty" plist:"type,omitempty"`
	// The first IP address in the range.
	First string `json:"first" plist:"first" required:"true"`
	// The last IP address in the range.
	Last string `json:"last" plist:"last" required:"true"`
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
	ParentSelectionPolicyUrlPathHash     ParentSelectionPolicy = "url-path-hash"
	ParentSelectionPolicyRandom          ParentSelectionPolicy = "random"
	ParentSelectionPolicyRoundRobin      ParentSelectionPolicy = "round-robin"
	ParentSelectionPolicyStickyAvailable ParentSelectionPolicy = "sticky-available"
)

// The payload that configures parental control dictionary restrictions.
// Parental controls dictionary restrictions.
type ParentalControlsDictionary struct {
	*CommonPayloadKeys
	// If `true`, enables parental controls dictionary restrictions.
	ParentalControl bool `json:"parentalControl" plist:"parentalControl" required:"true"`
}

func (p *ParentalControlsDictionary) PayloadType() string {
	return "com.apple.Dictionary"
}

// The payload that configures an Active Directory (AD) domain.
// In macOS 10.9 and later, a configuration profile can be used to configure macOS to join an Active Directory (AD) domain. Advanced AD options available via Directory Utility or the dsconfigad command line tool can also be set using a configuration profile.
type DirectoryService struct {
	*CommonPayloadKeys
	// The Active Directory domain to join.
	HostName string `json:"HostName" plist:"HostName" required:"true"`
	// The user name of the account for the domain.
	UserName *string `json:"UserName,omitempty" plist:"UserName,omitempty"`
	// The password of the account for the domain.
	Password *string `json:"Password,omitempty" plist:"Password,omitempty"`
	// The client's identifier.
	ClientID *string `json:"ClientID,omitempty" plist:"ClientID,omitempty"`
	// The directory service description.
	Description *string `json:"Description,omitempty" plist:"Description,omitempty"`
	// The organizational unit to add the joining computer object to.
	ADOrganizationalUnit *string `json:"ADOrganizationalUnit,omitempty" plist:"ADOrganizationalUnit,omitempty"`
	// The network home protocol to use: `afp` or `smb`.
	ADMountStyle *string `json:"ADMountStyle,omitempty" plist:"ADMountStyle,omitempty"`
	// If `true`, the system enables the `ADCreateMobileAccountAtLogin` key.
	ADCreateMobileAccountAtLoginFlag *bool `json:"ADCreateMobileAccountAtLoginFlag,omitempty" plist:"ADCreateMobileAccountAtLoginFlag,omitempty"`
	// If `true`, the system creates a mobile account at login.
	ADCreateMobileAccountAtLogin *bool `json:"ADCreateMobileAccountAtLogin,omitempty" plist:"ADCreateMobileAccountAtLogin,omitempty"`
	// If `true`, the system enables the `ADWarnUserBeforeCreatingMA` key.
	ADWarnUserBeforeCreatingMAFlag *bool `json:"ADWarnUserBeforeCreatingMAFlag,omitempty" plist:"ADWarnUserBeforeCreatingMAFlag,omitempty"`
	// If `true`, the system enables the warning before creating the mobile account.
	ADWarnUserBeforeCreatingMA *bool `json:"ADWarnUserBeforeCreatingMA,omitempty" plist:"ADWarnUserBeforeCreatingMA,omitempty"`
	// If `true`, the system enables the `ADForceHomeLocal` key.
	ADForceHomeLocalFlag *bool `json:"ADForceHomeLocalFlag,omitempty" plist:"ADForceHomeLocalFlag,omitempty"`
	// If `true`, the system forces a local home directory.
	ADForceHomeLocal *bool `json:"ADForceHomeLocal,omitempty" plist:"ADForceHomeLocal,omitempty"`
	// If `true`, the system enables the `ADUseWindowsUNCPath` key.
	ADUseWindowsUNCPathFlag *bool `json:"ADUseWindowsUNCPathFlag,omitempty" plist:"ADUseWindowsUNCPathFlag,omitempty"`
	// If `true`, the system uses the UNC path from Active Directory to derive the network home location.
	ADUseWindowsUNCPath *bool `json:"ADUseWindowsUNCPath,omitempty" plist:"ADUseWindowsUNCPath,omitempty"`
	// If `true`, the system enables the `ADAllowMultiDomainAuth` key.
	ADAllowMultiDomainAuthFlag *bool `json:"ADAllowMultiDomainAuthFlag,omitempty" plist:"ADAllowMultiDomainAuthFlag,omitempty"`
	// If `true`, the system allows authentication from any domain in the namespace.
	ADAllowMultiDomainAuth *bool `json:"ADAllowMultiDomainAuth,omitempty" plist:"ADAllowMultiDomainAuth,omitempty"`
	// If `true`, the system enables the `ADDefaultUserShell` key.
	ADDefaultUserShellFlag *bool `json:"ADDefaultUserShellFlag,omitempty" plist:"ADDefaultUserShellFlag,omitempty"`
	// The default user shell.
	ADDefaultUserShell *string `json:"ADDefaultUserShell,omitempty" plist:"ADDefaultUserShell,omitempty"`
	// If `true`, the system enables the `ADMapUIDAttribute` key.
	ADMapUIDAttributeFlag *bool `json:"ADMapUIDAttributeFlag,omitempty" plist:"ADMapUIDAttributeFlag,omitempty"`
	// The map UID to attribute.
	ADMapUIDAttribute *string `json:"ADMapUIDAttribute,omitempty" plist:"ADMapUIDAttribute,omitempty"`
	// If `true`, the system enables the `ADMapGIDAttribute` key.
	ADMapGIDAttributeFlag *bool `json:"ADMapGIDAttributeFlag,omitempty" plist:"ADMapGIDAttributeFlag,omitempty"`
	// The map GID to attribute.
	ADMapGIDAttribute *string `json:"ADMapGIDAttribute,omitempty" plist:"ADMapGIDAttribute,omitempty"`
	// If `true`, the system enables the `ADMapGGIDAttributeFlag` key.
	ADMapGGIDAttributeFlag *bool `json:"ADMapGGIDAttributeFlag,omitempty" plist:"ADMapGGIDAttributeFlag,omitempty"`
	// The map group GID to attribute.
	ADMapGGIDAttribute *string `json:"ADMapGGIDAttribute,omitempty" plist:"ADMapGGIDAttribute,omitempty"`
	// If `true`, the system enables the `ADPreferredDCServer` key.
	ADPreferredDCServerFlag *bool `json:"ADPreferredDCServerFlag,omitempty" plist:"ADPreferredDCServerFlag,omitempty"`
	// The preferred domain server.
	ADPreferredDCServer *string `json:"ADPreferredDCServer,omitempty" plist:"ADPreferredDCServer,omitempty"`
	// If `true`, the system enables the `ADDomainAdminGroupList` key.
	ADDomainAdminGroupListFlag *bool `json:"ADDomainAdminGroupListFlag,omitempty" plist:"ADDomainAdminGroupListFlag,omitempty"`
	// The list of Active Directory groups with admin access.
	ADDomainAdminGroupList *[]string `json:"ADDomainAdminGroupList,omitempty" plist:"ADDomainAdminGroupList,omitempty"`
	// If `true`, the system enables the `ADNamespace` key.
	ADNamespaceFlag *bool `json:"ADNamespaceFlag,omitempty" plist:"ADNamespaceFlag,omitempty"`
	// The primary user account naming convention; either `forest` or `domain`.
	ADNamespace *string `json:"ADNamespace,omitempty" plist:"ADNamespace,omitempty"`
	// If `true`, the system enables the `ADPacketSign` key.
	ADPacketSignFlag *bool `json:"ADPacketSignFlag,omitempty" plist:"ADPacketSignFlag,omitempty"`
	// The packet signing policy.
	ADPacketSign *string `json:"ADPacketSign,omitempty" plist:"ADPacketSign,omitempty"`
	// If `true`, the system enables the `ADPacketEncrypt` key.
	ADPacketEncryptFlag *bool `json:"ADPacketEncryptFlag,omitempty" plist:"ADPacketEncryptFlag,omitempty"`
	// The packet encryption policy.
	ADPacketEncrypt *string `json:"ADPacketEncrypt,omitempty" plist:"ADPacketEncrypt,omitempty"`
	// If `true`, the system enables the `ADRestrictDDNS` key.
	ADRestrictDDNSFlag *bool `json:"ADRestrictDDNSFlag,omitempty" plist:"ADRestrictDDNSFlag,omitempty"`
	// An array of strings that represent the interfaces allowed for dynamic DNS updates, such as en0 and en1.
	ADRestrictDDNS *[]string `json:"ADRestrictDDNS,omitempty" plist:"ADRestrictDDNS,omitempty"`
	// If `true`, the system enables the `ADTrustChangePassIntervalDays` key.
	ADTrustChangePassIntervalDaysFlag *bool `json:"ADTrustChangePassIntervalDaysFlag,omitempty" plist:"ADTrustChangePassIntervalDaysFlag,omitempty"`
	// The number of days before requiring a change of the computer trust account password. Set to `0` to disable the feature.
	ADTrustChangePassIntervalDays *int64 `json:"ADTrustChangePassIntervalDays,omitempty" plist:"ADTrustChangePassIntervalDays,omitempty"`
}

func (p *DirectoryService) PayloadType() string {
	return "com.apple.DirectoryService.managed"
}

// The payload that configures disc-burning settings.
type MediaManagementDiscBurning struct {
	*CommonPayloadKeys
	// Configure disc-burn. Allowed values:
	// - `off`: The system disables disc burning.
	// - `on`: The system allows normal default operation. Setting this key to `on` doesn't enable disc burn support if other mechanisms or preferences disabled it. Needs to be enabled with the `Finder` profile.
	// - `authenticate`: The system requires authentication.
	BurnSupport BurnSupport `json:"BurnSupport" plist:"BurnSupport" required:"true"`
}

func (p *MediaManagementDiscBurning) PayloadType() string {
	return "com.apple.DiscRecording"
}

// Configure disc-burn. Allowed values:
// - `off`: The system disables disc burning.
// - `on`: The system allows normal default operation. Setting this key to `on` doesn't enable disc burn support if other mechanisms or preferences disabled it. Needs to be enabled with the `Finder` profile.
// - `authenticate`: The system requires authentication.
type BurnSupport string

const (
	BurnSupportOff          BurnSupport = "off"
	BurnSupportAuthenticate BurnSupport = "authenticate"
	BurnSupportOn           BurnSupport = "on"
)

// The payload that configures guest accounts.
type Accounts struct {
	*CommonPayloadKeys
	// If `true`, the system enables the guest account.
	EnableGuestAccount *bool `json:"EnableGuestAccount,omitempty" plist:"EnableGuestAccount,omitempty"`
	// If `true`, the system disables the guest account. This property has no effect if `EnableGuestAccount` is `true`.
	DisableGuestAccount *bool `json:"DisableGuestAccount,omitempty" plist:"DisableGuestAccount,omitempty"`
}

func (p *Accounts) PayloadType() string {
	return "com.apple.MCX"
}

// The payload that configures Energy Saver settings.
type EnergySaver struct {
	*CommonPayloadKeys
	// The settings for a desktop computer.
	ComappleEnergySaverdesktopACPower *ComappleEnergySaverdesktopACPower `json:"com.apple.EnergySaver.desktop.ACPower,omitempty" plist:"com.apple.EnergySaver.desktop.ACPower,omitempty"`
	// The settings for a laptop computer using AC power.
	ComappleEnergySaverportableACPower *ComappleEnergySaverportableACPower `json:"com.apple.EnergySaver.portable.ACPower,omitempty" plist:"com.apple.EnergySaver.portable.ACPower,omitempty"`
	// The settings for a laptop computer using battery power.
	ComappleEnergySaverportableBatteryPower *ComappleEnergySaverportableBatteryPower `json:"com.apple.EnergySaver.portable.BatteryPower,omitempty" plist:"com.apple.EnergySaver.portable.BatteryPower,omitempty"`
	// The schedule for turning a computer on and off.
	ComappleEnergySaverdesktopSchedule *ComappleEnergySaverdesktopSchedule `json:"com.apple.EnergySaver.desktop.Schedule,omitempty" plist:"com.apple.EnergySaver.desktop.Schedule,omitempty"`
	// If `true`, disables sleep.
	SleepDisabled *bool `json:"SleepDisabled,omitempty" plist:"SleepDisabled,omitempty"`
}

func (p *EnergySaver) PayloadType() string {
	return "com.apple.MCX"
}

// The settings for a desktop computer.
type ComappleEnergySaverdesktopACPower struct {
	// The display sleep time, in minutes. A value of 0 means never.
	DisplaySleepTimer *int64 `json:"Display Sleep Timer,omitempty" plist:"Display Sleep Timer,omitempty"`
	// The disk sleep time, in minutes. A value of 0 means never.
	DiskSleepTimer *int64 `json:"Disk Sleep Timer,omitempty" plist:"Disk Sleep Timer,omitempty"`
	// System sleep time, in minutes. A value of 0 means never.
	SystemSleepTimer *int64 `json:"System Sleep Timer,omitempty" plist:"System Sleep Timer,omitempty"`
	// May not be available on all systems.
	ReduceProcessorSpeed *ReduceProcessorSpeed `json:"Reduce Processor Speed,omitempty" plist:"Reduce Processor Speed,omitempty"`
	// May not be available on all systems.
	DynamicPowerStep *DynamicPowerStep `json:"Dynamic Power Step,omitempty" plist:"Dynamic Power Step,omitempty"`
	// If `true`, enables "Wake for network access."
	WakeonLAN *WakeonLAN `json:"Wake on LAN,omitempty" plist:"Wake on LAN,omitempty"`
	// If `true`, enables "Wake for modem ring."
	WakeOnModemRing *WakeOnModemRing `json:"Wake On Modem Ring,omitempty" plist:"Wake On Modem Ring,omitempty"`
	// If `true`, enables "Start up automatically after a power failure."
	AutomaticRestartOnPowerLoss *AutomaticRestartOnPowerLoss `json:"Automatic Restart On Power Loss,omitempty" plist:"Automatic Restart On Power Loss,omitempty"`
}

// May not be available on all systems.
type ReduceProcessorSpeed int64

const (
	ReduceProcessorSpeed0 ReduceProcessorSpeed = 0
	ReduceProcessorSpeed1 ReduceProcessorSpeed = 1
)

// May not be available on all systems.
type DynamicPowerStep int64

const (
	DynamicPowerStep0 DynamicPowerStep = 0
	DynamicPowerStep1 DynamicPowerStep = 1
)

// If `true`, enables "Wake for network access."
type WakeonLAN int64

const (
	WakeonLAN0 WakeonLAN = 0
	WakeonLAN1 WakeonLAN = 1
)

// If `true`, enables "Wake for modem ring."
type WakeOnModemRing int64

const (
	WakeOnModemRing0 WakeOnModemRing = 0
	WakeOnModemRing1 WakeOnModemRing = 1
)

// If `true`, enables "Start up automatically after a power failure."
type AutomaticRestartOnPowerLoss int64

const (
	AutomaticRestartOnPowerLoss0 AutomaticRestartOnPowerLoss = 0
	AutomaticRestartOnPowerLoss1 AutomaticRestartOnPowerLoss = 1
)

// The settings for a laptop computer using AC power.
type ComappleEnergySaverportableACPower struct {
	// The display sleep time, in minutes. A value of 0 means never.
	DisplaySleepTimer *int64 `json:"Display Sleep Timer,omitempty" plist:"Display Sleep Timer,omitempty"`
	// The disk sleep time, in minutes. A value of 0 means never.
	DiskSleepTimer *int64 `json:"Disk Sleep Timer,omitempty" plist:"Disk Sleep Timer,omitempty"`
	// System sleep time, in minutes. A value of 0 means never.
	SystemSleepTimer *int64 `json:"System Sleep Timer,omitempty" plist:"System Sleep Timer,omitempty"`
	// May not be available on all systems.
	ReduceProcessorSpeed *ReduceProcessorSpeed `json:"Reduce Processor Speed,omitempty" plist:"Reduce Processor Speed,omitempty"`
	// May not be available on all systems.
	DynamicPowerStep *DynamicPowerStep `json:"Dynamic Power Step,omitempty" plist:"Dynamic Power Step,omitempty"`
	// If `true`, enables "Wake for network access."
	WakeonLAN *WakeonLAN `json:"Wake on LAN,omitempty" plist:"Wake on LAN,omitempty"`
	// If `true`, enables "Wake for modem ring."
	WakeOnModemRing *WakeOnModemRing `json:"Wake On Modem Ring,omitempty" plist:"Wake On Modem Ring,omitempty"`
	// If `true`, enables "Start up automatically after a power failure."
	AutomaticRestartOnPowerLoss *AutomaticRestartOnPowerLoss `json:"Automatic Restart On Power Loss,omitempty" plist:"Automatic Restart On Power Loss,omitempty"`
}

// The settings for a laptop computer using battery power.
type ComappleEnergySaverportableBatteryPower struct {
	// The display sleep time, in minutes. A value of 0 means never.
	DisplaySleepTimer *int64 `json:"Display Sleep Timer,omitempty" plist:"Display Sleep Timer,omitempty"`
	// The disk sleep time, in minutes. A value of 0 means never.
	DiskSleepTimer *int64 `json:"Disk Sleep Timer,omitempty" plist:"Disk Sleep Timer,omitempty"`
	// System sleep time, in minutes. A value of 0 means never.
	SystemSleepTimer *int64 `json:"System Sleep Timer,omitempty" plist:"System Sleep Timer,omitempty"`
	// May not be available on all systems.
	ReduceProcessorSpeed *ReduceProcessorSpeed `json:"Reduce Processor Speed,omitempty" plist:"Reduce Processor Speed,omitempty"`
	// May not be available on all systems.
	DynamicPowerStep *DynamicPowerStep `json:"Dynamic Power Step,omitempty" plist:"Dynamic Power Step,omitempty"`
	// If `true`, enables "Wake for network access."
	WakeonLAN *WakeonLAN `json:"Wake on LAN,omitempty" plist:"Wake on LAN,omitempty"`
	// If `true`, enables "Wake for modem ring."
	WakeOnModemRing *WakeOnModemRing `json:"Wake On Modem Ring,omitempty" plist:"Wake On Modem Ring,omitempty"`
	// If `true`, enables "Start up automatically after a power failure."
	AutomaticRestartOnPowerLoss *AutomaticRestartOnPowerLoss `json:"Automatic Restart On Power Loss,omitempty" plist:"Automatic Restart On Power Loss,omitempty"`
}

// The schedule for turning a computer on and off.
type ComappleEnergySaverdesktopSchedule struct {
	// The schedule for turning the device on.
	RepeatingPowerOn *RepeatingPowerOn `json:"RepeatingPowerOn,omitempty" plist:"RepeatingPowerOn,omitempty"`
	// The schedule for turning the device off.
	RepeatingPowerOff *RepeatingPowerOff `json:"RepeatingPowerOff,omitempty" plist:"RepeatingPowerOff,omitempty"`
}

// The schedule for turning the device on.
type RepeatingPowerOn struct {
	// The type of action defined by this schedule.
	Eventtype Eventtype `json:"eventtype" plist:"eventtype" required:"true"`
	// One or more days of the week in an unsigned integer bitmap:
	// - `1` = Mon
	// - `2` = Tue
	// - `4` = Wed
	// - `8` = Thu
	// - `16` = Fri
	// - `32` = Sat
	// - `64` = Sun
	Weekdays *int64 `json:"weekdays,omitempty" plist:"weekdays,omitempty"`
	// The time, in minutes, since midnight.
	Time *int64 `json:"time,omitempty" plist:"time,omitempty"`
}

// The type of action defined by this schedule.
type Eventtype string

const (
	EventtypeWake        Eventtype = "wake"
	EventtypePoweron     Eventtype = "poweron"
	EventtypeWakepoweron Eventtype = "wakepoweron"
	EventtypeSleep       Eventtype = "sleep"
	EventtypeShutdown    Eventtype = "shutdown"
	EventtypeRestart     Eventtype = "restart"
)

// The schedule for turning the device off.
type RepeatingPowerOff struct {
	// The type of action defined by this schedule.
	Eventtype Eventtype `json:"eventtype" plist:"eventtype" required:"true"`
	// One or more days of the week in an unsigned integer bitmap:
	// - `1` = Mon
	// - `2` = Tue
	// - `4` = Wed
	// - `8` = Thu
	// - `16` = Fri
	// - `32` = Sat
	// - `64` = Sun
	Weekdays *int64 `json:"weekdays,omitempty" plist:"weekdays,omitempty"`
	// The time, in minutes, since midnight.
	Time *int64 `json:"time,omitempty" plist:"time,omitempty"`
}

// The payload that configures FileVault options.
// The FileVault accounts payload sets up options for enabling FileVault.
type FDEFileVaultOptions struct {
	*CommonPayloadKeys
	// If `true`, the system won't disable FileVault.
	DontAllowFDEDisable *bool `json:"dontAllowFDEDisable,omitempty" plist:"dontAllowFDEDisable,omitempty"`
	// If `true`, the system won't enable FileVault.
	DontAllowFDEEnable *bool `json:"dontAllowFDEEnable,omitempty" plist:"dontAllowFDEEnable,omitempty"`
	// If `true`, the system won't store th FileVault key across restarts.
	DestroyFVKeyOnStandby *bool `json:"DestroyFVKeyOnStandby,omitempty" plist:"DestroyFVKeyOnStandby,omitempty"`
}

func (p *FDEFileVaultOptions) PayloadType() string {
	return "com.apple.MCX"
}

// The payload that configures mobile accounts on the device.
// Sets up mobile account options for network based user accounts.
type MobileAccounts struct {
	*CommonPayloadKeys
	// If `true`, the system creates the mobile account at login time.
	ComapplecachedaccountsCreateAtLogin *bool `json:"com.apple.cachedaccounts.CreateAtLogin,omitempty" plist:"com.apple.cachedaccounts.CreateAtLogin,omitempty"`
	// If `true`, the system asks the user whether to create the mobile account and it allows the user to not create it.
	ComapplecachedaccountsWarnOnCreate *bool `json:"com.apple.cachedaccounts.WarnOnCreate,omitempty" plist:"com.apple.cachedaccounts.WarnOnCreate,omitempty"`
	// If `true`, the system allows the user to stop the prompts about mobile account creation every time the user logs in. This key is only valid if `com.apple.cachedaccounts.WarnOnCreate` is `true`.
	CachedaccountsWarnOnCreateallowNever *bool `json:"cachedaccounts.WarnOnCreate.allowNever,omitempty" plist:"cachedaccounts.WarnOnCreate.allowNever,omitempty"`
	// The minimum number of seconds a mobile account can exist before the system makes an automatic attempt to remove the mobile account. Set to `0` to attempt removing it at the next login or logout. Set to `-1` to never attempt removing the mobile account.
	CachedaccountsexpirydeletedisusedSeconds *int64 `default:"-1" json:"cachedaccounts.expiry.delete.disusedSeconds,omitempty" plist:"cachedaccounts.expiry.delete.disusedSeconds,omitempty"`
	// If `true`, the system bypasses the secure token authorization dialog. This dialog only appears on APFS volumes.
	CachedaccountsaskForSecureTokenAuthBypass *bool `json:"cachedaccounts.askForSecureTokenAuthBypass,omitempty" plist:"cachedaccounts.askForSecureTokenAuthBypass,omitempty"`
}

func (p *MobileAccounts) PayloadType() string {
	return "com.apple.MCX"
}

// The payload that configures the time server.
// Settings for time zone and server. If multiple profiles with this payload are sent, the device's time server will be set to the value in the last payload installed. Removing the payload will not change the settings back to the prior settings.
type TimeServer struct {
	*CommonPayloadKeys
	// The NTP server to connect to. In macOS 10.13 and later, only one time server is supported.
	TimeServer *string `json:"timeServer,omitempty" plist:"timeServer,omitempty"`
	// The time zone path location string in `/usr/share/zoneinfo/`; for example, `America/Denver` or `Zulu`.
	TimeZone *string `json:"timeZone,omitempty" plist:"timeZone,omitempty"`
}

func (p *TimeServer) PayloadType() string {
	return "com.apple.MCX"
}

// The payload that configures managed Wi-Fi settings.
type WiFiManagedSettings struct {
	*CommonPayloadKeys
	// If `true`, requires administrator authorization to enable IBSS.
	RequireAdminForIBSS *bool `json:"RequireAdminForIBSS,omitempty" plist:"RequireAdminForIBSS,omitempty"`
	// If `true`, requires administrator authorization for network changes.
	RequireAdminForAirPortNetworkChange *bool `json:"RequireAdminForAirPortNetworkChange,omitempty" plist:"RequireAdminForAirPortNetworkChange,omitempty"`
	// If `true`, requires administrator authorization to turn Wi-Fi on or off.
	RequireAdminToTurnAirPortOnOff *bool `json:"RequireAdminToTurnAirPortOnOff,omitempty" plist:"RequireAdminToTurnAirPortOnOff,omitempty"`
}

func (p *WiFiManagedSettings) PayloadType() string {
	return "com.apple.MCX"
}

// The payload that configures FileVault.
// The FileVault payload only works on macOS to enable or disable FileVault. Starting with macOS 10.15, this payload requires UAMDM to enable FileVault.
type FDEFileVault struct {
	*CommonPayloadKeys
	// Set to `On` to enable FileVault and set to `Off` to disable FileVault. Payloads set to `On` sent through MDM need to either include full authentication information in the payload or have the `Defer` option set to `true`. When `Defer` is `true`, the system prompts for the authentication information when the user enables FileVault.
	Enable Enable `json:"Enable" plist:"Enable" required:"true"`
	// If `true`, the system defers enabling FileVault until the designated user logs out. For details, see `fdesetup(8)`. Only a local user or a mobile account user can enable FileVault.
	Defer *bool `json:"Defer,omitempty" plist:"Defer,omitempty"`
	// If `true`, the system enables a prompt for missing user name or password fields.
	UserEntersMissingInfo *bool `json:"UserEntersMissingInfo,omitempty" plist:"UserEntersMissingInfo,omitempty"`
	// If `true`, the system creates a personal recovery key and displays it to the user.
	UseRecoveryKey *bool `json:"UseRecoveryKey,omitempty" plist:"UseRecoveryKey,omitempty"`
	// If `false`, the system prevents display of the personal recovery key to the user after the system enables FileVault.
	ShowRecoveryKey *bool `json:"ShowRecoveryKey,omitempty" plist:"ShowRecoveryKey,omitempty"`
	// The path to the location of the recovery key and computer information property list.
	OutputPath *string `json:"OutputPath,omitempty" plist:"OutputPath,omitempty"`
	// The DER-encoded certificate data if the system creates an institutional recovery key. This key isn't supported on a Mac with Apple silicon.
	Certificate *[]byte `json:"Certificate,omitempty" plist:"Certificate,omitempty"`
	// The UUID of the payload within the same profile containing the asymmetric recovery key certificate payload.
	PayloadCertificateUUID *string `json:"PayloadCertificateUUID,omitempty" plist:"PayloadCertificateUUID,omitempty"`
	// The user name of the Open Directory user to add to FileVault.
	Username *string `json:"Username,omitempty" plist:"Username,omitempty"`
	// The password of the Open Directory user to add to FileVault. Use the `UserEntersMissingInfo` key to prompt for this information.
	Password *string `json:"Password,omitempty" plist:"Password,omitempty"`
	// If `true` and you don't include certificate information in this payload, the system uses the keychain created at `/Library/Keychains/FileVaultMaster.keychain` when it adds the institutional recovery key.
	UseKeychain *bool `json:"UseKeychain,omitempty" plist:"UseKeychain,omitempty"`
	// The maximum number of times users can bypass enabling FileVault before the system requires the user to enable it to log in. If the value is `0`, the system requires the user to enable FileVault the next time they attempt to log in. Set this key to `-1` to disable this feature.
	DeferForceAtUserLoginMaxBypassAttempts *int64 `json:"DeferForceAtUserLoginMaxBypassAttempts,omitempty" plist:"DeferForceAtUserLoginMaxBypassAttempts,omitempty"`
	// If `true`, the system prevents requests to enable FileVault at user logout time.
	DeferDontAskAtUserLogout *bool `json:"DeferDontAskAtUserLogout,omitempty" plist:"DeferDontAskAtUserLogout,omitempty"`
	// If `true`, and installation of this payload occurs after enrolling with MDM in Setup Assistant, the system requests Setup Assistant to enable FileVault at setup time.
	// To use this, enable the Await Device Configured DEP configuration option and send this profile with this key set, before sending the `DeviceConfiguredCommand`.
	// An admin SecureToken user is required, otherwise the FileVault pane does not appear.
	ForceEnableInSetupAssistant *bool `json:"ForceEnableInSetupAssistant,omitempty" plist:"ForceEnableInSetupAssistant,omitempty"`
}

func (p *FDEFileVault) PayloadType() string {
	return "com.apple.MCX.FileVault2"
}

// Set to `On` to enable FileVault and set to `Off` to disable FileVault. Payloads set to `On` sent through MDM need to either include full authentication information in the payload or have the `Defer` option set to `true`. When `Defer` is `true`, the system prompts for the authentication information when the user enables FileVault.
type Enable string

const (
	EnableOn  Enable = "On"
	EnableOff Enable = "Off"
)

// The payload that configures Time Machine.
type TimeMachine struct {
	*CommonPayloadKeys
	// If `true`, performs automatic backups at regular intervals.
	AutoBackup *bool `json:"AutoBackup,omitempty" plist:"AutoBackup,omitempty"`
	// If `true`, backs up only the startup volume by default.
	BackupAllVolumes *bool `json:"BackupAllVolumes,omitempty" plist:"BackupAllVolumes,omitempty"`
	// The URL of the backup destination.
	BackupDestURL string `json:"BackupDestURL" plist:"BackupDestURL" required:"true"`
	// The backup size limit, in megabytes. Set to 0 for unlimited.
	BackupSizeMB *int64 `default:"0" json:"BackupSizeMB,omitempty" plist:"BackupSizeMB,omitempty"`
	// If `true`, skips system files and folders by default.
	BackupSkipSys *bool `json:"BackupSkipSys,omitempty" plist:"BackupSkipSys,omitempty"`
	// If `true`, create local backup snapshots when not connected to the network.
	MobileBackups *bool `json:"MobileBackups,omitempty" plist:"MobileBackups,omitempty"`
	// The list of paths to back up besides the startup volume.
	BasePaths *[]string `json:"BasePaths,omitempty" plist:"BasePaths,omitempty"`
	// The path to skip from start volume.
	SkipPaths *[]string `json:"SkipPaths,omitempty" plist:"SkipPaths,omitempty"`
}

func (p *TimeMachine) PayloadType() string {
	return "com.apple.MCX.TimeMachine"
}

// The payload that configures managed preferences.
type ManagedPreferences struct {
	*CommonPayloadKeys
	// The dictionary containing app preference domains. The key names are application preference domain identifiers (for example, `com.example.my-app`), or the string `.GlobalPreferences` for the global domain. The values are the corresponding forced and set-once preferences.
	PayloadContent map[string]PreferenceDomain `json:"PayloadContent" plist:"PayloadContent" required:"true"`
}

func (p *ManagedPreferences) PayloadType() string {
	return "com.apple.ManagedClient.preferences"
}

// The dictionary containing app preference domains.
type PreferenceDomain struct {
	// The dictionary of forced settings.
	Forced *[]Settings `json:"Forced,omitempty" plist:"Forced,omitempty"`
	// The dictionary of one-time settings.
	SetOnce *[]Settings `json:"Set-Once,omitempty" plist:"Set-Once,omitempty"`
}
type Settings struct {
	// The dictionary of settings.
	McxPreferenceSettings map[string]any `json:"mcx_preference_settings" plist:"mcx_preference_settings" required:"true"`
}

// The payload that configures the extensions that the system allows or disallows to run on the device.
// Specifies which NSExtension extensions are to be allowed or disallowed on a system. Extensions can be managed by bundleID allow/deny lists and "extension points".
type NSExtensionManagement struct {
	*CommonPayloadKeys
	// An array of bundle identifiers for allowed extensions.
	AllowedExtensions *[]string `json:"AllowedExtensions,omitempty" plist:"AllowedExtensions,omitempty"`
	// An array of bundle identifiers for extensions that the system doesn't allow to run.
	DeniedExtensions *[]string `json:"DeniedExtensions,omitempty" plist:"DeniedExtensions,omitempty"`
	// An array of extension points for extensions that the system doesn't allow to run.
	DeniedExtensionPoints *[]string `json:"DeniedExtensionPoints,omitempty" plist:"DeniedExtensionPoints,omitempty"`
}

func (p *NSExtensionManagement) PayloadType() string {
	return "com.apple.NSExtension"
}

// The payload that configures Setup Assistant settings.
// On macOS, this payload can specify Setup Assistant options for either the system or particular users.
type SetupAssistant struct {
	*CommonPayloadKeys
	// If `true`, the system skips the Apple Account setup pane.
	SkipCloudSetup *bool `json:"SkipCloudSetup,omitempty" plist:"SkipCloudSetup,omitempty"`
	// If `true`, the system skips the Siri setup pane.
	SkipSiriSetup *bool `json:"SkipSiriSetup,omitempty" plist:"SkipSiriSetup,omitempty"`
	// If `true`, the system skips the Privacy consent pane.
	SkipPrivacySetup *bool `json:"SkipPrivacySetup,omitempty" plist:"SkipPrivacySetup,omitempty"`
	// If `true`, the system skips the iCloud Storage pane.
	SkipiCloudStorageSetup *bool `json:"SkipiCloudStorageSetup,omitempty" plist:"SkipiCloudStorageSetup,omitempty"`
	// If `true`, the system skips the True Tone Display pane.
	SkipTrueTone *bool `json:"SkipTrueTone,omitempty" plist:"SkipTrueTone,omitempty"`
	// If `true`, the system skips the Choose Your Look pane.
	SkipAppearance *bool `json:"SkipAppearance,omitempty" plist:"SkipAppearance,omitempty"`
	// If `true`, the system skips the Touch ID setup pane.
	SkipTouchIDSetup *bool `json:"SkipTouchIDSetup,omitempty" plist:"SkipTouchIDSetup,omitempty"`
	// If `true`, the system skips the Screen Time pane.
	SkipScreenTime *bool `json:"SkipScreenTime,omitempty" plist:"SkipScreenTime,omitempty"`
	// If `true`, the system skips the Accessibility pane.
	SkipAccessibility *bool `json:"SkipAccessibility,omitempty" plist:"SkipAccessibility,omitempty"`
	// An array of strings that describe the setup items to skip. `SkipKeys` provides a list of valid strings and their meanings. Available in iOS 14 and later, and macOS 15 and later.
	SkipSetupItems *[]string `json:"SkipSetupItems,omitempty" plist:"SkipSetupItems,omitempty"`
	// If `true`, the system skips the Unlock With Apple Watch pane.
	SkipUnlockWithWatch *bool `json:"SkipUnlockWithWatch,omitempty" plist:"SkipUnlockWithWatch,omitempty"`
	// If 'true', the system skips the Wallpaper selection window.
	SkipWallpaper *bool `json:"SkipWallpaper,omitempty" plist:"SkipWallpaper,omitempty"`
}

func (p *SetupAssistant) PayloadType() string {
	return "com.apple.SetupAssistant.managed"
}

// The payload that configures ShareKit.
// macOS only. Specifies which ShareKit plugin can be accessed on client. Both allow and disallow lists can be specified.
type ShareKit struct {
	*CommonPayloadKeys
	// The list of plugin IDs that show up in the user's Share menu. If this array exists, only these items are permitted.
	SHKAllowedShareServices *[]string `json:"SHKAllowedShareServices,omitempty" plist:"SHKAllowedShareServices,omitempty"`
	// The list of plugin IDs that won't show up in the user's Share menu. This key is used only if there is no `SHKAllowedShareServices` key.
	SHKDeniedShareServices *[]string `json:"SHKDeniedShareServices,omitempty" plist:"SHKDeniedShareServices,omitempty"`
}

func (p *ShareKit) PayloadType() string {
	return "com.apple.ShareKitHelper"
}

// The payload that configures the software update policy.
// Software update catalog options.
type SoftwareUpdate struct {
	*CommonPayloadKeys
	// The URL of the software update catalog. This property is not supported in macOS 11 and later.
	CatalogURL *string `json:"CatalogURL,omitempty" plist:"CatalogURL,omitempty"`
	// If `true`, prerelease software can be installed on this computer.
	AllowPreReleaseInstallation *bool `json:"AllowPreReleaseInstallation,omitempty" plist:"AllowPreReleaseInstallation,omitempty"`
	// If `true`, restrict app installations to admin users. This key has the same function as the  `restrict-store-require-admin-to-install` key in the `com.apple.appstore` payload.
	RestrictSoftwareUpdateRequireAdminToInstall *bool `json:"restrict-software-update-require-admin-to-install,omitempty" plist:"restrict-software-update-require-admin-to-install,omitempty"`
	// If `false`, restricts the "Install macOS Updates" option and prevents the user from changing the option.
	AutomaticallyInstallMacOSUpdates *bool `json:"AutomaticallyInstallMacOSUpdates,omitempty" plist:"AutomaticallyInstallMacOSUpdates,omitempty"`
	// If `false`, deselects the "Install app updates from the App Store" option and prevents the user from changing the option.
	AutomaticallyInstallAppUpdates *bool `json:"AutomaticallyInstallAppUpdates,omitempty" plist:"AutomaticallyInstallAppUpdates,omitempty"`
	// If `false`, deselects the "Check for updates" option and prevents the user from changing the option.
	AutomaticCheckEnabled *bool `json:"AutomaticCheckEnabled,omitempty" plist:"AutomaticCheckEnabled,omitempty"`
	// If `false`, deselects the "Download new updates when available from the App Store" option and prevents the user from changing the option.
	AutomaticDownload *bool `json:"AutomaticDownload,omitempty" plist:"AutomaticDownload,omitempty"`
	// If `false`, disables the automatic installation of critical updates and prevents the user from changing the "Install system data files and security updates" option.
	CriticalUpdateInstall *bool `json:"CriticalUpdateInstall,omitempty" plist:"CriticalUpdateInstall,omitempty"`
	// If `false`, restricts the automatic installation of configuration data.
	ConfigDataInstall *bool `json:"ConfigDataInstall,omitempty" plist:"ConfigDataInstall,omitempty"`
}

func (p *SoftwareUpdate) PayloadType() string {
	return "com.apple.SoftwareUpdate"
}

// The payload that configures network proxies for a device.
type NetworkProxyConfiguration struct {
	*CommonPayloadKeys
	// The dictionary containing all the proxies for this device.
	Proxies NetworkProxyConfigurationProxies `json:"Proxies" plist:"Proxies" required:"true"`
}

func (p *NetworkProxyConfiguration) PayloadType() string {
	return "com.apple.SystemConfiguration"
}

// The dictionary containing all the proxies for this device.
type NetworkProxyConfigurationProxies struct {
	// If `true`, enables FTP proxy.
	FTPEnable *int64 `json:"FTPEnable,omitempty" plist:"FTPEnable,omitempty"`
	// If `true`, enables passive FTP mode.
	FTPPassive *int64 `json:"FTPPassive,omitempty" plist:"FTPPassive,omitempty"`
	// The FTP proxy port.
	FTPPort *int64 `json:"FTPPort,omitempty" plist:"FTPPort,omitempty"`
	// The host name or IP address for the FTP proxy.
	FTPProxy *string `json:"FTPProxy,omitempty" plist:"FTPProxy,omitempty"`
	// If `true`, enables gopher proxy.
	GopherEnable *int64 `json:"GopherEnable,omitempty" plist:"GopherEnable,omitempty"`
	// The gopher proxy port.
	GopherPort *int64 `json:"GopherPort,omitempty" plist:"GopherPort,omitempty"`
	// The host name or IP address for the gopher proxy.
	GopherProxy *string `json:"GopherProxy,omitempty" plist:"GopherProxy,omitempty"`
	// If `true`, enables web proxy.
	HTTPEnable *int64 `json:"HTTPEnable,omitempty" plist:"HTTPEnable,omitempty"`
	// The web proxy port.
	HTTPPort *int64 `json:"HTTPPort,omitempty" plist:"HTTPPort,omitempty"`
	// The host name or IP address for the web proxy.
	HTTPProxy *string `json:"HTTPProxy,omitempty" plist:"HTTPProxy,omitempty"`
	// If `true`, enables secure web proxy.
	HTTPSEnable *int64 `json:"HTTPSEnable,omitempty" plist:"HTTPSEnable,omitempty"`
	// The secure web proxy port.
	HTTPSPort *int64 `json:"HTTPSPort,omitempty" plist:"HTTPSPort,omitempty"`
	// The host name or IP address for the secure web proxy.
	HTTPSProxy *string `json:"HTTPSProxy,omitempty" plist:"HTTPSProxy,omitempty"`
	// If `true`, enables automatic proxy configuration.
	ProxyAutoConfigEnable *int64 `json:"ProxyAutoConfigEnable,omitempty" plist:"ProxyAutoConfigEnable,omitempty"`
	// The automatic proxy configuration URL.
	ProxyAutoConfigURLString *string `json:"ProxyAutoConfigURLString,omitempty" plist:"ProxyAutoConfigURLString,omitempty"`
	// If 1, allows client to log into captive portal network.
	ProxyCaptiveLoginAllowed *int64 `json:"ProxyCaptiveLoginAllowed,omitempty" plist:"ProxyCaptiveLoginAllowed,omitempty"`
	// If `true`, enable streaming proxy.
	RTSPEnable *int64 `json:"RTSPEnable,omitempty" plist:"RTSPEnable,omitempty"`
	// The streaming proxy port.
	RTSPPort *int64 `json:"RTSPPort,omitempty" plist:"RTSPPort,omitempty"`
	// The host name or IP address for the streaming proxy.
	RTSPProxy *string `json:"RTSPProxy,omitempty" plist:"RTSPProxy,omitempty"`
	// If `true`, enable the SOCKS proxy.
	SOCKSEnable *int64 `json:"SOCKSEnable,omitempty" plist:"SOCKSEnable,omitempty"`
	// The SOCKS proxy port.
	SOCKSPortinteger *int64 `json:"SOCKSPortinteger,omitempty" plist:"SOCKSPortinteger,omitempty"`
	// The host name or IP address for the SOCKS proxy.
	SOCKSProxy *string `json:"SOCKSProxy,omitempty" plist:"SOCKSProxy,omitempty"`
	// If `1`, enables fallback. Default is `1`.
	// For managed devices, if not supplied, the default is `0`.
	FallBackAllowed *int64 `json:"FallBackAllowed,omitempty" plist:"FallBackAllowed,omitempty"`
	// The list of hosts and domains that should bypass proxy settings.
	ExceptionsList *[]string `json:"ExceptionsList,omitempty" plist:"ExceptionsList,omitempty"`
}

// The payload that configures privacy preferences.
type PrivacyPreferencesPolicyControl struct {
	*CommonPayloadKeys
	// A dictionary whose keys are limited to the privacy policy control services.  In the case of conflicting specifications, the most restrictive setting (deny) is used.
	Services Services `json:"Services" plist:"Services" required:"true"`
}

func (p *PrivacyPreferencesPolicyControl) PayloadType() string {
	return "com.apple.TCC.configuration-profile-policy"
}

// A dictionary whose keys are limited to the privacy policy control services.  In the case of conflicting specifications, the most restrictive setting (deny) is used.
type Services struct {
	// Specifies the policies for contact information managed by the Contacts.app.
	AddressBook *[]*IdentityDict `json:"AddressBook,omitempty" plist:"AddressBook,omitempty"`
	// Specifies the policies for calendar information managed by the Calendar.app.
	Calendar *[]*IdentityDict `json:"Calendar,omitempty" plist:"Calendar,omitempty"`
	// Specifies the policies for reminders information managed by the Reminders app.
	Reminders *[]*IdentityDict `json:"Reminders,omitempty" plist:"Reminders,omitempty"`
	// The pictures managed by the Photos app in `~/Pictures/.photoslibrary`.
	Photos *[]*IdentityDict `json:"Photos,omitempty" plist:"Photos,omitempty"`
	// A system camera. Access to the camera can't be given in a profile; it can only be denied.
	Camera *[]*IdentityDict `json:"Camera,omitempty" plist:"Camera,omitempty"`
	// A system microphone. Access to the microphone can't be given in a profile; it can only be denied.
	Microphone *[]*IdentityDict `json:"Microphone,omitempty" plist:"Microphone,omitempty"`
	// Specifies the policies for the app via the Accessibility subsystem. The ability to grant access by this profile is deprecated as of macOS 26.2, and will be removed in macOS 27.0.
	Accessibility *[]*IdentityDict `json:"Accessibility,omitempty" plist:"Accessibility,omitempty"`
	// Specifies the policies for the application to use CoreGraphics APIs to send CGEvents to the system event stream.
	PostEvent *[]*IdentityDict `json:"PostEvent,omitempty" plist:"PostEvent,omitempty"`
	// Allows the application access to all protected files, including system administration files.
	SystemPolicyAllFiles *[]*IdentityDict `json:"SystemPolicyAllFiles,omitempty" plist:"SystemPolicyAllFiles,omitempty"`
	// Allows the application access to some files used in system administration.
	SystemPolicySysAdminFiles *[]*IdentityDict `json:"SystemPolicySysAdminFiles,omitempty" plist:"SystemPolicySysAdminFiles,omitempty"`
	// Specifies the policies for the app sending restricted AppleEvents to another process.
	AppleEvents *[]*IdentityDict `json:"AppleEvents,omitempty" plist:"AppleEvents,omitempty"`
	// Allows the application to access Apple Music, music and video activity, and the media library.
	MediaLibrary *[]*IdentityDict `json:"MediaLibrary,omitempty" plist:"MediaLibrary,omitempty"`
	// Allows a File Provider application to know when the user is using files managed by the File Provider.
	FileProviderPresence *[]*IdentityDict `json:"FileProviderPresence,omitempty" plist:"FileProviderPresence,omitempty"`
	// Allows the application to use CoreGraphics and HID APIs to listen to (receive) CGEvents and HID events from all processes. Access to these events can't be given in a profile; it can only be denied.
	ListenEvent *[]*IdentityDict `json:"ListenEvent,omitempty" plist:"ListenEvent,omitempty"`
	// Allows the application to capture (read) the contents of the system display. Access to the contents can't be given in a profile; it can only be denied.
	ScreenCapture *[]*IdentityDict `json:"ScreenCapture,omitempty" plist:"ScreenCapture,omitempty"`
	// Allows the application to use the system Speech Recognition facility and to send speech data to Apple.
	SpeechRecognition *[]*IdentityDict `json:"SpeechRecognition,omitempty" plist:"SpeechRecognition,omitempty"`
	// Allows the application to access files in the user's Desktop folder.
	SystemPolicyDesktopFolder *[]*IdentityDict `json:"SystemPolicyDesktopFolder,omitempty" plist:"SystemPolicyDesktopFolder,omitempty"`
	// Allows the application to access files in the user's Documents folder.
	SystemPolicyDocumentsFolder *[]*IdentityDict `json:"SystemPolicyDocumentsFolder,omitempty" plist:"SystemPolicyDocumentsFolder,omitempty"`
	// Allows the application to access files in the user's Downloads folder.
	SystemPolicyDownloadsFolder *[]*IdentityDict `json:"SystemPolicyDownloadsFolder,omitempty" plist:"SystemPolicyDownloadsFolder,omitempty"`
	// Allows the application to access files on network volumes.
	SystemPolicyNetworkVolumes *[]*IdentityDict `json:"SystemPolicyNetworkVolumes,omitempty" plist:"SystemPolicyNetworkVolumes,omitempty"`
	// Allows the application to access files on removable volumes.
	SystemPolicyRemovableVolumes *[]*IdentityDict `json:"SystemPolicyRemovableVolumes,omitempty" plist:"SystemPolicyRemovableVolumes,omitempty"`
	// Allows the application to update or delete other apps. Available in macOS 13 and later.
	SystemPolicyAppBundles *[]*IdentityDict `json:"SystemPolicyAppBundles,omitempty" plist:"SystemPolicyAppBundles,omitempty"`
	// Specifies the policies for the app to access the data of other apps.
	SystemPolicyAppData *[]*IdentityDict `json:"SystemPolicyAppData,omitempty" plist:"SystemPolicyAppData,omitempty"`
	// Specifies the policies for the app to access Bluetooth devices.
	BluetoothAlways *[]*IdentityDict `json:"BluetoothAlways,omitempty" plist:"BluetoothAlways,omitempty"`
}

// A dictionary listing apps and the privacy policy to apply to them.
type IdentityDict struct {
	// The bundle ID or installation path of the binary.
	// > Note:
	// > This value is case-sensitive.
	Identifier string `json:"Identifier" plist:"Identifier" required:"true"`
	// The type of identifier value. Application bundles must be identified by bundle ID. Nonbundled binaries must be identified by installation path. Helper tools embedded within an application bundle automatically inherit the permissions of their enclosing app bundle.
	IdentifierType IdentifierType `json:"IdentifierType" plist:"IdentifierType" required:"true"`
	// Obtained via the command `codesign -display -r -`.
	CodeRequirement string `json:"CodeRequirement" plist:"CodeRequirement" required:"true"`
	// If `true`, statically validate the code requirement. Used only if the process invalidates its dynamic code signature.
	StaticCode *bool `json:"StaticCode,omitempty" plist:"StaticCode,omitempty"`
	// If `true`, access is granted; otherwise, the process doesn't have access. The user isn't prompted and can't change this value.
	// > Note:
	// > Every payload needs to include either `Authorization` or `Allowed`, but not both.
	Allowed *bool `json:"Allowed,omitempty" plist:"Allowed,omitempty"`
	// The `Authorization` key is an optional replacement for the `Allowed` key, which has one of the following possible values:
	// - `Allow`: Equivalent to a `true` value for the `Allowed` key
	// - `Deny`: Equivalent to a `false` value for the `Allowed` key
	// - `AllowStandardUserToSetSystemService`: Allows a standard (non-admin) user to configure the permissions for the specified app in the Privacy preferences for services that otherwise require admin authorization; only valid for the `ListenEvent` and `ScreenCapture` services
	// > Note:
	// > Every payload needs to include either `Authorization` or `Allowed`, but not both.
	// Available in macOS 11 and later.
	Authorization *Authorization `json:"Authorization,omitempty" plist:"Authorization,omitempty"`
	// Not used.
	Comment *string `json:"Comment,omitempty" plist:"Comment,omitempty"`
	// The identifier of the process receiving an AppleEvent sent by the Identifier process. This identifier is required for AppleEvents service; not valid for other services.
	AEReceiverIdentifier *string `json:"AEReceiverIdentifier,omitempty" plist:"AEReceiverIdentifier,omitempty"`
	// The type of AEReceiverIdentifier value, either `bundleID` or `path`. This setting is required for AppleEvents service; not valid for other services.
	AEReceiverIdentifierType *AEReceiverIdentifierType `json:"AEReceiverIdentifierType,omitempty" plist:"AEReceiverIdentifierType,omitempty"`
	// The code requirement for the receiving binary. This code requirement is required for AppleEvents service; not valid for other services.
	AEReceiverCodeRequirement *string `json:"AEReceiverCodeRequirement,omitempty" plist:"AEReceiverCodeRequirement,omitempty"`
}

// The type of identifier value. Application bundles must be identified by bundle ID. Nonbundled binaries must be identified by installation path. Helper tools embedded within an application bundle automatically inherit the permissions of their enclosing app bundle.
type IdentifierType string

const (
	IdentifierTypeBundleID IdentifierType = "bundleID"
	IdentifierTypePath     IdentifierType = "path"
)

// The `Authorization` key is an optional replacement for the `Allowed` key, which has one of the following possible values:
// - `Allow`: Equivalent to a `true` value for the `Allowed` key
// - `Deny`: Equivalent to a `false` value for the `Allowed` key
// - `AllowStandardUserToSetSystemService`: Allows a standard (non-admin) user to configure the permissions for the specified app in the Privacy preferences for services that otherwise require admin authorization; only valid for the `ListenEvent` and `ScreenCapture` services
// > Note:
// > Every payload needs to include either `Authorization` or `Allowed`, but not both.
// Available in macOS 11 and later.
type Authorization string

const (
	AuthorizationAllow                               Authorization = "Allow"
	AuthorizationDeny                                Authorization = "Deny"
	AuthorizationAllowStandardUserToSetSystemService Authorization = "AllowStandardUserToSetSystemService"
)

// The type of AEReceiverIdentifier value, either `bundleID` or `path`. This setting is required for AppleEvents service; not valid for other services.
type AEReceiverIdentifierType string

const (
	AEReceiverIdentifierTypeBundleID AEReceiverIdentifierType = "bundleID"
	AEReceiverIdentifierTypePath     AEReceiverIdentifierType = "path"
)

// The payload that configures Apple TV for a particular style of AirPlay security.
// Manages the AirPlay Security settings on Apple TV (Settings > AirPlay > Security). Use this payload to lock Apple TV to a particular style of AirPlay security. The setting can enable/disable an on-screen passcode, or require a specific password phrase.
type AirPlaySecurity struct {
	*CommonPayloadKeys
	// The security policy for AirPlay. Allowed values:
	// - `PASSCODE_ONCE`: Requires an onscreen passcode on first connection from a device. Subsequent connections from the same device aren't prompted.
	// - `PASSCODE_ALWAYS`: Requires an onscreen passcode for every AirPlay connection. After an AirPlay connection ends, the system allows reconnecting within 30 seconds without a password.
	// - `PASSWORD`: Requires the passphrase set for `Password`.
	// > Note:
	// > `NONE` was deprecated in tvOS 11.3. Existing profiles that use `NONE` get the `PASSWORD_ONCE` behavior.
	SecurityType SecurityType `json:"SecurityType" plist:"SecurityType" required:"true"`
	// The access policy for AirPlay.
	// `ANY` allows connections from both Ethernet, Wi-Fi, and Apple Wireless Direct Link.
	// `WIFI_ONLY` allows connections only from devices on the same Ethernet or Wi-Fi network as Apple TV.
	AccessType AccessType `json:"AccessType" plist:"AccessType" required:"true"`
	// The AirPlay password; required if `SecurityType` is `PASSWORD`.
	Password *string `json:"Password,omitempty" plist:"Password,omitempty"`
}

func (p *AirPlaySecurity) PayloadType() string {
	return "com.apple.airplay.security"
}

// The security policy for AirPlay. Allowed values:
// - `PASSCODE_ONCE`: Requires an onscreen passcode on first connection from a device. Subsequent connections from the same device aren't prompted.
// - `PASSCODE_ALWAYS`: Requires an onscreen passcode for every AirPlay connection. After an AirPlay connection ends, the system allows reconnecting within 30 seconds without a password.
// - `PASSWORD`: Requires the passphrase set for `Password`.
// > Note:
// > `NONE` was deprecated in tvOS 11.3. Existing profiles that use `NONE` get the `PASSWORD_ONCE` behavior.
type SecurityType string

const (
	SecurityTypePASSCODEONCE   SecurityType = "PASSCODE_ONCE"
	SecurityTypePASSCODEALWAYS SecurityType = "PASSCODE_ALWAYS"
	SecurityTypePASSWORD       SecurityType = "PASSWORD"
)

// The access policy for AirPlay.
// `ANY` allows connections from both Ethernet, Wi-Fi, and Apple Wireless Direct Link.
// `WIFI_ONLY` allows connections only from devices on the same Ethernet or Wi-Fi network as Apple TV.
type AccessType string

const (
	AccessTypeANY      AccessType = "ANY"
	AccessTypeWIFIONLY AccessType = "WIFI_ONLY"
)

// The payload that configures AirPlay settings.
// macOS supports more than one payload, iOS does not. Supported on the user channel for macOS only.
type AirPlay struct {
	*CommonPayloadKeys
	// If present, only AirPlay destinations in this list are available to the device. This allow list applies to supervised devices.
	AllowList *[]AllowListItem `json:"AllowList,omitempty" plist:"AllowList,omitempty"`
	// If present, sets passwords for known AirPlay destinations. Using multiple entries for the same destination, whether within the same payload or across multiple installed payloads, is an error and results in undefined behavior.
	Passwords *[]PasswordsItem `json:"Passwords,omitempty" plist:"Passwords,omitempty"`
	// Use `AllowList` instead. This key is deprecated in iOS 14.5 and macOS 11.3.
	Whitelist *[]AllowListItem `json:"Whitelist,omitempty" plist:"Whitelist,omitempty"`
}

func (p *AirPlay) PayloadType() string {
	return "com.apple.airplay"
}

type AllowListItem struct {
	// The device ID of the AirPlay destination in the format `xx:xx:xx:xx:xx:xx`. This field isn't case-sensitive.
	// The system limits the list of visible AirPlay destinations to devices that are present in the `AllowList` field of all installed AirPlay payloads.
	// Specifying the same MACAddress more than once, whether in the same payload across different payloads, results in undefined behavior.
	// As of tvOS 18, `DeviceID` isn't supported.
	DeviceID *string `json:"DeviceID,omitempty" plist:"DeviceID,omitempty"`
	// The name of the AirPlay device.
	// The system limits the list of visible AirPlay destinations to devices that are present in the `AllowList` field of all installed AirPlay payloads.
	DeviceName *string `json:"DeviceName,omitempty" plist:"DeviceName,omitempty"`
}
type PasswordsItem struct {
	// The name of the AirPlay destination; used in iOS, and available in macOS 15 and later.
	DeviceName *string `json:"DeviceName,omitempty" plist:"DeviceName,omitempty"`
	// The password for the AirPlay destination.
	Password string `json:"Password" plist:"Password" required:"true"`
	// The device ID of the AirPlay destination; used in macOS.
	// Deprecated in macOS 15 and later as tvOS 18 AirPlay destinations don't support it; use `DeviceName` instead.
	DeviceID *string `json:"DeviceID,omitempty" plist:"DeviceID,omitempty"`
}

// The payload that configures AirPrint printer discoverability in the user's printer list.
type AirPrint struct {
	*CommonPayloadKeys
	// An array of AirPrint printers that are presented to the user.
	AirPrint []*AirPrintItem `json:"AirPrint" plist:"AirPrint" required:"true"`
}

func (p *AirPrint) PayloadType() string {
	return "com.apple.airprint"
}

type AirPrintItem struct {
	// The IP address or hostname of the AirPrint destination.
	IPAddress string `json:"IPAddress" plist:"IPAddress" required:"true"`
	// The resource path associated with the printer. This path corresponds to the `rp` parameter of the `_ipps.tcp` Bonjour record. For example:
	// - `printers/Canon_MG5300_series`
	// - `printers/Xerox_Phaser_7600`
	// - `ipp/print`
	// - `Epson_IPP_Printer`
	ResourcePath string `json:"ResourcePath" plist:"ResourcePath" required:"true"`
	// The listening port of the AirPrint destination. Available only in iOS 11 and later.
	Port *int64 `json:"Port,omitempty" plist:"Port,omitempty"`
	// If `true`, AirPrint connections are secured by Transport Layer Security (TLS). Available only in iOS 11 and later.
	ForceTLS *bool `json:"ForceTLS,omitempty" plist:"ForceTLS,omitempty"`
}

// The payload that configures access point names.
// Not supported in macOS.
// This technically does install on watchOS but we are removing the supportedOS dictionary. The cellular payload should be used instead.
// Only applies to the preferred data SIM.
// Deprecated. Use Cellular instead.
type APN struct {
	*CommonPayloadKeys
	// The list of access point names (APNs).
	DefaultsData DefaultsData `json:"DefaultsData" plist:"DefaultsData" required:"true"`
	// The domain name.
	DefaultsDomainName DefaultsDomainName `json:"DefaultsDomainName" plist:"DefaultsDomainName" required:"true"`
}

func (p *APN) PayloadType() string {
	return "com.apple.apn.managed"
}

// The list of access point names (APNs).
type DefaultsData struct {
	// An array of APN dictionaries (\`APN.DefaultsData.Apns\`).
	Apns []*ApnsItem `json:"apns" plist:"apns" required:"true"`
}

// A dictionary that describes an APN configuration.
type ApnsItem struct {
	// The access point name.
	Apn string `json:"apn" plist:"apn" required:"true"`
	// The user name. If missing, the device prompts for it during profile installation.
	Username *string `json:"username,omitempty" plist:"username,omitempty"`
	// The password for the user. For obfuscation purposes, the system encodes the password. If missing, the device prompts for the password during profile installation.
	Password *[]byte `json:"password,omitempty" plist:"password,omitempty"`
	// The IP address or URL of the APN proxy.
	Proxy *string `json:"proxy,omitempty" plist:"proxy,omitempty"`
	// The port number of the APN proxy.
	ProxyPort *int64 `json:"proxyPort,omitempty" plist:"proxyPort,omitempty"`
}

// The domain name.
type DefaultsDomainName string

const (
	DefaultsDomainNameComapplemanagedCarrier DefaultsDomainName = "com.apple.managedCarrier"
)

// The payload that configures a device to run a single app.
type AppLock struct {
	*CommonPayloadKeys
	// A dictionary that contains information about the app.
	App App `json:"App" plist:"App" required:"true"`
}

func (p *AppLock) PayloadType() string {
	return "com.apple.app.lock"
}

// A dictionary that contains information about the app.
type App struct {
	// The app's bundle identifier.
	Identifier string `json:"Identifier" plist:"Identifier" required:"true"`
	// A dictionary of options that the user can't change.
	Options *Options `json:"Options,omitempty" plist:"Options,omitempty"`
	// A dictionary of user-editable options.
	UserEnabledOptions *UserEnabledOptions `json:"UserEnabledOptions,omitempty" plist:"UserEnabledOptions,omitempty"`
}

// A dictionary of options that the user can't change.
type Options struct {
	// If `true`, the system disables the touch screen. In tvOS, it disables the touch surface on the Apple TV Remote.
	DisableTouch *bool `json:"DisableTouch,omitempty" plist:"DisableTouch,omitempty"`
	// If `true`, the system disables device rotation sensing.
	DisableDeviceRotation *bool `json:"DisableDeviceRotation,omitempty" plist:"DisableDeviceRotation,omitempty"`
	// If `true`, the system disables the volume buttons.
	DisableVolumeButtons *bool `json:"DisableVolumeButtons,omitempty" plist:"DisableVolumeButtons,omitempty"`
	// If `true`, the system disables the ringer switch. When disabled, the ringer behavior depends on what position the switch was in when it was first disabled.
	DisableRingerSwitch *bool `json:"DisableRingerSwitch,omitempty" plist:"DisableRingerSwitch,omitempty"`
	// If `true`, the system disables the sleep/wake button.
	DisableSleepWakeButton *bool `json:"DisableSleepWakeButton,omitempty" plist:"DisableSleepWakeButton,omitempty"`
	// If `true`, the device doesn't automatically go to sleep after an idle period.
	DisableAutoLock *bool `json:"DisableAutoLock,omitempty" plist:"DisableAutoLock,omitempty"`
	// If `true`, the system enables VoiceOver.
	EnableVoiceOver *bool `json:"EnableVoiceOver,omitempty" plist:"EnableVoiceOver,omitempty"`
	// If `true`, the system enables Zoom.
	EnableZoom *bool `json:"EnableZoom,omitempty" plist:"EnableZoom,omitempty"`
	// If `true`, the system enables Invert Colors.
	EnableInvertColors *bool `json:"EnableInvertColors,omitempty" plist:"EnableInvertColors,omitempty"`
	// If `true`, the system enables AssistiveTouch.
	EnableAssistiveTouch *bool `json:"EnableAssistiveTouch,omitempty" plist:"EnableAssistiveTouch,omitempty"`
	// If `true`, the system enables Speak Selection.
	EnableSpeakSelection *bool `json:"EnableSpeakSelection,omitempty" plist:"EnableSpeakSelection,omitempty"`
	// If `true`, the system enables Mono Audio.
	EnableMonoAudio *bool `json:"EnableMonoAudio,omitempty" plist:"EnableMonoAudio,omitempty"`
	// If `true`, the system enables Voice Control.
	EnableVoiceControl *bool `json:"EnableVoiceControl,omitempty" plist:"EnableVoiceControl,omitempty"`
}

// A dictionary of user-editable options.
type UserEnabledOptions struct {
	// If `true`, the system allows the user to toggle Voice Control.
	VoiceControl *bool `json:"VoiceControl,omitempty" plist:"VoiceControl,omitempty"`
	// If `true`, the system allows the user to toggle VoiceOver.
	VoiceOver *bool `json:"VoiceOver,omitempty" plist:"VoiceOver,omitempty"`
	// If `true`, the system allows the user to toggle Zoom.
	Zoom *bool `json:"Zoom,omitempty" plist:"Zoom,omitempty"`
	// If `true`, the system allows the user to toggle Invert Colors.
	InvertColors *bool `json:"InvertColors,omitempty" plist:"InvertColors,omitempty"`
	// If `true`, the system allows the user to toggle AssistiveTouch.
	AssistiveTouch *bool `json:"AssistiveTouch,omitempty" plist:"AssistiveTouch,omitempty"`
}

// The payload that configures parental controls for apps.
// Parental controls application restrictions.
// Order of evaluation:
// (1) Certain system applications and utilities are always allowed to run
// (2) The "whiteList" is searched to see if a matching entry is found by bundleID. If a match is found, the "appID" and "detachedSignature"
// (if present) are used to verify the signature of the application being launched. If the signature is valid and matches the designated
// requirement (in the "appID" key), the application is allowed to launch.
// (3) (deprecated) If the path to the binary being launched matches (or is in a subdirectory) of a path in "pathBlackList", the binary is denied.
// (4) (deprecated) If the path to the binary being launched matches (or is a subdirectory) of a path in "pathWhiteList", the binary is allowed to launch.
// (5) The binary is denied permission to launch.
type ParentalControlsApplicationRestrictions struct {
	*CommonPayloadKeys
	// If `true`, enables app access restrictions.
	FamilyControlsEnabled bool `json:"familyControlsEnabled" plist:"familyControlsEnabled" required:"true"`
	// The allow list of app item dictionaries.
	WhiteList *[]*ParentalControlsApplicationRestrictionsWhiteListWhiteListItem `json:"whiteList,omitempty" plist:"whiteList,omitempty"`
	// The paths to apps in the deny list. This property is deprecated in macOS 10.15 and later.
	PathBlackList *[]string `json:"pathBlackList,omitempty" plist:"pathBlackList,omitempty"`
	// The paths to apps in the allow list. This property is deprecated in macOS 10.15 and later.
	PathWhiteList *[]string `json:"pathWhiteList,omitempty" plist:"pathWhiteList,omitempty"`
}

func (p *ParentalControlsApplicationRestrictions) PayloadType() string {
	return "com.apple.applicationaccess.new"
}

// A dictionary defining an app for parental control.
type ParentalControlsApplicationRestrictionsWhiteListWhiteListItem struct {
	// The bundle ID of the app.
	BundleID string `json:"bundleID" plist:"bundleID" required:"true"`
	// The identifier of the app. Obtain this value from the Security framework using `SecCodeCopyDesignatedRequirement`.
	AppID []byte `json:"appID" plist:"appID" required:"true"`
	// The signature for an unsigned binary.
	DetachedSignature *[]byte `json:"detachedSignature,omitempty" plist:"detachedSignature,omitempty"`
	// If `true`, this app isn't added to the allow list.
	Disabled *bool `json:"disabled,omitempty" plist:"disabled,omitempty"`
	// An array of nested helper applications.
	SubApps *[]*ParentalControlsApplicationRestrictionsWhiteListWhiteListItem `json:"subApps,omitempty" plist:"subApps,omitempty"`
	// The name used for display purposes.
	DisplayName *string `json:"displayName,omitempty" plist:"displayName,omitempty"`
}

// The payload that configures restrictions on a device.
type Restrictions struct {
	*CommonPayloadKeys
	// If `false`, the system disables modification of accounts, such as Apple Accounts, and internet-based accounts, such as Mail, Contacts, and Calendar.
	AllowAccountModification *bool `json:"allowAccountModification,omitempty" plist:"allowAccountModification,omitempty"`
	// If `false`, the system disables activity continuation. Support for this restriction on unsupervised devices and with Managed Apple Accounts is deprecated. In a future release, this restriction will begin requiring supervision and will apply to personal Apple Accounts only.
	AllowActivityContinuation *bool `json:"allowActivityContinuation,omitempty" plist:"allowActivityContinuation,omitempty"`
	// If `false`, the system prohibits adding friends to Game Center. Requires a supervised device in iOS 13 and later.
	AllowAddingGameCenterFriends *bool `json:"allowAddingGameCenterFriends,omitempty" plist:"allowAddingGameCenterFriends,omitempty"`
	// If `false`, the system disables AirDrop.
	AllowAirDrop *bool `json:"allowAirDrop,omitempty" plist:"allowAirDrop,omitempty"`
	// If `false`, the system disables incoming AirPlay requests.
	AllowAirPlayIncomingRequests *bool `json:"allowAirPlayIncomingRequests,omitempty" plist:"allowAirPlayIncomingRequests,omitempty"`
	// If `false`, the system disables AirPrint.
	AllowAirPrint *bool `json:"allowAirPrint,omitempty" plist:"allowAirPrint,omitempty"`
	// If `false`, the system disables Keychain storage of user name and password for AirPrint.
	AllowAirPrintCredentialsStorage *bool `json:"allowAirPrintCredentialsStorage,omitempty" plist:"allowAirPrintCredentialsStorage,omitempty"`
	// If `false`, the system disables iBeacon discovery of AirPrint printers, which prevents spurious AirPrint Bluetooth beacons from phishing for network traffic.
	AllowAirPrintiBeaconDiscovery *bool `json:"allowAirPrintiBeaconDiscovery,omitempty" plist:"allowAirPrintiBeaconDiscovery,omitempty"`
	// If `false`, the system disables changing settings for cellular data usage for apps.
	AllowAppCellularDataModification *bool `json:"allowAppCellularDataModification,omitempty" plist:"allowAppCellularDataModification,omitempty"`
	// If `false`, the system prevents a user from adding any App Clips, and removes any existing App Clips on the device.
	AllowAppClips *bool `json:"allowAppClips,omitempty" plist:"allowAppClips,omitempty"`
	// If `false`, the system disables the App Store and removes its icon from the Home Screen. Users are unable to install or update their apps. This applies to App Store apps, marketplace apps, and locally installed apps (using Configurator, Xcode, and so forth).
	// In iOS 10 and later, MDM commands can override this restriction. Requires a supervised device in iOS 13 and later.
	AllowAppInstallation *bool `json:"allowAppInstallation,omitempty" plist:"allowAppInstallation,omitempty"`
	// If `false`, the system disables Apple Intelligence reports.
	AllowAppleIntelligenceReport *bool `json:"allowAppleIntelligenceReport,omitempty" plist:"allowAppleIntelligenceReport,omitempty"`
	// If `false`, the system limits Apple personalized advertising.
	AllowApplePersonalizedAdvertising *bool `json:"allowApplePersonalizedAdvertising,omitempty" plist:"allowApplePersonalizedAdvertising,omitempty"`
	// If `false`, the system disables removal of apps from an iOS device. This applies to App Store apps, marketplace apps, and locally installed apps (using Configurator, Xcode, and so forth).
	AllowAppRemoval *bool `json:"allowAppRemoval,omitempty" plist:"allowAppRemoval,omitempty"`
	// If `false`, disables the ability for the user to hide apps. It doesn't affect the user's ability to leave it in the App Library, while removing it from the Home Screen.
	AllowAppsToBeHidden *bool `json:"allowAppsToBeHidden,omitempty" plist:"allowAppsToBeHidden,omitempty"`
	// If `false`, disables the ability for the user to lock apps. Because hiding apps also requires locking them, disallowing locking also disallows hiding.
	AllowAppsToBeLocked *bool `json:"allowAppsToBeLocked,omitempty" plist:"allowAppsToBeLocked,omitempty"`
	// If `false`, the system prevents modifying the Remote Management Sharing setting in System Settings.
	AllowARDRemoteManagementModification *bool `json:"allowARDRemoteManagementModification,omitempty" plist:"allowARDRemoteManagementModification,omitempty"`
	// If `false`, the system disables Siri.
	AllowAssistant *bool `json:"allowAssistant,omitempty" plist:"allowAssistant,omitempty"`
	// If `false`, the system prevents Siri from querying user-generated content from the web.
	AllowAssistantUserGeneratedContent *bool `json:"allowAssistantUserGeneratedContent,omitempty" plist:"allowAssistantUserGeneratedContent,omitempty"`
	// If `false`, the system disables Siri when the device is locked. The system ignores this restriction if the device doesn't have a passcode set.
	AllowAssistantWhileLocked *bool `json:"allowAssistantWhileLocked,omitempty" plist:"allowAssistantWhileLocked,omitempty"`
	// If `false`, the system disables keyboard autocorrection.
	AllowAutoCorrection *bool `json:"allowAutoCorrection,omitempty" plist:"allowAutoCorrection,omitempty"`
	// If `false`, disables auto dim on iPads with OLED displays.
	AllowAutoDim *bool `json:"allowAutoDim,omitempty" plist:"allowAutoDim,omitempty"`
	// If `false`, the system prevents automatic downloading of apps purchased on other devices. This setting doesn't affect updates to existing apps.
	AllowAutomaticAppDownloads *bool `json:"allowAutomaticAppDownloads,omitempty" plist:"allowAutomaticAppDownloads,omitempty"`
	// If `false`, the system disables Apple TV's automatic screen saver.
	AllowAutomaticScreenSaver *bool `json:"allowAutomaticScreenSaver,omitempty" plist:"allowAutomaticScreenSaver,omitempty"`
	// If `false`, the system disallows auto unlock. Support for this restriction on unsupervised devices is deprecated.
	AllowAutoUnlock *bool `json:"allowAutoUnlock,omitempty" plist:"allowAutoUnlock,omitempty"`
	// If `false`, the system prevents modification of Bluetooth settings.
	AllowBluetoothModification *bool `json:"allowBluetoothModification,omitempty" plist:"allowBluetoothModification,omitempty"`
	// If `false`, the system prevents modifying Bluetooth settings in System Settings.
	AllowBluetoothSharingModification *bool `json:"allowBluetoothSharingModification,omitempty" plist:"allowBluetoothSharingModification,omitempty"`
	// If `false`, the system removes the Book Store tab from the Books app.
	AllowBookstore *bool `json:"allowBookstore,omitempty" plist:"allowBookstore,omitempty"`
	// If `false`, the system prevents the user from downloading Apple Books media that's tagged as erotica. Support for this restriction on unsupervised devices is deprecated.
	AllowBookstoreErotica *bool `json:"allowBookstoreErotica,omitempty" plist:"allowBookstoreErotica,omitempty"`
	// If `false`, disables call recording.
	AllowCallRecording *bool `json:"allowCallRecording,omitempty" plist:"allowCallRecording,omitempty"`
	// If `false`, the system disables the camera and removes its icon from the Home Screen, and users are unable to take photographs. Support for this restriction on unsupervised devices is deprecated.
	AllowCamera *bool `json:"allowCamera,omitempty" plist:"allowCamera,omitempty"`
	// If `false`, the system prevents users from changing settings related to their cellular plan (available only on select carriers).
	AllowCellularPlanModification *bool `json:"allowCellularPlanModification,omitempty" plist:"allowCellularPlanModification,omitempty"`
	// If `false`, the system disables the use of iMessage with supervised devices. If the device supports text messaging, the user can still send and receive text messages.
	AllowChat *bool `json:"allowChat,omitempty" plist:"allowChat,omitempty"`
	// If `false`, the system disables iCloud Contacts services.
	AllowCloudAddressBook *bool `json:"allowCloudAddressBook,omitempty" plist:"allowCloudAddressBook,omitempty"`
	// If `false`, the system disables backing up the device to iCloud. Support for this restriction on unsupervised devices is deprecated.
	AllowCloudBackup *bool `json:"allowCloudBackup,omitempty" plist:"allowCloudBackup,omitempty"`
	// If `false`, the system disables iCloud Bookmark sync.
	AllowCloudBookmarks *bool `json:"allowCloudBookmarks,omitempty" plist:"allowCloudBookmarks,omitempty"`
	// If `false`, the system disables iCloud Calendar services.
	AllowCloudCalendar *bool `json:"allowCloudCalendar,omitempty" plist:"allowCloudCalendar,omitempty"`
	// If `false`, the system disables iCloud Desktop and Document services.
	AllowCloudDesktopAndDocuments *bool `json:"allowCloudDesktopAndDocuments,omitempty" plist:"allowCloudDesktopAndDocuments,omitempty"`
	// If `false`, the system disables document and key-value syncing to iCloud. Requires a supervised device in iOS 13 and later, and Shared iPad doesn't support it. Support for this restriction on unsupervised devices and with Managed Apple Accounts is deprecated.
	AllowCloudDocumentSync *bool `json:"allowCloudDocumentSync,omitempty" plist:"allowCloudDocumentSync,omitempty"`
	// If `false`, the system disallows iCloud Freeform services.
	AllowCloudFreeform *bool `json:"allowCloudFreeform,omitempty" plist:"allowCloudFreeform,omitempty"`
	// If `false`, the system disables iCloud Keychain synchronization. Support for this restriction on unsupervised devices and with Managed Apple Accounts is deprecated.
	AllowCloudKeychainSync *bool `json:"allowCloudKeychainSync,omitempty" plist:"allowCloudKeychainSync,omitempty"`
	// If `false`, the system disables iCloud Mail services.
	AllowCloudMail *bool `json:"allowCloudMail,omitempty" plist:"allowCloudMail,omitempty"`
	// If `false`, the system disables iCloud Notes services.
	AllowCloudNotes *bool `json:"allowCloudNotes,omitempty" plist:"allowCloudNotes,omitempty"`
	// If `false`, the system disables iCloud Photo Library. The system removes any photos from local storage that aren't fully downloaded from iCloud Photo Library to the device. Support for this restriction on unsupervised devices and with Managed Apple Accounts is deprecated.
	AllowCloudPhotoLibrary *bool `json:"allowCloudPhotoLibrary,omitempty" plist:"allowCloudPhotoLibrary,omitempty"`
	// If `false`, the system disables iCloud Private Relay. Support for this restriction on unsupervised devices and with Managed Apple Accounts is deprecated.
	AllowCloudPrivateRelay *bool `json:"allowCloudPrivateRelay,omitempty" plist:"allowCloudPrivateRelay,omitempty"`
	// If `false`, the system disables iCloud Reminder services.
	AllowCloudReminders *bool `json:"allowCloudReminders,omitempty" plist:"allowCloudReminders,omitempty"`
	// If `false`, the system disables content caching. This restriction is not supported on the user channel.
	AllowContentCaching *bool `json:"allowContentCaching,omitempty" plist:"allowContentCaching,omitempty"`
	// If `false`, the system disables QuickPath keyboard.
	AllowContinuousPathKeyboard *bool `json:"allowContinuousPathKeyboard,omitempty" plist:"allowContinuousPathKeyboard,omitempty"`
	// If `false`, disables default browser preference modification. The MDM Settings command to set the default browser preference still works when applying this.
	AllowDefaultBrowserModification *bool `json:"allowDefaultBrowserModification,omitempty" plist:"allowDefaultBrowserModification,omitempty"`
	// If `false`, disables default calling app preference modification. The MDM Settings command to set the default calling app preference still works when applying this.
	AllowDefaultCallingAppModification *bool `json:"allowDefaultCallingAppModification,omitempty" plist:"allowDefaultCallingAppModification,omitempty"`
	// If `false`, disables default messaging app preference modification. The MDM Settings command to set the default messaging app preference still works when applying this.
	AllowDefaultMessagingAppModification *bool `json:"allowDefaultMessagingAppModification,omitempty" plist:"allowDefaultMessagingAppModification,omitempty"`
	// If `false`, the system disables definition lookup.
	AllowDefinitionLookup *bool `json:"allowDefinitionLookup,omitempty" plist:"allowDefinitionLookup,omitempty"`
	// If `false`, the system prevents the user from changing the device name.
	AllowDeviceNameModification *bool `json:"allowDeviceNameModification,omitempty" plist:"allowDeviceNameModification,omitempty"`
	// If `false`, the system prevents the device from automatically sleeping.
	AllowDeviceSleep *bool `json:"allowDeviceSleep,omitempty" plist:"allowDeviceSleep,omitempty"`
	// If `false`, the system prevents the device from automatically submitting diagnostic reports to Apple.
	AllowDiagnosticSubmission *bool `json:"allowDiagnosticSubmission,omitempty" plist:"allowDiagnosticSubmission,omitempty"`
	// If `false`, the system disables changing the diagnostic submission and app analytics settings in the Diagnostics & Usage UI in Settings.
	AllowDiagnosticSubmissionModification *bool `json:"allowDiagnosticSubmissionModification,omitempty" plist:"allowDiagnosticSubmissionModification,omitempty"`
	// If `false`, the system disallows dictation input.
	AllowDictation *bool `json:"allowDictation,omitempty" plist:"allowDictation,omitempty"`
	// If present, the system exempts apps with bundle IDs in the array from the `allowCamera` restriction. The system doesn't grant these apps access to the camera automatically; they're only exempted from the `allowCamera` restriction. This key has no effect when the camera isn't restricted. Multiple payloads combine using an intersect operation. Requires a supervised device.
	AllowedCameraRestrictionBundleIDs *[]string `json:"allowedCameraRestrictionBundleIDs,omitempty" plist:"allowedCameraRestrictionBundleIDs,omitempty"`
	// An array of strings, but currently restricted to a single element. If present, Apple Intelligence allows use of only the given external integration workspace ID, and requires a sign-in to make requests. The user is required to sign in to integrations that support signing in. Multiple payloads combine using an intersect operation. This means the allowed set of workspace IDs can become the empty set if multiple payloads specify conflicting values.
	AllowedExternalIntelligenceWorkspaceIDs *[]string `json:"allowedExternalIntelligenceWorkspaceIDs,omitempty" plist:"allowedExternalIntelligenceWorkspaceIDs,omitempty"`
	// If `false`, the system disables the Enable Restrictions option in the Restrictions UI in Settings. If `false` in iOS 12 and later, the system disables the Enable ScreenTime option in the ScreenTime UI in Settings and disables ScreenTime if already enabled.
	AllowEnablingRestrictions *bool `json:"allowEnablingRestrictions,omitempty" plist:"allowEnablingRestrictions,omitempty"`
	// If `false`, the system removes the Trust Enterprise Developer button in Settings > General > VPN & Device Management, which prevents provisioning apps by universal provisioning profiles. This restriction applies to free developer accounts and enterprise app developers that aren't implicitly trusted by apps that install through MDM. This restriction doesn't revoke previously granted trust.
	AllowEnterpriseAppTrust *bool `json:"allowEnterpriseAppTrust,omitempty" plist:"allowEnterpriseAppTrust,omitempty"`
	// If `false`, the system disables backup of Enterprise books.
	AllowEnterpriseBookBackup *bool `json:"allowEnterpriseBookBackup,omitempty" plist:"allowEnterpriseBookBackup,omitempty"`
	// If `false`, the system disables sync of Enterprise books, notes, and highlights.
	AllowEnterpriseBookMetadataSync *bool `json:"allowEnterpriseBookMetadataSync,omitempty" plist:"allowEnterpriseBookMetadataSync,omitempty"`
	// If `false`, the system disables the Erase All Content and Settings option in the Reset UI.
	AllowEraseContentAndSettings *bool `json:"allowEraseContentAndSettings,omitempty" plist:"allowEraseContentAndSettings,omitempty"`
	// If `false`, the system disables modifications of eSIMs.
	AllowESIMModification *bool `json:"allowESIMModification,omitempty" plist:"allowESIMModification,omitempty"`
	// If `false`, prevents the transfer of an eSIM from the device on which the restriction is installed to a different device.
	AllowESIMOutgoingTransfers *bool `json:"allowESIMOutgoingTransfers,omitempty" plist:"allowESIMOutgoingTransfers,omitempty"`
	// If `false`, the system hides explicit music or video content purchased from the iTunes Store. The system marks explicit content as such by content providers, such as record labels, when sold through the iTunes Store. Explicit content in the News and Podcast apps is also hidden.
	// Requires a supervised device in iOS 13 and later. Support for this restriction on unsupervised devices is deprecated.
	AllowExplicitContent *bool `json:"allowExplicitContent,omitempty" plist:"allowExplicitContent,omitempty"`
	// If `false`, disables the use of external, cloud-based intelligence services with Siri. In iOS, this restriction is temporarily allowed on unsupervised and user enrollments. In a future release, this restriction will require supervision, and will be ignored on unsupervised devices.
	AllowExternalIntelligenceIntegrations *bool `json:"allowExternalIntelligenceIntegrations,omitempty" plist:"allowExternalIntelligenceIntegrations,omitempty"`
	// If `false`, forces external intelligence providers into anonymous mode. If a user is already signed in to an external intelligence provider, applying this restriction signs them out when attempting the next request.
	AllowExternalIntelligenceIntegrationsSignIn *bool `json:"allowExternalIntelligenceIntegrationsSignIn,omitempty" plist:"allowExternalIntelligenceIntegrationsSignIn,omitempty"`
	// If `false`, the system prevents modifying File Sharing setting in System Settings.
	AllowFileSharingModification *bool `json:"allowFileSharingModification,omitempty" plist:"allowFileSharingModification,omitempty"`
	// If `false`, the system prevents connecting to network drives in the Files app.
	AllowFilesNetworkDriveAccess *bool `json:"allowFilesNetworkDriveAccess,omitempty" plist:"allowFilesNetworkDriveAccess,omitempty"`
	// If `false`, the system prevents connecting to any connected USB devices in the Files app.
	AllowFilesUSBDriveAccess *bool `json:"allowFilesUSBDriveAccess,omitempty" plist:"allowFilesUSBDriveAccess,omitempty"`
	// If `false`, the system disables Find My Device in the Find My app.
	AllowFindMyDevice *bool `json:"allowFindMyDevice,omitempty" plist:"allowFindMyDevice,omitempty"`
	// If `false`, the system disables Find My Friends in the Find My app.
	AllowFindMyFriends *bool `json:"allowFindMyFriends,omitempty" plist:"allowFindMyFriends,omitempty"`
	// If `false`, the system disables changes to Find My Friends.
	AllowFindMyFriendsModification *bool `json:"allowFindMyFriendsModification,omitempty" plist:"allowFindMyFriendsModification,omitempty"`
	// If `false`, the system prevents Touch ID, Face ID, or Optic ID from unlocking a device. Support for this restriction on unsupervised devices is deprecated.
	AllowFingerprintForUnlock *bool `json:"allowFingerprintForUnlock,omitempty" plist:"allowFingerprintForUnlock,omitempty"`
	// If `false`, the system prevents the user from modifying Touch ID or Face ID.
	AllowFingerprintModification *bool `json:"allowFingerprintModification,omitempty" plist:"allowFingerprintModification,omitempty"`
	// If `false`, the system disables Game Center, and the system removes its icon from the Home Screen.
	AllowGameCenter *bool `json:"allowGameCenter,omitempty" plist:"allowGameCenter,omitempty"`
	// If `false`, prohibits creating new Genmoji.
	AllowGenmoji *bool `json:"allowGenmoji,omitempty" plist:"allowGenmoji,omitempty"`
	// If `false`, the system disables global background fetch activity when an iOS phone is roaming. Support for this restriction on unsupervised devices is deprecated.
	AllowGlobalBackgroundFetchWhenRoaming *bool `json:"allowGlobalBackgroundFetchWhenRoaming,omitempty" plist:"allowGlobalBackgroundFetchWhenRoaming,omitempty"`
	// If `false`, the system disables host pairing with the exception of the supervision host. If there's no configured supervision host certificate, the system disables all pairing. Host pairing lets the administrator control whether an iOS device can pair with a host Mac or PC.
	AllowHostPairing *bool `json:"allowHostPairing,omitempty" plist:"allowHostPairing,omitempty"`
	// If `false`, prohibits the use of image generation.
	AllowImagePlayground *bool `json:"allowImagePlayground,omitempty" plist:"allowImagePlayground,omitempty"`
	// If `false`, prohibits the use of Image Wand.
	AllowImageWand *bool `json:"allowImageWand,omitempty" plist:"allowImageWand,omitempty"`
	// If `false`, the system prohibits in-app purchasing. Support for this restriction on unsupervised devices is deprecated.
	AllowInAppPurchases *bool `json:"allowInAppPurchases,omitempty" plist:"allowInAppPurchases,omitempty"`
	// If `false`, the system prevents modifying the Internet Sharing setting in System Settings.
	AllowInternetSharingModification *bool `json:"allowInternetSharingModification,omitempty" plist:"allowInternetSharingModification,omitempty"`
	// If `false`, prohibits the use of iPhone Mirroring. In macOS, this prevents the Mac from mirroring any iPhone. In iOS, this prevents the iPhone from mirroring to any Mac.
	AllowiPhoneMirroring *bool `json:"allowiPhoneMirroring,omitempty" plist:"allowiPhoneMirroring,omitempty"`
	// If `false`, the system disallows iPhone widgets on a Mac that signs in with the same Apple Account for iCloud.
	AllowiPhoneWidgetsOnMac *bool `json:"allowiPhoneWidgetsOnMac,omitempty" plist:"allowiPhoneWidgetsOnMac,omitempty"`
	// If `false`, the system disables the iTunes Music Store and removes its icon from the Home Screen. Users can't preview, purchase, or download content. Requires a supervised device in iOS 13 and later.
	AllowiTunes *bool `json:"allowiTunes,omitempty" plist:"allowiTunes,omitempty"`
	// If `false`, the system disables iTunes file sharing services.
	AllowiTunesFileSharing *bool `json:"allowiTunesFileSharing,omitempty" plist:"allowiTunesFileSharing,omitempty"`
	// If `false`, the system disables keyboard shortcuts.
	AllowKeyboardShortcuts *bool `json:"allowKeyboardShortcuts,omitempty" plist:"allowKeyboardShortcuts,omitempty"`
	// If present, the system only shows or can launch apps with bundle IDs in the array. Include the value `com.apple.webapp` to allow all webclips. This applies to App Store apps, marketplace apps, and locally installed apps (using Configurator, Xcode, and so forth).
	AllowListedAppBundleIDs *[]string `json:"allowListedAppBundleIDs,omitempty" plist:"allowListedAppBundleIDs,omitempty"`
	// If `false`, the system disables live voicemail on the device.
	AllowLiveVoicemail *bool `json:"allowLiveVoicemail,omitempty" plist:"allowLiveVoicemail,omitempty"`
	// If `false`, the system prevents creating users in System Settings.
	AllowLocalUserCreation *bool `json:"allowLocalUserCreation,omitempty" plist:"allowLocalUserCreation,omitempty"`
	// If `false`, the system prevents Control Center from appearing on the Lock Screen.
	AllowLockScreenControlCenter *bool `json:"allowLockScreenControlCenter,omitempty" plist:"allowLockScreenControlCenter,omitempty"`
	// If `false`, the system disables the Notifications history view on the Lock Screen, so users can't view past notifications. However, they can still see notifications when they arrive.
	AllowLockScreenNotificationsView *bool `json:"allowLockScreenNotificationsView,omitempty" plist:"allowLockScreenNotificationsView,omitempty"`
	// If `false`, the system disables the Today view in Notification Center on the Lock Screen.
	AllowLockScreenTodayView *bool `json:"allowLockScreenTodayView,omitempty" plist:"allowLockScreenTodayView,omitempty"`
	// If `false`, the system disables Mail Privacy Protection on the device.
	AllowMailPrivacyProtection *bool `json:"allowMailPrivacyProtection,omitempty" plist:"allowMailPrivacyProtection,omitempty"`
	// If `false`, disables smart replies in Mail.
	AllowMailSmartReplies *bool `json:"allowMailSmartReplies,omitempty" plist:"allowMailSmartReplies,omitempty"`
	// If `false`, disables the ability to create summaries of email messages manually. This doesn't affect automatic summary generation.
	AllowMailSummary *bool `json:"allowMailSummary,omitempty" plist:"allowMailSummary,omitempty"`
	// If `false`, the system prevents managed apps from using iCloud sync.
	AllowManagedAppsCloudSync *bool `json:"allowManagedAppsCloudSync,omitempty" plist:"allowManagedAppsCloudSync,omitempty"`
	// If `true`, the system allows managed apps to write contacts to unmanaged accounts. If `allowOpenFromManagedToUnmanaged` is `true`, this restriction has no effect.
	// > Important:
	// > Use MDM to install profiles that contain this restriction.
	AllowManagedToWriteUnmanagedContacts *bool `json:"allowManagedToWriteUnmanagedContacts,omitempty" plist:"allowManagedToWriteUnmanagedContacts,omitempty"`
	// If `false`, the system prevents installation of alternative marketplace apps from the web and prevents any installed alternative marketplace apps from installing apps.
	AllowMarketplaceAppInstallation *bool `json:"allowMarketplaceAppInstallation,omitempty" plist:"allowMarketplaceAppInstallation,omitempty"`
	// If `false`, prevents modification of Media Sharing settings.
	AllowMediaSharingModification *bool `json:"allowMediaSharingModification,omitempty" plist:"allowMediaSharingModification,omitempty"`
	// If `false`, the system prohibits multiplayer gaming.
	AllowMultiplayerGaming *bool `json:"allowMultiplayerGaming,omitempty" plist:"allowMultiplayerGaming,omitempty"`
	// If `false`, the system disables the Music service, and the Music app reverts to classic mode.
	AllowMusicService *bool `json:"allowMusicService,omitempty" plist:"allowMusicService,omitempty"`
	// If `false`, the system disables News.
	AllowNews *bool `json:"allowNews,omitempty" plist:"allowNews,omitempty"`
	// If `false`, the system disables NFC.
	AllowNFC *bool `json:"allowNFC,omitempty" plist:"allowNFC,omitempty"`
	// If `false`, disables transcription in Notes.
	AllowNotesTranscription *bool `json:"allowNotesTranscription,omitempty" plist:"allowNotesTranscription,omitempty"`
	// If `false`, disables transcription summarization in Notes.
	AllowNotesTranscriptionSummary *bool `json:"allowNotesTranscriptionSummary,omitempty" plist:"allowNotesTranscriptionSummary,omitempty"`
	// If `false`, the system disables modification of notification settings.
	AllowNotificationsModification *bool `json:"allowNotificationsModification,omitempty" plist:"allowNotificationsModification,omitempty"`
	// If `false`, documents in managed apps and accounts open only in other managed apps and accounts.
	AllowOpenFromManagedToUnmanaged *bool `json:"allowOpenFromManagedToUnmanaged,omitempty" plist:"allowOpenFromManagedToUnmanaged,omitempty"`
	// If `false`, documents in unmanaged apps and accounts open only in other unmanaged apps and accounts.
	AllowOpenFromUnmanagedToManaged *bool `json:"allowOpenFromUnmanagedToManaged,omitempty" plist:"allowOpenFromUnmanagedToManaged,omitempty"`
	// If `false`, the system disables over-the-air PKI updates. Setting this restriction to `false` doesn't disable CRL and OCSP checks.
	AllowOTAPKIUpdates *bool `json:"allowOTAPKIUpdates,omitempty" plist:"allowOTAPKIUpdates,omitempty"`
	// If `false`, the system disables pairing with an Apple Watch, and the system unpairs any currently paired Apple Watch and erases its content.
	AllowPairedWatch *bool `json:"allowPairedWatch,omitempty" plist:"allowPairedWatch,omitempty"`
	// If `false`, the system hides Passbook notifications from the Lock Screen.
	AllowPassbookWhileLocked *bool `json:"allowPassbookWhileLocked,omitempty" plist:"allowPassbookWhileLocked,omitempty"`
	// If `false`, the system prevents adding, changing, or removing the passcode. The system ignores this restriction on Shared iPad.
	AllowPasscodeModification *bool `json:"allowPasscodeModification,omitempty" plist:"allowPasscodeModification,omitempty"`
	// If `false`, the system disables:
	// - The AutoFill Passwords feature in iOS, with Keychain and third-party password managers
	// - Prompting the user to use a saved password in Safari or in apps
	// - Automatic strong passwords
	// - Suggesting strong passwords to users
	// However, if `false`, the system doesn't prevent AutoFill for contact info and credit cards in Safari.
	AllowPasswordAutoFill *bool `json:"allowPasswordAutoFill,omitempty" plist:"allowPasswordAutoFill,omitempty"`
	// If `false`, the system disables requesting passwords from nearby devices.
	AllowPasswordProximityRequests *bool `json:"allowPasswordProximityRequests,omitempty" plist:"allowPasswordProximityRequests,omitempty"`
	// If `false`, the system disables sharing passwords with the AirDrop passwords feature, or with the Passwords app.
	AllowPasswordSharing *bool `json:"allowPasswordSharing,omitempty" plist:"allowPasswordSharing,omitempty"`
	// If `false`, the system disables modifications of the personal hotspot setting.
	AllowPersonalHotspotModification *bool `json:"allowPersonalHotspotModification,omitempty" plist:"allowPersonalHotspotModification,omitempty"`
	// If false, prevents the system from generating text in the user's handwriting.
	AllowPersonalizedHandwritingResults *bool `json:"allowPersonalizedHandwritingResults,omitempty" plist:"allowPersonalizedHandwritingResults,omitempty"`
	// If `false`, the system disables Photo Stream.
	AllowPhotoStream *bool `json:"allowPhotoStream,omitempty" plist:"allowPhotoStream,omitempty"`
	// If `false`, the system disables podcasts.
	AllowPodcasts *bool `json:"allowPodcasts,omitempty" plist:"allowPodcasts,omitempty"`
	// If `false`, the system disables predictive keyboards.
	AllowPredictiveKeyboard *bool `json:"allowPredictiveKeyboard,omitempty" plist:"allowPredictiveKeyboard,omitempty"`
	// If `false`, the system prevents modifying Printer Sharing settings in System Settings.
	AllowPrinterSharingModification *bool `json:"allowPrinterSharingModification,omitempty" plist:"allowPrinterSharingModification,omitempty"`
	// If `false`, disables the prompt to set up new devices that are nearby.
	AllowProximitySetupToNewDevice *bool `json:"allowProximitySetupToNewDevice,omitempty" plist:"allowProximitySetupToNewDevice,omitempty"`
	// If `false`, the system disables Apple Music Radio.
	AllowRadioService *bool `json:"allowRadioService,omitempty" plist:"allowRadioService,omitempty"`
	// If `false`, the system prohibits installation of Background Security Improvements.
	AllowRapidSecurityResponseInstallation *bool `json:"allowRapidSecurityResponseInstallation,omitempty" plist:"allowRapidSecurityResponseInstallation,omitempty"`
	// If `false`, the system prohibits removal of Background Security Improvements.
	AllowRapidSecurityResponseRemoval *bool `json:"allowRapidSecurityResponseRemoval,omitempty" plist:"allowRapidSecurityResponseRemoval,omitempty"`
	// If `false`, prevents the use of RCS messaging.
	AllowRCSMessaging *bool `json:"allowRCSMessaging,omitempty" plist:"allowRCSMessaging,omitempty"`
	// If `false`, the system prevents modifying Remote Apple Events Sharing settings in System Settings.
	AllowRemoteAppleEventsModification *bool `json:"allowRemoteAppleEventsModification,omitempty" plist:"allowRemoteAppleEventsModification,omitempty"`
	// If `false`, the system disables pairing Apple TV for use with the Control Center widget.
	AllowRemoteAppPairing *bool `json:"allowRemoteAppPairing,omitempty" plist:"allowRemoteAppPairing,omitempty"`
	// If `false`, the system disables remote screen observation by the Classroom app. Nest this key beneath `allowScreenShot` as a subrestriction. If `allowScreenShot` is `false`, the Classroom app doesn't observe remote screens. Requires a supervised device until iOS 13 and macOS 10.15. Allowed for user enrollments in macOS 12 and later.
	AllowRemoteScreenObservation *bool `json:"allowRemoteScreenObservation,omitempty" plist:"allowRemoteScreenObservation,omitempty"`
	// If `false`, the system disables the Safari web browser app, and the system removes its icon from the Home Screen. This setting also prevents users from opening web clips. Requires a supervised device in iOS 13 and later.
	AllowSafari *bool `json:"allowSafari,omitempty" plist:"allowSafari,omitempty"`
	// If `false`, the system disables the ability to clear browsing history in Safari.
	AllowSafariHistoryClearing *bool `json:"allowSafariHistoryClearing,omitempty" plist:"allowSafariHistoryClearing,omitempty"`
	// If `false`, the system disables the ability to use private browsing in Safari.
	AllowSafariPrivateBrowsing *bool `json:"allowSafariPrivateBrowsing,omitempty" plist:"allowSafariPrivateBrowsing,omitempty"`
	// If `false`, the system disables the ability to summarize content in Safari.
	AllowSafariSummary *bool `json:"allowSafariSummary,omitempty" plist:"allowSafariSummary,omitempty"`
	// If `false`, the system prohibits the connection to and use of satellite services.
	AllowSatelliteConnection *bool `json:"allowSatelliteConnection,omitempty" plist:"allowSatelliteConnection,omitempty"`
	// If `false`, the system disables saving a screenshot of the display and capturing a screen recording. It also disables the Classroom app from observing remote screens.
	AllowScreenShot *bool `json:"allowScreenShot,omitempty" plist:"allowScreenShot,omitempty"`
	// If `false`, the system makes temporary sessions unavailable on Shared iPad.
	AllowSharedDeviceTemporarySession *bool `json:"allowSharedDeviceTemporarySession,omitempty" plist:"allowSharedDeviceTemporarySession,omitempty"`
	// If `false`, the system disables Shared Photo Stream. Support for this restriction on unsupervised devices is deprecated.
	AllowSharedStream *bool `json:"allowSharedStream,omitempty" plist:"allowSharedStream,omitempty"`
	// If `false`, the system disables the keyboard spell checker.
	AllowSpellCheck *bool `json:"allowSpellCheck,omitempty" plist:"allowSpellCheck,omitempty"`
	// If `false`, the system disables Spotlight Internet search results in Siri Suggestions. Support for this restriction on unsupervised devices is deprecated.
	AllowSpotlightInternetResults *bool `json:"allowSpotlightInternetResults,omitempty" plist:"allowSpotlightInternetResults,omitempty"`
	// If `false`, the system prevents modification of Startup Disk settings in System Settings.
	AllowStartupDiskModification *bool `json:"allowStartupDiskModification,omitempty" plist:"allowStartupDiskModification,omitempty"`
	// If `false`, the system disables the removal of system apps from the device.
	AllowSystemAppRemoval *bool `json:"allowSystemAppRemoval,omitempty" plist:"allowSystemAppRemoval,omitempty"`
	// If `false`, the system prevents modification of Time Machine settings in System Settings. This restriction is not supported on the user channel.
	AllowTimeMachineBackup *bool `json:"allowTimeMachineBackup,omitempty" plist:"allowTimeMachineBackup,omitempty"`
	// If `false`, the system disables the App Store and removes its icon from the Home Screen. However, users can continue to install or update their apps either locally (via Configurator, Xcode, and so forth), or using alternative marketplace apps.
	// In iOS 10 and later, MDM commands can override this restriction.
	AllowUIAppInstallation *bool `json:"allowUIAppInstallation,omitempty" plist:"allowUIAppInstallation,omitempty"`
	// If `false`, the system prohibits the user from installing configuration profiles and certificates interactively.
	AllowUIConfigurationProfileInstallation *bool `json:"allowUIConfigurationProfileInstallation,omitempty" plist:"allowUIConfigurationProfileInstallation,omitempty"`
	// If `false`, the system disables Universal Control.
	AllowUniversalControl *bool `json:"allowUniversalControl,omitempty" plist:"allowUniversalControl,omitempty"`
	// If `true`, the system allows unmanaged apps to read from managed contacts accounts. If `allowOpenFromManagedToUnmanaged` is `true`, this restriction has no effect.
	// > Important:
	// > Use MDM to install profiles that contain this restriction.
	AllowUnmanagedToReadManagedContacts *bool `json:"allowUnmanagedToReadManagedContacts,omitempty" plist:"allowUnmanagedToReadManagedContacts,omitempty"`
	// If `true`, the system allows unpaired devices to boot devices into recovery.
	AllowUnpairedExternalBootToRecovery *bool `json:"allowUnpairedExternalBootToRecovery,omitempty" plist:"allowUnpairedExternalBootToRecovery,omitempty"`
	// If `false`, the system automatically rejects untrusted HTTPS certificates without prompting the user.
	AllowUntrustedTLSPrompt *bool `json:"allowUntrustedTLSPrompt,omitempty" plist:"allowUntrustedTLSPrompt,omitempty"`
	// If `false`, the system allows iOS devices to always connect to USB accessories while locked. In macOS, allows new USB and Thunderbolt accessories, and SD cards to connect without authorization. If the system has Lockdown mode enabled, it ignores this value. This restriction is not supported on the user channel.
	AllowUSBRestrictedMode *bool `json:"allowUSBRestrictedMode,omitempty" plist:"allowUSBRestrictedMode,omitempty"`
	// If `false`, the system hides the FaceTime app. Requires a supervised device in iOS 13 and later.
	AllowVideoConferencing *bool `json:"allowVideoConferencing,omitempty" plist:"allowVideoConferencing,omitempty"`
	// If `false`, disables the ability for a remote FaceTime session to request control of the device.
	AllowVideoConferencingRemoteControl *bool `json:"allowVideoConferencingRemoteControl,omitempty" plist:"allowVideoConferencingRemoteControl,omitempty"`
	// If `false`, the system disables visual intelligence summarization.
	AllowVisualIntelligenceSummary *bool `json:"allowVisualIntelligenceSummary,omitempty" plist:"allowVisualIntelligenceSummary,omitempty"`
	// If `false`, the system disables voice dialing if the device is locked with a passcode.
	AllowVoiceDialing *bool `json:"allowVoiceDialing,omitempty" plist:"allowVoiceDialing,omitempty"`
	// If `false`, the system allows only managed apps to create VPN configurations. Prior to iOS 18, the system also allows unmanaged apps to create VPN configurations.
	AllowVPNCreation *bool `json:"allowVPNCreation,omitempty" plist:"allowVPNCreation,omitempty"`
	// If `false`, the system prevents changing the wallpaper.
	AllowWallpaperModification *bool `json:"allowWallpaperModification,omitempty" plist:"allowWallpaperModification,omitempty"`
	// If `false`, the device prevents installation of apps directly from the web.
	AllowWebDistributionAppInstallation *bool `json:"allowWebDistributionAppInstallation,omitempty" plist:"allowWebDistributionAppInstallation,omitempty"`
	// If `false`, disables Apple Intelligence writing tools.
	AllowWritingTools *bool `json:"allowWritingTools,omitempty" plist:"allowWritingTools,omitempty"`
	// If present, the system allows apps identified by the bundle IDs listed in the array to autonomously enter Single App Mode.
	AutonomousSingleAppModePermittedAppIDs *[]string `json:"autonomousSingleAppModePermittedAppIDs,omitempty" plist:"autonomousSingleAppModePermittedAppIDs,omitempty"`
	// Use `blockedAppBundleIDs` instead.
	BlacklistedAppBundleIDs *[]string `json:"blacklistedAppBundleIDs,omitempty" plist:"blacklistedAppBundleIDs,omitempty"`
	// If present, the system prevents showing or launching apps with bundle IDs in the array. Include the value `com.apple.webapp` to restrict all webclips. This applies to App Store apps, marketplace apps, and locally installed apps (using Configurator, Xcode, and so forth).
	// > Note:
	// > Denying system apps may disable other functionality. For example, denying the App Store app may prevent users from accepting the terms and conditions for the user-based Volume Purchase Program (VPP).
	BlockedAppBundleIDs *[]string `json:"blockedAppBundleIDs,omitempty" plist:"blockedAppBundleIDs,omitempty"`
	// An array of strings representing ICCIDs of cellular plans. The device prevents use of any matching cellular networks in iMessage and FaceTime. The array must contain no more than 4 ICCID strings.
	DeniedICCIDsForiMessageFaceTime *[]string `json:"deniedICCIDsForiMessageFaceTime,omitempty" plist:"deniedICCIDsForiMessageFaceTime,omitempty"`
	// An array of strings representing ICCIDs of cellular plans. The device prevents use of any matching cellular networks with RCS messaging. The array must contain no more than 4 ICCID strings.
	DeniedICCIDsForRCS *[]string `json:"deniedICCIDsForRCS,omitempty" plist:"deniedICCIDsForRCS,omitempty"`
	// The value, in seconds, after which the fingerprint unlock requires a password to authenticate. The default value is 48 hours.
	EnforcedFingerprintTimeout *int64 `default:"172800" json:"enforcedFingerprintTimeout,omitempty" plist:"enforcedFingerprintTimeout,omitempty"`
	// How many days to delay a software update on the device. With this restriction in place, the user doesn't see a software update until the specified number of days after the software update release date. The restrictions `forceDelayedAppSoftwareUpdates` and `forceDelayedSoftwareUpdates` use this value.
	EnforcedSoftwareUpdateDelay *int64 `default:"30" json:"enforcedSoftwareUpdateDelay,omitempty" plist:"enforcedSoftwareUpdateDelay,omitempty"`
	// This restriction allows the administrator to set the number of days to delay a major software upgrade on the device. When this restriction is in place, the user sees a software upgrade only after the specified delay after the release of the software upgrade. This value controls the delay for `forceDelayedMajorSoftwareUpdates`.
	EnforcedSoftwareUpdateMajorOSDeferredInstallDelay *int64 `default:"30" json:"enforcedSoftwareUpdateMajorOSDeferredInstallDelay,omitempty" plist:"enforcedSoftwareUpdateMajorOSDeferredInstallDelay,omitempty"`
	// This restriction allows the administrator to set the number of days to delay a minor OS software update on the device. When this restriction is in place, the user sees a software update only after the specified delay after the release of the software update. This value controls the delay for `forceDelayedSoftwareUpdates`.
	EnforcedSoftwareUpdateMinorOSDeferredInstallDelay *int64 `default:"30" json:"enforcedSoftwareUpdateMinorOSDeferredInstallDelay,omitempty" plist:"enforcedSoftwareUpdateMinorOSDeferredInstallDelay,omitempty"`
	// This restriction allows the administrator to set the number of days to delay an app software update on the device. When this restriction is in place, the user sees a non-OS software update only after the specified delay after the release of the software. This value controls the delay for `forceDelayedAppSoftwareUpdates`.
	EnforcedSoftwareUpdateNonOSDeferredInstallDelay *int64 `default:"30" json:"enforcedSoftwareUpdateNonOSDeferredInstallDelay,omitempty" plist:"enforcedSoftwareUpdateNonOSDeferredInstallDelay,omitempty"`
	// If `true`, the system considers AirDrop to be an unmanaged drop target.
	ForceAirDropUnmanaged *bool `json:"forceAirDropUnmanaged,omitempty" plist:"forceAirDropUnmanaged,omitempty"`
	// If `true`, the system forces all devices sending AirPlay requests to this device to use a pairing password. This key isn't supported in tvOS 10.2 and later. Use the AirPlay Security Payload instead.
	ForceAirPlayIncomingRequestsPairingPassword *bool `json:"forceAirPlayIncomingRequestsPairingPassword,omitempty" plist:"forceAirPlayIncomingRequestsPairingPassword,omitempty"`
	// If `true`, the system forces all devices receiving AirPlay requests from this device to use a pairing password.
	ForceAirPlayOutgoingRequestsPairingPassword *bool `json:"forceAirPlayOutgoingRequestsPairingPassword,omitempty" plist:"forceAirPlayOutgoingRequestsPairingPassword,omitempty"`
	// If `true`, the system requires trusted certificates for TLS printing communication.
	ForceAirPrintTrustedTLSRequirement *bool `json:"forceAirPrintTrustedTLSRequirement,omitempty" plist:"forceAirPrintTrustedTLSRequirement,omitempty"`
	// If `true`, the system forces the use of the profanity filter for Siri and dictation. Requires a supervised device in iOS.
	ForceAssistantProfanityFilter *bool `json:"forceAssistantProfanityFilter,omitempty" plist:"forceAssistantProfanityFilter,omitempty"`
	// If `true`, the user needs to authenticate before the system can autofill passwords or credit card information in Safari and apps. If this restriction isn't enforced, the user can toggle this feature in Settings. Only supported on devices with Face ID or Touch ID.
	ForceAuthenticationBeforeAutoFill *bool `json:"forceAuthenticationBeforeAutoFill,omitempty" plist:"forceAuthenticationBeforeAutoFill,omitempty"`
	// If `true`, the system enables the Set Automatically feature in Date & Time and the user can't disable it. The system updates the device's time zone only when the device can determine its location using a cellular connection or Wi-Fi with location services enabled.
	ForceAutomaticDateAndTime *bool `json:"forceAutomaticDateAndTime,omitempty" plist:"forceAutomaticDateAndTime,omitempty"`
	// If `true`, then the system bypasses the presentation of a screen capture alert.
	ForceBypassScreenCaptureAlert *bool `json:"forceBypassScreenCaptureAlert,omitempty" plist:"forceBypassScreenCaptureAlert,omitempty"`
	// If `true`, the system automatically gives permission to the teacher's requests without prompting the student.
	ForceClassroomAutomaticallyJoinClasses *bool `json:"forceClassroomAutomaticallyJoinClasses,omitempty" plist:"forceClassroomAutomaticallyJoinClasses,omitempty"`
	// If `true`, a student enrolled in an unmanaged course through Classroom needs to request permission from the teacher to leave the course.
	ForceClassroomRequestPermissionToLeaveClasses *bool `json:"forceClassroomRequestPermissionToLeaveClasses,omitempty" plist:"forceClassroomRequestPermissionToLeaveClasses,omitempty"`
	// If `true`, the system allows the teacher to lock apps or the device without prompting the student.
	ForceClassroomUnpromptedAppAndDeviceLock *bool `json:"forceClassroomUnpromptedAppAndDeviceLock,omitempty" plist:"forceClassroomUnpromptedAppAndDeviceLock,omitempty"`
	// If `true` and `ScreenObservationPermissionModificationAllowed` is also `true` in the Education payload, a student enrolled in a managed course through the Classroom app automatically gives permission to that course teacher's requests to observe the student's screen without prompting the student.
	ForceClassroomUnpromptedScreenObservation *bool `json:"forceClassroomUnpromptedScreenObservation,omitempty" plist:"forceClassroomUnpromptedScreenObservation,omitempty"`
	// If `true`, the system delays user visibility of non-OS software updates. Control visibility of operating system updates through `forceDelayedSoftwareUpdates`. The delay is 30 days unless you set `enforcedSoftwareUpdateDelay` to another value.
	ForceDelayedAppSoftwareUpdates *bool `json:"forceDelayedAppSoftwareUpdates,omitempty" plist:"forceDelayedAppSoftwareUpdates,omitempty"`
	// If `true`, the system delays user visibility of major OS updates.
	ForceDelayedMajorSoftwareUpdates *bool `json:"forceDelayedMajorSoftwareUpdates,omitempty" plist:"forceDelayedMajorSoftwareUpdates,omitempty"`
	// If `true`, the system delays user visibility of software updates. In macOS, the system allows seed build updates without delay. The delay is 30 days unless you set `enforcedSoftwareUpdateDelay` to another value.
	ForceDelayedSoftwareUpdates *bool `json:"forceDelayedSoftwareUpdates,omitempty" plist:"forceDelayedSoftwareUpdates,omitempty"`
	// If `true`, the system encrypts all backups.
	ForceEncryptedBackup *bool `json:"forceEncryptedBackup,omitempty" plist:"forceEncryptedBackup,omitempty"`
	// If `true`, the system forces the user to enter their iTunes password for each transaction.
	ForceITunesStorePasswordEntry *bool `json:"forceITunesStorePasswordEntry,omitempty" plist:"forceITunesStorePasswordEntry,omitempty"`
	// If `true`, the system limits ad tracking. Additionally, it disables app tracking and the Allow Apps to Request to Track setting.
	ForceLimitAdTracking *bool `json:"forceLimitAdTracking,omitempty" plist:"forceLimitAdTracking,omitempty"`
	// If `true`, the system disables connections to Siri servers for the purposes of dictation.
	ForceOnDeviceOnlyDictation *bool `json:"forceOnDeviceOnlyDictation,omitempty" plist:"forceOnDeviceOnlyDictation,omitempty"`
	// If `true`, the device can't connect to Siri servers for the purposes of translation.
	ForceOnDeviceOnlyTranslation *bool `json:"forceOnDeviceOnlyTranslation,omitempty" plist:"forceOnDeviceOnlyTranslation,omitempty"`
	// If `true`, the system preserves eSIM when it erases the device due to too many failed password attempts or the Erase All Content and Settings option in Settings > General > Reset.
	// > Note:
	// > The system doesn't preserve eSIM if Find My initiates erasing the device.
	ForcePreserveESIMOnErase *bool `json:"forcePreserveESIMOnErase,omitempty" plist:"forcePreserveESIMOnErase,omitempty"`
	// If `true`, the system forces a paired Apple Watch to use Wrist Detection.
	ForceWatchWristDetection *bool `json:"forceWatchWristDetection,omitempty" plist:"forceWatchWristDetection,omitempty"`
	// If `true`, the system prevents turning off Wi-Fi in Settings or Control Center, even by entering or leaving Airplane Mode. It doesn't prevent selecting which Wi-Fi network to use. and later.
	ForceWiFiPowerOn *bool `json:"forceWiFiPowerOn,omitempty" plist:"forceWiFiPowerOn,omitempty"`
	// If `true`, the system limits the device to only join Wi-Fi networks set up through a configuration profile.
	ForceWiFiToAllowedNetworksOnly *bool `json:"forceWiFiToAllowedNetworksOnly,omitempty" plist:"forceWiFiToAllowedNetworksOnly,omitempty"`
	// Use `forceWiFiToAllowedNetworksOnly` instead.
	ForceWiFiWhitelisting *bool `json:"forceWiFiWhitelisting,omitempty" plist:"forceWiFiWhitelisting,omitempty"`
	// The maximum level of app content allowed on the device. Preinstalled (first-party) apps ignore this restriction.
	// Possible values, with the U.S. description of the rating level:
	// - `1000`: All
	// - `600`: 17+
	// - `300`: 12+
	// - `200`: 9+
	// - `100`: 4+
	// - `0`: None
	// Age bands and the number of discrete age values vary by region, but the values are consistent across regions. For example, in a region that defines rating level 14+, its value is guaranteed to be larger than 300 (12+) and smaller than 600 (17+). Also, the value of rating level 15+ is guaranteed to be larger than the assigned value of rating level 14+. For more information about age ratings, see [Age ratings values and definitions](https://developer.apple.com/help/app-store-connect/reference/age-ratings-values-and-definitions).
	// Examples of values in other regions include:
	// - `1000`: All
	// - `621`: 21+
	// - `620`: 20+
	// - `619`: 19+
	// - `618`: 18+
	// - `600`: 17+
	// - `416`: 16+
	// - `415`: 15+
	// - `314`: 14+
	// - `313`: 13+
	// - `300`: 12+
	// - `211`: 11+
	// - `210`: 10+
	// - `200`: 9+
	// - `108`: 8+
	// - `107`: 7+
	// - `106`: 6+
	// - `105`: 5+
	// - `100`: 4+
	// - `3`: 3+
	// - `2`: 2+
	// - `1`: 1+
	// - `0`: None
	// This restriction will require supervision in a future release.
	RatingApps *int64 `default:"1000" json:"ratingApps,omitempty" plist:"ratingApps,omitempty"`
	// If present, the system exempts apps with bundle IDs in the array from age-based rating restrictions. The system uses intersection combine rules to combine multiple payloads and any exceptions that parental control apps provide, including ScreenTime.
	RatingAppsExemptedBundleIDs *[]string `json:"ratingAppsExemptedBundleIDs,omitempty" plist:"ratingAppsExemptedBundleIDs,omitempty"`
	// The maximum level of movie content allowed on the device. Support for this restriction on unsupervised devices is deprecated.
	// Possible values, with the U.S. description of the rating level:
	// - `1000`: All
	// - `500`: NC-17
	// - `400`: R
	// - `300`: PG-13
	// - `200`: PG
	// - `100`: G
	// - `0`: None
	RatingMovies *int64 `default:"1000" json:"ratingMovies,omitempty" plist:"ratingMovies,omitempty"`
	// The two-letter key that profile tools use to display the proper ratings for the given region. The client doesn't recognize or report this data.
	RatingRegion *RatingRegion `json:"ratingRegion,omitempty" plist:"ratingRegion,omitempty"`
	// The maximum level of TV content allowed on the device. Support for this restriction on unsupervised devices is deprecated.
	// Possible values, with the U.S. description of the rating level:
	// - `1000`: All
	// - `600`: TV-MA
	// - `500`: TV-14
	// - `400`: TV-PG
	// - `300`: TV-G
	// - `200`: TV-Y7
	// - `100`: TV-Y
	// - `0`: None
	RatingTVShows *int64 `default:"1000" json:"ratingTVShows,omitempty" plist:"ratingTVShows,omitempty"`
	// If `true`, copy-and-paste functionality is limited by the `allowOpenFromManagedToUnmanaged` and `allowOpenFromUnmanagedToManaged` restrictions.
	RequireManagedPasteboard *bool `json:"requireManagedPasteboard,omitempty" plist:"requireManagedPasteboard,omitempty"`
	// Defines the conditions under which the device accepts cookies. The user-facing settings changed in iOS 11, although the possible values remain the same. Support for this restriction on unsupervised devices is deprecated. Allowed values:
	// - `0`: Enables Prevent Cross-Site Tracking and Block All Cookies, and the user canʼt disable either setting.
	// - `1` or `1.5`: Enables Prevent Cross-Site Tracking, and the user canʼt disable it. Doesn't enable Block All Cookies, but the user can enable it.
	// - `2`: Enables Prevent Cross-Site Tracking, but doesn't enable Block All Cookies. The user can toggle either setting.
	SafariAcceptCookies *SafariAcceptCookies `default:"2" json:"safariAcceptCookies,omitempty" plist:"safariAcceptCookies,omitempty"`
	// If `false`, the system disables Safari AutoFill for passwords, contact info, and credit cards, and also prevents using the Keychain for AutoFill. Requires a supervised device in iOS 13 and later.
	// > Note:
	// > The system still allows third-party password managers, and apps can use AutoFill.
	SafariAllowAutoFill *bool `json:"safariAllowAutoFill,omitempty" plist:"safariAllowAutoFill,omitempty"`
	// If `false`, Safari doesn't execute JavaScript. This restriction will require supervision in a future release.
	SafariAllowJavaScript *bool `json:"safariAllowJavaScript,omitempty" plist:"safariAllowJavaScript,omitempty"`
	// If `false`, Safari doesn't allow pop-up windows. Support for this restriction on unsupervised devices is deprecated.
	SafariAllowPopups *bool `json:"safariAllowPopups,omitempty" plist:"safariAllowPopups,omitempty"`
	// If `true`, the system enables Safari fraud warning.
	SafariForceFraudWarning *bool `json:"safariForceFraudWarning,omitempty" plist:"safariForceFraudWarning,omitempty"`
	// Use `allowListedAppBundleIDs` instead.
	WhitelistedAppBundleIDs *[]string `json:"whitelistedAppBundleIDs,omitempty" plist:"whitelistedAppBundleIDs,omitempty"`
}

func (p *Restrictions) PayloadType() string {
	return "com.apple.applicationaccess"
}

// The two-letter key that profile tools use to display the proper ratings for the given region. The client doesn't recognize or report this data.
type RatingRegion string

const (
	RatingRegionUs RatingRegion = "us"
	RatingRegionAu RatingRegion = "au"
	RatingRegionCa RatingRegion = "ca"
	RatingRegionDe RatingRegion = "de"
	RatingRegionFr RatingRegion = "fr"
	RatingRegionIe RatingRegion = "ie"
	RatingRegionJp RatingRegion = "jp"
	RatingRegionNz RatingRegion = "nz"
	RatingRegionGb RatingRegion = "gb"
)

// Defines the conditions under which the device accepts cookies. The user-facing settings changed in iOS 11, although the possible values remain the same. Support for this restriction on unsupervised devices is deprecated. Allowed values:
// - `0`: Enables Prevent Cross-Site Tracking and Block All Cookies, and the user canʼt disable either setting.
// - `1` or `1.5`: Enables Prevent Cross-Site Tracking, and the user canʼt disable it. Doesn't enable Block All Cookies, but the user can enable it.
// - `2`: Enables Prevent Cross-Site Tracking, but doesn't enable Block All Cookies. The user can toggle either setting.
type SafariAcceptCookies float64

const (
	SafariAcceptCookies0  SafariAcceptCookies = 0.0
	SafariAcceptCookies1  SafariAcceptCookies = 1.0
	SafariAcceptCookies15 SafariAcceptCookies = 1.5
	SafariAcceptCookies2  SafariAcceptCookies = 2.0
)

// The payload that configures macOS App Store restrictions.
// Use this payload to set restrictions used by the Mac App Store.
type AppStore struct {
	*CommonPayloadKeys
	// If `true`, the system restricts app installations to admin users only. Deprecated in macOS 10.14. Use the `com.apple.SoftwareUpdate` payload key `restrict-software-update-require-admin-to-install` instead.
	RestrictStoreRequireAdminToInstall *bool `json:"restrict-store-require-admin-to-install,omitempty" plist:"restrict-store-require-admin-to-install,omitempty"`
	// If `true`, the system prevents App Store from launching. Available in macOS 10.14 and later. Restricts installations to software updates only in macOS 10.10 through 10.13.
	RestrictStoreSoftwareupdateOnly *bool `json:"restrict-store-softwareupdate-only,omitempty" plist:"restrict-store-softwareupdate-only,omitempty"`
	// If `true`, the system disables app adoption by users. Available in macOS 10.10 and later.
	RestrictStoreDisableAppAdoption *bool `json:"restrict-store-disable-app-adoption,omitempty" plist:"restrict-store-disable-app-adoption,omitempty"`
	// If `true`, the system disables software update notifications. Available in macOS 10.10 and later.
	DisableSoftwareUpdateNotifications *bool `json:"DisableSoftwareUpdateNotifications,omitempty" plist:"DisableSoftwareUpdateNotifications,omitempty"`
}

func (p *AppStore) PayloadType() string {
	return "com.apple.appstore"
}

// The payload that configures Autonomous Single App mode.
type AutonomousSingleAppMode struct {
	*CommonPayloadKeys
	// An array of dictionaries that specifies the apps that the system grants access to the Accessibility APIs.
	AllowedApplications []*AllowedApplicationsItem `json:"AllowedApplications" plist:"AllowedApplications" required:"true"`
}

func (p *AutonomousSingleAppMode) PayloadType() string {
	return "com.apple.asam"
}

// A dictionary that specifies an app that can be granted access to the Accessibilty APIs.
type AllowedApplicationsItem struct {
	// The unique bundle identifier. If two dictionaries contain the same `BundleIdentifier` value but a different `TeamIdentifier` value, an error occurs and the profile won't be installed.
	BundleIdentifier string `json:"BundleIdentifier" plist:"BundleIdentifier" required:"true"`
	// The developer's team identifier that the system used when it signed the app.
	TeamIdentifier string `json:"TeamIdentifier" plist:"TeamIdentifier" required:"true"`
}

// The payload that configures associated domains.
// Configures Associated Domains to be used with features such as Extensible AppSSO, universal links and Password AutoFill. Settings are per-user. The effective settings for a user will be the union of payloads installed for the device and the user. Users on a system that are not managed by the MDM will not have any effective settings, not even those from device payloads.
type AssociatedDomains struct {
	*CommonPayloadKeys
	// A dictionary that maps apps to their associated domains.
	Configuration []*ConfigurationItem `json:"Configuration" plist:"Configuration" required:"true"`
}

func (p *AssociatedDomains) PayloadType() string {
	return "com.apple.associated-domains"
}

// A dictionary that maps apps to their associated domains.
type ConfigurationItem struct {
	// The app identifier to associate the domains with.
	ApplicationIdentifier string `json:"ApplicationIdentifier" plist:"ApplicationIdentifier" required:"true"`
	// The domains to associate with the app. Each string is in the form of `service:domain`. Use fully qualified hostnames, such as `www.example.com`. See `Supporting associated domains` for more information.
	AssociatedDomains []string `json:"AssociatedDomains" plist:"AssociatedDomains" required:"true"`
	// If `true`, the system enables direct download of data for this domain instead of through a CDN. Set the entitlement value for this domain to `service:domain?mode=managed`; otherwise, the system ignores this value. Available in macOS 11 and later.
	EnableDirectDownloads *bool `json:"EnableDirectDownloads,omitempty" plist:"EnableDirectDownloads,omitempty"`
}

// The payload that configures a Calendar account.
type CalDAV struct {
	*CommonPayloadKeys
	// The description of the account.
	CalDAVAccountDescription *string `json:"CalDAVAccountDescription,omitempty" plist:"CalDAVAccountDescription,omitempty"`
	// The server's address.
	CalDAVHostName string `json:"CalDAVHostName" plist:"CalDAVHostName" required:"true"`
	// The user name for logins. If this profile is part of a non-interactive install, the system requires this field.
	CalDAVUsername *string `json:"CalDAVUsername,omitempty" plist:"CalDAVUsername,omitempty"`
	// The user's password. Only use this in encrypted profiles.
	CalDAVPassword *string `json:"CalDAVPassword,omitempty" plist:"CalDAVPassword,omitempty"`
	// The base URL to the user's calendar.
	CalDAVPrincipalURL *string `json:"CalDAVPrincipalURL,omitempty" plist:"CalDAVPrincipalURL,omitempty"`
	// If `true`, the system enables SSL.
	CalDAVUseSSL *bool `json:"CalDAVUseSSL,omitempty" plist:"CalDAVUseSSL,omitempty"`
	// The server's port.
	CalDAVPort *int64 `json:"CalDAVPort,omitempty" plist:"CalDAVPort,omitempty"`
	// The VPNUUID of the per-app VPN the account uses for network communication. Available in iOS 14 and later.
	VPNUUID *string `json:"VPNUUID,omitempty" plist:"VPNUUID,omitempty"`
}

func (p *CalDAV) PayloadType() string {
	return "com.apple.caldav.account"
}

// The payload that configures a Contacts account.
type CardDAV struct {
	*CommonPayloadKeys
	// The description of the account.
	CardDAVAccountDescription *string `json:"CardDAVAccountDescription,omitempty" plist:"CardDAVAccountDescription,omitempty"`
	// The server's address.
	CardDAVHostName string `json:"CardDAVHostName" plist:"CardDAVHostName" required:"true"`
	// The user name for logins.
	CardDAVUsername *string `json:"CardDAVUsername,omitempty" plist:"CardDAVUsername,omitempty"`
	// The user's password. Only use this in encrypted profiles.
	CardDAVPassword *string `json:"CardDAVPassword,omitempty" plist:"CardDAVPassword,omitempty"`
	// The base URL to the user's address book.
	CardDAVPrincipalURL *string `json:"CardDAVPrincipalURL,omitempty" plist:"CardDAVPrincipalURL,omitempty"`
	// If `true`, the system enables SSL.
	CardDAVUseSSL *bool `json:"CardDAVUseSSL,omitempty" plist:"CardDAVUseSSL,omitempty"`
	// The server's port.
	CardDAVPort *int64 `json:"CardDAVPort,omitempty" plist:"CardDAVPort,omitempty"`
	// An array of communication service rules for this account.
	CommunicationServiceRules *CardDAVCommunicationServiceRules `json:"CommunicationServiceRules,omitempty" plist:"CommunicationServiceRules,omitempty"`
	// The VPNUUID of the per-app VPN the account uses for network communication. Available in iOS 14 and later.
	VPNUUID *string `json:"VPNUUID,omitempty" plist:"VPNUUID,omitempty"`
}

func (p *CardDAV) PayloadType() string {
	return "com.apple.carddav.account"
}

// An array of communication service rules for this account.
type CardDAVCommunicationServiceRules struct {
	// A dictionary of service handlers for contacts from this account.
	DefaultServiceHandlers *CardDAVCommunicationServiceRulesDefaultServiceHandlers `json:"DefaultServiceHandlers,omitempty" plist:"DefaultServiceHandlers,omitempty"`
}

// A dictionary of service handlers for contacts from this account.
type CardDAVCommunicationServiceRulesDefaultServiceHandlers struct {
	// The bundle identifier for the default application that handles audio calls to contacts from this account.
	AudioCall *string `json:"AudioCall,omitempty" plist:"AudioCall,omitempty"`
}

// The payload that configures cellular settings.
// This payload cannot be installed if an APN payload is already installed.
// This payload only applies to the preferred data SIM. There is no way to have a cellular payload affect a different SIM.
// This payload replaces the com.apple.managedCarrier payload. The latter payload is supported, but deprecated.
type Cellular struct {
	*CommonPayloadKeys
	// A configuration dictionary.
	AttachAPN *AttachAPN `json:"AttachAPN,omitempty" plist:"AttachAPN,omitempty"`
	// An array of access point name (APN) dictionaries.
	APNs *[]*APNsItem `json:"APNs,omitempty" plist:"APNs,omitempty"`
}

func (p *Cellular) PayloadType() string {
	return "com.apple.cellular"
}

// A configuration dictionary.
type AttachAPN struct {
	// The name for this configuration.
	Name string `json:"Name" plist:"Name" required:"true"`
	// The authentication type.
	AuthenticationType *AttachAPNAuthenticationType `default:"PAP" json:"AuthenticationType,omitempty" plist:"AuthenticationType,omitempty"`
	// The user name.
	Username *string `json:"Username,omitempty" plist:"Username,omitempty"`
	// The password for the user.
	Password *string `json:"Password,omitempty" plist:"Password,omitempty"`
	// The Internet Protocol versions that the system supports. Allowed values:
	// - `1`: IPv4
	// - `2`: IPv6
	// - `3`: Both
	AllowedProtocolMask *AttachAPNAllowedProtocolMask `json:"AllowedProtocolMask,omitempty" plist:"AllowedProtocolMask,omitempty"`
}

// The authentication type.
type AttachAPNAuthenticationType string

const (
	AttachAPNAuthenticationTypeCHAP AttachAPNAuthenticationType = "CHAP"
	AttachAPNAuthenticationTypePAP  AttachAPNAuthenticationType = "PAP"
)

// The Internet Protocol versions that the system supports. Allowed values:
// - `1`: IPv4
// - `2`: IPv6
// - `3`: Both
type AttachAPNAllowedProtocolMask int64

const (
	AttachAPNAllowedProtocolMask1 AttachAPNAllowedProtocolMask = 1
	AttachAPNAllowedProtocolMask2 AttachAPNAllowedProtocolMask = 2
	AttachAPNAllowedProtocolMask3 AttachAPNAllowedProtocolMask = 3
)

// A dictionary that contains details about an access point name (APN) configuration.
type APNsItem struct {
	// The name for this configuration.
	Name string `json:"Name" plist:"Name" required:"true"`
	// The authentication type for logging in.
	AuthenticationType *APNsItemAuthenticationType `default:"PAP" json:"AuthenticationType,omitempty" plist:"AuthenticationType,omitempty"`
	// The user name for the APN.
	Username *string `json:"Username,omitempty" plist:"Username,omitempty"`
	// The user's password for the APN.
	Password *string `json:"Password,omitempty" plist:"Password,omitempty"`
	// The proxy server's address.
	ProxyServer *string `json:"ProxyServer,omitempty" plist:"ProxyServer,omitempty"`
	// The proxy server's port number.
	ProxyPort *int64 `json:"ProxyPort,omitempty" plist:"ProxyPort,omitempty"`
	// The default Internet Protocol versions. Available in iOS 10.3 but no longer used in iOS 11 and later. Allowed values:
	// - `1`: IPv4
	// - `2`: IPv6
	// - `3`: Both
	DefaultProtocolMask *DefaultProtocolMask `json:"DefaultProtocolMask,omitempty" plist:"DefaultProtocolMask,omitempty"`
	// The Internet Protocol versions that the system supports. Available in iOS 10.3 and later. Allowed values:
	// - `1`: IPv4
	// - `2`: IPv6
	// - `3`: Both
	AllowedProtocolMask *APNsItemAllowedProtocolMask `json:"AllowedProtocolMask,omitempty" plist:"AllowedProtocolMask,omitempty"`
	// The Internet Protocol versions that the system supports while roaming. Available in iOS 10.3 and later. Allowed values:
	// - `1`: IPv4
	// - `2`: IPv6
	// - `3`: Both
	AllowedProtocolMaskInRoaming *AllowedProtocolMaskInRoaming `json:"AllowedProtocolMaskInRoaming,omitempty" plist:"AllowedProtocolMaskInRoaming,omitempty"`
	// The Internet Protocol versions that the system supports while roaming. Available in iOS 10.3 and later. Allowed values:
	// - `1`: IPv4
	// - `2`: IPv6
	// - `3`: Both
	AllowedProtocolMaskInDomesticRoaming *AllowedProtocolMaskInDomesticRoaming `json:"AllowedProtocolMaskInDomesticRoaming,omitempty" plist:"AllowedProtocolMaskInDomesticRoaming,omitempty"`
	// If `true`, the system enables XLAT464. Available in iOS 16 and later and watchOS 9 and later.
	EnableXLAT464 *bool `json:"EnableXLAT464,omitempty" plist:"EnableXLAT464,omitempty"`
}

// The authentication type for logging in.
type APNsItemAuthenticationType string

const (
	APNsItemAuthenticationTypeCHAP APNsItemAuthenticationType = "CHAP"
	APNsItemAuthenticationTypePAP  APNsItemAuthenticationType = "PAP"
)

// The default Internet Protocol versions. Available in iOS 10.3 but no longer used in iOS 11 and later. Allowed values:
// - `1`: IPv4
// - `2`: IPv6
// - `3`: Both
type DefaultProtocolMask int64

const (
	DefaultProtocolMask1 DefaultProtocolMask = 1
	DefaultProtocolMask2 DefaultProtocolMask = 2
	DefaultProtocolMask3 DefaultProtocolMask = 3
)

// The Internet Protocol versions that the system supports. Available in iOS 10.3 and later. Allowed values:
// - `1`: IPv4
// - `2`: IPv6
// - `3`: Both
type APNsItemAllowedProtocolMask int64

const (
	APNsItemAllowedProtocolMask1 APNsItemAllowedProtocolMask = 1
	APNsItemAllowedProtocolMask2 APNsItemAllowedProtocolMask = 2
	APNsItemAllowedProtocolMask3 APNsItemAllowedProtocolMask = 3
)

// The Internet Protocol versions that the system supports while roaming. Available in iOS 10.3 and later. Allowed values:
// - `1`: IPv4
// - `2`: IPv6
// - `3`: Both
type AllowedProtocolMaskInRoaming int64

const (
	AllowedProtocolMaskInRoaming1 AllowedProtocolMaskInRoaming = 1
	AllowedProtocolMaskInRoaming2 AllowedProtocolMaskInRoaming = 2
	AllowedProtocolMaskInRoaming3 AllowedProtocolMaskInRoaming = 3
)

// The Internet Protocol versions that the system supports while roaming. Available in iOS 10.3 and later. Allowed values:
// - `1`: IPv4
// - `2`: IPv6
// - `3`: Both
type AllowedProtocolMaskInDomesticRoaming int64

const (
	AllowedProtocolMaskInDomesticRoaming1 AllowedProtocolMaskInDomesticRoaming = 1
	AllowedProtocolMaskInDomesticRoaming2 AllowedProtocolMaskInDomesticRoaming = 2
	AllowedProtocolMaskInDomesticRoaming3 AllowedProtocolMaskInDomesticRoaming = 3
)

// The payload that provides device info on private network deployments, including geographical location, preference over Wi-Fi, and network deployment type.
// Payload can be used to provide device info on private network deployments including geographical location, preference over Wi-Fi, and network deployment type. Only five Cellular Private Networks can be configured simultaneously.
type CellularPrivateNetwork struct {
	*CommonPayloadKeys
	// A list of up to 1000 geofences for private networks. Geofencing is only used on iPhone.
	Geofences *[]*GeofenceItem `json:"Geofences,omitempty" plist:"Geofences,omitempty"`
	// The name of the private network configuration data set.
	DataSetName string `json:"DataSetName" plist:"DataSetName" required:"true"`
	// The version number of this dataset that the system uses to track updates.
	VersionNumber string `json:"VersionNumber" plist:"VersionNumber" required:"true"`
	// Set to `true` to prefer this private network over Wi-Fi.
	CellularDataPreferred *bool `json:"CellularDataPreferred,omitempty" plist:"CellularDataPreferred,omitempty"`
	// Set to `true` if this private network is NR Standalone.
	EnableNRStandalone *bool `json:"EnableNRStandalone,omitempty" plist:"EnableNRStandalone,omitempty"`
	// A string using the 3GPP "Coordinated NID" (option 1 or option 2) format (defined in 3GPP 31.102, Section 12.7.1). The device uses this value to match a SIM present on the device.
	// All combinations of `NetworkIdentifier` and `CsgNetworkIdentifier` must be unique across all profiles installed on the device.
	NetworkIdentifier *string `json:"NetworkIdentifier,omitempty" plist:"NetworkIdentifier,omitempty"`
	// A string using the 3GPP "CSG_ID" format (defined in 3GPP 23.003, Section 4.7). The device uses this value to match a SIM present on the device.
	// All combinations of `NetworkIdentifier` and `CsgNetworkIdentifier` must be unique across all profiles installed on the device.
	CsgNetworkIdentifier *string `json:"CsgNetworkIdentifier,omitempty" plist:"CsgNetworkIdentifier,omitempty"`
}

func (p *CellularPrivateNetwork) PayloadType() string {
	return "com.apple.cellularprivatenetwork.managed"
}

// A geofence for a private network.
type GeofenceItem struct {
	// The longitude of the geofence.
	Longitude float64 `json:"Longitude" plist:"Longitude" required:"true"`
	// The latitude of the geofence.
	Latitude float64 `json:"Latitude" plist:"Latitude" required:"true"`
	// Specifies the radius of the geofence in meters. Set this value slightly greater than the private cellular network coverage area.
	Radius float64 `json:"Radius" plist:"Radius" required:"true"`
	// A geofence identifier that's unique within a list of geofences.
	GeofenceId string `json:"GeofenceId" plist:"GeofenceId" required:"true"`
}

// The payload that configures Conference Room Display mode for Apple TV.
// Configures an Apple TV to enter Conference Room Display mode, and restrictions exit from that mode
type ConferenceRoomDisplay struct {
	*CommonPayloadKeys
	// The custom message displayed on the screen in Conference Room Display mode.
	Message *string `json:"Message,omitempty" plist:"Message,omitempty"`
}

func (p *ConferenceRoomDisplay) PayloadType() string {
	return "com.apple.conferenceroomdisplay"
}

// The payload that configures the names of the account user.
// This payload can be used on the device or user channel depending on what payload it is paired with.
// Device channel:
// *com.apple.MCX.FileVault2
// *com.apple.ADCertificate.managed
// *com.apple.DirectoryService.managed
// User channel:
// *com.apple.caldav.account
// *com.apple.carddav.account
// *com.apple.ews.account
// *com.apple.ldap.account
// *com.apple.mail.managed
type Identification struct {
	*CommonPayloadKeys
	// The dictionary that contains details about the user.
	PayloadIdentification PayloadIdentification `json:"PayloadIdentification" plist:"PayloadIdentification" required:"true"`
}

func (p *Identification) PayloadType() string {
	return "com.apple.configurationprofile.identification"
}

// The dictionary that contains details about the user.
type PayloadIdentification struct {
	// The UNIX user name for the accounts.
	UserName string `json:"UserName" plist:"UserName" required:"true"`
	// The full name of the account.
	FullName string `json:"FullName" plist:"FullName" required:"true"`
	// The address for the account.
	EmailAddress string `json:"EmailAddress" plist:"EmailAddress" required:"true"`
	// The authorization method. Either the profile contains the password or the user provides it.
	AuthMethod AuthMethod `json:"AuthMethod" plist:"AuthMethod" required:"true"`
	// The password for the account. Required when the `AuthMethod` is `Password`.
	Password string `json:"Password" plist:"Password" required:"true"`
	// The custom instructions for the user, if needed.
	Prompt *string `json:"Prompt,omitempty" plist:"Prompt,omitempty"`
	// The additional descriptive text for the user prompt.
	PromptMessage *string `json:"PromptMessage,omitempty" plist:"PromptMessage,omitempty"`
}

// The authorization method. Either the profile contains the password or the user provides it.
type AuthMethod string

const (
	AuthMethodPassword            AuthMethod = "Password"
	AuthMethodUserEnteredPassword AuthMethod = "UserEnteredPassword"
)

// The payload that configures allowed dashboard widgets.
// Widget restrictions.
type ParentalControlsDashboardWidgetRestrictions struct {
	*CommonPayloadKeys
	// If `true`, enables the widget allow list.
	WhiteListEnabled bool `json:"whiteListEnabled" plist:"whiteListEnabled" required:"true"`
	// An array of widget item dictionaries that are allowed.
	WhiteList []*ParentalControlsDashboardWidgetRestrictionsWhiteListWhiteListItem `json:"WhiteList" plist:"WhiteList" required:"true"`
}

func (p *ParentalControlsDashboardWidgetRestrictions) PayloadType() string {
	return "com.apple.dashboard"
}

// The widget item dictionary.
type ParentalControlsDashboardWidgetRestrictionsWhiteListWhiteListItem struct {
	// The type of allow list item. Set to `bundleID` to use a widget's bundle ID as its main ID.
	Type string `json:"Type" plist:"Type" required:"true"`
	// The bundle ID of a widget.
	ID string `json:"ID" plist:"ID" required:"true"`
}

// The payload that applies a set of declarations to the device through the Settings app.
type Declarations struct {
	*CommonPayloadKeys
	// The set of declarations to apply. The array items are Base64-encoded data representations of the declaration JSON data.
	Declarations [][]byte `json:"Declarations" plist:"Declarations" required:"true"`
}

func (p *Declarations) PayloadType() string {
	return "com.apple.declarations"
}

// The payload that configures the desktop wallpaper.
type Desktop struct {
	*CommonPayloadKeys
	// If `true`, locks the desktop picture. Replaced with allowWallpaperModification in macOS 10.13.
	Locked *bool `json:"locked,omitempty" plist:"locked,omitempty"`
	// The path to the desktop picture. If set, this picture is always locked.
	OverridePicturePath *string `json:"override-picture-path,omitempty" plist:"override-picture-path,omitempty"`
}

func (p *Desktop) PayloadType() string {
	return "com.apple.desktop"
}

// The payload that configures DNS proxies.
// As of iOS 15.0 this payload can be installed on unsupervised devices via MDM and can only be installed via MDM. As of iOS 16.0, this can be installed on user enrollments via MDM if DNSProxyUUID is specified.
type DNSProxy struct {
	*CommonPayloadKeys
	// The bundle identifier of the app containing the DNS proxy network extension.
	AppBundleIdentifier string `json:"AppBundleIdentifier" plist:"AppBundleIdentifier" required:"true"`
	// The bundle identifier of the DNS proxy network extension to use. Declaring the bundle identifier is useful for apps that contain more than one DNS proxy extension.
	ProviderBundleIdentifier *string `json:"ProviderBundleIdentifier,omitempty" plist:"ProviderBundleIdentifier,omitempty"`
	// The dictionary of vendor-specific configuration items.
	ProviderConfiguration *map[string]any `json:"ProviderConfiguration,omitempty" plist:"ProviderConfiguration,omitempty"`
	// A globally unique identifier for this DNS proxy configuration. The proxy processes DNS lookups traffic for managed apps with the same `DNSProxyUUID` in their app attributes. This key is required for user enrollment.
	DNSProxyUUID *string `json:"DNSProxyUUID,omitempty" plist:"DNSProxyUUID,omitempty"`
}

func (p *DNSProxy) PayloadType() string {
	return "com.apple.dnsProxy.managed"
}

// The payload that configures encrypted DNS settings.
type DNSSettings1 struct {
	*CommonPayloadKeys
	// A dictionary that defines a configuration for an encrypted DNS server.
	DNSSettings DNSSettingsDNSSettings `json:"DNSSettings" plist:"DNSSettings" required:"true"`
	// An array of rules that define the DNS settings. If not set, the system always applies the DNS settings. These rules are identical to the `OnDemandRules` array in VPN payloads.
	OnDemandRules *[]*DNSSettingsOnDemandRulesOnDemandRulesElement `json:"OnDemandRules,omitempty" plist:"OnDemandRules,omitempty"`
	// If `true`, the system prohibits users from disabling DNS settings. This key is only available on supervised devices.
	ProhibitDisablement *bool `json:"ProhibitDisablement,omitempty" plist:"ProhibitDisablement,omitempty"`
}

func (p *DNSSettings1) PayloadType() string {
	return "com.apple.dnsSettings.managed"
}

// A dictionary that defines a configuration for an encrypted DNS server.
type DNSSettingsDNSSettings struct {
	// The encrypted transport protocol used to communicate with the DNS server.
	DNSProtocol DNSSettingsDNSProtocol `json:"DNSProtocol" plist:"DNSProtocol" required:"true"`
	// The URI template of a DNS-over-HTTPS server, as defined in RFC 8484. This URL needs to use the `https://` scheme, and the system uses the hostname or address in the URL to validate the server certificate. If no `ServerAddresses` are provided, the system uses the hostname or address in the URL to determine the server addresses. Required if `DNSProtocol` is `HTTPS`.
	ServerURL *string `json:"ServerURL,omitempty" plist:"ServerURL,omitempty"`
	// The hostname of a DNS-over-TLS server used to validate the server certificate, as defined in RFC 7858. If no `ServerAddresses` are provided, the system uses the hostname to determine the server addresses. This key must be present only if the DNSProtocol is `TLS`.
	ServerName *string `json:"ServerName,omitempty" plist:"ServerName,omitempty"`
	// An unordered list of DNS server IP address strings. These IP addresses can be a mixture of IPv4 and IPv6 addresses.
	ServerAddresses *[]string `json:"ServerAddresses,omitempty" plist:"ServerAddresses,omitempty"`
	// If `true`, the device allows failover to the default system DNS resolver.
	AllowFailover *bool `json:"AllowFailover,omitempty" plist:"AllowFailover,omitempty"`
	// The UUID that points to an identity certificate payload. The system uses this identity to authenticate the user to the DNS resolver.
	PayloadCertificateUUID *string `json:"PayloadCertificateUUID,omitempty" plist:"PayloadCertificateUUID,omitempty"`
	// A list of domain strings used to determine which DNS queries use the DNS server. If not set, all domains use the DNS server.
	// The system supports a single wildcard (`*`) prefix, but it's not required. For example, both `*.example.com` and `example.com` match against `mydomain.example.com` and `your.domain.example.com`, but don't match against `mydomain-example.com`.
	SupplementalMatchDomains *[]string `json:"SupplementalMatchDomains,omitempty" plist:"SupplementalMatchDomains,omitempty"`
}

// The encrypted transport protocol used to communicate with the DNS server.
type DNSSettingsDNSProtocol string

const (
	DNSSettingsDNSProtocolHTTPS DNSSettingsDNSProtocol = "HTTPS"
	DNSSettingsDNSProtocolTLS   DNSSettingsDNSProtocol = "TLS"
)

type DNSSettingsOnDemandRulesOnDemandRulesElement struct {
	// The action to take if this dictionary matches the current network. Allowed values:
	// - `Connect`: Apply DNS Settings when the dictionary matches.
	// - `Disconnect`: Don't apply DNS Settings when the dictionary matches.
	// - `EvaluateConnection`: Apply DNS Settings with per-domain exceptions when the dictionary matches.
	Action DNSSettingsOnDemandRulesOnDemandRulesElementAction `json:"Action" plist:"Action" required:"true"`
	// An array of dictionaries that provide per-connection rules. The system uses this array only for settings where the `Action` value is `EvaluateConnection`.
	ActionParameters *[]*DNSSettingsOnDemandRulesOnDemandRulesElementActionParametersActionParameter `json:"ActionParameters,omitempty" plist:"ActionParameters,omitempty"`
	// An array of domain names. This rule matches if any of the domain names in the specified list matches any domain in the device's search domains list.
	// The system supports a single wildcard (`*`) prefix, but it's not required. For example, both `*.example.com` and `example.com` match against `mydomain.example.com` and `your.domain.example.com`, but don't match against `mydomain-example.com`.
	DNSDomainMatch *[]string `json:"DNSDomainMatch,omitempty" plist:"DNSDomainMatch,omitempty"`
	// An array of IP addresses. This rule matches if any of the network's specified DNS servers match any entry in the array.
	// The system supports matching with a single wildcard. For example, `17.*` matches any DNS server in the 17.0.0.0/8 subnet.
	DNSServerAddressMatch *[]string `json:"DNSServerAddressMatch,omitempty" plist:"DNSServerAddressMatch,omitempty"`
	// An interface type. If specified, this rule matches only if the primary network interface hardware matches the specified type.
	InterfaceTypeMatch *DNSSettingsOnDemandRulesOnDemandRulesElementInterfaceTypeMatch `json:"InterfaceTypeMatch,omitempty" plist:"InterfaceTypeMatch,omitempty"`
	// An array of SSIDs to match against the current network. If the network isn't a Wi-Fi network or if the SSID doesn't appear in this array, the match fails. Omit this key and the corresponding array to match against any SSID.
	SSIDMatch *[]string `json:"SSIDMatch,omitempty" plist:"SSIDMatch,omitempty"`
	// A URL to probe. This rule matches if this URL is successfully fetched and returns a 200 HTTP status code without redirection.
	URLStringProbe *string `json:"URLStringProbe,omitempty" plist:"URLStringProbe,omitempty"`
}

// The action to take if this dictionary matches the current network. Allowed values:
// - `Connect`: Apply DNS Settings when the dictionary matches.
// - `Disconnect`: Don't apply DNS Settings when the dictionary matches.
// - `EvaluateConnection`: Apply DNS Settings with per-domain exceptions when the dictionary matches.
type DNSSettingsOnDemandRulesOnDemandRulesElementAction string

const (
	DNSSettingsOnDemandRulesOnDemandRulesElementActionConnect            DNSSettingsOnDemandRulesOnDemandRulesElementAction = "Connect"
	DNSSettingsOnDemandRulesOnDemandRulesElementActionDisconnect         DNSSettingsOnDemandRulesOnDemandRulesElementAction = "Disconnect"
	DNSSettingsOnDemandRulesOnDemandRulesElementActionEvaluateConnection DNSSettingsOnDemandRulesOnDemandRulesElementAction = "EvaluateConnection"
)

// A dictionary that provides per-connection rules.
// The keys allowed in each dictionary are described below. Note: This array is only for dictionaries in which `EvaluateConnection` is the `Action` value.
type DNSSettingsOnDemandRulesOnDemandRulesElementActionParametersActionParameter struct {
	// The domains for which this evaluation applies.
	Domains []string `json:"Domains" plist:"Domains" required:"true"`
	// The DNS settings behavior for the specified domains. Allowed values:
	// * 'NeverConnect': Don't use the DNS Settings for the specified domains.
	// * 'ConnectIfNeeded': Allow using the DNS Settings for the specified domains.
	DomainAction DNSSettingsOnDemandRulesOnDemandRulesElementActionParametersActionParameterDomainAction `json:"DomainAction" plist:"DomainAction" required:"true"`
}

// The DNS settings behavior for the specified domains. Allowed values:
// * 'NeverConnect': Don't use the DNS Settings for the specified domains.
// * 'ConnectIfNeeded': Allow using the DNS Settings for the specified domains.
type DNSSettingsOnDemandRulesOnDemandRulesElementActionParametersActionParameterDomainAction string

const (
	DNSSettingsOnDemandRulesOnDemandRulesElementActionParametersActionParameterDomainActionNeverConnect    DNSSettingsOnDemandRulesOnDemandRulesElementActionParametersActionParameterDomainAction = "NeverConnect"
	DNSSettingsOnDemandRulesOnDemandRulesElementActionParametersActionParameterDomainActionConnectIfNeeded DNSSettingsOnDemandRulesOnDemandRulesElementActionParametersActionParameterDomainAction = "ConnectIfNeeded"
)

// An interface type. If specified, this rule matches only if the primary network interface hardware matches the specified type.
type DNSSettingsOnDemandRulesOnDemandRulesElementInterfaceTypeMatch string

const (
	DNSSettingsOnDemandRulesOnDemandRulesElementInterfaceTypeMatchEthernet DNSSettingsOnDemandRulesOnDemandRulesElementInterfaceTypeMatch = "Ethernet"
	DNSSettingsOnDemandRulesOnDemandRulesElementInterfaceTypeMatchWiFi     DNSSettingsOnDemandRulesOnDemandRulesElementInterfaceTypeMatch = "WiFi"
	DNSSettingsOnDemandRulesOnDemandRulesElementInterfaceTypeMatchCellular DNSSettingsOnDemandRulesOnDemandRulesElementInterfaceTypeMatch = "Cellular"
)

// The payload that configures the Dock.
type Dock struct {
	*CommonPayloadKeys
	// The tile size. Values must be in the range from 16 to 128.
	Tilesize *int64 `json:"tilesize,omitempty" plist:"tilesize,omitempty"`
	// If `true`, locks the size slider.
	SizeImmutable *bool `json:"size-immutable,omitempty" plist:"size-immutable,omitempty"`
	// If `true`, enables magnification.
	Magnification *bool `json:"magnification,omitempty" plist:"magnification,omitempty"`
	// If `true`, locks magnification.
	MagnifyImmutable *bool `json:"magnify-immutable,omitempty" plist:"magnify-immutable,omitempty"`
	// The size of the largest magnification.
	Largesize *int64 `json:"largesize,omitempty" plist:"largesize,omitempty"`
	// If `true`, locks the magnification slider.
	MagsizeImmutable *bool `json:"magsize-immutable,omitempty" plist:"magsize-immutable,omitempty"`
	// The orientation of the Dock.
	Orientation *Orientation `json:"orientation,omitempty" plist:"orientation,omitempty"`
	// If `true`, locks the position.
	PositionImmutable *bool `json:"position-immutable,omitempty" plist:"position-immutable,omitempty"`
	// The minimize effect.
	Mineffect *Mineffect `json:"mineffect,omitempty" plist:"mineffect,omitempty"`
	// If `true`, locks "Minimize windows using."
	MineffectImmutable *bool `json:"mineffect-immutable,omitempty" plist:"mineffect-immutable,omitempty"`
	// Set the "Prefer tabs when opening documents" to the provided value.
	Windowtabbing *Windowtabbing `json:"windowtabbing,omitempty" plist:"windowtabbing,omitempty"`
	// If `true`, disables "Prefer tabs when opening documents" checkbox.
	WindowtabbingImmutable *bool `json:"windowtabbing-immutable,omitempty" plist:"windowtabbing-immutable,omitempty"`
	// The behavior when the window's title bar is double-clicked.
	Dblclickbehavior *Dblclickbehavior `json:"dblclickbehavior,omitempty" plist:"dblclickbehavior,omitempty"`
	// If `true`, locks "Double-click a window's title bar."
	DblclickbehaviorImmutable *bool `json:"dblclickbehavior-immutable,omitempty" plist:"dblclickbehavior-immutable,omitempty"`
	// If `true`, enables "Minimize windows into application icon."
	MinimizeToApplication *bool `json:"minimize-to-application,omitempty" plist:"minimize-to-application,omitempty"`
	// If `true`, disables the "Minimize windows into application icon" checkbox.
	MinintoappImmutable *bool `json:"minintoapp-immutable,omitempty" plist:"minintoapp-immutable,omitempty"`
	// If `true`, enables "Animate opening applications."
	Launchanim *bool `json:"launchanim,omitempty" plist:"launchanim,omitempty"`
	// If `true`, locks "Animate opening applications."
	LaunchanimImmutable *bool `json:"launchanim-immutable,omitempty" plist:"launchanim-immutable,omitempty"`
	// If `true`, enables "Automatically hide and show the Dock."
	Autohide *bool `json:"autohide,omitempty" plist:"autohide,omitempty"`
	// If `true`, locks "Automatically hide."
	AutohideImmutable *bool `json:"autohide-immutable,omitempty" plist:"autohide-immutable,omitempty"`
	// If true, shows the process indicator.
	ShowProcessIndicators *bool `json:"show-process-indicators,omitempty" plist:"show-process-indicators,omitempty"`
	// If `true`, locks "Show indicators."
	ShowindicatorsImmutable *bool `json:"showindicators-immutable,omitempty" plist:"showindicators-immutable,omitempty"`
	// If `true`, enables "Show recent items."
	ShowRecents *bool `json:"show-recents,omitempty" plist:"show-recents,omitempty"`
	// If `true`, disables "Show recent applications" checkbox.
	ShowrecentsImmutable *bool `json:"showrecents-immutable,omitempty" plist:"showrecents-immutable,omitempty"`
	// If `true`, disables changes to the Dock.
	ContentsImmutable *bool `json:"contents-immutable,omitempty" plist:"contents-immutable,omitempty"`
	// One or more special folders that may be created at user login time and placed in the Dock.
	// The "My Applications" item is only used for Simple Finder environments. The "Original Network Home" item is only used for mobile account users.
	MCXDockSpecialFolders *[]MCXDockSpecialFolders `json:"MCXDockSpecialFolders,omitempty" plist:"MCXDockSpecialFolders,omitempty"`
	// If `true`, use the file in `/Library/Preferences/com.apple.dockfixup.plist` when a new user or migrated user logs in. This option has no effect for existing users. Available in macOS 10.12 and later. Only available on the device channel.
	AllowDockFixupOverride *bool `json:"AllowDockFixupOverride,omitempty" plist:"AllowDockFixupOverride,omitempty"`
	// If `true`, uses the `static-apps` and `static-others` dictionaries for the Dock and ignores any items in the `persistent-apps` and `persistent-others` dictionaries. If `false`, the contents are merged with the static items listed first.
	StaticOnly *bool `json:"static-only,omitempty" plist:"static-only,omitempty"`
	// An array of items located on the Documents side of the Dock and cannot be removed from that location.
	StaticOthers *[]*StaticItem `json:"static-others,omitempty" plist:"static-others,omitempty"`
	// An array of items located on the Applications side of the Dock and cannot be removed from that location.
	StaticApps *[]*StaticItem `json:"static-apps,omitempty" plist:"static-apps,omitempty"`
	// An array of items located on the Applications side of the Dock that can be removed from the Dock.
	PersistentApps *[]*StaticItem `json:"persistent-apps,omitempty" plist:"persistent-apps,omitempty"`
	// An array of items located on the Documents side of the Dock that can be removed from the Dock.
	PersistentOthers *[]*StaticItem `json:"persistent-others,omitempty" plist:"persistent-others,omitempty"`
}

func (p *Dock) PayloadType() string {
	return "com.apple.dock"
}

// The orientation of the Dock.
type Orientation string

const (
	OrientationBottom Orientation = "bottom"
	OrientationLeft   Orientation = "left"
	OrientationRight  Orientation = "right"
)

// The minimize effect.
type Mineffect string

const (
	MineffectGenie Mineffect = "genie"
	MineffectScale Mineffect = "scale"
)

// Set the "Prefer tabs when opening documents" to the provided value.
type Windowtabbing string

const (
	WindowtabbingManual     Windowtabbing = "manual"
	WindowtabbingAlways     Windowtabbing = "always"
	WindowtabbingFullscreen Windowtabbing = "fullscreen"
)

// The behavior when the window's title bar is double-clicked.
type Dblclickbehavior string

const (
	DblclickbehaviorMinimize Dblclickbehavior = "minimize"
	DblclickbehaviorMaximize Dblclickbehavior = "maximize"
	DblclickbehaviorNone     Dblclickbehavior = "none"
)

// One or more special folders that may be created at user login time and placed in the Dock.
// The "My Applications" item is only used for Simple Finder environments. The "Original Network Home" item is only used for mobile account users.
type MCXDockSpecialFolders string

const (
	MCXDockSpecialFoldersAddDockMCXMyApplicationsFolder      MCXDockSpecialFolders = "AddDockMCXMyApplicationsFolder"
	MCXDockSpecialFoldersAddDockMCXDocumentsFolder           MCXDockSpecialFolders = "AddDockMCXDocumentsFolder"
	MCXDockSpecialFoldersAddDockMCXSharedFolder              MCXDockSpecialFolders = "AddDockMCXSharedFolder"
	MCXDockSpecialFoldersAddDockMCXOriginalNetworkHomeFolder MCXDockSpecialFolders = "AddDockMCXOriginalNetworkHomeFolder"
)

// Items that are located on the Documents side of the Dock and cannot be removed from that location.
type StaticItem struct {
	// The information about the Dock item.
	TileData TileData `json:"tile-data" plist:"tile-data" required:"true"`
	// The type of tile.
	TileType TileType `json:"tile-type" plist:"tile-type" required:"true"`
}

// The information about the Dock item.
type TileData struct {
	// The label of the Dock item.
	Label string `json:"label" plist:"label" required:"true"`
	// The URL string.
	Url *string `json:"url,omitempty" plist:"url,omitempty"`
	// The type of tile:
	// - `0`: URL
	// - `1`: File
	// - `3`: Directory
	FileType FileType `json:"file-type" plist:"file-type" required:"true"`
	// The data in a file. For Apple use only.
	FileData *map[string]any `json:"file-data,omitempty" plist:"file-data,omitempty"`
}

// The type of tile:
// - `0`: URL
// - `1`: File
// - `3`: Directory
type FileType int64

const (
	FileType0 FileType = 0
	FileType1 FileType = 1
	FileType3 FileType = 3
)

// The type of tile.
type TileType string

const (
	TileTypeFileTile      TileType = "file-tile"
	TileTypeDirectoryTile TileType = "directory-tile"
	TileTypeUrlTile       TileType = "url-tile"
)

// The payload that configures the domains under an organization's management.
// This payload defines web domains that are under an enterprise's management.
type Domains struct {
	*CommonPayloadKeys
	// An array of domains. Mail marks in red all email addresses that lack a suffix matching any of these strings.
	// Available in iOS 8 and later and macOS 10.10 and later.
	EmailDomains *[]string `json:"EmailDomains,omitempty" plist:"EmailDomains,omitempty"`
	// An array of domains. The system considers URLs matching the patterns listed in this property managed.
	// Available in iOS 9.3 and later.
	WebDomains *[]string `json:"WebDomains,omitempty" plist:"WebDomains,omitempty"`
	// An array of domains. Users can only save passwords in Safari from URLs matching the patterns listed here. This property doesn't disable the autofill feature itself.
	// Supervised devices or Shared iPads need this property to enable saving passwords in Safari.
	// Available in iOS 9.3 and later.
	SafariPasswordAutoFillDomains *[]string `json:"SafariPasswordAutoFillDomains,omitempty" plist:"SafariPasswordAutoFillDomains,omitempty"`
	// An array of up to 10 strings. URLs matching the patterns listed here have relaxed enforcement of cross-site tracking prevention.
	// Available in iOS 16.2 and later and macOS 13.1 and later.
	CrossSiteTrackingPreventionRelaxedDomains *[]string `json:"CrossSiteTrackingPreventionRelaxedDomains,omitempty" plist:"CrossSiteTrackingPreventionRelaxedDomains,omitempty"`
	// An array of up to 10 strings representing app bundle-ids. Apps matching the bundle-ids listed here have relaxed enforcement of cross-site tracking prevention for the domains listed in `CrossSiteTrackingPreventionRelaxedDomains`.
	// Available in iOS 18 and later and macOS 15 and later.
	CrossSiteTrackingPreventionRelaxedApps *[]string `json:"CrossSiteTrackingPreventionRelaxedApps,omitempty" plist:"CrossSiteTrackingPreventionRelaxedApps,omitempty"`
}

func (p *Domains) PayloadType() string {
	return "com.apple.domains"
}

// The payload that configures Exchange ActiveSync accounts.
// This payload configures an Exchange Active Sync account on an iOS device for Mail, Contacts, Calendars, Reminders, and Notes.
// Updating this payload overrides any settings that the user customized, such as EnableMail/Contacts/Calendars/Reminders/Notes and MailNumberOfPastDaysToSync.
type ExchangeActiveSync struct {
	*CommonPayloadKeys
	// The full email address for the account. If not present in the payload, the device prompts for this string during profile installation.
	EmailAddress *string `json:"EmailAddress,omitempty" plist:"EmailAddress,omitempty"`
	// The Exchange server host name or IP address.
	Host *string `json:"Host,omitempty" plist:"Host,omitempty"`
	// If `true`, the system enables SSL for authentication.
	SSL *bool `json:"SSL,omitempty" plist:"SSL,omitempty"`
	// If `true`, enables OAuth for authentication. If enabled, don't specify a password.
	// Available only in iOS 12.0 and above.
	OAuth *bool `json:"OAuth,omitempty" plist:"OAuth,omitempty"`
	// This user name for this Exchange account. Required for noninteractive installations like MDM in iOS.
	UserName *string `json:"UserName,omitempty" plist:"UserName,omitempty"`
	// The password of the account. Use only with encrypted profiles.
	Password *string `json:"Password,omitempty" plist:"Password,omitempty"`
	// The `.p12` identity certificate in NSData blob format, for accounts that allow authentication via certificate.
	Certificate *[]byte `json:"Certificate,omitempty" plist:"Certificate,omitempty"`
	// The name or description of the certificate.
	CertificateName *string `json:"CertificateName,omitempty" plist:"CertificateName,omitempty"`
	// The password necessary for the `.p12` identity certificate. Used with mandatory encryption of profiles.
	CertificatePassword *string `json:"CertificatePassword,omitempty" plist:"CertificatePassword,omitempty"`
	// If `true`, the system prevents moving messages from out of this email account into another account. This setting also prevents forwarding or replying from an account other than the recipient of the message.
	PreventMove *bool `json:"PreventMove,omitempty" plist:"PreventMove,omitempty"`
	// If `true`, prevents this account from sending mail in any app other than the Apple Mail app.
	PreventAppSheet *bool `json:"PreventAppSheet,omitempty" plist:"PreventAppSheet,omitempty"`
	// The UUID of the certificate payload within the same profile to use for the identity credential. If this field is present, the Certificate field isn't used.
	PayloadCertificateUUID *string `json:"PayloadCertificateUUID,omitempty" plist:"PayloadCertificateUUID,omitempty"`
	// If `true`, the system enables S/MIME encryption. In iOS 10.0 and later, this key is ignored. Use `SMIMESigningEnabled` instead.
	SMIMEEnabled *bool `json:"SMIMEEnabled,omitempty" plist:"SMIMEEnabled,omitempty"`
	// If `true`, the system enables S/MIME signing for this account. Available in iOS 10.0 and later.
	SMIMESigningEnabled *bool `json:"SMIMESigningEnabled,omitempty" plist:"SMIMESigningEnabled,omitempty"`
	// The UUID of the identity certificate used to sign messages sent from this account.
	SMIMESigningCertificateUUID *string `json:"SMIMESigningCertificateUUID,omitempty" plist:"SMIMESigningCertificateUUID,omitempty"`
	// If `true`, the system enables S/MIME encryption for this account. Available in iOS 10.0 and later. As of iOS 12.0, this key is deprecated. Use `SMIMEEncryptByDefault` instead.
	SMIMEEncryptionEnabled *bool `json:"SMIMEEncryptionEnabled,omitempty" plist:"SMIMEEncryptionEnabled,omitempty"`
	// The payload UUID of the identity certificate used to decrypt messages sent to this account. The system attaches the public certificate to outgoing mail to allow the user to receive encrypted mail. When the user sends encrypted mail, the system uses the public certificate to encrypt the copy of the mail in the user's Sent mailbox.
	SMIMEEncryptionCertificateUUID *string `json:"SMIMEEncryptionCertificateUUID,omitempty" plist:"SMIMEEncryptionCertificateUUID,omitempty"`
	// If `true`, the system displays the per-message encryption switch in the Mail Compose UI.
	// Available in iOS 8.0 and later. As of iOS 12.0, this key is deprecated. Use `SMIMEEnableEncryptionPerMessageSwitch` instead.
	SMIMEEnablePerMessageSwitch *bool `json:"SMIMEEnablePerMessageSwitch,omitempty" plist:"SMIMEEnablePerMessageSwitch,omitempty"`
	// If `true`, the system excludes this account from Recent Addresses syncing.
	DisableMailRecentsSyncing *bool `json:"disableMailRecentsSyncing,omitempty" plist:"disableMailRecentsSyncing,omitempty"`
	// The number of days in the past to sync mail on the device.
	// For no limit, use the value `0`.
	MailNumberOfPastDaysToSync *MailNumberOfPastDaysToSync `default:"7" json:"MailNumberOfPastDaysToSync,omitempty" plist:"MailNumberOfPastDaysToSync,omitempty"`
	// The value of the `X-Apple-Config-Magic` header in each EAS HTTP request.
	HeaderMagic *string `json:"HeaderMagic,omitempty" plist:"HeaderMagic,omitempty"`
	// The communication service handler rules for this account.
	CommunicationServiceRules *ExchangeActiveSyncCommunicationServiceRules `json:"CommunicationServiceRules,omitempty" plist:"CommunicationServiceRules,omitempty"`
	// If `true`, the system enables this account to use Mail Drop.
	AllowMailDrop *bool `json:"allowMailDrop,omitempty" plist:"allowMailDrop,omitempty"`
	// If `true`, the user can turn S/MIME signing on or off in Settings. Available in iOS 12.0 and later.
	SMIMESigningUserOverrideable *bool `json:"SMIMESigningUserOverrideable,omitempty" plist:"SMIMESigningUserOverrideable,omitempty"`
	// If `true`, the user can select the signing identity. Available in iOS 12.0 and later.
	SMIMESigningCertificateUUIDUserOverrideable *bool `json:"SMIMESigningCertificateUUIDUserOverrideable,omitempty" plist:"SMIMESigningCertificateUUIDUserOverrideable,omitempty"`
	// If `true`, the system enables S/MIME encryption by default. If `SMIMEEnableEncryptionPerMessageSwitch` is `false`, the user can't change this default. Available in iOS 12.0 and later.
	SMIMEEncryptByDefault *bool `json:"SMIMEEncryptByDefault,omitempty" plist:"SMIMEEncryptByDefault,omitempty"`
	// If `true`, the system enables encryption by default and the user can't change it. Available in iOS 12.0 and later.
	SMIMEEncryptByDefaultUserOverrideable *bool `json:"SMIMEEncryptByDefaultUserOverrideable,omitempty" plist:"SMIMEEncryptByDefaultUserOverrideable,omitempty"`
	// If `true`, the user can select the S/MIME encryption identity, and encryption is on.Available in iOS 12.0 and later.
	SMIMEEncryptionCertificateUUIDUserOverrideable *bool `json:"SMIMEEncryptionCertificateUUIDUserOverrideable,omitempty" plist:"SMIMEEncryptionCertificateUUIDUserOverrideable,omitempty"`
	// If `true`, the system displays the per-message encryption switch in the Mail Compose UI. Available in iOS 12.0 and later.
	SMIMEEnableEncryptionPerMessageSwitch *bool `json:"SMIMEEnableEncryptionPerMessageSwitch,omitempty" plist:"SMIMEEnableEncryptionPerMessageSwitch,omitempty"`
	// If `false`, the system disables the Mail service for this account. The user can reenable Mail service in Settings unless `EnableMailUserOverridable` is `false`.
	// > Note:
	// > At least of the following fields needs to be `true`: `EnableMail`, `EnableContacts`, `EnableCalendars`, `EnableReminders`, and `EnableNotes`.
	EnableMail *bool `json:"EnableMail,omitempty" plist:"EnableMail,omitempty"`
	// If `false`, the system disables the Contacts service for this account. The user can reenable Contacts service in Settings unless `EnableContactsUserOverridable` is `false`.
	// > Note:
	// > At least of the following fields needs to be `true`: `EnableMail`, `EnableContacts`, `EnableCalendars`, `EnableReminders`, and `EnableNotes`.
	EnableContacts *bool `json:"EnableContacts,omitempty" plist:"EnableContacts,omitempty"`
	// If `false`, the system disables the Calendars service for this account. The user can reenable Calendars service in Settings unless `EnableCalendarsUserOverridable` is `false`.
	// > Note:
	// > At least of the following fields needs to be `true`: `EnableMail`, `EnableContacts`, `EnableCalendars`, `EnableReminders`, and `EnableNotes`.
	EnableCalendars *bool `json:"EnableCalendars,omitempty" plist:"EnableCalendars,omitempty"`
	// If `false`, the system disables the Reminders service for this account. The user can reenable Reminders service in Settings unless `EnableRemindersUserOverridable` is `false`.
	// > Note:
	// > At least of the following fields needs to be `true`: `EnableMail`, `EnableContacts`, `EnableCalendars`, `EnableReminders`, and `EnableNotes`.
	EnableReminders *bool `json:"EnableReminders,omitempty" plist:"EnableReminders,omitempty"`
	// If `false`, the system disables the Notes service for this account. The user can reenable Notes service in Settings unless `EnableNotesUserOverridable` is `false`.
	// > Note:
	// > At least of the following fields needs to be `true`: `EnableMail`, `EnableContacts`, `EnableCalendars`, `EnableReminders`, and `EnableNotes`.
	EnableNotes *bool `json:"EnableNotes,omitempty" plist:"EnableNotes,omitempty"`
	// If `false`, the system prevents the user from changing the state of the Mail service for this account in Settings.
	EnableMailUserOverridable *bool `json:"EnableMailUserOverridable,omitempty" plist:"EnableMailUserOverridable,omitempty"`
	// If `false`, the system prevents the user from changing the state of the Contacts service for this account in Settings.
	EnableContactsUserOverridable *bool `json:"EnableContactsUserOverridable,omitempty" plist:"EnableContactsUserOverridable,omitempty"`
	// If `false`, the system prevents the user from changing the state of the Calendars service for this account in Settings.
	EnableCalendarsUserOverridable *bool `json:"EnableCalendarsUserOverridable,omitempty" plist:"EnableCalendarsUserOverridable,omitempty"`
	// If `false`, the system prevents the user from changing the state of the Reminders service for this account in Settings.
	EnableRemindersUserOverridable *bool `json:"EnableRemindersUserOverridable,omitempty" plist:"EnableRemindersUserOverridable,omitempty"`
	// If `false`, prevents the user from changing the state of the Notes service for this account in Settings.
	EnableNotesUserOverridable *bool `json:"EnableNotesUserOverridable,omitempty" plist:"EnableNotesUserOverridable,omitempty"`
	// The URL that this account should use for signing in through OAuth. Ignored unless `OAuth` is `true`. If you specify this URL, auto-discovery isn't used for this account, so you need to also specify a host.
	OAuthSignInURL *string `json:"OAuthSignInURL,omitempty" plist:"OAuthSignInURL,omitempty"`
	// The URL that this account should use for token requests through OAuth. Ignored unless `OAuth` is `true`.
	OAuthTokenRequestURL *string `json:"OAuthTokenRequestURL,omitempty" plist:"OAuthTokenRequestURL,omitempty"`
	// If `true`, the system overrides the previous user/EAS password with the new EAS password in the payload. Available in iOS 14 and later.
	OverridePreviousPassword *bool `json:"OverridePreviousPassword,omitempty" plist:"OverridePreviousPassword,omitempty"`
	// The VPNUUID of the per-app VPN the account uses for network communication. Available in iOS 14 and later.
	VPNUUID *string `json:"VPNUUID,omitempty" plist:"VPNUUID,omitempty"`
}

func (p *ExchangeActiveSync) PayloadType() string {
	return "com.apple.eas.account"
}

// The number of days in the past to sync mail on the device.
// For no limit, use the value `0`.
type MailNumberOfPastDaysToSync int64

const (
	MailNumberOfPastDaysToSync0  MailNumberOfPastDaysToSync = 0
	MailNumberOfPastDaysToSync1  MailNumberOfPastDaysToSync = 1
	MailNumberOfPastDaysToSync3  MailNumberOfPastDaysToSync = 3
	MailNumberOfPastDaysToSync7  MailNumberOfPastDaysToSync = 7
	MailNumberOfPastDaysToSync14 MailNumberOfPastDaysToSync = 14
	MailNumberOfPastDaysToSync31 MailNumberOfPastDaysToSync = 31
)

// The communication service handler rules for this account.
type ExchangeActiveSyncCommunicationServiceRules struct {
	// The default handlers to use for contacts from this account.
	DefaultServiceHandlers *ExchangeActiveSyncCommunicationServiceRulesDefaultServiceHandlers `json:"DefaultServiceHandlers,omitempty" plist:"DefaultServiceHandlers,omitempty"`
}

// The default handlers to use for contacts from this account.
type ExchangeActiveSyncCommunicationServiceRulesDefaultServiceHandlers struct {
	// The bundle identifier of the default application to use for audio calls made to contacts from this account.
	AudioCall *string `json:"AudioCall,omitempty" plist:"AudioCall,omitempty"`
}

// The payload that configures the users, groups, and departments within an educational organization.
// This payload is used to configure Classroom students, Classroom instructors, and the Shared iPad login screen. These do not necessarily require the same set of keys to be present in their payloads, so make sure to include all keys that are required for the education product you are configuring.
type EducationConfiguration struct {
	*CommonPayloadKeys
	// The organization's UUID identifier. This identifier can be any valid UUID. All teacher and student devices that need to communicate with one another must have the same organization UUID, particularly if they originated from different Device Enrollment Programs.
	OrganizationUUID string `json:"OrganizationUUID" plist:"OrganizationUUID" required:"true"`
	// The organization's display name. The system displays this name in the iOS login screen.
	OrganizationName string `json:"OrganizationName" plist:"OrganizationName" required:"true"`
	// The UUID of an identity certificate payload within the same profile to use for performing client authentication with other devices. This property supports PKCS12 certificates.
	// Required to configure Classroom. Has no effect on the configuration of the Shared iPad login screen.
	PayloadCertificateUUID *string `json:"PayloadCertificateUUID,omitempty" plist:"PayloadCertificateUUID,omitempty"`
	// The array of UUIDs referring to certificate payloads within the same profile that the system uses to authorize leader peer certificate identities. This array needs to contain all necessary certificates to validate the entire chain of trust. Leader certificates needs to have the common name prefix leader, which is case insensitive.
	// This property doesn't support identity payloads or PKCS12 certificates.
	// Required when configuring a student device for Classroom, and ignored when configuring an instructor device. Has no effect on the configuration of the Shared iPad login screen.
	LeaderPayloadCertificateAnchorUUID *[]string `json:"LeaderPayloadCertificateAnchorUUID,omitempty" plist:"LeaderPayloadCertificateAnchorUUID,omitempty"`
	// The array of UUIDs referring to certificate payloads within the same profile that the system uses to authorize group member peer certificate identities. This array must contain all certificates needed to validate the entire chain of trust. Member certificates must have the common name prefix member (case insensitive).
	// This property doesn't support identity payloads or PKCS12 certificates.
	// Required when configuring a student device for Classroom, and ignored when configuring an instructor device. Has no effect on the configuration of the Shared iPad login screen.
	MemberPayloadCertificateAnchorUUID *[]string `json:"MemberPayloadCertificateAnchorUUID,omitempty" plist:"MemberPayloadCertificateAnchorUUID,omitempty"`
	// The UUID of an identity certificate payload within the same profile that the system uses to perform client authentication when fetching additional resources, such as student images.
	// If set, the system uses this key to configure both Classroom and the Shared iPad login screen. If not set, the system uses MDM client identity.
	ResourcePayloadCertificateUUID *string `json:"ResourcePayloadCertificateUUID,omitempty" plist:"ResourcePayloadCertificateUUID,omitempty"`
	// The unique string that identifies the user of this device within the organization.
	// Don't set this value in payloads intended to configure the Shared iPad login screen.
	UserIdentifier string `json:"UserIdentifier" plist:"UserIdentifier" required:"true"`
	// _For Shared iPad profiles:_ The array of dictionaries that defines which departments the system displays in the Shared iPad login screen. If set, the system uses this key to configure both Classroom and the Shared iPad login screen.
	Departments *[]*DepartmentsItem `json:"Departments,omitempty" plist:"Departments,omitempty"`
	// _For Shared iPad profiles:_ The array of dictionaries that defines which groups the user can select in the Login Window.
	// _For leader/teacher profiles:_ The array of dictionaries that defines the groups that the user can control.
	// _For member/student profiles:_ The array of dictionaries that defines the groups where the user is a member.
	Groups []*GroupsItem `json:"Groups" plist:"Groups" required:"true"`
	// For Shared iPad profiles: The array of dictionaries that define the users that the system displays in the iOS Login Window.
	// _For leader/teacher profiles:_ The array of dictionaries that define users that are members of the teacher's groups.
	// _For member/student profiles:_ The array of dictionaries that needs to contain the definition of the user specified in the `UserIdentifier` key. With one-to-one member devices, this key should include only the device user and the teacher but not other class members.
	Users []*UsersItem `json:"Users" plist:"Users" required:"true"`
	// _For leader/teacher profiles:_ The array of dictionaries that defines which device groups the leader can assign devices to. Not included in member payloads.
	DeviceGroups *[]*DeviceGroupsItem `json:"DeviceGroups,omitempty" plist:"DeviceGroups,omitempty"`
	// If `true`, the system allows students enrolled in managed classes to modify their teacher's permissions for screen observation on their device.
	ScreenObservationPermissionModificationAllowed *bool `json:"ScreenObservationPermissionModificationAllowed,omitempty" plist:"ScreenObservationPermissionModificationAllowed,omitempty"`
}

func (p *EducationConfiguration) PayloadType() string {
	return "com.apple.education"
}

// A department in the organization.
type DepartmentsItem struct {
	// The display name of the department.
	Name string `json:"Name" plist:"Name" required:"true"`
	// The group beacon identifiers that are members of this department.
	GroupBeaconIDs []int64 `json:"GroupBeaconIDs" plist:"GroupBeaconIDs" required:"true"`
}

// An array of dictionaries defining groups.
type GroupsItem struct {
	// An unsigned 16 bit integer specifying this group's unique beacon ID.
	BeaconID int64 `json:"BeaconID" plist:"BeaconID" required:"true"`
	// The display name of the group.
	Name string `json:"Name" plist:"Name" required:"true"`
	// The description of the group.
	Description *string `json:"Description,omitempty" plist:"Description,omitempty"`
	// Deprecated in iOS 9.3.1 and later. The URL of an image for the group.
	ImageURL *string `json:"ImageURL,omitempty" plist:"ImageURL,omitempty"`
	// The source that provided this group, such as SIS, or MDM.
	ConfigurationSource *string `json:"ConfigurationSource,omitempty" plist:"ConfigurationSource,omitempty"`
	// The user identifiers that are leaders of this group.
	LeaderIdentifiers *[]string `json:"LeaderIdentifiers,omitempty" plist:"LeaderIdentifiers,omitempty"`
	// The entries in the Users array that are members of the group.
	MemberIdentifiers []string `json:"MemberIdentifiers" plist:"MemberIdentifiers" required:"true"`
	// The identifiers that refer to entries in the `DeviceGroups` array to which the instructor can assign users from this class.
	// Has no effect on the configuration of the Shared iPad login screen.
	DeviceGroupIdentifiers *[]string `json:"DeviceGroupIdentifiers,omitempty" plist:"DeviceGroupIdentifiers,omitempty"`
}

// A user in the organization.
type UsersItem struct {
	// The unique identifier for a user in the organization.
	Identifier string `json:"Identifier" plist:"Identifier" required:"true"`
	// The name of the user.
	Name string `json:"Name" plist:"Name" required:"true"`
	// The given name of the user.
	GivenName *string `json:"GivenName,omitempty" plist:"GivenName,omitempty"`
	// The family name of the user.
	FamilyName *string `json:"FamilyName,omitempty" plist:"FamilyName,omitempty"`
	// The user's phonetic given name. The system uses this name to sort users in the Classroom app and the Shared iPad Login Screen.
	PhoneticGivenName *string `json:"PhoneticGivenName,omitempty" plist:"PhoneticGivenName,omitempty"`
	// The user's phonetic family name. The system uses this name to sort users in the Classroom app and the Shared iPad login screen.
	PhoneticFamilyName *string `json:"PhoneticFamilyName,omitempty" plist:"PhoneticFamilyName,omitempty"`
	// A string that contains a URL pointing to an image of the user. The system displays this image in the iOS login screen and in the Classroom app. The recommended resolution is 256 x 256 pixels (512 x 512 pixels on a 2x device). The recommended formats are JPEG, PNG, and TIFF. The system uses the `ResourcePayloadCertificateUUID` identity certificate or the MDM client identity to perform authentication when fetching the image.
	ImageURL *string `json:"ImageURL,omitempty" plist:"ImageURL,omitempty"`
	// Deprecated in iOS 9.3.1 and later. The URL pointing to an image of the user. The system uses the  `ResourcePayloadCertificateUUID` identity certificate or the MDM client identity to perform authentication when fetching the specified resource.
	FullScreenImageURL *string `json:"FullScreenImageURL,omitempty" plist:"FullScreenImageURL,omitempty"`
	// The Managed Apple Account for this user.
	// Not required to configure Classroom, but if set the system uses it.
	// Required to configure the Shared iPad login screen.
	AppleID *string `json:"AppleID,omitempty" plist:"AppleID,omitempty"`
	// The type of passcode UI to show when the user is at the Login Window.
	PasscodeType *PasscodeType `json:"PasscodeType,omitempty" plist:"PasscodeType,omitempty"`
}

// The type of passcode UI to show when the user is at the Login Window.
type PasscodeType string

const (
	PasscodeTypeComplex PasscodeType = "complex"
	PasscodeTypeFour    PasscodeType = "four"
	PasscodeTypeSix     PasscodeType = "six"
)

// A device group in the organization.
type DeviceGroupsItem struct {
	// The unique identifier for the device group in the organization.
	Identifier string `json:"Identifier" plist:"Identifier" required:"true"`
	// The name of the device group, which must be unique in the organization.
	Name string `json:"Name" plist:"Name" required:"true"`
	// The serial numbers of the devices in the group.
	SerialNumbers []string `json:"SerialNumbers" plist:"SerialNumbers" required:"true"`
}

// The payload that configures an Exchange Web Services accounts.
// For macOS 10.9 and higher, an Exchange Web services (EWS) account is configured with support for Mail, Contacts, Calendar, Notes and Reminders. macOS 10.7-10.8 only supported Contacts.
type ExchangeWebServices struct {
	*CommonPayloadKeys
	// The full email address for the account. If the email address string isn't present in the payload, the device prompts for it during profile installation.
	EmailAddress *string `json:"EmailAddress,omitempty" plist:"EmailAddress,omitempty"`
	// The Exchange server host name or IP address. Ignored if using OAuth.
	Host *string `json:"Host,omitempty" plist:"Host,omitempty"`
	// If `true`, the system enables SSL.
	SSL *bool `json:"SSL,omitempty" plist:"SSL,omitempty"`
	// If `true`, the system enables OAuth for authentication. Don't specify a password if `OAuth` is `true`. Available in macOS 10.14 and later
	OAuth *bool `json:"OAuth,omitempty" plist:"OAuth,omitempty"`
	// The URL to load into a web view for authentication through OAuth when autodiscovery isn't used. This setting requires a `Host` value.
	OAuthSignInURL *string `json:"OAuthSignInURL,omitempty" plist:"OAuthSignInURL,omitempty"`
	// The user name for this Exchange account. Required for noninteractive installation, such as through MDM. If missing, the system prompts the user for it during interactive profile installation.
	UserName *string `json:"UserName,omitempty" plist:"UserName,omitempty"`
	// The password of the account. Use only with encrypted profiles.
	Password *string `json:"Password,omitempty" plist:"Password,omitempty"`
	// The UUID of the certificate payload within the same profile to use for the identity credential. Supported on macOS 10.12 or later.
	PayloadCertificateUUID *string `json:"PayloadCertificateUUID,omitempty" plist:"PayloadCertificateUUID,omitempty"`
	// The UUID of the certificate payload within the same profile to use for the identity credential. Supported on macOS 10.11 or later. On macOS 10.12 or later use the PayloadCertificateUUID.
	AuthenticationCertificateUUID *string `json:"AuthenticationCertificateUUID,omitempty" plist:"AuthenticationCertificateUUID,omitempty"`
	// If `true`, the system enables Mail Drop.
	AllowMailDrop *bool `json:"allowMailDrop,omitempty" plist:"allowMailDrop,omitempty"`
	// The server path.
	Path *string `json:"Path,omitempty" plist:"Path,omitempty"`
	// The server port number.
	Port *int64 `json:"Port,omitempty" plist:"Port,omitempty"`
	// The external server address.
	ExternalHost *string `json:"ExternalHost,omitempty" plist:"ExternalHost,omitempty"`
	// If `true`, the system enables SSL for connections to the external server.
	ExternalSSL *bool `json:"ExternalSSL,omitempty" plist:"ExternalSSL,omitempty"`
	// The external server path.
	ExternalPath *string `json:"ExternalPath,omitempty" plist:"ExternalPath,omitempty"`
	// The external server port number.
	ExternalPort *int64 `json:"ExternalPort,omitempty" plist:"ExternalPort,omitempty"`
}

func (p *ExchangeWebServices) PayloadType() string {
	return "com.apple.ews.account"
}

// The payload that configures an app extension that performs single sign-on with the Kerberos extension.
// Configures the included Kerberos extension that performs SSO on behalf of specified hosts. User channel support was added in macOS 11.0.
type ExtensibleSingleSignOnKerberos struct {
	*CommonPayloadKeys
	// Set this to `com.apple.AppSSOKerberos.KerberosExtension` for this extension.
	ExtensionIdentifier ExtensionIdentifier `json:"ExtensionIdentifier" plist:"ExtensionIdentifier" required:"true"`
	// Set this to `apple` for this extension.
	TeamIdentifier TeamIdentifier `json:"TeamIdentifier" plist:"TeamIdentifier" required:"true"`
	// Set this to `Credential` for this extension.
	Type ExtensibleSingleSignOnKerberosType `json:"Type" plist:"Type" required:"true"`
	// The Kerberos realm. Use proper capitalization for this value. If in an Active Directory forest, this is the realm where the user logs in.
	Realm string `json:"Realm" plist:"Realm" required:"true"`
	// This is the dictionary used by the Apple built-in Kerberos extension.
	ExtensionData *ExtensionData `json:"ExtensionData,omitempty" plist:"ExtensionData,omitempty"`
	// One or more host or domain names for which the app extension performs SSO.
	// The system:
	// - Matches host or domain names case-insensitively
	// - Requires that all the host and domain names of all installed Extensible SSO payloads are unique
	// > Note:
	// > Host names that begin with a "." are wildcard suffixes that match all subdomains; otherwise the host name needs be an exact match.
	Hosts *[]string `json:"Hosts,omitempty" plist:"Hosts,omitempty"`
}

func (p *ExtensibleSingleSignOnKerberos) PayloadType() string {
	return "com.apple.extensiblesso"
}

// Set this to `com.apple.AppSSOKerberos.KerberosExtension` for this extension.
type ExtensionIdentifier string

const (
	ExtensionIdentifierComappleAppSSOKerberosKerberosExtension ExtensionIdentifier = "com.apple.AppSSOKerberos.KerberosExtension"
)

// Set this to `apple` for this extension.
type TeamIdentifier string

const (
	TeamIdentifierApple TeamIdentifier = "apple"
)

// Set this to `Credential` for this extension.
type ExtensibleSingleSignOnKerberosType string

const (
	ExtensibleSingleSignOnKerberosTypeCredential ExtensibleSingleSignOnKerberosType = "Credential"
)

// This is the dictionary used by the Apple built-in Kerberos extension.
type ExtensionData struct {
	// The GSS name of the Kerberos cache to use. Rarely set by an administrator.
	CacheName *string `json:"cacheName,omitempty" plist:"cacheName,omitempty"`
	// The principal (username) to use. You don't need to include the realm.
	PrincipalName *string `json:"principalName,omitempty" plist:"principalName,omitempty"`
	// The name of the Active Directory site the Kerberos extension should use. Most administrators don't need to modify this value, as the Kerberos extension can normally find the site automatically.
	SiteCode *string `json:"siteCode,omitempty" plist:"siteCode,omitempty"`
	// The PayloadUUID of a PKINIT certificate.
	CertificateUUID *string `json:"certificateUUID,omitempty" plist:"certificateUUID,omitempty"`
	// If `false`, the Kerberos extension doesn't automatically use LDAP and DNS to determine its AD site name.
	UseSiteAutoDiscovery *bool `json:"useSiteAutoDiscovery,omitempty" plist:"useSiteAutoDiscovery,omitempty"`
	// A list of bundle IDs allowed to access the ticket-granting ticket (TGT).
	CredentialBundleIdACL *[]string `json:"credentialBundleIdACL,omitempty" plist:"credentialBundleIdACL,omitempty"`
	// If `true`, the Kerberos extension allows only managed apps to access and use the credential. This is in addition to the `credentialBundleIDACL`, if you specify that value. Available in iOS 14 and later, and macOS 12 and later.
	IncludeManagedAppsInBundleIdACL *bool `json:"includeManagedAppsInBundleIdACL,omitempty" plist:"includeManagedAppsInBundleIdACL,omitempty"`
	// If `true`, the Kerberos extension allows the standard Kerberos utilities including `TicketViewer` and `klist` to access and use the credential. This is in addition to `includeManagedAppsInBundleIdACL` or the `credentialBundleIdACL`, if you specify those values. Available in macOS 12 and later.
	IncludeKerberosAppsInBundleIdACL *bool `json:"includeKerberosAppsInBundleIdACL,omitempty" plist:"includeKerberosAppsInBundleIdACL,omitempty"`
	// A custom domain-realm mapping for Kerberos. The system uses this when the DNS name of hosts doesn't match the realm name. Most administrators don't need to customize this.
	DomainRealmMapping *DomainRealmMapping `json:"domainRealmMapping,omitempty" plist:"domainRealmMapping,omitempty"`
	// Specifies whether this is the default realm if there's more than one Kerberos extension configuration.
	IsDefaultRealm *bool `json:"isDefaultRealm,omitempty" plist:"isDefaultRealm,omitempty"`
	// The custom user name label used in the Kerberos extension instead of "Username," such as "Company ID". Available in macOS 11 and later.
	CustomUsernameLabel *string `json:"customUsernameLabel,omitempty" plist:"customUsernameLabel,omitempty"`
	// The text to display to the user at the bottom of the Kerberos Login Window. You can also use this to display help information or disclaimer text. Available in iOS 14 and later, and macOS 11 and later.
	HelpText *string `json:"helpText,omitempty" plist:"helpText,omitempty"`
	// If `false`, the system disables password changes. Available in macOS 10.15 and later.
	AllowPasswordChange *bool `json:"allowPasswordChange,omitempty" plist:"allowPasswordChange,omitempty"`
	// If `false`, the system doesn't allow saving passwords in the keychain.
	AllowAutomaticLogin *bool `json:"allowAutomaticLogin,omitempty" plist:"allowAutomaticLogin,omitempty"`
	// If `true`, the system requires the user to provide Touch ID, Face ID or their passcode to access the keychain entry.
	RequireUserPresence *bool `json:"requireUserPresence,omitempty" plist:"requireUserPresence,omitempty"`
	// The number of days that the system allows using passwords on this domain. For most domains, this calculation is automatic. Available in macOS 10.15 and later.
	PwExpireOverride *int64 `json:"pwExpireOverride,omitempty" plist:"pwExpireOverride,omitempty"`
	// The number of days prior to password expiration when the system sends a notification of password expiration to the user. Available in macOS 10.15 and later.
	PwNotificationDays *int64 `default:"15" json:"pwNotificationDays,omitempty" plist:"pwNotificationDays,omitempty"`
	// The minimum length of passwords on the domain.Available in macOS 10.15 and later.
	PwReqLength *int64 `json:"pwReqLength,omitempty" plist:"pwReqLength,omitempty"`
	// If `true`, the system requires passwords to meet Active Directory's definition of "complex". Available in macOS 10.15 and later.
	PwReqComplexity *bool `json:"pwReqComplexity,omitempty" plist:"pwReqComplexity,omitempty"`
	// The minimum age of passwords before the system allows changing them on this domain. Available in macOS 10.15 and later.
	PwReqMinAge *int64 `json:"pwReqMinAge,omitempty" plist:"pwReqMinAge,omitempty"`
	// The number of prior passwords that the system disallows reuse on this domain. Available in macOS 10.15 and later.
	PwReqHistory *int64 `json:"pwReqHistory,omitempty" plist:"pwReqHistory,omitempty"`
	// The text version of the domain's password requirements. Only for use if `pwReqComplexity` or `pwReqLength` aren't specified. Available in macOS 10.15 and later.
	PwReqText *string `json:"pwReqText,omitempty" plist:"pwReqText,omitempty"`
	// The RTF file formatted version of the domain's password requirements. Only for use if `pwReqComplexity` or `pwReqLength` aren't specified. Available in macOS 15 and later.
	PwReqRTFData *[]byte `json:"pwReqRTFData,omitempty" plist:"pwReqRTFData,omitempty"`
	// This URL will launch in the user's default web browser when they initiate a password change. Available in macOS 10.15 and later.
	PwChangeURL *string `json:"pwChangeURL,omitempty" plist:"pwChangeURL,omitempty"`
	// If `false`, the system disables password sync. Note that this will not work if the user is logged in with a mobile account. Available in macOS 10.15 and later.
	SyncLocalPassword *bool `json:"syncLocalPassword,omitempty" plist:"syncLocalPassword,omitempty"`
	// The time, in seconds, required to replicate changes in the Active Directory domain. The Kerberos extension uses this when checking password age after a change. Available in macOS 11 and later.
	ReplicationTime *int64 `default:"900" json:"replicationTime,omitempty" plist:"replicationTime,omitempty"`
	// If `true`, the system doesn't prompt the user to setup the Kerberos extension until either the administrator enables it with the `app-sso` tool or the system receives a Kerberos challenge. Available in macOS 11 and later.
	DelayUserSetup *bool `json:"delayUserSetup,omitempty" plist:"delayUserSetup,omitempty"`
	// If `false`, the system requests the credential on the next matching Kerberos challenge or network state change. If the credential is expired or missing, the system creates a new one. Available in macOS 11 and later.
	MonitorCredentialsCache *bool `json:"monitorCredentialsCache,omitempty" plist:"monitorCredentialsCache,omitempty"`
	// Require that LDAP connections use TLS. Available in macOS 11 and later.
	RequireTLSForLDAP *bool `json:"requireTLSForLDAP,omitempty" plist:"requireTLSForLDAP,omitempty"`
	// This setting affects how other processes use the Kerberos Extension credential. Allowed values:
	// - `always`: The system always uses the credential if the SPN matches the Kerberos Extension `Hosts` array and the caller hasn't specified another credential. However, the system won't use the credential if the calling app isn't in the `credentialBundleIDACL`.
	// - `whenNotSpecified`: The system only uses the extension credential if the SPN matches the Kerberos Extension `Hosts` array. However, the system won't use the credential if the calling app isn't in the `credentialBundleIDACL`.
	// - `kerberosDefault`: The system uses the default Kerberos processes to select credentials, and normally uses the default Kerberos credential. This is the same as turning off this capability.
	// Available in macOS 11 and later.
	CredentialUseMode *CredentialUseMode `default:"always" json:"credentialUseMode,omitempty" plist:"credentialUseMode,omitempty"`
	// The ordered list of preferred Key Distribution Centers (KDCs) to use for Kerberos traffic. Use this if the servers aren't discoverable through DNS. If the servers are specified, then the system uses them for both connectivity checks and attempts to use them first for Kerberos traffic. If the servers don't respond, the device falls back to DNS discovery. Format each entry the same as it would be in a `krb5.conf` file, for example:
	// - `adserver1.example.com`
	// - `tcp/adserver1.example.com:88`
	// - `kkdcp://kerberosproxy.example.com:443/kkdcp`
	PreferredKDCs *[]string `json:"preferredKDCs,omitempty" plist:"preferredKDCs,omitempty"`
	// If `true`, the system requires this configuration uses a TGT from Platform SSO instead of requesting a new one. Available in macOS 13 and later.
	UsePlatformSSOTGT *bool `json:"usePlatformSSOTGT,omitempty" plist:"usePlatformSSOTGT,omitempty"`
	// If `true` and `usePlatformSSOTGT` is `true`, the system allows the user to manually sign in. Available in macOS 13 and later.
	AllowPlatformSSOAuthFallback *bool `json:"allowPlatformSSOAuthFallback,omitempty" plist:"allowPlatformSSOAuthFallback,omitempty"`
	// If `true`, the Kerberos Extension handles Kerberos requests only. It doesn't check for password expiration, show the password expiration in the menu, check for external password changes, perform password sync, or retrieve the home directory. Available in macOS 13 and later.
	PerformKerberosOnly *bool `json:"performKerberosOnly,omitempty" plist:"performKerberosOnly,omitempty"`
	// A string with wildcards that can use used to filter the list of available SmartCards by issuer. e.g "\*My CA2\*". If there is one remaining, it will be auto-selected. If there more than one remaining, then the list is shorter. Available in macOS 15 and later.
	IdentityIssuerAutoSelectFilter *string `json:"identityIssuerAutoSelectFilter,omitempty" plist:"identityIssuerAutoSelectFilter,omitempty"`
	// If `true`, allow the user to switch the user interface to SmartCard mode. Available in macOS 15 and later.
	AllowSmartCard *bool `json:"allowSmartCard,omitempty" plist:"allowSmartCard,omitempty"`
	// If `true`, allow the user to switch the user interface to Password mode. Available in macOS 15 and later.
	AllowPassword *bool `json:"allowPassword,omitempty" plist:"allowPassword,omitempty"`
	// If `true`, the user interface will start in SmartCard mode. Available in macOS 15 and later.
	StartInSmartCardMode *bool `json:"startInSmartCardMode,omitempty" plist:"startInSmartCardMode,omitempty"`
}

// A custom domain-realm mapping for Kerberos. The system uses this when the DNS name of hosts doesn't match the realm name. Most administrators don't need to customize this.
type DomainRealmMapping struct {
	// The key should be the name of the realm, and the value is an array of DNS suffixes that map to the realm.
	Realm *[]string `json:"Realm,omitempty" plist:"Realm,omitempty"`
}

// This setting affects how other processes use the Kerberos Extension credential. Allowed values:
// - `always`: The system always uses the credential if the SPN matches the Kerberos Extension `Hosts` array and the caller hasn't specified another credential. However, the system won't use the credential if the calling app isn't in the `credentialBundleIDACL`.
// - `whenNotSpecified`: The system only uses the extension credential if the SPN matches the Kerberos Extension `Hosts` array. However, the system won't use the credential if the calling app isn't in the `credentialBundleIDACL`.
// - `kerberosDefault`: The system uses the default Kerberos processes to select credentials, and normally uses the default Kerberos credential. This is the same as turning off this capability.
// Available in macOS 11 and later.
type CredentialUseMode string

const (
	CredentialUseModeAlways           CredentialUseMode = "always"
	CredentialUseModeWhenNotSpecified CredentialUseMode = "whenNotSpecified"
	CredentialUseModeKerberosDefault  CredentialUseMode = "kerberosDefault"
)

// The payload that configures an app extension that performs single sign-on (SSO).
// Configures an app extension that performs SSO on behalf of certain URLs. User channel support was added in macOS 11.0.
type ExtensibleSingleSignOn struct {
	*CommonPayloadKeys
	// The bundle identifier of the app extension that performs SSO for the specified URLs.
	ExtensionIdentifier string `json:"ExtensionIdentifier" plist:"ExtensionIdentifier" required:"true"`
	// The team identifier of the app extension. This key is required on macOS and ignored elsewhere.
	TeamIdentifier *string `json:"TeamIdentifier,omitempty" plist:"TeamIdentifier,omitempty"`
	// The type of SSO.
	Type ExtensibleSingleSignOnType `json:"Type" plist:"Type" required:"true"`
	// The realm name for `Credential` payloads. Use proper capitalization for this value. Ignored for `Redirect` payloads.
	Realm *string `json:"Realm,omitempty" plist:"Realm,omitempty"`
	// A dictionary of arbitrary data passed through to the app extension.
	ExtensionData *map[string]any `json:"ExtensionData,omitempty" plist:"ExtensionData,omitempty"`
	// An array of URL prefixes of identity providers where the app extension performs SSO.
	// Required for `Redirect` payloads. Ignored for `Credential` payloads.
	// The URLs need to begin with `http://` or `https://`.
	// The system:
	// - Matches scheme and host name case-insensitively
	// - Doesn't allow query parameters and URL fragments
	// - Requires that the URLs of all installed Extensible SSO payloads are unique
	URLs *[]string `json:"URLs,omitempty" plist:"URLs,omitempty"`
	// An array of host or domain names that apps can authenticate through the app extension.
	// Required for `Credential` payloads. Ignored for `Redirect` payloads.
	// The system:
	// - Matches host or domain names case-insensitively
	// - Requires that all the host and domain names of all installed Extensible SSO payloads are unique
	// > Note:
	// > Host names that begin with a "." are wildcard suffixes that match all subdomains; otherwise the host name needs be an exact match.
	Hosts *[]string `json:"Hosts,omitempty" plist:"Hosts,omitempty"`
	// If set to `Cancel`, the system cancels authentication requests when the screen is locked. If set to `DoNotHandle`, the request continues without SSO instead. This doesn't apply to requests where `userInterfaceEnabled` is `false`, or for background `URLSession` requests. Available in iOS 15 and later, and macOS 12 and later.
	ScreenLockedBehavior *ScreenLockedBehavior `default:"Cancel" json:"ScreenLockedBehavior,omitempty" plist:"ScreenLockedBehavior,omitempty"`
	// An array of bundle identifiers of apps that don't use SSO provided by this extension. Available in iOS 15 and later, and macOS 12 and later.
	DeniedBundleIdentifiers *[]string `json:"DeniedBundleIdentifiers,omitempty" plist:"DeniedBundleIdentifiers,omitempty"`
	// The Platform SSO authentication method the extension uses. Requires that the SSO Extension also supports the method. Available in macOS 13 and later, and deprecated in macOS 14.
	AuthenticationMethod *ExtensibleSingleSignOnAuthenticationMethod `json:"AuthenticationMethod,omitempty" plist:"AuthenticationMethod,omitempty"`
	// The token this device uses for registration with Platform SSO. Use it for silent registration with the Identity Provider. Requires that `AuthenticationMethod` in `PlatformSSO` isn't empty. Available in macOS 13 and later.
	RegistrationToken *string `json:"RegistrationToken,omitempty" plist:"RegistrationToken,omitempty"`
	// The dictionary to configure Platform SSO. Requires `Type` to be set to `Redirect`.
	PlatformSSO *PlatformSSO `json:"PlatformSSO,omitempty" plist:"PlatformSSO,omitempty"`
}

func (p *ExtensibleSingleSignOn) PayloadType() string {
	return "com.apple.extensiblesso"
}

// The type of SSO.
type ExtensibleSingleSignOnType string

const (
	ExtensibleSingleSignOnTypeCredential ExtensibleSingleSignOnType = "Credential"
	ExtensibleSingleSignOnTypeRedirect   ExtensibleSingleSignOnType = "Redirect"
)

// If set to `Cancel`, the system cancels authentication requests when the screen is locked. If set to `DoNotHandle`, the request continues without SSO instead. This doesn't apply to requests where `userInterfaceEnabled` is `false`, or for background `URLSession` requests. Available in iOS 15 and later, and macOS 12 and later.
type ScreenLockedBehavior string

const (
	ScreenLockedBehaviorCancel      ScreenLockedBehavior = "Cancel"
	ScreenLockedBehaviorDoNotHandle ScreenLockedBehavior = "DoNotHandle"
)

// The Platform SSO authentication method the extension uses. Requires that the SSO Extension also supports the method. Available in macOS 13 and later, and deprecated in macOS 14.
type ExtensibleSingleSignOnAuthenticationMethod string

const (
	ExtensibleSingleSignOnAuthenticationMethodPassword             ExtensibleSingleSignOnAuthenticationMethod = "Password"
	ExtensibleSingleSignOnAuthenticationMethodUserSecureEnclaveKey ExtensibleSingleSignOnAuthenticationMethod = "UserSecureEnclaveKey"
)

// The dictionary to configure Platform SSO. Requires `Type` to be set to `Redirect`.
type PlatformSSO struct {
	// The Platform SSO authentication method to use with the extension. Requires that the SSO Extension also support the method.
	AuthenticationMethod *PlatformSSOAuthenticationMethod `json:"AuthenticationMethod,omitempty" plist:"AuthenticationMethod,omitempty"`
	// If `true`, the system uses the same signing and encryption keys for all users. Only supported on the device channel.
	UseSharedDeviceKeys *bool `json:"UseSharedDeviceKeys,omitempty" plist:"UseSharedDeviceKeys,omitempty"`
	// The display name for the account in notifications and authentication requests.
	AccountDisplayName *string `json:"AccountDisplayName,omitempty" plist:"AccountDisplayName,omitempty"`
	// The duration, in seconds, until the system requires a full login instead of a refresh. The default value is 64,800 (18 hours). The minimum value is 3600 (1 hour).
	LoginFrequency *int64 `default:"64800" json:"LoginFrequency,omitempty" plist:"LoginFrequency,omitempty"`
	// Enables creating users at the Login Window with an `AuthenticationMethod` of either `Password` or `SmartCard`. Requires that `UseSharedDeviceKeys` is `true`.
	EnableCreateUserAtLogin *bool `json:"EnableCreateUserAtLogin,omitempty" plist:"EnableCreateUserAtLogin,omitempty"`
	// If `true`, the device uses Platform SSO to create the first user account on the Mac during `Setup Assistant`.
	EnableCreateFirstUserDuringSetup *bool `json:"EnableCreateFirstUserDuringSetup,omitempty" plist:"EnableCreateFirstUserDuringSetup,omitempty"`
	// Enables using identity provider accounts at authorization prompts. Requires that `UseSharedDeviceKeys` is `true`. The system assigns groups using `AdministratorGroups`, `AdditionalGroups`, or `AuthorizationGroups`.
	EnableAuthorization *bool `json:"EnableAuthorization,omitempty" plist:"EnableAuthorization,omitempty"`
	// The attribute mapping to use when creating users, or for authorization.
	TokenToUserMapping *TokenToUserMapping `json:"TokenToUserMapping,omitempty" plist:"TokenToUserMapping,omitempty"`
	// The set of authentication methods to use for newly created accounts at login or during `Setup Assistant`. The system uses `Password` and `SmartCard` if this key isn't present.
	NewUserAuthenticationMethods *[]NewUserAuthenticationMethods `json:"NewUserAuthenticationMethods,omitempty" plist:"NewUserAuthenticationMethods,omitempty"`
	// The permission to apply to newly created accounts at login. Allowed values:
	// - `Standard`: The account is a standard user.
	// - `Admin`: The system adds the account to the local administrators group.
	// - `Groups`: The system assigns groups to the account using `AdministratorGroups`, `AdditionalGroups`, or `AuthorizationGroups`.
	// - `Temporary`: The system uses a temporary session configuration for newly created accounts at login.
	NewUserAuthorizationMode *NewUserAuthorizationMode `json:"NewUserAuthorizationMode,omitempty" plist:"NewUserAuthorizationMode,omitempty"`
	// The permission to apply to an account each time the user authenticates. Allowed values:
	// - `Standard`: The account is a standard user.
	// - `Admin`: The system adds the account to the local administrators group.
	// - `Groups`: The system assigns group to the account using `AdministratorGroups`, `AdditionalGroups`, or `AuthorizationGroups`.
	UserAuthorizationMode *UserAuthorizationMode `json:"UserAuthorizationMode,omitempty" plist:"UserAuthorizationMode,omitempty"`
	// The list of groups to use for administrator access. The system requests membership during authentication.
	AdministratorGroups *[]string `json:"AdministratorGroups,omitempty" plist:"AdministratorGroups,omitempty"`
	// The list of created groups that don't have administrator access.
	AdditionalGroups *[]string `json:"AdditionalGroups,omitempty" plist:"AdditionalGroups,omitempty"`
	// The pairing of Authorization Rights to group names. When using this, the system updates the Authorization Right to use the group.
	AuthorizationGroups *map[string]string `json:"AuthorizationGroups,omitempty" plist:"AuthorizationGroups,omitempty"`
	// The reader group identifier for use with the `AccessKey`. The value needs to match the configured access key. Required if `NewUserAuthenticationMethods` contains `AccessKey`.
	AccessKeyReaderGroupIdentifier *[]byte `json:"AccessKeyReaderGroupIdentifier,omitempty" plist:"AccessKeyReaderGroupIdentifier,omitempty"`
	// The `PayloadUUID` of an identity payload to use as the `Terminal` identity of the access key. The identity needs to be trusted by the access key. Required if `NewUserAuthenticationMethods` includes `AccessKey`. Allowed identity payload types:
	// - `com.apple.security.pkcs12`
	// - `com.apple.security.acme`
	// - `com.apple.security.scep`
	AccessKeyTerminalIdentityUUID *string `json:"AccessKeyTerminalIdentityUUID,omitempty" plist:"AccessKeyTerminalIdentityUUID,omitempty"`
	// The `PayloadUUID` of a certificate payload for the issuer certificate of the `Terminal` identity of the access key. Other specifications refer to the key as the "Reader CA Public Key". The key must be an elliptic curve key. Required if `NewUserAuthenticationMethods` includes `AccessKey`. The issuer of the Terminal identity of the access key needs to match this certificate, otherwise the device fails the authentication.
	AccessKeyReaderIssuerCertificateUUID *string `json:"AccessKeyReaderIssuerCertificateUUID,omitempty" plist:"AccessKeyReaderIssuerCertificateUUID,omitempty"`
	// If `true`, the system uses the access key in express mode, and doesn't require authentication before use.
	AllowAccessKeyExpressMode *bool `json:"AllowAccessKeyExpressMode,omitempty" plist:"AllowAccessKeyExpressMode,omitempty"`
	// The policy to apply when using Platform SSO at FileVault unlock on a Mac with Apple silicon. Applies when `AuthenticationMethod` is `Password`. Available in macOS 15 and later.
	FileVaultPolicy *[]FileVaultPolicy `json:"FileVaultPolicy,omitempty" plist:"FileVaultPolicy,omitempty"`
	// The policy to apply when using Platform SSO at the Login Window. Applies when `AuthenticationMethod` is `Password`. Available in macOS 15 and later.
	LoginPolicy *[]LoginPolicy `json:"LoginPolicy,omitempty" plist:"LoginPolicy,omitempty"`
	// The policy to apply when using Platform SSO at screensaver unlock. Applies when `AuthenticationMethod` is `Password`. Available in macOS 15 and later.
	UnlockPolicy *[]UnlockPolicy `json:"UnlockPolicy,omitempty" plist:"UnlockPolicy,omitempty"`
	// The amount of time after the last successful Platform SSO login for using a local account password offline. Required when setting `AllowOfflineGracePeriod`. Available in macOS 15 and later.
	OfflineGracePeriod *int64 `json:"OfflineGracePeriod,omitempty" plist:"OfflineGracePeriod,omitempty"`
	// The amount of time after receiving or updating a `FileVaultPolicy`, `LoginPolicy`, or `UnlockPolicy` that the system can use unregistered local accounts. Required when `AllowAuthenticationGracePeriod` is set. Available in macOS 15 and later.
	AuthenticationGracePeriod *int64 `json:"AuthenticationGracePeriod,omitempty" plist:"AuthenticationGracePeriod,omitempty"`
	// The list of local accounts that aren't subject to the `FileVaultPolicy`, `LoginPolicy`, or `UnlockPolicy`. The accounts don't receive a prompt to register for Platform SSO. Available in macOS 15 and later.
	NonPlatformSSOAccounts *[]string `json:"NonPlatformSSOAccounts,omitempty" plist:"NonPlatformSSOAccounts,omitempty"`
	// If `true`, the system includes the device UDID and serial number in Platform SSO attestations.
	AllowDeviceIdentifiersInAttestation *bool `json:"AllowDeviceIdentifiersInAttestation,omitempty" plist:"AllowDeviceIdentifiersInAttestation,omitempty"`
	// If `true`, the system requests the user's profile picture from the SSO extension.
	SynchronizeProfilePicture *bool `json:"SynchronizeProfilePicture,omitempty" plist:"SynchronizeProfilePicture,omitempty"`
	// If `true`, the system uses a quicker Authenticated Guest Mode login to Mac behavior. The system erases user data from only select locations in the user home directory after each session completes. Once every eight hours the system erases the full user home directory after a session completes. Turn this on for shared environments that have a high frequency of short sessions.
	TemporarySessionQuickLogin *bool `json:"TemporarySessionQuickLogin,omitempty" plist:"TemporarySessionQuickLogin,omitempty"`
	// If `true`, the system enables the PlatformSSO registration process during Setup Assistant on devices running macOS 26 and later. Set this key to `true` when configuring PlatformSSO before enrollment using the `com.apple.psso.required` error response.
	EnableRegistrationDuringSetup *bool `json:"EnableRegistrationDuringSetup,omitempty" plist:"EnableRegistrationDuringSetup,omitempty"`
}

// The Platform SSO authentication method to use with the extension. Requires that the SSO Extension also support the method.
type PlatformSSOAuthenticationMethod string

const (
	PlatformSSOAuthenticationMethodPassword             PlatformSSOAuthenticationMethod = "Password"
	PlatformSSOAuthenticationMethodUserSecureEnclaveKey PlatformSSOAuthenticationMethod = "UserSecureEnclaveKey"
	PlatformSSOAuthenticationMethodSmartCard            PlatformSSOAuthenticationMethod = "SmartCard"
)

// The attribute mapping to use when creating users, or for authorization.
type TokenToUserMapping struct {
	// The claim name to use for the user's account name.
	AccountName *string `json:"AccountName,omitempty" plist:"AccountName,omitempty"`
	// The claim name to use for the user's full name.
	FullName *string `json:"FullName,omitempty" plist:"FullName,omitempty"`
}

// The set of authentication methods to use for newly created accounts at login or during `Setup Assistant`. The system uses `Password` and `SmartCard` if this key isn't present.
// An authentication method to use for newly created accounts at login or during `Setup Assistant`. Allowed values:
// - `Password`: The account uses a password for authentication.
// - `SmartCard`: The account uses a smart card for authentication.
// - `AccessKey`: The account uses an access key for authentication.
type NewUserAuthenticationMethods string

const (
	NewUserAuthenticationMethodsPassword  NewUserAuthenticationMethods = "Password"
	NewUserAuthenticationMethodsSmartCard NewUserAuthenticationMethods = "SmartCard"
	NewUserAuthenticationMethodsAccessKey NewUserAuthenticationMethods = "AccessKey"
)

// The permission to apply to newly created accounts at login. Allowed values:
// - `Standard`: The account is a standard user.
// - `Admin`: The system adds the account to the local administrators group.
// - `Groups`: The system assigns groups to the account using `AdministratorGroups`, `AdditionalGroups`, or `AuthorizationGroups`.
// - `Temporary`: The system uses a temporary session configuration for newly created accounts at login.
type NewUserAuthorizationMode string

const (
	NewUserAuthorizationModeStandard  NewUserAuthorizationMode = "Standard"
	NewUserAuthorizationModeAdmin     NewUserAuthorizationMode = "Admin"
	NewUserAuthorizationModeGroups    NewUserAuthorizationMode = "Groups"
	NewUserAuthorizationModeTemporary NewUserAuthorizationMode = "Temporary"
)

// The permission to apply to an account each time the user authenticates. Allowed values:
// - `Standard`: The account is a standard user.
// - `Admin`: The system adds the account to the local administrators group.
// - `Groups`: The system assigns group to the account using `AdministratorGroups`, `AdditionalGroups`, or `AuthorizationGroups`.
type UserAuthorizationMode string

const (
	UserAuthorizationModeStandard UserAuthorizationMode = "Standard"
	UserAuthorizationModeAdmin    UserAuthorizationMode = "Admin"
	UserAuthorizationModeGroups   UserAuthorizationMode = "Groups"
)

// The policy to apply when using Platform SSO at FileVault unlock on a Mac with Apple silicon. Applies when `AuthenticationMethod` is `Password`. Available in macOS 15 and later.
// * AttemptAuthentication
// Platform SSO authentication is attempted before proceeding. If offline, unlock will continue
// if the local account password matches. If online and the credential is incorrect, then a
// successful Platform SSO authentication is required to proceed, even if taken offline.
// * RequireAuthentication
// Platform SSO authentication is required before proceeding.  If the device is offline and
// `AllowOfflineGracePeriod` is enabled, then the offline `OfflineGracePeriod` is used to determine
// if the user can proceed or not. If online and the credential is incorrect, then a valid Platform
// SSO authentication is required to proceed regardless of the `OfflineGracePeriod`. If the account
// is not registered for Platform SSO and `AllowAuthenticationGracePeriod` is enabled, then the
// `AuthenticationGracePeriod` is used to determine if the user can proceed or not.
// * AllowOfflineGracePeriod
// Allow the use of the `OfflineGracePeriod` when `RequireAuthentication` is enabled.  If
// `AllowOfflineGracePeriod` is not set, then offline access is denied.
// * AllowAuthenticationGracePeriod
// Allow the use of the `AuthenticationGracePeriod` for other local accounts when `RequireAuthentication`
// is enabled. The `AuthenticationGracePeriod` starts when any of the policies have been updated. If
// `AllowAuthenticationGracePeriod` is not set, then unregistered account access is denied.
type FileVaultPolicy string

const (
	FileVaultPolicyAttemptAuthentication          FileVaultPolicy = "AttemptAuthentication"
	FileVaultPolicyRequireAuthentication          FileVaultPolicy = "RequireAuthentication"
	FileVaultPolicyAllowOfflineGracePeriod        FileVaultPolicy = "AllowOfflineGracePeriod"
	FileVaultPolicyAllowAuthenticationGracePeriod FileVaultPolicy = "AllowAuthenticationGracePeriod"
)

// The policy to apply when using Platform SSO at the Login Window. Applies when `AuthenticationMethod` is `Password`. Available in macOS 15 and later.
// * AttemptAuthentication
// Platform SSO authentication is attempted before proceeding. If offline, unlock will continue
// if the local account password matches. If online and the credential is incorrect, then a
// successful Platform SSO authentication is required to proceed, even if taken offline.
// * RequireAuthentication
// Platform SSO authentication is required before proceeding.  If the device is offline and
// `AllowOfflineGracePeriod` is enabled, then the offline `OfflineGracePeriod` is used to determine
// if the user can proceed or not. If online and the credential is incorrect, then a valid Platform
// SSO authentication is required to proceed regardless of the `OfflineGracePeriod`. If the account
// is not registered for Platform SSO and `AllowAuthenticationGracePeriod` is enabled, then the
// `AuthenticationGracePeriod` is used to determine if the user can proceed or not.
// * AllowOfflineGracePeriod
// Allow the use of the `OfflineGracePeriod` when `RequireAuthentication` is enabled.  If
// `AllowOfflineGracePeriod` is not set, then offline access is denied.
// * AllowAuthenticationGracePeriod
// Allow the use of the `AuthenticationGracePeriod` for other local accounts when `RequireAuthentication`
// is enabled. The `AuthenticationGracePeriod` starts when any of the policies have been updated. If
// `AllowAuthenticationGracePeriod` is not set, then unregistered account access is denied.
type LoginPolicy string

const (
	LoginPolicyAttemptAuthentication          LoginPolicy = "AttemptAuthentication"
	LoginPolicyRequireAuthentication          LoginPolicy = "RequireAuthentication"
	LoginPolicyAllowOfflineGracePeriod        LoginPolicy = "AllowOfflineGracePeriod"
	LoginPolicyAllowAuthenticationGracePeriod LoginPolicy = "AllowAuthenticationGracePeriod"
)

// The policy to apply when using Platform SSO at screensaver unlock. Applies when `AuthenticationMethod` is `Password`. Available in macOS 15 and later.
// * AttemptAuthentication
// Platform SSO authentication is attempted before proceeding. If offline, unlock will continue
// if the local account password matches. If online and the credential is incorrect, then a
// successful Platform SSO authentication is required to proceed, even if taken offline.
// * RequireAuthentication
// Platform SSO authentication is required before proceeding.  If the device is offline and
// `AllowOfflineGracePeriod` is enabled, then the offline `OfflineGracePeriod` is used to determine
// if the user can proceed or not. If online and the credential is incorrect, then a valid Platform
// SSO authentication is required to proceed regardless of the `OfflineGracePeriod`. If the account
// is not registered for Platform SSO and `AllowAuthenticationGracePeriod` is enabled, then the
// `AuthenticationGracePeriod` is used to determine if the user can proceed or not.
// * AllowOfflineGracePeriod
// Allow the use of the `OfflineGracePeriod` when `RequireAuthentication` is enabled.  If
// `AllowOfflineGracePeriod` is not set, then offline access is denied.
// * AllowAuthenticationGracePeriod
// Allow the use of the `AuthenticationGracePeriod` for other local accounts when `RequireAuthentication`
// is enabled. The `AuthenticationGracePeriod` starts when any of the policies have been updated. If
// `AllowAuthenticationGracePeriod` is not set, then unregistered account access is denied.
// * AllowTouchIDOrWatchForUnlock
// Allow TouchID or Watch to unlock the screensaver instead of Platform SSO authentication when
// `RequireAuthentication` is enabled.
type UnlockPolicy string

const (
	UnlockPolicyAttemptAuthentication          UnlockPolicy = "AttemptAuthentication"
	UnlockPolicyRequireAuthentication          UnlockPolicy = "RequireAuthentication"
	UnlockPolicyAllowOfflineGracePeriod        UnlockPolicy = "AllowOfflineGracePeriod"
	UnlockPolicyAllowAuthenticationGracePeriod UnlockPolicy = "AllowAuthenticationGracePeriod"
	UnlockPolicyAllowTouchIDOrWatchForUnlock   UnlockPolicy = "AllowTouchIDOrWatchForUnlock"
)

// The payload that configures the parental control web content filters.
// Parental controls web filter.
type ParentalControlsContentFilter struct {
	*CommonPayloadKeys
	// If `true`, enables web content filters.
	RestrictWeb bool `json:"restrictWeb" plist:"restrictWeb" required:"true"`
	// If `true`, filters content automatically.
	UseContentFilter *bool `json:"useContentFilter,omitempty" plist:"useContentFilter,omitempty"`
	// If `true`, enables web content filters.
	AllowlistEnabled *bool `json:"allowlistEnabled,omitempty" plist:"allowlistEnabled,omitempty"`
	// Use `allowlistEnabled` instead.
	WhitelistEnabled *bool `json:"whitelistEnabled,omitempty" plist:"whitelistEnabled,omitempty"`
	// An array of sites that defines an allow list. If specified, this defines additional allowed sites besides those in the automated allow list and deny list, including disallowed adult sites.
	// This key is required if `allowlistEnabled` is `true`.
	SiteAllowlist *[]*SiteAllowlistItem `json:"siteAllowlist,omitempty" plist:"siteAllowlist,omitempty"`
	// Use `siteAllowlist` instead.
	SiteWhitelist *[]*SiteWhitelistItem `json:"siteWhitelist,omitempty" plist:"siteWhitelist,omitempty"`
	// The array of URLs that defines an allow list. When `restrictWeb` and `useContentFilter` are enabled, only URLs in the allow list are available to the user.
	FilterAllowlist *[]string `json:"filterAllowlist,omitempty" plist:"filterAllowlist,omitempty"`
	// Use `filterAllowlist` instead.
	FilterWhitelist *[]string `json:"filterWhitelist,omitempty" plist:"filterWhitelist,omitempty"`
	// The array of URLs that defines a deny list. When `restrictWeb` and `useContentFilter` are enabled, no URLs in the deny list are available to the user.
	FilterDenylist *[]string `json:"filterDenylist,omitempty" plist:"filterDenylist,omitempty"`
	// Use `filterDenylist` instead.
	FilterBlacklist *[]string `json:"filterBlacklist,omitempty" plist:"filterBlacklist,omitempty"`
}

func (p *ParentalControlsContentFilter) PayloadType() string {
	return "com.apple.familycontrols.contentfilter"
}

// A dictionary defining a site for the allow list.
type SiteAllowlistItem struct {
	// The site prefix, including the `http(s)` scheme.
	Address string `json:"address" plist:"address" required:"true"`
	// The site page title.
	PageTitle *string `json:"pageTitle,omitempty" plist:"pageTitle,omitempty"`
}

// A dictionary defining a site for the allow list.
type SiteWhitelistItem struct {
	// The site prefix, including http(s) scheme.
	Address string `json:"address" plist:"address" required:"true"`
	// The site page title.
	PageTitle *string `json:"pageTitle,omitempty" plist:"pageTitle,omitempty"`
}

// The payload that configures parental control time limits.
// Parental controls time limits.
type ParentalControlsTimeLimits struct {
	*CommonPayloadKeys
	// If `true`, enables time limits.
	FamilyControlsEnabled bool `json:"familyControlsEnabled" plist:"familyControlsEnabled" required:"true"`
	// The time limits to enforce if `familyControlsEnabled` is enabled.
	TimeLimits *TimeLimits `json:"time-limits,omitempty" plist:"time-limits,omitempty"`
}

func (p *ParentalControlsTimeLimits) PayloadType() string {
	return "com.apple.familycontrols.timelimits.v2"
}

// The time limits to enforce if `familyControlsEnabled` is enabled.
type TimeLimits struct {
	// The weekday allowance settings.
	WeekdayAllowance *WeekdayAllowance `json:"weekday-allowance,omitempty" plist:"weekday-allowance,omitempty"`
	// The weekday curfew settings.
	WeekdayCurfew *WeekdayCurfew `json:"weekday-curfew,omitempty" plist:"weekday-curfew,omitempty"`
	// The weekend allowance settings.
	WeekendAllowance *WeekendAllowance `json:"weekend-allowance,omitempty" plist:"weekend-allowance,omitempty"`
	// The weekend curfew settings.
	WeekendCurfew *WeekendCurfew `json:"weekend-curfew,omitempty" plist:"weekend-curfew,omitempty"`
}

// The weekday allowance settings.
type WeekdayAllowance struct {
	// If `true`, enable these settings.
	Enabled bool `json:"enabled" plist:"enabled" required:"true"`
	// The type of day range, which has the following possible values:
	// - `0`: Weekday
	// - `1`: Weekend
	RangeType RangeType `json:"rangeType" plist:"rangeType" required:"true"`
	// The curfew start time, in the format '%d:%d:%d'.
	Start *string `json:"start,omitempty" plist:"start,omitempty"`
	// The curfew end time, in the format `%d:%d:%d`.
	End *string `json:"end,omitempty" plist:"end,omitempty"`
	// The allowance for that day, in seconds.
	SecondsPerDay *int64 `json:"secondsPerDay,omitempty" plist:"secondsPerDay,omitempty"`
}

// The type of day range, which has the following possible values:
// - `0`: Weekday
// - `1`: Weekend
type RangeType int64

const (
	RangeType0 RangeType = 0
	RangeType1 RangeType = 1
)

// The weekday curfew settings.
type WeekdayCurfew struct {
	// If `true`, enable these settings.
	Enabled bool `json:"enabled" plist:"enabled" required:"true"`
	// The type of day range, which has the following possible values:
	// - `0`: Weekday
	// - `1`: Weekend
	RangeType RangeType `json:"rangeType" plist:"rangeType" required:"true"`
	// The curfew start time, in the format '%d:%d:%d'.
	Start *string `json:"start,omitempty" plist:"start,omitempty"`
	// The curfew end time, in the format `%d:%d:%d`.
	End *string `json:"end,omitempty" plist:"end,omitempty"`
	// The allowance for that day, in seconds.
	SecondsPerDay *int64 `json:"secondsPerDay,omitempty" plist:"secondsPerDay,omitempty"`
}

// The weekend allowance settings.
type WeekendAllowance struct {
	// If `true`, enable these settings.
	Enabled bool `json:"enabled" plist:"enabled" required:"true"`
	// The type of day range, which has the following possible values:
	// - `0`: Weekday
	// - `1`: Weekend
	RangeType RangeType `json:"rangeType" plist:"rangeType" required:"true"`
	// The curfew start time, in the format '%d:%d:%d'.
	Start *string `json:"start,omitempty" plist:"start,omitempty"`
	// The curfew end time, in the format `%d:%d:%d`.
	End *string `json:"end,omitempty" plist:"end,omitempty"`
	// The allowance for that day, in seconds.
	SecondsPerDay *int64 `json:"secondsPerDay,omitempty" plist:"secondsPerDay,omitempty"`
}

// The weekend curfew settings.
type WeekendCurfew struct {
	// If `true`, enable these settings.
	Enabled bool `json:"enabled" plist:"enabled" required:"true"`
	// The type of day range, which has the following possible values:
	// - `0`: Weekday
	// - `1`: Weekend
	RangeType RangeType `json:"rangeType" plist:"rangeType" required:"true"`
	// The curfew start time, in the format '%d:%d:%d'.
	Start *string `json:"start,omitempty" plist:"start,omitempty"`
	// The curfew end time, in the format `%d:%d:%d`.
	End *string `json:"end,omitempty" plist:"end,omitempty"`
	// The allowance for that day, in seconds.
	SecondsPerDay *int64 `json:"secondsPerDay,omitempty" plist:"secondsPerDay,omitempty"`
}

// The payload that configures file provider settings.
type FileProvider struct {
	*CommonPayloadKeys
	// If `true`, enables file providers access to the path of the requesting process.
	AllowManagedFileProvidersToRequestAttribution *bool `json:"AllowManagedFileProvidersToRequestAttribution,omitempty" plist:"AllowManagedFileProvidersToRequestAttribution,omitempty"`
	// If `false`, the device prevents the File Provider extension from using desktop and documents synchronization in any app. This does not impact the ability for apps to utilize the File Provider extension for file and folder syncing with remote storage.
	ManagementAllowsKnownFolderSyncing *bool `json:"ManagementAllowsKnownFolderSyncing,omitempty" plist:"ManagementAllowsKnownFolderSyncing,omitempty"`
	// An array of strings representing the composed identifiers of apps. The device allows the corresponding apps to use File Provider extension desktop and documents synchronization. If present, and `ManagementAllowsKnownFolderSyncing` is set to `true`, the device allows only the apps in this list to use desktop and documents synchronization. This key is ignored if `ManagementAllowsKnownFolderSyncing` is set to `false`. This setting does not impact the ability for apps to use File Provider extension volume access. The format of the app identifiers is "Bundle-ID (Team-ID)", for example `com.example.app (ABCD1234)`.
	ManagementKnownFolderSyncingAllowList *[]string `json:"ManagementKnownFolderSyncingAllowList,omitempty" plist:"ManagementKnownFolderSyncingAllowList,omitempty"`
}

func (p *FileProvider) PayloadType() string {
	return "com.apple.fileproviderd"
}

// The payload that configures Finder settings.
type Finder struct {
	*CommonPayloadKeys
	// If `true`, the system disables the Finder's burn support.
	ProhibitBurn *bool `json:"ProhibitBurn,omitempty" plist:"ProhibitBurn,omitempty"`
	// Specifies whether Finder should operate in Simple or Full mode.
	InterfaceLevel *InterfaceLevel `json:"InterfaceLevel,omitempty" plist:"InterfaceLevel,omitempty"`
	// If `true`, the system disables Connect to Server.
	ProhibitConnectTo *bool `json:"ProhibitConnectTo,omitempty" plist:"ProhibitConnectTo,omitempty"`
	// If `true`, the system disables Eject.
	ProhibitEject *bool `json:"ProhibitEject,omitempty" plist:"ProhibitEject,omitempty"`
	// If `true`, the system disables Go to Folder.
	ProhibitGoToFolder *bool `json:"ProhibitGoToFolder,omitempty" plist:"ProhibitGoToFolder,omitempty"`
	// If `false`, the system doesn't show external hard drives on the Desktop.
	ShowExternalHardDrivesOnDesktop *bool `json:"ShowExternalHardDrivesOnDesktop,omitempty" plist:"ShowExternalHardDrivesOnDesktop,omitempty"`
	// If `false`, the system doesn't show internal hard drives on the Desktop.
	ShowHardDrivesOnDesktop *bool `json:"ShowHardDrivesOnDesktop,omitempty" plist:"ShowHardDrivesOnDesktop,omitempty"`
	// If `false`, the system doesn't show mounted file servers on the Desktop.
	ShowMountedServersOnDesktop *bool `json:"ShowMountedServersOnDesktop,omitempty" plist:"ShowMountedServersOnDesktop,omitempty"`
	// If `false`, the system doesn't show removable media items on the Desktop.
	ShowRemovableMediaOnDesktop *bool `json:"ShowRemovableMediaOnDesktop,omitempty" plist:"ShowRemovableMediaOnDesktop,omitempty"`
	// If `false`, the system doesn't warn the user before emptying the trash.
	WarnOnEmptyTrash *bool `json:"WarnOnEmptyTrash,omitempty" plist:"WarnOnEmptyTrash,omitempty"`
}

func (p *Finder) PayloadType() string {
	return "com.apple.finder"
}

// Specifies whether Finder should operate in Simple or Full mode.
type InterfaceLevel string

const (
	InterfaceLevelSimple InterfaceLevel = "Simple"
	InterfaceLevelFull   InterfaceLevel = "Full"
)

// The payload that configures the first wired, active Ethernet interface.
type Net8021XFirstActiveEthernet struct {
	*CommonPayloadKeys
	Properties map[string]any
}

func (p *Net8021XFirstActiveEthernet) PayloadType() string {
	return "com.apple.firstactiveethernet.managed"
}

// The payload that configures the first wired Ethernet interface.
type Net8021XFirstEthernet struct {
	*CommonPayloadKeys
	Properties map[string]any
}

func (p *Net8021XFirstEthernet) PayloadType() string {
	return "com.apple.firstethernet.managed"
}

// The payload that configures fonts.
// Each payload may contain one font file. Font files may be in TrueType (.ttf) or OpenType (.otf) file format. Collection types (.ttc or .otc) formats are not supported.
// Fonts are uniquely identified internally by their embedded PostScript name. Two fonts with the same PostScript name will be considered the same font, even if their contents differ. Installing two different fonts with the same PostScript name is not supported, and it is undefined which font will remain installed.
// Supported on the Shared iPad user channel as of iPadOS 18.0. Earlier versions of iPadOS erroneously accepted the Font payload on the device channel but installed it for the currently logged in user.
type Font struct {
	*CommonPayloadKeys
	// The user-visible name for the font. This field is replaced by the actual name of the font after installation. Each payload must contain exactly one font file in trueType (.ttf) or OpenType (.otf) format. Collection formats (.ttc or .otc) are not supported.
	// Fonts are identified by their embedded PostScript names. Two fonts with the same PostScript name are considered to be the same font even if their contents differ. Installing two different fonts with the same PostScript name isn't supported, and the resulting behavior is undefined.
	Name *string `default:"" json:"Name,omitempty" plist:"Name,omitempty"`
	// The contents of the font file.
	Font []byte `json:"Font" plist:"Font" required:"true"`
}

func (p *Font) PayloadType() string {
	return "com.apple.font"
}

// The payload that configures Game Center parental controls.
// Parental controls Game Center restrictions.
type ParentalControlsGameCenter struct {
	*CommonPayloadKeys
	// If `true`, enables Game Center.
	GKFeatureGameCenterAllowed *bool `json:"GKFeatureGameCenterAllowed,omitempty" plist:"GKFeatureGameCenterAllowed,omitempty"`
	// If `true`, allows account modifications.
	GKFeatureAccountModificationAllowed *bool `json:"GKFeatureAccountModificationAllowed,omitempty" plist:"GKFeatureAccountModificationAllowed,omitempty"`
	// If `true`, allows adding Game Center friends.
	GKFeatureAddingGameCenterFriendsAllowed *bool `json:"GKFeatureAddingGameCenterFriendsAllowed,omitempty" plist:"GKFeatureAddingGameCenterFriendsAllowed,omitempty"`
	// If `true`, allows multiplayer gaming.
	GKFeatureMultiplayerGamingAllowed *bool `json:"GKFeatureMultiplayerGamingAllowed,omitempty" plist:"GKFeatureMultiplayerGamingAllowed,omitempty"`
}

func (p *ParentalControlsGameCenter) PayloadType() string {
	return "com.apple.gamed"
}

// The payload that configures the default fallback global Ethernet interface.
type Net8021XGlobalEthernet struct {
	*CommonPayloadKeys
	Properties map[string]any
}

func (p *Net8021XGlobalEthernet) PayloadType() string {
	return "com.apple.globalethernet.managed"
}

// The payload that configures a Google account.
// A Google account payload sets up a Google email address as well as any other Google services the user enables after authentication. Google accounts must be installed via MDM or by Apple Configurator 2 (if the device is supervised). The payload never contains credentials and the user will be prompted to enter their credentials shortly after the payload successfully installs. On Shared iPads, this payload can only be installed on the MDM user channel.
type GoogleAccount struct {
	*CommonPayloadKeys
	// A user-visible description of the Google account, shown in the Mail and Settings apps.
	AccountDescription *string `json:"AccountDescription,omitempty" plist:"AccountDescription,omitempty"`
	// The user's full name for the Google account. This name appears in sent messages.
	AccountName *string `json:"AccountName,omitempty" plist:"AccountName,omitempty"`
	// The full Google email address for the account.
	EmailAddress string `json:"EmailAddress" plist:"EmailAddress" required:"true"`
	// The communication service handler rules for this account.
	CommunicationServiceRules *GoogleAccountCommunicationServiceRules `json:"CommunicationServiceRules,omitempty" plist:"CommunicationServiceRules,omitempty"`
	// The VPNUUID of the per-app VPN the account uses for network communication. Available in iOS 14 and later.
	VPNUUID *string `json:"VPNUUID,omitempty" plist:"VPNUUID,omitempty"`
}

func (p *GoogleAccount) PayloadType() string {
	return "com.apple.google-oauth"
}

// The communication service handler rules for this account.
type GoogleAccountCommunicationServiceRules struct {
	// A dictionary that defines which app to use for audio calls from this account.
	DefaultServiceHandlers *GoogleAccountCommunicationServiceRulesDefaultServiceHandlers `json:"DefaultServiceHandlers,omitempty" plist:"DefaultServiceHandlers,omitempty"`
}

// A dictionary that defines which app to use for audio calls from this account.
type GoogleAccountCommunicationServiceRulesDefaultServiceHandlers struct {
	// The bundle identifier for the default application that handles audio calls to contacts from this account.
	AudioCall *string `json:"AudioCall,omitempty" plist:"AudioCall,omitempty"`
}

// The payload that configures the Home Screen layout.
// The payload defines a layout of apps, folders, & web clips for the Home screen.
type HomeScreenLayout struct {
	*CommonPayloadKeys
	// An array of dictionaries, each of which must conform to the icon dictionary format. If this key isn't present, the user's Dock is empty.
	Dock *[]*IconItem `json:"Dock,omitempty" plist:"Dock,omitempty"`
	// An array of arrays of dictionaries, each of which must conform to the icon dictionary format.
	Pages [][]*IconItem `json:"Pages" plist:"Pages" required:"true"`
}

func (p *HomeScreenLayout) PayloadType() string {
	return "com.apple.homescreenlayout"
}

// An array of dictionaries that conform to the icon dictionary format.
type IconItem struct {
	// The type of the Dock item.
	Type IconItemType `json:"Type" plist:"Type" required:"true"`
	// The human-readable string shown to the user. This setting is valid only if the type is `Folder`.
	DisplayName *string `json:"DisplayName,omitempty" plist:"DisplayName,omitempty"`
	// The bundle identifier of the app. This setting is required if the type is `Application`.
	BundleID *string `json:"BundleID,omitempty" plist:"BundleID,omitempty"`
	// An array of arrays of dictionaries, each conforming to the icon dictionary format. This setting is valid only if the type is `Folder`.
	Pages *[][]*IconItem `json:"Pages,omitempty" plist:"Pages,omitempty"`
	// The URL of the existing web clip for this item. This setting is required if `type` is `WebClip`. If more than one web clip exists with the same URL, the behavior is undefined.
	// Specifying a web clip in this payload doesn't create the web clip. Use the `WebClip` payload to create a web clip.
	URL *string `json:"URL,omitempty" plist:"URL,omitempty"`
}

// The type of the Dock item.
type IconItemType string

const (
	IconItemTypeApplication IconItemType = "Application"
	IconItemTypeFolder      IconItemType = "Folder"
	IconItemTypeWebClip     IconItemType = "WebClip"
)

// The payload that configures parental control for dictation and profanity.
type ParentalControlDictationandProfanity struct {
	*CommonPayloadKeys
	// If `false`, suppresses profanity. Use `forceAssistantProfanityFilter` in Restrictions instead.
	ProfanityAllowed *bool `json:"Profanity Allowed,omitempty" plist:"Profanity Allowed,omitempty"`
	// If `false`, disables dictation. Use `allowDictation` in Restrictions instead.
	IronwoodAllowed *bool `json:"Ironwood Allowed,omitempty" plist:"Ironwood Allowed,omitempty"`
}

func (p *ParentalControlDictationandProfanity) PayloadType() string {
	return "com.apple.ironwood.support"
}

// The payload that configures a Jabber account.
// A Jabber payload creates a Jabber account on the device.
type JabberAccount struct {
	*CommonPayloadKeys
	// The description of the account.
	JabberAccountDescription *string `json:"JabberAccountDescription,omitempty" plist:"JabberAccountDescription,omitempty"`
	// The server's address.
	JabberHostName string `json:"JabberHostName" plist:"JabberHostName" required:"true"`
	// The user's user name.
	JabberUserName *string `json:"JabberUserName,omitempty" plist:"JabberUserName,omitempty"`
	// The user's password.
	JabberPassword *string `json:"JabberPassword,omitempty" plist:"JabberPassword,omitempty"`
	// If `true`, enables SSL.
	JabberUseSSL *bool `json:"JabberUseSSL,omitempty" plist:"JabberUseSSL,omitempty"`
	// The server's port.
	JabberPort *int64 `default:"5222" json:"JabberPort,omitempty" plist:"JabberPort,omitempty"`
	// The authentication method for the account.
	JabberAuthentication JabberAuthentication `json:"JabberAuthentication" plist:"JabberAuthentication" required:"true"`
}

func (p *JabberAccount) PayloadType() string {
	return "com.apple.jabber.account"
}

// The authentication method for the account.
type JabberAuthentication string

const (
	JabberAuthenticationJabberAuthPassword JabberAuthentication = "JabberAuthPassword"
)

// The payload that configures a Lightweight Directory Access Protocol (LDAP) account.
type LDAP struct {
	*CommonPayloadKeys
	// The description of the account.
	LDAPAccountDescription *string `json:"LDAPAccountDescription,omitempty" plist:"LDAPAccountDescription,omitempty"`
	// The server's address.
	LDAPAccountHostName string `json:"LDAPAccountHostName" plist:"LDAPAccountHostName" required:"true"`
	// The user's user name.
	LDAPAccountUserName *string `json:"LDAPAccountUserName,omitempty" plist:"LDAPAccountUserName,omitempty"`
	// The user's password. Only use this in encrypted profiles.
	LDAPAccountPassword *string `json:"LDAPAccountPassword,omitempty" plist:"LDAPAccountPassword,omitempty"`
	// If `true`, the system enables SSL.
	LDAPAccountUseSSL *bool `json:"LDAPAccountUseSSL,omitempty" plist:"LDAPAccountUseSSL,omitempty"`
	// An array of search settings dictionaries.
	LDAPSearchSettings *[]*LDAPSearchSettingsItem `json:"LDAPSearchSettings,omitempty" plist:"LDAPSearchSettings,omitempty"`
	// The VPNUUID of the per-app VPN the account uses for network communication. Available in iOS 14 and later.
	VPNUUID *string `json:"VPNUUID,omitempty" plist:"VPNUUID,omitempty"`
}

func (p *LDAP) PayloadType() string {
	return "com.apple.ldap.account"
}

type LDAPSearchSettingsItem struct {
	// The description of this search setting.
	LDAPSearchSettingDescription *string `json:"LDAPSearchSettingDescription,omitempty" plist:"LDAPSearchSettingDescription,omitempty"`
	// The path to the node where a search should start.
	LDAPSearchSettingSearchBase string `json:"LDAPSearchSettingSearchBase" plist:"LDAPSearchSettingSearchBase" required:"true"`
	// The type of recursion to use in the search:
	// - `LDAPSearchSettingScopeBase`: The search uses only the immediate node that the search base points to.
	// - `LDAPSearchSettingScopeOneLevel`: The search uses the node plus its immediate children.
	// - `LDAPSearchSettingScopeSubtree`: The search uses the node plus all children, regardless of depth.
	LDAPSearchSettingScope *LDAPSearchSettingScope `default:"LDAPSearchSettingScopeSubtree" json:"LDAPSearchSettingScope,omitempty" plist:"LDAPSearchSettingScope,omitempty"`
}

// The type of recursion to use in the search:
// - `LDAPSearchSettingScopeBase`: The search uses only the immediate node that the search base points to.
// - `LDAPSearchSettingScopeOneLevel`: The search uses the node plus its immediate children.
// - `LDAPSearchSettingScopeSubtree`: The search uses the node plus all children, regardless of depth.
type LDAPSearchSettingScope string

const (
	LDAPSearchSettingScopeLDAPSearchSettingScopeBase     LDAPSearchSettingScope = "LDAPSearchSettingScopeBase"
	LDAPSearchSettingScopeLDAPSearchSettingScopeOneLevel LDAPSearchSettingScope = "LDAPSearchSettingScopeOneLevel"
	LDAPSearchSettingScopeLDAPSearchSettingScopeSubtree  LDAPSearchSettingScope = "LDAPSearchSettingScopeSubtree"
)

// The payload that configures a device's login items.
// This payload handles login items usage on macOS.
type LoginItemsManagedItems struct {
	*CommonPayloadKeys
	// An array of login item dictionaries.
	AutoLaunchedApplicationDictionaryManaged []LoginItem `json:"AutoLaunchedApplicationDictionary-managed" plist:"AutoLaunchedApplicationDictionary-managed" required:"true"`
}

func (p *LoginItemsManagedItems) PayloadType() string {
	return "com.apple.loginitems.managed"
}

// A login item.
type LoginItem struct {
	// The URL or path string to the item's location.
	Path string `json:"Path" plist:"Path" required:"true"`
	// If `true`, the system hides this item in the Users & Groups login items list.
	Hide *bool `json:"Hide,omitempty" plist:"Hide,omitempty"`
}

// The payload that configures Login Window behavior.
// The com.apple.loginwindow payload creates managed preferences on macOS for system/device profiles.
type LoginWindow struct {
	*CommonPayloadKeys
	// If `true`, the system shows the name and password dialog. If `false`, the system displays a list of users.
	SHOWFULLNAME *bool `json:"SHOWFULLNAME,omitempty" plist:"SHOWFULLNAME,omitempty"`
	// If `true`, the system shows only network and system users when showing a user list.
	HideLocalUsers *bool `json:"HideLocalUsers,omitempty" plist:"HideLocalUsers,omitempty"`
	// If `true`, the system shows network users when showing a user list.
	IncludeNetworkUser *bool `json:"IncludeNetworkUser,omitempty" plist:"IncludeNetworkUser,omitempty"`
	// If `true`, the system hides administrator users when showing a user list.
	HideAdminUsers *bool `json:"HideAdminUsers,omitempty" plist:"HideAdminUsers,omitempty"`
	// If `true`, the system displays "Other..." when it shows a list of users.
	SHOWOTHERUSERSMANAGED *bool `json:"SHOWOTHERUSERS_MANAGED,omitempty" plist:"SHOWOTHERUSERS_MANAGED,omitempty"`
	// The admin host info. If present in the payload, the system displays its value in the Login Window as additional computer information. Before macOS 10.10, this string could only contain host name, system version, or IP address. After macOS 10.10, setting this key to any value allows the user to click the time area of the menu bar to toggle through various computer information values.
	AdminHostInfo *AdminHostInfo `json:"AdminHostInfo,omitempty" plist:"AdminHostInfo,omitempty"`
	// The list of user GUIDs or group GUIDs of users that the system allows to log in. An asterisk (`*`) string specifies all users or groups. This only applies to network accounts and mobile accounts.
	AllowList *[]string `json:"AllowList,omitempty" plist:"AllowList,omitempty"`
	// The list of user GUIDs or group GUIDs of users that the system disallows to log in. This list takes priority over the list in the `AllowList` key. This only applies to network accounts and mobile accounts.
	DenyList *[]string `json:"DenyList,omitempty" plist:"DenyList,omitempty"`
	// If `true`, the system hides mobile account users in a user list. In some cases, mobile users show up as network users.
	HideMobileAccounts *bool `json:"HideMobileAccounts,omitempty" plist:"HideMobileAccounts,omitempty"`
	// If `true`, the system disables the Shut Down button.
	ShutDownDisabled *bool `json:"ShutDownDisabled,omitempty" plist:"ShutDownDisabled,omitempty"`
	// If `true`, the system disables the Restart item.
	RestartDisabled *bool `json:"RestartDisabled,omitempty" plist:"RestartDisabled,omitempty"`
	// If `true`, the system disables the Sleep button.
	SleepDisabled *bool `json:"SleepDisabled,omitempty" plist:"SleepDisabled,omitempty"`
	// If `true`, the system disregards the `>console` special user name, which provides a command line UI.
	DisableConsoleAccess *bool `json:"DisableConsoleAccess,omitempty" plist:"DisableConsoleAccess,omitempty"`
	// The text to display in the Login Window.
	LoginwindowText *string `json:"LoginwindowText,omitempty" plist:"LoginwindowText,omitempty"`
	// If `true`, the system disables the Shut Down menu item when the user is logged in.
	ShutDownDisabledWhileLoggedIn *bool `json:"ShutDownDisabledWhileLoggedIn,omitempty" plist:"ShutDownDisabledWhileLoggedIn,omitempty"`
	// If `true`, the system disables the Restart menu item when the user is logged in.
	RestartDisabledWhileLoggedIn *bool `json:"RestartDisabledWhileLoggedIn,omitempty" plist:"RestartDisabledWhileLoggedIn,omitempty"`
	// If `true`, the system disables the Power Off menu item when the user is logged in.
	PowerOffDisabledWhileLoggedIn *bool `json:"PowerOffDisabledWhileLoggedIn,omitempty" plist:"PowerOffDisabledWhileLoggedIn,omitempty"`
	// If `true`, the system disables the Log Out menu item when the user is logged in. Available in macOS 10.13 and later.
	LogOutDisabledWhileLoggedIn *bool `json:"LogOutDisabledWhileLoggedIn,omitempty" plist:"LogOutDisabledWhileLoggedIn,omitempty"`
	// If `true`, the system disables the immediate Screen Lock functions. Available in macOS 10.13 and later.
	DisableScreenLockImmediate *bool `json:"DisableScreenLockImmediate,omitempty" plist:"DisableScreenLockImmediate,omitempty"`
	// If `true`, the system shows the Input Menu in the Login Window.
	ShowInputMenu *bool `json:"showInputMenu,omitempty" plist:"showInputMenu,omitempty"`
	// If `true`, the system disables the automatic login option when using FileVault.
	DisableFDEAutoLogin *bool `json:"DisableFDEAutoLogin,omitempty" plist:"DisableFDEAutoLogin,omitempty"`
	// The user short name for an existing user to set up auto login.
	AutologinUsername *string `json:"AutologinUsername,omitempty" plist:"AutologinUsername,omitempty"`
	// An optional user password to set up auto login. This must match the `AutologinUsername` user's current password.
	AutologinPassword *string `json:"AutologinPassword,omitempty" plist:"AutologinPassword,omitempty"`
}

func (p *LoginWindow) PayloadType() string {
	return "com.apple.loginwindow"
}

// The admin host info. If present in the payload, the system displays its value in the Login Window as additional computer information. Before macOS 10.10, this string could only contain host name, system version, or IP address. After macOS 10.10, setting this key to any value allows the user to click the time area of the menu bar to toggle through various computer information values.
type AdminHostInfo string

const (
	AdminHostInfoHostName      AdminHostInfo = "HostName"
	AdminHostInfoSystemVersion AdminHostInfo = "SystemVersion"
	AdminHostInfoIPAddress     AdminHostInfo = "IPAddress"
)

// The payload that configures lights-out management (LOM) settings.
// Configures a computer to send or receive "PowerON". "PowerOFF", "Reset" requests.
type LightsOutManagementLOM struct {
	*CommonPayloadKeys
	// The UUID certificate for the device. This key indicates the device can receive `PowerON`, `PowerOFF`, and `Reset` requests from a LOM controller. This certificate must contain the Key Usage attributes of Digital Signature, Key Encipherment and Data Encipherment. As well as the Extended Key Usage attributes of Server Authentication and Client Authentication.
	DeviceCertificateUUID *string `json:"DeviceCertificateUUID,omitempty" plist:"DeviceCertificateUUID,omitempty"`
	// The UUID certificate for the LOM controller. This key configures the device to accept the `LOMDeviceRequestCommand` from MDM and then send it to the target device.
	ControllerCertificateUUID *string `json:"ControllerCertificateUUID,omitempty" plist:"ControllerCertificateUUID,omitempty"`
	// An array of payload UUIDs containing CA certificates that controllers use to evaluate trust of device certificates.
	DeviceCACertificateUUIDs *[]string `json:"DeviceCACertificateUUIDs,omitempty" plist:"DeviceCACertificateUUIDs,omitempty"`
	// An array of payload UUIDs containing CA certificates that devices use to evaluate trust of controller certificates.
	// This key configures the device to accept the `LOMDeviceRequestCommand` from MDM and then send it to the target device. This certificate must contain the Key Usage attributes of Digital Signature, Key Encipherment and Data Encipherment. As well as the Extended Key Usage attributes of Server Authentication and Client Authentication.
	ControllerCACertificateUUIDs *[]string `json:"ControllerCACertificateUUIDs,omitempty" plist:"ControllerCACertificateUUIDs,omitempty"`
}

func (p *LightsOutManagementLOM) PayloadType() string {
	return "com.apple.lom"
}

// The payload that configures a Mail account.
// An email payload creates an email account on the device.
type Mail struct {
	*CommonPayloadKeys
	// A user-visible description of the email account, shown in the Mail and Settings applications.
	EmailAccountDescription *string `json:"EmailAccountDescription,omitempty" plist:"EmailAccountDescription,omitempty"`
	// The full user name for the account. The system displays this name in sent messages.
	EmailAccountName *string `json:"EmailAccountName,omitempty" plist:"EmailAccountName,omitempty"`
	// Defines the protocol to use for the account.
	EmailAccountType EmailAccountType `json:"EmailAccountType" plist:"EmailAccountType" required:"true"`
	// The full email address for the account. If this string isn't present in the payload, the device prompts the user for this string during interactive profile installation in Settings or System Preferences.
	EmailAddress *string `json:"EmailAddress,omitempty" plist:"EmailAddress,omitempty"`
	// The authentication scheme for incoming mail.
	IncomingMailServerAuthentication IncomingMailServerAuthentication `json:"IncomingMailServerAuthentication" plist:"IncomingMailServerAuthentication" required:"true"`
	// The incoming mail server host name.
	IncomingMailServerHostName string `json:"IncomingMailServerHostName" plist:"IncomingMailServerHostName" required:"true"`
	// The incoming mail server port number. If not set, the system uses the default port for a given protocol.
	IncomingMailServerPortNumber *int64 `json:"IncomingMailServerPortNumber,omitempty" plist:"IncomingMailServerPortNumber,omitempty"`
	// If `true`, the system enables SSL for authentication on the incoming mail server.
	IncomingMailServerUseSSL *bool `json:"IncomingMailServerUseSSL,omitempty" plist:"IncomingMailServerUseSSL,omitempty"`
	// The user name for the email account, usually the same as the email address up to the "@" character. If not set and the account requires authentication for incoming email, the device prompts the user for this string during interactive profile installation in Settings or System Preferences.
	IncomingMailServerUsername *string `json:"IncomingMailServerUsername,omitempty" plist:"IncomingMailServerUsername,omitempty"`
	// The password for the incoming mail server. Only use this in encrypted profiles.
	IncomingPassword *string `json:"IncomingPassword,omitempty" plist:"IncomingPassword,omitempty"`
	// The password for the outgoing mail server. Only use this in encrypted profiles.
	OutgoingPassword *string `json:"OutgoingPassword,omitempty" plist:"OutgoingPassword,omitempty"`
	// If `true`, the system prompts the user only once for the password, which it uses for both outgoing and incoming mail.
	// This setting is only supported by interactive profile installations. Not supported by non-interactive installations, such as MDM on iOS.
	OutgoingPasswordSameAsIncomingPassword *bool `json:"OutgoingPasswordSameAsIncomingPassword,omitempty" plist:"OutgoingPasswordSameAsIncomingPassword,omitempty"`
	// The authentication scheme for outgoing mail.
	OutgoingMailServerAuthentication OutgoingMailServerAuthentication `json:"OutgoingMailServerAuthentication" plist:"OutgoingMailServerAuthentication" required:"true"`
	// The outgoing mail server host name.
	OutgoingMailServerHostName string `json:"OutgoingMailServerHostName" plist:"OutgoingMailServerHostName" required:"true"`
	// The outgoing mail server port number. If not set, the system uses ports 25, 587, and 465, in that order.
	OutgoingMailServerPortNumber *int64 `json:"OutgoingMailServerPortNumber,omitempty" plist:"OutgoingMailServerPortNumber,omitempty"`
	// If `true`, the system enables SSL authentication on the outgoing mail server.
	OutgoingMailServerUseSSL *bool `json:"OutgoingMailServerUseSSL,omitempty" plist:"OutgoingMailServerUseSSL,omitempty"`
	// The user name for the email account, usually the same as the email address up to the "@" character. If not set and the account requires authentication for outgoing email, the device prompts the user for this string during interactive profile installation in Settings or System Preferences.
	OutgoingMailServerUsername *string `json:"OutgoingMailServerUsername,omitempty" plist:"OutgoingMailServerUsername,omitempty"`
	// If `true`, the system prevents moving messages out of this email account and into another account. It also prevents forwarding or replying from an account other than the recipient of the message.
	PreventMove *bool `json:"PreventMove,omitempty" plist:"PreventMove,omitempty"`
	// If `true`, the system prevents this account from sending mail in any app other than the Apple Mail app.
	PreventAppSheet *bool `json:"PreventAppSheet,omitempty" plist:"PreventAppSheet,omitempty"`
	// If `true`, the system enables S/MIME encryption. The system ignores this key in iOS 10.0 and later.
	SMIMEEnabled *bool `json:"SMIMEEnabled,omitempty" plist:"SMIMEEnabled,omitempty"`
	// If `true`, the system enables S/MIME signing for this account.
	SMIMESigningEnabled *bool `json:"SMIMESigningEnabled,omitempty" plist:"SMIMESigningEnabled,omitempty"`
	// The payload UUID of the identity certificate used to sign messages sent from this account.
	SMIMESigningCertificateUUID *string `json:"SMIMESigningCertificateUUID,omitempty" plist:"SMIMESigningCertificateUUID,omitempty"`
	// If `true`, the system enables S/MIME encryption for this account.
	SMIMEEncryptionEnabled *bool `json:"SMIMEEncryptionEnabled,omitempty" plist:"SMIMEEncryptionEnabled,omitempty"`
	// The UUID of the identity certificate used to decrypt messages sent to this account. The system attaches the public certificate to outgoing mail to allow the user to receive encrypted mail. When the user sends encrypted mail, the system uses the public certificate to encrypt the copy of the mail in their Sent mailbox.
	SMIMEEncryptionCertificateUUID *string `json:"SMIMEEncryptionCertificateUUID,omitempty" plist:"SMIMEEncryptionCertificateUUID,omitempty"`
	// If `true`, the system displays the per-message encryption switch in the Mail Compose UI. Deprecated in iOS 12.0. Use `SMIMEEnableEncryptionPerMessageSwitch` instead.
	SMIMEEnablePerMessageSwitch *bool `json:"SMIMEEnablePerMessageSwitch,omitempty" plist:"SMIMEEnablePerMessageSwitch,omitempty"`
	// If `true`, the system excludes this account from Recent Addresses syncing.
	DisableMailRecentsSyncing *bool `json:"disableMailRecentsSyncing,omitempty" plist:"disableMailRecentsSyncing,omitempty"`
	// If `true`, the system enables this account to use Mail Drop.
	AllowMailDrop *bool `json:"allowMailDrop,omitempty" plist:"allowMailDrop,omitempty"`
	// The path prefix for the IMAP mail server.
	IncomingMailServerIMAPPathPrefix *string `json:"IncomingMailServerIMAPPathPrefix,omitempty" plist:"IncomingMailServerIMAPPathPrefix,omitempty"`
	// If `true`, the user can turn S/MIME signing on or off in Settings.
	SMIMESigningUserOverrideable *bool `json:"SMIMESigningUserOverrideable,omitempty" plist:"SMIMESigningUserOverrideable,omitempty"`
	// If `true`, the user can select the signing identity.
	SMIMESigningCertificateUUIDUserOverrideable *bool `json:"SMIMESigningCertificateUUIDUserOverrideable,omitempty" plist:"SMIMESigningCertificateUUIDUserOverrideable,omitempty"`
	// If `true`, the system enables S/MIME encryption by default.
	SMIMEEncryptByDefault *bool `json:"SMIMEEncryptByDefault,omitempty" plist:"SMIMEEncryptByDefault,omitempty"`
	// If `true`, the user can turn encryption by default on/off, and encryption is on.
	SMIMEEncryptByDefaultUserOverrideable *bool `json:"SMIMEEncryptByDefaultUserOverrideable,omitempty" plist:"SMIMEEncryptByDefaultUserOverrideable,omitempty"`
	// If `true`, the user can select the S/MIME encryption identity, and encryption is on.
	SMIMEEncryptionCertificateUUIDUserOverrideable *bool `json:"SMIMEEncryptionCertificateUUIDUserOverrideable,omitempty" plist:"SMIMEEncryptionCertificateUUIDUserOverrideable,omitempty"`
	// If `true`, the system displays the per-message encryption switch in the Mail Compose UI.
	SMIMEEnableEncryptionPerMessageSwitch *bool `json:"SMIMEEnableEncryptionPerMessageSwitch,omitempty" plist:"SMIMEEnableEncryptionPerMessageSwitch,omitempty"`
	// The VPNUUID of the per-app VPN the account uses for network communication. Available in iOS 14 and later.
	VPNUUID *string `json:"VPNUUID,omitempty" plist:"VPNUUID,omitempty"`
}

func (p *Mail) PayloadType() string {
	return "com.apple.mail.managed"
}

// Defines the protocol to use for the account.
type EmailAccountType string

const (
	EmailAccountTypeEmailTypeIMAP EmailAccountType = "EmailTypeIMAP"
	EmailAccountTypeEmailTypePOP  EmailAccountType = "EmailTypePOP"
)

// The authentication scheme for incoming mail.
type IncomingMailServerAuthentication string

const (
	IncomingMailServerAuthenticationEmailAuthNone     IncomingMailServerAuthentication = "EmailAuthNone"
	IncomingMailServerAuthenticationEmailAuthPassword IncomingMailServerAuthentication = "EmailAuthPassword"
	IncomingMailServerAuthenticationEmailAuthCRAMMD5  IncomingMailServerAuthentication = "EmailAuthCRAMMD5"
	IncomingMailServerAuthenticationEmailAuthNTLM     IncomingMailServerAuthentication = "EmailAuthNTLM"
	IncomingMailServerAuthenticationEmailAuthHTTPMD5  IncomingMailServerAuthentication = "EmailAuthHTTPMD5"
)

// The authentication scheme for outgoing mail.
type OutgoingMailServerAuthentication string

const (
	OutgoingMailServerAuthenticationEmailAuthNone     OutgoingMailServerAuthentication = "EmailAuthNone"
	OutgoingMailServerAuthenticationEmailAuthPassword OutgoingMailServerAuthentication = "EmailAuthPassword"
	OutgoingMailServerAuthenticationEmailAuthCRAMMD5  OutgoingMailServerAuthentication = "EmailAuthCRAMMD5"
	OutgoingMailServerAuthenticationEmailAuthNTLM     OutgoingMailServerAuthentication = "EmailAuthNTLM"
	OutgoingMailServerAuthenticationEmailAuthHTTPMD5  OutgoingMailServerAuthentication = "EmailAuthHTTPMD5"
)

// The payload that configures menu extras.
// Specified menu extras will be added or removed from the menu bar
// after user login. Standard menu extra may be specified by file
// name. Non-standard menu extras are specified by full path.
type ManagedMenuExtras struct {
	*CommonPayloadKeys
	// The number of seconds to delay after login before adding or removing menu extras. If the delay is too short, the menu extras don't appear, or disappear from the menu bar.
	DelaySeconds *float64 `default:"2.5" json:"delaySeconds,omitempty" plist:"delaySeconds,omitempty"`
	// The maximum wait, in seconds, for all menu extras to be added or removed.
	MaxWaitSeconds *float64 `default:"20" json:"maxWaitSeconds,omitempty" plist:"maxWaitSeconds,omitempty"`
	// If `true`, enables the AirPort menu extra.
	AirPortmenu *bool `json:"AirPort.menu,omitempty" plist:"AirPort.menu,omitempty"`
	// If `true`, enables the Battery menu extra.
	Batterymenu *bool `json:"Battery.menu,omitempty" plist:"Battery.menu,omitempty"`
	// If `true`, enables the Bluetooth menu extra.
	Bluetoothmenu *bool `json:"Bluetooth.menu,omitempty" plist:"Bluetooth.menu,omitempty"`
	// If `true`, enables the CPU menu extra.
	CPUmenu *bool `json:"CPU.menu,omitempty" plist:"CPU.menu,omitempty"`
	// If `true`, enables the Clock menu extra.
	Clockmenu *bool `json:"Clock.menu,omitempty" plist:"Clock.menu,omitempty"`
	// If `true`, enables the Displays menu extra.
	Displaysmenu *bool `json:"Displays.menu,omitempty" plist:"Displays.menu,omitempty"`
	// If `true`, enables the Eject menu extra.
	Ejectmenu *bool `json:"Eject.menu,omitempty" plist:"Eject.menu,omitempty"`
	// If `true`, enables the Fax menu extra.
	Faxmenu *bool `json:"Fax.menu,omitempty" plist:"Fax.menu,omitempty"`
	// If `true`, enables the HomeSync menu extra.
	HomeSyncmenu *bool `json:"HomeSync.menu,omitempty" plist:"HomeSync.menu,omitempty"`
	// If `true`, enables the iChat menu extra.
	IChatmenu *bool `json:"iChat.menu,omitempty" plist:"iChat.menu,omitempty"`
	// If `true`, enables the Ink menu extra.
	Inkmenu *bool `json:"Ink.menu,omitempty" plist:"Ink.menu,omitempty"`
	// If `true`, enables the IrDA menu extra.
	IrDAmenu *bool `json:"IrDA.menu,omitempty" plist:"IrDA.menu,omitempty"`
	// If `true`, enables the PCCard menu extra.
	PCCardmenu *bool `json:"PCCard.menu,omitempty" plist:"PCCard.menu,omitempty"`
	// If `true`, enables the PPP menu extra.
	PPPmenu *bool `json:"PPP.menu,omitempty" plist:"PPP.menu,omitempty"`
	// If `true`, enables the PPPoE menu extra.
	PPPoEmenu *bool `json:"PPPoE.menu,omitempty" plist:"PPPoE.menu,omitempty"`
	// If `true`, enables the Remote Desktop menu extra.
	RemoteDesktopmenu *bool `json:"RemoteDesktop.menu,omitempty" plist:"RemoteDesktop.menu,omitempty"`
	// If `true`, enables the Script menu extra.
	ScriptMenumenu *bool `json:"Script Menu.menu,omitempty" plist:"Script Menu.menu,omitempty"`
	// If `true`, enables the Spaces menu extra.
	Spacesmenu *bool `json:"Spaces.menu,omitempty" plist:"Spaces.menu,omitempty"`
	// If `true`, enables the Sync menu extra.
	Syncmenu *bool `json:"Sync.menu,omitempty" plist:"Sync.menu,omitempty"`
	// If `true`, enables the Text Input menu extra.
	TextInputmenu *bool `json:"TextInput.menu,omitempty" plist:"TextInput.menu,omitempty"`
	// If `true`, enables the TimeMachine menu extra.
	TimeMachinemenu *bool `json:"TimeMachine.menu,omitempty" plist:"TimeMachine.menu,omitempty"`
	// If `true`, enables the Universal Access menu extra.
	UniversalAccessmenu *bool `json:"UniversalAccess.menu,omitempty" plist:"UniversalAccess.menu,omitempty"`
	// If `true`, enables the User menu extra.
	Usermenu *bool `json:"User.menu,omitempty" plist:"User.menu,omitempty"`
	// If `true`, enables the VPN menu extra.
	VPNmenu *bool `json:"VPN.menu,omitempty" plist:"VPN.menu,omitempty"`
	// If `true`, enables the Volume menu extra.
	Volumemenu *bool `json:"Volume.menu,omitempty" plist:"Volume.menu,omitempty"`
	// If `true`, enables the WWAN menu extra.
	WWANmenu *bool `json:"WWAN.menu,omitempty" plist:"WWAN.menu,omitempty"`
}

func (p *ManagedMenuExtras) PayloadType() string {
	return "com.apple.mcxMenuExtras"
}

// The payload that configures scripts to run at login and logout.
// Login and logout managed script handling
type LoginWindowScripts struct {
	*CommonPayloadKeys
	// An array of one or more dictionaries of scripts to run at user login time.
	Loginscripts *[]*ScriptsItems `json:"loginscripts,omitempty" plist:"loginscripts,omitempty"`
	// An array of one or more dictionaries of scripts to run at user logout time.
	Logoutscripts *[]*ScriptsItems `json:"logoutscripts,omitempty" plist:"logoutscripts,omitempty"`
	// If `true`, the system doesn't execute the login scripts during login.
	SkipLoginHook *bool `json:"skipLoginHook,omitempty" plist:"skipLoginHook,omitempty"`
	// If `true`, the system doesn't execute the logout scripts during logout.
	SkipLogoutHook *bool `json:"skipLogoutHook,omitempty" plist:"skipLogoutHook,omitempty"`
}

func (p *LoginWindowScripts) PayloadType() string {
	return "com.apple.mcxloginscripts"
}

// A dictionary of login scripts.
type ScriptsItems struct {
	// The filename for display purposes.
	Filename string `json:"filename" plist:"filename" required:"true"`
	// The UTF-8 encoded data object representing the executable script.
	Filedata []byte `json:"filedata" plist:"filedata" required:"true"`
}

// The payload that configures printers.
type Printing struct {
	*CommonPayloadKeys
	// If `true`, requires an administrator password to add printers.
	RequireAdminToAddPrinters *bool `json:"RequireAdminToAddPrinters,omitempty" plist:"RequireAdminToAddPrinters,omitempty"`
	// If `true`, allows printers that connect directly to a user's computer.
	AllowLocalPrinters *bool `json:"AllowLocalPrinters,omitempty" plist:"AllowLocalPrinters,omitempty"`
	// If `true`, requires an administrator password to print locally.
	RequireAdminToPrintLocally *bool `json:"RequireAdminToPrintLocally,omitempty" plist:"RequireAdminToPrintLocally,omitempty"`
	// If `true`, shows only managed printers.
	ShowOnlyManagedPrinters *bool `json:"ShowOnlyManagedPrinters,omitempty" plist:"ShowOnlyManagedPrinters,omitempty"`
	// If `true`, prints the page footer (including the user name and date).
	PrintFooter *bool `json:"PrintFooter,omitempty" plist:"PrintFooter,omitempty"`
	// If `true`, includes the MAC address.
	PrintMACAddress *bool `json:"PrintMACAddress,omitempty" plist:"PrintMACAddress,omitempty"`
	// The footer font size.
	FooterFontSize *string `json:"FooterFontSize,omitempty" plist:"FooterFontSize,omitempty"`
	// The footer font name.
	FooterFontName *string `json:"FooterFontName,omitempty" plist:"FooterFontName,omitempty"`
	// The default printer for the user.
	DefaultPrinter *DefaultPrinter `json:"DefaultPrinter,omitempty" plist:"DefaultPrinter,omitempty"`
	// The printers available to a user.
	UserPrinterList *UserPrinterList `json:"UserPrinterList,omitempty" plist:"UserPrinterList,omitempty"`
}

func (p *Printing) PayloadType() string {
	return "com.apple.mcxprinting"
}

// The default printer for the user.
type DefaultPrinter struct {
	// The device URI.
	DeviceURI *string `json:"DeviceURI,omitempty" plist:"DeviceURI,omitempty"`
	// The display name.
	DisplayName *string `json:"DisplayName,omitempty" plist:"DisplayName,omitempty"`
}

// The printers available to a user.
type UserPrinterList struct {
	// A dictionary of printer details.
	Printer *Printer `json:"Printer,omitempty" plist:"Printer,omitempty"`
}

// A dictionary of printer details.
type Printer struct {
	// The device URI.
	DeviceURI *string `json:"DeviceURI,omitempty" plist:"DeviceURI,omitempty"`
	// The display name.
	DisplayName *string `json:"DisplayName,omitempty" plist:"DisplayName,omitempty"`
	// The printer's location.
	Location *string `json:"Location,omitempty" plist:"Location,omitempty"`
	// The printer's model.
	Model *string `json:"Model,omitempty" plist:"Model,omitempty"`
	// If `true`, locks the printer.
	PrinterLocked *bool `json:"PrinterLocked,omitempty" plist:"PrinterLocked,omitempty"`
	// The printer's PPDURL.
	PPDURL *string `json:"PPDURL,omitempty" plist:"PPDURL,omitempty"`
}

// The payload that configures mobile device management (MDM) settings.
type MDM struct {
	*CommonPayloadKeys
	// The UUID of the certificate payload for the device's identity. It may also point to a SCEP payload.
	IdentityCertificateUUID string `json:"IdentityCertificateUUID" plist:"IdentityCertificateUUID" required:"true"`
	// The topic that MDM listens to for push notifications. The certificate that the server uses to send push notifications must have the same topic in its subject. The topic must begin with the 'com.apple.mgmt.' prefix.
	// > Note:
	// > When updating the payload, the value of this key must not change. Any change is an error, and the update is rejected.
	Topic string `json:"Topic" plist:"Topic" required:"true"`
	// The URL that the device contacts to retrieve device management instructions. The URL must begin with the `https://` URL scheme, and may contain a port number
	// (`:1234`, for example).
	// > Note:
	// > When updating the payload, the value of this key must not change. Any change is an error, and the update is rejected.
	ServerURL string `json:"ServerURL" plist:"ServerURL" required:"true"`
	// The URL that the device should use to check in during installation. The URL must begin with the `https://` URL scheme and may contain a port number (`:1234`, for example). If not set, the system uses `ServerURL`.
	// > Note:
	// > When updating the payload, the value of this key must not change. Any change is an error, and the update is rejected.
	CheckInURL *string `json:"CheckInURL,omitempty" plist:"CheckInURL,omitempty"`
	// If 'true', each message coming from the device carries the additional 'Mdm-Signature' HTTP header.
	SignMessage *bool `json:"SignMessage,omitempty" plist:"SignMessage,omitempty"`
	// Logical OR of the following bit flags:
	// - `1`: Allow inspection of installed configuration profiles.
	// - `2`: Allow installation and removal of configuration profiles.
	// - `4`: Allow device lock and passcode removal.
	// - `8`: Allow device erase.
	// - `16`: Allow query of device information (device capacity, serial number).
	// - `32`: Allow query of network information (phone/SIM numbers, MAC addresses).
	// - `64`: Allow inspection of installed provisioning profiles.
	// - `128`: Allow installation and removal of provisioning profiles.
	// - `256`: Allow inspection of installed applications.
	// - `512`: Allow restriction-related queries.
	// - `1024`: Allow security-related queries.
	// - `2048`: Allow manipulation of settings.
	// - `4096`: Allow app management.
	// Don't set to `0`. Specify `1` if you specify `2`. Specify `64` if you specify `128`. Ignored if you set a value for `ManagedAppleID`.
	// > Note:
	// > When updating the payload, the addition of any access right is an error, and the update is rejected.
	AccessRights *int64 `json:"AccessRights,omitempty" plist:"AccessRights,omitempty"`
	// If 'true', the device uses the development APNS servers. Otherwise, the device uses the production servers.
	// Set to 'false' if your Apple Push Notification Service certificate was issued by the Apple Push Certificate Portal ('https://identity.apple.com/pushcert'). That portal only issues certificates for the production push environment.
	UseDevelopmentAPNS *bool `json:"UseDevelopmentAPNS,omitempty" plist:"UseDevelopmentAPNS,omitempty"`
	// The Managed Apple Account of the user. Previously required for profile-driven user enrollment.
	// Removed as of iOS 18 and macOS 15.
	ManagedAppleID *string `json:"ManagedAppleID,omitempty" plist:"ManagedAppleID,omitempty"`
	// The Managed Apple Account pre-assigned to the authenticated user. Required for account-driven enrollments. Available in iOS 15 and later, and macOS 14 and later.
	// > Note:
	// > When updating the payload, the value of this key must not change. Any change is an error, and the update is rejected.
	AssignedManagedAppleID *string `json:"AssignedManagedAppleID,omitempty" plist:"AssignedManagedAppleID,omitempty"`
	// The enrollment mode the server indicates to use when enrolling. Required for account-driven enrollment. Available in iOS 15 and macOS 14, and later.
	// > Note:
	// > When updating the payload, the value of this key must not change. Any change is an error, and the update is rejected.
	EnrollmentMode *EnrollmentMode `json:"EnrollmentMode,omitempty" plist:"EnrollmentMode,omitempty"`
	// An array of strings, each containing the UUID of a certificate to use when evaluating trust to the '.../connect/' URLs of MDM servers.
	ServerURLPinningCertificateUUIDs *[]string `json:"ServerURLPinningCertificateUUIDs,omitempty" plist:"ServerURLPinningCertificateUUIDs,omitempty"`
	// An array of strings, each containing the payload UUID of a certificate to use when evaluating trust to the '.../checkin/' URLs of MDM servers.
	CheckInURLPinningCertificateUUIDs *[]string `json:"CheckInURLPinningCertificateUUIDs,omitempty" plist:"CheckInURLPinningCertificateUUIDs,omitempty"`
	// If 'true', the system fails the connection attempt unless it obtains a verified positive response during certificate revocation checks.
	// If 'false', the system performs revocation checks on a best-attempt basis, where failure to reach the server isn't considered fatal.
	PinningRevocationCheckRequired *bool `json:"PinningRevocationCheckRequired,omitempty" plist:"PinningRevocationCheckRequired,omitempty"`
	// A unique array of strings indicating server capabilities:
	// - `com.apple.mdm.per-user-connections`: Indicates that the server supports both device and user connections. This must be present when managing Shared iPad or macOS devices.
	// - `com.apple.mdm.bootstraptoken`: Indicates that the server supports escrowing the bootstrap token. This must be present for the device to create a bootstrap token and send it to the server. Available in iOS 26 and later, macOS 11 and later, and visionOS 26 and later.
	// - `com.apple.mdm.token`: Indicates that the server supports the `Get-Token` CheckIn message type. This must be present for the device to use `Get-Token` CheckIn message when appropriate.
	// > Note:
	// > When updating the payload, the `com.apple.mdm.per-user-connections` capability must not be added or removed. Any such change is an error, and the update is rejected.
	ServerCapabilities *[]ServerCapabilities `json:"ServerCapabilities,omitempty" plist:"ServerCapabilities,omitempty"`
	// If 'true', the device attempts to send a `Check-Out` message to the 'CheckInURL' when the profile is removed.
	CheckOutWhenRemoved *bool `json:"CheckOutWhenRemoved,omitempty" plist:"CheckOutWhenRemoved,omitempty"`
	// This property specifies an iTunes Store ID for an app the system can install with the InstallApplicationCommand, without any approval from the user. The MDM vendor or managing organization generally provides this app, which enhances the management experience for the user. The device shows the user details about this app in the account-driven enrollment process prior to installing the MDM profile. Use this property with account-driven MDM enrollments that normally require user approval for app installs through MDM.
	// Only account-driven enrollments support this property and other enrollment types ignore it.
	// Available in iOS 15.1 and later.
	// > Note:
	// > When updating the payload, the value of this key must not change. Any change is an error, and the update is rejected.
	RequiredAppIDForMDM *int64 `json:"RequiredAppIDForMDM,omitempty" plist:"RequiredAppIDForMDM,omitempty"`
	// If 'true', the system warns the user that they need to reboot into RecoveryOS and allow the MDM to use the bootstrap token for authentication for certain sensitive operations such as enabling kernel extensions or installing some types of software updates. If the MDM doesn't need to perform these operations, it can leave this key set to 'false', and the user isn't notified.
	// The SettingsCommand.Command.Settings.MDMOptions.MDMOptions command overrides this default value.
	// This setting only applies to devices that have 'BootstrapTokenRequiredForSoftwareUpdate' or 'BootstrapTokenRequiredForKernelExtensionApproval' set to 'true' in their SecurityInfoResponse.SecurityInfo.
	// DEP-enrolled devices are automatically allowed to use the bootstrap token for authentication.
	// Available in macOS 11 and later.
	PromptUserToAllowBootstrapTokenForAuthentication *bool `json:"PromptUserToAllowBootstrapTokenForAuthentication,omitempty" plist:"PromptUserToAllowBootstrapTokenForAuthentication,omitempty"`
}

func (p *MDM) PayloadType() string {
	return "com.apple.mdm"
}

// The enrollment mode the server indicates to use when enrolling. Required for account-driven enrollment. Available in iOS 15 and macOS 14, and later.
// > Note:
// > When updating the payload, the value of this key must not change. Any change is an error, and the update is rejected.
type EnrollmentMode string

const (
	EnrollmentModeBYOD EnrollmentMode = "BYOD"
	EnrollmentModeADDE EnrollmentMode = "ADDE"
)

// A unique array of strings indicating server capabilities:
// - `com.apple.mdm.per-user-connections`: Indicates that the server supports both device and user connections. This must be present when managing Shared iPad or macOS devices.
// - `com.apple.mdm.bootstraptoken`: Indicates that the server supports escrowing the bootstrap token. This must be present for the device to create a bootstrap token and send it to the server. Available in iOS 26 and later, macOS 11 and later, and visionOS 26 and later.
// - `com.apple.mdm.token`: Indicates that the server supports the `Get-Token` CheckIn message type. This must be present for the device to use `Get-Token` CheckIn message when appropriate.
// > Note:
// > When updating the payload, the `com.apple.mdm.per-user-connections` capability must not be added or removed. Any such change is an error, and the update is rejected.
type ServerCapabilities string

const (
	ServerCapabilitiesComapplemdmperUserConnections ServerCapabilities = "com.apple.mdm.per-user-connections"
	ServerCapabilitiesComapplemdmbootstraptoken     ServerCapabilities = "com.apple.mdm.bootstraptoken"
	ServerCapabilitiesComapplemdmtoken              ServerCapabilities = "com.apple.mdm.token"
)

// The payload that configures a passcode policy.
type Passcode struct {
	*CommonPayloadKeys
	// If `false`, the system prevents use of a simple passcode. A simple passcode contains repeated characters, or increasing or decreasing characters, such as `123` or `CBA`.
	AllowSimple *bool `json:"allowSimple,omitempty" plist:"allowSimple,omitempty"`
	// If `true`, the system forces the user to enter a PIN.
	ForcePIN *bool `json:"forcePIN,omitempty" plist:"forcePIN,omitempty"`
	// The number of failed passcode attempts that the system allows the user before it erases or locks the device. After six failed attempts, the device imposes a time delay before the user can enter a passcode again. The time delay increases with each failed attempt. On macOS, set `minutesUntilFailedLoginReset` to define the time delay. The time delay begins after the sixth attempt, so if `MaximumFailedAttempts` is six or lower, the system has no time delay and triggers the erase or lock as soon as the user exceeds the limit.
	// After the final failed attempt, the system locks a macOS device, or securely erases all data and settings from an iOS, visionOS, or watchOS device.
	MaxFailedAttempts *int64 `default:"11" json:"maxFailedAttempts,omitempty" plist:"maxFailedAttempts,omitempty"`
	// The maximum number of minutes for which the device can be idle without the user unlocking it, before the system locks it. When this limit is reached, the system locks the device and the passcode is required to unlock it. The user can edit this setting, but the value can't exceed the `maxInactivity` value.
	// On macOS, the system translates this inactivity value to screen-saver settings. The maximum value for macOS is `60`.
	// Setting this key removes the `never` option in the Settings UI on user enrolled devices.
	MaxInactivity *int64 `json:"maxInactivity,omitempty" plist:"maxInactivity,omitempty"`
	// The number of days for which the passcode can remain unchanged. After this number of days, the system forces the user to change the passcode before it unlocks the device.
	MaxPINAgeInDays *int64 `json:"maxPINAgeInDays,omitempty" plist:"maxPINAgeInDays,omitempty"`
	// The minimum number of complex characters that a passcode needs to contain. A _complex_ character is a character other than a number or a letter, such as `&`, `%`, `$`, and `#`.
	// The system ignores this property for user enrollments.
	MinComplexChars *int64 `default:"0" json:"minComplexChars,omitempty" plist:"minComplexChars,omitempty"`
	// The minimum overall length of the passcode. This value is independent of the value for `minComplexChars`.
	MinLength *int64 `default:"0" json:"minLength,omitempty" plist:"minLength,omitempty"`
	// If `true`, the system requires alphabetic characters instead of only numeric characters.
	RequireAlphanumeric *bool `json:"requireAlphanumeric,omitempty" plist:"requireAlphanumeric,omitempty"`
	// This value defines _N_, where the new passcode must be unique within the last _N_ entries in the passcode history.
	PinHistory *int64 `json:"pinHistory,omitempty" plist:"pinHistory,omitempty"`
	// The maximum grace period, in minutes, to unlock the phone without entering a passcode. The default is `0`, which is no grace period and requires a passcode immediately. On macOS, the system translates this grace period value to screen-saver settings.
	MaxGracePeriod *int64 `default:"0" json:"maxGracePeriod,omitempty" plist:"maxGracePeriod,omitempty"`
	// The number of minutes before the system resets the login after the maximum number of unsuccessful login attempts is reached. This key requires setting `maxFailedAttempts`. Available in macOS 10.10 and later.
	MinutesUntilFailedLoginReset *int64 `json:"minutesUntilFailedLoginReset,omitempty" plist:"minutesUntilFailedLoginReset,omitempty"`
	// If `true`, the system causes a password reset to occur the next time the user tries to authenticate. If this key is set in a device profile, the setting takes effect for all users, and admin authentications may fail until the admin user password is also reset. Available in macOS 10.13 and later.
	ChangeAtNextAuth *bool `json:"changeAtNextAuth,omitempty" plist:"changeAtNextAuth,omitempty"`
	// Specifies a regular expression, and its description, used to enforce password compliance. Use the simpler passcode restrictions whenever possible, and rely on regular expression matching only when necessary. Mistakes in regular expressions can lead to frustrating user experiences, such as unsatisfiable passcode policies, or policy descriptions that don't match the enforced policy.
	// Available in macOS 14 and later.
	CustomRegex *CustomRegex `json:"customRegex,omitempty" plist:"customRegex,omitempty"`
}

func (p *Passcode) PayloadType() string {
	return "com.apple.mobiledevice.passwordpolicy"
}

// Specifies a regular expression, and its description, used to enforce password compliance. Use the simpler passcode restrictions whenever possible, and rely on regular expression matching only when necessary. Mistakes in regular expressions can lead to frustrating user experiences, such as unsatisfiable passcode policies, or policy descriptions that don't match the enforced policy.
// Available in macOS 14 and later.
type CustomRegex struct {
	// A regular expression string that the system matches against the password to determine whether it complies with a policy. The regular expression uses the ICU syntax ([https://unicode-org.github.io/icu/userguide/strings/regexp.html](https://unicode-org.github.io/icu/userguide/strings/regexp.html)). The string must not exceed 2048 characters in length.
	PasswordContentRegex string `json:"passwordContentRegex" plist:"passwordContentRegex" required:"true"`
	// Contains a dictionary of keys for supported OS language IDs (for example, "en-US"), and whose values represent a localized description of the policy enforced by the regular expression. Use the special `default` key can for languages that aren't contained in the dictionary.
	PasswordContentDescription *map[string]string `json:"passwordContentDescription,omitempty" plist:"passwordContentDescription,omitempty"`
}

// The payload that configures network-usage rules.
// Network Usage Rules allow enterprises to specify how devices use networks, such as cellular data networks. iOS 9-12 support only ApplicationRules. In iOS 13, ApplicationRules, SIMRules, or both must be present.
type NetworkUsageRules struct {
	*CommonPayloadKeys
	// An array of application rules, that apply to only managed apps.
	ApplicationRules *[]*ApplicationRulesItem `json:"ApplicationRules,omitempty" plist:"ApplicationRules,omitempty"`
	// An array of SIM rules, that apply to all apps.
	SIMRules *[]*SIMRulesItem `json:"SIMRules,omitempty" plist:"SIMRules,omitempty"`
}

func (p *NetworkUsageRules) PayloadType() string {
	return "com.apple.networkusagerules"
}

// The application rules dictionary.
type ApplicationRulesItem struct {
	// A list of managed app identifiers, as strings, that must follow the associated rules. If this key is missing, the rules apply to all managed apps on the device.
	// Each string in the `AppIdentifierMatches` array may either be an exact app identifier match (for example, `com.mycompany.myapp`) or it may specify a prefix match for the bundle ID by using the \* wildcard character. If used, this character must appear after a period (.) and may only appear once, at the end of the string; for example, `com.mycompany.*`.
	AppIdentifierMatches *[]string `json:"AppIdentifierMatches,omitempty" plist:"AppIdentifierMatches,omitempty"`
	// If `false`, disables cellular data while roaming for all matching managed apps.
	AllowRoamingCellularData *bool `json:"AllowRoamingCellularData,omitempty" plist:"AllowRoamingCellularData,omitempty"`
	// If `false`, disables cellular data for all matching managed apps.
	AllowCellularData *bool `json:"AllowCellularData,omitempty" plist:"AllowCellularData,omitempty"`
}

// The policy for individual SIM cards.
type SIMRulesItem struct {
	// One or more ICCIDs of SIM cards for which the `WiFiAssistPolicy` applies. All ICCIDs in all installed Network Usage Rules payloads must be unique. An example ICCID is `89310410106543789301`.
	ICCIDs []string `json:"ICCIDs" plist:"ICCIDs" required:"true"`
	// The Wi-Fi Assist policy to apply to the SIM cards specified in the ICCIDs. Allowed values:
	// - `2`: Use the default system policy for the specified SIM card(s).
	// - `3`: Make Wi-Fi Assist switch more aggressively from a poor Wi-Fi connection to cellular data for the specified SIM card(s). This setting may increase cellular data use and may impact battery life.
	// For more information, see [About Wi-Fi Assist](https://support.apple.com/en-us/HT205296).
	WiFiAssistPolicy WiFiAssistPolicy `json:"WiFiAssistPolicy" plist:"WiFiAssistPolicy" required:"true"`
}

// The Wi-Fi Assist policy to apply to the SIM cards specified in the ICCIDs. Allowed values:
// - `2`: Use the default system policy for the specified SIM card(s).
// - `3`: Make Wi-Fi Assist switch more aggressively from a poor Wi-Fi connection to cellular data for the specified SIM card(s). This setting may increase cellular data use and may impact battery life.
// For more information, see [About Wi-Fi Assist](https://support.apple.com/en-us/HT205296).
type WiFiAssistPolicy int64

const (
	WiFiAssistPolicy2 WiFiAssistPolicy = 2
	WiFiAssistPolicy3 WiFiAssistPolicy = 3
)

// The payload that configures notifications.
// A notification settings payload specifies the restriction enforced notification settings for apps using their bundle identifier. The profile specifies notification settings by bundle identifier (even for apps that aren’t installed on the device yet), and those settings will always be enforced.
type Notifications struct {
	*CommonPayloadKeys
	// An array of notification settings dictionaries.
	NotificationSettings []*NotificationSettingsItem `json:"NotificationSettings" plist:"NotificationSettings" required:"true"`
}

func (p *Notifications) PayloadType() string {
	return "com.apple.notificationsettings"
}

type NotificationSettingsItem struct {
	// The bundle identifier of the app to which to apply these notification settings.
	// Available in iOS 9.3 and later and macOS 10.15 and later.
	BundleIdentifier string `json:"BundleIdentifier" plist:"BundleIdentifier" required:"true"`
	// If `true`, enables notifications for this app.
	// Available in iOS 9.3 and later and macOS 10.15 and later.
	NotificationsEnabled *bool `json:"NotificationsEnabled,omitempty" plist:"NotificationsEnabled,omitempty"`
	// If `true`, enables notifications in the notification center for this app.
	// Available in iOS 9.3 and later and macOS 10.15 and later.
	ShowInNotificationCenter *bool `json:"ShowInNotificationCenter,omitempty" plist:"ShowInNotificationCenter,omitempty"`
	// If `true`, enables notifications on the Lock Screen for this app.
	// Available in iOS 9.3 and later and macOS 10.15 and later.
	ShowInLockScreen *bool `json:"ShowInLockScreen,omitempty" plist:"ShowInLockScreen,omitempty"`
	// The type of alert for notifications for this app:
	// - `0`: None
	// - `1`: Temporary Banner
	// - `2`: Persistent Banner
	// Available in iOS 9.3 and later and macOS 10.15 and later.
	AlertType *AlertType `default:"1" json:"AlertType,omitempty" plist:"AlertType,omitempty"`
	// If `true`, enables badges for this app.
	// Available in iOS 9.3 and later and macOS 10.15 and later.
	BadgesEnabled *bool `json:"BadgesEnabled,omitempty" plist:"BadgesEnabled,omitempty"`
	// If `true`, enables sounds for this app.
	SoundsEnabled *bool `json:"SoundsEnabled,omitempty" plist:"SoundsEnabled,omitempty"`
	// If `true`, enables notifications in CarPlay for this app.
	// Available in iOS 12 and later.
	ShowInCarPlay *bool `json:"ShowInCarPlay,omitempty" plist:"ShowInCarPlay,omitempty"`
	// If `true`, enables critical alerts that can ignore Do Not Disturb and ringer settings for this app.
	// Available in iOS 12 and later and macOS 10.15 and later.
	CriticalAlertEnabled *bool `json:"CriticalAlertEnabled,omitempty" plist:"CriticalAlertEnabled,omitempty"`
	// The type of grouping for notifications for this app:
	// - `0`: Automatic: Group notifications into app-specified groups.
	// - `1`: By app: Group notifications into one group.
	// - `2`: Off: Don't group notifications.
	// Available in iOS 12 and later.
	GroupingType *GroupingType `default:"0" json:"GroupingType,omitempty" plist:"GroupingType,omitempty"`
	// The type previews for notifications. This key overrides the value at Settings>Notifications>Show Previews.
	// - `0` - Always: Previews will be shown when the device is locked and unlocked
	// - `1` - When Unlocked: Previews will only be shown when the device is unlocked
	// - `2` - Never: Previews will never be shown
	// Available in iOS 14 and later.
	PreviewType *PreviewType `json:"PreviewType,omitempty" plist:"PreviewType,omitempty"`
}

// The type of alert for notifications for this app:
// - `0`: None
// - `1`: Temporary Banner
// - `2`: Persistent Banner
// Available in iOS 9.3 and later and macOS 10.15 and later.
type AlertType int64

const (
	AlertType0 AlertType = 0
	AlertType1 AlertType = 1
	AlertType2 AlertType = 2
)

// The type of grouping for notifications for this app:
// - `0`: Automatic: Group notifications into app-specified groups.
// - `1`: By app: Group notifications into one group.
// - `2`: Off: Don't group notifications.
// Available in iOS 12 and later.
type GroupingType int64

const (
	GroupingType0 GroupingType = 0
	GroupingType1 GroupingType = 1
	GroupingType2 GroupingType = 2
)

// The type previews for notifications. This key overrides the value at Settings>Notifications>Show Previews.
// - `0` - Always: Previews will be shown when the device is locked and unlocked
// - `1` - When Unlocked: Previews will only be shown when the device is unlocked
// - `2` - Never: Previews will never be shown
// Available in iOS 14 and later.
type PreviewType int64

const (
	PreviewType0 PreviewType = 0
	PreviewType1 PreviewType = 1
	PreviewType2 PreviewType = 2
)

// The payload that configures a macOS Server account.
type MacOSServerAccount struct {
	*CommonPayloadKeys
	// The server's address.
	HostName string `json:"HostName" plist:"HostName" required:"true"`
	// The user's user name.
	UserName string `json:"UserName" plist:"UserName" required:"true"`
	// The user's password.
	Password *string `json:"Password,omitempty" plist:"Password,omitempty"`
	// The description of the account.
	AccountDescription *string `json:"AccountDescription,omitempty" plist:"AccountDescription,omitempty"`
	// An array of dictionaries containing configured account types and relevant settings
	ConfiguredAccounts []*ConfiguredAccountsItem `json:"ConfiguredAccounts" plist:"ConfiguredAccounts" required:"true"`
}

func (p *MacOSServerAccount) PayloadType() string {
	return "com.apple.osxserver.account"
}

type ConfiguredAccountsItem struct {
	// com.apple.osxserver.documents (the Documents account type).
	Type ConfiguredAccountsItemType `json:"Type" plist:"Type" required:"true"`
	// Designates the port number to use when contacting the server. If no port number is specified, the default port is used.
	Port *int64 `json:"Port,omitempty" plist:"Port,omitempty"`
}

// com.apple.osxserver.documents (the Documents account type).
type ConfiguredAccountsItemType string

const (
	ConfiguredAccountsItemTypeComappleosxserverdocuments ConfiguredAccountsItemType = "com.apple.osxserver.documents"
)

// The payload that configures security preferences.
type SecurityPreferences struct {
	*CommonPayloadKeys
	// If `true`, disables user changes to the password.
	DontAllowPasswordResetUI *bool `json:"dontAllowPasswordResetUI,omitempty" plist:"dontAllowPasswordResetUI,omitempty"`
	// If `true`, disables user changes to the lock message.
	DontAllowLockMessageUI *bool `json:"dontAllowLockMessageUI,omitempty" plist:"dontAllowLockMessageUI,omitempty"`
	// If `true`, disables user changes to the firewall settings.
	DontAllowFireWallUI *bool `json:"dontAllowFireWallUI,omitempty" plist:"dontAllowFireWallUI,omitempty"`
}

func (p *SecurityPreferences) PayloadType() string {
	return "com.apple.preference.security"
}

// The payload that configures iCloud password preferences.
type UserPreferences struct {
	*CommonPayloadKeys
	// If `true`, disables the iCloud password for local accounts.
	DisableUsingiCloudPassword *bool `json:"DisableUsingiCloudPassword,omitempty" plist:"DisableUsingiCloudPassword,omitempty"`
}

func (p *UserPreferences) PayloadType() string {
	return "com.apple.preference.users"
}

// The payload that configures profile removal.
type ProfileRemovalPassword struct {
	*CommonPayloadKeys
	// The password to allow removing the profile.
	RemovalPassword *string `json:"RemovalPassword,omitempty" plist:"RemovalPassword,omitempty"`
}

func (p *ProfileRemovalPassword) PayloadType() string {
	return "com.apple.profileRemovalPassword"
}

// The payload that configures a global HTTP proxy.
// PEM-encoded cer
type GlobalHTTPProxy struct {
	*CommonPayloadKeys
	// The proxy type. For a manual proxy type, the profile contains the proxy server address, including its port, and optionally a user name and password. For an auto proxy type, you can enter a PAC URL.
	ProxyType *GlobalHTTPProxyProxyType `default:"Manual" json:"ProxyType,omitempty" plist:"ProxyType,omitempty"`
	// The proxy server's network address. The device requires this if `ProxyType` is set to `Manual`, and ignores it if `ProxyType` is set to `Automatic`.
	ProxyServer *string `json:"ProxyServer,omitempty" plist:"ProxyServer,omitempty"`
	// The proxy server's port number. The device requires this if `ProxyType` is set to `Manual`, and ignores this if `ProxyType` is set to `Automatic`.
	ProxyServerPort *int64 `json:"ProxyServerPort,omitempty" plist:"ProxyServerPort,omitempty"`
	// The user name used to authenticate to the proxy server. The device only uses this if `ProxyType` is set to `Manual`, and ignores it if `ProxyType` is set to `Automatic`.
	ProxyUsername *string `json:"ProxyUsername,omitempty" plist:"ProxyUsername,omitempty"`
	// The password used to authenticate to the proxy server. The device only uses this if `ProxyType` is set to `Manual`, and ignores it if `ProxyType` is set to `Automatic`.
	ProxyPassword *string `json:"ProxyPassword,omitempty" plist:"ProxyPassword,omitempty"`
	// The URL of the PAC file that defines the proxy configuration. Starting in iOS 13 and macOS 10.15, only URLs that begin with `http://` or `https://` are allowed. This is only used if `ProxyType` is set to `Automatic`, and is ignored if `ProxyType` is set to `Manual`.
	ProxyPACURL *string `json:"ProxyPACURL,omitempty" plist:"ProxyPACURL,omitempty"`
	// If `true`, allows connecting directly to the destination if the proxy autoconfiguration (PAC) file is unreachable.
	ProxyPACFallbackAllowed *bool `json:"ProxyPACFallbackAllowed,omitempty" plist:"ProxyPACFallbackAllowed,omitempty"`
	// If `true`, allows the device to bypass the proxy server to display the login page for captive networks.
	ProxyCaptiveLoginAllowed *bool `json:"ProxyCaptiveLoginAllowed,omitempty" plist:"ProxyCaptiveLoginAllowed,omitempty"`
}

func (p *GlobalHTTPProxy) PayloadType() string {
	return "com.apple.proxy.http.global"
}

// The proxy type. For a manual proxy type, the profile contains the proxy server address, including its port, and optionally a user name and password. For an auto proxy type, you can enter a PAC URL.
type GlobalHTTPProxyProxyType string

const (
	GlobalHTTPProxyProxyTypeManual GlobalHTTPProxyProxyType = "Manual"
	GlobalHTTPProxyProxyTypeAuto   GlobalHTTPProxyProxyType = "Auto"
)

// The payload that configures relay settings.
type Relay1 struct {
	*CommonPayloadKeys
	// An array of dictionaries that describe one or more relay servers that the system can chain together.
	Relays []*RelaysRelay `json:"Relays" plist:"Relays" required:"true"`
	// A list of domain strings that the system uses to determine which connection to route through the servers in `Relays`.
	// Any connection that matches a domain in the list exactly or is a subdomain of the listed domain uses the relay servers, unless it matches a domain in `ExcludedDomains`.
	// If this list and `MatchFQDNs` are empty, the system routes traffic to all domains to the relay servers, except those that match an excluded domain or excluded FQDN.
	MatchDomains *[]string `json:"MatchDomains,omitempty" plist:"MatchDomains,omitempty"`
	// A list of domain strings to exclude from routing through the servers in `Relays`. Any connection that matches a domain in the list exactly or is a subdomain of the listed domain won't use the relay server.
	ExcludedDomains *[]string `json:"ExcludedDomains,omitempty" plist:"ExcludedDomains,omitempty"`
	// A list of Fully Qualified Domain Names (FQDNs) to be routed through the servers contained in `Relays`. Any connection that matches an FQDN in the list exactly uses the relay servers. If this list and `MatchDomains` are empty, the system routes traffic to all domains to the relay servers, except those that match an excluded domain or excluded FQDN.
	MatchFQDNs *[]string `json:"MatchFQDNs,omitempty" plist:"MatchFQDNs,omitempty"`
	// A list of Fully Qualified Domain Names (FQDNs) to exclude from routing through the servers contained in `Relays`. Any connection that matches an FQDN in the list exactly won't use the relay server. When `MatchDomains` is also present, any FQDN listed in the list should be a subdomain of at least one `MatchDomain` value, otherwise it will not have any effect.
	ExcludedFQDNs *[]string `json:"ExcludedFQDNs,omitempty" plist:"ExcludedFQDNs,omitempty"`
	// A globally unique identifier for this relay configuration. The system uses this UUID to route managed apps through the servers in `Relays`. This key is required for user enrollment.
	RelayUUID *string `json:"RelayUUID,omitempty" plist:"RelayUUID,omitempty"`
	// If `true`, the device allows the user to disable this network relay configuration.
	UIToggleEnabled *bool `json:"UIToggleEnabled,omitempty" plist:"UIToggleEnabled,omitempty"`
	// If `true`, the device allows the relay to failover to the default system DNS resolver.
	AllowDNSFailover *bool `json:"AllowDNSFailover,omitempty" plist:"AllowDNSFailover,omitempty"`
}

func (p *Relay1) PayloadType() string {
	return "com.apple.relay.managed"
}

type RelaysRelay struct {
	// The URL or URI template, as defined in RFC 9298, of a relay server that's reachable using HTTP/3 and supports proxying TCP and UDP using the CONNECT method.
	// Each relay needs to include either `HTTP2RelayURL` or `HTTP3RelayURL`, or it can include both.
	HTTP3RelayURL *string `json:"HTTP3RelayURL,omitempty" plist:"HTTP3RelayURL,omitempty"`
	// The URL or URI template, as defined in RFC 9298, of a relay server that's reachable using HTTP/2 and supports proxying TCP and UDP using the CONNECT method.
	// Each relay needs to include either `HTTP2RelayURL` or `HTTP3RelayURL`, or it can include both.
	HTTP2RelayURL *string `json:"HTTP2RelayURL,omitempty" plist:"HTTP2RelayURL,omitempty"`
	// A dictionary that contains custom HTTP header keys and values to add to each request. The dictionary key name represents the HTTP header field name to use, and the dictionary value is the string to use as the HTTP header field value.
	AdditionalHTTPHeaderFields *map[string]string `json:"AdditionalHTTPHeaderFields,omitempty" plist:"AdditionalHTTPHeaderFields,omitempty"`
	// The UUID that points to an identity certificate payload, which the system uses to authenticate the user to the relay server.
	PayloadCertificateUUID *string `json:"PayloadCertificateUUID,omitempty" plist:"PayloadCertificateUUID,omitempty"`
	// An array of DER-encoded raw public keys that the system uses to authenticate the server during a TLS handshake. The server needs to use one of the keys in the handshake to authenticate.
	// If this array is empty, the system uses the default TLS trust evaluation.
	RawPublicKeys *[][]byte `json:"RawPublicKeys,omitempty" plist:"RawPublicKeys,omitempty"`
}

// The payload that configures a user's screen saver settings.
// Specifies *user* screen saver settings. (Settings for Login Window screen saver use a different payload)
type ScreensaverUser struct {
	*CommonPayloadKeys
	// The name of the screen saver module.
	ModuleName string `json:"moduleName" plist:"moduleName" required:"true"`
	// A full path to the screen saver module to use.
	ModulePath *string `json:"modulePath,omitempty" plist:"modulePath,omitempty"`
	// The number of seconds of inactivity before the screen saver activates (`0` = Never activate).
	IdleTime *int64 `json:"idleTime,omitempty" plist:"idleTime,omitempty"`
}

func (p *ScreensaverUser) PayloadType() string {
	return "com.apple.screensaver.user"
}

// The payload that configures the screen saver.
// Specifies grace period for screensaver locking
type Screensaver struct {
	*CommonPayloadKeys
	// If `true`, the user is prompted for a password when the screen saver is unlocked or stopped. When you use this prompt, you must also provide `askForPasswordDelay`. Available in macOS 10.13 and later.
	AskForPassword *bool `json:"askForPassword,omitempty" plist:"askForPassword,omitempty"`
	// The number of seconds to delay before the password will be required to unlock or stop the screen saver (the grace period). A value of `2147483647` (for example, `0x7FFFFFFF`) disables this requirement. To use this option, you must set `askForPassword` to `true`. Available in macOS 10.13 and later.
	AskForPasswordDelay *int64 `json:"askForPasswordDelay,omitempty" plist:"askForPasswordDelay,omitempty"`
	// The number of seconds of inactivity before the screen saver activates (0 = Never activate).
	IdleTime *int64 `json:"idleTime,omitempty" plist:"idleTime,omitempty"`
	// The full path to the screen-saver module to use.
	LoginWindowModulePath *string `json:"loginWindowModulePath,omitempty" plist:"loginWindowModulePath,omitempty"`
	// The name of the screen saver module.
	ModuleName string `json:"moduleName" plist:"moduleName" required:"true"`
}

func (p *Screensaver) PayloadType() string {
	return "com.apple.screensaver"
}

// The payload that configures the second wired, active Ethernet interface.
type Net8021XSecondActiveEthernet struct {
	*CommonPayloadKeys
	Properties map[string]any
}

func (p *Net8021XSecondActiveEthernet) PayloadType() string {
	return "com.apple.secondactiveethernet.managed"
}

// The payload that configures the second wired Ethernet interface.
type Net8021XSecondEthernet struct {
	*CommonPayloadKeys
	Properties map[string]any
}

func (p *Net8021XSecondEthernet) PayloadType() string {
	return "com.apple.secondethernet.managed"
}

// The payload that configures FileVault recovery key escrow.
// If FileVault is enabled after this payload is installed on the system, the FileVault PRK will be encrypted with the specified certificate, wrapped with a CMS envelope and stored at:
// /var/db/FileVaultPRK.dat
// The encrypted data will be made available to the MDM server as part of the SecurityInfo command. Alternatively, if a site uses their own administration software, they can extract the PRK from the above location at any time. As the PRK will be encrypted using the certificate provided in the profile, only the author of the profile can extract the data.
// Notes:
// * The payload must exist in a "system" scoped profile.
// * It will be an error to install more than one payload of this type per machine.
// * The old payload ("com.apple.security.FDERecoveryRedirect") will no longer be supported. It will still be allowed to be installed, but will be ignored. (This is so servers can send out the same profile to old and new clients).
// * If only an old-style redirection payload is installed at the time FileVault is turned on (via Security Pref pane), an error will be displayed and FileVault will not be allowed to be enabled.
// * No warning/error will be provided if FileVault is already enabled and an old-style payload is installed. In this case, it's assumed the recovery key has already been escrowed with the server.
type FDERecoveryKeyEscrow struct {
	*CommonPayloadKeys
	// The description of the location where the system escrows the recovery key. The system inserts this text into the message the user sees when it enables FileVault.
	Location string `json:"Location" plist:"Location" required:"true"`
	// The UUID of a payload within the same profile that contains the certificate that the system uses to encrypt the recovery key. The referenced payload must be of type `com.apple.security.pkcs1`.
	EncryptCertPayloadUUID string `json:"EncryptCertPayloadUUID" plist:"EncryptCertPayloadUUID" required:"true"`
	// The string that's included in help text if the user appears to have forgotten the password. Site admins can use this key to look up the escrowed key for the particular computer.
	// This key replaces the `RecordNumber` key used in the previous escrow mechanism. If the key is missing, the system uses the device serial number instead.
	DeviceKey *string `json:"DeviceKey,omitempty" plist:"DeviceKey,omitempty"`
}

func (p *FDERecoveryKeyEscrow) PayloadType() string {
	return "com.apple.security.FDERecoveryKeyEscrow"
}

// The payload that configures FileVault recovery key redirection.
// *** This payload will be ignored on macOS 10.13 and later. See "com.apple.security.FDERecoveryKeyEscrow" payload. ***
// Old notes:
// Once installed, this payload will cause any FDE (Full Disk Encryption) recovery keys to be redirected to the specified URL instead of being sent to Apple. This will require sites to implement their own HTTPS server that will receive the recovery keys via a POST request. Details of the data sent to the server will be provided in a different document.
// Notes:
// * The payload must exist in a "system" scoped profile.
// * It will be an error to install more than one payload of this type per machine.
type FDERecoveryKeyRedirection struct {
	*CommonPayloadKeys
	// The URL to which FDE recovery keys should be sent instead of to Apple. The URL must begin with https://.
	RedirectURL string `json:"RedirectURL" plist:"RedirectURL" required:"true"`
	// The UUID of a payload within the same profile that contains a certificate used to encrypt the recovery key when it's sent to the redirected URL. The referenced payload must be of type \`com.apple.security.pkcs1\`.
	EncryptCertPayloadUUID string `json:"EncryptCertPayloadUUID" plist:"EncryptCertPayloadUUID" required:"true"`
}

func (p *FDERecoveryKeyRedirection) PayloadType() string {
	return "com.apple.security.FDERecoveryRedirect"
}

// The payload that configures Automated Certificate Management Environment (ACME) settings.
type ACMECertificate struct {
	*CommonPayloadKeys
	// The directory URL of the ACME server. The URL must use the https scheme.
	DirectoryURL string `json:"DirectoryURL" plist:"DirectoryURL" required:"true"`
	// A unique string identifying a specific device. The server may use this as an anti-replay code to prevent issuing multiple certificates. This identifier also indicates to the ACME server that the device has access to a valid client identifier issued by the enterprise infrastructure. This can help the ACME server determine whether to trust the device. Though this is a relatively weak indication because of the risk that an attacker can intercept the client identifier.
	ClientIdentifier string `json:"ClientIdentifier" plist:"ClientIdentifier" required:"true"`
	// The valid values for `KeySize` depend on the values of `KeyType` and `HardwareBound`. See those keys for specific requirements.
	KeySize int64 `json:"KeySize" plist:"KeySize" required:"true"`
	// The type of key pair to generate. Allowed values:
	// - `RSA`: Specifies an RSA key pair. RSA key pairs need to have a `KeySize` that's a multiple of 8 in the range of 1024 through 4096 (inclusive), and `HardwareBound` needs to be `false`.
	// - `ECSECPrimeRandom`: Specifies a key pair on the P-192, P-256, P-384, or P-521 curves as defined in FIPS Pub 186-4. `KeySize` defines the particular curve, which needs to be `192`, `256`, `384`, or `521`. Hardware bound keys only support values of `256` and `384`.
	// > Note:
	// > The key size is `521`, not `512`, even though the other key sizes are multiples of 64.
	KeyType KeyType `json:"KeyType" plist:"KeyType" required:"true"`
	// If `false`, the private key isn't bound to the device.
	// If `true`, the private key is bound to the device. The Secure Enclave generates the key pair, and the private key is cryptographically entangled with a system key. This prevents the system from exporting the private key.
	// If `true`, `KeyType` must be `ECSECPrimeRandom` and `KeySize` must be 256 or 384.
	// Setting this key to `true` is supported as of macOS 14 on Apple Silicon and Intel devices that have a T2 chip. Older macOS versions or other Mac devices require this key but it must have a value of `false`.
	HardwareBound bool `json:"HardwareBound" plist:"HardwareBound" required:"true"`
	// The device requests this subject for the certificate that the ACME server issues. The ACME server may override or ignore this field in the certificate it issues.
	// The representation of a X.500 name represented as an array of OID and value. For example, `/C=US/O=Apple Inc./CN=foo/1.2.5.3=bar` corresponds to:
	// `[ [ ["C", "US"] ], [ ["O", "Apple Inc."] ], ..., [ [ "1.2.5.3", "bar" ] ] ]`
	// Dotted numbers can represent OIDs , with shortcuts for country (C), locality (L), state (ST), organization (O), organizational unit (OU), and common name (CN).
	Subject [][][]string `json:"Subject" plist:"Subject" required:"true"`
	// The Subject Alt Name that the device requests for the certificate that the ACME server issues. The ACME server may override or ignore this field in the certificate it issues.
	SubjectAltName *ACMECertificateSubjectAltName `json:"SubjectAltName,omitempty" plist:"SubjectAltName,omitempty"`
	// This value is a bit field.
	// - Bit `0x01` indicates digital signature.
	// - Bit `0x04` indicates encryption.
	// The device requests this key for the certificate that the ACME server issues. The ACME server may override or ignore this field in the certificate it issues.
	UsageFlags *int64 `json:"UsageFlags,omitempty" plist:"UsageFlags,omitempty"`
	// The value is an array of strings. Each string is an OID in dotted notation. For instance, `["1.3.6.1.5.5.7.3.2", "1.3.6.1.5.5.7.3.4"]` indicates client authentication and email protection.
	// The device requests this field for the certificate that the ACME server issues. The ACME server may override or ignore this field in the certificate it issues.
	ExtendedKeyUsage *[]string `json:"ExtendedKeyUsage,omitempty" plist:"ExtendedKeyUsage,omitempty"`
	// If `true`, the device provides attestations that describe the device and the generated key to the ACME server. The server can use the attestations as strong evidence that the key is bound to the device, and that the device has properties listed in the attestation. The server can use that as part of a trust score to decide whether to issue the requested certificate.
	// When `Attest` is `true`, `HardwareBound` also needs to be `true`.
	// Setting this key to `true` is supported as of macOS 14. Older macOS versions require this key but it must have a value of `false`. See below for hardware requirements.
	Attest *bool `json:"Attest,omitempty" plist:"Attest,omitempty"`
	// If `true`, the private key of the identity obtained through Automated Certificate Management Environment (ACME) needs to be tagged as "non-extractable" in the keychain.
	KeyIsExtractable *bool `json:"KeyIsExtractable,omitempty" plist:"KeyIsExtractable,omitempty"`
	// If `true`, all apps have access to the private key.
	AllowAllAppsAccess *bool `json:"AllowAllAppsAccess,omitempty" plist:"AllowAllAppsAccess,omitempty"`
}

func (p *ACMECertificate) PayloadType() string {
	return "com.apple.security.acme"
}

// The type of key pair to generate. Allowed values:
// - `RSA`: Specifies an RSA key pair. RSA key pairs need to have a `KeySize` that's a multiple of 8 in the range of 1024 through 4096 (inclusive), and `HardwareBound` needs to be `false`.
// - `ECSECPrimeRandom`: Specifies a key pair on the P-192, P-256, P-384, or P-521 curves as defined in FIPS Pub 186-4. `KeySize` defines the particular curve, which needs to be `192`, `256`, `384`, or `521`. Hardware bound keys only support values of `256` and `384`.
// > Note:
// > The key size is `521`, not `512`, even though the other key sizes are multiples of 64.
type KeyType string

const (
	KeyTypeRSA              KeyType = "RSA"
	KeyTypeECSECPrimeRandom KeyType = "ECSECPrimeRandom"
)

// The Subject Alt Name that the device requests for the certificate that the ACME server issues. The ACME server may override or ignore this field in the certificate it issues.
type ACMECertificateSubjectAltName struct {
	// The RFC 822 (email address) string.
	Rfc822Name *string `json:"rfc822Name,omitempty" plist:"rfc822Name,omitempty"`
	// The DNS name.
	DNSName *string `json:"dNSName,omitempty" plist:"dNSName,omitempty"`
	// The Uniform Resource Identifier.
	UniformResourceIdentifier *string `json:"uniformResourceIdentifier,omitempty" plist:"uniformResourceIdentifier,omitempty"`
	// The NT principal name. Use an other name OID set to `1.3.6.1.4.1.311.20.2.3`.
	NtPrincipalName *string `json:"ntPrincipalName,omitempty" plist:"ntPrincipalName,omitempty"`
}

// The payload that configures a certificate preference.
// Defines a Certificate Preference item in the user's keychain that references a certificate payload included in the same profile. Can only appear in a user profile (not a device profile). See also "com.apple.security.identitypreference" for setting up identity preferences.
type CertificatePreference struct {
	*CommonPayloadKeys
	// An email address (in RFC 822 format) or other name for which a preferred certificate is requested.
	Name string `json:"Name" plist:"Name" required:"true"`
	// The UUID of the certificate payload within the same profile to use for the identity credential.
	PayloadCertificateUUID string `json:"PayloadCertificateUUID" plist:"PayloadCertificateUUID" required:"true"`
}

func (p *CertificatePreference) PayloadType() string {
	return "com.apple.security.certificatepreference"
}

// The payload that configures certificate revocation checking.
// Policies that affect system-wide certificate revocation checking.
type CertificateRevocation struct {
	*CommonPayloadKeys
	// An array of certificates that the system checks for revocation.
	// Specifying a certificate authority (CA) enables revocation checking for all certificates chaining up to that CA.
	// It's not necessary to specify trusted root certificates because they're implicitly specified. See [https://support.apple.com/en-us/HT209143](https://support.apple.com/en-us/HT209143) for the available trusted root certificates for Apple operating systems.
	EnabledForCerts *[]*EnabledForCertsSubjectPublicKeyInfoHashDict `json:"EnabledForCerts,omitempty" plist:"EnabledForCerts,omitempty"`
}

func (p *CertificateRevocation) PayloadType() string {
	return "com.apple.security.certificaterevocation"
}

// A dictionary of hashed public keys.
type EnabledForCertsSubjectPublicKeyInfoHashDict struct {
	// The algorithm must be `sha256`.
	Algorithm EnabledForCertsSubjectPublicKeyInfoHashDictAlgorithm `json:"Algorithm" plist:"Algorithm" required:"true"`
	// The hash of the DER-encoding of the certificate's `subjectPublicKeyInfo`.
	// The hash field requires the data (`subjectPublicKeyInfo` hash) in a specific format: a Base64 encoded (binary) SHA-256 hash of the certificate's public key.
	Hash []byte `json:"Hash" plist:"Hash" required:"true"`
}

// The algorithm must be `sha256`.
type EnabledForCertsSubjectPublicKeyInfoHashDictAlgorithm string

const (
	EnabledForCertsSubjectPublicKeyInfoHashDictAlgorithmSha256 EnabledForCertsSubjectPublicKeyInfoHashDictAlgorithm = "sha256"
)

// The payload that configures certificate transparency enforcement.
// Policies that affect system-wide certificate transparency enforcement.
type CertificateTransparency struct {
	*CommonPayloadKeys
	// An array of certificates for which certificate transparency is disabled. One of the following conditions needs to be met to disable certificate transparency enforcement when this policy is set:
	// - The hash is of the server certificate's `subjectPublicKeyInfo`.
	// - The hash is of a `subjectPublicKeyInfo` that appears in a CA certificate in the certificate chain; the CA certificate is constrained through the X.509v3 `nameConstraints` extension. One or more `directoryName` `nameConstraints` are present in the `permittedSubtrees`, and the `directoryName` contains an `organizationName` attribute.
	// - The hash is of a `subjectPublicKeyInfo` that appears in a CA certificate in the certificate chain. The CA certificate has one or more `organizationName` attributes in the certificate `Subject`, and the server's certificate contains the same number of `organizationName` attributes, in the same order, and with byte-for-byte identical values.
	DisabledForCerts *[]*DisabledForCertsSubjectPublicKeyInfoHashDict `json:"DisabledForCerts,omitempty" plist:"DisabledForCerts,omitempty"`
	// An array of strings that represent the domains to exclude from certificate transparency enforcement. The system supports using a leading period (`.`) to signify subdomains. However, the system doesn't support wildcards. If you include a leading period, the domain can't be a top-level domain, such as `.com` and `.co.uk`.
	DisabledForDomains *[]string `json:"DisabledForDomains,omitempty" plist:"DisabledForDomains,omitempty"`
}

func (p *CertificateTransparency) PayloadType() string {
	return "com.apple.security.certificatetransparency"
}

// A dictionary of hashed public keys.
type DisabledForCertsSubjectPublicKeyInfoHashDict struct {
	// The algorithm must be `sha256`.
	Algorithm DisabledForCertsSubjectPublicKeyInfoHashDictAlgorithm `json:"Algorithm" plist:"Algorithm" required:"true"`
	// The hash of the DER-encoding of the certificate's `subjectPublicKeyInfo`.
	// The hash field requires the data (`subjectPublicKeyInfo` hash) in a specific format: a Base64 encoded (binary) SHA-256 hash of the certificate's public key.
	Hash []byte `json:"Hash" plist:"Hash" required:"true"`
}

// The algorithm must be `sha256`.
type DisabledForCertsSubjectPublicKeyInfoHashDictAlgorithm string

const (
	DisabledForCertsSubjectPublicKeyInfoHashDictAlgorithmSha256 DisabledForCertsSubjectPublicKeyInfoHashDictAlgorithm = "sha256"
)

// The payload that configures the firewall.
// Manages the Application Firewall settings (e.g. Security pref pane -> Firewall).
// Notes:
// * The payload must exist in a "system" scoped profile.
// * If more than one profile contains this payload, the most restrictive union of settings will be used.
type Firewall struct {
	*CommonPayloadKeys
	// If `true`, the system enables the firewall.
	EnableFirewall bool `json:"EnableFirewall" plist:"EnableFirewall" required:"true"`
	// If `true`, the system enables blocking all incoming connections.
	BlockAllIncoming *bool `json:"BlockAllIncoming,omitempty" plist:"BlockAllIncoming,omitempty"`
	// If `true`, the system enables stealth mode.
	EnableStealthMode *bool `json:"EnableStealthMode,omitempty" plist:"EnableStealthMode,omitempty"`
	// The list of apps with connections that the firewall controls.
	Applications *[]*ApplicationsItem `json:"Applications,omitempty" plist:"Applications,omitempty"`
	// If `true`, the system enables logging. Available in macOS 12 through macOS 14.6.
	EnableLogging *bool `json:"EnableLogging,omitempty" plist:"EnableLogging,omitempty"`
	// The type of logging. Available in macOS 12 and through macOS 14.6.
	LoggingOption *LoggingOption `json:"LoggingOption,omitempty" plist:"LoggingOption,omitempty"`
	// If `true`, the system allows built-in software to receive incoming connections. Available in macOS 12.3 and later.
	// > Note:
	// > The system ensures that `AllowSigned` always has a value. If missing from the payload, the system sets it to `true`.
	AllowSigned *bool `json:"AllowSigned,omitempty" plist:"AllowSigned,omitempty"`
	// If `true`, the system allows downloaded signed software to receive incoming connections. Available in macOS 12.3 and later.
	// > Note:
	// > The system ensures that `AllowSignedApp` always has a value. If missing from the payload, the system sets it to `true`.
	AllowSignedApp *bool `json:"AllowSignedApp,omitempty" plist:"AllowSignedApp,omitempty"`
}

func (p *Firewall) PayloadType() string {
	return "com.apple.security.firewall"
}

type ApplicationsItem struct {
	// The bundle identifier for the app.
	BundleID string `json:"BundleID" plist:"BundleID" required:"true"`
	// If `true`, the system allows connections for the app.
	Allowed bool `json:"Allowed" plist:"Allowed" required:"true"`
}

// The type of logging. Available in macOS 12 and through macOS 14.6.
type LoggingOption string

const (
	LoggingOptionThrottled LoggingOption = "throttled"
	LoggingOptionBrief     LoggingOption = "brief"
	LoggingOptionDetail    LoggingOption = "detail"
)

// The payload that configures the user's identity on the device.
// Defines an Identity Preference item in the user's keychain that references a identity payload included in the same profile. Can only appear in a user profile (not a device profile). See also "com.apple.security.certificatepreference" for setting up certificate preferences.
type IdentityPreference struct {
	*CommonPayloadKeys
	// The email address (in RFC 822 format), DNS host name, or other name that uniquely identifies a service requiring this identity.
	Name string `json:"Name" plist:"Name" required:"true"`
	// The UUID of the certificate payload within the same profile to use for the identity credential.
	PayloadCertificateUUID string `json:"PayloadCertificateUUID" plist:"PayloadCertificateUUID" required:"true"`
}

func (p *IdentityPreference) PayloadType() string {
	return "com.apple.security.identitypreference"
}

// The payload that configures a PEM-formatted certificate.
// PEM-encoded certificate without private key. May contain root certificates.
type CertificatePEM struct {
	*CommonPayloadKeys
	// The file name of the enclosed certificate.
	PayloadCertificateFileName *string `json:"PayloadCertificateFileName,omitempty" plist:"PayloadCertificateFileName,omitempty"`
	// The binary representation of the payload, encoded in Base64.
	PayloadContent []byte `json:"PayloadContent" plist:"PayloadContent" required:"true"`
}

func (p *CertificatePEM) PayloadType() string {
	return "com.apple.security.pem"
}

// The payload that configures a PKCS #1-formatted certificate.
// DER-encoded certificate without private key. May contain root certificates.
type CertificatePKCS1 struct {
	*CommonPayloadKeys
	// The file name of the enclosed certificate.
	PayloadCertificateFileName *string `json:"PayloadCertificateFileName,omitempty" plist:"PayloadCertificateFileName,omitempty"`
	// The binary representation of the payload, encoded in Base64.
	PayloadContent []byte `json:"PayloadContent" plist:"PayloadContent" required:"true"`
}

func (p *CertificatePKCS1) PayloadType() string {
	return "com.apple.security.pkcs1"
}

// The payload that configures a PKCS #12-formatted certificate.
// Password-protected identity certificate. Only one certificate may be included.
type CertificatePKCS12 struct {
	*CommonPayloadKeys
	// The file name of the enclosed certificate.
	PayloadCertificateFileName *string `json:"PayloadCertificateFileName,omitempty" plist:"PayloadCertificateFileName,omitempty"`
	// The binary representation of the payload, encoded in Base64.
	PayloadContent []byte `json:"PayloadContent" plist:"PayloadContent" required:"true"`
	// The password to the identity.
	Password *string `json:"Password,omitempty" plist:"Password,omitempty"`
	// If `true`, the system allows apps access to the private key. Available in macOS 10.10 and later.
	AllowAllAppsAccess *bool `json:"AllowAllAppsAccess,omitempty" plist:"AllowAllAppsAccess,omitempty"`
	// If `false`, the system doesn't tag the private key data as extractable in the keychain.
	KeyIsExtractable *bool `json:"KeyIsExtractable,omitempty" plist:"KeyIsExtractable,omitempty"`
}

func (p *CertificatePKCS12) PayloadType() string {
	return "com.apple.security.pkcs12"
}

// The payload that configures a root certificate.
// Alias for com.apple.security.pkcs1.
type CertificateRoot struct {
	*CommonPayloadKeys
	// The file name of the enclosed certificate.
	PayloadCertificateFileName *string `json:"PayloadCertificateFileName,omitempty" plist:"PayloadCertificateFileName,omitempty"`
	// The binary representation of the payload encoded in base64.
	PayloadContent []byte `json:"PayloadContent" plist:"PayloadContent" required:"true"`
}

func (p *CertificateRoot) PayloadType() string {
	return "com.apple.security.root"
}

// The payload that configures Simple Certificate Enrollment Protocol (SCEP) settings.
type SCEP struct {
	*CommonPayloadKeys
	// A dictionary containing the SCEP information.
	PayloadContent PayloadContent `json:"PayloadContent" plist:"PayloadContent" required:"true"`
}

func (p *SCEP) PayloadType() string {
	return "com.apple.security.scep"
}

// A dictionary containing the SCEP information.
type PayloadContent struct {
	// The SCEP URL. See Over-the-Air Profile Delivery and Configuration for more information about SCEP.
	URL string `json:"URL" plist:"URL" required:"true"`
	// A string that's understood by the SCEP server; for example, a domain name like example.org. If a certificate authority has multiple CA certificates, this field can be used to distinguish which is required.
	Name *string `json:"Name,omitempty" plist:"Name,omitempty"`
	// The representation of an X.500 name as an array of OID and value.
	// For example, `/C=US/O=Apple Inc./CN=foo/1.2.5.3=bar` translates to `[ [ ["C", "US"] ], [ ["O", "Apple Inc."] ], …, [ [ "1.2.5.3", "bar" ] ] ]`.
	// OIDs can be represented as dotted numbers, with shortcuts for country (C), locality (L), state (ST), organization (O), organizational unit (OU), and common name (CN).
	Subject *[][][]string `json:"Subject,omitempty" plist:"Subject,omitempty"`
	// A preshared secret.
	Challenge *string `json:"Challenge,omitempty" plist:"Challenge,omitempty"`
	// The key size, in bits.
	Keysize *Keysize `default:"1024" json:"Keysize,omitempty" plist:"Keysize,omitempty"`
	// Always `RSA`.
	KeyType *string `default:"RSA" json:"Key Type,omitempty" plist:"Key Type,omitempty"`
	// A bitmask indicating the use of the key. Possible values:
	// - `1`: Signing
	// - `4`: Encryption
	// Some certificate authorities, such as Windows CA, support only encryption or signing, but not both at the same time.
	KeyUsage *int64 `default:"0" json:"Key Usage,omitempty" plist:"Key Usage,omitempty"`
	// The fingerprint of the Certificate Authority certificate.
	CAFingerprint *[]byte `json:"CAFingerprint,omitempty" plist:"CAFingerprint,omitempty"`
	// The number of times the device should retry if the server sends a PENDING response.
	Retries *int64 `default:"3" json:"Retries,omitempty" plist:"Retries,omitempty"`
	// The number of seconds to wait between subsequent retries. The first retry is attempted without this delay.
	RetryDelay *int64 `default:"10" json:"RetryDelay,omitempty" plist:"RetryDelay,omitempty"`
	// The SCEP payload can specify an optional `SubjectAltName` dictionary that provides values required by the CA for issuing a certificate. You can specify a single string or an array of strings for each key. The values you specify depend on the CA you're using, but might include DNS name, URL, or email values. For an example, see Sample Configuration Profile or Over-the-Air Profile Delivery and Configuration.
	SubjectAltName *PayloadContentSubjectAltName `json:"SubjectAltName,omitempty" plist:"SubjectAltName,omitempty"`
	// If `false`, the system disables exporting the private key from the keychain.
	KeyIsExtractable *bool `json:"KeyIsExtractable,omitempty" plist:"KeyIsExtractable,omitempty"`
	// If `true`, all apps have access to the private key.
	AllowAllAppsAccess *bool `json:"AllowAllAppsAccess,omitempty" plist:"AllowAllAppsAccess,omitempty"`
}

// The key size, in bits.
type Keysize int64

const (
	Keysize1024 Keysize = 1024
	Keysize2048 Keysize = 2048
	Keysize4096 Keysize = 4096
)

// The SCEP payload can specify an optional `SubjectAltName` dictionary that provides values required by the CA for issuing a certificate. You can specify a single string or an array of strings for each key. The values you specify depend on the CA you're using, but might include DNS name, URL, or email values. For an example, see Sample Configuration Profile or Over-the-Air Profile Delivery and Configuration.
type PayloadContentSubjectAltName struct {
	// The RFC 822 (email address) string.
	Rfc822Name *string `json:"rfc822Name,omitempty" plist:"rfc822Name,omitempty"`
	// The DNS name.
	DNSName *string `json:"dNSName,omitempty" plist:"dNSName,omitempty"`
	// The Uniform Resource Identifier.
	UniformResourceIdentifier *string `json:"uniformResourceIdentifier,omitempty" plist:"uniformResourceIdentifier,omitempty"`
	// The NT principal name. Use an other name OID set to `1.3.6.1.4.1.311.20.2.3`.
	NtPrincipalName *string `json:"ntPrincipalName,omitempty" plist:"ntPrincipalName,omitempty"`
}

// The payload that configures a smart card.
// Restrictions and settings for smart card pairing on macOS
type SmartCard struct {
	*CommonPayloadKeys
	// If `false`, users don't get the pairing dialog, although existing pairings still work.
	UserPairing *bool `json:"UserPairing,omitempty" plist:"UserPairing,omitempty"`
	// If `false`, the system disables smart cards for logins, authorizations, and screen saver unlocking. It is still allowed for other functions, such as signing emails and accessing the web. A restart is required for a setting change to take effect.
	AllowSmartCard *bool `json:"allowSmartCard,omitempty" plist:"allowSmartCard,omitempty"`
	// Configures the certificate trust check and has one of the following possible values:
	// - `0`: Turns off certificate trust check.
	// - `1`: Turns on certificate trust check. A standard validity check is performed but doesn't include additional revocation checks.
	// - `2`: Turns on certificate trust check. A soft revocation check is also performed. Until the certificate is explicitly rejected by CRL/OCSP, it's considered valid. This setting means that unavailable or unreachable CRL/OCSP allow this check to succeed.
	// - `3`: Turns on certificate trust check. A hard revocation check is also performed. Unless CRL/OCSP explicitly says "This certificate is OK," it's considered invalid. This option is the most secure.
	CheckCertificateTrust *CheckCertificateTrust `default:"0" json:"checkCertificateTrust,omitempty" plist:"checkCertificateTrust,omitempty"`
	// If `true`, a user can pair with only one smart card, although existing pairings are allowed if already set up.
	OneCardPerUser *bool `json:"oneCardPerUser,omitempty" plist:"oneCardPerUser,omitempty"`
	// If `1`, the system enables the screen saver when the smart card is removed. Available in macOS 10.13.4 and later.
	TokenRemovalAction *TokenRemovalAction `default:"0" json:"tokenRemovalAction,omitempty" plist:"tokenRemovalAction,omitempty"`
	// If `true`, a user can only log in or authenticate with a smart card. Available in macOS 10.13.2 and later.
	EnforceSmartCard *bool `json:"enforceSmartCard,omitempty" plist:"enforceSmartCard,omitempty"`
}

func (p *SmartCard) PayloadType() string {
	return "com.apple.security.smartcard"
}

// Configures the certificate trust check and has one of the following possible values:
// - `0`: Turns off certificate trust check.
// - `1`: Turns on certificate trust check. A standard validity check is performed but doesn't include additional revocation checks.
// - `2`: Turns on certificate trust check. A soft revocation check is also performed. Until the certificate is explicitly rejected by CRL/OCSP, it's considered valid. This setting means that unavailable or unreachable CRL/OCSP allow this check to succeed.
// - `3`: Turns on certificate trust check. A hard revocation check is also performed. Unless CRL/OCSP explicitly says "This certificate is OK," it's considered invalid. This option is the most secure.
type CheckCertificateTrust int64

const (
	CheckCertificateTrust0 CheckCertificateTrust = 0
	CheckCertificateTrust1 CheckCertificateTrust = 1
	CheckCertificateTrust2 CheckCertificateTrust = 2
	CheckCertificateTrust3 CheckCertificateTrust = 3
)

// If `1`, the system enables the screen saver when the smart card is removed. Available in macOS 10.13.4 and later.
type TokenRemovalAction int64

const (
	TokenRemovalAction0 TokenRemovalAction = 0
	TokenRemovalAction1 TokenRemovalAction = 1
)

// This payload that configures managed login items, which auto-enables and auto-allows matched items.
// This payload defines rules for tagging login items as managed, which will auto-enable and auto-allow matched items.
type ServiceManagementManagedLoginItems struct {
	*CommonPayloadKeys
	// An array of service management rules.
	Rules []Rule `json:"Rules" plist:"Rules" required:"true"`
}

func (p *ServiceManagementManagedLoginItems) PayloadType() string {
	return "com.apple.servicemanagement"
}

// A specification for matching one or more login items.
type Rule struct {
	// The type of comparison to make.
	RuleType RuleType `json:"RuleType" plist:"RuleType" required:"true"`
	// The value to compare with each login item's value, to determine if this rule is a match.
	RuleValue string `json:"RuleValue" plist:"RuleValue" required:"true"`
	// An optional description of the rule.
	Comment *string `json:"Comment,omitempty" plist:"Comment,omitempty"`
	// An additional constraint to limit the scope of the rule that the system tests after matching the `RuleType` and `RuleValue`.
	TeamIdentifier *string `json:"TeamIdentifier,omitempty" plist:"TeamIdentifier,omitempty"`
}

// The type of comparison to make.
type RuleType string

const (
	RuleTypeBundleIdentifier       RuleType = "BundleIdentifier"
	RuleTypeBundleIdentifierPrefix RuleType = "BundleIdentifierPrefix"
	RuleTypeLabel                  RuleType = "Label"
	RuleTypeLabelPrefix            RuleType = "LabelPrefix"
	RuleTypeTeamIdentifier         RuleType = "TeamIdentifier"
)

// The payload that configures a Lock Screen message.
// Allows admins to specify optional text displayed on the Login Window and Lock Screen (i.e. a footnote and Asset Tag Information).
type LockScreenMessage struct {
	*CommonPayloadKeys
	// The asset tag information for the device, displayed in the Login Window and Lock Screen.
	AssetTagInformation *string `json:"AssetTagInformation,omitempty" plist:"AssetTagInformation,omitempty"`
	// Deprecated. Use `LockScreenFootnote` instead.
	IfLostReturnToMessage *string `json:"IfLostReturnToMessage,omitempty" plist:"IfLostReturnToMessage,omitempty"`
	// The footnote displayed in the Login Window and Lock Screen.
	LockScreenFootnote *string `json:"LockScreenFootnote,omitempty" plist:"LockScreenFootnote,omitempty"`
}

func (p *LockScreenMessage) PayloadType() string {
	return "com.apple.shareddeviceconfiguration"
}

// The payload that configures single sign-on (SSO).
type SingleSignOn struct {
	*CommonPayloadKeys
	// The human-readable name for the account.
	Name string `json:"Name" plist:"Name" required:"true"`
	// The Kerberos dictionary.
	Kerberos *Kerberos `json:"Kerberos,omitempty" plist:"Kerberos,omitempty"`
}

func (p *SingleSignOn) PayloadType() string {
	return "com.apple.sso"
}

// The Kerberos dictionary.
type Kerberos struct {
	// The principal name. If not provided, the system prompts the user for one during profile installation. Required for MDM installation.
	PrincipalName *string `json:"PrincipalName,omitempty" plist:"PrincipalName,omitempty"`
	// The `PayloadUUID` of an identity certificate payload that the system can use to renew the Kerberos credential without user interaction. Set the payload type to either `com.apple.security.pkcs12` or `com.apple.security.scep` in the certificate payload. The configuration file needs to contain both the SSO payload and the identity certificate payload.
	PayloadCertificateUUID *string `json:"PayloadCertificateUUID,omitempty" plist:"PayloadCertificateUUID,omitempty"`
	// The properly capitalized realm name.
	Realm string `json:"Realm" plist:"Realm" required:"true"`
	// The list of URL prefixes to match in order to use this account for Kerberos authentication over HTTP. If this key is missing, the system makes the account eligible to match all `http://` and `https://` URLs.
	// Begin the URL matching patterns with either `http://` or `https://`. The system performs a simple string match, so the URL prefix `http://www.apple.com/` doesn't match `http://www.apple.com:80/`. However, if a matching pattern doesn't end in `/`, the system automatically append a `/` to it.
	URLPrefixMatches *[]string `json:"URLPrefixMatches,omitempty" plist:"URLPrefixMatches,omitempty"`
	// The list of app identifiers that the system allows to use this login. If this field missing, the system matches all app identifiers with this login.
	// Don't set an empty array. The array needs to contain strings that match App Bundle IDs. These strings can be exact matches such as `com.mycompany.myapp`, or they may specify a prefix match on the Bundle ID by using the `*` wildcard character. The wildcard character needs to appear after a period (`.`), and may only appear once, at the end of the string, for example, `com.mycompany.*`. When you provide a wildcard, the system grants access to the account to any app with a Bundle ID that begins with the prefix.
	AppIdentifierMatches *[]string `json:"AppIdentifierMatches,omitempty" plist:"AppIdentifierMatches,omitempty"`
}

// The payload that configures subscribed calendars.
type SubscribedCalendars struct {
	*CommonPayloadKeys
	// The description of the account.
	SubCalAccountDescription *string `json:"SubCalAccountDescription,omitempty" plist:"SubCalAccountDescription,omitempty"`
	// The server's address.
	SubCalAccountHostName string `json:"SubCalAccountHostName" plist:"SubCalAccountHostName" required:"true"`
	// The user's user name.
	SubCalAccountUsername *string `json:"SubCalAccountUsername,omitempty" plist:"SubCalAccountUsername,omitempty"`
	// The user's password.
	SubCalAccountPassword *string `json:"SubCalAccountPassword,omitempty" plist:"SubCalAccountPassword,omitempty"`
	// If `true`, the system enables SSL.
	SubCalAccountUseSSL *bool `json:"SubCalAccountUseSSL,omitempty" plist:"SubCalAccountUseSSL,omitempty"`
	// The VPNUUID of the per-app VPN the account uses for network communication. Available in iOS 14 and later.
	VPNUUID *string `json:"VPNUUID,omitempty" plist:"VPNUUID,omitempty"`
}

func (p *SubscribedCalendars) PayloadType() string {
	return "com.apple.subscribedcalendar.account"
}

// The payload that configures the kernel extension policies.
// Provides a way of enabling a set of team identifiers or specific kernel extensions for loading without user approval. Also provides a way to block users from approving additional kernel extensions. Payload must be user-approved only.
type SystemPolicyKernelExtensions struct {
	*CommonPayloadKeys
	// If `true`, nonadministrative users can approve additional kernel extensions in the Security & Privacy preferences.
	// Available in macOS 11 and later.
	AllowNonAdminUserApprovals *bool `json:"AllowNonAdminUserApprovals,omitempty" plist:"AllowNonAdminUserApprovals,omitempty"`
	// If `true`, users can approve additional kernel extensions that configuration profiles don't explicitly allow.
	AllowUserOverrides *bool `json:"AllowUserOverrides,omitempty" plist:"AllowUserOverrides,omitempty"`
	// The array of team identifiers that define which validly signed kernel extensions can load.
	AllowedTeamIdentifiers *[]string `json:"AllowedTeamIdentifiers,omitempty" plist:"AllowedTeamIdentifiers,omitempty"`
	// The dictionary that represents a set of kernel extensions that the system always allows to load on the computer. The dictionary maps team identifiers (keys) to arrays of bundle identifiers.
	AllowedKernelExtensions *map[string][]string `json:"AllowedKernelExtensions,omitempty" plist:"AllowedKernelExtensions,omitempty"`
}

func (p *SystemPolicyKernelExtensions) PayloadType() string {
	return "com.apple.syspolicy.kernel-extension-policy"
}

// The payload that configures system extensions.
// Provides a way of enabling a set of team identifiers or specific system extensions for loading without user approval. Also provides a way to block users from approving additional system extensions. Payload must be user-approved only. Starting in macOS 11.3, installing or removing this payload can change the state of system extensions on the machine. If a system extension has been activated by its containing application but is still in a pending state, installing a payload which specifies that extension is Allowed will complete the activation process. If a system extension is active, removing a payload which specified that extension was Allowed will deactivate the extension.
type SystemExtensions struct {
	*CommonPayloadKeys
	// If `false`, restricts users from approving additional system extensions that configuration profiles don't explicitly allow.
	AllowUserOverrides *bool `json:"AllowUserOverrides,omitempty" plist:"AllowUserOverrides,omitempty"`
	// An array of team identifiers that defines valid, signed system extensions that are allowable to load. Approved system extensions are those signed with any of the specified team identifiers.
	// To avoid requiring an administrator to authorize the operation, you can activate system extensions that this key specifies using `OSSystemExtensionActivationRequest API`.
	// It's an error for the same team identifier to appear in both this array and as a key in the `AllowedSystemExtensions` dictionary.
	AllowedTeamIdentifiers *[]string `json:"AllowedTeamIdentifiers,omitempty" plist:"AllowedTeamIdentifiers,omitempty"`
	// A dictionary that maps a team identifier to an array of strings, where each string is a type of system extension that you can install for that team identifier. The allowed extension types are `DriverExtension`, `NetworkExtension`, and `EndpointSecurityExtension`.
	// If there's no entry for a specified team identifier in the dictionary, the system allows all extension types.
	AllowedSystemExtensionTypes *map[string][]string `json:"AllowedSystemExtensionTypes,omitempty" plist:"AllowedSystemExtensionTypes,omitempty"`
	// A dictionary of approved system extensions on the computer. The dictionary maps the team identifiers (keys) to arrays of bundle identifiers, where the bundle identifier defines the system extension to install.
	// To avoid requiring an administrator to authorize the operation, you can activate system extensions that this key specifies using `OSSystemExtensionActivationRequest API`.
	// It's an error for the same team identifier to appear in both the `AllowedTeamIdentifiers` array and as a key in this dictionary.
	AllowedSystemExtensions *map[string][]string `json:"AllowedSystemExtensions,omitempty" plist:"AllowedSystemExtensions,omitempty"`
	// A dictionary of system extensions that are allowed to remove themselves from the machine. The dictionary maps team identifiers (keys) to arrays of bundle identifiers, where the bundle identifier defines the system extension. An application using the `OSSystemExtensionDeactivationRequest` API can deactivate the specified system extensions without requiring an administrator to authorize the operation.
	// Available in macOS 12 and later.
	RemovableSystemExtensions *map[string][]string `json:"RemovableSystemExtensions,omitempty" plist:"RemovableSystemExtensions,omitempty"`
	// A dictionary of system extensions on the computer. The dictionary maps the team identifiers (keys) to arrays of bundle identifiers, where the bundle identifier defines the system extension which can't be disabled or uninstalled when SIP is enabled. It's an error for the same mapping to appear in the dictionary values corresponding to `RemovableSystemExtensions` and `NonRemovableSystemExtensions` keys.
	NonRemovableSystemExtensions *map[string][]string `json:"NonRemovableSystemExtensions,omitempty" plist:"NonRemovableSystemExtensions,omitempty"`
	// A dictionary of system extensions on the computer. The dictionary maps the team identifiers (keys) to arrays of bundle identifiers, where the bundle identifier defines the system extension which can't be disabled or uninstalled from System Settings or Finder. The set of system extensions between `RemovableSystemExtensions` and `NonRemovableFromUISystemExtensions` can to overlap.
	NonRemovableFromUISystemExtensions *map[string][]string `json:"NonRemovableFromUISystemExtensions,omitempty" plist:"NonRemovableFromUISystemExtensions,omitempty"`
}

func (p *SystemExtensions) PayloadType() string {
	return "com.apple.system-extension-policy"
}

// The payload that configures system logging.
type SystemLogging struct {
	*CommonPayloadKeys
	// Not to be used.
	Processes *map[string]any `json:"Processes,omitempty" plist:"Processes,omitempty"`
	// A dictionary enabling the logging level for subsystems. See `Customizing Logging Behavior While Debugging` for more details about the format of the dictionary.
	Subsystems *map[string]any `json:"Subsystems,omitempty" plist:"Subsystems,omitempty"`
	// This dictionary has one key, `Enable-Private-Data`. Setting that value to `true` enables private data logging for the entire system.
	System *map[string]any `json:"System,omitempty" plist:"System,omitempty"`
}

func (p *SystemLogging) PayloadType() string {
	return "com.apple.system.logging"
}

// The payload that configures system migration.
// Provides a way of customizing items migrated during System Migration.
type SystemMigration struct {
	*CommonPayloadKeys
	// The list of custom behavior dictionaries.
	CustomBehavior *[]*CustomBehaviorItem `json:"CustomBehavior,omitempty" plist:"CustomBehavior,omitempty"`
}

func (p *SystemMigration) PayloadType() string {
	return "com.apple.systemmigration"
}

// The custom behavior dictionary.
type CustomBehaviorItem struct {
	// The context that custom paths apply to.
	Context string `json:"Context" plist:"Context" required:"true"`
	// The list of custom behavior path dictionaries.
	Paths []*PathsItem `json:"Paths" plist:"Paths" required:"true"`
}

// The custom behavior path dictionary.
type PathsItem struct {
	// The path to the migrating file or directory on the source system.
	SourcePath string `json:"SourcePath" plist:"SourcePath" required:"true"`
	// If `true`, the source path is located within a user home directory.
	SourcePathInUserHome bool `json:"SourcePathInUserHome" plist:"SourcePathInUserHome" required:"true"`
	// The path to the destination file or directory on the target system.
	TargetPath string `json:"TargetPath" plist:"TargetPath" required:"true"`
	// If `true`, the target path is located within a user home directory.
	TargetPathInUserHome bool `json:"TargetPathInUserHome" plist:"TargetPathInUserHome" required:"true"`
}

// The payload that configures the system policy for assessments.
// Provides a way of enabling System Policy assessment processing. This corresponds to the Gatekeeper UI in the Security pref pane.
type SystemPolicyControl struct {
	*CommonPayloadKeys
	// If `true`, enables Gatekeeper. If `false`, disables Gatekeeper.
	EnableAssessment *bool `json:"EnableAssessment,omitempty" plist:"EnableAssessment,omitempty"`
	// If `true`, enables Gatekeeper's "Mac App Store and identified developers" option.
	// If `false`, enables Gatekeeper's "Mac App Store" option.
	// If the value of `EnableAssessment` isn't set to `true`, this key has no effect.
	AllowIdentifiedDevelopers *bool `json:"AllowIdentifiedDevelopers,omitempty" plist:"AllowIdentifiedDevelopers,omitempty"`
	// If `false`, prevents Gatekeeper from prompting the user to upload blocked malware to Apple for purposes of improving malware detection.
	EnableXProtectMalwareUpload *bool `json:"EnableXProtectMalwareUpload,omitempty" plist:"EnableXProtectMalwareUpload,omitempty"`
}

func (p *SystemPolicyControl) PayloadType() string {
	return "com.apple.systempolicy.control"
}

// The payload that configures the Finder's contextual menu to bypass the system policy.
// Provides a way of disabling the Finder's contextual menu that allows bypass of System Policy restrictions.
type SystemPolicyManaged struct {
	*CommonPayloadKeys
	// If `true`, disables the Finder's contextual menu item.
	DisableOverride *bool `json:"DisableOverride,omitempty" plist:"DisableOverride,omitempty"`
}

func (p *SystemPolicyManaged) PayloadType() string {
	return "com.apple.systempolicy.managed"
}

// The payload that configures the system policy.
// This payload allows control over Gatekeeper's system policy rules. The keys and functionality are tightly related to the spctl command line tool. For more information, see the manual page for spctl.
type SystemPolicyRule struct {
	*CommonPayloadKeys
	// The policy requirement. This key must follow the syntax described in [Code Signing Requirement Language](https://developer.apple.com/library/archive/documentation/Security/Conceptual/CodeSigningGuide/RequirementLang/RequirementLang.html#//apple_ref/doc/uid/TP40005929-CH5).
	Requirement *string `json:"Requirement,omitempty" plist:"Requirement,omitempty"`
	// This string appears in the System Policy UI. If it's missing, `PayloadDisplayName` or `PayloadDescription` is entered into this field before the rule is added to the System Policy database.
	Comment *string `json:"Comment,omitempty" plist:"Comment,omitempty"`
	// The rule's priority.
	Priority *float64 `json:"Priority,omitempty" plist:"Priority,omitempty"`
	// The expiration date for rules being processed.
	Expiration *time.Time `json:"Expiration,omitempty" plist:"Expiration,omitempty"`
	// The type of operation.
	OperationType *OperationType `default:"operation:execute" json:"OperationType,omitempty" plist:"OperationType,omitempty"`
	// The single leaf certificate for the app that is in the allow list.
	LeafCertificate *[]byte `json:"LeafCertificate,omitempty" plist:"LeafCertificate,omitempty"`
}

func (p *SystemPolicyRule) PayloadType() string {
	return "com.apple.systempolicy.rule"
}

// The type of operation.
type OperationType string

const (
	OperationTypeOperationExecute OperationType = "operation:execute"
	OperationTypeOperationInstall OperationType = "operation:install"
	OperationTypeOperationLsopen  OperationType = "operation:lsopen"
)

// The payload that configures the preference panes.
// Hide and show individual System Preferences panes.
// The following preference pane items are no longer supported on macOS 10.14:
// • com.apple.preferences.appstore
// The following preference pane items are no longer supported on macOS 10.15:
// • com.apple.preference.ink
// • com.apple.preferences.icloud
// • com.apple.preferences.parentalcontrols
// This payload is deprecated as of macOS 13. When new restrictions become available to control functionality exposed through System Settings, those restrictions should be used instead of disabling the entire preference pane. This allows for user transparency even when the user's control has been disabled by a restriction.
// macOS 13 introduces a new DisabledSystemSettings key for controlling macOS 13 and newer System Settings extensions. However, note that System Settings extensions within the Privacy & Security section cannot be disabled. If DisabledSystemSettings is not provided, the system will attempt to honor the EnabledPreferencePanes and DisabledPreferencePanes by mapping the old preference pane value to one or more new settings extension values whose content was wholly contained in the old preference pane value.
type SystemPreferences struct {
	*CommonPayloadKeys
	// The list of enabled System Preferences panes.
	EnabledPreferencePanes *[]EnabledPreferencePanes `json:"EnabledPreferencePanes,omitempty" plist:"EnabledPreferencePanes,omitempty"`
	// The list of disabled System Preferences panes.
	DisabledPreferencePanes *[]DisabledPreferencePanes `json:"DisabledPreferencePanes,omitempty" plist:"DisabledPreferencePanes,omitempty"`
	// The list of disabled System Settings extensions. All other items will be enabled. When `DisabledSystemSettings` is specified, the device ignores `DisabledPreferencePanes` and `EnabledPreferencePanes`.
	// > Note:
	// > A given System Settings extension can supply more than one section in System Settings; disabling such an extension disables all sections it supplies.
	DisabledSystemSettings *[]DisabledSystemSettings `json:"DisabledSystemSettings,omitempty" plist:"DisabledSystemSettings,omitempty"`
}

func (p *SystemPreferences) PayloadType() string {
	return "com.apple.systempreferences"
}

// The list of enabled System Preferences panes.
type EnabledPreferencePanes string

const (
	EnabledPreferencePanesComappleClassroomSettings                EnabledPreferencePanes = "com.apple.ClassroomSettings"
	EnabledPreferencePanesComappleLocalization                     EnabledPreferencePanes = "com.apple.Localization"
	EnabledPreferencePanesComapplepreferencedatetime               EnabledPreferencePanes = "com.apple.preference.datetime"
	EnabledPreferencePanesComapplepreferencedesktopscreeneffect    EnabledPreferencePanes = "com.apple.preference.desktopscreeneffect"
	EnabledPreferencePanesComapplepreferencedigihubdiscs           EnabledPreferencePanes = "com.apple.preference.digihub.discs"
	EnabledPreferencePanesComapplepreferencedisplays               EnabledPreferencePanes = "com.apple.preference.displays"
	EnabledPreferencePanesComapplepreferencedock                   EnabledPreferencePanes = "com.apple.preference.dock"
	EnabledPreferencePanesComapplepreferenceenergysaver            EnabledPreferencePanes = "com.apple.preference.energysaver"
	EnabledPreferencePanesComapplepreferenceexpose                 EnabledPreferencePanes = "com.apple.preference.expose"
	EnabledPreferencePanesComapplepreferencegeneral                EnabledPreferencePanes = "com.apple.preference.general"
	EnabledPreferencePanesComapplepreferenceink                    EnabledPreferencePanes = "com.apple.preference.ink"
	EnabledPreferencePanesComapplepreferencekeyboard               EnabledPreferencePanes = "com.apple.preference.keyboard"
	EnabledPreferencePanesComapplepreferencemouse                  EnabledPreferencePanes = "com.apple.preference.mouse"
	EnabledPreferencePanesComapplepreferencenetwork                EnabledPreferencePanes = "com.apple.preference.network"
	EnabledPreferencePanesComapplepreferencenotifications          EnabledPreferencePanes = "com.apple.preference.notifications"
	EnabledPreferencePanesComapplepreferenceprintfax               EnabledPreferencePanes = "com.apple.preference.printfax"
	EnabledPreferencePanesComapplepreferencescreentime             EnabledPreferencePanes = "com.apple.preference.screentime"
	EnabledPreferencePanesComapplepreferencesecurity               EnabledPreferencePanes = "com.apple.preference.security"
	EnabledPreferencePanesComapplepreferencesidecar                EnabledPreferencePanes = "com.apple.preference.sidecar"
	EnabledPreferencePanesComapplepreferencesound                  EnabledPreferencePanes = "com.apple.preference.sound"
	EnabledPreferencePanesComapplepreferencespeech                 EnabledPreferencePanes = "com.apple.preference.speech"
	EnabledPreferencePanesComapplepreferencespotlight              EnabledPreferencePanes = "com.apple.preference.spotlight"
	EnabledPreferencePanesComapplepreferencestartupdisk            EnabledPreferencePanes = "com.apple.preference.startupdisk"
	EnabledPreferencePanesComapplepreferencetrackpad               EnabledPreferencePanes = "com.apple.preference.trackpad"
	EnabledPreferencePanesComapplepreferenceuniversalaccess        EnabledPreferencePanes = "com.apple.preference.universalaccess"
	EnabledPreferencePanesComapplepreferencesAppleIDPrefPane       EnabledPreferencePanes = "com.apple.preferences.AppleIDPrefPane"
	EnabledPreferencePanesComapplepreferencesappstore              EnabledPreferencePanes = "com.apple.preferences.appstore"
	EnabledPreferencePanesComapplepreferencesBluetooth             EnabledPreferencePanes = "com.apple.preferences.Bluetooth"
	EnabledPreferencePanesComapplepreferencesconfigurationprofiles EnabledPreferencePanes = "com.apple.preferences.configurationprofiles"
	EnabledPreferencePanesComapplepreferencesextensions            EnabledPreferencePanes = "com.apple.preferences.extensions"
	EnabledPreferencePanesComapplepreferencesFamilySharingPrefPane EnabledPreferencePanes = "com.apple.preferences.FamilySharingPrefPane"
	EnabledPreferencePanesComapplepreferencesicloud                EnabledPreferencePanes = "com.apple.preferences.icloud"
	EnabledPreferencePanesComapplepreferencesinternetaccounts      EnabledPreferencePanes = "com.apple.preferences.internetaccounts"
	EnabledPreferencePanesComapplepreferencesparentalcontrols      EnabledPreferencePanes = "com.apple.preferences.parentalcontrols"
	EnabledPreferencePanesComapplepreferencespassword              EnabledPreferencePanes = "com.apple.preferences.password"
	EnabledPreferencePanesComapplepreferencessharing               EnabledPreferencePanes = "com.apple.preferences.sharing"
	EnabledPreferencePanesComapplepreferencessoftwareupdate        EnabledPreferencePanes = "com.apple.preferences.softwareupdate"
	EnabledPreferencePanesComapplepreferencesusers                 EnabledPreferencePanes = "com.apple.preferences.users"
	EnabledPreferencePanesComapplepreferenceswallet                EnabledPreferencePanes = "com.apple.preferences.wallet"
	EnabledPreferencePanesComappleprefpanelfibrechannel            EnabledPreferencePanes = "com.apple.prefpanel.fibrechannel"
	EnabledPreferencePanesComappleprefsbackup                      EnabledPreferencePanes = "com.apple.prefs.backup"
	EnabledPreferencePanesComappleXsan                             EnabledPreferencePanes = "com.apple.Xsan"
)

// The list of disabled System Preferences panes.
type DisabledPreferencePanes string

const (
	DisabledPreferencePanesComappleClassroomSettings                DisabledPreferencePanes = "com.apple.ClassroomSettings"
	DisabledPreferencePanesComappleLocalization                     DisabledPreferencePanes = "com.apple.Localization"
	DisabledPreferencePanesComapplepreferencedatetime               DisabledPreferencePanes = "com.apple.preference.datetime"
	DisabledPreferencePanesComapplepreferencedesktopscreeneffect    DisabledPreferencePanes = "com.apple.preference.desktopscreeneffect"
	DisabledPreferencePanesComapplepreferencedigihubdiscs           DisabledPreferencePanes = "com.apple.preference.digihub.discs"
	DisabledPreferencePanesComapplepreferencedisplays               DisabledPreferencePanes = "com.apple.preference.displays"
	DisabledPreferencePanesComapplepreferencedock                   DisabledPreferencePanes = "com.apple.preference.dock"
	DisabledPreferencePanesComapplepreferenceenergysaver            DisabledPreferencePanes = "com.apple.preference.energysaver"
	DisabledPreferencePanesComapplepreferenceexpose                 DisabledPreferencePanes = "com.apple.preference.expose"
	DisabledPreferencePanesComapplepreferencegeneral                DisabledPreferencePanes = "com.apple.preference.general"
	DisabledPreferencePanesComapplepreferenceink                    DisabledPreferencePanes = "com.apple.preference.ink"
	DisabledPreferencePanesComapplepreferencekeyboard               DisabledPreferencePanes = "com.apple.preference.keyboard"
	DisabledPreferencePanesComapplepreferencemouse                  DisabledPreferencePanes = "com.apple.preference.mouse"
	DisabledPreferencePanesComapplepreferencenetwork                DisabledPreferencePanes = "com.apple.preference.network"
	DisabledPreferencePanesComapplepreferencenotifications          DisabledPreferencePanes = "com.apple.preference.notifications"
	DisabledPreferencePanesComapplepreferenceprintfax               DisabledPreferencePanes = "com.apple.preference.printfax"
	DisabledPreferencePanesComapplepreferencescreentime             DisabledPreferencePanes = "com.apple.preference.screentime"
	DisabledPreferencePanesComapplepreferencesecurity               DisabledPreferencePanes = "com.apple.preference.security"
	DisabledPreferencePanesComapplepreferencesidecar                DisabledPreferencePanes = "com.apple.preference.sidecar"
	DisabledPreferencePanesComapplepreferencesound                  DisabledPreferencePanes = "com.apple.preference.sound"
	DisabledPreferencePanesComapplepreferencespeech                 DisabledPreferencePanes = "com.apple.preference.speech"
	DisabledPreferencePanesComapplepreferencespotlight              DisabledPreferencePanes = "com.apple.preference.spotlight"
	DisabledPreferencePanesComapplepreferencestartupdisk            DisabledPreferencePanes = "com.apple.preference.startupdisk"
	DisabledPreferencePanesComapplepreferencetrackpad               DisabledPreferencePanes = "com.apple.preference.trackpad"
	DisabledPreferencePanesComapplepreferenceuniversalaccess        DisabledPreferencePanes = "com.apple.preference.universalaccess"
	DisabledPreferencePanesComapplepreferencesAppleIDPrefPane       DisabledPreferencePanes = "com.apple.preferences.AppleIDPrefPane"
	DisabledPreferencePanesComapplepreferencesappstore              DisabledPreferencePanes = "com.apple.preferences.appstore"
	DisabledPreferencePanesComapplepreferencesBluetooth             DisabledPreferencePanes = "com.apple.preferences.Bluetooth"
	DisabledPreferencePanesComapplepreferencesconfigurationprofiles DisabledPreferencePanes = "com.apple.preferences.configurationprofiles"
	DisabledPreferencePanesComapplepreferencesextensions            DisabledPreferencePanes = "com.apple.preferences.extensions"
	DisabledPreferencePanesComapplepreferencesFamilySharingPrefPane DisabledPreferencePanes = "com.apple.preferences.FamilySharingPrefPane"
	DisabledPreferencePanesComapplepreferencesicloud                DisabledPreferencePanes = "com.apple.preferences.icloud"
	DisabledPreferencePanesComapplepreferencesinternetaccounts      DisabledPreferencePanes = "com.apple.preferences.internetaccounts"
	DisabledPreferencePanesComapplepreferencesparentalcontrols      DisabledPreferencePanes = "com.apple.preferences.parentalcontrols"
	DisabledPreferencePanesComapplepreferencespassword              DisabledPreferencePanes = "com.apple.preferences.password"
	DisabledPreferencePanesComapplepreferencessharing               DisabledPreferencePanes = "com.apple.preferences.sharing"
	DisabledPreferencePanesComapplepreferencessoftwareupdate        DisabledPreferencePanes = "com.apple.preferences.softwareupdate"
	DisabledPreferencePanesComapplepreferencesusers                 DisabledPreferencePanes = "com.apple.preferences.users"
	DisabledPreferencePanesComapplepreferenceswallet                DisabledPreferencePanes = "com.apple.preferences.wallet"
	DisabledPreferencePanesComappleprefpanelfibrechannel            DisabledPreferencePanes = "com.apple.prefpanel.fibrechannel"
	DisabledPreferencePanesComappleprefsbackup                      DisabledPreferencePanes = "com.apple.prefs.backup"
	DisabledPreferencePanesComappleXsan                             DisabledPreferencePanes = "com.apple.Xsan"
)

// The list of disabled System Settings extensions. All other items will be enabled. When `DisabledSystemSettings` is specified, the device ignores `DisabledPreferencePanes` and `EnabledPreferencePanes`.
// > Note:
// > A given System Settings extension can supply more than one section in System Settings; disabling such an extension disables all sections it supplies.
type DisabledSystemSettings string

const (
	DisabledSystemSettingsComappleAccessibilitySettingsextension                  DisabledSystemSettings = "com.apple.Accessibility-Settings.extension"
	DisabledSystemSettingsComappleAirDropHandoffSettingsextension                 DisabledSystemSettings = "com.apple.AirDrop-Handoff-Settings.extension"
	DisabledSystemSettingsComappleBatterySettingsextension                        DisabledSystemSettings = "com.apple.Battery-Settings.extension"
	DisabledSystemSettingsComappleBluetoothSettings                               DisabledSystemSettings = "com.apple.BluetoothSettings"
	DisabledSystemSettingsComappleCDDVDSettingsextension                          DisabledSystemSettings = "com.apple.CD-DVD-Settings.extension"
	DisabledSystemSettingsComappleClassKitSettingsextension                       DisabledSystemSettings = "com.apple.ClassKit-Settings.extension"
	DisabledSystemSettingsComappleClassroomSettingsextension                      DisabledSystemSettings = "com.apple.Classroom-Settings.extension"
	DisabledSystemSettingsComappleControlCenterSettingsextension                  DisabledSystemSettings = "com.apple.ControlCenter-Settings.extension"
	DisabledSystemSettingsComappleDateTimeSettingsextension                       DisabledSystemSettings = "com.apple.Date-Time-Settings.extension"
	DisabledSystemSettingsComappleDesktopSettingsextension                        DisabledSystemSettings = "com.apple.Desktop-Settings.extension"
	DisabledSystemSettingsComappleDisplaysSettingsextension                       DisabledSystemSettings = "com.apple.Displays-Settings.extension"
	DisabledSystemSettingsComappleExtensionsPreferences                           DisabledSystemSettings = "com.apple.ExtensionsPreferences"
	DisabledSystemSettingsComappleFamilySettingsextension                         DisabledSystemSettings = "com.apple.Family-Settings.extension"
	DisabledSystemSettingsComappleFocusSettingsextension                          DisabledSystemSettings = "com.apple.Focus-Settings.extension"
	DisabledSystemSettingsComappleGameCenterSettingsextension                     DisabledSystemSettings = "com.apple.Game-Center-Settings.extension"
	DisabledSystemSettingsComappleGameControllerSettingsextension                 DisabledSystemSettings = "com.apple.Game-Controller-Settings.extension"
	DisabledSystemSettingsComappleHeadphoneSettings                               DisabledSystemSettings = "com.apple.HeadphoneSettings"
	DisabledSystemSettingsComappleInternetAccountsSettingsextension               DisabledSystemSettings = "com.apple.Internet-Accounts-Settings.extension"
	DisabledSystemSettingsComappleKeyboardSettingsextension                       DisabledSystemSettings = "com.apple.Keyboard-Settings.extension"
	DisabledSystemSettingsComappleLocalizationSettingsextension                   DisabledSystemSettings = "com.apple.Localization-Settings.extension"
	DisabledSystemSettingsComappleLockScreenSettingsextension                     DisabledSystemSettings = "com.apple.Lock-Screen-Settings.extension"
	DisabledSystemSettingsComappleLoginItemsSettingsextension                     DisabledSystemSettings = "com.apple.LoginItems-Settings.extension"
	DisabledSystemSettingsComappleMouseSettingsextension                          DisabledSystemSettings = "com.apple.Mouse-Settings.extension"
	DisabledSystemSettingsComappleNetworkSettingsextension                        DisabledSystemSettings = "com.apple.Network-Settings.extension"
	DisabledSystemSettingsComappleNetworkExtensionSettingsUINESettingsUIExtension DisabledSystemSettings = "com.apple.NetworkExtensionSettingsUI.NESettingsUIExtension"
	DisabledSystemSettingsComappleNotificationsSettingsextension                  DisabledSystemSettings = "com.apple.Notifications-Settings.extension"
	DisabledSystemSettingsComapplePasswordsSettingsextension                      DisabledSystemSettings = "com.apple.Passwords-Settings.extension"
	DisabledSystemSettingsComapplePrintScanSettingsextension                      DisabledSystemSettings = "com.apple.Print-Scan-Settings.extension"
	DisabledSystemSettingsComappleScreenTimeSettingsextension                     DisabledSystemSettings = "com.apple.Screen-Time-Settings.extension"
	DisabledSystemSettingsComappleScreenSaverSettingsextension                    DisabledSystemSettings = "com.apple.ScreenSaver-Settings.extension"
	DisabledSystemSettingsComappleSharingSettingsextension                        DisabledSystemSettings = "com.apple.Sharing-Settings.extension"
	DisabledSystemSettingsComappleSiriSettingsextension                           DisabledSystemSettings = "com.apple.Siri-Settings.extension"
	DisabledSystemSettingsComappleSoftwareUpdateSettingsextension                 DisabledSystemSettings = "com.apple.Software-Update-Settings.extension"
	DisabledSystemSettingsComappleSoundSettingsextension                          DisabledSystemSettings = "com.apple.Sound-Settings.extension"
	DisabledSystemSettingsComappleStartupDiskSettingsextension                    DisabledSystemSettings = "com.apple.Startup-Disk-Settings.extension"
	DisabledSystemSettingsComappleTimeMachineSettingsextension                    DisabledSystemSettings = "com.apple.Time-Machine-Settings.extension"
	DisabledSystemSettingsComappleTouchIDSettingsextension                        DisabledSystemSettings = "com.apple.Touch-ID-Settings.extension"
	DisabledSystemSettingsComappleTrackpadSettingsextension                       DisabledSystemSettings = "com.apple.Trackpad-Settings.extension"
	DisabledSystemSettingsComappleTransferResetSettingsextension                  DisabledSystemSettings = "com.apple.Transfer-Reset-Settings.extension"
	DisabledSystemSettingsComappleUsersGroupsSettingsextension                    DisabledSystemSettings = "com.apple.Users-Groups-Settings.extension"
	DisabledSystemSettingsComappleWalletSettingsExtension                         DisabledSystemSettings = "com.apple.WalletSettingsExtension"
	DisabledSystemSettingsComappleWallpaperSettingsextension                      DisabledSystemSettings = "com.apple.Wallpaper-Settings.extension"
	DisabledSystemSettingsComapplesettingsStorage                                 DisabledSystemSettings = "com.apple.settings.Storage"
	DisabledSystemSettingsComapplesystempreferencesAppleIDSettings                DisabledSystemSettings = "com.apple.systempreferences.AppleIDSettings"
	DisabledSystemSettingsComapplewifiSettingsExtension                           DisabledSystemSettings = "com.apple.wifi-settings-extension"
)

// The payload that configures media management.
type MediaManagementAllowedMedia struct {
	*CommonPayloadKeys
	// The media type dictionary that defines volumes to eject when the user logs out.
	LogoutEject *LogoutEject `json:"logout-eject,omitempty" plist:"logout-eject,omitempty"`
	// The media type dictionary that controls volume mounting.
	MountControls *MountControls `json:"mount-controls,omitempty" plist:"mount-controls,omitempty"`
	// The media type dictionary that controls volume unmounting.
	UnmountControls *UnmountControls `json:"unmount-controls,omitempty" plist:"unmount-controls,omitempty"`
}

func (p *MediaManagementAllowedMedia) PayloadType() string {
	return "com.apple.systemuiserver"
}

// The media type dictionary that defines volumes to eject when the user logs out.
type LogoutEject struct {
	// Unused; set to an empty string.
	AllMedia *string `json:"all-media,omitempty" plist:"all-media,omitempty"`
	// A media action string or an array of media action strings.
	Cd *[]Cd `json:"cd,omitempty" plist:"cd,omitempty"`
	// A media action string or an array of media action strings.
	Dvd *[]Dvd `json:"dvd,omitempty" plist:"dvd,omitempty"`
	// A media action string or an array of media action strings.
	Bd *[]Bd `json:"bd,omitempty" plist:"bd,omitempty"`
	// A media action string or an array of media action strings.
	Blankcd *[]Blankcd `json:"blankcd,omitempty" plist:"blankcd,omitempty"`
	// A media action string or an array of media action strings.
	Blankdvd *[]Blankdvd `json:"blankdvd,omitempty" plist:"blankdvd,omitempty"`
	// A media action string or an array of media action strings.
	Blankbd *[]Blankbd `json:"blankbd,omitempty" plist:"blankbd,omitempty"`
	// A media action string or an array of media action strings.
	Dvdram *[]Dvdram `json:"dvdram,omitempty" plist:"dvdram,omitempty"`
	// A media action string or an array of media action strings.
	DiskImage *[]DiskImage `json:"disk-image,omitempty" plist:"disk-image,omitempty"`
	// A media action string or an array of media action strings.
	HarddiskInternal *[]HarddiskInternal `json:"harddisk-internal,omitempty" plist:"harddisk-internal,omitempty"`
	// A string or an array of media action strings. Internally installed SD cards and USB flash drives are included in the hard disk-external category.
	// This key is the default for media types that don't fall into other categories.
	HarddiskExternal *[]HarddiskExternal `json:"harddisk-external,omitempty" plist:"harddisk-external,omitempty"`
	// A media action string or an array of media action strings.
	Networkdisk *[]Networkdisk `json:"networkdisk,omitempty" plist:"networkdisk,omitempty"`
}

// A media action string or an array of media action strings.
// One of the following values:
// * authenticate - User will be authenticated before media is mounted
// * read-only - The media will be mounted read-only. Not valid for unmount-controls.
// * deny - The media will not be mounted.
// * eject - The media will not be mounted and it will be ejected if possible. Note that some volumes are not defined as ejectable, so using the deny key may be the best solution. Not valid for unmount-controls.
type Cd string

const (
	CdAuthenticate Cd = "authenticate"
	CdReadOnly     Cd = "read-only"
	CdDeny         Cd = "deny"
	CdEject        Cd = "eject"
)

// A media action string or an array of media action strings.
// One of the following values:
// * authenticate - User will be authenticated before media is mounted
// * read-only - The media will be mounted read-only. Not valid for unmount-controls.
// * deny - The media will not be mounted.
// * eject - The media will not be mounted and it will be ejected if possible. Note that some volumes are not defined as ejectable, so using the deny key may be the best solution. Not valid for unmount-controls.
type Dvd string

const (
	DvdAuthenticate Dvd = "authenticate"
	DvdReadOnly     Dvd = "read-only"
	DvdDeny         Dvd = "deny"
	DvdEject        Dvd = "eject"
)

// A media action string or an array of media action strings.
// One of the following values:
// * authenticate - User will be authenticated before media is mounted
// * read-only - The media will be mounted read-only. Not valid for unmount-controls.
// * deny - The media will not be mounted.
// * eject - The media will not be mounted and it will be ejected if possible. Note that some volumes are not defined as ejectable, so using the deny key may be the best solution. Not valid for unmount-controls.
type Bd string

const (
	BdAuthenticate Bd = "authenticate"
	BdReadOnly     Bd = "read-only"
	BdDeny         Bd = "deny"
	BdEject        Bd = "eject"
)

// A media action string or an array of media action strings.
// One of the following values:
// * authenticate - User will be authenticated before media is mounted
// * read-only - The media will be mounted read-only. Not valid for unmount-controls.
// * deny - The media will not be mounted.
// * eject - The media will not be mounted and it will be ejected if possible. Note that some volumes are not defined as ejectable, so using the deny key may be the best solution. Not valid for unmount-controls.
type Blankcd string

const (
	BlankcdAuthenticate Blankcd = "authenticate"
	BlankcdReadOnly     Blankcd = "read-only"
	BlankcdDeny         Blankcd = "deny"
	BlankcdEject        Blankcd = "eject"
)

// A media action string or an array of media action strings.
// One of the following values:
// * authenticate - User will be authenticated before media is mounted
// * read-only - The media will be mounted read-only. Not valid for unmount-controls.
// * deny - The media will not be mounted.
// * eject - The media will not be mounted and it will be ejected if possible. Note that some volumes are not defined as ejectable, so using the deny key may be the best solution. Not valid for unmount-controls.
type Blankdvd string

const (
	BlankdvdAuthenticate Blankdvd = "authenticate"
	BlankdvdReadOnly     Blankdvd = "read-only"
	BlankdvdDeny         Blankdvd = "deny"
	BlankdvdEject        Blankdvd = "eject"
)

// A media action string or an array of media action strings.
// One of the following values:
// * authenticate - User will be authenticated before media is mounted
// * read-only - The media will be mounted read-only. Not valid for unmount-controls.
// * deny - The media will not be mounted.
// * eject - The media will not be mounted and it will be ejected if possible. Note that some volumes are not defined as ejectable, so using the deny key may be the best solution. Not valid for unmount-controls.
type Blankbd string

const (
	BlankbdAuthenticate Blankbd = "authenticate"
	BlankbdReadOnly     Blankbd = "read-only"
	BlankbdDeny         Blankbd = "deny"
	BlankbdEject        Blankbd = "eject"
)

// A media action string or an array of media action strings.
// One of the following values:
// * authenticate - User will be authenticated before media is mounted
// * read-only - The media will be mounted read-only. Not valid for unmount-controls.
// * deny - The media will not be mounted.
// * eject - The media will not be mounted and it will be ejected if possible. Note that some volumes are not defined as ejectable, so using the deny key may be the best solution. Not valid for unmount-controls.
type Dvdram string

const (
	DvdramAuthenticate Dvdram = "authenticate"
	DvdramReadOnly     Dvdram = "read-only"
	DvdramDeny         Dvdram = "deny"
	DvdramEject        Dvdram = "eject"
)

// A media action string or an array of media action strings.
// One of the following values:
// * authenticate - User will be authenticated before media is mounted
// * read-only - The media will be mounted read-only. Not valid for unmount-controls.
// * deny - The media will not be mounted.
// * eject - The media will not be mounted and it will be ejected if possible. Note that some volumes are not defined as ejectable, so using the deny key may be the best solution. Not valid for unmount-controls.
type DiskImage string

const (
	DiskImageAuthenticate DiskImage = "authenticate"
	DiskImageReadOnly     DiskImage = "read-only"
	DiskImageDeny         DiskImage = "deny"
	DiskImageEject        DiskImage = "eject"
)

// A media action string or an array of media action strings.
// One of the following values:
// * authenticate - User will be authenticated before media is mounted
// * read-only - The media will be mounted read-only. Not valid for unmount-controls.
// * deny - The media will not be mounted.
// * eject - The media will not be mounted and it will be ejected if possible. Note that some volumes are not defined as ejectable, so using the deny key may be the best solution. Not valid for unmount-controls.
type HarddiskInternal string

const (
	HarddiskInternalAuthenticate HarddiskInternal = "authenticate"
	HarddiskInternalReadOnly     HarddiskInternal = "read-only"
	HarddiskInternalDeny         HarddiskInternal = "deny"
	HarddiskInternalEject        HarddiskInternal = "eject"
)

// A string or an array of media action strings. Internally installed SD cards and USB flash drives are included in the hard disk-external category.
// This key is the default for media types that don't fall into other categories.
// One of the following values:
// * authenticate - User will be authenticated before media is mounted
// * read-only - The media will be mounted read-only. Not valid for unmount-controls.
// * deny - The media will not be mounted.
// * eject - The media will not be mounted and it will be ejected if possible. Note that some volumes are not defined as ejectable, so using the deny key may be the best solution. Not valid for unmount-controls.
type HarddiskExternal string

const (
	HarddiskExternalAuthenticate HarddiskExternal = "authenticate"
	HarddiskExternalReadOnly     HarddiskExternal = "read-only"
	HarddiskExternalDeny         HarddiskExternal = "deny"
	HarddiskExternalEject        HarddiskExternal = "eject"
)

// A media action string or an array of media action strings.
// One of the following values:
// * authenticate - User will be authenticated before media is mounted
// * read-only - The media will be mounted read-only. Not valid for unmount-controls.
// * deny - The media will not be mounted.
// * eject - The media will not be mounted and it will be ejected if possible. Note that some volumes are not defined as ejectable, so using the deny key may be the best solution. Not valid for unmount-controls.
type Networkdisk string

const (
	NetworkdiskAuthenticate Networkdisk = "authenticate"
	NetworkdiskReadOnly     Networkdisk = "read-only"
	NetworkdiskDeny         Networkdisk = "deny"
	NetworkdiskEject        Networkdisk = "eject"
)

// The media type dictionary that controls volume mounting.
type MountControls struct {
	// Unused; set to an empty string.
	AllMedia *string `json:"all-media,omitempty" plist:"all-media,omitempty"`
	// A media action string or an array of media action strings.
	Cd *[]Cd `json:"cd,omitempty" plist:"cd,omitempty"`
	// A media action string or an array of media action strings.
	Dvd *[]Dvd `json:"dvd,omitempty" plist:"dvd,omitempty"`
	// A media action string or an array of media action strings.
	Bd *[]Bd `json:"bd,omitempty" plist:"bd,omitempty"`
	// A media action string or an array of media action strings.
	Blankcd *[]Blankcd `json:"blankcd,omitempty" plist:"blankcd,omitempty"`
	// A media action string or an array of media action strings.
	Blankdvd *[]Blankdvd `json:"blankdvd,omitempty" plist:"blankdvd,omitempty"`
	// A media action string or an array of media action strings.
	Blankbd *[]Blankbd `json:"blankbd,omitempty" plist:"blankbd,omitempty"`
	// A media action string or an array of media action strings.
	Dvdram *[]Dvdram `json:"dvdram,omitempty" plist:"dvdram,omitempty"`
	// A media action string or an array of media action strings.
	DiskImage *[]DiskImage `json:"disk-image,omitempty" plist:"disk-image,omitempty"`
	// A media action string or an array of media action strings.
	HarddiskInternal *[]HarddiskInternal `json:"harddisk-internal,omitempty" plist:"harddisk-internal,omitempty"`
	// A string or an array of media action strings. Internally installed SD cards and USB flash drives are included in the hard disk-external category.
	// This key is the default for media types that don't fall into other categories.
	HarddiskExternal *[]HarddiskExternal `json:"harddisk-external,omitempty" plist:"harddisk-external,omitempty"`
	// A media action string or an array of media action strings.
	Networkdisk *[]Networkdisk `json:"networkdisk,omitempty" plist:"networkdisk,omitempty"`
}

// The media type dictionary that controls volume unmounting.
type UnmountControls struct {
	// Unused; set to an empty string.
	AllMedia *string `json:"all-media,omitempty" plist:"all-media,omitempty"`
	// A media action string or an array of media action strings.
	Cd *[]Cd `json:"cd,omitempty" plist:"cd,omitempty"`
	// A media action string or an array of media action strings.
	Dvd *[]Dvd `json:"dvd,omitempty" plist:"dvd,omitempty"`
	// A media action string or an array of media action strings.
	Bd *[]Bd `json:"bd,omitempty" plist:"bd,omitempty"`
	// A media action string or an array of media action strings.
	Blankcd *[]Blankcd `json:"blankcd,omitempty" plist:"blankcd,omitempty"`
	// A media action string or an array of media action strings.
	Blankdvd *[]Blankdvd `json:"blankdvd,omitempty" plist:"blankdvd,omitempty"`
	// A media action string or an array of media action strings.
	Blankbd *[]Blankbd `json:"blankbd,omitempty" plist:"blankbd,omitempty"`
	// A media action string or an array of media action strings.
	Dvdram *[]Dvdram `json:"dvdram,omitempty" plist:"dvdram,omitempty"`
	// A media action string or an array of media action strings.
	DiskImage *[]DiskImage `json:"disk-image,omitempty" plist:"disk-image,omitempty"`
	// A media action string or an array of media action strings.
	HarddiskInternal *[]HarddiskInternal `json:"harddisk-internal,omitempty" plist:"harddisk-internal,omitempty"`
	// A string or an array of media action strings. Internally installed SD cards and USB flash drives are included in the hard disk-external category.
	// This key is the default for media types that don't fall into other categories.
	HarddiskExternal *[]HarddiskExternal `json:"harddisk-external,omitempty" plist:"harddisk-external,omitempty"`
	// A media action string or an array of media action strings.
	Networkdisk *[]Networkdisk `json:"networkdisk,omitempty" plist:"networkdisk,omitempty"`
}

// The payload that configures the third wired, active Ethernet interface.
type Net8021XThirdActiveEthernet struct {
	*CommonPayloadKeys
	Properties map[string]any
}

func (p *Net8021XThirdActiveEthernet) PayloadType() string {
	return "com.apple.thirdactiveethernet.managed"
}

// The payload that configures the third wired Ethernet interface.
type Net8021XThirdEthernet struct {
	*CommonPayloadKeys
	Properties map[string]any
}

func (p *Net8021XThirdEthernet) PayloadType() string {
	return "com.apple.thirdethernet.managed"
}

// The payload that configures the Apple TV remote.
type TVRemote struct {
	*CommonPayloadKeys
	// The array of valid devices that Apple TV can connect to.
	AllowedRemotes *[]*AllowedRemotesItem `json:"AllowedRemotes,omitempty" plist:"AllowedRemotes,omitempty"`
	// The array of valid Apple TV identifiers that the remote can connect to.
	AllowedTVs *[]*AllowedTVsItem `json:"AllowedTVs,omitempty" plist:"AllowedTVs,omitempty"`
}

func (p *TVRemote) PayloadType() string {
	return "com.apple.tvremote"
}

// The array of valid devices that Apple TV can connect to.
type AllowedRemotesItem struct {
	// The MAC address of a permitted iOS device that can control this Apple TV. Use the format `xx:xx:xx:xx:xx:xx`, which isn't case-sensitive.
	RemoteDeviceID string `json:"RemoteDeviceID" plist:"RemoteDeviceID" required:"true"`
}

// The array of valid Apple TV identifiers that the remote can connect to.
type AllowedTVsItem struct {
	// The MAC address of an Apple TV device that the system permits this iOS device to control. Use the format `xx:xx:xx:xx:xx:xx`, which isn't case-sensitive.
	TVDeviceID string `json:"TVDeviceID" plist:"TVDeviceID" required:"true"`
	// The name of an Apple TV device that the system permits this iOS device to control.
	TVDeviceName *string `json:"TVDeviceName,omitempty" plist:"TVDeviceName,omitempty"`
}

// The payload that configures the accessibility features of the device.
type Accessibility struct {
	*CommonPayloadKeys
	// The minimum zoom level in the Zoom options.
	CloseViewFarPoint *int64 `json:"closeViewFarPoint,omitempty" plist:"closeViewFarPoint,omitempty"`
	// If `true`, enables "Use keyboard shortcuts" in the Zoom options.
	CloseViewHotkeysEnabled *bool `json:"closeViewHotkeysEnabled,omitempty" plist:"closeViewHotkeysEnabled,omitempty"`
	// The maximum zoom level in the Zoom options.
	CloseViewNearPoint *int64 `json:"closeViewNearPoint,omitempty" plist:"closeViewNearPoint,omitempty"`
	// If `true`, enables "Use scroll gesture" in the Zoom options.
	CloseViewScrollWheelToggle *bool `json:"closeViewScrollWheelToggle,omitempty" plist:"closeViewScrollWheelToggle,omitempty"`
	// If `true`, enables "Show preview rectangle" in the Zoom options. Only available in macOS 10.15 and earlier.
	CloseViewShowPreview *bool `json:"closeViewShowPreview,omitempty" plist:"closeViewShowPreview,omitempty"`
	// If `true`, enables "Smooth images" in the Zoom options.
	CloseViewSmoothImages *bool `json:"closeViewSmoothImages,omitempty" plist:"closeViewSmoothImages,omitempty"`
	// The contrast value in the Display options.
	Contrast *float64 `json:"contrast,omitempty" plist:"contrast,omitempty"`
	// If `true`, enables "Flash the screen" in the Audio options.
	FlashScreen *bool `json:"flashScreen,omitempty" plist:"flashScreen,omitempty"`
	// If `true`, enables "Use grayscale" in the Display options.
	// This option is deprecated in macOS 11.
	Grayscale *bool `json:"grayscale,omitempty" plist:"grayscale,omitempty"`
	// If `true`, enables Mouse Keys in the Mouse & Trackpad options.
	MouseDriver *bool `json:"mouseDriver,omitempty" plist:"mouseDriver,omitempty"`
	// The size of the cursor.
	MouseDriverCursorSize *int64 `json:"mouseDriverCursorSize,omitempty" plist:"mouseDriverCursorSize,omitempty"`
	// If `true`, ignores the built-in trackpad.
	MouseDriverIgnoreTrackpad *bool `json:"mouseDriverIgnoreTrackpad,omitempty" plist:"mouseDriverIgnoreTrackpad,omitempty"`
	// The initial delay before moving the mouse with Mouse Keys.
	MouseDriverInitialDelay *int64 `json:"mouseDriverInitialDelay,omitempty" plist:"mouseDriverInitialDelay,omitempty"`
	// The maximum speed for the cursor when using Mouse Keys.
	MouseDriverMaxSpeed *int64 `json:"mouseDriverMaxSpeed,omitempty" plist:"mouseDriverMaxSpeed,omitempty"`
	// If `true`, enables "Slow Keys" in the Keyboard options.
	SlowKey *bool `json:"slowKey,omitempty" plist:"slowKey,omitempty"`
	// If `true`, enables "click key sounds" for Slow Keys.
	SlowKeyBeepOn *bool `json:"slowKeyBeepOn,omitempty" plist:"slowKeyBeepOn,omitempty"`
	// The acceptance delay, in milliseconds, for Slow Keys.
	SlowKeyDelay *int64 `json:"slowKeyDelay,omitempty" plist:"slowKeyDelay,omitempty"`
	// If `true`, plays stereo audio as mono.
	StereoAsMono *bool `json:"stereoAsMono,omitempty" plist:"stereoAsMono,omitempty"`
	// If `true`, enables Sticky Keys in the Keyboard options.
	StickyKey *bool `json:"stickyKey,omitempty" plist:"stickyKey,omitempty"`
	// If `true`, enables the beep when a modifier key is set for Sticky Keys.
	StickyKeyBeepOnModifier *bool `json:"stickyKeyBeepOnModifier,omitempty" plist:"stickyKeyBeepOnModifier,omitempty"`
	// If `true`, enables "Display pressed keys on screen" for Sticky Keys.
	StickyKeyShowWindow *bool `json:"stickyKeyShowWindow,omitempty" plist:"stickyKeyShowWindow,omitempty"`
	// If `true`, enables Voice Over.
	VoiceOverOnOffKey *bool `json:"voiceOverOnOffKey,omitempty" plist:"voiceOverOnOffKey,omitempty"`
	// If `true`, enables Invert Colors in Display Accommodations.
	WhiteOnBlack *bool `json:"whiteOnBlack,omitempty" plist:"whiteOnBlack,omitempty"`
}

func (p *Accessibility) PayloadType() string {
	return "com.apple.universalaccess"
}

// The payload that configures a per-app VPN.
// The fields in this payload are the same as the VPN payload, with the addition of the fields shown below. On watchOS, only the IKEv2 VPN type is supported.
type AppLayerVPN struct {
	*CommonPayloadKeys
	// A globally unique identifier for this VPN configuration.
	VPNUUID string `json:"VPNUUID" plist:"VPNUUID" required:"true"`
	// A string representing the data network name (DNN) or app category identifying a Cellular Slice. The device forces the VPN tunnel to use the specified Cellular Slice.
	CellularSliceUUID *string `json:"CellularSliceUUID,omitempty" plist:"CellularSliceUUID,omitempty"`
	// An array with entries that must each specify a domain that triggers the VPN connection in Safari. Each entry is in the format `www.apple.com`.
	SafariDomains *[]string `json:"SafariDomains,omitempty" plist:"SafariDomains,omitempty"`
	// An array with entries that must each specify a domain that triggers this VPN connection in Mail. Each entry is in the format `www.apple.com`.
	// This property is deprecated in iOS 13.4 and later; use the `VPNUUID` property of the `Mail` or `ExchangeActiveSync` payload instead.
	MailDomains *[]string `json:"MailDomains,omitempty" plist:"MailDomains,omitempty"`
	// An array with entries that must each specify a domain that triggers this VPN connection in Calendar. Each entry is in the format `www.apple.com`.
	// This property is deprecated in iOS 13.4 and later; use the `VPNUUID` property of the `CalDAV` payload instead.
	CalendarDomains *[]string `json:"CalendarDomains,omitempty" plist:"CalendarDomains,omitempty"`
	// An array with entries that must each specify a domain that triggers this VPN connection in Contacts. Each entry is in the format `www.apple.com`.
	// This property is deprecated in iOS 13.4 and later; use the `VPNUUID` property of the `CardDAV` payload instead.
	ContactsDomains *[]string `json:"ContactsDomains,omitempty" plist:"ContactsDomains,omitempty"`
	// An array with entries that must each specify a domain that triggers this VPN. The domains must also be part of the `apple-app-site-association` file, as described in `Supporting associated domains`.
	// Available in iOS 14 and later, and macOS 11 and later.
	AssociatedDomains *[]string `json:"AssociatedDomains,omitempty" plist:"AssociatedDomains,omitempty"`
	// An array with entries that each specify a domain that doesn't trigger this VPN for connections to the domain.
	// Available in iOS 14 and later, and macOS 11 and later.
	ExcludedDomains *[]string `json:"ExcludedDomains,omitempty" plist:"ExcludedDomains,omitempty"`
	// If `true`, automatically connects the VPN when associated apps for this per-app VPN service initiate network communication. Otherwise, the user must initiate the connection manually before those apps can initiate network communication. If this key isn't present, the value of the `OnDemandEnabled` key determines the status of per-app VPN On Demand.
	OnDemandMatchAppEnabled *bool `json:"OnDemandMatchAppEnabled,omitempty" plist:"OnDemandMatchAppEnabled,omitempty"`
	// An array of SMB domains that's accessible through this VPN connection.
	// Available in iOS 13 and later.
	SMBDomains *[]string `json:"SMBDomains,omitempty" plist:"SMBDomains,omitempty"`
}

func (p *AppLayerVPN) PayloadType() string {
	return "com.apple.vpn.managed.applayer"
}

// The payload that configures per-app VPN settings.
// This payload is only valid on macOS.
type AppToAppLayerVPNMapping struct {
	*CommonPayloadKeys
	// The array of VPN mapping dictionaries.
	AppLayerVPNMapping []*AppLayerVPNMappingItem `json:"AppLayerVPNMapping" plist:"AppLayerVPNMapping" required:"true"`
}

func (p *AppToAppLayerVPNMapping) PayloadType() string {
	return "com.apple.vpn.managed.appmapping"
}

// A dictionary defining a per-app VPN relationship.
type AppLayerVPNMappingItem struct {
	// The bundle identifier of the app using the per-app VPN.
	Identifier string `json:"Identifier" plist:"Identifier" required:"true"`
	// The identifier of the per-app VPN payload, which defines the per-app VPN that the app uses. See the `VPNUUID` key of the `AppLayerVPN` payload.
	VPNUUID string `json:"VPNUUID" plist:"VPNUUID" required:"true"`
	// The code signature designated requirement of the app using the per-app VPN.
	DesignatedRequirement string `json:"DesignatedRequirement" plist:"DesignatedRequirement" required:"true"`
	// The code signature signing identifier of the app using the per-app VPN.
	SigningIdentifier string `json:"SigningIdentifier" plist:"SigningIdentifier" required:"true"`
	// The file-system path of the executable using the per-app VPN.
	Path *string `json:"Path,omitempty" plist:"Path,omitempty"`
	// An array of dictionaries. Each dictionary specifies a per-app VPN rule. Use this property to restrict this per-app VPN rule to only match the app's spawned _helper tool_ network traffic.
	// For example, to match network traffic that the `curl` command generates when run from the Terminal.app, create an app mapping payload for Terminal.app and set the payload's `MatchTools` key to an array that contains a dictionary that matches the `curl` command-line tool.
	// If you don't specify the `MatchTools` key, this per-app VPN rule matches all network traffic that the matching app and its spawned helper tools generate.
	MatchTools *[]*MatchToolsItem `json:"MatchTools,omitempty" plist:"MatchTools,omitempty"`
}

// Specifies a per-app VPN rule to match network traffic that the app's spawned command-line tool generates.
type MatchToolsItem struct {
	// The code signature designated requirement of the command-line tool using the per-app VPN.
	DesignatedRequirement string `json:"DesignatedRequirement" plist:"DesignatedRequirement" required:"true"`
	// The code signature signing identifier of the command-line tool using the per-app VPN.
	SigningIdentifier string `json:"SigningIdentifier" plist:"SigningIdentifier" required:"true"`
	// The file-system path of the command-line tool using the per-app VPN.
	Path *string `json:"Path,omitempty" plist:"Path,omitempty"`
}

// The payload that configures a VPN.
type VPN1 struct {
	*CommonPayloadKeys
	// The type of the VPN, which defines which settings are appropriate for this VPN payload.
	// If the type is `VPN` or `TransparentProxy`, then the system requires a value for `VPNSubType`.
	// `TransparentProxy` is only available in macOS. `L2TP` and `IPSec` aren't available in tvOS. `AlwaysOn` is only available on iOS and Apple Watch pairing isn't supported with `AlwaysOn`. For a previously paired Apple Watch, all phone-watch communications cease when `AlwaysOn` is enabled. Not available in watchOS.
	VPNType VPNType `json:"VPNType" plist:"VPNType" required:"true"`
	// An identifier for a vendor-specified configuration dictionary when the value for `VPNType` is `VPN`.
	// If `VPNType` is `VPN`, the system requires this field. If the configuration targets a VPN solution that uses a VPN plugin, then this field contains the bundle identifier of the plugin. Here are some examples:
	// - Cisco AnyConnect: `com.cisco.anyconnect.applevpn.plugin`
	// - Juniper SSL: `net.juniper.sslvpn`
	// - F5 SSL: `com.f5.F5-Edge-Client.vpnplugin`
	// - SonicWALL Mobile Connect: `com.sonicwall.SonicWALL-SSLVPN.vpnplugin`
	// - “Aruba VIA: `com.arubanetworks.aruba-via.vpnplugin`
	// If the configuration targets a VPN solution that uses a network extension provider, then this field contains the bundle identifier of the app that contains the provider. Contact the VPN solution vendor for the value of the identifier.
	// If `VPNType` is `IKEv2`, then the `VPNSubType` field is optional and reserved for future use. If it's specified, it needs to contain an empty string.
	// Not available in watchOS.
	VPNSubType *string `json:"VPNSubType,omitempty" plist:"VPNSubType,omitempty"`
	// The description of the VPN connection that the system displays on the device. Not available in watchOS.
	UserDefinedName string `json:"UserDefinedName" plist:"UserDefinedName" required:"true"`
	// The vendor-specific configuration dictionary, which the system reads only when `VPNSubType` has a value. Not available in watchOS.
	VendorConfig *VendorConfig `json:"VendorConfig,omitempty" plist:"VendorConfig,omitempty"`
	// The dictionary to use when `VPNType` is `VPN`.
	VPN *VPNVPN `json:"VPN,omitempty" plist:"VPN,omitempty"`
	// The dictionary that contains IPv4 settings. Not available in watchOS.
	IPv4 *IPv4 `json:"IPv4,omitempty" plist:"IPv4,omitempty"`
	// The dictionary to use when `VPNType` is `L2TP` or `PTPP`. Not available in watchOS.
	PPP *PPP `json:"PPP,omitempty" plist:"PPP,omitempty"`
	// The dictionary that contains IPSec settings. Not available in watchOS.
	IPSec *IPSec `json:"IPSec,omitempty" plist:"IPSec,omitempty"`
	// The dictionary to use when `VPNType` is `IKEv2`.
	IKEv2 *IKEv2 `json:"IKEv2,omitempty" plist:"IKEv2,omitempty"`
	// A dictionary to use for all VPN types.
	DNS *DNS `json:"DNS,omitempty" plist:"DNS,omitempty"`
	// The dictionary to use to configure `Proxies` for use with `VPN`.
	Proxies *VPNProxies `json:"Proxies,omitempty" plist:"Proxies,omitempty"`
	// The dictionary to use when `VPNType` is `AlwaysOn`. Not available in tvOS or watchOS.
	AlwaysOn *AlwaysOn `json:"AlwaysOn,omitempty" plist:"AlwaysOn,omitempty"`
	// The dictionary to use when `VPNType` is `TransparentProxy`. Available in macOS 14 and later.
	TransparentProxy *TransparentProxy `json:"TransparentProxy,omitempty" plist:"TransparentProxy,omitempty"`
}

func (p *VPN1) PayloadType() string {
	return "com.apple.vpn.managed"
}

// The type of the VPN, which defines which settings are appropriate for this VPN payload.
// If the type is `VPN` or `TransparentProxy`, then the system requires a value for `VPNSubType`.
// `TransparentProxy` is only available in macOS. `L2TP` and `IPSec` aren't available in tvOS. `AlwaysOn` is only available on iOS and Apple Watch pairing isn't supported with `AlwaysOn`. For a previously paired Apple Watch, all phone-watch communications cease when `AlwaysOn` is enabled. Not available in watchOS.
type VPNType string

const (
	VPNTypeVPN              VPNType = "VPN"
	VPNTypeL2TP             VPNType = "L2TP"
	VPNTypeIPSec            VPNType = "IPSec"
	VPNTypeIKEv2            VPNType = "IKEv2"
	VPNTypeAlwaysOn         VPNType = "AlwaysOn"
	VPNTypeTransparentProxy VPNType = "TransparentProxy"
)

// The vendor-specific configuration dictionary, which the system reads only when `VPNSubType` has a value. Not available in watchOS.
type VendorConfig struct {
	// The Kerberos realm name, which needs to be properly capitalized. Valid only for Juniper SSL and Pulse Secure. Not available in watchOS.
	Realm *string `json:"Realm,omitempty" plist:"Realm,omitempty"`
	// The role to select when connecting to the server. Valid only for Juniper SSL and Pulse Secure. Not available in watchOS.
	Role *string `json:"Role,omitempty" plist:"Role,omitempty"`
	// The group to connect to on the head end. Valid for Cisco AnyConnect and Cisco Legacy AnyConnect. Not available in watchOS.
	Group *string `json:"Group,omitempty" plist:"Group,omitempty"`
	// The login group or domain. Valid only for SonicWALL Mobile Connect. Not available in watchOS.
	LoginGroupOrDomain *string `json:"LoginGroupOrDomain,omitempty" plist:"LoginGroupOrDomain,omitempty"`
}

// The dictionary to use when `VPNType` is `VPN`.
type VPNVPN struct {
	// The VPN account username.
	AuthName *string `json:"AuthName,omitempty" plist:"AuthName,omitempty"`
	// The VPN account password. Only use this if `AuthenticationMethod` is set to `Password`.
	AuthPassword *string `json:"AuthPassword,omitempty" plist:"AuthPassword,omitempty"`
	// The IP address or hostname of the VPN server.
	RemoteAddress string `json:"RemoteAddress" plist:"RemoteAddress" required:"true"`
	// The authentication method to use.
	AuthenticationMethod *VPNAuthenticationMethod `default:"Password" json:"AuthenticationMethod,omitempty" plist:"AuthenticationMethod,omitempty"`
	// The UUID of the certificate payload within the same profile to use for account credentials.
	PayloadCertificateUUID *string `json:"PayloadCertificateUUID,omitempty" plist:"PayloadCertificateUUID,omitempty"`
	// The bundle identifier for the VPN provider. Not available in watchOS.
	ProviderBundleIdentifier *string `json:"ProviderBundleIdentifier,omitempty" plist:"ProviderBundleIdentifier,omitempty"`
	// If the VPN provider is implemented as a system extension, this field is required. Not available in watchOS.
	ProviderDesignatedRequirement *string `json:"ProviderDesignatedRequirement,omitempty" plist:"ProviderDesignatedRequirement,omitempty"`
	// If `1`, disconnects after an on-demand connection idles.
	DisconnectOnIdle *VPNDisconnectOnIdle `default:"0" json:"DisconnectOnIdle,omitempty" plist:"DisconnectOnIdle,omitempty"`
	// The length of time to wait, in seconds, before disconnecting an on-demand connection. In watchOS, the maximum allowed value is `15`.
	DisconnectOnIdleTimer *int64 `json:"DisconnectOnIdleTimer,omitempty" plist:"DisconnectOnIdleTimer,omitempty"`
	// The type of VPN service. If the value is `app-proxy`, the service tunnels traffic at the app level. If the value is `packet-tunnel`, the service tunnels traffic at the IP layer. Not available in watchOS.
	ProviderType *VPNProviderType `default:"packet-tunnel" json:"ProviderType,omitempty" plist:"ProviderType,omitempty"`
	// If `1“, routes all traffic through the VPN, with some exclusions. Several of the exclusions can be controlled with the `ExcludeLocalNetworks`, `ExcludeCellularServices`, `ExcludeAPNs` and `ExcludeDeviceCommunication` properties. The following traffic is always excluded from the tunnel:
	// - Traffic necessary for connecting and maintaining the device's network connection, such as DHCP.
	// - Traffic necessary for connecting to captive networks.
	// - Certain cellular services traffic that is not routable over the internet and is instead directly routed to the cellular network. See the ExcludeCellularServices property for more details.
	// - Network communication with a companion device such as a watchOS device.
	// Not available in watchOS.
	IncludeAllNetworks *VPNIncludeAllNetworks `default:"0" json:"IncludeAllNetworks,omitempty" plist:"IncludeAllNetworks,omitempty"`
	// If `1`, all the VPN's non-default routes take precedence over any locally defined routes.
	// If `IncludeAllNetworks` is `1`, the system ignores the value of `EnforceRoutes`.
	// Available in iOS 14.2 and later, and macOS 11 and later. Not available in watchOS.
	EnforceRoutes *VPNEnforceRoutes `default:"0" json:"EnforceRoutes,omitempty" plist:"EnforceRoutes,omitempty"`
	// If `1` and `IncludeAllNetworks` is `1`, routes all local network traffic outside the VPN. Not available in watchOS.
	ExcludeLocalNetworks *VPNExcludeLocalNetworks `json:"ExcludeLocalNetworks,omitempty" plist:"ExcludeLocalNetworks,omitempty"`
	// If `1` and `IncludeAllNetworks` is `1`, then the system excludes internet-routable network traffic for cellular services (VoLTE, Wi-Fi Calling, IMS, MMS, Visual Voicemail, etc.) from the tunnel. Note that some cellular carriers route cellular services traffic directly to the carrier network, bypassing the internet. Such cellular services traffic is always excluded from the tunnel. Not available in watchOS.
	ExcludeCellularServices *VPNExcludeCellularServices `default:"1" json:"ExcludeCellularServices,omitempty" plist:"ExcludeCellularServices,omitempty"`
	// If `1` and `IncludeAllNetworks` is `1`, then the system excludes the network traffic for the Apple Push Notification service (APNs) from the tunnel. Not available in watchOS.
	ExcludeAPNs *VPNExcludeAPNs `default:"1" json:"ExcludeAPNs,omitempty" plist:"ExcludeAPNs,omitempty"`
	// If set to `1` and `IncludeAllNetworks` is set to `1`, the device excludes network traffic used for communicating with devices connected via USB or Wi-Fi from the tunnel.
	ExcludeDeviceCommunication *VPNExcludeDeviceCommunication `default:"1" json:"ExcludeDeviceCommunication,omitempty" plist:"ExcludeDeviceCommunication,omitempty"`
	// If `1`, enables VPN On Demand.
	OnDemandEnabled *VPNOnDemandEnabled `default:"0" json:"OnDemandEnabled,omitempty" plist:"OnDemandEnabled,omitempty"`
	// If `1`, the Connect On Demand toggle in Settings is disabled for this configuration. Available in iOS 14 and later. Not available in watchOS.
	OnDemandUserOverrideDisabled *VPNOnDemandUserOverrideDisabled `default:"0" json:"OnDemandUserOverrideDisabled,omitempty" plist:"OnDemandUserOverrideDisabled,omitempty"`
	// A list of domain names. The system treats associated domain names as though they're associated with the `OnDemandMatchDomainsOnRetry` key. This behavior can be overridden by `OnDemandRules`.
	// In iOS 7 and later, this key is deprecated (but still supported) in favor of `EvaluateConnection` actions in the `OnDemandRules` dictionaries.
	// Not available in watchOS.
	OnDemandMatchDomainsAlways *[]string `json:"OnDemandMatchDomainsAlways,omitempty" plist:"OnDemandMatchDomainsAlways,omitempty"`
	// A list of domain names. If the host name ends with one of these domain names, the system doesn't start the VPN automatically. The system uses this value to exclude a subdomain within an included domain.
	// In iOS 7 and later, this key is deprecated (but still supported) in favor of `EvaluateConnection` actions in the `OnDemandRules` dictionaries.
	// Not available in watchOS.
	OnDemandMatchDomainsNever *[]string `json:"OnDemandMatchDomainsNever,omitempty" plist:"OnDemandMatchDomainsNever,omitempty"`
	// A list of domain names. If the host name ends with one of these domain names and a DNS query for that domain name fails, the system starts the VPN automatically.
	// In iOS 7 and later, this key is deprecated (but still supported) in favor of `EvaluateConnection` actions in the `OnDemandRules` dictionaries.
	// Not available in watchOS.
	OnDemandMatchDomainsOnRetry *[]string `json:"OnDemandMatchDomainsOnRetry,omitempty" plist:"OnDemandMatchDomainsOnRetry,omitempty"`
	// An array of dictionaries defining On Demand Rules.
	OnDemandRules *[]*VPNOnDemandRulesOnDemandRulesElement `json:"OnDemandRules,omitempty" plist:"OnDemandRules,omitempty"`
}

// The authentication method to use.
type VPNAuthenticationMethod string

const (
	VPNAuthenticationMethodPassword            VPNAuthenticationMethod = "Password"
	VPNAuthenticationMethodCertificate         VPNAuthenticationMethod = "Certificate"
	VPNAuthenticationMethodPasswordCertificate VPNAuthenticationMethod = "Password+Certificate"
)

// If `1`, disconnects after an on-demand connection idles.
type VPNDisconnectOnIdle int64

const (
	VPNDisconnectOnIdle0 VPNDisconnectOnIdle = 0
	VPNDisconnectOnIdle1 VPNDisconnectOnIdle = 1
)

// The type of VPN service. If the value is `app-proxy`, the service tunnels traffic at the app level. If the value is `packet-tunnel`, the service tunnels traffic at the IP layer. Not available in watchOS.
type VPNProviderType string

const (
	VPNProviderTypePacketTunnel VPNProviderType = "packet-tunnel"
	VPNProviderTypeAppProxy     VPNProviderType = "app-proxy"
)

// If `1“, routes all traffic through the VPN, with some exclusions. Several of the exclusions can be controlled with the `ExcludeLocalNetworks`, `ExcludeCellularServices`, `ExcludeAPNs` and `ExcludeDeviceCommunication` properties. The following traffic is always excluded from the tunnel:
// - Traffic necessary for connecting and maintaining the device's network connection, such as DHCP.
// - Traffic necessary for connecting to captive networks.
// - Certain cellular services traffic that is not routable over the internet and is instead directly routed to the cellular network. See the ExcludeCellularServices property for more details.
// - Network communication with a companion device such as a watchOS device.
// Not available in watchOS.
type VPNIncludeAllNetworks int64

const (
	VPNIncludeAllNetworks0 VPNIncludeAllNetworks = 0
	VPNIncludeAllNetworks1 VPNIncludeAllNetworks = 1
)

// If `1`, all the VPN's non-default routes take precedence over any locally defined routes.
// If `IncludeAllNetworks` is `1`, the system ignores the value of `EnforceRoutes`.
// Available in iOS 14.2 and later, and macOS 11 and later. Not available in watchOS.
type VPNEnforceRoutes int64

const (
	VPNEnforceRoutes0 VPNEnforceRoutes = 0
	VPNEnforceRoutes1 VPNEnforceRoutes = 1
)

// If `1` and `IncludeAllNetworks` is `1`, routes all local network traffic outside the VPN. Not available in watchOS.
type VPNExcludeLocalNetworks int64

const (
	VPNExcludeLocalNetworks0 VPNExcludeLocalNetworks = 0
	VPNExcludeLocalNetworks1 VPNExcludeLocalNetworks = 1
)

// If `1` and `IncludeAllNetworks` is `1`, then the system excludes internet-routable network traffic for cellular services (VoLTE, Wi-Fi Calling, IMS, MMS, Visual Voicemail, etc.) from the tunnel. Note that some cellular carriers route cellular services traffic directly to the carrier network, bypassing the internet. Such cellular services traffic is always excluded from the tunnel. Not available in watchOS.
type VPNExcludeCellularServices int64

const (
	VPNExcludeCellularServices0 VPNExcludeCellularServices = 0
	VPNExcludeCellularServices1 VPNExcludeCellularServices = 1
)

// If `1` and `IncludeAllNetworks` is `1`, then the system excludes the network traffic for the Apple Push Notification service (APNs) from the tunnel. Not available in watchOS.
type VPNExcludeAPNs int64

const (
	VPNExcludeAPNs0 VPNExcludeAPNs = 0
	VPNExcludeAPNs1 VPNExcludeAPNs = 1
)

// If set to `1` and `IncludeAllNetworks` is set to `1`, the device excludes network traffic used for communicating with devices connected via USB or Wi-Fi from the tunnel.
type VPNExcludeDeviceCommunication int64

const (
	VPNExcludeDeviceCommunication0 VPNExcludeDeviceCommunication = 0
	VPNExcludeDeviceCommunication1 VPNExcludeDeviceCommunication = 1
)

// If `1`, enables VPN On Demand.
type VPNOnDemandEnabled int64

const (
	VPNOnDemandEnabled0 VPNOnDemandEnabled = 0
	VPNOnDemandEnabled1 VPNOnDemandEnabled = 1
)

// If `1`, the Connect On Demand toggle in Settings is disabled for this configuration. Available in iOS 14 and later. Not available in watchOS.
type VPNOnDemandUserOverrideDisabled int64

const (
	VPNOnDemandUserOverrideDisabled0 VPNOnDemandUserOverrideDisabled = 0
	VPNOnDemandUserOverrideDisabled1 VPNOnDemandUserOverrideDisabled = 1
)

type VPNOnDemandRulesOnDemandRulesElement struct {
	// The action to take if this dictionary matches the current network. Possible values are:
	// - `Allow`: Deprecated. Allow VPN On Demand to connect if triggered.
	// - `Connect`: Unconditionally initiate a VPN connection on the next network attempt.
	// - `Disconnect`: Tear down the VPN connection and don't reconnect on demand as long as this dictionary matches.
	// - `EvaluateConnection`: Evaluate the ActionParameters array for each connection attempt.
	// - `Ignore`: Leave any existing VPN connection up, but don't reconnect on demand as long as this dictionary matches.
	// Only the `Disconnect` action is available on watchOS 10 and later.
	Action VPNOnDemandRulesOnDemandRulesElementAction `json:"Action" plist:"Action" required:"true"`
	// An array of dictionaries that provides rules similar to the `OnDemandRules` dictionary, but evaluated on each connection instead of when the network changes. This value is only for use with dictionaries in which the `Action` value is `EvaluateConnection`. The system evaluates these dictionaries in order and the first dictionary that matches determines the behavior. Not available in watchOS.
	ActionParameters *[]*VPNOnDemandRulesOnDemandRulesElementActionParametersActionParameter `json:"ActionParameters,omitempty" plist:"ActionParameters,omitempty"`
	// An array of domain names. This rule matches if any of the domain names in the specified list matches any domain in the device's search domains list.
	// The system supports a wildcard (`*`) prefix. For example, `*.example.com` matches against either `mydomain.example.com` or `yourdomain.example.com`.
	DNSDomainMatch *[]string `json:"DNSDomainMatch,omitempty" plist:"DNSDomainMatch,omitempty"`
	// An array of IP addresses. This rule matches if any of the network's specified DNS servers match any entry in the array.
	// The system supports matching with a single wildcard. For example, `17.*` matches any DNS server in the `17.0.0.0/8` subnet.
	DNSServerAddressMatch *[]string `json:"DNSServerAddressMatch,omitempty" plist:"DNSServerAddressMatch,omitempty"`
	// An interface type. If specified, this rule matches only if the primary network interface hardware matches the specified type.
	InterfaceTypeMatch *VPNOnDemandRulesOnDemandRulesElementInterfaceTypeMatch `json:"InterfaceTypeMatch,omitempty" plist:"InterfaceTypeMatch,omitempty"`
	// An array of SSIDs to match against the current network. If the network isn't a Wi-Fi network or if the SSID doesn't appear in this array, the match fails.
	// Omit this key and the corresponding array to match against any SSID.
	SSIDMatch *[]string `json:"SSIDMatch,omitempty" plist:"SSIDMatch,omitempty"`
	// A URL to probe. This rule matches when this URL is successfully fetched (returns a `200` HTTP status code) without redirection. Not available in watchOS.
	URLStringProbe *string `json:"URLStringProbe,omitempty" plist:"URLStringProbe,omitempty"`
}

// The action to take if this dictionary matches the current network. Possible values are:
// - `Allow`: Deprecated. Allow VPN On Demand to connect if triggered.
// - `Connect`: Unconditionally initiate a VPN connection on the next network attempt.
// - `Disconnect`: Tear down the VPN connection and don't reconnect on demand as long as this dictionary matches.
// - `EvaluateConnection`: Evaluate the ActionParameters array for each connection attempt.
// - `Ignore`: Leave any existing VPN connection up, but don't reconnect on demand as long as this dictionary matches.
// Only the `Disconnect` action is available on watchOS 10 and later.
type VPNOnDemandRulesOnDemandRulesElementAction string

const (
	VPNOnDemandRulesOnDemandRulesElementActionAllow              VPNOnDemandRulesOnDemandRulesElementAction = "Allow"
	VPNOnDemandRulesOnDemandRulesElementActionConnect            VPNOnDemandRulesOnDemandRulesElementAction = "Connect"
	VPNOnDemandRulesOnDemandRulesElementActionDisconnect         VPNOnDemandRulesOnDemandRulesElementAction = "Disconnect"
	VPNOnDemandRulesOnDemandRulesElementActionEvaluateConnection VPNOnDemandRulesOnDemandRulesElementAction = "EvaluateConnection"
	VPNOnDemandRulesOnDemandRulesElementActionIgnore             VPNOnDemandRulesOnDemandRulesElementAction = "Ignore"
)

// A dictionary that provides rules similar to the OnDemandRules dictionary, but evaluated on each connection instead of when the network changes. These dictionaries are evaluated in order, and the behavior is determined by the first dictionary that matches.
// The keys allowed in each dictionary are described below. Note: This array is used only for dictionaries in which EvaluateConnection is the Action value.
type VPNOnDemandRulesOnDemandRulesElementActionParametersActionParameter struct {
	// The domains to apply this evaluation.
	Domains []string `json:"Domains" plist:"Domains" required:"true"`
	// Defines the VPN behavior for the specified domains. Allowed values are:
	// * 'ConnectIfNeeded': The specified domains should trigger a VPN connection attempt if domain name resolution fails, such as when the DNS server indicates that it can't resolve the domain, responds with a redirection to a different server, or fails to respond (timeout).
	// * 'NeverConnect': The specified domains should never trigger a VPN connection attempt.
	DomainAction VPNOnDemandRulesOnDemandRulesElementActionParametersActionParameterDomainAction `json:"DomainAction" plist:"DomainAction" required:"true"`
	// An array of IP addresses of DNS servers to use for resolving the specified domains. These servers don't need to be part of the device's current network configuration. If these DNS servers aren't reachable, the system establishes a VPN connection. These DNS servers need to be either internal DNS servers or trusted external DNS servers.
	// This key is valid only if the value of 'DomainAction' is 'ConnectIfNeeded'.
	RequiredDNSServers *[]string `json:"RequiredDNSServers,omitempty" plist:"RequiredDNSServers,omitempty"`
	// An HTTP or HTTPS (preferred) URL to probe, using a GET request. If the URL's hostname can't be resolved, if the server is unreachable, or if the server doesn't respond with a 200 HTTP status code, a VPN connection is established in response.
	// This key is valid only if the value of 'DomainAction' is 'ConnectIfNeeded'.
	RequiredURLStringProbe *string `json:"RequiredURLStringProbe,omitempty" plist:"RequiredURLStringProbe,omitempty"`
}

// Defines the VPN behavior for the specified domains. Allowed values are:
// * 'ConnectIfNeeded': The specified domains should trigger a VPN connection attempt if domain name resolution fails, such as when the DNS server indicates that it can't resolve the domain, responds with a redirection to a different server, or fails to respond (timeout).
// * 'NeverConnect': The specified domains should never trigger a VPN connection attempt.
type VPNOnDemandRulesOnDemandRulesElementActionParametersActionParameterDomainAction string

const (
	VPNOnDemandRulesOnDemandRulesElementActionParametersActionParameterDomainActionConnectIfNeeded VPNOnDemandRulesOnDemandRulesElementActionParametersActionParameterDomainAction = "ConnectIfNeeded"
	VPNOnDemandRulesOnDemandRulesElementActionParametersActionParameterDomainActionNeverConnect    VPNOnDemandRulesOnDemandRulesElementActionParametersActionParameterDomainAction = "NeverConnect"
)

// An interface type. If specified, this rule matches only if the primary network interface hardware matches the specified type.
type VPNOnDemandRulesOnDemandRulesElementInterfaceTypeMatch string

const (
	VPNOnDemandRulesOnDemandRulesElementInterfaceTypeMatchEthernet VPNOnDemandRulesOnDemandRulesElementInterfaceTypeMatch = "Ethernet"
	VPNOnDemandRulesOnDemandRulesElementInterfaceTypeMatchWiFi     VPNOnDemandRulesOnDemandRulesElementInterfaceTypeMatch = "WiFi"
	VPNOnDemandRulesOnDemandRulesElementInterfaceTypeMatchCellular VPNOnDemandRulesOnDemandRulesElementInterfaceTypeMatch = "Cellular"
)

// The dictionary that contains IPv4 settings. Not available in watchOS.
type IPv4 struct {
	// If `1`, the system sends all network traffic over VPN. Only applies to Cisco IPsec and L2TP VPN types.
	OverridePrimary *OverridePrimary `default:"0" json:"OverridePrimary,omitempty" plist:"OverridePrimary,omitempty"`
}

// If `1`, the system sends all network traffic over VPN. Only applies to Cisco IPsec and L2TP VPN types.
type OverridePrimary int64

const (
	OverridePrimary0 OverridePrimary = 0
	OverridePrimary1 OverridePrimary = 1
)

// The dictionary to use when `VPNType` is `L2TP` or `PTPP`. Not available in watchOS.
type PPP struct {
	// The VPN account user name. This key is for use with L2TP and PPTP networks.
	AuthName *string `json:"AuthName,omitempty" plist:"AuthName,omitempty"`
	// If `TokenCard` is `1`, use this password for authentication. This key is for use with L2TP and PPTP networks.
	AuthPassword *string `json:"AuthPassword,omitempty" plist:"AuthPassword,omitempty"`
	// If `1`, uses a token card such as an RSA SecurID card for connecting. This key is for use with L2TP networks.
	TokenCard *TokenCard `default:"0" json:"TokenCard,omitempty" plist:"TokenCard,omitempty"`
	// The IP address or host name of VPN server. This key is for use with L2TP and PPTP networks.
	CommRemoteAddress *string `json:"CommRemoteAddress,omitempty" plist:"CommRemoteAddress,omitempty"`
	// An array of authentication plugins. For use of RSA SecurID, this array should only have one value: `EAP-RSA`. This key is for use with L2TP and PPTP networks.
	AuthEAPPlugins *[]AuthEAPPlugins `json:"AuthEAPPlugins,omitempty" plist:"AuthEAPPlugins,omitempty"`
	// An array of authentication protocols. For use of RSA SecurID, this array should have one value, `EAP`. This key is for use with L2TP and PPTP networks.
	AuthProtocol *[]AuthProtocol `json:"AuthProtocol,omitempty" plist:"AuthProtocol,omitempty"`
	// If `1` and `CCPEnabled` is also `1`, enables CCPMPPE128 encryption.
	CCPMPPE40Enabled *CCPMPPE40Enabled `json:"CCPMPPE40Enabled,omitempty" plist:"CCPMPPE40Enabled,omitempty"`
	// If `1` and `CCPEnabled` is also `1`, enables CCPMPPE40 encryption.
	CCPMPPE128Enabled *CCPMPPE128Enabled `json:"CCPMPPE128Enabled,omitempty" plist:"CCPMPPE128Enabled,omitempty"`
	// If `1`, enables encryption on the connection. This key is for use with PPTP networks.
	CCPEnabled *CCPEnabled `json:"CCPEnabled,omitempty" plist:"CCPEnabled,omitempty"`
	// If `1`, disconnects after an on demand connection idles.
	DisconnectOnIdle *PPPDisconnectOnIdle `default:"0" json:"DisconnectOnIdle,omitempty" plist:"DisconnectOnIdle,omitempty"`
	// The length of time to wait before disconnecting an on demand connection
	DisconnectOnIdleTimer *int64 `json:"DisconnectOnIdleTimer,omitempty" plist:"DisconnectOnIdleTimer,omitempty"`
}

// If `1`, uses a token card such as an RSA SecurID card for connecting. This key is for use with L2TP networks.
type TokenCard int64

const (
	TokenCard0 TokenCard = 0
	TokenCard1 TokenCard = 1
)

// An array of authentication plugins. For use of RSA SecurID, this array should only have one value: `EAP-RSA`. This key is for use with L2TP and PPTP networks.
type AuthEAPPlugins string

const (
	AuthEAPPluginsEAPRSA AuthEAPPlugins = "EAP-RSA"
	AuthEAPPluginsEAPTLS AuthEAPPlugins = "EAP-TLS"
	AuthEAPPluginsEAPKRB AuthEAPPlugins = "EAP-KRB"
)

// An array of authentication protocols. For use of RSA SecurID, this array should have one value, `EAP`. This key is for use with L2TP and PPTP networks.
type AuthProtocol string

const (
	AuthProtocolEAP AuthProtocol = "EAP"
)

// If `1` and `CCPEnabled` is also `1`, enables CCPMPPE128 encryption.
type CCPMPPE40Enabled int64

const (
	CCPMPPE40Enabled0 CCPMPPE40Enabled = 0
	CCPMPPE40Enabled1 CCPMPPE40Enabled = 1
)

// If `1` and `CCPEnabled` is also `1`, enables CCPMPPE40 encryption.
type CCPMPPE128Enabled int64

const (
	CCPMPPE128Enabled0 CCPMPPE128Enabled = 0
	CCPMPPE128Enabled1 CCPMPPE128Enabled = 1
)

// If `1`, enables encryption on the connection. This key is for use with PPTP networks.
type CCPEnabled int64

const (
	CCPEnabled0 CCPEnabled = 0
	CCPEnabled1 CCPEnabled = 1
)

// If `1`, disconnects after an on demand connection idles.
type PPPDisconnectOnIdle int64

const (
	PPPDisconnectOnIdle0 PPPDisconnectOnIdle = 0
	PPPDisconnectOnIdle1 PPPDisconnectOnIdle = 1
)

// The dictionary that contains IPSec settings. Not available in watchOS.
type IPSec struct {
	// The IP address or host name of the VPN server.
	RemoteAddress *string `json:"RemoteAddress,omitempty" plist:"RemoteAddress,omitempty"`
	// The authentication method for L2TP and Cisco IPSec.
	AuthenticationMethod *IPSecAuthenticationMethod `default:"SharedSecret" json:"AuthenticationMethod,omitempty" plist:"AuthenticationMethod,omitempty"`
	// The user name for the VPN account for Cisco IPSec.
	XAuthName *string `json:"XAuthName,omitempty" plist:"XAuthName,omitempty"`
	// The VPN account password for Cisco IPSec.
	XAuthPassword *string `json:"XAuthPassword,omitempty" plist:"XAuthPassword,omitempty"`
	// If `1`, enables Xauth for Cisco IPSec VPNs.
	XAuthEnabled *XAuthEnabled `json:"XAuthEnabled,omitempty" plist:"XAuthEnabled,omitempty"`
	// A string that either has the value "Prompt" or isn't present.
	XAuthPasswordEncryption *XAuthPasswordEncryption `json:"XAuthPasswordEncryption,omitempty" plist:"XAuthPasswordEncryption,omitempty"`
	// The name of the group. For hybrid authentication, the string needs to end with "hybrid".
	// Present only for Cisco IPSec if `AuthenticationMethod` is `SharedSecret`.
	LocalIdentifier *string `json:"LocalIdentifier,omitempty" plist:"LocalIdentifier,omitempty"`
	// Present only if `AuthenticationMethod` is `SharedSecret`. The value is `KeyID`. The system uses this value for L2TP and Cisco IPSec VPNs.
	LocalIdentifierType *LocalIdentifierType `default:"KeyID" json:"LocalIdentifierType,omitempty" plist:"LocalIdentifierType,omitempty"`
	// The shared secret for this VPN account.
	// Only use this with L2TP and Cisco IPSec VPNs and if the `AuthenticationMethod` key is to `SharedSecret`.
	SharedSecret *[]byte `json:"SharedSecret,omitempty" plist:"SharedSecret,omitempty"`
	// The UUID of the certificate payload within the same profile to use for the account credentials.
	// Only use this with Cisco IPSec VPNs and if the `AuthenticationMethod` key is to `Certificate`.
	PayloadCertificateUUID *string `json:"PayloadCertificateUUID,omitempty" plist:"PayloadCertificateUUID,omitempty"`
	// If `true`, prompts for a PIN when connecting to Cisco IPSec VPNs.
	PromptForVPNPIN *bool `json:"PromptForVPNPIN,omitempty" plist:"PromptForVPNPIN,omitempty"`
	// If `1`, disconnect after an on-demand connection idles.
	DisconnectOnIdle *IPSecDisconnectOnIdle `default:"0" json:"DisconnectOnIdle,omitempty" plist:"DisconnectOnIdle,omitempty"`
	// The length of time to wait before disconnecting an on-demand connection.
	DisconnectOnIdleTimer *int64 `json:"DisconnectOnIdleTimer,omitempty" plist:"DisconnectOnIdleTimer,omitempty"`
	// If `1`, enables bringing the VPN connection up on demand.
	OnDemandEnabled *IPSecOnDemandEnabled `default:"0" json:"OnDemandEnabled,omitempty" plist:"OnDemandEnabled,omitempty"`
	// Deprecated. A list of domain names. In iOS 7 and later, if this key is present, the system treats associated domain names as though they're associated with the `OnDemandMatchDomainsOnRetry` key. This behavior can be overridden by `OnDemandRules`.
	OnDemandMatchDomainsAlways *[]string `json:"OnDemandMatchDomainsAlways,omitempty" plist:"OnDemandMatchDomainsAlways,omitempty"`
	// Deprecated. A list of domain names. In iOS 7 and later, this key is deprecated (but still supported) in favor of `EvaluateConnection` actions in the `OnDemandRules` dictionaries.
	OnDemandMatchDomainsNever *[]string `json:"OnDemandMatchDomainsNever,omitempty" plist:"OnDemandMatchDomainsNever,omitempty"`
	// Deprecated. A list of domain names. In iOS 7 and later, this field is deprecated (but still supported) in favor of `EvaluateConnection` actions in the `OnDemandRules` dictionaries.
	OnDemandMatchDomainsOnRetry *[]string `json:"OnDemandMatchDomainsOnRetry,omitempty" plist:"OnDemandMatchDomainsOnRetry,omitempty"`
	// The on-demand rules dictionary.
	OnDemandRules *[]*VPNOnDemandRulesOnDemandRulesElement `json:"OnDemandRules,omitempty" plist:"OnDemandRules,omitempty"`
}

// The authentication method for L2TP and Cisco IPSec.
type IPSecAuthenticationMethod string

const (
	IPSecAuthenticationMethodSharedSecret IPSecAuthenticationMethod = "SharedSecret"
	IPSecAuthenticationMethodCertificate  IPSecAuthenticationMethod = "Certificate"
)

// If `1`, enables Xauth for Cisco IPSec VPNs.
type XAuthEnabled int64

const (
	XAuthEnabled0 XAuthEnabled = 0
	XAuthEnabled1 XAuthEnabled = 1
)

// A string that either has the value "Prompt" or isn't present.
type XAuthPasswordEncryption string

const (
	XAuthPasswordEncryptionPrompt XAuthPasswordEncryption = "Prompt"
)

// Present only if `AuthenticationMethod` is `SharedSecret`. The value is `KeyID`. The system uses this value for L2TP and Cisco IPSec VPNs.
type LocalIdentifierType string

const (
	LocalIdentifierTypeKeyID LocalIdentifierType = "KeyID"
)

// If `1`, disconnect after an on-demand connection idles.
type IPSecDisconnectOnIdle int64

const (
	IPSecDisconnectOnIdle0 IPSecDisconnectOnIdle = 0
	IPSecDisconnectOnIdle1 IPSecDisconnectOnIdle = 1
)

// If `1`, enables bringing the VPN connection up on demand.
type IPSecOnDemandEnabled int64

const (
	IPSecOnDemandEnabled0 IPSecOnDemandEnabled = 0
	IPSecOnDemandEnabled1 IPSecOnDemandEnabled = 1
)

// The dictionary to use when `VPNType` is `IKEv2`.
type IKEv2 struct {
	// The IP address or host name of the VPN server.
	RemoteAddress string `json:"RemoteAddress" plist:"RemoteAddress" required:"true"`
	// Identifier of the IKEv2 client.
	LocalIdentifier string `json:"LocalIdentifier" plist:"LocalIdentifier" required:"true"`
	// The remote identifier.
	RemoteIdentifier string `json:"RemoteIdentifier" plist:"RemoteIdentifier" required:"true"`
	// The type of authentication method for the VPN.
	// To enable EAP-only authentication, set this to `None` and `ExtendedAuthEnabled` to `1`. If this is `None` and the `ExtendedAuthEnabled` key isn't set, the authentication configuration defaults to `SharedSecret`.
	AuthenticationMethod IKEv2AuthenticationMethod `json:"AuthenticationMethod" plist:"AuthenticationMethod" required:"true"`
	// The type of `PayloadCertificateUUID` to use for IKEv2 machine authentication. If this key is included, the system requires a value for `ServerCertificateIssuerCommonName`.
	CertificateType *CertificateType `default:"RSA" json:"CertificateType,omitempty" plist:"CertificateType,omitempty"`
	// The UUID of the certificate payload within the same profile to use as the account credential. If the value of `AuthenticationMethod` is `Certificate`, the system sends this certificate out for IKEv2 machine authentication. If extended authentication (EAP) is used, the system sends this certificate out for EAP-TLS authentication.
	PayloadCertificateUUID *string `json:"PayloadCertificateUUID,omitempty" plist:"PayloadCertificateUUID,omitempty"`
	// The password to use for the account credentials. Only used if `AuthenticationMethod` is `Password`.
	Password *string `json:"Password,omitempty" plist:"Password,omitempty"`
	// If the VPNSubType field contains the bundle identifier of an app that contains multiple VPN providers of the same type (app-proxy or packet-tunnel), then the system uses this field to choose which provider to use for this configuration. If the VPN provider is implemented as a System Extension, then this field is required.
	ProviderBundleIdentifier *string `json:"ProviderBundleIdentifier,omitempty" plist:"ProviderBundleIdentifier,omitempty"`
	// If the VPN provider is implemented as a System Extension, then this field is required. Available in macOS 10.15 and later, tvOS 17 and later, and watchOS 10 and later.
	ProviderDesignatedRequirement *string `json:"ProviderDesignatedRequirement,omitempty" plist:"ProviderDesignatedRequirement,omitempty"`
	// If `AuthenticationMethod` is `SharedSecret`, this value is used for IKE authentication.
	SharedSecret *string `json:"SharedSecret,omitempty" plist:"SharedSecret,omitempty"`
	// If `1`, enables EAP-only authentication.
	ExtendedAuthEnabled *ExtendedAuthEnabled `default:"0" json:"ExtendedAuthEnabled,omitempty" plist:"ExtendedAuthEnabled,omitempty"`
	// The user name to use for authentication.
	AuthName *string `json:"AuthName,omitempty" plist:"AuthName,omitempty"`
	// The password to use for authentication.
	AuthPassword *string `json:"AuthPassword,omitempty" plist:"AuthPassword,omitempty"`
	// If `1`, enables VPN up on demand.
	OnDemandEnabled *IKEv2OnDemandEnabled `default:"0" json:"OnDemandEnabled,omitempty" plist:"OnDemandEnabled,omitempty"`
	// If `1`, the system disables the Connect On Demand toggle in Settings for this configuration.
	OnDemandUserOverrideDisabled *IKEv2OnDemandUserOverrideDisabled `default:"0" json:"OnDemandUserOverrideDisabled,omitempty" plist:"OnDemandUserOverrideDisabled,omitempty"`
	// A list of rules that determine when and how to use an OnDemand VPN.
	OnDemandRules *[]*VPNOnDemandRulesOnDemandRulesElement `json:"OnDemandRules,omitempty" plist:"OnDemandRules,omitempty"`
	// One of the following:
	// - `None`: No keepalive.
	// - `Low`: Send keepalive every 30 minutes.
	// - `Medium`: Send keepalive every 10 minutes.
	// - `High`: Send keepalive every 1 minute.
	// Not available in watchOS.
	DeadPeerDetectionRate *DeadPeerDetectionRate `default:"Medium" json:"DeadPeerDetectionRate,omitempty" plist:"DeadPeerDetectionRate,omitempty"`
	// Common Name of the server certificate issuer. If set, this field causes IKE to send a certificate request based on this certificate issuer to the server. This key is required if the `CertificateType` key is included and the `ExtendedAuthEnabled` key is `1`.
	ServerCertificateIssuerCommonName *string `json:"ServerCertificateIssuerCommonName,omitempty" plist:"ServerCertificateIssuerCommonName,omitempty"`
	// The common name of the server certificate. The system uses this name to validate the certificate sent by the IKE server. If not set, the system uses the remote identifier to validate the certificate.
	ServerCertificateCommonName *string `json:"ServerCertificateCommonName,omitempty" plist:"ServerCertificateCommonName,omitempty"`
	// The minimum TLS version to use with EAP-TLS authentication.
	TLSMinimumVersion *IKEv2TLSMinimumVersion `default:"1.0" json:"TLSMinimumVersion,omitempty" plist:"TLSMinimumVersion,omitempty"`
	// The maximum TLS version to use with EAP-TLS authentication.
	TLSMaximumVersion *IKEv2TLSMaximumVersion `default:"1.2" json:"TLSMaximumVersion,omitempty" plist:"TLSMaximumVersion,omitempty"`
	// If `1`, negotiations should use IKEv2 Configuration Attribute `INTERNAL_IP4_SUBNET` and `INTERNAL_IP6_SUBNET`.
	UseConfigurationAttributeInternalIPSubnet *UseConfigurationAttributeInternalIPSubnet `default:"0" json:"UseConfigurationAttributeInternalIPSubnet,omitempty" plist:"UseConfigurationAttributeInternalIPSubnet,omitempty"`
	// If `1`, the system disables MOBIKE.
	DisableMOBIKE *DisableMOBIKE `default:"0" json:"DisableMOBIKE,omitempty" plist:"DisableMOBIKE,omitempty"`
	// If `1`, the system disables IKEv2 redirect. If not set, the system redirects an IKEv2 connection when it receives a redirect request from the server.
	DisableRedirect *DisableRedirect `default:"0" json:"DisableRedirect,omitempty" plist:"DisableRedirect,omitempty"`
	// If `1`, the VPN disconnects automatically after a period defined by `DisconnectOnIdleTimer`.
	DisconnectOnIdle *IKEv2DisconnectOnIdle `default:"0" json:"DisconnectOnIdle,omitempty" plist:"DisconnectOnIdle,omitempty"`
	// Only used if `DisconnectOnIdle` is `1`. The number of seconds before the VPN disconnects. On watchOS, maximum allowed value is 15 seconds
	DisconnectOnIdleTimer *int64 `json:"DisconnectOnIdleTimer,omitempty" plist:"DisconnectOnIdleTimer,omitempty"`
	// If `1`, enables NAT keepalive offload for Always On VPN IKEv2 connections. The device sends keepalive packets to maintain NAT mappings for IKEv2 connections that have a NAT on the path. It sends keepalive packets at regular intervals when the device is awake. If `NATKeepAliveOffloadEnable` is `1`, the system offloads keepalive packets to hardware while the device is asleep.
	// NAT keepalive offload has an impact on the battery life due to the extra workload during sleep. The default interval for the keepalive offload packets is 20 seconds over Wi-Fi and 110 seconds over Cellular interface. The default NAT keepalive works well on networks with small NAT mapping timeouts but imposes a potential battery impact. If a network has larger NAT mapping timeouts, larger keepalive intervals may be safely used to minimize battery impact. Modify the keepalive interval through the `NATKeepAliveInterval` key.
	NATKeepAliveOffloadEnable *NATKeepAliveOffloadEnable `default:"1" json:"NATKeepAliveOffloadEnable,omitempty" plist:"NATKeepAliveOffloadEnable,omitempty"`
	// The NAT Keepalive interval for Always On VPN IKEv2 connections. This value controls the interval that the device sends keepalive offload packets. The minimum value is 20 seconds. If no key is specified, the default is 20 seconds over Wi-Fi and 110 seconds over a cellular interface.
	NATKeepAliveInterval *int64 `default:"20" json:"NATKeepAliveInterval,omitempty" plist:"NATKeepAliveInterval,omitempty"`
	// If `1`,  enables Perfect Forward Secrecy (PFS) for IKEv2 Connections.
	EnablePFS *EnablePFS `default:"0" json:"EnablePFS,omitempty" plist:"EnablePFS,omitempty"`
	// If `1`, the system performs a certificate revocation check for IKEv2 connections. This is a best-effort revocation check and server response timeouts won't cause it to fail.
	EnableCertificateRevocationCheck *EnableCertificateRevocationCheck `default:"0" json:"EnableCertificateRevocationCheck,omitempty" plist:"EnableCertificateRevocationCheck,omitempty"`
	// If `1`, the system enables a tunnel over cellular data to carry traffic that's eligible for Wi-Fi Assist and also requires VPN.
	// Enabling fallback requires that the server support multiple tunnels for a single user.
	// This field is available in iOS 13 and later, and tvOS 17 and later. Not available in watchOS.
	EnableFallback *EnableFallback `default:"0" json:"EnableFallback,omitempty" plist:"EnableFallback,omitempty"`
	// The Maximum Transmission Unit (MTU) specifies the maximum size in bytes of each packet that the system sends over the IKEv2 VPN interface. Available in iOS 14 and later, and macOS 11 and later.
	MTU *int64 `default:"1280" json:"MTU,omitempty" plist:"MTU,omitempty"`
	// If the value of this key is `app-proxy`, the VPN service tunnels traffic at the application layer. If the value of this key is `packet-tunnel`, the VPN service tunnels traffic at the IP layer.
	ProviderType *IKEv2ProviderType `default:"packet-tunnel" json:"ProviderType,omitempty" plist:"ProviderType,omitempty"`
	// If `1`, then the system routes all network traffic through the VPN, with some controllable exclusions, such as `ExcludeLocalNetworks`, `ExcludeCellularServices`, and `ExcludeAPNs` properties. The system always excludes the following traffic from the tunnel:
	// - Traffic necessary for connecting and maintaining the device's network connection, such as DHCP.
	// - Traffic necessary for connecting to captive networks.
	// - Certain cellular services traffic that's not routable over the internet and is instead directly routed to the cellular network. See the `ExcludeCellularServices` field for more information.
	// - Network communication with a companion device such as a watchOS device.
	IncludeAllNetworks *IKEv2IncludeAllNetworks `default:"0" json:"IncludeAllNetworks,omitempty" plist:"IncludeAllNetworks,omitempty"`
	// If `1`, all the VPN's non-default routes take precedence over any locally-defined routes. If `IncludeAllNetworks` is `1`, the system ignores `EnforceRoutes`.
	EnforceRoutes *IKEv2EnforceRoutes `default:"0" json:"EnforceRoutes,omitempty" plist:"EnforceRoutes,omitempty"`
	// If `1` and either `IncludeAllNetworks` or `EnforceRoutes` are `1`, then the system routes local network traffic outside of the VPN. The default for this value is `0` on macOS and `1` on iOS.
	ExcludeLocalNetworks *IKEv2ExcludeLocalNetworks `json:"ExcludeLocalNetworks,omitempty" plist:"ExcludeLocalNetworks,omitempty"`
	// If `1` and `IncludeAllNetworks` is `1`, the system excludes internet-routable network traffic for cellular services (VoLTE, Wi-Fi Calling, IMS, MMS, Visual Voicemail, etc.) from the tunnel. Note that some cellular carriers route cellular services traffic directly to the carrier network, bypassing the internet. Such cellular services traffic is always excluded from the tunnel.
	ExcludeCellularServices *IKEv2ExcludeCellularServices `default:"1" json:"ExcludeCellularServices,omitempty" plist:"ExcludeCellularServices,omitempty"`
	// If `1` and `IncludeAllNetworks` is `1`, the system excludes network traffic for the Apple Push Notification service (APNs) from the tunnel.
	ExcludeAPNs *IKEv2ExcludeAPNs `default:"1" json:"ExcludeAPNs,omitempty" plist:"ExcludeAPNs,omitempty"`
	// If set to `1` and `IncludeAllNetworks` is set to `1`, the device excludes network traffic used for communicating with devices connected via USB or Wi-Fi from the tunnel.
	ExcludeDeviceCommunication *IKEv2ExcludeDeviceCommunication `default:"1" json:"ExcludeDeviceCommunication,omitempty" plist:"ExcludeDeviceCommunication,omitempty"`
	// The Post-quantum Pre-shared key (PPK) the device uses for this VPN. This key is is used with VPN servers that support RFC 8784. If this key is present `PPKIdentifier` must also be present.
	PPK *[]byte `json:"PPK,omitempty" plist:"PPK,omitempty"`
	// The identifier for the Post-quantum Pre-shared key (PPK) the device uses for this VPN. This key is is used with VPN servers that support RFC 8784. If this key is present `PPK` must also be present.
	PPKIdentifier *string `json:"PPKIdentifier,omitempty" plist:"PPKIdentifier,omitempty"`
	// If set to `1`, the VPN doesn't establish a connection if the server doesn't support RFC 8784 or doesn't accept the PPK identifier specified in `PPKIdentifier`. The device ignores this key if `PPK` and `PPKIdentifier` are not present.
	PPKMandatory *PPKMandatory `default:"1" json:"PPKMandatory,omitempty" plist:"PPKMandatory,omitempty"`
	// If set to `0`, the VPN doesn't establish a connection if the server does not support or doesn't allow post-quantum key exchanges. Thd device ignores this key if `PostQuantumKeyExchangeMethods` is not present in `IKESecurityAssociationParameters` or `ChildSecurityAssociationParameters`.
	AllowPostQuantumKeyExchangeFallback *AllowPostQuantumKeyExchangeFallback `default:"0" json:"AllowPostQuantumKeyExchangeFallback,omitempty" plist:"AllowPostQuantumKeyExchangeFallback,omitempty"`
	// If set to `1`, the device doesn't allow DES, 3DES, and Diffie-Hellman groups less than 14. Also the device requires the encryption algorithm specified for the IKE SA to be at least as cryptographically strong as the algorithm specified for the child SA. The device rejects this profile payload if these requirements are not met.
	EnforceStrictAlgorithmSelection *EnforceStrictAlgorithmSelection `default:"0" json:"EnforceStrictAlgorithmSelection,omitempty" plist:"EnforceStrictAlgorithmSelection,omitempty"`
	// These parameters apply to Child Security Association unless `ChildSecurityAssociationParameters` is specified.
	IKESecurityAssociationParameters *IKESecurityAssociationParameters `json:"IKESecurityAssociationParameters,omitempty" plist:"IKESecurityAssociationParameters,omitempty"`
	// The `ChildSecurityAssociationParameters` dictionaries.
	ChildSecurityAssociationParameters *ChildSecurityAssociationParameters `json:"ChildSecurityAssociationParameters,omitempty" plist:"ChildSecurityAssociationParameters,omitempty"`
}

// The type of authentication method for the VPN.
// To enable EAP-only authentication, set this to `None` and `ExtendedAuthEnabled` to `1`. If this is `None` and the `ExtendedAuthEnabled` key isn't set, the authentication configuration defaults to `SharedSecret`.
type IKEv2AuthenticationMethod string

const (
	IKEv2AuthenticationMethodNone         IKEv2AuthenticationMethod = "None"
	IKEv2AuthenticationMethodSharedSecret IKEv2AuthenticationMethod = "SharedSecret"
	IKEv2AuthenticationMethodCertificate  IKEv2AuthenticationMethod = "Certificate"
)

// The type of `PayloadCertificateUUID` to use for IKEv2 machine authentication. If this key is included, the system requires a value for `ServerCertificateIssuerCommonName`.
type CertificateType string

const (
	CertificateTypeRSA      CertificateType = "RSA"
	CertificateTypeECDSA256 CertificateType = "ECDSA256"
	CertificateTypeECDSA384 CertificateType = "ECDSA384"
	CertificateTypeECDSA521 CertificateType = "ECDSA521"
	CertificateTypeRSAPSS   CertificateType = "RSA-PSS"
)

// If `1`, enables EAP-only authentication.
type ExtendedAuthEnabled int64

const (
	ExtendedAuthEnabled0 ExtendedAuthEnabled = 0
	ExtendedAuthEnabled1 ExtendedAuthEnabled = 1
)

// If `1`, enables VPN up on demand.
type IKEv2OnDemandEnabled int64

const (
	IKEv2OnDemandEnabled0 IKEv2OnDemandEnabled = 0
	IKEv2OnDemandEnabled1 IKEv2OnDemandEnabled = 1
)

// If `1`, the system disables the Connect On Demand toggle in Settings for this configuration.
type IKEv2OnDemandUserOverrideDisabled int64

const (
	IKEv2OnDemandUserOverrideDisabled0 IKEv2OnDemandUserOverrideDisabled = 0
	IKEv2OnDemandUserOverrideDisabled1 IKEv2OnDemandUserOverrideDisabled = 1
)

// One of the following:
// - `None`: No keepalive.
// - `Low`: Send keepalive every 30 minutes.
// - `Medium`: Send keepalive every 10 minutes.
// - `High`: Send keepalive every 1 minute.
// Not available in watchOS.
type DeadPeerDetectionRate string

const (
	DeadPeerDetectionRateNone   DeadPeerDetectionRate = "None"
	DeadPeerDetectionRateLow    DeadPeerDetectionRate = "Low"
	DeadPeerDetectionRateMedium DeadPeerDetectionRate = "Medium"
	DeadPeerDetectionRateHigh   DeadPeerDetectionRate = "High"
)

// The minimum TLS version to use with EAP-TLS authentication.
type IKEv2TLSMinimumVersion string

const (
	IKEv2TLSMinimumVersion10 IKEv2TLSMinimumVersion = "1.0"
	IKEv2TLSMinimumVersion11 IKEv2TLSMinimumVersion = "1.1"
	IKEv2TLSMinimumVersion12 IKEv2TLSMinimumVersion = "1.2"
)

// The maximum TLS version to use with EAP-TLS authentication.
type IKEv2TLSMaximumVersion string

const (
	IKEv2TLSMaximumVersion10 IKEv2TLSMaximumVersion = "1.0"
	IKEv2TLSMaximumVersion11 IKEv2TLSMaximumVersion = "1.1"
	IKEv2TLSMaximumVersion12 IKEv2TLSMaximumVersion = "1.2"
)

// If `1`, negotiations should use IKEv2 Configuration Attribute `INTERNAL_IP4_SUBNET` and `INTERNAL_IP6_SUBNET`.
type UseConfigurationAttributeInternalIPSubnet int64

const (
	UseConfigurationAttributeInternalIPSubnet0 UseConfigurationAttributeInternalIPSubnet = 0
	UseConfigurationAttributeInternalIPSubnet1 UseConfigurationAttributeInternalIPSubnet = 1
)

// If `1`, the system disables MOBIKE.
type DisableMOBIKE int64

const (
	DisableMOBIKE0 DisableMOBIKE = 0
	DisableMOBIKE1 DisableMOBIKE = 1
)

// If `1`, the system disables IKEv2 redirect. If not set, the system redirects an IKEv2 connection when it receives a redirect request from the server.
type DisableRedirect int64

const (
	DisableRedirect0 DisableRedirect = 0
	DisableRedirect1 DisableRedirect = 1
)

// If `1`, the VPN disconnects automatically after a period defined by `DisconnectOnIdleTimer`.
type IKEv2DisconnectOnIdle int64

const (
	IKEv2DisconnectOnIdle0 IKEv2DisconnectOnIdle = 0
	IKEv2DisconnectOnIdle1 IKEv2DisconnectOnIdle = 1
)

// If `1`, enables NAT keepalive offload for Always On VPN IKEv2 connections. The device sends keepalive packets to maintain NAT mappings for IKEv2 connections that have a NAT on the path. It sends keepalive packets at regular intervals when the device is awake. If `NATKeepAliveOffloadEnable` is `1`, the system offloads keepalive packets to hardware while the device is asleep.
// NAT keepalive offload has an impact on the battery life due to the extra workload during sleep. The default interval for the keepalive offload packets is 20 seconds over Wi-Fi and 110 seconds over Cellular interface. The default NAT keepalive works well on networks with small NAT mapping timeouts but imposes a potential battery impact. If a network has larger NAT mapping timeouts, larger keepalive intervals may be safely used to minimize battery impact. Modify the keepalive interval through the `NATKeepAliveInterval` key.
type NATKeepAliveOffloadEnable int64

const (
	NATKeepAliveOffloadEnable0 NATKeepAliveOffloadEnable = 0
	NATKeepAliveOffloadEnable1 NATKeepAliveOffloadEnable = 1
)

// If `1`,  enables Perfect Forward Secrecy (PFS) for IKEv2 Connections.
type EnablePFS int64

const (
	EnablePFS0 EnablePFS = 0
	EnablePFS1 EnablePFS = 1
)

// If `1`, the system performs a certificate revocation check for IKEv2 connections. This is a best-effort revocation check and server response timeouts won't cause it to fail.
type EnableCertificateRevocationCheck int64

const (
	EnableCertificateRevocationCheck0 EnableCertificateRevocationCheck = 0
	EnableCertificateRevocationCheck1 EnableCertificateRevocationCheck = 1
)

// If `1`, the system enables a tunnel over cellular data to carry traffic that's eligible for Wi-Fi Assist and also requires VPN.
// Enabling fallback requires that the server support multiple tunnels for a single user.
// This field is available in iOS 13 and later, and tvOS 17 and later. Not available in watchOS.
type EnableFallback int64

const (
	EnableFallback0 EnableFallback = 0
	EnableFallback1 EnableFallback = 1
)

// If the value of this key is `app-proxy`, the VPN service tunnels traffic at the application layer. If the value of this key is `packet-tunnel`, the VPN service tunnels traffic at the IP layer.
type IKEv2ProviderType string

const (
	IKEv2ProviderTypePacketTunnel IKEv2ProviderType = "packet-tunnel"
	IKEv2ProviderTypeAppProxy     IKEv2ProviderType = "app-proxy"
)

// If `1`, then the system routes all network traffic through the VPN, with some controllable exclusions, such as `ExcludeLocalNetworks`, `ExcludeCellularServices`, and `ExcludeAPNs` properties. The system always excludes the following traffic from the tunnel:
// - Traffic necessary for connecting and maintaining the device's network connection, such as DHCP.
// - Traffic necessary for connecting to captive networks.
// - Certain cellular services traffic that's not routable over the internet and is instead directly routed to the cellular network. See the `ExcludeCellularServices` field for more information.
// - Network communication with a companion device such as a watchOS device.
type IKEv2IncludeAllNetworks int64

const (
	IKEv2IncludeAllNetworks0 IKEv2IncludeAllNetworks = 0
	IKEv2IncludeAllNetworks1 IKEv2IncludeAllNetworks = 1
)

// If `1`, all the VPN's non-default routes take precedence over any locally-defined routes. If `IncludeAllNetworks` is `1`, the system ignores `EnforceRoutes`.
type IKEv2EnforceRoutes int64

const (
	IKEv2EnforceRoutes0 IKEv2EnforceRoutes = 0
	IKEv2EnforceRoutes1 IKEv2EnforceRoutes = 1
)

// If `1` and either `IncludeAllNetworks` or `EnforceRoutes` are `1`, then the system routes local network traffic outside of the VPN. The default for this value is `0` on macOS and `1` on iOS.
type IKEv2ExcludeLocalNetworks int64

const (
	IKEv2ExcludeLocalNetworks0 IKEv2ExcludeLocalNetworks = 0
	IKEv2ExcludeLocalNetworks1 IKEv2ExcludeLocalNetworks = 1
)

// If `1` and `IncludeAllNetworks` is `1`, the system excludes internet-routable network traffic for cellular services (VoLTE, Wi-Fi Calling, IMS, MMS, Visual Voicemail, etc.) from the tunnel. Note that some cellular carriers route cellular services traffic directly to the carrier network, bypassing the internet. Such cellular services traffic is always excluded from the tunnel.
type IKEv2ExcludeCellularServices int64

const (
	IKEv2ExcludeCellularServices0 IKEv2ExcludeCellularServices = 0
	IKEv2ExcludeCellularServices1 IKEv2ExcludeCellularServices = 1
)

// If `1` and `IncludeAllNetworks` is `1`, the system excludes network traffic for the Apple Push Notification service (APNs) from the tunnel.
type IKEv2ExcludeAPNs int64

const (
	IKEv2ExcludeAPNs0 IKEv2ExcludeAPNs = 0
	IKEv2ExcludeAPNs1 IKEv2ExcludeAPNs = 1
)

// If set to `1` and `IncludeAllNetworks` is set to `1`, the device excludes network traffic used for communicating with devices connected via USB or Wi-Fi from the tunnel.
type IKEv2ExcludeDeviceCommunication int64

const (
	IKEv2ExcludeDeviceCommunication0 IKEv2ExcludeDeviceCommunication = 0
	IKEv2ExcludeDeviceCommunication1 IKEv2ExcludeDeviceCommunication = 1
)

// If set to `1`, the VPN doesn't establish a connection if the server doesn't support RFC 8784 or doesn't accept the PPK identifier specified in `PPKIdentifier`. The device ignores this key if `PPK` and `PPKIdentifier` are not present.
type PPKMandatory int64

const (
	PPKMandatory0 PPKMandatory = 0
	PPKMandatory1 PPKMandatory = 1
)

// If set to `0`, the VPN doesn't establish a connection if the server does not support or doesn't allow post-quantum key exchanges. Thd device ignores this key if `PostQuantumKeyExchangeMethods` is not present in `IKESecurityAssociationParameters` or `ChildSecurityAssociationParameters`.
type AllowPostQuantumKeyExchangeFallback int64

const (
	AllowPostQuantumKeyExchangeFallback0 AllowPostQuantumKeyExchangeFallback = 0
	AllowPostQuantumKeyExchangeFallback1 AllowPostQuantumKeyExchangeFallback = 1
)

// If set to `1`, the device doesn't allow DES, 3DES, and Diffie-Hellman groups less than 14. Also the device requires the encryption algorithm specified for the IKE SA to be at least as cryptographically strong as the algorithm specified for the child SA. The device rejects this profile payload if these requirements are not met.
type EnforceStrictAlgorithmSelection int64

const (
	EnforceStrictAlgorithmSelection0 EnforceStrictAlgorithmSelection = 0
	EnforceStrictAlgorithmSelection1 EnforceStrictAlgorithmSelection = 1
)

// These parameters apply to Child Security Association unless `ChildSecurityAssociationParameters` is specified.
type IKESecurityAssociationParameters struct {
	// The encryption algorithm.
	// In watchOS and tvOS, the default value is `AES-256-GCM`.
	// `DES` and `3DES` are available only in iOS, macOS, and visionOS prior to iOS 26, macOS 26, and visionOS 26.
	EncryptionAlgorithm *EncryptionAlgorithm `default:"AES-256" json:"EncryptionAlgorithm,omitempty" plist:"EncryptionAlgorithm,omitempty"`
	// The integrity algorithm.
	// `SHA1-96` and `SHA1-160` are available only in iOS, macOS, and visionOS prior to iOS 26, macOS 26, and visionOS 26.
	IntegrityAlgorithm *IntegrityAlgorithm `default:"SHA2-256" json:"IntegrityAlgorithm,omitempty" plist:"IntegrityAlgorithm,omitempty"`
	// The Diffie-Hellman group.
	// For `AlwaysOn` VPN in iOS 14.2 and later, the minimum allowed value is `14`.
	// `1`, `2`, and `5` are available only in iOS, macOS, and visionOS prior to iOS 26, macOS 26, and visionOS 26.
	DiffieHellmanGroup *DiffieHellmanGroup `default:"14" json:"DiffieHellmanGroup,omitempty" plist:"DiffieHellmanGroup,omitempty"`
	// An array of strings representing postquantum key exchange methods the device uses during SA establishment and rekey. You can specify up to seven items, which correspond to ADDKE1 - ADDKE7 from RFC 9370.
	PostQuantumKeyExchangeMethods *[]PostQuantumKeyExchangeMethods `json:"PostQuantumKeyExchangeMethods,omitempty" plist:"PostQuantumKeyExchangeMethods,omitempty"`
	// The SA lifetime (rekey interval) in minutes.
	LifeTimeInMinutes *int64 `default:"1440" json:"LifeTimeInMinutes,omitempty" plist:"LifeTimeInMinutes,omitempty"`
}

// The encryption algorithm.
// In watchOS and tvOS, the default value is `AES-256-GCM`.
// `DES` and `3DES` are available only in iOS, macOS, and visionOS prior to iOS 26, macOS 26, and visionOS 26.
type EncryptionAlgorithm string

const (
	EncryptionAlgorithmDES              EncryptionAlgorithm = "DES"
	EncryptionAlgorithm3DES             EncryptionAlgorithm = "3DES"
	EncryptionAlgorithmAES128           EncryptionAlgorithm = "AES-128"
	EncryptionAlgorithmAES256           EncryptionAlgorithm = "AES-256"
	EncryptionAlgorithmAES128GCM        EncryptionAlgorithm = "AES-128-GCM"
	EncryptionAlgorithmAES256GCM        EncryptionAlgorithm = "AES-256-GCM"
	EncryptionAlgorithmChaCha20Poly1305 EncryptionAlgorithm = "ChaCha20Poly1305"
)

// The integrity algorithm.
// `SHA1-96` and `SHA1-160` are available only in iOS, macOS, and visionOS prior to iOS 26, macOS 26, and visionOS 26.
type IntegrityAlgorithm string

const (
	IntegrityAlgorithmSHA196  IntegrityAlgorithm = "SHA1-96"
	IntegrityAlgorithmSHA1160 IntegrityAlgorithm = "SHA1-160"
	IntegrityAlgorithmSHA2256 IntegrityAlgorithm = "SHA2-256"
	IntegrityAlgorithmSHA2384 IntegrityAlgorithm = "SHA2-384"
	IntegrityAlgorithmSHA2512 IntegrityAlgorithm = "SHA2-512"
)

// The Diffie-Hellman group.
// For `AlwaysOn` VPN in iOS 14.2 and later, the minimum allowed value is `14`.
// `1`, `2`, and `5` are available only in iOS, macOS, and visionOS prior to iOS 26, macOS 26, and visionOS 26.
type DiffieHellmanGroup int64

const (
	DiffieHellmanGroup1  DiffieHellmanGroup = 1
	DiffieHellmanGroup2  DiffieHellmanGroup = 2
	DiffieHellmanGroup5  DiffieHellmanGroup = 5
	DiffieHellmanGroup14 DiffieHellmanGroup = 14
	DiffieHellmanGroup15 DiffieHellmanGroup = 15
	DiffieHellmanGroup16 DiffieHellmanGroup = 16
	DiffieHellmanGroup17 DiffieHellmanGroup = 17
	DiffieHellmanGroup18 DiffieHellmanGroup = 18
	DiffieHellmanGroup19 DiffieHellmanGroup = 19
	DiffieHellmanGroup20 DiffieHellmanGroup = 20
	DiffieHellmanGroup21 DiffieHellmanGroup = 21
	DiffieHellmanGroup31 DiffieHellmanGroup = 31
	DiffieHellmanGroup32 DiffieHellmanGroup = 32
)

// An array of strings representing postquantum key exchange methods the device uses during SA establishment and rekey. You can specify up to seven items, which correspond to ADDKE1 - ADDKE7 from RFC 9370.
type PostQuantumKeyExchangeMethods int64

const (
	PostQuantumKeyExchangeMethods0  PostQuantumKeyExchangeMethods = 0
	PostQuantumKeyExchangeMethods36 PostQuantumKeyExchangeMethods = 36
	PostQuantumKeyExchangeMethods37 PostQuantumKeyExchangeMethods = 37
)

// The `ChildSecurityAssociationParameters` dictionaries.
type ChildSecurityAssociationParameters struct {
	// The encryption algorithm.
	// In watchOS and tvOS, the default value is `AES-256-GCM`.
	// `DES` and `3DES` are available only in iOS, macOS, and visionOS prior to iOS 26, macOS 26, and visionOS 26.
	EncryptionAlgorithm *EncryptionAlgorithm `default:"AES-256" json:"EncryptionAlgorithm,omitempty" plist:"EncryptionAlgorithm,omitempty"`
	// The integrity algorithm.
	// `SHA1-96` and `SHA1-160` are available only in iOS, macOS, and visionOS prior to iOS 26, macOS 26, and visionOS 26.
	IntegrityAlgorithm *IntegrityAlgorithm `default:"SHA2-256" json:"IntegrityAlgorithm,omitempty" plist:"IntegrityAlgorithm,omitempty"`
	// The Diffie-Hellman group.
	// For `AlwaysOn` VPN in iOS 14.2 and later, the minimum allowed value is `14`.
	// `1`, `2`, and `5` are available only in iOS, macOS, and visionOS prior to iOS 26, macOS 26, and visionOS 26.
	DiffieHellmanGroup *DiffieHellmanGroup `default:"14" json:"DiffieHellmanGroup,omitempty" plist:"DiffieHellmanGroup,omitempty"`
	// An array of strings representing postquantum key exchange methods the device uses during SA establishment and rekey. You can specify up to seven items, which correspond to ADDKE1 - ADDKE7 from RFC 9370.
	PostQuantumKeyExchangeMethods *[]PostQuantumKeyExchangeMethods `json:"PostQuantumKeyExchangeMethods,omitempty" plist:"PostQuantumKeyExchangeMethods,omitempty"`
	// The SA lifetime (rekey interval) in minutes.
	LifeTimeInMinutes *int64 `default:"1440" json:"LifeTimeInMinutes,omitempty" plist:"LifeTimeInMinutes,omitempty"`
}

// A dictionary to use for all VPN types.
type DNS struct {
	// The transport protocol to communicate with the DNS server.
	DNSProtocol DNSDNSProtocol `json:"DNSProtocol" plist:"DNSProtocol" required:"true"`
	// The URI template of a DNS-over-HTTPS server, as defined in RFC 8484, which needs to use the `https://` scheme. The system uses the hostname or address in the URL to validate the server certificate. If `ServerAddresses` isn't specified, the system uses the hostname or address in the URL to determine the server addresses. This key is required if the `DNSProtocol` is `HTTPS`.
	ServerURL *string `json:"ServerURL,omitempty" plist:"ServerURL,omitempty"`
	// The hostname of a DNS-over-TLS server to validate the server certificate, as defined in RFC 7858. If `ServerAddresses` isn't specified, the system uses the hostname to determine the server addresses. This key is required if the `DNSProtocol` is `TLS`.
	ServerName *string `json:"ServerName,omitempty" plist:"ServerName,omitempty"`
	// The array of DNS server IP address strings. These IP addresses can be a mixture of IPv4 and IPv6 addresses.
	ServerAddresses []string `json:"ServerAddresses" plist:"ServerAddresses" required:"true"`
	// The list of domain strings used to fully qualify single-label host names.
	SearchDomains *[]string `json:"SearchDomains,omitempty" plist:"SearchDomains,omitempty"`
	// The primary domain of the tunnel.
	DomainName *string `json:"DomainName,omitempty" plist:"DomainName,omitempty"`
	// The list of domain strings used to determine which DNS queries use the DNS resolver settings in `ServerAddresses`. The system uses this key to create a split DNS configuration where it resolves only hosts in certain domains using the tunnel's DNS resolver. The system uses the default resolver for hosts that aren't in one of the domains in this list.
	// If `SupplementalMatchDomains` contains the empty string it becomes the default domain.
	// Split-tunnel configurations can direct all DNS queries to the VPN DNS servers before the primary DNS servers. If the VPN tunnel becomes the network's default route, the servers listed in `ServerAddresses` become the default resolver and the system ignores the `SupplementalMatchDomains` list.
	SupplementalMatchDomains *[]string `json:"SupplementalMatchDomains,omitempty" plist:"SupplementalMatchDomains,omitempty"`
	// If `0`, append the domains in the `SupplementalMatchDomains` list to the resolver's list of search domains.
	SupplementalMatchDomainsNoSearch *SupplementalMatchDomainsNoSearch `default:"0" json:"SupplementalMatchDomainsNoSearch,omitempty" plist:"SupplementalMatchDomainsNoSearch,omitempty"`
	// That UUID that points to an identity certificate payload. The system uses this identity to authenticate the user to the DNS resolver.
	PayloadCertificateUUID *string `json:"PayloadCertificateUUID,omitempty" plist:"PayloadCertificateUUID,omitempty"`
}

// The transport protocol to communicate with the DNS server.
type DNSDNSProtocol string

const (
	DNSDNSProtocolCleartext DNSDNSProtocol = "Cleartext"
	DNSDNSProtocolHTTPS     DNSDNSProtocol = "HTTPS"
	DNSDNSProtocolTLS       DNSDNSProtocol = "TLS"
)

// If `0`, append the domains in the `SupplementalMatchDomains` list to the resolver's list of search domains.
type SupplementalMatchDomainsNoSearch int64

const (
	SupplementalMatchDomainsNoSearch0 SupplementalMatchDomainsNoSearch = 0
	SupplementalMatchDomainsNoSearch1 SupplementalMatchDomainsNoSearch = 1
)

// The dictionary to use to configure `Proxies` for use with `VPN`.
type VPNProxies struct {
	// If `true`, enables automatic proxy configuration.
	ProxyAutoConfigEnable *ProxyAutoConfigEnable `json:"ProxyAutoConfigEnable,omitempty" plist:"ProxyAutoConfigEnable,omitempty"`
	// If `true`, enables proxy auto discovery.
	ProxyAutoDiscoveryEnable *ProxyAutoDiscoveryEnable `default:"1" json:"ProxyAutoDiscoveryEnable,omitempty" plist:"ProxyAutoDiscoveryEnable,omitempty"`
	// The URL to the location of the proxy auto-configuration file. Used only when `ProxyAutoConfigEnable` is `true`.
	ProxyAutoConfigURLString *string `json:"ProxyAutoConfigURLString,omitempty" plist:"ProxyAutoConfigURLString,omitempty"`
	// An array of domains that defines which hosts use proxy settings for hosts.
	SupplementalMatchDomains *[]string `json:"SupplementalMatchDomains,omitempty" plist:"SupplementalMatchDomains,omitempty"`
	// If `1`, enables proxy for HTTP traffic.
	HTTPEnable *HTTPEnable `default:"0" json:"HTTPEnable,omitempty" plist:"HTTPEnable,omitempty"`
	// The host name of the HTTP proxy.
	HTTPProxy *string `json:"HTTPProxy,omitempty" plist:"HTTPProxy,omitempty"`
	// The port number of the HTTP proxy. This field is required if `HTTPProxy` is specified.
	HTTPPort *int64 `json:"HTTPPort,omitempty" plist:"HTTPPort,omitempty"`
	// The user name used for authentication.
	HTTPProxyUsername *string `json:"HTTPProxyUsername,omitempty" plist:"HTTPProxyUsername,omitempty"`
	// The password used for authentication.
	HTTPProxyPassword *string `json:"HTTPProxyPassword,omitempty" plist:"HTTPProxyPassword,omitempty"`
	// If `true`, enables proxy for HTTPS traffic.
	HTTPSEnable *HTTPSEnable `default:"0" json:"HTTPSEnable,omitempty" plist:"HTTPSEnable,omitempty"`
	// The host name of the HTTPS proxy.
	HTTPSProxy *string `json:"HTTPSProxy,omitempty" plist:"HTTPSProxy,omitempty"`
	// The port number of the HTTPS proxy. This field is required if `HTTPSProxy` is specified.
	HTTPSPort *int64 `json:"HTTPSPort,omitempty" plist:"HTTPSPort,omitempty"`
}

// If `true`, enables automatic proxy configuration.
type ProxyAutoConfigEnable int64

const (
	ProxyAutoConfigEnable0 ProxyAutoConfigEnable = 0
	ProxyAutoConfigEnable1 ProxyAutoConfigEnable = 1
)

// If `true`, enables proxy auto discovery.
type ProxyAutoDiscoveryEnable int64

const (
	ProxyAutoDiscoveryEnable0 ProxyAutoDiscoveryEnable = 0
	ProxyAutoDiscoveryEnable1 ProxyAutoDiscoveryEnable = 1
)

// If `1`, enables proxy for HTTP traffic.
type HTTPEnable int64

const (
	HTTPEnable0 HTTPEnable = 0
	HTTPEnable1 HTTPEnable = 1
)

// If `true`, enables proxy for HTTPS traffic.
type HTTPSEnable int64

const (
	HTTPSEnable0 HTTPSEnable = 0
	HTTPSEnable1 HTTPSEnable = 1
)

// The dictionary to use when `VPNType` is `AlwaysOn`. Not available in tvOS or watchOS.
type AlwaysOn struct {
	// If `1`, allows the user to disable the VPN configuration.
	UIToggleEnabled *UIToggleEnabled `default:"0" json:"UIToggleEnabled,omitempty" plist:"UIToggleEnabled,omitempty"`
	// An array that contains an arbitrary number of tunnel configurations.
	TunnelConfigurations []*TunnelConfigurationElement `json:"TunnelConfigurations" plist:"TunnelConfigurations" required:"true"`
	// An array that contains an arbitrary number of service exceptions.
	ServiceExceptions *[]*ServiceExceptionElement `json:"ServiceExceptions,omitempty" plist:"ServiceExceptions,omitempty"`
	// An array that contains an arbitrary number of apps whose connections occur outside the VPN.
	ApplicationExceptions *[]*ApplicationExceptionElement `json:"ApplicationExceptions,omitempty" plist:"ApplicationExceptions,omitempty"`
	// If `1`, allows traffic from Captive Web Sheet outside the VPN tunnel.
	AllowCaptiveWebSheet *AllowCaptiveWebSheet `default:"0" json:"AllowCaptiveWebSheet,omitempty" plist:"AllowCaptiveWebSheet,omitempty"`
	// If `1`, allows traffic from all captive networking apps outside the VPN tunnel to perform captive network handling.
	AllowAllCaptiveNetworkPlugins *AllowAllCaptiveNetworkPlugins `default:"0" json:"AllowAllCaptiveNetworkPlugins,omitempty" plist:"AllowAllCaptiveNetworkPlugins,omitempty"`
	// The array of captive networking apps whose traffic is allowed outside the VPN tunnel, to perform captive network handling. Used only when `AllowAllCaptiveNetworkPlugins` is `false`.
	AllowedCaptiveNetworkPlugins *[]*AllowedCaptiveNetworkPluginElement `json:"AllowedCaptiveNetworkPlugins,omitempty" plist:"AllowedCaptiveNetworkPlugins,omitempty"`
}

// If `1`, allows the user to disable the VPN configuration.
type UIToggleEnabled int64

const (
	UIToggleEnabled0 UIToggleEnabled = 0
	UIToggleEnabled1 UIToggleEnabled = 1
)

type TunnelConfigurationElement struct {
	// The type of connection, which needs to be `IKEv2`.
	ProtocolType ProtocolType `json:"ProtocolType" plist:"ProtocolType" required:"true"`
	// The interfaces to apply this configuration to.
	Interfaces *[]Interfaces `json:"Interfaces,omitempty" plist:"Interfaces,omitempty"`
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

type ServiceExceptionElement struct {
	// The name of a service that's exempt from Always On VPN.
	// `CellularServices` is available in iOS 11.3 and later; it exempts `VoLTE`, `IMS` and `MMS`. WiFiCalling is exempted in iOS 13.4 and later.
	// `DeviceCommunication` is available in iOS 17.4 and later; it exempts network traffic used for communicating with devices connected via USB or Wi-Fi.
	ServiceName ServiceName `json:"ServiceName" plist:"ServiceName" required:"true"`
	// The action to take with network connections from the named service.
	Action ServiceExceptionElementAction `json:"Action" plist:"Action" required:"true"`
}

// The name of a service that's exempt from Always On VPN.
// `CellularServices` is available in iOS 11.3 and later; it exempts `VoLTE`, `IMS` and `MMS`. WiFiCalling is exempted in iOS 13.4 and later.
// `DeviceCommunication` is available in iOS 17.4 and later; it exempts network traffic used for communicating with devices connected via USB or Wi-Fi.
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
	BundleIdentifier string `json:"BundleIdentifier" plist:"BundleIdentifier" required:"true"`
	// Limit the exception to only the specified list of protocols, with support for `UDP` only.
	LimitToProtocols *[]LimitToProtocols `json:"LimitToProtocols,omitempty" plist:"LimitToProtocols,omitempty"`
}

// Limit the exception to only the specified list of protocols, with support for `UDP` only.
type LimitToProtocols string

const (
	LimitToProtocolsUDP LimitToProtocols = "UDP"
)

// If `1`, allows traffic from Captive Web Sheet outside the VPN tunnel.
type AllowCaptiveWebSheet int64

const (
	AllowCaptiveWebSheet0 AllowCaptiveWebSheet = 0
	AllowCaptiveWebSheet1 AllowCaptiveWebSheet = 1
)

// If `1`, allows traffic from all captive networking apps outside the VPN tunnel to perform captive network handling.
type AllowAllCaptiveNetworkPlugins int64

const (
	AllowAllCaptiveNetworkPlugins0 AllowAllCaptiveNetworkPlugins = 0
	AllowAllCaptiveNetworkPlugins1 AllowAllCaptiveNetworkPlugins = 1
)

type AllowedCaptiveNetworkPluginElement struct {
	// The bundle identifier for the app that's allowed on the captive network.
	BundleIdentifier string `json:"BundleIdentifier" plist:"BundleIdentifier" required:"true"`
}

// The dictionary to use when `VPNType` is `TransparentProxy`. Available in macOS 14 and later.
type TransparentProxy struct {
	// The type of authentication method to use: `Password`, `Certificate`, or `Password+Certificate`.
	// Available in macOS 14 and later.
	AuthenticationMethod *TransparentProxyAuthenticationMethod `default:"Password" json:"AuthenticationMethod,omitempty" plist:"AuthenticationMethod,omitempty"`
	// If `1`, the VPN disconnects automatically disconnect after a period defined by `DisconnectOnIdleTimer`.
	// Available in macOS 14 and later.
	DisconnectOnIdle *TransparentProxyDisconnectOnIdle `default:"0" json:"DisconnectOnIdle,omitempty" plist:"DisconnectOnIdle,omitempty"`
	// The number of seconds before the VPN disconnects. This value is only used if `DisconnectOnIdle` is `1`.
	// Available in macOS 14 and later.
	DisconnectOnIdleTimer *int64 `json:"DisconnectOnIdleTimer,omitempty" plist:"DisconnectOnIdleTimer,omitempty"`
	// If `1`, then all the VPN's non-default routes take precedence over any locally-defined routes. If `IncludeAllNetworks` is `1`, the system ignores the value of `EnforceRoutes`.
	// Available in macOS 14 and later.
	EnforceRoutes *TransparentProxyEnforceRoutes `default:"0" json:"EnforceRoutes,omitempty" plist:"EnforceRoutes,omitempty"`
	// If `1`, the system brings up the VPN on demand.
	// Available in macOS 14 and later.
	OnDemandEnabled *TransparentProxyOnDemandEnabled `default:"0" json:"OnDemandEnabled,omitempty" plist:"OnDemandEnabled,omitempty"`
	// Determines when and how the system uses an OnDemand VPN.
	// Available in macOS 14 and later.
	OnDemandRules *[]*VPNOnDemandRulesOnDemandRulesElement `json:"OnDemandRules,omitempty" plist:"OnDemandRules,omitempty"`
	// The UUID of the identity certificate as the account credential. If `AuthenticationMethod` is `Certificate`, and extended authentication (EAP) isn't used, this certificate is sent out for IKE client authentication. If extended authentication is used, this certificate can be used for EAP-TLS.
	// Available in macOS 14 and later.
	PayloadCertificateUUID *string `json:"PayloadCertificateUUID,omitempty" plist:"PayloadCertificateUUID,omitempty"`
	// The password to use for the account credentials. Only used if `AuthenticationMethod` is `Password`.
	// Available in macOS 14 and later.
	Password *string `json:"Password,omitempty" plist:"Password,omitempty"`
	// If the VPNSubType field contains the bundle identifier of an app that contains multiple VPN providers of the same type (app-proxy or packet-tunnel), then the system uses this field to choose which provider to use for this configuration. If the VPN provider is implemented as a System Extension, then this field is required.
	// Available in macOS 14 and later.
	ProviderBundleIdentifier *string `json:"ProviderBundleIdentifier,omitempty" plist:"ProviderBundleIdentifier,omitempty"`
	// If the VPN provider is implemented as a System Extension, then this field is required.
	// Available in macOS 14 and later.
	ProviderDesignatedRequirement *string `json:"ProviderDesignatedRequirement,omitempty" plist:"ProviderDesignatedRequirement,omitempty"`
	// If the value of this key is `app-proxy`, the VPN service tunnels traffic at the application layer. If the value of this key is `packet-tunnel`, the VPN service tunnels traffic at the IP layer.
	// Available in macOS 14 and later.
	ProviderType *TransparentProxyProviderType `default:"packet-tunnel" json:"ProviderType,omitempty" plist:"ProviderType,omitempty"`
	// A positive integer.
	// Available in macOS 14 and later.
	Order *int64 `json:"Order,omitempty" plist:"Order,omitempty"`
}

// The type of authentication method to use: `Password`, `Certificate`, or `Password+Certificate`.
// Available in macOS 14 and later.
type TransparentProxyAuthenticationMethod string

const (
	TransparentProxyAuthenticationMethodPassword            TransparentProxyAuthenticationMethod = "Password"
	TransparentProxyAuthenticationMethodCertificate         TransparentProxyAuthenticationMethod = "Certificate"
	TransparentProxyAuthenticationMethodPasswordCertificate TransparentProxyAuthenticationMethod = "Password+Certificate"
)

// If `1`, the VPN disconnects automatically disconnect after a period defined by `DisconnectOnIdleTimer`.
// Available in macOS 14 and later.
type TransparentProxyDisconnectOnIdle int64

const (
	TransparentProxyDisconnectOnIdle0 TransparentProxyDisconnectOnIdle = 0
	TransparentProxyDisconnectOnIdle1 TransparentProxyDisconnectOnIdle = 1
)

// If `1`, then all the VPN's non-default routes take precedence over any locally-defined routes. If `IncludeAllNetworks` is `1`, the system ignores the value of `EnforceRoutes`.
// Available in macOS 14 and later.
type TransparentProxyEnforceRoutes int64

const (
	TransparentProxyEnforceRoutes0 TransparentProxyEnforceRoutes = 0
	TransparentProxyEnforceRoutes1 TransparentProxyEnforceRoutes = 1
)

// If `1`, the system brings up the VPN on demand.
// Available in macOS 14 and later.
type TransparentProxyOnDemandEnabled int64

const (
	TransparentProxyOnDemandEnabled0 TransparentProxyOnDemandEnabled = 0
	TransparentProxyOnDemandEnabled1 TransparentProxyOnDemandEnabled = 1
)

// If the value of this key is `app-proxy`, the VPN service tunnels traffic at the application layer. If the value of this key is `packet-tunnel`, the VPN service tunnels traffic at the IP layer.
// Available in macOS 14 and later.
type TransparentProxyProviderType string

const (
	TransparentProxyProviderTypePacketTunnel TransparentProxyProviderType = "packet-tunnel"
	TransparentProxyProviderTypeAppProxy     TransparentProxyProviderType = "app-proxy"
)

// The profile that configures web clips on the device.
type WebClip struct {
	*CommonPayloadKeys
	// If `true`, the system prevents SpringBoard from adding shine to the icon.
	Precomposed *bool `json:"Precomposed,omitempty" plist:"Precomposed,omitempty"`
	// If `true`, the system launches the web clip as a full-screen web app.
	FullScreen *bool `json:"FullScreen,omitempty" plist:"FullScreen,omitempty"`
	// The URL of the web clip.
	URL string `json:"URL" plist:"URL" required:"true"`
	// The PNG icon to show on the Home Screen. If not set, the system displays a white square. For best results, provide a square image that's no larger than 400 x 400 pixels and less than 1 MB when uncompressed. The graphics file is automatically scaled and cropped to fit, if necessary, and converted to PNG format. Web clip icons are 144 x 144 pixels for iPad devices with a Retina display, and 114 x 114 pixels for iPhone devices. To prevent the device from adding a shine to the image, set `Precomposed` to `true`.
	Icon *[]byte `json:"Icon,omitempty" plist:"Icon,omitempty"`
	// If `true`, the system enables removing the web clip.
	IsRemovable *bool `json:"IsRemovable,omitempty" plist:"IsRemovable,omitempty"`
	// The name of the web clip that the system displays on the Home Screen.
	Label string `json:"Label" plist:"Label" required:"true"`
	// If `true`, a full screen web clip can navigate to an external web site without showing Safari UI. Otherwise, Safari UI appears when navigating away from the web clip's URL. This key has no effect when `FullScreen` is `false`. Available in iOS 14 and later.
	IgnoreManifestScope *bool `json:"IgnoreManifestScope,omitempty" plist:"IgnoreManifestScope,omitempty"`
	// The application bundle identifier of the application that opens the URL. To use this property, install the profile through MDM. Available in iOS 14 and later.
	TargetApplicationBundleIdentifier *string `json:"TargetApplicationBundleIdentifier,omitempty" plist:"TargetApplicationBundleIdentifier,omitempty"`
}

func (p *WebClip) PayloadType() string {
	return "com.apple.webClip.managed"
}

// The payload that configures web content filters.
// As of iOS 16.0 and visionOS 1.1, this can be installed on unsupervised devices and user enrollments if ContentFilterUUID is specified. Previously it could only be installed on supervised devices.
type WebContentFilter struct {
	*CommonPayloadKeys
	// The type of filter, built-in or plug-in. In macOS, the system only supports the plug-in value.
	FilterType *FilterType `default:"BuiltIn" json:"FilterType,omitempty" plist:"FilterType,omitempty"`
	// If `true`, this payload enforces a policy which requires retention of browsing history. This causes Safari to disable clearing of browsing history, and prevents the use of private browsing mode because that mode doesn't keep browsing history.
	SafariHistoryRetentionEnabled *bool `json:"SafariHistoryRetentionEnabled,omitempty" plist:"SafariHistoryRetentionEnabled,omitempty"`
	// If `true`, the system enables automatic filtering. Use when `FilterType` is `BuiltIn`.
	AutoFilterEnabled *bool `json:"AutoFilterEnabled,omitempty" plist:"AutoFilterEnabled,omitempty"`
	// An array or URLs that are accessible whether or not the automatic filter allows access. Use when `FilterType` is `BuiltIn`. Requires that `AutoFilterEnabled` is `true`.
	PermittedURLs *[]string `json:"PermittedURLs,omitempty" plist:"PermittedURLs,omitempty"`
	// Use `DenyListURLs` instead.
	BlacklistedURLs *[]string `json:"BlacklistedURLs,omitempty" plist:"BlacklistedURLs,omitempty"`
	// An array of URLs that are inaccessible. Use when `FilterType` is `BuiltIn`. Limit the number of these URLs to no more than 500.
	DenyListURLs *[]string `json:"DenyListURLs,omitempty" plist:"DenyListURLs,omitempty"`
	// If `true`, the device hides the `DenyListURLs` item in the profiles that display in Settings > General > VPN & Device Management.
	HideDenyListURLs *bool `json:"HideDenyListURLs,omitempty" plist:"HideDenyListURLs,omitempty"`
	// Use `AllowListBookmarks` instead.
	WhitelistedBookmarks *[]*WhitelistedBookmarksItem `json:"WhitelistedBookmarks,omitempty" plist:"WhitelistedBookmarks,omitempty"`
	// An array of dictionaries that define the pages that the user can bookmark or visit. Use when `FilterType` is `BuiltIn`.
	AllowListBookmarks *[]*AllowListBookmarksItem `json:"AllowListBookmarks,omitempty" plist:"AllowListBookmarks,omitempty"`
	// The display name for this filtering configuration. Required when `FilterType` is `Plugin`.
	UserDefinedName *string `json:"UserDefinedName,omitempty" plist:"UserDefinedName,omitempty"`
	// The bundle ID of the plug-in that provides filtering service. Required when `FilterType` is `Plugin`. Otherwise, it ignores this value. Consult your filtering solution vendor to determine what to specify for this value. Required when `FilterType` is `Plugin`.
	PluginBundleID *string `json:"PluginBundleID,omitempty" plist:"PluginBundleID,omitempty"`
	// The server address, which may be the IP address, hostname, or URL. Use when `FilterType` is `Plugin`.
	ServerAddress *string `json:"ServerAddress,omitempty" plist:"ServerAddress,omitempty"`
	// The user name for the service. Use when `FilterType` is `Plugin`.
	UserName *string `json:"UserName,omitempty" plist:"UserName,omitempty"`
	// The password for the service. Use when `FilterType` is `Plugin`.
	Password *string `json:"Password,omitempty" plist:"Password,omitempty"`
	// The UUID of the certificate payload within the same profile that the system uses to authenticate the user. Use when `FilterType` is `Plugin`.
	PayloadCertificateUUID *string `json:"PayloadCertificateUUID,omitempty" plist:"PayloadCertificateUUID,omitempty"`
	// The organization string to pass to the third-party plug-in. Use when `FilterType` is `Plugin`.
	Organization *string `json:"Organization,omitempty" plist:"Organization,omitempty"`
	// The custom dictionary that the filtering service plug-in needs. Use when `FilterType` is `Plugin`.
	VendorConfig *map[string]any `json:"VendorConfig,omitempty" plist:"VendorConfig,omitempty"`
	// If `true`, the system enables filtering WebKit traffic. Use when `FilterType` is `Plugin`.
	// > Note:
	// > At least one of `FilterBrowsers` or `FilterSockets` needs to be `true`.
	FilterBrowsers *bool `json:"FilterBrowsers,omitempty" plist:"FilterBrowsers,omitempty"`
	// If `true`, enables the filtering of socket traffic. Use when `FilterType` is `Plugin`.
	// > Note:
	// > At least one of `FilterBrowsers` or `FilterSockets` needs to be `true`.
	FilterSockets *bool `json:"FilterSockets,omitempty" plist:"FilterSockets,omitempty"`
	// The designated requirement string that the system embeds in the code signature of the filter data provider system extension. This string identifies the filter data provider when the filter starts running. Required if `FilterSockets` is `true`.
	FilterDataProviderDesignatedRequirement *string `json:"FilterDataProviderDesignatedRequirement,omitempty" plist:"FilterDataProviderDesignatedRequirement,omitempty"`
	// The bundle identifier string of the filter data provider system extension. This string identifies the filter data provider when the filter starts running. Required if `FilterSockets` is `true`.
	FilterDataProviderBundleIdentifier *string `json:"FilterDataProviderBundleIdentifier,omitempty" plist:"FilterDataProviderBundleIdentifier,omitempty"`
	// If `true` and `FilterType` is `Plugin`, the system enables filtering network packets. Use when `FilterType` is `Plugin`.
	// > Note:
	// > At least one of `FilterPackets` or `FilterSockets` needs to be `true`.
	FilterPackets *bool `json:"FilterPackets,omitempty" plist:"FilterPackets,omitempty"`
	// The designated requirement string that the system embeds in the code signature of the filter packet provider system extension. This string identifies the filter packet provider when the filter starts running. Required if `FilterPackets` is `true`.
	FilterPacketProviderDesignatedRequirement *string `json:"FilterPacketProviderDesignatedRequirement,omitempty" plist:"FilterPacketProviderDesignatedRequirement,omitempty"`
	// The bundle identifier string of the filter packet provider system extension. This string identifies the filter packet provider when the filter starts running. Required if `FilterPackets` is `true`.
	FilterPacketProviderBundleIdentifier *string `json:"FilterPacketProviderBundleIdentifier,omitempty" plist:"FilterPacketProviderBundleIdentifier,omitempty"`
	// The system uses this value to derive the relative order of content filters. Filters with a grade of `firewall` see network traffic before filters with a grade of `inspector`. However, the system doesn't define the order of filters within a grade.
	FilterGrade *FilterGrade `default:"firewall" json:"FilterGrade,omitempty" plist:"FilterGrade,omitempty"`
	// A globally unique identifier for this content filter configuration. The content filter processes network traffic for managed apps with the same `ContentFilterUUID` in their app attributes. Use when `FilterType` is `Plugin`.This key must be present for unsupervised devices and user enrollment.
	ContentFilterUUID *string `json:"ContentFilterUUID,omitempty" plist:"ContentFilterUUID,omitempty"`
	// If `true`, the system filters URL requests. Use when `FilterType` is `Plugin`. Available in iOS 26 and macOS 26, and later.
	FilterURLs *bool `json:"FilterURLs,omitempty" plist:"FilterURLs,omitempty"`
	// A dictionary containing URL filter parameters. Required when `FilterURLs` is `true`. Available in iOS 26 and macOS 26 and later.
	URLFilterParameters *URLFilterParameters `json:"URLFilterParameters,omitempty" plist:"URLFilterParameters,omitempty"`
}

func (p *WebContentFilter) PayloadType() string {
	return "com.apple.webcontent-filter"
}

// The type of filter, built-in or plug-in. In macOS, the system only supports the plug-in value.
type FilterType string

const (
	FilterTypeBuiltIn FilterType = "BuiltIn"
	FilterTypePlugin  FilterType = "Plugin"
)

type WhitelistedBookmarksItem struct {
	// The URL of the bookmark in the allow list.
	URL string `json:"URL" plist:"URL" required:"true"`
	// The title of the bookmark.
	Title string `json:"Title" plist:"Title" required:"true"`
}
type AllowListBookmarksItem struct {
	// The URL of the bookmark in the allow list.
	URL string `json:"URL" plist:"URL" required:"true"`
	// The title of the bookmark.
	Title string `json:"Title" plist:"Title" required:"true"`
}

// The system uses this value to derive the relative order of content filters. Filters with a grade of `firewall` see network traffic before filters with a grade of `inspector`. However, the system doesn't define the order of filters within a grade.
type FilterGrade string

const (
	FilterGradeFirewall  FilterGrade = "firewall"
	FilterGradeInspector FilterGrade = "inspector"
)

// A dictionary containing URL filter parameters. Required when `FilterURLs` is `true`. Available in iOS 26 and macOS 26 and later.
type URLFilterParameters struct {
	// The designated requirement string in the code signature of the URL filter control provider app extension. The system uses this string to identify the URL filter control provider when the filter starts running. Required in macOS.
	URLFilterControlProviderDesignatedRequirement *string `json:"URLFilterControlProviderDesignatedRequirement,omitempty" plist:"URLFilterControlProviderDesignatedRequirement,omitempty"`
	// The bundle identifier string of the URL filter control provider app extension. The system uses this string to identify the URL filter control provider when the filter starts running.
	URLFilterControlProviderBundleIdentifier string `json:"URLFilterControlProviderBundleIdentifier" plist:"URLFilterControlProviderBundleIdentifier" required:"true"`
	// The URL containing the domain name of the private information retrieval server.
	PIRServerURL string `json:"PIRServerURL" plist:"PIRServerURL" required:"true"`
	// The URL containing the domain name of Privacy Pass Issuer.
	PIRPrivacyPassIssuerURL string `json:"PIRPrivacyPassIssuerURL" plist:"PIRPrivacyPassIssuerURL" required:"true"`
	// The per-user authentication token string, which is an HTTP bearer token for the person using your app. The system uses this token to attest that it is a valid user when requesting anonymous authentication tokens for PIR exchanges.
	PIRAuthenticationToken string `json:"PIRAuthenticationToken" plist:"PIRAuthenticationToken" required:"true"`
	// If `true`, the system blocks URLs if the filter is enabled, but it fails to make any filtering decision; for example, if there's a communication failure with the PIR server. If `false`, the system allows URLs if the filter is enabled, but it fails to make any filtering decision.
	URLFilterFailClosed *bool `json:"URLFilterFailClosed,omitempty" plist:"URLFilterFailClosed,omitempty"`
	// The time interval in seconds that the system uses to periodically run the `NEURLFilterControlProvider` app extension. The default value is 86400 seconds (1 day). The minimum allowed value is 2700 seconds (45 minutes). The system allows `NEURLFilterControlProvider` implementations to download prefilter Bloom filter data onto the device periodically at the specified interval. Implementations need to allow for a slight difference between the scheduled time and the actual runtime of the task, due to the scheduling mechanism on the system.
	URLPrefilterFetchFrequency *int64 `default:"86400" json:"URLPrefilterFetchFrequency,omitempty" plist:"URLPrefilterFetchFrequency,omitempty"`
}

// The payload that configures Wi-Fi settings.
type WiFi struct {
	*CommonPayloadKeys
	// If `true`, the device joins the network automatically.
	// If `false`, the user must tap the network name to join it.
	AutoJoin *bool `json:"AutoJoin,omitempty" plist:"AutoJoin,omitempty"`
	// The SSID of the Wi-Fi network to use. In iOS 7.0 and later, the SSID is optional if a value exists for `DomainName` value.
	SSIDSTR *string `json:"SSID_STR,omitempty" plist:"SSID_STR,omitempty"`
	// If `true`, defines this network as hidden.
	HIDDENNETWORK *bool `json:"HIDDEN_NETWORK,omitempty" plist:"HIDDEN_NETWORK,omitempty"`
	// The proxy type, if any, to use. If you choose the manual proxy type, you need the proxy server address, including its port and optionally a user name and password into the proxy server. If you choose the auto proxy type, you can enter a proxy autoconfiguration (PAC) URL.
	ProxyType *WiFiProxyType `default:"None" json:"ProxyType,omitempty" plist:"ProxyType,omitempty"`
	// The encryption type for the network.
	// If set to anything except `None`, the payload may contain the following three keys: `Password`, `PayloadCertificateUUID`, or `EAPClientConfiguration`.
	// As of iOS 16, tvOS 16, watchOS 9, and macOS 13:
	// - `WPA` allows joining WPA or WPA2 networks
	// - `WPA2` allows joining WPA2 or WPA3 networks
	// - `WPA3` allows joining WPA3 networks only
	// - `Any` allows joining WPA, WPA2, WPA3, and WEP networks
	// Prior to iOS 16, tvOS 16, and watchOS 9, specifying `WPA`, `WPA2`, and `WPA3` were equivalent and would allow joining any WPA network.
	// Prior to macOS 13, the encryption type, if specified explicitly, needed to match the encryption type of the network exactly.
	EncryptionType *EncryptionType `default:"Any" json:"EncryptionType,omitempty" plist:"EncryptionType,omitempty"`
	// The password for the access point.
	Password *string `json:"Password,omitempty" plist:"Password,omitempty"`
	// The UUID of the certificate payload within the same profile to use for the client credential.
	PayloadCertificateUUID *string `json:"PayloadCertificateUUID,omitempty" plist:"PayloadCertificateUUID,omitempty"`
	// The enterprise network configuration.
	EAPClientConfiguration *EAPClientConfiguration `json:"EAPClientConfiguration,omitempty" plist:"EAPClientConfiguration,omitempty"`
	// The operator name to display when connected to this network. Used only with Wi-Fi Hotspot 2.0 access points.
	DisplayedOperatorName *string `json:"DisplayedOperatorName,omitempty" plist:"DisplayedOperatorName,omitempty"`
	// The primary domain of the tunnel.
	DomainName *string `json:"DomainName,omitempty" plist:"DomainName,omitempty"`
	// An array of Roaming Consortium Organization Identifiers used for Wi-Fi Hotspot 2.0 negotiation.
	RoamingConsortiumOIs *[]string `json:"RoamingConsortiumOIs,omitempty" plist:"RoamingConsortiumOIs,omitempty"`
	// If `true`, allows connection to roaming service providers.
	ServiceProviderRoamingEnabled *bool `json:"ServiceProviderRoamingEnabled,omitempty" plist:"ServiceProviderRoamingEnabled,omitempty"`
	// If `true`, the device treats the network as a hotspot.
	IsHotspot *bool `json:"IsHotspot,omitempty" plist:"IsHotspot,omitempty"`
	// The HESSID used for Wi-Fi Hotspot 2.0 negotiation.
	HESSID *string `json:"HESSID,omitempty" plist:"HESSID,omitempty"`
	// An array of Network Access Identifier Realm names used for Wi-Fi Hotspot 2.0 negotiation.
	NAIRealmNames *[]string `json:"NAIRealmNames,omitempty" plist:"NAIRealmNames,omitempty"`
	// An array of Mobile Country Code/Mobile Network Code (MCC/MNC) pairs used for Wi-Fi Hotspot 2.0 negotiation. Each string must contain exactly six digits.
	MCCAndMNCs *[]string `json:"MCCAndMNCs,omitempty" plist:"MCCAndMNCs,omitempty"`
	// If `true`, the system bypasses Captive Network detection when the device connects to the network.
	CaptiveBypass *bool `json:"CaptiveBypass,omitempty" plist:"CaptiveBypass,omitempty"`
	// A dictionary that contains the list of apps that the system allows to benefit from L2 and L3 marking. When this dictionary isn't present, the system allows all apps to use L2 and L3 marking when the Wi-Fi network supports Cisco QoS fast lane.
	QoSMarkingPolicy *QoSMarkingPolicy `json:"QoSMarkingPolicy,omitempty" plist:"QoSMarkingPolicy,omitempty"`
	// An array of strings that contain the type of connection mode to attach.
	SetupModes *[]SetupModes `json:"SetupModes,omitempty" plist:"SetupModes,omitempty"`
	// If `true`, enables IPv6 on this interface.
	EnableIPv6 *bool `json:"EnableIPv6,omitempty" plist:"EnableIPv6,omitempty"`
	// If `true`, allows for two-factor authentication for EAP-TTLS, PEAP, or EAP-FAST. If `false`, allows for zero-factor authentication for EAP-TLS.
	TLSCertificateRequired *bool `json:"TLSCertificateRequired,omitempty" plist:"TLSCertificateRequired,omitempty"`
	// The proxy server's network address.
	ProxyServer *string `json:"ProxyServer,omitempty" plist:"ProxyServer,omitempty"`
	// The proxy server's port number.
	ProxyServerPort *int64 `json:"ProxyServerPort,omitempty" plist:"ProxyServerPort,omitempty"`
	// The user name used to authenticate to the proxy server.
	ProxyUsername *string `json:"ProxyUsername,omitempty" plist:"ProxyUsername,omitempty"`
	// The password used to authenticate to the proxy server.
	ProxyPassword *string `json:"ProxyPassword,omitempty" plist:"ProxyPassword,omitempty"`
	// The URL of the PAC file that defines the proxy configuration.
	ProxyPACURL *string `json:"ProxyPACURL,omitempty" plist:"ProxyPACURL,omitempty"`
	// If `true`, allows connecting directly to the destination if the PAC file is unreachable.
	ProxyPACFallbackAllowed *bool `json:"ProxyPACFallbackAllowed,omitempty" plist:"ProxyPACFallbackAllowed,omitempty"`
	// If `true,` disables MAC address randomization for a Wi-Fi network while associated with that network. This feature also shows a privacy warning in Settings indicating that the network has reduced privacy protections.
	// If `false`, then the system enables MAC address randomization on iOS, watchOS, and visionOS.
	// This value is only locked when MDM installs the profile. If the profile is manually installed, the system sets the value but the user can change it.
	DisableAssociationMACRandomization *bool `json:"DisableAssociationMACRandomization,omitempty" plist:"DisableAssociationMACRandomization,omitempty"`
	// If `true`, the device makes this network available for joining before the device is unlocked for the first time following a reboot, on a device configured for return to service. Any network credentials are placed into Class D storage within the keychain, and information about the network is stored on disk in Class D.
	// There are several restrictions on the use of this flag:
	// - This property is only available in the return to service mode.
	// - Only one network can be designated as available before first unlock.
	// - `EAPClientConfiguration` must not be present.
	// - If `IsHotspot` is present, it must be set to `false`.
	// - `QoSMarkingPolicy` must not be present.
	// - If `ProxyType` is present, it must be set to `None`.
	// The device fails to install the profile payload if any of these conditions are not met.
	AllowJoinBeforeFirstUnlock *bool `json:"AllowJoinBeforeFirstUnlock,omitempty" plist:"AllowJoinBeforeFirstUnlock,omitempty"`
}

func (p *WiFi) PayloadType() string {
	return "com.apple.wifi.managed"
}

// The proxy type, if any, to use. If you choose the manual proxy type, you need the proxy server address, including its port and optionally a user name and password into the proxy server. If you choose the auto proxy type, you can enter a proxy autoconfiguration (PAC) URL.
type WiFiProxyType string

const (
	WiFiProxyTypeNone   WiFiProxyType = "None"
	WiFiProxyTypeManual WiFiProxyType = "Manual"
	WiFiProxyTypeAuto   WiFiProxyType = "Auto"
)

// The encryption type for the network.
// If set to anything except `None`, the payload may contain the following three keys: `Password`, `PayloadCertificateUUID`, or `EAPClientConfiguration`.
// As of iOS 16, tvOS 16, watchOS 9, and macOS 13:
// - `WPA` allows joining WPA or WPA2 networks
// - `WPA2` allows joining WPA2 or WPA3 networks
// - `WPA3` allows joining WPA3 networks only
// - `Any` allows joining WPA, WPA2, WPA3, and WEP networks
// Prior to iOS 16, tvOS 16, and watchOS 9, specifying `WPA`, `WPA2`, and `WPA3` were equivalent and would allow joining any WPA network.
// Prior to macOS 13, the encryption type, if specified explicitly, needed to match the encryption type of the network exactly.
type EncryptionType string

const (
	EncryptionTypeWEP  EncryptionType = "WEP"
	EncryptionTypeWPA  EncryptionType = "WPA"
	EncryptionTypeWPA2 EncryptionType = "WPA2"
	EncryptionTypeWPA3 EncryptionType = "WPA3"
	EncryptionTypeAny  EncryptionType = "Any"
	EncryptionTypeNone EncryptionType = "None"
)

// The enterprise network configuration.
type EAPClientConfiguration struct {
	// The EAP types that the system accepts. Allowed values:
	// - `13`: EAP-TLS
	// - `17`: LEAP
	// - `18`: EAP-SIM
	// - `21`: EAP-TTLS
	// - `23`: EAP-AKA
	// - `25`: PEAPv0/v1
	// - `43`: EAP-FAST
	// For EAP-TLS authentication without a network payload, install the necessary identity certificates and have your users select EAP-TLS mode in the 802.1X credentials dialog that appears when they connect to the network. For other EAP types, a network payload is necessary and must specify the correct settings for the network.
	AcceptEAPTypes []AcceptEAPTypes `json:"AcceptEAPTypes" plist:"AcceptEAPTypes" required:"true"`
	// The user name for the account. If you don't specify a value, the system prompts the user during login.
	UserName *string `json:"UserName,omitempty" plist:"UserName,omitempty"`
	// The user's password. If you don't specify a value, the system prompts the user during login.
	UserPassword *string `json:"UserPassword,omitempty" plist:"UserPassword,omitempty"`
	// An array of the UUID of each certificate payload in the same profile to trust for authentication. Use this key to prevent the device from asking the user whether to trust the listed certificates. Dynamic trust (the certificate dialogue) is in a disabled state if you specify this property without also enabling 'TLSAllowTrustExceptions'.
	PayloadCertificateAnchorUUID *[]string `json:"PayloadCertificateAnchorUUID,omitempty" plist:"PayloadCertificateAnchorUUID,omitempty"`
	// An array of trusted certificates. Each entry in the array must contain certificate data that represents an anchor certificate used for verifying the server certificate.
	TLSTrustedCertificates *[]string `json:"TLSTrustedCertificates,omitempty" plist:"TLSTrustedCertificates,omitempty"`
	// The list of accepted server certificate common names. If a server presents a certificate that isn't in this list, the system doesn't trust it.
	// If you specify this property, the system disables dynamic trust (the certificate dialog) unless you also specify 'TLSAllowTrustExceptions' with the value 'true'.
	// If necessary, use wildcards to specify the name, such as 'wpa.*.example.com'.
	TLSTrustedServerNames *[]string `json:"TLSTrustedServerNames,omitempty" plist:"TLSTrustedServerNames,omitempty"`
	// If 'true', allows a dynamic trust decision by the user. The dynamic trust is the certificate dialogue that appears when the system doesn't trust a certificate.
	// If 'false', the authentication fails if the system doesn't already trust the certificate.
	// As of iOS 8, Apple no longer supports this key.
	TLSAllowTrustExceptions *bool `json:"TLSAllowTrustExceptions,omitempty" plist:"TLSAllowTrustExceptions,omitempty"`
	// If 'true', allows for two-factor authentication for EAP-TTLS, PEAP, or EAP-FAST. If 'false', allows for zero-factor authentication for EAP-TLS.
	// If you don't specify a value, the default is 'true' for EAP-TLS, and 'false' for other EAP types.
	TLSCertificateIsRequired *bool `json:"TLSCertificateIsRequired,omitempty" plist:"TLSCertificateIsRequired,omitempty"`
	// The inner authentication that the TTLS module uses.
	TTLSInnerAuthentication *TTLSInnerAuthentication `default:"MSCHAPv2" json:"TTLSInnerAuthentication,omitempty" plist:"TTLSInnerAuthentication,omitempty"`
	// The minimum TLS version for EAP authentication.
	TLSMinimumVersion *EAPClientConfigurationTLSMinimumVersion `default:"1.0" json:"TLSMinimumVersion,omitempty" plist:"TLSMinimumVersion,omitempty"`
	// The maximum TLS version for EAP authentication.
	TLSMaximumVersion *EAPClientConfigurationTLSMaximumVersion `default:"1.2" json:"TLSMaximumVersion,omitempty" plist:"TLSMaximumVersion,omitempty"`
	// A name that hides the user's true name. The user's actual name appears only inside the encrypted tunnel. For example, you might set this to anonymous or anon, or anon@mycompany.net. It can increase security because an attacker can't see the authenticating user's name in the clear.
	// This key is only relevant to TTLS, PEAP, and EAP-FAST.
	// This field is required if 'TLSMinimumVersion' is '1.3'.
	OuterIdentity *string `json:"OuterIdentity,omitempty" plist:"OuterIdentity,omitempty"`
	// If 'true', the device uses an existing PAC if it's present. Otherwise, the server must present its identity using a certificate.
	EAPFASTUsePAC *bool `json:"EAPFASTUsePAC,omitempty" plist:"EAPFASTUsePAC,omitempty"`
	// If 'true', allows PAC provisioning.
	// This value is only applicable if 'EAPFASTUsePAC' is 'true'. This value must be 'true' for EAP-FAST PAC usage to succeed because there's no other way to provision a PAC.
	EAPFASTProvisionPAC *bool `json:"EAPFASTProvisionPAC,omitempty" plist:"EAPFASTProvisionPAC,omitempty"`
	// If 'true', provisions the device anonymously. Note that there are known machine-in-the-middle attacks for anonymous provisioning.
	EAPFASTProvisionPACAnonymously *bool `json:"EAPFASTProvisionPACAnonymously,omitempty" plist:"EAPFASTProvisionPACAnonymously,omitempty"`
	// The minimum number of RAND values to accept from the server.
	// For use with EAP-SIM only.
	EAPSIMNumberOfRANDs *EAPSIMNumberOfRANDs `default:"3" json:"EAPSIMNumberOfRANDs,omitempty" plist:"EAPSIMNumberOfRANDs,omitempty"`
	// Set this string to 'ActiveDirectory' to use the AD computer name and password credentials.
	// If using this property, you can't use 'SystemModeUseOpenDirectoryCredentials'.
	SystemModeCredentialsSource *string `json:"SystemModeCredentialsSource,omitempty" plist:"SystemModeCredentialsSource,omitempty"`
	// If 'true', the system mode connection tries to use the Open Directory credentials.
	// If using this property, you can't use 'SystemModeCredentialsSource'.
	SystemModeUseOpenDirectoryCredentials *bool `json:"SystemModeUseOpenDirectoryCredentials,omitempty" plist:"SystemModeUseOpenDirectoryCredentials,omitempty"`
	// If 'true', the user receives a prompt for a password each time they connect to the network.
	OneTimeUserPassword *bool `json:"OneTimeUserPassword,omitempty" plist:"OneTimeUserPassword,omitempty"`
}

// The EAP types that the system accepts. Allowed values:
// - `13`: EAP-TLS
// - `17`: LEAP
// - `18`: EAP-SIM
// - `21`: EAP-TTLS
// - `23`: EAP-AKA
// - `25`: PEAPv0/v1
// - `43`: EAP-FAST
// For EAP-TLS authentication without a network payload, install the necessary identity certificates and have your users select EAP-TLS mode in the 802.1X credentials dialog that appears when they connect to the network. For other EAP types, a network payload is necessary and must specify the correct settings for the network.
type AcceptEAPTypes int64

const (
	AcceptEAPTypes13 AcceptEAPTypes = 13
	AcceptEAPTypes17 AcceptEAPTypes = 17
	AcceptEAPTypes18 AcceptEAPTypes = 18
	AcceptEAPTypes21 AcceptEAPTypes = 21
	AcceptEAPTypes23 AcceptEAPTypes = 23
	AcceptEAPTypes25 AcceptEAPTypes = 25
	AcceptEAPTypes43 AcceptEAPTypes = 43
)

// The inner authentication that the TTLS module uses.
type TTLSInnerAuthentication string

const (
	TTLSInnerAuthenticationPAP      TTLSInnerAuthentication = "PAP"
	TTLSInnerAuthenticationEAP      TTLSInnerAuthentication = "EAP"
	TTLSInnerAuthenticationCHAP     TTLSInnerAuthentication = "CHAP"
	TTLSInnerAuthenticationMSCHAP   TTLSInnerAuthentication = "MSCHAP"
	TTLSInnerAuthenticationMSCHAPv2 TTLSInnerAuthentication = "MSCHAPv2"
)

// The minimum TLS version for EAP authentication.
type EAPClientConfigurationTLSMinimumVersion string

const (
	EAPClientConfigurationTLSMinimumVersion10 EAPClientConfigurationTLSMinimumVersion = "1.0"
	EAPClientConfigurationTLSMinimumVersion11 EAPClientConfigurationTLSMinimumVersion = "1.1"
	EAPClientConfigurationTLSMinimumVersion12 EAPClientConfigurationTLSMinimumVersion = "1.2"
	EAPClientConfigurationTLSMinimumVersion13 EAPClientConfigurationTLSMinimumVersion = "1.3"
)

// The maximum TLS version for EAP authentication.
type EAPClientConfigurationTLSMaximumVersion string

const (
	EAPClientConfigurationTLSMaximumVersion10 EAPClientConfigurationTLSMaximumVersion = "1.0"
	EAPClientConfigurationTLSMaximumVersion11 EAPClientConfigurationTLSMaximumVersion = "1.1"
	EAPClientConfigurationTLSMaximumVersion12 EAPClientConfigurationTLSMaximumVersion = "1.2"
	EAPClientConfigurationTLSMaximumVersion13 EAPClientConfigurationTLSMaximumVersion = "1.3"
)

// The minimum number of RAND values to accept from the server.
// For use with EAP-SIM only.
type EAPSIMNumberOfRANDs int64

const (
	EAPSIMNumberOfRANDs2 EAPSIMNumberOfRANDs = 2
	EAPSIMNumberOfRANDs3 EAPSIMNumberOfRANDs = 3
)

// A dictionary that contains the list of apps that the system allows to benefit from L2 and L3 marking. When this dictionary isn't present, the system allows all apps to use L2 and L3 marking when the Wi-Fi network supports Cisco QoS fast lane.
type QoSMarkingPolicy struct {
	// An array of app bundle identifiers that defines the allow list for L2 and L3 marking for traffic that goes to the Wi-Fi network. If the array isn't present, but the `QoSMarkingPolicy` key is present — even empty — no apps can use L2 and L3 marking.
	QoSMarkingAllowListAppIdentifiers *[]string `json:"QoSMarkingAllowListAppIdentifiers,omitempty" plist:"QoSMarkingAllowListAppIdentifiers,omitempty"`
	// Use `QoSMarkingAllowListAppIdentifiers` instead.
	QoSMarkingWhitelistedAppIdentifiers *[]string `json:"QoSMarkingWhitelistedAppIdentifiers,omitempty" plist:"QoSMarkingWhitelistedAppIdentifiers,omitempty"`
	// If `true`, adds audio and video traffic of built-in audio or video services, such as FaceTime and Wi-Fi Calling, to the allow list for L2 and L3 marking for traffic that goes to the Wi-Fi network.
	QoSMarkingAppleAudioVideoCalls *bool `json:"QoSMarkingAppleAudioVideoCalls,omitempty" plist:"QoSMarkingAppleAudioVideoCalls,omitempty"`
	// If `true`, disables L3 marking and only uses L2 marking for traffic that goes to the Wi-Fi network.
	// If `false`, the system behaves as if Wi-Fi doesn't have an association with a Cisco QoS fast lane network.
	QoSMarkingEnabled *bool `json:"QoSMarkingEnabled,omitempty" plist:"QoSMarkingEnabled,omitempty"`
}

// An array of strings that contain the type of connection mode to attach.
// A type of connection mode.
type SetupModes string

const (
	SetupModesSystem      SetupModes = "System"
	SetupModesLoginwindow SetupModes = "Loginwindow"
)

// The payload that configures the Xsan preferences that define the volumes that automatically mount at startup.
// The Xsan preferences payload can be used to configure which volumes automatically mount at startup. For StorNext volumes this payload also determines whether the mount uses Fibre Channel or Distributed LAN Client (DLC).
type XsanPreferences struct {
	*CommonPayloadKeys
	// An array of Xsan or StorNext volume names. The Xsan client attempts to automatically mount these volumes at startup. The system administrator can mount additional volumes manually by using the `xsanctl(8)` mount command.
	OnlyMount *[]string `json:"onlyMount,omitempty" plist:"onlyMount,omitempty"`
	// An array of Xsan or StorNext volume names. If no `onlyMount` array is present, the Xsan client automatically attempts to mount all SAN volumes except the volumes in this array. The system administrator can mount those volumes manually by using the `xsanctl(8)` mount command.
	DenyMount *[]string `json:"denyMount,omitempty" plist:"denyMount,omitempty"`
	// An array of StorNext volume names. If the Xsan client is attempting to mount a volume named in this array, the client only mounts the volume if its logical units (LUNs) are available through Fibre Channel. It doesn't attempt to mount the volume using Distributed LAN Client (DLC).
	DenyDLC *[]string `json:"denyDLC,omitempty" plist:"denyDLC,omitempty"`
	// An array of StorNext volume names. If the Xsan client is attempting to mount a volume named in this array, the Xsan client attempts to mount the volume using DLC. If DLC isn't available, the client attempts to mount the volume if its LUNs are available through Fibre Channel. The volume name must not also appear in `denyDLC`.
	PreferDLC *[]string `json:"preferDLC,omitempty" plist:"preferDLC,omitempty"`
	// If `true`, use the DLC for all volumes.
	UseDLC *bool `json:"useDLC,omitempty" plist:"useDLC,omitempty"`
}

func (p *XsanPreferences) PayloadType() string {
	return "com.apple.xsan.preferences"
}

// The payload that configures an Xsan client system.
// Sets up Xsan clients and controls certain Xsan volume mount behaviors. The payload should include either sanConfigURLs or fsnameservers, but not both.
type Xsan struct {
	*CommonPayloadKeys
	// The name of the SAN. This key is required for all Xsan SANs. The name must exactly match the name of the SAN defined in the metadata server.
	SanName string `json:"sanName" plist:"sanName" required:"true"`
	// An array of LDAP URLs where Xsan systems can obtain SAN configuration updates. There should be one entry for each Xsan MDC.
	// This key is required for all Xsan SANs.
	// Example URL: `ldaps://mdc1.example.com:389`.
	SanConfigURLs *[]string `json:"sanConfigURLs,omitempty" plist:"sanConfigURLs,omitempty"`
	// An array of storage area network (SAN) File System Name Server coordinators. The list should contain the same addresses in the same order as the metadata controller (MDC) `/Library/Preferences/Xsan/fsnameservers` file.
	// This key is required for StorNext SANs.
	Fsnameservers *[]string `json:"fsnameservers,omitempty" plist:"fsnameservers,omitempty"`
	// The authentication method for the SAN. This key is required for all Xsan SANs. It's optional for StorNext SANs but should be set if the StorNext SAN uses an `auth_secret` file.
	// Only one value is accepted: `auth_secret`
	SanAuthMethod *SanAuthMethod `json:"sanAuthMethod,omitempty" plist:"sanAuthMethod,omitempty"`
	// The shared secret used for Xsan network authentication. This key is required when the `sanAuthMethod` key is present. The value should equal the content of the MDC's `/Library/Preferences/Xsan/.auth_secret` file.
	SharedSecret string `json:"sharedSecret" plist:"sharedSecret" required:"true"`
}

func (p *Xsan) PayloadType() string {
	return "com.apple.xsan"
}

// The authentication method for the SAN. This key is required for all Xsan SANs. It's optional for StorNext SANs but should be set if the StorNext SAN uses an `auth_secret` file.
// Only one value is accepted: `auth_secret`
type SanAuthMethod string

const (
	SanAuthMethodAuthSecret SanAuthMethod = "auth_secret"
)

// The payload that configures login behavior.
// This payload handles login items management.
type LoginWindowLoginItems struct {
	*CommonPayloadKeys
	// If `true`, the system prevents the user from disabling login item launches by using the Shift key.
	DisableLoginItemsSuppression *bool `json:"DisableLoginItemsSuppression,omitempty" plist:"DisableLoginItemsSuppression,omitempty"`
}

func (p *LoginWindowLoginItems) PayloadType() string {
	return "loginwindow"
}
