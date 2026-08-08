// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package manageddevopspool

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-provider-azurerm-go/azurerm/v18/jsii"

	"github.com/cdktn-io/cdktn-provider-azurerm-go/azurerm/v18/manageddevopspool/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type ManagedDevopsPoolStatefulAgentManualResourcePredictionOutputReference interface {
	cdktn.ComplexObject
	AllWeekSchedule() *float64
	SetAllWeekSchedule(val *float64)
	AllWeekScheduleInput() *float64
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
	FridaySchedule() ManagedDevopsPoolStatefulAgentManualResourcePredictionFridayScheduleList
	FridayScheduleInput() interface{}
	InternalValue() *ManagedDevopsPoolStatefulAgentManualResourcePrediction
	SetInternalValue(val *ManagedDevopsPoolStatefulAgentManualResourcePrediction)
	MondaySchedule() ManagedDevopsPoolStatefulAgentManualResourcePredictionMondayScheduleList
	MondayScheduleInput() interface{}
	SaturdaySchedule() ManagedDevopsPoolStatefulAgentManualResourcePredictionSaturdayScheduleList
	SaturdayScheduleInput() interface{}
	SundaySchedule() ManagedDevopsPoolStatefulAgentManualResourcePredictionSundayScheduleList
	SundayScheduleInput() interface{}
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktn.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktn.IInterpolatingParent)
	ThursdaySchedule() ManagedDevopsPoolStatefulAgentManualResourcePredictionThursdayScheduleList
	ThursdayScheduleInput() interface{}
	TimeZoneName() *string
	SetTimeZoneName(val *string)
	TimeZoneNameInput() *string
	TuesdaySchedule() ManagedDevopsPoolStatefulAgentManualResourcePredictionTuesdayScheduleList
	TuesdayScheduleInput() interface{}
	WednesdaySchedule() ManagedDevopsPoolStatefulAgentManualResourcePredictionWednesdayScheduleList
	WednesdayScheduleInput() interface{}
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
	PutFridaySchedule(value interface{})
	PutMondaySchedule(value interface{})
	PutSaturdaySchedule(value interface{})
	PutSundaySchedule(value interface{})
	PutThursdaySchedule(value interface{})
	PutTuesdaySchedule(value interface{})
	PutWednesdaySchedule(value interface{})
	ResetAllWeekSchedule()
	ResetFridaySchedule()
	ResetMondaySchedule()
	ResetSaturdaySchedule()
	ResetSundaySchedule()
	ResetThursdaySchedule()
	ResetTimeZoneName()
	ResetTuesdaySchedule()
	ResetWednesdaySchedule()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktn.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for ManagedDevopsPoolStatefulAgentManualResourcePredictionOutputReference
