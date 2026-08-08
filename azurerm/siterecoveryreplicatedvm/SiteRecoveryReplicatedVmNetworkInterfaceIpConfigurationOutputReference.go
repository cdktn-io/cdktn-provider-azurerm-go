// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package siterecoveryreplicatedvm

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-provider-azurerm-go/azurerm/v18/jsii"

	"github.com/cdktn-io/cdktn-provider-azurerm-go/azurerm/v18/siterecoveryreplicatedvm/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type SiteRecoveryReplicatedVmNetworkInterfaceIpConfigurationOutputReference interface {
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
	FailoverTestPublicIpAddressId() *string
	SetFailoverTestPublicIpAddressId(val *string)
	FailoverTestPublicIpAddressIdInput() *string
	FailoverTestStaticIp() *string
	SetFailoverTestStaticIp(val *string)
	FailoverTestStaticIpInput() *string
	FailoverTestSubnetName() *string
	SetFailoverTestSubnetName(val *string)
	FailoverTestSubnetNameInput() *string
	// Experimental.
	Fqn() *string
	InternalValue() interface{}
	SetInternalValue(val interface{})
	Name() *string
	SetName(val *string)
	NameInput() *string
	Primary() interface{}
	SetPrimary(val interface{})
	PrimaryInput() interface{}
	RecoveryLoadBalancerBackendAddressPoolIds() *[]*string
	SetRecoveryLoadBalancerBackendAddressPoolIds(val *[]*string)
	RecoveryLoadBalancerBackendAddressPoolIdsInput() *[]*string
	RecoveryPublicIpAddressId() *string
	SetRecoveryPublicIpAddressId(val *string)
	RecoveryPublicIpAddressIdInput() *string
	TargetStaticIp() *string
	SetTargetStaticIp(val *string)
	TargetStaticIpInput() *string
	TargetSubnetName() *string
	SetTargetSubnetName(val *string)
	TargetSubnetNameInput() *string
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
	ResetFailoverTestPublicIpAddressId()
	ResetFailoverTestStaticIp()
	ResetFailoverTestSubnetName()
	ResetName()
	ResetPrimary()
	ResetRecoveryLoadBalancerBackendAddressPoolIds()
	ResetRecoveryPublicIpAddressId()
	ResetTargetStaticIp()
	ResetTargetSubnetName()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktn.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for SiteRecoveryReplicatedVmNetworkInterfaceIpConfigurationOutputReference
type jsiiProxy_SiteRecoveryReplicatedVmNetworkInterfaceIpConfigurationOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_SiteRecoveryReplicatedVmNetworkInterfaceIpConfigurationOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SiteRecoveryReplicatedVmNetworkInterfaceIpConfigurationOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SiteRecoveryReplicatedVmNetworkInterfaceIpConfigurationOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SiteRecoveryReplicatedVmNetworkInterfaceIpConfigurationOutputReference) FailoverTestPublicIpAddressId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"failoverTestPublicIpAddressId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SiteRecoveryReplicatedVmNetworkInterfaceIpConfigurationOutputReference) FailoverTestPublicIpAddressIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"failoverTestPublicIpAddressIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SiteRecoveryReplicatedVmNetworkInterfaceIpConfigurationOutputReference) FailoverTestStaticIp() *string {
	var returns *string
	_jsii_.Get(
		j,
		"failoverTestStaticIp",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SiteRecoveryReplicatedVmNetworkInterfaceIpConfigurationOutputReference) FailoverTestStaticIpInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"failoverTestStaticIpInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SiteRecoveryReplicatedVmNetworkInterfaceIpConfigurationOutputReference) FailoverTestSubnetName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"failoverTestSubnetName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SiteRecoveryReplicatedVmNetworkInterfaceIpConfigurationOutputReference) FailoverTestSubnetNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"failoverTestSubnetNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SiteRecoveryReplicatedVmNetworkInterfaceIpConfigurationOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SiteRecoveryReplicatedVmNetworkInterfaceIpConfigurationOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SiteRecoveryReplicatedVmNetworkInterfaceIpConfigurationOutputReference) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SiteRecoveryReplicatedVmNetworkInterfaceIpConfigurationOutputReference) NameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"nameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SiteRecoveryReplicatedVmNetworkInterfaceIpConfigurationOutputReference) Primary() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"primary",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SiteRecoveryReplicatedVmNetworkInterfaceIpConfigurationOutputReference) PrimaryInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"primaryInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SiteRecoveryReplicatedVmNetworkInterfaceIpConfigurationOutputReference) RecoveryLoadBalancerBackendAddressPoolIds() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"recoveryLoadBalancerBackendAddressPoolIds",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SiteRecoveryReplicatedVmNetworkInterfaceIpConfigurationOutputReference) RecoveryLoadBalancerBackendAddressPoolIdsInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"recoveryLoadBalancerBackendAddressPoolIdsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SiteRecoveryReplicatedVmNetworkInterfaceIpConfigurationOutputReference) RecoveryPublicIpAddressId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"recoveryPublicIpAddressId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SiteRecoveryReplicatedVmNetworkInterfaceIpConfigurationOutputReference) RecoveryPublicIpAddressIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"recoveryPublicIpAddressIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SiteRecoveryReplicatedVmNetworkInterfaceIpConfigurationOutputReference) TargetStaticIp() *string {
	var returns *string
	_jsii_.Get(
		j,
		"targetStaticIp",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SiteRecoveryReplicatedVmNetworkInterfaceIpConfigurationOutputReference) TargetStaticIpInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"targetStaticIpInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SiteRecoveryReplicatedVmNetworkInterfaceIpConfigurationOutputReference) TargetSubnetName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"targetSubnetName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SiteRecoveryReplicatedVmNetworkInterfaceIpConfigurationOutputReference) TargetSubnetNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"targetSubnetNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SiteRecoveryReplicatedVmNetworkInterfaceIpConfigurationOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SiteRecoveryReplicatedVmNetworkInterfaceIpConfigurationOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewSiteRecoveryReplicatedVmNetworkInterfaceIpConfigurationOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) SiteRecoveryReplicatedVmNetworkInterfaceIpConfigurationOutputReference {
	_init_.Initialize()

	if err := validateNewSiteRecoveryReplicatedVmNetworkInterfaceIpConfigurationOutputReferenceParameters(terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet); err != nil {
		panic(err)
	}
	j := jsiiProxy_SiteRecoveryReplicatedVmNetworkInterfaceIpConfigurationOutputReference{}

	_jsii_.Create(
		"@cdktn/provider-azurerm.siteRecoveryReplicatedVm.SiteRecoveryReplicatedVmNetworkInterfaceIpConfigurationOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		&j,
	)

	return &j
}

