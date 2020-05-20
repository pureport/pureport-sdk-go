/*
 * Pureport Control Plane
 *
 * Pureport API
 *
 * API version: 1.0.0
 * Contact: support@pureport.com
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package client

// Ikev2IkeConfig IKE Version 2 Internet Key Exchange configuration.
type Ikev2IkeConfig struct {
	DhGroup Ikev2DiffieHellmanGroup `json:"dhGroup"`
	// The Dead Peering Detection (DPD) retry interval.
	DpdDelay int32 `json:"dpdDelay,omitempty"`
	// The Dead Peering Detection (DPD) interval.
	DpdTimeout int32                       `json:"dpdTimeout,omitempty"`
	Encryption Ikev2IkeEncryptionAlgorithm `json:"encryption"`
	Integrity  Ikev2IkeIntegrityAlgorithm  `json:"integrity,omitempty"`
	// The lifetime of the SA.
	Lifetime int32                     `json:"lifetime,omitempty"`
	Prf      Ikev2PseudoRandomFunction `json:"prf,omitempty"`
}
