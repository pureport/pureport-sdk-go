/*
 * Purport Control Plane
 *
 * Pureport API
 *
 * API version: 1.0.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package swagger

type AccountBilling struct {
	Account        *Link  `json:"account,omitempty"`
	Email          string `json:"email"`
	Href           string `json:"href,omitempty"`
	StripeExpiry   string `json:"stripeExpiry"`
	StripeLastFour string `json:"stripeLastFour"`
	StripeToken    string `json:"stripeToken"`
}
