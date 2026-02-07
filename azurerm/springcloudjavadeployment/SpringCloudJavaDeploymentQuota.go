// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package springcloudjavadeployment


type SpringCloudJavaDeploymentQuota struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.59.0/docs/resources/spring_cloud_java_deployment#cpu SpringCloudJavaDeployment#cpu}.
	Cpu *string `field:"optional" json:"cpu" yaml:"cpu"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.59.0/docs/resources/spring_cloud_java_deployment#memory SpringCloudJavaDeployment#memory}.
	Memory *string `field:"optional" json:"memory" yaml:"memory"`
}

