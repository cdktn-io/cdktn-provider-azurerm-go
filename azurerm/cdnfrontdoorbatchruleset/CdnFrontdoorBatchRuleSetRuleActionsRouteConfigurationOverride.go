// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package cdnfrontdoorbatchruleset


type CdnFrontdoorBatchRuleSetRuleActionsRouteConfigurationOverride struct {
	// caching block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/5.4.0/docs/resources/cdn_frontdoor_batch_rule_set#caching CdnFrontdoorBatchRuleSet#caching}
	Caching *CdnFrontdoorBatchRuleSetRuleActionsRouteConfigurationOverrideCaching `field:"required" json:"caching" yaml:"caching"`
	// origin_group block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/5.4.0/docs/resources/cdn_frontdoor_batch_rule_set#origin_group CdnFrontdoorBatchRuleSet#origin_group}
	OriginGroup *CdnFrontdoorBatchRuleSetRuleActionsRouteConfigurationOverrideOriginGroup `field:"optional" json:"originGroup" yaml:"originGroup"`
}

