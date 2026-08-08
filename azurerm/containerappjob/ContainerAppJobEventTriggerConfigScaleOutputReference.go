// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package containerappjob

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-provider-azurerm-go/azurerm/v18/jsii"

	"github.com/cdktn-io/cdktn-provider-azurerm-go/azurerm/v18/containerappjob/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type ContainerAppJobEventTriggerConfigScaleOutputReference interface {
	cdktn.ComplexObject
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
	// The creation stack of this resolvable which will be appended to errors thrown during resolution.
	//
	// If this returns an empty array the stack will not be attached.
	// Experimental.
	CreationStack() *[]*string
	// Experimental.
	Fqn() *string
	InternalValue() interface{}
	SetInternalValue(val interface{})
	MaxExecutions() *float64
	SetMaxExecutions(val *float64)
	MaxExecutionsInput() *float64
	MinExecutions() *float64
	SetMinExecutions(val *float64)
	MinExecutionsInput() *float64
	PollingIntervalInSeconds() *float64
	SetPollingIntervalInSeconds(val *float64)
	PollingIntervalInSecondsInput() *float64
	Rules() ContainerAppJobEventTriggerConfigScaleRulesList
	RulesInput() interface{}
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
	PutRules(value interface{})
	ResetMaxExecutions()
	ResetMinExecutions()
	ResetPollingIntervalInSeconds()
	ResetRules()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktn.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for ContainerAppJobEventTriggerConfigScaleOutputReference
type jsiiProxy_ContainerAppJobEventTriggerConfigScaleOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_ContainerAppJobEventTriggerConfigScaleOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ContainerAppJobEventTriggerConfigScaleOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ContainerAppJobEventTriggerConfigScaleOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ContainerAppJobEventTriggerConfigScaleOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ContainerAppJobEventTriggerConfigScaleOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ContainerAppJobEventTriggerConfigScaleOutputReference) MaxExecutions() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"maxExecutions",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ContainerAppJobEventTriggerConfigScaleOutputReference) MaxExecutionsInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"maxExecutionsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ContainerAppJobEventTriggerConfigScaleOutputReference) MinExecutions() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"minExecutions",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ContainerAppJobEventTriggerConfigScaleOutputReference) MinExecutionsInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"minExecutionsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ContainerAppJobEventTriggerConfigScaleOutputReference) PollingIntervalInSeconds() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"pollingIntervalInSeconds",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ContainerAppJobEventTriggerConfigScaleOutputReference) PollingIntervalInSecondsInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"pollingIntervalInSecondsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ContainerAppJobEventTriggerConfigScaleOutputReference) Rules() ContainerAppJobEventTriggerConfigScaleRulesList {
	var returns ContainerAppJobEventTriggerConfigScaleRulesList
	_jsii_.Get(
		j,
		"rules",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ContainerAppJobEventTriggerConfigScaleOutputReference) RulesInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rulesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ContainerAppJobEventTriggerConfigScaleOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ContainerAppJobEventTriggerConfigScaleOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewContainerAppJobEventTriggerConfigScaleOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) ContainerAppJobEventTriggerConfigScaleOutputReference {
	_init_.Initialize()

	if err := validateNewContainerAppJobEventTriggerConfigScaleOutputReferenceParameters(terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet); err != nil {
		panic(err)
	}
	j := jsiiProxy_ContainerAppJobEventTriggerConfigScaleOutputReference{}

	_jsii_.Create(
		"@cdktn/provider-azurerm.containerAppJob.ContainerAppJobEventTriggerConfigScaleOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		&j,
	)

	return &j
}

func NewContainerAppJobEventTriggerConfigScaleOutputReference_Override(c ContainerAppJobEventTriggerConfigScaleOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/provider-azurerm.containerAppJob.ContainerAppJobEventTriggerConfigScaleOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		c,
	)
}

func (j *jsiiProxy_ContainerAppJobEventTriggerConfigScaleOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_ContainerAppJobEventTriggerConfigScaleOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_ContainerAppJobEventTriggerConfigScaleOutputReference)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_ContainerAppJobEventTriggerConfigScaleOutputReference)SetMaxExecutions(val *float64) {
	if err := j.validateSetMaxExecutionsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"maxExecutions",
		val,
	)
}

func (j *jsiiProxy_ContainerAppJobEventTriggerConfigScaleOutputReference)SetMinExecutions(val *float64) {
	if err := j.validateSetMinExecutionsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"minExecutions",
		val,
	)
}

func (j *jsiiProxy_ContainerAppJobEventTriggerConfigScaleOutputReference)SetPollingIntervalInSeconds(val *float64) {
	if err := j.validateSetPollingIntervalInSecondsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"pollingIntervalInSeconds",
		val,
	)
}

func (j *jsiiProxy_ContainerAppJobEventTriggerConfigScaleOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_ContainerAppJobEventTriggerConfigScaleOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (c *jsiiProxy_ContainerAppJobEventTriggerConfigScaleOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_ContainerAppJobEventTriggerConfigScaleOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (c *jsiiProxy_ContainerAppJobEventTriggerConfigScaleOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (c *jsiiProxy_ContainerAppJobEventTriggerConfigScaleOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (c *jsiiProxy_ContainerAppJobEventTriggerConfigScaleOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (c *jsiiProxy_ContainerAppJobEventTriggerConfigScaleOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (c *jsiiProxy_ContainerAppJobEventTriggerConfigScaleOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (c *jsiiProxy_ContainerAppJobEventTriggerConfigScaleOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (c *jsiiProxy_ContainerAppJobEventTriggerConfigScaleOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (c *jsiiProxy_ContainerAppJobEventTriggerConfigScaleOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (c *jsiiProxy_ContainerAppJobEventTriggerConfigScaleOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		c,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_ContainerAppJobEventTriggerConfigScaleOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (c *jsiiProxy_ContainerAppJobEventTriggerConfigScaleOutputReference) PutRules(value interface{}) {
	if err := c.validatePutRulesParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"putRules",
		[]interface{}{value},
	)
}

func (c *jsiiProxy_ContainerAppJobEventTriggerConfigScaleOutputReference) ResetMaxExecutions() {
	_jsii_.InvokeVoid(
		c,
		"resetMaxExecutions",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ContainerAppJobEventTriggerConfigScaleOutputReference) ResetMinExecutions() {
	_jsii_.InvokeVoid(
		c,
		"resetMinExecutions",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ContainerAppJobEventTriggerConfigScaleOutputReference) ResetPollingIntervalInSeconds() {
	_jsii_.InvokeVoid(
		c,
		"resetPollingIntervalInSeconds",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ContainerAppJobEventTriggerConfigScaleOutputReference) ResetRules() {
	_jsii_.InvokeVoid(
		c,
		"resetRules",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ContainerAppJobEventTriggerConfigScaleOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
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

func (c *jsiiProxy_ContainerAppJobEventTriggerConfigScaleOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

