// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package containerappenvironmentcertificate


type ContainerAppEnvironmentCertificateCertificateKeyVault struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.67.0/docs/resources/container_app_environment_certificate#key_vault_secret_id ContainerAppEnvironmentCertificate#key_vault_secret_id}.
	KeyVaultSecretId *string `field:"required" json:"keyVaultSecretId" yaml:"keyVaultSecretId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.67.0/docs/resources/container_app_environment_certificate#identity ContainerAppEnvironmentCertificate#identity}.
	Identity *string `field:"optional" json:"identity" yaml:"identity"`
}

