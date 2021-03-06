# PacketFabricConnection

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
**BgpPasswordConfiguration** | [**BgpPasswordConfiguration**](BGPPasswordConfiguration.md) |  | [optional] 
**CloudRegion** | [**Link**](Link.md) |  | 
**Customer** | [**PacketFabricCustomer**](PacketFabricCustomer.md) |  | 
**CustomerPrimaryId** | **string** | An ID representing the primary connection with the 3rd party provider selected in the “customer” field. | 
**CustomerSecondaryId** | **string** | An ID representing the secondary connection with the 3rd party provider selected in the “customer” field. | [optional] 
**Peering** | [**PeeringConfiguration**](PeeringConfiguration.md) |  | 
**PrimaryVlan** | **int32** | The primary VLAN ID. | [optional] 
**SecondaryVlan** | **int32** | The secondary VLAN ID if this is an HA connection. | [optional] 

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


