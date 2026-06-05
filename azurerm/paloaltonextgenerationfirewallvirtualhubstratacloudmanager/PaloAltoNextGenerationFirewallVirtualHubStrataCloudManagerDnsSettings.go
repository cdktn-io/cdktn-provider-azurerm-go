// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package paloaltonextgenerationfirewallvirtualhubstratacloudmanager


type PaloAltoNextGenerationFirewallVirtualHubStrataCloudManagerDnsSettings struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.76.0/docs/resources/palo_alto_next_generation_firewall_virtual_hub_strata_cloud_manager#dns_servers PaloAltoNextGenerationFirewallVirtualHubStrataCloudManager#dns_servers}.
	DnsServers *[]*string `field:"optional" json:"dnsServers" yaml:"dnsServers"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.76.0/docs/resources/palo_alto_next_generation_firewall_virtual_hub_strata_cloud_manager#use_azure_dns PaloAltoNextGenerationFirewallVirtualHubStrataCloudManager#use_azure_dns}.
	UseAzureDns interface{} `field:"optional" json:"useAzureDns" yaml:"useAzureDns"`
}

