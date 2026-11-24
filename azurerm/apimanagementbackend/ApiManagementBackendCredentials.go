// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package apimanagementbackend


type ApiManagementBackendCredentials struct {
	// authorization block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/5.15.0/docs/resources/api_management_backend#authorization ApiManagementBackend#authorization}
	Authorization *ApiManagementBackendCredentialsAuthorization `field:"optional" json:"authorization" yaml:"authorization"`
	// A list of client certificate thumbprints to present to the backend host.
	//
	// The certificates must exist within the API Management Service.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/5.15.0/docs/resources/api_management_backend#certificate ApiManagementBackend#certificate}
	Certificate *[]*string `field:"optional" json:"certificate" yaml:"certificate"`
	// A mapping of header parameters to pass to the backend host.
	//
	// The keys are the header names and the values are a comma separated string of header values. This is converted to a list before being passed to the API.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/5.15.0/docs/resources/api_management_backend#header ApiManagementBackend#header}
	Header *map[string]*string `field:"optional" json:"header" yaml:"header"`
	// A mapping of query parameters to pass to the backend host.
	//
	// The keys are the query names and the values are a comma separated string of query values. This is converted to a list before being passed to the API.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/5.15.0/docs/resources/api_management_backend#query ApiManagementBackend#query}
	Query *map[string]*string `field:"optional" json:"query" yaml:"query"`
}

