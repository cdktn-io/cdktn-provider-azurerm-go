// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package kubernetesautomaticcluster


type KubernetesAutomaticClusterServiceMeshCertificateAuthority struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/5.0.1/docs/resources/kubernetes_automatic_cluster#certificate_chain_object_name KubernetesAutomaticCluster#certificate_chain_object_name}.
	CertificateChainObjectName *string `field:"required" json:"certificateChainObjectName" yaml:"certificateChainObjectName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/5.0.1/docs/resources/kubernetes_automatic_cluster#certificate_object_name KubernetesAutomaticCluster#certificate_object_name}.
	CertificateObjectName *string `field:"required" json:"certificateObjectName" yaml:"certificateObjectName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/5.0.1/docs/resources/kubernetes_automatic_cluster#key_object_name KubernetesAutomaticCluster#key_object_name}.
	KeyObjectName *string `field:"required" json:"keyObjectName" yaml:"keyObjectName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/5.0.1/docs/resources/kubernetes_automatic_cluster#key_vault_id KubernetesAutomaticCluster#key_vault_id}.
	KeyVaultId *string `field:"required" json:"keyVaultId" yaml:"keyVaultId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/5.0.1/docs/resources/kubernetes_automatic_cluster#root_certificate_object_name KubernetesAutomaticCluster#root_certificate_object_name}.
	RootCertificateObjectName *string `field:"required" json:"rootCertificateObjectName" yaml:"rootCertificateObjectName"`
}

