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

// StandardGateway struct for StandardGateway
type StandardGateway struct {
	AvailabilityDomain        string    `json:"availabilityDomain,omitempty"`
	BgpConfig                 BgpConfig `json:"bgpConfig,omitempty"`
	Connection                Link      `json:"connection,omitempty"`
	Description               string    `json:"description,omitempty"`
	ErrorCode                 string    `json:"errorCode,omitempty"`
	ErrorMessage              string    `json:"errorMessage,omitempty"`
	Href                      string    `json:"href,omitempty"`
	Id                        string    `json:"id,omitempty"`
	LinkState                 string    `json:"linkState,omitempty"`
	Name                      string    `json:"name,omitempty"`
	PureportInternalGatewayIP string    `json:"pureportInternalGatewayIP,omitempty"`
	RemoteId                  string    `json:"remoteId,omitempty"`
	State                     string    `json:"state,omitempty"`
	Type                      string    `json:"type,omitempty"`
	Version                   string    `json:"version,omitempty"`
	Vlan                      int32     `json:"vlan,omitempty"`
}
