// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package cdnfrontdoorrule

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-provider-azurerm-go/azurerm/v18/jsii"

	"github.com/cdktn-io/cdktn-provider-azurerm-go/azurerm/v18/cdnfrontdoorrule/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type CdnFrontdoorRuleActionsRouteConfigurationOverrideOutputReference interface {
	cdktn.ComplexObject
	Caching() CdnFrontdoorRuleActionsRouteConfigurationOverrideCachingOutputReference
	CachingInput() *CdnFrontdoorRuleActionsRouteConfigurationOverrideCaching
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
	InternalValue() *CdnFrontdoorRuleActionsRouteConfigurationOverride
	SetInternalValue(val *CdnFrontdoorRuleActionsRouteConfigurationOverride)
	OriginGroup() CdnFrontdoorRuleActionsRouteConfigurationOverrideOriginGroupOutputReference
	OriginGroupInput() *CdnFrontdoorRuleActionsRouteConfigurationOverrideOriginGroup
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
	PutCaching(value *CdnFrontdoorRuleActionsRouteConfigurationOverrideCaching)
	PutOriginGroup(value *CdnFrontdoorRuleActionsRouteConfigurationOverrideOriginGroup)
	ResetOriginGroup()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktn.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for CdnFrontdoorRuleActionsRouteConfigurationOverrideOutputReference
type jsiiProxy_CdnFrontdoorRuleActionsRouteConfigurationOverrideOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_CdnFrontdoorRuleActionsRouteConfigurationOverrideOutputReference) Caching() CdnFrontdoorRuleActionsRouteConfigurationOverrideCachingOutputReference {
	var returns CdnFrontdoorRuleActionsRouteConfigurationOverrideCachingOutputReference
	_jsii_.Get(
		j,
		"caching",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CdnFrontdoorRuleActionsRouteConfigurationOverrideOutputReference) CachingInput() *CdnFrontdoorRuleActionsRouteConfigurationOverrideCaching {
	var returns *CdnFrontdoorRuleActionsRouteConfigurationOverrideCaching
	_jsii_.Get(
		j,
		"cachingInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CdnFrontdoorRuleActionsRouteConfigurationOverrideOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CdnFrontdoorRuleActionsRouteConfigurationOverrideOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CdnFrontdoorRuleActionsRouteConfigurationOverrideOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CdnFrontdoorRuleActionsRouteConfigurationOverrideOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CdnFrontdoorRuleActionsRouteConfigurationOverrideOutputReference) InternalValue() *CdnFrontdoorRuleActionsRouteConfigurationOverride {
	var returns *CdnFrontdoorRuleActionsRouteConfigurationOverride
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CdnFrontdoorRuleActionsRouteConfigurationOverrideOutputReference) OriginGroup() CdnFrontdoorRuleActionsRouteConfigurationOverrideOriginGroupOutputReference {
	var returns CdnFrontdoorRuleActionsRouteConfigurationOverrideOriginGroupOutputReference
	_jsii_.Get(
		j,
		"originGroup",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CdnFrontdoorRuleActionsRouteConfigurationOverrideOutputReference) OriginGroupInput() *CdnFrontdoorRuleActionsRouteConfigurationOverrideOriginGroup {
	var returns *CdnFrontdoorRuleActionsRouteConfigurationOverrideOriginGroup
	_jsii_.Get(
		j,
		"originGroupInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CdnFrontdoorRuleActionsRouteConfigurationOverrideOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CdnFrontdoorRuleActionsRouteConfigurationOverrideOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewCdnFrontdoorRuleActionsRouteConfigurationOverrideOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) CdnFrontdoorRuleActionsRouteConfigurationOverrideOutputReference {
	_init_.Initialize()

	if err := validateNewCdnFrontdoorRuleActionsRouteConfigurationOverrideOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_CdnFrontdoorRuleActionsRouteConfigurationOverrideOutputReference{}

	_jsii_.Create(
		"@cdktn/provider-azurerm.cdnFrontdoorRule.CdnFrontdoorRuleActionsRouteConfigurationOverrideOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewCdnFrontdoorRuleActionsRouteConfigurationOverrideOutputReference_Override(c CdnFrontdoorRuleActionsRouteConfigurationOverrideOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/provider-azurerm.cdnFrontdoorRule.CdnFrontdoorRuleActionsRouteConfigurationOverrideOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		c,
	)
}

func (j *jsiiProxy_CdnFrontdoorRuleActionsRouteConfigurationOverrideOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_CdnFrontdoorRuleActionsRouteConfigurationOverrideOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_CdnFrontdoorRuleActionsRouteConfigurationOverrideOutputReference)SetInternalValue(val *CdnFrontdoorRuleActionsRouteConfigurationOverride) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_CdnFrontdoorRuleActionsRouteConfigurationOverrideOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_CdnFrontdoorRuleActionsRouteConfigurationOverrideOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (c *jsiiProxy_CdnFrontdoorRuleActionsRouteConfigurationOverrideOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CdnFrontdoorRuleActionsRouteConfigurationOverrideOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (c *jsiiProxy_CdnFrontdoorRuleActionsRouteConfigurationOverrideOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (c *jsiiProxy_CdnFrontdoorRuleActionsRouteConfigurationOverrideOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (c *jsiiProxy_CdnFrontdoorRuleActionsRouteConfigurationOverrideOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (c *jsiiProxy_CdnFrontdoorRuleActionsRouteConfigurationOverrideOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (c *jsiiProxy_CdnFrontdoorRuleActionsRouteConfigurationOverrideOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (c *jsiiProxy_CdnFrontdoorRuleActionsRouteConfigurationOverrideOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (c *jsiiProxy_CdnFrontdoorRuleActionsRouteConfigurationOverrideOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (c *jsiiProxy_CdnFrontdoorRuleActionsRouteConfigurationOverrideOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (c *jsiiProxy_CdnFrontdoorRuleActionsRouteConfigurationOverrideOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		c,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CdnFrontdoorRuleActionsRouteConfigurationOverrideOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (c *jsiiProxy_CdnFrontdoorRuleActionsRouteConfigurationOverrideOutputReference) PutCaching(value *CdnFrontdoorRuleActionsRouteConfigurationOverrideCaching) {
	if err := c.validatePutCachingParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"putCaching",
		[]interface{}{value},
	)
}

func (c *jsiiProxy_CdnFrontdoorRuleActionsRouteConfigurationOverrideOutputReference) PutOriginGroup(value *CdnFrontdoorRuleActionsRouteConfigurationOverrideOriginGroup) {
	if err := c.validatePutOriginGroupParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"putOriginGroup",
		[]interface{}{value},
	)
}

func (c *jsiiProxy_CdnFrontdoorRuleActionsRouteConfigurationOverrideOutputReference) ResetOriginGroup() {
	_jsii_.InvokeVoid(
		c,
		"resetOriginGroup",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CdnFrontdoorRuleActionsRouteConfigurationOverrideOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
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

func (c *jsiiProxy_CdnFrontdoorRuleActionsRouteConfigurationOverrideOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

