// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package apimanagementapi


type ApiManagementApiImportWsdlSelector struct {
	// The name of endpoint (port) to import from WSDL.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/5.15.0/docs/resources/api_management_api#endpoint_name ApiManagementApi#endpoint_name}
	EndpointName *string `field:"required" json:"endpointName" yaml:"endpointName"`
	// The name of service to import from WSDL.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/5.15.0/docs/resources/api_management_api#service_name ApiManagementApi#service_name}
	ServiceName *string `field:"required" json:"serviceName" yaml:"serviceName"`
}

