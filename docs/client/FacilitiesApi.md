# pureport\client\FacilitiesApi

All URIs are relative to *https://api.pureport.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**FindFacilities**](FacilitiesApi.md#FindFacilities) | **Get** /facilities | List facilities
[**GetFacility**](FacilitiesApi.md#GetFacility) | **Get** /facilities/{facilityId} | Get facility details
[**ListFacilitiesAsKml**](FacilitiesApi.md#ListFacilitiesAsKml) | **Get** /facilities/kml | List facilities (as KML)


# **FindFacilities**
> []Facility FindFacilities(ctx, )
List facilities



### Required Parameters
This endpoint does not need any parameter.

### Return type

[**[]Facility**](Facility.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetFacility**
> Facility GetFacility(ctx, facilityId)
Get facility details



### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **facilityId** | **string**|  | 

### Return type

[**Facility**](Facility.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ListFacilitiesAsKml**
> ListFacilitiesAsKml(ctx, )
List facilities (as KML)



### Required Parameters
This endpoint does not need any parameter.

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/vnd.google-earth.kml+xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

