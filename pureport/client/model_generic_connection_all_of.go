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

// GenericConnectionAllOf struct for GenericConnectionAllOf
type GenericConnectionAllOf struct {
	BgpPasswordConfiguration *BgpPasswordConfiguration `json:"bgpPasswordConfiguration,omitempty"`
	GatewayCidr              string                    `json:"gatewayCidr,omitempty"`
	Peering                  *PeeringConfiguration     `json:"peering,omitempty"`
	PrimaryGatewayIP         string                    `json:"primaryGatewayIP,omitempty"`
	// The primary VLAN ID.
	PrimaryVlan        int32       `json:"primaryVlan,omitempty"`
	RoutingType        RoutingType `json:"routingType,omitempty"`
	SecondaryGatewayIP string      `json:"secondaryGatewayIP,omitempty"`
	// The secondary VLAN ID if this is an HA connection.
	SecondaryVlan int32 `json:"secondaryVlan,omitempty"`
	// The user configured static routes.
	StaticRoutes     []StaticRoute `json:"staticRoutes,omitempty"`
	VirtualGatewayIP string        `json:"virtualGatewayIP,omitempty"`
}
