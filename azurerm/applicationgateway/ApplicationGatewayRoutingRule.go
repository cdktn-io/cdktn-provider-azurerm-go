// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package applicationgateway


type ApplicationGatewayRoutingRule struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.75.0/docs/resources/application_gateway#backend_address_pool_name ApplicationGateway#backend_address_pool_name}.
	BackendAddressPoolName *string `field:"required" json:"backendAddressPoolName" yaml:"backendAddressPoolName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.75.0/docs/resources/application_gateway#backend_name ApplicationGateway#backend_name}.
	BackendName *string `field:"required" json:"backendName" yaml:"backendName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.75.0/docs/resources/application_gateway#listener_name ApplicationGateway#listener_name}.
	ListenerName *string `field:"required" json:"listenerName" yaml:"listenerName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.75.0/docs/resources/application_gateway#name ApplicationGateway#name}.
	Name *string `field:"required" json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.75.0/docs/resources/application_gateway#priority ApplicationGateway#priority}.
	Priority *float64 `field:"required" json:"priority" yaml:"priority"`
}

