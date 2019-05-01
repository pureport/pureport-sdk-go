//go:generate stringer -linecomment=true -tags=ENC -type=EncryptionAlgorithm
package ike

type EncryptionAlgorithm int

const (
	NULL            EncryptionAlgorithm = iota // NULL
	AES_128                                    // AES_128
	AES_192                                    // AES_192
	AES_256                                    // AES_256
	AES_128_CTR                                // AES_128_CTR
	AES_192_CTR                                // AES_192_CTR
	AES_256_CTR                                // AES_256_CTR
	AES_128_GCM_64                             // AES_128_GCM_64
	AES_192_GCM_64                             // AES_192_GCM_64
	AES_256_GCM_64                             // AES_256_GCM_64
	AES_128_GCM_96                             // AES_128_GCM_96
	AES_192_GCM_96                             // AES_192_GCM_96
	AES_256_GCM_96                             // AES_256_GCM_96
	AES_128_GCM_128                            // AES_128_GCM_128
	AES_192_GCM_128                            // AES_192_GCM_128
	AES_256_GCM_128                            // AES_256_GCM_128
)
