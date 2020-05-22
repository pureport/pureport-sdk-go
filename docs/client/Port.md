# Port

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Account** | [**Link**](Link.md) |  | 
**AvailabilityDomain** | [**AvailabilityDomain**](AvailabilityDomain.md) |  | 
**AvailableToChildAccounts** | **bool** | Whether this port is available for use in child accounts. | [optional] 
**BillingPlan** | [**BillingPlan**](BillingPlan.md) |  | [optional] 
**BillingTerm** | [**BillingTerm**](BillingTerm.md) |  | 
**Description** | **string** | The description. | [optional] 
**DeviceId** | **string** | The id of the device hosting this port. | [optional] [readonly] 
**Facility** | [**Link**](Link.md) |  | 
**Href** | **string** | The URI of the Pureport asset. | [optional] [readonly] 
**Id** | **string** | The id is a unique identifier representing the port. | [optional] 
**InterfaceId** | **string** | The id of the interface the port is using. | [optional] [readonly] 
**LoaCustomerName** | **string** | The customer name to appear on a Letter of Authorization. | [optional] 
**MediaType** | **string** | The media type for this port. | 
**Name** | **string** | The name. | 
**PatchPanelId** | **string** | The id of the patch panel the port is using. | [optional] [readonly] 
**PatchPanelPortNumbers** | **[]int32** | The ports on the patch panel the port is using. | [optional] [readonly] 
**Provider** | [**PortProvider**](PortProvider.md) |  | 
**Speed** | [**PortSpeed**](PortSpeed.md) |  | 
**State** | [**PortState**](PortState.md) |  | [optional] 
**Tags** | **map[string]string** | Key-value pairs to associate with the Pureport asset. | [optional] 

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


