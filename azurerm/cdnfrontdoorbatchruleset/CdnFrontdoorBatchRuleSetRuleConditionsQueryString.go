// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package cdnfrontdoorbatchruleset


type CdnFrontdoorBatchRuleSetRuleConditionsQueryString struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/5.1.0/docs/resources/cdn_frontdoor_batch_rule_set#operator CdnFrontdoorBatchRuleSet#operator}.
	Operator *string `field:"required" json:"operator" yaml:"operator"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/5.1.0/docs/resources/cdn_frontdoor_batch_rule_set#transforms CdnFrontdoorBatchRuleSet#transforms}.
	Transforms *[]*string `field:"optional" json:"transforms" yaml:"transforms"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/5.1.0/docs/resources/cdn_frontdoor_batch_rule_set#values CdnFrontdoorBatchRuleSet#values}.
	Values *[]*string `field:"optional" json:"values" yaml:"values"`
}

