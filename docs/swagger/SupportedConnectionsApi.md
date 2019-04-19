# \SupportedConnectionsApi

All URIs are relative to *http://api.pureport.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**Get17**](SupportedConnectionsApi.md#Get17) | **Get** /supportedConnections/{supportedConnectionId} | Get supported connection details
[**GetSupportedConnections**](SupportedConnectionsApi.md#GetSupportedConnections) | **Get** /supportedConnections | List supported connections


# **Get17**
> SupportedConnection Get17(ctx, supportedConnectionId)
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

# **GetSupportedConnections**
> []SupportedConnection GetSupportedConnections(ctx, )
List supported connections



### Required Parameters
This endpoint does not need any parameter.

### Return type

[**[]SupportedConnection**](SupportedConnection.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

