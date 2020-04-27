# SiteIpSecVpnConnectionAllOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AuthType** | **string** |  | [optional] 
**BgpPasswordConfiguration** | [**BgpPasswordConfiguration**](BGPPasswordConfiguration.md) |  | [optional] 
**EnableBGPPassword** | **bool** |  | [optional] 
**IkeV1** | [**Ikev1Config**](IKEV1Config.md) |  | [optional] 
**IkeV2** | [**Ikev2Config**](IKEV2Config.md) |  | [optional] 
**IkeVersion** | **string** |  | [optional] 
**PhysicalAddress** | [**PhysicalAddress**](PhysicalAddress.md) |  | [optional] 
**PrimaryCustomerRouterIP** | **string** |  | [optional] 
**PrimaryKey** | **string** | IPsec pre-shared key override for the primary gateway | [optional] 
**RoutingType** | **string** |  | [optional] 
**SecondaryCustomerRouterIP** | **string** |  | [optional] 
**SecondaryKey** | **string** | IPsec pre-shared key override for the secondary gateway | [optional] 
**TrafficSelectors** | [**[]TrafficSelectorMapping**](TrafficSelectorMapping.md) |  | [optional] 

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


