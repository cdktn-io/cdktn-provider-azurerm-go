// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package mongocluster


type MongoClusterCustomerManagedKey struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.66.0/docs/resources/mongo_cluster#key_vault_key_id MongoCluster#key_vault_key_id}.
	KeyVaultKeyId *string `field:"required" json:"keyVaultKeyId" yaml:"keyVaultKeyId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.66.0/docs/resources/mongo_cluster#user_assigned_identity_id MongoCluster#user_assigned_identity_id}.
	UserAssignedIdentityId *string `field:"required" json:"userAssignedIdentityId" yaml:"userAssignedIdentityId"`
}

