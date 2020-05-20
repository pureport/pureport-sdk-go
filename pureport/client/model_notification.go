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

import (
	"time"
)

// Notification An informational message for Pureport users.
type Notification struct {
	AvailabilityDomain AvailabilityDomain `json:"availabilityDomain,omitempty"`
	// A list of connection types to be affected.
	ConnectionTypes []ConnectionType `json:"connectionTypes,omitempty"`
	CreatedAt       time.Time        `json:"createdAt,omitempty"`
	// The description.
	Description string    `json:"description,omitempty"`
	End         time.Time `json:"end"`
	// The URI of the Pureport asset.
	Href string `json:"href,omitempty"`
	// The id is a unique identifier representing the notification.
	Id string `json:"id,omitempty"`
	// The list of Pureport services to be affected.
	ImpactedServices []NotificationImpactedService `json:"impactedServices"`
	// A list of asset links of Pureport locations to be affected.
	Locations []Link            `json:"locations,omitempty"`
	Start     time.Time         `json:"start"`
	State     NotificationState `json:"state,omitempty"`
	Type      NotificationType  `json:"type"`
	UpdatedAt time.Time         `json:"updatedAt,omitempty"`
}
