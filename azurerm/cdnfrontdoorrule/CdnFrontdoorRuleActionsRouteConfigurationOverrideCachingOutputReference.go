// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package cdnfrontdoorrule

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-provider-azurerm-go/azurerm/v18/jsii"

	"github.com/cdktn-io/cdktn-provider-azurerm-go/azurerm/v18/cdnfrontdoorrule/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type CdnFrontdoorRuleActionsRouteConfigurationOverrideCachingOutputReference interface {
	cdktn.ComplexObject
	Behaviour() *string
	SetBehaviour(val *string)
	BehaviourInput() *string
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
	CompressionEnabled() interface{}
	SetCompressionEnabled(val interface{})
	CompressionEnabledInput() interface{}
	// The creation stack of this resolvable which will be appended to errors thrown during resolution.
	//
	// If this returns an empty array the stack will not be attached.
	// Experimental.
	CreationStack() *[]*string
	Duration() *string
	SetDuration(val *string)
	DurationInput() *string
	// Experimental.
	Fqn() *string
	InternalValue() *CdnFrontdoorRuleActionsRouteConfigurationOverrideCaching
	SetInternalValue(val *CdnFrontdoorRuleActionsRouteConfigurationOverrideCaching)
	QueryStringBehaviour() *string
	SetQueryStringBehaviour(val *string)
	QueryStringBehaviourInput() *string
	QueryStringParameters() *[]*string
	SetQueryStringParameters(val *[]*string)
	QueryStringParametersInput() *[]*string
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
	ResetCompressionEnabled()
	ResetDuration()
	ResetQueryStringBehaviour()
	ResetQueryStringParameters()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktn.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for CdnFrontdoorRuleActionsRouteConfigurationOverrideCachingOutputReference
type jsiiProxy_CdnFrontdoorRuleActionsRouteConfigurationOverrideCachingOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_CdnFrontdoorRuleActionsRouteConfigurationOverrideCachingOutputReference) Behaviour() *string {
	var returns *string
	_jsii_.Get(
		j,
		"behaviour",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CdnFrontdoorRuleActionsRouteConfigurationOverrideCachingOutputReference) BehaviourInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"behaviourInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CdnFrontdoorRuleActionsRouteConfigurationOverrideCachingOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CdnFrontdoorRuleActionsRouteConfigurationOverrideCachingOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CdnFrontdoorRuleActionsRouteConfigurationOverrideCachingOutputReference) CompressionEnabled() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"compressionEnabled",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CdnFrontdoorRuleActionsRouteConfigurationOverrideCachingOutputReference) CompressionEnabledInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"compressionEnabledInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CdnFrontdoorRuleActionsRouteConfigurationOverrideCachingOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CdnFrontdoorRuleActionsRouteConfigurationOverrideCachingOutputReference) Duration() *string {
	var returns *string
	_jsii_.Get(
		j,
		"duration",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CdnFrontdoorRuleActionsRouteConfigurationOverrideCachingOutputReference) DurationInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"durationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CdnFrontdoorRuleActionsRouteConfigurationOverrideCachingOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CdnFrontdoorRuleActionsRouteConfigurationOverrideCachingOutputReference) InternalValue() *CdnFrontdoorRuleActionsRouteConfigurationOverrideCaching {
	var returns *CdnFrontdoorRuleActionsRouteConfigurationOverrideCaching
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CdnFrontdoorRuleActionsRouteConfigurationOverrideCachingOutputReference) QueryStringBehaviour() *string {
	var returns *string
	_jsii_.Get(
		j,
		"queryStringBehaviour",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CdnFrontdoorRuleActionsRouteConfigurationOverrideCachingOutputReference) QueryStringBehaviourInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"queryStringBehaviourInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CdnFrontdoorRuleActionsRouteConfigurationOverrideCachingOutputReference) QueryStringParameters() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"queryStringParameters",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CdnFrontdoorRuleActionsRouteConfigurationOverrideCachingOutputReference) QueryStringParametersInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"queryStringParametersInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CdnFrontdoorRuleActionsRouteConfigurationOverrideCachingOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CdnFrontdoorRuleActionsRouteConfigurationOverrideCachingOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewCdnFrontdoorRuleActionsRouteConfigurationOverrideCachingOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) CdnFrontdoorRuleActionsRouteConfigurationOverrideCachingOutputReference {
	_init_.Initialize()

	if err := validateNewCdnFrontdoorRuleActionsRouteConfigurationOverrideCachingOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_CdnFrontdoorRuleActionsRouteConfigurationOverrideCachingOutputReference{}

	_jsii_.Create(
		"@cdktn/provider-azurerm.cdnFrontdoorRule.CdnFrontdoorRuleActionsRouteConfigurationOverrideCachingOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewCdnFrontdoorRuleActionsRouteConfigurationOverrideCachingOutputReference_Override(c CdnFrontdoorRuleActionsRouteConfigurationOverrideCachingOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/provider-azurerm.cdnFrontdoorRule.CdnFrontdoorRuleActionsRouteConfigurationOverrideCachingOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		c,
	)
}

