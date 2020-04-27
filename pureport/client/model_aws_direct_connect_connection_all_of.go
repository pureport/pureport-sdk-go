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

// AwsDirectConnectConnectionAllOf struct for AwsDirectConnectConnectionAllOf
type AwsDirectConnectConnectionAllOf struct {
	AwsAccountId                 string                    `json:"awsAccountId,omitempty"`
	AwsRegion                    string                    `json:"awsRegion,omitempty"`
	BgpPasswordConfiguration     *BgpPasswordConfiguration `json:"bgpPasswordConfiguration,omitempty"`
	CloudRegion                  *Link                     `json:"cloudRegion,omitempty"`
	CloudServices                []Link                    `json:"cloudServices,omitempty"`
	CloudServicesPrefixWhitelist []string                  `json:"cloudServicesPrefixWhitelist,omitempty"`
	Peering                      PeeringConfiguration      `json:"peering,omitempty"`
}
