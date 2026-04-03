// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package emailcommunicationservicedomain

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-provider-azurerm-go/azurerm/v15/jsii"

	"github.com/cdktn-io/cdktn-provider-azurerm-go/azurerm/v15/emailcommunicationservicedomain/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type EmailCommunicationServiceDomainVerificationRecordsSpfList interface {
	cdktn.ComplexList
	// The creation stack of this resolvable which will be appended to errors thrown during resolution.
	//
	// If this returns an empty array the stack will not be attached.
	// Experimental.
	CreationStack() *[]*string
	// Experimental.
	Fqn() *string
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
	Get(index *float64) EmailCommunicationServiceDomainVerificationRecordsSpfOutputReference
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktn.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for EmailCommunicationServiceDomainVerificationRecordsSpfList
type jsiiProxy_EmailCommunicationServiceDomainVerificationRecordsSpfList struct {
	internal.Type__cdktnComplexList
}

func (j *jsiiProxy_EmailCommunicationServiceDomainVerificationRecordsSpfList) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EmailCommunicationServiceDomainVerificationRecordsSpfList) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EmailCommunicationServiceDomainVerificationRecordsSpfList) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EmailCommunicationServiceDomainVerificationRecordsSpfList) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EmailCommunicationServiceDomainVerificationRecordsSpfList) WrapsSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"wrapsSet",
		&returns,
	)
	return returns
}


func NewEmailCommunicationServiceDomainVerificationRecordsSpfList(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) EmailCommunicationServiceDomainVerificationRecordsSpfList {
	_init_.Initialize()

	if err := validateNewEmailCommunicationServiceDomainVerificationRecordsSpfListParameters(terraformResource, terraformAttribute, wrapsSet); err != nil {
		panic(err)
	}
	j := jsiiProxy_EmailCommunicationServiceDomainVerificationRecordsSpfList{}

	_jsii_.Create(
		"@cdktn/provider-azurerm.emailCommunicationServiceDomain.EmailCommunicationServiceDomainVerificationRecordsSpfList",
		[]interface{}{terraformResource, terraformAttribute, wrapsSet},
		&j,
	)

	return &j
}

func NewEmailCommunicationServiceDomainVerificationRecordsSpfList_Override(e EmailCommunicationServiceDomainVerificationRecordsSpfList, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/provider-azurerm.emailCommunicationServiceDomain.EmailCommunicationServiceDomainVerificationRecordsSpfList",
		[]interface{}{terraformResource, terraformAttribute, wrapsSet},
		e,
	)
}

func (j *jsiiProxy_EmailCommunicationServiceDomainVerificationRecordsSpfList)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_EmailCommunicationServiceDomainVerificationRecordsSpfList)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_EmailCommunicationServiceDomainVerificationRecordsSpfList)SetWrapsSet(val *bool) {
	if err := j.validateSetWrapsSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"wrapsSet",
		val,
	)
}

func (e *jsiiProxy_EmailCommunicationServiceDomainVerificationRecordsSpfList) AllWithMapKey(mapKeyAttributeName *string) cdktn.DynamicListTerraformIterator {
	if err := e.validateAllWithMapKeyParameters(mapKeyAttributeName); err != nil {
		panic(err)
	}
	var returns cdktn.DynamicListTerraformIterator

	_jsii_.Invoke(
		e,
		"allWithMapKey",
		[]interface{}{mapKeyAttributeName},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_EmailCommunicationServiceDomainVerificationRecordsSpfList) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		e,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (e *jsiiProxy_EmailCommunicationServiceDomainVerificationRecordsSpfList) Get(index *float64) EmailCommunicationServiceDomainVerificationRecordsSpfOutputReference {
	if err := e.validateGetParameters(index); err != nil {
		panic(err)
	}
	var returns EmailCommunicationServiceDomainVerificationRecordsSpfOutputReference

	_jsii_.Invoke(
		e,
		"get",
		[]interface{}{index},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_EmailCommunicationServiceDomainVerificationRecordsSpfList) Resolve(context cdktn.IResolveContext) interface{} {
	if err := e.validateResolveParameters(context); err != nil {
		panic(err)
	}
	var returns interface{}

	_jsii_.Invoke(
		e,
		"resolve",
		[]interface{}{context},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_EmailCommunicationServiceDomainVerificationRecordsSpfList) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		e,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

