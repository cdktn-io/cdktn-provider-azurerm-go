// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package paloaltonextgenerationfirewallvirtualhubstratacloudmanager


type PaloAltoNextGenerationFirewallVirtualHubStrataCloudManagerIdentity struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.74.0/docs/resources/palo_alto_next_generation_firewall_virtual_hub_strata_cloud_manager#identity_ids PaloAltoNextGenerationFirewallVirtualHubStrataCloudManager#identity_ids}.
	IdentityIds *[]*string `field:"required" json:"identityIds" yaml:"identityIds"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.74.0/docs/resources/palo_alto_next_generation_firewall_virtual_hub_strata_cloud_manager#type PaloAltoNextGenerationFirewallVirtualHubStrataCloudManager#type}.
	Type *string `field:"required" json:"type" yaml:"type"`
}

