// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package cdnfrontdoorbatchruleset

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-provider-azurerm-go/azurerm/v16/jsii"

	"github.com/cdktn-io/cdktn-provider-azurerm-go/azurerm/v16/cdnfrontdoorbatchruleset/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type CdnFrontdoorBatchRuleSetRuleActionsUrlRedirectOutputReference interface {
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
	InternalValue() *CdnFrontdoorBatchRuleSetRuleActionsUrlRedirect
	SetInternalValue(val *CdnFrontdoorBatchRuleSetRuleActionsUrlRedirect)
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

// The jsii proxy struct for CdnFrontdoorBatchRuleSetRuleActionsUrlRedirectOutputReference
type jsiiProxy_CdnFrontdoorBatchRuleSetRuleActionsUrlRedirectOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_CdnFrontdoorBatchRuleSetRuleActionsUrlRedirectOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CdnFrontdoorBatchRuleSetRuleActionsUrlRedirectOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CdnFrontdoorBatchRuleSetRuleActionsUrlRedirectOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CdnFrontdoorBatchRuleSetRuleActionsUrlRedirectOutputReference) DestinationFragment() *string {
	var returns *string
	_jsii_.Get(
		j,
		"destinationFragment",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CdnFrontdoorBatchRuleSetRuleActionsUrlRedirectOutputReference) DestinationFragmentInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"destinationFragmentInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CdnFrontdoorBatchRuleSetRuleActionsUrlRedirectOutputReference) DestinationHostName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"destinationHostName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CdnFrontdoorBatchRuleSetRuleActionsUrlRedirectOutputReference) DestinationHostNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"destinationHostNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CdnFrontdoorBatchRuleSetRuleActionsUrlRedirectOutputReference) DestinationPath() *string {
	var returns *string
	_jsii_.Get(
		j,
		"destinationPath",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CdnFrontdoorBatchRuleSetRuleActionsUrlRedirectOutputReference) DestinationPathInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"destinationPathInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CdnFrontdoorBatchRuleSetRuleActionsUrlRedirectOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CdnFrontdoorBatchRuleSetRuleActionsUrlRedirectOutputReference) InternalValue() *CdnFrontdoorBatchRuleSetRuleActionsUrlRedirect {
	var returns *CdnFrontdoorBatchRuleSetRuleActionsUrlRedirect
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CdnFrontdoorBatchRuleSetRuleActionsUrlRedirectOutputReference) QueryString() *string {
	var returns *string
	_jsii_.Get(
		j,
		"queryString",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CdnFrontdoorBatchRuleSetRuleActionsUrlRedirectOutputReference) QueryStringInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"queryStringInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CdnFrontdoorBatchRuleSetRuleActionsUrlRedirectOutputReference) RedirectProtocol() *string {
	var returns *string
	_jsii_.Get(
		j,
		"redirectProtocol",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CdnFrontdoorBatchRuleSetRuleActionsUrlRedirectOutputReference) RedirectProtocolInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"redirectProtocolInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CdnFrontdoorBatchRuleSetRuleActionsUrlRedirectOutputReference) RedirectType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"redirectType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CdnFrontdoorBatchRuleSetRuleActionsUrlRedirectOutputReference) RedirectTypeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"redirectTypeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CdnFrontdoorBatchRuleSetRuleActionsUrlRedirectOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CdnFrontdoorBatchRuleSetRuleActionsUrlRedirectOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewCdnFrontdoorBatchRuleSetRuleActionsUrlRedirectOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) CdnFrontdoorBatchRuleSetRuleActionsUrlRedirectOutputReference {
	_init_.Initialize()

	if err := validateNewCdnFrontdoorBatchRuleSetRuleActionsUrlRedirectOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_CdnFrontdoorBatchRuleSetRuleActionsUrlRedirectOutputReference{}

	_jsii_.Create(
		"@cdktn/provider-azurerm.cdnFrontdoorBatchRuleSet.CdnFrontdoorBatchRuleSetRuleActionsUrlRedirectOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewCdnFrontdoorBatchRuleSetRuleActionsUrlRedirectOutputReference_Override(c CdnFrontdoorBatchRuleSetRuleActionsUrlRedirectOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/provider-azurerm.cdnFrontdoorBatchRuleSet.CdnFrontdoorBatchRuleSetRuleActionsUrlRedirectOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		c,
	)
}

