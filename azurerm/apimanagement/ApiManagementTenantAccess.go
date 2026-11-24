// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package apimanagement


type ApiManagementTenantAccess struct {
	// Should the access to the management API be enabled?
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/5.15.0/docs/resources/api_management#enabled ApiManagement#enabled}
	Enabled interface{} `field:"required" json:"enabled" yaml:"enabled"`
}

