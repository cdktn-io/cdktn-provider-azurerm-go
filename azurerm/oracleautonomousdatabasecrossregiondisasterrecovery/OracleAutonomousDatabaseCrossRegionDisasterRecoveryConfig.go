// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package oracleautonomousdatabasecrossregiondisasterrecovery

import (
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type OracleAutonomousDatabaseCrossRegionDisasterRecoveryConfig struct {
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
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/5.8.0/docs/resources/oracle_autonomous_database_cross_region_disaster_recovery#display_name OracleAutonomousDatabaseCrossRegionDisasterRecovery#display_name}.
	DisplayName *string `field:"required" json:"displayName" yaml:"displayName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/5.8.0/docs/resources/oracle_autonomous_database_cross_region_disaster_recovery#location OracleAutonomousDatabaseCrossRegionDisasterRecovery#location}.
	Location *string `field:"required" json:"location" yaml:"location"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/5.8.0/docs/resources/oracle_autonomous_database_cross_region_disaster_recovery#name OracleAutonomousDatabaseCrossRegionDisasterRecovery#name}.
	Name *string `field:"required" json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/5.8.0/docs/resources/oracle_autonomous_database_cross_region_disaster_recovery#resource_group_name OracleAutonomousDatabaseCrossRegionDisasterRecovery#resource_group_name}.
	ResourceGroupName *string `field:"required" json:"resourceGroupName" yaml:"resourceGroupName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/5.8.0/docs/resources/oracle_autonomous_database_cross_region_disaster_recovery#source_autonomous_database_id OracleAutonomousDatabaseCrossRegionDisasterRecovery#source_autonomous_database_id}.
	SourceAutonomousDatabaseId *string `field:"required" json:"sourceAutonomousDatabaseId" yaml:"sourceAutonomousDatabaseId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/5.8.0/docs/resources/oracle_autonomous_database_cross_region_disaster_recovery#subnet_id OracleAutonomousDatabaseCrossRegionDisasterRecovery#subnet_id}.
	SubnetId *string `field:"required" json:"subnetId" yaml:"subnetId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/5.8.0/docs/resources/oracle_autonomous_database_cross_region_disaster_recovery#id OracleAutonomousDatabaseCrossRegionDisasterRecovery#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/5.8.0/docs/resources/oracle_autonomous_database_cross_region_disaster_recovery#replicate_automatic_backups_enabled OracleAutonomousDatabaseCrossRegionDisasterRecovery#replicate_automatic_backups_enabled}.
	ReplicateAutomaticBackupsEnabled interface{} `field:"optional" json:"replicateAutomaticBackupsEnabled" yaml:"replicateAutomaticBackupsEnabled"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/5.8.0/docs/resources/oracle_autonomous_database_cross_region_disaster_recovery#tags OracleAutonomousDatabaseCrossRegionDisasterRecovery#tags}.
	Tags *map[string]*string `field:"optional" json:"tags" yaml:"tags"`
	// timeouts block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/5.8.0/docs/resources/oracle_autonomous_database_cross_region_disaster_recovery#timeouts OracleAutonomousDatabaseCrossRegionDisasterRecovery#timeouts}
	Timeouts *OracleAutonomousDatabaseCrossRegionDisasterRecoveryTimeouts `field:"optional" json:"timeouts" yaml:"timeouts"`
}

