// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package redhatopenshiftcluster


type RedhatOpenshiftClusterNetworkProfileLoadBalancerProfile struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/5.3.0/docs/resources/redhat_openshift_cluster#managed_outbound_ip_count RedhatOpenshiftCluster#managed_outbound_ip_count}.
	ManagedOutboundIpCount *float64 `field:"required" json:"managedOutboundIpCount" yaml:"managedOutboundIpCount"`
}

