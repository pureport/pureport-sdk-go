# SupportedConnection

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**BillingPlans** | [**[]BillingPlan**](BillingPlan.md) | Billing plans available for this configuration. | [optional] [readonly] 
**BillingProductIds** | **map[string]string** | Association between billing providers and the product to bill under for this configuration. | 
**Groups** | [**[]Link**](Link.md) | The groups this connection configuration belongs to. | [optional] 
**HighAvailability** | **bool** | Whether connections using this configuration have redundant gateways for failover. | 
**Href** | **string** | The URI of the Pureport asset. | [optional] [readonly] 
**Id** | **string** | The id is a unique identifier representing the supported connection configuration. | [optional] 
**Location** | [**Link**](Link.md) |  | 
**PeeringType** | **string** | The peering type of this configuration. | 
**Pending** | **bool** | Whether this configuration is currently available or requires additional steps from Pureport. | 
**ReachableCloudRegions** | [**[]Link**](Link.md) | The cloud regions this configuration allows. | [optional] 
**RequiresApproval** | **bool** | Whether connections using this configuration require approval from Pureport to provision. | 
**Speed** | **int32** | The connection speed in Mbps of this configuration. | 
**Type** | **string** | The connection type of this configuration. | 

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


