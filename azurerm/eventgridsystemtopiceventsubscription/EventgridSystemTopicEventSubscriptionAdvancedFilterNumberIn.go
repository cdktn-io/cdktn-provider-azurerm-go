// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package eventgridsystemtopiceventsubscription


type EventgridSystemTopicEventSubscriptionAdvancedFilterNumberIn struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.80.0/docs/resources/eventgrid_system_topic_event_subscription#key EventgridSystemTopicEventSubscription#key}.
	Key *string `field:"required" json:"key" yaml:"key"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.80.0/docs/resources/eventgrid_system_topic_event_subscription#values EventgridSystemTopicEventSubscription#values}.
	Values *[]*float64 `field:"required" json:"values" yaml:"values"`
}

