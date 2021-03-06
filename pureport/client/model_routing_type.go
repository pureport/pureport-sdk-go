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

// RoutingType The method to use for determining network routes.
type RoutingType string

// List of RoutingType
const (
	ROUTING_TYPE__STATIC RoutingType = "STATIC"
	ROUTING_TYPE__BGP    RoutingType = "BGP"
)
