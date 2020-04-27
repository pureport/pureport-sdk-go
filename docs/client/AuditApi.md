# pureport\client\AuditApi

All URIs are relative to *https://api.pureport.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**FindAuditLogs**](AuditApi.md#FindAuditLogs) | **Get** /accounts/{accountId}/auditLog | 



## FindAuditLogs

> PageAuditEntry FindAuditLogs(ctx, accountId, optional)



### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**accountId** | **string**|  | 
 **optional** | ***FindAuditLogsOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a FindAuditLogsOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **pageNumber** | **optional.Int32**|  | [default to 0]
 **pageSize** | **optional.Int32**|  | [default to 100]
 **sort** | **optional.String**|  | 
 **sortDirection** | **optional.String**|  | 
 **startTime** | **optional.Time**|  | 
 **endTime** | **optional.Time**|  | 
 **includeChildAccounts** | **optional.Bool**|  | [default to false]
 **includeChildSubjects** | **optional.Bool**| include child object records if specified subject is of type NETWORK or CONNECTION | [default to false]
 **eventTypes** | [**optional.Interface of []string**](string.md)|  | 
 **result** | **optional.String**|  | 
 **principalId** | **optional.String**|  | 
 **ipAddress** | **optional.String**|  | 
 **correlationId** | **optional.String**|  | 
 **subjectId** | **optional.String**|  | 
 **subjectType** | **optional.String**|  | 

### Return type

[**PageAuditEntry**](PageAuditEntry.md)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

