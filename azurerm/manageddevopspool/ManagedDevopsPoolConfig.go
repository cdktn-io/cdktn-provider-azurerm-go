// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package manageddevopspool

import (
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type ManagedDevopsPoolConfig struct {
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
	// azure_devops_organization block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.79.0/docs/resources/managed_devops_pool#azure_devops_organization ManagedDevopsPool#azure_devops_organization}
	AzureDevopsOrganization *ManagedDevopsPoolAzureDevopsOrganization `field:"required" json:"azureDevopsOrganization" yaml:"azureDevopsOrganization"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.79.0/docs/resources/managed_devops_pool#dev_center_project_id ManagedDevopsPool#dev_center_project_id}.
	DevCenterProjectId *string `field:"required" json:"devCenterProjectId" yaml:"devCenterProjectId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.79.0/docs/resources/managed_devops_pool#location ManagedDevopsPool#location}.
	Location *string `field:"required" json:"location" yaml:"location"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.79.0/docs/resources/managed_devops_pool#maximum_concurrency ManagedDevopsPool#maximum_concurrency}.
	MaximumConcurrency *float64 `field:"required" json:"maximumConcurrency" yaml:"maximumConcurrency"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.79.0/docs/resources/managed_devops_pool#name ManagedDevopsPool#name}.
	Name *string `field:"required" json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.79.0/docs/resources/managed_devops_pool#resource_group_name ManagedDevopsPool#resource_group_name}.
	ResourceGroupName *string `field:"required" json:"resourceGroupName" yaml:"resourceGroupName"`
	// virtual_machine_scale_set_fabric block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.79.0/docs/resources/managed_devops_pool#virtual_machine_scale_set_fabric ManagedDevopsPool#virtual_machine_scale_set_fabric}
	VirtualMachineScaleSetFabric *ManagedDevopsPoolVirtualMachineScaleSetFabric `field:"required" json:"virtualMachineScaleSetFabric" yaml:"virtualMachineScaleSetFabric"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.79.0/docs/resources/managed_devops_pool#id ManagedDevopsPool#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// identity block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.79.0/docs/resources/managed_devops_pool#identity ManagedDevopsPool#identity}
	Identity *ManagedDevopsPoolIdentity `field:"optional" json:"identity" yaml:"identity"`
	// stateful_agent block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.79.0/docs/resources/managed_devops_pool#stateful_agent ManagedDevopsPool#stateful_agent}
	StatefulAgent *ManagedDevopsPoolStatefulAgent `field:"optional" json:"statefulAgent" yaml:"statefulAgent"`
	// stateless_agent block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.79.0/docs/resources/managed_devops_pool#stateless_agent ManagedDevopsPool#stateless_agent}
	StatelessAgent *ManagedDevopsPoolStatelessAgent `field:"optional" json:"statelessAgent" yaml:"statelessAgent"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.79.0/docs/resources/managed_devops_pool#tags ManagedDevopsPool#tags}.
	Tags *map[string]*string `field:"optional" json:"tags" yaml:"tags"`
	// timeouts block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.79.0/docs/resources/managed_devops_pool#timeouts ManagedDevopsPool#timeouts}
	Timeouts *ManagedDevopsPoolTimeouts `field:"optional" json:"timeouts" yaml:"timeouts"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.79.0/docs/resources/managed_devops_pool#work_folder ManagedDevopsPool#work_folder}.
	WorkFolder *string `field:"optional" json:"workFolder" yaml:"workFolder"`
}

