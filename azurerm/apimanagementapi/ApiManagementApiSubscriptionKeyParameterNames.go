// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package apimanagementapi


type ApiManagementApiSubscriptionKeyParameterNames struct {
	// The name of the HTTP Header which should be used for the Subscription Key.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/5.15.0/docs/resources/api_management_api#header ApiManagementApi#header}
	Header *string `field:"required" json:"header" yaml:"header"`
	// The name of the QueryString parameter which should be used for the Subscription Key.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/5.15.0/docs/resources/api_management_api#query ApiManagementApi#query}
	Query *string `field:"required" json:"query" yaml:"query"`
}

