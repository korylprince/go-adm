// DO NOT EDIT
// generated from https://github.com/apple/device-management.git:0a4527c5ea21825fd23e08273ccdb9e2302458ce/declarative/declarations/assets

package assets

const DeviceManagementGenerateHash = "0a4527c5ea21825fd23e08273ccdb9e2302458ce"

var DeclarationMap = map[string]any{
	"com.apple.asset.credential.acme":         AssetCredentialACME{},
	"com.apple.asset.credential.certificate":  AssetCredentialCertificate{},
	"com.apple.asset.credential.identity":     AssetCredentialIdentity{},
	"com.apple.asset.credential.scep":         AssetCredentialSCEP{},
	"com.apple.asset.credential.userpassword": AssetCredentialUserNameandPassword{},
	"com.apple.asset.data":                    AssetData{},
	"com.apple.asset.useridentity":            AssetUserIdentity{},
}

// A reference to an ACME identity.
type AssetCredentialACME struct {
	// The external reference. Ensure that the asset data:
	// - Is a JSON document that represents the `com.apple.credential.acme` credential type
	// - Uses a media type of `application/json`, and if it includes a `ContentType` sub-key, that sub-key media type is also `application/json`
	Reference AssetCredentialACMEReference `json:"Reference" plist:"Reference" required:"true"`
	// The server authentication details.
	Authentication *AssetCredentialACMEAuthentication `json:"Authentication,omitempty" plist:"Authentication,omitempty"`
	// The keychain accessibility that determines when the keychain item is available for use, which has these allowed values:
	// - `Default`: The most restrictive accessibility that still satisfies all uses of the asset by configurations that reference it.
	// - `AfterFirstUnlock`: The keychain item is only available after the first unlock of the device.
	Accessible *AssetCredentialACMEAccessible `default:"Default" json:"Accessible,omitempty" plist:"Accessible,omitempty"`
}

func (p *AssetCredentialACME) DeclarationType() string {
	return "com.apple.asset.credential.acme"
}

// The external reference. Ensure that the asset data:
// - Is a JSON document that represents the `com.apple.credential.acme` credential type
// - Uses a media type of `application/json`, and if it includes a `ContentType` sub-key, that sub-key media type is also `application/json`
type AssetCredentialACMEReference struct {
	// The URL to retrieve data, which needs to start with `https://`.
	DataURL string `json:"DataURL" plist:"DataURL" required:"true"`
	// The media type that describes the data. If present, the system checks the actual media type of the downloaded data, and an error occurs if the values don't match.
	ContentType *string `json:"ContentType,omitempty" plist:"ContentType,omitempty"`
	// The size of the data. Set the size to `0` if there's no expectation of a response body. If present, the system checks the actual size of the downloaded data, and an error occurs if the values don't match.
	Size *int64 `json:"Size,omitempty" plist:"Size,omitempty"`
	// A SHA-256 hash of the data stored at the `DataURL`. Don't set this value if `Size` is `0` as the client ignores it. However, if present, the system checks the actual hash of the downloaded data, and an error occurs if the values don't match.
	HashSHA256 *string `json:"Hash-SHA-256,omitempty" plist:"Hash-SHA-256,omitempty"`
}

// The server authentication details.
type AssetCredentialACMEAuthentication struct {
	// The type of authentication, which has these allowed values:
	// - `MDM`: A request that uses MDM semantics, which includes the device-identity certificate, and any user authentication. This is equivalent to an MDM request made to the `CheckInURL` or `ServerURL`. This option is only available through declarative device management.
	// - `None`: A standard GET request.
	Type AssetCredentialACMEAuthenticationType `json:"Type" plist:"Type" required:"true"`
}

// The type of authentication, which has these allowed values:
// - `MDM`: A request that uses MDM semantics, which includes the device-identity certificate, and any user authentication. This is equivalent to an MDM request made to the `CheckInURL` or `ServerURL`. This option is only available through declarative device management.
// - `None`: A standard GET request.
type AssetCredentialACMEAuthenticationType string

const (
	AssetCredentialACMEAuthenticationTypeMDM  AssetCredentialACMEAuthenticationType = "MDM"
	AssetCredentialACMEAuthenticationTypeNone AssetCredentialACMEAuthenticationType = "None"
)

