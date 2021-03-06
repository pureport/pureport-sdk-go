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

// TaskState the model 'TaskState'
type TaskState string

// List of TaskState
const (
	TASK_STATE__CREATED   TaskState = "CREATED"
	TASK_STATE__RUNNING   TaskState = "RUNNING"
	TASK_STATE__COMPLETED TaskState = "COMPLETED"
	TASK_STATE__FAILED    TaskState = "FAILED"
	TASK_STATE__DELETED   TaskState = "DELETED"
	TASK_STATE__AGGREGATE TaskState = "AGGREGATE"
)
