// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package vmwareprivatecloud


type VmwarePrivateCloudManagementCluster struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/5.4.0/docs/resources/vmware_private_cloud#size VmwarePrivateCloud#size}.
	Size *float64 `field:"required" json:"size" yaml:"size"`
}

