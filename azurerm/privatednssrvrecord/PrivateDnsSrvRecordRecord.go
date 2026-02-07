// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package privatednssrvrecord


type PrivateDnsSrvRecordRecord struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.59.0/docs/resources/private_dns_srv_record#port PrivateDnsSrvRecord#port}.
	Port *float64 `field:"required" json:"port" yaml:"port"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.59.0/docs/resources/private_dns_srv_record#priority PrivateDnsSrvRecord#priority}.
	Priority *float64 `field:"required" json:"priority" yaml:"priority"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.59.0/docs/resources/private_dns_srv_record#target PrivateDnsSrvRecord#target}.
	Target *string `field:"required" json:"target" yaml:"target"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.59.0/docs/resources/private_dns_srv_record#weight PrivateDnsSrvRecord#weight}.
	Weight *float64 `field:"required" json:"weight" yaml:"weight"`
}

