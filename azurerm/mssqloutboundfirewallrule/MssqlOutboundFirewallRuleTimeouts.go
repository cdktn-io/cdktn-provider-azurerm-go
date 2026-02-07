// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package mssqloutboundfirewallrule


type MssqlOutboundFirewallRuleTimeouts struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.59.0/docs/resources/mssql_outbound_firewall_rule#create MssqlOutboundFirewallRule#create}.
	Create *string `field:"optional" json:"create" yaml:"create"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.59.0/docs/resources/mssql_outbound_firewall_rule#delete MssqlOutboundFirewallRule#delete}.
	Delete *string `field:"optional" json:"delete" yaml:"delete"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.59.0/docs/resources/mssql_outbound_firewall_rule#read MssqlOutboundFirewallRule#read}.
	Read *string `field:"optional" json:"read" yaml:"read"`
}

