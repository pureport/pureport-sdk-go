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

// AuditEntryEventType the model 'AuditEntryEventType'
type AuditEntryEventType string

// List of AuditEntryEventType
const (
	USER_LOGIN                  AuditEntryEventType = "USER_LOGIN"
	USER_FORGOT_PASSWORD        AuditEntryEventType = "USER_FORGOT_PASSWORD"
	API_LOGIN                   AuditEntryEventType = "API_LOGIN"
	ACCOUNT_CREATE              AuditEntryEventType = "ACCOUNT_CREATE"
	ACCOUNT_UPDATE              AuditEntryEventType = "ACCOUNT_UPDATE"
	ACCOUNT_DELETE              AuditEntryEventType = "ACCOUNT_DELETE"
	ACCOUNT_BILLING_CREATE      AuditEntryEventType = "ACCOUNT_BILLING_CREATE"
	ACCOUNT_BILLING_UPDATE      AuditEntryEventType = "ACCOUNT_BILLING_UPDATE"
	ACCOUNT_BILLING_DELETE      AuditEntryEventType = "ACCOUNT_BILLING_DELETE"
	NETWORK_CREATE              AuditEntryEventType = "NETWORK_CREATE"
	NETWORK_UPDATE              AuditEntryEventType = "NETWORK_UPDATE"
	NETWORK_DELETE              AuditEntryEventType = "NETWORK_DELETE"
	CONNECTION_CREATE           AuditEntryEventType = "CONNECTION_CREATE"
	CONNECTION_UPDATE           AuditEntryEventType = "CONNECTION_UPDATE"
	CONNECTION_DELETE           AuditEntryEventType = "CONNECTION_DELETE"
	GATEWAY_CREATE              AuditEntryEventType = "GATEWAY_CREATE"
	GATEWAY_UPDATE              AuditEntryEventType = "GATEWAY_UPDATE"
	GATEWAY_DELETE              AuditEntryEventType = "GATEWAY_DELETE"
	API_KEY_CREATE              AuditEntryEventType = "API_KEY_CREATE"
	API_KEY_UPDATE              AuditEntryEventType = "API_KEY_UPDATE"
	API_KEY_DELETE              AuditEntryEventType = "API_KEY_DELETE"
	ROLE_CREATE                 AuditEntryEventType = "ROLE_CREATE"
	ROLE_UPDATE                 AuditEntryEventType = "ROLE_UPDATE"
	ROLE_DELETE                 AuditEntryEventType = "ROLE_DELETE"
	USER_CREATE                 AuditEntryEventType = "USER_CREATE"
	USER_UPDATE                 AuditEntryEventType = "USER_UPDATE"
	USER_DELETE                 AuditEntryEventType = "USER_DELETE"
	USER_DOMAIN_CREATE          AuditEntryEventType = "USER_DOMAIN_CREATE"
	USER_DOMAIN_UPDATE          AuditEntryEventType = "USER_DOMAIN_UPDATE"
	USER_DOMAIN_DELETE          AuditEntryEventType = "USER_DOMAIN_DELETE"
	PORT_CREATE                 AuditEntryEventType = "PORT_CREATE"
	PORT_UPDATE                 AuditEntryEventType = "PORT_UPDATE"
	PORT_DELETE                 AuditEntryEventType = "PORT_DELETE"
	MEMBER_INVITE_CREATE        AuditEntryEventType = "MEMBER_INVITE_CREATE"
	MEMBER_INVITE_ACCEPT        AuditEntryEventType = "MEMBER_INVITE_ACCEPT"
	MEMBER_INVITE_UPDATE        AuditEntryEventType = "MEMBER_INVITE_UPDATE"
	MEMBER_INVITE_DELETE        AuditEntryEventType = "MEMBER_INVITE_DELETE"
	ACCOUNT_MEMBER_CREATE       AuditEntryEventType = "ACCOUNT_MEMBER_CREATE"
	ACCOUNT_MEMBER_UPDATE       AuditEntryEventType = "ACCOUNT_MEMBER_UPDATE"
	ACCOUNT_MEMBER_DELETE       AuditEntryEventType = "ACCOUNT_MEMBER_DELETE"
	CONNECTION_STATE_CHANGE     AuditEntryEventType = "CONNECTION_STATE_CHANGE"
	GATEWAY_STATE_CHANGE        AuditEntryEventType = "GATEWAY_STATE_CHANGE"
	GATEWAY_BGP_STATUS_CHANGE   AuditEntryEventType = "GATEWAY_BGP_STATUS_CHANGE"
	GATEWAY_IPSEC_STATUS_CHANGE AuditEntryEventType = "GATEWAY_IPSEC_STATUS_CHANGE"
	NOTIFICATION_CREATE         AuditEntryEventType = "NOTIFICATION_CREATE"
	NOTIFICATION_UPDATE         AuditEntryEventType = "NOTIFICATION_UPDATE"
	NOTIFICATION_DELETE         AuditEntryEventType = "NOTIFICATION_DELETE"
	TASK_CREATE                 AuditEntryEventType = "TASK_CREATE"
	TASK_UPDATE                 AuditEntryEventType = "TASK_UPDATE"
	TASK_DELETE                 AuditEntryEventType = "TASK_DELETE"
)
