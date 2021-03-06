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

// Ikev1Config IKE Version 1 configuration
type Ikev1Config struct {
	Esp Ikev1EspConfig `json:"esp"`
	Ike Ikev1IkeConfig `json:"ike"`
}
