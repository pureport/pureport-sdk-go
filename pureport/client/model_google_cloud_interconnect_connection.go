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

// GoogleCloudInterconnectConnection struct for GoogleCloudInterconnectConnection
type GoogleCloudInterconnectConnection struct {
	ActiveAt            time.Time         `json:"activeAt,omitempty"`
	BillingPlan         BillingPlan       `json:"billingPlan,omitempty"`
	BillingTerm         string            `json:"billingTerm"`
	CreatedAt           time.Time         `json:"createdAt,omitempty"`
	CustomerASN         int64             `json:"customerASN,omitempty"`
	CustomerNetworks    []CustomerNetwork `json:"customerNetworks,omitempty"`
	DeletedAt           time.Time         `json:"deletedAt,omitempty"`
	Description         string            `json:"description,omitempty"`
	ErrorCode           string            `json:"errorCode,omitempty"`
	ErrorMessage        string            `json:"errorMessage,omitempty"`
	HighAvailability    bool              `json:"highAvailability,omitempty"`
	Href                string            `json:"href,omitempty"`
	Id                  string            `json:"id,omitempty"`
	Location            Link              `json:"location"`
	Name                string            `json:"name"`
	Nat                 *NatConfig        `json:"nat,omitempty"`
	Network             Link              `json:"network,omitempty"`
	PrimaryGateway      *StandardGateway  `json:"primaryGateway,omitempty"`
	SecondaryGateway    *StandardGateway  `json:"secondaryGateway,omitempty"`
	Speed               int32             `json:"speed"`
	State               string            `json:"state,omitempty"`
	Tags                map[string]string `json:"tags,omitempty"`
	Type                string            `json:"type"`
	PrimaryPairingKey   string            `json:"primaryPairingKey"`
	SecondaryPairingKey string            `json:"secondaryPairingKey,omitempty"`
}
