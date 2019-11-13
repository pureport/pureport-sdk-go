# \SupportedPortsApi

All URIs are relative to *https://api.pureport.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetSupportedPorts**](SupportedPortsApi.md#GetSupportedPorts) | **Get** /accounts/{accountId}/supportedPorts | List supported ports



## GetSupportedPorts

> []SupportedPort GetSupportedPorts(ctx, facility, accountId, optional)

List supported ports

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**facility** | **string**|  | 
**accountId** | **string**|  | 
 **optional** | ***GetSupportedPortsOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a GetSupportedPortsOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **includeProvisioned** | **optional.Bool**|  | 

### Return type

[**[]SupportedPort**](SupportedPort.md)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

