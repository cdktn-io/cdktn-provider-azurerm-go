// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package netappvolumebucketwithserver


type NetappVolumeBucketWithServerTimeouts struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.79.0/docs/resources/netapp_volume_bucket_with_server#create NetappVolumeBucketWithServer#create}.
	Create *string `field:"optional" json:"create" yaml:"create"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.79.0/docs/resources/netapp_volume_bucket_with_server#delete NetappVolumeBucketWithServer#delete}.
	Delete *string `field:"optional" json:"delete" yaml:"delete"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.79.0/docs/resources/netapp_volume_bucket_with_server#read NetappVolumeBucketWithServer#read}.
	Read *string `field:"optional" json:"read" yaml:"read"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.79.0/docs/resources/netapp_volume_bucket_with_server#update NetappVolumeBucketWithServer#update}.
	Update *string `field:"optional" json:"update" yaml:"update"`
}

