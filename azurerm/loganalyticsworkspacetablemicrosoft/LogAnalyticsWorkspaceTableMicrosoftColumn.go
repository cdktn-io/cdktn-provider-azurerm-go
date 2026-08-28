// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package loganalyticsworkspacetablemicrosoft


type LogAnalyticsWorkspaceTableMicrosoftColumn struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/5.3.0/docs/resources/log_analytics_workspace_table_microsoft#name LogAnalyticsWorkspaceTableMicrosoft#name}.
	Name *string `field:"required" json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/5.3.0/docs/resources/log_analytics_workspace_table_microsoft#type LogAnalyticsWorkspaceTableMicrosoft#type}.
	Type *string `field:"required" json:"type" yaml:"type"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/5.3.0/docs/resources/log_analytics_workspace_table_microsoft#description LogAnalyticsWorkspaceTableMicrosoft#description}.
	Description *string `field:"optional" json:"description" yaml:"description"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/5.3.0/docs/resources/log_analytics_workspace_table_microsoft#display_by_default LogAnalyticsWorkspaceTableMicrosoft#display_by_default}.
	DisplayByDefault interface{} `field:"optional" json:"displayByDefault" yaml:"displayByDefault"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/5.3.0/docs/resources/log_analytics_workspace_table_microsoft#display_name LogAnalyticsWorkspaceTableMicrosoft#display_name}.
	DisplayName *string `field:"optional" json:"displayName" yaml:"displayName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/5.3.0/docs/resources/log_analytics_workspace_table_microsoft#hidden LogAnalyticsWorkspaceTableMicrosoft#hidden}.
	Hidden interface{} `field:"optional" json:"hidden" yaml:"hidden"`
}

