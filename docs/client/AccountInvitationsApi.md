# pureport\client\AccountInvitationsApi

All URIs are relative to *https://api.pureport.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AcceptAllInvites**](AccountInvitationsApi.md#AcceptAllInvites) | **Post** /invites | Accept all invites for the current user
[**DeleteAccountInvite**](AccountInvitationsApi.md#DeleteAccountInvite) | **Delete** /accounts/{accountId}/invites/{inviteId} | Delete account invite
[**FindAccountInvites**](AccountInvitationsApi.md#FindAccountInvites) | **Get** /accounts/{accountId}/invites | List account invites
[**GetAccountInvite**](AccountInvitationsApi.md#GetAccountInvite) | **Get** /accounts/{accountId}/invites/{inviteId} | Get account invite
[**InviteAccount**](AccountInvitationsApi.md#InviteAccount) | **Post** /accounts/{accountId}/invites | Invite account member
[**UpdateAccountInvite**](AccountInvitationsApi.md#UpdateAccountInvite) | **Put** /accounts/{accountId}/invites/{inviteId} | Update account invite


# **AcceptAllInvites**
> []AccountInvite AcceptAllInvites(ctx, )
Accept all invites for the current user



### Required Parameters
This endpoint does not need any parameter.

### Return type

[**[]AccountInvite**](AccountInvite.md)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteAccountInvite**
> DeleteAccountInvite(ctx, inviteId, accountId)
Delete account invite



### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **inviteId** | **string**|  | 
  **accountId** | **string**|  | 

### Return type

 (empty response body)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **FindAccountInvites**
> []AccountInvite FindAccountInvites(ctx, accountId)
List account invites



### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **accountId** | **string**|  | 

### Return type

[**[]AccountInvite**](AccountInvite.md)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetAccountInvite**
> AccountInvite GetAccountInvite(ctx, inviteId, accountId)
Get account invite



### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **inviteId** | **string**|  | 
  **accountId** | **string**|  | 

### Return type

[**AccountInvite**](AccountInvite.md)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **InviteAccount**
> AccountInvite InviteAccount(ctx, accountId, optional)
Invite account member



### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **accountId** | **string**|  | 
 **optional** | ***InviteAccountOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a InviteAccountOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**optional.Interface of AccountInvite**](AccountInvite.md)|  | 

### Return type

[**AccountInvite**](AccountInvite.md)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdateAccountInvite**
> AccountInvite UpdateAccountInvite(ctx, inviteId, accountId, optional)
Update account invite



### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **inviteId** | **string**|  | 
  **accountId** | **string**|  | 
 **optional** | ***UpdateAccountInviteOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a UpdateAccountInviteOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **reinvite** | **optional.Bool**|  | 
 **body** | [**optional.Interface of AccountInvite**](AccountInvite.md)|  | 

### Return type

[**AccountInvite**](AccountInvite.md)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

