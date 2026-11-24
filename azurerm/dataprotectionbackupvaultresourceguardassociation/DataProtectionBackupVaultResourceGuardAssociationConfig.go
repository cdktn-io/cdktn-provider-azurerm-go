// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package dataprotectionbackupvaultresourceguardassociation

import (
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type DataProtectionBackupVaultResourceGuardAssociationConfig struct {
	// Experimental.
	Connection interface{} `field:"optional" json:"connection" yaml:"connection"`
	// Experimental.
	Count interface{} `field:"optional" json:"count" yaml:"count"`
	// Experimental.
	DependsOn *[]cdktn.ITerraformDependable `field:"optional" json:"dependsOn" yaml:"dependsOn"`
	// Experimental.
	ForEach cdktn.ITerraformIterator `field:"optional" json:"forEach" yaml:"forEach"`
	// Experimental.
	Lifecycle *cdktn.TerraformResourceLifecycle `field:"optional" json:"lifecycle" yaml:"lifecycle"`
	// Experimental.
	Provider cdktn.TerraformProvider `field:"optional" json:"provider" yaml:"provider"`
	// Experimental.
	Provisioners *[]interface{} `field:"optional" json:"provisioners" yaml:"provisioners"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/5.15.0/docs/resources/data_protection_backup_vault_resource_guard_association#data_protection_backup_vault_id DataProtectionBackupVaultResourceGuardAssociation#data_protection_backup_vault_id}.
	DataProtectionBackupVaultId *string `field:"required" json:"dataProtectionBackupVaultId" yaml:"dataProtectionBackupVaultId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/5.15.0/docs/resources/data_protection_backup_vault_resource_guard_association#data_protection_resource_guard_id DataProtectionBackupVaultResourceGuardAssociation#data_protection_resource_guard_id}.
	DataProtectionResourceGuardId *string `field:"required" json:"dataProtectionResourceGuardId" yaml:"dataProtectionResourceGuardId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/5.15.0/docs/resources/data_protection_backup_vault_resource_guard_association#id DataProtectionBackupVaultResourceGuardAssociation#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// timeouts block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/5.15.0/docs/resources/data_protection_backup_vault_resource_guard_association#timeouts DataProtectionBackupVaultResourceGuardAssociation#timeouts}
	Timeouts *DataProtectionBackupVaultResourceGuardAssociationTimeouts `field:"optional" json:"timeouts" yaml:"timeouts"`
}

