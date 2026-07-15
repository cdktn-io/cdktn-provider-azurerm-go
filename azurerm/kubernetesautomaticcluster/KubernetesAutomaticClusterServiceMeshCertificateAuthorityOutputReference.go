// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package kubernetesautomaticcluster

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-provider-azurerm-go/azurerm/v16/jsii"

	"github.com/cdktn-io/cdktn-provider-azurerm-go/azurerm/v16/kubernetesautomaticcluster/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type KubernetesAutomaticClusterServiceMeshCertificateAuthorityOutputReference interface {
	cdktn.ComplexObject
	CertificateChainObjectName() *string
	SetCertificateChainObjectName(val *string)
	CertificateChainObjectNameInput() *string
	CertificateObjectName() *string
	SetCertificateObjectName(val *string)
	CertificateObjectNameInput() *string
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
	InternalValue() *KubernetesAutomaticClusterServiceMeshCertificateAuthority
	SetInternalValue(val *KubernetesAutomaticClusterServiceMeshCertificateAuthority)
	KeyObjectName() *string
	SetKeyObjectName(val *string)
	KeyObjectNameInput() *string
	KeyVaultId() *string
	SetKeyVaultId(val *string)
	KeyVaultIdInput() *string
	RootCertificateObjectName() *string
	SetRootCertificateObjectName(val *string)
	RootCertificateObjectNameInput() *string
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

// The jsii proxy struct for KubernetesAutomaticClusterServiceMeshCertificateAuthorityOutputReference
type jsiiProxy_KubernetesAutomaticClusterServiceMeshCertificateAuthorityOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_KubernetesAutomaticClusterServiceMeshCertificateAuthorityOutputReference) CertificateChainObjectName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"certificateChainObjectName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KubernetesAutomaticClusterServiceMeshCertificateAuthorityOutputReference) CertificateChainObjectNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"certificateChainObjectNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KubernetesAutomaticClusterServiceMeshCertificateAuthorityOutputReference) CertificateObjectName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"certificateObjectName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KubernetesAutomaticClusterServiceMeshCertificateAuthorityOutputReference) CertificateObjectNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"certificateObjectNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KubernetesAutomaticClusterServiceMeshCertificateAuthorityOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KubernetesAutomaticClusterServiceMeshCertificateAuthorityOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KubernetesAutomaticClusterServiceMeshCertificateAuthorityOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KubernetesAutomaticClusterServiceMeshCertificateAuthorityOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KubernetesAutomaticClusterServiceMeshCertificateAuthorityOutputReference) InternalValue() *KubernetesAutomaticClusterServiceMeshCertificateAuthority {
	var returns *KubernetesAutomaticClusterServiceMeshCertificateAuthority
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KubernetesAutomaticClusterServiceMeshCertificateAuthorityOutputReference) KeyObjectName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"keyObjectName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KubernetesAutomaticClusterServiceMeshCertificateAuthorityOutputReference) KeyObjectNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"keyObjectNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KubernetesAutomaticClusterServiceMeshCertificateAuthorityOutputReference) KeyVaultId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"keyVaultId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KubernetesAutomaticClusterServiceMeshCertificateAuthorityOutputReference) KeyVaultIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"keyVaultIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KubernetesAutomaticClusterServiceMeshCertificateAuthorityOutputReference) RootCertificateObjectName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"rootCertificateObjectName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KubernetesAutomaticClusterServiceMeshCertificateAuthorityOutputReference) RootCertificateObjectNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"rootCertificateObjectNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KubernetesAutomaticClusterServiceMeshCertificateAuthorityOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KubernetesAutomaticClusterServiceMeshCertificateAuthorityOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewKubernetesAutomaticClusterServiceMeshCertificateAuthorityOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) KubernetesAutomaticClusterServiceMeshCertificateAuthorityOutputReference {
	_init_.Initialize()

	if err := validateNewKubernetesAutomaticClusterServiceMeshCertificateAuthorityOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_KubernetesAutomaticClusterServiceMeshCertificateAuthorityOutputReference{}

	_jsii_.Create(
		"@cdktn/provider-azurerm.kubernetesAutomaticCluster.KubernetesAutomaticClusterServiceMeshCertificateAuthorityOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewKubernetesAutomaticClusterServiceMeshCertificateAuthorityOutputReference_Override(k KubernetesAutomaticClusterServiceMeshCertificateAuthorityOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/provider-azurerm.kubernetesAutomaticCluster.KubernetesAutomaticClusterServiceMeshCertificateAuthorityOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		k,
	)
}

