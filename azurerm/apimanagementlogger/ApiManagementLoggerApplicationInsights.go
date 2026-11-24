// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package apimanagementlogger


type ApiManagementLoggerApplicationInsights struct {
	// The connection string of Application Insights.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/5.15.0/docs/resources/api_management_logger#connection_string ApiManagementLogger#connection_string}
	ConnectionString *string `field:"optional" json:"connectionString" yaml:"connectionString"`
	// The Client ID of the User Assigned Identity, or `SystemAssigned` to use the System Assigned Identity, that has the "Monitoring Metrics Publisher" role on the target Application Insights resource.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/5.15.0/docs/resources/api_management_logger#identity_client_id ApiManagementLogger#identity_client_id}
	IdentityClientId *string `field:"optional" json:"identityClientId" yaml:"identityClientId"`
	// The instrumentation key used to push data to Application Insights.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/5.15.0/docs/resources/api_management_logger#instrumentation_key ApiManagementLogger#instrumentation_key}
	InstrumentationKey *string `field:"optional" json:"instrumentationKey" yaml:"instrumentationKey"`
}

