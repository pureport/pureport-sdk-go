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

// ConnectionProvider The provider of connections using this link.
type ConnectionProvider string

// List of ConnectionProvider
const (
	CONNECTION_PROVIDER__AWS                    ConnectionProvider = "AWS"
	CONNECTION_PROVIDER__GOOGLE_CLOUD           ConnectionProvider = "GOOGLE_CLOUD"
	CONNECTION_PROVIDER__EQUINIX_CLOUD_EXCHANGE ConnectionProvider = "EQUINIX_CLOUD_EXCHANGE"
	CONNECTION_PROVIDER__PUREPORT               ConnectionProvider = "PUREPORT"
	CONNECTION_PROVIDER__PACKET_FABRIC          ConnectionProvider = "PACKET_FABRIC"
	CONNECTION_PROVIDER__COMCAST                ConnectionProvider = "COMCAST"
	CONNECTION_PROVIDER__COX_COMMUNICATIONS     ConnectionProvider = "COX_COMMUNICATIONS"
	CONNECTION_PROVIDER__ELEMENT_CRITICAL       ConnectionProvider = "ELEMENT_CRITICAL"
	CONNECTION_PROVIDER__PACKET                 ConnectionProvider = "PACKET"
	CONNECTION_PROVIDER__ZAYO                   ConnectionProvider = "ZAYO"
)