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

// TrafficSelectorMapping A pair of local and remote addresses for use in VPN Traffic Selectors.
type TrafficSelectorMapping struct {
	CustomerSide string `json:"customerSide"`
	PureportSide string `json:"pureportSide"`
}
