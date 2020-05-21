# Account

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Demo** | **bool** | Whether this account is for demonstration purposes. | [optional] 
**Description** | **string** | The description. | [optional] 
**HasChildren** | **bool** | Whether this account is a parent account for any other accounts. | [optional] [readonly] 
**Href** | **string** | The URI of the Pureport asset. | [optional] [readonly] 
**Id** | **string** | The id is a unique identifier representing the account. | [optional] 
**Name** | **string** | The name. | 
**Parent** | [**Link**](Link.md) |  | [optional] 
**PricingHidden** | **bool** | Whether pricing information is restricted on this account. | [optional] [readonly] 
**ShowChildAccountPricing** | **bool** | Whether to show pricing information to child accounts of this account. | [optional] 
**SupportedConnectionGroups** | [**[]Link**](Link.md) | A collection of asset links for which Supported Connection Groups this account has access to. | [optional] 
**Tags** | **map[string]string** | Key-value pairs to associate with the Pureport asset. | [optional] 
**TechnicalContactEmails** | **[]string** | Email addresses of technical contacts for this account. | [optional] 
**Verified** | **bool** | Whether this account has been verified by Pureport operations. | [optional] 

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


