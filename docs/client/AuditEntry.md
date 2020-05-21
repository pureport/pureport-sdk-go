# AuditEntry

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Account** | [**Link**](Link.md) |  | [optional] 
**Changes** | [**[]ChangeObject**](ChangeObject.md) | The changes to the Pureport entity. | [optional] [readonly] 
**CorrelationId** | **string** | An identifier to group multiple related audit entries. | [optional] [readonly] 
**EventType** | [**AuditEntryEventType**](AuditEntryEventType.md) |  | [optional] 
**IpAddress** | **string** |  | [optional] [readonly] 
**Principal** | [**Link**](Link.md) |  | [optional] 
**Request** | [**AuditEntryRequest**](AuditEntryRequest.md) |  | [optional] 
**Response** | [**AuditEntryResponse**](AuditEntryResponse.md) |  | [optional] 
**Result** | [**AuditEntryResult**](AuditEntryResult.md) |  | [optional] 
**Source** | [**AuditEntrySource**](AuditEntrySource.md) |  | [optional] 
**Subject** | [**Link**](Link.md) |  | [optional] 
**SubjectType** | [**AuditEntrySubjectType**](AuditEntrySubjectType.md) |  | [optional] 
**Timestamp** | [**time.Time**](time.Time.md) |  | [optional] [readonly] 
**UserAgent** | **string** | The User-Agent of the request that triggered this entry. | [optional] [readonly] 

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


