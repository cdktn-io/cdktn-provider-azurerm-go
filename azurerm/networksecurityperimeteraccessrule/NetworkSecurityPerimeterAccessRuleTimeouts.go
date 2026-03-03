// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package networksecurityperimeteraccessrule


type NetworkSecurityPerimeterAccessRuleTimeouts struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.62.1/docs/resources/network_security_perimeter_access_rule#create NetworkSecurityPerimeterAccessRule#create}.
	Create *string `field:"optional" json:"create" yaml:"create"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.62.1/docs/resources/network_security_perimeter_access_rule#delete NetworkSecurityPerimeterAccessRule#delete}.
	Delete *string `field:"optional" json:"delete" yaml:"delete"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.62.1/docs/resources/network_security_perimeter_access_rule#read NetworkSecurityPerimeterAccessRule#read}.
	Read *string `field:"optional" json:"read" yaml:"read"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.62.1/docs/resources/network_security_perimeter_access_rule#update NetworkSecurityPerimeterAccessRule#update}.
	Update *string `field:"optional" json:"update" yaml:"update"`
}

