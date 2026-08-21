// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package redhatopenshiftcluster

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-provider-azurerm-go/azurerm/v18/jsii"

	"github.com/cdktn-io/cdktn-provider-azurerm-go/azurerm/v18/redhatopenshiftcluster/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type RedhatOpenshiftClusterPlatformWorkloadIdentityProfileOutputReference interface {
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
	InternalValue() *RedhatOpenshiftClusterPlatformWorkloadIdentityProfile
	SetInternalValue(val *RedhatOpenshiftClusterPlatformWorkloadIdentityProfile)
	PlatformWorkloadIdentity() RedhatOpenshiftClusterPlatformWorkloadIdentityProfilePlatformWorkloadIdentityList
	PlatformWorkloadIdentityInput() interface{}
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktn.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktn.IInterpolatingParent)
	UpgradeableTo() *string
	SetUpgradeableTo(val *string)
	UpgradeableToInput() *string
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
	PutPlatformWorkloadIdentity(value interface{})
	ResetUpgradeableTo()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktn.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for RedhatOpenshiftClusterPlatformWorkloadIdentityProfileOutputReference
type jsiiProxy_RedhatOpenshiftClusterPlatformWorkloadIdentityProfileOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_RedhatOpenshiftClusterPlatformWorkloadIdentityProfileOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RedhatOpenshiftClusterPlatformWorkloadIdentityProfileOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RedhatOpenshiftClusterPlatformWorkloadIdentityProfileOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RedhatOpenshiftClusterPlatformWorkloadIdentityProfileOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RedhatOpenshiftClusterPlatformWorkloadIdentityProfileOutputReference) InternalValue() *RedhatOpenshiftClusterPlatformWorkloadIdentityProfile {
	var returns *RedhatOpenshiftClusterPlatformWorkloadIdentityProfile
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RedhatOpenshiftClusterPlatformWorkloadIdentityProfileOutputReference) PlatformWorkloadIdentity() RedhatOpenshiftClusterPlatformWorkloadIdentityProfilePlatformWorkloadIdentityList {
	var returns RedhatOpenshiftClusterPlatformWorkloadIdentityProfilePlatformWorkloadIdentityList
	_jsii_.Get(
		j,
		"platformWorkloadIdentity",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RedhatOpenshiftClusterPlatformWorkloadIdentityProfileOutputReference) PlatformWorkloadIdentityInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"platformWorkloadIdentityInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RedhatOpenshiftClusterPlatformWorkloadIdentityProfileOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RedhatOpenshiftClusterPlatformWorkloadIdentityProfileOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RedhatOpenshiftClusterPlatformWorkloadIdentityProfileOutputReference) UpgradeableTo() *string {
	var returns *string
	_jsii_.Get(
		j,
		"upgradeableTo",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RedhatOpenshiftClusterPlatformWorkloadIdentityProfileOutputReference) UpgradeableToInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"upgradeableToInput",
		&returns,
	)
	return returns
}


func NewRedhatOpenshiftClusterPlatformWorkloadIdentityProfileOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) RedhatOpenshiftClusterPlatformWorkloadIdentityProfileOutputReference {
	_init_.Initialize()

	if err := validateNewRedhatOpenshiftClusterPlatformWorkloadIdentityProfileOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_RedhatOpenshiftClusterPlatformWorkloadIdentityProfileOutputReference{}

	_jsii_.Create(
		"@cdktn/provider-azurerm.redhatOpenshiftCluster.RedhatOpenshiftClusterPlatformWorkloadIdentityProfileOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewRedhatOpenshiftClusterPlatformWorkloadIdentityProfileOutputReference_Override(r RedhatOpenshiftClusterPlatformWorkloadIdentityProfileOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/provider-azurerm.redhatOpenshiftCluster.RedhatOpenshiftClusterPlatformWorkloadIdentityProfileOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		r,
	)
}

