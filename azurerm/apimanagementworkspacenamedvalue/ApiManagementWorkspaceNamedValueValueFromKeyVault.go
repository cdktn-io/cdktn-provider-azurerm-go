// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package apimanagementworkspacenamedvalue


type ApiManagementWorkspaceNamedValueValueFromKeyVault struct {
	// The resource ID of the Key Vault Secret.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/5.15.0/docs/resources/api_management_workspace_named_value#secret_id ApiManagementWorkspaceNamedValue#secret_id}
	SecretId *string `field:"required" json:"secretId" yaml:"secretId"`
	// The client ID of the User Assigned Identity, for the API Management Service, which will be used to access the Key Vault Secret.
	//
	// The System Assigned Identity will be used if not specified.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/5.15.0/docs/resources/api_management_workspace_named_value#identity_client_id ApiManagementWorkspaceNamedValue#identity_client_id}
	IdentityClientId *string `field:"optional" json:"identityClientId" yaml:"identityClientId"`
}

