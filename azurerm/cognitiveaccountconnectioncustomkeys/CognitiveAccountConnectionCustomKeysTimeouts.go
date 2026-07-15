// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package cognitiveaccountconnectioncustomkeys


type CognitiveAccountConnectionCustomKeysTimeouts struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.81.0/docs/resources/cognitive_account_connection_custom_keys#create CognitiveAccountConnectionCustomKeys#create}.
	Create *string `field:"optional" json:"create" yaml:"create"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.81.0/docs/resources/cognitive_account_connection_custom_keys#delete CognitiveAccountConnectionCustomKeys#delete}.
	Delete *string `field:"optional" json:"delete" yaml:"delete"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.81.0/docs/resources/cognitive_account_connection_custom_keys#read CognitiveAccountConnectionCustomKeys#read}.
	Read *string `field:"optional" json:"read" yaml:"read"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.81.0/docs/resources/cognitive_account_connection_custom_keys#update CognitiveAccountConnectionCustomKeys#update}.
	Update *string `field:"optional" json:"update" yaml:"update"`
}

