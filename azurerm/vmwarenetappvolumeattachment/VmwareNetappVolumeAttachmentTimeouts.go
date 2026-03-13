// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package vmwarenetappvolumeattachment


type VmwareNetappVolumeAttachmentTimeouts struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.64.0/docs/resources/vmware_netapp_volume_attachment#create VmwareNetappVolumeAttachment#create}.
	Create *string `field:"optional" json:"create" yaml:"create"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.64.0/docs/resources/vmware_netapp_volume_attachment#delete VmwareNetappVolumeAttachment#delete}.
	Delete *string `field:"optional" json:"delete" yaml:"delete"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.64.0/docs/resources/vmware_netapp_volume_attachment#read VmwareNetappVolumeAttachment#read}.
	Read *string `field:"optional" json:"read" yaml:"read"`
}

