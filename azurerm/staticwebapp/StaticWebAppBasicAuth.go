// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package staticwebapp


type StaticWebAppBasicAuth struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.69.0/docs/resources/static_web_app#environments StaticWebApp#environments}.
	Environments *string `field:"required" json:"environments" yaml:"environments"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.69.0/docs/resources/static_web_app#password StaticWebApp#password}.
	Password *string `field:"required" json:"password" yaml:"password"`
}

