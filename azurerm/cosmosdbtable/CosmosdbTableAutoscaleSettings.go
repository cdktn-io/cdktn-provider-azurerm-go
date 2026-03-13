// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package cosmosdbtable


type CosmosdbTableAutoscaleSettings struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.64.0/docs/resources/cosmosdb_table#max_throughput CosmosdbTable#max_throughput}.
	MaxThroughput *float64 `field:"optional" json:"maxThroughput" yaml:"maxThroughput"`
}

