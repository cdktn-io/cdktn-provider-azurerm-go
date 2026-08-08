// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package cdnfrontdoorsecuritypolicy

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-provider-azurerm-go/azurerm/v18/jsii"

	"github.com/cdktn-io/cdktn-provider-azurerm-go/azurerm/v18/cdnfrontdoorsecuritypolicy/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type CdnFrontdoorSecurityPolicySecurityPoliciesFirewallAssociationDomainList interface {
	cdktn.ComplexList
	// The creation stack of this resolvable which will be appended to errors thrown during resolution.
	//
	// If this returns an empty array the stack will not be attached.
	// Experimental.
	CreationStack() *[]*string
	// Experimental.
	Fqn() *string
	InternalValue() interface{}
	SetInternalValue(val interface{})
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktn.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktn.IInterpolatingParent)
	// Experimental.
	WrapsSet() *bool
	// Experimental.
	SetWrapsSet(val *bool)
	// Creating an iterator for this complex list.
	//
	// The list will be converted into a map with the mapKeyAttributeName as the key.
	// Experimental.
	AllWithMapKey(mapKeyAttributeName *string) cdktn.DynamicListTerraformIterator
	// Experimental.
	ComputeFqn() *string
	Get(index *float64) CdnFrontdoorSecurityPolicySecurityPoliciesFirewallAssociationDomainOutputReference
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktn.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for CdnFrontdoorSecurityPolicySecurityPoliciesFirewallAssociationDomainList
type jsiiProxy_CdnFrontdoorSecurityPolicySecurityPoliciesFirewallAssociationDomainList struct {
	internal.Type__cdktnComplexList
}

func (j *jsiiProxy_CdnFrontdoorSecurityPolicySecurityPoliciesFirewallAssociationDomainList) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CdnFrontdoorSecurityPolicySecurityPoliciesFirewallAssociationDomainList) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CdnFrontdoorSecurityPolicySecurityPoliciesFirewallAssociationDomainList) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CdnFrontdoorSecurityPolicySecurityPoliciesFirewallAssociationDomainList) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CdnFrontdoorSecurityPolicySecurityPoliciesFirewallAssociationDomainList) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CdnFrontdoorSecurityPolicySecurityPoliciesFirewallAssociationDomainList) WrapsSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"wrapsSet",
		&returns,
	)
	return returns
}


func NewCdnFrontdoorSecurityPolicySecurityPoliciesFirewallAssociationDomainList(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) CdnFrontdoorSecurityPolicySecurityPoliciesFirewallAssociationDomainList {
	_init_.Initialize()

	if err := validateNewCdnFrontdoorSecurityPolicySecurityPoliciesFirewallAssociationDomainListParameters(terraformResource, terraformAttribute, wrapsSet); err != nil {
		panic(err)
	}
	j := jsiiProxy_CdnFrontdoorSecurityPolicySecurityPoliciesFirewallAssociationDomainList{}

	_jsii_.Create(
		"@cdktn/provider-azurerm.cdnFrontdoorSecurityPolicy.CdnFrontdoorSecurityPolicySecurityPoliciesFirewallAssociationDomainList",
		[]interface{}{terraformResource, terraformAttribute, wrapsSet},
		&j,
	)

	return &j
}

func NewCdnFrontdoorSecurityPolicySecurityPoliciesFirewallAssociationDomainList_Override(c CdnFrontdoorSecurityPolicySecurityPoliciesFirewallAssociationDomainList, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/provider-azurerm.cdnFrontdoorSecurityPolicy.CdnFrontdoorSecurityPolicySecurityPoliciesFirewallAssociationDomainList",
		[]interface{}{terraformResource, terraformAttribute, wrapsSet},
		c,
	)
}

func (j *jsiiProxy_CdnFrontdoorSecurityPolicySecurityPoliciesFirewallAssociationDomainList)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_CdnFrontdoorSecurityPolicySecurityPoliciesFirewallAssociationDomainList)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_CdnFrontdoorSecurityPolicySecurityPoliciesFirewallAssociationDomainList)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_CdnFrontdoorSecurityPolicySecurityPoliciesFirewallAssociationDomainList)SetWrapsSet(val *bool) {
	if err := j.validateSetWrapsSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"wrapsSet",
		val,
	)
}

func (c *jsiiProxy_CdnFrontdoorSecurityPolicySecurityPoliciesFirewallAssociationDomainList) AllWithMapKey(mapKeyAttributeName *string) cdktn.DynamicListTerraformIterator {
	if err := c.validateAllWithMapKeyParameters(mapKeyAttributeName); err != nil {
		panic(err)
	}
	var returns cdktn.DynamicListTerraformIterator

	_jsii_.Invoke(
		c,
		"allWithMapKey",
		[]interface{}{mapKeyAttributeName},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CdnFrontdoorSecurityPolicySecurityPoliciesFirewallAssociationDomainList) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CdnFrontdoorSecurityPolicySecurityPoliciesFirewallAssociationDomainList) Get(index *float64) CdnFrontdoorSecurityPolicySecurityPoliciesFirewallAssociationDomainOutputReference {
	if err := c.validateGetParameters(index); err != nil {
		panic(err)
	}
	var returns CdnFrontdoorSecurityPolicySecurityPoliciesFirewallAssociationDomainOutputReference

	_jsii_.Invoke(
		c,
		"get",
		[]interface{}{index},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CdnFrontdoorSecurityPolicySecurityPoliciesFirewallAssociationDomainList) Resolve(context cdktn.IResolveContext) interface{} {
	if err := c.validateResolveParameters(context); err != nil {
		panic(err)
	}
	var returns interface{}

	_jsii_.Invoke(
		c,
		"resolve",
		[]interface{}{context},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CdnFrontdoorSecurityPolicySecurityPoliciesFirewallAssociationDomainList) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

