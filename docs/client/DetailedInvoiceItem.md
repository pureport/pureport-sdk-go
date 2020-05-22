# DetailedInvoiceItem

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Account** | [**Link**](Link.md) |  | [optional] 
**BillableObject** | [**Link**](Link.md) |  | [optional] 
**BillableObjectType** | **string** | The type of the Pureport entity billed for. | [optional] [readonly] 
**BillingPlan** | [**BillingPlan**](BillingPlan.md) |  | [optional] 
**Description** | **string** | The description of this line item. | [optional] [readonly] 
**Quantity** | **float32** | The amount billed for. | [optional] [readonly] 
**TimeRanges** | [**[]ClosedRangeInstant**](ClosedRangeInstant.md) | The time ranges billed for this Pureport entity. | [optional] [readonly] 

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


