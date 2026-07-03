// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package netappvolumebucket


type NetappVolumeBucketFileSystemNfsUser struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.80.0/docs/resources/netapp_volume_bucket#group_id NetappVolumeBucket#group_id}.
	GroupId *float64 `field:"required" json:"groupId" yaml:"groupId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.80.0/docs/resources/netapp_volume_bucket#user_id NetappVolumeBucket#user_id}.
	UserId *float64 `field:"required" json:"userId" yaml:"userId"`
}

