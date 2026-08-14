// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package kustoattacheddatabaseconfiguration

import (
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type KustoAttachedDatabaseConfigurationConfig struct {
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
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/5.1.0/docs/resources/kusto_attached_database_configuration#cluster_id KustoAttachedDatabaseConfiguration#cluster_id}.
	ClusterId *string `field:"required" json:"clusterId" yaml:"clusterId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/5.1.0/docs/resources/kusto_attached_database_configuration#cluster_name KustoAttachedDatabaseConfiguration#cluster_name}.
	ClusterName *string `field:"required" json:"clusterName" yaml:"clusterName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/5.1.0/docs/resources/kusto_attached_database_configuration#database_name KustoAttachedDatabaseConfiguration#database_name}.
	DatabaseName *string `field:"required" json:"databaseName" yaml:"databaseName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/5.1.0/docs/resources/kusto_attached_database_configuration#location KustoAttachedDatabaseConfiguration#location}.
	Location *string `field:"required" json:"location" yaml:"location"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/5.1.0/docs/resources/kusto_attached_database_configuration#name KustoAttachedDatabaseConfiguration#name}.
	Name *string `field:"required" json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/5.1.0/docs/resources/kusto_attached_database_configuration#resource_group_name KustoAttachedDatabaseConfiguration#resource_group_name}.
	ResourceGroupName *string `field:"required" json:"resourceGroupName" yaml:"resourceGroupName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/5.1.0/docs/resources/kusto_attached_database_configuration#database_name_override KustoAttachedDatabaseConfiguration#database_name_override}.
	DatabaseNameOverride *string `field:"optional" json:"databaseNameOverride" yaml:"databaseNameOverride"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/5.1.0/docs/resources/kusto_attached_database_configuration#database_name_prefix KustoAttachedDatabaseConfiguration#database_name_prefix}.
	DatabaseNamePrefix *string `field:"optional" json:"databaseNamePrefix" yaml:"databaseNamePrefix"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/5.1.0/docs/resources/kusto_attached_database_configuration#default_principal_modification_kind KustoAttachedDatabaseConfiguration#default_principal_modification_kind}.
	DefaultPrincipalModificationKind *string `field:"optional" json:"defaultPrincipalModificationKind" yaml:"defaultPrincipalModificationKind"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/5.1.0/docs/resources/kusto_attached_database_configuration#id KustoAttachedDatabaseConfiguration#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// sharing block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/5.1.0/docs/resources/kusto_attached_database_configuration#sharing KustoAttachedDatabaseConfiguration#sharing}
	Sharing *KustoAttachedDatabaseConfigurationSharing `field:"optional" json:"sharing" yaml:"sharing"`
	// timeouts block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/5.1.0/docs/resources/kusto_attached_database_configuration#timeouts KustoAttachedDatabaseConfiguration#timeouts}
	Timeouts *KustoAttachedDatabaseConfigurationTimeouts `field:"optional" json:"timeouts" yaml:"timeouts"`
}

