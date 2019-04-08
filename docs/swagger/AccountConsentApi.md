# \AccountConsentApi

All URIs are relative to *http://localhost:8080*

Method | HTTP request | Description
------------- | ------------- | -------------
[**Consent**](AccountConsentApi.md#Consent) | **Post** /accounts/{accountId}/consent | Consent to Master Service Agreement and Acceptable Use Policy for this account
[**HasConsent**](AccountConsentApi.md#HasConsent) | **Get** /accounts/{accountId}/consent | Account consent to Master Service Agreement and Acceptable Use Policy


# **Consent**
> Consent(ctx, accountId)
Consent to Master Service Agreement and Acceptable Use Policy for this account



### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **accountId** | **string**|  | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **HasConsent**
> AccountConsent HasConsent(ctx, accountId)
Account consent to Master Service Agreement and Acceptable Use Policy



### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **accountId** | **string**|  | 

### Return type

[**AccountConsent**](AccountConsent.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

