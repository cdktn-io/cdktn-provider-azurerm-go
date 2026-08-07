// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package cdnfrontdoorbatchruleset

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-provider-azurerm-go/azurerm/v17/jsii"

	"github.com/cdktn-io/cdktn-provider-azurerm-go/azurerm/v17/cdnfrontdoorbatchruleset/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type CdnFrontdoorBatchRuleSetRuleActionsOutputReference interface {
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
	InternalValue() *CdnFrontdoorBatchRuleSetRuleActions
	SetInternalValue(val *CdnFrontdoorBatchRuleSetRuleActions)
	ModifyRequestHeader() CdnFrontdoorBatchRuleSetRuleActionsModifyRequestHeaderList
	ModifyRequestHeaderInput() interface{}
	ModifyResponseHeader() CdnFrontdoorBatchRuleSetRuleActionsModifyResponseHeaderList
	ModifyResponseHeaderInput() interface{}
	RouteConfigurationOverride() CdnFrontdoorBatchRuleSetRuleActionsRouteConfigurationOverrideOutputReference
	RouteConfigurationOverrideInput() *CdnFrontdoorBatchRuleSetRuleActionsRouteConfigurationOverride
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktn.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktn.IInterpolatingParent)
	UrlRedirect() CdnFrontdoorBatchRuleSetRuleActionsUrlRedirectOutputReference
	UrlRedirectInput() *CdnFrontdoorBatchRuleSetRuleActionsUrlRedirect
	UrlRewrite() CdnFrontdoorBatchRuleSetRuleActionsUrlRewriteOutputReference
	UrlRewriteInput() *CdnFrontdoorBatchRuleSetRuleActionsUrlRewrite
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
	PutRouteConfigurationOverride(value *CdnFrontdoorBatchRuleSetRuleActionsRouteConfigurationOverride)
	PutUrlRedirect(value *CdnFrontdoorBatchRuleSetRuleActionsUrlRedirect)
	PutUrlRewrite(value *CdnFrontdoorBatchRuleSetRuleActionsUrlRewrite)
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

