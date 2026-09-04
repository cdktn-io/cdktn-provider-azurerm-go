// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package cdnfrontdoorbatchruleset


type CdnFrontdoorBatchRuleSetRuleActions struct {
	// modify_request_header block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/5.4.0/docs/resources/cdn_frontdoor_batch_rule_set#modify_request_header CdnFrontdoorBatchRuleSet#modify_request_header}
	ModifyRequestHeader interface{} `field:"optional" json:"modifyRequestHeader" yaml:"modifyRequestHeader"`
	// modify_response_header block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/5.4.0/docs/resources/cdn_frontdoor_batch_rule_set#modify_response_header CdnFrontdoorBatchRuleSet#modify_response_header}
	ModifyResponseHeader interface{} `field:"optional" json:"modifyResponseHeader" yaml:"modifyResponseHeader"`
	// route_configuration_override block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/5.4.0/docs/resources/cdn_frontdoor_batch_rule_set#route_configuration_override CdnFrontdoorBatchRuleSet#route_configuration_override}
	RouteConfigurationOverride *CdnFrontdoorBatchRuleSetRuleActionsRouteConfigurationOverride `field:"optional" json:"routeConfigurationOverride" yaml:"routeConfigurationOverride"`
	// url_redirect block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/5.4.0/docs/resources/cdn_frontdoor_batch_rule_set#url_redirect CdnFrontdoorBatchRuleSet#url_redirect}
	UrlRedirect *CdnFrontdoorBatchRuleSetRuleActionsUrlRedirect `field:"optional" json:"urlRedirect" yaml:"urlRedirect"`
	// url_rewrite block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/5.4.0/docs/resources/cdn_frontdoor_batch_rule_set#url_rewrite CdnFrontdoorBatchRuleSet#url_rewrite}
	UrlRewrite *CdnFrontdoorBatchRuleSetRuleActionsUrlRewrite `field:"optional" json:"urlRewrite" yaml:"urlRewrite"`
}

