// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package cdnfrontdoorbatchruleset

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-provider-azurerm-go/azurerm/v18/jsii"

	"github.com/cdktn-io/cdktn-provider-azurerm-go/azurerm/v18/cdnfrontdoorbatchruleset/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type CdnFrontdoorBatchRuleSetRuleOutputReference interface {
	cdktn.ComplexObject
	Actions() CdnFrontdoorBatchRuleSetRuleActionsOutputReference
	ActionsInput() *CdnFrontdoorBatchRuleSetRuleActions
	BehaviourOnMatch() *string
	SetBehaviourOnMatch(val *string)
	BehaviourOnMatchInput() *string
	// the index of the complex object in a list.
	// Experimental.
	ComplexObjectIndex() interface{}
	// Experimental.
	SetComplexObjectIndex(val interface{})
	// set to true if this item is from inside a set and needs tolist() for accessing it set to "0" for single list items.
	// Experimental.
	ComplexObjectIsFromSet() *bool
	// Experimental.
	SetComplexObjectIsFromSet(val *bool)
	Conditions() CdnFrontdoorBatchRuleSetRuleConditionsOutputReference
	ConditionsInput() *CdnFrontdoorBatchRuleSetRuleConditions
	// The creation stack of this resolvable which will be appended to errors thrown during resolution.
	//
	// If this returns an empty array the stack will not be attached.
	// Experimental.
	CreationStack() *[]*string
	// Experimental.
	Fqn() *string
	InternalValue() interface{}
	SetInternalValue(val interface{})
	Name() *string
	SetName(val *string)
	NameInput() *string
	Order() *float64
	SetOrder(val *float64)
	OrderInput() *float64
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktn.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktn.IInterpolatingParent)
	// Experimental.
	ComputeFqn() *string
	// Experimental.
	GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{}
	// Experimental.
	GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable
	// Experimental.
	GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool
	// Experimental.
	GetListAttribute(terraformAttribute *string) *[]*string
	// Experimental.
	GetNumberAttribute(terraformAttribute *string) *float64
	// Experimental.
	GetNumberListAttribute(terraformAttribute *string) *[]*float64
	// Experimental.
	GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64
	// Experimental.
	GetStringAttribute(terraformAttribute *string) *string
	// Experimental.
	GetStringMapAttribute(terraformAttribute *string) *map[string]*string
	// Experimental.
	InterpolationAsList() cdktn.IResolvable
	// Experimental.
	InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable
	PutActions(value *CdnFrontdoorBatchRuleSetRuleActions)
	PutConditions(value *CdnFrontdoorBatchRuleSetRuleConditions)
	ResetBehaviourOnMatch()
	ResetConditions()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktn.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for CdnFrontdoorBatchRuleSetRuleOutputReference
type jsiiProxy_CdnFrontdoorBatchRuleSetRuleOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_CdnFrontdoorBatchRuleSetRuleOutputReference) Actions() CdnFrontdoorBatchRuleSetRuleActionsOutputReference {
	var returns CdnFrontdoorBatchRuleSetRuleActionsOutputReference
	_jsii_.Get(
		j,
		"actions",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CdnFrontdoorBatchRuleSetRuleOutputReference) ActionsInput() *CdnFrontdoorBatchRuleSetRuleActions {
	var returns *CdnFrontdoorBatchRuleSetRuleActions
	_jsii_.Get(
		j,
		"actionsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CdnFrontdoorBatchRuleSetRuleOutputReference) BehaviourOnMatch() *string {
	var returns *string
	_jsii_.Get(
		j,
		"behaviourOnMatch",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CdnFrontdoorBatchRuleSetRuleOutputReference) BehaviourOnMatchInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"behaviourOnMatchInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CdnFrontdoorBatchRuleSetRuleOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CdnFrontdoorBatchRuleSetRuleOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CdnFrontdoorBatchRuleSetRuleOutputReference) Conditions() CdnFrontdoorBatchRuleSetRuleConditionsOutputReference {
	var returns CdnFrontdoorBatchRuleSetRuleConditionsOutputReference
	_jsii_.Get(
		j,
		"conditions",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CdnFrontdoorBatchRuleSetRuleOutputReference) ConditionsInput() *CdnFrontdoorBatchRuleSetRuleConditions {
	var returns *CdnFrontdoorBatchRuleSetRuleConditions
	_jsii_.Get(
		j,
		"conditionsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CdnFrontdoorBatchRuleSetRuleOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CdnFrontdoorBatchRuleSetRuleOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CdnFrontdoorBatchRuleSetRuleOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CdnFrontdoorBatchRuleSetRuleOutputReference) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CdnFrontdoorBatchRuleSetRuleOutputReference) NameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"nameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CdnFrontdoorBatchRuleSetRuleOutputReference) Order() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"order",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CdnFrontdoorBatchRuleSetRuleOutputReference) OrderInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"orderInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CdnFrontdoorBatchRuleSetRuleOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CdnFrontdoorBatchRuleSetRuleOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewCdnFrontdoorBatchRuleSetRuleOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) CdnFrontdoorBatchRuleSetRuleOutputReference {
	_init_.Initialize()

	if err := validateNewCdnFrontdoorBatchRuleSetRuleOutputReferenceParameters(terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet); err != nil {
		panic(err)
	}
	j := jsiiProxy_CdnFrontdoorBatchRuleSetRuleOutputReference{}

	_jsii_.Create(
		"@cdktn/provider-azurerm.cdnFrontdoorBatchRuleSet.CdnFrontdoorBatchRuleSetRuleOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		&j,
	)

	return &j
}

func NewCdnFrontdoorBatchRuleSetRuleOutputReference_Override(c CdnFrontdoorBatchRuleSetRuleOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/provider-azurerm.cdnFrontdoorBatchRuleSet.CdnFrontdoorBatchRuleSetRuleOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		c,
	)
}

func (j *jsiiProxy_CdnFrontdoorBatchRuleSetRuleOutputReference)SetBehaviourOnMatch(val *string) {
	if err := j.validateSetBehaviourOnMatchParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"behaviourOnMatch",
		val,
	)
}

