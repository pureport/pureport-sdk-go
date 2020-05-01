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

// PortConnectionAllOf struct for PortConnectionAllOf
type PortConnectionAllOf struct {
	BgpPasswordConfiguration BgpPasswordConfiguration `json:"bgpPasswordConfiguration,omitempty"`
	// The long haul billing configuration for this connection.
	BillingLongHaul string `json:"billingLongHaul,omitempty"`
	GatewayCidr     string `json:"gatewayCidr,omitempty"`
	// The VLAN ID of the primary gateway.
	PrimaryCustomerVlan int32  `json:"primaryCustomerVlan,omitempty"`
	PrimaryGatewayIP    string `json:"primaryGatewayIP,omitempty"`
	PrimaryPort         Link   `json:"primaryPort,omitempty"`
	// The method to use for determining network routes.
	RoutingType string `json:"routingType,omitempty"`
	// The VLAN ID of the secondary gateway if this is an HA connection.
	SecondaryCustomerVlan int32  `json:"secondaryCustomerVlan,omitempty"`
	SecondaryGatewayIP    string `json:"secondaryGatewayIP,omitempty"`
	SecondaryPort         Link   `json:"secondaryPort,omitempty"`
	// The user configured static routes.
	StaticRoutes     []StaticRoute `json:"staticRoutes,omitempty"`
	VirtualGatewayIP string        `json:"virtualGatewayIP,omitempty"`
}
