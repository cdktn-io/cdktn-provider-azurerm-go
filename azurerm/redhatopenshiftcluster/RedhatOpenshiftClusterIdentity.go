// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package redhatopenshiftcluster


type RedhatOpenshiftClusterIdentity struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/5.3.0/docs/resources/redhat_openshift_cluster#identity_ids RedhatOpenshiftCluster#identity_ids}.
	IdentityIds *[]*string `field:"required" json:"identityIds" yaml:"identityIds"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/5.3.0/docs/resources/redhat_openshift_cluster#type RedhatOpenshiftCluster#type}.
	Type *string `field:"required" json:"type" yaml:"type"`
}

