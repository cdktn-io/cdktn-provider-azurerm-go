// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package provider


type AzurermProviderFeaturesEnhancedValidation struct {
	// Should the AzureRM Provider validate location arguments against the list of supported Azure Locations?
	//
	// When enabled, invalid locations are caught at plan time; when disabled, they are caught at apply time.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/5.3.0/docs#locations AzurermProvider#locations}
	Locations interface{} `field:"optional" json:"locations" yaml:"locations"`
	// Should the AzureRM Provider call the Azure Preflight Validation API at plan time to check the request payload for each Preflight-supported resource is valid.
	//
	// Note: requires valid credentials and external Azure API access at plan-time.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/5.3.0/docs#preflight_enabled AzurermProvider#preflight_enabled}
	PreflightEnabled interface{} `field:"optional" json:"preflightEnabled" yaml:"preflightEnabled"`
	// The Azure location to use as a fallback when Preflight Validation is enabled and a resource does not specify a location.
	//
	// This is typically used for resources that derive their location from a dependency that has not yet been created.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/5.3.0/docs#preflight_location_fallback AzurermProvider#preflight_location_fallback}
	PreflightLocationFallback *string `field:"optional" json:"preflightLocationFallback" yaml:"preflightLocationFallback"`
	// Should the AzureRM Provider validate Resource Provider arguments against the list of supported Resource Providers?
	//
	// When enabled, invalid resource providers are caught at plan time; when disabled, they are caught at apply time.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/5.3.0/docs#resource_providers AzurermProvider#resource_providers}
	ResourceProviders interface{} `field:"optional" json:"resourceProviders" yaml:"resourceProviders"`
}

