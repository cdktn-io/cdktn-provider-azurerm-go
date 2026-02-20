// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package paloaltonextgenerationfirewallvirtualhubstratacloudmanager


type PaloAltoNextGenerationFirewallVirtualHubStrataCloudManagerDestinationNat struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.61.0/docs/resources/palo_alto_next_generation_firewall_virtual_hub_strata_cloud_manager#name PaloAltoNextGenerationFirewallVirtualHubStrataCloudManager#name}.
	Name *string `field:"required" json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.61.0/docs/resources/palo_alto_next_generation_firewall_virtual_hub_strata_cloud_manager#protocol PaloAltoNextGenerationFirewallVirtualHubStrataCloudManager#protocol}.
	Protocol *string `field:"required" json:"protocol" yaml:"protocol"`
	// backend_config block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.61.0/docs/resources/palo_alto_next_generation_firewall_virtual_hub_strata_cloud_manager#backend_config PaloAltoNextGenerationFirewallVirtualHubStrataCloudManager#backend_config}
	BackendConfig *PaloAltoNextGenerationFirewallVirtualHubStrataCloudManagerDestinationNatBackendConfig `field:"optional" json:"backendConfig" yaml:"backendConfig"`
	// frontend_config block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.61.0/docs/resources/palo_alto_next_generation_firewall_virtual_hub_strata_cloud_manager#frontend_config PaloAltoNextGenerationFirewallVirtualHubStrataCloudManager#frontend_config}
	FrontendConfig *PaloAltoNextGenerationFirewallVirtualHubStrataCloudManagerDestinationNatFrontendConfig `field:"optional" json:"frontendConfig" yaml:"frontendConfig"`
}

