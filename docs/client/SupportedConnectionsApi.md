# pureport\client\SupportedConnectionsApi

All URIs are relative to *http://localhost:8080*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetAccountSupportedConnections**](SupportedConnectionsApi.md#GetAccountSupportedConnections) | **Get** /accounts/{accountId}/supportedConnections | List supported connections
[**GetSupportedConnection**](SupportedConnectionsApi.md#GetSupportedConnection) | **Get** /supportedConnections/{supportedConnectionId} | Get supported connection details


# **GetAccountSupportedConnections**
> []SupportedConnection GetAccountSupportedConnections(ctx, accountId)
List supported connections



### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **accountId** | **string**|  | 

### Return type

[**[]SupportedConnection**](SupportedConnection.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetSupportedConnection**
> SupportedConnection GetSupportedConnection(ctx, supportedConnectionId)
Get supported connection details



### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **supportedConnectionId** | **string**|  | 

### Return type

[**SupportedConnection**](SupportedConnection.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

