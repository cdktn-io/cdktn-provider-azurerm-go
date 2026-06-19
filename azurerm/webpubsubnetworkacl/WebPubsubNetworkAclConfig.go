// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package webpubsubnetworkacl

import (
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type WebPubsubNetworkAclConfig struct {
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
	// public_network block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.78.0/docs/resources/web_pubsub_network_acl#public_network WebPubsubNetworkAcl#public_network}
	PublicNetwork *WebPubsubNetworkAclPublicNetwork `field:"required" json:"publicNetwork" yaml:"publicNetwork"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.78.0/docs/resources/web_pubsub_network_acl#web_pubsub_id WebPubsubNetworkAcl#web_pubsub_id}.
	WebPubsubId *string `field:"required" json:"webPubsubId" yaml:"webPubsubId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.78.0/docs/resources/web_pubsub_network_acl#default_action WebPubsubNetworkAcl#default_action}.
	DefaultAction *string `field:"optional" json:"defaultAction" yaml:"defaultAction"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.78.0/docs/resources/web_pubsub_network_acl#id WebPubsubNetworkAcl#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// private_endpoint block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.78.0/docs/resources/web_pubsub_network_acl#private_endpoint WebPubsubNetworkAcl#private_endpoint}
	PrivateEndpoint interface{} `field:"optional" json:"privateEndpoint" yaml:"privateEndpoint"`
	// timeouts block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.78.0/docs/resources/web_pubsub_network_acl#timeouts WebPubsubNetworkAcl#timeouts}
	Timeouts *WebPubsubNetworkAclTimeouts `field:"optional" json:"timeouts" yaml:"timeouts"`
}

