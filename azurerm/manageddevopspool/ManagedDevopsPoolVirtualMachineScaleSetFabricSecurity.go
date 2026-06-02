// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package manageddevopspool


type ManagedDevopsPoolVirtualMachineScaleSetFabricSecurity struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.75.0/docs/resources/managed_devops_pool#interactive_logon_enabled ManagedDevopsPool#interactive_logon_enabled}.
	InteractiveLogonEnabled interface{} `field:"optional" json:"interactiveLogonEnabled" yaml:"interactiveLogonEnabled"`
	// key_vault_management block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.75.0/docs/resources/managed_devops_pool#key_vault_management ManagedDevopsPool#key_vault_management}
	KeyVaultManagement *ManagedDevopsPoolVirtualMachineScaleSetFabricSecurityKeyVaultManagement `field:"optional" json:"keyVaultManagement" yaml:"keyVaultManagement"`
}

