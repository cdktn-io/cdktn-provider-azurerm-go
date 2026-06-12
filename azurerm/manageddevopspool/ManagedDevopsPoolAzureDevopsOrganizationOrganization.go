// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package manageddevopspool


type ManagedDevopsPoolAzureDevopsOrganizationOrganization struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.77.0/docs/resources/managed_devops_pool#parallelism ManagedDevopsPool#parallelism}.
	Parallelism *float64 `field:"required" json:"parallelism" yaml:"parallelism"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.77.0/docs/resources/managed_devops_pool#url ManagedDevopsPool#url}.
	Url *string `field:"required" json:"url" yaml:"url"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.77.0/docs/resources/managed_devops_pool#projects ManagedDevopsPool#projects}.
	Projects *[]*string `field:"optional" json:"projects" yaml:"projects"`
}

