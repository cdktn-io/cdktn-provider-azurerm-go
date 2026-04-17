// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package dataprotectionbackuppolicydatalakestorage


type DataProtectionBackupPolicyDataLakeStorageRetentionRule struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.69.0/docs/resources/data_protection_backup_policy_data_lake_storage#duration DataProtectionBackupPolicyDataLakeStorage#duration}.
	Duration *string `field:"required" json:"duration" yaml:"duration"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.69.0/docs/resources/data_protection_backup_policy_data_lake_storage#name DataProtectionBackupPolicyDataLakeStorage#name}.
	Name *string `field:"required" json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.69.0/docs/resources/data_protection_backup_policy_data_lake_storage#absolute_criteria DataProtectionBackupPolicyDataLakeStorage#absolute_criteria}.
	AbsoluteCriteria *string `field:"optional" json:"absoluteCriteria" yaml:"absoluteCriteria"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.69.0/docs/resources/data_protection_backup_policy_data_lake_storage#days_of_week DataProtectionBackupPolicyDataLakeStorage#days_of_week}.
	DaysOfWeek *[]*string `field:"optional" json:"daysOfWeek" yaml:"daysOfWeek"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.69.0/docs/resources/data_protection_backup_policy_data_lake_storage#months_of_year DataProtectionBackupPolicyDataLakeStorage#months_of_year}.
	MonthsOfYear *[]*string `field:"optional" json:"monthsOfYear" yaml:"monthsOfYear"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.69.0/docs/resources/data_protection_backup_policy_data_lake_storage#scheduled_backup_times DataProtectionBackupPolicyDataLakeStorage#scheduled_backup_times}.
	ScheduledBackupTimes *[]*string `field:"optional" json:"scheduledBackupTimes" yaml:"scheduledBackupTimes"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.69.0/docs/resources/data_protection_backup_policy_data_lake_storage#weeks_of_month DataProtectionBackupPolicyDataLakeStorage#weeks_of_month}.
	WeeksOfMonth *[]*string `field:"optional" json:"weeksOfMonth" yaml:"weeksOfMonth"`
}

