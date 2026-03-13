// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package virtualmachineautomanageconfigurationassignment


type VirtualMachineAutomanageConfigurationAssignmentTimeouts struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.64.0/docs/resources/virtual_machine_automanage_configuration_assignment#create VirtualMachineAutomanageConfigurationAssignment#create}.
	Create *string `field:"optional" json:"create" yaml:"create"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.64.0/docs/resources/virtual_machine_automanage_configuration_assignment#delete VirtualMachineAutomanageConfigurationAssignment#delete}.
	Delete *string `field:"optional" json:"delete" yaml:"delete"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.64.0/docs/resources/virtual_machine_automanage_configuration_assignment#read VirtualMachineAutomanageConfigurationAssignment#read}.
	Read *string `field:"optional" json:"read" yaml:"read"`
}

