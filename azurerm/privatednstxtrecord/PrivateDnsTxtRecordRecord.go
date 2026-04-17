// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package privatednstxtrecord


type PrivateDnsTxtRecordRecord struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.69.0/docs/resources/private_dns_txt_record#value PrivateDnsTxtRecord#value}.
	Value *string `field:"required" json:"value" yaml:"value"`
}

