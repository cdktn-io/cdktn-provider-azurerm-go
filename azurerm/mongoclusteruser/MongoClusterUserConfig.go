// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package mongoclusteruser

import (
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type MongoClusterUserConfig struct {
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
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.80.0/docs/resources/mongo_cluster_user#identity_provider_type MongoClusterUser#identity_provider_type}.
	IdentityProviderType *string `field:"required" json:"identityProviderType" yaml:"identityProviderType"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.80.0/docs/resources/mongo_cluster_user#mongo_cluster_id MongoClusterUser#mongo_cluster_id}.
	MongoClusterId *string `field:"required" json:"mongoClusterId" yaml:"mongoClusterId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.80.0/docs/resources/mongo_cluster_user#object_id MongoClusterUser#object_id}.
	ObjectId *string `field:"required" json:"objectId" yaml:"objectId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.80.0/docs/resources/mongo_cluster_user#principal_type MongoClusterUser#principal_type}.
	PrincipalType *string `field:"required" json:"principalType" yaml:"principalType"`
	// role block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.80.0/docs/resources/mongo_cluster_user#role MongoClusterUser#role}
	Role interface{} `field:"required" json:"role" yaml:"role"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.80.0/docs/resources/mongo_cluster_user#id MongoClusterUser#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// timeouts block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.80.0/docs/resources/mongo_cluster_user#timeouts MongoClusterUser#timeouts}
	Timeouts *MongoClusterUserTimeouts `field:"optional" json:"timeouts" yaml:"timeouts"`
}

