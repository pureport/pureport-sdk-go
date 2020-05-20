# ProviderLink

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AvailabilityDomain** | [**AvailabilityDomain**](AvailabilityDomain.md) |  | 
**Href** | **string** | The URI of the Pureport asset. | [optional] [readonly] 
**Id** | **string** | The id is a unique identifier representing the provider link. | 
**MaxAvailableBandwidth** | **int32** | The maximum available bandwidth of the link in Mbps. | 
**Pod** | [**Link**](Link.md) |  | 
**Provider** | [**ConnectionProvider**](ConnectionProvider.md) |  | 
**ProviderId** | **string** | The id of the provider side resource. | 
**ProviderRegion** | **string** | The region of the provider side resource. | [optional] 
**Speed** | **int32** | The speed of the link in Mbps. | 
**Status** | [**ProviderLinkStatus**](ProviderLinkStatus.md) |  | 
**SupportedConnectionTypes** | [**[]ConnectionType**](ConnectionType.md) | The types of connection this provider link will be used for. | 
**Vlans** | **string** | A range of vlans available on this link. | 

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


