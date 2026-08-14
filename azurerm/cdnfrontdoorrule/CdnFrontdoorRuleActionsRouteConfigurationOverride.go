// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package cdnfrontdoorrule


type CdnFrontdoorRuleActionsRouteConfigurationOverride struct {
	// caching block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/5.1.0/docs/resources/cdn_frontdoor_rule#caching CdnFrontdoorRule#caching}
	Caching *CdnFrontdoorRuleActionsRouteConfigurationOverrideCaching `field:"required" json:"caching" yaml:"caching"`
	// origin_group block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/5.1.0/docs/resources/cdn_frontdoor_rule#origin_group CdnFrontdoorRule#origin_group}
	OriginGroup *CdnFrontdoorRuleActionsRouteConfigurationOverrideOriginGroup `field:"optional" json:"originGroup" yaml:"originGroup"`
}

