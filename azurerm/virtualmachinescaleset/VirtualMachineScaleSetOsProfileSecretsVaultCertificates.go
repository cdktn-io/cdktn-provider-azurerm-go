// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package virtualmachinescaleset


type VirtualMachineScaleSetOsProfileSecretsVaultCertificates struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.69.0/docs/resources/virtual_machine_scale_set#certificate_url VirtualMachineScaleSet#certificate_url}.
	CertificateUrl *string `field:"required" json:"certificateUrl" yaml:"certificateUrl"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.69.0/docs/resources/virtual_machine_scale_set#certificate_store VirtualMachineScaleSet#certificate_store}.
	CertificateStore *string `field:"optional" json:"certificateStore" yaml:"certificateStore"`
}

