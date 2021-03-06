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

// VpnGatewayAllOf struct for VpnGatewayAllOf
type VpnGatewayAllOf struct {
	Auth              VpnAuthConfig `json:"auth,omitempty"`
	CustomerGatewayIP string        `json:"customerGatewayIP,omitempty"`
	CustomerVtiIP     string        `json:"customerVtiIP,omitempty"`
	IpsecStatus       IkeState      `json:"ipsecStatus,omitempty"`
	PureportGatewayIP string        `json:"pureportGatewayIP,omitempty"`
	PureportVtiIP     string        `json:"pureportVtiIP,omitempty"`
}
