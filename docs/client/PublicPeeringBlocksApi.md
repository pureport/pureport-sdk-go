# pureport\client\PublicPeeringBlocksApi

All URIs are relative to *https://api.pureport.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreatePublicPeeringBlock**](PublicPeeringBlocksApi.md#CreatePublicPeeringBlock) | **Post** /publicPeeringBlocks | Add Public Peering CIDR Blocks
[**DeletePublicPeeringBlock**](PublicPeeringBlocksApi.md#DeletePublicPeeringBlock) | **Delete** /publicPeeringBlocks/{publicPeeringBlockId} | Delete Public Peering CIDR
[**GetPublicPeeringBlock**](PublicPeeringBlocksApi.md#GetPublicPeeringBlock) | **Get** /publicPeeringBlocks/{publicPeeringBlockId} | Get Public Peering CIDR details
[**GetPublicPeeringBlocks**](PublicPeeringBlocksApi.md#GetPublicPeeringBlocks) | **Get** /publicPeeringBlocks | List Public Peering CIDR Blocks
[**UpdatePublicPeeringBlock**](PublicPeeringBlocksApi.md#UpdatePublicPeeringBlock) | **Put** /publicPeeringBlocks/{publicPeeringBlockId} | Update Public Peering CIDR



## CreatePublicPeeringBlock

> PublicPeeringBlock CreatePublicPeeringBlock(ctx, optional)

Add Public Peering CIDR Blocks

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***CreatePublicPeeringBlockOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a CreatePublicPeeringBlockOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **publicPeeringBlock** | [**optional.Interface of PublicPeeringBlock**](PublicPeeringBlock.md)|  | 

### Return type

[**PublicPeeringBlock**](PublicPeeringBlock.md)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeletePublicPeeringBlock

> DeletePublicPeeringBlock(ctx, publicPeeringBlockId)

Delete Public Peering CIDR

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**publicPeeringBlockId** | **string**|  | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetPublicPeeringBlock

> PublicPeeringBlock GetPublicPeeringBlock(ctx, publicPeeringBlockId)

Get Public Peering CIDR details

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**publicPeeringBlockId** | **string**|  | 

### Return type

[**PublicPeeringBlock**](PublicPeeringBlock.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetPublicPeeringBlocks

> []PublicPeeringBlock GetPublicPeeringBlocks(ctx, )

List Public Peering CIDR Blocks

### Required Parameters

This endpoint does not need any parameter.

### Return type

[**[]PublicPeeringBlock**](PublicPeeringBlock.md)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdatePublicPeeringBlock

> PublicPeeringBlock UpdatePublicPeeringBlock(ctx, publicPeeringBlockId, optional)

Update Public Peering CIDR

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**publicPeeringBlockId** | **string**|  | 
 **optional** | ***UpdatePublicPeeringBlockOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a UpdatePublicPeeringBlockOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **publicPeeringBlock** | [**optional.Interface of PublicPeeringBlock**](PublicPeeringBlock.md)|  | 

### Return type

[**PublicPeeringBlock**](PublicPeeringBlock.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

