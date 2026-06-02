// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package dataprotectionbackupinstancedatalakestorage

import (
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type DataProtectionBackupInstanceDataLakeStorageConfig struct {
	// Experimental.
	Connection interface{} `field:"optional" json:"connection" yaml:"connection"`
	// Experimental.
	Count interface{} `field:"optional" json:"count" yaml:"count"`
	// Experimental.
	DependsOn *[]cdktn.ITerraformDependable `field:"optional" json:"dependsOn" yaml:"dependsOn"`
	// Experimental.
	ForEach cdktn.ITerraformIterator `field:"optional" json:"forEach" yaml:"forEach"`
	// Experimental.
	Lifecycle *cdktn.TerraformResourceLifecycle `field:"optional" json:"lifecycle" yaml:"lifecycle"`
	// Experimental.
	Provider cdktn.TerraformProvider `field:"optional" json:"provider" yaml:"provider"`
	// Experimental.
	Provisioners *[]interface{} `field:"optional" json:"provisioners" yaml:"provisioners"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.75.0/docs/resources/data_protection_backup_instance_data_lake_storage#backup_policy_data_lake_storage_id DataProtectionBackupInstanceDataLakeStorage#backup_policy_data_lake_storage_id}.
	BackupPolicyDataLakeStorageId *string `field:"required" json:"backupPolicyDataLakeStorageId" yaml:"backupPolicyDataLakeStorageId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.75.0/docs/resources/data_protection_backup_instance_data_lake_storage#data_protection_backup_vault_id DataProtectionBackupInstanceDataLakeStorage#data_protection_backup_vault_id}.
	DataProtectionBackupVaultId *string `field:"required" json:"dataProtectionBackupVaultId" yaml:"dataProtectionBackupVaultId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.75.0/docs/resources/data_protection_backup_instance_data_lake_storage#location DataProtectionBackupInstanceDataLakeStorage#location}.
	Location *string `field:"required" json:"location" yaml:"location"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.75.0/docs/resources/data_protection_backup_instance_data_lake_storage#name DataProtectionBackupInstanceDataLakeStorage#name}.
	Name *string `field:"required" json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.75.0/docs/resources/data_protection_backup_instance_data_lake_storage#storage_account_id DataProtectionBackupInstanceDataLakeStorage#storage_account_id}.
	StorageAccountId *string `field:"required" json:"storageAccountId" yaml:"storageAccountId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.75.0/docs/resources/data_protection_backup_instance_data_lake_storage#storage_container_names DataProtectionBackupInstanceDataLakeStorage#storage_container_names}.
	StorageContainerNames *[]*string `field:"required" json:"storageContainerNames" yaml:"storageContainerNames"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.75.0/docs/resources/data_protection_backup_instance_data_lake_storage#id DataProtectionBackupInstanceDataLakeStorage#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// timeouts block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.75.0/docs/resources/data_protection_backup_instance_data_lake_storage#timeouts DataProtectionBackupInstanceDataLakeStorage#timeouts}
	Timeouts *DataProtectionBackupInstanceDataLakeStorageTimeouts `field:"optional" json:"timeouts" yaml:"timeouts"`
}

