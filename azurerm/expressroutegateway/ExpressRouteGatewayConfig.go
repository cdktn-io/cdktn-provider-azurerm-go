// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package expressroutegateway

import (
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type ExpressRouteGatewayConfig struct {
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
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.62.1/docs/resources/express_route_gateway#location ExpressRouteGateway#location}.
	Location *string `field:"required" json:"location" yaml:"location"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.62.1/docs/resources/express_route_gateway#name ExpressRouteGateway#name}.
	Name *string `field:"required" json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.62.1/docs/resources/express_route_gateway#resource_group_name ExpressRouteGateway#resource_group_name}.
	ResourceGroupName *string `field:"required" json:"resourceGroupName" yaml:"resourceGroupName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.62.1/docs/resources/express_route_gateway#scale_units ExpressRouteGateway#scale_units}.
	ScaleUnits *float64 `field:"required" json:"scaleUnits" yaml:"scaleUnits"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.62.1/docs/resources/express_route_gateway#virtual_hub_id ExpressRouteGateway#virtual_hub_id}.
	VirtualHubId *string `field:"required" json:"virtualHubId" yaml:"virtualHubId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.62.1/docs/resources/express_route_gateway#allow_non_virtual_wan_traffic ExpressRouteGateway#allow_non_virtual_wan_traffic}.
	AllowNonVirtualWanTraffic interface{} `field:"optional" json:"allowNonVirtualWanTraffic" yaml:"allowNonVirtualWanTraffic"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.62.1/docs/resources/express_route_gateway#id ExpressRouteGateway#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.62.1/docs/resources/express_route_gateway#tags ExpressRouteGateway#tags}.
	Tags *map[string]*string `field:"optional" json:"tags" yaml:"tags"`
	// timeouts block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.62.1/docs/resources/express_route_gateway#timeouts ExpressRouteGateway#timeouts}
	Timeouts *ExpressRouteGatewayTimeouts `field:"optional" json:"timeouts" yaml:"timeouts"`
}

