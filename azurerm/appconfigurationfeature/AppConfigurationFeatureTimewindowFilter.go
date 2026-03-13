// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package appconfigurationfeature


type AppConfigurationFeatureTimewindowFilter struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.64.0/docs/resources/app_configuration_feature#end AppConfigurationFeature#end}.
	End *string `field:"optional" json:"end" yaml:"end"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.64.0/docs/resources/app_configuration_feature#start AppConfigurationFeature#start}.
	Start *string `field:"optional" json:"start" yaml:"start"`
}

