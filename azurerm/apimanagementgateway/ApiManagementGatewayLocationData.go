// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package apimanagementgateway


type ApiManagementGatewayLocationData struct {
	// A canonical name for the geographic or physical location.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/5.15.0/docs/resources/api_management_gateway#name ApiManagementGateway#name}
	Name *string `field:"required" json:"name" yaml:"name"`
	// The city or locality where the resource is located.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/5.15.0/docs/resources/api_management_gateway#city ApiManagementGateway#city}
	City *string `field:"optional" json:"city" yaml:"city"`
	// The district, state, or province where the resource is located.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/5.15.0/docs/resources/api_management_gateway#district ApiManagementGateway#district}
	District *string `field:"optional" json:"district" yaml:"district"`
	// The country or region where the resource is located.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/5.15.0/docs/resources/api_management_gateway#region ApiManagementGateway#region}
	Region *string `field:"optional" json:"region" yaml:"region"`
}

