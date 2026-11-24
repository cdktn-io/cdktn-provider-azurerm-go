// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package apimanagementbackend


type ApiManagementBackendCircuitBreakerRule struct {
	// failure_condition block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/5.15.0/docs/resources/api_management_backend#failure_condition ApiManagementBackend#failure_condition}
	FailureCondition *ApiManagementBackendCircuitBreakerRuleFailureCondition `field:"required" json:"failureCondition" yaml:"failureCondition"`
	// The name of the circuit breaker rule.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/5.15.0/docs/resources/api_management_backend#name ApiManagementBackend#name}
	Name *string `field:"required" json:"name" yaml:"name"`
	// The duration for which the circuit remains open before retrying, in ISO 8601 format.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/5.15.0/docs/resources/api_management_backend#trip_duration ApiManagementBackend#trip_duration}
	TripDuration *string `field:"required" json:"tripDuration" yaml:"tripDuration"`
	// Whether the circuit breaker should honour `Retry-After` requests.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/5.15.0/docs/resources/api_management_backend#accept_retry_after_enabled ApiManagementBackend#accept_retry_after_enabled}
	AcceptRetryAfterEnabled interface{} `field:"optional" json:"acceptRetryAfterEnabled" yaml:"acceptRetryAfterEnabled"`
}

