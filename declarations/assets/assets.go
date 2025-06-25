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

// An ACME identity that the device generates.
type ACMECredential struct {
	// Specifies the directory URL of the ACME server. Use the `https` scheme for the URL.
	DirectoryURL string `json:"DirectoryURL" plist:"DirectoryURL" required:"true"`
	// The server can use this as a one-time code to prevent issuing multiple certificates. It also indicates to the ACME server that the device has access to a valid client identifier that the enterprise infrastructure issued. This can help the ACME server determine whether to trust the device, however this is a relatively weak indication because of the risk that an attacker may intercept and duplicate the client identifier.
	ClientIdentifier string `json:"ClientIdentifier" plist:"ClientIdentifier" required:"true"`
	// The valid values for `KeySize` depend on the values of `KeyType` and `HardwareBound`. See those keys for specific requirements.
	KeySize int64 `json:"KeySize" plist:"KeySize" required:"true"`
	// Specifies the type of key pair to generate.
	// `RSA` specifies an RSA key pair. If you set this value to `RSA`, set `KeySize` in the range `[1024..4096]` inclusive and a multiple of `8`, and set `HardwareBound` to `false`.
	// `ECSECPrimeRandom` specifies a key pair on the P-256, P-384 or P-521 curves as defined in FIPS Pub 186-4, and `KeySize` determines the specific curve. If you set this value to `ECSECPrimeRandom`, set `KeySize` to `256`, `384`, or `521`. The system only supports `256` and `384` for hardware bound keys.
	// > Note:
	// > The key size is `521`, not `512`, even though the other key sizes are multiples of `64`.
	KeyType KeyType `json:"KeyType" plist:"KeyType" required:"true"`
	// If `false`, the private key isn't bound to the device.
	// If `true`, the private key is bound to the device. The Secure Enclave generates the key pair, and the private key is cryptographically entangled with a system key. This protects the private key from being exported.
	// If `true`, `KeyType` needs to be `ECSECPrimeRandom` and `KeySize` needs to be `256` or `384`.
	// On macOS, this is a required key. Set the value to `false`.
	HardwareBound bool `json:"HardwareBound" plist:"HardwareBound" required:"true"`
	// The device requests this subject for the certificate that the ACME server issues. The ACME server may override or ignore this field in the certificate it issues.
	// The representation of an X.500 name is an array of OID and value. For example, `/C=US/O=Apple Inc./CN=foo/1.2.5.3=bar` corresponds to:
	// `[ [ ["C", "US"] ], [ ["O", "Apple Inc."] ], [ [ "CN", "foo"] ], [ [ "1.2.5.3", "bar" ] ] ]`
	// You can represent OIDs as dotted numbers or use shortcuts for country (`C`), locality (`L`), state (`ST`), organization (`O`), organizational unit (`OU`), and common name (`CN`).
	Subject [][][]string `json:"Subject" plist:"Subject" required:"true"`
	// Specifies the subject's alternative name that the device requests for the certificate that the ACME server issues. The ACME server may override or ignore this field in the certificate it issues.
	SubjectAltName *ACMECredentialSubjectAltName `json:"SubjectAltName,omitempty" plist:"SubjectAltName,omitempty"`
	// The device requests this key usage for the certificate that the ACME server issues. The ACME server may override or ignore this field in the certificate it issues.
	// The value is a bit field. Bit `0x01` indicates digital signature, and bit `0x04` indicates key encipherment.
	UsageFlags *int64 `json:"UsageFlags,omitempty" plist:"UsageFlags,omitempty"`
	// The device requests this extended key usage for the certificate that the ACME server issues. The ACME server may override or ignore this field in the certificate it issues.
	// The value is an array of strings. Each string is an OID in dotted notation. For example, `["1.3.6.1.5.5.7.3.2", "1.3.6.1.5.5.7.3.4"]` indicates client authentication and email protection.
	ExtendedKeyUsage *[]string `json:"ExtendedKeyUsage,omitempty" plist:"ExtendedKeyUsage,omitempty"`
	// If `true`, the device provides attestations that describe the device and the generated key to the ACME server. The server can use the attestations as strong evidence that the key is bound to the device, and that the device has properties listed in the attestation. The server can use that as part of a trust score to decide whether to issue the requested certificate. When `Attest` is `true`, set `HardwareBound` to `true`. On macOS, set this key, if present, to `false`. The hardware requirements for attestation are described below.
	Attest *bool `json:"Attest,omitempty" plist:"Attest,omitempty"`
}

// Specifies the type of key pair to generate.
// `RSA` specifies an RSA key pair. If you set this value to `RSA`, set `KeySize` in the range `[1024..4096]` inclusive and a multiple of `8`, and set `HardwareBound` to `false`.
// `ECSECPrimeRandom` specifies a key pair on the P-256, P-384 or P-521 curves as defined in FIPS Pub 186-4, and `KeySize` determines the specific curve. If you set this value to `ECSECPrimeRandom`, set `KeySize` to `256`, `384`, or `521`. The system only supports `256` and `384` for hardware bound keys.
// > Note:
// > The key size is `521`, not `512`, even though the other key sizes are multiples of `64`.
type KeyType string

