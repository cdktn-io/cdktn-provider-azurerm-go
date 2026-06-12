// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package linuxvirtualmachine


type LinuxVirtualMachineSecretCertificate struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.77.0/docs/resources/linux_virtual_machine#url LinuxVirtualMachine#url}.
	Url *string `field:"required" json:"url" yaml:"url"`
}

