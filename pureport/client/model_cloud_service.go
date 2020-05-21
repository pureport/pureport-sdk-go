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

import (
	"time"
)

// CloudService A cloud-provided service networking configuration
type CloudService struct {
	CloudRegion Link `json:"cloudRegion"`
	// Whether this cloud service configuration has been deactivated.
	Deactivated   bool      `json:"deactivated"`
	DeactivatedAt time.Time `json:"deactivatedAt,omitempty"`
	// The URI of the Pureport asset.
	Href string `json:"href,omitempty"`
	// The id is a unique identifier representing the cloud service.
	Id string `json:"id,omitempty"`
	// The network CIDRs associated with this cloud service configuration.
	IpPrefixes []string `json:"ipPrefixes,omitempty"`
	// The count of IPv4 CIDRs associated with this cloud service configuration.
	Ipv4PrefixCount int32 `json:"ipv4PrefixCount,omitempty"`
	// The count of IPv6 CIDRs associated with this cloud service configuration.
	Ipv6PrefixCount int32 `json:"ipv6PrefixCount,omitempty"`
	// The name of the cloud service configuration.
	Name     string        `json:"name"`
	Provider CloudProvider `json:"provider"`
	// The service name.
	Service string `json:"service"`
}
