// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package cdnfrontdoorsecuritypolicy


type CdnFrontdoorSecurityPolicySecurityPolicies struct {
	// firewall block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/5.0.1/docs/resources/cdn_frontdoor_security_policy#firewall CdnFrontdoorSecurityPolicy#firewall}
	Firewall *CdnFrontdoorSecurityPolicySecurityPoliciesFirewall `field:"required" json:"firewall" yaml:"firewall"`
}

