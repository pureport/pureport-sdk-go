# \TasksApi

All URIs are relative to *https://api.pureport.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateConnectionTask**](TasksApi.md#CreateConnectionTask) | **Post** /connections/{connectionId}/tasks | Add a new task
[**CreateGatewayTask**](TasksApi.md#CreateGatewayTask) | **Post** /gateways/{gatewayId}/tasks | Add a new task
[**DeleteTask**](TasksApi.md#DeleteTask) | **Delete** /tasks/{taskId} | Delete Task
[**GetConnectionTasks**](TasksApi.md#GetConnectionTasks) | **Get** /connections/{connectionId}/tasks | List connection tasks
[**GetGatewayTasks**](TasksApi.md#GetGatewayTasks) | **Get** /gateways/{gatewayId}/tasks | List connection tasks
[**GetTask**](TasksApi.md#GetTask) | **Get** /tasks/{taskId} | Get Task details
[**GetTasks**](TasksApi.md#GetTasks) | **Get** /tasks | List Tasks



## CreateConnectionTask

> Task CreateConnectionTask(ctx, connectionId, optional)

Add a new task

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**connectionId** | **string**|  | 
 **optional** | ***CreateConnectionTaskOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a CreateConnectionTaskOpts struct


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


## DeleteTask

> DeleteTask(ctx, taskId)

Delete Task

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**taskId** | **string**|  | 

### Return type

 (empty response body)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetConnectionTasks

> []Task GetConnectionTasks(ctx, connectionId)

List connection tasks

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**connectionId** | **string**|  | 

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


## GetTask

> Task GetTask(ctx, taskId)

Get Task details

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**taskId** | **string**|  | 

### Return type

[**Task**](Task.md)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetTasks

> []Task GetTasks(ctx, )

List Tasks

### Required Parameters

This endpoint does not need any parameter.

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

