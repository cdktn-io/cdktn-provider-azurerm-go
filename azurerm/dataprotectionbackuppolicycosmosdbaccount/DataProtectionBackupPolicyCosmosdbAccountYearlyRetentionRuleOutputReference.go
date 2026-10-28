// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package dataprotectionbackuppolicycosmosdbaccount

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-provider-azurerm-go/azurerm/v19/jsii"

	"github.com/cdktn-io/cdktn-provider-azurerm-go/azurerm/v19/dataprotectionbackuppolicycosmosdbaccount/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type DataProtectionBackupPolicyCosmosdbAccountYearlyRetentionRuleOutputReference interface {
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
	DaysOfWeek() *[]*string
	SetDaysOfWeek(val *[]*string)
	DaysOfWeekInput() *[]*string
	Duration() *string
	SetDuration(val *string)
	DurationInput() *string
	// Experimental.
	Fqn() *string
	InternalValue() *DataProtectionBackupPolicyCosmosdbAccountYearlyRetentionRule
	SetInternalValue(val *DataProtectionBackupPolicyCosmosdbAccountYearlyRetentionRule)
	MonthsOfYear() *[]*string
	SetMonthsOfYear(val *[]*string)
	MonthsOfYearInput() *[]*string
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
	ResetDaysOfWeek()
	ResetMonthsOfYear()
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

// The jsii proxy struct for DataProtectionBackupPolicyCosmosdbAccountYearlyRetentionRuleOutputReference
type jsiiProxy_DataProtectionBackupPolicyCosmosdbAccountYearlyRetentionRuleOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_DataProtectionBackupPolicyCosmosdbAccountYearlyRetentionRuleOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataProtectionBackupPolicyCosmosdbAccountYearlyRetentionRuleOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataProtectionBackupPolicyCosmosdbAccountYearlyRetentionRuleOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataProtectionBackupPolicyCosmosdbAccountYearlyRetentionRuleOutputReference) DaysOfWeek() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"daysOfWeek",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataProtectionBackupPolicyCosmosdbAccountYearlyRetentionRuleOutputReference) DaysOfWeekInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"daysOfWeekInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataProtectionBackupPolicyCosmosdbAccountYearlyRetentionRuleOutputReference) Duration() *string {
	var returns *string
	_jsii_.Get(
		j,
		"duration",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataProtectionBackupPolicyCosmosdbAccountYearlyRetentionRuleOutputReference) DurationInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"durationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataProtectionBackupPolicyCosmosdbAccountYearlyRetentionRuleOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataProtectionBackupPolicyCosmosdbAccountYearlyRetentionRuleOutputReference) InternalValue() *DataProtectionBackupPolicyCosmosdbAccountYearlyRetentionRule {
	var returns *DataProtectionBackupPolicyCosmosdbAccountYearlyRetentionRule
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataProtectionBackupPolicyCosmosdbAccountYearlyRetentionRuleOutputReference) MonthsOfYear() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"monthsOfYear",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataProtectionBackupPolicyCosmosdbAccountYearlyRetentionRuleOutputReference) MonthsOfYearInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"monthsOfYearInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataProtectionBackupPolicyCosmosdbAccountYearlyRetentionRuleOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataProtectionBackupPolicyCosmosdbAccountYearlyRetentionRuleOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataProtectionBackupPolicyCosmosdbAccountYearlyRetentionRuleOutputReference) WeeksOfMonth() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"weeksOfMonth",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataProtectionBackupPolicyCosmosdbAccountYearlyRetentionRuleOutputReference) WeeksOfMonthInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"weeksOfMonthInput",
		&returns,
	)
	return returns
}


func NewDataProtectionBackupPolicyCosmosdbAccountYearlyRetentionRuleOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) DataProtectionBackupPolicyCosmosdbAccountYearlyRetentionRuleOutputReference {
	_init_.Initialize()

	if err := validateNewDataProtectionBackupPolicyCosmosdbAccountYearlyRetentionRuleOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_DataProtectionBackupPolicyCosmosdbAccountYearlyRetentionRuleOutputReference{}

	_jsii_.Create(
		"@cdktn/provider-azurerm.dataProtectionBackupPolicyCosmosdbAccount.DataProtectionBackupPolicyCosmosdbAccountYearlyRetentionRuleOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewDataProtectionBackupPolicyCosmosdbAccountYearlyRetentionRuleOutputReference_Override(d DataProtectionBackupPolicyCosmosdbAccountYearlyRetentionRuleOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/provider-azurerm.dataProtectionBackupPolicyCosmosdbAccount.DataProtectionBackupPolicyCosmosdbAccountYearlyRetentionRuleOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		d,
	)
}

