// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package dataprotectionbackuppolicyblobstorage


type DataProtectionBackupPolicyBlobStorageRetentionRuleLifeCycle struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.75.0/docs/resources/data_protection_backup_policy_blob_storage#data_store_type DataProtectionBackupPolicyBlobStorage#data_store_type}.
	DataStoreType *string `field:"required" json:"dataStoreType" yaml:"dataStoreType"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.75.0/docs/resources/data_protection_backup_policy_blob_storage#duration DataProtectionBackupPolicyBlobStorage#duration}.
	Duration *string `field:"required" json:"duration" yaml:"duration"`
}

