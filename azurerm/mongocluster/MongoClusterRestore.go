// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package mongocluster


type MongoClusterRestore struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.80.0/docs/resources/mongo_cluster#point_in_time_utc MongoCluster#point_in_time_utc}.
	PointInTimeUtc *string `field:"required" json:"pointInTimeUtc" yaml:"pointInTimeUtc"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.80.0/docs/resources/mongo_cluster#source_id MongoCluster#source_id}.
	SourceId *string `field:"required" json:"sourceId" yaml:"sourceId"`
}

