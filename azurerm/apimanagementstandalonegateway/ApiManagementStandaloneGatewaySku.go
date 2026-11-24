// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package apimanagementstandalonegateway


type ApiManagementStandaloneGatewaySku struct {
	// The name of the SKU.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/5.15.0/docs/resources/api_management_standalone_gateway#name ApiManagementStandaloneGateway#name}
	Name *string `field:"required" json:"name" yaml:"name"`
	// The number of deployed units of the SKU.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/5.15.0/docs/resources/api_management_standalone_gateway#capacity ApiManagementStandaloneGateway#capacity}
	Capacity *float64 `field:"optional" json:"capacity" yaml:"capacity"`
}

