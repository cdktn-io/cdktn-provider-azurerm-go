// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package apimanagementpolicy

import (
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type ApiManagementPolicyConfig struct {
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
	// The ID of the API Management service.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/5.15.0/docs/resources/api_management_policy#api_management_id ApiManagementPolicy#api_management_id}
	ApiManagementId *string `field:"required" json:"apiManagementId" yaml:"apiManagementId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/5.15.0/docs/resources/api_management_policy#id ApiManagementPolicy#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// timeouts block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/5.15.0/docs/resources/api_management_policy#timeouts ApiManagementPolicy#timeouts}
	Timeouts *ApiManagementPolicyTimeouts `field:"optional" json:"timeouts" yaml:"timeouts"`
	// The XML Content for this Policy as a string.
	//
	// An XML file can be used here with Terraform's [file function](https://developer.hashicorp.com/terraform/language/functions/file) that is similar to Microsoft's `PolicyFilePath` option. To integrate frontend and backend services in Azure API Management, utilize the [`set-backend-service`](https://learn.microsoft.com/azure/api-management/set-backend-service-policy) policy, specifying the `base-url` value. Typically, this value corresponds to the `url` property defined in the [`azurerm_api_management_backend`](https://registry.terraform.io/providers/hashicorp/azurerm/latest/docs/resources/api_management_backend) configuration.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/5.15.0/docs/resources/api_management_policy#xml_content ApiManagementPolicy#xml_content}
	XmlContent *string `field:"optional" json:"xmlContent" yaml:"xmlContent"`
	// A link to a Policy XML Document, which must be publicly available.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/5.15.0/docs/resources/api_management_policy#xml_link ApiManagementPolicy#xml_link}
	XmlLink *string `field:"optional" json:"xmlLink" yaml:"xmlLink"`
}

