// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package apimanagementapi

import (
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type ApiManagementApiConfig struct {
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
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/5.15.0/docs/resources/api_management_api#api_management_name ApiManagementApi#api_management_name}.
	ApiManagementName *string `field:"required" json:"apiManagementName" yaml:"apiManagementName"`
	// The name of the API Management API.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/5.15.0/docs/resources/api_management_api#name ApiManagementApi#name}
	Name *string `field:"required" json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/5.15.0/docs/resources/api_management_api#resource_group_name ApiManagementApi#resource_group_name}.
	ResourceGroupName *string `field:"required" json:"resourceGroupName" yaml:"resourceGroupName"`
	// The Revision which used for this API.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/5.15.0/docs/resources/api_management_api#revision ApiManagementApi#revision}
	Revision *string `field:"required" json:"revision" yaml:"revision"`
	// Type of API.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/5.15.0/docs/resources/api_management_api#api_type ApiManagementApi#api_type}
	ApiType *string `field:"optional" json:"apiType" yaml:"apiType"`
	// contact block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/5.15.0/docs/resources/api_management_api#contact ApiManagementApi#contact}
	Contact *ApiManagementApiContact `field:"optional" json:"contact" yaml:"contact"`
	// A description of the API Management API, which may include HTML formatting tags.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/5.15.0/docs/resources/api_management_api#description ApiManagementApi#description}
	Description *string `field:"optional" json:"description" yaml:"description"`
	// The display name of the API.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/5.15.0/docs/resources/api_management_api#display_name ApiManagementApi#display_name}
	DisplayName *string `field:"optional" json:"displayName" yaml:"displayName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/5.15.0/docs/resources/api_management_api#id ApiManagementApi#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// import block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/5.15.0/docs/resources/api_management_api#import ApiManagementApi#import}
	Import *ApiManagementApiImport `field:"optional" json:"import" yaml:"import"`
	// license block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/5.15.0/docs/resources/api_management_api#license ApiManagementApi#license}
	License *ApiManagementApiLicense `field:"optional" json:"license" yaml:"license"`
	// oauth2_authorization block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/5.15.0/docs/resources/api_management_api#oauth2_authorization ApiManagementApi#oauth2_authorization}
	Oauth2Authorization *ApiManagementApiOauth2Authorization `field:"optional" json:"oauth2Authorization" yaml:"oauth2Authorization"`
	// openid_authentication block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/5.15.0/docs/resources/api_management_api#openid_authentication ApiManagementApi#openid_authentication}
	OpenidAuthentication *ApiManagementApiOpenidAuthentication `field:"optional" json:"openidAuthentication" yaml:"openidAuthentication"`
	// The Path for this API Management API, which is a relative URL which uniquely identifies this API and all of its resource paths within the API Management Service.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/5.15.0/docs/resources/api_management_api#path ApiManagementApi#path}
	Path *string `field:"optional" json:"path" yaml:"path"`
	// A list of protocols the operations in this API can be invoked.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/5.15.0/docs/resources/api_management_api#protocols ApiManagementApi#protocols}
	Protocols *[]*string `field:"optional" json:"protocols" yaml:"protocols"`
	// The description of the API Revision of the API Management API.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/5.15.0/docs/resources/api_management_api#revision_description ApiManagementApi#revision_description}
	RevisionDescription *string `field:"optional" json:"revisionDescription" yaml:"revisionDescription"`
	// Absolute URL of the backend service implementing this API.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/5.15.0/docs/resources/api_management_api#service_url ApiManagementApi#service_url}
	ServiceUrl *string `field:"optional" json:"serviceUrl" yaml:"serviceUrl"`
	// The API ID of the source API, which could be in format `azurerm_api_management_api.example.id` or in format `azurerm_api_management_api.example.id;rev=1`.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/5.15.0/docs/resources/api_management_api#source_api_id ApiManagementApi#source_api_id}
	SourceApiId *string `field:"optional" json:"sourceApiId" yaml:"sourceApiId"`
	// subscription_key_parameter_names block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/5.15.0/docs/resources/api_management_api#subscription_key_parameter_names ApiManagementApi#subscription_key_parameter_names}
	SubscriptionKeyParameterNames *ApiManagementApiSubscriptionKeyParameterNames `field:"optional" json:"subscriptionKeyParameterNames" yaml:"subscriptionKeyParameterNames"`
	// Should this API require a subscription key?
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/5.15.0/docs/resources/api_management_api#subscription_required ApiManagementApi#subscription_required}
	SubscriptionRequired interface{} `field:"optional" json:"subscriptionRequired" yaml:"subscriptionRequired"`
	// Absolute URL of the Terms of Service for the API.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/5.15.0/docs/resources/api_management_api#terms_of_service_url ApiManagementApi#terms_of_service_url}
	TermsOfServiceUrl *string `field:"optional" json:"termsOfServiceUrl" yaml:"termsOfServiceUrl"`
	// timeouts block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/5.15.0/docs/resources/api_management_api#timeouts ApiManagementApi#timeouts}
	Timeouts *ApiManagementApiTimeouts `field:"optional" json:"timeouts" yaml:"timeouts"`
	// The Version number of this API, if this API is versioned.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/5.15.0/docs/resources/api_management_api#version ApiManagementApi#version}
	Version *string `field:"optional" json:"version" yaml:"version"`
	// The description of the API Version of the API Management API.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/5.15.0/docs/resources/api_management_api#version_description ApiManagementApi#version_description}
	VersionDescription *string `field:"optional" json:"versionDescription" yaml:"versionDescription"`
	// The ID of the Version Set which this API is associated with.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/5.15.0/docs/resources/api_management_api#version_set_id ApiManagementApi#version_set_id}
	VersionSetId *string `field:"optional" json:"versionSetId" yaml:"versionSetId"`
}

