//go:generate stringer -linecomment=true -tags=ENC -type=PseudoRandomFunction
package ikev2

type PseudoRandomFunction int

const (
	MD5      PseudoRandomFunction = iota // MD5
	SHA_1                                // SHA_1
	AES_XCBC                             // AES_XCBC
	SHA_256                              // SHA_256
	SHA_384                              // SHA_384
	SHA_512                              // SHA_512
)