func (j *jsiiProxy_RedhatOpenshiftClusterPlatformWorkloadIdentityProfileOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_RedhatOpenshiftClusterPlatformWorkloadIdentityProfileOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_RedhatOpenshiftClusterPlatformWorkloadIdentityProfileOutputReference)SetInternalValue(val *RedhatOpenshiftClusterPlatformWorkloadIdentityProfile) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_RedhatOpenshiftClusterPlatformWorkloadIdentityProfileOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_RedhatOpenshiftClusterPlatformWorkloadIdentityProfileOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_RedhatOpenshiftClusterPlatformWorkloadIdentityProfileOutputReference)SetUpgradeableTo(val *string) {
	if err := j.validateSetUpgradeableToParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"upgradeableTo",
		val,
	)
}

func (r *jsiiProxy_RedhatOpenshiftClusterPlatformWorkloadIdentityProfileOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		r,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (r *jsiiProxy_RedhatOpenshiftClusterPlatformWorkloadIdentityProfileOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	if err := r.validateGetAnyMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]interface{}

	_jsii_.Invoke(
		r,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (r *jsiiProxy_RedhatOpenshiftClusterPlatformWorkloadIdentityProfileOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
	if err := r.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		r,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (r *jsiiProxy_RedhatOpenshiftClusterPlatformWorkloadIdentityProfileOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	if err := r.validateGetBooleanMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*bool

	_jsii_.Invoke(
		r,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (r *jsiiProxy_RedhatOpenshiftClusterPlatformWorkloadIdentityProfileOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	if err := r.validateGetListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*string

	_jsii_.Invoke(
		r,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (r *jsiiProxy_RedhatOpenshiftClusterPlatformWorkloadIdentityProfileOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	if err := r.validateGetNumberAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *float64

	_jsii_.Invoke(
		r,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (r *jsiiProxy_RedhatOpenshiftClusterPlatformWorkloadIdentityProfileOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	if err := r.validateGetNumberListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*float64

	_jsii_.Invoke(
		r,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (r *jsiiProxy_RedhatOpenshiftClusterPlatformWorkloadIdentityProfileOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	if err := r.validateGetNumberMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*float64

	_jsii_.Invoke(
		r,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (r *jsiiProxy_RedhatOpenshiftClusterPlatformWorkloadIdentityProfileOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	if err := r.validateGetStringAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		r,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (r *jsiiProxy_RedhatOpenshiftClusterPlatformWorkloadIdentityProfileOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	if err := r.validateGetStringMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*string

	_jsii_.Invoke(
		r,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (r *jsiiProxy_RedhatOpenshiftClusterPlatformWorkloadIdentityProfileOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		r,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (r *jsiiProxy_RedhatOpenshiftClusterPlatformWorkloadIdentityProfileOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
	if err := r.validateInterpolationForAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		r,
		"interpolationForAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (r *jsiiProxy_RedhatOpenshiftClusterPlatformWorkloadIdentityProfileOutputReference) PutPlatformWorkloadIdentity(value interface{}) {
	if err := r.validatePutPlatformWorkloadIdentityParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		r,
		"putPlatformWorkloadIdentity",
		[]interface{}{value},
	)
}

func (r *jsiiProxy_RedhatOpenshiftClusterPlatformWorkloadIdentityProfileOutputReference) ResetUpgradeableTo() {
	_jsii_.InvokeVoid(
		r,
		"resetUpgradeableTo",
		nil, // no parameters
	)
}

func (r *jsiiProxy_RedhatOpenshiftClusterPlatformWorkloadIdentityProfileOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
	if err := r.validateResolveParameters(context); err != nil {
		panic(err)
	}
	var returns interface{}

	_jsii_.Invoke(
		r,
		"resolve",
		[]interface{}{context},
		&returns,
	)

	return returns
}

func (r *jsiiProxy_RedhatOpenshiftClusterPlatformWorkloadIdentityProfileOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		r,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

