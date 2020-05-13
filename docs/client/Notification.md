# Notification

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AvailabilityDomain** | **string** | The availability domain to be affected. | [optional] 
**ConnectionTypes** | **[]string** | A list of connection types to be affected. | [optional] 
**CreatedAt** | [**time.Time**](time.Time.md) |  | [optional] [readonly] 
**Description** | **string** | The description. | [optional] 
**End** | [**time.Time**](time.Time.md) |  | 
**Href** | **string** | The URI of the Pureport asset. | [optional] [readonly] 
**Id** | **string** | The id is a unique identifier representing the notification. | [optional] 
**ImpactedServices** | **[]string** | The list of Pureport services to be affected. | 
**Locations** | [**[]Link**](Link.md) | A list of asset links of Pureport locations to be affected. | [optional] 
**Start** | [**time.Time**](time.Time.md) |  | 
**State** | **string** | The current state. | [optional] 
**Type** | **string** | The type. | 
**UpdatedAt** | [**time.Time**](time.Time.md) |  | [optional] [readonly] 

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


