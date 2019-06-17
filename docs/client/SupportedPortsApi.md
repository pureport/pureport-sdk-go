# pureport\client\SupportedPortsApi

All URIs are relative to *http://api.pureport.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetSupportedPorts**](SupportedPortsApi.md#GetSupportedPorts) | **Get** /accounts/{accountId}/supportedPorts | List supported ports


# **GetSupportedPorts**
> []SupportedPort GetSupportedPorts(ctx, facility, accountId)
List supported ports



### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **facility** | **string**|  | 
  **accountId** | **string**|  | 

### Return type

[**[]SupportedPort**](SupportedPort.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

