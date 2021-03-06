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

// PskAuthConfig PSK authentication configuration
type PskAuthConfig struct {
	Type IpSecAuthenticationType `json:"type,omitempty"`
	// The Pre-Shared Key.
	Key string `json:"key,omitempty"`
}
