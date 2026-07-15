// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package cognitiveaccountconnectionaccountkey

import (
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type CognitiveAccountConnectionAccountKeyConfig struct {
	// Experimental.
	Connection interface{} `field:"optional" json:"connection" yaml:"connection"`
	// Experimental.
	Count interface{} `field:"optional" json:"count" yaml:"count"`
	// Experimental.
	DependsOn *[]cdktn.ITerraformDependable `field:"optional" json:"dependsOn" yaml:"dependsOn"`
	// Experimental.
	ForEach cdktn.ITerraformIterator `field:"optional" json:"forEach" yaml:"forEach"`
	// Experimental.
	Lifecycle *cdktn.TerraformResourceLifecycle `field:"optional" json:"lifecycle" yaml:"lifecycle"`
	// Experimental.
	Provider cdktn.TerraformProvider `field:"optional" json:"provider" yaml:"provider"`
	// Experimental.
	Provisioners *[]interface{} `field:"optional" json:"provisioners" yaml:"provisioners"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.81.0/docs/resources/cognitive_account_connection_account_key#account_key CognitiveAccountConnectionAccountKey#account_key}.
	AccountKey *string `field:"required" json:"accountKey" yaml:"accountKey"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.81.0/docs/resources/cognitive_account_connection_account_key#category CognitiveAccountConnectionAccountKey#category}.
	Category *string `field:"required" json:"category" yaml:"category"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.81.0/docs/resources/cognitive_account_connection_account_key#cognitive_account_id CognitiveAccountConnectionAccountKey#cognitive_account_id}.
	CognitiveAccountId *string `field:"required" json:"cognitiveAccountId" yaml:"cognitiveAccountId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.81.0/docs/resources/cognitive_account_connection_account_key#metadata CognitiveAccountConnectionAccountKey#metadata}.
	Metadata *map[string]*string `field:"required" json:"metadata" yaml:"metadata"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.81.0/docs/resources/cognitive_account_connection_account_key#name CognitiveAccountConnectionAccountKey#name}.
	Name *string `field:"required" json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.81.0/docs/resources/cognitive_account_connection_account_key#target CognitiveAccountConnectionAccountKey#target}.
	Target *string `field:"required" json:"target" yaml:"target"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.81.0/docs/resources/cognitive_account_connection_account_key#id CognitiveAccountConnectionAccountKey#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// timeouts block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.81.0/docs/resources/cognitive_account_connection_account_key#timeouts CognitiveAccountConnectionAccountKey#timeouts}
	Timeouts *CognitiveAccountConnectionAccountKeyTimeouts `field:"optional" json:"timeouts" yaml:"timeouts"`
}

