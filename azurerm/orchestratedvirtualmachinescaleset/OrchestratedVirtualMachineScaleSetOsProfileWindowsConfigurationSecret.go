// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package orchestratedvirtualmachinescaleset


type OrchestratedVirtualMachineScaleSetOsProfileWindowsConfigurationSecret struct {
	// certificate block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.59.0/docs/resources/orchestrated_virtual_machine_scale_set#certificate OrchestratedVirtualMachineScaleSet#certificate}
	Certificate interface{} `field:"required" json:"certificate" yaml:"certificate"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.59.0/docs/resources/orchestrated_virtual_machine_scale_set#key_vault_id OrchestratedVirtualMachineScaleSet#key_vault_id}.
	KeyVaultId *string `field:"required" json:"keyVaultId" yaml:"keyVaultId"`
}

