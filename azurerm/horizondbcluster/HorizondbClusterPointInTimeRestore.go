// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package horizondbcluster


type HorizondbClusterPointInTimeRestore struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/5.15.0/docs/resources/horizondb_cluster#point_in_time_utc HorizondbCluster#point_in_time_utc}.
	PointInTimeUtc *string `field:"required" json:"pointInTimeUtc" yaml:"pointInTimeUtc"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/5.15.0/docs/resources/horizondb_cluster#source_horizondb_cluster_id HorizondbCluster#source_horizondb_cluster_id}.
	SourceHorizondbClusterId *string `field:"required" json:"sourceHorizondbClusterId" yaml:"sourceHorizondbClusterId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/5.15.0/docs/resources/horizondb_cluster#source_location HorizondbCluster#source_location}.
	SourceLocation *string `field:"optional" json:"sourceLocation" yaml:"sourceLocation"`
}

