# pureport\client\GatewaysApi

All URIs are relative to *https://api.pureport.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetGatewayBGPRoutes**](GatewaysApi.md#GetGatewayBGPRoutes) | **Get** /gateways/{gatewayId}/bgpRoutes | Get gateway bgp routes


# **GetGatewayBGPRoutes**
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

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

