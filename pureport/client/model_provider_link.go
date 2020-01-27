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

// ProviderLink struct for ProviderLink
type ProviderLink struct {
	AvailabilityDomain       string   `json:"availabilityDomain,omitempty"`
	Href                     string   `json:"href,omitempty"`
	Id                       string   `json:"id"`
	MaxAvailableBandwidth    int32    `json:"maxAvailableBandwidth,omitempty"`
	Pod                      Link     `json:"pod,omitempty"`
	Provider                 string   `json:"provider,omitempty"`
	ProviderId               string   `json:"providerId,omitempty"`
	ProviderRegion           string   `json:"providerRegion,omitempty"`
	Speed                    int32    `json:"speed,omitempty"`
	Status                   string   `json:"status,omitempty"`
	SupportedConnectionTypes []string `json:"supportedConnectionTypes,omitempty"`
	Vlans                    string   `json:"vlans,omitempty"`
}
