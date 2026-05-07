// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package paloaltonextgenerationfirewallvirtualhubstratacloudmanager

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-provider-azurerm-go/azurerm/v16/jsii"

	"github.com/cdktn-io/cdktn-provider-azurerm-go/azurerm/v16/paloaltonextgenerationfirewallvirtualhubstratacloudmanager/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type PaloAltoNextGenerationFirewallVirtualHubStrataCloudManagerNetworkProfileOutputReference interface {
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
	EgressNatIpAddresses() *[]*string
	EgressNatIpAddressIds() *[]*string
	SetEgressNatIpAddressIds(val *[]*string)
	EgressNatIpAddressIdsInput() *[]*string
	// Experimental.
	Fqn() *string
	InternalValue() *PaloAltoNextGenerationFirewallVirtualHubStrataCloudManagerNetworkProfile
	SetInternalValue(val *PaloAltoNextGenerationFirewallVirtualHubStrataCloudManagerNetworkProfile)
	IpOfTrustForUserDefinedRoutes() *string
	NetworkVirtualApplianceId() *string
	SetNetworkVirtualApplianceId(val *string)
	NetworkVirtualApplianceIdInput() *string
	PublicIpAddresses() *[]*string
	PublicIpAddressIds() *[]*string
	SetPublicIpAddressIds(val *[]*string)
	PublicIpAddressIdsInput() *[]*string
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktn.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktn.IInterpolatingParent)
	TrustedAddressRanges() *[]*string
	SetTrustedAddressRanges(val *[]*string)
	TrustedAddressRangesInput() *[]*string
	TrustedSubnetId() *string
	UntrustedSubnetId() *string
	VirtualHubId() *string
	SetVirtualHubId(val *string)
	VirtualHubIdInput() *string
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
	ResetEgressNatIpAddressIds()
	ResetTrustedAddressRanges()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktn.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for PaloAltoNextGenerationFirewallVirtualHubStrataCloudManagerNetworkProfileOutputReference
type jsiiProxy_PaloAltoNextGenerationFirewallVirtualHubStrataCloudManagerNetworkProfileOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_PaloAltoNextGenerationFirewallVirtualHubStrataCloudManagerNetworkProfileOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PaloAltoNextGenerationFirewallVirtualHubStrataCloudManagerNetworkProfileOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PaloAltoNextGenerationFirewallVirtualHubStrataCloudManagerNetworkProfileOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PaloAltoNextGenerationFirewallVirtualHubStrataCloudManagerNetworkProfileOutputReference) EgressNatIpAddresses() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"egressNatIpAddresses",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PaloAltoNextGenerationFirewallVirtualHubStrataCloudManagerNetworkProfileOutputReference) EgressNatIpAddressIds() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"egressNatIpAddressIds",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PaloAltoNextGenerationFirewallVirtualHubStrataCloudManagerNetworkProfileOutputReference) EgressNatIpAddressIdsInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"egressNatIpAddressIdsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PaloAltoNextGenerationFirewallVirtualHubStrataCloudManagerNetworkProfileOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PaloAltoNextGenerationFirewallVirtualHubStrataCloudManagerNetworkProfileOutputReference) InternalValue() *PaloAltoNextGenerationFirewallVirtualHubStrataCloudManagerNetworkProfile {
	var returns *PaloAltoNextGenerationFirewallVirtualHubStrataCloudManagerNetworkProfile
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PaloAltoNextGenerationFirewallVirtualHubStrataCloudManagerNetworkProfileOutputReference) IpOfTrustForUserDefinedRoutes() *string {
	var returns *string
	_jsii_.Get(
		j,
		"ipOfTrustForUserDefinedRoutes",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PaloAltoNextGenerationFirewallVirtualHubStrataCloudManagerNetworkProfileOutputReference) NetworkVirtualApplianceId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"networkVirtualApplianceId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PaloAltoNextGenerationFirewallVirtualHubStrataCloudManagerNetworkProfileOutputReference) NetworkVirtualApplianceIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"networkVirtualApplianceIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PaloAltoNextGenerationFirewallVirtualHubStrataCloudManagerNetworkProfileOutputReference) PublicIpAddresses() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"publicIpAddresses",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PaloAltoNextGenerationFirewallVirtualHubStrataCloudManagerNetworkProfileOutputReference) PublicIpAddressIds() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"publicIpAddressIds",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PaloAltoNextGenerationFirewallVirtualHubStrataCloudManagerNetworkProfileOutputReference) PublicIpAddressIdsInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"publicIpAddressIdsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PaloAltoNextGenerationFirewallVirtualHubStrataCloudManagerNetworkProfileOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PaloAltoNextGenerationFirewallVirtualHubStrataCloudManagerNetworkProfileOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PaloAltoNextGenerationFirewallVirtualHubStrataCloudManagerNetworkProfileOutputReference) TrustedAddressRanges() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"trustedAddressRanges",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PaloAltoNextGenerationFirewallVirtualHubStrataCloudManagerNetworkProfileOutputReference) TrustedAddressRangesInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"trustedAddressRangesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PaloAltoNextGenerationFirewallVirtualHubStrataCloudManagerNetworkProfileOutputReference) TrustedSubnetId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"trustedSubnetId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PaloAltoNextGenerationFirewallVirtualHubStrataCloudManagerNetworkProfileOutputReference) UntrustedSubnetId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"untrustedSubnetId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PaloAltoNextGenerationFirewallVirtualHubStrataCloudManagerNetworkProfileOutputReference) VirtualHubId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"virtualHubId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PaloAltoNextGenerationFirewallVirtualHubStrataCloudManagerNetworkProfileOutputReference) VirtualHubIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"virtualHubIdInput",
		&returns,
	)
	return returns
}


