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

// LoginRequest Authenticate to the Pureport API
type LoginRequest struct {
	// The key to authenticate with.
	Key string `json:"key"`
	// The secret for the key to authenticate with.
	Secret string `json:"secret"`
}
