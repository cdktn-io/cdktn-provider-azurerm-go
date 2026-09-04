// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package cosmosdbcassandratable


type CosmosdbCassandraTableAutoscaleSettings struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/5.4.0/docs/resources/cosmosdb_cassandra_table#max_throughput CosmosdbCassandraTable#max_throughput}.
	MaxThroughput *float64 `field:"optional" json:"maxThroughput" yaml:"maxThroughput"`
}

