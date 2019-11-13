# \CloudServicesApi

All URIs are relative to *https://api.pureport.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetCloudService**](CloudServicesApi.md#GetCloudService) | **Get** /cloudServices/{cloudServiceId} | Get cloud service details
[**GetCloudServices**](CloudServicesApi.md#GetCloudServices) | **Get** /cloudServices | List cloud services



## GetCloudService

> CloudService GetCloudService(ctx, cloudServiceId)

Get cloud service details

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**cloudServiceId** | **string**|  | 

### Return type

[**CloudService**](CloudService.md)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetCloudServices

> []CloudService GetCloudServices(ctx, )

List cloud services

### Required Parameters

This endpoint does not need any parameter.

### Return type

[**[]CloudService**](CloudService.md)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

