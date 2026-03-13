// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package customprovider


type CustomProviderValidation struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.64.0/docs/resources/custom_provider#specification CustomProvider#specification}.
	Specification *string `field:"required" json:"specification" yaml:"specification"`
}

