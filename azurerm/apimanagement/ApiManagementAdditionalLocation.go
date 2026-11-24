// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package apimanagement


type ApiManagementAdditionalLocation struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/5.15.0/docs/resources/api_management#location ApiManagement#location}.
	Location *string `field:"required" json:"location" yaml:"location"`
	// The number of compute units in this region.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/5.15.0/docs/resources/api_management#capacity ApiManagement#capacity}
	Capacity *float64 `field:"optional" json:"capacity" yaml:"capacity"`
	// Only valid for an API Management service deployed in multiple locations.
	//
	// This can be used to disable the gateway in this additional location.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/5.15.0/docs/resources/api_management#gateway_disabled ApiManagement#gateway_disabled}
	GatewayDisabled interface{} `field:"optional" json:"gatewayDisabled" yaml:"gatewayDisabled"`
	// ID of a standard SKU IPv4 Public IP.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/5.15.0/docs/resources/api_management#public_ip_address_id ApiManagement#public_ip_address_id}
	PublicIpAddressId *string `field:"optional" json:"publicIpAddressId" yaml:"publicIpAddressId"`
	// virtual_network_configuration block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/5.15.0/docs/resources/api_management#virtual_network_configuration ApiManagement#virtual_network_configuration}
	VirtualNetworkConfiguration *ApiManagementAdditionalLocationVirtualNetworkConfiguration `field:"optional" json:"virtualNetworkConfiguration" yaml:"virtualNetworkConfiguration"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/5.15.0/docs/resources/api_management#zones ApiManagement#zones}.
	Zones *[]*string `field:"optional" json:"zones" yaml:"zones"`
}

