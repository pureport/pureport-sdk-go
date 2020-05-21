# AccountInvite

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Account** | [**Link**](Link.md) |  | 
**Email** | **string** | The email address of the user being invited to the Pureport account. | 
**Expired** | **bool** |  | [optional] [readonly] 
**Href** | **string** | The URI of the Pureport asset. | [optional] [readonly] 
**Id** | **string** | The id is a unique identifier representing the invite. | [optional] 
**InvitedAt** | [**time.Time**](time.Time.md) |  | [optional] 
**InvitedBy** | [**Link**](Link.md) |  | [optional] 
**Roles** | [**[]Link**](Link.md) | The Pureport roles the invited user will have permissions for on the account. | 

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


