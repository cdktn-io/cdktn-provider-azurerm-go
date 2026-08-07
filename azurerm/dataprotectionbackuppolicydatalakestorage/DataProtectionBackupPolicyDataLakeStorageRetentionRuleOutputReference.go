// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package dataprotectionbackuppolicydatalakestorage

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-provider-azurerm-go/azurerm/v17/jsii"

	"github.com/cdktn-io/cdktn-provider-azurerm-go/azurerm/v17/dataprotectionbackuppolicydatalakestorage/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type DataProtectionBackupPolicyDataLakeStorageRetentionRuleOutputReference interface {
	cdktn.ComplexObject
	AbsoluteCriteria() *string
	SetAbsoluteCriteria(val *string)
	AbsoluteCriteriaInput() *string
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
	DaysOfWeek() *[]*string
	SetDaysOfWeek(val *[]*string)
	DaysOfWeekInput() *[]*string
	Duration() *string
	SetDuration(val *string)
	DurationInput() *string
	// Experimental.
	Fqn() *string
	InternalValue() interface{}
	SetInternalValue(val interface{})
	MonthsOfYear() *[]*string
	SetMonthsOfYear(val *[]*string)
	MonthsOfYearInput() *[]*string
	Name() *string
	SetName(val *string)
	NameInput() *string
	ScheduledBackupTimes() *[]*string
	SetScheduledBackupTimes(val *[]*string)
	ScheduledBackupTimesInput() *[]*string
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktn.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktn.IInterpolatingParent)
	WeeksOfMonth() *[]*string
	SetWeeksOfMonth(val *[]*string)
	WeeksOfMonthInput() *[]*string
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
	ResetAbsoluteCriteria()
	ResetDaysOfWeek()
	ResetMonthsOfYear()
	ResetScheduledBackupTimes()
	ResetWeeksOfMonth()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktn.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for DataProtectionBackupPolicyDataLakeStorageRetentionRuleOutputReference
type jsiiProxy_DataProtectionBackupPolicyDataLakeStorageRetentionRuleOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_DataProtectionBackupPolicyDataLakeStorageRetentionRuleOutputReference) AbsoluteCriteria() *string {
	var returns *string
	_jsii_.Get(
		j,
		"absoluteCriteria",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataProtectionBackupPolicyDataLakeStorageRetentionRuleOutputReference) AbsoluteCriteriaInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"absoluteCriteriaInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataProtectionBackupPolicyDataLakeStorageRetentionRuleOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataProtectionBackupPolicyDataLakeStorageRetentionRuleOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataProtectionBackupPolicyDataLakeStorageRetentionRuleOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataProtectionBackupPolicyDataLakeStorageRetentionRuleOutputReference) DaysOfWeek() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"daysOfWeek",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataProtectionBackupPolicyDataLakeStorageRetentionRuleOutputReference) DaysOfWeekInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"daysOfWeekInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataProtectionBackupPolicyDataLakeStorageRetentionRuleOutputReference) Duration() *string {
	var returns *string
	_jsii_.Get(
		j,
		"duration",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataProtectionBackupPolicyDataLakeStorageRetentionRuleOutputReference) DurationInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"durationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataProtectionBackupPolicyDataLakeStorageRetentionRuleOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataProtectionBackupPolicyDataLakeStorageRetentionRuleOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataProtectionBackupPolicyDataLakeStorageRetentionRuleOutputReference) MonthsOfYear() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"monthsOfYear",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataProtectionBackupPolicyDataLakeStorageRetentionRuleOutputReference) MonthsOfYearInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"monthsOfYearInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataProtectionBackupPolicyDataLakeStorageRetentionRuleOutputReference) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataProtectionBackupPolicyDataLakeStorageRetentionRuleOutputReference) NameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"nameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataProtectionBackupPolicyDataLakeStorageRetentionRuleOutputReference) ScheduledBackupTimes() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"scheduledBackupTimes",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataProtectionBackupPolicyDataLakeStorageRetentionRuleOutputReference) ScheduledBackupTimesInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"scheduledBackupTimesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataProtectionBackupPolicyDataLakeStorageRetentionRuleOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataProtectionBackupPolicyDataLakeStorageRetentionRuleOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataProtectionBackupPolicyDataLakeStorageRetentionRuleOutputReference) WeeksOfMonth() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"weeksOfMonth",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataProtectionBackupPolicyDataLakeStorageRetentionRuleOutputReference) WeeksOfMonthInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"weeksOfMonthInput",
		&returns,
	)
	return returns
}


func NewDataProtectionBackupPolicyDataLakeStorageRetentionRuleOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) DataProtectionBackupPolicyDataLakeStorageRetentionRuleOutputReference {
	_init_.Initialize()

	if err := validateNewDataProtectionBackupPolicyDataLakeStorageRetentionRuleOutputReferenceParameters(terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet); err != nil {
		panic(err)
	}
	j := jsiiProxy_DataProtectionBackupPolicyDataLakeStorageRetentionRuleOutputReference{}

	_jsii_.Create(
		"@cdktn/provider-azurerm.dataProtectionBackupPolicyDataLakeStorage.DataProtectionBackupPolicyDataLakeStorageRetentionRuleOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		&j,
	)

	return &j
}

