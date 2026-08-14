// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package manageddevopspool


type ManagedDevopsPoolStatefulAgent struct {
	// automatic_resource_prediction block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/5.1.0/docs/resources/managed_devops_pool#automatic_resource_prediction ManagedDevopsPool#automatic_resource_prediction}
	AutomaticResourcePrediction *ManagedDevopsPoolStatefulAgentAutomaticResourcePrediction `field:"optional" json:"automaticResourcePrediction" yaml:"automaticResourcePrediction"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/5.1.0/docs/resources/managed_devops_pool#grace_period_time_span ManagedDevopsPool#grace_period_time_span}.
	GracePeriodTimeSpan *string `field:"optional" json:"gracePeriodTimeSpan" yaml:"gracePeriodTimeSpan"`
	// manual_resource_prediction block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/5.1.0/docs/resources/managed_devops_pool#manual_resource_prediction ManagedDevopsPool#manual_resource_prediction}
	ManualResourcePrediction *ManagedDevopsPoolStatefulAgentManualResourcePrediction `field:"optional" json:"manualResourcePrediction" yaml:"manualResourcePrediction"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/5.1.0/docs/resources/managed_devops_pool#maximum_agent_lifetime ManagedDevopsPool#maximum_agent_lifetime}.
	MaximumAgentLifetime *string `field:"optional" json:"maximumAgentLifetime" yaml:"maximumAgentLifetime"`
}

