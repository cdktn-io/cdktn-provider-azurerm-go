// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package dataazurermmanageddevopspool

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-provider-azurerm-go/azurerm/v17/jsii"

	"github.com/cdktn-io/cdktn-provider-azurerm-go/azurerm/v17/dataazurermmanageddevopspool/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type DataAzurermManagedDevopsPoolStatelessAgentManualResourcePredictionOutputReference interface {
	cdktn.ComplexObject
	AllWeekSchedule() *float64
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
	FridaySchedule() DataAzurermManagedDevopsPoolStatelessAgentManualResourcePredictionFridayScheduleList
	InternalValue() *DataAzurermManagedDevopsPoolStatelessAgentManualResourcePrediction
	SetInternalValue(val *DataAzurermManagedDevopsPoolStatelessAgentManualResourcePrediction)
	MondaySchedule() DataAzurermManagedDevopsPoolStatelessAgentManualResourcePredictionMondayScheduleList
	SaturdaySchedule() DataAzurermManagedDevopsPoolStatelessAgentManualResourcePredictionSaturdayScheduleList
	SundaySchedule() DataAzurermManagedDevopsPoolStatelessAgentManualResourcePredictionSundayScheduleList
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktn.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktn.IInterpolatingParent)
	ThursdaySchedule() DataAzurermManagedDevopsPoolStatelessAgentManualResourcePredictionThursdayScheduleList
	TimeZoneName() *string
	TuesdaySchedule() DataAzurermManagedDevopsPoolStatelessAgentManualResourcePredictionTuesdayScheduleList
	WednesdaySchedule() DataAzurermManagedDevopsPoolStatelessAgentManualResourcePredictionWednesdayScheduleList
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
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktn.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for DataAzurermManagedDevopsPoolStatelessAgentManualResourcePredictionOutputReference
type jsiiProxy_DataAzurermManagedDevopsPoolStatelessAgentManualResourcePredictionOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_DataAzurermManagedDevopsPoolStatelessAgentManualResourcePredictionOutputReference) AllWeekSchedule() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"allWeekSchedule",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAzurermManagedDevopsPoolStatelessAgentManualResourcePredictionOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAzurermManagedDevopsPoolStatelessAgentManualResourcePredictionOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAzurermManagedDevopsPoolStatelessAgentManualResourcePredictionOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAzurermManagedDevopsPoolStatelessAgentManualResourcePredictionOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAzurermManagedDevopsPoolStatelessAgentManualResourcePredictionOutputReference) FridaySchedule() DataAzurermManagedDevopsPoolStatelessAgentManualResourcePredictionFridayScheduleList {
	var returns DataAzurermManagedDevopsPoolStatelessAgentManualResourcePredictionFridayScheduleList
	_jsii_.Get(
		j,
		"fridaySchedule",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAzurermManagedDevopsPoolStatelessAgentManualResourcePredictionOutputReference) InternalValue() *DataAzurermManagedDevopsPoolStatelessAgentManualResourcePrediction {
	var returns *DataAzurermManagedDevopsPoolStatelessAgentManualResourcePrediction
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAzurermManagedDevopsPoolStatelessAgentManualResourcePredictionOutputReference) MondaySchedule() DataAzurermManagedDevopsPoolStatelessAgentManualResourcePredictionMondayScheduleList {
	var returns DataAzurermManagedDevopsPoolStatelessAgentManualResourcePredictionMondayScheduleList
	_jsii_.Get(
		j,
		"mondaySchedule",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAzurermManagedDevopsPoolStatelessAgentManualResourcePredictionOutputReference) SaturdaySchedule() DataAzurermManagedDevopsPoolStatelessAgentManualResourcePredictionSaturdayScheduleList {
	var returns DataAzurermManagedDevopsPoolStatelessAgentManualResourcePredictionSaturdayScheduleList
	_jsii_.Get(
		j,
		"saturdaySchedule",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAzurermManagedDevopsPoolStatelessAgentManualResourcePredictionOutputReference) SundaySchedule() DataAzurermManagedDevopsPoolStatelessAgentManualResourcePredictionSundayScheduleList {
	var returns DataAzurermManagedDevopsPoolStatelessAgentManualResourcePredictionSundayScheduleList
	_jsii_.Get(
		j,
		"sundaySchedule",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAzurermManagedDevopsPoolStatelessAgentManualResourcePredictionOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAzurermManagedDevopsPoolStatelessAgentManualResourcePredictionOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAzurermManagedDevopsPoolStatelessAgentManualResourcePredictionOutputReference) ThursdaySchedule() DataAzurermManagedDevopsPoolStatelessAgentManualResourcePredictionThursdayScheduleList {
	var returns DataAzurermManagedDevopsPoolStatelessAgentManualResourcePredictionThursdayScheduleList
	_jsii_.Get(
		j,
		"thursdaySchedule",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAzurermManagedDevopsPoolStatelessAgentManualResourcePredictionOutputReference) TimeZoneName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"timeZoneName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAzurermManagedDevopsPoolStatelessAgentManualResourcePredictionOutputReference) TuesdaySchedule() DataAzurermManagedDevopsPoolStatelessAgentManualResourcePredictionTuesdayScheduleList {
	var returns DataAzurermManagedDevopsPoolStatelessAgentManualResourcePredictionTuesdayScheduleList
	_jsii_.Get(
		j,
		"tuesdaySchedule",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAzurermManagedDevopsPoolStatelessAgentManualResourcePredictionOutputReference) WednesdaySchedule() DataAzurermManagedDevopsPoolStatelessAgentManualResourcePredictionWednesdayScheduleList {
	var returns DataAzurermManagedDevopsPoolStatelessAgentManualResourcePredictionWednesdayScheduleList
	_jsii_.Get(
		j,
		"wednesdaySchedule",
		&returns,
	)
	return returns
}


func NewDataAzurermManagedDevopsPoolStatelessAgentManualResourcePredictionOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) DataAzurermManagedDevopsPoolStatelessAgentManualResourcePredictionOutputReference {
	_init_.Initialize()

	if err := validateNewDataAzurermManagedDevopsPoolStatelessAgentManualResourcePredictionOutputReferenceParameters(terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet); err != nil {
		panic(err)
	}
	j := jsiiProxy_DataAzurermManagedDevopsPoolStatelessAgentManualResourcePredictionOutputReference{}

	_jsii_.Create(
		"@cdktn/provider-azurerm.dataAzurermManagedDevopsPool.DataAzurermManagedDevopsPoolStatelessAgentManualResourcePredictionOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		&j,
	)

	return &j
}

func NewDataAzurermManagedDevopsPoolStatelessAgentManualResourcePredictionOutputReference_Override(d DataAzurermManagedDevopsPoolStatelessAgentManualResourcePredictionOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/provider-azurerm.dataAzurermManagedDevopsPool.DataAzurermManagedDevopsPoolStatelessAgentManualResourcePredictionOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		d,
	)
}

func (j *jsiiProxy_DataAzurermManagedDevopsPoolStatelessAgentManualResourcePredictionOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_DataAzurermManagedDevopsPoolStatelessAgentManualResourcePredictionOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_DataAzurermManagedDevopsPoolStatelessAgentManualResourcePredictionOutputReference)SetInternalValue(val *DataAzurermManagedDevopsPoolStatelessAgentManualResourcePrediction) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_DataAzurermManagedDevopsPoolStatelessAgentManualResourcePredictionOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_DataAzurermManagedDevopsPoolStatelessAgentManualResourcePredictionOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (d *jsiiProxy_DataAzurermManagedDevopsPoolStatelessAgentManualResourcePredictionOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataAzurermManagedDevopsPoolStatelessAgentManualResourcePredictionOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	if err := d.validateGetAnyMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]interface{}

	_jsii_.Invoke(
		d,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataAzurermManagedDevopsPoolStatelessAgentManualResourcePredictionOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
	if err := d.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		d,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataAzurermManagedDevopsPoolStatelessAgentManualResourcePredictionOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	if err := d.validateGetBooleanMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*bool

	_jsii_.Invoke(
		d,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataAzurermManagedDevopsPoolStatelessAgentManualResourcePredictionOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	if err := d.validateGetListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*string

	_jsii_.Invoke(
		d,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataAzurermManagedDevopsPoolStatelessAgentManualResourcePredictionOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	if err := d.validateGetNumberAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *float64

	_jsii_.Invoke(
		d,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataAzurermManagedDevopsPoolStatelessAgentManualResourcePredictionOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	if err := d.validateGetNumberListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*float64

	_jsii_.Invoke(
		d,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataAzurermManagedDevopsPoolStatelessAgentManualResourcePredictionOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	if err := d.validateGetNumberMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*float64

	_jsii_.Invoke(
		d,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataAzurermManagedDevopsPoolStatelessAgentManualResourcePredictionOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	if err := d.validateGetStringAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		d,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataAzurermManagedDevopsPoolStatelessAgentManualResourcePredictionOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	if err := d.validateGetStringMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*string

	_jsii_.Invoke(
		d,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataAzurermManagedDevopsPoolStatelessAgentManualResourcePredictionOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		d,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataAzurermManagedDevopsPoolStatelessAgentManualResourcePredictionOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
	if err := d.validateInterpolationForAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		d,
		"interpolationForAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataAzurermManagedDevopsPoolStatelessAgentManualResourcePredictionOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
	if err := d.validateResolveParameters(context); err != nil {
		panic(err)
	}
	var returns interface{}

	_jsii_.Invoke(
		d,
		"resolve",
		[]interface{}{context},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataAzurermManagedDevopsPoolStatelessAgentManualResourcePredictionOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

