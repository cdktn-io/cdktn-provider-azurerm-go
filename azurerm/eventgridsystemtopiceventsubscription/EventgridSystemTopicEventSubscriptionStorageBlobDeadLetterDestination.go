// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package eventgridsystemtopiceventsubscription


type EventgridSystemTopicEventSubscriptionStorageBlobDeadLetterDestination struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/5.1.0/docs/resources/eventgrid_system_topic_event_subscription#storage_account_id EventgridSystemTopicEventSubscription#storage_account_id}.
	StorageAccountId *string `field:"required" json:"storageAccountId" yaml:"storageAccountId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/5.1.0/docs/resources/eventgrid_system_topic_event_subscription#storage_blob_container_name EventgridSystemTopicEventSubscription#storage_blob_container_name}.
	StorageBlobContainerName *string `field:"required" json:"storageBlobContainerName" yaml:"storageBlobContainerName"`
}

