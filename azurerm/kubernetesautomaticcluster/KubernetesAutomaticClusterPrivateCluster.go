// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package kubernetesautomaticcluster


type KubernetesAutomaticClusterPrivateCluster struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/5.3.0/docs/resources/kubernetes_automatic_cluster#private_dns_zone_id KubernetesAutomaticCluster#private_dns_zone_id}.
	PrivateDnsZoneId *string `field:"optional" json:"privateDnsZoneId" yaml:"privateDnsZoneId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/5.3.0/docs/resources/kubernetes_automatic_cluster#public_fully_qualified_domain_name_enabled KubernetesAutomaticCluster#public_fully_qualified_domain_name_enabled}.
	PublicFullyQualifiedDomainNameEnabled interface{} `field:"optional" json:"publicFullyQualifiedDomainNameEnabled" yaml:"publicFullyQualifiedDomainNameEnabled"`
}