type jsiiProxy_ManagedDevopsPoolStatefulAgentManualResourcePredictionOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_ManagedDevopsPoolStatefulAgentManualResourcePredictionOutputReference) AllWeekSchedule() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"allWeekSchedule",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ManagedDevopsPoolStatefulAgentManualResourcePredictionOutputReference) AllWeekScheduleInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"allWeekScheduleInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ManagedDevopsPoolStatefulAgentManualResourcePredictionOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ManagedDevopsPoolStatefulAgentManualResourcePredictionOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ManagedDevopsPoolStatefulAgentManualResourcePredictionOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ManagedDevopsPoolStatefulAgentManualResourcePredictionOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ManagedDevopsPoolStatefulAgentManualResourcePredictionOutputReference) FridaySchedule() ManagedDevopsPoolStatefulAgentManualResourcePredictionFridayScheduleList {
	var returns ManagedDevopsPoolStatefulAgentManualResourcePredictionFridayScheduleList
	_jsii_.Get(
		j,
		"fridaySchedule",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ManagedDevopsPoolStatefulAgentManualResourcePredictionOutputReference) FridayScheduleInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"fridayScheduleInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ManagedDevopsPoolStatefulAgentManualResourcePredictionOutputReference) InternalValue() *ManagedDevopsPoolStatefulAgentManualResourcePrediction {
	var returns *ManagedDevopsPoolStatefulAgentManualResourcePrediction
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ManagedDevopsPoolStatefulAgentManualResourcePredictionOutputReference) MondaySchedule() ManagedDevopsPoolStatefulAgentManualResourcePredictionMondayScheduleList {
	var returns ManagedDevopsPoolStatefulAgentManualResourcePredictionMondayScheduleList
	_jsii_.Get(
		j,
		"mondaySchedule",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ManagedDevopsPoolStatefulAgentManualResourcePredictionOutputReference) MondayScheduleInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"mondayScheduleInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ManagedDevopsPoolStatefulAgentManualResourcePredictionOutputReference) SaturdaySchedule() ManagedDevopsPoolStatefulAgentManualResourcePredictionSaturdayScheduleList {
	var returns ManagedDevopsPoolStatefulAgentManualResourcePredictionSaturdayScheduleList
	_jsii_.Get(
		j,
		"saturdaySchedule",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ManagedDevopsPoolStatefulAgentManualResourcePredictionOutputReference) SaturdayScheduleInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"saturdayScheduleInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ManagedDevopsPoolStatefulAgentManualResourcePredictionOutputReference) SundaySchedule() ManagedDevopsPoolStatefulAgentManualResourcePredictionSundayScheduleList {
	var returns ManagedDevopsPoolStatefulAgentManualResourcePredictionSundayScheduleList
	_jsii_.Get(
		j,
		"sundaySchedule",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ManagedDevopsPoolStatefulAgentManualResourcePredictionOutputReference) SundayScheduleInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"sundayScheduleInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ManagedDevopsPoolStatefulAgentManualResourcePredictionOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ManagedDevopsPoolStatefulAgentManualResourcePredictionOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ManagedDevopsPoolStatefulAgentManualResourcePredictionOutputReference) ThursdaySchedule() ManagedDevopsPoolStatefulAgentManualResourcePredictionThursdayScheduleList {
	var returns ManagedDevopsPoolStatefulAgentManualResourcePredictionThursdayScheduleList
	_jsii_.Get(
		j,
		"thursdaySchedule",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ManagedDevopsPoolStatefulAgentManualResourcePredictionOutputReference) ThursdayScheduleInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"thursdayScheduleInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ManagedDevopsPoolStatefulAgentManualResourcePredictionOutputReference) TimeZoneName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"timeZoneName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ManagedDevopsPoolStatefulAgentManualResourcePredictionOutputReference) TimeZoneNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"timeZoneNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ManagedDevopsPoolStatefulAgentManualResourcePredictionOutputReference) TuesdaySchedule() ManagedDevopsPoolStatefulAgentManualResourcePredictionTuesdayScheduleList {
	var returns ManagedDevopsPoolStatefulAgentManualResourcePredictionTuesdayScheduleList
	_jsii_.Get(
		j,
		"tuesdaySchedule",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ManagedDevopsPoolStatefulAgentManualResourcePredictionOutputReference) TuesdayScheduleInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"tuesdayScheduleInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ManagedDevopsPoolStatefulAgentManualResourcePredictionOutputReference) WednesdaySchedule() ManagedDevopsPoolStatefulAgentManualResourcePredictionWednesdayScheduleList {
	var returns ManagedDevopsPoolStatefulAgentManualResourcePredictionWednesdayScheduleList
	_jsii_.Get(
		j,
		"wednesdaySchedule",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ManagedDevopsPoolStatefulAgentManualResourcePredictionOutputReference) WednesdayScheduleInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"wednesdayScheduleInput",
		&returns,
	)
	return returns
}


func NewManagedDevopsPoolStatefulAgentManualResourcePredictionOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) ManagedDevopsPoolStatefulAgentManualResourcePredictionOutputReference {
	_init_.Initialize()

	if err := validateNewManagedDevopsPoolStatefulAgentManualResourcePredictionOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_ManagedDevopsPoolStatefulAgentManualResourcePredictionOutputReference{}

	_jsii_.Create(
		"@cdktn/provider-azurerm.managedDevopsPool.ManagedDevopsPoolStatefulAgentManualResourcePredictionOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewManagedDevopsPoolStatefulAgentManualResourcePredictionOutputReference_Override(m ManagedDevopsPoolStatefulAgentManualResourcePredictionOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/provider-azurerm.managedDevopsPool.ManagedDevopsPoolStatefulAgentManualResourcePredictionOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		m,
	)
}

func (j *jsiiProxy_ManagedDevopsPoolStatefulAgentManualResourcePredictionOutputReference)SetAllWeekSchedule(val *float64) {
	if err := j.validateSetAllWeekScheduleParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"allWeekSchedule",
		val,
	)
}

func (j *jsiiProxy_ManagedDevopsPoolStatefulAgentManualResourcePredictionOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_ManagedDevopsPoolStatefulAgentManualResourcePredictionOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_ManagedDevopsPoolStatefulAgentManualResourcePredictionOutputReference)SetInternalValue(val *ManagedDevopsPoolStatefulAgentManualResourcePrediction) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_ManagedDevopsPoolStatefulAgentManualResourcePredictionOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_ManagedDevopsPoolStatefulAgentManualResourcePredictionOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_ManagedDevopsPoolStatefulAgentManualResourcePredictionOutputReference)SetTimeZoneName(val *string) {
	if err := j.validateSetTimeZoneNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"timeZoneName",
		val,
	)
}

