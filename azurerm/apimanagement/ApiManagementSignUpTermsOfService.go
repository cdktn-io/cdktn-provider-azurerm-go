// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package apimanagement


type ApiManagementSignUpTermsOfService struct {
	// Should the user be asked for consent during sign up?
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/5.15.0/docs/resources/api_management#consent_required ApiManagement#consent_required}
	ConsentRequired interface{} `field:"required" json:"consentRequired" yaml:"consentRequired"`
	// Should Terms of Service be displayed during sign up?
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/5.15.0/docs/resources/api_management#enabled ApiManagement#enabled}
	Enabled interface{} `field:"required" json:"enabled" yaml:"enabled"`
	// The Terms of Service which users are required to agree to in order to sign up.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/5.15.0/docs/resources/api_management#text ApiManagement#text}
	Text *string `field:"optional" json:"text" yaml:"text"`
}

