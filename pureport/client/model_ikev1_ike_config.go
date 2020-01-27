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

// Ikev1IkeConfig struct for Ikev1IkeConfig
type Ikev1IkeConfig struct {
	DhGroup    string `json:"dhGroup"`
	DpdDelay   int32  `json:"dpdDelay,omitempty"`
	DpdTimeout int32  `json:"dpdTimeout,omitempty"`
	Encryption string `json:"encryption"`
	Integrity  string `json:"integrity"`
	Lifetime   int32  `json:"lifetime,omitempty"`
}
