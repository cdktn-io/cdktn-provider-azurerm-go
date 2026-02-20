// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package paloaltonextgenerationfirewallvirtualhubstratacloudmanager

import (
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type PaloAltoNextGenerationFirewallVirtualHubStrataCloudManagerConfig struct {
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
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.61.0/docs/resources/palo_alto_next_generation_firewall_virtual_hub_strata_cloud_manager#location PaloAltoNextGenerationFirewallVirtualHubStrataCloudManager#location}.
	Location *string `field:"required" json:"location" yaml:"location"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.61.0/docs/resources/palo_alto_next_generation_firewall_virtual_hub_strata_cloud_manager#name PaloAltoNextGenerationFirewallVirtualHubStrataCloudManager#name}.
	Name *string `field:"required" json:"name" yaml:"name"`
	// network_profile block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.61.0/docs/resources/palo_alto_next_generation_firewall_virtual_hub_strata_cloud_manager#network_profile PaloAltoNextGenerationFirewallVirtualHubStrataCloudManager#network_profile}
	NetworkProfile *PaloAltoNextGenerationFirewallVirtualHubStrataCloudManagerNetworkProfile `field:"required" json:"networkProfile" yaml:"networkProfile"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.61.0/docs/resources/palo_alto_next_generation_firewall_virtual_hub_strata_cloud_manager#resource_group_name PaloAltoNextGenerationFirewallVirtualHubStrataCloudManager#resource_group_name}.
	ResourceGroupName *string `field:"required" json:"resourceGroupName" yaml:"resourceGroupName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.61.0/docs/resources/palo_alto_next_generation_firewall_virtual_hub_strata_cloud_manager#strata_cloud_manager_tenant_name PaloAltoNextGenerationFirewallVirtualHubStrataCloudManager#strata_cloud_manager_tenant_name}.
	StrataCloudManagerTenantName *string `field:"required" json:"strataCloudManagerTenantName" yaml:"strataCloudManagerTenantName"`
	// destination_nat block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.61.0/docs/resources/palo_alto_next_generation_firewall_virtual_hub_strata_cloud_manager#destination_nat PaloAltoNextGenerationFirewallVirtualHubStrataCloudManager#destination_nat}
	DestinationNat interface{} `field:"optional" json:"destinationNat" yaml:"destinationNat"`
	// dns_settings block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.61.0/docs/resources/palo_alto_next_generation_firewall_virtual_hub_strata_cloud_manager#dns_settings PaloAltoNextGenerationFirewallVirtualHubStrataCloudManager#dns_settings}
	DnsSettings *PaloAltoNextGenerationFirewallVirtualHubStrataCloudManagerDnsSettings `field:"optional" json:"dnsSettings" yaml:"dnsSettings"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.61.0/docs/resources/palo_alto_next_generation_firewall_virtual_hub_strata_cloud_manager#id PaloAltoNextGenerationFirewallVirtualHubStrataCloudManager#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// identity block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.61.0/docs/resources/palo_alto_next_generation_firewall_virtual_hub_strata_cloud_manager#identity PaloAltoNextGenerationFirewallVirtualHubStrataCloudManager#identity}
	Identity *PaloAltoNextGenerationFirewallVirtualHubStrataCloudManagerIdentity `field:"optional" json:"identity" yaml:"identity"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.61.0/docs/resources/palo_alto_next_generation_firewall_virtual_hub_strata_cloud_manager#marketplace_offer_id PaloAltoNextGenerationFirewallVirtualHubStrataCloudManager#marketplace_offer_id}.
	MarketplaceOfferId *string `field:"optional" json:"marketplaceOfferId" yaml:"marketplaceOfferId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.61.0/docs/resources/palo_alto_next_generation_firewall_virtual_hub_strata_cloud_manager#plan_id PaloAltoNextGenerationFirewallVirtualHubStrataCloudManager#plan_id}.
	PlanId *string `field:"optional" json:"planId" yaml:"planId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.61.0/docs/resources/palo_alto_next_generation_firewall_virtual_hub_strata_cloud_manager#tags PaloAltoNextGenerationFirewallVirtualHubStrataCloudManager#tags}.
	Tags *map[string]*string `field:"optional" json:"tags" yaml:"tags"`
	// timeouts block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.61.0/docs/resources/palo_alto_next_generation_firewall_virtual_hub_strata_cloud_manager#timeouts PaloAltoNextGenerationFirewallVirtualHubStrataCloudManager#timeouts}
	Timeouts *PaloAltoNextGenerationFirewallVirtualHubStrataCloudManagerTimeouts `field:"optional" json:"timeouts" yaml:"timeouts"`
}

