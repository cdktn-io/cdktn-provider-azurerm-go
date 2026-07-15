// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package automationruntimeenvironmentpackage

import (
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type AutomationRuntimeEnvironmentPackageConfig struct {
	// Experimental.
	Connection interface{} `field:"optional" json:"connection" yaml:"connection"`
	// Experimental.
	Count interface{} `field:"optional" json:"count" yaml:"count"`
	// Experimental.
	DependsOn *[]cdktn.ITerraformDependable `field:"optional" json:"dependsOn" yaml:"dependsOn"`
	// Experimental.
	ForEach cdktn.ITerraformIterator `field:"optional" json:"forEach" yaml:"forEach"`
	// Experimental.
	Lifecycle *cdktn.TerraformResourceLifecycle `field:"optional" json:"lifecycle" yaml:"lifecycle"`
	// Experimental.
	Provider cdktn.TerraformProvider `field:"optional" json:"provider" yaml:"provider"`
	// Experimental.
	Provisioners *[]interface{} `field:"optional" json:"provisioners" yaml:"provisioners"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.81.0/docs/resources/automation_runtime_environment_package#automation_runtime_environment_id AutomationRuntimeEnvironmentPackage#automation_runtime_environment_id}.
	AutomationRuntimeEnvironmentId *string `field:"required" json:"automationRuntimeEnvironmentId" yaml:"automationRuntimeEnvironmentId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.81.0/docs/resources/automation_runtime_environment_package#content_uri AutomationRuntimeEnvironmentPackage#content_uri}.
	ContentUri *string `field:"required" json:"contentUri" yaml:"contentUri"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.81.0/docs/resources/automation_runtime_environment_package#name AutomationRuntimeEnvironmentPackage#name}.
	Name *string `field:"required" json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.81.0/docs/resources/automation_runtime_environment_package#content_version AutomationRuntimeEnvironmentPackage#content_version}.
	ContentVersion *string `field:"optional" json:"contentVersion" yaml:"contentVersion"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.81.0/docs/resources/automation_runtime_environment_package#hash_algorithm AutomationRuntimeEnvironmentPackage#hash_algorithm}.
	HashAlgorithm *string `field:"optional" json:"hashAlgorithm" yaml:"hashAlgorithm"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.81.0/docs/resources/automation_runtime_environment_package#hash_value AutomationRuntimeEnvironmentPackage#hash_value}.
	HashValue *string `field:"optional" json:"hashValue" yaml:"hashValue"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.81.0/docs/resources/automation_runtime_environment_package#id AutomationRuntimeEnvironmentPackage#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// timeouts block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.81.0/docs/resources/automation_runtime_environment_package#timeouts AutomationRuntimeEnvironmentPackage#timeouts}
	Timeouts *AutomationRuntimeEnvironmentPackageTimeouts `field:"optional" json:"timeouts" yaml:"timeouts"`
}

