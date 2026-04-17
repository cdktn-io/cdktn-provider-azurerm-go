// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package mssqlserversecurityalertpolicy


type MssqlServerSecurityAlertPolicyTimeouts struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.69.0/docs/resources/mssql_server_security_alert_policy#create MssqlServerSecurityAlertPolicy#create}.
	Create *string `field:"optional" json:"create" yaml:"create"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.69.0/docs/resources/mssql_server_security_alert_policy#delete MssqlServerSecurityAlertPolicy#delete}.
	Delete *string `field:"optional" json:"delete" yaml:"delete"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.69.0/docs/resources/mssql_server_security_alert_policy#read MssqlServerSecurityAlertPolicy#read}.
	Read *string `field:"optional" json:"read" yaml:"read"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.69.0/docs/resources/mssql_server_security_alert_policy#update MssqlServerSecurityAlertPolicy#update}.
	Update *string `field:"optional" json:"update" yaml:"update"`
}

