# CloudService

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CloudRegion** | [**Link**](Link.md) |  | 
**Deactivated** | **bool** | Whether this cloud service configuration has been deactivated. | [readonly] 
**DeactivatedAt** | [**time.Time**](time.Time.md) |  | [optional] [readonly] 
**Href** | **string** | The URI of the Pureport asset. | [optional] [readonly] 
**Id** | **string** | The id is a unique identifier representing the cloud service. | [optional] 
**IpPrefixes** | **[]string** | The network CIDRs associated with this cloud service configuration. | [optional] [readonly] 
**Ipv4PrefixCount** | **int32** | The count of IPv4 CIDRs associated with this cloud service configuration. | [optional] [readonly] 
**Ipv6PrefixCount** | **int32** | The count of IPv6 CIDRs associated with this cloud service configuration. | [optional] [readonly] 
**Name** | **string** | The name of the cloud service configuration. | 
**Provider** | [**CloudProvider**](CloudProvider.md) |  | 
**Service** | **string** | The service name. | 

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


