// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package linuxwebapp


type LinuxWebAppTimeouts struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.64.0/docs/resources/linux_web_app#create LinuxWebApp#create}.
	Create *string `field:"optional" json:"create" yaml:"create"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.64.0/docs/resources/linux_web_app#delete LinuxWebApp#delete}.
	Delete *string `field:"optional" json:"delete" yaml:"delete"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.64.0/docs/resources/linux_web_app#read LinuxWebApp#read}.
	Read *string `field:"optional" json:"read" yaml:"read"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.64.0/docs/resources/linux_web_app#update LinuxWebApp#update}.
	Update *string `field:"optional" json:"update" yaml:"update"`
}