func NewPaloAltoNextGenerationFirewallVirtualHubStrataCloudManagerNetworkProfileOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) PaloAltoNextGenerationFirewallVirtualHubStrataCloudManagerNetworkProfileOutputReference {
	_init_.Initialize()

	if err := validateNewPaloAltoNextGenerationFirewallVirtualHubStrataCloudManagerNetworkProfileOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_PaloAltoNextGenerationFirewallVirtualHubStrataCloudManagerNetworkProfileOutputReference{}

	_jsii_.Create(
		"@cdktn/provider-azurerm.paloAltoNextGenerationFirewallVirtualHubStrataCloudManager.PaloAltoNextGenerationFirewallVirtualHubStrataCloudManagerNetworkProfileOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewPaloAltoNextGenerationFirewallVirtualHubStrataCloudManagerNetworkProfileOutputReference_Override(p PaloAltoNextGenerationFirewallVirtualHubStrataCloudManagerNetworkProfileOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/provider-azurerm.paloAltoNextGenerationFirewallVirtualHubStrataCloudManager.PaloAltoNextGenerationFirewallVirtualHubStrataCloudManagerNetworkProfileOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		p,
	)
}

func (j *jsiiProxy_PaloAltoNextGenerationFirewallVirtualHubStrataCloudManagerNetworkProfileOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_PaloAltoNextGenerationFirewallVirtualHubStrataCloudManagerNetworkProfileOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_PaloAltoNextGenerationFirewallVirtualHubStrataCloudManagerNetworkProfileOutputReference)SetEgressNatIpAddressIds(val *[]*string) {
	if err := j.validateSetEgressNatIpAddressIdsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"egressNatIpAddressIds",
		val,
	)
}

func (j *jsiiProxy_PaloAltoNextGenerationFirewallVirtualHubStrataCloudManagerNetworkProfileOutputReference)SetInternalValue(val *PaloAltoNextGenerationFirewallVirtualHubStrataCloudManagerNetworkProfile) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_PaloAltoNextGenerationFirewallVirtualHubStrataCloudManagerNetworkProfileOutputReference)SetNetworkVirtualApplianceId(val *string) {
	if err := j.validateSetNetworkVirtualApplianceIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"networkVirtualApplianceId",
		val,
	)
}

func (j *jsiiProxy_PaloAltoNextGenerationFirewallVirtualHubStrataCloudManagerNetworkProfileOutputReference)SetPublicIpAddressIds(val *[]*string) {
	if err := j.validateSetPublicIpAddressIdsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"publicIpAddressIds",
		val,
	)
}

