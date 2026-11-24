// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package apimanagement


type ApiManagementDelegation struct {
	// Should subscription requests be delegated to an external URL?
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/5.15.0/docs/resources/api_management#subscriptions_enabled ApiManagement#subscriptions_enabled}
	SubscriptionsEnabled interface{} `field:"optional" json:"subscriptionsEnabled" yaml:"subscriptionsEnabled"`
	// The delegation URL.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/5.15.0/docs/resources/api_management#url ApiManagement#url}
	Url *string `field:"optional" json:"url" yaml:"url"`
	// Should user registration requests be delegated to an external URL?
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/5.15.0/docs/resources/api_management#user_registration_enabled ApiManagement#user_registration_enabled}
	UserRegistrationEnabled interface{} `field:"optional" json:"userRegistrationEnabled" yaml:"userRegistrationEnabled"`
	// A base64-encoded validation key to validate, that a request is coming from Azure API Management.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/5.15.0/docs/resources/api_management#validation_key ApiManagement#validation_key}
	ValidationKey *string `field:"optional" json:"validationKey" yaml:"validationKey"`
}

