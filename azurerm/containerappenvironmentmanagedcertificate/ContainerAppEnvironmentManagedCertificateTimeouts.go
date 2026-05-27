// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package containerappenvironmentmanagedcertificate


type ContainerAppEnvironmentManagedCertificateTimeouts struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.74.0/docs/resources/container_app_environment_managed_certificate#create ContainerAppEnvironmentManagedCertificate#create}.
	Create *string `field:"optional" json:"create" yaml:"create"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.74.0/docs/resources/container_app_environment_managed_certificate#delete ContainerAppEnvironmentManagedCertificate#delete}.
	Delete *string `field:"optional" json:"delete" yaml:"delete"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.74.0/docs/resources/container_app_environment_managed_certificate#read ContainerAppEnvironmentManagedCertificate#read}.
	Read *string `field:"optional" json:"read" yaml:"read"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.74.0/docs/resources/container_app_environment_managed_certificate#update ContainerAppEnvironmentManagedCertificate#update}.
	Update *string `field:"optional" json:"update" yaml:"update"`
}

