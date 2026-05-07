// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package applicationgateway

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-provider-azurerm-go/azurerm/v16/jsii"

	"github.com/cdktn-io/cdktn-provider-azurerm-go/azurerm/v16/applicationgateway/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type ApplicationGatewayListenerOutputReference interface {
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
	FrontendIpConfigurationId() *string
	FrontendIpConfigurationName() *string
	SetFrontendIpConfigurationName(val *string)
	FrontendIpConfigurationNameInput() *string
	FrontendPortId() *string
	FrontendPortName() *string
	SetFrontendPortName(val *string)
	FrontendPortNameInput() *string
	HostNames() *[]*string
	SetHostNames(val *[]*string)
	HostNamesInput() *[]*string
	Id() *string
	InternalValue() interface{}
	SetInternalValue(val interface{})
	Name() *string
	SetName(val *string)
	NameInput() *string
	Protocol() *string
	SetProtocol(val *string)
	ProtocolInput() *string
	SslCertificateId() *string
	SslCertificateName() *string
	SetSslCertificateName(val *string)
	SslCertificateNameInput() *string
	SslProfileId() *string
	SslProfileName() *string
	SetSslProfileName(val *string)
	SslProfileNameInput() *string
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
	ResetHostNames()
	ResetSslCertificateName()
	ResetSslProfileName()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktn.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for ApplicationGatewayListenerOutputReference
type jsiiProxy_ApplicationGatewayListenerOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_ApplicationGatewayListenerOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApplicationGatewayListenerOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApplicationGatewayListenerOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApplicationGatewayListenerOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApplicationGatewayListenerOutputReference) FrontendIpConfigurationId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"frontendIpConfigurationId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApplicationGatewayListenerOutputReference) FrontendIpConfigurationName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"frontendIpConfigurationName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApplicationGatewayListenerOutputReference) FrontendIpConfigurationNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"frontendIpConfigurationNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApplicationGatewayListenerOutputReference) FrontendPortId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"frontendPortId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApplicationGatewayListenerOutputReference) FrontendPortName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"frontendPortName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApplicationGatewayListenerOutputReference) FrontendPortNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"frontendPortNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApplicationGatewayListenerOutputReference) HostNames() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"hostNames",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApplicationGatewayListenerOutputReference) HostNamesInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"hostNamesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApplicationGatewayListenerOutputReference) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApplicationGatewayListenerOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApplicationGatewayListenerOutputReference) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApplicationGatewayListenerOutputReference) NameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"nameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApplicationGatewayListenerOutputReference) Protocol() *string {
	var returns *string
	_jsii_.Get(
		j,
		"protocol",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApplicationGatewayListenerOutputReference) ProtocolInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"protocolInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApplicationGatewayListenerOutputReference) SslCertificateId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"sslCertificateId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApplicationGatewayListenerOutputReference) SslCertificateName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"sslCertificateName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApplicationGatewayListenerOutputReference) SslCertificateNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"sslCertificateNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApplicationGatewayListenerOutputReference) SslProfileId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"sslProfileId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApplicationGatewayListenerOutputReference) SslProfileName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"sslProfileName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApplicationGatewayListenerOutputReference) SslProfileNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"sslProfileNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApplicationGatewayListenerOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApplicationGatewayListenerOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewApplicationGatewayListenerOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) ApplicationGatewayListenerOutputReference {
	_init_.Initialize()

	if err := validateNewApplicationGatewayListenerOutputReferenceParameters(terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet); err != nil {
		panic(err)
	}
	j := jsiiProxy_ApplicationGatewayListenerOutputReference{}

	_jsii_.Create(
		"@cdktn/provider-azurerm.applicationGateway.ApplicationGatewayListenerOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		&j,
	)

	return &j
}

func NewApplicationGatewayListenerOutputReference_Override(a ApplicationGatewayListenerOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/provider-azurerm.applicationGateway.ApplicationGatewayListenerOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		a,
	)
}

