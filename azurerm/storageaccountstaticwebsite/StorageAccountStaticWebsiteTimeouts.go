// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package storageaccountstaticwebsite


type StorageAccountStaticWebsiteTimeouts struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/5.0.1/docs/resources/storage_account_static_website#create StorageAccountStaticWebsite#create}.
	Create *string `field:"optional" json:"create" yaml:"create"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/5.0.1/docs/resources/storage_account_static_website#delete StorageAccountStaticWebsite#delete}.
	Delete *string `field:"optional" json:"delete" yaml:"delete"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/5.0.1/docs/resources/storage_account_static_website#read StorageAccountStaticWebsite#read}.
	Read *string `field:"optional" json:"read" yaml:"read"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/5.0.1/docs/resources/storage_account_static_website#update StorageAccountStaticWebsite#update}.
	Update *string `field:"optional" json:"update" yaml:"update"`
}

