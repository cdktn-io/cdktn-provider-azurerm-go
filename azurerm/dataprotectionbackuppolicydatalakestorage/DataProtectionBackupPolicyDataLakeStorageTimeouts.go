// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package dataprotectionbackuppolicydatalakestorage


type DataProtectionBackupPolicyDataLakeStorageTimeouts struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.71.0/docs/resources/data_protection_backup_policy_data_lake_storage#create DataProtectionBackupPolicyDataLakeStorage#create}.
	Create *string `field:"optional" json:"create" yaml:"create"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.71.0/docs/resources/data_protection_backup_policy_data_lake_storage#delete DataProtectionBackupPolicyDataLakeStorage#delete}.
	Delete *string `field:"optional" json:"delete" yaml:"delete"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.71.0/docs/resources/data_protection_backup_policy_data_lake_storage#read DataProtectionBackupPolicyDataLakeStorage#read}.
	Read *string `field:"optional" json:"read" yaml:"read"`
}

