# pureport\client\ThirdPartyApi

All URIs are relative to *https://api.pureport.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetPacketFabric**](ThirdPartyApi.md#GetPacketFabric) | **Get** /thirdParty/packetFabric/customers/{marketCode} | 



## GetPacketFabric

> []PacketFabricCustomer GetPacketFabric(ctx, marketCode)



### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**marketCode** | **string**|  | 

### Return type

[**[]PacketFabricCustomer**](PacketFabricCustomer.md)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

