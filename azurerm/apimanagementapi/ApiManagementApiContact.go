// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package apimanagementapi


type ApiManagementApiContact struct {
	// The email address of the contact person/organization.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/5.15.0/docs/resources/api_management_api#email ApiManagementApi#email}
	Email *string `field:"optional" json:"email" yaml:"email"`
	// The name of the contact person/organization.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/5.15.0/docs/resources/api_management_api#name ApiManagementApi#name}
	Name *string `field:"optional" json:"name" yaml:"name"`
	// Absolute URL of the contact information.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/5.15.0/docs/resources/api_management_api#url ApiManagementApi#url}
	Url *string `field:"optional" json:"url" yaml:"url"`
}

