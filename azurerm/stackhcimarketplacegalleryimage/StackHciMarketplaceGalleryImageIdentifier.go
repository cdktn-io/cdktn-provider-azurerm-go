// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package stackhcimarketplacegalleryimage


type StackHciMarketplaceGalleryImageIdentifier struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.78.0/docs/resources/stack_hci_marketplace_gallery_image#offer StackHciMarketplaceGalleryImage#offer}.
	Offer *string `field:"required" json:"offer" yaml:"offer"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.78.0/docs/resources/stack_hci_marketplace_gallery_image#publisher StackHciMarketplaceGalleryImage#publisher}.
	Publisher *string `field:"required" json:"publisher" yaml:"publisher"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.78.0/docs/resources/stack_hci_marketplace_gallery_image#sku StackHciMarketplaceGalleryImage#sku}.
	Sku *string `field:"required" json:"sku" yaml:"sku"`
}

