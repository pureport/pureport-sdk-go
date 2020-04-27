# pureport\client\GatewayMetricsApi

All URIs are relative to *https://api.pureport.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ConnectivityOverTime**](GatewayMetricsApi.md#ConnectivityOverTime) | **Post** /gateways/{gatewayId}/metrics/connectivity | Retrieve connectivity details over time for gateway
[**CurrentConnectivity**](GatewayMetricsApi.md#CurrentConnectivity) | **Get** /gateways/{gatewayId}/metrics/connectivity/current | Retrieve latest connectivity details for gateway



## ConnectivityOverTime

> []ConnectivityByGateway ConnectivityOverTime(ctx, gatewayId, optional)

Retrieve connectivity details over time for gateway

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**gatewayId** | **string**|  | 
 **optional** | ***ConnectivityOverTimeOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a ConnectivityOverTimeOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **dateFilter** | [**optional.Interface of DateFilter**](DateFilter.md)|  | 

### Return type

[**[]ConnectivityByGateway**](ConnectivityByGateway.md)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json, text/csv

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CurrentConnectivity

> []ConnectivityByGateway CurrentConnectivity(ctx, gatewayId)

Retrieve latest connectivity details for gateway

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**gatewayId** | **string**|  | 

### Return type

[**[]ConnectivityByGateway**](ConnectivityByGateway.md)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, text/csv

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

