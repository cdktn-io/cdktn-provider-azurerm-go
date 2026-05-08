// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package mongoclusteruser


type MongoClusterUserRole struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.72.0/docs/resources/mongo_cluster_user#database MongoClusterUser#database}.
	Database *string `field:"required" json:"database" yaml:"database"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.72.0/docs/resources/mongo_cluster_user#name MongoClusterUser#name}.
	Name *string `field:"required" json:"name" yaml:"name"`
}

