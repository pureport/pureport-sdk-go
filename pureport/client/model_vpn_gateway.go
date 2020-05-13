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

// VpnGateway Pureport network gateway for VPN connections.
type VpnGateway struct {
	// The availability domain assigned to the gateway.
	AvailabilityDomain string    `json:"availabilityDomain,omitempty"`
	BgpConfig          BgpConfig `json:"bgpConfig,omitempty"`
	Connection         Link      `json:"connection,omitempty"`
	// The description.
	Description string `json:"description,omitempty"`
	// The HTTP error code if the gateway is in an error state.
	ErrorCode string `json:"errorCode,omitempty"`
	// The error message if the gateway is in an error state.
	ErrorMessage string `json:"errorMessage,omitempty"`
	// The URI of the Pureport asset.
	Href string `json:"href,omitempty"`
	// The id is a unique identifier representing the gateway.
	Id string `json:"id,omitempty"`
	// The link state.
	LinkState string `json:"linkState,omitempty"`
	// The name.
	Name                      string `json:"name,omitempty"`
	PureportInternalGatewayIP string `json:"pureportInternalGatewayIP,omitempty"`
	// The ID of any resources attached to the gateway.
	RemoteId string `json:"remoteId,omitempty"`
	// The current state of the gateway.
	State string `json:"state,omitempty"`
	// The type.
	Type string `json:"type,omitempty"`
	// The version of the gateway container.
	Version           string        `json:"version,omitempty"`
	Auth              VpnAuthConfig `json:"auth,omitempty"`
	CustomerGatewayIP string        `json:"customerGatewayIP,omitempty"`
	CustomerVtiIP     string        `json:"customerVtiIP,omitempty"`
	OsNetworkId       string        `json:"osNetworkId,omitempty"`
	PureportGatewayIP string        `json:"pureportGatewayIP,omitempty"`
	PureportVtiIP     string        `json:"pureportVtiIP,omitempty"`
}
