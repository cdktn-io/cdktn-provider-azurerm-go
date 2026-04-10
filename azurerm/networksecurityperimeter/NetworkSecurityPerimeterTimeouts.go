// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package networksecurityperimeter


type NetworkSecurityPerimeterTimeouts struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.68.0/docs/resources/network_security_perimeter#create NetworkSecurityPerimeter#create}.
	Create *string `field:"optional" json:"create" yaml:"create"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.68.0/docs/resources/network_security_perimeter#delete NetworkSecurityPerimeter#delete}.
	Delete *string `field:"optional" json:"delete" yaml:"delete"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.68.0/docs/resources/network_security_perimeter#read NetworkSecurityPerimeter#read}.
	Read *string `field:"optional" json:"read" yaml:"read"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.68.0/docs/resources/network_security_perimeter#update NetworkSecurityPerimeter#update}.
	Update *string `field:"optional" json:"update" yaml:"update"`
}