// The keychain accessibility that determines when the keychain item is available for use, which has these allowed values:
// - `Default`: The most restrictive accessibility that still satisfies all uses of the asset by configurations that reference it.
// - `AfterFirstUnlock`: The keychain item is only available after the first unlock of the device.
type AssetCredentialACMEAccessible string

const (
	AssetCredentialACMEAccessibleDefault          AssetCredentialACMEAccessible = "Default"
	AssetCredentialACMEAccessibleAfterFirstUnlock AssetCredentialACMEAccessible = "AfterFirstUnlock"
)

// A reference to a PKCS #1 or PEM encoded certificate.
type AssetCredentialCertificate struct {
	// The external reference. Ensure that the asset data uses a media type of `application/pkcs1` or `application/pem` to correctly identify the type of encoded certificate. If the asset data includes a `ContentType` sub-key, set it to the corresponding media type.
	Reference AssetCredentialCertificateReference `json:"Reference" plist:"Reference" required:"true"`
	// The server authentication details.
	Authentication *AssetCredentialCertificateAuthentication `json:"Authentication,omitempty" plist:"Authentication,omitempty"`
}

func (p *AssetCredentialCertificate) DeclarationType() string {
	return "com.apple.asset.credential.certificate"
}

// The external reference. Ensure that the asset data uses a media type of `application/pkcs1` or `application/pem` to correctly identify the type of encoded certificate. If the asset data includes a `ContentType` sub-key, set it to the corresponding media type.
type AssetCredentialCertificateReference struct {
	// The URL to retrieve data, which needs to start with `https://`.
	DataURL string `json:"DataURL" plist:"DataURL" required:"true"`
	// The media type that describes the data. If present, the system checks the actual media type of the downloaded data, and an error occurs if the values don't match.
	ContentType *string `json:"ContentType,omitempty" plist:"ContentType,omitempty"`
	// The size of the data. Set the size to `0` if there's no expectation of a response body. If present, the system checks the actual size of the downloaded data, and an error occurs if the values don't match.
	Size *int64 `json:"Size,omitempty" plist:"Size,omitempty"`
	// A SHA-256 hash of the data stored at the `DataURL`. Don't set this value if `Size` is `0` as the client ignores it. However, if present, the system checks the actual hash of the downloaded data, and an error occurs if the values don't match.
	HashSHA256 *string `json:"Hash-SHA-256,omitempty" plist:"Hash-SHA-256,omitempty"`
}

// The server authentication details.
type AssetCredentialCertificateAuthentication struct {
	// The type of authentication, which has these allowed values:
	// - `MDM`: A request that uses MDM semantics, which includes the device-identity certificate, and any user authentication. This is equivalent to an MDM request made to the `CheckInURL` or `ServerURL`. This option is only available through declarative device management.
	// - `None`: A standard GET request.
	Type AssetCredentialCertificateAuthenticationType `json:"Type" plist:"Type" required:"true"`
}

// The type of authentication, which has these allowed values:
// - `MDM`: A request that uses MDM semantics, which includes the device-identity certificate, and any user authentication. This is equivalent to an MDM request made to the `CheckInURL` or `ServerURL`. This option is only available through declarative device management.
// - `None`: A standard GET request.
type AssetCredentialCertificateAuthenticationType string

const (
	AssetCredentialCertificateAuthenticationTypeMDM  AssetCredentialCertificateAuthenticationType = "MDM"
	AssetCredentialCertificateAuthenticationTypeNone AssetCredentialCertificateAuthenticationType = "None"
)

// A reference to a PKCS #12 password-protected identity.
type AssetCredentialIdentity struct {
	// The external reference. Ensure that the asset data:
	// - Is a JSON document that represents the `com.apple.credential.identity` credential type
	// - Uses a media type of `application/json`, and if it includes a `ContentType` sub-key, that sub-key media type is also `application/json`
	Reference AssetCredentialIdentityReference `json:"Reference" plist:"Reference" required:"true"`
	// The server authentication details.
	Authentication *AssetCredentialIdentityAuthentication `json:"Authentication,omitempty" plist:"Authentication,omitempty"`
	// The keychain accessibility that determines when the keychain item is available for use, which has these allowed values:
	// - `Default`: The most restrictive accessibility that still satisfies all uses of the asset by configurations that reference it.
	// - `AfterFirstUnlock`: The keychain item is only available after the first unlock of the device.
	Accessible *AssetCredentialIdentityAccessible `default:"Default" json:"Accessible,omitempty" plist:"Accessible,omitempty"`
}

