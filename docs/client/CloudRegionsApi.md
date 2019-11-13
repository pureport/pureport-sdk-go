# \CloudRegionsApi

All URIs are relative to *https://api.pureport.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateCloudRegion**](CloudRegionsApi.md#CreateCloudRegion) | **Post** /cloudRegions | Add cloud region
[**DeleteCloudRegion**](CloudRegionsApi.md#DeleteCloudRegion) | **Delete** /cloudRegions/{cloudRegionId} | Delete cloud region
[**GetCloudRegion**](CloudRegionsApi.md#GetCloudRegion) | **Get** /cloudRegions/{cloudRegionId} | Get cloud region details
[**GetCloudRegions**](CloudRegionsApi.md#GetCloudRegions) | **Get** /cloudRegions | List cloud regions
[**UpdateCloudRegion**](CloudRegionsApi.md#UpdateCloudRegion) | **Put** /cloudRegions/{cloudRegionId} | Update cloud region



## CreateCloudRegion

> CloudRegion CreateCloudRegion(ctx, optional)

Add cloud region

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***CreateCloudRegionOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a CreateCloudRegionOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **cloudRegion** | [**optional.Interface of CloudRegion**](CloudRegion.md)|  | 

### Return type

[**CloudRegion**](CloudRegion.md)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteCloudRegion

> DeleteCloudRegion(ctx, cloudRegionId)

Delete cloud region

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**cloudRegionId** | **string**|  | 

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


## GetCloudRegion

> CloudRegion GetCloudRegion(ctx, cloudRegionId)

Get cloud region details

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**cloudRegionId** | **string**|  | 

### Return type

[**CloudRegion**](CloudRegion.md)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetCloudRegions

> []CloudRegion GetCloudRegions(ctx, )

List cloud regions

### Required Parameters

This endpoint does not need any parameter.

### Return type

[**[]CloudRegion**](CloudRegion.md)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateCloudRegion

> CloudRegion UpdateCloudRegion(ctx, cloudRegionId, optional)

Update cloud region

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**cloudRegionId** | **string**|  | 
 **optional** | ***UpdateCloudRegionOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a UpdateCloudRegionOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **cloudRegion** | [**optional.Interface of CloudRegion**](CloudRegion.md)|  | 

### Return type

[**CloudRegion**](CloudRegion.md)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

