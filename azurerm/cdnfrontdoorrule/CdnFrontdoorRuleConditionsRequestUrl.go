// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package cdnfrontdoorrule


type CdnFrontdoorRuleConditionsRequestUrl struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/5.0.1/docs/resources/cdn_frontdoor_rule#operator CdnFrontdoorRule#operator}.
	Operator *string `field:"required" json:"operator" yaml:"operator"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/5.0.1/docs/resources/cdn_frontdoor_rule#transforms CdnFrontdoorRule#transforms}.
	Transforms *[]*string `field:"optional" json:"transforms" yaml:"transforms"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/5.0.1/docs/resources/cdn_frontdoor_rule#values CdnFrontdoorRule#values}.
	Values *[]*string `field:"optional" json:"values" yaml:"values"`
}

