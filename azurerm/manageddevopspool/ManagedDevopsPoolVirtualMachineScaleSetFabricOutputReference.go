// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package manageddevopspool

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-provider-azurerm-go/azurerm/v18/jsii"

	"github.com/cdktn-io/cdktn-provider-azurerm-go/azurerm/v18/manageddevopspool/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type ManagedDevopsPoolVirtualMachineScaleSetFabricOutputReference interface {
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
	// Experimental.
	Fqn() *string
	Image() ManagedDevopsPoolVirtualMachineScaleSetFabricImageList
	ImageInput() interface{}
	InternalValue() *ManagedDevopsPoolVirtualMachineScaleSetFabric
	SetInternalValue(val *ManagedDevopsPoolVirtualMachineScaleSetFabric)
	OsDiskStorageAccountType() *string
	SetOsDiskStorageAccountType(val *string)
	OsDiskStorageAccountTypeInput() *string
	Security() ManagedDevopsPoolVirtualMachineScaleSetFabricSecurityOutputReference
	SecurityInput() *ManagedDevopsPoolVirtualMachineScaleSetFabricSecurity
	SkuName() *string
	SetSkuName(val *string)
	SkuNameInput() *string
	Storage() ManagedDevopsPoolVirtualMachineScaleSetFabricStorageOutputReference
	StorageInput() *ManagedDevopsPoolVirtualMachineScaleSetFabricStorage
	SubnetId() *string
	SetSubnetId(val *string)
	SubnetIdInput() *string
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
	PutImage(value interface{})
	PutSecurity(value *ManagedDevopsPoolVirtualMachineScaleSetFabricSecurity)
	PutStorage(value *ManagedDevopsPoolVirtualMachineScaleSetFabricStorage)
	ResetOsDiskStorageAccountType()
	ResetSecurity()
	ResetStorage()
	ResetSubnetId()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktn.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for ManagedDevopsPoolVirtualMachineScaleSetFabricOutputReference
type jsiiProxy_ManagedDevopsPoolVirtualMachineScaleSetFabricOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_ManagedDevopsPoolVirtualMachineScaleSetFabricOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ManagedDevopsPoolVirtualMachineScaleSetFabricOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ManagedDevopsPoolVirtualMachineScaleSetFabricOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ManagedDevopsPoolVirtualMachineScaleSetFabricOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ManagedDevopsPoolVirtualMachineScaleSetFabricOutputReference) Image() ManagedDevopsPoolVirtualMachineScaleSetFabricImageList {
	var returns ManagedDevopsPoolVirtualMachineScaleSetFabricImageList
	_jsii_.Get(
		j,
		"image",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ManagedDevopsPoolVirtualMachineScaleSetFabricOutputReference) ImageInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"imageInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ManagedDevopsPoolVirtualMachineScaleSetFabricOutputReference) InternalValue() *ManagedDevopsPoolVirtualMachineScaleSetFabric {
	var returns *ManagedDevopsPoolVirtualMachineScaleSetFabric
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ManagedDevopsPoolVirtualMachineScaleSetFabricOutputReference) OsDiskStorageAccountType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"osDiskStorageAccountType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ManagedDevopsPoolVirtualMachineScaleSetFabricOutputReference) OsDiskStorageAccountTypeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"osDiskStorageAccountTypeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ManagedDevopsPoolVirtualMachineScaleSetFabricOutputReference) Security() ManagedDevopsPoolVirtualMachineScaleSetFabricSecurityOutputReference {
	var returns ManagedDevopsPoolVirtualMachineScaleSetFabricSecurityOutputReference
	_jsii_.Get(
		j,
		"security",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ManagedDevopsPoolVirtualMachineScaleSetFabricOutputReference) SecurityInput() *ManagedDevopsPoolVirtualMachineScaleSetFabricSecurity {
	var returns *ManagedDevopsPoolVirtualMachineScaleSetFabricSecurity
	_jsii_.Get(
		j,
		"securityInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ManagedDevopsPoolVirtualMachineScaleSetFabricOutputReference) SkuName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"skuName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ManagedDevopsPoolVirtualMachineScaleSetFabricOutputReference) SkuNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"skuNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ManagedDevopsPoolVirtualMachineScaleSetFabricOutputReference) Storage() ManagedDevopsPoolVirtualMachineScaleSetFabricStorageOutputReference {
	var returns ManagedDevopsPoolVirtualMachineScaleSetFabricStorageOutputReference
	_jsii_.Get(
		j,
		"storage",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ManagedDevopsPoolVirtualMachineScaleSetFabricOutputReference) StorageInput() *ManagedDevopsPoolVirtualMachineScaleSetFabricStorage {
	var returns *ManagedDevopsPoolVirtualMachineScaleSetFabricStorage
	_jsii_.Get(
		j,
		"storageInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ManagedDevopsPoolVirtualMachineScaleSetFabricOutputReference) SubnetId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"subnetId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ManagedDevopsPoolVirtualMachineScaleSetFabricOutputReference) SubnetIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"subnetIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ManagedDevopsPoolVirtualMachineScaleSetFabricOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ManagedDevopsPoolVirtualMachineScaleSetFabricOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewManagedDevopsPoolVirtualMachineScaleSetFabricOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) ManagedDevopsPoolVirtualMachineScaleSetFabricOutputReference {
	_init_.Initialize()

	if err := validateNewManagedDevopsPoolVirtualMachineScaleSetFabricOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_ManagedDevopsPoolVirtualMachineScaleSetFabricOutputReference{}

	_jsii_.Create(
		"@cdktn/provider-azurerm.managedDevopsPool.ManagedDevopsPoolVirtualMachineScaleSetFabricOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewManagedDevopsPoolVirtualMachineScaleSetFabricOutputReference_Override(m ManagedDevopsPoolVirtualMachineScaleSetFabricOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/provider-azurerm.managedDevopsPool.ManagedDevopsPoolVirtualMachineScaleSetFabricOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		m,
	)
}

func (j *jsiiProxy_ManagedDevopsPoolVirtualMachineScaleSetFabricOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_ManagedDevopsPoolVirtualMachineScaleSetFabricOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_ManagedDevopsPoolVirtualMachineScaleSetFabricOutputReference)SetInternalValue(val *ManagedDevopsPoolVirtualMachineScaleSetFabric) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_ManagedDevopsPoolVirtualMachineScaleSetFabricOutputReference)SetOsDiskStorageAccountType(val *string) {
	if err := j.validateSetOsDiskStorageAccountTypeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"osDiskStorageAccountType",
		val,
	)
}

