// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package applicationgateway


type ApplicationGatewayBackendHttpSettings struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/5.2.0/docs/resources/application_gateway#cookie_based_affinity ApplicationGateway#cookie_based_affinity}.
	CookieBasedAffinity *string `field:"required" json:"cookieBasedAffinity" yaml:"cookieBasedAffinity"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/5.2.0/docs/resources/application_gateway#name ApplicationGateway#name}.
	Name *string `field:"required" json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/5.2.0/docs/resources/application_gateway#port ApplicationGateway#port}.
	Port *float64 `field:"required" json:"port" yaml:"port"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/5.2.0/docs/resources/application_gateway#protocol ApplicationGateway#protocol}.
	Protocol *string `field:"required" json:"protocol" yaml:"protocol"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/5.2.0/docs/resources/application_gateway#affinity_cookie_name ApplicationGateway#affinity_cookie_name}.
	AffinityCookieName *string `field:"optional" json:"affinityCookieName" yaml:"affinityCookieName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/5.2.0/docs/resources/application_gateway#certificate_chain_validation_enabled ApplicationGateway#certificate_chain_validation_enabled}.
	CertificateChainValidationEnabled interface{} `field:"optional" json:"certificateChainValidationEnabled" yaml:"certificateChainValidationEnabled"`
	// connection_draining block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/5.2.0/docs/resources/application_gateway#connection_draining ApplicationGateway#connection_draining}
	ConnectionDraining *ApplicationGatewayBackendHttpSettingsConnectionDraining `field:"optional" json:"connectionDraining" yaml:"connectionDraining"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/5.2.0/docs/resources/application_gateway#dedicated_backend_connection_enabled ApplicationGateway#dedicated_backend_connection_enabled}.
	DedicatedBackendConnectionEnabled interface{} `field:"optional" json:"dedicatedBackendConnectionEnabled" yaml:"dedicatedBackendConnectionEnabled"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/5.2.0/docs/resources/application_gateway#host_name ApplicationGateway#host_name}.
	HostName *string `field:"optional" json:"hostName" yaml:"hostName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/5.2.0/docs/resources/application_gateway#path ApplicationGateway#path}.
	Path *string `field:"optional" json:"path" yaml:"path"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/5.2.0/docs/resources/application_gateway#pick_host_name_from_backend_address ApplicationGateway#pick_host_name_from_backend_address}.
	PickHostNameFromBackendAddress interface{} `field:"optional" json:"pickHostNameFromBackendAddress" yaml:"pickHostNameFromBackendAddress"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/5.2.0/docs/resources/application_gateway#probe_name ApplicationGateway#probe_name}.
	ProbeName *string `field:"optional" json:"probeName" yaml:"probeName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/5.2.0/docs/resources/application_gateway#request_timeout ApplicationGateway#request_timeout}.
	RequestTimeout *float64 `field:"optional" json:"requestTimeout" yaml:"requestTimeout"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/5.2.0/docs/resources/application_gateway#sni_name ApplicationGateway#sni_name}.
	SniName *string `field:"optional" json:"sniName" yaml:"sniName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/5.2.0/docs/resources/application_gateway#sni_validation_enabled ApplicationGateway#sni_validation_enabled}.
	SniValidationEnabled interface{} `field:"optional" json:"sniValidationEnabled" yaml:"sniValidationEnabled"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/5.2.0/docs/resources/application_gateway#trusted_root_certificate_names ApplicationGateway#trusted_root_certificate_names}.
	TrustedRootCertificateNames *[]*string `field:"optional" json:"trustedRootCertificateNames" yaml:"trustedRootCertificateNames"`
}

