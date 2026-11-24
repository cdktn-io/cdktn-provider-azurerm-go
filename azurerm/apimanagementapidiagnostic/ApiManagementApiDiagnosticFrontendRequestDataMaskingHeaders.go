// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package apimanagementapidiagnostic


type ApiManagementApiDiagnosticFrontendRequestDataMaskingHeaders struct {
	// The data masking mode.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/5.15.0/docs/resources/api_management_api_diagnostic#mode ApiManagementApiDiagnostic#mode}
	Mode *string `field:"required" json:"mode" yaml:"mode"`
	// The name of the header or the query parameter to mask.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/5.15.0/docs/resources/api_management_api_diagnostic#value ApiManagementApiDiagnostic#value}
	Value *string `field:"required" json:"value" yaml:"value"`
}

