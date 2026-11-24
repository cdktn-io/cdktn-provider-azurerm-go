// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package apimanagement


type ApiManagementSecurity struct {
	// Should SSL 3.0 be enabled on the backend of the gateway?
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/5.15.0/docs/resources/api_management#backend_ssl30_enabled ApiManagement#backend_ssl30_enabled}
	BackendSsl30Enabled interface{} `field:"optional" json:"backendSsl30Enabled" yaml:"backendSsl30Enabled"`
	// Should TLS 1.0 be enabled on the backend of the gateway?
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/5.15.0/docs/resources/api_management#backend_tls10_enabled ApiManagement#backend_tls10_enabled}
	BackendTls10Enabled interface{} `field:"optional" json:"backendTls10Enabled" yaml:"backendTls10Enabled"`
	// Should TLS 1.1 be enabled on the backend of the gateway?
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/5.15.0/docs/resources/api_management#backend_tls11_enabled ApiManagement#backend_tls11_enabled}
	BackendTls11Enabled interface{} `field:"optional" json:"backendTls11Enabled" yaml:"backendTls11Enabled"`
	// Should SSL 3.0 be enabled on the frontend of the gateway?
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/5.15.0/docs/resources/api_management#frontend_ssl30_enabled ApiManagement#frontend_ssl30_enabled}
	FrontendSsl30Enabled interface{} `field:"optional" json:"frontendSsl30Enabled" yaml:"frontendSsl30Enabled"`
	// Should TLS 1.0 be enabled on the frontend of the gateway?
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/5.15.0/docs/resources/api_management#frontend_tls10_enabled ApiManagement#frontend_tls10_enabled}
	FrontendTls10Enabled interface{} `field:"optional" json:"frontendTls10Enabled" yaml:"frontendTls10Enabled"`
	// Should TLS 1.1 be enabled on the frontend of the gateway?
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/5.15.0/docs/resources/api_management#frontend_tls11_enabled ApiManagement#frontend_tls11_enabled}
	FrontendTls11Enabled interface{} `field:"optional" json:"frontendTls11Enabled" yaml:"frontendTls11Enabled"`
	// Should the `TLS_ECDHE_ECDSA_WITH_AES_128_CBC_SHA` cipher be enabled?
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/5.15.0/docs/resources/api_management#tls_ecdhe_ecdsa_with_aes128_cbc_sha_ciphers_enabled ApiManagement#tls_ecdhe_ecdsa_with_aes128_cbc_sha_ciphers_enabled}
	TlsEcdheEcdsaWithAes128CbcShaCiphersEnabled interface{} `field:"optional" json:"tlsEcdheEcdsaWithAes128CbcShaCiphersEnabled" yaml:"tlsEcdheEcdsaWithAes128CbcShaCiphersEnabled"`
	// Should the `TLS_ECDHE_ECDSA_WITH_AES_256_CBC_SHA` cipher be enabled?
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/5.15.0/docs/resources/api_management#tls_ecdhe_ecdsa_with_aes256_cbc_sha_ciphers_enabled ApiManagement#tls_ecdhe_ecdsa_with_aes256_cbc_sha_ciphers_enabled}
	TlsEcdheEcdsaWithAes256CbcShaCiphersEnabled interface{} `field:"optional" json:"tlsEcdheEcdsaWithAes256CbcShaCiphersEnabled" yaml:"tlsEcdheEcdsaWithAes256CbcShaCiphersEnabled"`
	// Should the `TLS_ECDHE_RSA_WITH_AES_128_CBC_SHA` cipher be enabled?
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/5.15.0/docs/resources/api_management#tls_ecdhe_rsa_with_aes128_cbc_sha_ciphers_enabled ApiManagement#tls_ecdhe_rsa_with_aes128_cbc_sha_ciphers_enabled}
	TlsEcdheRsaWithAes128CbcShaCiphersEnabled interface{} `field:"optional" json:"tlsEcdheRsaWithAes128CbcShaCiphersEnabled" yaml:"tlsEcdheRsaWithAes128CbcShaCiphersEnabled"`
	// Should the `TLS_ECDHE_RSA_WITH_AES_256_CBC_SHA` cipher be enabled?
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/5.15.0/docs/resources/api_management#tls_ecdhe_rsa_with_aes256_cbc_sha_ciphers_enabled ApiManagement#tls_ecdhe_rsa_with_aes256_cbc_sha_ciphers_enabled}
	TlsEcdheRsaWithAes256CbcShaCiphersEnabled interface{} `field:"optional" json:"tlsEcdheRsaWithAes256CbcShaCiphersEnabled" yaml:"tlsEcdheRsaWithAes256CbcShaCiphersEnabled"`
	// Should the `TLS_RSA_WITH_AES_128_CBC_SHA256` cipher be enabled?
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/5.15.0/docs/resources/api_management#tls_rsa_with_aes128_cbc_sha256_ciphers_enabled ApiManagement#tls_rsa_with_aes128_cbc_sha256_ciphers_enabled}
	TlsRsaWithAes128CbcSha256CiphersEnabled interface{} `field:"optional" json:"tlsRsaWithAes128CbcSha256CiphersEnabled" yaml:"tlsRsaWithAes128CbcSha256CiphersEnabled"`
	// Should the `TLS_RSA_WITH_AES_128_CBC_SHA` cipher be enabled?
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/5.15.0/docs/resources/api_management#tls_rsa_with_aes128_cbc_sha_ciphers_enabled ApiManagement#tls_rsa_with_aes128_cbc_sha_ciphers_enabled}
	TlsRsaWithAes128CbcShaCiphersEnabled interface{} `field:"optional" json:"tlsRsaWithAes128CbcShaCiphersEnabled" yaml:"tlsRsaWithAes128CbcShaCiphersEnabled"`
	// Should the `TLS_RSA_WITH_AES_128_GCM_SHA256` cipher be enabled?
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/5.15.0/docs/resources/api_management#tls_rsa_with_aes128_gcm_sha256_ciphers_enabled ApiManagement#tls_rsa_with_aes128_gcm_sha256_ciphers_enabled}
	TlsRsaWithAes128GcmSha256CiphersEnabled interface{} `field:"optional" json:"tlsRsaWithAes128GcmSha256CiphersEnabled" yaml:"tlsRsaWithAes128GcmSha256CiphersEnabled"`
	// Should the `TLS_RSA_WITH_AES_256_CBC_SHA256` cipher be enabled?
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/5.15.0/docs/resources/api_management#tls_rsa_with_aes256_cbc_sha256_ciphers_enabled ApiManagement#tls_rsa_with_aes256_cbc_sha256_ciphers_enabled}
	TlsRsaWithAes256CbcSha256CiphersEnabled interface{} `field:"optional" json:"tlsRsaWithAes256CbcSha256CiphersEnabled" yaml:"tlsRsaWithAes256CbcSha256CiphersEnabled"`
	// Should the `TLS_RSA_WITH_AES_256_CBC_SHA` cipher be enabled?
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/5.15.0/docs/resources/api_management#tls_rsa_with_aes256_cbc_sha_ciphers_enabled ApiManagement#tls_rsa_with_aes256_cbc_sha_ciphers_enabled}
	TlsRsaWithAes256CbcShaCiphersEnabled interface{} `field:"optional" json:"tlsRsaWithAes256CbcShaCiphersEnabled" yaml:"tlsRsaWithAes256CbcShaCiphersEnabled"`
	// Should the `TLS_RSA_WITH_AES_256_GCM_SHA384` cipher be enabled?
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/5.15.0/docs/resources/api_management#tls_rsa_with_aes256_gcm_sha384_ciphers_enabled ApiManagement#tls_rsa_with_aes256_gcm_sha384_ciphers_enabled}
	TlsRsaWithAes256GcmSha384CiphersEnabled interface{} `field:"optional" json:"tlsRsaWithAes256GcmSha384CiphersEnabled" yaml:"tlsRsaWithAes256GcmSha384CiphersEnabled"`
	// Should the `TLS_RSA_WITH_3DES_EDE_CBC_SHA` cipher be enabled for alL TLS versions (1.0, 1.1 and 1.2)?
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/5.15.0/docs/resources/api_management#triple_des_ciphers_enabled ApiManagement#triple_des_ciphers_enabled}
	TripleDesCiphersEnabled interface{} `field:"optional" json:"tripleDesCiphersEnabled" yaml:"tripleDesCiphersEnabled"`
}

