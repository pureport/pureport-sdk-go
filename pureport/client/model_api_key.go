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

// ApiKey Key for authenticating with the Pureport API
type ApiKey struct {
	Account Link `json:"account,omitempty"`
	// The description.
	Description string `json:"description,omitempty"`
	// The URI of the Pureport asset.
	Href string `json:"href,omitempty"`
	// The public key for authenticating.
	Key string `json:"key,omitempty"`
	// The name.
	Name string `json:"name"`
	// The Pureport roles users of this API key have permissions for.
	Roles []Link `json:"roles"`
	// The private secret for authenticating. Note that this field is only populated on creation.
	Secret string `json:"secret,omitempty"`
}
