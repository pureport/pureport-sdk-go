# GenericConnectionAllOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**BgpPasswordConfiguration** | [**BgpPasswordConfiguration**](BGPPasswordConfiguration.md) |  | [optional] 
**GatewayCidr** | **string** |  | [optional] 
**Peering** | [**PeeringConfiguration**](PeeringConfiguration.md) |  | [optional] 
**PrimaryGatewayIP** | **string** |  | [optional] 
**PrimaryVlan** | **int32** | The primary VLAN ID. | [optional] 
**RoutingType** | [**RoutingType**](RoutingType.md) |  | [optional] 
**SecondaryGatewayIP** | **string** |  | [optional] 
**SecondaryVlan** | **int32** | The secondary VLAN ID if this is an HA connection. | [optional] 
**StaticRoutes** | [**[]StaticRoute**](StaticRoute.md) | The user configured static routes. | [optional] 
**VirtualGatewayIP** | **string** |  | [optional] 

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


