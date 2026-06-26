// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package netappvolumebucketwithserver

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-provider-azurerm-go/azurerm/v16/jsii"

	"github.com/cdktn-io/cdktn-provider-azurerm-go/azurerm/v16/netappvolumebucketwithserver/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type NetappVolumeBucketWithServerKeyVaultOutputReference interface {
	cdktn.ComplexObject
	CertificateKeyVaultUri() *string
	SetCertificateKeyVaultUri(val *string)
	CertificateKeyVaultUriInput() *string
	CertificateName() *string
	SetCertificateName(val *string)
	CertificateNameInput() *string
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
	CredentialsKeyVaultUri() *string
	SetCredentialsKeyVaultUri(val *string)
	CredentialsKeyVaultUriInput() *string
	CredentialsSecretName() *string
	SetCredentialsSecretName(val *string)
	CredentialsSecretNameInput() *string
	// Experimental.
	Fqn() *string
	InternalValue() *NetappVolumeBucketWithServerKeyVault
	SetInternalValue(val *NetappVolumeBucketWithServerKeyVault)
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
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktn.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for NetappVolumeBucketWithServerKeyVaultOutputReference
type jsiiProxy_NetappVolumeBucketWithServerKeyVaultOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_NetappVolumeBucketWithServerKeyVaultOutputReference) CertificateKeyVaultUri() *string {
	var returns *string
	_jsii_.Get(
		j,
		"certificateKeyVaultUri",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetappVolumeBucketWithServerKeyVaultOutputReference) CertificateKeyVaultUriInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"certificateKeyVaultUriInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetappVolumeBucketWithServerKeyVaultOutputReference) CertificateName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"certificateName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetappVolumeBucketWithServerKeyVaultOutputReference) CertificateNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"certificateNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetappVolumeBucketWithServerKeyVaultOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetappVolumeBucketWithServerKeyVaultOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetappVolumeBucketWithServerKeyVaultOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetappVolumeBucketWithServerKeyVaultOutputReference) CredentialsKeyVaultUri() *string {
	var returns *string
	_jsii_.Get(
		j,
		"credentialsKeyVaultUri",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetappVolumeBucketWithServerKeyVaultOutputReference) CredentialsKeyVaultUriInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"credentialsKeyVaultUriInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetappVolumeBucketWithServerKeyVaultOutputReference) CredentialsSecretName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"credentialsSecretName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetappVolumeBucketWithServerKeyVaultOutputReference) CredentialsSecretNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"credentialsSecretNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetappVolumeBucketWithServerKeyVaultOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetappVolumeBucketWithServerKeyVaultOutputReference) InternalValue() *NetappVolumeBucketWithServerKeyVault {
	var returns *NetappVolumeBucketWithServerKeyVault
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetappVolumeBucketWithServerKeyVaultOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetappVolumeBucketWithServerKeyVaultOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewNetappVolumeBucketWithServerKeyVaultOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) NetappVolumeBucketWithServerKeyVaultOutputReference {
	_init_.Initialize()

	if err := validateNewNetappVolumeBucketWithServerKeyVaultOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_NetappVolumeBucketWithServerKeyVaultOutputReference{}

	_jsii_.Create(
		"@cdktn/provider-azurerm.netappVolumeBucketWithServer.NetappVolumeBucketWithServerKeyVaultOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewNetappVolumeBucketWithServerKeyVaultOutputReference_Override(n NetappVolumeBucketWithServerKeyVaultOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/provider-azurerm.netappVolumeBucketWithServer.NetappVolumeBucketWithServerKeyVaultOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		n,
	)
}

func (j *jsiiProxy_NetappVolumeBucketWithServerKeyVaultOutputReference)SetCertificateKeyVaultUri(val *string) {
	if err := j.validateSetCertificateKeyVaultUriParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"certificateKeyVaultUri",
		val,
	)
}

func (j *jsiiProxy_NetappVolumeBucketWithServerKeyVaultOutputReference)SetCertificateName(val *string) {
	if err := j.validateSetCertificateNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"certificateName",
		val,
	)
}

func (j *jsiiProxy_NetappVolumeBucketWithServerKeyVaultOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_NetappVolumeBucketWithServerKeyVaultOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_NetappVolumeBucketWithServerKeyVaultOutputReference)SetCredentialsKeyVaultUri(val *string) {
	if err := j.validateSetCredentialsKeyVaultUriParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"credentialsKeyVaultUri",
		val,
	)
}

func (j *jsiiProxy_NetappVolumeBucketWithServerKeyVaultOutputReference)SetCredentialsSecretName(val *string) {
	if err := j.validateSetCredentialsSecretNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"credentialsSecretName",
		val,
	)
}

func (j *jsiiProxy_NetappVolumeBucketWithServerKeyVaultOutputReference)SetInternalValue(val *NetappVolumeBucketWithServerKeyVault) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_NetappVolumeBucketWithServerKeyVaultOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_NetappVolumeBucketWithServerKeyVaultOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (n *jsiiProxy_NetappVolumeBucketWithServerKeyVaultOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		n,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (n *jsiiProxy_NetappVolumeBucketWithServerKeyVaultOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	if err := n.validateGetAnyMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]interface{}

	_jsii_.Invoke(
		n,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (n *jsiiProxy_NetappVolumeBucketWithServerKeyVaultOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
	if err := n.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		n,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (n *jsiiProxy_NetappVolumeBucketWithServerKeyVaultOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	if err := n.validateGetBooleanMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*bool

	_jsii_.Invoke(
		n,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (n *jsiiProxy_NetappVolumeBucketWithServerKeyVaultOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	if err := n.validateGetListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*string

	_jsii_.Invoke(
		n,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (n *jsiiProxy_NetappVolumeBucketWithServerKeyVaultOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	if err := n.validateGetNumberAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *float64

	_jsii_.Invoke(
		n,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (n *jsiiProxy_NetappVolumeBucketWithServerKeyVaultOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	if err := n.validateGetNumberListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*float64

	_jsii_.Invoke(
		n,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (n *jsiiProxy_NetappVolumeBucketWithServerKeyVaultOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	if err := n.validateGetNumberMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*float64

	_jsii_.Invoke(
		n,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (n *jsiiProxy_NetappVolumeBucketWithServerKeyVaultOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	if err := n.validateGetStringAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		n,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (n *jsiiProxy_NetappVolumeBucketWithServerKeyVaultOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	if err := n.validateGetStringMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*string

	_jsii_.Invoke(
		n,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (n *jsiiProxy_NetappVolumeBucketWithServerKeyVaultOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		n,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (n *jsiiProxy_NetappVolumeBucketWithServerKeyVaultOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
	if err := n.validateInterpolationForAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		n,
		"interpolationForAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (n *jsiiProxy_NetappVolumeBucketWithServerKeyVaultOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
	if err := n.validateResolveParameters(context); err != nil {
		panic(err)
	}
	var returns interface{}

	_jsii_.Invoke(
		n,
		"resolve",
		[]interface{}{context},
		&returns,
	)

	return returns
}

func (n *jsiiProxy_NetappVolumeBucketWithServerKeyVaultOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		n,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

