// DO NOT EDIT
// generated from https://github.com/apple/device-management.git:b838baacf2e790db729b6ca3f52724adc8bfb96d/declarative/declarations/assets

package assets

const DeviceManagementGenerateHash = "b838baacf2e790db729b6ca3f52724adc8bfb96d"

// Type of authentication:
// * MDM - a request using MDM semantics (includes the device identity certificate, and any user authentication). Equivalent to an MDM request made to the CheckInURL or ServerURL. This option can only be used when using declarative device management.
// * None - a standard GET request is carried out.
type AssetCredentialACMEAuthenticationType string

const (
	AssetCredentialACMEAuthenticationTypeMDM  AssetCredentialACMEAuthenticationType = "MDM"
	AssetCredentialACMEAuthenticationTypeNone AssetCredentialACMEAuthenticationType = "None"
)

// The keychain accessibility that determines when the keychain item is available for use.
// * Default - the most restrictive accessibility that still satisfies all uses of the asset by configurations that reference it
// * AfterFirstUnlock - the keychain item is only available after first unlock of the device
type AssetCredentialACMEAccessible string

const (
	AssetCredentialACMEAccessibleDefault          AssetCredentialACMEAccessible = "Default"
	AssetCredentialACMEAccessibleAfterFirstUnlock AssetCredentialACMEAccessible = "AfterFirstUnlock"
)

// Type of authentication:
// * MDM - a request using MDM semantics (includes the device identity certificate, and any user authentication). Equivalent to an MDM request made to the CheckInURL or ServerURL. This option can only be used when using declarative device management.
// * None - a standard GET request is carried out.
type AssetCredentialCertificateAuthenticationType string

const (
	AssetCredentialCertificateAuthenticationTypeMDM  AssetCredentialCertificateAuthenticationType = "MDM"
	AssetCredentialCertificateAuthenticationTypeNone AssetCredentialCertificateAuthenticationType = "None"
)

// Type of authentication:
// * MDM - a request using MDM semantics (includes the device identity certificate, and any user authentication). Equivalent to an MDM request made to the CheckInURL or ServerURL. This option can only be used when using declarative device management.
// * None - a standard GET request is carried out.
type AssetCredentialIdentityAuthenticationType string

const (
	AssetCredentialIdentityAuthenticationTypeMDM  AssetCredentialIdentityAuthenticationType = "MDM"
	AssetCredentialIdentityAuthenticationTypeNone AssetCredentialIdentityAuthenticationType = "None"
)

// The keychain accessibility that determines when the keychain item is available for use.
// * Default - the most restrictive accessibility that still satisfies all uses of the asset by configurations that reference it
// * AfterFirstUnlock - the keychain item is only available after first unlock of the device
type AssetCredentialIdentityAccessible string

const (
	AssetCredentialIdentityAccessibleDefault          AssetCredentialIdentityAccessible = "Default"
	AssetCredentialIdentityAccessibleAfterFirstUnlock AssetCredentialIdentityAccessible = "AfterFirstUnlock"
)

// Type of authentication:
// * MDM - a request using MDM semantics (includes the device identity certificate, and any user authentication). Equivalent to an MDM request made to the CheckInURL or ServerURL. This option can only be used when using declarative device management.
// * None - a standard GET request is carried out.
type AssetCredentialSCEPAuthenticationType string

const (
	AssetCredentialSCEPAuthenticationTypeMDM  AssetCredentialSCEPAuthenticationType = "MDM"
	AssetCredentialSCEPAuthenticationTypeNone AssetCredentialSCEPAuthenticationType = "None"
)

// The keychain accessibility that determines when the keychain item is available for use.
// * Default - the most restrictive accessibility that still satisfies all uses of the asset by configurations that reference it
// * AfterFirstUnlock - the keychain item is only available after first unlock of the device
type AssetCredentialSCEPAccessible string

const (
	AssetCredentialSCEPAccessibleDefault          AssetCredentialSCEPAccessible = "Default"
	AssetCredentialSCEPAccessibleAfterFirstUnlock AssetCredentialSCEPAccessible = "AfterFirstUnlock"
)

// Type of authentication:
// * MDM - a request using MDM semantics (includes the device identity certificate, and any user authentication). Equivalent to an MDM request made to the CheckInURL or ServerURL. This option can only be used when using declarative device management.
// * None - a standard GET request is carried out.
type AssetCredentialUserNameandPasswordAuthenticationType string

