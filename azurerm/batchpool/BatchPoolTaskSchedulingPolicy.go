// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package batchpool


type BatchPoolTaskSchedulingPolicy struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.79.0/docs/resources/batch_pool#node_fill_type BatchPool#node_fill_type}.
	NodeFillType *string `field:"optional" json:"nodeFillType" yaml:"nodeFillType"`
}

