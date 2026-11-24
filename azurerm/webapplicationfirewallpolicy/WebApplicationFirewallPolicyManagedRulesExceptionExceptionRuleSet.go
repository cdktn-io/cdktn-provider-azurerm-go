// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package webapplicationfirewallpolicy


type WebApplicationFirewallPolicyManagedRulesExceptionExceptionRuleSet struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/5.15.0/docs/resources/web_application_firewall_policy#type WebApplicationFirewallPolicy#type}.
	Type *string `field:"required" json:"type" yaml:"type"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/5.15.0/docs/resources/web_application_firewall_policy#version WebApplicationFirewallPolicy#version}.
	Version *string `field:"required" json:"version" yaml:"version"`
	// rule_group block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/5.15.0/docs/resources/web_application_firewall_policy#rule_group WebApplicationFirewallPolicy#rule_group}
	RuleGroup interface{} `field:"optional" json:"ruleGroup" yaml:"ruleGroup"`
}

