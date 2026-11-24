// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package horizondbcluster

import (
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type HorizondbClusterConfig struct {
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
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/5.15.0/docs/resources/horizondb_cluster#administrator_login HorizondbCluster#administrator_login}.
	AdministratorLogin *string `field:"required" json:"administratorLogin" yaml:"administratorLogin"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/5.15.0/docs/resources/horizondb_cluster#location HorizondbCluster#location}.
	Location *string `field:"required" json:"location" yaml:"location"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/5.15.0/docs/resources/horizondb_cluster#name HorizondbCluster#name}.
	Name *string `field:"required" json:"name" yaml:"name"`
	// read_write_pool block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/5.15.0/docs/resources/horizondb_cluster#read_write_pool HorizondbCluster#read_write_pool}
	ReadWritePool *HorizondbClusterReadWritePool `field:"required" json:"readWritePool" yaml:"readWritePool"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/5.15.0/docs/resources/horizondb_cluster#resource_group_name HorizondbCluster#resource_group_name}.
	ResourceGroupName *string `field:"required" json:"resourceGroupName" yaml:"resourceGroupName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/5.15.0/docs/resources/horizondb_cluster#administrator_password HorizondbCluster#administrator_password}.
	AdministratorPassword *string `field:"optional" json:"administratorPassword" yaml:"administratorPassword"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/5.15.0/docs/resources/horizondb_cluster#administrator_password_wo HorizondbCluster#administrator_password_wo}.
	AdministratorPasswordWo *string `field:"optional" json:"administratorPasswordWo" yaml:"administratorPasswordWo"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/5.15.0/docs/resources/horizondb_cluster#administrator_password_wo_version HorizondbCluster#administrator_password_wo_version}.
	AdministratorPasswordWoVersion *float64 `field:"optional" json:"administratorPasswordWoVersion" yaml:"administratorPasswordWoVersion"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/5.15.0/docs/resources/horizondb_cluster#backup_retention_in_days HorizondbCluster#backup_retention_in_days}.
	BackupRetentionInDays *float64 `field:"optional" json:"backupRetentionInDays" yaml:"backupRetentionInDays"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/5.15.0/docs/resources/horizondb_cluster#geo_redundant_backup_enabled HorizondbCluster#geo_redundant_backup_enabled}.
	GeoRedundantBackupEnabled interface{} `field:"optional" json:"geoRedundantBackupEnabled" yaml:"geoRedundantBackupEnabled"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/5.15.0/docs/resources/horizondb_cluster#id HorizondbCluster#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// identity block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/5.15.0/docs/resources/horizondb_cluster#identity HorizondbCluster#identity}
	Identity *HorizondbClusterIdentity `field:"optional" json:"identity" yaml:"identity"`
	// point_in_time_restore block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/5.15.0/docs/resources/horizondb_cluster#point_in_time_restore HorizondbCluster#point_in_time_restore}
	PointInTimeRestore *HorizondbClusterPointInTimeRestore `field:"optional" json:"pointInTimeRestore" yaml:"pointInTimeRestore"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/5.15.0/docs/resources/horizondb_cluster#postgresql_version HorizondbCluster#postgresql_version}.
	PostgresqlVersion *string `field:"optional" json:"postgresqlVersion" yaml:"postgresqlVersion"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/5.15.0/docs/resources/horizondb_cluster#public_network_access HorizondbCluster#public_network_access}.
	PublicNetworkAccess *string `field:"optional" json:"publicNetworkAccess" yaml:"publicNetworkAccess"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/5.15.0/docs/resources/horizondb_cluster#tags HorizondbCluster#tags}.
	Tags *map[string]*string `field:"optional" json:"tags" yaml:"tags"`
	// timeouts block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/5.15.0/docs/resources/horizondb_cluster#timeouts HorizondbCluster#timeouts}
	Timeouts *HorizondbClusterTimeouts `field:"optional" json:"timeouts" yaml:"timeouts"`
}

