// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package netappaccount


type NetappAccountActiveDirectoryLdapSearchScope struct {
	// Specifies the group DN, which overrides the base DN for group lookups.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/5.14.0/docs/resources/netapp_account#group_dn NetappAccount#group_dn}
	GroupDn *string `field:"optional" json:"groupDn" yaml:"groupDn"`
	// Specifies the custom LDAP search filter to be used when looking up group membership from LDAP server.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/5.14.0/docs/resources/netapp_account#group_membership_filter NetappAccount#group_membership_filter}
	GroupMembershipFilter *string `field:"optional" json:"groupMembershipFilter" yaml:"groupMembershipFilter"`
	// Specifies the user DN, which overrides the base DN for user lookups.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/5.14.0/docs/resources/netapp_account#user_dn NetappAccount#user_dn}
	UserDn *string `field:"optional" json:"userDn" yaml:"userDn"`
}

