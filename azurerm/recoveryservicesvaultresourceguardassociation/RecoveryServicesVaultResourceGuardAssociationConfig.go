// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package recoveryservicesvaultresourceguardassociation

import (
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type RecoveryServicesVaultResourceGuardAssociationConfig struct {
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
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.69.0/docs/resources/recovery_services_vault_resource_guard_association#resource_guard_id RecoveryServicesVaultResourceGuardAssociation#resource_guard_id}.
	ResourceGuardId *string `field:"required" json:"resourceGuardId" yaml:"resourceGuardId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.69.0/docs/resources/recovery_services_vault_resource_guard_association#vault_id RecoveryServicesVaultResourceGuardAssociation#vault_id}.
	VaultId *string `field:"required" json:"vaultId" yaml:"vaultId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.69.0/docs/resources/recovery_services_vault_resource_guard_association#id RecoveryServicesVaultResourceGuardAssociation#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// timeouts block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.69.0/docs/resources/recovery_services_vault_resource_guard_association#timeouts RecoveryServicesVaultResourceGuardAssociation#timeouts}
	Timeouts *RecoveryServicesVaultResourceGuardAssociationTimeouts `field:"optional" json:"timeouts" yaml:"timeouts"`
}