// The jsii proxy struct for CdnFrontdoorBatchRuleSetRuleActionsOutputReference
type jsiiProxy_CdnFrontdoorBatchRuleSetRuleActionsOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_CdnFrontdoorBatchRuleSetRuleActionsOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CdnFrontdoorBatchRuleSetRuleActionsOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CdnFrontdoorBatchRuleSetRuleActionsOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CdnFrontdoorBatchRuleSetRuleActionsOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CdnFrontdoorBatchRuleSetRuleActionsOutputReference) InternalValue() *CdnFrontdoorBatchRuleSetRuleActions {
	var returns *CdnFrontdoorBatchRuleSetRuleActions
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CdnFrontdoorBatchRuleSetRuleActionsOutputReference) ModifyRequestHeader() CdnFrontdoorBatchRuleSetRuleActionsModifyRequestHeaderList {
	var returns CdnFrontdoorBatchRuleSetRuleActionsModifyRequestHeaderList
	_jsii_.Get(
		j,
		"modifyRequestHeader",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CdnFrontdoorBatchRuleSetRuleActionsOutputReference) ModifyRequestHeaderInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"modifyRequestHeaderInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CdnFrontdoorBatchRuleSetRuleActionsOutputReference) ModifyResponseHeader() CdnFrontdoorBatchRuleSetRuleActionsModifyResponseHeaderList {
	var returns CdnFrontdoorBatchRuleSetRuleActionsModifyResponseHeaderList
	_jsii_.Get(
		j,
		"modifyResponseHeader",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CdnFrontdoorBatchRuleSetRuleActionsOutputReference) ModifyResponseHeaderInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"modifyResponseHeaderInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CdnFrontdoorBatchRuleSetRuleActionsOutputReference) RouteConfigurationOverride() CdnFrontdoorBatchRuleSetRuleActionsRouteConfigurationOverrideOutputReference {
	var returns CdnFrontdoorBatchRuleSetRuleActionsRouteConfigurationOverrideOutputReference
	_jsii_.Get(
		j,
		"routeConfigurationOverride",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CdnFrontdoorBatchRuleSetRuleActionsOutputReference) RouteConfigurationOverrideInput() *CdnFrontdoorBatchRuleSetRuleActionsRouteConfigurationOverride {
	var returns *CdnFrontdoorBatchRuleSetRuleActionsRouteConfigurationOverride
	_jsii_.Get(
		j,
		"routeConfigurationOverrideInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CdnFrontdoorBatchRuleSetRuleActionsOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CdnFrontdoorBatchRuleSetRuleActionsOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CdnFrontdoorBatchRuleSetRuleActionsOutputReference) UrlRedirect() CdnFrontdoorBatchRuleSetRuleActionsUrlRedirectOutputReference {
	var returns CdnFrontdoorBatchRuleSetRuleActionsUrlRedirectOutputReference
	_jsii_.Get(
		j,
		"urlRedirect",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CdnFrontdoorBatchRuleSetRuleActionsOutputReference) UrlRedirectInput() *CdnFrontdoorBatchRuleSetRuleActionsUrlRedirect {
	var returns *CdnFrontdoorBatchRuleSetRuleActionsUrlRedirect
	_jsii_.Get(
		j,
		"urlRedirectInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CdnFrontdoorBatchRuleSetRuleActionsOutputReference) UrlRewrite() CdnFrontdoorBatchRuleSetRuleActionsUrlRewriteOutputReference {
	var returns CdnFrontdoorBatchRuleSetRuleActionsUrlRewriteOutputReference
	_jsii_.Get(
		j,
		"urlRewrite",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CdnFrontdoorBatchRuleSetRuleActionsOutputReference) UrlRewriteInput() *CdnFrontdoorBatchRuleSetRuleActionsUrlRewrite {
	var returns *CdnFrontdoorBatchRuleSetRuleActionsUrlRewrite
	_jsii_.Get(
		j,
		"urlRewriteInput",
		&returns,
	)
	return returns
}


func NewCdnFrontdoorBatchRuleSetRuleActionsOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) CdnFrontdoorBatchRuleSetRuleActionsOutputReference {
	_init_.Initialize()

	if err := validateNewCdnFrontdoorBatchRuleSetRuleActionsOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_CdnFrontdoorBatchRuleSetRuleActionsOutputReference{}

	_jsii_.Create(
		"@cdktn/provider-azurerm.cdnFrontdoorBatchRuleSet.CdnFrontdoorBatchRuleSetRuleActionsOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewCdnFrontdoorBatchRuleSetRuleActionsOutputReference_Override(c CdnFrontdoorBatchRuleSetRuleActionsOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/provider-azurerm.cdnFrontdoorBatchRuleSet.CdnFrontdoorBatchRuleSetRuleActionsOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		c,
	)
}

func (j *jsiiProxy_CdnFrontdoorBatchRuleSetRuleActionsOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_CdnFrontdoorBatchRuleSetRuleActionsOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_CdnFrontdoorBatchRuleSetRuleActionsOutputReference)SetInternalValue(val *CdnFrontdoorBatchRuleSetRuleActions) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_CdnFrontdoorBatchRuleSetRuleActionsOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_CdnFrontdoorBatchRuleSetRuleActionsOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (c *jsiiProxy_CdnFrontdoorBatchRuleSetRuleActionsOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CdnFrontdoorBatchRuleSetRuleActionsOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (c *jsiiProxy_CdnFrontdoorBatchRuleSetRuleActionsOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (c *jsiiProxy_CdnFrontdoorBatchRuleSetRuleActionsOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (c *jsiiProxy_CdnFrontdoorBatchRuleSetRuleActionsOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (c *jsiiProxy_CdnFrontdoorBatchRuleSetRuleActionsOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (c *jsiiProxy_CdnFrontdoorBatchRuleSetRuleActionsOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (c *jsiiProxy_CdnFrontdoorBatchRuleSetRuleActionsOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (c *jsiiProxy_CdnFrontdoorBatchRuleSetRuleActionsOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (c *jsiiProxy_CdnFrontdoorBatchRuleSetRuleActionsOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (c *jsiiProxy_CdnFrontdoorBatchRuleSetRuleActionsOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		c,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CdnFrontdoorBatchRuleSetRuleActionsOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (c *jsiiProxy_CdnFrontdoorBatchRuleSetRuleActionsOutputReference) PutModifyRequestHeader(value interface{}) {
	if err := c.validatePutModifyRequestHeaderParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"putModifyRequestHeader",
		[]interface{}{value},
	)
}

func (c *jsiiProxy_CdnFrontdoorBatchRuleSetRuleActionsOutputReference) PutModifyResponseHeader(value interface{}) {
	if err := c.validatePutModifyResponseHeaderParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"putModifyResponseHeader",
		[]interface{}{value},
	)
}

func (c *jsiiProxy_CdnFrontdoorBatchRuleSetRuleActionsOutputReference) PutRouteConfigurationOverride(value *CdnFrontdoorBatchRuleSetRuleActionsRouteConfigurationOverride) {
	if err := c.validatePutRouteConfigurationOverrideParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"putRouteConfigurationOverride",
		[]interface{}{value},
	)
}

func (c *jsiiProxy_CdnFrontdoorBatchRuleSetRuleActionsOutputReference) PutUrlRedirect(value *CdnFrontdoorBatchRuleSetRuleActionsUrlRedirect) {
	if err := c.validatePutUrlRedirectParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"putUrlRedirect",
		[]interface{}{value},
	)
}

func (c *jsiiProxy_CdnFrontdoorBatchRuleSetRuleActionsOutputReference) PutUrlRewrite(value *CdnFrontdoorBatchRuleSetRuleActionsUrlRewrite) {
	if err := c.validatePutUrlRewriteParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"putUrlRewrite",
		[]interface{}{value},
	)
}

func (c *jsiiProxy_CdnFrontdoorBatchRuleSetRuleActionsOutputReference) ResetModifyRequestHeader() {
	_jsii_.InvokeVoid(
		c,
		"resetModifyRequestHeader",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CdnFrontdoorBatchRuleSetRuleActionsOutputReference) ResetModifyResponseHeader() {
	_jsii_.InvokeVoid(
		c,
		"resetModifyResponseHeader",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CdnFrontdoorBatchRuleSetRuleActionsOutputReference) ResetRouteConfigurationOverride() {
	_jsii_.InvokeVoid(
		c,
		"resetRouteConfigurationOverride",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CdnFrontdoorBatchRuleSetRuleActionsOutputReference) ResetUrlRedirect() {
	_jsii_.InvokeVoid(
		c,
		"resetUrlRedirect",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CdnFrontdoorBatchRuleSetRuleActionsOutputReference) ResetUrlRewrite() {
	_jsii_.InvokeVoid(
		c,
		"resetUrlRewrite",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CdnFrontdoorBatchRuleSetRuleActionsOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
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

func (c *jsiiProxy_CdnFrontdoorBatchRuleSetRuleActionsOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

