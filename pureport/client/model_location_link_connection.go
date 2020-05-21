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

// LocationLinkConnection Link between Pureport Locations across the Pureport backbone.
type LocationLinkConnection struct {
	Location Link            `json:"location"`
	Speed    ConnectionSpeed `json:"speed"`
}
