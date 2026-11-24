// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package apimanagementbackend


type ApiManagementBackendCredentialsAuthorization struct {
	// The authentication Parameter value.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/5.15.0/docs/resources/api_management_backend#parameter ApiManagementBackend#parameter}
	Parameter *string `field:"optional" json:"parameter" yaml:"parameter"`
	// The authentication Scheme name.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/5.15.0/docs/resources/api_management_backend#scheme ApiManagementBackend#scheme}
	Scheme *string `field:"optional" json:"scheme" yaml:"scheme"`
}

