// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package cdnfrontdoorbatchruleset


type CdnFrontdoorBatchRuleSetRule struct {
	// actions block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/5.3.0/docs/resources/cdn_frontdoor_batch_rule_set#actions CdnFrontdoorBatchRuleSet#actions}
	Actions *CdnFrontdoorBatchRuleSetRuleActions `field:"required" json:"actions" yaml:"actions"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/5.3.0/docs/resources/cdn_frontdoor_batch_rule_set#name CdnFrontdoorBatchRuleSet#name}.
	Name *string `field:"required" json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/5.3.0/docs/resources/cdn_frontdoor_batch_rule_set#order CdnFrontdoorBatchRuleSet#order}.
	Order *float64 `field:"required" json:"order" yaml:"order"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/5.3.0/docs/resources/cdn_frontdoor_batch_rule_set#behaviour_on_match CdnFrontdoorBatchRuleSet#behaviour_on_match}.
	BehaviourOnMatch *string `field:"optional" json:"behaviourOnMatch" yaml:"behaviourOnMatch"`
	// conditions block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/5.3.0/docs/resources/cdn_frontdoor_batch_rule_set#conditions CdnFrontdoorBatchRuleSet#conditions}
	Conditions *CdnFrontdoorBatchRuleSetRuleConditions `field:"optional" json:"conditions" yaml:"conditions"`
}

