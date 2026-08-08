// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package appservicesourcecontrol


type AppServiceSourceControlTimeouts struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/5.0.1/docs/resources/app_service_source_control#create AppServiceSourceControl#create}.
	Create *string `field:"optional" json:"create" yaml:"create"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/5.0.1/docs/resources/app_service_source_control#delete AppServiceSourceControl#delete}.
	Delete *string `field:"optional" json:"delete" yaml:"delete"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/5.0.1/docs/resources/app_service_source_control#read AppServiceSourceControl#read}.
	Read *string `field:"optional" json:"read" yaml:"read"`
}