func (j *jsiiProxy_KubernetesAutomaticClusterServiceMeshCertificateAuthorityOutputReference)SetCertificateChainObjectName(val *string) {
	if err := j.validateSetCertificateChainObjectNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"certificateChainObjectName",
		val,
	)
}

func (j *jsiiProxy_KubernetesAutomaticClusterServiceMeshCertificateAuthorityOutputReference)SetCertificateObjectName(val *string) {
	if err := j.validateSetCertificateObjectNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"certificateObjectName",
		val,
	)
}

func (j *jsiiProxy_KubernetesAutomaticClusterServiceMeshCertificateAuthorityOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_KubernetesAutomaticClusterServiceMeshCertificateAuthorityOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_KubernetesAutomaticClusterServiceMeshCertificateAuthorityOutputReference)SetInternalValue(val *KubernetesAutomaticClusterServiceMeshCertificateAuthority) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_KubernetesAutomaticClusterServiceMeshCertificateAuthorityOutputReference)SetKeyObjectName(val *string) {
	if err := j.validateSetKeyObjectNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"keyObjectName",
		val,
	)
}

func (j *jsiiProxy_KubernetesAutomaticClusterServiceMeshCertificateAuthorityOutputReference)SetKeyVaultId(val *string) {
	if err := j.validateSetKeyVaultIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"keyVaultId",
		val,
	)
}

func (j *jsiiProxy_KubernetesAutomaticClusterServiceMeshCertificateAuthorityOutputReference)SetRootCertificateObjectName(val *string) {
	if err := j.validateSetRootCertificateObjectNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"rootCertificateObjectName",
		val,
	)
}

func (j *jsiiProxy_KubernetesAutomaticClusterServiceMeshCertificateAuthorityOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_KubernetesAutomaticClusterServiceMeshCertificateAuthorityOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (k *jsiiProxy_KubernetesAutomaticClusterServiceMeshCertificateAuthorityOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		k,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (k *jsiiProxy_KubernetesAutomaticClusterServiceMeshCertificateAuthorityOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (k *jsiiProxy_KubernetesAutomaticClusterServiceMeshCertificateAuthorityOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (k *jsiiProxy_KubernetesAutomaticClusterServiceMeshCertificateAuthorityOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (k *jsiiProxy_KubernetesAutomaticClusterServiceMeshCertificateAuthorityOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (k *jsiiProxy_KubernetesAutomaticClusterServiceMeshCertificateAuthorityOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (k *jsiiProxy_KubernetesAutomaticClusterServiceMeshCertificateAuthorityOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (k *jsiiProxy_KubernetesAutomaticClusterServiceMeshCertificateAuthorityOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (k *jsiiProxy_KubernetesAutomaticClusterServiceMeshCertificateAuthorityOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (k *jsiiProxy_KubernetesAutomaticClusterServiceMeshCertificateAuthorityOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (k *jsiiProxy_KubernetesAutomaticClusterServiceMeshCertificateAuthorityOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		k,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (k *jsiiProxy_KubernetesAutomaticClusterServiceMeshCertificateAuthorityOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (k *jsiiProxy_KubernetesAutomaticClusterServiceMeshCertificateAuthorityOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
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

func (k *jsiiProxy_KubernetesAutomaticClusterServiceMeshCertificateAuthorityOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		k,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

