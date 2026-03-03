// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package dataazurermconsumptionbudgetresourcegroup

import (
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type DataAzurermConsumptionBudgetResourceGroupConfig struct {
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
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.62.1/docs/data-sources/consumption_budget_resource_group#name DataAzurermConsumptionBudgetResourceGroup#name}.
	Name *string `field:"required" json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.62.1/docs/data-sources/consumption_budget_resource_group#resource_group_id DataAzurermConsumptionBudgetResourceGroup#resource_group_id}.
	ResourceGroupId *string `field:"required" json:"resourceGroupId" yaml:"resourceGroupId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.62.1/docs/data-sources/consumption_budget_resource_group#id DataAzurermConsumptionBudgetResourceGroup#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// timeouts block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.62.1/docs/data-sources/consumption_budget_resource_group#timeouts DataAzurermConsumptionBudgetResourceGroup#timeouts}
	Timeouts *DataAzurermConsumptionBudgetResourceGroupTimeouts `field:"optional" json:"timeouts" yaml:"timeouts"`
}

