// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package orchestratedvirtualmachinescaleset

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-provider-azurerm-go/azurerm/v18/jsii"

	"github.com/cdktn-io/cdktn-provider-azurerm-go/azurerm/v18/orchestratedvirtualmachinescaleset/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type OrchestratedVirtualMachineScaleSetNetworkInterfaceOutputReference interface {
	cdktn.ComplexObject
	AcceleratedNetworkingEnabled() interface{}
	SetAcceleratedNetworkingEnabled(val interface{})
	AcceleratedNetworkingEnabledInput() interface{}
	AuxiliaryMode() *string
	SetAuxiliaryMode(val *string)
	AuxiliaryModeInput() *string
	AuxiliarySku() *string
	SetAuxiliarySku(val *string)
	AuxiliarySkuInput() *string
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
	DnsServers() *[]*string
	SetDnsServers(val *[]*string)
	DnsServersInput() *[]*string
	// Experimental.
	Fqn() *string
	InternalValue() interface{}
	SetInternalValue(val interface{})
	IpConfiguration() OrchestratedVirtualMachineScaleSetNetworkInterfaceIpConfigurationList
	IpConfigurationInput() interface{}
	IpForwardingEnabled() interface{}
	SetIpForwardingEnabled(val interface{})
	IpForwardingEnabledInput() interface{}
	Name() *string
	SetName(val *string)
	NameInput() *string
	NetworkSecurityGroupId() *string
	SetNetworkSecurityGroupId(val *string)
	NetworkSecurityGroupIdInput() *string
	Primary() interface{}
	SetPrimary(val interface{})
	PrimaryInput() interface{}
	Tags() *map[string]*string
	SetTags(val *map[string]*string)
	TagsInput() *map[string]*string
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
	PutIpConfiguration(value interface{})
	ResetAcceleratedNetworkingEnabled()
	ResetAuxiliaryMode()
	ResetAuxiliarySku()
	ResetDnsServers()
	ResetIpForwardingEnabled()
	ResetNetworkSecurityGroupId()
	ResetPrimary()
	ResetTags()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktn.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for OrchestratedVirtualMachineScaleSetNetworkInterfaceOutputReference
type jsiiProxy_OrchestratedVirtualMachineScaleSetNetworkInterfaceOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_OrchestratedVirtualMachineScaleSetNetworkInterfaceOutputReference) AcceleratedNetworkingEnabled() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"acceleratedNetworkingEnabled",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OrchestratedVirtualMachineScaleSetNetworkInterfaceOutputReference) AcceleratedNetworkingEnabledInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"acceleratedNetworkingEnabledInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OrchestratedVirtualMachineScaleSetNetworkInterfaceOutputReference) AuxiliaryMode() *string {
	var returns *string
	_jsii_.Get(
		j,
		"auxiliaryMode",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OrchestratedVirtualMachineScaleSetNetworkInterfaceOutputReference) AuxiliaryModeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"auxiliaryModeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OrchestratedVirtualMachineScaleSetNetworkInterfaceOutputReference) AuxiliarySku() *string {
	var returns *string
	_jsii_.Get(
		j,
		"auxiliarySku",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OrchestratedVirtualMachineScaleSetNetworkInterfaceOutputReference) AuxiliarySkuInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"auxiliarySkuInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OrchestratedVirtualMachineScaleSetNetworkInterfaceOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OrchestratedVirtualMachineScaleSetNetworkInterfaceOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OrchestratedVirtualMachineScaleSetNetworkInterfaceOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OrchestratedVirtualMachineScaleSetNetworkInterfaceOutputReference) DnsServers() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dnsServers",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OrchestratedVirtualMachineScaleSetNetworkInterfaceOutputReference) DnsServersInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dnsServersInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OrchestratedVirtualMachineScaleSetNetworkInterfaceOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OrchestratedVirtualMachineScaleSetNetworkInterfaceOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OrchestratedVirtualMachineScaleSetNetworkInterfaceOutputReference) IpConfiguration() OrchestratedVirtualMachineScaleSetNetworkInterfaceIpConfigurationList {
	var returns OrchestratedVirtualMachineScaleSetNetworkInterfaceIpConfigurationList
	_jsii_.Get(
		j,
		"ipConfiguration",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OrchestratedVirtualMachineScaleSetNetworkInterfaceOutputReference) IpConfigurationInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"ipConfigurationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OrchestratedVirtualMachineScaleSetNetworkInterfaceOutputReference) IpForwardingEnabled() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"ipForwardingEnabled",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OrchestratedVirtualMachineScaleSetNetworkInterfaceOutputReference) IpForwardingEnabledInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"ipForwardingEnabledInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OrchestratedVirtualMachineScaleSetNetworkInterfaceOutputReference) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OrchestratedVirtualMachineScaleSetNetworkInterfaceOutputReference) NameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"nameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OrchestratedVirtualMachineScaleSetNetworkInterfaceOutputReference) NetworkSecurityGroupId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"networkSecurityGroupId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OrchestratedVirtualMachineScaleSetNetworkInterfaceOutputReference) NetworkSecurityGroupIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"networkSecurityGroupIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OrchestratedVirtualMachineScaleSetNetworkInterfaceOutputReference) Primary() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"primary",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OrchestratedVirtualMachineScaleSetNetworkInterfaceOutputReference) PrimaryInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"primaryInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OrchestratedVirtualMachineScaleSetNetworkInterfaceOutputReference) Tags() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"tags",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OrchestratedVirtualMachineScaleSetNetworkInterfaceOutputReference) TagsInput() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"tagsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OrchestratedVirtualMachineScaleSetNetworkInterfaceOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OrchestratedVirtualMachineScaleSetNetworkInterfaceOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewOrchestratedVirtualMachineScaleSetNetworkInterfaceOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) OrchestratedVirtualMachineScaleSetNetworkInterfaceOutputReference {
	_init_.Initialize()

	if err := validateNewOrchestratedVirtualMachineScaleSetNetworkInterfaceOutputReferenceParameters(terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet); err != nil {
		panic(err)
	}
	j := jsiiProxy_OrchestratedVirtualMachineScaleSetNetworkInterfaceOutputReference{}

	_jsii_.Create(
		"@cdktn/provider-azurerm.orchestratedVirtualMachineScaleSet.OrchestratedVirtualMachineScaleSetNetworkInterfaceOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		&j,
	)

	return &j
}

