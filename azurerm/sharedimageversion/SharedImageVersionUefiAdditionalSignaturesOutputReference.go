// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package sharedimageversion

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-provider-azurerm-go/azurerm/v19/jsii"

	"github.com/cdktn-io/cdktn-provider-azurerm-go/azurerm/v19/sharedimageversion/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type SharedImageVersionUefiAdditionalSignaturesOutputReference interface {
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
	Db() SharedImageVersionUefiAdditionalSignaturesDbList
	DbInput() interface{}
	Dbx() SharedImageVersionUefiAdditionalSignaturesDbxList
	DbxInput() interface{}
	// Experimental.
	Fqn() *string
	InternalValue() *SharedImageVersionUefiAdditionalSignatures
	SetInternalValue(val *SharedImageVersionUefiAdditionalSignatures)
	Kek() SharedImageVersionUefiAdditionalSignaturesKekList
	KekInput() interface{}
	Pk() SharedImageVersionUefiAdditionalSignaturesPkOutputReference
	PkInput() *SharedImageVersionUefiAdditionalSignaturesPk
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
	PutDb(value interface{})
	PutDbx(value interface{})
	PutKek(value interface{})
	PutPk(value *SharedImageVersionUefiAdditionalSignaturesPk)
	ResetDb()
	ResetDbx()
	ResetKek()
	ResetPk()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktn.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for SharedImageVersionUefiAdditionalSignaturesOutputReference
type jsiiProxy_SharedImageVersionUefiAdditionalSignaturesOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_SharedImageVersionUefiAdditionalSignaturesOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SharedImageVersionUefiAdditionalSignaturesOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SharedImageVersionUefiAdditionalSignaturesOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SharedImageVersionUefiAdditionalSignaturesOutputReference) Db() SharedImageVersionUefiAdditionalSignaturesDbList {
	var returns SharedImageVersionUefiAdditionalSignaturesDbList
	_jsii_.Get(
		j,
		"db",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SharedImageVersionUefiAdditionalSignaturesOutputReference) DbInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"dbInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SharedImageVersionUefiAdditionalSignaturesOutputReference) Dbx() SharedImageVersionUefiAdditionalSignaturesDbxList {
	var returns SharedImageVersionUefiAdditionalSignaturesDbxList
	_jsii_.Get(
		j,
		"dbx",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SharedImageVersionUefiAdditionalSignaturesOutputReference) DbxInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"dbxInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SharedImageVersionUefiAdditionalSignaturesOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SharedImageVersionUefiAdditionalSignaturesOutputReference) InternalValue() *SharedImageVersionUefiAdditionalSignatures {
	var returns *SharedImageVersionUefiAdditionalSignatures
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SharedImageVersionUefiAdditionalSignaturesOutputReference) Kek() SharedImageVersionUefiAdditionalSignaturesKekList {
	var returns SharedImageVersionUefiAdditionalSignaturesKekList
	_jsii_.Get(
		j,
		"kek",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SharedImageVersionUefiAdditionalSignaturesOutputReference) KekInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"kekInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SharedImageVersionUefiAdditionalSignaturesOutputReference) Pk() SharedImageVersionUefiAdditionalSignaturesPkOutputReference {
	var returns SharedImageVersionUefiAdditionalSignaturesPkOutputReference
	_jsii_.Get(
		j,
		"pk",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SharedImageVersionUefiAdditionalSignaturesOutputReference) PkInput() *SharedImageVersionUefiAdditionalSignaturesPk {
	var returns *SharedImageVersionUefiAdditionalSignaturesPk
	_jsii_.Get(
		j,
		"pkInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SharedImageVersionUefiAdditionalSignaturesOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SharedImageVersionUefiAdditionalSignaturesOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewSharedImageVersionUefiAdditionalSignaturesOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) SharedImageVersionUefiAdditionalSignaturesOutputReference {
	_init_.Initialize()

	if err := validateNewSharedImageVersionUefiAdditionalSignaturesOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_SharedImageVersionUefiAdditionalSignaturesOutputReference{}

	_jsii_.Create(
		"@cdktn/provider-azurerm.sharedImageVersion.SharedImageVersionUefiAdditionalSignaturesOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewSharedImageVersionUefiAdditionalSignaturesOutputReference_Override(s SharedImageVersionUefiAdditionalSignaturesOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/provider-azurerm.sharedImageVersion.SharedImageVersionUefiAdditionalSignaturesOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		s,
	)
}

func (j *jsiiProxy_SharedImageVersionUefiAdditionalSignaturesOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_SharedImageVersionUefiAdditionalSignaturesOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_SharedImageVersionUefiAdditionalSignaturesOutputReference)SetInternalValue(val *SharedImageVersionUefiAdditionalSignatures) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_SharedImageVersionUefiAdditionalSignaturesOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_SharedImageVersionUefiAdditionalSignaturesOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (s *jsiiProxy_SharedImageVersionUefiAdditionalSignaturesOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		s,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SharedImageVersionUefiAdditionalSignaturesOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (s *jsiiProxy_SharedImageVersionUefiAdditionalSignaturesOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (s *jsiiProxy_SharedImageVersionUefiAdditionalSignaturesOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (s *jsiiProxy_SharedImageVersionUefiAdditionalSignaturesOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (s *jsiiProxy_SharedImageVersionUefiAdditionalSignaturesOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (s *jsiiProxy_SharedImageVersionUefiAdditionalSignaturesOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (s *jsiiProxy_SharedImageVersionUefiAdditionalSignaturesOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (s *jsiiProxy_SharedImageVersionUefiAdditionalSignaturesOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (s *jsiiProxy_SharedImageVersionUefiAdditionalSignaturesOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (s *jsiiProxy_SharedImageVersionUefiAdditionalSignaturesOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		s,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SharedImageVersionUefiAdditionalSignaturesOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (s *jsiiProxy_SharedImageVersionUefiAdditionalSignaturesOutputReference) PutDb(value interface{}) {
	if err := s.validatePutDbParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"putDb",
		[]interface{}{value},
	)
}

func (s *jsiiProxy_SharedImageVersionUefiAdditionalSignaturesOutputReference) PutDbx(value interface{}) {
	if err := s.validatePutDbxParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"putDbx",
		[]interface{}{value},
	)
}

func (s *jsiiProxy_SharedImageVersionUefiAdditionalSignaturesOutputReference) PutKek(value interface{}) {
	if err := s.validatePutKekParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"putKek",
		[]interface{}{value},
	)
}

func (s *jsiiProxy_SharedImageVersionUefiAdditionalSignaturesOutputReference) PutPk(value *SharedImageVersionUefiAdditionalSignaturesPk) {
	if err := s.validatePutPkParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"putPk",
		[]interface{}{value},
	)
}

func (s *jsiiProxy_SharedImageVersionUefiAdditionalSignaturesOutputReference) ResetDb() {
	_jsii_.InvokeVoid(
		s,
		"resetDb",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SharedImageVersionUefiAdditionalSignaturesOutputReference) ResetDbx() {
	_jsii_.InvokeVoid(
		s,
		"resetDbx",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SharedImageVersionUefiAdditionalSignaturesOutputReference) ResetKek() {
	_jsii_.InvokeVoid(
		s,
		"resetKek",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SharedImageVersionUefiAdditionalSignaturesOutputReference) ResetPk() {
	_jsii_.InvokeVoid(
		s,
		"resetPk",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SharedImageVersionUefiAdditionalSignaturesOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
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

func (s *jsiiProxy_SharedImageVersionUefiAdditionalSignaturesOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		s,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

