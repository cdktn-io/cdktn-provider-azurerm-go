// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package cosmosdbaccount


type CosmosdbAccountGeoLocation struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.80.0/docs/resources/cosmosdb_account#failover_priority CosmosdbAccount#failover_priority}.
	FailoverPriority *float64 `field:"required" json:"failoverPriority" yaml:"failoverPriority"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.80.0/docs/resources/cosmosdb_account#location CosmosdbAccount#location}.
	Location *string `field:"required" json:"location" yaml:"location"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.80.0/docs/resources/cosmosdb_account#zone_redundant CosmosdbAccount#zone_redundant}.
	ZoneRedundant interface{} `field:"optional" json:"zoneRedundant" yaml:"zoneRedundant"`
}