func (j *jsiiProxy_CdnFrontdoorRuleActionsRouteConfigurationOverrideCachingOutputReference)SetBehaviour(val *string) {
	if err := j.validateSetBehaviourParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"behaviour",
		val,
	)
}

func (j *jsiiProxy_CdnFrontdoorRuleActionsRouteConfigurationOverrideCachingOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_CdnFrontdoorRuleActionsRouteConfigurationOverrideCachingOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_CdnFrontdoorRuleActionsRouteConfigurationOverrideCachingOutputReference)SetCompressionEnabled(val interface{}) {
	if err := j.validateSetCompressionEnabledParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"compressionEnabled",
		val,
	)
}

func (j *jsiiProxy_CdnFrontdoorRuleActionsRouteConfigurationOverrideCachingOutputReference)SetDuration(val *string) {
	if err := j.validateSetDurationParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"duration",
		val,
	)
}

func (j *jsiiProxy_CdnFrontdoorRuleActionsRouteConfigurationOverrideCachingOutputReference)SetInternalValue(val *CdnFrontdoorRuleActionsRouteConfigurationOverrideCaching) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_CdnFrontdoorRuleActionsRouteConfigurationOverrideCachingOutputReference)SetQueryStringBehaviour(val *string) {
	if err := j.validateSetQueryStringBehaviourParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"queryStringBehaviour",
		val,
	)
}

func (j *jsiiProxy_CdnFrontdoorRuleActionsRouteConfigurationOverrideCachingOutputReference)SetQueryStringParameters(val *[]*string) {
	if err := j.validateSetQueryStringParametersParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"queryStringParameters",
		val,
	)
}

func (j *jsiiProxy_CdnFrontdoorRuleActionsRouteConfigurationOverrideCachingOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_CdnFrontdoorRuleActionsRouteConfigurationOverrideCachingOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (c *jsiiProxy_CdnFrontdoorRuleActionsRouteConfigurationOverrideCachingOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CdnFrontdoorRuleActionsRouteConfigurationOverrideCachingOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (c *jsiiProxy_CdnFrontdoorRuleActionsRouteConfigurationOverrideCachingOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (c *jsiiProxy_CdnFrontdoorRuleActionsRouteConfigurationOverrideCachingOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (c *jsiiProxy_CdnFrontdoorRuleActionsRouteConfigurationOverrideCachingOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (c *jsiiProxy_CdnFrontdoorRuleActionsRouteConfigurationOverrideCachingOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (c *jsiiProxy_CdnFrontdoorRuleActionsRouteConfigurationOverrideCachingOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (c *jsiiProxy_CdnFrontdoorRuleActionsRouteConfigurationOverrideCachingOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (c *jsiiProxy_CdnFrontdoorRuleActionsRouteConfigurationOverrideCachingOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (c *jsiiProxy_CdnFrontdoorRuleActionsRouteConfigurationOverrideCachingOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (c *jsiiProxy_CdnFrontdoorRuleActionsRouteConfigurationOverrideCachingOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		c,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CdnFrontdoorRuleActionsRouteConfigurationOverrideCachingOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (c *jsiiProxy_CdnFrontdoorRuleActionsRouteConfigurationOverrideCachingOutputReference) ResetCompressionEnabled() {
	_jsii_.InvokeVoid(
		c,
		"resetCompressionEnabled",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CdnFrontdoorRuleActionsRouteConfigurationOverrideCachingOutputReference) ResetDuration() {
	_jsii_.InvokeVoid(
		c,
		"resetDuration",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CdnFrontdoorRuleActionsRouteConfigurationOverrideCachingOutputReference) ResetQueryStringBehaviour() {
	_jsii_.InvokeVoid(
		c,
		"resetQueryStringBehaviour",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CdnFrontdoorRuleActionsRouteConfigurationOverrideCachingOutputReference) ResetQueryStringParameters() {
	_jsii_.InvokeVoid(
		c,
		"resetQueryStringParameters",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CdnFrontdoorRuleActionsRouteConfigurationOverrideCachingOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
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

func (c *jsiiProxy_CdnFrontdoorRuleActionsRouteConfigurationOverrideCachingOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

