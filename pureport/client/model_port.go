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

// Port Networking port
type Port struct {
	Account            Link               `json:"account"`
	AvailabilityDomain AvailabilityDomain `json:"availabilityDomain"`
	// Whether this port is available for use in child accounts.
	AvailableToChildAccounts bool        `json:"availableToChildAccounts,omitempty"`
	BillingPlan              BillingPlan `json:"billingPlan,omitempty"`
	BillingTerm              BillingTerm `json:"billingTerm"`
	// The description.
	Description string `json:"description,omitempty"`
	// The id of the device hosting this port.
	DeviceId string `json:"deviceId,omitempty"`
	Facility Link   `json:"facility"`
	// The URI of the Pureport asset.
	Href string `json:"href,omitempty"`
	// The id is a unique identifier representing the port.
	Id string `json:"id,omitempty"`
	// The id of the interface the port is using.
	InterfaceId string `json:"interfaceId,omitempty"`
	// The customer name to appear on a Letter of Authorization.
	LoaCustomerName string `json:"loaCustomerName,omitempty"`
	// The media type for this port.
	MediaType string `json:"mediaType"`
	// The name.
	Name string `json:"name"`
	// The id of the patch panel the port is using.
	PatchPanelId string `json:"patchPanelId,omitempty"`
	// The ports on the patch panel the port is using.
	PatchPanelPortNumbers []int32      `json:"patchPanelPortNumbers,omitempty"`
	Provider              PortProvider `json:"provider"`
	Speed                 PortSpeed    `json:"speed"`
	State                 PortState    `json:"state,omitempty"`
	// Key-value pairs to associate with the Pureport asset.
	Tags map[string]string `json:"tags,omitempty"`
}
