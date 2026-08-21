// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package cdnfrontdoorcustomdomain


type CdnFrontdoorCustomDomainTls struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/5.2.0/docs/resources/cdn_frontdoor_custom_domain#cdn_frontdoor_secret_id CdnFrontdoorCustomDomain#cdn_frontdoor_secret_id}.
	CdnFrontdoorSecretId *string `field:"optional" json:"cdnFrontdoorSecretId" yaml:"cdnFrontdoorSecretId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/5.2.0/docs/resources/cdn_frontdoor_custom_domain#certificate_type CdnFrontdoorCustomDomain#certificate_type}.
	CertificateType *string `field:"optional" json:"certificateType" yaml:"certificateType"`
	// cipher_suite block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/5.2.0/docs/resources/cdn_frontdoor_custom_domain#cipher_suite CdnFrontdoorCustomDomain#cipher_suite}
	CipherSuite *CdnFrontdoorCustomDomainTlsCipherSuite `field:"optional" json:"cipherSuite" yaml:"cipherSuite"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/5.2.0/docs/resources/cdn_frontdoor_custom_domain#minimum_version CdnFrontdoorCustomDomain#minimum_version}.
	MinimumVersion *string `field:"optional" json:"minimumVersion" yaml:"minimumVersion"`
}

