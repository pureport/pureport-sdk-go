# \HelpApi

All URIs are relative to *https://api.pureport.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetHelpUrl**](HelpApi.md#GetHelpUrl) | **Get** /help | Retrieve the URL for the help portal using an SSO login



## GetHelpUrl

> string GetHelpUrl(ctx, optional)

Retrieve the URL for the help portal using an SSO login

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***GetHelpUrlOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a GetHelpUrlOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **redirectTo** | **optional.String**|  | 

### Return type

**string**

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: text/plain

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

