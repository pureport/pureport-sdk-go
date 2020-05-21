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

// ConnectionType The connection type of this configuration.
type ConnectionType string

// List of ConnectionType
const (
	CONNECTION_TYPE__AWS_DIRECT_CONNECT        ConnectionType = "AWS_DIRECT_CONNECT"
	CONNECTION_TYPE__AZURE_EXPRESS_ROUTE       ConnectionType = "AZURE_EXPRESS_ROUTE"
	CONNECTION_TYPE__EQUINIX_CLOUD_EXCHANGE    ConnectionType = "EQUINIX_CLOUD_EXCHANGE"
	CONNECTION_TYPE__GOOGLE_CLOUD_INTERCONNECT ConnectionType = "GOOGLE_CLOUD_INTERCONNECT"
	CONNECTION_TYPE__ORACLE_FAST_CONNECT       ConnectionType = "ORACLE_FAST_CONNECT"
	CONNECTION_TYPE__SITE_IPSEC_VPN            ConnectionType = "SITE_IPSEC_VPN"
	CONNECTION_TYPE__PORT                      ConnectionType = "PORT"
	CONNECTION_TYPE__PACKET_FABRIC             ConnectionType = "PACKET_FABRIC"
	CONNECTION_TYPE__COMCAST                   ConnectionType = "COMCAST"
	CONNECTION_TYPE__COX_COMMUNICATIONS        ConnectionType = "COX_COMMUNICATIONS"
	CONNECTION_TYPE__ELEMENT_CRITICAL          ConnectionType = "ELEMENT_CRITICAL"
	CONNECTION_TYPE__PACKET                    ConnectionType = "PACKET"
	CONNECTION_TYPE__ZAYO                      ConnectionType = "ZAYO"
)