// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package storageaccounttableproperties


type StorageAccountTablePropertiesHourMetrics struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/5.0.1/docs/resources/storage_account_table_properties#version StorageAccountTableProperties#version}.
	Version *string `field:"required" json:"version" yaml:"version"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/5.0.1/docs/resources/storage_account_table_properties#include_apis StorageAccountTableProperties#include_apis}.
	IncludeApis interface{} `field:"optional" json:"includeApis" yaml:"includeApis"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/5.0.1/docs/resources/storage_account_table_properties#retention_policy_days StorageAccountTableProperties#retention_policy_days}.
	RetentionPolicyDays *float64 `field:"optional" json:"retentionPolicyDays" yaml:"retentionPolicyDays"`
}

