// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package webapplicationfirewallpolicy

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-provider-azurerm-go/azurerm/v18/jsii"

	"github.com/cdktn-io/cdktn-provider-azurerm-go/azurerm/v18/webapplicationfirewallpolicy/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type WebApplicationFirewallPolicyPolicySettingsLogScrubbingOutputReference interface {
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
	Enabled() interface{}
	SetEnabled(val interface{})
	EnabledInput() interface{}
	// Experimental.
	Fqn() *string
	InternalValue() *WebApplicationFirewallPolicyPolicySettingsLogScrubbing
	SetInternalValue(val *WebApplicationFirewallPolicyPolicySettingsLogScrubbing)
	Rule() WebApplicationFirewallPolicyPolicySettingsLogScrubbingRuleList
	RuleInput() interface{}
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
	PutRule(value interface{})
	ResetEnabled()
	ResetRule()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktn.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for WebApplicationFirewallPolicyPolicySettingsLogScrubbingOutputReference
type jsiiProxy_WebApplicationFirewallPolicyPolicySettingsLogScrubbingOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_WebApplicationFirewallPolicyPolicySettingsLogScrubbingOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WebApplicationFirewallPolicyPolicySettingsLogScrubbingOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WebApplicationFirewallPolicyPolicySettingsLogScrubbingOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WebApplicationFirewallPolicyPolicySettingsLogScrubbingOutputReference) Enabled() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"enabled",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WebApplicationFirewallPolicyPolicySettingsLogScrubbingOutputReference) EnabledInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"enabledInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WebApplicationFirewallPolicyPolicySettingsLogScrubbingOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WebApplicationFirewallPolicyPolicySettingsLogScrubbingOutputReference) InternalValue() *WebApplicationFirewallPolicyPolicySettingsLogScrubbing {
	var returns *WebApplicationFirewallPolicyPolicySettingsLogScrubbing
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WebApplicationFirewallPolicyPolicySettingsLogScrubbingOutputReference) Rule() WebApplicationFirewallPolicyPolicySettingsLogScrubbingRuleList {
	var returns WebApplicationFirewallPolicyPolicySettingsLogScrubbingRuleList
	_jsii_.Get(
		j,
		"rule",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WebApplicationFirewallPolicyPolicySettingsLogScrubbingOutputReference) RuleInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"ruleInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WebApplicationFirewallPolicyPolicySettingsLogScrubbingOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WebApplicationFirewallPolicyPolicySettingsLogScrubbingOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewWebApplicationFirewallPolicyPolicySettingsLogScrubbingOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) WebApplicationFirewallPolicyPolicySettingsLogScrubbingOutputReference {
	_init_.Initialize()

	if err := validateNewWebApplicationFirewallPolicyPolicySettingsLogScrubbingOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_WebApplicationFirewallPolicyPolicySettingsLogScrubbingOutputReference{}

	_jsii_.Create(
		"@cdktn/provider-azurerm.webApplicationFirewallPolicy.WebApplicationFirewallPolicyPolicySettingsLogScrubbingOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewWebApplicationFirewallPolicyPolicySettingsLogScrubbingOutputReference_Override(w WebApplicationFirewallPolicyPolicySettingsLogScrubbingOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/provider-azurerm.webApplicationFirewallPolicy.WebApplicationFirewallPolicyPolicySettingsLogScrubbingOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		w,
	)
}

func (j *jsiiProxy_WebApplicationFirewallPolicyPolicySettingsLogScrubbingOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_WebApplicationFirewallPolicyPolicySettingsLogScrubbingOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_WebApplicationFirewallPolicyPolicySettingsLogScrubbingOutputReference)SetEnabled(val interface{}) {
	if err := j.validateSetEnabledParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"enabled",
		val,
	)
}

func (j *jsiiProxy_WebApplicationFirewallPolicyPolicySettingsLogScrubbingOutputReference)SetInternalValue(val *WebApplicationFirewallPolicyPolicySettingsLogScrubbing) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_WebApplicationFirewallPolicyPolicySettingsLogScrubbingOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_WebApplicationFirewallPolicyPolicySettingsLogScrubbingOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (w *jsiiProxy_WebApplicationFirewallPolicyPolicySettingsLogScrubbingOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		w,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (w *jsiiProxy_WebApplicationFirewallPolicyPolicySettingsLogScrubbingOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	if err := w.validateGetAnyMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]interface{}

	_jsii_.Invoke(
		w,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (w *jsiiProxy_WebApplicationFirewallPolicyPolicySettingsLogScrubbingOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
	if err := w.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		w,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (w *jsiiProxy_WebApplicationFirewallPolicyPolicySettingsLogScrubbingOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	if err := w.validateGetBooleanMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*bool

	_jsii_.Invoke(
		w,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (w *jsiiProxy_WebApplicationFirewallPolicyPolicySettingsLogScrubbingOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	if err := w.validateGetListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*string

	_jsii_.Invoke(
		w,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (w *jsiiProxy_WebApplicationFirewallPolicyPolicySettingsLogScrubbingOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	if err := w.validateGetNumberAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *float64

	_jsii_.Invoke(
		w,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (w *jsiiProxy_WebApplicationFirewallPolicyPolicySettingsLogScrubbingOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	if err := w.validateGetNumberListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*float64

	_jsii_.Invoke(
		w,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (w *jsiiProxy_WebApplicationFirewallPolicyPolicySettingsLogScrubbingOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	if err := w.validateGetNumberMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*float64

	_jsii_.Invoke(
		w,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (w *jsiiProxy_WebApplicationFirewallPolicyPolicySettingsLogScrubbingOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	if err := w.validateGetStringAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		w,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (w *jsiiProxy_WebApplicationFirewallPolicyPolicySettingsLogScrubbingOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	if err := w.validateGetStringMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*string

	_jsii_.Invoke(
		w,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (w *jsiiProxy_WebApplicationFirewallPolicyPolicySettingsLogScrubbingOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		w,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (w *jsiiProxy_WebApplicationFirewallPolicyPolicySettingsLogScrubbingOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
	if err := w.validateInterpolationForAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		w,
		"interpolationForAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (w *jsiiProxy_WebApplicationFirewallPolicyPolicySettingsLogScrubbingOutputReference) PutRule(value interface{}) {
	if err := w.validatePutRuleParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		w,
		"putRule",
		[]interface{}{value},
	)
}

func (w *jsiiProxy_WebApplicationFirewallPolicyPolicySettingsLogScrubbingOutputReference) ResetEnabled() {
	_jsii_.InvokeVoid(
		w,
		"resetEnabled",
		nil, // no parameters
	)
}

func (w *jsiiProxy_WebApplicationFirewallPolicyPolicySettingsLogScrubbingOutputReference) ResetRule() {
	_jsii_.InvokeVoid(
		w,
		"resetRule",
		nil, // no parameters
	)
}

func (w *jsiiProxy_WebApplicationFirewallPolicyPolicySettingsLogScrubbingOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
	if err := w.validateResolveParameters(context); err != nil {
		panic(err)
	}
	var returns interface{}

	_jsii_.Invoke(
		w,
		"resolve",
		[]interface{}{context},
		&returns,
	)

	return returns
}

func (w *jsiiProxy_WebApplicationFirewallPolicyPolicySettingsLogScrubbingOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		w,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

