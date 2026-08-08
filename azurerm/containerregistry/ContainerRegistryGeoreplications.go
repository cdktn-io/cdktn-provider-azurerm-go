// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package containerregistry


type ContainerRegistryGeoreplications struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/5.0.1/docs/resources/container_registry#global_endpoint_routing_enabled ContainerRegistry#global_endpoint_routing_enabled}.
	GlobalEndpointRoutingEnabled interface{} `field:"required" json:"globalEndpointRoutingEnabled" yaml:"globalEndpointRoutingEnabled"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/5.0.1/docs/resources/container_registry#location ContainerRegistry#location}.
	Location *string `field:"required" json:"location" yaml:"location"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/5.0.1/docs/resources/container_registry#tags ContainerRegistry#tags}.
	Tags *map[string]*string `field:"optional" json:"tags" yaml:"tags"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/5.0.1/docs/resources/container_registry#zone_redundancy_enabled ContainerRegistry#zone_redundancy_enabled}.
	ZoneRedundancyEnabled interface{} `field:"optional" json:"zoneRedundancyEnabled" yaml:"zoneRedundancyEnabled"`
}

