// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package apimanagementdiagnostic

import (
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type ApiManagementDiagnosticConfig struct {
	// Experimental.
	Connection interface{} `field:"optional" json:"connection" yaml:"connection"`
	// Experimental.
	Count interface{} `field:"optional" json:"count" yaml:"count"`
	// Experimental.
	DependsOn *[]cdktn.ITerraformDependable `field:"optional" json:"dependsOn" yaml:"dependsOn"`
	// Experimental.
	ForEach cdktn.ITerraformIterator `field:"optional" json:"forEach" yaml:"forEach"`
	// Experimental.
	Lifecycle *cdktn.TerraformResourceLifecycle `field:"optional" json:"lifecycle" yaml:"lifecycle"`
	// Experimental.
	Provider cdktn.TerraformProvider `field:"optional" json:"provider" yaml:"provider"`
	// Experimental.
	Provisioners *[]interface{} `field:"optional" json:"provisioners" yaml:"provisioners"`
	// The ID of the target API Management Logger where the API Management Diagnostic should be saved.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/5.15.0/docs/resources/api_management_diagnostic#api_management_logger_id ApiManagementDiagnostic#api_management_logger_id}
	ApiManagementLoggerId *string `field:"required" json:"apiManagementLoggerId" yaml:"apiManagementLoggerId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/5.15.0/docs/resources/api_management_diagnostic#api_management_name ApiManagementDiagnostic#api_management_name}.
	ApiManagementName *string `field:"required" json:"apiManagementName" yaml:"apiManagementName"`
	// The diagnostic identifier for the API Management Service. At this time the.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/5.15.0/docs/resources/api_management_diagnostic#identifier ApiManagementDiagnostic#identifier}
	Identifier *string `field:"required" json:"identifier" yaml:"identifier"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/5.15.0/docs/resources/api_management_diagnostic#resource_group_name ApiManagementDiagnostic#resource_group_name}.
	ResourceGroupName *string `field:"required" json:"resourceGroupName" yaml:"resourceGroupName"`
	// Always log errors. Send telemetry if there is an erroneous condition, regardless of sampling settings.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/5.15.0/docs/resources/api_management_diagnostic#always_log_errors ApiManagementDiagnostic#always_log_errors}
	AlwaysLogErrors interface{} `field:"optional" json:"alwaysLogErrors" yaml:"alwaysLogErrors"`
	// backend_request block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/5.15.0/docs/resources/api_management_diagnostic#backend_request ApiManagementDiagnostic#backend_request}
	BackendRequest *ApiManagementDiagnosticBackendRequest `field:"optional" json:"backendRequest" yaml:"backendRequest"`
	// backend_response block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/5.15.0/docs/resources/api_management_diagnostic#backend_response ApiManagementDiagnostic#backend_response}
	BackendResponse *ApiManagementDiagnosticBackendResponse `field:"optional" json:"backendResponse" yaml:"backendResponse"`
	// frontend_request block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/5.15.0/docs/resources/api_management_diagnostic#frontend_request ApiManagementDiagnostic#frontend_request}
	FrontendRequest *ApiManagementDiagnosticFrontendRequest `field:"optional" json:"frontendRequest" yaml:"frontendRequest"`
	// frontend_response block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/5.15.0/docs/resources/api_management_diagnostic#frontend_response ApiManagementDiagnostic#frontend_response}
	FrontendResponse *ApiManagementDiagnosticFrontendResponse `field:"optional" json:"frontendResponse" yaml:"frontendResponse"`
	// The HTTP Correlation Protocol to use.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/5.15.0/docs/resources/api_management_diagnostic#http_correlation_protocol ApiManagementDiagnostic#http_correlation_protocol}
	HttpCorrelationProtocol *string `field:"optional" json:"httpCorrelationProtocol" yaml:"httpCorrelationProtocol"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/5.15.0/docs/resources/api_management_diagnostic#id ApiManagementDiagnostic#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// Log client IP address.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/5.15.0/docs/resources/api_management_diagnostic#log_client_ip ApiManagementDiagnostic#log_client_ip}
	LogClientIp interface{} `field:"optional" json:"logClientIp" yaml:"logClientIp"`
	// Whether custom metrics emitted via the `emit-metric` policy are sent to Application Insights.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/5.15.0/docs/resources/api_management_diagnostic#metrics_enabled ApiManagementDiagnostic#metrics_enabled}
	MetricsEnabled interface{} `field:"optional" json:"metricsEnabled" yaml:"metricsEnabled"`
	// The format of the Operation Name for Application Insights telemetries.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/5.15.0/docs/resources/api_management_diagnostic#operation_name_format ApiManagementDiagnostic#operation_name_format}
	OperationNameFormat *string `field:"optional" json:"operationNameFormat" yaml:"operationNameFormat"`
	// Sampling (%). For high traffic APIs, please read this [documentation](https://learn.microsoft.com/azure/api-management/api-management-howto-app-insights#performance-implications-and-log-sampling) to understand performance implications and log sampling.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/5.15.0/docs/resources/api_management_diagnostic#sampling_percentage ApiManagementDiagnostic#sampling_percentage}
	SamplingPercentage *float64 `field:"optional" json:"samplingPercentage" yaml:"samplingPercentage"`
	// timeouts block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/5.15.0/docs/resources/api_management_diagnostic#timeouts ApiManagementDiagnostic#timeouts}
	Timeouts *ApiManagementDiagnosticTimeouts `field:"optional" json:"timeouts" yaml:"timeouts"`
	// Logging verbosity.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/5.15.0/docs/resources/api_management_diagnostic#verbosity ApiManagementDiagnostic#verbosity}
	Verbosity *string `field:"optional" json:"verbosity" yaml:"verbosity"`
}