func (j *jsiiProxy_CdnFrontdoorBatchRuleSetRuleOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_CdnFrontdoorBatchRuleSetRuleOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_CdnFrontdoorBatchRuleSetRuleOutputReference)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_CdnFrontdoorBatchRuleSetRuleOutputReference)SetName(val *string) {
	if err := j.validateSetNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"name",
		val,
	)
}

func (j *jsiiProxy_CdnFrontdoorBatchRuleSetRuleOutputReference)SetOrder(val *float64) {
	if err := j.validateSetOrderParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"order",
		val,
	)
}

func (j *jsiiProxy_CdnFrontdoorBatchRuleSetRuleOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_CdnFrontdoorBatchRuleSetRuleOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (c *jsiiProxy_CdnFrontdoorBatchRuleSetRuleOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CdnFrontdoorBatchRuleSetRuleOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	if err := c.validateGetAnyMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]interface{}

	_jsii_.Invoke(
		c,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CdnFrontdoorBatchRuleSetRuleOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
	if err := c.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		c,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CdnFrontdoorBatchRuleSetRuleOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	if err := c.validateGetBooleanMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*bool

	_jsii_.Invoke(
		c,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CdnFrontdoorBatchRuleSetRuleOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	if err := c.validateGetListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*string

	_jsii_.Invoke(
		c,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CdnFrontdoorBatchRuleSetRuleOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	if err := c.validateGetNumberAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *float64

	_jsii_.Invoke(
		c,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CdnFrontdoorBatchRuleSetRuleOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	if err := c.validateGetNumberListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*float64

	_jsii_.Invoke(
		c,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CdnFrontdoorBatchRuleSetRuleOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	if err := c.validateGetNumberMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*float64

	_jsii_.Invoke(
		c,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CdnFrontdoorBatchRuleSetRuleOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	if err := c.validateGetStringAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		c,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CdnFrontdoorBatchRuleSetRuleOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	if err := c.validateGetStringMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*string

	_jsii_.Invoke(
		c,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CdnFrontdoorBatchRuleSetRuleOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		c,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CdnFrontdoorBatchRuleSetRuleOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
	if err := c.validateInterpolationForAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		c,
		"interpolationForAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CdnFrontdoorBatchRuleSetRuleOutputReference) PutActions(value *CdnFrontdoorBatchRuleSetRuleActions) {
	if err := c.validatePutActionsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"putActions",
		[]interface{}{value},
	)
}

func (c *jsiiProxy_CdnFrontdoorBatchRuleSetRuleOutputReference) PutConditions(value *CdnFrontdoorBatchRuleSetRuleConditions) {
	if err := c.validatePutConditionsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"putConditions",
		[]interface{}{value},
	)
}

func (c *jsiiProxy_CdnFrontdoorBatchRuleSetRuleOutputReference) ResetBehaviourOnMatch() {
	_jsii_.InvokeVoid(
		c,
		"resetBehaviourOnMatch",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CdnFrontdoorBatchRuleSetRuleOutputReference) ResetConditions() {
	_jsii_.InvokeVoid(
		c,
		"resetConditions",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CdnFrontdoorBatchRuleSetRuleOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
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

func (c *jsiiProxy_CdnFrontdoorBatchRuleSetRuleOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

