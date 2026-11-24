// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package webapplicationfirewallpolicy


type WebApplicationFirewallPolicyManagedRulesException struct {
	// exception_rule_set block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/5.15.0/docs/resources/web_application_firewall_policy#exception_rule_set WebApplicationFirewallPolicy#exception_rule_set}
	ExceptionRuleSet interface{} `field:"required" json:"exceptionRuleSet" yaml:"exceptionRuleSet"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/5.15.0/docs/resources/web_application_firewall_policy#match_variable WebApplicationFirewallPolicy#match_variable}.
	MatchVariable *string `field:"required" json:"matchVariable" yaml:"matchVariable"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/5.15.0/docs/resources/web_application_firewall_policy#value_match_operator WebApplicationFirewallPolicy#value_match_operator}.
	ValueMatchOperator *string `field:"required" json:"valueMatchOperator" yaml:"valueMatchOperator"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/5.15.0/docs/resources/web_application_firewall_policy#values WebApplicationFirewallPolicy#values}.
	Values *[]*string `field:"required" json:"values" yaml:"values"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/5.15.0/docs/resources/web_application_firewall_policy#selector WebApplicationFirewallPolicy#selector}.
	Selector *string `field:"optional" json:"selector" yaml:"selector"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/5.15.0/docs/resources/web_application_firewall_policy#selector_match_operator WebApplicationFirewallPolicy#selector_match_operator}.
	SelectorMatchOperator *string `field:"optional" json:"selectorMatchOperator" yaml:"selectorMatchOperator"`
}

