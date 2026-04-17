// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package orbitalcontact


type OrbitalContactTimeouts struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.69.0/docs/resources/orbital_contact#create OrbitalContact#create}.
	Create *string `field:"optional" json:"create" yaml:"create"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.69.0/docs/resources/orbital_contact#delete OrbitalContact#delete}.
	Delete *string `field:"optional" json:"delete" yaml:"delete"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.69.0/docs/resources/orbital_contact#read OrbitalContact#read}.
	Read *string `field:"optional" json:"read" yaml:"read"`
}

