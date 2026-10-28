// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package sharedimageversion


type SharedImageVersionUefi struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/5.11.0/docs/resources/shared_image_version#signature_template_names SharedImageVersion#signature_template_names}.
	SignatureTemplateNames *[]*string `field:"required" json:"signatureTemplateNames" yaml:"signatureTemplateNames"`
	// additional_signatures block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/5.11.0/docs/resources/shared_image_version#additional_signatures SharedImageVersion#additional_signatures}
	AdditionalSignatures *SharedImageVersionUefiAdditionalSignatures `field:"optional" json:"additionalSignatures" yaml:"additionalSignatures"`
}

