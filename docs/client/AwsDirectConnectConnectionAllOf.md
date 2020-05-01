# AwsDirectConnectConnectionAllOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AwsAccountId** | **string** | The AWS Account ID, a 12-digit number that you use to construct Amazon Resource Names (ARNs). | [optional] 
**AwsRegion** | **string** | The AWS Region for the connection. This has been deprecated in favor of cloudRegion. | [optional] 
**BgpPasswordConfiguration** | [**BgpPasswordConfiguration**](BGPPasswordConfiguration.md) |  | [optional] 
**CloudRegion** | [**Link**](Link.md) |  | [optional] 
**CloudServices** | [**[]Link**](Link.md) | The asset link for the clouds services associated with this connection. | [optional] 
**CloudServicesPrefixWhitelist** | **[]string** | The whitelisted cloud services CIDR blocks. | [optional] 
**Peering** | [**PeeringConfiguration**](PeeringConfiguration.md) |  | [optional] 

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