func NewSiteRecoveryReplicatedVmNetworkInterfaceIpConfigurationOutputReference_Override(s SiteRecoveryReplicatedVmNetworkInterfaceIpConfigurationOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/provider-azurerm.siteRecoveryReplicatedVm.SiteRecoveryReplicatedVmNetworkInterfaceIpConfigurationOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		s,
	)
}

func (j *jsiiProxy_SiteRecoveryReplicatedVmNetworkInterfaceIpConfigurationOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_SiteRecoveryReplicatedVmNetworkInterfaceIpConfigurationOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_SiteRecoveryReplicatedVmNetworkInterfaceIpConfigurationOutputReference)SetFailoverTestPublicIpAddressId(val *string) {
	if err := j.validateSetFailoverTestPublicIpAddressIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"failoverTestPublicIpAddressId",
		val,
	)
}

func (j *jsiiProxy_SiteRecoveryReplicatedVmNetworkInterfaceIpConfigurationOutputReference)SetFailoverTestStaticIp(val *string) {
	if err := j.validateSetFailoverTestStaticIpParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"failoverTestStaticIp",
		val,
	)
}

func (j *jsiiProxy_SiteRecoveryReplicatedVmNetworkInterfaceIpConfigurationOutputReference)SetFailoverTestSubnetName(val *string) {
	if err := j.validateSetFailoverTestSubnetNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"failoverTestSubnetName",
		val,
	)
}

func (j *jsiiProxy_SiteRecoveryReplicatedVmNetworkInterfaceIpConfigurationOutputReference)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_SiteRecoveryReplicatedVmNetworkInterfaceIpConfigurationOutputReference)SetName(val *string) {
	if err := j.validateSetNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"name",
		val,
	)
}

func (j *jsiiProxy_SiteRecoveryReplicatedVmNetworkInterfaceIpConfigurationOutputReference)SetPrimary(val interface{}) {
	if err := j.validateSetPrimaryParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"primary",
		val,
	)
}

func (j *jsiiProxy_SiteRecoveryReplicatedVmNetworkInterfaceIpConfigurationOutputReference)SetRecoveryLoadBalancerBackendAddressPoolIds(val *[]*string) {
	if err := j.validateSetRecoveryLoadBalancerBackendAddressPoolIdsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"recoveryLoadBalancerBackendAddressPoolIds",
		val,
	)
}

func (j *jsiiProxy_SiteRecoveryReplicatedVmNetworkInterfaceIpConfigurationOutputReference)SetRecoveryPublicIpAddressId(val *string) {
	if err := j.validateSetRecoveryPublicIpAddressIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"recoveryPublicIpAddressId",
		val,
	)
}

func (j *jsiiProxy_SiteRecoveryReplicatedVmNetworkInterfaceIpConfigurationOutputReference)SetTargetStaticIp(val *string) {
	if err := j.validateSetTargetStaticIpParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"targetStaticIp",
		val,
	)
}

