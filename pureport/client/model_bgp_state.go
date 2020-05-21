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

// BgpState The current BGP state.
type BgpState string

// List of BGPState
const (
	BGP_STATE__UP      BgpState = "UP"
	BGP_STATE__DOWN    BgpState = "DOWN"
	BGP_STATE__PENDING BgpState = "PENDING"
)
