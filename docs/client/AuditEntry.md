# AuditEntry

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Account** | [**Link**](Link.md) |  | [optional] 
**Changes** | [**[]ChangeObject**](ChangeObject.md) |  | [optional] [readonly] 
**CorrelationId** | **string** |  | [optional] [readonly] 
**EventType** | **string** |  | [optional] [readonly] 
**IpAddress** | **string** |  | [optional] [readonly] 
**Principal** | [**Link**](Link.md) |  | [optional] 
**Request** | [**AuditEntryRequest**](AuditEntryRequest.md) |  | [optional] 
**Response** | [**AuditEntryResponse**](AuditEntryResponse.md) |  | [optional] 
**Result** | **string** |  | [optional] [readonly] 
**Source** | **string** |  | [optional] [readonly] 
**Subject** | [**Link**](Link.md) |  | [optional] 
**SubjectType** | **string** |  | [optional] [readonly] 
**Timestamp** | [**time.Time**](time.Time.md) |  | [optional] [readonly] 
**UserAgent** | **string** |  | [optional] [readonly] 

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


