// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package cognitiveaccountconnectioncustomkeys

import (
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type CognitiveAccountConnectionCustomKeysConfig struct {
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
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.81.0/docs/resources/cognitive_account_connection_custom_keys#category CognitiveAccountConnectionCustomKeys#category}.
	Category *string `field:"required" json:"category" yaml:"category"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.81.0/docs/resources/cognitive_account_connection_custom_keys#cognitive_account_id CognitiveAccountConnectionCustomKeys#cognitive_account_id}.
	CognitiveAccountId *string `field:"required" json:"cognitiveAccountId" yaml:"cognitiveAccountId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.81.0/docs/resources/cognitive_account_connection_custom_keys#custom_keys CognitiveAccountConnectionCustomKeys#custom_keys}.
	CustomKeys *map[string]*string `field:"required" json:"customKeys" yaml:"customKeys"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.81.0/docs/resources/cognitive_account_connection_custom_keys#name CognitiveAccountConnectionCustomKeys#name}.
	Name *string `field:"required" json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.81.0/docs/resources/cognitive_account_connection_custom_keys#target CognitiveAccountConnectionCustomKeys#target}.
	Target *string `field:"required" json:"target" yaml:"target"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.81.0/docs/resources/cognitive_account_connection_custom_keys#id CognitiveAccountConnectionCustomKeys#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.81.0/docs/resources/cognitive_account_connection_custom_keys#metadata CognitiveAccountConnectionCustomKeys#metadata}.
	Metadata *map[string]*string `field:"optional" json:"metadata" yaml:"metadata"`
	// timeouts block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.81.0/docs/resources/cognitive_account_connection_custom_keys#timeouts CognitiveAccountConnectionCustomKeys#timeouts}
	Timeouts *CognitiveAccountConnectionCustomKeysTimeouts `field:"optional" json:"timeouts" yaml:"timeouts"`
}

