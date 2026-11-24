// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package apimanagementbackend


type ApiManagementBackendTls struct {
	// Flag indicating whether SSL certificate chain validation should be done when using self-signed certificates for the backend host.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/5.15.0/docs/resources/api_management_backend#validate_certificate_chain ApiManagementBackend#validate_certificate_chain}
	ValidateCertificateChain interface{} `field:"optional" json:"validateCertificateChain" yaml:"validateCertificateChain"`
	// Flag indicating whether SSL certificate name validation should be done when using self-signed certificates for the backend host.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/5.15.0/docs/resources/api_management_backend#validate_certificate_name ApiManagementBackend#validate_certificate_name}
	ValidateCertificateName interface{} `field:"optional" json:"validateCertificateName" yaml:"validateCertificateName"`
}

