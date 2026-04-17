// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package securitycenterautomation


type SecurityCenterAutomationSourceRuleSet struct {
	// rule block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.69.0/docs/resources/security_center_automation#rule SecurityCenterAutomation#rule}
	Rule interface{} `field:"required" json:"rule" yaml:"rule"`
}

