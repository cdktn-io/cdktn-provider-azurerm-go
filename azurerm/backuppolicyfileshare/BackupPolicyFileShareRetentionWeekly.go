// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package backuppolicyfileshare


type BackupPolicyFileShareRetentionWeekly struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.66.0/docs/resources/backup_policy_file_share#count BackupPolicyFileShare#count}.
	Count *float64 `field:"required" json:"count" yaml:"count"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.66.0/docs/resources/backup_policy_file_share#weekdays BackupPolicyFileShare#weekdays}.
	Weekdays *[]*string `field:"required" json:"weekdays" yaml:"weekdays"`
}

