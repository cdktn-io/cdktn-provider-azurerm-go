// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package kubernetesautomaticcluster


type KubernetesAutomaticClusterHostedSystem struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.81.0/docs/resources/kubernetes_automatic_cluster#node_subnet_id KubernetesAutomaticCluster#node_subnet_id}.
	NodeSubnetId *string `field:"required" json:"nodeSubnetId" yaml:"nodeSubnetId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.81.0/docs/resources/kubernetes_automatic_cluster#system_node_subnet_id KubernetesAutomaticCluster#system_node_subnet_id}.
	SystemNodeSubnetId *string `field:"required" json:"systemNodeSubnetId" yaml:"systemNodeSubnetId"`
}

