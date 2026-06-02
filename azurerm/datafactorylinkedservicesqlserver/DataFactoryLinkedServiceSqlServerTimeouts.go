// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package datafactorylinkedservicesqlserver


type DataFactoryLinkedServiceSqlServerTimeouts struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.75.0/docs/resources/data_factory_linked_service_sql_server#create DataFactoryLinkedServiceSqlServer#create}.
	Create *string `field:"optional" json:"create" yaml:"create"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.75.0/docs/resources/data_factory_linked_service_sql_server#delete DataFactoryLinkedServiceSqlServer#delete}.
	Delete *string `field:"optional" json:"delete" yaml:"delete"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.75.0/docs/resources/data_factory_linked_service_sql_server#read DataFactoryLinkedServiceSqlServer#read}.
	Read *string `field:"optional" json:"read" yaml:"read"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.75.0/docs/resources/data_factory_linked_service_sql_server#update DataFactoryLinkedServiceSqlServer#update}.
	Update *string `field:"optional" json:"update" yaml:"update"`
}

