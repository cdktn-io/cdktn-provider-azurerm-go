// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package apimanagementapi


type ApiManagementApiImport struct {
	// The format of the content from which the API Definition should be imported.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/5.15.0/docs/resources/api_management_api#content_format ApiManagementApi#content_format}
	ContentFormat *string `field:"required" json:"contentFormat" yaml:"contentFormat"`
	// The Content from which the API Definition should be imported.
	//
	// When a `content_format` of `*-link-*` is specified this must be a URL, otherwise this must be defined inline. The URL must be accessible and return a valid document; otherwise, deployment may fail.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/5.15.0/docs/resources/api_management_api#content_value ApiManagementApi#content_value}
	ContentValue *string `field:"required" json:"contentValue" yaml:"contentValue"`
	// wsdl_selector block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/5.15.0/docs/resources/api_management_api#wsdl_selector ApiManagementApi#wsdl_selector}
	WsdlSelector *ApiManagementApiImportWsdlSelector `field:"optional" json:"wsdlSelector" yaml:"wsdlSelector"`
}

