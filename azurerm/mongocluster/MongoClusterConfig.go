// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package mongocluster

import (
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type MongoClusterConfig struct {
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
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.81.0/docs/resources/mongo_cluster#location MongoCluster#location}.
	Location *string `field:"required" json:"location" yaml:"location"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.81.0/docs/resources/mongo_cluster#name MongoCluster#name}.
	Name *string `field:"required" json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.81.0/docs/resources/mongo_cluster#resource_group_name MongoCluster#resource_group_name}.
	ResourceGroupName *string `field:"required" json:"resourceGroupName" yaml:"resourceGroupName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.81.0/docs/resources/mongo_cluster#administrator_password MongoCluster#administrator_password}.
	AdministratorPassword *string `field:"optional" json:"administratorPassword" yaml:"administratorPassword"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.81.0/docs/resources/mongo_cluster#administrator_username MongoCluster#administrator_username}.
	AdministratorUsername *string `field:"optional" json:"administratorUsername" yaml:"administratorUsername"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.81.0/docs/resources/mongo_cluster#authentication_methods MongoCluster#authentication_methods}.
	AuthenticationMethods *[]*string `field:"optional" json:"authenticationMethods" yaml:"authenticationMethods"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.81.0/docs/resources/mongo_cluster#compute_tier MongoCluster#compute_tier}.
	ComputeTier *string `field:"optional" json:"computeTier" yaml:"computeTier"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.81.0/docs/resources/mongo_cluster#create_mode MongoCluster#create_mode}.
	CreateMode *string `field:"optional" json:"createMode" yaml:"createMode"`
	// customer_managed_key block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.81.0/docs/resources/mongo_cluster#customer_managed_key MongoCluster#customer_managed_key}
	CustomerManagedKey *MongoClusterCustomerManagedKey `field:"optional" json:"customerManagedKey" yaml:"customerManagedKey"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.81.0/docs/resources/mongo_cluster#data_api_mode_enabled MongoCluster#data_api_mode_enabled}.
	DataApiModeEnabled interface{} `field:"optional" json:"dataApiModeEnabled" yaml:"dataApiModeEnabled"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.81.0/docs/resources/mongo_cluster#high_availability_mode MongoCluster#high_availability_mode}.
	HighAvailabilityMode *string `field:"optional" json:"highAvailabilityMode" yaml:"highAvailabilityMode"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.81.0/docs/resources/mongo_cluster#id MongoCluster#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// identity block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.81.0/docs/resources/mongo_cluster#identity MongoCluster#identity}
	Identity *MongoClusterIdentity `field:"optional" json:"identity" yaml:"identity"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.81.0/docs/resources/mongo_cluster#preview_features MongoCluster#preview_features}.
	PreviewFeatures *[]*string `field:"optional" json:"previewFeatures" yaml:"previewFeatures"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.81.0/docs/resources/mongo_cluster#public_network_access MongoCluster#public_network_access}.
	PublicNetworkAccess *string `field:"optional" json:"publicNetworkAccess" yaml:"publicNetworkAccess"`
	// restore block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.81.0/docs/resources/mongo_cluster#restore MongoCluster#restore}
	Restore *MongoClusterRestore `field:"optional" json:"restore" yaml:"restore"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.81.0/docs/resources/mongo_cluster#shard_count MongoCluster#shard_count}.
	ShardCount *float64 `field:"optional" json:"shardCount" yaml:"shardCount"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.81.0/docs/resources/mongo_cluster#source_location MongoCluster#source_location}.
	SourceLocation *string `field:"optional" json:"sourceLocation" yaml:"sourceLocation"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.81.0/docs/resources/mongo_cluster#source_server_id MongoCluster#source_server_id}.
	SourceServerId *string `field:"optional" json:"sourceServerId" yaml:"sourceServerId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.81.0/docs/resources/mongo_cluster#storage_size_in_gb MongoCluster#storage_size_in_gb}.
	StorageSizeInGb *float64 `field:"optional" json:"storageSizeInGb" yaml:"storageSizeInGb"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.81.0/docs/resources/mongo_cluster#storage_type MongoCluster#storage_type}.
	StorageType *string `field:"optional" json:"storageType" yaml:"storageType"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.81.0/docs/resources/mongo_cluster#tags MongoCluster#tags}.
	Tags *map[string]*string `field:"optional" json:"tags" yaml:"tags"`
	// timeouts block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.81.0/docs/resources/mongo_cluster#timeouts MongoCluster#timeouts}
	Timeouts *MongoClusterTimeouts `field:"optional" json:"timeouts" yaml:"timeouts"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.81.0/docs/resources/mongo_cluster#version MongoCluster#version}.
	Version *string `field:"optional" json:"version" yaml:"version"`
}

