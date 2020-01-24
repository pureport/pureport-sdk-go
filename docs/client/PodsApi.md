# pureport\client\PodsApi

All URIs are relative to *https://api.pureport.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreatePod**](PodsApi.md#CreatePod) | **Post** /pods | Add new pod
[**DeletePodResource**](PodsApi.md#DeletePodResource) | **Delete** /pods/{podId} | Delete pod
[**FindPods**](PodsApi.md#FindPods) | **Get** /pods | List all pods
[**GetPodResource**](PodsApi.md#GetPodResource) | **Get** /pods/{podId} | Get pod details
[**UpdatePodResource**](PodsApi.md#UpdatePodResource) | **Put** /pods/{podId} | Update pod



## CreatePod

> Pod CreatePod(ctx, optional)

Add new pod

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***CreatePodOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a CreatePodOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **pod** | [**optional.Interface of Pod**](Pod.md)|  | 

### Return type

[**Pod**](Pod.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeletePodResource

> DeletePodResource(ctx, podId)

Delete pod

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**podId** | **string**|  | 

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


## FindPods

> []Pod FindPods(ctx, )

List all pods

### Required Parameters

This endpoint does not need any parameter.

### Return type

[**[]Pod**](Pod.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetPodResource

> Pod GetPodResource(ctx, podId)

Get pod details

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**podId** | **string**|  | 

### Return type

[**Pod**](Pod.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdatePodResource

> Pod UpdatePodResource(ctx, podId, optional)

Update pod

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**podId** | **string**|  | 
 **optional** | ***UpdatePodResourceOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a UpdatePodResourceOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **pod** | [**optional.Interface of Pod**](Pod.md)|  | 

### Return type

[**Pod**](Pod.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

