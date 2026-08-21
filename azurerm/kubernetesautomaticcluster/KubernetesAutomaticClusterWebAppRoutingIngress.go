// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package kubernetesautomaticcluster


type KubernetesAutomaticClusterWebAppRoutingIngress struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/5.2.0/docs/resources/kubernetes_automatic_cluster#default_nginx_controller KubernetesAutomaticCluster#default_nginx_controller}.
	DefaultNginxController *string `field:"optional" json:"defaultNginxController" yaml:"defaultNginxController"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/5.2.0/docs/resources/kubernetes_automatic_cluster#dns_zone_ids KubernetesAutomaticCluster#dns_zone_ids}.
	DnsZoneIds *[]*string `field:"optional" json:"dnsZoneIds" yaml:"dnsZoneIds"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/5.2.0/docs/resources/kubernetes_automatic_cluster#istio_enabled KubernetesAutomaticCluster#istio_enabled}.
	IstioEnabled interface{} `field:"optional" json:"istioEnabled" yaml:"istioEnabled"`
}

