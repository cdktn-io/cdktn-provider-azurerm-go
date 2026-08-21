// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package kubernetesclusterdeploymentsafeguard

import (
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type KubernetesClusterDeploymentSafeguardConfig struct {
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
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/5.2.0/docs/resources/kubernetes_cluster_deployment_safeguard#kubernetes_cluster_id KubernetesClusterDeploymentSafeguard#kubernetes_cluster_id}.
	KubernetesClusterId *string `field:"required" json:"kubernetesClusterId" yaml:"kubernetesClusterId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/5.2.0/docs/resources/kubernetes_cluster_deployment_safeguard#level KubernetesClusterDeploymentSafeguard#level}.
	Level *string `field:"required" json:"level" yaml:"level"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/5.2.0/docs/resources/kubernetes_cluster_deployment_safeguard#excluded_namespaces KubernetesClusterDeploymentSafeguard#excluded_namespaces}.
	ExcludedNamespaces *[]*string `field:"optional" json:"excludedNamespaces" yaml:"excludedNamespaces"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/5.2.0/docs/resources/kubernetes_cluster_deployment_safeguard#id KubernetesClusterDeploymentSafeguard#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/5.2.0/docs/resources/kubernetes_cluster_deployment_safeguard#pod_security_standards_level KubernetesClusterDeploymentSafeguard#pod_security_standards_level}.
	PodSecurityStandardsLevel *string `field:"optional" json:"podSecurityStandardsLevel" yaml:"podSecurityStandardsLevel"`
	// timeouts block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/5.2.0/docs/resources/kubernetes_cluster_deployment_safeguard#timeouts KubernetesClusterDeploymentSafeguard#timeouts}
	Timeouts *KubernetesClusterDeploymentSafeguardTimeouts `field:"optional" json:"timeouts" yaml:"timeouts"`
}

