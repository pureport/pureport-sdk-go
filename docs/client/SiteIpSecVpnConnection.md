# SiteIpSecVpnConnection

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ActiveAt** | [**time.Time**](time.Time.md) |  | [optional] 
**BillingPlan** | [***BillingPlan**](BillingPlan.md) |  | [optional] 
**BillingTerm** | **string** |  | 
**CreatedAt** | [**time.Time**](time.Time.md) |  | [optional] 
**CustomerASN** | **int64** |  | [optional] 
**CustomerNetworks** | [**[]CustomerNetwork**](CustomerNetwork.md) |  | [optional] 
**DeletedAt** | [**time.Time**](time.Time.md) |  | [optional] 
**Description** | **string** |  | [optional] 
**ErrorCode** | **string** |  | [optional] 
**ErrorMessage** | **string** |  | [optional] 
**HighAvailability** | **bool** |  | [optional] 
**Href** | **string** |  | [optional] 
**Id** | **string** |  | [optional] 
**Location** | [***Link**](Link.md) |  | 
**Name** | **string** |  | 
**Nat** | [***NatConfig**](NATConfig.md) |  | [optional] 
**Network** | [***Link**](Link.md) |  | [optional] 
**PrimaryGateway** | [***Gateway**](Gateway.md) |  | [optional] 
**SecondaryGateway** | [***Gateway**](Gateway.md) |  | [optional] 
**Speed** | **int32** |  | 
**State** | **string** |  | [optional] 
**Type_** | **string** |  | 
**AuthType** | **string** |  | 
**EnableBGPPassword** | **bool** |  | [optional] 
**IkeV1** | [***Ikev1Config**](IKEV1Config.md) |  | [optional] 
**IkeV2** | [***Ikev2Config**](IKEV2Config.md) |  | [optional] 
**IkeVersion** | **string** |  | 
**PhysicalAddress** | [***PhysicalAddress**](PhysicalAddress.md) |  | [optional] 
**PrimaryCustomerRouterIP** | **string** |  | 
**PrimaryKey** | **string** | IPsec pre-shared key override for the primary gateway | [optional] 
**RoutingType** | **string** |  | 
**SecondaryCustomerRouterIP** | **string** |  | [optional] 
**SecondaryKey** | **string** | IPsec pre-shared key override for the secondary gateway | [optional] 
**TrafficSelectors** | [**[]TrafficSelectorMapping**](TrafficSelectorMapping.md) |  | [optional] 

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


