// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package cdnfrontdoorcustomdomain


type CdnFrontdoorCustomDomainTlsCipherSuite struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/5.2.0/docs/resources/cdn_frontdoor_custom_domain#type CdnFrontdoorCustomDomain#type}.
	Type *string `field:"required" json:"type" yaml:"type"`
	// custom_ciphers block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/5.2.0/docs/resources/cdn_frontdoor_custom_domain#custom_ciphers CdnFrontdoorCustomDomain#custom_ciphers}
	CustomCiphers *CdnFrontdoorCustomDomainTlsCipherSuiteCustomCiphers `field:"optional" json:"customCiphers" yaml:"customCiphers"`
}

