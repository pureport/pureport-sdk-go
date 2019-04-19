/*
 * Pureport Control Plane
 *
 * Pureport API
 *
 * API version: 1.0.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package swagger

import (
	"time"
)

type PublicPeeringPoolStats struct {
	Time           time.Time `json:"time,omitempty"`
	TotalAllocated int32     `json:"totalAllocated,omitempty"`
	TotalBlocks    int32     `json:"totalBlocks,omitempty"`
	TotalDeleted   int32     `json:"totalDeleted,omitempty"`
	TotalFree      int32     `json:"totalFree,omitempty"`
	TotalSubnets   int32     `json:"totalSubnets,omitempty"`
}