func (m *jsiiProxy_ManagedDevopsPoolStatefulAgentManualResourcePredictionOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		m,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (m *jsiiProxy_ManagedDevopsPoolStatefulAgentManualResourcePredictionOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (m *jsiiProxy_ManagedDevopsPoolStatefulAgentManualResourcePredictionOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (m *jsiiProxy_ManagedDevopsPoolStatefulAgentManualResourcePredictionOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (m *jsiiProxy_ManagedDevopsPoolStatefulAgentManualResourcePredictionOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (m *jsiiProxy_ManagedDevopsPoolStatefulAgentManualResourcePredictionOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (m *jsiiProxy_ManagedDevopsPoolStatefulAgentManualResourcePredictionOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (m *jsiiProxy_ManagedDevopsPoolStatefulAgentManualResourcePredictionOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (m *jsiiProxy_ManagedDevopsPoolStatefulAgentManualResourcePredictionOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (m *jsiiProxy_ManagedDevopsPoolStatefulAgentManualResourcePredictionOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (m *jsiiProxy_ManagedDevopsPoolStatefulAgentManualResourcePredictionOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		m,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (m *jsiiProxy_ManagedDevopsPoolStatefulAgentManualResourcePredictionOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (m *jsiiProxy_ManagedDevopsPoolStatefulAgentManualResourcePredictionOutputReference) PutFridaySchedule(value interface{}) {
	if err := m.validatePutFridayScheduleParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		m,
		"putFridaySchedule",
		[]interface{}{value},
	)
}

func (m *jsiiProxy_ManagedDevopsPoolStatefulAgentManualResourcePredictionOutputReference) PutMondaySchedule(value interface{}) {
	if err := m.validatePutMondayScheduleParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		m,
		"putMondaySchedule",
		[]interface{}{value},
	)
}

func (m *jsiiProxy_ManagedDevopsPoolStatefulAgentManualResourcePredictionOutputReference) PutSaturdaySchedule(value interface{}) {
	if err := m.validatePutSaturdayScheduleParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		m,
		"putSaturdaySchedule",
		[]interface{}{value},
	)
}

func (m *jsiiProxy_ManagedDevopsPoolStatefulAgentManualResourcePredictionOutputReference) PutSundaySchedule(value interface{}) {
	if err := m.validatePutSundayScheduleParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		m,
		"putSundaySchedule",
		[]interface{}{value},
	)
}

func (m *jsiiProxy_ManagedDevopsPoolStatefulAgentManualResourcePredictionOutputReference) PutThursdaySchedule(value interface{}) {
	if err := m.validatePutThursdayScheduleParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		m,
		"putThursdaySchedule",
		[]interface{}{value},
	)
}

func (m *jsiiProxy_ManagedDevopsPoolStatefulAgentManualResourcePredictionOutputReference) PutTuesdaySchedule(value interface{}) {
	if err := m.validatePutTuesdayScheduleParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		m,
		"putTuesdaySchedule",
		[]interface{}{value},
	)
}

func (m *jsiiProxy_ManagedDevopsPoolStatefulAgentManualResourcePredictionOutputReference) PutWednesdaySchedule(value interface{}) {
	if err := m.validatePutWednesdayScheduleParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		m,
		"putWednesdaySchedule",
		[]interface{}{value},
	)
}

func (m *jsiiProxy_ManagedDevopsPoolStatefulAgentManualResourcePredictionOutputReference) ResetAllWeekSchedule() {
	_jsii_.InvokeVoid(
		m,
		"resetAllWeekSchedule",
		nil, // no parameters
	)
}

func (m *jsiiProxy_ManagedDevopsPoolStatefulAgentManualResourcePredictionOutputReference) ResetFridaySchedule() {
	_jsii_.InvokeVoid(
		m,
		"resetFridaySchedule",
		nil, // no parameters
	)
}

func (m *jsiiProxy_ManagedDevopsPoolStatefulAgentManualResourcePredictionOutputReference) ResetMondaySchedule() {
	_jsii_.InvokeVoid(
		m,
		"resetMondaySchedule",
		nil, // no parameters
	)
}

func (m *jsiiProxy_ManagedDevopsPoolStatefulAgentManualResourcePredictionOutputReference) ResetSaturdaySchedule() {
	_jsii_.InvokeVoid(
		m,
		"resetSaturdaySchedule",
		nil, // no parameters
	)
}

func (m *jsiiProxy_ManagedDevopsPoolStatefulAgentManualResourcePredictionOutputReference) ResetSundaySchedule() {
	_jsii_.InvokeVoid(
		m,
		"resetSundaySchedule",
		nil, // no parameters
	)
}

func (m *jsiiProxy_ManagedDevopsPoolStatefulAgentManualResourcePredictionOutputReference) ResetThursdaySchedule() {
	_jsii_.InvokeVoid(
		m,
		"resetThursdaySchedule",
		nil, // no parameters
	)
}

func (m *jsiiProxy_ManagedDevopsPoolStatefulAgentManualResourcePredictionOutputReference) ResetTimeZoneName() {
	_jsii_.InvokeVoid(
		m,
		"resetTimeZoneName",
		nil, // no parameters
	)
}

func (m *jsiiProxy_ManagedDevopsPoolStatefulAgentManualResourcePredictionOutputReference) ResetTuesdaySchedule() {
	_jsii_.InvokeVoid(
		m,
		"resetTuesdaySchedule",
		nil, // no parameters
	)
}

func (m *jsiiProxy_ManagedDevopsPoolStatefulAgentManualResourcePredictionOutputReference) ResetWednesdaySchedule() {
	_jsii_.InvokeVoid(
		m,
		"resetWednesdaySchedule",
		nil, // no parameters
	)
}

func (m *jsiiProxy_ManagedDevopsPoolStatefulAgentManualResourcePredictionOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
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

func (m *jsiiProxy_ManagedDevopsPoolStatefulAgentManualResourcePredictionOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		m,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

