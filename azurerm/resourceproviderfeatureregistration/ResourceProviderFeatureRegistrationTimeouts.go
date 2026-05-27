// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package resourceproviderfeatureregistration


type ResourceProviderFeatureRegistrationTimeouts struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.74.0/docs/resources/resource_provider_feature_registration#create ResourceProviderFeatureRegistration#create}.
	Create *string `field:"optional" json:"create" yaml:"create"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.74.0/docs/resources/resource_provider_feature_registration#delete ResourceProviderFeatureRegistration#delete}.
	Delete *string `field:"optional" json:"delete" yaml:"delete"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.74.0/docs/resources/resource_provider_feature_registration#read ResourceProviderFeatureRegistration#read}.
	Read *string `field:"optional" json:"read" yaml:"read"`
}

