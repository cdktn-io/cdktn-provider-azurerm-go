// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package webapplicationfirewallpolicy


type WebApplicationFirewallPolicyPolicySettingsLogScrubbing struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.78.0/docs/resources/web_application_firewall_policy#enabled WebApplicationFirewallPolicy#enabled}.
	Enabled interface{} `field:"optional" json:"enabled" yaml:"enabled"`
	// rule block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.78.0/docs/resources/web_application_firewall_policy#rule WebApplicationFirewallPolicy#rule}
	Rule interface{} `field:"optional" json:"rule" yaml:"rule"`
}

