// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package eventhub


type EventhubCaptureDescriptionDestination struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.79.0/docs/resources/eventhub#archive_name_format Eventhub#archive_name_format}.
	ArchiveNameFormat *string `field:"required" json:"archiveNameFormat" yaml:"archiveNameFormat"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.79.0/docs/resources/eventhub#blob_container_name Eventhub#blob_container_name}.
	BlobContainerName *string `field:"required" json:"blobContainerName" yaml:"blobContainerName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.79.0/docs/resources/eventhub#name Eventhub#name}.
	Name *string `field:"required" json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.79.0/docs/resources/eventhub#storage_account_id Eventhub#storage_account_id}.
	StorageAccountId *string `field:"required" json:"storageAccountId" yaml:"storageAccountId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.79.0/docs/resources/eventhub#storage_authentication_id Eventhub#storage_authentication_id}.
	StorageAuthenticationId *string `field:"optional" json:"storageAuthenticationId" yaml:"storageAuthenticationId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.79.0/docs/resources/eventhub#storage_authentication_type Eventhub#storage_authentication_type}.
	StorageAuthenticationType *string `field:"optional" json:"storageAuthenticationType" yaml:"storageAuthenticationType"`
}