func (j *jsiiProxy_DataProtectionBackupPolicyCosmosdbAccountYearlyRetentionRuleOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_DataProtectionBackupPolicyCosmosdbAccountYearlyRetentionRuleOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_DataProtectionBackupPolicyCosmosdbAccountYearlyRetentionRuleOutputReference)SetDaysOfWeek(val *[]*string) {
	if err := j.validateSetDaysOfWeekParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"daysOfWeek",
		val,
	)
}

func (j *jsiiProxy_DataProtectionBackupPolicyCosmosdbAccountYearlyRetentionRuleOutputReference)SetDuration(val *string) {
	if err := j.validateSetDurationParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"duration",
		val,
	)
}

func (j *jsiiProxy_DataProtectionBackupPolicyCosmosdbAccountYearlyRetentionRuleOutputReference)SetInternalValue(val *DataProtectionBackupPolicyCosmosdbAccountYearlyRetentionRule) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_DataProtectionBackupPolicyCosmosdbAccountYearlyRetentionRuleOutputReference)SetMonthsOfYear(val *[]*string) {
	if err := j.validateSetMonthsOfYearParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"monthsOfYear",
		val,
	)
}

func (j *jsiiProxy_DataProtectionBackupPolicyCosmosdbAccountYearlyRetentionRuleOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_DataProtectionBackupPolicyCosmosdbAccountYearlyRetentionRuleOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_DataProtectionBackupPolicyCosmosdbAccountYearlyRetentionRuleOutputReference)SetWeeksOfMonth(val *[]*string) {
	if err := j.validateSetWeeksOfMonthParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"weeksOfMonth",
		val,
	)
}

func (d *jsiiProxy_DataProtectionBackupPolicyCosmosdbAccountYearlyRetentionRuleOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataProtectionBackupPolicyCosmosdbAccountYearlyRetentionRuleOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (d *jsiiProxy_DataProtectionBackupPolicyCosmosdbAccountYearlyRetentionRuleOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (d *jsiiProxy_DataProtectionBackupPolicyCosmosdbAccountYearlyRetentionRuleOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (d *jsiiProxy_DataProtectionBackupPolicyCosmosdbAccountYearlyRetentionRuleOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (d *jsiiProxy_DataProtectionBackupPolicyCosmosdbAccountYearlyRetentionRuleOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (d *jsiiProxy_DataProtectionBackupPolicyCosmosdbAccountYearlyRetentionRuleOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (d *jsiiProxy_DataProtectionBackupPolicyCosmosdbAccountYearlyRetentionRuleOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (d *jsiiProxy_DataProtectionBackupPolicyCosmosdbAccountYearlyRetentionRuleOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (d *jsiiProxy_DataProtectionBackupPolicyCosmosdbAccountYearlyRetentionRuleOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (d *jsiiProxy_DataProtectionBackupPolicyCosmosdbAccountYearlyRetentionRuleOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		d,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataProtectionBackupPolicyCosmosdbAccountYearlyRetentionRuleOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (d *jsiiProxy_DataProtectionBackupPolicyCosmosdbAccountYearlyRetentionRuleOutputReference) ResetDaysOfWeek() {
	_jsii_.InvokeVoid(
		d,
		"resetDaysOfWeek",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataProtectionBackupPolicyCosmosdbAccountYearlyRetentionRuleOutputReference) ResetMonthsOfYear() {
	_jsii_.InvokeVoid(
		d,
		"resetMonthsOfYear",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataProtectionBackupPolicyCosmosdbAccountYearlyRetentionRuleOutputReference) ResetWeeksOfMonth() {
	_jsii_.InvokeVoid(
		d,
		"resetWeeksOfMonth",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataProtectionBackupPolicyCosmosdbAccountYearlyRetentionRuleOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
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

func (d *jsiiProxy_DataProtectionBackupPolicyCosmosdbAccountYearlyRetentionRuleOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

