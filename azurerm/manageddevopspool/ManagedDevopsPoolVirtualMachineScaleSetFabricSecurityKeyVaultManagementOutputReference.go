// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package manageddevopspool

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-provider-azurerm-go/azurerm/v16/jsii"

	"github.com/cdktn-io/cdktn-provider-azurerm-go/azurerm/v16/manageddevopspool/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type ManagedDevopsPoolVirtualMachineScaleSetFabricSecurityKeyVaultManagementOutputReference interface {
	cdktn.ComplexObject
	CertificateStoreLocation() *string
	SetCertificateStoreLocation(val *string)
	CertificateStoreLocationInput() *string
	CertificateStoreName() *string
	SetCertificateStoreName(val *string)
	CertificateStoreNameInput() *string
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
	InternalValue() *ManagedDevopsPoolVirtualMachineScaleSetFabricSecurityKeyVaultManagement
	SetInternalValue(val *ManagedDevopsPoolVirtualMachineScaleSetFabricSecurityKeyVaultManagement)
	KeyExportEnabled() interface{}
	SetKeyExportEnabled(val interface{})
	KeyExportEnabledInput() interface{}
	KeyVaultCertificateIds() *[]*string
	SetKeyVaultCertificateIds(val *[]*string)
	KeyVaultCertificateIdsInput() *[]*string
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
	ResetCertificateStoreLocation()
	ResetCertificateStoreName()
	ResetKeyExportEnabled()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktn.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for ManagedDevopsPoolVirtualMachineScaleSetFabricSecurityKeyVaultManagementOutputReference
type jsiiProxy_ManagedDevopsPoolVirtualMachineScaleSetFabricSecurityKeyVaultManagementOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_ManagedDevopsPoolVirtualMachineScaleSetFabricSecurityKeyVaultManagementOutputReference) CertificateStoreLocation() *string {
	var returns *string
	_jsii_.Get(
		j,
		"certificateStoreLocation",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ManagedDevopsPoolVirtualMachineScaleSetFabricSecurityKeyVaultManagementOutputReference) CertificateStoreLocationInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"certificateStoreLocationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ManagedDevopsPoolVirtualMachineScaleSetFabricSecurityKeyVaultManagementOutputReference) CertificateStoreName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"certificateStoreName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ManagedDevopsPoolVirtualMachineScaleSetFabricSecurityKeyVaultManagementOutputReference) CertificateStoreNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"certificateStoreNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ManagedDevopsPoolVirtualMachineScaleSetFabricSecurityKeyVaultManagementOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ManagedDevopsPoolVirtualMachineScaleSetFabricSecurityKeyVaultManagementOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ManagedDevopsPoolVirtualMachineScaleSetFabricSecurityKeyVaultManagementOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ManagedDevopsPoolVirtualMachineScaleSetFabricSecurityKeyVaultManagementOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ManagedDevopsPoolVirtualMachineScaleSetFabricSecurityKeyVaultManagementOutputReference) InternalValue() *ManagedDevopsPoolVirtualMachineScaleSetFabricSecurityKeyVaultManagement {
	var returns *ManagedDevopsPoolVirtualMachineScaleSetFabricSecurityKeyVaultManagement
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ManagedDevopsPoolVirtualMachineScaleSetFabricSecurityKeyVaultManagementOutputReference) KeyExportEnabled() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"keyExportEnabled",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ManagedDevopsPoolVirtualMachineScaleSetFabricSecurityKeyVaultManagementOutputReference) KeyExportEnabledInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"keyExportEnabledInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ManagedDevopsPoolVirtualMachineScaleSetFabricSecurityKeyVaultManagementOutputReference) KeyVaultCertificateIds() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"keyVaultCertificateIds",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ManagedDevopsPoolVirtualMachineScaleSetFabricSecurityKeyVaultManagementOutputReference) KeyVaultCertificateIdsInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"keyVaultCertificateIdsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ManagedDevopsPoolVirtualMachineScaleSetFabricSecurityKeyVaultManagementOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ManagedDevopsPoolVirtualMachineScaleSetFabricSecurityKeyVaultManagementOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewManagedDevopsPoolVirtualMachineScaleSetFabricSecurityKeyVaultManagementOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) ManagedDevopsPoolVirtualMachineScaleSetFabricSecurityKeyVaultManagementOutputReference {
	_init_.Initialize()

	if err := validateNewManagedDevopsPoolVirtualMachineScaleSetFabricSecurityKeyVaultManagementOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_ManagedDevopsPoolVirtualMachineScaleSetFabricSecurityKeyVaultManagementOutputReference{}

	_jsii_.Create(
		"@cdktn/provider-azurerm.managedDevopsPool.ManagedDevopsPoolVirtualMachineScaleSetFabricSecurityKeyVaultManagementOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewManagedDevopsPoolVirtualMachineScaleSetFabricSecurityKeyVaultManagementOutputReference_Override(m ManagedDevopsPoolVirtualMachineScaleSetFabricSecurityKeyVaultManagementOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/provider-azurerm.managedDevopsPool.ManagedDevopsPoolVirtualMachineScaleSetFabricSecurityKeyVaultManagementOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		m,
	)
}

