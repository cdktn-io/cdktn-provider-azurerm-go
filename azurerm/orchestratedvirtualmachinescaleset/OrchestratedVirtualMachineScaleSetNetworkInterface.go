// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package orchestratedvirtualmachinescaleset


type OrchestratedVirtualMachineScaleSetNetworkInterface struct {
	// ip_configuration block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/5.1.0/docs/resources/orchestrated_virtual_machine_scale_set#ip_configuration OrchestratedVirtualMachineScaleSet#ip_configuration}
	IpConfiguration interface{} `field:"required" json:"ipConfiguration" yaml:"ipConfiguration"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/5.1.0/docs/resources/orchestrated_virtual_machine_scale_set#name OrchestratedVirtualMachineScaleSet#name}.
	Name *string `field:"required" json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/5.1.0/docs/resources/orchestrated_virtual_machine_scale_set#accelerated_networking_enabled OrchestratedVirtualMachineScaleSet#accelerated_networking_enabled}.
	AcceleratedNetworkingEnabled interface{} `field:"optional" json:"acceleratedNetworkingEnabled" yaml:"acceleratedNetworkingEnabled"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/5.1.0/docs/resources/orchestrated_virtual_machine_scale_set#auxiliary_mode OrchestratedVirtualMachineScaleSet#auxiliary_mode}.
	AuxiliaryMode *string `field:"optional" json:"auxiliaryMode" yaml:"auxiliaryMode"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/5.1.0/docs/resources/orchestrated_virtual_machine_scale_set#auxiliary_sku OrchestratedVirtualMachineScaleSet#auxiliary_sku}.
	AuxiliarySku *string `field:"optional" json:"auxiliarySku" yaml:"auxiliarySku"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/5.1.0/docs/resources/orchestrated_virtual_machine_scale_set#dns_servers OrchestratedVirtualMachineScaleSet#dns_servers}.
	DnsServers *[]*string `field:"optional" json:"dnsServers" yaml:"dnsServers"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/5.1.0/docs/resources/orchestrated_virtual_machine_scale_set#ip_forwarding_enabled OrchestratedVirtualMachineScaleSet#ip_forwarding_enabled}.
	IpForwardingEnabled interface{} `field:"optional" json:"ipForwardingEnabled" yaml:"ipForwardingEnabled"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/5.1.0/docs/resources/orchestrated_virtual_machine_scale_set#network_security_group_id OrchestratedVirtualMachineScaleSet#network_security_group_id}.
	NetworkSecurityGroupId *string `field:"optional" json:"networkSecurityGroupId" yaml:"networkSecurityGroupId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/5.1.0/docs/resources/orchestrated_virtual_machine_scale_set#primary OrchestratedVirtualMachineScaleSet#primary}.
	Primary interface{} `field:"optional" json:"primary" yaml:"primary"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/5.1.0/docs/resources/orchestrated_virtual_machine_scale_set#tags OrchestratedVirtualMachineScaleSet#tags}.
	Tags *map[string]*string `field:"optional" json:"tags" yaml:"tags"`
}

