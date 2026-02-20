// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package networksecurityperimeterassociation


type NetworkSecurityPerimeterAssociationTimeouts struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.61.0/docs/resources/network_security_perimeter_association#create NetworkSecurityPerimeterAssociation#create}.
	Create *string `field:"optional" json:"create" yaml:"create"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.61.0/docs/resources/network_security_perimeter_association#delete NetworkSecurityPerimeterAssociation#delete}.
	Delete *string `field:"optional" json:"delete" yaml:"delete"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.61.0/docs/resources/network_security_perimeter_association#read NetworkSecurityPerimeterAssociation#read}.
	Read *string `field:"optional" json:"read" yaml:"read"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.61.0/docs/resources/network_security_perimeter_association#update NetworkSecurityPerimeterAssociation#update}.
	Update *string `field:"optional" json:"update" yaml:"update"`
}