func (j *jsiiProxy_ManagedDevopsPoolVirtualMachineScaleSetFabricSecurityKeyVaultManagementOutputReference)SetCertificateStoreLocation(val *string) {
	if err := j.validateSetCertificateStoreLocationParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"certificateStoreLocation",
		val,
	)
}

func (j *jsiiProxy_ManagedDevopsPoolVirtualMachineScaleSetFabricSecurityKeyVaultManagementOutputReference)SetCertificateStoreName(val *string) {
	if err := j.validateSetCertificateStoreNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"certificateStoreName",
		val,
	)
}

func (j *jsiiProxy_ManagedDevopsPoolVirtualMachineScaleSetFabricSecurityKeyVaultManagementOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_ManagedDevopsPoolVirtualMachineScaleSetFabricSecurityKeyVaultManagementOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_ManagedDevopsPoolVirtualMachineScaleSetFabricSecurityKeyVaultManagementOutputReference)SetInternalValue(val *ManagedDevopsPoolVirtualMachineScaleSetFabricSecurityKeyVaultManagement) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_ManagedDevopsPoolVirtualMachineScaleSetFabricSecurityKeyVaultManagementOutputReference)SetKeyExportEnabled(val interface{}) {
	if err := j.validateSetKeyExportEnabledParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"keyExportEnabled",
		val,
	)
}

func (j *jsiiProxy_ManagedDevopsPoolVirtualMachineScaleSetFabricSecurityKeyVaultManagementOutputReference)SetKeyVaultCertificateIds(val *[]*string) {
	if err := j.validateSetKeyVaultCertificateIdsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"keyVaultCertificateIds",
		val,
	)
}

func (j *jsiiProxy_ManagedDevopsPoolVirtualMachineScaleSetFabricSecurityKeyVaultManagementOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_ManagedDevopsPoolVirtualMachineScaleSetFabricSecurityKeyVaultManagementOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (m *jsiiProxy_ManagedDevopsPoolVirtualMachineScaleSetFabricSecurityKeyVaultManagementOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		m,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (m *jsiiProxy_ManagedDevopsPoolVirtualMachineScaleSetFabricSecurityKeyVaultManagementOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (m *jsiiProxy_ManagedDevopsPoolVirtualMachineScaleSetFabricSecurityKeyVaultManagementOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (m *jsiiProxy_ManagedDevopsPoolVirtualMachineScaleSetFabricSecurityKeyVaultManagementOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (m *jsiiProxy_ManagedDevopsPoolVirtualMachineScaleSetFabricSecurityKeyVaultManagementOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (m *jsiiProxy_ManagedDevopsPoolVirtualMachineScaleSetFabricSecurityKeyVaultManagementOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (m *jsiiProxy_ManagedDevopsPoolVirtualMachineScaleSetFabricSecurityKeyVaultManagementOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (m *jsiiProxy_ManagedDevopsPoolVirtualMachineScaleSetFabricSecurityKeyVaultManagementOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (m *jsiiProxy_ManagedDevopsPoolVirtualMachineScaleSetFabricSecurityKeyVaultManagementOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (m *jsiiProxy_ManagedDevopsPoolVirtualMachineScaleSetFabricSecurityKeyVaultManagementOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (m *jsiiProxy_ManagedDevopsPoolVirtualMachineScaleSetFabricSecurityKeyVaultManagementOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		m,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (m *jsiiProxy_ManagedDevopsPoolVirtualMachineScaleSetFabricSecurityKeyVaultManagementOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (m *jsiiProxy_ManagedDevopsPoolVirtualMachineScaleSetFabricSecurityKeyVaultManagementOutputReference) ResetCertificateStoreLocation() {
	_jsii_.InvokeVoid(
		m,
		"resetCertificateStoreLocation",
		nil, // no parameters
	)
}

func (m *jsiiProxy_ManagedDevopsPoolVirtualMachineScaleSetFabricSecurityKeyVaultManagementOutputReference) ResetCertificateStoreName() {
	_jsii_.InvokeVoid(
		m,
		"resetCertificateStoreName",
		nil, // no parameters
	)
}

func (m *jsiiProxy_ManagedDevopsPoolVirtualMachineScaleSetFabricSecurityKeyVaultManagementOutputReference) ResetKeyExportEnabled() {
	_jsii_.InvokeVoid(
		m,
		"resetKeyExportEnabled",
		nil, // no parameters
	)
}

func (m *jsiiProxy_ManagedDevopsPoolVirtualMachineScaleSetFabricSecurityKeyVaultManagementOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
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

func (m *jsiiProxy_ManagedDevopsPoolVirtualMachineScaleSetFabricSecurityKeyVaultManagementOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		m,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

