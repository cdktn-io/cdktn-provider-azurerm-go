// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package springcloudgateway


type SpringCloudGatewayQuota struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.79.0/docs/resources/spring_cloud_gateway#cpu SpringCloudGateway#cpu}.
	Cpu *string `field:"optional" json:"cpu" yaml:"cpu"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.79.0/docs/resources/spring_cloud_gateway#memory SpringCloudGateway#memory}.
	Memory *string `field:"optional" json:"memory" yaml:"memory"`
}

