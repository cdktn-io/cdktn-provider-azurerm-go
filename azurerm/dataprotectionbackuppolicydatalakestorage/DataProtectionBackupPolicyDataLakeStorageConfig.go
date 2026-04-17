// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package dataprotectionbackuppolicydatalakestorage

import (
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type DataProtectionBackupPolicyDataLakeStorageConfig struct {
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
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.69.0/docs/resources/data_protection_backup_policy_data_lake_storage#backup_schedule DataProtectionBackupPolicyDataLakeStorage#backup_schedule}.
	BackupSchedule *[]*string `field:"required" json:"backupSchedule" yaml:"backupSchedule"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.69.0/docs/resources/data_protection_backup_policy_data_lake_storage#data_protection_backup_vault_id DataProtectionBackupPolicyDataLakeStorage#data_protection_backup_vault_id}.
	DataProtectionBackupVaultId *string `field:"required" json:"dataProtectionBackupVaultId" yaml:"dataProtectionBackupVaultId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.69.0/docs/resources/data_protection_backup_policy_data_lake_storage#default_retention_duration DataProtectionBackupPolicyDataLakeStorage#default_retention_duration}.
	DefaultRetentionDuration *string `field:"required" json:"defaultRetentionDuration" yaml:"defaultRetentionDuration"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.69.0/docs/resources/data_protection_backup_policy_data_lake_storage#name DataProtectionBackupPolicyDataLakeStorage#name}.
	Name *string `field:"required" json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.69.0/docs/resources/data_protection_backup_policy_data_lake_storage#id DataProtectionBackupPolicyDataLakeStorage#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// retention_rule block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.69.0/docs/resources/data_protection_backup_policy_data_lake_storage#retention_rule DataProtectionBackupPolicyDataLakeStorage#retention_rule}
	RetentionRule interface{} `field:"optional" json:"retentionRule" yaml:"retentionRule"`
	// timeouts block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.69.0/docs/resources/data_protection_backup_policy_data_lake_storage#timeouts DataProtectionBackupPolicyDataLakeStorage#timeouts}
	Timeouts *DataProtectionBackupPolicyDataLakeStorageTimeouts `field:"optional" json:"timeouts" yaml:"timeouts"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.69.0/docs/resources/data_protection_backup_policy_data_lake_storage#time_zone DataProtectionBackupPolicyDataLakeStorage#time_zone}.
	TimeZone *string `field:"optional" json:"timeZone" yaml:"timeZone"`
}

