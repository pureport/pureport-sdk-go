# pureport\client\InvoicesApi

All URIs are relative to *https://api.pureport.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ListInvoicesForAccount**](InvoicesApi.md#ListInvoicesForAccount) | **Post** /accounts/{accountId}/invoices | List invoices
[**ListUpcomingInvoicesForAccount**](InvoicesApi.md#ListUpcomingInvoicesForAccount) | **Get** /accounts/{accountId}/invoices/upcoming | List upcoming invoices



## ListInvoicesForAccount

> []AccountInvoice ListInvoicesForAccount(ctx, accountId, optional)

List invoices

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**accountId** | **string**|  | 
 **optional** | ***ListInvoicesForAccountOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a ListInvoicesForAccountOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **requestBody** | [**optional.Interface of map[string]map[string]interface{}**](map[string]interface{}.md)|  | 

### Return type

[**[]AccountInvoice**](AccountInvoice.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListUpcomingInvoicesForAccount

> []AccountInvoice ListUpcomingInvoicesForAccount(ctx, accountId)

List upcoming invoices

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**accountId** | **string**|  | 

### Return type

[**[]AccountInvoice**](AccountInvoice.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

