// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package sharedimageversion


type SharedImageVersionUefiAdditionalSignaturesDb struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/5.10.0/docs/resources/shared_image_version#type SharedImageVersion#type}.
	Type *string `field:"required" json:"type" yaml:"type"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/5.10.0/docs/resources/shared_image_version#values_base64 SharedImageVersion#values_base64}.
	ValuesBase64 *[]*string `field:"required" json:"valuesBase64" yaml:"valuesBase64"`
}

