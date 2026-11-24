// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package apimanagementbackend


type ApiManagementBackendCircuitBreakerRuleFailureCondition struct {
	// The time window over which failures are counted, in ISO 8601 format.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/5.15.0/docs/resources/api_management_backend#interval_duration ApiManagementBackend#interval_duration}
	IntervalDuration *string `field:"required" json:"intervalDuration" yaml:"intervalDuration"`
	// The number of failures within the specified interval that will trigger the circuit breaker.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/5.15.0/docs/resources/api_management_backend#count ApiManagementBackend#count}
	Count *float64 `field:"optional" json:"count" yaml:"count"`
	// A list of error reasons to consider as failures.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/5.15.0/docs/resources/api_management_backend#error_reasons ApiManagementBackend#error_reasons}
	ErrorReasons *[]*string `field:"optional" json:"errorReasons" yaml:"errorReasons"`
	// The percentage of failures within the specified interval that will trigger the circuit breaker.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/5.15.0/docs/resources/api_management_backend#percentage ApiManagementBackend#percentage}
	Percentage *float64 `field:"optional" json:"percentage" yaml:"percentage"`
	// status_code_range block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/5.15.0/docs/resources/api_management_backend#status_code_range ApiManagementBackend#status_code_range}
	StatusCodeRange interface{} `field:"optional" json:"statusCodeRange" yaml:"statusCodeRange"`
}

