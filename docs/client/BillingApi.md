# pureport\client\BillingApi

All URIs are relative to *https://api.pureport.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AddPaymentInformation**](BillingApi.md#AddPaymentInformation) | **Post** /accounts/{accountId}/billing | Add payment method
[**DeletePaymentInformation**](BillingApi.md#DeletePaymentInformation) | **Delete** /accounts/{accountId}/billing | Delete payment method
[**FindBillingForAccount**](BillingApi.md#FindBillingForAccount) | **Get** /accounts/{accountId}/billing | Payment method for account
[**IsBillingConfiguredForAccount**](BillingApi.md#IsBillingConfiguredForAccount) | **Get** /accounts/{accountId}/billing/configured | Payment method configured for account
[**UpdatePaymentInformation**](BillingApi.md#UpdatePaymentInformation) | **Put** /accounts/{accountId}/billing | Update payment method



## AddPaymentInformation

> AccountBilling AddPaymentInformation(ctx, accountId, optional)

Add payment method

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**accountId** | **string**|  | 
 **optional** | ***AddPaymentInformationOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a AddPaymentInformationOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **accountBilling** | [**optional.Interface of AccountBilling**](AccountBilling.md)|  | 

### Return type

[**AccountBilling**](AccountBilling.md)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeletePaymentInformation

> string DeletePaymentInformation(ctx, accountId)

Delete payment method

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**accountId** | **string**|  | 

### Return type

**string**

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## FindBillingForAccount

> AccountBilling FindBillingForAccount(ctx, accountId)

Payment method for account

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**accountId** | **string**|  | 

### Return type

[**AccountBilling**](AccountBilling.md)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## IsBillingConfiguredForAccount

> bool IsBillingConfiguredForAccount(ctx, accountId, optional)

Payment method configured for account

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**accountId** | **string**|  | 
 **optional** | ***IsBillingConfiguredForAccountOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a IsBillingConfiguredForAccountOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **ancestorOnly** | **optional.Bool**|  | [default to false]

### Return type

**bool**

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdatePaymentInformation

> AccountBilling UpdatePaymentInformation(ctx, accountId, optional)

Update payment method

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**accountId** | **string**|  | 
 **optional** | ***UpdatePaymentInformationOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a UpdatePaymentInformationOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **accountBilling** | [**optional.Interface of AccountBilling**](AccountBilling.md)|  | 

### Return type

[**AccountBilling**](AccountBilling.md)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

