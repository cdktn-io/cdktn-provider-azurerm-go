// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package dataprotectionbackuppolicycosmosdbaccount

import (
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type DataProtectionBackupPolicyCosmosdbAccountConfig struct {
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
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/5.11.0/docs/resources/data_protection_backup_policy_cosmosdb_account#backup_schedule DataProtectionBackupPolicyCosmosdbAccount#backup_schedule}.
	BackupSchedule *string `field:"required" json:"backupSchedule" yaml:"backupSchedule"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/5.11.0/docs/resources/data_protection_backup_policy_cosmosdb_account#data_protection_backup_vault_id DataProtectionBackupPolicyCosmosdbAccount#data_protection_backup_vault_id}.
	DataProtectionBackupVaultId *string `field:"required" json:"dataProtectionBackupVaultId" yaml:"dataProtectionBackupVaultId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/5.11.0/docs/resources/data_protection_backup_policy_cosmosdb_account#default_retention_duration DataProtectionBackupPolicyCosmosdbAccount#default_retention_duration}.
	DefaultRetentionDuration *string `field:"required" json:"defaultRetentionDuration" yaml:"defaultRetentionDuration"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/5.11.0/docs/resources/data_protection_backup_policy_cosmosdb_account#name DataProtectionBackupPolicyCosmosdbAccount#name}.
	Name *string `field:"required" json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/5.11.0/docs/resources/data_protection_backup_policy_cosmosdb_account#backup_frequency DataProtectionBackupPolicyCosmosdbAccount#backup_frequency}.
	BackupFrequency *string `field:"optional" json:"backupFrequency" yaml:"backupFrequency"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/5.11.0/docs/resources/data_protection_backup_policy_cosmosdb_account#id DataProtectionBackupPolicyCosmosdbAccount#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// monthly_retention_rule block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/5.11.0/docs/resources/data_protection_backup_policy_cosmosdb_account#monthly_retention_rule DataProtectionBackupPolicyCosmosdbAccount#monthly_retention_rule}
	MonthlyRetentionRule *DataProtectionBackupPolicyCosmosdbAccountMonthlyRetentionRule `field:"optional" json:"monthlyRetentionRule" yaml:"monthlyRetentionRule"`
	// timeouts block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/5.11.0/docs/resources/data_protection_backup_policy_cosmosdb_account#timeouts DataProtectionBackupPolicyCosmosdbAccount#timeouts}
	Timeouts *DataProtectionBackupPolicyCosmosdbAccountTimeouts `field:"optional" json:"timeouts" yaml:"timeouts"`
	// weekly_retention_rule block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/5.11.0/docs/resources/data_protection_backup_policy_cosmosdb_account#weekly_retention_rule DataProtectionBackupPolicyCosmosdbAccount#weekly_retention_rule}
	WeeklyRetentionRule *DataProtectionBackupPolicyCosmosdbAccountWeeklyRetentionRule `field:"optional" json:"weeklyRetentionRule" yaml:"weeklyRetentionRule"`
	// yearly_retention_rule block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/5.11.0/docs/resources/data_protection_backup_policy_cosmosdb_account#yearly_retention_rule DataProtectionBackupPolicyCosmosdbAccount#yearly_retention_rule}
	YearlyRetentionRule *DataProtectionBackupPolicyCosmosdbAccountYearlyRetentionRule `field:"optional" json:"yearlyRetentionRule" yaml:"yearlyRetentionRule"`
}