func (j *jsiiProxy_SiteRecoveryReplicatedVmNetworkInterfaceIpConfigurationOutputReference)SetTargetSubnetName(val *string) {
	if err := j.validateSetTargetSubnetNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"targetSubnetName",
		val,
	)
}

func (j *jsiiProxy_SiteRecoveryReplicatedVmNetworkInterfaceIpConfigurationOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_SiteRecoveryReplicatedVmNetworkInterfaceIpConfigurationOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (s *jsiiProxy_SiteRecoveryReplicatedVmNetworkInterfaceIpConfigurationOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		s,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SiteRecoveryReplicatedVmNetworkInterfaceIpConfigurationOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	if err := s.validateGetAnyMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]interface{}

	_jsii_.Invoke(
		s,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SiteRecoveryReplicatedVmNetworkInterfaceIpConfigurationOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
	if err := s.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		s,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SiteRecoveryReplicatedVmNetworkInterfaceIpConfigurationOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	if err := s.validateGetBooleanMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*bool

	_jsii_.Invoke(
		s,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SiteRecoveryReplicatedVmNetworkInterfaceIpConfigurationOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	if err := s.validateGetListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*string

	_jsii_.Invoke(
		s,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SiteRecoveryReplicatedVmNetworkInterfaceIpConfigurationOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	if err := s.validateGetNumberAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *float64

	_jsii_.Invoke(
		s,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SiteRecoveryReplicatedVmNetworkInterfaceIpConfigurationOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	if err := s.validateGetNumberListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*float64

	_jsii_.Invoke(
		s,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SiteRecoveryReplicatedVmNetworkInterfaceIpConfigurationOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	if err := s.validateGetNumberMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*float64

	_jsii_.Invoke(
		s,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SiteRecoveryReplicatedVmNetworkInterfaceIpConfigurationOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	if err := s.validateGetStringAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		s,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SiteRecoveryReplicatedVmNetworkInterfaceIpConfigurationOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	if err := s.validateGetStringMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*string

	_jsii_.Invoke(
		s,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SiteRecoveryReplicatedVmNetworkInterfaceIpConfigurationOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		s,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SiteRecoveryReplicatedVmNetworkInterfaceIpConfigurationOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
	if err := s.validateInterpolationForAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		s,
		"interpolationForAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SiteRecoveryReplicatedVmNetworkInterfaceIpConfigurationOutputReference) ResetFailoverTestPublicIpAddressId() {
	_jsii_.InvokeVoid(
		s,
		"resetFailoverTestPublicIpAddressId",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SiteRecoveryReplicatedVmNetworkInterfaceIpConfigurationOutputReference) ResetFailoverTestStaticIp() {
	_jsii_.InvokeVoid(
		s,
		"resetFailoverTestStaticIp",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SiteRecoveryReplicatedVmNetworkInterfaceIpConfigurationOutputReference) ResetFailoverTestSubnetName() {
	_jsii_.InvokeVoid(
		s,
		"resetFailoverTestSubnetName",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SiteRecoveryReplicatedVmNetworkInterfaceIpConfigurationOutputReference) ResetName() {
	_jsii_.InvokeVoid(
		s,
		"resetName",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SiteRecoveryReplicatedVmNetworkInterfaceIpConfigurationOutputReference) ResetPrimary() {
	_jsii_.InvokeVoid(
		s,
		"resetPrimary",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SiteRecoveryReplicatedVmNetworkInterfaceIpConfigurationOutputReference) ResetRecoveryLoadBalancerBackendAddressPoolIds() {
	_jsii_.InvokeVoid(
		s,
		"resetRecoveryLoadBalancerBackendAddressPoolIds",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SiteRecoveryReplicatedVmNetworkInterfaceIpConfigurationOutputReference) ResetRecoveryPublicIpAddressId() {
	_jsii_.InvokeVoid(
		s,
		"resetRecoveryPublicIpAddressId",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SiteRecoveryReplicatedVmNetworkInterfaceIpConfigurationOutputReference) ResetTargetStaticIp() {
	_jsii_.InvokeVoid(
		s,
		"resetTargetStaticIp",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SiteRecoveryReplicatedVmNetworkInterfaceIpConfigurationOutputReference) ResetTargetSubnetName() {
	_jsii_.InvokeVoid(
		s,
		"resetTargetSubnetName",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SiteRecoveryReplicatedVmNetworkInterfaceIpConfigurationOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
	if err := s.validateResolveParameters(context); err != nil {
		panic(err)
	}
	var returns interface{}

	_jsii_.Invoke(
		s,
		"resolve",
		[]interface{}{context},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SiteRecoveryReplicatedVmNetworkInterfaceIpConfigurationOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		s,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

