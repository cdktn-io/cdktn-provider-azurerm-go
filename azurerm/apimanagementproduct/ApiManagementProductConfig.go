// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package apimanagementproduct

import (
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type ApiManagementProductConfig struct {
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
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/5.15.0/docs/resources/api_management_product#api_management_name ApiManagementProduct#api_management_name}.
	ApiManagementName *string `field:"required" json:"apiManagementName" yaml:"apiManagementName"`
	// The Display Name for this API Management Product.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/5.15.0/docs/resources/api_management_product#display_name ApiManagementProduct#display_name}
	DisplayName *string `field:"required" json:"displayName" yaml:"displayName"`
	// The Identifier for this Product, which must be unique within the API Management Service.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/5.15.0/docs/resources/api_management_product#product_id ApiManagementProduct#product_id}
	ProductId *string `field:"required" json:"productId" yaml:"productId"`
	// Is this Product Published?
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/5.15.0/docs/resources/api_management_product#published ApiManagementProduct#published}
	Published interface{} `field:"required" json:"published" yaml:"published"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/5.15.0/docs/resources/api_management_product#resource_group_name ApiManagementProduct#resource_group_name}.
	ResourceGroupName *string `field:"required" json:"resourceGroupName" yaml:"resourceGroupName"`
	// Do subscribers need to be approved prior to being able to use the Product?
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/5.15.0/docs/resources/api_management_product#approval_required ApiManagementProduct#approval_required}
	ApprovalRequired interface{} `field:"optional" json:"approvalRequired" yaml:"approvalRequired"`
	// A description of this Product, which may include HTML formatting tags.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/5.15.0/docs/resources/api_management_product#description ApiManagementProduct#description}
	Description *string `field:"optional" json:"description" yaml:"description"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/5.15.0/docs/resources/api_management_product#id ApiManagementProduct#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// Is a Subscription required to access API's included in this Product?
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/5.15.0/docs/resources/api_management_product#subscription_required ApiManagementProduct#subscription_required}
	SubscriptionRequired interface{} `field:"optional" json:"subscriptionRequired" yaml:"subscriptionRequired"`
	// The number of subscriptions a user can have to this Product at the same time.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/5.15.0/docs/resources/api_management_product#subscriptions_limit ApiManagementProduct#subscriptions_limit}
	SubscriptionsLimit *float64 `field:"optional" json:"subscriptionsLimit" yaml:"subscriptionsLimit"`
	// The Terms and Conditions for this Product, which must be accepted by Developers before they can begin the Subscription process.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/5.15.0/docs/resources/api_management_product#terms ApiManagementProduct#terms}
	Terms *string `field:"optional" json:"terms" yaml:"terms"`
	// timeouts block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/5.15.0/docs/resources/api_management_product#timeouts ApiManagementProduct#timeouts}
	Timeouts *ApiManagementProductTimeouts `field:"optional" json:"timeouts" yaml:"timeouts"`
}