func (j *jsiiProxy_PaloAltoNextGenerationFirewallVirtualHubStrataCloudManagerNetworkProfileOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_PaloAltoNextGenerationFirewallVirtualHubStrataCloudManagerNetworkProfileOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_PaloAltoNextGenerationFirewallVirtualHubStrataCloudManagerNetworkProfileOutputReference)SetTrustedAddressRanges(val *[]*string) {
	if err := j.validateSetTrustedAddressRangesParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"trustedAddressRanges",
		val,
	)
}

func (j *jsiiProxy_PaloAltoNextGenerationFirewallVirtualHubStrataCloudManagerNetworkProfileOutputReference)SetVirtualHubId(val *string) {
	if err := j.validateSetVirtualHubIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"virtualHubId",
		val,
	)
}

func (p *jsiiProxy_PaloAltoNextGenerationFirewallVirtualHubStrataCloudManagerNetworkProfileOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		p,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (p *jsiiProxy_PaloAltoNextGenerationFirewallVirtualHubStrataCloudManagerNetworkProfileOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	if err := p.validateGetAnyMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]interface{}

	_jsii_.Invoke(
		p,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (p *jsiiProxy_PaloAltoNextGenerationFirewallVirtualHubStrataCloudManagerNetworkProfileOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
	if err := p.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		p,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (p *jsiiProxy_PaloAltoNextGenerationFirewallVirtualHubStrataCloudManagerNetworkProfileOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	if err := p.validateGetBooleanMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*bool

	_jsii_.Invoke(
		p,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (p *jsiiProxy_PaloAltoNextGenerationFirewallVirtualHubStrataCloudManagerNetworkProfileOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	if err := p.validateGetListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*string

	_jsii_.Invoke(
		p,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (p *jsiiProxy_PaloAltoNextGenerationFirewallVirtualHubStrataCloudManagerNetworkProfileOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	if err := p.validateGetNumberAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *float64

	_jsii_.Invoke(
		p,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (p *jsiiProxy_PaloAltoNextGenerationFirewallVirtualHubStrataCloudManagerNetworkProfileOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	if err := p.validateGetNumberListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*float64

	_jsii_.Invoke(
		p,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (p *jsiiProxy_PaloAltoNextGenerationFirewallVirtualHubStrataCloudManagerNetworkProfileOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	if err := p.validateGetNumberMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*float64

	_jsii_.Invoke(
		p,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (p *jsiiProxy_PaloAltoNextGenerationFirewallVirtualHubStrataCloudManagerNetworkProfileOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	if err := p.validateGetStringAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		p,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (p *jsiiProxy_PaloAltoNextGenerationFirewallVirtualHubStrataCloudManagerNetworkProfileOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	if err := p.validateGetStringMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*string

	_jsii_.Invoke(
		p,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (p *jsiiProxy_PaloAltoNextGenerationFirewallVirtualHubStrataCloudManagerNetworkProfileOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		p,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (p *jsiiProxy_PaloAltoNextGenerationFirewallVirtualHubStrataCloudManagerNetworkProfileOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
	if err := p.validateInterpolationForAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		p,
		"interpolationForAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (p *jsiiProxy_PaloAltoNextGenerationFirewallVirtualHubStrataCloudManagerNetworkProfileOutputReference) ResetEgressNatIpAddressIds() {
	_jsii_.InvokeVoid(
		p,
		"resetEgressNatIpAddressIds",
		nil, // no parameters
	)
}

func (p *jsiiProxy_PaloAltoNextGenerationFirewallVirtualHubStrataCloudManagerNetworkProfileOutputReference) ResetTrustedAddressRanges() {
	_jsii_.InvokeVoid(
		p,
		"resetTrustedAddressRanges",
		nil, // no parameters
	)
}

func (p *jsiiProxy_PaloAltoNextGenerationFirewallVirtualHubStrataCloudManagerNetworkProfileOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
	if err := p.validateResolveParameters(context); err != nil {
		panic(err)
	}
	var returns interface{}

	_jsii_.Invoke(
		p,
		"resolve",
		[]interface{}{context},
		&returns,
	)

	return returns
}

func (p *jsiiProxy_PaloAltoNextGenerationFirewallVirtualHubStrataCloudManagerNetworkProfileOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		p,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

