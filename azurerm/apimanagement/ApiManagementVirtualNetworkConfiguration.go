// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package apimanagement


type ApiManagementVirtualNetworkConfiguration struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.80.0/docs/resources/api_management#subnet_id ApiManagement#subnet_id}.
	SubnetId *string `field:"required" json:"subnetId" yaml:"subnetId"`
}

