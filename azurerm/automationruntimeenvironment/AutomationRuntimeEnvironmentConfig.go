// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package automationruntimeenvironment

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type AutomationRuntimeEnvironmentConfig struct {
	// Experimental.
	Connection interface{} `field:"optional" json:"connection" yaml:"connection"`
	// Experimental.
	Count interface{} `field:"optional" json:"count" yaml:"count"`
	// Experimental.
	DependsOn *[]cdktf.ITerraformDependable `field:"optional" json:"dependsOn" yaml:"dependsOn"`
	// Experimental.
	ForEach cdktf.ITerraformIterator `field:"optional" json:"forEach" yaml:"forEach"`
	// Experimental.
	Lifecycle *cdktf.TerraformResourceLifecycle `field:"optional" json:"lifecycle" yaml:"lifecycle"`
	// Experimental.
	Provider cdktf.TerraformProvider `field:"optional" json:"provider" yaml:"provider"`
	// Experimental.
	Provisioners *[]interface{} `field:"optional" json:"provisioners" yaml:"provisioners"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.59.0/docs/resources/automation_runtime_environment#automation_account_id AutomationRuntimeEnvironment#automation_account_id}.
	AutomationAccountId *string `field:"required" json:"automationAccountId" yaml:"automationAccountId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.59.0/docs/resources/automation_runtime_environment#location AutomationRuntimeEnvironment#location}.
	Location *string `field:"required" json:"location" yaml:"location"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.59.0/docs/resources/automation_runtime_environment#name AutomationRuntimeEnvironment#name}.
	Name *string `field:"required" json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.59.0/docs/resources/automation_runtime_environment#runtime_language AutomationRuntimeEnvironment#runtime_language}.
	RuntimeLanguage *string `field:"required" json:"runtimeLanguage" yaml:"runtimeLanguage"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.59.0/docs/resources/automation_runtime_environment#runtime_version AutomationRuntimeEnvironment#runtime_version}.
	RuntimeVersion *string `field:"required" json:"runtimeVersion" yaml:"runtimeVersion"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.59.0/docs/resources/automation_runtime_environment#description AutomationRuntimeEnvironment#description}.
	Description *string `field:"optional" json:"description" yaml:"description"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.59.0/docs/resources/automation_runtime_environment#id AutomationRuntimeEnvironment#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.59.0/docs/resources/automation_runtime_environment#runtime_default_packages AutomationRuntimeEnvironment#runtime_default_packages}.
	RuntimeDefaultPackages *map[string]*string `field:"optional" json:"runtimeDefaultPackages" yaml:"runtimeDefaultPackages"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.59.0/docs/resources/automation_runtime_environment#tags AutomationRuntimeEnvironment#tags}.
	Tags *map[string]*string `field:"optional" json:"tags" yaml:"tags"`
	// timeouts block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.59.0/docs/resources/automation_runtime_environment#timeouts AutomationRuntimeEnvironment#timeouts}
	Timeouts *AutomationRuntimeEnvironmentTimeouts `field:"optional" json:"timeouts" yaml:"timeouts"`
}

