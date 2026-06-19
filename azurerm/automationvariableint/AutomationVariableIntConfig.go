// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package automationvariableint

import (
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type AutomationVariableIntConfig struct {
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
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.78.0/docs/resources/automation_variable_int#automation_account_name AutomationVariableInt#automation_account_name}.
	AutomationAccountName *string `field:"required" json:"automationAccountName" yaml:"automationAccountName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.78.0/docs/resources/automation_variable_int#name AutomationVariableInt#name}.
	Name *string `field:"required" json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.78.0/docs/resources/automation_variable_int#resource_group_name AutomationVariableInt#resource_group_name}.
	ResourceGroupName *string `field:"required" json:"resourceGroupName" yaml:"resourceGroupName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.78.0/docs/resources/automation_variable_int#description AutomationVariableInt#description}.
	Description *string `field:"optional" json:"description" yaml:"description"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.78.0/docs/resources/automation_variable_int#encrypted AutomationVariableInt#encrypted}.
	Encrypted interface{} `field:"optional" json:"encrypted" yaml:"encrypted"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.78.0/docs/resources/automation_variable_int#id AutomationVariableInt#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// timeouts block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.78.0/docs/resources/automation_variable_int#timeouts AutomationVariableInt#timeouts}
	Timeouts *AutomationVariableIntTimeouts `field:"optional" json:"timeouts" yaml:"timeouts"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.78.0/docs/resources/automation_variable_int#value AutomationVariableInt#value}.
	Value *float64 `field:"optional" json:"value" yaml:"value"`
}

