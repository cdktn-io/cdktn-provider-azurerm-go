// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package backuppolicyvm


type BackupPolicyVmTieringPolicy struct {
	// archived_restore_point block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.74.0/docs/resources/backup_policy_vm#archived_restore_point BackupPolicyVm#archived_restore_point}
	ArchivedRestorePoint *BackupPolicyVmTieringPolicyArchivedRestorePoint `field:"required" json:"archivedRestorePoint" yaml:"archivedRestorePoint"`
}

