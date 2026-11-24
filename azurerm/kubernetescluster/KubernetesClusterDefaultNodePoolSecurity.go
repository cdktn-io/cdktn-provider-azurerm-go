// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package kubernetescluster


type KubernetesClusterDefaultNodePoolSecurity struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/5.15.0/docs/resources/kubernetes_cluster#secure_boot_enabled KubernetesCluster#secure_boot_enabled}.
	SecureBootEnabled interface{} `field:"optional" json:"secureBootEnabled" yaml:"secureBootEnabled"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/5.15.0/docs/resources/kubernetes_cluster#vtpm_enabled KubernetesCluster#vtpm_enabled}.
	VtpmEnabled interface{} `field:"optional" json:"vtpmEnabled" yaml:"vtpmEnabled"`
}

