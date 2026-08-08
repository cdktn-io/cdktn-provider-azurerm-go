// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package iothub


type IothubNetworkRuleSetIpRule struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/5.0.1/docs/resources/iothub#ip_mask Iothub#ip_mask}.
	IpMask *string `field:"required" json:"ipMask" yaml:"ipMask"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/5.0.1/docs/resources/iothub#name Iothub#name}.
	Name *string `field:"required" json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/5.0.1/docs/resources/iothub#action Iothub#action}.
	Action *string `field:"optional" json:"action" yaml:"action"`
}

