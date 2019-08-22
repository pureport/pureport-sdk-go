/*
 * Pureport Control Plane
 *
 * Pureport API
 *
 * API version: 1.0.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package client

type VpnGateway struct {
	AvailabilityDomain string         `json:"availabilityDomain,omitempty"`
	BgpConfig          *BgpConfig     `json:"bgpConfig,omitempty"`
	Description        string         `json:"description,omitempty"`
	ErrorCode          string         `json:"errorCode,omitempty"`
	ErrorMessage       string         `json:"errorMessage,omitempty"`
	Id                 string         `json:"id,omitempty"`
	LinkState          string         `json:"linkState,omitempty"`
	Name               string         `json:"name,omitempty"`
	RemoteId           string         `json:"remoteId,omitempty"`
	State              string         `json:"state,omitempty"`
	Type_              string         `json:"type,omitempty"`
	Auth               *PskAuthConfig `json:"auth,omitempty"`
	CustomerGatewayIP  string         `json:"customerGatewayIP,omitempty"`
	CustomerVtiIP      string         `json:"customerVtiIP,omitempty"`
	IpsecStatus        string         `json:"ipsecStatus,omitempty"`
	PureportGatewayIP  string         `json:"pureportGatewayIP,omitempty"`
	PureportVtiIP      string         `json:"pureportVtiIP,omitempty"`
}
