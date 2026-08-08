// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package machinelearninginferencecluster


type MachineLearningInferenceClusterIdentity struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/5.0.1/docs/resources/machine_learning_inference_cluster#type MachineLearningInferenceCluster#type}.
	Type *string `field:"required" json:"type" yaml:"type"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/5.0.1/docs/resources/machine_learning_inference_cluster#identity_ids MachineLearningInferenceCluster#identity_ids}.
	IdentityIds *[]*string `field:"optional" json:"identityIds" yaml:"identityIds"`
}

