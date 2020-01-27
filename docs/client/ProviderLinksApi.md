# pureport\client\ProviderLinksApi

All URIs are relative to *https://api.pureport.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateLink**](ProviderLinksApi.md#CreateLink) | **Post** /providerLinks | Add new provider link
[**DeleteProviderLink**](ProviderLinksApi.md#DeleteProviderLink) | **Delete** /providerLinks/{linkId} | Delete provider link
[**FindAllProviderLinks**](ProviderLinksApi.md#FindAllProviderLinks) | **Get** /providerLinks | List all provider links
[**GetProviderLink**](ProviderLinksApi.md#GetProviderLink) | **Get** /providerLinks/{linkId} | Get provider link details
[**GetVlans**](ProviderLinksApi.md#GetVlans) | **Get** /providerLinks/{linkId}/vlans | Get provider link VLANs
[**UpdateProviderLink**](ProviderLinksApi.md#UpdateProviderLink) | **Put** /providerLinks/{linkId} | Update provider link



## CreateLink

> ProviderLink CreateLink(ctx, optional)

Add new provider link

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***CreateLinkOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a CreateLinkOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **providerLink** | [**optional.Interface of ProviderLink**](ProviderLink.md)|  | 

### Return type

[**ProviderLink**](ProviderLink.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteProviderLink

> DeleteProviderLink(ctx, linkId)

Delete provider link

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**linkId** | **string**|  | 

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


## FindAllProviderLinks

> []ProviderLink FindAllProviderLinks(ctx, )

List all provider links

### Required Parameters

This endpoint does not need any parameter.

### Return type

[**[]ProviderLink**](ProviderLink.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetProviderLink

> ProviderLink GetProviderLink(ctx, linkId)

Get provider link details

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**linkId** | **string**|  | 

### Return type

[**ProviderLink**](ProviderLink.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetVlans

> []ProviderVlan GetVlans(ctx, linkId)

Get provider link VLANs

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**linkId** | **string**|  | 

### Return type

[**[]ProviderVlan**](ProviderVlan.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateProviderLink

> ProviderLink UpdateProviderLink(ctx, linkId, optional)

Update provider link

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**linkId** | **string**|  | 
 **optional** | ***UpdateProviderLinkOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a UpdateProviderLinkOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **providerLink** | [**optional.Interface of ProviderLink**](ProviderLink.md)|  | 

### Return type

[**ProviderLink**](ProviderLink.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

