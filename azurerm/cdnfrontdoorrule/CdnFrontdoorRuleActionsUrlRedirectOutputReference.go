// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package cdnfrontdoorrule

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-provider-azurerm-go/azurerm/v18/jsii"

	"github.com/cdktn-io/cdktn-provider-azurerm-go/azurerm/v18/cdnfrontdoorrule/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type CdnFrontdoorRuleActionsUrlRedirectOutputReference interface {
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
	DestinationFragment() *string
	SetDestinationFragment(val *string)
	DestinationFragmentInput() *string
	DestinationHostName() *string
	SetDestinationHostName(val *string)
	DestinationHostNameInput() *string
	DestinationPath() *string
	SetDestinationPath(val *string)
	DestinationPathInput() *string
	// Experimental.
	Fqn() *string
	InternalValue() *CdnFrontdoorRuleActionsUrlRedirect
	SetInternalValue(val *CdnFrontdoorRuleActionsUrlRedirect)
	QueryString() *string
	SetQueryString(val *string)
	QueryStringInput() *string
	RedirectProtocol() *string
	SetRedirectProtocol(val *string)
	RedirectProtocolInput() *string
	RedirectType() *string
	SetRedirectType(val *string)
	RedirectTypeInput() *string
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
	ResetDestinationFragment()
	ResetDestinationHostName()
	ResetDestinationPath()
	ResetQueryString()
	ResetRedirectProtocol()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktn.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for CdnFrontdoorRuleActionsUrlRedirectOutputReference
type jsiiProxy_CdnFrontdoorRuleActionsUrlRedirectOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_CdnFrontdoorRuleActionsUrlRedirectOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CdnFrontdoorRuleActionsUrlRedirectOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CdnFrontdoorRuleActionsUrlRedirectOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CdnFrontdoorRuleActionsUrlRedirectOutputReference) DestinationFragment() *string {
	var returns *string
	_jsii_.Get(
		j,
		"destinationFragment",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CdnFrontdoorRuleActionsUrlRedirectOutputReference) DestinationFragmentInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"destinationFragmentInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CdnFrontdoorRuleActionsUrlRedirectOutputReference) DestinationHostName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"destinationHostName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CdnFrontdoorRuleActionsUrlRedirectOutputReference) DestinationHostNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"destinationHostNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CdnFrontdoorRuleActionsUrlRedirectOutputReference) DestinationPath() *string {
	var returns *string
	_jsii_.Get(
		j,
		"destinationPath",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CdnFrontdoorRuleActionsUrlRedirectOutputReference) DestinationPathInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"destinationPathInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CdnFrontdoorRuleActionsUrlRedirectOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CdnFrontdoorRuleActionsUrlRedirectOutputReference) InternalValue() *CdnFrontdoorRuleActionsUrlRedirect {
	var returns *CdnFrontdoorRuleActionsUrlRedirect
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CdnFrontdoorRuleActionsUrlRedirectOutputReference) QueryString() *string {
	var returns *string
	_jsii_.Get(
		j,
		"queryString",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CdnFrontdoorRuleActionsUrlRedirectOutputReference) QueryStringInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"queryStringInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CdnFrontdoorRuleActionsUrlRedirectOutputReference) RedirectProtocol() *string {
	var returns *string
	_jsii_.Get(
		j,
		"redirectProtocol",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CdnFrontdoorRuleActionsUrlRedirectOutputReference) RedirectProtocolInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"redirectProtocolInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CdnFrontdoorRuleActionsUrlRedirectOutputReference) RedirectType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"redirectType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CdnFrontdoorRuleActionsUrlRedirectOutputReference) RedirectTypeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"redirectTypeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CdnFrontdoorRuleActionsUrlRedirectOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CdnFrontdoorRuleActionsUrlRedirectOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewCdnFrontdoorRuleActionsUrlRedirectOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) CdnFrontdoorRuleActionsUrlRedirectOutputReference {
	_init_.Initialize()

	if err := validateNewCdnFrontdoorRuleActionsUrlRedirectOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_CdnFrontdoorRuleActionsUrlRedirectOutputReference{}

	_jsii_.Create(
		"@cdktn/provider-azurerm.cdnFrontdoorRule.CdnFrontdoorRuleActionsUrlRedirectOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewCdnFrontdoorRuleActionsUrlRedirectOutputReference_Override(c CdnFrontdoorRuleActionsUrlRedirectOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/provider-azurerm.cdnFrontdoorRule.CdnFrontdoorRuleActionsUrlRedirectOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		c,
	)
}

