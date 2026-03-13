// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package apimanagementworkspacenamedvalue

import (
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type ApiManagementWorkspaceNamedValueConfig struct {
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
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.64.0/docs/resources/api_management_workspace_named_value#api_management_workspace_id ApiManagementWorkspaceNamedValue#api_management_workspace_id}.
	ApiManagementWorkspaceId *string `field:"required" json:"apiManagementWorkspaceId" yaml:"apiManagementWorkspaceId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.64.0/docs/resources/api_management_workspace_named_value#display_name ApiManagementWorkspaceNamedValue#display_name}.
	DisplayName *string `field:"required" json:"displayName" yaml:"displayName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.64.0/docs/resources/api_management_workspace_named_value#name ApiManagementWorkspaceNamedValue#name}.
	Name *string `field:"required" json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.64.0/docs/resources/api_management_workspace_named_value#id ApiManagementWorkspaceNamedValue#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.64.0/docs/resources/api_management_workspace_named_value#secret ApiManagementWorkspaceNamedValue#secret}.
	Secret interface{} `field:"optional" json:"secret" yaml:"secret"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.64.0/docs/resources/api_management_workspace_named_value#tags ApiManagementWorkspaceNamedValue#tags}.
	Tags *[]*string `field:"optional" json:"tags" yaml:"tags"`
	// timeouts block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.64.0/docs/resources/api_management_workspace_named_value#timeouts ApiManagementWorkspaceNamedValue#timeouts}
	Timeouts *ApiManagementWorkspaceNamedValueTimeouts `field:"optional" json:"timeouts" yaml:"timeouts"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.64.0/docs/resources/api_management_workspace_named_value#value ApiManagementWorkspaceNamedValue#value}.
	Value *string `field:"optional" json:"value" yaml:"value"`
	// value_from_key_vault block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.64.0/docs/resources/api_management_workspace_named_value#value_from_key_vault ApiManagementWorkspaceNamedValue#value_from_key_vault}
	ValueFromKeyVault *ApiManagementWorkspaceNamedValueValueFromKeyVault `field:"optional" json:"valueFromKeyVault" yaml:"valueFromKeyVault"`
}

