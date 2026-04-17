// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package natgatewaypublicipprefixassociation


type NatGatewayPublicIpPrefixAssociationTimeouts struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.69.0/docs/resources/nat_gateway_public_ip_prefix_association#create NatGatewayPublicIpPrefixAssociation#create}.
	Create *string `field:"optional" json:"create" yaml:"create"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.69.0/docs/resources/nat_gateway_public_ip_prefix_association#delete NatGatewayPublicIpPrefixAssociation#delete}.
	Delete *string `field:"optional" json:"delete" yaml:"delete"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.69.0/docs/resources/nat_gateway_public_ip_prefix_association#read NatGatewayPublicIpPrefixAssociation#read}.
	Read *string `field:"optional" json:"read" yaml:"read"`
}

