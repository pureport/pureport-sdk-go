# pureport\client\ConnectionsApi

All URIs are relative to *https://api.pureport.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AddConnection**](ConnectionsApi.md#AddConnection) | **Post** /networks/{networkId}/connections | Add new connection
[**CreateConnectionTask**](ConnectionsApi.md#CreateConnectionTask) | **Post** /connections/{connectionId}/tasks | Add a new task
[**DeleteConnection**](ConnectionsApi.md#DeleteConnection) | **Delete** /connections/{connectionId} | Delete connection
[**FindConnections**](ConnectionsApi.md#FindConnections) | **Get** /accounts/{accountId}/connections | List connections across all networks for the account
[**GetConnection**](ConnectionsApi.md#GetConnection) | **Get** /connections/{connectionId} | Get connection details
[**GetConnectionTasks**](ConnectionsApi.md#GetConnectionTasks) | **Get** /connections/{connectionId}/tasks | List connection tasks
[**GetConnections**](ConnectionsApi.md#GetConnections) | **Get** /networks/{networkId}/connections | List network connections
[**UpdateConnection**](ConnectionsApi.md#UpdateConnection) | **Put** /connections/{connectionId} | Update connection


# **AddConnection**
> Connection AddConnection(ctx, networkId, optional)
Add new connection



### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **networkId** | **string**|  | 
 **optional** | ***AddConnectionOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a AddConnectionOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**optional.Interface of Connection**](Connection.md)|  | 

### Return type

[**Connection**](Connection.md)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **CreateConnectionTask**
> Task CreateConnectionTask(ctx, connectionId, optional)
Add a new task



### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **connectionId** | **string**|  | 
 **optional** | ***CreateConnectionTaskOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a CreateConnectionTaskOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**optional.Interface of Task**](Task.md)|  | 

### Return type

[**Task**](Task.md)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteConnection**
> Connection DeleteConnection(ctx, connectionId)
Delete connection



### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **connectionId** | **string**|  | 

### Return type

[**Connection**](Connection.md)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **FindConnections**
> []Connection FindConnections(ctx, accountId)
List connections across all networks for the account



### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **accountId** | **string**|  | 

### Return type

[**[]Connection**](Connection.md)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetConnection**
> Connection GetConnection(ctx, connectionId)
Get connection details



### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **connectionId** | **string**|  | 

### Return type

[**Connection**](Connection.md)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetConnectionTasks**
> []Task GetConnectionTasks(ctx, connectionId)
List connection tasks



### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **connectionId** | **string**|  | 

### Return type

[**[]Task**](Task.md)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetConnections**
> []Connection GetConnections(ctx, networkId)
List network connections



### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **networkId** | **string**|  | 

### Return type

[**[]Connection**](Connection.md)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdateConnection**
> Connection UpdateConnection(ctx, connectionId, optional)
Update connection



### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **connectionId** | **string**|  | 
 **optional** | ***UpdateConnectionOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a UpdateConnectionOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**optional.Interface of Connection**](Connection.md)|  | 

### Return type

[**Connection**](Connection.md)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

