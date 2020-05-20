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

// Sort struct for Sort
type Sort struct {
	Direction SortDirection `json:"direction,omitempty"`
	Property  string        `json:"property,omitempty"`
}
