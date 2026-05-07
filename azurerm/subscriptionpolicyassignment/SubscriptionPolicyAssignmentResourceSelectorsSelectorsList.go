// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package subscriptionpolicyassignment

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-provider-azurerm-go/azurerm/v16/jsii"

	"github.com/cdktn-io/cdktn-provider-azurerm-go/azurerm/v16/subscriptionpolicyassignment/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type SubscriptionPolicyAssignmentResourceSelectorsSelectorsList interface {
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
	Get(index *float64) SubscriptionPolicyAssignmentResourceSelectorsSelectorsOutputReference
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktn.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for SubscriptionPolicyAssignmentResourceSelectorsSelectorsList
type jsiiProxy_SubscriptionPolicyAssignmentResourceSelectorsSelectorsList struct {
	internal.Type__cdktnComplexList
}

func (j *jsiiProxy_SubscriptionPolicyAssignmentResourceSelectorsSelectorsList) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SubscriptionPolicyAssignmentResourceSelectorsSelectorsList) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SubscriptionPolicyAssignmentResourceSelectorsSelectorsList) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SubscriptionPolicyAssignmentResourceSelectorsSelectorsList) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SubscriptionPolicyAssignmentResourceSelectorsSelectorsList) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SubscriptionPolicyAssignmentResourceSelectorsSelectorsList) WrapsSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"wrapsSet",
		&returns,
	)
	return returns
}


func NewSubscriptionPolicyAssignmentResourceSelectorsSelectorsList(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) SubscriptionPolicyAssignmentResourceSelectorsSelectorsList {
	_init_.Initialize()

	if err := validateNewSubscriptionPolicyAssignmentResourceSelectorsSelectorsListParameters(terraformResource, terraformAttribute, wrapsSet); err != nil {
		panic(err)
	}
	j := jsiiProxy_SubscriptionPolicyAssignmentResourceSelectorsSelectorsList{}

	_jsii_.Create(
		"@cdktn/provider-azurerm.subscriptionPolicyAssignment.SubscriptionPolicyAssignmentResourceSelectorsSelectorsList",
		[]interface{}{terraformResource, terraformAttribute, wrapsSet},
		&j,
	)

	return &j
}

func NewSubscriptionPolicyAssignmentResourceSelectorsSelectorsList_Override(s SubscriptionPolicyAssignmentResourceSelectorsSelectorsList, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/provider-azurerm.subscriptionPolicyAssignment.SubscriptionPolicyAssignmentResourceSelectorsSelectorsList",
		[]interface{}{terraformResource, terraformAttribute, wrapsSet},
		s,
	)
}

func (j *jsiiProxy_SubscriptionPolicyAssignmentResourceSelectorsSelectorsList)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_SubscriptionPolicyAssignmentResourceSelectorsSelectorsList)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_SubscriptionPolicyAssignmentResourceSelectorsSelectorsList)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_SubscriptionPolicyAssignmentResourceSelectorsSelectorsList)SetWrapsSet(val *bool) {
	if err := j.validateSetWrapsSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"wrapsSet",
		val,
	)
}

func (s *jsiiProxy_SubscriptionPolicyAssignmentResourceSelectorsSelectorsList) AllWithMapKey(mapKeyAttributeName *string) cdktn.DynamicListTerraformIterator {
	if err := s.validateAllWithMapKeyParameters(mapKeyAttributeName); err != nil {
		panic(err)
	}
	var returns cdktn.DynamicListTerraformIterator

	_jsii_.Invoke(
		s,
		"allWithMapKey",
		[]interface{}{mapKeyAttributeName},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SubscriptionPolicyAssignmentResourceSelectorsSelectorsList) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		s,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SubscriptionPolicyAssignmentResourceSelectorsSelectorsList) Get(index *float64) SubscriptionPolicyAssignmentResourceSelectorsSelectorsOutputReference {
	if err := s.validateGetParameters(index); err != nil {
		panic(err)
	}
	var returns SubscriptionPolicyAssignmentResourceSelectorsSelectorsOutputReference

	_jsii_.Invoke(
		s,
		"get",
		[]interface{}{index},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SubscriptionPolicyAssignmentResourceSelectorsSelectorsList) Resolve(context cdktn.IResolveContext) interface{} {
	if err := s.validateResolveParameters(context); err != nil {
		panic(err)
	}
	var returns interface{}

	_jsii_.Invoke(
		s,
		"resolve",
		[]interface{}{context},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SubscriptionPolicyAssignmentResourceSelectorsSelectorsList) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		s,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

