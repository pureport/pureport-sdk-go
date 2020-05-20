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

// VpnRoutingType The VPN Routing Type.
type VpnRoutingType string

// List of VPNRoutingType
const (
	POLICY_BASED       VpnRoutingType = "POLICY_BASED"
	ROUTE_BASED_STATIC VpnRoutingType = "ROUTE_BASED_STATIC"
	ROUTE_BASED_BGP    VpnRoutingType = "ROUTE_BASED_BGP"
)