const (
	AssetCredentialUserNameandPasswordAuthenticationTypeMDM  AssetCredentialUserNameandPasswordAuthenticationType = "MDM"
	AssetCredentialUserNameandPasswordAuthenticationTypeNone AssetCredentialUserNameandPasswordAuthenticationType = "None"
)

// Specifies the type of key pair to generate.
// "RSA" specifies an RSA key pair. If this is specified, KeySize must be in the range [1024..4096] inclusive and a multiple of 8, and HardwareBound must be false.
// "ECSECPrimeRandom" indicates a key pair on the P-256, P-384 or P-521 curves as defined in FIPS Pub 186-4. The specific curve is selected by the KeySize, which must be 256, 384 or 521. Only 256 and 384 are supported for hardware bound keys. (Note that the key size is 521, not 512, even though the other key sizes are multiples of 64.)
type ACMECredentialKeyType string

const (
	ACMECredentialKeyTypeRSA              ACMECredentialKeyType = "RSA"
	ACMECredentialKeyTypeECSECPrimeRandom ACMECredentialKeyType = "ECSECPrimeRandom"
)

// The key size in bits, either 1024, 2048 or 4096.
type SCEPCredentialKeysize int64

const (
	SCEPCredentialKeysize1024 SCEPCredentialKeysize = 1024
	SCEPCredentialKeysize2048 SCEPCredentialKeysize = 2048
	SCEPCredentialKeysize4096 SCEPCredentialKeysize = 4096
)

// Type of authentication:
// * MDM - a request using MDM semantics (includes the device identity certificate, and any user authentication). Equivalent to an MDM request made to the CheckInURL or ServerURL. This option can only be used when using declarative device management.
// * None - a standard GET request is carried out.
type AssetDataAuthenticationType string

const (
	AssetDataAuthenticationTypeMDM  AssetDataAuthenticationType = "MDM"
	AssetDataAuthenticationTypeNone AssetDataAuthenticationType = "None"
)

// The external reference. The asset data must be a JSON document representing the "com.apple.credential.acme" credential type. The asset data must be returned using a media type of "application/json". If a "ContentType" sub-key is included, it must be set to "application/json".
type AssetCredentialACMEReference struct {
	// The URL that hosts the credential data. The URL must start with 'https://'.
	DataURL string `json:"DataURL"`
	// The media type that describes the data.
	ContentType *string `json:"ContentType,omitempty"`
	// The size of the data at the 'DataURL'. Use this value to verify that the returned data is the expected data. Use this value to detect when the data changes.
	Size *int64 `json:"Size,omitempty"`
	// A SHA-256 hash of the data at the 'DataURL'. Use this value to verify that the returned data is the expected data. Use this value to detect when the data changes.
	// If 'Size' is '0', clients need to ignore this value or set it to an empty string.
	HashSHA256 *string `json:"Hash-SHA-256,omitempty"`
}

// The server authentication details.
type AssetCredentialACMEAuthentication struct {
	// Type of authentication:
	// * MDM - a request using MDM semantics (includes the device identity certificate, and any user authentication). Equivalent to an MDM request made to the CheckInURL or ServerURL. This option can only be used when using declarative device management.
	// * None - a standard GET request is carried out.
	Type string `json:"Type"`
}

// A reference to an ACME identity.
type AssetCredentialACME struct {
	// The external reference. The asset data must be a JSON document representing the "com.apple.credential.acme" credential type. The asset data must be returned using a media type of "application/json". If a "ContentType" sub-key is included, it must be set to "application/json".
	Reference *AssetCredentialACMEReference `json:"Reference"`
	// The server authentication details.
	Authentication *AssetCredentialACMEAuthentication `json:"Authentication,omitempty"`
	// The keychain accessibility that determines when the keychain item is available for use.
	// * Default - the most restrictive accessibility that still satisfies all uses of the asset by configurations that reference it
	// * AfterFirstUnlock - the keychain item is only available after first unlock of the device
	Accessible *string `default:"Default" json:"Accessible,omitempty"`
}

func (p *AssetCredentialACME) DeclarationType() string {
	return "com.apple.asset.credential.acme"
}

