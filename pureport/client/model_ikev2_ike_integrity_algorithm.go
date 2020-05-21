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

// Ikev2IkeIntegrityAlgorithm The integrity algorithm.
type Ikev2IkeIntegrityAlgorithm string

// List of IKEV2IKEIntegrityAlgorithm
const (
	IKEV2_IKE__MD5_HMAC    Ikev2IkeIntegrityAlgorithm = "MD5_HMAC"
	IKEV2_IKE__SHA1_HMAC   Ikev2IkeIntegrityAlgorithm = "SHA1_HMAC"
	IKEV2_IKE__SHA256_HMAC Ikev2IkeIntegrityAlgorithm = "SHA256_HMAC"
	IKEV2_IKE__SHA384_HMAC Ikev2IkeIntegrityAlgorithm = "SHA384_HMAC"
	IKEV2_IKE__SHA512_HMAC Ikev2IkeIntegrityAlgorithm = "SHA512_HMAC"
	IKEV2_IKE__AES_XCBC    Ikev2IkeIntegrityAlgorithm = "AES_XCBC"
)
