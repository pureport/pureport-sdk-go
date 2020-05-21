# IpRoute

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DirectlyConnected** | **bool** | Whether this route is directly connected. | [optional] [readonly] 
**Distance** | **int64** | The route distance. | [optional] [readonly] 
**Fib** | **bool** | Whether this route is a part of the FIB (Forwarding) table. | [optional] [readonly] 
**InterfaceName** | **string** | The interface name. | [optional] [readonly] 
**Metric** | **int64** | The metric for the route. | [optional] [readonly] 
**NextHop** | **string** |  | [optional] [readonly] 
**NextHopConnection** | [**Link**](Link.md) |  | [optional] 
**NextHopGateway** | [**Link**](Link.md) |  | [optional] 
**Prefix** | **string** |  | [optional] [readonly] 
**Protocol** | **string** | The protocol for the route. | [optional] [readonly] 
**Selected** | **bool** | Whether this route is marked as selected. | [optional] [readonly] 
**Uptime** | **string** | The uptime for this route in hours:minutes:seconds. | [optional] [readonly] 

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


