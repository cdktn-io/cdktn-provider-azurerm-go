// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package paloaltonextgenerationfirewallvirtualnetworkstratacloudmanager


type PaloAltoNextGenerationFirewallVirtualNetworkStrataCloudManagerNetworkProfileVnetConfiguration struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.62.0/docs/resources/palo_alto_next_generation_firewall_virtual_network_strata_cloud_manager#virtual_network_id PaloAltoNextGenerationFirewallVirtualNetworkStrataCloudManager#virtual_network_id}.
	VirtualNetworkId *string `field:"required" json:"virtualNetworkId" yaml:"virtualNetworkId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.62.0/docs/resources/palo_alto_next_generation_firewall_virtual_network_strata_cloud_manager#trusted_subnet_id PaloAltoNextGenerationFirewallVirtualNetworkStrataCloudManager#trusted_subnet_id}.
	TrustedSubnetId *string `field:"optional" json:"trustedSubnetId" yaml:"trustedSubnetId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.62.0/docs/resources/palo_alto_next_generation_firewall_virtual_network_strata_cloud_manager#untrusted_subnet_id PaloAltoNextGenerationFirewallVirtualNetworkStrataCloudManager#untrusted_subnet_id}.
	UntrustedSubnetId *string `field:"optional" json:"untrustedSubnetId" yaml:"untrustedSubnetId"`
}

