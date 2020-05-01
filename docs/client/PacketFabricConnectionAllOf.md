# PacketFabricConnectionAllOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**BgpPasswordConfiguration** | [**BgpPasswordConfiguration**](BGPPasswordConfiguration.md) |  | [optional] 
**CloudRegion** | [**Link**](Link.md) |  | [optional] 
**Customer** | [**PacketFabricCustomer**](PacketFabricCustomer.md) |  | [optional] 
**CustomerPrimaryId** | **string** | An ID representing the primary connection with the 3rd party provider selected in the “customer” field. | [optional] 
**CustomerSecondaryId** | **string** | An ID representing the secondary connection with the 3rd party provider selected in the “customer” field. | [optional] 
**Peering** | [**PeeringConfiguration**](PeeringConfiguration.md) |  | [optional] 
**PrimaryVlan** | **int32** | The primary VLAN ID. | [optional] 
**SecondaryVlan** | **int32** | The secondary VLAN ID if this is an HA connection. | [optional] 

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


