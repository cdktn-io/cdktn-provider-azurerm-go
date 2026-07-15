// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package netappvolumebucketwithserver

import (
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type NetappVolumeBucketWithServerConfig struct {
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
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.81.0/docs/resources/netapp_volume_bucket_with_server#name NetappVolumeBucketWithServer#name}.
	Name *string `field:"required" json:"name" yaml:"name"`
	// server block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.81.0/docs/resources/netapp_volume_bucket_with_server#server NetappVolumeBucketWithServer#server}
	Server *NetappVolumeBucketWithServerServer `field:"required" json:"server" yaml:"server"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.81.0/docs/resources/netapp_volume_bucket_with_server#volume_id NetappVolumeBucketWithServer#volume_id}.
	VolumeId *string `field:"required" json:"volumeId" yaml:"volumeId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.81.0/docs/resources/netapp_volume_bucket_with_server#file_system_cifs_username NetappVolumeBucketWithServer#file_system_cifs_username}.
	FileSystemCifsUsername *string `field:"optional" json:"fileSystemCifsUsername" yaml:"fileSystemCifsUsername"`
	// file_system_nfs_user block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.81.0/docs/resources/netapp_volume_bucket_with_server#file_system_nfs_user NetappVolumeBucketWithServer#file_system_nfs_user}
	FileSystemNfsUser *NetappVolumeBucketWithServerFileSystemNfsUser `field:"optional" json:"fileSystemNfsUser" yaml:"fileSystemNfsUser"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.81.0/docs/resources/netapp_volume_bucket_with_server#id NetappVolumeBucketWithServer#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// key_vault block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.81.0/docs/resources/netapp_volume_bucket_with_server#key_vault NetappVolumeBucketWithServer#key_vault}
	KeyVault *NetappVolumeBucketWithServerKeyVault `field:"optional" json:"keyVault" yaml:"keyVault"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.81.0/docs/resources/netapp_volume_bucket_with_server#path NetappVolumeBucketWithServer#path}.
	Path *string `field:"optional" json:"path" yaml:"path"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.81.0/docs/resources/netapp_volume_bucket_with_server#permissions NetappVolumeBucketWithServer#permissions}.
	Permissions *string `field:"optional" json:"permissions" yaml:"permissions"`
	// timeouts block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.81.0/docs/resources/netapp_volume_bucket_with_server#timeouts NetappVolumeBucketWithServer#timeouts}
	Timeouts *NetappVolumeBucketWithServerTimeouts `field:"optional" json:"timeouts" yaml:"timeouts"`
}

