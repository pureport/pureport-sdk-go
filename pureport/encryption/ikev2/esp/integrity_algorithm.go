//go:generate stringer -linecomment=true -tags=ENC -type=IntegrityAlgorithm
package esp

type IntegrityAlgorithm int

const (
	MD5_HMAC    IntegrityAlgorithm = iota // MD5_HMAC
	SHA1_HMAC                             // SHA1_HMAC
	SHA256_HMAC                           // SHA256_HMAC
	SHA384_HMAC                           // SHA384_HMAC
	SHA512_HMAC                           // SHA512_HMAC
	AES_XCBC                              // AES_XCBC
)
