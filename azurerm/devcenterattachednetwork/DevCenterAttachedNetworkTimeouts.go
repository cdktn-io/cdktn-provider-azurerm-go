// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package devcenterattachednetwork


type DevCenterAttachedNetworkTimeouts struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.80.0/docs/resources/dev_center_attached_network#create DevCenterAttachedNetwork#create}.
	Create *string `field:"optional" json:"create" yaml:"create"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.80.0/docs/resources/dev_center_attached_network#delete DevCenterAttachedNetwork#delete}.
	Delete *string `field:"optional" json:"delete" yaml:"delete"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.80.0/docs/resources/dev_center_attached_network#read DevCenterAttachedNetwork#read}.
	Read *string `field:"optional" json:"read" yaml:"read"`
}

