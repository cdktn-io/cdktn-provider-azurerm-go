// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package dataprotectionbackuppolicycosmosdbaccount


type DataProtectionBackupPolicyCosmosdbAccountWeeklyRetentionRule struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/5.11.0/docs/resources/data_protection_backup_policy_cosmosdb_account#duration DataProtectionBackupPolicyCosmosdbAccount#duration}.
	Duration *string `field:"required" json:"duration" yaml:"duration"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/5.11.0/docs/resources/data_protection_backup_policy_cosmosdb_account#days_of_week DataProtectionBackupPolicyCosmosdbAccount#days_of_week}.
	DaysOfWeek *[]*string `field:"optional" json:"daysOfWeek" yaml:"daysOfWeek"`
}

