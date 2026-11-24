// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package apimanagement


type ApiManagementHostnameConfigurationPortal struct {
	// The host name to use.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/5.15.0/docs/resources/api_management#host_name ApiManagement#host_name}
	HostName *string `field:"required" json:"hostName" yaml:"hostName"`
	// The Base64-encoded certificate.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/5.15.0/docs/resources/api_management#certificate ApiManagement#certificate}
	Certificate *string `field:"optional" json:"certificate" yaml:"certificate"`
	// The password associated with the certificate provided above.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/5.15.0/docs/resources/api_management#certificate_password ApiManagement#certificate_password}
	CertificatePassword *string `field:"optional" json:"certificatePassword" yaml:"certificatePassword"`
	// The ID of the Key Vault Secret containing the SSL Certificate, which must be of the type `application/x-pkcs12`.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/5.15.0/docs/resources/api_management#key_vault_certificate_id ApiManagement#key_vault_certificate_id}
	KeyVaultCertificateId *string `field:"optional" json:"keyVaultCertificateId" yaml:"keyVaultCertificateId"`
	// Should Client Certificate Negotiation be enabled for this Hostname?
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/5.15.0/docs/resources/api_management#negotiate_client_certificate ApiManagement#negotiate_client_certificate}
	NegotiateClientCertificate interface{} `field:"optional" json:"negotiateClientCertificate" yaml:"negotiateClientCertificate"`
	// The Managed Service Identity Client ID, which has `GET` access to the Key Vault containing the SSL certificate.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/5.15.0/docs/resources/api_management#ssl_keyvault_identity_client_id ApiManagement#ssl_keyvault_identity_client_id}
	SslKeyvaultIdentityClientId *string `field:"optional" json:"sslKeyvaultIdentityClientId" yaml:"sslKeyvaultIdentityClientId"`
}

