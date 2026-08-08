// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package cdnfrontdoorrule

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-provider-azurerm-go/azurerm/v18/jsii"

	"github.com/cdktn-io/cdktn-provider-azurerm-go/azurerm/v18/cdnfrontdoorrule/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type CdnFrontdoorRuleActionsOutputReference interface {
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
	InternalValue() *CdnFrontdoorRuleActions
	SetInternalValue(val *CdnFrontdoorRuleActions)
	ModifyRequestHeader() CdnFrontdoorRuleActionsModifyRequestHeaderList
	ModifyRequestHeaderInput() interface{}
	ModifyResponseHeader() CdnFrontdoorRuleActionsModifyResponseHeaderList
	ModifyResponseHeaderInput() interface{}
	RouteConfigurationOverride() CdnFrontdoorRuleActionsRouteConfigurationOverrideOutputReference
	RouteConfigurationOverrideInput() *CdnFrontdoorRuleActionsRouteConfigurationOverride
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktn.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktn.IInterpolatingParent)
	UrlRedirect() CdnFrontdoorRuleActionsUrlRedirectOutputReference
	UrlRedirectInput() *CdnFrontdoorRuleActionsUrlRedirect
	UrlRewrite() CdnFrontdoorRuleActionsUrlRewriteOutputReference
	UrlRewriteInput() *CdnFrontdoorRuleActionsUrlRewrite
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
	PutModifyRequestHeader(value interface{})
	PutModifyResponseHeader(value interface{})
	PutRouteConfigurationOverride(value *CdnFrontdoorRuleActionsRouteConfigurationOverride)
	PutUrlRedirect(value *CdnFrontdoorRuleActionsUrlRedirect)
	PutUrlRewrite(value *CdnFrontdoorRuleActionsUrlRewrite)
	ResetModifyRequestHeader()
	ResetModifyResponseHeader()
	ResetRouteConfigurationOverride()
	ResetUrlRedirect()
	ResetUrlRewrite()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktn.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for CdnFrontdoorRuleActionsOutputReference
