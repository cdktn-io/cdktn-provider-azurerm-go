// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package apimanagementapi


type ApiManagementApiOpenidAuthentication struct {
	// OpenID Connect provider identifier. The name of an [OpenID Connect Provider](https://registry.terraform.io/providers/hashicorp/azurerm/latest/docs/resources/api_management_openid_connect_provider).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/5.15.0/docs/resources/api_management_api#openid_provider_name ApiManagementApi#openid_provider_name}
	OpenidProviderName *string `field:"required" json:"openidProviderName" yaml:"openidProviderName"`
	// How to send token to the server. A list of zero or more methods.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/5.15.0/docs/resources/api_management_api#bearer_token_sending_methods ApiManagementApi#bearer_token_sending_methods}
	BearerTokenSendingMethods *[]*string `field:"optional" json:"bearerTokenSendingMethods" yaml:"bearerTokenSendingMethods"`
}

