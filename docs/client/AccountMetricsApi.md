# pureport\client\AccountMetricsApi

All URIs are relative to *https://api.pureport.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**UsageByConnection**](AccountMetricsApi.md#UsageByConnection) | **Post** /accounts/{accountId}/metrics/usageByConnection | Retrieve usage by connection
[**UsageByConnectionAndTime**](AccountMetricsApi.md#UsageByConnectionAndTime) | **Post** /accounts/{accountId}/metrics/usageByConnectionAndTime | Retrieve usage by connection and time
[**UsageByNetworkAndTime**](AccountMetricsApi.md#UsageByNetworkAndTime) | **Post** /accounts/{accountId}/metrics/usageByNetworkAndTime | Retrieve usage by network over time



## UsageByConnection

> []NetworkConnectionEgressIngress UsageByConnection(ctx, accountId, optional)

Retrieve usage by connection

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**accountId** | **string**|  | 
 **optional** | ***UsageByConnectionOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a UsageByConnectionOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **usageByConnectionOptions** | [**optional.Interface of UsageByConnectionOptions**](UsageByConnectionOptions.md)|  | 

### Return type

[**[]NetworkConnectionEgressIngress**](NetworkConnectionEgressIngress.md)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, text/csv

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UsageByConnectionAndTime

> []ConnectionTimeEgressIngress UsageByConnectionAndTime(ctx, accountId, optional)

Retrieve usage by connection and time

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**accountId** | **string**|  | 
 **optional** | ***UsageByConnectionAndTimeOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a UsageByConnectionAndTimeOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **usageByConnectionAndTimeOptions** | [**optional.Interface of UsageByConnectionAndTimeOptions**](UsageByConnectionAndTimeOptions.md)|  | 

### Return type

[**[]ConnectionTimeEgressIngress**](ConnectionTimeEgressIngress.md)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, text/csv

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UsageByNetworkAndTime

> []NetworkTimeUsage UsageByNetworkAndTime(ctx, accountId, optional)

Retrieve usage by network over time

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**accountId** | **string**|  | 
 **optional** | ***UsageByNetworkAndTimeOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a UsageByNetworkAndTimeOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **usageByNetworkAndTimeOptions** | [**optional.Interface of UsageByNetworkAndTimeOptions**](UsageByNetworkAndTimeOptions.md)|  | 

### Return type

[**[]NetworkTimeUsage**](NetworkTimeUsage.md)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json, text/csv

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

