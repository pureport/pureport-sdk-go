/*
 * Pureport Control Plane
 *
 * Pureport API
 *
 * API version: 1.0.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package swagger

type Ikev1IkeConfig struct {
	DhGroup    string `json:"dhGroup"`
	Encryption string `json:"encryption"`
	Integrity  string `json:"integrity"`
}
