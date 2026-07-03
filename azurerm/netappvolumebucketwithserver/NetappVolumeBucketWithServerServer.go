// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package netappvolumebucketwithserver


type NetappVolumeBucketWithServerServer struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.80.0/docs/resources/netapp_volume_bucket_with_server#fqdn NetappVolumeBucketWithServer#fqdn}.
	Fqdn *string `field:"required" json:"fqdn" yaml:"fqdn"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.80.0/docs/resources/netapp_volume_bucket_with_server#certificate_pem NetappVolumeBucketWithServer#certificate_pem}.
	CertificatePem *string `field:"optional" json:"certificatePem" yaml:"certificatePem"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.80.0/docs/resources/netapp_volume_bucket_with_server#on_certificate_conflict_action NetappVolumeBucketWithServer#on_certificate_conflict_action}.
	OnCertificateConflictAction *string `field:"optional" json:"onCertificateConflictAction" yaml:"onCertificateConflictAction"`
}

