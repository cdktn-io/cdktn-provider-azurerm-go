// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package recoveryservicesvault


type RecoveryServicesVaultMonitoring struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/5.2.0/docs/resources/recovery_services_vault#alerts_for_all_failover_issues_enabled RecoveryServicesVault#alerts_for_all_failover_issues_enabled}.
	AlertsForAllFailoverIssuesEnabled interface{} `field:"optional" json:"alertsForAllFailoverIssuesEnabled" yaml:"alertsForAllFailoverIssuesEnabled"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/5.2.0/docs/resources/recovery_services_vault#alerts_for_all_job_failures_enabled RecoveryServicesVault#alerts_for_all_job_failures_enabled}.
	AlertsForAllJobFailuresEnabled interface{} `field:"optional" json:"alertsForAllJobFailuresEnabled" yaml:"alertsForAllJobFailuresEnabled"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/5.2.0/docs/resources/recovery_services_vault#alerts_for_all_replication_issues_enabled RecoveryServicesVault#alerts_for_all_replication_issues_enabled}.
	AlertsForAllReplicationIssuesEnabled interface{} `field:"optional" json:"alertsForAllReplicationIssuesEnabled" yaml:"alertsForAllReplicationIssuesEnabled"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/5.2.0/docs/resources/recovery_services_vault#alerts_for_critical_operation_failures_enabled RecoveryServicesVault#alerts_for_critical_operation_failures_enabled}.
	AlertsForCriticalOperationFailuresEnabled interface{} `field:"optional" json:"alertsForCriticalOperationFailuresEnabled" yaml:"alertsForCriticalOperationFailuresEnabled"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/5.2.0/docs/resources/recovery_services_vault#email_notifications_for_site_recovery_enabled RecoveryServicesVault#email_notifications_for_site_recovery_enabled}.
	EmailNotificationsForSiteRecoveryEnabled interface{} `field:"optional" json:"emailNotificationsForSiteRecoveryEnabled" yaml:"emailNotificationsForSiteRecoveryEnabled"`
}

