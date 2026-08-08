// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package loganalyticsworkspacetablemicrosoft

import (
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type LogAnalyticsWorkspaceTableMicrosoftConfig struct {
	// Experimental.
	Connection interface{} `field:"optional" json:"connection" yaml:"connection"`
	// Experimental.
	Count interface{} `field:"optional" json:"count" yaml:"count"`
	// Experimental.
	DependsOn *[]cdktn.ITerraformDependable `field:"optional" json:"dependsOn" yaml:"dependsOn"`
	// Experimental.
	ForEach cdktn.ITerraformIterator `field:"optional" json:"forEach" yaml:"forEach"`
	// Experimental.
	Lifecycle *cdktn.TerraformResourceLifecycle `field:"optional" json:"lifecycle" yaml:"lifecycle"`
	// Experimental.
	Provider cdktn.TerraformProvider `field:"optional" json:"provider" yaml:"provider"`
	// Experimental.
	Provisioners *[]interface{} `field:"optional" json:"provisioners" yaml:"provisioners"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/5.0.1/docs/resources/log_analytics_workspace_table_microsoft#name LogAnalyticsWorkspaceTableMicrosoft#name}.
	Name *string `field:"required" json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/5.0.1/docs/resources/log_analytics_workspace_table_microsoft#workspace_id LogAnalyticsWorkspaceTableMicrosoft#workspace_id}.
	WorkspaceId *string `field:"required" json:"workspaceId" yaml:"workspaceId"`
	// column block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/5.0.1/docs/resources/log_analytics_workspace_table_microsoft#column LogAnalyticsWorkspaceTableMicrosoft#column}
	Column interface{} `field:"optional" json:"column" yaml:"column"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/5.0.1/docs/resources/log_analytics_workspace_table_microsoft#description LogAnalyticsWorkspaceTableMicrosoft#description}.
	Description *string `field:"optional" json:"description" yaml:"description"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/5.0.1/docs/resources/log_analytics_workspace_table_microsoft#display_name LogAnalyticsWorkspaceTableMicrosoft#display_name}.
	DisplayName *string `field:"optional" json:"displayName" yaml:"displayName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/5.0.1/docs/resources/log_analytics_workspace_table_microsoft#id LogAnalyticsWorkspaceTableMicrosoft#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/5.0.1/docs/resources/log_analytics_workspace_table_microsoft#labels LogAnalyticsWorkspaceTableMicrosoft#labels}.
	Labels *[]*string `field:"optional" json:"labels" yaml:"labels"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/5.0.1/docs/resources/log_analytics_workspace_table_microsoft#retention_in_days LogAnalyticsWorkspaceTableMicrosoft#retention_in_days}.
	RetentionInDays *float64 `field:"optional" json:"retentionInDays" yaml:"retentionInDays"`
	// timeouts block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/5.0.1/docs/resources/log_analytics_workspace_table_microsoft#timeouts LogAnalyticsWorkspaceTableMicrosoft#timeouts}
	Timeouts *LogAnalyticsWorkspaceTableMicrosoftTimeouts `field:"optional" json:"timeouts" yaml:"timeouts"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/5.0.1/docs/resources/log_analytics_workspace_table_microsoft#total_retention_in_days LogAnalyticsWorkspaceTableMicrosoft#total_retention_in_days}.
	TotalRetentionInDays *float64 `field:"optional" json:"totalRetentionInDays" yaml:"totalRetentionInDays"`
}

