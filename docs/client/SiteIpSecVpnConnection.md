# SiteIpSecVpnConnection

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ActiveAt** | [**time.Time**](time.Time.md) |  | [optional] [readonly] 
**AdvertiseInternalRoutes** | **bool** | If the connection is advertising internal routes, which allows the customer the option of probing and tracing these routes. | [optional] 
**BillingPlan** | [**BillingPlan**](BillingPlan.md) |  | [optional] 
**BillingProvider** | **string** | The provider used for billing this connection. | [optional] [readonly] 
**BillingTerm** | **string** | The licensed billing term for the connection. | 
**CreatedAt** | [**time.Time**](time.Time.md) |  | [optional] [readonly] 
**CustomerASN** | **int64** | The customer side ASN. This can either be a public or private ASN. If this is a public ASN, you must own it to prevent conflicts. | [optional] 
**CustomerNetworks** | [**[]CustomerNetwork**](CustomerNetwork.md) | Set of customer Networks for this connection. | [optional] 
**DeletedAt** | [**time.Time**](time.Time.md) |  | [optional] [readonly] 
**Description** | **string** | The user defined description for the connection. | [optional] 
**ErrorCode** | **string** | Error Code assigned to the connection if it is an error state. | [optional] [readonly] 
**ErrorMessage** | **string** | Error message assigned to the connection if it is an error state. | [optional] [readonly] 
**HighAvailability** | **bool** | Whether this connection has redundant gateways for failover. | 
**Href** | **string** | The URI of the Pureport asset. | [optional] [readonly] 
**Id** | **string** | The id is a unique identifier representing the connection. This can be provided during creation, but if left empty, will be generated. | [optional] 
**Location** | [**Link**](Link.md) |  | 
**Name** | **string** | The user specified name for the connection. | 
**Nat** | [**NatConfig**](NATConfig.md) |  | [optional] 
**Network** | [**Link**](Link.md) |  | [optional] 
**PrimaryGateway** | [**Gateway**](Gateway.md) |  | [optional] 
**SecondaryGateway** | [**Gateway**](Gateway.md) |  | [optional] 
**Speed** | **int32** | The connection speed in Mbps. | 
**State** | **string** | The current state of the connection. | [optional] [readonly] 
**Tags** | **map[string]string** | Key-value pairs to associate with the Pureport asset. | [optional] 
**Type** | **string** | The connection type. | 
**AuthType** | **string** | The authentication type. | 
**BgpPasswordConfiguration** | [**BgpPasswordConfiguration**](BGPPasswordConfiguration.md) |  | [optional] 
**EnableBGPPassword** | **bool** |  | [optional] 
**IkeV1** | [**Ikev1Config**](IKEV1Config.md) |  | [optional] 
**IkeV2** | [**Ikev2Config**](IKEV2Config.md) |  | [optional] 
**IkeVersion** | **string** | The IKE version. | 
**PhysicalAddress** | [**PhysicalAddress**](PhysicalAddress.md) |  | [optional] 
**PrimaryCustomerRouterIP** | **string** |  | 
**PrimaryKey** | **string** | IPsec pre-shared key (PSK) override for the primary gateway. | [optional] 
**RoutingType** | **string** | The VPN Routing Type. | 
**SecondaryCustomerRouterIP** | **string** |  | [optional] 
**SecondaryKey** | **string** | IPsec pre-shared key (PSK) override for the secondary gateway. | [optional] 
**TrafficSelectors** | [**[]TrafficSelectorMapping**](TrafficSelectorMapping.md) | The traffic selectors to apply for routing. | [optional] 

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


