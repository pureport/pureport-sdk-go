# AwsDirectConnectConnection

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ActiveAt** | [**time.Time**](time.Time.md) |  | [optional] [readonly] 
**AdvertiseInternalRoutes** | **bool** | If the connection is advertising internal routes, which allows the customer the option of probing and tracing these routes. | [optional] 
**BillingPlan** | [**BillingPlan**](BillingPlan.md) |  | [optional] 
**BillingProvider** | [**BillingProvider**](BillingProvider.md) |  | [optional] 
**BillingTerm** | [**BillingTerm**](BillingTerm.md) |  | 
**CreatedAt** | [**time.Time**](time.Time.md) |  | [optional] [readonly] 
**CustomerASN** | **int64** | The customer side ASN. This can either be a public or private ASN. If this is a public ASN, you must own it to prevent conflicts. | [optional] 
**CustomerNetworks** | [**[]CustomerNetwork**](CustomerNetwork.md) | Set of customer Networks for this connection. | [optional] 
**DeletedAt** | [**time.Time**](time.Time.md) |  | [optional] [readonly] 
**Description** | **string** | The user defined description for the connection. | [optional] 
**ErrorCode** | [**ErrorCode**](ErrorCode.md) |  | [optional] 
**ErrorMessage** | **string** | Error message assigned to the connection if it is an error state. | [optional] [readonly] 
**HighAvailability** | **bool** | Whether this connection has redundant gateways for failover. | 
**Href** | **string** | The URI of the Pureport asset. | [optional] [readonly] 
**Id** | **string** | The id is a unique identifier representing the connection. This can be provided during creation, but if left empty, will be generated. | [optional] 
**Location** | [**Link**](Link.md) |  | 
**Name** | **string** | The user specified name for the connection. | 
**Nat** | [**NatConfig**](NATConfig.md) |  | [optional] 
**Network** | [**Link**](Link.md) |  | [optional] 
**PrimaryGateway** | [**Gateway**](Gateway.md) |  | [optional] 
**SecondaryGateway** | [**Gateway**](Gateway.md) |  | [optional] 
**Speed** | [**ConnectionSpeed**](ConnectionSpeed.md) |  | 
**State** | [**ConnectionState**](ConnectionState.md) |  | [optional] 
**Tags** | **map[string]string** | Key-value pairs to associate with the Pureport asset. | [optional] 
**Type** | [**ConnectionType**](ConnectionType.md) |  | 
**AwsAccountId** | **string** | The AWS Account ID, a 12-digit number that you use to construct Amazon Resource Names (ARNs). | 
**AwsRegion** | **string** | The AWS Region for the connection. This has been deprecated in favor of cloudRegion. | [optional] 
**BgpPasswordConfiguration** | [**BgpPasswordConfiguration**](BGPPasswordConfiguration.md) |  | [optional] 
**CloudRegion** | [**Link**](Link.md) |  | [optional] 
**CloudServices** | [**[]Link**](Link.md) | The asset link for the clouds services associated with this connection. | [optional] 
**CloudServicesPrefixWhitelist** | **[]string** | The whitelisted cloud services CIDR blocks. | [optional] 
**Peering** | [**PeeringConfiguration**](PeeringConfiguration.md) |  | [optional] 

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


