// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package loganalyticsworkspacetablecustomlog


type LogAnalyticsWorkspaceTableCustomLogTimeouts struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.75.0/docs/resources/log_analytics_workspace_table_custom_log#create LogAnalyticsWorkspaceTableCustomLog#create}.
	Create *string `field:"optional" json:"create" yaml:"create"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.75.0/docs/resources/log_analytics_workspace_table_custom_log#delete LogAnalyticsWorkspaceTableCustomLog#delete}.
	Delete *string `field:"optional" json:"delete" yaml:"delete"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.75.0/docs/resources/log_analytics_workspace_table_custom_log#read LogAnalyticsWorkspaceTableCustomLog#read}.
	Read *string `field:"optional" json:"read" yaml:"read"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.75.0/docs/resources/log_analytics_workspace_table_custom_log#update LogAnalyticsWorkspaceTableCustomLog#update}.
	Update *string `field:"optional" json:"update" yaml:"update"`
}