const (
	KeyTypeRSA              KeyType = "RSA"
	KeyTypeECSECPrimeRandom KeyType = "ECSECPrimeRandom"
)

// Specifies the subject's alternative name that the device requests for the certificate that the ACME server issues. The ACME server may override or ignore this field in the certificate it issues.
type ACMECredentialSubjectAltName struct {
	// The RFC 822 email address.
	Rfc822Name *string `json:"rfc822Name,omitempty" plist:"rfc822Name,omitempty"`
	// The DNS name.
	DNSName *string `json:"dNSName,omitempty" plist:"dNSName,omitempty"`
	// The uniform resource identifier.
	UniformResourceIdentifier *string `json:"uniformResourceIdentifier,omitempty" plist:"uniformResourceIdentifier,omitempty"`
	// The NT principal name.
	NtPrincipalName *string `json:"ntPrincipalName,omitempty" plist:"ntPrincipalName,omitempty"`
}

// The data for a PKCS #12 password-protected identity.
type IdentityCredential struct {
	// The password required to decrypt the PKCS #12 identity data.
	Password string `json:"Password" plist:"Password" required:"true"`
	// The PKCS #12 identity data.
	Identity []byte `json:"Identity" plist:"Identity" required:"true"`
}

// A SCEP identity that the device generates.
type SCEPCredential struct {
	// The SCEP URL.
	URL string `json:"URL" plist:"URL" required:"true"`
	// Any string that the SCEP server recognizes. For example, it could be a domain name such as `example.org`. If a certificate authority has multiple CA certificates, you can use this field to specify the required certificate.
	Name *string `json:"Name,omitempty" plist:"Name,omitempty"`
	// The representation of an X.500 name is an array of OID and value. For example, `/C=US/O=Apple Inc./CN=foo/1.2.5.3=bar` corresponds to:
	// `[ [ ["C", "US"] ], [ ["O", "Apple Inc."] ], [ [ "CN", "foo"] ], [ [ "1.2.5.3", "bar" ] ] ]`
	// You can represent OIDs as dotted numbers or use shortcuts for country (`C`), locality (`L`), state (`ST`), organization (`O`), organizational unit (`OU`), and common name (`CN`).
	Subject *[][][]string `json:"Subject,omitempty" plist:"Subject,omitempty"`
	// A preshared secret.
	Challenge *string `json:"Challenge,omitempty" plist:"Challenge,omitempty"`
	// The key size in bits, either `1024`, `2048`, or `4096`.
	Keysize *Keysize `default:"1024" json:"Keysize,omitempty" plist:"Keysize,omitempty"`
	// The key type, which always has the value `RSA`.
	KeyType *string `default:"RSA" json:"Key Type,omitempty" plist:"Key Type,omitempty"`
	// A bitmask that specifies the use of the key: `1` is signing, `4` is encryption, and `5` is both signing and encryption. Some certificate authorities, such as Windows CA, support only encryption or signing, but not both at the same time.
	KeyUsage *int64 `default:"0" json:"Key Usage,omitempty" plist:"Key Usage,omitempty"`
	// The fingerprint of the Certificate Authority certificate.
	CAFingerprint *[]byte `json:"CAFingerprint,omitempty" plist:"CAFingerprint,omitempty"`
	// The number of times the device should retry if the server sends a `PENDING` response.
	Retries *int64 `default:"3" json:"Retries,omitempty" plist:"Retries,omitempty"`
	// The number of seconds to wait between subsequent retries. The system makes the first retry without this delay.
	RetryDelay *int64 `default:"10" json:"RetryDelay,omitempty" plist:"RetryDelay,omitempty"`
	// The subject's alternative name for the certificate.
	SubjectAltName *SCEPCredentialSubjectAltName `json:"SubjectAltName,omitempty" plist:"SubjectAltName,omitempty"`
}

// The key size in bits, either `1024`, `2048`, or `4096`.
type Keysize int64

const (
	Keysize1024 Keysize = 1024
	Keysize2048 Keysize = 2048
	Keysize4096 Keysize = 4096
)

// The subject's alternative name for the certificate.
type SCEPCredentialSubjectAltName struct {
	// The RFC 822 email address.
	Rfc822Name *string `json:"rfc822Name,omitempty" plist:"rfc822Name,omitempty"`
	// The DNS name.
	DNSName *string `json:"dNSName,omitempty" plist:"dNSName,omitempty"`
	// The uniform resource identifier.
	UniformResourceIdentifier *string `json:"uniformResourceIdentifier,omitempty" plist:"uniformResourceIdentifier,omitempty"`
	// The NT principal name.
	NtPrincipalName *string `json:"ntPrincipalName,omitempty" plist:"ntPrincipalName,omitempty"`
}

// Data that describes a credential that represents a user name and password.
type UserNameandPasswordCredential struct {
	// The user name for this credential.
	UserName string `json:"UserName" plist:"UserName" required:"true"`
	// The password for this credential.
	Password *string `json:"Password,omitempty" plist:"Password,omitempty"`
}

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
