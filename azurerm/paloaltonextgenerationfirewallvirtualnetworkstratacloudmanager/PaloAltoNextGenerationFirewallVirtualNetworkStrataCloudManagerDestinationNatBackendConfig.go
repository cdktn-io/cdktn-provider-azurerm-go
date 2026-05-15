// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package paloaltonextgenerationfirewallvirtualnetworkstratacloudmanager


type PaloAltoNextGenerationFirewallVirtualNetworkStrataCloudManagerDestinationNatBackendConfig struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.73.0/docs/resources/palo_alto_next_generation_firewall_virtual_network_strata_cloud_manager#port PaloAltoNextGenerationFirewallVirtualNetworkStrataCloudManager#port}.
	Port *float64 `field:"required" json:"port" yaml:"port"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.73.0/docs/resources/palo_alto_next_generation_firewall_virtual_network_strata_cloud_manager#public_ip_address PaloAltoNextGenerationFirewallVirtualNetworkStrataCloudManager#public_ip_address}.
	PublicIpAddress *string `field:"required" json:"publicIpAddress" yaml:"publicIpAddress"`
}

