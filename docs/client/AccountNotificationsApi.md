# pureport\client\AccountNotificationsApi

All URIs are relative to *https://api.pureport.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**FindAccountNotifications**](AccountNotificationsApi.md#FindAccountNotifications) | **Get** /accounts/{accountId}/notifications | Retrieve all notifications for this account
[**UpdateAccountNotificationsStatus**](AccountNotificationsApi.md#UpdateAccountNotificationsStatus) | **Post** /accounts/{accountId}/notifications | Update an authenticated user&#39;s read status for account notifications



## FindAccountNotifications

> []AccountNotification FindAccountNotifications(ctx, accountId, optional)

Retrieve all notifications for this account

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**accountId** | **string**|  | 
 **optional** | ***FindAccountNotificationsOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a FindAccountNotificationsOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **startTime** | **optional.Time**|  | 
 **endTime** | **optional.Time**|  | 
 **status** | [**optional.Interface of AccountNotificationStatus**](.md)|  | 

### Return type

[**[]AccountNotification**](AccountNotification.md)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateAccountNotificationsStatus

> UpdateAccountNotificationsStatus(ctx, accountId, optional)

Update an authenticated user's read status for account notifications

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**accountId** | **string**|  | 
 **optional** | ***UpdateAccountNotificationsStatusOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a UpdateAccountNotificationsStatusOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **requestBody** | [**optional.Interface of map[string]AccountNotificationStatus**](AccountNotificationStatus.md)|  | 

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

