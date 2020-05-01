# PortConnectionAllOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**BgpPasswordConfiguration** | [**BgpPasswordConfiguration**](BGPPasswordConfiguration.md) |  | [optional] 
**BillingLongHaul** | **string** | The long haul billing configuration for this connection. | [optional] [readonly] 
**GatewayCidr** | **string** |  | [optional] 
**PrimaryCustomerVlan** | **int32** | The VLAN ID of the primary gateway. | [optional] 
**PrimaryGatewayIP** | **string** |  | [optional] 
**PrimaryPort** | [**Link**](Link.md) |  | [optional] 
**RoutingType** | **string** | The method to use for determining network routes. | [optional] 
**SecondaryCustomerVlan** | **int32** | The VLAN ID of the secondary gateway if this is an HA connection. | [optional] 
**SecondaryGatewayIP** | **string** |  | [optional] 
**SecondaryPort** | [**Link**](Link.md) |  | [optional] 
**StaticRoutes** | [**[]StaticRoute**](StaticRoute.md) | The user configured static routes. | [optional] 
**VirtualGatewayIP** | **string** |  | [optional] 

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


