// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package siterecoveryreplicatedvm


type SiteRecoveryReplicatedVmNetworkInterface struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/5.0.1/docs/resources/site_recovery_replicated_vm#ip_configuration SiteRecoveryReplicatedVm#ip_configuration}.
	IpConfiguration interface{} `field:"optional" json:"ipConfiguration" yaml:"ipConfiguration"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/5.0.1/docs/resources/site_recovery_replicated_vm#source_network_interface_id SiteRecoveryReplicatedVm#source_network_interface_id}.
	SourceNetworkInterfaceId *string `field:"optional" json:"sourceNetworkInterfaceId" yaml:"sourceNetworkInterfaceId"`
}

