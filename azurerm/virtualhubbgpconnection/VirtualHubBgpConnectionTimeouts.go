// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package virtualhubbgpconnection


type VirtualHubBgpConnectionTimeouts struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.81.0/docs/resources/virtual_hub_bgp_connection#create VirtualHubBgpConnection#create}.
	Create *string `field:"optional" json:"create" yaml:"create"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.81.0/docs/resources/virtual_hub_bgp_connection#delete VirtualHubBgpConnection#delete}.
	Delete *string `field:"optional" json:"delete" yaml:"delete"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.81.0/docs/resources/virtual_hub_bgp_connection#read VirtualHubBgpConnection#read}.
	Read *string `field:"optional" json:"read" yaml:"read"`
}

