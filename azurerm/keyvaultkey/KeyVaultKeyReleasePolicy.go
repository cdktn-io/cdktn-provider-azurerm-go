// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package keyvaultkey


type KeyVaultKeyReleasePolicy struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.78.0/docs/resources/key_vault_key#json KeyVaultKey#json}.
	Json *string `field:"required" json:"json" yaml:"json"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.78.0/docs/resources/key_vault_key#immutable KeyVaultKey#immutable}.
	Immutable interface{} `field:"optional" json:"immutable" yaml:"immutable"`
}

