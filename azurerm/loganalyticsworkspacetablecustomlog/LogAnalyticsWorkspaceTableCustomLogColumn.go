// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package loganalyticsworkspacetablecustomlog


type LogAnalyticsWorkspaceTableCustomLogColumn struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.78.0/docs/resources/log_analytics_workspace_table_custom_log#name LogAnalyticsWorkspaceTableCustomLog#name}.
	Name *string `field:"required" json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.78.0/docs/resources/log_analytics_workspace_table_custom_log#type LogAnalyticsWorkspaceTableCustomLog#type}.
	Type *string `field:"required" json:"type" yaml:"type"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.78.0/docs/resources/log_analytics_workspace_table_custom_log#description LogAnalyticsWorkspaceTableCustomLog#description}.
	Description *string `field:"optional" json:"description" yaml:"description"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.78.0/docs/resources/log_analytics_workspace_table_custom_log#display_name LogAnalyticsWorkspaceTableCustomLog#display_name}.
	DisplayName *string `field:"optional" json:"displayName" yaml:"displayName"`
}

