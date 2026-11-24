// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package apimanagementbackend


type ApiManagementBackendProxy struct {
	// The URL of the proxy server.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/5.15.0/docs/resources/api_management_backend#url ApiManagementBackend#url}
	Url *string `field:"required" json:"url" yaml:"url"`
	// The password to connect to the proxy server.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/5.15.0/docs/resources/api_management_backend#password ApiManagementBackend#password}
	Password *string `field:"optional" json:"password" yaml:"password"`
	// The username to connect to the proxy server.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/5.15.0/docs/resources/api_management_backend#username ApiManagementBackend#username}
	Username *string `field:"optional" json:"username" yaml:"username"`
}

