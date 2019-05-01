//go:generate stringer -linecomment=true -tags=ENC -type=DiffieHellmanGroup
package ikev2

type DiffieHellmanGroup int

const (
	MODP_1024 DiffieHellmanGroup = iota // MODP_1024
	MODP_1536                           // MODP_1536
	MODP_2048                           // MODP_2048
	MODP_3072                           // MODP_3072
	MODP_4096                           // MODP_4096
	MODP_6144                           // MODP_6144
	MODP_8192                           // MODP_8192
	ECP_192                             // ECP_192
	ECP_224                             // ECP_224
	ECP_256                             // ECP_256
	ECP_384                             // ECP_384
	ECP_521                             // ECP_521
)
