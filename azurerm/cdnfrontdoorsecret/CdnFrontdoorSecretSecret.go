// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package cdnfrontdoorsecret


type CdnFrontdoorSecretSecret struct {
	// customer_certificate block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/5.1.0/docs/resources/cdn_frontdoor_secret#customer_certificate CdnFrontdoorSecret#customer_certificate}
	CustomerCertificate interface{} `field:"required" json:"customerCertificate" yaml:"customerCertificate"`
}

