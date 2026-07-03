// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package containerappenvironmentmanagedcertificate

import (
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type ContainerAppEnvironmentManagedCertificateConfig struct {
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
	// The Container App Managed Environment ID to configure this Managed Certificate on.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.80.0/docs/resources/container_app_environment_managed_certificate#container_app_environment_id ContainerAppEnvironmentManagedCertificate#container_app_environment_id}
	ContainerAppEnvironmentId *string `field:"required" json:"containerAppEnvironmentId" yaml:"containerAppEnvironmentId"`
	// The name of the Container Apps Managed Certificate.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.80.0/docs/resources/container_app_environment_managed_certificate#name ContainerAppEnvironmentManagedCertificate#name}
	Name *string `field:"required" json:"name" yaml:"name"`
	// The Subject Name of the Certificate. Must be a valid domain name.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.80.0/docs/resources/container_app_environment_managed_certificate#subject_name ContainerAppEnvironmentManagedCertificate#subject_name}
	SubjectName *string `field:"required" json:"subjectName" yaml:"subjectName"`
	// The domain control validation type for the managed certificate. Possible values are `CNAME` and `HTTP`. Defaults to `HTTP`.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.80.0/docs/resources/container_app_environment_managed_certificate#domain_control_validation ContainerAppEnvironmentManagedCertificate#domain_control_validation}
	DomainControlValidation *string `field:"optional" json:"domainControlValidation" yaml:"domainControlValidation"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.80.0/docs/resources/container_app_environment_managed_certificate#id ContainerAppEnvironmentManagedCertificate#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.80.0/docs/resources/container_app_environment_managed_certificate#tags ContainerAppEnvironmentManagedCertificate#tags}.
	Tags *map[string]*string `field:"optional" json:"tags" yaml:"tags"`
	// timeouts block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.80.0/docs/resources/container_app_environment_managed_certificate#timeouts ContainerAppEnvironmentManagedCertificate#timeouts}
	Timeouts *ContainerAppEnvironmentManagedCertificateTimeouts `field:"optional" json:"timeouts" yaml:"timeouts"`
}

