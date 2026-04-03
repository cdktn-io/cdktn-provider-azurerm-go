// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package datadogmonitorssoconfiguration

import (
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type DatadogMonitorSsoConfigurationConfig struct {
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
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.67.0/docs/resources/datadog_monitor_sso_configuration#datadog_monitor_id DatadogMonitorSsoConfiguration#datadog_monitor_id}.
	DatadogMonitorId *string `field:"required" json:"datadogMonitorId" yaml:"datadogMonitorId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.67.0/docs/resources/datadog_monitor_sso_configuration#enterprise_application_id DatadogMonitorSsoConfiguration#enterprise_application_id}.
	EnterpriseApplicationId *string `field:"required" json:"enterpriseApplicationId" yaml:"enterpriseApplicationId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.67.0/docs/resources/datadog_monitor_sso_configuration#id DatadogMonitorSsoConfiguration#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.67.0/docs/resources/datadog_monitor_sso_configuration#name DatadogMonitorSsoConfiguration#name}.
	Name *string `field:"optional" json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.67.0/docs/resources/datadog_monitor_sso_configuration#single_sign_on DatadogMonitorSsoConfiguration#single_sign_on}.
	SingleSignOn *string `field:"optional" json:"singleSignOn" yaml:"singleSignOn"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.67.0/docs/resources/datadog_monitor_sso_configuration#single_sign_on_enabled DatadogMonitorSsoConfiguration#single_sign_on_enabled}.
	SingleSignOnEnabled *string `field:"optional" json:"singleSignOnEnabled" yaml:"singleSignOnEnabled"`
	// timeouts block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.67.0/docs/resources/datadog_monitor_sso_configuration#timeouts DatadogMonitorSsoConfiguration#timeouts}
	Timeouts *DatadogMonitorSsoConfigurationTimeouts `field:"optional" json:"timeouts" yaml:"timeouts"`
}