// The external reference. The asset data must be returned using a media type of "application/pkcs1" or "application/pem" to correctly identify the type of encoded certificate. If a "ContentType" sub-key is included, it must be set to the corresponding media type.
type AssetCredentialCertificateReference struct {
	// The URL that hosts the credential data. The URL must start with 'https://'.
	DataURL string `json:"DataURL"`
	// The media type that describes the data.
	ContentType *string `json:"ContentType,omitempty"`
	// The size of the data at the 'DataURL'. Use this value to verify that the returned data is the expected data. Use this value to detect when the data changes.
	Size *int64 `json:"Size,omitempty"`
	// A SHA-256 hash of the data at the 'DataURL'. Use this value to verify that the returned data is the expected data. Use this value to detect when the data changes.
	// If 'Size' is '0', clients need to ignore this value or set it to an empty string.
	HashSHA256 *string `json:"Hash-SHA-256,omitempty"`
}

// The server authentication details.
type AssetCredentialCertificateAuthentication struct {
	// Type of authentication:
	// * MDM - a request using MDM semantics (includes the device identity certificate, and any user authentication). Equivalent to an MDM request made to the CheckInURL or ServerURL. This option can only be used when using declarative device management.
	// * None - a standard GET request is carried out.
	Type string `json:"Type"`
}

// A reference to a PKCS #1 or PEM encoded certificate.
type AssetCredentialCertificate struct {
	// The external reference. The asset data must be returned using a media type of "application/pkcs1" or "application/pem" to correctly identify the type of encoded certificate. If a "ContentType" sub-key is included, it must be set to the corresponding media type.
	Reference *AssetCredentialCertificateReference `json:"Reference"`
	// The server authentication details.
	Authentication *AssetCredentialCertificateAuthentication `json:"Authentication,omitempty"`
}

func (p *AssetCredentialCertificate) DeclarationType() string {
	return "com.apple.asset.credential.certificate"
}

// The external reference. The asset data must be a JSON document representing the "com.apple.credential.identity" credential type. The asset data must be returned using a media type of "application/json". If a "ContentType" sub-key is included, it must be set to "application/json".
type AssetCredentialIdentityReference struct {
	// The URL that hosts the credential data. The URL must start with 'https://'.
	DataURL string `json:"DataURL"`
	// The media type that describes the data.
	ContentType *string `json:"ContentType,omitempty"`
	// The size of the data at the 'DataURL'. Use this value to verify that the returned data is the expected data. Use this value to detect when the data changes.
	Size *int64 `json:"Size,omitempty"`
	// A SHA-256 hash of the data at the 'DataURL'. Use this value to verify that the returned data is the expected data. Use this value to detect when the data changes.
	// If 'Size' is '0', clients need to ignore this value or set it to an empty string.
	HashSHA256 *string `json:"Hash-SHA-256,omitempty"`
}

// The server authentication details.
type AssetCredentialIdentityAuthentication struct {
	// Type of authentication:
	// * MDM - a request using MDM semantics (includes the device identity certificate, and any user authentication). Equivalent to an MDM request made to the CheckInURL or ServerURL. This option can only be used when using declarative device management.
	// * None - a standard GET request is carried out.
	Type string `json:"Type"`
}

// A reference to a PKCS #12 password-protected identity.
type AssetCredentialIdentity struct {
	// The external reference. The asset data must be a JSON document representing the "com.apple.credential.identity" credential type. The asset data must be returned using a media type of "application/json". If a "ContentType" sub-key is included, it must be set to "application/json".
	Reference *AssetCredentialIdentityReference `json:"Reference"`
	// The server authentication details.
	Authentication *AssetCredentialIdentityAuthentication `json:"Authentication,omitempty"`
	// The keychain accessibility that determines when the keychain item is available for use.
	// * Default - the most restrictive accessibility that still satisfies all uses of the asset by configurations that reference it
	// * AfterFirstUnlock - the keychain item is only available after first unlock of the device
	Accessible *string `default:"Default" json:"Accessible,omitempty"`
}

func (p *AssetCredentialIdentity) DeclarationType() string {
	return "com.apple.asset.credential.identity"
}

