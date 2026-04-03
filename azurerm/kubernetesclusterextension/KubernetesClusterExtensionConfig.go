// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package kubernetesclusterextension

import (
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type KubernetesClusterExtensionConfig struct {
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
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.67.0/docs/resources/kubernetes_cluster_extension#cluster_id KubernetesClusterExtension#cluster_id}.
	ClusterId *string `field:"required" json:"clusterId" yaml:"clusterId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.67.0/docs/resources/kubernetes_cluster_extension#extension_type KubernetesClusterExtension#extension_type}.
	ExtensionType *string `field:"required" json:"extensionType" yaml:"extensionType"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.67.0/docs/resources/kubernetes_cluster_extension#name KubernetesClusterExtension#name}.
	Name *string `field:"required" json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.67.0/docs/resources/kubernetes_cluster_extension#configuration_protected_settings KubernetesClusterExtension#configuration_protected_settings}.
	ConfigurationProtectedSettings *map[string]*string `field:"optional" json:"configurationProtectedSettings" yaml:"configurationProtectedSettings"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.67.0/docs/resources/kubernetes_cluster_extension#configuration_settings KubernetesClusterExtension#configuration_settings}.
	ConfigurationSettings *map[string]*string `field:"optional" json:"configurationSettings" yaml:"configurationSettings"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.67.0/docs/resources/kubernetes_cluster_extension#id KubernetesClusterExtension#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// plan block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.67.0/docs/resources/kubernetes_cluster_extension#plan KubernetesClusterExtension#plan}
	Plan *KubernetesClusterExtensionPlan `field:"optional" json:"plan" yaml:"plan"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.67.0/docs/resources/kubernetes_cluster_extension#release_namespace KubernetesClusterExtension#release_namespace}.
	ReleaseNamespace *string `field:"optional" json:"releaseNamespace" yaml:"releaseNamespace"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.67.0/docs/resources/kubernetes_cluster_extension#release_train KubernetesClusterExtension#release_train}.
	ReleaseTrain *string `field:"optional" json:"releaseTrain" yaml:"releaseTrain"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.67.0/docs/resources/kubernetes_cluster_extension#target_namespace KubernetesClusterExtension#target_namespace}.
	TargetNamespace *string `field:"optional" json:"targetNamespace" yaml:"targetNamespace"`
	// timeouts block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.67.0/docs/resources/kubernetes_cluster_extension#timeouts KubernetesClusterExtension#timeouts}
	Timeouts *KubernetesClusterExtensionTimeouts `field:"optional" json:"timeouts" yaml:"timeouts"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.67.0/docs/resources/kubernetes_cluster_extension#version KubernetesClusterExtension#version}.
	Version *string `field:"optional" json:"version" yaml:"version"`
}

