// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package manageddevopspool

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-provider-azurerm-go/azurerm/v17/jsii"

	"github.com/cdktn-io/cdktn-provider-azurerm-go/azurerm/v17/manageddevopspool/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type ManagedDevopsPoolVirtualMachineScaleSetFabricStorageOutputReference interface {
	cdktn.ComplexObject
	Caching() *string
	SetCaching(val *string)
	CachingInput() *string
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
	DiskSizeInGb() *float64
	SetDiskSizeInGb(val *float64)
	DiskSizeInGbInput() *float64
	DriveLetter() *string
	SetDriveLetter(val *string)
	DriveLetterInput() *string
	// Experimental.
	Fqn() *string
	InternalValue() *ManagedDevopsPoolVirtualMachineScaleSetFabricStorage
	SetInternalValue(val *ManagedDevopsPoolVirtualMachineScaleSetFabricStorage)
	StorageAccountType() *string
	SetStorageAccountType(val *string)
	StorageAccountTypeInput() *string
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
	ResetCaching()
	ResetDriveLetter()
	ResetStorageAccountType()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktn.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for ManagedDevopsPoolVirtualMachineScaleSetFabricStorageOutputReference
type jsiiProxy_ManagedDevopsPoolVirtualMachineScaleSetFabricStorageOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_ManagedDevopsPoolVirtualMachineScaleSetFabricStorageOutputReference) Caching() *string {
	var returns *string
	_jsii_.Get(
		j,
		"caching",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ManagedDevopsPoolVirtualMachineScaleSetFabricStorageOutputReference) CachingInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"cachingInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ManagedDevopsPoolVirtualMachineScaleSetFabricStorageOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ManagedDevopsPoolVirtualMachineScaleSetFabricStorageOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ManagedDevopsPoolVirtualMachineScaleSetFabricStorageOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ManagedDevopsPoolVirtualMachineScaleSetFabricStorageOutputReference) DiskSizeInGb() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"diskSizeInGb",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ManagedDevopsPoolVirtualMachineScaleSetFabricStorageOutputReference) DiskSizeInGbInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"diskSizeInGbInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ManagedDevopsPoolVirtualMachineScaleSetFabricStorageOutputReference) DriveLetter() *string {
	var returns *string
	_jsii_.Get(
		j,
		"driveLetter",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ManagedDevopsPoolVirtualMachineScaleSetFabricStorageOutputReference) DriveLetterInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"driveLetterInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ManagedDevopsPoolVirtualMachineScaleSetFabricStorageOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ManagedDevopsPoolVirtualMachineScaleSetFabricStorageOutputReference) InternalValue() *ManagedDevopsPoolVirtualMachineScaleSetFabricStorage {
	var returns *ManagedDevopsPoolVirtualMachineScaleSetFabricStorage
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ManagedDevopsPoolVirtualMachineScaleSetFabricStorageOutputReference) StorageAccountType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"storageAccountType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ManagedDevopsPoolVirtualMachineScaleSetFabricStorageOutputReference) StorageAccountTypeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"storageAccountTypeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ManagedDevopsPoolVirtualMachineScaleSetFabricStorageOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ManagedDevopsPoolVirtualMachineScaleSetFabricStorageOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewManagedDevopsPoolVirtualMachineScaleSetFabricStorageOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) ManagedDevopsPoolVirtualMachineScaleSetFabricStorageOutputReference {
	_init_.Initialize()

	if err := validateNewManagedDevopsPoolVirtualMachineScaleSetFabricStorageOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_ManagedDevopsPoolVirtualMachineScaleSetFabricStorageOutputReference{}

	_jsii_.Create(
		"@cdktn/provider-azurerm.managedDevopsPool.ManagedDevopsPoolVirtualMachineScaleSetFabricStorageOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewManagedDevopsPoolVirtualMachineScaleSetFabricStorageOutputReference_Override(m ManagedDevopsPoolVirtualMachineScaleSetFabricStorageOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/provider-azurerm.managedDevopsPool.ManagedDevopsPoolVirtualMachineScaleSetFabricStorageOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		m,
	)
}

