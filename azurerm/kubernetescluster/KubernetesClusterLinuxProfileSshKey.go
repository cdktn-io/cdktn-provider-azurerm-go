// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package kubernetescluster


type KubernetesClusterLinuxProfileSshKey struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.62.1/docs/resources/kubernetes_cluster#key_data KubernetesCluster#key_data}.
	KeyData *string `field:"required" json:"keyData" yaml:"keyData"`
}

