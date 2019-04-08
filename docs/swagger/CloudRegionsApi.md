# \CloudRegionsApi

All URIs are relative to *http://localhost:8080*

Method | HTTP request | Description
------------- | ------------- | -------------
[**Get10**](CloudRegionsApi.md#Get10) | **Get** /cloudRegions/{cloudRegionId} | Get cloud region details
[**GetCloudRegions**](CloudRegionsApi.md#GetCloudRegions) | **Get** /cloudRegions | List cloud regions


# **Get10**
> CloudRegion Get10(ctx, cloudRegionId)
Get cloud region details



### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **cloudRegionId** | **string**|  | 

### Return type

[**CloudRegion**](CloudRegion.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetCloudRegions**
> []CloudRegion GetCloudRegions(ctx, )
List cloud regions



### Required Parameters
This endpoint does not need any parameter.

### Return type

[**[]CloudRegion**](CloudRegion.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

