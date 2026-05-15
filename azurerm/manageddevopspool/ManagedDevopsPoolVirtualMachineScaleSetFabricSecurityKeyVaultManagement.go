// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package manageddevopspool


type ManagedDevopsPoolVirtualMachineScaleSetFabricSecurityKeyVaultManagement struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.73.0/docs/resources/managed_devops_pool#key_vault_certificate_ids ManagedDevopsPool#key_vault_certificate_ids}.
	KeyVaultCertificateIds *[]*string `field:"required" json:"keyVaultCertificateIds" yaml:"keyVaultCertificateIds"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.73.0/docs/resources/managed_devops_pool#certificate_store_location ManagedDevopsPool#certificate_store_location}.
	CertificateStoreLocation *string `field:"optional" json:"certificateStoreLocation" yaml:"certificateStoreLocation"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.73.0/docs/resources/managed_devops_pool#certificate_store_name ManagedDevopsPool#certificate_store_name}.
	CertificateStoreName *string `field:"optional" json:"certificateStoreName" yaml:"certificateStoreName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.73.0/docs/resources/managed_devops_pool#key_export_enabled ManagedDevopsPool#key_export_enabled}.
	KeyExportEnabled interface{} `field:"optional" json:"keyExportEnabled" yaml:"keyExportEnabled"`
}

