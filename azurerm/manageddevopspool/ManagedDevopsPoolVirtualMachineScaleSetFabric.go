// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package manageddevopspool


type ManagedDevopsPoolVirtualMachineScaleSetFabric struct {
	// image block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.81.0/docs/resources/managed_devops_pool#image ManagedDevopsPool#image}
	Image interface{} `field:"required" json:"image" yaml:"image"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.81.0/docs/resources/managed_devops_pool#sku_name ManagedDevopsPool#sku_name}.
	SkuName *string `field:"required" json:"skuName" yaml:"skuName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.81.0/docs/resources/managed_devops_pool#os_disk_storage_account_type ManagedDevopsPool#os_disk_storage_account_type}.
	OsDiskStorageAccountType *string `field:"optional" json:"osDiskStorageAccountType" yaml:"osDiskStorageAccountType"`
	// security block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.81.0/docs/resources/managed_devops_pool#security ManagedDevopsPool#security}
	Security *ManagedDevopsPoolVirtualMachineScaleSetFabricSecurity `field:"optional" json:"security" yaml:"security"`
	// storage block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.81.0/docs/resources/managed_devops_pool#storage ManagedDevopsPool#storage}
	Storage *ManagedDevopsPoolVirtualMachineScaleSetFabricStorage `field:"optional" json:"storage" yaml:"storage"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.81.0/docs/resources/managed_devops_pool#subnet_id ManagedDevopsPool#subnet_id}.
	SubnetId *string `field:"optional" json:"subnetId" yaml:"subnetId"`
}

