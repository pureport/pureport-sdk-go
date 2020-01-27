# AzureExpressRouteConnection

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ActiveAt** | [**time.Time**](time.Time.md) |  | [optional] [readonly] 
**BillingPlan** | [**BillingPlan**](BillingPlan.md) |  | [optional] 
**BillingTerm** | **string** |  | 
**CreatedAt** | [**time.Time**](time.Time.md) |  | [optional] [readonly] 
**CustomerASN** | **int64** |  | [optional] 
**CustomerNetworks** | [**[]CustomerNetwork**](CustomerNetwork.md) |  | [optional] 
**DeletedAt** | [**time.Time**](time.Time.md) |  | [optional] [readonly] 
**Description** | **string** |  | [optional] 
**ErrorCode** | **string** |  | [optional] [readonly] 
**ErrorMessage** | **string** |  | [optional] [readonly] 
**HighAvailability** | **bool** |  | [optional] 
**Href** | **string** |  | [optional] 
**Id** | **string** |  | [optional] 
**Location** | [**Link**](Link.md) |  | 
**Name** | **string** |  | 
**Nat** | [**NatConfig**](NATConfig.md) |  | [optional] 
**Network** | [**Link**](Link.md) |  | [optional] 
**PrimaryGateway** | [**Gateway**](Gateway.md) |  | [optional] 
**SecondaryGateway** | [**Gateway**](Gateway.md) |  | [optional] 
**Speed** | **int32** |  | 
**State** | **string** |  | [optional] [readonly] 
**Tags** | **map[string]string** |  | [optional] 
**Type** | **string** |  | 
**Peering** | [**PeeringConfiguration**](PeeringConfiguration.md) |  | [optional] 
**ServiceKey** | **string** |  | 

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


