// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package datafactorylinkedservicesqlmanagedinstance

import (
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type DataFactoryLinkedServiceSqlManagedInstanceConfig struct {
	// Experimental.
	Connection interface{} `field:"optional" json:"connection" yaml:"connection"`
	// Experimental.
	Count interface{} `field:"optional" json:"count" yaml:"count"`
	// Experimental.
	DependsOn *[]cdktn.ITerraformDependable `field:"optional" json:"dependsOn" yaml:"dependsOn"`
	// Experimental.
	ForEach cdktn.ITerraformIterator `field:"optional" json:"forEach" yaml:"forEach"`
	// Experimental.
	Lifecycle *cdktn.TerraformResourceLifecycle `field:"optional" json:"lifecycle" yaml:"lifecycle"`
	// Experimental.
	Provider cdktn.TerraformProvider `field:"optional" json:"provider" yaml:"provider"`
	// Experimental.
	Provisioners *[]interface{} `field:"optional" json:"provisioners" yaml:"provisioners"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.80.0/docs/resources/data_factory_linked_service_sql_managed_instance#data_factory_id DataFactoryLinkedServiceSqlManagedInstance#data_factory_id}.
	DataFactoryId *string `field:"required" json:"dataFactoryId" yaml:"dataFactoryId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.80.0/docs/resources/data_factory_linked_service_sql_managed_instance#name DataFactoryLinkedServiceSqlManagedInstance#name}.
	Name *string `field:"required" json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.80.0/docs/resources/data_factory_linked_service_sql_managed_instance#annotations DataFactoryLinkedServiceSqlManagedInstance#annotations}.
	Annotations *[]*string `field:"optional" json:"annotations" yaml:"annotations"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.80.0/docs/resources/data_factory_linked_service_sql_managed_instance#connection_string DataFactoryLinkedServiceSqlManagedInstance#connection_string}.
	ConnectionString *string `field:"optional" json:"connectionString" yaml:"connectionString"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.80.0/docs/resources/data_factory_linked_service_sql_managed_instance#description DataFactoryLinkedServiceSqlManagedInstance#description}.
	Description *string `field:"optional" json:"description" yaml:"description"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.80.0/docs/resources/data_factory_linked_service_sql_managed_instance#id DataFactoryLinkedServiceSqlManagedInstance#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.80.0/docs/resources/data_factory_linked_service_sql_managed_instance#integration_runtime_name DataFactoryLinkedServiceSqlManagedInstance#integration_runtime_name}.
	IntegrationRuntimeName *string `field:"optional" json:"integrationRuntimeName" yaml:"integrationRuntimeName"`
	// key_vault_connection_string block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.80.0/docs/resources/data_factory_linked_service_sql_managed_instance#key_vault_connection_string DataFactoryLinkedServiceSqlManagedInstance#key_vault_connection_string}
	KeyVaultConnectionString *DataFactoryLinkedServiceSqlManagedInstanceKeyVaultConnectionString `field:"optional" json:"keyVaultConnectionString" yaml:"keyVaultConnectionString"`
	// key_vault_password block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.80.0/docs/resources/data_factory_linked_service_sql_managed_instance#key_vault_password DataFactoryLinkedServiceSqlManagedInstance#key_vault_password}
	KeyVaultPassword *DataFactoryLinkedServiceSqlManagedInstanceKeyVaultPassword `field:"optional" json:"keyVaultPassword" yaml:"keyVaultPassword"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.80.0/docs/resources/data_factory_linked_service_sql_managed_instance#parameters DataFactoryLinkedServiceSqlManagedInstance#parameters}.
	Parameters *map[string]*string `field:"optional" json:"parameters" yaml:"parameters"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.80.0/docs/resources/data_factory_linked_service_sql_managed_instance#service_principal_id DataFactoryLinkedServiceSqlManagedInstance#service_principal_id}.
	ServicePrincipalId *string `field:"optional" json:"servicePrincipalId" yaml:"servicePrincipalId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.80.0/docs/resources/data_factory_linked_service_sql_managed_instance#service_principal_key DataFactoryLinkedServiceSqlManagedInstance#service_principal_key}.
	ServicePrincipalKey *string `field:"optional" json:"servicePrincipalKey" yaml:"servicePrincipalKey"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.80.0/docs/resources/data_factory_linked_service_sql_managed_instance#tenant DataFactoryLinkedServiceSqlManagedInstance#tenant}.
	Tenant *string `field:"optional" json:"tenant" yaml:"tenant"`
	// timeouts block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.80.0/docs/resources/data_factory_linked_service_sql_managed_instance#timeouts DataFactoryLinkedServiceSqlManagedInstance#timeouts}
	Timeouts *DataFactoryLinkedServiceSqlManagedInstanceTimeouts `field:"optional" json:"timeouts" yaml:"timeouts"`
}

