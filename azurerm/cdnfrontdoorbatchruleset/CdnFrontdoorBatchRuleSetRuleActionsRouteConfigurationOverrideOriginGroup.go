// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package cdnfrontdoorbatchruleset


type CdnFrontdoorBatchRuleSetRuleActionsRouteConfigurationOverrideOriginGroup struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/5.2.0/docs/resources/cdn_frontdoor_batch_rule_set#cdn_frontdoor_origin_group_id CdnFrontdoorBatchRuleSet#cdn_frontdoor_origin_group_id}.
	CdnFrontdoorOriginGroupId *string `field:"required" json:"cdnFrontdoorOriginGroupId" yaml:"cdnFrontdoorOriginGroupId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/5.2.0/docs/resources/cdn_frontdoor_batch_rule_set#forwarding_protocol CdnFrontdoorBatchRuleSet#forwarding_protocol}.
	ForwardingProtocol *string `field:"required" json:"forwardingProtocol" yaml:"forwardingProtocol"`
}