func (j *jsiiProxy_ManagedDevopsPoolVirtualMachineScaleSetFabricOutputReference)SetSkuName(val *string) {
	if err := j.validateSetSkuNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"skuName",
		val,
	)
}

func (j *jsiiProxy_ManagedDevopsPoolVirtualMachineScaleSetFabricOutputReference)SetSubnetId(val *string) {
	if err := j.validateSetSubnetIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"subnetId",
		val,
	)
}

func (j *jsiiProxy_ManagedDevopsPoolVirtualMachineScaleSetFabricOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_ManagedDevopsPoolVirtualMachineScaleSetFabricOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (m *jsiiProxy_ManagedDevopsPoolVirtualMachineScaleSetFabricOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		m,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (m *jsiiProxy_ManagedDevopsPoolVirtualMachineScaleSetFabricOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (m *jsiiProxy_ManagedDevopsPoolVirtualMachineScaleSetFabricOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (m *jsiiProxy_ManagedDevopsPoolVirtualMachineScaleSetFabricOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (m *jsiiProxy_ManagedDevopsPoolVirtualMachineScaleSetFabricOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (m *jsiiProxy_ManagedDevopsPoolVirtualMachineScaleSetFabricOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (m *jsiiProxy_ManagedDevopsPoolVirtualMachineScaleSetFabricOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (m *jsiiProxy_ManagedDevopsPoolVirtualMachineScaleSetFabricOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (m *jsiiProxy_ManagedDevopsPoolVirtualMachineScaleSetFabricOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (m *jsiiProxy_ManagedDevopsPoolVirtualMachineScaleSetFabricOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (m *jsiiProxy_ManagedDevopsPoolVirtualMachineScaleSetFabricOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		m,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (m *jsiiProxy_ManagedDevopsPoolVirtualMachineScaleSetFabricOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (m *jsiiProxy_ManagedDevopsPoolVirtualMachineScaleSetFabricOutputReference) PutImage(value interface{}) {
	if err := m.validatePutImageParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		m,
		"putImage",
		[]interface{}{value},
	)
}

func (m *jsiiProxy_ManagedDevopsPoolVirtualMachineScaleSetFabricOutputReference) PutSecurity(value *ManagedDevopsPoolVirtualMachineScaleSetFabricSecurity) {
	if err := m.validatePutSecurityParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		m,
		"putSecurity",
		[]interface{}{value},
	)
}

func (m *jsiiProxy_ManagedDevopsPoolVirtualMachineScaleSetFabricOutputReference) PutStorage(value *ManagedDevopsPoolVirtualMachineScaleSetFabricStorage) {
	if err := m.validatePutStorageParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		m,
		"putStorage",
		[]interface{}{value},
	)
}

func (m *jsiiProxy_ManagedDevopsPoolVirtualMachineScaleSetFabricOutputReference) ResetOsDiskStorageAccountType() {
	_jsii_.InvokeVoid(
		m,
		"resetOsDiskStorageAccountType",
		nil, // no parameters
	)
}

func (m *jsiiProxy_ManagedDevopsPoolVirtualMachineScaleSetFabricOutputReference) ResetSecurity() {
	_jsii_.InvokeVoid(
		m,
		"resetSecurity",
		nil, // no parameters
	)
}

func (m *jsiiProxy_ManagedDevopsPoolVirtualMachineScaleSetFabricOutputReference) ResetStorage() {
	_jsii_.InvokeVoid(
		m,
		"resetStorage",
		nil, // no parameters
	)
}

func (m *jsiiProxy_ManagedDevopsPoolVirtualMachineScaleSetFabricOutputReference) ResetSubnetId() {
	_jsii_.InvokeVoid(
		m,
		"resetSubnetId",
		nil, // no parameters
	)
}

func (m *jsiiProxy_ManagedDevopsPoolVirtualMachineScaleSetFabricOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
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

func (m *jsiiProxy_ManagedDevopsPoolVirtualMachineScaleSetFabricOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		m,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

