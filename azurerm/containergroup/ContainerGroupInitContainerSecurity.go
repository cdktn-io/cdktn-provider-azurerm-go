// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package containergroup


type ContainerGroupInitContainerSecurity struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.81.0/docs/resources/container_group#privilege_enabled ContainerGroup#privilege_enabled}.
	PrivilegeEnabled interface{} `field:"required" json:"privilegeEnabled" yaml:"privilegeEnabled"`
}