func NewOrchestratedVirtualMachineScaleSetNetworkInterfaceOutputReference_Override(o OrchestratedVirtualMachineScaleSetNetworkInterfaceOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/provider-azurerm.orchestratedVirtualMachineScaleSet.OrchestratedVirtualMachineScaleSetNetworkInterfaceOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		o,
	)
}

func (j *jsiiProxy_OrchestratedVirtualMachineScaleSetNetworkInterfaceOutputReference)SetAcceleratedNetworkingEnabled(val interface{}) {
	if err := j.validateSetAcceleratedNetworkingEnabledParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"acceleratedNetworkingEnabled",
		val,
	)
}

func (j *jsiiProxy_OrchestratedVirtualMachineScaleSetNetworkInterfaceOutputReference)SetAuxiliaryMode(val *string) {
	if err := j.validateSetAuxiliaryModeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"auxiliaryMode",
		val,
	)
}

func (j *jsiiProxy_OrchestratedVirtualMachineScaleSetNetworkInterfaceOutputReference)SetAuxiliarySku(val *string) {
	if err := j.validateSetAuxiliarySkuParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"auxiliarySku",
		val,
	)
}

func (j *jsiiProxy_OrchestratedVirtualMachineScaleSetNetworkInterfaceOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_OrchestratedVirtualMachineScaleSetNetworkInterfaceOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_OrchestratedVirtualMachineScaleSetNetworkInterfaceOutputReference)SetDnsServers(val *[]*string) {
	if err := j.validateSetDnsServersParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"dnsServers",
		val,
	)
}

func (j *jsiiProxy_OrchestratedVirtualMachineScaleSetNetworkInterfaceOutputReference)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_OrchestratedVirtualMachineScaleSetNetworkInterfaceOutputReference)SetIpForwardingEnabled(val interface{}) {
	if err := j.validateSetIpForwardingEnabledParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"ipForwardingEnabled",
		val,
	)
}

func (j *jsiiProxy_OrchestratedVirtualMachineScaleSetNetworkInterfaceOutputReference)SetName(val *string) {
	if err := j.validateSetNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"name",
		val,
	)
}

func (j *jsiiProxy_OrchestratedVirtualMachineScaleSetNetworkInterfaceOutputReference)SetNetworkSecurityGroupId(val *string) {
	if err := j.validateSetNetworkSecurityGroupIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"networkSecurityGroupId",
		val,
	)
}

func (j *jsiiProxy_OrchestratedVirtualMachineScaleSetNetworkInterfaceOutputReference)SetPrimary(val interface{}) {
	if err := j.validateSetPrimaryParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"primary",
		val,
	)
}

func (j *jsiiProxy_OrchestratedVirtualMachineScaleSetNetworkInterfaceOutputReference)SetTags(val *map[string]*string) {
	if err := j.validateSetTagsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"tags",
		val,
	)
}

