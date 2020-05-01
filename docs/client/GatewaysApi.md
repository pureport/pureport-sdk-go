# pureport\client\GatewaysApi

All URIs are relative to *https://api.pureport.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ConnectivityOverTime**](GatewaysApi.md#ConnectivityOverTime) | **Post** /gateways/{gatewayId}/metrics/connectivity | Retrieve connectivity details over time for gateway
[**CreateGatewayTask**](GatewaysApi.md#CreateGatewayTask) | **Post** /gateways/{gatewayId}/tasks | Add a new task
[**CurrentConnectivity**](GatewaysApi.md#CurrentConnectivity) | **Get** /gateways/{gatewayId}/metrics/connectivity/current | Retrieve latest connectivity details for gateway
[**GetGateway**](GatewaysApi.md#GetGateway) | **Get** /gateways/{gatewayId} | Get gateway details
[**GetGatewayBGPRoutes**](GatewaysApi.md#GetGatewayBGPRoutes) | **Get** /gateways/{gatewayId}/bgpRoutes | Get gateway bgp routes
[**GetGatewayIPRoutes**](GatewaysApi.md#GetGatewayIPRoutes) | **Get** /gateways/{gatewayId}/ipRoutes | Get gateway ip routes
[**GetGatewayTasks**](GatewaysApi.md#GetGatewayTasks) | **Get** /gateways/{gatewayId}/tasks | List connection tasks
[**RespawnGateway**](GatewaysApi.md#RespawnGateway) | **Post** /gateways/{gatewayId}/respawn | Respawn gateway
[**StartGateway**](GatewaysApi.md#StartGateway) | **Post** /gateways/{gatewayId}/start | Start gateway
[**StopGateway**](GatewaysApi.md#StopGateway) | **Post** /gateways/{gatewayId}/stop | Stop gateway



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


## CreateGatewayTask

> Task CreateGatewayTask(ctx, gatewayId, optional)

Add a new task

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**gatewayId** | **string**|  | 
 **optional** | ***CreateGatewayTaskOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a CreateGatewayTaskOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **task** | [**optional.Interface of Task**](Task.md)|  | 

### Return type

[**Task**](Task.md)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

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


## GetGateway

> Gateway GetGateway(ctx, gatewayId)

Get gateway details

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**gatewayId** | **string**|  | 

### Return type

[**Gateway**](Gateway.md)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetGatewayBGPRoutes

> []BgpRoute GetGatewayBGPRoutes(ctx, gatewayId)

Get gateway bgp routes

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**gatewayId** | **string**|  | 

### Return type

[**[]BgpRoute**](BGPRoute.md)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetGatewayIPRoutes

> []IpRoute GetGatewayIPRoutes(ctx, gatewayId)

Get gateway ip routes

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**gatewayId** | **string**|  | 

### Return type

[**[]IpRoute**](IPRoute.md)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetGatewayTasks

> []Task GetGatewayTasks(ctx, gatewayId)

List connection tasks

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**gatewayId** | **string**|  | 

### Return type

[**[]Task**](Task.md)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RespawnGateway

> Gateway RespawnGateway(ctx, gatewayId)

Respawn gateway

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**gatewayId** | **string**|  | 

### Return type

[**Gateway**](Gateway.md)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## StartGateway

> Gateway StartGateway(ctx, gatewayId)

Start gateway

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**gatewayId** | **string**|  | 

### Return type

[**Gateway**](Gateway.md)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## StopGateway

> Gateway StopGateway(ctx, gatewayId)

Stop gateway

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**gatewayId** | **string**|  | 

### Return type

[**Gateway**](Gateway.md)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

