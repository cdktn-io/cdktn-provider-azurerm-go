// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package sharedimageversion

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-provider-azurerm-go/azurerm/v18/jsii"

	"github.com/cdktn-io/cdktn-provider-azurerm-go/azurerm/v18/sharedimageversion/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type SharedImageVersionUefiAdditionalSignaturesDbxList interface {
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
	Get(index *float64) SharedImageVersionUefiAdditionalSignaturesDbxOutputReference
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktn.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for SharedImageVersionUefiAdditionalSignaturesDbxList
type jsiiProxy_SharedImageVersionUefiAdditionalSignaturesDbxList struct {
	internal.Type__cdktnComplexList
}

func (j *jsiiProxy_SharedImageVersionUefiAdditionalSignaturesDbxList) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SharedImageVersionUefiAdditionalSignaturesDbxList) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SharedImageVersionUefiAdditionalSignaturesDbxList) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SharedImageVersionUefiAdditionalSignaturesDbxList) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SharedImageVersionUefiAdditionalSignaturesDbxList) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SharedImageVersionUefiAdditionalSignaturesDbxList) WrapsSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"wrapsSet",
		&returns,
	)
	return returns
}


func NewSharedImageVersionUefiAdditionalSignaturesDbxList(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) SharedImageVersionUefiAdditionalSignaturesDbxList {
	_init_.Initialize()

	if err := validateNewSharedImageVersionUefiAdditionalSignaturesDbxListParameters(terraformResource, terraformAttribute, wrapsSet); err != nil {
		panic(err)
	}
	j := jsiiProxy_SharedImageVersionUefiAdditionalSignaturesDbxList{}

	_jsii_.Create(
		"@cdktn/provider-azurerm.sharedImageVersion.SharedImageVersionUefiAdditionalSignaturesDbxList",
		[]interface{}{terraformResource, terraformAttribute, wrapsSet},
		&j,
	)

	return &j
}

func NewSharedImageVersionUefiAdditionalSignaturesDbxList_Override(s SharedImageVersionUefiAdditionalSignaturesDbxList, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/provider-azurerm.sharedImageVersion.SharedImageVersionUefiAdditionalSignaturesDbxList",
		[]interface{}{terraformResource, terraformAttribute, wrapsSet},
		s,
	)
}

func (j *jsiiProxy_SharedImageVersionUefiAdditionalSignaturesDbxList)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_SharedImageVersionUefiAdditionalSignaturesDbxList)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_SharedImageVersionUefiAdditionalSignaturesDbxList)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_SharedImageVersionUefiAdditionalSignaturesDbxList)SetWrapsSet(val *bool) {
	if err := j.validateSetWrapsSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"wrapsSet",
		val,
	)
}

func (s *jsiiProxy_SharedImageVersionUefiAdditionalSignaturesDbxList) AllWithMapKey(mapKeyAttributeName *string) cdktn.DynamicListTerraformIterator {
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

func (s *jsiiProxy_SharedImageVersionUefiAdditionalSignaturesDbxList) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		s,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SharedImageVersionUefiAdditionalSignaturesDbxList) Get(index *float64) SharedImageVersionUefiAdditionalSignaturesDbxOutputReference {
	if err := s.validateGetParameters(index); err != nil {
		panic(err)
	}
	var returns SharedImageVersionUefiAdditionalSignaturesDbxOutputReference

	_jsii_.Invoke(
		s,
		"get",
		[]interface{}{index},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SharedImageVersionUefiAdditionalSignaturesDbxList) Resolve(context cdktn.IResolveContext) interface{} {
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

func (s *jsiiProxy_SharedImageVersionUefiAdditionalSignaturesDbxList) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		s,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

