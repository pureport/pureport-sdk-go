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

// ConnectionState The current state of the connection.
type ConnectionState string

// List of ConnectionState
const (
	INITIALIZING         ConnectionState = "INITIALIZING"
	WAITING_TO_PROVISION ConnectionState = "WAITING_TO_PROVISION"
	PENDING_APPROVAL     ConnectionState = "PENDING_APPROVAL"
	APPROVED             ConnectionState = "APPROVED"
	PROVISIONING         ConnectionState = "PROVISIONING"
	FAILED_TO_PROVISION  ConnectionState = "FAILED_TO_PROVISION"
	ACTIVE               ConnectionState = "ACTIVE"
	DOWN                 ConnectionState = "DOWN"
	UPDATING             ConnectionState = "UPDATING"
	FAILED_TO_UPDATE     ConnectionState = "FAILED_TO_UPDATE"
	DELETING             ConnectionState = "DELETING"
	FAILED_TO_DELETE     ConnectionState = "FAILED_TO_DELETE"
	DELETED              ConnectionState = "DELETED"
)
