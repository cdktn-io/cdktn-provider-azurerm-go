// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package apimanagementopenidconnectprovider

import (
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type ApiManagementOpenidConnectProviderConfig struct {
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
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/5.15.0/docs/resources/api_management_openid_connect_provider#api_management_name ApiManagementOpenidConnectProvider#api_management_name}.
	ApiManagementName *string `field:"required" json:"apiManagementName" yaml:"apiManagementName"`
	// The Client ID used for the Client Application.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/5.15.0/docs/resources/api_management_openid_connect_provider#client_id ApiManagementOpenidConnectProvider#client_id}
	ClientId *string `field:"required" json:"clientId" yaml:"clientId"`
	// The Client Secret used for the Client Application.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/5.15.0/docs/resources/api_management_openid_connect_provider#client_secret ApiManagementOpenidConnectProvider#client_secret}
	ClientSecret *string `field:"required" json:"clientSecret" yaml:"clientSecret"`
	// A user-friendly name for this OpenID Connect Provider.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/5.15.0/docs/resources/api_management_openid_connect_provider#display_name ApiManagementOpenidConnectProvider#display_name}
	DisplayName *string `field:"required" json:"displayName" yaml:"displayName"`
	// The URI of the Metadata endpoint.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/5.15.0/docs/resources/api_management_openid_connect_provider#metadata_endpoint ApiManagementOpenidConnectProvider#metadata_endpoint}
	MetadataEndpoint *string `field:"required" json:"metadataEndpoint" yaml:"metadataEndpoint"`
	// The name of the OpenID Connect Provider which should be created within the API Management Service.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/5.15.0/docs/resources/api_management_openid_connect_provider#name ApiManagementOpenidConnectProvider#name}
	Name *string `field:"required" json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/5.15.0/docs/resources/api_management_openid_connect_provider#resource_group_name ApiManagementOpenidConnectProvider#resource_group_name}.
	ResourceGroupName *string `field:"required" json:"resourceGroupName" yaml:"resourceGroupName"`
	// A description of this OpenID Connect Provider.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/5.15.0/docs/resources/api_management_openid_connect_provider#description ApiManagementOpenidConnectProvider#description}
	Description *string `field:"optional" json:"description" yaml:"description"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/5.15.0/docs/resources/api_management_openid_connect_provider#id ApiManagementOpenidConnectProvider#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// timeouts block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/5.15.0/docs/resources/api_management_openid_connect_provider#timeouts ApiManagementOpenidConnectProvider#timeouts}
	Timeouts *ApiManagementOpenidConnectProviderTimeouts `field:"optional" json:"timeouts" yaml:"timeouts"`
}

