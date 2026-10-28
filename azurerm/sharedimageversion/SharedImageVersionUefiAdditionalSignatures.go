// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package sharedimageversion


type SharedImageVersionUefiAdditionalSignatures struct {
	// db block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/5.11.0/docs/resources/shared_image_version#db SharedImageVersion#db}
	Db interface{} `field:"optional" json:"db" yaml:"db"`
	// dbx block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/5.11.0/docs/resources/shared_image_version#dbx SharedImageVersion#dbx}
	Dbx interface{} `field:"optional" json:"dbx" yaml:"dbx"`
	// kek block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/5.11.0/docs/resources/shared_image_version#kek SharedImageVersion#kek}
	Kek interface{} `field:"optional" json:"kek" yaml:"kek"`
	// pk block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/5.11.0/docs/resources/shared_image_version#pk SharedImageVersion#pk}
	Pk *SharedImageVersionUefiAdditionalSignaturesPk `field:"optional" json:"pk" yaml:"pk"`
}

