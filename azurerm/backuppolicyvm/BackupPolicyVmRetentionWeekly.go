// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package backuppolicyvm


type BackupPolicyVmRetentionWeekly struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/5.1.0/docs/resources/backup_policy_vm#count BackupPolicyVm#count}.
	Count *float64 `field:"required" json:"count" yaml:"count"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/5.1.0/docs/resources/backup_policy_vm#weekdays BackupPolicyVm#weekdays}.
	Weekdays *[]*string `field:"required" json:"weekdays" yaml:"weekdays"`
}

