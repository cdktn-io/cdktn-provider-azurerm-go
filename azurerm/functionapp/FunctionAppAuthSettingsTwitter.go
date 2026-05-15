// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package functionapp


type FunctionAppAuthSettingsTwitter struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.73.0/docs/resources/function_app#consumer_key FunctionApp#consumer_key}.
	ConsumerKey *string `field:"required" json:"consumerKey" yaml:"consumerKey"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.73.0/docs/resources/function_app#consumer_secret FunctionApp#consumer_secret}.
	ConsumerSecret *string `field:"required" json:"consumerSecret" yaml:"consumerSecret"`
}

