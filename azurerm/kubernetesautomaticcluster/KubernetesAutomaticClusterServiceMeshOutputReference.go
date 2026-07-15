// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package kubernetesautomaticcluster

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-provider-azurerm-go/azurerm/v16/jsii"

	"github.com/cdktn-io/cdktn-provider-azurerm-go/azurerm/v16/kubernetesautomaticcluster/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type KubernetesAutomaticClusterServiceMeshOutputReference interface {
	cdktn.ComplexObject
	CertificateAuthority() KubernetesAutomaticClusterServiceMeshCertificateAuthorityOutputReference
	CertificateAuthorityInput() *KubernetesAutomaticClusterServiceMeshCertificateAuthority
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
	ExternalIngressGatewayEnabled() interface{}
	SetExternalIngressGatewayEnabled(val interface{})
	ExternalIngressGatewayEnabledInput() interface{}
	// Experimental.
	Fqn() *string
	InternalIngressGatewayEnabled() interface{}
	SetInternalIngressGatewayEnabled(val interface{})
	InternalIngressGatewayEnabledInput() interface{}
	InternalValue() *KubernetesAutomaticClusterServiceMesh
	SetInternalValue(val *KubernetesAutomaticClusterServiceMesh)
	ProxyRedirectMechanism() *string
	SetProxyRedirectMechanism(val *string)
	ProxyRedirectMechanismInput() *string
	Revisions() *[]*string
	SetRevisions(val *[]*string)
	RevisionsInput() *[]*string
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
	PutCertificateAuthority(value *KubernetesAutomaticClusterServiceMeshCertificateAuthority)
	ResetCertificateAuthority()
	ResetExternalIngressGatewayEnabled()
	ResetInternalIngressGatewayEnabled()
	ResetProxyRedirectMechanism()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktn.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for KubernetesAutomaticClusterServiceMeshOutputReference
type jsiiProxy_KubernetesAutomaticClusterServiceMeshOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_KubernetesAutomaticClusterServiceMeshOutputReference) CertificateAuthority() KubernetesAutomaticClusterServiceMeshCertificateAuthorityOutputReference {
	var returns KubernetesAutomaticClusterServiceMeshCertificateAuthorityOutputReference
	_jsii_.Get(
		j,
		"certificateAuthority",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KubernetesAutomaticClusterServiceMeshOutputReference) CertificateAuthorityInput() *KubernetesAutomaticClusterServiceMeshCertificateAuthority {
	var returns *KubernetesAutomaticClusterServiceMeshCertificateAuthority
	_jsii_.Get(
		j,
		"certificateAuthorityInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KubernetesAutomaticClusterServiceMeshOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KubernetesAutomaticClusterServiceMeshOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KubernetesAutomaticClusterServiceMeshOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KubernetesAutomaticClusterServiceMeshOutputReference) ExternalIngressGatewayEnabled() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"externalIngressGatewayEnabled",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KubernetesAutomaticClusterServiceMeshOutputReference) ExternalIngressGatewayEnabledInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"externalIngressGatewayEnabledInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KubernetesAutomaticClusterServiceMeshOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KubernetesAutomaticClusterServiceMeshOutputReference) InternalIngressGatewayEnabled() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalIngressGatewayEnabled",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KubernetesAutomaticClusterServiceMeshOutputReference) InternalIngressGatewayEnabledInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalIngressGatewayEnabledInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KubernetesAutomaticClusterServiceMeshOutputReference) InternalValue() *KubernetesAutomaticClusterServiceMesh {
	var returns *KubernetesAutomaticClusterServiceMesh
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KubernetesAutomaticClusterServiceMeshOutputReference) ProxyRedirectMechanism() *string {
	var returns *string
	_jsii_.Get(
		j,
		"proxyRedirectMechanism",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KubernetesAutomaticClusterServiceMeshOutputReference) ProxyRedirectMechanismInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"proxyRedirectMechanismInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KubernetesAutomaticClusterServiceMeshOutputReference) Revisions() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"revisions",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KubernetesAutomaticClusterServiceMeshOutputReference) RevisionsInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"revisionsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KubernetesAutomaticClusterServiceMeshOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KubernetesAutomaticClusterServiceMeshOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewKubernetesAutomaticClusterServiceMeshOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) KubernetesAutomaticClusterServiceMeshOutputReference {
	_init_.Initialize()

	if err := validateNewKubernetesAutomaticClusterServiceMeshOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_KubernetesAutomaticClusterServiceMeshOutputReference{}

	_jsii_.Create(
		"@cdktn/provider-azurerm.kubernetesAutomaticCluster.KubernetesAutomaticClusterServiceMeshOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewKubernetesAutomaticClusterServiceMeshOutputReference_Override(k KubernetesAutomaticClusterServiceMeshOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/provider-azurerm.kubernetesAutomaticCluster.KubernetesAutomaticClusterServiceMeshOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		k,
	)
}

func (j *jsiiProxy_KubernetesAutomaticClusterServiceMeshOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_KubernetesAutomaticClusterServiceMeshOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_KubernetesAutomaticClusterServiceMeshOutputReference)SetExternalIngressGatewayEnabled(val interface{}) {
	if err := j.validateSetExternalIngressGatewayEnabledParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"externalIngressGatewayEnabled",
		val,
	)
}

