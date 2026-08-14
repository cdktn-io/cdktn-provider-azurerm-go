// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package cdnfrontdoorbatchruleset


type CdnFrontdoorBatchRuleSetRuleActionsUrlRewrite struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/5.1.0/docs/resources/cdn_frontdoor_batch_rule_set#destination_path CdnFrontdoorBatchRuleSet#destination_path}.
	DestinationPath *string `field:"required" json:"destinationPath" yaml:"destinationPath"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/5.1.0/docs/resources/cdn_frontdoor_batch_rule_set#source_pattern CdnFrontdoorBatchRuleSet#source_pattern}.
	SourcePattern *string `field:"required" json:"sourcePattern" yaml:"sourcePattern"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/5.1.0/docs/resources/cdn_frontdoor_batch_rule_set#preserve_unmatched_path_enabled CdnFrontdoorBatchRuleSet#preserve_unmatched_path_enabled}.
	PreserveUnmatchedPathEnabled interface{} `field:"optional" json:"preserveUnmatchedPathEnabled" yaml:"preserveUnmatchedPathEnabled"`
}

