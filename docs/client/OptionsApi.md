# pureport\client\OptionsApi

All URIs are relative to *https://api.pureport.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetOptions**](OptionsApi.md#GetOptions) | **Get** /options | Get available options
[**IsBackbone**](OptionsApi.md#IsBackbone) | **Get** /options/isBackbone | Get whether this configuration will have longHaul related charges.



## GetOptions

> map[string][]Option GetOptions(ctx, optional)

Get available options

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***GetOptionsOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a GetOptionsOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **type_** | [**optional.Interface of []string**](string.md)|  | 

### Return type

[**map[string][]Option**](array.md)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## IsBackbone

> map[string]bool IsBackbone(ctx, optional)

Get whether this configuration will have longHaul related charges.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***IsBackboneOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a IsBackboneOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **portId** | **optional.String**|  | 
 **locationId** | **optional.String**|  | 
 **speed** | **optional.Int32**|  | 
 **billingTerm** | [**optional.Interface of BillingTerm**](.md)|  | 

### Return type

**map[string]bool**

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

