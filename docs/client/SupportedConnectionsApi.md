# pureport\client\SupportedConnectionsApi

All URIs are relative to *https://api.pureport.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateSupportedConnection**](SupportedConnectionsApi.md#CreateSupportedConnection) | **Post** /supportedConnections | Add supported connection
[**CreateSupportedConnectionGroup**](SupportedConnectionsApi.md#CreateSupportedConnectionGroup) | **Post** /supportedConnections/groups | Add supported connection group
[**DeleteSupportedConnection**](SupportedConnectionsApi.md#DeleteSupportedConnection) | **Delete** /supportedConnections/{supportedConnectionId} | Delete supported connection
[**DeleteSupportedConnectionGroup**](SupportedConnectionsApi.md#DeleteSupportedConnectionGroup) | **Delete** /supportedConnections/groups/{supportedConnectionGroupId} | Delete supported connection group
[**GetAccountSupportedConnections**](SupportedConnectionsApi.md#GetAccountSupportedConnections) | **Get** /accounts/{accountId}/supportedConnections | List supported connections
[**GetSupportedConnection**](SupportedConnectionsApi.md#GetSupportedConnection) | **Get** /supportedConnections/{supportedConnectionId} | Get supported connection details
[**GetSupportedConnectionGroup**](SupportedConnectionsApi.md#GetSupportedConnectionGroup) | **Get** /supportedConnections/groups/{supportedConnectionGroupId} | Get supported connection group details
[**GetSupportedConnectionGroups**](SupportedConnectionsApi.md#GetSupportedConnectionGroups) | **Get** /supportedConnections/groups | List supported connections groups
[**GetSupportedConnections**](SupportedConnectionsApi.md#GetSupportedConnections) | **Get** /supportedConnections | List supported connections
[**UpdateSupportedConnection**](SupportedConnectionsApi.md#UpdateSupportedConnection) | **Put** /supportedConnections/{supportedConnectionId} | Update supported connection
[**UpdateSupportedConnectionGroup**](SupportedConnectionsApi.md#UpdateSupportedConnectionGroup) | **Put** /supportedConnections/groups/{supportedConnectionGroupId} | Update supported connection group



## CreateSupportedConnection

> SupportedConnection CreateSupportedConnection(ctx, optional)

Add supported connection

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***CreateSupportedConnectionOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a CreateSupportedConnectionOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **supportedConnection** | [**optional.Interface of SupportedConnection**](SupportedConnection.md)|  | 

### Return type

[**SupportedConnection**](SupportedConnection.md)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateSupportedConnectionGroup

> SupportedConnectionGroup CreateSupportedConnectionGroup(ctx, optional)

Add supported connection group

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***CreateSupportedConnectionGroupOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a CreateSupportedConnectionGroupOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **supportedConnectionGroup** | [**optional.Interface of SupportedConnectionGroup**](SupportedConnectionGroup.md)|  | 

### Return type

[**SupportedConnectionGroup**](SupportedConnectionGroup.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteSupportedConnection

> DeleteSupportedConnection(ctx, supportedConnectionId)

Delete supported connection

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**supportedConnectionId** | **string**|  | 

### Return type

 (empty response body)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteSupportedConnectionGroup

> DeleteSupportedConnectionGroup(ctx, supportedConnectionGroupId)

Delete supported connection group

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**supportedConnectionGroupId** | **string**|  | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetAccountSupportedConnections

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

[OAuth2](../README.md#OAuth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetSupportedConnection

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

[OAuth2](../README.md#OAuth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetSupportedConnectionGroup

> SupportedConnectionGroup GetSupportedConnectionGroup(ctx, supportedConnectionGroupId)

Get supported connection group details

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**supportedConnectionGroupId** | **string**|  | 

### Return type

[**SupportedConnectionGroup**](SupportedConnectionGroup.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetSupportedConnectionGroups

> []SupportedConnectionGroup GetSupportedConnectionGroups(ctx, )

List supported connections groups

### Required Parameters

This endpoint does not need any parameter.

### Return type

[**[]SupportedConnectionGroup**](SupportedConnectionGroup.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetSupportedConnections

> []SupportedConnection GetSupportedConnections(ctx, )

List supported connections

### Required Parameters

This endpoint does not need any parameter.

### Return type

[**[]SupportedConnection**](SupportedConnection.md)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateSupportedConnection

> SupportedConnection UpdateSupportedConnection(ctx, supportedConnectionId, optional)

Update supported connection

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**supportedConnectionId** | **string**|  | 
 **optional** | ***UpdateSupportedConnectionOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a UpdateSupportedConnectionOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **supportedConnection** | [**optional.Interface of SupportedConnection**](SupportedConnection.md)|  | 

### Return type

[**SupportedConnection**](SupportedConnection.md)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateSupportedConnectionGroup

> SupportedConnectionGroup UpdateSupportedConnectionGroup(ctx, supportedConnectionGroupId, optional)

Update supported connection group

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**supportedConnectionGroupId** | **string**|  | 
 **optional** | ***UpdateSupportedConnectionGroupOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a UpdateSupportedConnectionGroupOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **supportedConnectionGroup** | [**optional.Interface of SupportedConnectionGroup**](SupportedConnectionGroup.md)|  | 

### Return type

[**SupportedConnectionGroup**](SupportedConnectionGroup.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

