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

// Ikev1EspConfig IKE Version 1 Encapsulating Security Payload configuration.
type Ikev1EspConfig struct {
	DhGroup    Ikev1DiffieHellmanGroup     `json:"dhGroup"`
	Encryption Ikev1EspEncryptionAlgorithm `json:"encryption"`
	Integrity  Ikev1EspIntegrityAlgorithm  `json:"integrity,omitempty"`
	// The lifetime of the Security Association (SA).
	Lifetime int32 `json:"lifetime,omitempty"`
}