type jsiiProxy_CdnFrontdoorRuleActionsOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_CdnFrontdoorRuleActionsOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CdnFrontdoorRuleActionsOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CdnFrontdoorRuleActionsOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CdnFrontdoorRuleActionsOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CdnFrontdoorRuleActionsOutputReference) InternalValue() *CdnFrontdoorRuleActions {
	var returns *CdnFrontdoorRuleActions
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CdnFrontdoorRuleActionsOutputReference) ModifyRequestHeader() CdnFrontdoorRuleActionsModifyRequestHeaderList {
	var returns CdnFrontdoorRuleActionsModifyRequestHeaderList
	_jsii_.Get(
		j,
		"modifyRequestHeader",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CdnFrontdoorRuleActionsOutputReference) ModifyRequestHeaderInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"modifyRequestHeaderInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CdnFrontdoorRuleActionsOutputReference) ModifyResponseHeader() CdnFrontdoorRuleActionsModifyResponseHeaderList {
	var returns CdnFrontdoorRuleActionsModifyResponseHeaderList
	_jsii_.Get(
		j,
		"modifyResponseHeader",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CdnFrontdoorRuleActionsOutputReference) ModifyResponseHeaderInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"modifyResponseHeaderInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CdnFrontdoorRuleActionsOutputReference) RouteConfigurationOverride() CdnFrontdoorRuleActionsRouteConfigurationOverrideOutputReference {
	var returns CdnFrontdoorRuleActionsRouteConfigurationOverrideOutputReference
	_jsii_.Get(
		j,
		"routeConfigurationOverride",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CdnFrontdoorRuleActionsOutputReference) RouteConfigurationOverrideInput() *CdnFrontdoorRuleActionsRouteConfigurationOverride {
	var returns *CdnFrontdoorRuleActionsRouteConfigurationOverride
	_jsii_.Get(
		j,
		"routeConfigurationOverrideInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CdnFrontdoorRuleActionsOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CdnFrontdoorRuleActionsOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CdnFrontdoorRuleActionsOutputReference) UrlRedirect() CdnFrontdoorRuleActionsUrlRedirectOutputReference {
	var returns CdnFrontdoorRuleActionsUrlRedirectOutputReference
	_jsii_.Get(
		j,
		"urlRedirect",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CdnFrontdoorRuleActionsOutputReference) UrlRedirectInput() *CdnFrontdoorRuleActionsUrlRedirect {
	var returns *CdnFrontdoorRuleActionsUrlRedirect
	_jsii_.Get(
		j,
		"urlRedirectInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CdnFrontdoorRuleActionsOutputReference) UrlRewrite() CdnFrontdoorRuleActionsUrlRewriteOutputReference {
	var returns CdnFrontdoorRuleActionsUrlRewriteOutputReference
	_jsii_.Get(
		j,
		"urlRewrite",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CdnFrontdoorRuleActionsOutputReference) UrlRewriteInput() *CdnFrontdoorRuleActionsUrlRewrite {
	var returns *CdnFrontdoorRuleActionsUrlRewrite
	_jsii_.Get(
		j,
		"urlRewriteInput",
		&returns,
	)
	return returns
}


func NewCdnFrontdoorRuleActionsOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) CdnFrontdoorRuleActionsOutputReference {
	_init_.Initialize()

	if err := validateNewCdnFrontdoorRuleActionsOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_CdnFrontdoorRuleActionsOutputReference{}

	_jsii_.Create(
		"@cdktn/provider-azurerm.cdnFrontdoorRule.CdnFrontdoorRuleActionsOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewCdnFrontdoorRuleActionsOutputReference_Override(c CdnFrontdoorRuleActionsOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/provider-azurerm.cdnFrontdoorRule.CdnFrontdoorRuleActionsOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		c,
	)
}

func (j *jsiiProxy_CdnFrontdoorRuleActionsOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_CdnFrontdoorRuleActionsOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_CdnFrontdoorRuleActionsOutputReference)SetInternalValue(val *CdnFrontdoorRuleActions) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_CdnFrontdoorRuleActionsOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_CdnFrontdoorRuleActionsOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (c *jsiiProxy_CdnFrontdoorRuleActionsOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CdnFrontdoorRuleActionsOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (c *jsiiProxy_CdnFrontdoorRuleActionsOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (c *jsiiProxy_CdnFrontdoorRuleActionsOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (c *jsiiProxy_CdnFrontdoorRuleActionsOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (c *jsiiProxy_CdnFrontdoorRuleActionsOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (c *jsiiProxy_CdnFrontdoorRuleActionsOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (c *jsiiProxy_CdnFrontdoorRuleActionsOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (c *jsiiProxy_CdnFrontdoorRuleActionsOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (c *jsiiProxy_CdnFrontdoorRuleActionsOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (c *jsiiProxy_CdnFrontdoorRuleActionsOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		c,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CdnFrontdoorRuleActionsOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (c *jsiiProxy_CdnFrontdoorRuleActionsOutputReference) PutModifyRequestHeader(value interface{}) {
	if err := c.validatePutModifyRequestHeaderParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"putModifyRequestHeader",
		[]interface{}{value},
	)
}

func (c *jsiiProxy_CdnFrontdoorRuleActionsOutputReference) PutModifyResponseHeader(value interface{}) {
	if err := c.validatePutModifyResponseHeaderParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"putModifyResponseHeader",
		[]interface{}{value},
	)
}

func (c *jsiiProxy_CdnFrontdoorRuleActionsOutputReference) PutRouteConfigurationOverride(value *CdnFrontdoorRuleActionsRouteConfigurationOverride) {
	if err := c.validatePutRouteConfigurationOverrideParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"putRouteConfigurationOverride",
		[]interface{}{value},
	)
}

func (c *jsiiProxy_CdnFrontdoorRuleActionsOutputReference) PutUrlRedirect(value *CdnFrontdoorRuleActionsUrlRedirect) {
	if err := c.validatePutUrlRedirectParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"putUrlRedirect",
		[]interface{}{value},
	)
}

func (c *jsiiProxy_CdnFrontdoorRuleActionsOutputReference) PutUrlRewrite(value *CdnFrontdoorRuleActionsUrlRewrite) {
	if err := c.validatePutUrlRewriteParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"putUrlRewrite",
		[]interface{}{value},
	)
}

func (c *jsiiProxy_CdnFrontdoorRuleActionsOutputReference) ResetModifyRequestHeader() {
	_jsii_.InvokeVoid(
		c,
		"resetModifyRequestHeader",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CdnFrontdoorRuleActionsOutputReference) ResetModifyResponseHeader() {
	_jsii_.InvokeVoid(
		c,
		"resetModifyResponseHeader",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CdnFrontdoorRuleActionsOutputReference) ResetRouteConfigurationOverride() {
	_jsii_.InvokeVoid(
		c,
		"resetRouteConfigurationOverride",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CdnFrontdoorRuleActionsOutputReference) ResetUrlRedirect() {
	_jsii_.InvokeVoid(
		c,
		"resetUrlRedirect",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CdnFrontdoorRuleActionsOutputReference) ResetUrlRewrite() {
	_jsii_.InvokeVoid(
		c,
		"resetUrlRewrite",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CdnFrontdoorRuleActionsOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
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

func (c *jsiiProxy_CdnFrontdoorRuleActionsOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

