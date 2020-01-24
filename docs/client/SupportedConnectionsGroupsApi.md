# pureport\client\SupportedConnectionsGroupsApi

All URIs are relative to *https://api.pureport.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateSupportedConnectionGroup**](SupportedConnectionsGroupsApi.md#CreateSupportedConnectionGroup) | **Post** /supportedConnections/groups | Add supported connection group
[**DeleteSupportedConnectionGroup**](SupportedConnectionsGroupsApi.md#DeleteSupportedConnectionGroup) | **Delete** /supportedConnections/groups/{supportedConnectionGroupId} | Delete supported connection group
[**GetSupportedConnectionGroup**](SupportedConnectionsGroupsApi.md#GetSupportedConnectionGroup) | **Get** /supportedConnections/groups/{supportedConnectionGroupId} | Get supported connection group details
[**GetSupportedConnectionGroups**](SupportedConnectionsGroupsApi.md#GetSupportedConnectionGroups) | **Get** /supportedConnections/groups | List supported connections groups
[**UpdateSupportedConnectionGroup**](SupportedConnectionsGroupsApi.md#UpdateSupportedConnectionGroup) | **Put** /supportedConnections/groups/{supportedConnectionGroupId} | Update supported connection group



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

