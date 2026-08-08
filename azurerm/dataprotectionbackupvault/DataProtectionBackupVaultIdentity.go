// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package dataprotectionbackupvault


type DataProtectionBackupVaultIdentity struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/5.0.1/docs/resources/data_protection_backup_vault#type DataProtectionBackupVault#type}.
	Type *string `field:"required" json:"type" yaml:"type"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/5.0.1/docs/resources/data_protection_backup_vault#identity_ids DataProtectionBackupVault#identity_ids}.
	IdentityIds *[]*string `field:"optional" json:"identityIds" yaml:"identityIds"`
}