// The external reference. The asset data must be a JSON document representing the "com.apple.credential.scep" credential type. The asset data must be returned using a media type of "application/json". If a "ContentType" sub-key is included, it must be set to "application/json".
type AssetCredentialSCEPReference struct {
	// The URL that hosts the credential data. The URL must start with 'https://'.
	DataURL string `json:"DataURL"`
	// The media type that describes the data.
	ContentType *string `json:"ContentType,omitempty"`
	// The size of the data at the 'DataURL'. Use this value to verify that the returned data is the expected data. Use this value to detect when the data changes.
	Size *int64 `json:"Size,omitempty"`
	// A SHA-256 hash of the data at the 'DataURL'. Use this value to verify that the returned data is the expected data. Use this value to detect when the data changes.
	// If 'Size' is '0', clients need to ignore this value or set it to an empty string.
	HashSHA256 *string `json:"Hash-SHA-256,omitempty"`
}

// The server authentication details.
type AssetCredentialSCEPAuthentication struct {
	// Type of authentication:
	// * MDM - a request using MDM semantics (includes the device identity certificate, and any user authentication). Equivalent to an MDM request made to the CheckInURL or ServerURL. This option can only be used when using declarative device management.
	// * None - a standard GET request is carried out.
	Type string `json:"Type"`
}

// A reference to a SCEP identity.
type AssetCredentialSCEP struct {
	// The external reference. The asset data must be a JSON document representing the "com.apple.credential.scep" credential type. The asset data must be returned using a media type of "application/json". If a "ContentType" sub-key is included, it must be set to "application/json".
	Reference *AssetCredentialSCEPReference `json:"Reference"`
	// The server authentication details.
	Authentication *AssetCredentialSCEPAuthentication `json:"Authentication,omitempty"`
	// The keychain accessibility that determines when the keychain item is available for use.
	// * Default - the most restrictive accessibility that still satisfies all uses of the asset by configurations that reference it
	// * AfterFirstUnlock - the keychain item is only available after first unlock of the device
	Accessible *string `default:"Default" json:"Accessible,omitempty"`
}

func (p *AssetCredentialSCEP) DeclarationType() string {
	return "com.apple.asset.credential.scep"
}

// The reference to the credential.
type AssetCredentialUserNameandPasswordReference struct {
	// The URL that hosts the credential data. The URL must start with 'https://'.
	DataURL string `json:"DataURL"`
	// The media type that describes the data.
	ContentType *string `json:"ContentType,omitempty"`
	// The size of the data at the 'DataURL'. Use this value to verify that the returned data is the expected data. Use this value to detect when the data changes.
	Size *int64 `json:"Size,omitempty"`
	// A SHA-256 hash of the data at the 'DataURL'. Use this value to verify that the returned data is the expected data. Use this value to detect when the data changes.
	// If 'Size' is '0', clients need to ignore this value or set it to an empty string.
	HashSHA256 *string `json:"Hash-SHA-256,omitempty"`
}

// The server authentication details.
type AssetCredentialUserNameandPasswordAuthentication struct {
	// Type of authentication:
	// * MDM - a request using MDM semantics (includes the device identity certificate, and any user authentication). Equivalent to an MDM request made to the CheckInURL or ServerURL. This option can only be used when using declarative device management.
	// * None - a standard GET request is carried out.
	Type string `json:"Type"`
}

// A reference to data describing a credential representing a user name and password.
type AssetCredentialUserNameandPassword struct {
	// The reference to the credential.
	Reference *AssetCredentialUserNameandPasswordReference `json:"Reference"`
	// The server authentication details.
	Authentication *AssetCredentialUserNameandPasswordAuthentication `json:"Authentication,omitempty"`
}

func (p *AssetCredentialUserNameandPassword) DeclarationType() string {
	return "com.apple.asset.credential.userpassword"
}

// Specifies the Subject Alt Name that the device will request for the certificate that the ACME server issues. The ACME server may override or ignore this field in the certificate it issues.
type ACMECredentialSubjectAltName struct {
	// RFC 822 (email address) string
	Rfc822Name *string `json:"rfc822Name,omitempty"`
	// The DNS name
	DNSName *string `json:"dNSName,omitempty"`
	// The Uniform Resource Identifier
	UniformResourceIdentifier *string `json:"uniformResourceIdentifier,omitempty"`
	// The NT principal name
	NtPrincipalName *string `json:"ntPrincipalName,omitempty"`
}

