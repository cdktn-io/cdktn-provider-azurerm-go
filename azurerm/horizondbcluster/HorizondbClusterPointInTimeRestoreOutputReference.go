// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package horizondbcluster

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-provider-azurerm-go/azurerm/v19/jsii"

	"github.com/cdktn-io/cdktn-provider-azurerm-go/azurerm/v19/horizondbcluster/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type HorizondbClusterPointInTimeRestoreOutputReference interface {
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
	InternalValue() *HorizondbClusterPointInTimeRestore
	SetInternalValue(val *HorizondbClusterPointInTimeRestore)
	PointInTimeUtc() *string
	SetPointInTimeUtc(val *string)
	PointInTimeUtcInput() *string
	SourceHorizondbClusterId() *string
	SetSourceHorizondbClusterId(val *string)
	SourceHorizondbClusterIdInput() *string
	SourceLocation() *string
	SetSourceLocation(val *string)
	SourceLocationInput() *string
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
	ResetSourceLocation()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktn.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for HorizondbClusterPointInTimeRestoreOutputReference
type jsiiProxy_HorizondbClusterPointInTimeRestoreOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_HorizondbClusterPointInTimeRestoreOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HorizondbClusterPointInTimeRestoreOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HorizondbClusterPointInTimeRestoreOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HorizondbClusterPointInTimeRestoreOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HorizondbClusterPointInTimeRestoreOutputReference) InternalValue() *HorizondbClusterPointInTimeRestore {
	var returns *HorizondbClusterPointInTimeRestore
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HorizondbClusterPointInTimeRestoreOutputReference) PointInTimeUtc() *string {
	var returns *string
	_jsii_.Get(
		j,
		"pointInTimeUtc",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HorizondbClusterPointInTimeRestoreOutputReference) PointInTimeUtcInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"pointInTimeUtcInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HorizondbClusterPointInTimeRestoreOutputReference) SourceHorizondbClusterId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"sourceHorizondbClusterId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HorizondbClusterPointInTimeRestoreOutputReference) SourceHorizondbClusterIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"sourceHorizondbClusterIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HorizondbClusterPointInTimeRestoreOutputReference) SourceLocation() *string {
	var returns *string
	_jsii_.Get(
		j,
		"sourceLocation",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HorizondbClusterPointInTimeRestoreOutputReference) SourceLocationInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"sourceLocationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HorizondbClusterPointInTimeRestoreOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HorizondbClusterPointInTimeRestoreOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewHorizondbClusterPointInTimeRestoreOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) HorizondbClusterPointInTimeRestoreOutputReference {
	_init_.Initialize()

	if err := validateNewHorizondbClusterPointInTimeRestoreOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_HorizondbClusterPointInTimeRestoreOutputReference{}

	_jsii_.Create(
		"@cdktn/provider-azurerm.horizondbCluster.HorizondbClusterPointInTimeRestoreOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewHorizondbClusterPointInTimeRestoreOutputReference_Override(h HorizondbClusterPointInTimeRestoreOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/provider-azurerm.horizondbCluster.HorizondbClusterPointInTimeRestoreOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		h,
	)
}

func (j *jsiiProxy_HorizondbClusterPointInTimeRestoreOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_HorizondbClusterPointInTimeRestoreOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_HorizondbClusterPointInTimeRestoreOutputReference)SetInternalValue(val *HorizondbClusterPointInTimeRestore) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_HorizondbClusterPointInTimeRestoreOutputReference)SetPointInTimeUtc(val *string) {
	if err := j.validateSetPointInTimeUtcParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"pointInTimeUtc",
		val,
	)
}

func (j *jsiiProxy_HorizondbClusterPointInTimeRestoreOutputReference)SetSourceHorizondbClusterId(val *string) {
	if err := j.validateSetSourceHorizondbClusterIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"sourceHorizondbClusterId",
		val,
	)
}

func (j *jsiiProxy_HorizondbClusterPointInTimeRestoreOutputReference)SetSourceLocation(val *string) {
	if err := j.validateSetSourceLocationParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"sourceLocation",
		val,
	)
}

func (j *jsiiProxy_HorizondbClusterPointInTimeRestoreOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_HorizondbClusterPointInTimeRestoreOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (h *jsiiProxy_HorizondbClusterPointInTimeRestoreOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		h,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (h *jsiiProxy_HorizondbClusterPointInTimeRestoreOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	if err := h.validateGetAnyMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]interface{}

	_jsii_.Invoke(
		h,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (h *jsiiProxy_HorizondbClusterPointInTimeRestoreOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
	if err := h.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		h,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (h *jsiiProxy_HorizondbClusterPointInTimeRestoreOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	if err := h.validateGetBooleanMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*bool

	_jsii_.Invoke(
		h,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (h *jsiiProxy_HorizondbClusterPointInTimeRestoreOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	if err := h.validateGetListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*string

	_jsii_.Invoke(
		h,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (h *jsiiProxy_HorizondbClusterPointInTimeRestoreOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	if err := h.validateGetNumberAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *float64

	_jsii_.Invoke(
		h,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (h *jsiiProxy_HorizondbClusterPointInTimeRestoreOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	if err := h.validateGetNumberListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*float64

	_jsii_.Invoke(
		h,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (h *jsiiProxy_HorizondbClusterPointInTimeRestoreOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	if err := h.validateGetNumberMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*float64

	_jsii_.Invoke(
		h,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (h *jsiiProxy_HorizondbClusterPointInTimeRestoreOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	if err := h.validateGetStringAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		h,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (h *jsiiProxy_HorizondbClusterPointInTimeRestoreOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	if err := h.validateGetStringMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*string

	_jsii_.Invoke(
		h,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (h *jsiiProxy_HorizondbClusterPointInTimeRestoreOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		h,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (h *jsiiProxy_HorizondbClusterPointInTimeRestoreOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
	if err := h.validateInterpolationForAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		h,
		"interpolationForAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (h *jsiiProxy_HorizondbClusterPointInTimeRestoreOutputReference) ResetSourceLocation() {
	_jsii_.InvokeVoid(
		h,
		"resetSourceLocation",
		nil, // no parameters
	)
}

func (h *jsiiProxy_HorizondbClusterPointInTimeRestoreOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
	if err := h.validateResolveParameters(context); err != nil {
		panic(err)
	}
	var returns interface{}

	_jsii_.Invoke(
		h,
		"resolve",
		[]interface{}{context},
		&returns,
	)

	return returns
}

func (h *jsiiProxy_HorizondbClusterPointInTimeRestoreOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		h,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

