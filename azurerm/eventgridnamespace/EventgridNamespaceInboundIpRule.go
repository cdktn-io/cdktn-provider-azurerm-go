// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package eventgridnamespace


type EventgridNamespaceInboundIpRule struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.59.0/docs/resources/eventgrid_namespace#ip_mask EventgridNamespace#ip_mask}.
	IpMask *string `field:"required" json:"ipMask" yaml:"ipMask"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.59.0/docs/resources/eventgrid_namespace#action EventgridNamespace#action}.
	Action *string `field:"optional" json:"action" yaml:"action"`
}

