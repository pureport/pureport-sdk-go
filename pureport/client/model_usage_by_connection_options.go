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

// UsageByConnectionOptions struct for UsageByConnectionOptions
type UsageByConnectionOptions struct {
	Date                 DateFilter `json:"date,omitempty"`
	IncludeChildAccounts bool       `json:"includeChildAccounts,omitempty"`
	TimeUnit             string     `json:"timeUnit,omitempty"`
	TrafficType          string     `json:"trafficType,omitempty"`
}
