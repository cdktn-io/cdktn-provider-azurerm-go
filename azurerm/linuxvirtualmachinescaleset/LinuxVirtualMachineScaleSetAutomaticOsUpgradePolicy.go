// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package linuxvirtualmachinescaleset


type LinuxVirtualMachineScaleSetAutomaticOsUpgradePolicy struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/5.0.1/docs/resources/linux_virtual_machine_scale_set#automatic_os_upgrade_enabled LinuxVirtualMachineScaleSet#automatic_os_upgrade_enabled}.
	AutomaticOsUpgradeEnabled interface{} `field:"required" json:"automaticOsUpgradeEnabled" yaml:"automaticOsUpgradeEnabled"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/5.0.1/docs/resources/linux_virtual_machine_scale_set#automatic_rollback_enabled LinuxVirtualMachineScaleSet#automatic_rollback_enabled}.
	AutomaticRollbackEnabled interface{} `field:"required" json:"automaticRollbackEnabled" yaml:"automaticRollbackEnabled"`
}