func (j *jsiiProxy_CdnFrontdoorBatchRuleSetRuleActionsUrlRedirectOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_CdnFrontdoorBatchRuleSetRuleActionsUrlRedirectOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_CdnFrontdoorBatchRuleSetRuleActionsUrlRedirectOutputReference)SetDestinationFragment(val *string) {
	if err := j.validateSetDestinationFragmentParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"destinationFragment",
		val,
	)
}

func (j *jsiiProxy_CdnFrontdoorBatchRuleSetRuleActionsUrlRedirectOutputReference)SetDestinationHostName(val *string) {
	if err := j.validateSetDestinationHostNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"destinationHostName",
		val,
	)
}

func (j *jsiiProxy_CdnFrontdoorBatchRuleSetRuleActionsUrlRedirectOutputReference)SetDestinationPath(val *string) {
	if err := j.validateSetDestinationPathParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"destinationPath",
		val,
	)
}

func (j *jsiiProxy_CdnFrontdoorBatchRuleSetRuleActionsUrlRedirectOutputReference)SetInternalValue(val *CdnFrontdoorBatchRuleSetRuleActionsUrlRedirect) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_CdnFrontdoorBatchRuleSetRuleActionsUrlRedirectOutputReference)SetQueryString(val *string) {
	if err := j.validateSetQueryStringParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"queryString",
		val,
	)
}

func (j *jsiiProxy_CdnFrontdoorBatchRuleSetRuleActionsUrlRedirectOutputReference)SetRedirectProtocol(val *string) {
	if err := j.validateSetRedirectProtocolParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"redirectProtocol",
		val,
	)
}

func (j *jsiiProxy_CdnFrontdoorBatchRuleSetRuleActionsUrlRedirectOutputReference)SetRedirectType(val *string) {
	if err := j.validateSetRedirectTypeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"redirectType",
		val,
	)
}

func (j *jsiiProxy_CdnFrontdoorBatchRuleSetRuleActionsUrlRedirectOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_CdnFrontdoorBatchRuleSetRuleActionsUrlRedirectOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (c *jsiiProxy_CdnFrontdoorBatchRuleSetRuleActionsUrlRedirectOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CdnFrontdoorBatchRuleSetRuleActionsUrlRedirectOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (c *jsiiProxy_CdnFrontdoorBatchRuleSetRuleActionsUrlRedirectOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (c *jsiiProxy_CdnFrontdoorBatchRuleSetRuleActionsUrlRedirectOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (c *jsiiProxy_CdnFrontdoorBatchRuleSetRuleActionsUrlRedirectOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (c *jsiiProxy_CdnFrontdoorBatchRuleSetRuleActionsUrlRedirectOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (c *jsiiProxy_CdnFrontdoorBatchRuleSetRuleActionsUrlRedirectOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (c *jsiiProxy_CdnFrontdoorBatchRuleSetRuleActionsUrlRedirectOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (c *jsiiProxy_CdnFrontdoorBatchRuleSetRuleActionsUrlRedirectOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (c *jsiiProxy_CdnFrontdoorBatchRuleSetRuleActionsUrlRedirectOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (c *jsiiProxy_CdnFrontdoorBatchRuleSetRuleActionsUrlRedirectOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		c,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CdnFrontdoorBatchRuleSetRuleActionsUrlRedirectOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (c *jsiiProxy_CdnFrontdoorBatchRuleSetRuleActionsUrlRedirectOutputReference) ResetDestinationFragment() {
	_jsii_.InvokeVoid(
		c,
		"resetDestinationFragment",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CdnFrontdoorBatchRuleSetRuleActionsUrlRedirectOutputReference) ResetDestinationHostName() {
	_jsii_.InvokeVoid(
		c,
		"resetDestinationHostName",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CdnFrontdoorBatchRuleSetRuleActionsUrlRedirectOutputReference) ResetDestinationPath() {
	_jsii_.InvokeVoid(
		c,
		"resetDestinationPath",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CdnFrontdoorBatchRuleSetRuleActionsUrlRedirectOutputReference) ResetQueryString() {
	_jsii_.InvokeVoid(
		c,
		"resetQueryString",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CdnFrontdoorBatchRuleSetRuleActionsUrlRedirectOutputReference) ResetRedirectProtocol() {
	_jsii_.InvokeVoid(
		c,
		"resetRedirectProtocol",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CdnFrontdoorBatchRuleSetRuleActionsUrlRedirectOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
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

func (c *jsiiProxy_CdnFrontdoorBatchRuleSetRuleActionsUrlRedirectOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

