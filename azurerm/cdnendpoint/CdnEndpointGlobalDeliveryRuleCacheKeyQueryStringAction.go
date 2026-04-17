// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package cdnendpoint


type CdnEndpointGlobalDeliveryRuleCacheKeyQueryStringAction struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.69.0/docs/resources/cdn_endpoint#behavior CdnEndpoint#behavior}.
	Behavior *string `field:"required" json:"behavior" yaml:"behavior"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.69.0/docs/resources/cdn_endpoint#parameters CdnEndpoint#parameters}.
	Parameters *string `field:"optional" json:"parameters" yaml:"parameters"`
}

