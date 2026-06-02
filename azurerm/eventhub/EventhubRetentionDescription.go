// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package eventhub


type EventhubRetentionDescription struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.75.0/docs/resources/eventhub#cleanup_policy Eventhub#cleanup_policy}.
	CleanupPolicy *string `field:"required" json:"cleanupPolicy" yaml:"cleanupPolicy"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.75.0/docs/resources/eventhub#retention_time_in_hours Eventhub#retention_time_in_hours}.
	RetentionTimeInHours *float64 `field:"optional" json:"retentionTimeInHours" yaml:"retentionTimeInHours"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.75.0/docs/resources/eventhub#tombstone_retention_time_in_hours Eventhub#tombstone_retention_time_in_hours}.
	TombstoneRetentionTimeInHours *float64 `field:"optional" json:"tombstoneRetentionTimeInHours" yaml:"tombstoneRetentionTimeInHours"`
}

