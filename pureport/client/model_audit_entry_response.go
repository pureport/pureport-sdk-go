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

// AuditEntryResponse The response to the request that triggered this entry.
type AuditEntryResponse struct {
	// The body of the response.
	Body string `json:"body,omitempty"`
	// Execution time in milliseconds.
	ExecutionTime int64 `json:"executionTime,omitempty"`
	// The HTTP status code of the response.
	StatusCode int32 `json:"statusCode,omitempty"`
	// The family of the HTTP status code.
	StatusCodeFamily string `json:"statusCodeFamily,omitempty"`
}
