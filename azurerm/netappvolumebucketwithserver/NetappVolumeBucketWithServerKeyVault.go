// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package netappvolumebucketwithserver


type NetappVolumeBucketWithServerKeyVault struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/5.1.0/docs/resources/netapp_volume_bucket_with_server#certificate_key_vault_uri NetappVolumeBucketWithServer#certificate_key_vault_uri}.
	CertificateKeyVaultUri *string `field:"required" json:"certificateKeyVaultUri" yaml:"certificateKeyVaultUri"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/5.1.0/docs/resources/netapp_volume_bucket_with_server#certificate_name NetappVolumeBucketWithServer#certificate_name}.
	CertificateName *string `field:"required" json:"certificateName" yaml:"certificateName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/5.1.0/docs/resources/netapp_volume_bucket_with_server#credentials_key_vault_uri NetappVolumeBucketWithServer#credentials_key_vault_uri}.
	CredentialsKeyVaultUri *string `field:"required" json:"credentialsKeyVaultUri" yaml:"credentialsKeyVaultUri"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/5.1.0/docs/resources/netapp_volume_bucket_with_server#credentials_secret_name NetappVolumeBucketWithServer#credentials_secret_name}.
	CredentialsSecretName *string `field:"required" json:"credentialsSecretName" yaml:"credentialsSecretName"`
}

