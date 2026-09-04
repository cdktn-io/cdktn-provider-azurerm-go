// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package virtualnetwork


type VirtualNetworkSubnetServiceEndpoint struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/5.4.0/docs/resources/virtual_network#network_identifier VirtualNetwork#network_identifier}.
	NetworkIdentifier *string `field:"optional" json:"networkIdentifier" yaml:"networkIdentifier"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/5.4.0/docs/resources/virtual_network#service VirtualNetwork#service}.
	Service *string `field:"optional" json:"service" yaml:"service"`
}

