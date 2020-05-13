# BgpRoute

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Best** | **bool** | Whether the route has been designated as the best route. | [optional] [readonly] 
**Internal** | **bool** | Whether the route path origin is internal. | [optional] [readonly] 
**LocPref** | **int32** | The local preference assigned to the route. | [optional] [readonly] 
**Metric** | **int32** | The metric for the route. | [optional] [readonly] 
**MultiPath** | **bool** | Whether the route is a multipath route. | [optional] [readonly] 
**Network** | **string** |  | [optional] [readonly] 
**NextHop** | **string** |  | [optional] [readonly] 
**NextHopConnection** | [**Link**](Link.md) |  | [optional] 
**NextHopGateway** | [**Link**](Link.md) |  | [optional] 
**Origin** | **string** | The BGP origin code for the route. | [optional] [readonly] 
**Path** | **[]int32** | List of AS paths for this route. | [optional] [readonly] 
**Valid** | **bool** | Whether the route is valid. | [optional] [readonly] 
**Weight** | **int32** | The weight assigned to the route. | [optional] [readonly] 

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


