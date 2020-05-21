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

// Ikev2EspIntegrityAlgorithm The integrity algorithm.
type Ikev2EspIntegrityAlgorithm string

// List of IKEV2ESPIntegrityAlgorithm
const (
	MD5_HMAC    Ikev2EspIntegrityAlgorithm = "MD5_HMAC"
	SHA1_HMAC   Ikev2EspIntegrityAlgorithm = "SHA1_HMAC"
	SHA256_HMAC Ikev2EspIntegrityAlgorithm = "SHA256_HMAC"
	SHA384_HMAC Ikev2EspIntegrityAlgorithm = "SHA384_HMAC"
	SHA512_HMAC Ikev2EspIntegrityAlgorithm = "SHA512_HMAC"
	AES_XCBC    Ikev2EspIntegrityAlgorithm = "AES_XCBC"
)
