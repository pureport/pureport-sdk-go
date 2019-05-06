//go:generate stringer -linecomment=true -tags=ENC -type=EncryptionAlgorithm
package ike

type EncryptionAlgorithm int

const (
	AES_128 EncryptionAlgorithm = iota // AES_128
	AES_192                            // AES_192
	AES_256                            // AES_256
)
