// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package storageaccounttableproperties


type StorageAccountTablePropertiesLogging struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/5.0.1/docs/resources/storage_account_table_properties#delete StorageAccountTableProperties#delete}.
	Delete interface{} `field:"required" json:"delete" yaml:"delete"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/5.0.1/docs/resources/storage_account_table_properties#read StorageAccountTableProperties#read}.
	Read interface{} `field:"required" json:"read" yaml:"read"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/5.0.1/docs/resources/storage_account_table_properties#version StorageAccountTableProperties#version}.
	Version *string `field:"required" json:"version" yaml:"version"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/5.0.1/docs/resources/storage_account_table_properties#write StorageAccountTableProperties#write}.
	Write interface{} `field:"required" json:"write" yaml:"write"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/5.0.1/docs/resources/storage_account_table_properties#retention_policy_days StorageAccountTableProperties#retention_policy_days}.
	RetentionPolicyDays *float64 `field:"optional" json:"retentionPolicyDays" yaml:"retentionPolicyDays"`
}

