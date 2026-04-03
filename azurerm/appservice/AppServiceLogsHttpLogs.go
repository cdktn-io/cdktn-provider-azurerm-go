// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package appservice


type AppServiceLogsHttpLogs struct {
	// azure_blob_storage block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.67.0/docs/resources/app_service#azure_blob_storage AppService#azure_blob_storage}
	AzureBlobStorage *AppServiceLogsHttpLogsAzureBlobStorage `field:"optional" json:"azureBlobStorage" yaml:"azureBlobStorage"`
	// file_system block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.67.0/docs/resources/app_service#file_system AppService#file_system}
	FileSystem *AppServiceLogsHttpLogsFileSystem `field:"optional" json:"fileSystem" yaml:"fileSystem"`
}

