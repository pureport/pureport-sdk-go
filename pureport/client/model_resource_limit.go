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

// ResourceLimit A representation of limitations between two Pureport resources.
type ResourceLimit struct {
	ChildResourceType ResourceType `json:"childResourceType"`
	// The URI of the Pureport asset.
	Href string `json:"href,omitempty"`
	// The number of child resources to limit to for the specified parent.
	Limit int32 `json:"limit"`
	// The unique identifier of the parent resource.
	ParentResourceId   string       `json:"parentResourceId,omitempty"`
	ParentResourceType ResourceType `json:"parentResourceType,omitempty"`
}
