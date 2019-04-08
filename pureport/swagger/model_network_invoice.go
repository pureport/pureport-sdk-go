/*
 * Purport Control Plane
 *
 * Pureport API
 *
 * API version: 1.0.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package swagger

type NetworkInvoice struct {
	Account *Link `json:"account,omitempty"`
	// The Stripe Invoice object (https://stripe.com/docs/api/invoices/object)
	Invoice *interface{} `json:"invoice,omitempty"`
	Network *Link        `json:"network,omitempty"`
}
