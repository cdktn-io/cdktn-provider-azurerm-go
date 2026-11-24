// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package analysisservicesserver

import (
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type AnalysisServicesServerConfig struct {
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
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/5.15.0/docs/resources/analysis_services_server#location AnalysisServicesServer#location}.
	Location *string `field:"required" json:"location" yaml:"location"`
	// The name of the Analysis Services Server. Only lowercase Alphanumeric characters allowed, starting with a letter.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/5.15.0/docs/resources/analysis_services_server#name AnalysisServicesServer#name}
	Name *string `field:"required" json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/5.15.0/docs/resources/analysis_services_server#resource_group_name AnalysisServicesServer#resource_group_name}.
	ResourceGroupName *string `field:"required" json:"resourceGroupName" yaml:"resourceGroupName"`
	// SKU for the Analysis Services Server.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/5.15.0/docs/resources/analysis_services_server#sku AnalysisServicesServer#sku}
	Sku *string `field:"required" json:"sku" yaml:"sku"`
	// List of email addresses of admin users.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/5.15.0/docs/resources/analysis_services_server#admin_users AnalysisServicesServer#admin_users}
	AdminUsers *[]*string `field:"optional" json:"adminUsers" yaml:"adminUsers"`
	// URI and SAS token for a blob container to store backups.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/5.15.0/docs/resources/analysis_services_server#backup_blob_container_uri AnalysisServicesServer#backup_blob_container_uri}
	BackupBlobContainerUri *string `field:"optional" json:"backupBlobContainerUri" yaml:"backupBlobContainerUri"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/5.15.0/docs/resources/analysis_services_server#id AnalysisServicesServer#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// ipv4_firewall_rule block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/5.15.0/docs/resources/analysis_services_server#ipv4_firewall_rule AnalysisServicesServer#ipv4_firewall_rule}
	Ipv4FirewallRule interface{} `field:"optional" json:"ipv4FirewallRule" yaml:"ipv4FirewallRule"`
	// Whether the Power BI service is allowed to access or not.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/5.15.0/docs/resources/analysis_services_server#power_bi_service_enabled AnalysisServicesServer#power_bi_service_enabled}
	PowerBiServiceEnabled interface{} `field:"optional" json:"powerBiServiceEnabled" yaml:"powerBiServiceEnabled"`
	// Controls how the read-write server is used in the query pool.
	//
	// If this value is set to `All` then read-write servers are also used for queries. Otherwise with `ReadOnly` these servers do not participate in query operations.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/5.15.0/docs/resources/analysis_services_server#querypool_connection_mode AnalysisServicesServer#querypool_connection_mode}
	QuerypoolConnectionMode *string `field:"optional" json:"querypoolConnectionMode" yaml:"querypoolConnectionMode"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/5.15.0/docs/resources/analysis_services_server#tags AnalysisServicesServer#tags}.
	Tags *map[string]*string `field:"optional" json:"tags" yaml:"tags"`
	// timeouts block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/5.15.0/docs/resources/analysis_services_server#timeouts AnalysisServicesServer#timeouts}
	Timeouts *AnalysisServicesServerTimeouts `field:"optional" json:"timeouts" yaml:"timeouts"`
}

