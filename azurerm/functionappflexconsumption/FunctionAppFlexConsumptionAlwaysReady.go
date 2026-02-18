// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package functionappflexconsumption


type FunctionAppFlexConsumptionAlwaysReady struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.60.0/docs/resources/function_app_flex_consumption#name FunctionAppFlexConsumption#name}.
	Name *string `field:"required" json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.60.0/docs/resources/function_app_flex_consumption#instance_count FunctionAppFlexConsumption#instance_count}.
	InstanceCount *float64 `field:"optional" json:"instanceCount" yaml:"instanceCount"`
}

