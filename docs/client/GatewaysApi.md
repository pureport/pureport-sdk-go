# \GatewaysApi

All URIs are relative to *https://api.pureport.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateGatewayTask**](GatewaysApi.md#CreateGatewayTask) | **Post** /gateways/{gatewayId}/tasks | Add a new task
[**GetGatewayBGPRoutes**](GatewaysApi.md#GetGatewayBGPRoutes) | **Get** /gateways/{gatewayId}/bgpRoutes | Get gateway bgp routes
[**GetGatewayTasks**](GatewaysApi.md#GetGatewayTasks) | **Get** /gateways/{gatewayId}/tasks | List connection tasks



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

