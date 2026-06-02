// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package mysqlflexibleserver


type MysqlFlexibleServerHighAvailability struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.75.0/docs/resources/mysql_flexible_server#mode MysqlFlexibleServer#mode}.
	Mode *string `field:"required" json:"mode" yaml:"mode"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.75.0/docs/resources/mysql_flexible_server#standby_availability_zone MysqlFlexibleServer#standby_availability_zone}.
	StandbyAvailabilityZone *string `field:"optional" json:"standbyAvailabilityZone" yaml:"standbyAvailabilityZone"`
}

