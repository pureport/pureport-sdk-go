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

import (
	"time"
)

// Task struct for Task
type Task struct {
	Children    []Link    `json:"children,omitempty"`
	CompletedAt time.Time `json:"completedAt,omitempty"`
	CreatedAt   time.Time `json:"createdAt,omitempty"`
	Description string    `json:"description,omitempty"`
	Href        string    `json:"href,omitempty"`
	Id          string    `json:"id,omitempty"`
	Parent      Link      `json:"parent,omitempty"`
	Result      string    `json:"result,omitempty"`
	StartedAt   time.Time `json:"startedAt,omitempty"`
	State       string    `json:"state,omitempty"`
	Type        string    `json:"type"`
	UpdatedAt   time.Time `json:"updatedAt,omitempty"`
}
