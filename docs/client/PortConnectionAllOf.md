# PortConnectionAllOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**BgpPasswordConfiguration** | [**BgpPasswordConfiguration**](BGPPasswordConfiguration.md) |  | [optional] 
**BillingLongHaul** | [**BillingLongHaul**](BillingLongHaul.md) |  | [optional] 
**GatewayCidr** | **string** |  | [optional] 
**PrimaryCustomerVlan** | **int32** | The VLAN ID of the primary gateway. | [optional] 
**PrimaryGatewayIP** | **string** |  | [optional] 
**PrimaryPort** | [**Link**](Link.md) |  | [optional] 
**RoutingType** | [**RoutingType**](RoutingType.md) |  | [optional] 
**SecondaryCustomerVlan** | **int32** | The VLAN ID of the secondary gateway if this is an HA connection. | [optional] 
**SecondaryGatewayIP** | **string** |  | [optional] 
**SecondaryPort** | [**Link**](Link.md) |  | [optional] 
**StaticRoutes** | [**[]StaticRoute**](StaticRoute.md) | The user configured static routes. | [optional] 
**VirtualGatewayIP** | **string** |  | [optional] 

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