func (j *jsiiProxy_KubernetesAutomaticClusterServiceMeshOutputReference)SetInternalIngressGatewayEnabled(val interface{}) {
	if err := j.validateSetInternalIngressGatewayEnabledParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalIngressGatewayEnabled",
		val,
	)
}

func (j *jsiiProxy_KubernetesAutomaticClusterServiceMeshOutputReference)SetInternalValue(val *KubernetesAutomaticClusterServiceMesh) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_KubernetesAutomaticClusterServiceMeshOutputReference)SetProxyRedirectMechanism(val *string) {
	if err := j.validateSetProxyRedirectMechanismParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"proxyRedirectMechanism",
		val,
	)
}

func (j *jsiiProxy_KubernetesAutomaticClusterServiceMeshOutputReference)SetRevisions(val *[]*string) {
	if err := j.validateSetRevisionsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"revisions",
		val,
	)
}

func (j *jsiiProxy_KubernetesAutomaticClusterServiceMeshOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_KubernetesAutomaticClusterServiceMeshOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (k *jsiiProxy_KubernetesAutomaticClusterServiceMeshOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		k,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (k *jsiiProxy_KubernetesAutomaticClusterServiceMeshOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	if err := k.validateGetAnyMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]interface{}

	_jsii_.Invoke(
		k,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (k *jsiiProxy_KubernetesAutomaticClusterServiceMeshOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
	if err := k.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		k,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (k *jsiiProxy_KubernetesAutomaticClusterServiceMeshOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	if err := k.validateGetBooleanMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*bool

	_jsii_.Invoke(
		k,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (k *jsiiProxy_KubernetesAutomaticClusterServiceMeshOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	if err := k.validateGetListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*string

	_jsii_.Invoke(
		k,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (k *jsiiProxy_KubernetesAutomaticClusterServiceMeshOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	if err := k.validateGetNumberAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *float64

	_jsii_.Invoke(
		k,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (k *jsiiProxy_KubernetesAutomaticClusterServiceMeshOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	if err := k.validateGetNumberListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*float64

	_jsii_.Invoke(
		k,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (k *jsiiProxy_KubernetesAutomaticClusterServiceMeshOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	if err := k.validateGetNumberMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*float64

	_jsii_.Invoke(
		k,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (k *jsiiProxy_KubernetesAutomaticClusterServiceMeshOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	if err := k.validateGetStringAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		k,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (k *jsiiProxy_KubernetesAutomaticClusterServiceMeshOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	if err := k.validateGetStringMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*string

	_jsii_.Invoke(
		k,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (k *jsiiProxy_KubernetesAutomaticClusterServiceMeshOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		k,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (k *jsiiProxy_KubernetesAutomaticClusterServiceMeshOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
	if err := k.validateInterpolationForAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		k,
		"interpolationForAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (k *jsiiProxy_KubernetesAutomaticClusterServiceMeshOutputReference) PutCertificateAuthority(value *KubernetesAutomaticClusterServiceMeshCertificateAuthority) {
	if err := k.validatePutCertificateAuthorityParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		k,
		"putCertificateAuthority",
		[]interface{}{value},
	)
}

func (k *jsiiProxy_KubernetesAutomaticClusterServiceMeshOutputReference) ResetCertificateAuthority() {
	_jsii_.InvokeVoid(
		k,
		"resetCertificateAuthority",
		nil, // no parameters
	)
}

func (k *jsiiProxy_KubernetesAutomaticClusterServiceMeshOutputReference) ResetExternalIngressGatewayEnabled() {
	_jsii_.InvokeVoid(
		k,
		"resetExternalIngressGatewayEnabled",
		nil, // no parameters
	)
}

func (k *jsiiProxy_KubernetesAutomaticClusterServiceMeshOutputReference) ResetInternalIngressGatewayEnabled() {
	_jsii_.InvokeVoid(
		k,
		"resetInternalIngressGatewayEnabled",
		nil, // no parameters
	)
}

func (k *jsiiProxy_KubernetesAutomaticClusterServiceMeshOutputReference) ResetProxyRedirectMechanism() {
	_jsii_.InvokeVoid(
		k,
		"resetProxyRedirectMechanism",
		nil, // no parameters
	)
}

func (k *jsiiProxy_KubernetesAutomaticClusterServiceMeshOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
	if err := k.validateResolveParameters(context); err != nil {
		panic(err)
	}
	var returns interface{}

	_jsii_.Invoke(
		k,
		"resolve",
		[]interface{}{context},
		&returns,
	)

	return returns
}

func (k *jsiiProxy_KubernetesAutomaticClusterServiceMeshOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		k,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

