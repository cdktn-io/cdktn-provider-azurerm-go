// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package netappvolume


type NetappVolumeDataProtectionAdvancedRansomware struct {
	// Enable or disable the Advanced Ransomware Protection feature.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/5.3.0/docs/resources/netapp_volume#protection_enabled NetappVolume#protection_enabled}
	ProtectionEnabled interface{} `field:"required" json:"protectionEnabled" yaml:"protectionEnabled"`
}

