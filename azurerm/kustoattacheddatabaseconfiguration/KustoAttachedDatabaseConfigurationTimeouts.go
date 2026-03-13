// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package kustoattacheddatabaseconfiguration


type KustoAttachedDatabaseConfigurationTimeouts struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.64.0/docs/resources/kusto_attached_database_configuration#create KustoAttachedDatabaseConfiguration#create}.
	Create *string `field:"optional" json:"create" yaml:"create"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.64.0/docs/resources/kusto_attached_database_configuration#delete KustoAttachedDatabaseConfiguration#delete}.
	Delete *string `field:"optional" json:"delete" yaml:"delete"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.64.0/docs/resources/kusto_attached_database_configuration#read KustoAttachedDatabaseConfiguration#read}.
	Read *string `field:"optional" json:"read" yaml:"read"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.64.0/docs/resources/kusto_attached_database_configuration#update KustoAttachedDatabaseConfiguration#update}.
	Update *string `field:"optional" json:"update" yaml:"update"`
}

