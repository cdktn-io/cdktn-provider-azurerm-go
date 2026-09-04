// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package redhatopenshiftcluster


type RedhatOpenshiftClusterPlatformWorkloadIdentityProfilePlatformWorkloadIdentity struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/5.4.0/docs/resources/redhat_openshift_cluster#identity_id RedhatOpenshiftCluster#identity_id}.
	IdentityId *string `field:"required" json:"identityId" yaml:"identityId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/5.4.0/docs/resources/redhat_openshift_cluster#name RedhatOpenshiftCluster#name}.
	Name *string `field:"required" json:"name" yaml:"name"`
}

