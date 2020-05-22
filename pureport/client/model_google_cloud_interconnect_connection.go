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

// GoogleCloudInterconnectConnection Connection using Google Cloud Interconnect.
type GoogleCloudInterconnectConnection struct {
	ActiveAt time.Time `json:"activeAt,omitempty"`
	// If the connection is advertising internal routes, which allows the customer the option of probing and tracing these routes.
	AdvertiseInternalRoutes bool            `json:"advertiseInternalRoutes,omitempty"`
	BillingPlan             *BillingPlan    `json:"billingPlan,omitempty"`
	BillingProvider         BillingProvider `json:"billingProvider,omitempty"`
	BillingTerm             BillingTerm     `json:"billingTerm"`
	CreatedAt               time.Time       `json:"createdAt,omitempty"`
	// The customer side ASN. This can either be a public or private ASN. If this is a public ASN, you must own it to prevent conflicts.
	CustomerASN int64 `json:"customerASN,omitempty"`
	// Set of customer Networks for this connection.
	CustomerNetworks []CustomerNetwork `json:"customerNetworks,omitempty"`
	DeletedAt        time.Time         `json:"deletedAt,omitempty"`
	// The user defined description for the connection.
	Description string    `json:"description,omitempty"`
	ErrorCode   ErrorCode `json:"errorCode,omitempty"`
	// Error message assigned to the connection if it is an error state.
	ErrorMessage string `json:"errorMessage,omitempty"`
	// Whether this connection has redundant gateways for failover.
	HighAvailability bool `json:"highAvailability"`
	// The URI of the Pureport asset.
	Href string `json:"href,omitempty"`
	// The id is a unique identifier representing the connection. This can be provided during creation, but if left empty, will be generated.
	Id       string `json:"id,omitempty"`
	Location Link   `json:"location"`
	// The user specified name for the connection.
	Name             string           `json:"name"`
	Nat              *NatConfig       `json:"nat,omitempty"`
	Network          Link             `json:"network,omitempty"`
	PrimaryGateway   *StandardGateway `json:"primaryGateway,omitempty"`
	SecondaryGateway *StandardGateway `json:"secondaryGateway,omitempty"`
	Speed            ConnectionSpeed  `json:"speed"`
	State            ConnectionState  `json:"state,omitempty"`
	// Key-value pairs to associate with the Pureport asset.
	Tags map[string]string `json:"tags,omitempty"`
	Type ConnectionType    `json:"type"`
	// The Pairing Key for the primary gateway of the Google Cloud Interconnect.
	PrimaryPairingKey string `json:"primaryPairingKey"`
	// The Pairing Key for the secondary gateway of the Google Cloud Interconnect.
	SecondaryPairingKey string `json:"secondaryPairingKey,omitempty"`
}
