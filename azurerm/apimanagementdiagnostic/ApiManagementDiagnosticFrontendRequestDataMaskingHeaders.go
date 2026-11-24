// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package apimanagementdiagnostic


type ApiManagementDiagnosticFrontendRequestDataMaskingHeaders struct {
	// The data masking mode.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/5.15.0/docs/resources/api_management_diagnostic#mode ApiManagementDiagnostic#mode}
	Mode *string `field:"required" json:"mode" yaml:"mode"`
	// The name of the header or the query parameter to mask.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/5.15.0/docs/resources/api_management_diagnostic#value ApiManagementDiagnostic#value}
	Value *string `field:"required" json:"value" yaml:"value"`
}

