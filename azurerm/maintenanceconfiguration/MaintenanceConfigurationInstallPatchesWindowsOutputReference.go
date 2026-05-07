// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package maintenanceconfiguration

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-provider-azurerm-go/azurerm/v16/jsii"

	"github.com/cdktn-io/cdktn-provider-azurerm-go/azurerm/v16/maintenanceconfiguration/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type MaintenanceConfigurationInstallPatchesWindowsOutputReference interface {
	cdktn.ComplexObject
	ClassificationsToInclude() *[]*string
	SetClassificationsToInclude(val *[]*string)
	ClassificationsToIncludeInput() *[]*string
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
	InternalValue() interface{}
	SetInternalValue(val interface{})
	KbNumbersToExclude() *[]*string
	SetKbNumbersToExclude(val *[]*string)
	KbNumbersToExcludeInput() *[]*string
	KbNumbersToInclude() *[]*string
	SetKbNumbersToInclude(val *[]*string)
	KbNumbersToIncludeInput() *[]*string
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
	ResetClassificationsToInclude()
	ResetKbNumbersToExclude()
	ResetKbNumbersToInclude()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktn.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for MaintenanceConfigurationInstallPatchesWindowsOutputReference
type jsiiProxy_MaintenanceConfigurationInstallPatchesWindowsOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_MaintenanceConfigurationInstallPatchesWindowsOutputReference) ClassificationsToInclude() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"classificationsToInclude",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MaintenanceConfigurationInstallPatchesWindowsOutputReference) ClassificationsToIncludeInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"classificationsToIncludeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MaintenanceConfigurationInstallPatchesWindowsOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MaintenanceConfigurationInstallPatchesWindowsOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MaintenanceConfigurationInstallPatchesWindowsOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MaintenanceConfigurationInstallPatchesWindowsOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MaintenanceConfigurationInstallPatchesWindowsOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MaintenanceConfigurationInstallPatchesWindowsOutputReference) KbNumbersToExclude() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"kbNumbersToExclude",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MaintenanceConfigurationInstallPatchesWindowsOutputReference) KbNumbersToExcludeInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"kbNumbersToExcludeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MaintenanceConfigurationInstallPatchesWindowsOutputReference) KbNumbersToInclude() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"kbNumbersToInclude",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MaintenanceConfigurationInstallPatchesWindowsOutputReference) KbNumbersToIncludeInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"kbNumbersToIncludeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MaintenanceConfigurationInstallPatchesWindowsOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MaintenanceConfigurationInstallPatchesWindowsOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewMaintenanceConfigurationInstallPatchesWindowsOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) MaintenanceConfigurationInstallPatchesWindowsOutputReference {
	_init_.Initialize()

	if err := validateNewMaintenanceConfigurationInstallPatchesWindowsOutputReferenceParameters(terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet); err != nil {
		panic(err)
	}
	j := jsiiProxy_MaintenanceConfigurationInstallPatchesWindowsOutputReference{}

	_jsii_.Create(
		"@cdktn/provider-azurerm.maintenanceConfiguration.MaintenanceConfigurationInstallPatchesWindowsOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		&j,
	)

	return &j
}

func NewMaintenanceConfigurationInstallPatchesWindowsOutputReference_Override(m MaintenanceConfigurationInstallPatchesWindowsOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/provider-azurerm.maintenanceConfiguration.MaintenanceConfigurationInstallPatchesWindowsOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		m,
	)
}

func (j *jsiiProxy_MaintenanceConfigurationInstallPatchesWindowsOutputReference)SetClassificationsToInclude(val *[]*string) {
	if err := j.validateSetClassificationsToIncludeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"classificationsToInclude",
		val,
	)
}

func (j *jsiiProxy_MaintenanceConfigurationInstallPatchesWindowsOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_MaintenanceConfigurationInstallPatchesWindowsOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_MaintenanceConfigurationInstallPatchesWindowsOutputReference)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_MaintenanceConfigurationInstallPatchesWindowsOutputReference)SetKbNumbersToExclude(val *[]*string) {
	if err := j.validateSetKbNumbersToExcludeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"kbNumbersToExclude",
		val,
	)
}

func (j *jsiiProxy_MaintenanceConfigurationInstallPatchesWindowsOutputReference)SetKbNumbersToInclude(val *[]*string) {
	if err := j.validateSetKbNumbersToIncludeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"kbNumbersToInclude",
		val,
	)
}

func (j *jsiiProxy_MaintenanceConfigurationInstallPatchesWindowsOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_MaintenanceConfigurationInstallPatchesWindowsOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (m *jsiiProxy_MaintenanceConfigurationInstallPatchesWindowsOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		m,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MaintenanceConfigurationInstallPatchesWindowsOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (m *jsiiProxy_MaintenanceConfigurationInstallPatchesWindowsOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (m *jsiiProxy_MaintenanceConfigurationInstallPatchesWindowsOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (m *jsiiProxy_MaintenanceConfigurationInstallPatchesWindowsOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (m *jsiiProxy_MaintenanceConfigurationInstallPatchesWindowsOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (m *jsiiProxy_MaintenanceConfigurationInstallPatchesWindowsOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (m *jsiiProxy_MaintenanceConfigurationInstallPatchesWindowsOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (m *jsiiProxy_MaintenanceConfigurationInstallPatchesWindowsOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (m *jsiiProxy_MaintenanceConfigurationInstallPatchesWindowsOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (m *jsiiProxy_MaintenanceConfigurationInstallPatchesWindowsOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		m,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MaintenanceConfigurationInstallPatchesWindowsOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (m *jsiiProxy_MaintenanceConfigurationInstallPatchesWindowsOutputReference) ResetClassificationsToInclude() {
	_jsii_.InvokeVoid(
		m,
		"resetClassificationsToInclude",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MaintenanceConfigurationInstallPatchesWindowsOutputReference) ResetKbNumbersToExclude() {
	_jsii_.InvokeVoid(
		m,
		"resetKbNumbersToExclude",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MaintenanceConfigurationInstallPatchesWindowsOutputReference) ResetKbNumbersToInclude() {
	_jsii_.InvokeVoid(
		m,
		"resetKbNumbersToInclude",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MaintenanceConfigurationInstallPatchesWindowsOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
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

func (m *jsiiProxy_MaintenanceConfigurationInstallPatchesWindowsOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		m,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

