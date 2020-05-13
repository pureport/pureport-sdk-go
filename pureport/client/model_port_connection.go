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

import (
	"time"
)

// PortConnection Connection using a Pureport Access Port or Packet Fabric Port.
type PortConnection struct {
	ActiveAt time.Time `json:"activeAt,omitempty"`
	// If the connection is advertising internal routes, which allows the customer the option of probing and tracing these routes.
	AdvertiseInternalRoutes bool        `json:"advertiseInternalRoutes,omitempty"`
	BillingPlan             BillingPlan `json:"billingPlan,omitempty"`
	// The provider used for billing this connection.
	BillingProvider string `json:"billingProvider,omitempty"`
	// The licensed billing term for the connection.
	BillingTerm string    `json:"billingTerm"`
	CreatedAt   time.Time `json:"createdAt,omitempty"`
	// The customer side ASN. This can either be a public or private ASN. If this is a public ASN, you must own it to prevent conflicts.
	CustomerASN int64 `json:"customerASN,omitempty"`
	// Set of customer Networks for this connection.
	CustomerNetworks []CustomerNetwork `json:"customerNetworks,omitempty"`
	DeletedAt        time.Time         `json:"deletedAt,omitempty"`
	// The user defined description for the connection.
	Description string `json:"description,omitempty"`
	// Error Code assigned to the connection if it is an error state.
	ErrorCode string `json:"errorCode,omitempty"`
	// Error message assigned to the connection if it is an error state.
	ErrorMessage string `json:"errorMessage,omitempty"`
	// Whether this connection has redundant gateways for failover.
	HighAvailability bool `json:"highAvailability"`
	// The URI of the Pureport asset.
	Href string `json:"href,omitempty"`
	// The id is a unique identifier representing the connection. This can be provided during creation, but if left empty, will be generated.
	Id       string `json:"id,omitempty"`
	Location Link   `json:"location"`
	// The user specified name for the connection.
	Name             string    `json:"name"`
	Nat              NatConfig `json:"nat,omitempty"`
	Network          Link      `json:"network,omitempty"`
	PrimaryGateway   Gateway   `json:"primaryGateway,omitempty"`
	SecondaryGateway Gateway   `json:"secondaryGateway,omitempty"`
	// The connection speed in Mbps.
	Speed int32 `json:"speed"`
	// The current state of the connection.
	State string `json:"state,omitempty"`
	// Key-value pairs to associate with the Pureport asset.
	Tags map[string]string `json:"tags,omitempty"`
	// The connection type.
	Type                     string                   `json:"type"`
	BgpPasswordConfiguration BgpPasswordConfiguration `json:"bgpPasswordConfiguration,omitempty"`
	// The long haul billing configuration for this connection.
	BillingLongHaul string `json:"billingLongHaul,omitempty"`
	GatewayCidr     string `json:"gatewayCidr,omitempty"`
	// The VLAN ID of the primary gateway.
	PrimaryCustomerVlan int32  `json:"primaryCustomerVlan"`
	PrimaryGatewayIP    string `json:"primaryGatewayIP,omitempty"`
	PrimaryPort         Link   `json:"primaryPort"`
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
