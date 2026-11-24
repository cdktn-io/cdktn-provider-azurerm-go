// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package apimanagementlogger


type ApiManagementLoggerEventhub struct {
	// The name of an EventHub.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/5.15.0/docs/resources/api_management_logger#name ApiManagementLogger#name}
	Name *string `field:"required" json:"name" yaml:"name"`
	// The connection string of an EventHub Namespace.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/5.15.0/docs/resources/api_management_logger#connection_string ApiManagementLogger#connection_string}
	ConnectionString *string `field:"optional" json:"connectionString" yaml:"connectionString"`
	// The endpoint address of an EventHub Namespace.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/5.15.0/docs/resources/api_management_logger#endpoint_uri ApiManagementLogger#endpoint_uri}
	EndpointUri *string `field:"optional" json:"endpointUri" yaml:"endpointUri"`
	// The Client ID of the User Assigned Identity with the "Azure Event Hubs Data Sender" role to the target EventHub Namespace.
	//
	// If not specified the System Assigned Identity will be used.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/5.15.0/docs/resources/api_management_logger#user_assigned_identity_client_id ApiManagementLogger#user_assigned_identity_client_id}
	UserAssignedIdentityClientId *string `field:"optional" json:"userAssignedIdentityClientId" yaml:"userAssignedIdentityClientId"`
}

