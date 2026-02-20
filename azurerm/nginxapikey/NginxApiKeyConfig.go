// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package nginxapikey

import (
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type NginxApiKeyConfig struct {
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
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.61.0/docs/resources/nginx_api_key#end_date_time NginxApiKey#end_date_time}.
	EndDateTime *string `field:"required" json:"endDateTime" yaml:"endDateTime"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.61.0/docs/resources/nginx_api_key#name NginxApiKey#name}.
	Name *string `field:"required" json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.61.0/docs/resources/nginx_api_key#nginx_deployment_id NginxApiKey#nginx_deployment_id}.
	NginxDeploymentId *string `field:"required" json:"nginxDeploymentId" yaml:"nginxDeploymentId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.61.0/docs/resources/nginx_api_key#secret_text NginxApiKey#secret_text}.
	SecretText *string `field:"required" json:"secretText" yaml:"secretText"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.61.0/docs/resources/nginx_api_key#id NginxApiKey#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// timeouts block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.61.0/docs/resources/nginx_api_key#timeouts NginxApiKey#timeouts}
	Timeouts *NginxApiKeyTimeouts `field:"optional" json:"timeouts" yaml:"timeouts"`
}

