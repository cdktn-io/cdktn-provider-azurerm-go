// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package manageddevopspool


type ManagedDevopsPoolStatelessAgent struct {
	// automatic_resource_prediction block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.79.0/docs/resources/managed_devops_pool#automatic_resource_prediction ManagedDevopsPool#automatic_resource_prediction}
	AutomaticResourcePrediction *ManagedDevopsPoolStatelessAgentAutomaticResourcePrediction `field:"optional" json:"automaticResourcePrediction" yaml:"automaticResourcePrediction"`
	// manual_resource_prediction block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.79.0/docs/resources/managed_devops_pool#manual_resource_prediction ManagedDevopsPool#manual_resource_prediction}
	ManualResourcePrediction *ManagedDevopsPoolStatelessAgentManualResourcePrediction `field:"optional" json:"manualResourcePrediction" yaml:"manualResourcePrediction"`
}

