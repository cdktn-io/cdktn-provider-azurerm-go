// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package manageddevopspool


type ManagedDevopsPoolIdentity struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.78.0/docs/resources/managed_devops_pool#identity_ids ManagedDevopsPool#identity_ids}.
	IdentityIds *[]*string `field:"required" json:"identityIds" yaml:"identityIds"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.78.0/docs/resources/managed_devops_pool#type ManagedDevopsPool#type}.
	Type *string `field:"required" json:"type" yaml:"type"`
}

