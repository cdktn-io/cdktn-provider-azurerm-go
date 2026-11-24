// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package horizondbcluster


type HorizondbClusterReadWritePoolCompute struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/5.15.0/docs/resources/horizondb_cluster#generation HorizondbCluster#generation}.
	Generation *float64 `field:"required" json:"generation" yaml:"generation"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/5.15.0/docs/resources/horizondb_cluster#processor HorizondbCluster#processor}.
	Processor *string `field:"required" json:"processor" yaml:"processor"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/5.15.0/docs/resources/horizondb_cluster#vcores HorizondbCluster#vcores}.
	Vcores *float64 `field:"required" json:"vcores" yaml:"vcores"`
}

