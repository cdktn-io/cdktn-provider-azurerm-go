// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package manageddevopspool

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-provider-azurerm-go/azurerm/v16/jsii"

	"github.com/cdktn-io/cdktn-provider-azurerm-go/azurerm/v16/manageddevopspool/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type ManagedDevopsPoolStatelessAgentOutputReference interface {
	cdktn.ComplexObject
	AutomaticResourcePrediction() ManagedDevopsPoolStatelessAgentAutomaticResourcePredictionOutputReference
	AutomaticResourcePredictionInput() *ManagedDevopsPoolStatelessAgentAutomaticResourcePrediction
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
	InternalValue() *ManagedDevopsPoolStatelessAgent
	SetInternalValue(val *ManagedDevopsPoolStatelessAgent)
	ManualResourcePrediction() ManagedDevopsPoolStatelessAgentManualResourcePredictionOutputReference
	ManualResourcePredictionInput() *ManagedDevopsPoolStatelessAgentManualResourcePrediction
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
	PutAutomaticResourcePrediction(value *ManagedDevopsPoolStatelessAgentAutomaticResourcePrediction)
	PutManualResourcePrediction(value *ManagedDevopsPoolStatelessAgentManualResourcePrediction)
	ResetAutomaticResourcePrediction()
	ResetManualResourcePrediction()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktn.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for ManagedDevopsPoolStatelessAgentOutputReference
type jsiiProxy_ManagedDevopsPoolStatelessAgentOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_ManagedDevopsPoolStatelessAgentOutputReference) AutomaticResourcePrediction() ManagedDevopsPoolStatelessAgentAutomaticResourcePredictionOutputReference {
	var returns ManagedDevopsPoolStatelessAgentAutomaticResourcePredictionOutputReference
	_jsii_.Get(
		j,
		"automaticResourcePrediction",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ManagedDevopsPoolStatelessAgentOutputReference) AutomaticResourcePredictionInput() *ManagedDevopsPoolStatelessAgentAutomaticResourcePrediction {
	var returns *ManagedDevopsPoolStatelessAgentAutomaticResourcePrediction
	_jsii_.Get(
		j,
		"automaticResourcePredictionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ManagedDevopsPoolStatelessAgentOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ManagedDevopsPoolStatelessAgentOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ManagedDevopsPoolStatelessAgentOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ManagedDevopsPoolStatelessAgentOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ManagedDevopsPoolStatelessAgentOutputReference) InternalValue() *ManagedDevopsPoolStatelessAgent {
	var returns *ManagedDevopsPoolStatelessAgent
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ManagedDevopsPoolStatelessAgentOutputReference) ManualResourcePrediction() ManagedDevopsPoolStatelessAgentManualResourcePredictionOutputReference {
	var returns ManagedDevopsPoolStatelessAgentManualResourcePredictionOutputReference
	_jsii_.Get(
		j,
		"manualResourcePrediction",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ManagedDevopsPoolStatelessAgentOutputReference) ManualResourcePredictionInput() *ManagedDevopsPoolStatelessAgentManualResourcePrediction {
	var returns *ManagedDevopsPoolStatelessAgentManualResourcePrediction
	_jsii_.Get(
		j,
		"manualResourcePredictionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ManagedDevopsPoolStatelessAgentOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ManagedDevopsPoolStatelessAgentOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewManagedDevopsPoolStatelessAgentOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) ManagedDevopsPoolStatelessAgentOutputReference {
	_init_.Initialize()

	if err := validateNewManagedDevopsPoolStatelessAgentOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_ManagedDevopsPoolStatelessAgentOutputReference{}

	_jsii_.Create(
		"@cdktn/provider-azurerm.managedDevopsPool.ManagedDevopsPoolStatelessAgentOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewManagedDevopsPoolStatelessAgentOutputReference_Override(m ManagedDevopsPoolStatelessAgentOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/provider-azurerm.managedDevopsPool.ManagedDevopsPoolStatelessAgentOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		m,
	)
}

func (j *jsiiProxy_ManagedDevopsPoolStatelessAgentOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_ManagedDevopsPoolStatelessAgentOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_ManagedDevopsPoolStatelessAgentOutputReference)SetInternalValue(val *ManagedDevopsPoolStatelessAgent) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_ManagedDevopsPoolStatelessAgentOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_ManagedDevopsPoolStatelessAgentOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (m *jsiiProxy_ManagedDevopsPoolStatelessAgentOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		m,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (m *jsiiProxy_ManagedDevopsPoolStatelessAgentOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	if err := m.validateGetAnyMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]interface{}

	_jsii_.Invoke(
		m,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_ManagedDevopsPoolStatelessAgentOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
	if err := m.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		m,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_ManagedDevopsPoolStatelessAgentOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	if err := m.validateGetBooleanMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*bool

	_jsii_.Invoke(
		m,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_ManagedDevopsPoolStatelessAgentOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	if err := m.validateGetListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*string

	_jsii_.Invoke(
		m,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_ManagedDevopsPoolStatelessAgentOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	if err := m.validateGetNumberAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *float64

	_jsii_.Invoke(
		m,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_ManagedDevopsPoolStatelessAgentOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	if err := m.validateGetNumberListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*float64

	_jsii_.Invoke(
		m,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_ManagedDevopsPoolStatelessAgentOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	if err := m.validateGetNumberMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*float64

	_jsii_.Invoke(
		m,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_ManagedDevopsPoolStatelessAgentOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	if err := m.validateGetStringAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		m,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_ManagedDevopsPoolStatelessAgentOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	if err := m.validateGetStringMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*string

	_jsii_.Invoke(
		m,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_ManagedDevopsPoolStatelessAgentOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		m,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (m *jsiiProxy_ManagedDevopsPoolStatelessAgentOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
	if err := m.validateInterpolationForAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		m,
		"interpolationForAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_ManagedDevopsPoolStatelessAgentOutputReference) PutAutomaticResourcePrediction(value *ManagedDevopsPoolStatelessAgentAutomaticResourcePrediction) {
	if err := m.validatePutAutomaticResourcePredictionParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		m,
		"putAutomaticResourcePrediction",
		[]interface{}{value},
	)
}

func (m *jsiiProxy_ManagedDevopsPoolStatelessAgentOutputReference) PutManualResourcePrediction(value *ManagedDevopsPoolStatelessAgentManualResourcePrediction) {
	if err := m.validatePutManualResourcePredictionParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		m,
		"putManualResourcePrediction",
		[]interface{}{value},
	)
}

func (m *jsiiProxy_ManagedDevopsPoolStatelessAgentOutputReference) ResetAutomaticResourcePrediction() {
	_jsii_.InvokeVoid(
		m,
		"resetAutomaticResourcePrediction",
		nil, // no parameters
	)
}

func (m *jsiiProxy_ManagedDevopsPoolStatelessAgentOutputReference) ResetManualResourcePrediction() {
	_jsii_.InvokeVoid(
		m,
		"resetManualResourcePrediction",
		nil, // no parameters
	)
}

func (m *jsiiProxy_ManagedDevopsPoolStatelessAgentOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
	if err := m.validateResolveParameters(context); err != nil {
		panic(err)
	}
	var returns interface{}

	_jsii_.Invoke(
		m,
		"resolve",
		[]interface{}{context},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_ManagedDevopsPoolStatelessAgentOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		m,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

