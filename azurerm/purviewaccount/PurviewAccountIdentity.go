// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package purviewaccount


type PurviewAccountIdentity struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.78.0/docs/resources/purview_account#type PurviewAccount#type}.
	Type *string `field:"required" json:"type" yaml:"type"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.78.0/docs/resources/purview_account#identity_ids PurviewAccount#identity_ids}.
	IdentityIds *[]*string `field:"optional" json:"identityIds" yaml:"identityIds"`
}

