/*
 * Pureport Control Plane
 *
 * Pureport API
 *
 * API version: 1.0.0
 * Contact: support@pureport.com
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package client

// DetailedInvoiceItem Line item for an account detailed invoice.
type DetailedInvoiceItem struct {
	Account        Link `json:"account,omitempty"`
	BillableObject Link `json:"billableObject,omitempty"`
	// The type of the Pureport entity billed for.
	BillableObjectType string      `json:"billableObjectType,omitempty"`
	BillingPlan        BillingPlan `json:"billingPlan,omitempty"`
	// The description of this line item.
	Description string `json:"description,omitempty"`
	// The amount billed for.
	Quantity float32 `json:"quantity,omitempty"`
	// The time ranges billed for this Pureport entity.
	TimeRanges []ClosedRangeInstant `json:"timeRanges,omitempty"`
}
