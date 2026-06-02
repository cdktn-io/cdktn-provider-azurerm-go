// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package linuxwebappslot


type LinuxWebAppSlotSiteConfigAutoHealSettingTriggerRequests struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.75.0/docs/resources/linux_web_app_slot#count LinuxWebAppSlot#count}.
	Count *float64 `field:"required" json:"count" yaml:"count"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.75.0/docs/resources/linux_web_app_slot#interval LinuxWebAppSlot#interval}.
	Interval *string `field:"required" json:"interval" yaml:"interval"`
}

