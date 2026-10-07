// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package oracleautonomousdatabasecrossregiondisasterrecovery


type OracleAutonomousDatabaseCrossRegionDisasterRecoveryTimeouts struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/5.8.0/docs/resources/oracle_autonomous_database_cross_region_disaster_recovery#create OracleAutonomousDatabaseCrossRegionDisasterRecovery#create}.
	Create *string `field:"optional" json:"create" yaml:"create"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/5.8.0/docs/resources/oracle_autonomous_database_cross_region_disaster_recovery#delete OracleAutonomousDatabaseCrossRegionDisasterRecovery#delete}.
	Delete *string `field:"optional" json:"delete" yaml:"delete"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/5.8.0/docs/resources/oracle_autonomous_database_cross_region_disaster_recovery#read OracleAutonomousDatabaseCrossRegionDisasterRecovery#read}.
	Read *string `field:"optional" json:"read" yaml:"read"`
}

