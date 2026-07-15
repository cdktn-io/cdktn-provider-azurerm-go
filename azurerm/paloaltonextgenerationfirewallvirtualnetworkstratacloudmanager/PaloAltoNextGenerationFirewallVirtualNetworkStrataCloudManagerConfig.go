// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package paloaltonextgenerationfirewallvirtualnetworkstratacloudmanager

import (
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type PaloAltoNextGenerationFirewallVirtualNetworkStrataCloudManagerConfig struct {
	// Experimental.
	Connection interface{} `field:"optional" json:"connection" yaml:"connection"`
	// Experimental.
	Count interface{} `field:"optional" json:"count" yaml:"count"`
	// Experimental.
	DependsOn *[]cdktn.ITerraformDependable `field:"optional" json:"dependsOn" yaml:"dependsOn"`
	// Experimental.
	ForEach cdktn.ITerraformIterator `field:"optional" json:"forEach" yaml:"forEach"`
	// Experimental.
	Lifecycle *cdktn.TerraformResourceLifecycle `field:"optional" json:"lifecycle" yaml:"lifecycle"`
	// Experimental.
	Provider cdktn.TerraformProvider `field:"optional" json:"provider" yaml:"provider"`
	// Experimental.
	Provisioners *[]interface{} `field:"optional" json:"provisioners" yaml:"provisioners"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.81.0/docs/resources/palo_alto_next_generation_firewall_virtual_network_strata_cloud_manager#location PaloAltoNextGenerationFirewallVirtualNetworkStrataCloudManager#location}.
	Location *string `field:"required" json:"location" yaml:"location"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.81.0/docs/resources/palo_alto_next_generation_firewall_virtual_network_strata_cloud_manager#name PaloAltoNextGenerationFirewallVirtualNetworkStrataCloudManager#name}.
	Name *string `field:"required" json:"name" yaml:"name"`
	// network_profile block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.81.0/docs/resources/palo_alto_next_generation_firewall_virtual_network_strata_cloud_manager#network_profile PaloAltoNextGenerationFirewallVirtualNetworkStrataCloudManager#network_profile}
	NetworkProfile *PaloAltoNextGenerationFirewallVirtualNetworkStrataCloudManagerNetworkProfile `field:"required" json:"networkProfile" yaml:"networkProfile"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.81.0/docs/resources/palo_alto_next_generation_firewall_virtual_network_strata_cloud_manager#resource_group_name PaloAltoNextGenerationFirewallVirtualNetworkStrataCloudManager#resource_group_name}.
	ResourceGroupName *string `field:"required" json:"resourceGroupName" yaml:"resourceGroupName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.81.0/docs/resources/palo_alto_next_generation_firewall_virtual_network_strata_cloud_manager#strata_cloud_manager_tenant_name PaloAltoNextGenerationFirewallVirtualNetworkStrataCloudManager#strata_cloud_manager_tenant_name}.
	StrataCloudManagerTenantName *string `field:"required" json:"strataCloudManagerTenantName" yaml:"strataCloudManagerTenantName"`
	// destination_nat block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.81.0/docs/resources/palo_alto_next_generation_firewall_virtual_network_strata_cloud_manager#destination_nat PaloAltoNextGenerationFirewallVirtualNetworkStrataCloudManager#destination_nat}
	DestinationNat interface{} `field:"optional" json:"destinationNat" yaml:"destinationNat"`
	// dns_settings block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.81.0/docs/resources/palo_alto_next_generation_firewall_virtual_network_strata_cloud_manager#dns_settings PaloAltoNextGenerationFirewallVirtualNetworkStrataCloudManager#dns_settings}
	DnsSettings *PaloAltoNextGenerationFirewallVirtualNetworkStrataCloudManagerDnsSettings `field:"optional" json:"dnsSettings" yaml:"dnsSettings"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.81.0/docs/resources/palo_alto_next_generation_firewall_virtual_network_strata_cloud_manager#id PaloAltoNextGenerationFirewallVirtualNetworkStrataCloudManager#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// identity block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.81.0/docs/resources/palo_alto_next_generation_firewall_virtual_network_strata_cloud_manager#identity PaloAltoNextGenerationFirewallVirtualNetworkStrataCloudManager#identity}
	Identity *PaloAltoNextGenerationFirewallVirtualNetworkStrataCloudManagerIdentity `field:"optional" json:"identity" yaml:"identity"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.81.0/docs/resources/palo_alto_next_generation_firewall_virtual_network_strata_cloud_manager#marketplace_offer_id PaloAltoNextGenerationFirewallVirtualNetworkStrataCloudManager#marketplace_offer_id}.
	MarketplaceOfferId *string `field:"optional" json:"marketplaceOfferId" yaml:"marketplaceOfferId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.81.0/docs/resources/palo_alto_next_generation_firewall_virtual_network_strata_cloud_manager#plan_id PaloAltoNextGenerationFirewallVirtualNetworkStrataCloudManager#plan_id}.
	PlanId *string `field:"optional" json:"planId" yaml:"planId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.81.0/docs/resources/palo_alto_next_generation_firewall_virtual_network_strata_cloud_manager#tags PaloAltoNextGenerationFirewallVirtualNetworkStrataCloudManager#tags}.
	Tags *map[string]*string `field:"optional" json:"tags" yaml:"tags"`
	// timeouts block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.81.0/docs/resources/palo_alto_next_generation_firewall_virtual_network_strata_cloud_manager#timeouts PaloAltoNextGenerationFirewallVirtualNetworkStrataCloudManager#timeouts}
	Timeouts *PaloAltoNextGenerationFirewallVirtualNetworkStrataCloudManagerTimeouts `field:"optional" json:"timeouts" yaml:"timeouts"`
}

