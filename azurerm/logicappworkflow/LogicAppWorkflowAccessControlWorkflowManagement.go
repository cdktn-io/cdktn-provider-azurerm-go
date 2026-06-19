// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package logicappworkflow


type LogicAppWorkflowAccessControlWorkflowManagement struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.78.0/docs/resources/logic_app_workflow#allowed_caller_ip_address_range LogicAppWorkflow#allowed_caller_ip_address_range}.
	AllowedCallerIpAddressRange *[]*string `field:"required" json:"allowedCallerIpAddressRange" yaml:"allowedCallerIpAddressRange"`
}