func (p *AssetCredentialIdentity) DeclarationType() string {
	return "com.apple.asset.credential.identity"
}

// The external reference. Ensure that the asset data:
// - Is a JSON document that represents the `com.apple.credential.identity` credential type
// - Uses a media type of `application/json`, and if it includes a `ContentType` sub-key, that sub-key media type is also `application/json`
type AssetCredentialIdentityReference struct {
	// The URL to retrieve data, which needs to start with `https://`.
	DataURL string `json:"DataURL" plist:"DataURL" required:"true"`
	// The media type that describes the data. If present, the system checks the actual media type of the downloaded data, and an error occurs if the values don't match.
	ContentType *string `json:"ContentType,omitempty" plist:"ContentType,omitempty"`
	// The size of the data. Set the size to `0` if there's no expectation of a response body. If present, the system checks the actual size of the downloaded data, and an error occurs if the values don't match.
	Size *int64 `json:"Size,omitempty" plist:"Size,omitempty"`
	// A SHA-256 hash of the data stored at the `DataURL`. Don't set this value if `Size` is `0` as the client ignores it. However, if present, the system checks the actual hash of the downloaded data, and an error occurs if the values don't match.
	HashSHA256 *string `json:"Hash-SHA-256,omitempty" plist:"Hash-SHA-256,omitempty"`
}

// The server authentication details.
type AssetCredentialIdentityAuthentication struct {
	// The type of authentication, which has these allowed values:
	// - `MDM`: A request that uses MDM semantics, which includes the device-identity certificate, and any user authentication. This is equivalent to an MDM request made to the `CheckInURL` or `ServerURL`. This option is only available through declarative device management.
	// - `None`: A standard GET request.
	Type AssetCredentialIdentityAuthenticationType `json:"Type" plist:"Type" required:"true"`
}

// The type of authentication, which has these allowed values:
// - `MDM`: A request that uses MDM semantics, which includes the device-identity certificate, and any user authentication. This is equivalent to an MDM request made to the `CheckInURL` or `ServerURL`. This option is only available through declarative device management.
// - `None`: A standard GET request.
type AssetCredentialIdentityAuthenticationType string

const (
	AssetCredentialIdentityAuthenticationTypeMDM  AssetCredentialIdentityAuthenticationType = "MDM"
	AssetCredentialIdentityAuthenticationTypeNone AssetCredentialIdentityAuthenticationType = "None"
)

// The keychain accessibility that determines when the keychain item is available for use, which has these allowed values:
// - `Default`: The most restrictive accessibility that still satisfies all uses of the asset by configurations that reference it.
// - `AfterFirstUnlock`: The keychain item is only available after the first unlock of the device.
type AssetCredentialIdentityAccessible string

const (
	AssetCredentialIdentityAccessibleDefault          AssetCredentialIdentityAccessible = "Default"
	AssetCredentialIdentityAccessibleAfterFirstUnlock AssetCredentialIdentityAccessible = "AfterFirstUnlock"
)

// A reference to a SCEP identity.
type AssetCredentialSCEP struct {
	// The external reference. Ensure that the asset data:
	// - Is a JSON document that represents the `com.apple.credential.scep` credential type
	// - Uses a media type of `application/json`, and if it includes a `ContentType` sub-key, that sub-key media type is also `application/json`
	Reference AssetCredentialSCEPReference `json:"Reference" plist:"Reference" required:"true"`
	// The server authentication details.
	Authentication *AssetCredentialSCEPAuthentication `json:"Authentication,omitempty" plist:"Authentication,omitempty"`
	// The keychain accessibility that determines when the keychain item is available for use, which has these allowed values:
	// - `Default`: The most restrictive accessibility that still satisfies all uses of the asset by configurations that reference it.
	// - `AfterFirstUnlock`: The keychain item is only available after the first unlock of the device.
	Accessible *AssetCredentialSCEPAccessible `default:"Default" json:"Accessible,omitempty" plist:"Accessible,omitempty"`
}

func (p *AssetCredentialSCEP) DeclarationType() string {
	return "com.apple.asset.credential.scep"
}

