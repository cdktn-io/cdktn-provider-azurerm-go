// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package paloaltonextgenerationfirewallvirtualhubpanorama


type PaloAltoNextGenerationFirewallVirtualHubPanoramaDestinationNatBackendConfig struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.59.0/docs/resources/palo_alto_next_generation_firewall_virtual_hub_panorama#port PaloAltoNextGenerationFirewallVirtualHubPanorama#port}.
	Port *float64 `field:"required" json:"port" yaml:"port"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.59.0/docs/resources/palo_alto_next_generation_firewall_virtual_hub_panorama#public_ip_address PaloAltoNextGenerationFirewallVirtualHubPanorama#public_ip_address}.
	PublicIpAddress *string `field:"required" json:"publicIpAddress" yaml:"publicIpAddress"`
}

