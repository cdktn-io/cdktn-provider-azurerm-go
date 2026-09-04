// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package subnet


type SubnetServiceEndpoint struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/5.4.0/docs/resources/subnet#service Subnet#service}.
	Service *string `field:"required" json:"service" yaml:"service"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/5.4.0/docs/resources/subnet#network_identifier Subnet#network_identifier}.
	NetworkIdentifier *string `field:"optional" json:"networkIdentifier" yaml:"networkIdentifier"`
}

