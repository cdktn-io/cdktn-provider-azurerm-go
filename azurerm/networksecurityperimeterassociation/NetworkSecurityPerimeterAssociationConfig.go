// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package networksecurityperimeterassociation

import (
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type NetworkSecurityPerimeterAssociationConfig struct {
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
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/5.0.1/docs/resources/network_security_perimeter_association#access_mode NetworkSecurityPerimeterAssociation#access_mode}.
	AccessMode *string `field:"required" json:"accessMode" yaml:"accessMode"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/5.0.1/docs/resources/network_security_perimeter_association#name NetworkSecurityPerimeterAssociation#name}.
	Name *string `field:"required" json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/5.0.1/docs/resources/network_security_perimeter_association#network_security_perimeter_profile_id NetworkSecurityPerimeterAssociation#network_security_perimeter_profile_id}.
	NetworkSecurityPerimeterProfileId *string `field:"required" json:"networkSecurityPerimeterProfileId" yaml:"networkSecurityPerimeterProfileId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/5.0.1/docs/resources/network_security_perimeter_association#resource_id NetworkSecurityPerimeterAssociation#resource_id}.
	ResourceId *string `field:"required" json:"resourceId" yaml:"resourceId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/5.0.1/docs/resources/network_security_perimeter_association#id NetworkSecurityPerimeterAssociation#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// timeouts block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/5.0.1/docs/resources/network_security_perimeter_association#timeouts NetworkSecurityPerimeterAssociation#timeouts}
	Timeouts *NetworkSecurityPerimeterAssociationTimeouts `field:"optional" json:"timeouts" yaml:"timeouts"`
}

