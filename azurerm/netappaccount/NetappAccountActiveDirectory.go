// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package netappaccount


type NetappAccountActiveDirectory struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/5.15.0/docs/resources/netapp_account#dns_servers NetappAccount#dns_servers}.
	DnsServers *[]*string `field:"required" json:"dnsServers" yaml:"dnsServers"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/5.15.0/docs/resources/netapp_account#domain NetappAccount#domain}.
	Domain *string `field:"required" json:"domain" yaml:"domain"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/5.15.0/docs/resources/netapp_account#password NetappAccount#password}.
	Password *string `field:"required" json:"password" yaml:"password"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/5.15.0/docs/resources/netapp_account#smb_server_name NetappAccount#smb_server_name}.
	SmbServerName *string `field:"required" json:"smbServerName" yaml:"smbServerName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/5.15.0/docs/resources/netapp_account#username NetappAccount#username}.
	Username *string `field:"required" json:"username" yaml:"username"`
	// Domain users to be added to the built-in Administrators Active Directory group.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/5.15.0/docs/resources/netapp_account#administrator_users NetappAccount#administrator_users}
	AdministratorUsers *[]*string `field:"optional" json:"administratorUsers" yaml:"administratorUsers"`
	// If enabled, AES encryption will be enabled for SMB communication.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/5.15.0/docs/resources/netapp_account#aes_encryption_enabled NetappAccount#aes_encryption_enabled}
	AesEncryptionEnabled interface{} `field:"optional" json:"aesEncryptionEnabled" yaml:"aesEncryptionEnabled"`
	// Domain users to be added to the built-in Backup Operators Active Directory group.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/5.15.0/docs/resources/netapp_account#backup_users NetappAccount#backup_users}
	BackupUsers *[]*string `field:"optional" json:"backupUsers" yaml:"backupUsers"`
	// Whether traffic between the Domain Controller and SMB server is encrypted.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/5.15.0/docs/resources/netapp_account#domain_controller_connection_encryption_enabled NetappAccount#domain_controller_connection_encryption_enabled}
	DomainControllerConnectionEncryptionEnabled interface{} `field:"optional" json:"domainControllerConnectionEncryptionEnabled" yaml:"domainControllerConnectionEncryptionEnabled"`
	// Name of the Active Directory machine. This optional parameter is used only while creating kerberos volume.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/5.15.0/docs/resources/netapp_account#kerberos_ad_name NetappAccount#kerberos_ad_name}
	KerberosAdName *string `field:"optional" json:"kerberosAdName" yaml:"kerberosAdName"`
	// IP address of the KDC server (usually same the DC).
	//
	// This optional parameter is used only while creating kerberos volume.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/5.15.0/docs/resources/netapp_account#kerberos_kdc_ip NetappAccount#kerberos_kdc_ip}
	KerberosKdcIp *string `field:"optional" json:"kerberosKdcIp" yaml:"kerberosKdcIp"`
	// The preferred LDAP server IPs for the LDAP client.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/5.15.0/docs/resources/netapp_account#ldap_client_preferred_servers NetappAccount#ldap_client_preferred_servers}
	LdapClientPreferredServers *[]*string `field:"optional" json:"ldapClientPreferredServers" yaml:"ldapClientPreferredServers"`
	// Whether or not the LDAP traffic needs to be secured via TLS.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/5.15.0/docs/resources/netapp_account#ldap_over_tls_enabled NetappAccount#ldap_over_tls_enabled}
	LdapOverTlsEnabled interface{} `field:"optional" json:"ldapOverTlsEnabled" yaml:"ldapOverTlsEnabled"`
	// ldap_search_scope block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/5.15.0/docs/resources/netapp_account#ldap_search_scope NetappAccount#ldap_search_scope}
	LdapSearchScope *NetappAccountActiveDirectoryLdapSearchScope `field:"optional" json:"ldapSearchScope" yaml:"ldapSearchScope"`
	// Whether or not the LDAP traffic needs to be signed.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/5.15.0/docs/resources/netapp_account#ldap_signing_enabled NetappAccount#ldap_signing_enabled}
	LdapSigningEnabled interface{} `field:"optional" json:"ldapSigningEnabled" yaml:"ldapSigningEnabled"`
	// If enabled, NFS client local users can also (in addition to LDAP users) access the NFS volumes.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/5.15.0/docs/resources/netapp_account#local_nfs_users_with_ldap_allowed NetappAccount#local_nfs_users_with_ldap_allowed}
	LocalNfsUsersWithLdapAllowed interface{} `field:"optional" json:"localNfsUsersWithLdapAllowed" yaml:"localNfsUsersWithLdapAllowed"`
	// The Organizational Unit (OU) within the Windows Active Directory where machines will be created.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/5.15.0/docs/resources/netapp_account#organizational_unit NetappAccount#organizational_unit}
	OrganizationalUnit *string `field:"optional" json:"organizationalUnit" yaml:"organizationalUnit"`
	// Domain users to be given SeSecurityPrivilege privilege.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/5.15.0/docs/resources/netapp_account#security_users NetappAccount#security_users}
	SecurityUsers *[]*string `field:"optional" json:"securityUsers" yaml:"securityUsers"`
	// When LDAP over SSL/TLS is enabled, the LDAP client is required to have base64 encoded Active Directory Certificate Service's self-signed root CA certificate, this optional parameter is used only for dual protocol with LDAP user-mapping volumes.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/5.15.0/docs/resources/netapp_account#server_root_ca_certificate NetappAccount#server_root_ca_certificate}
	ServerRootCaCertificate *string `field:"optional" json:"serverRootCaCertificate" yaml:"serverRootCaCertificate"`
	// The Active Directory site the service will limit Domain Controller discovery to.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/5.15.0/docs/resources/netapp_account#site_name NetappAccount#site_name}
	SiteName *string `field:"optional" json:"siteName" yaml:"siteName"`
}

