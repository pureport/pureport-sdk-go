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

// SiteIpSecVpnConnection struct for SiteIpSecVpnConnection
type SiteIpSecVpnConnection struct {
	ActiveAt                  time.Time                 `json:"activeAt,omitempty"`
	AdvertiseInternalRoutes   bool                      `json:"advertiseInternalRoutes,omitempty"`
	BillingPlan               *BillingPlan              `json:"billingPlan,omitempty"`
	BillingProvider           string                    `json:"billingProvider,omitempty"`
	BillingTerm               string                    `json:"billingTerm"`
	CreatedAt                 time.Time                 `json:"createdAt,omitempty"`
	CustomerASN               int64                     `json:"customerASN,omitempty"`
	CustomerNetworks          []CustomerNetwork         `json:"customerNetworks,omitempty"`
	DeletedAt                 time.Time                 `json:"deletedAt,omitempty"`
	Description               string                    `json:"description,omitempty"`
	ErrorCode                 string                    `json:"errorCode,omitempty"`
	ErrorMessage              string                    `json:"errorMessage,omitempty"`
	HighAvailability          bool                      `json:"highAvailability"`
	Href                      string                    `json:"href,omitempty"`
	Id                        string                    `json:"id,omitempty"`
	Location                  Link                      `json:"location"`
	Name                      string                    `json:"name"`
	Nat                       *NatConfig                `json:"nat,omitempty"`
	Network                   Link                      `json:"network,omitempty"`
	PrimaryGateway            *VpnGateway               `json:"primaryGateway,omitempty"`
	SecondaryGateway          *VpnGateway               `json:"secondaryGateway,omitempty"`
	Speed                     int32                     `json:"speed"`
	State                     string                    `json:"state,omitempty"`
	Tags                      map[string]string         `json:"tags,omitempty"`
	Type                      string                    `json:"type"`
	AuthType                  string                    `json:"authType"`
	BgpPasswordConfiguration  *BgpPasswordConfiguration `json:"bgpPasswordConfiguration,omitempty"`
	EnableBGPPassword         bool                      `json:"enableBGPPassword,omitempty"`
	IkeV1                     *Ikev1Config              `json:"ikeV1,omitempty"`
	IkeV2                     *Ikev2Config              `json:"ikeV2,omitempty"`
	IkeVersion                string                    `json:"ikeVersion"`
	PhysicalAddress           PhysicalAddress           `json:"physicalAddress,omitempty"`
	PrimaryCustomerRouterIP   string                    `json:"primaryCustomerRouterIP"`
	PrimaryKey                string                    `json:"primaryKey,omitempty"`
	RoutingType               string                    `json:"routingType"`
	SecondaryCustomerRouterIP string                    `json:"secondaryCustomerRouterIP,omitempty"`
	SecondaryKey              string                    `json:"secondaryKey,omitempty"`
	TrafficSelectors          []TrafficSelectorMapping  `json:"trafficSelectors,omitempty"`
}
