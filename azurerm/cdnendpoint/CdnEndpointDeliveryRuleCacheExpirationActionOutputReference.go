// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package cdnendpoint

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-provider-azurerm-go/azurerm/v18/jsii"

	"github.com/cdktn-io/cdktn-provider-azurerm-go/azurerm/v18/cdnendpoint/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type CdnEndpointDeliveryRuleCacheExpirationActionOutputReference interface {
	cdktn.ComplexObject
	Behavior() *string
	SetBehavior(val *string)
	BehaviorInput() *string
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
	Duration() *string
	SetDuration(val *string)
	DurationInput() *string
	// Experimental.
	Fqn() *string
	InternalValue() *CdnEndpointDeliveryRuleCacheExpirationAction
	SetInternalValue(val *CdnEndpointDeliveryRuleCacheExpirationAction)
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
	ResetDuration()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktn.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for CdnEndpointDeliveryRuleCacheExpirationActionOutputReference
type jsiiProxy_CdnEndpointDeliveryRuleCacheExpirationActionOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_CdnEndpointDeliveryRuleCacheExpirationActionOutputReference) Behavior() *string {
	var returns *string
	_jsii_.Get(
		j,
		"behavior",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CdnEndpointDeliveryRuleCacheExpirationActionOutputReference) BehaviorInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"behaviorInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CdnEndpointDeliveryRuleCacheExpirationActionOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CdnEndpointDeliveryRuleCacheExpirationActionOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CdnEndpointDeliveryRuleCacheExpirationActionOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CdnEndpointDeliveryRuleCacheExpirationActionOutputReference) Duration() *string {
	var returns *string
	_jsii_.Get(
		j,
		"duration",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CdnEndpointDeliveryRuleCacheExpirationActionOutputReference) DurationInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"durationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CdnEndpointDeliveryRuleCacheExpirationActionOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CdnEndpointDeliveryRuleCacheExpirationActionOutputReference) InternalValue() *CdnEndpointDeliveryRuleCacheExpirationAction {
	var returns *CdnEndpointDeliveryRuleCacheExpirationAction
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CdnEndpointDeliveryRuleCacheExpirationActionOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CdnEndpointDeliveryRuleCacheExpirationActionOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewCdnEndpointDeliveryRuleCacheExpirationActionOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) CdnEndpointDeliveryRuleCacheExpirationActionOutputReference {
	_init_.Initialize()

	if err := validateNewCdnEndpointDeliveryRuleCacheExpirationActionOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_CdnEndpointDeliveryRuleCacheExpirationActionOutputReference{}

	_jsii_.Create(
		"@cdktn/provider-azurerm.cdnEndpoint.CdnEndpointDeliveryRuleCacheExpirationActionOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewCdnEndpointDeliveryRuleCacheExpirationActionOutputReference_Override(c CdnEndpointDeliveryRuleCacheExpirationActionOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/provider-azurerm.cdnEndpoint.CdnEndpointDeliveryRuleCacheExpirationActionOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		c,
	)
}

func (j *jsiiProxy_CdnEndpointDeliveryRuleCacheExpirationActionOutputReference)SetBehavior(val *string) {
	if err := j.validateSetBehaviorParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"behavior",
		val,
	)
}

func (j *jsiiProxy_CdnEndpointDeliveryRuleCacheExpirationActionOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_CdnEndpointDeliveryRuleCacheExpirationActionOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_CdnEndpointDeliveryRuleCacheExpirationActionOutputReference)SetDuration(val *string) {
	if err := j.validateSetDurationParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"duration",
		val,
	)
}

func (j *jsiiProxy_CdnEndpointDeliveryRuleCacheExpirationActionOutputReference)SetInternalValue(val *CdnEndpointDeliveryRuleCacheExpirationAction) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_CdnEndpointDeliveryRuleCacheExpirationActionOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_CdnEndpointDeliveryRuleCacheExpirationActionOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (c *jsiiProxy_CdnEndpointDeliveryRuleCacheExpirationActionOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CdnEndpointDeliveryRuleCacheExpirationActionOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (c *jsiiProxy_CdnEndpointDeliveryRuleCacheExpirationActionOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (c *jsiiProxy_CdnEndpointDeliveryRuleCacheExpirationActionOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (c *jsiiProxy_CdnEndpointDeliveryRuleCacheExpirationActionOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (c *jsiiProxy_CdnEndpointDeliveryRuleCacheExpirationActionOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (c *jsiiProxy_CdnEndpointDeliveryRuleCacheExpirationActionOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (c *jsiiProxy_CdnEndpointDeliveryRuleCacheExpirationActionOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (c *jsiiProxy_CdnEndpointDeliveryRuleCacheExpirationActionOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (c *jsiiProxy_CdnEndpointDeliveryRuleCacheExpirationActionOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (c *jsiiProxy_CdnEndpointDeliveryRuleCacheExpirationActionOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		c,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CdnEndpointDeliveryRuleCacheExpirationActionOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (c *jsiiProxy_CdnEndpointDeliveryRuleCacheExpirationActionOutputReference) ResetDuration() {
	_jsii_.InvokeVoid(
		c,
		"resetDuration",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CdnEndpointDeliveryRuleCacheExpirationActionOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
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

func (c *jsiiProxy_CdnEndpointDeliveryRuleCacheExpirationActionOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

