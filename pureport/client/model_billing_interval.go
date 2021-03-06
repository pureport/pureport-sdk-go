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

// BillingInterval The interval on which recurring charges will be charged.
type BillingInterval string

// List of BillingInterval
const (
	BILLING_INTERVAL__ONCE  BillingInterval = "ONCE"
	BILLING_INTERVAL__DAY   BillingInterval = "DAY"
	BILLING_INTERVAL__WEEK  BillingInterval = "WEEK"
	BILLING_INTERVAL__MONTH BillingInterval = "MONTH"
	BILLING_INTERVAL__YEAR  BillingInterval = "YEAR"
)
