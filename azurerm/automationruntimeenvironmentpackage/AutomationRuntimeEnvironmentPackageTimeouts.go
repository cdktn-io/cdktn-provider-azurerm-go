// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package automationruntimeenvironmentpackage


type AutomationRuntimeEnvironmentPackageTimeouts struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/5.4.0/docs/resources/automation_runtime_environment_package#create AutomationRuntimeEnvironmentPackage#create}.
	Create *string `field:"optional" json:"create" yaml:"create"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/5.4.0/docs/resources/automation_runtime_environment_package#delete AutomationRuntimeEnvironmentPackage#delete}.
	Delete *string `field:"optional" json:"delete" yaml:"delete"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/5.4.0/docs/resources/automation_runtime_environment_package#read AutomationRuntimeEnvironmentPackage#read}.
	Read *string `field:"optional" json:"read" yaml:"read"`
}

