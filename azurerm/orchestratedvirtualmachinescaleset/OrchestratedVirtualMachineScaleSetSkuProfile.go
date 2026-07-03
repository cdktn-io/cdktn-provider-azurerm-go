// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package orchestratedvirtualmachinescaleset


type OrchestratedVirtualMachineScaleSetSkuProfile struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.80.0/docs/resources/orchestrated_virtual_machine_scale_set#allocation_strategy OrchestratedVirtualMachineScaleSet#allocation_strategy}.
	AllocationStrategy *string `field:"required" json:"allocationStrategy" yaml:"allocationStrategy"`
	// virtual_machine_size block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.80.0/docs/resources/orchestrated_virtual_machine_scale_set#virtual_machine_size OrchestratedVirtualMachineScaleSet#virtual_machine_size}
	VirtualMachineSize interface{} `field:"optional" json:"virtualMachineSize" yaml:"virtualMachineSize"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.80.0/docs/resources/orchestrated_virtual_machine_scale_set#vm_sizes OrchestratedVirtualMachineScaleSet#vm_sizes}.
	VmSizes *[]*string `field:"optional" json:"vmSizes" yaml:"vmSizes"`
}

