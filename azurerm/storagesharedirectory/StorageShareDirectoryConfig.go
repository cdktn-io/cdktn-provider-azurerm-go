// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package storagesharedirectory

import (
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type StorageShareDirectoryConfig struct {
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
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/5.0.1/docs/resources/storage_share_directory#name StorageShareDirectory#name}.
	Name *string `field:"required" json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/5.0.1/docs/resources/storage_share_directory#storage_share_url StorageShareDirectory#storage_share_url}.
	StorageShareUrl *string `field:"required" json:"storageShareUrl" yaml:"storageShareUrl"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/5.0.1/docs/resources/storage_share_directory#id StorageShareDirectory#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/5.0.1/docs/resources/storage_share_directory#metadata StorageShareDirectory#metadata}.
	Metadata *map[string]*string `field:"optional" json:"metadata" yaml:"metadata"`
	// timeouts block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/5.0.1/docs/resources/storage_share_directory#timeouts StorageShareDirectory#timeouts}
	Timeouts *StorageShareDirectoryTimeouts `field:"optional" json:"timeouts" yaml:"timeouts"`
}

