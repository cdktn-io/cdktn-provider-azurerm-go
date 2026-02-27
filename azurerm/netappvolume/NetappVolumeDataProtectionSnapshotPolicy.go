// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package netappvolume


type NetappVolumeDataProtectionSnapshotPolicy struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.62.0/docs/resources/netapp_volume#snapshot_policy_id NetappVolume#snapshot_policy_id}.
	SnapshotPolicyId *string `field:"required" json:"snapshotPolicyId" yaml:"snapshotPolicyId"`
}

