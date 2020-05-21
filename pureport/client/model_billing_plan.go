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

// BillingPlan A billing plan for a Pureport asset
type BillingPlan struct {
	// The recurring amount in cents that will be charged on the billing interval.
	Amount          float32         `json:"amount,omitempty"`
	BillingInterval BillingInterval `json:"billingInterval"`
	BillingLongHaul BillingLongHaul `json:"billingLongHaul,omitempty"`
	// The id is a unique identifier representing the billing plan.
	Id string `json:"id,omitempty"`
	// The nonrecurring amount in cents that will be charged on activation of the Pureport asset.
	SetupAmount float32     `json:"setupAmount,omitempty"`
	Term        BillingTerm `json:"term"`
}
