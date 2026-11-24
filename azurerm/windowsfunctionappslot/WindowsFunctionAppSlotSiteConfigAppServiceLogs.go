// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package windowsfunctionappslot


type WindowsFunctionAppSlotSiteConfigAppServiceLogs struct {
	// The amount of disk space to use for logs.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/5.15.0/docs/resources/windows_function_app_slot#disk_quota_mb WindowsFunctionAppSlot#disk_quota_mb}
	DiskQuotaMb *float64 `field:"optional" json:"diskQuotaMb" yaml:"diskQuotaMb"`
	// The retention period for logs in days. Defaults to `0` (never delete).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/5.15.0/docs/resources/windows_function_app_slot#retention_period_days WindowsFunctionAppSlot#retention_period_days}
	RetentionPeriodDays *float64 `field:"optional" json:"retentionPeriodDays" yaml:"retentionPeriodDays"`
}