// An ACME identity that should be generated by the device.
type ACMECredential struct {
	// Specifies the directory URL of the ACME server. The URL must use the https scheme.
	DirectoryURL string `json:"DirectoryURL"`
	// The server may use this as a nonce to prevent issuing multiple certificates. It also indicates to the ACME server that the device has access to a valid client identifier that was issued by the enterprise infrastructure. This may help the ACME server determine whether to trust the device, however this is a relatively weak indication because of the risk that the client identifier may be intercepted and duplicated by an attacker.
	ClientIdentifier string `json:"ClientIdentifier"`
	// The valid values for KeySize depend on the values of KeyType and HardwareBound. See those keys for specific requirements.
	KeySize int64 `json:"KeySize"`
	// Specifies the type of key pair to generate.
	// "RSA" specifies an RSA key pair. If this is specified, KeySize must be in the range [1024..4096] inclusive and a multiple of 8, and HardwareBound must be false.
	// "ECSECPrimeRandom" indicates a key pair on the P-256, P-384 or P-521 curves as defined in FIPS Pub 186-4. The specific curve is selected by the KeySize, which must be 256, 384 or 521. Only 256 and 384 are supported for hardware bound keys. (Note that the key size is 521, not 512, even though the other key sizes are multiples of 64.)
	KeyType string `json:"KeyType"`
	// If false, the private key is not bound to the device.
	// If true, the private key is bound to the device. The Secure Enclave generates the key pair, and the private key is cryptographically entangled with a system key. This protects the private key from being exported.
	// If true, KeyType must be ECSECPrimeRandom and KeySize must be 256 or 384.
	// On macOS, this key is required but must have a value of false.
	HardwareBound bool `json:"HardwareBound"`
	// The device requests this subject for the certificate that the ACME server issues. The ACME server may override or ignore this field in the certificate it issues.
	// The representation of a X.500 name represented as an array of OID and value. For example, /C=US/O=Apple Inc./CN=foo/1.2.5.3=bar corresponds to:
	// [ [ ["C", "US"] ], [ ["O", "Apple Inc."] ], ..., [ [ "1.2.5.3", "bar" ] ] ]
	// OIDs can be represented as dotted numbers, with shortcuts for country (C), locality (L), state (ST), organization (O), organizational unit (OU), and common name (CN).
	Subject [][][]string `json:"Subject"`
	// Specifies the Subject Alt Name that the device will request for the certificate that the ACME server issues. The ACME server may override or ignore this field in the certificate it issues.
	SubjectAltName *ACMECredentialSubjectAltName `json:"SubjectAltName,omitempty"`
	// The device requests this key usage for the certificate that the ACME server issues. The ACME server may override or ignore this field in the certificate it issues.
	// The value is a bit field. Bit 0x01 indicates digital signature, and bit 0x04 indicates key encipherment.
	UsageFlags *int64 `json:"UsageFlags,omitempty"`
	// The device requests this extended key usage for the certificate that the ACME server issues. The ACME server may override or ignore this field in the certificate it issues.
	// The value is an array of strings. Each string is an OID in dotted notation. For instance, ["1.3.6.1.5.5.7.3.2", "1.3.6.1.5.5.7.3.4"] indicates client authentication and email protection.
	ExtendedKeyUsage *[]string `json:"ExtendedKeyUsage,omitempty"`
	// When true, the device provides attestations describing the device and the generated key to the ACME server. The server can use the attestations as strong evidence that the key is bound to the device, and that the device has properties listed in the attestation. The server can use that as part of a trust score to decide whether to issue the requested certificate. When Attest is true, HardwareBound must also be true. On macOS, this key, if present, must have a value of false.
	Attest *bool `default:"false" json:"Attest,omitempty"`
}

// Data for a PKCS #12 password-protected identity.
type IdentityCredential struct {
	// The password required to decrypt the PKCS #12 identity data.
	Password string `json:"Password"`
	// The PKCS #12 identity data.
	Identity []byte `json:"Identity"`
}

// Specifies the Subject Alt Name for the certificate
type SCEPCredentialSubjectAltName struct {
	// RFC 822 (email address) string
	Rfc822Name *string `json:"rfc822Name,omitempty"`
	// The DNS name
	DNSName *string `json:"dNSName,omitempty"`
	// The URI
	UniformResourceIdentifier *string `json:"uniformResourceIdentifier,omitempty"`
	// The NT principal name
	NtPrincipalName *string `json:"ntPrincipalName,omitempty"`
}

