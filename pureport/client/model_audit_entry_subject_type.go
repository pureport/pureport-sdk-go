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

// AuditEntrySubjectType the model 'AuditEntrySubjectType'
type AuditEntrySubjectType string

// List of AuditEntrySubjectType
const (
	AUDIT_ENTRY_SUBJECT_TYPE__ACCOUNT         AuditEntrySubjectType = "ACCOUNT"
	AUDIT_ENTRY_SUBJECT_TYPE__CONNECTION      AuditEntrySubjectType = "CONNECTION"
	AUDIT_ENTRY_SUBJECT_TYPE__NETWORK         AuditEntrySubjectType = "NETWORK"
	AUDIT_ENTRY_SUBJECT_TYPE__USER            AuditEntrySubjectType = "USER"
	AUDIT_ENTRY_SUBJECT_TYPE__USER_DOMAIN     AuditEntrySubjectType = "USER_DOMAIN"
	AUDIT_ENTRY_SUBJECT_TYPE__ROLE            AuditEntrySubjectType = "ROLE"
	AUDIT_ENTRY_SUBJECT_TYPE__API_KEY         AuditEntrySubjectType = "API_KEY"
	AUDIT_ENTRY_SUBJECT_TYPE__GATEWAY         AuditEntrySubjectType = "GATEWAY"
	AUDIT_ENTRY_SUBJECT_TYPE__NOTIFICATION    AuditEntrySubjectType = "NOTIFICATION"
	AUDIT_ENTRY_SUBJECT_TYPE__ACCOUNT_INVITE  AuditEntrySubjectType = "ACCOUNT_INVITE"
	AUDIT_ENTRY_SUBJECT_TYPE__ACCOUNT_BILLING AuditEntrySubjectType = "ACCOUNT_BILLING"
	AUDIT_ENTRY_SUBJECT_TYPE__PORT            AuditEntrySubjectType = "PORT"
	AUDIT_ENTRY_SUBJECT_TYPE__ACCOUNT_MEMBER  AuditEntrySubjectType = "ACCOUNT_MEMBER"
	AUDIT_ENTRY_SUBJECT_TYPE__TASK            AuditEntrySubjectType = "TASK"
)
