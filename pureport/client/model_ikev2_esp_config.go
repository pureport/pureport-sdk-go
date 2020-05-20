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

// Ikev2EspConfig IKE Version 2 Encapsulating Security Payload configuration.
type Ikev2EspConfig struct {
	DhGroup    Ikev2DiffieHellmanGroup     `json:"dhGroup"`
	Encryption Ikev2EspEncryptionAlgorithm `json:"encryption"`
	Integrity  Ikev2EspIntegrityAlgorithm  `json:"integrity,omitempty"`
	// The lifetime of the Security Association (SA).
	Lifetime int32 `json:"lifetime,omitempty"`
}