// The external reference. Ensure that the asset data:
// - Is a JSON document that represents the `com.apple.credential.scep` credential type
// - Uses a media type of `application/json`, and if it includes a `ContentType` sub-key, that sub-key media type is also `application/json`
type AssetCredentialSCEPReference struct {
	// The URL to retrieve data, which needs to start with `https://`.
	DataURL string `json:"DataURL" plist:"DataURL" required:"true"`
	// The media type that describes the data. If present, the system checks the actual media type of the downloaded data, and an error occurs if the values don't match.
	ContentType *string `json:"ContentType,omitempty" plist:"ContentType,omitempty"`
	// The size of the data. Set the size to `0` if there's no expectation of a response body. If present, the system checks the actual size of the downloaded data, and an error occurs if the values don't match.
	Size *int64 `json:"Size,omitempty" plist:"Size,omitempty"`
	// A SHA-256 hash of the data stored at the `DataURL`. Don't set this value if `Size` is `0` as the client ignores it. However, if present, the system checks the actual hash of the downloaded data, and an error occurs if the values don't match.
	HashSHA256 *string `json:"Hash-SHA-256,omitempty" plist:"Hash-SHA-256,omitempty"`
}

// The server authentication details.
type AssetCredentialSCEPAuthentication struct {
	// The type of authentication, which has these allowed values:
	// - `MDM`: A request that uses MDM semantics, which includes the device-identity certificate, and any user authentication. This is equivalent to an MDM request made to the `CheckInURL` or `ServerURL`. This option is only available through declarative device management.
	// - `None`: A standard GET request.
	Type AssetCredentialSCEPAuthenticationType `json:"Type" plist:"Type" required:"true"`
}

// The type of authentication, which has these allowed values:
// - `MDM`: A request that uses MDM semantics, which includes the device-identity certificate, and any user authentication. This is equivalent to an MDM request made to the `CheckInURL` or `ServerURL`. This option is only available through declarative device management.
// - `None`: A standard GET request.
type AssetCredentialSCEPAuthenticationType string

const (
	AssetCredentialSCEPAuthenticationTypeMDM  AssetCredentialSCEPAuthenticationType = "MDM"
	AssetCredentialSCEPAuthenticationTypeNone AssetCredentialSCEPAuthenticationType = "None"
)

// The keychain accessibility that determines when the keychain item is available for use, which has these allowed values:
// - `Default`: The most restrictive accessibility that still satisfies all uses of the asset by configurations that reference it.
// - `AfterFirstUnlock`: The keychain item is only available after the first unlock of the device.
type AssetCredentialSCEPAccessible string

const (
	AssetCredentialSCEPAccessibleDefault          AssetCredentialSCEPAccessible = "Default"
	AssetCredentialSCEPAccessibleAfterFirstUnlock AssetCredentialSCEPAccessible = "AfterFirstUnlock"
)

// A reference to data that describes a credential that represents a user name and password.
type AssetCredentialUserNameandPassword struct {
	// The external reference. Ensure that the asset data:
	// - Is a JSON document that represents the `com.apple.credential.usernameandpassword` credential type
	// - Uses a media type of `application/json`, and if it includes a `ContentType` sub-key, that sub-key media type is also `application/json`
	Reference AssetCredentialUserNameandPasswordReference `json:"Reference" plist:"Reference" required:"true"`
	// The server authentication details.
	Authentication *AssetCredentialUserNameandPasswordAuthentication `json:"Authentication,omitempty" plist:"Authentication,omitempty"`
}

func (p *AssetCredentialUserNameandPassword) DeclarationType() string {
	return "com.apple.asset.credential.userpassword"
}

// The external reference. Ensure that the asset data:
// - Is a JSON document that represents the `com.apple.credential.usernameandpassword` credential type
// - Uses a media type of `application/json`, and if it includes a `ContentType` sub-key, that sub-key media type is also `application/json`
type AssetCredentialUserNameandPasswordReference struct {
	// The URL to retrieve data, which needs to start with `https://`.
	DataURL string `json:"DataURL" plist:"DataURL" required:"true"`
	// The media type that describes the data. If present, the system checks the actual media type of the downloaded data, and an error occurs if the values don't match.
	ContentType *string `json:"ContentType,omitempty" plist:"ContentType,omitempty"`
	// The size of the data. Set the size to `0` if there's no expectation of a response body. If present, the system checks the actual size of the downloaded data, and an error occurs if the values don't match.
	Size *int64 `json:"Size,omitempty" plist:"Size,omitempty"`
	// A SHA-256 hash of the data stored at the `DataURL`. Don't set this value if `Size` is `0` as the client ignores it. However, if present, the system checks the actual hash of the downloaded data, and an error occurs if the values don't match.
	HashSHA256 *string `json:"Hash-SHA-256,omitempty" plist:"Hash-SHA-256,omitempty"`
}

