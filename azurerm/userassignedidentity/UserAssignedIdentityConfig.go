// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package userassignedidentity

import (
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type UserAssignedIdentityConfig struct {
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
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.60.0/docs/resources/user_assigned_identity#location UserAssignedIdentity#location}.
	Location *string `field:"required" json:"location" yaml:"location"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.60.0/docs/resources/user_assigned_identity#name UserAssignedIdentity#name}.
	Name *string `field:"required" json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.60.0/docs/resources/user_assigned_identity#resource_group_name UserAssignedIdentity#resource_group_name}.
	ResourceGroupName *string `field:"required" json:"resourceGroupName" yaml:"resourceGroupName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.60.0/docs/resources/user_assigned_identity#id UserAssignedIdentity#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.60.0/docs/resources/user_assigned_identity#isolation_scope UserAssignedIdentity#isolation_scope}.
	IsolationScope *string `field:"optional" json:"isolationScope" yaml:"isolationScope"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.60.0/docs/resources/user_assigned_identity#tags UserAssignedIdentity#tags}.
	Tags *map[string]*string `field:"optional" json:"tags" yaml:"tags"`
	// timeouts block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.60.0/docs/resources/user_assigned_identity#timeouts UserAssignedIdentity#timeouts}
	Timeouts *UserAssignedIdentityTimeouts `field:"optional" json:"timeouts" yaml:"timeouts"`
}

