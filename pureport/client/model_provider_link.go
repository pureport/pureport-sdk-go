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

// ProviderLink Link between Pureport infrastructure and an external connection provider.
type ProviderLink struct {
	// The availability domain.
	AvailabilityDomain string `json:"availabilityDomain"`
	// The URI of the Pureport asset.
	Href string `json:"href,omitempty"`
	// The id is a unique identifier representing the provider link.
	Id string `json:"id"`
	// The maximum available bandwidth of the link in Mbps.
	MaxAvailableBandwidth int32 `json:"maxAvailableBandwidth"`
	Pod                   Link  `json:"pod"`
	// The provider of connections using this link.
	Provider string `json:"provider"`
	// The id of the provider side resource.
	ProviderId string `json:"providerId"`
	// The region of the provider side resource.
	ProviderRegion string `json:"providerRegion,omitempty"`
	// The speed of the link in Mbps.
	Speed int32 `json:"speed"`
	// The current status of this link.
	Status string `json:"status"`
	// The types of connection this provider link will be used for.
	SupportedConnectionTypes []string `json:"supportedConnectionTypes"`
	// A range of vlans available on this link.
	Vlans string `json:"vlans"`
}
