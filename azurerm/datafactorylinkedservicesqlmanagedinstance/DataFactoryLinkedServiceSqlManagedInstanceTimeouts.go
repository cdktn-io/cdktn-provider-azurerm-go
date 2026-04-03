// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package datafactorylinkedservicesqlmanagedinstance


type DataFactoryLinkedServiceSqlManagedInstanceTimeouts struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.67.0/docs/resources/data_factory_linked_service_sql_managed_instance#create DataFactoryLinkedServiceSqlManagedInstance#create}.
	Create *string `field:"optional" json:"create" yaml:"create"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.67.0/docs/resources/data_factory_linked_service_sql_managed_instance#delete DataFactoryLinkedServiceSqlManagedInstance#delete}.
	Delete *string `field:"optional" json:"delete" yaml:"delete"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.67.0/docs/resources/data_factory_linked_service_sql_managed_instance#read DataFactoryLinkedServiceSqlManagedInstance#read}.
	Read *string `field:"optional" json:"read" yaml:"read"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.67.0/docs/resources/data_factory_linked_service_sql_managed_instance#update DataFactoryLinkedServiceSqlManagedInstance#update}.
	Update *string `field:"optional" json:"update" yaml:"update"`
}

