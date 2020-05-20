# SiteIpSecVpnConnectionAllOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AuthType** | [**IpSecAuthenticationType**](IPSecAuthenticationType.md) |  | [optional] 
**BgpPasswordConfiguration** | [**BgpPasswordConfiguration**](BGPPasswordConfiguration.md) |  | [optional] 
**EnableBGPPassword** | **bool** |  | [optional] 
**IkeV1** | [**Ikev1Config**](IKEV1Config.md) |  | [optional] 
**IkeV2** | [**Ikev2Config**](IKEV2Config.md) |  | [optional] 
**IkeVersion** | [**IkeVersion**](IKEVersion.md) |  | [optional] 
**PhysicalAddress** | [**PhysicalAddress**](PhysicalAddress.md) |  | [optional] 
**PrimaryCustomerRouterIP** | **string** |  | [optional] 
**PrimaryKey** | **string** | IPsec pre-shared key (PSK) override for the primary gateway. | [optional] 
**RoutingType** | [**VpnRoutingType**](VPNRoutingType.md) |  | [optional] 
**SecondaryCustomerRouterIP** | **string** |  | [optional] 
**SecondaryKey** | **string** | IPsec pre-shared key (PSK) override for the secondary gateway. | [optional] 
**TrafficSelectors** | [**[]TrafficSelectorMapping**](TrafficSelectorMapping.md) | The traffic selectors to apply for routing. | [optional] 

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


