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

// NotificationState The current state.
type NotificationState string

// List of NotificationState
const (
	NOTIFICATION_STATE__DRAFT     NotificationState = "DRAFT"
	NOTIFICATION_STATE__PUBLISHED NotificationState = "PUBLISHED"
	NOTIFICATION_STATE__CANCELED  NotificationState = "CANCELED"
	NOTIFICATION_STATE__DELETED   NotificationState = "DELETED"
)