// The server authentication details.
type AssetCredentialUserNameandPasswordAuthentication struct {
	// The type of authentication, which has these allowed values:
	// - `MDM`: A request that uses MDM semantics, which includes the device-identity certificate, and any user authentication. This is equivalent to an MDM request made to the `CheckInURL` or `ServerURL`. This option is only available through declarative device management.
	// - `None`: A standard GET request.
	Type AssetCredentialUserNameandPasswordAuthenticationType `json:"Type" plist:"Type" required:"true"`
}

// The type of authentication, which has these allowed values:
// - `MDM`: A request that uses MDM semantics, which includes the device-identity certificate, and any user authentication. This is equivalent to an MDM request made to the `CheckInURL` or `ServerURL`. This option is only available through declarative device management.
// - `None`: A standard GET request.
type AssetCredentialUserNameandPasswordAuthenticationType string

const (
	AssetCredentialUserNameandPasswordAuthenticationTypeMDM  AssetCredentialUserNameandPasswordAuthenticationType = "MDM"
	AssetCredentialUserNameandPasswordAuthenticationTypeNone AssetCredentialUserNameandPasswordAuthenticationType = "None"
)

// A reference to arbitrary data with a specific media type.
type AssetData struct {
	// The external reference.
	Reference AssetDataReference `json:"Reference" plist:"Reference" required:"true"`
	// The server authentication details.
	Authentication *AssetDataAuthentication `json:"Authentication,omitempty" plist:"Authentication,omitempty"`
}

func (p *AssetData) DeclarationType() string {
	return "com.apple.asset.data"
}

// The external reference.
type AssetDataReference struct {
	// The URL to retrieve data, which needs to start with `https://`.
	DataURL string `json:"DataURL" plist:"DataURL" required:"true"`
	// The media type that describes the data. If present, the system checks the actual media type of the downloaded data, and an error occurs if the values don't match.
	ContentType *string `json:"ContentType,omitempty" plist:"ContentType,omitempty"`
	// The size of the data. Set the size to `0` if there's no expectation of a response body. If present, the system checks the actual size of the downloaded data, and an error occurs if the values don't match.
	Size *int64 `json:"Size,omitempty" plist:"Size,omitempty"`
	// A SHA-256 hash of the data stored at the `DataURL`. Don't set this value if `Size` is `0` as the client ignores it. However, if present, the system checks the actual hash of the downloaded data, and an error occurs if the values don't match.
	HashSHA256 *string `json:"Hash-SHA-256,omitempty" plist:"Hash-SHA-256,omitempty"`
}

// The server authentication details.
type AssetDataAuthentication struct {
	// The type of authentication, which has these allowed values:
	// - `MDM`: A request that uses MDM semantics, which includes the device-identity certificate, and any user authentication. This is equivalent to an MDM request made to the `CheckInURL` or `ServerURL`. This option is only available through declarative device management.
	// - `None`: A standard GET request.
	Type AssetDataAuthenticationType `json:"Type" plist:"Type" required:"true"`
}

// The type of authentication, which has these allowed values:
// - `MDM`: A request that uses MDM semantics, which includes the device-identity certificate, and any user authentication. This is equivalent to an MDM request made to the `CheckInURL` or `ServerURL`. This option is only available through declarative device management.
// - `None`: A standard GET request.
type AssetDataAuthenticationType string

const (
	AssetDataAuthenticationTypeMDM  AssetDataAuthenticationType = "MDM"
	AssetDataAuthenticationTypeNone AssetDataAuthenticationType = "None"
)

// The user-identity data.
type AssetUserIdentity struct {
	// The user's full name.
	FullName *string `json:"FullName,omitempty" plist:"FullName,omitempty"`
	// The email address of the user.
	EmailAddress *string `json:"EmailAddress,omitempty" plist:"EmailAddress,omitempty"`
}

func (p *AssetUserIdentity) DeclarationType() string {
	return "com.apple.asset.useridentity"
}
