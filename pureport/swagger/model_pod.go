/*
 * Pureport Control Plane
 *
 * Pureport API
 *
 * API version: 1.0.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package swagger

type Pod struct {
	Href     string `json:"href,omitempty"`
	Id       string `json:"id"`
	Location *Link  `json:"location,omitempty"`
	Status   string `json:"status,omitempty"`
}
