/*
 * Purport Control Plane
 *
 * Pureport API
 *
 * API version: 1.0.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package swagger

type SiteIpSecVpnConnection struct {
	ActiveAt                  int64                    `json:"activeAt,omitempty"`
	BillingPlan               *BillingPlan             `json:"billingPlan,omitempty"`
	BillingTerm               string                   `json:"billingTerm,omitempty"`
	CreatedAt                 int64                    `json:"createdAt,omitempty"`
	CustomerASN               int64                    `json:"customerASN,omitempty"`
	CustomerNetworks          []CustomerNetwork        `json:"customerNetworks,omitempty"`
	DeletedAt                 int64                    `json:"deletedAt,omitempty"`
	Description               string                   `json:"description,omitempty"`
	ErrorCode                 string                   `json:"errorCode,omitempty"`
	ErrorMessage              string                   `json:"errorMessage,omitempty"`
	HighAvailability          bool                     `json:"highAvailability,omitempty"`
	Href                      string                   `json:"href,omitempty"`
	Id                        string                   `json:"id,omitempty"`
	Location                  *Link                    `json:"location"`
	Name                      string                   `json:"name"`
	Nat                       *NatConfig               `json:"nat,omitempty"`
	Network                   *Link                    `json:"network,omitempty"`
	PrimaryGateway            *Gateway                 `json:"primaryGateway,omitempty"`
	SecondaryGateway          *Gateway                 `json:"secondaryGateway,omitempty"`
	Speed                     int32                    `json:"speed"`
	State                     string                   `json:"state,omitempty"`
	Type_                     string                   `json:"type"`
	AuthType                  string                   `json:"authType"`
	EnableBGPPassword         bool                     `json:"enableBGPPassword,omitempty"`
	IkeV1                     *Ikev1Config             `json:"ikeV1,omitempty"`
	IkeV2                     *Ikev2Config             `json:"ikeV2,omitempty"`
	IkeVersion                string                   `json:"ikeVersion"`
	PrimaryCustomerRouterIP   string                   `json:"primaryCustomerRouterIP"`
	PrimaryKey                string                   `json:"primaryKey,omitempty"`
	RoutingType               string                   `json:"routingType"`
	SecondaryCustomerRouterIP string                   `json:"secondaryCustomerRouterIP,omitempty"`
	SecondaryKey              string                   `json:"secondaryKey,omitempty"`
	TrafficSelectors          []TrafficSelectorMapping `json:"trafficSelectors,omitempty"`
}
