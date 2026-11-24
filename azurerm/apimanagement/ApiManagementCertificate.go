// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package apimanagement


type ApiManagementCertificate struct {
	// The Base64 Encoded PFX or Base64 Encoded X.509 Certificate.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/5.15.0/docs/resources/api_management#encoded_certificate ApiManagement#encoded_certificate}
	EncodedCertificate *string `field:"required" json:"encodedCertificate" yaml:"encodedCertificate"`
	// The name of the Certificate Store where this certificate should be stored.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/5.15.0/docs/resources/api_management#store_name ApiManagement#store_name}
	StoreName *string `field:"required" json:"storeName" yaml:"storeName"`
	// The password for the certificate.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/5.15.0/docs/resources/api_management#certificate_password ApiManagement#certificate_password}
	CertificatePassword *string `field:"optional" json:"certificatePassword" yaml:"certificatePassword"`
}