func (j *jsiiProxy_OrchestratedVirtualMachineScaleSetNetworkInterfaceOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_OrchestratedVirtualMachineScaleSetNetworkInterfaceOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (o *jsiiProxy_OrchestratedVirtualMachineScaleSetNetworkInterfaceOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		o,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (o *jsiiProxy_OrchestratedVirtualMachineScaleSetNetworkInterfaceOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	if err := o.validateGetAnyMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]interface{}

	_jsii_.Invoke(
		o,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (o *jsiiProxy_OrchestratedVirtualMachineScaleSetNetworkInterfaceOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
	if err := o.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		o,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (o *jsiiProxy_OrchestratedVirtualMachineScaleSetNetworkInterfaceOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	if err := o.validateGetBooleanMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*bool

	_jsii_.Invoke(
		o,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (o *jsiiProxy_OrchestratedVirtualMachineScaleSetNetworkInterfaceOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	if err := o.validateGetListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*string

	_jsii_.Invoke(
		o,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (o *jsiiProxy_OrchestratedVirtualMachineScaleSetNetworkInterfaceOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	if err := o.validateGetNumberAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *float64

	_jsii_.Invoke(
		o,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (o *jsiiProxy_OrchestratedVirtualMachineScaleSetNetworkInterfaceOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	if err := o.validateGetNumberListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*float64

	_jsii_.Invoke(
		o,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (o *jsiiProxy_OrchestratedVirtualMachineScaleSetNetworkInterfaceOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	if err := o.validateGetNumberMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*float64

	_jsii_.Invoke(
		o,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (o *jsiiProxy_OrchestratedVirtualMachineScaleSetNetworkInterfaceOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	if err := o.validateGetStringAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		o,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (o *jsiiProxy_OrchestratedVirtualMachineScaleSetNetworkInterfaceOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	if err := o.validateGetStringMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*string

	_jsii_.Invoke(
		o,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (o *jsiiProxy_OrchestratedVirtualMachineScaleSetNetworkInterfaceOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		o,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (o *jsiiProxy_OrchestratedVirtualMachineScaleSetNetworkInterfaceOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
	if err := o.validateInterpolationForAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		o,
		"interpolationForAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (o *jsiiProxy_OrchestratedVirtualMachineScaleSetNetworkInterfaceOutputReference) PutIpConfiguration(value interface{}) {
	if err := o.validatePutIpConfigurationParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		o,
		"putIpConfiguration",
		[]interface{}{value},
	)
}

func (o *jsiiProxy_OrchestratedVirtualMachineScaleSetNetworkInterfaceOutputReference) ResetAcceleratedNetworkingEnabled() {
	_jsii_.InvokeVoid(
		o,
		"resetAcceleratedNetworkingEnabled",
		nil, // no parameters
	)
}

func (o *jsiiProxy_OrchestratedVirtualMachineScaleSetNetworkInterfaceOutputReference) ResetAuxiliaryMode() {
	_jsii_.InvokeVoid(
		o,
		"resetAuxiliaryMode",
		nil, // no parameters
	)
}

func (o *jsiiProxy_OrchestratedVirtualMachineScaleSetNetworkInterfaceOutputReference) ResetAuxiliarySku() {
	_jsii_.InvokeVoid(
		o,
		"resetAuxiliarySku",
		nil, // no parameters
	)
}

func (o *jsiiProxy_OrchestratedVirtualMachineScaleSetNetworkInterfaceOutputReference) ResetDnsServers() {
	_jsii_.InvokeVoid(
		o,
		"resetDnsServers",
		nil, // no parameters
	)
}

func (o *jsiiProxy_OrchestratedVirtualMachineScaleSetNetworkInterfaceOutputReference) ResetIpForwardingEnabled() {
	_jsii_.InvokeVoid(
		o,
		"resetIpForwardingEnabled",
		nil, // no parameters
	)
}

func (o *jsiiProxy_OrchestratedVirtualMachineScaleSetNetworkInterfaceOutputReference) ResetNetworkSecurityGroupId() {
	_jsii_.InvokeVoid(
		o,
		"resetNetworkSecurityGroupId",
		nil, // no parameters
	)
}

func (o *jsiiProxy_OrchestratedVirtualMachineScaleSetNetworkInterfaceOutputReference) ResetPrimary() {
	_jsii_.InvokeVoid(
		o,
		"resetPrimary",
		nil, // no parameters
	)
}

func (o *jsiiProxy_OrchestratedVirtualMachineScaleSetNetworkInterfaceOutputReference) ResetTags() {
	_jsii_.InvokeVoid(
		o,
		"resetTags",
		nil, // no parameters
	)
}

func (o *jsiiProxy_OrchestratedVirtualMachineScaleSetNetworkInterfaceOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
	if err := o.validateResolveParameters(context); err != nil {
		panic(err)
	}
	var returns interface{}

	_jsii_.Invoke(
		o,
		"resolve",
		[]interface{}{context},
		&returns,
	)

	return returns
}

func (o *jsiiProxy_OrchestratedVirtualMachineScaleSetNetworkInterfaceOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		o,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

