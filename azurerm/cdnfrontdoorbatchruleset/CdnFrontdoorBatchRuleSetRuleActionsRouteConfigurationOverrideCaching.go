// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package cdnfrontdoorbatchruleset


type CdnFrontdoorBatchRuleSetRuleActionsRouteConfigurationOverrideCaching struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.81.0/docs/resources/cdn_frontdoor_batch_rule_set#behaviour CdnFrontdoorBatchRuleSet#behaviour}.
	Behaviour *string `field:"required" json:"behaviour" yaml:"behaviour"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.81.0/docs/resources/cdn_frontdoor_batch_rule_set#compression_enabled CdnFrontdoorBatchRuleSet#compression_enabled}.
	CompressionEnabled interface{} `field:"optional" json:"compressionEnabled" yaml:"compressionEnabled"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.81.0/docs/resources/cdn_frontdoor_batch_rule_set#duration CdnFrontdoorBatchRuleSet#duration}.
	Duration *string `field:"optional" json:"duration" yaml:"duration"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.81.0/docs/resources/cdn_frontdoor_batch_rule_set#query_string_behaviour CdnFrontdoorBatchRuleSet#query_string_behaviour}.
	QueryStringBehaviour *string `field:"optional" json:"queryStringBehaviour" yaml:"queryStringBehaviour"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.81.0/docs/resources/cdn_frontdoor_batch_rule_set#query_string_parameters CdnFrontdoorBatchRuleSet#query_string_parameters}.
	QueryStringParameters *[]*string `field:"optional" json:"queryStringParameters" yaml:"queryStringParameters"`
}

