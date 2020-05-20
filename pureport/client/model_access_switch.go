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

// AccessSwitch Access switch installed in a Pureport pod.
type AccessSwitch struct {
	AvailabilityDomain AvailabilityDomain `json:"availabilityDomain"`
	// The URI of the Pureport asset.
	Href string `json:"href,omitempty"`
	// The id is a unique identifier representing the access switch.
	Id     string             `json:"id"`
	Pod    Link               `json:"pod"`
	Status AccessSwitchStatus `json:"status"`
}
