// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package networksecurityperimeteraccessrule

import (
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type NetworkSecurityPerimeterAccessRuleConfig struct {
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
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.75.0/docs/resources/network_security_perimeter_access_rule#direction NetworkSecurityPerimeterAccessRule#direction}.
	Direction *string `field:"required" json:"direction" yaml:"direction"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.75.0/docs/resources/network_security_perimeter_access_rule#name NetworkSecurityPerimeterAccessRule#name}.
	Name *string `field:"required" json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.75.0/docs/resources/network_security_perimeter_access_rule#network_security_perimeter_profile_id NetworkSecurityPerimeterAccessRule#network_security_perimeter_profile_id}.
	NetworkSecurityPerimeterProfileId *string `field:"required" json:"networkSecurityPerimeterProfileId" yaml:"networkSecurityPerimeterProfileId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.75.0/docs/resources/network_security_perimeter_access_rule#address_prefixes NetworkSecurityPerimeterAccessRule#address_prefixes}.
	AddressPrefixes *[]*string `field:"optional" json:"addressPrefixes" yaml:"addressPrefixes"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.75.0/docs/resources/network_security_perimeter_access_rule#fqdns NetworkSecurityPerimeterAccessRule#fqdns}.
	Fqdns *[]*string `field:"optional" json:"fqdns" yaml:"fqdns"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.75.0/docs/resources/network_security_perimeter_access_rule#id NetworkSecurityPerimeterAccessRule#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.75.0/docs/resources/network_security_perimeter_access_rule#service_tags NetworkSecurityPerimeterAccessRule#service_tags}.
	ServiceTags *[]*string `field:"optional" json:"serviceTags" yaml:"serviceTags"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.75.0/docs/resources/network_security_perimeter_access_rule#subscription_ids NetworkSecurityPerimeterAccessRule#subscription_ids}.
	SubscriptionIds *[]*string `field:"optional" json:"subscriptionIds" yaml:"subscriptionIds"`
	// timeouts block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.75.0/docs/resources/network_security_perimeter_access_rule#timeouts NetworkSecurityPerimeterAccessRule#timeouts}
	Timeouts *NetworkSecurityPerimeterAccessRuleTimeouts `field:"optional" json:"timeouts" yaml:"timeouts"`
}

