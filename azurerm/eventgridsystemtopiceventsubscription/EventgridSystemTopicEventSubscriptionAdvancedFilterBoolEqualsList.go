// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package eventgridsystemtopiceventsubscription

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-provider-azurerm-go/azurerm/v16/jsii"

	"github.com/cdktn-io/cdktn-provider-azurerm-go/azurerm/v16/eventgridsystemtopiceventsubscription/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type EventgridSystemTopicEventSubscriptionAdvancedFilterBoolEqualsList interface {
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
	Get(index *float64) EventgridSystemTopicEventSubscriptionAdvancedFilterBoolEqualsOutputReference
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktn.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for EventgridSystemTopicEventSubscriptionAdvancedFilterBoolEqualsList
type jsiiProxy_EventgridSystemTopicEventSubscriptionAdvancedFilterBoolEqualsList struct {
	internal.Type__cdktnComplexList
}

func (j *jsiiProxy_EventgridSystemTopicEventSubscriptionAdvancedFilterBoolEqualsList) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EventgridSystemTopicEventSubscriptionAdvancedFilterBoolEqualsList) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EventgridSystemTopicEventSubscriptionAdvancedFilterBoolEqualsList) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EventgridSystemTopicEventSubscriptionAdvancedFilterBoolEqualsList) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EventgridSystemTopicEventSubscriptionAdvancedFilterBoolEqualsList) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EventgridSystemTopicEventSubscriptionAdvancedFilterBoolEqualsList) WrapsSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"wrapsSet",
		&returns,
	)
	return returns
}


func NewEventgridSystemTopicEventSubscriptionAdvancedFilterBoolEqualsList(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) EventgridSystemTopicEventSubscriptionAdvancedFilterBoolEqualsList {
	_init_.Initialize()

	if err := validateNewEventgridSystemTopicEventSubscriptionAdvancedFilterBoolEqualsListParameters(terraformResource, terraformAttribute, wrapsSet); err != nil {
		panic(err)
	}
	j := jsiiProxy_EventgridSystemTopicEventSubscriptionAdvancedFilterBoolEqualsList{}

	_jsii_.Create(
		"@cdktn/provider-azurerm.eventgridSystemTopicEventSubscription.EventgridSystemTopicEventSubscriptionAdvancedFilterBoolEqualsList",
		[]interface{}{terraformResource, terraformAttribute, wrapsSet},
		&j,
	)

	return &j
}

func NewEventgridSystemTopicEventSubscriptionAdvancedFilterBoolEqualsList_Override(e EventgridSystemTopicEventSubscriptionAdvancedFilterBoolEqualsList, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/provider-azurerm.eventgridSystemTopicEventSubscription.EventgridSystemTopicEventSubscriptionAdvancedFilterBoolEqualsList",
		[]interface{}{terraformResource, terraformAttribute, wrapsSet},
		e,
	)
}

func (j *jsiiProxy_EventgridSystemTopicEventSubscriptionAdvancedFilterBoolEqualsList)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_EventgridSystemTopicEventSubscriptionAdvancedFilterBoolEqualsList)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_EventgridSystemTopicEventSubscriptionAdvancedFilterBoolEqualsList)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_EventgridSystemTopicEventSubscriptionAdvancedFilterBoolEqualsList)SetWrapsSet(val *bool) {
	if err := j.validateSetWrapsSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"wrapsSet",
		val,
	)
}

func (e *jsiiProxy_EventgridSystemTopicEventSubscriptionAdvancedFilterBoolEqualsList) AllWithMapKey(mapKeyAttributeName *string) cdktn.DynamicListTerraformIterator {
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

func (e *jsiiProxy_EventgridSystemTopicEventSubscriptionAdvancedFilterBoolEqualsList) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		e,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (e *jsiiProxy_EventgridSystemTopicEventSubscriptionAdvancedFilterBoolEqualsList) Get(index *float64) EventgridSystemTopicEventSubscriptionAdvancedFilterBoolEqualsOutputReference {
	if err := e.validateGetParameters(index); err != nil {
		panic(err)
	}
	var returns EventgridSystemTopicEventSubscriptionAdvancedFilterBoolEqualsOutputReference

	_jsii_.Invoke(
		e,
		"get",
		[]interface{}{index},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_EventgridSystemTopicEventSubscriptionAdvancedFilterBoolEqualsList) Resolve(context cdktn.IResolveContext) interface{} {
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

func (e *jsiiProxy_EventgridSystemTopicEventSubscriptionAdvancedFilterBoolEqualsList) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		e,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

