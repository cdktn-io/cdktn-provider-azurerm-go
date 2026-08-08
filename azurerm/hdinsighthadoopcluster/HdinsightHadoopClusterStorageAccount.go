// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package hdinsighthadoopcluster


type HdinsightHadoopClusterStorageAccount struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/5.0.1/docs/resources/hdinsight_hadoop_cluster#is_default HdinsightHadoopCluster#is_default}.
	IsDefault interface{} `field:"required" json:"isDefault" yaml:"isDefault"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/5.0.1/docs/resources/hdinsight_hadoop_cluster#storage_account_key HdinsightHadoopCluster#storage_account_key}.
	StorageAccountKey *string `field:"required" json:"storageAccountKey" yaml:"storageAccountKey"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/5.0.1/docs/resources/hdinsight_hadoop_cluster#storage_container_url HdinsightHadoopCluster#storage_container_url}.
	StorageContainerUrl *string `field:"required" json:"storageContainerUrl" yaml:"storageContainerUrl"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/5.0.1/docs/resources/hdinsight_hadoop_cluster#storage_account_id HdinsightHadoopCluster#storage_account_id}.
	StorageAccountId *string `field:"optional" json:"storageAccountId" yaml:"storageAccountId"`
}

