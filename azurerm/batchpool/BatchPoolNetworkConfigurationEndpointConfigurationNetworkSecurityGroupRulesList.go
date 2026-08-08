// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package batchpool

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-provider-azurerm-go/azurerm/v18/jsii"

	"github.com/cdktn-io/cdktn-provider-azurerm-go/azurerm/v18/batchpool/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type BatchPoolNetworkConfigurationEndpointConfigurationNetworkSecurityGroupRulesList interface {
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
	Get(index *float64) BatchPoolNetworkConfigurationEndpointConfigurationNetworkSecurityGroupRulesOutputReference
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktn.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for BatchPoolNetworkConfigurationEndpointConfigurationNetworkSecurityGroupRulesList
type jsiiProxy_BatchPoolNetworkConfigurationEndpointConfigurationNetworkSecurityGroupRulesList struct {
	internal.Type__cdktnComplexList
}

func (j *jsiiProxy_BatchPoolNetworkConfigurationEndpointConfigurationNetworkSecurityGroupRulesList) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BatchPoolNetworkConfigurationEndpointConfigurationNetworkSecurityGroupRulesList) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BatchPoolNetworkConfigurationEndpointConfigurationNetworkSecurityGroupRulesList) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BatchPoolNetworkConfigurationEndpointConfigurationNetworkSecurityGroupRulesList) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BatchPoolNetworkConfigurationEndpointConfigurationNetworkSecurityGroupRulesList) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BatchPoolNetworkConfigurationEndpointConfigurationNetworkSecurityGroupRulesList) WrapsSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"wrapsSet",
		&returns,
	)
	return returns
}


func NewBatchPoolNetworkConfigurationEndpointConfigurationNetworkSecurityGroupRulesList(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) BatchPoolNetworkConfigurationEndpointConfigurationNetworkSecurityGroupRulesList {
	_init_.Initialize()

	if err := validateNewBatchPoolNetworkConfigurationEndpointConfigurationNetworkSecurityGroupRulesListParameters(terraformResource, terraformAttribute, wrapsSet); err != nil {
		panic(err)
	}
	j := jsiiProxy_BatchPoolNetworkConfigurationEndpointConfigurationNetworkSecurityGroupRulesList{}

	_jsii_.Create(
		"@cdktn/provider-azurerm.batchPool.BatchPoolNetworkConfigurationEndpointConfigurationNetworkSecurityGroupRulesList",
		[]interface{}{terraformResource, terraformAttribute, wrapsSet},
		&j,
	)

	return &j
}

func NewBatchPoolNetworkConfigurationEndpointConfigurationNetworkSecurityGroupRulesList_Override(b BatchPoolNetworkConfigurationEndpointConfigurationNetworkSecurityGroupRulesList, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/provider-azurerm.batchPool.BatchPoolNetworkConfigurationEndpointConfigurationNetworkSecurityGroupRulesList",
		[]interface{}{terraformResource, terraformAttribute, wrapsSet},
		b,
	)
}

func (j *jsiiProxy_BatchPoolNetworkConfigurationEndpointConfigurationNetworkSecurityGroupRulesList)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_BatchPoolNetworkConfigurationEndpointConfigurationNetworkSecurityGroupRulesList)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_BatchPoolNetworkConfigurationEndpointConfigurationNetworkSecurityGroupRulesList)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_BatchPoolNetworkConfigurationEndpointConfigurationNetworkSecurityGroupRulesList)SetWrapsSet(val *bool) {
	if err := j.validateSetWrapsSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"wrapsSet",
		val,
	)
}

func (b *jsiiProxy_BatchPoolNetworkConfigurationEndpointConfigurationNetworkSecurityGroupRulesList) AllWithMapKey(mapKeyAttributeName *string) cdktn.DynamicListTerraformIterator {
	if err := b.validateAllWithMapKeyParameters(mapKeyAttributeName); err != nil {
		panic(err)
	}
	var returns cdktn.DynamicListTerraformIterator

	_jsii_.Invoke(
		b,
		"allWithMapKey",
		[]interface{}{mapKeyAttributeName},
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BatchPoolNetworkConfigurationEndpointConfigurationNetworkSecurityGroupRulesList) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		b,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BatchPoolNetworkConfigurationEndpointConfigurationNetworkSecurityGroupRulesList) Get(index *float64) BatchPoolNetworkConfigurationEndpointConfigurationNetworkSecurityGroupRulesOutputReference {
	if err := b.validateGetParameters(index); err != nil {
		panic(err)
	}
	var returns BatchPoolNetworkConfigurationEndpointConfigurationNetworkSecurityGroupRulesOutputReference

	_jsii_.Invoke(
		b,
		"get",
		[]interface{}{index},
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BatchPoolNetworkConfigurationEndpointConfigurationNetworkSecurityGroupRulesList) Resolve(context cdktn.IResolveContext) interface{} {
	if err := b.validateResolveParameters(context); err != nil {
		panic(err)
	}
	var returns interface{}

	_jsii_.Invoke(
		b,
		"resolve",
		[]interface{}{context},
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BatchPoolNetworkConfigurationEndpointConfigurationNetworkSecurityGroupRulesList) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		b,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

