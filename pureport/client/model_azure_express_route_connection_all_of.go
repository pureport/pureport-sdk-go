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

// AzureExpressRouteConnectionAllOf struct for AzureExpressRouteConnectionAllOf
type AzureExpressRouteConnectionAllOf struct {
	BgpPasswordConfiguration BgpPasswordConfiguration `json:"bgpPasswordConfiguration,omitempty"`
	Peering                  PeeringConfiguration     `json:"peering,omitempty"`
	// The Azure ExpressRoute service key, a public and unique standard GUID representing identifying the ExpressRoute circuit.
	ServiceKey string `json:"serviceKey,omitempty"`
}
