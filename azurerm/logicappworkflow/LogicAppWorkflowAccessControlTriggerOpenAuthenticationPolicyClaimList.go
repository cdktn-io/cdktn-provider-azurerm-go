// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package logicappworkflow

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-provider-azurerm-go/azurerm/v18/jsii"

	"github.com/cdktn-io/cdktn-provider-azurerm-go/azurerm/v18/logicappworkflow/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type LogicAppWorkflowAccessControlTriggerOpenAuthenticationPolicyClaimList interface {
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
	Get(index *float64) LogicAppWorkflowAccessControlTriggerOpenAuthenticationPolicyClaimOutputReference
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktn.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for LogicAppWorkflowAccessControlTriggerOpenAuthenticationPolicyClaimList
type jsiiProxy_LogicAppWorkflowAccessControlTriggerOpenAuthenticationPolicyClaimList struct {
	internal.Type__cdktnComplexList
}

func (j *jsiiProxy_LogicAppWorkflowAccessControlTriggerOpenAuthenticationPolicyClaimList) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LogicAppWorkflowAccessControlTriggerOpenAuthenticationPolicyClaimList) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LogicAppWorkflowAccessControlTriggerOpenAuthenticationPolicyClaimList) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LogicAppWorkflowAccessControlTriggerOpenAuthenticationPolicyClaimList) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LogicAppWorkflowAccessControlTriggerOpenAuthenticationPolicyClaimList) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LogicAppWorkflowAccessControlTriggerOpenAuthenticationPolicyClaimList) WrapsSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"wrapsSet",
		&returns,
	)
	return returns
}


func NewLogicAppWorkflowAccessControlTriggerOpenAuthenticationPolicyClaimList(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) LogicAppWorkflowAccessControlTriggerOpenAuthenticationPolicyClaimList {
	_init_.Initialize()

	if err := validateNewLogicAppWorkflowAccessControlTriggerOpenAuthenticationPolicyClaimListParameters(terraformResource, terraformAttribute, wrapsSet); err != nil {
		panic(err)
	}
	j := jsiiProxy_LogicAppWorkflowAccessControlTriggerOpenAuthenticationPolicyClaimList{}

	_jsii_.Create(
		"@cdktn/provider-azurerm.logicAppWorkflow.LogicAppWorkflowAccessControlTriggerOpenAuthenticationPolicyClaimList",
		[]interface{}{terraformResource, terraformAttribute, wrapsSet},
		&j,
	)

	return &j
}

func NewLogicAppWorkflowAccessControlTriggerOpenAuthenticationPolicyClaimList_Override(l LogicAppWorkflowAccessControlTriggerOpenAuthenticationPolicyClaimList, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/provider-azurerm.logicAppWorkflow.LogicAppWorkflowAccessControlTriggerOpenAuthenticationPolicyClaimList",
		[]interface{}{terraformResource, terraformAttribute, wrapsSet},
		l,
	)
}

func (j *jsiiProxy_LogicAppWorkflowAccessControlTriggerOpenAuthenticationPolicyClaimList)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_LogicAppWorkflowAccessControlTriggerOpenAuthenticationPolicyClaimList)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_LogicAppWorkflowAccessControlTriggerOpenAuthenticationPolicyClaimList)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_LogicAppWorkflowAccessControlTriggerOpenAuthenticationPolicyClaimList)SetWrapsSet(val *bool) {
	if err := j.validateSetWrapsSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"wrapsSet",
		val,
	)
}

func (l *jsiiProxy_LogicAppWorkflowAccessControlTriggerOpenAuthenticationPolicyClaimList) AllWithMapKey(mapKeyAttributeName *string) cdktn.DynamicListTerraformIterator {
	if err := l.validateAllWithMapKeyParameters(mapKeyAttributeName); err != nil {
		panic(err)
	}
	var returns cdktn.DynamicListTerraformIterator

	_jsii_.Invoke(
		l,
		"allWithMapKey",
		[]interface{}{mapKeyAttributeName},
		&returns,
	)

	return returns
}

func (l *jsiiProxy_LogicAppWorkflowAccessControlTriggerOpenAuthenticationPolicyClaimList) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		l,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (l *jsiiProxy_LogicAppWorkflowAccessControlTriggerOpenAuthenticationPolicyClaimList) Get(index *float64) LogicAppWorkflowAccessControlTriggerOpenAuthenticationPolicyClaimOutputReference {
	if err := l.validateGetParameters(index); err != nil {
		panic(err)
	}
	var returns LogicAppWorkflowAccessControlTriggerOpenAuthenticationPolicyClaimOutputReference

	_jsii_.Invoke(
		l,
		"get",
		[]interface{}{index},
		&returns,
	)

	return returns
}

func (l *jsiiProxy_LogicAppWorkflowAccessControlTriggerOpenAuthenticationPolicyClaimList) Resolve(context cdktn.IResolveContext) interface{} {
	if err := l.validateResolveParameters(context); err != nil {
		panic(err)
	}
	var returns interface{}

	_jsii_.Invoke(
		l,
		"resolve",
		[]interface{}{context},
		&returns,
	)

	return returns
}

func (l *jsiiProxy_LogicAppWorkflowAccessControlTriggerOpenAuthenticationPolicyClaimList) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		l,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

