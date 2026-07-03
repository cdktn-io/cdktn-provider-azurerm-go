// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package datafactorytriggertumblingwindow


type DataFactoryTriggerTumblingWindowPipeline struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.80.0/docs/resources/data_factory_trigger_tumbling_window#name DataFactoryTriggerTumblingWindow#name}.
	Name *string `field:"required" json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.80.0/docs/resources/data_factory_trigger_tumbling_window#parameters DataFactoryTriggerTumblingWindow#parameters}.
	Parameters *map[string]*string `field:"optional" json:"parameters" yaml:"parameters"`
}

