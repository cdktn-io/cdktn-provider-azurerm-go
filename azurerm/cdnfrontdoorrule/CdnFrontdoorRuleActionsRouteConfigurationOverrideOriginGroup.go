// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package cdnfrontdoorrule


type CdnFrontdoorRuleActionsRouteConfigurationOverrideOriginGroup struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/5.0.1/docs/resources/cdn_frontdoor_rule#cdn_frontdoor_origin_group_id CdnFrontdoorRule#cdn_frontdoor_origin_group_id}.
	CdnFrontdoorOriginGroupId *string `field:"required" json:"cdnFrontdoorOriginGroupId" yaml:"cdnFrontdoorOriginGroupId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/5.0.1/docs/resources/cdn_frontdoor_rule#forwarding_protocol CdnFrontdoorRule#forwarding_protocol}.
	ForwardingProtocol *string `field:"required" json:"forwardingProtocol" yaml:"forwardingProtocol"`
}

