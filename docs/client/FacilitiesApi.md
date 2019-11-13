# \FacilitiesApi

All URIs are relative to *https://api.pureport.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateFacility**](FacilitiesApi.md#CreateFacility) | **Post** /facilities | Add new facility
[**DeleteFacility**](FacilitiesApi.md#DeleteFacility) | **Delete** /facilities/{facilityId} | Delete facility
[**FindFacilities**](FacilitiesApi.md#FindFacilities) | **Get** /facilities | List facilities
[**GetFacility**](FacilitiesApi.md#GetFacility) | **Get** /facilities/{facilityId} | Get facility details
[**ListFacilitiesAsKml**](FacilitiesApi.md#ListFacilitiesAsKml) | **Get** /facilities/kml | List facilities (as KML)
[**UpdateFacility**](FacilitiesApi.md#UpdateFacility) | **Put** /facilities/{facilityId} | Update facility



## CreateFacility

> Facility CreateFacility(ctx, optional)

Add new facility

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***CreateFacilityOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a CreateFacilityOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **facility** | [**optional.Interface of Facility**](Facility.md)|  | 

### Return type

[**Facility**](Facility.md)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteFacility

> DeleteFacility(ctx, facilityId)

Delete facility

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**facilityId** | **string**|  | 

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


## FindFacilities

> []Facility FindFacilities(ctx, )

List facilities

### Required Parameters

This endpoint does not need any parameter.

### Return type

[**[]Facility**](Facility.md)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetFacility

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

[OAuth2](../README.md#OAuth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListFacilitiesAsKml

> string ListFacilitiesAsKml(ctx, )

List facilities (as KML)

### Required Parameters

This endpoint does not need any parameter.

### Return type

**string**

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/vnd.google-earth.kml+xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateFacility

> Facility UpdateFacility(ctx, facilityId, optional)

Update facility

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**facilityId** | **string**|  | 
 **optional** | ***UpdateFacilityOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a UpdateFacilityOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **facility** | [**optional.Interface of Facility**](Facility.md)|  | 

### Return type

[**Facility**](Facility.md)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

