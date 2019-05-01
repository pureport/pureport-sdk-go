/*
 * Pureport Control Plane
 *
 * Pureport API
 *
 * API version: 1.0.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package client

type ResourceLimit struct {
	ChildResourceType  string `json:"childResourceType"`
	Href               string `json:"href,omitempty"`
	Limit              int32  `json:"limit"`
	ParentResourceId   string `json:"parentResourceId,omitempty"`
	ParentResourceType string `json:"parentResourceType,omitempty"`
}
