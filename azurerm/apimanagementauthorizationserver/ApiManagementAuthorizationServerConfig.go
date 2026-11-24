// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package apimanagementauthorizationserver

import (
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type ApiManagementAuthorizationServerConfig struct {
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
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/5.15.0/docs/resources/api_management_authorization_server#api_management_name ApiManagementAuthorizationServer#api_management_name}.
	ApiManagementName *string `field:"required" json:"apiManagementName" yaml:"apiManagementName"`
	// The OAuth Authorization Endpoint.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/5.15.0/docs/resources/api_management_authorization_server#authorization_endpoint ApiManagementAuthorizationServer#authorization_endpoint}
	AuthorizationEndpoint *string `field:"required" json:"authorizationEndpoint" yaml:"authorizationEndpoint"`
	// The HTTP Verbs supported by the Authorization Endpoint.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/5.15.0/docs/resources/api_management_authorization_server#authorization_methods ApiManagementAuthorizationServer#authorization_methods}
	AuthorizationMethods *[]*string `field:"required" json:"authorizationMethods" yaml:"authorizationMethods"`
	// The Client/App ID registered with this Authorization Server.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/5.15.0/docs/resources/api_management_authorization_server#client_id ApiManagementAuthorizationServer#client_id}
	ClientId *string `field:"required" json:"clientId" yaml:"clientId"`
	// The URI of page where Client/App Registration is performed for this Authorization Server.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/5.15.0/docs/resources/api_management_authorization_server#client_registration_endpoint ApiManagementAuthorizationServer#client_registration_endpoint}
	ClientRegistrationEndpoint *string `field:"required" json:"clientRegistrationEndpoint" yaml:"clientRegistrationEndpoint"`
	// The user-friendly name of this Authorization Server.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/5.15.0/docs/resources/api_management_authorization_server#display_name ApiManagementAuthorizationServer#display_name}
	DisplayName *string `field:"required" json:"displayName" yaml:"displayName"`
	// Form of Authorization Grants required when requesting an Access Token.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/5.15.0/docs/resources/api_management_authorization_server#grant_types ApiManagementAuthorizationServer#grant_types}
	GrantTypes *[]*string `field:"required" json:"grantTypes" yaml:"grantTypes"`
	// The name of this Authorization Server.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/5.15.0/docs/resources/api_management_authorization_server#name ApiManagementAuthorizationServer#name}
	Name *string `field:"required" json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/5.15.0/docs/resources/api_management_authorization_server#resource_group_name ApiManagementAuthorizationServer#resource_group_name}.
	ResourceGroupName *string `field:"required" json:"resourceGroupName" yaml:"resourceGroupName"`
	// The mechanism by which Access Tokens are passed to the API.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/5.15.0/docs/resources/api_management_authorization_server#bearer_token_sending_methods ApiManagementAuthorizationServer#bearer_token_sending_methods}
	BearerTokenSendingMethods *[]*string `field:"optional" json:"bearerTokenSendingMethods" yaml:"bearerTokenSendingMethods"`
	// The Authentication Methods supported by the Token endpoint of this Authorization Server.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/5.15.0/docs/resources/api_management_authorization_server#client_authentication_method ApiManagementAuthorizationServer#client_authentication_method}
	ClientAuthenticationMethod *[]*string `field:"optional" json:"clientAuthenticationMethod" yaml:"clientAuthenticationMethod"`
	// The Client/App Secret registered with this Authorization Server.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/5.15.0/docs/resources/api_management_authorization_server#client_secret ApiManagementAuthorizationServer#client_secret}
	ClientSecret *string `field:"optional" json:"clientSecret" yaml:"clientSecret"`
	// The Default Scope used when requesting an Access Token, specified as a string containing space-delimited values.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/5.15.0/docs/resources/api_management_authorization_server#default_scope ApiManagementAuthorizationServer#default_scope}
	DefaultScope *string `field:"optional" json:"defaultScope" yaml:"defaultScope"`
	// A description of the Authorization Server, which may contain HTML formatting tags.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/5.15.0/docs/resources/api_management_authorization_server#description ApiManagementAuthorizationServer#description}
	Description *string `field:"optional" json:"description" yaml:"description"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/5.15.0/docs/resources/api_management_authorization_server#id ApiManagementAuthorizationServer#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// The password associated with the Resource Owner.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/5.15.0/docs/resources/api_management_authorization_server#resource_owner_password ApiManagementAuthorizationServer#resource_owner_password}
	ResourceOwnerPassword *string `field:"optional" json:"resourceOwnerPassword" yaml:"resourceOwnerPassword"`
	// The username associated with the Resource Owner.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/5.15.0/docs/resources/api_management_authorization_server#resource_owner_username ApiManagementAuthorizationServer#resource_owner_username}
	ResourceOwnerUsername *string `field:"optional" json:"resourceOwnerUsername" yaml:"resourceOwnerUsername"`
	// Does this Authorization Server support State?
	//
	// If this is set to `true` the client may use the state parameter to raise protocol security.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/5.15.0/docs/resources/api_management_authorization_server#support_state ApiManagementAuthorizationServer#support_state}
	SupportState interface{} `field:"optional" json:"supportState" yaml:"supportState"`
	// timeouts block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/5.15.0/docs/resources/api_management_authorization_server#timeouts ApiManagementAuthorizationServer#timeouts}
	Timeouts *ApiManagementAuthorizationServerTimeouts `field:"optional" json:"timeouts" yaml:"timeouts"`
	// token_body_parameter block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/5.15.0/docs/resources/api_management_authorization_server#token_body_parameter ApiManagementAuthorizationServer#token_body_parameter}
	TokenBodyParameter interface{} `field:"optional" json:"tokenBodyParameter" yaml:"tokenBodyParameter"`
	// The OAuth Token Endpoint.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/5.15.0/docs/resources/api_management_authorization_server#token_endpoint ApiManagementAuthorizationServer#token_endpoint}
	TokenEndpoint *string `field:"optional" json:"tokenEndpoint" yaml:"tokenEndpoint"`
}

