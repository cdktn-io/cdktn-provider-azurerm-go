// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package provider


type AzurermProviderFeaturesServicebus struct {
	// When enabled, the $Default rule is automatically deleted after creating a Service Bus subscription, preventing unfiltered message delivery.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.81.0/docs#auto_delete_subscription_default_rule AzurermProvider#auto_delete_subscription_default_rule}
	AutoDeleteSubscriptionDefaultRule interface{} `field:"optional" json:"autoDeleteSubscriptionDefaultRule" yaml:"autoDeleteSubscriptionDefaultRule"`
}

