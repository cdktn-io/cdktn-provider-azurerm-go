// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package storageaccount


type StorageAccountBlobPropertiesRestorePolicy struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.64.0/docs/resources/storage_account#days StorageAccount#days}.
	Days *float64 `field:"required" json:"days" yaml:"days"`
}

