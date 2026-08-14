// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package kubernetesautomaticcluster


type KubernetesAutomaticClusterServiceMesh struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/5.1.0/docs/resources/kubernetes_automatic_cluster#revisions KubernetesAutomaticCluster#revisions}.
	Revisions *[]*string `field:"required" json:"revisions" yaml:"revisions"`
	// certificate_authority block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/5.1.0/docs/resources/kubernetes_automatic_cluster#certificate_authority KubernetesAutomaticCluster#certificate_authority}
	CertificateAuthority *KubernetesAutomaticClusterServiceMeshCertificateAuthority `field:"optional" json:"certificateAuthority" yaml:"certificateAuthority"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/5.1.0/docs/resources/kubernetes_automatic_cluster#external_ingress_gateway_enabled KubernetesAutomaticCluster#external_ingress_gateway_enabled}.
	ExternalIngressGatewayEnabled interface{} `field:"optional" json:"externalIngressGatewayEnabled" yaml:"externalIngressGatewayEnabled"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/5.1.0/docs/resources/kubernetes_automatic_cluster#internal_ingress_gateway_enabled KubernetesAutomaticCluster#internal_ingress_gateway_enabled}.
	InternalIngressGatewayEnabled interface{} `field:"optional" json:"internalIngressGatewayEnabled" yaml:"internalIngressGatewayEnabled"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/5.1.0/docs/resources/kubernetes_automatic_cluster#proxy_redirect_mechanism KubernetesAutomaticCluster#proxy_redirect_mechanism}.
	ProxyRedirectMechanism *string `field:"optional" json:"proxyRedirectMechanism" yaml:"proxyRedirectMechanism"`
}

