// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package webpubsubhub


type WebPubsubHubTimeouts struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.64.0/docs/resources/web_pubsub_hub#create WebPubsubHub#create}.
	Create *string `field:"optional" json:"create" yaml:"create"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.64.0/docs/resources/web_pubsub_hub#delete WebPubsubHub#delete}.
	Delete *string `field:"optional" json:"delete" yaml:"delete"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.64.0/docs/resources/web_pubsub_hub#read WebPubsubHub#read}.
	Read *string `field:"optional" json:"read" yaml:"read"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.64.0/docs/resources/web_pubsub_hub#update WebPubsubHub#update}.
	Update *string `field:"optional" json:"update" yaml:"update"`
}

