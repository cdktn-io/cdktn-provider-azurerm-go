// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package apimanagementapipolicy

import (
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type ApiManagementApiPolicyConfig struct {
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
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/5.15.0/docs/resources/api_management_api_policy#api_management_name ApiManagementApiPolicy#api_management_name}.
	ApiManagementName *string `field:"required" json:"apiManagementName" yaml:"apiManagementName"`
	// The ID of the API Management API within the API Management Service.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/5.15.0/docs/resources/api_management_api_policy#api_name ApiManagementApiPolicy#api_name}
	ApiName *string `field:"required" json:"apiName" yaml:"apiName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/5.15.0/docs/resources/api_management_api_policy#resource_group_name ApiManagementApiPolicy#resource_group_name}.
	ResourceGroupName *string `field:"required" json:"resourceGroupName" yaml:"resourceGroupName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/5.15.0/docs/resources/api_management_api_policy#id ApiManagementApiPolicy#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// timeouts block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/5.15.0/docs/resources/api_management_api_policy#timeouts ApiManagementApiPolicy#timeouts}
	Timeouts *ApiManagementApiPolicyTimeouts `field:"optional" json:"timeouts" yaml:"timeouts"`
	// The XML Content for this Policy as a string.
	//
	// An XML file can be used here with Terraform's [file function](https://developer.hashicorp.com/terraform/language/functions/file) that is similar to Microsoft's `PolicyFilePath` option. If you need to pass variables into your XML file, use Terraform's [templatefile function](https://developer.hashicorp.com/terraform/language/functions/templatefile).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/5.15.0/docs/resources/api_management_api_policy#xml_content ApiManagementApiPolicy#xml_content}
	XmlContent *string `field:"optional" json:"xmlContent" yaml:"xmlContent"`
	// A link to a Policy XML Document, which must be publicly available.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/5.15.0/docs/resources/api_management_api_policy#xml_link ApiManagementApiPolicy#xml_link}
	XmlLink *string `field:"optional" json:"xmlLink" yaml:"xmlLink"`
}

