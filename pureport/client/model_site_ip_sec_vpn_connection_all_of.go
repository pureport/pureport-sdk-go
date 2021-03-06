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

// SiteIpSecVpnConnectionAllOf struct for SiteIpSecVpnConnectionAllOf
type SiteIpSecVpnConnectionAllOf struct {
	AuthType                 IpSecAuthenticationType   `json:"authType,omitempty"`
	BgpPasswordConfiguration *BgpPasswordConfiguration `json:"bgpPasswordConfiguration,omitempty"`
	EnableBGPPassword        bool                      `json:"enableBGPPassword,omitempty"`
	IkeV1                    *Ikev1Config              `json:"ikeV1,omitempty"`
	IkeV2                    *Ikev2Config              `json:"ikeV2,omitempty"`
	IkeVersion               IkeVersion                `json:"ikeVersion,omitempty"`
	PhysicalAddress          PhysicalAddress           `json:"physicalAddress,omitempty"`
	PrimaryCustomerRouterIP  string                    `json:"primaryCustomerRouterIP,omitempty"`
	// IPsec pre-shared key (PSK) override for the primary gateway.
	PrimaryKey                string         `json:"primaryKey,omitempty"`
	RoutingType               VpnRoutingType `json:"routingType,omitempty"`
	SecondaryCustomerRouterIP string         `json:"secondaryCustomerRouterIP,omitempty"`
	// IPsec pre-shared key (PSK) override for the secondary gateway.
	SecondaryKey string `json:"secondaryKey,omitempty"`
	// The traffic selectors to apply for routing.
	TrafficSelectors []TrafficSelectorMapping `json:"trafficSelectors,omitempty"`
}