func NewDataProtectionBackupPolicyDataLakeStorageRetentionRuleOutputReference_Override(d DataProtectionBackupPolicyDataLakeStorageRetentionRuleOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/provider-azurerm.dataProtectionBackupPolicyDataLakeStorage.DataProtectionBackupPolicyDataLakeStorageRetentionRuleOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		d,
	)
}

func (j *jsiiProxy_DataProtectionBackupPolicyDataLakeStorageRetentionRuleOutputReference)SetAbsoluteCriteria(val *string) {
	if err := j.validateSetAbsoluteCriteriaParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"absoluteCriteria",
		val,
	)
}

func (j *jsiiProxy_DataProtectionBackupPolicyDataLakeStorageRetentionRuleOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_DataProtectionBackupPolicyDataLakeStorageRetentionRuleOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_DataProtectionBackupPolicyDataLakeStorageRetentionRuleOutputReference)SetDaysOfWeek(val *[]*string) {
	if err := j.validateSetDaysOfWeekParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"daysOfWeek",
		val,
	)
}

func (j *jsiiProxy_DataProtectionBackupPolicyDataLakeStorageRetentionRuleOutputReference)SetDuration(val *string) {
	if err := j.validateSetDurationParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"duration",
		val,
	)
}

func (j *jsiiProxy_DataProtectionBackupPolicyDataLakeStorageRetentionRuleOutputReference)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_DataProtectionBackupPolicyDataLakeStorageRetentionRuleOutputReference)SetMonthsOfYear(val *[]*string) {
	if err := j.validateSetMonthsOfYearParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"monthsOfYear",
		val,
	)
}

func (j *jsiiProxy_DataProtectionBackupPolicyDataLakeStorageRetentionRuleOutputReference)SetName(val *string) {
	if err := j.validateSetNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"name",
		val,
	)
}

func (j *jsiiProxy_DataProtectionBackupPolicyDataLakeStorageRetentionRuleOutputReference)SetScheduledBackupTimes(val *[]*string) {
	if err := j.validateSetScheduledBackupTimesParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"scheduledBackupTimes",
		val,
	)
}

func (j *jsiiProxy_DataProtectionBackupPolicyDataLakeStorageRetentionRuleOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_DataProtectionBackupPolicyDataLakeStorageRetentionRuleOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_DataProtectionBackupPolicyDataLakeStorageRetentionRuleOutputReference)SetWeeksOfMonth(val *[]*string) {
	if err := j.validateSetWeeksOfMonthParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"weeksOfMonth",
		val,
	)
}

func (d *jsiiProxy_DataProtectionBackupPolicyDataLakeStorageRetentionRuleOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataProtectionBackupPolicyDataLakeStorageRetentionRuleOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (d *jsiiProxy_DataProtectionBackupPolicyDataLakeStorageRetentionRuleOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (d *jsiiProxy_DataProtectionBackupPolicyDataLakeStorageRetentionRuleOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (d *jsiiProxy_DataProtectionBackupPolicyDataLakeStorageRetentionRuleOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (d *jsiiProxy_DataProtectionBackupPolicyDataLakeStorageRetentionRuleOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (d *jsiiProxy_DataProtectionBackupPolicyDataLakeStorageRetentionRuleOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (d *jsiiProxy_DataProtectionBackupPolicyDataLakeStorageRetentionRuleOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (d *jsiiProxy_DataProtectionBackupPolicyDataLakeStorageRetentionRuleOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (d *jsiiProxy_DataProtectionBackupPolicyDataLakeStorageRetentionRuleOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (d *jsiiProxy_DataProtectionBackupPolicyDataLakeStorageRetentionRuleOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		d,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataProtectionBackupPolicyDataLakeStorageRetentionRuleOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (d *jsiiProxy_DataProtectionBackupPolicyDataLakeStorageRetentionRuleOutputReference) ResetAbsoluteCriteria() {
	_jsii_.InvokeVoid(
		d,
		"resetAbsoluteCriteria",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataProtectionBackupPolicyDataLakeStorageRetentionRuleOutputReference) ResetDaysOfWeek() {
	_jsii_.InvokeVoid(
		d,
		"resetDaysOfWeek",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataProtectionBackupPolicyDataLakeStorageRetentionRuleOutputReference) ResetMonthsOfYear() {
	_jsii_.InvokeVoid(
		d,
		"resetMonthsOfYear",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataProtectionBackupPolicyDataLakeStorageRetentionRuleOutputReference) ResetScheduledBackupTimes() {
	_jsii_.InvokeVoid(
		d,
		"resetScheduledBackupTimes",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataProtectionBackupPolicyDataLakeStorageRetentionRuleOutputReference) ResetWeeksOfMonth() {
	_jsii_.InvokeVoid(
		d,
		"resetWeeksOfMonth",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataProtectionBackupPolicyDataLakeStorageRetentionRuleOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
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

func (d *jsiiProxy_DataProtectionBackupPolicyDataLakeStorageRetentionRuleOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

