// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package storageaccounttableproperties


type StorageAccountTablePropertiesCorsRule struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/5.0.1/docs/resources/storage_account_table_properties#allowed_headers StorageAccountTableProperties#allowed_headers}.
	AllowedHeaders *[]*string `field:"required" json:"allowedHeaders" yaml:"allowedHeaders"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/5.0.1/docs/resources/storage_account_table_properties#allowed_methods StorageAccountTableProperties#allowed_methods}.
	AllowedMethods *[]*string `field:"required" json:"allowedMethods" yaml:"allowedMethods"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/5.0.1/docs/resources/storage_account_table_properties#allowed_origins StorageAccountTableProperties#allowed_origins}.
	AllowedOrigins *[]*string `field:"required" json:"allowedOrigins" yaml:"allowedOrigins"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/5.0.1/docs/resources/storage_account_table_properties#exposed_headers StorageAccountTableProperties#exposed_headers}.
	ExposedHeaders *[]*string `field:"required" json:"exposedHeaders" yaml:"exposedHeaders"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/5.0.1/docs/resources/storage_account_table_properties#max_age_in_seconds StorageAccountTableProperties#max_age_in_seconds}.
	MaxAgeInSeconds *float64 `field:"required" json:"maxAgeInSeconds" yaml:"maxAgeInSeconds"`
}

