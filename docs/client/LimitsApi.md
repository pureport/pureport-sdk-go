# pureport\client\LimitsApi

All URIs are relative to *https://api.pureport.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateResourceLimit**](LimitsApi.md#CreateResourceLimit) | **Post** /accounts/{accountId}/limits/{childResourceType} | Create a resource limit override
[**CreateResourceLimit1**](LimitsApi.md#CreateResourceLimit1) | **Post** /networks/{networkId}/limits/{childResourceType} | Create a resource limit override
[**DeleteResourceLimit**](LimitsApi.md#DeleteResourceLimit) | **Delete** /accounts/{accountId}/limits/{childResourceType} | Delete a resource limit override
[**DeleteResourceLimit1**](LimitsApi.md#DeleteResourceLimit1) | **Delete** /networks/{networkId}/limits/{childResourceType} | Delete a resource limit override
[**GetResourceLimit**](LimitsApi.md#GetResourceLimit) | **Get** /accounts/{accountId}/limits | Get resource limits including default values
[**GetResourceLimit1**](LimitsApi.md#GetResourceLimit1) | **Get** /accounts/{accountId}/limits/{childResourceType} | Get resource limit override. Does not return default.
[**GetResourceLimit2**](LimitsApi.md#GetResourceLimit2) | **Get** /networks/{networkId}/limits | Get resource limits including default values
[**GetResourceLimit3**](LimitsApi.md#GetResourceLimit3) | **Get** /networks/{networkId}/limits/{childResourceType} | Get resource limit override. Does not return default.
[**UpdateResourceLimit**](LimitsApi.md#UpdateResourceLimit) | **Put** /accounts/{accountId}/limits/{childResourceType} | Update a resource limit override
[**UpdateResourceLimit1**](LimitsApi.md#UpdateResourceLimit1) | **Put** /networks/{networkId}/limits/{childResourceType} | Update a resource limit override



## CreateResourceLimit

> ResourceLimit CreateResourceLimit(ctx, childResourceType, accountId, optional)

Create a resource limit override

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**childResourceType** | **string**|  | 
**accountId** | **string**|  | 
 **optional** | ***CreateResourceLimitOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a CreateResourceLimitOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **resourceLimit** | [**optional.Interface of ResourceLimit**](ResourceLimit.md)|  | 

### Return type

[**ResourceLimit**](ResourceLimit.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateResourceLimit1

> ResourceLimit CreateResourceLimit1(ctx, childResourceType, networkId, optional)

Create a resource limit override

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**childResourceType** | **string**|  | 
**networkId** | **string**|  | 
 **optional** | ***CreateResourceLimit1Opts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a CreateResourceLimit1Opts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **resourceLimit** | [**optional.Interface of ResourceLimit**](ResourceLimit.md)|  | 

### Return type

[**ResourceLimit**](ResourceLimit.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteResourceLimit

> DeleteResourceLimit(ctx, childResourceType, accountId)

Delete a resource limit override

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**childResourceType** | **string**|  | 
**accountId** | **string**|  | 

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


## DeleteResourceLimit1

> DeleteResourceLimit1(ctx, childResourceType, networkId)

Delete a resource limit override

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**childResourceType** | **string**|  | 
**networkId** | **string**|  | 

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


## GetResourceLimit

> []ResourceLimit GetResourceLimit(ctx, accountId)

Get resource limits including default values

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**accountId** | **string**|  | 

### Return type

[**[]ResourceLimit**](ResourceLimit.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetResourceLimit1

> ResourceLimit GetResourceLimit1(ctx, childResourceType, accountId)

Get resource limit override. Does not return default.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**childResourceType** | **string**|  | 
**accountId** | **string**|  | 

### Return type

[**ResourceLimit**](ResourceLimit.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetResourceLimit2

> []ResourceLimit GetResourceLimit2(ctx, networkId)

Get resource limits including default values

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**networkId** | **string**|  | 

### Return type

[**[]ResourceLimit**](ResourceLimit.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetResourceLimit3

> ResourceLimit GetResourceLimit3(ctx, childResourceType, networkId)

Get resource limit override. Does not return default.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**childResourceType** | **string**|  | 
**networkId** | **string**|  | 

### Return type

[**ResourceLimit**](ResourceLimit.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateResourceLimit

> ResourceLimit UpdateResourceLimit(ctx, childResourceType, accountId, optional)

Update a resource limit override

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**childResourceType** | **string**|  | 
**accountId** | **string**|  | 
 **optional** | ***UpdateResourceLimitOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a UpdateResourceLimitOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **resourceLimit** | [**optional.Interface of ResourceLimit**](ResourceLimit.md)|  | 

### Return type

[**ResourceLimit**](ResourceLimit.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateResourceLimit1

> ResourceLimit UpdateResourceLimit1(ctx, childResourceType, networkId, optional)

Update a resource limit override

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**childResourceType** | **string**|  | 
**networkId** | **string**|  | 
 **optional** | ***UpdateResourceLimit1Opts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a UpdateResourceLimit1Opts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **resourceLimit** | [**optional.Interface of ResourceLimit**](ResourceLimit.md)|  | 

### Return type

[**ResourceLimit**](ResourceLimit.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

