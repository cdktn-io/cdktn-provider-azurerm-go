// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package paloaltolocalrulestack


type PaloAltoLocalRulestackTimeouts struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/5.4.0/docs/resources/palo_alto_local_rulestack#create PaloAltoLocalRulestack#create}.
	Create *string `field:"optional" json:"create" yaml:"create"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/5.4.0/docs/resources/palo_alto_local_rulestack#delete PaloAltoLocalRulestack#delete}.
	Delete *string `field:"optional" json:"delete" yaml:"delete"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/5.4.0/docs/resources/palo_alto_local_rulestack#read PaloAltoLocalRulestack#read}.
	Read *string `field:"optional" json:"read" yaml:"read"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/5.4.0/docs/resources/palo_alto_local_rulestack#update PaloAltoLocalRulestack#update}.
	Update *string `field:"optional" json:"update" yaml:"update"`
}

