// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package manageddevopspool


type ManagedDevopsPoolAzureDevopsOrganizationPermissionAdministratorAccount struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.76.0/docs/resources/managed_devops_pool#groups ManagedDevopsPool#groups}.
	Groups *[]*string `field:"optional" json:"groups" yaml:"groups"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.76.0/docs/resources/managed_devops_pool#users ManagedDevopsPool#users}.
	Users *[]*string `field:"optional" json:"users" yaml:"users"`
}

