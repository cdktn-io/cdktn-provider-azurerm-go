// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package kubernetescluster


type KubernetesClusterNodeProvisioningProfile struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/5.3.0/docs/resources/kubernetes_cluster#default_node_pools KubernetesCluster#default_node_pools}.
	DefaultNodePools *string `field:"optional" json:"defaultNodePools" yaml:"defaultNodePools"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/5.3.0/docs/resources/kubernetes_cluster#mode KubernetesCluster#mode}.
	Mode *string `field:"optional" json:"mode" yaml:"mode"`
}

