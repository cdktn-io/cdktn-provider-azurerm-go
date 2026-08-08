// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package paloaltonextgenerationfirewallvirtualhubstratacloudmanager


type PaloAltoNextGenerationFirewallVirtualHubStrataCloudManagerNetworkProfile struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/5.0.1/docs/resources/palo_alto_next_generation_firewall_virtual_hub_strata_cloud_manager#network_virtual_appliance_id PaloAltoNextGenerationFirewallVirtualHubStrataCloudManager#network_virtual_appliance_id}.
	NetworkVirtualApplianceId *string `field:"required" json:"networkVirtualApplianceId" yaml:"networkVirtualApplianceId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/5.0.1/docs/resources/palo_alto_next_generation_firewall_virtual_hub_strata_cloud_manager#public_ip_address_ids PaloAltoNextGenerationFirewallVirtualHubStrataCloudManager#public_ip_address_ids}.
	PublicIpAddressIds *[]*string `field:"required" json:"publicIpAddressIds" yaml:"publicIpAddressIds"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/5.0.1/docs/resources/palo_alto_next_generation_firewall_virtual_hub_strata_cloud_manager#virtual_hub_id PaloAltoNextGenerationFirewallVirtualHubStrataCloudManager#virtual_hub_id}.
	VirtualHubId *string `field:"required" json:"virtualHubId" yaml:"virtualHubId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/5.0.1/docs/resources/palo_alto_next_generation_firewall_virtual_hub_strata_cloud_manager#egress_nat_ip_address_ids PaloAltoNextGenerationFirewallVirtualHubStrataCloudManager#egress_nat_ip_address_ids}.
	EgressNatIpAddressIds *[]*string `field:"optional" json:"egressNatIpAddressIds" yaml:"egressNatIpAddressIds"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/5.0.1/docs/resources/palo_alto_next_generation_firewall_virtual_hub_strata_cloud_manager#trusted_address_ranges PaloAltoNextGenerationFirewallVirtualHubStrataCloudManager#trusted_address_ranges}.
	TrustedAddressRanges *[]*string `field:"optional" json:"trustedAddressRanges" yaml:"trustedAddressRanges"`
}