func (j *jsiiProxy_ManagedDevopsPoolVirtualMachineScaleSetFabricStorageOutputReference)SetCaching(val *string) {
	if err := j.validateSetCachingParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"caching",
		val,
	)
}

func (j *jsiiProxy_ManagedDevopsPoolVirtualMachineScaleSetFabricStorageOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_ManagedDevopsPoolVirtualMachineScaleSetFabricStorageOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_ManagedDevopsPoolVirtualMachineScaleSetFabricStorageOutputReference)SetDiskSizeInGb(val *float64) {
	if err := j.validateSetDiskSizeInGbParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"diskSizeInGb",
		val,
	)
}

func (j *jsiiProxy_ManagedDevopsPoolVirtualMachineScaleSetFabricStorageOutputReference)SetDriveLetter(val *string) {
	if err := j.validateSetDriveLetterParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"driveLetter",
		val,
	)
}

func (j *jsiiProxy_ManagedDevopsPoolVirtualMachineScaleSetFabricStorageOutputReference)SetInternalValue(val *ManagedDevopsPoolVirtualMachineScaleSetFabricStorage) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_ManagedDevopsPoolVirtualMachineScaleSetFabricStorageOutputReference)SetStorageAccountType(val *string) {
	if err := j.validateSetStorageAccountTypeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"storageAccountType",
		val,
	)
}

func (j *jsiiProxy_ManagedDevopsPoolVirtualMachineScaleSetFabricStorageOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_ManagedDevopsPoolVirtualMachineScaleSetFabricStorageOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (m *jsiiProxy_ManagedDevopsPoolVirtualMachineScaleSetFabricStorageOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		m,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (m *jsiiProxy_ManagedDevopsPoolVirtualMachineScaleSetFabricStorageOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (m *jsiiProxy_ManagedDevopsPoolVirtualMachineScaleSetFabricStorageOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (m *jsiiProxy_ManagedDevopsPoolVirtualMachineScaleSetFabricStorageOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (m *jsiiProxy_ManagedDevopsPoolVirtualMachineScaleSetFabricStorageOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (m *jsiiProxy_ManagedDevopsPoolVirtualMachineScaleSetFabricStorageOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (m *jsiiProxy_ManagedDevopsPoolVirtualMachineScaleSetFabricStorageOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (m *jsiiProxy_ManagedDevopsPoolVirtualMachineScaleSetFabricStorageOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (m *jsiiProxy_ManagedDevopsPoolVirtualMachineScaleSetFabricStorageOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (m *jsiiProxy_ManagedDevopsPoolVirtualMachineScaleSetFabricStorageOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (m *jsiiProxy_ManagedDevopsPoolVirtualMachineScaleSetFabricStorageOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		m,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (m *jsiiProxy_ManagedDevopsPoolVirtualMachineScaleSetFabricStorageOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (m *jsiiProxy_ManagedDevopsPoolVirtualMachineScaleSetFabricStorageOutputReference) ResetCaching() {
	_jsii_.InvokeVoid(
		m,
		"resetCaching",
		nil, // no parameters
	)
}

func (m *jsiiProxy_ManagedDevopsPoolVirtualMachineScaleSetFabricStorageOutputReference) ResetDriveLetter() {
	_jsii_.InvokeVoid(
		m,
		"resetDriveLetter",
		nil, // no parameters
	)
}

func (m *jsiiProxy_ManagedDevopsPoolVirtualMachineScaleSetFabricStorageOutputReference) ResetStorageAccountType() {
	_jsii_.InvokeVoid(
		m,
		"resetStorageAccountType",
		nil, // no parameters
	)
}

func (m *jsiiProxy_ManagedDevopsPoolVirtualMachineScaleSetFabricStorageOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
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

func (m *jsiiProxy_ManagedDevopsPoolVirtualMachineScaleSetFabricStorageOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		m,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

