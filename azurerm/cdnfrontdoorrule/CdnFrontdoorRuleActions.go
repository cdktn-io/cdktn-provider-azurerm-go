// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package cdnfrontdoorrule


type CdnFrontdoorRuleActions struct {
	// modify_request_header block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/5.1.0/docs/resources/cdn_frontdoor_rule#modify_request_header CdnFrontdoorRule#modify_request_header}
	ModifyRequestHeader interface{} `field:"optional" json:"modifyRequestHeader" yaml:"modifyRequestHeader"`
	// modify_response_header block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/5.1.0/docs/resources/cdn_frontdoor_rule#modify_response_header CdnFrontdoorRule#modify_response_header}
	ModifyResponseHeader interface{} `field:"optional" json:"modifyResponseHeader" yaml:"modifyResponseHeader"`
	// route_configuration_override block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/5.1.0/docs/resources/cdn_frontdoor_rule#route_configuration_override CdnFrontdoorRule#route_configuration_override}
	RouteConfigurationOverride *CdnFrontdoorRuleActionsRouteConfigurationOverride `field:"optional" json:"routeConfigurationOverride" yaml:"routeConfigurationOverride"`
	// url_redirect block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/5.1.0/docs/resources/cdn_frontdoor_rule#url_redirect CdnFrontdoorRule#url_redirect}
	UrlRedirect *CdnFrontdoorRuleActionsUrlRedirect `field:"optional" json:"urlRedirect" yaml:"urlRedirect"`
	// url_rewrite block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/5.1.0/docs/resources/cdn_frontdoor_rule#url_rewrite CdnFrontdoorRule#url_rewrite}
	UrlRewrite *CdnFrontdoorRuleActionsUrlRewrite `field:"optional" json:"urlRewrite" yaml:"urlRewrite"`
}