func (j *jsiiProxy_CdnFrontdoorRuleActionsUrlRedirectOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_CdnFrontdoorRuleActionsUrlRedirectOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_CdnFrontdoorRuleActionsUrlRedirectOutputReference)SetDestinationFragment(val *string) {
	if err := j.validateSetDestinationFragmentParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"destinationFragment",
		val,
	)
}

func (j *jsiiProxy_CdnFrontdoorRuleActionsUrlRedirectOutputReference)SetDestinationHostName(val *string) {
	if err := j.validateSetDestinationHostNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"destinationHostName",
		val,
	)
}

func (j *jsiiProxy_CdnFrontdoorRuleActionsUrlRedirectOutputReference)SetDestinationPath(val *string) {
	if err := j.validateSetDestinationPathParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"destinationPath",
		val,
	)
}

func (j *jsiiProxy_CdnFrontdoorRuleActionsUrlRedirectOutputReference)SetInternalValue(val *CdnFrontdoorRuleActionsUrlRedirect) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_CdnFrontdoorRuleActionsUrlRedirectOutputReference)SetQueryString(val *string) {
	if err := j.validateSetQueryStringParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"queryString",
		val,
	)
}

func (j *jsiiProxy_CdnFrontdoorRuleActionsUrlRedirectOutputReference)SetRedirectProtocol(val *string) {
	if err := j.validateSetRedirectProtocolParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"redirectProtocol",
		val,
	)
}

func (j *jsiiProxy_CdnFrontdoorRuleActionsUrlRedirectOutputReference)SetRedirectType(val *string) {
	if err := j.validateSetRedirectTypeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"redirectType",
		val,
	)
}

func (j *jsiiProxy_CdnFrontdoorRuleActionsUrlRedirectOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_CdnFrontdoorRuleActionsUrlRedirectOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (c *jsiiProxy_CdnFrontdoorRuleActionsUrlRedirectOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CdnFrontdoorRuleActionsUrlRedirectOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (c *jsiiProxy_CdnFrontdoorRuleActionsUrlRedirectOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (c *jsiiProxy_CdnFrontdoorRuleActionsUrlRedirectOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (c *jsiiProxy_CdnFrontdoorRuleActionsUrlRedirectOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (c *jsiiProxy_CdnFrontdoorRuleActionsUrlRedirectOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (c *jsiiProxy_CdnFrontdoorRuleActionsUrlRedirectOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (c *jsiiProxy_CdnFrontdoorRuleActionsUrlRedirectOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (c *jsiiProxy_CdnFrontdoorRuleActionsUrlRedirectOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (c *jsiiProxy_CdnFrontdoorRuleActionsUrlRedirectOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (c *jsiiProxy_CdnFrontdoorRuleActionsUrlRedirectOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		c,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CdnFrontdoorRuleActionsUrlRedirectOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (c *jsiiProxy_CdnFrontdoorRuleActionsUrlRedirectOutputReference) ResetDestinationFragment() {
	_jsii_.InvokeVoid(
		c,
		"resetDestinationFragment",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CdnFrontdoorRuleActionsUrlRedirectOutputReference) ResetDestinationHostName() {
	_jsii_.InvokeVoid(
		c,
		"resetDestinationHostName",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CdnFrontdoorRuleActionsUrlRedirectOutputReference) ResetDestinationPath() {
	_jsii_.InvokeVoid(
		c,
		"resetDestinationPath",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CdnFrontdoorRuleActionsUrlRedirectOutputReference) ResetQueryString() {
	_jsii_.InvokeVoid(
		c,
		"resetQueryString",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CdnFrontdoorRuleActionsUrlRedirectOutputReference) ResetRedirectProtocol() {
	_jsii_.InvokeVoid(
		c,
		"resetRedirectProtocol",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CdnFrontdoorRuleActionsUrlRedirectOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
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

func (c *jsiiProxy_CdnFrontdoorRuleActionsUrlRedirectOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

