// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package containerregistrytask


type ContainerRegistryTaskAgentSetting struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.64.0/docs/resources/container_registry_task#cpu ContainerRegistryTask#cpu}.
	Cpu *float64 `field:"required" json:"cpu" yaml:"cpu"`
}

