// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package provider


type AzurermProviderEnhancedValidation struct {
	// Should the AzureRM Provider validate location arguments against the list of supported Azure Locations?
	//
	// When enabled, invalid locations are caught at plan time; when disabled, they are caught at apply time.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.79.0/docs#locations AzurermProvider#locations}
	Locations interface{} `field:"optional" json:"locations" yaml:"locations"`
	// Should the AzureRM Provider validate Resource Provider arguments against the list of supported Resource Providers?
	//
	// When enabled, invalid resource providers are caught at plan time; when disabled, they are caught at apply time.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.79.0/docs#resource_providers AzurermProvider#resource_providers}
	ResourceProviders interface{} `field:"optional" json:"resourceProviders" yaml:"resourceProviders"`
}

