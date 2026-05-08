// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package manageddevopspool


type ManagedDevopsPoolAzureDevopsOrganizationPermission struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.72.0/docs/resources/managed_devops_pool#kind ManagedDevopsPool#kind}.
	Kind *string `field:"required" json:"kind" yaml:"kind"`
	// administrator_account block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.72.0/docs/resources/managed_devops_pool#administrator_account ManagedDevopsPool#administrator_account}
	AdministratorAccount *ManagedDevopsPoolAzureDevopsOrganizationPermissionAdministratorAccount `field:"optional" json:"administratorAccount" yaml:"administratorAccount"`
}

