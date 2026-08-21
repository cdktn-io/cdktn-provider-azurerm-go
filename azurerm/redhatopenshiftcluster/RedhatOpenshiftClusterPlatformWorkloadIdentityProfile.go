// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package redhatopenshiftcluster


type RedhatOpenshiftClusterPlatformWorkloadIdentityProfile struct {
	// platform_workload_identity block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/5.2.0/docs/resources/redhat_openshift_cluster#platform_workload_identity RedhatOpenshiftCluster#platform_workload_identity}
	PlatformWorkloadIdentity interface{} `field:"required" json:"platformWorkloadIdentity" yaml:"platformWorkloadIdentity"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/5.2.0/docs/resources/redhat_openshift_cluster#upgradeable_to RedhatOpenshiftCluster#upgradeable_to}.
	UpgradeableTo *string `field:"optional" json:"upgradeableTo" yaml:"upgradeableTo"`
}

