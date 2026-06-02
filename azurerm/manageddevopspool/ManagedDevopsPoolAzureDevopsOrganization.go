// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package manageddevopspool


type ManagedDevopsPoolAzureDevopsOrganization struct {
	// organization block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.75.0/docs/resources/managed_devops_pool#organization ManagedDevopsPool#organization}
	Organization interface{} `field:"required" json:"organization" yaml:"organization"`
	// permission block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.75.0/docs/resources/managed_devops_pool#permission ManagedDevopsPool#permission}
	Permission *ManagedDevopsPoolAzureDevopsOrganizationPermission `field:"optional" json:"permission" yaml:"permission"`
}

