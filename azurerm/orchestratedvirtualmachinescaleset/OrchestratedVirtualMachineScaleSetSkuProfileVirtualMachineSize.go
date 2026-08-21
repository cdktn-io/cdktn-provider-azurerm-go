// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package orchestratedvirtualmachinescaleset


type OrchestratedVirtualMachineScaleSetSkuProfileVirtualMachineSize struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/5.2.0/docs/resources/orchestrated_virtual_machine_scale_set#name OrchestratedVirtualMachineScaleSet#name}.
	Name *string `field:"required" json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/5.2.0/docs/resources/orchestrated_virtual_machine_scale_set#rank OrchestratedVirtualMachineScaleSet#rank}.
	Rank *float64 `field:"optional" json:"rank" yaml:"rank"`
}

