/*
 * Purport Control Plane
 *
 * Pureport API
 *
 * API version: 1.0.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package swagger

type Location struct {
	GeoCoordinates *GeoCoordinates          `json:"geoCoordinates"`
	Href           string                   `json:"href,omitempty"`
	Id             string                   `json:"id,omitempty"`
	LocationLinks  []LocationLinkConnection `json:"locationLinks,omitempty"`
	Name           string                   `json:"name"`
}
