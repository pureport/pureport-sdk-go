# \AccountInvitationsApi

All URIs are relative to *http://api.pureport.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AcceptAllInvites**](AccountInvitationsApi.md#AcceptAllInvites) | **Post** /invites | Accept all invites for the current user
[**Delete7**](AccountInvitationsApi.md#Delete7) | **Delete** /accounts/{accountId}/invites/{inviteId} | Delete account invite
[**FindAccountInvites**](AccountInvitationsApi.md#FindAccountInvites) | **Get** /accounts/{accountId}/invites | List account invites
[**Get9**](AccountInvitationsApi.md#Get9) | **Get** /accounts/{accountId}/invites/{inviteId} | Get account invite
[**InviteAccount**](AccountInvitationsApi.md#InviteAccount) | **Post** /accounts/{accountId}/invites | Invite account member
[**Update7**](AccountInvitationsApi.md#Update7) | **Put** /accounts/{accountId}/invites/{inviteId} | Update account invite


# **AcceptAllInvites**
> []AccountInvite AcceptAllInvites(ctx, )
Accept all invites for the current user



### Required Parameters
This endpoint does not need any parameter.

### Return type

[**[]AccountInvite**](AccountInvite.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **Delete7**
> Delete7(ctx, inviteId, accountId)
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

No authorization required

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

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **Get9**
> AccountInvite Get9(ctx, inviteId, accountId)
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

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **InviteAccount**
> InviteAccount(ctx, accountId, optional)
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

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **Update7**
> AccountInvite Update7(ctx, inviteId, accountId, optional)
Update account invite



### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **inviteId** | **string**|  | 
  **accountId** | **string**|  | 
 **optional** | ***Update7Opts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a Update7Opts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **reinvite** | **optional.Bool**|  | 
 **body** | [**optional.Interface of AccountInvite**](AccountInvite.md)|  | 

### Return type

[**AccountInvite**](AccountInvite.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

