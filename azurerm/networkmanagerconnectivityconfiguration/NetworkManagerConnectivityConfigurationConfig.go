// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package networkmanagerconnectivityconfiguration

import (
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type NetworkManagerConnectivityConfigurationConfig struct {
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
	// applies_to_group block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.80.0/docs/resources/network_manager_connectivity_configuration#applies_to_group NetworkManagerConnectivityConfiguration#applies_to_group}
	AppliesToGroup interface{} `field:"required" json:"appliesToGroup" yaml:"appliesToGroup"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.80.0/docs/resources/network_manager_connectivity_configuration#connectivity_topology NetworkManagerConnectivityConfiguration#connectivity_topology}.
	ConnectivityTopology *string `field:"required" json:"connectivityTopology" yaml:"connectivityTopology"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.80.0/docs/resources/network_manager_connectivity_configuration#name NetworkManagerConnectivityConfiguration#name}.
	Name *string `field:"required" json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.80.0/docs/resources/network_manager_connectivity_configuration#network_manager_id NetworkManagerConnectivityConfiguration#network_manager_id}.
	NetworkManagerId *string `field:"required" json:"networkManagerId" yaml:"networkManagerId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.80.0/docs/resources/network_manager_connectivity_configuration#connected_group_address_overlap_enabled NetworkManagerConnectivityConfiguration#connected_group_address_overlap_enabled}.
	ConnectedGroupAddressOverlapEnabled interface{} `field:"optional" json:"connectedGroupAddressOverlapEnabled" yaml:"connectedGroupAddressOverlapEnabled"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.80.0/docs/resources/network_manager_connectivity_configuration#connected_group_private_endpoints_scale NetworkManagerConnectivityConfiguration#connected_group_private_endpoints_scale}.
	ConnectedGroupPrivateEndpointsScale *string `field:"optional" json:"connectedGroupPrivateEndpointsScale" yaml:"connectedGroupPrivateEndpointsScale"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.80.0/docs/resources/network_manager_connectivity_configuration#delete_existing_peering_enabled NetworkManagerConnectivityConfiguration#delete_existing_peering_enabled}.
	DeleteExistingPeeringEnabled interface{} `field:"optional" json:"deleteExistingPeeringEnabled" yaml:"deleteExistingPeeringEnabled"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.80.0/docs/resources/network_manager_connectivity_configuration#description NetworkManagerConnectivityConfiguration#description}.
	Description *string `field:"optional" json:"description" yaml:"description"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.80.0/docs/resources/network_manager_connectivity_configuration#global_mesh_enabled NetworkManagerConnectivityConfiguration#global_mesh_enabled}.
	GlobalMeshEnabled interface{} `field:"optional" json:"globalMeshEnabled" yaml:"globalMeshEnabled"`
	// hub block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.80.0/docs/resources/network_manager_connectivity_configuration#hub NetworkManagerConnectivityConfiguration#hub}
	Hub *NetworkManagerConnectivityConfigurationHub `field:"optional" json:"hub" yaml:"hub"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.80.0/docs/resources/network_manager_connectivity_configuration#id NetworkManagerConnectivityConfiguration#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.80.0/docs/resources/network_manager_connectivity_configuration#peering_enforcement_enabled NetworkManagerConnectivityConfiguration#peering_enforcement_enabled}.
	PeeringEnforcementEnabled interface{} `field:"optional" json:"peeringEnforcementEnabled" yaml:"peeringEnforcementEnabled"`
	// timeouts block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.80.0/docs/resources/network_manager_connectivity_configuration#timeouts NetworkManagerConnectivityConfiguration#timeouts}
	Timeouts *NetworkManagerConnectivityConfigurationTimeouts `field:"optional" json:"timeouts" yaml:"timeouts"`
}