func (j *jsiiProxy_ApplicationGatewayListenerOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_ApplicationGatewayListenerOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_ApplicationGatewayListenerOutputReference)SetFrontendIpConfigurationName(val *string) {
	if err := j.validateSetFrontendIpConfigurationNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"frontendIpConfigurationName",
		val,
	)
}

func (j *jsiiProxy_ApplicationGatewayListenerOutputReference)SetFrontendPortName(val *string) {
	if err := j.validateSetFrontendPortNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"frontendPortName",
		val,
	)
}

func (j *jsiiProxy_ApplicationGatewayListenerOutputReference)SetHostNames(val *[]*string) {
	if err := j.validateSetHostNamesParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"hostNames",
		val,
	)
}

func (j *jsiiProxy_ApplicationGatewayListenerOutputReference)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_ApplicationGatewayListenerOutputReference)SetName(val *string) {
	if err := j.validateSetNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"name",
		val,
	)
}

func (j *jsiiProxy_ApplicationGatewayListenerOutputReference)SetProtocol(val *string) {
	if err := j.validateSetProtocolParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"protocol",
		val,
	)
}

func (j *jsiiProxy_ApplicationGatewayListenerOutputReference)SetSslCertificateName(val *string) {
	if err := j.validateSetSslCertificateNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"sslCertificateName",
		val,
	)
}

func (j *jsiiProxy_ApplicationGatewayListenerOutputReference)SetSslProfileName(val *string) {
	if err := j.validateSetSslProfileNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"sslProfileName",
		val,
	)
}

func (j *jsiiProxy_ApplicationGatewayListenerOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_ApplicationGatewayListenerOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (a *jsiiProxy_ApplicationGatewayListenerOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_ApplicationGatewayListenerOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	if err := a.validateGetAnyMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]interface{}

	_jsii_.Invoke(
		a,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_ApplicationGatewayListenerOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
	if err := a.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		a,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_ApplicationGatewayListenerOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	if err := a.validateGetBooleanMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*bool

	_jsii_.Invoke(
		a,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_ApplicationGatewayListenerOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	if err := a.validateGetListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*string

	_jsii_.Invoke(
		a,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_ApplicationGatewayListenerOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	if err := a.validateGetNumberAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *float64

	_jsii_.Invoke(
		a,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_ApplicationGatewayListenerOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	if err := a.validateGetNumberListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*float64

	_jsii_.Invoke(
		a,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_ApplicationGatewayListenerOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	if err := a.validateGetNumberMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*float64

	_jsii_.Invoke(
		a,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_ApplicationGatewayListenerOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	if err := a.validateGetStringAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		a,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_ApplicationGatewayListenerOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	if err := a.validateGetStringMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*string

	_jsii_.Invoke(
		a,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_ApplicationGatewayListenerOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		a,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_ApplicationGatewayListenerOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
	if err := a.validateInterpolationForAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		a,
		"interpolationForAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_ApplicationGatewayListenerOutputReference) ResetHostNames() {
	_jsii_.InvokeVoid(
		a,
		"resetHostNames",
		nil, // no parameters
	)
}

func (a *jsiiProxy_ApplicationGatewayListenerOutputReference) ResetSslCertificateName() {
	_jsii_.InvokeVoid(
		a,
		"resetSslCertificateName",
		nil, // no parameters
	)
}

func (a *jsiiProxy_ApplicationGatewayListenerOutputReference) ResetSslProfileName() {
	_jsii_.InvokeVoid(
		a,
		"resetSslProfileName",
		nil, // no parameters
	)
}

func (a *jsiiProxy_ApplicationGatewayListenerOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
	if err := a.validateResolveParameters(context); err != nil {
		panic(err)
	}
	var returns interface{}

	_jsii_.Invoke(
		a,
		"resolve",
		[]interface{}{context},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_ApplicationGatewayListenerOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

