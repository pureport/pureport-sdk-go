# AuditEntry

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Account** | [**Link**](Link.md) |  | [optional] 
**Changes** | [**[]ChangeObject**](ChangeObject.md) | The changes to the Pureport entity. | [optional] [readonly] 
**CorrelationId** | **string** | An identifier to group multiple related audit entries. | [optional] [readonly] 
**EventType** | **string** | The type of event for this entry. | [optional] [readonly] 
**IpAddress** | **string** |  | [optional] [readonly] 
**Principal** | [**Link**](Link.md) |  | [optional] 
**Request** | [**AuditEntryRequest**](AuditEntryRequest.md) |  | [optional] 
**Response** | [**AuditEntryResponse**](AuditEntryResponse.md) |  | [optional] 
**Result** | **string** | The result of the request that triggered this entry. | [optional] [readonly] 
**Source** | **string** | The source of the request that triggered this entry. | [optional] [readonly] 
**Subject** | [**Link**](Link.md) |  | [optional] 
**SubjectType** | **string** | The type of the Pureport entity associated with this entry. | [optional] [readonly] 
**Timestamp** | [**time.Time**](time.Time.md) |  | [optional] [readonly] 
**UserAgent** | **string** | The User-Agent of the request that triggered this entry. | [optional] [readonly] 

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


