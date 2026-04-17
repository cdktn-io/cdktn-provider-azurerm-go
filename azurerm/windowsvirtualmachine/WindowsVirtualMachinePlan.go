// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package windowsvirtualmachine


type WindowsVirtualMachinePlan struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.69.0/docs/resources/windows_virtual_machine#name WindowsVirtualMachine#name}.
	Name *string `field:"required" json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.69.0/docs/resources/windows_virtual_machine#product WindowsVirtualMachine#product}.
	Product *string `field:"required" json:"product" yaml:"product"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.69.0/docs/resources/windows_virtual_machine#publisher WindowsVirtualMachine#publisher}.
	Publisher *string `field:"required" json:"publisher" yaml:"publisher"`
}

