// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package apimanagement


type ApiManagementSignIn struct {
	// Should anonymous users be redirected to the sign in page?
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/5.15.0/docs/resources/api_management#enabled ApiManagement#enabled}
	Enabled interface{} `field:"required" json:"enabled" yaml:"enabled"`
}