// A SCEP identity that should be generated by the device.
type SCEPCredential struct {
	// The SCEP URL. See Over-the-Air Profile Delivery and Configuration for more information about SCEP.
	URL string `json:"URL"`
	// Any string that is understood by the SCEP server. For example, it could be a domain name like example.org. If a certificate authority has multiple CA certificates this field can be used to distinguish which is required.
	Name *string `json:"Name,omitempty"`
	// The representation of a X.500 name represented as an array of OID and value. For example, /C=US/O=Apple Inc./CN=foo/1.2.5.3=bar, which would translate to:
	// [ [ ["C", "US"] ], [ ["O", "Apple Inc."] ], ..., [ [ "1.2.5.3", "bar" ] ] ] OIDs can be represented as dotted numbers, with shortcuts for country (C), locality (L), state (ST), organization (O), organizational unit (OU), and common name (CN).
	Subject *[][][]string `json:"Subject,omitempty"`
	// A pre-shared secret.
	Challenge *string `json:"Challenge,omitempty"`
	// The key size in bits, either 1024, 2048 or 4096.
	Keysize *int64 `default:"1024" json:"Keysize,omitempty"`
	// Currently always "RSA".
	KeyType *string `default:"RSA" json:"Key Type,omitempty"`
	// A bitmask indicating the use of the key. 1 is signing, 4 is encryption, 5 is both signing and encryption. Some certificate authorities, such as Windows CA, support only encryption or signing, but not both at the same time.
	KeyUsage *int64 `default:"0" json:"Key Usage,omitempty"`
	// The fingerprint of the Certificate Authority certificate.
	CAFingerprint *[]byte `json:"CAFingerprint,omitempty"`
	// The number of times the device should retry if the server sends a PENDING response.
	Retries *int64 `default:"3" json:"Retries,omitempty"`
	// Number of seconds to wait between subsequent retries. The first retry is attempted without this delay.
	RetryDelay *int64 `default:"10" json:"RetryDelay,omitempty"`
	// Specifies the Subject Alt Name for the certificate
	SubjectAltName *SCEPCredentialSubjectAltName `json:"SubjectAltName,omitempty"`
}

// Data describing a credential representing a user name and password.
type UserNameandPasswordCredential struct {
	// The user's user name for the credential.
	UserName string `json:"UserName"`
	// The user's password for the credential.
	Password *string `json:"Password,omitempty"`
}

// The reference to the data.
type AssetDataReference struct {
	// The URL that hosts the credential data. The URL must start with 'https://'.
	DataURL string `json:"DataURL"`
	// The media type that describes the data.
	ContentType *string `json:"ContentType,omitempty"`
	// The size of the data at the 'DataURL'. Use this value to verify that the returned data is the expected data. Use this value to detect when the data changes.
	Size *int64 `json:"Size,omitempty"`
	// A SHA-256 hash of the data at the 'DataURL'. Use this value to verify that the returned data is the expected data. Use this value to detect when the data changes.
	// If 'Size' is '0', clients need to ignore this value or set it to an empty string.
	HashSHA256 *string `json:"Hash-SHA-256,omitempty"`
}

// The server authentication details.
type AssetDataAuthentication struct {
	// Type of authentication:
	// * MDM - a request using MDM semantics (includes the device identity certificate, and any user authentication). Equivalent to an MDM request made to the CheckInURL or ServerURL. This option can only be used when using declarative device management.
	// * None - a standard GET request is carried out.
	Type string `json:"Type"`
}

// A reference to arbitrary data with a specific media type.
type AssetData struct {
	// The reference to the data.
	Reference *AssetDataReference `json:"Reference"`
	// The server authentication details.
	Authentication *AssetDataAuthentication `json:"Authentication,omitempty"`
}

func (p *AssetData) DeclarationType() string {
	return "com.apple.asset.data"
}

// User identity data.
type AssetUserIdentity struct {
	// The user's full name.
	FullName *string `json:"FullName,omitempty"`
	// The user's email address.
	EmailAddress *string `json:"EmailAddress,omitempty"`
}

func (p *AssetUserIdentity) DeclarationType() string {
	return "com.apple.asset.useridentity"
}
