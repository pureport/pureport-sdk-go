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

// Ikev1EspIntegrityAlgorithm The integrity algorithm.
type Ikev1EspIntegrityAlgorithm string

// List of IKEV1ESPIntegrityAlgorithm
const (
	MD5_HMAC    Ikev1EspIntegrityAlgorithm = "MD5_HMAC"
	SHA1_HMAC   Ikev1EspIntegrityAlgorithm = "SHA1_HMAC"
	SHA256_HMAC Ikev1EspIntegrityAlgorithm = "SHA256_HMAC"
	SHA384_HMAC Ikev1EspIntegrityAlgorithm = "SHA384_HMAC"
	SHA512_HMAC Ikev1EspIntegrityAlgorithm = "SHA512_HMAC"
	AES_XCBC    Ikev1EspIntegrityAlgorithm = "AES_XCBC"
)
