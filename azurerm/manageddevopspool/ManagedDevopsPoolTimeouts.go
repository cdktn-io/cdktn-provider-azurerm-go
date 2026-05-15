// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package manageddevopspool


type ManagedDevopsPoolTimeouts struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.73.0/docs/resources/managed_devops_pool#create ManagedDevopsPool#create}.
	Create *string `field:"optional" json:"create" yaml:"create"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.73.0/docs/resources/managed_devops_pool#delete ManagedDevopsPool#delete}.
	Delete *string `field:"optional" json:"delete" yaml:"delete"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.73.0/docs/resources/managed_devops_pool#read ManagedDevopsPool#read}.
	Read *string `field:"optional" json:"read" yaml:"read"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.73.0/docs/resources/managed_devops_pool#update ManagedDevopsPool#update}.
	Update *string `field:"optional" json:"update" yaml:"update"`
}

