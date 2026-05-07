// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package elasticcloudelasticsearch

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-provider-azurerm-go/azurerm/v16/jsii"

	"github.com/cdktn-io/cdktn-provider-azurerm-go/azurerm/v16/elasticcloudelasticsearch/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type ElasticCloudElasticsearchLogsOutputReference interface {
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
	FilteringTag() ElasticCloudElasticsearchLogsFilteringTagList
	FilteringTagInput() interface{}
	// Experimental.
	Fqn() *string
	InternalValue() *ElasticCloudElasticsearchLogs
	SetInternalValue(val *ElasticCloudElasticsearchLogs)
	SendActivityLogs() interface{}
	SetSendActivityLogs(val interface{})
	SendActivityLogsInput() interface{}
	SendAzureadLogs() interface{}
	SetSendAzureadLogs(val interface{})
	SendAzureadLogsInput() interface{}
	SendSubscriptionLogs() interface{}
	SetSendSubscriptionLogs(val interface{})
	SendSubscriptionLogsInput() interface{}
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
	PutFilteringTag(value interface{})
	ResetFilteringTag()
	ResetSendActivityLogs()
	ResetSendAzureadLogs()
	ResetSendSubscriptionLogs()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktn.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for ElasticCloudElasticsearchLogsOutputReference
type jsiiProxy_ElasticCloudElasticsearchLogsOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_ElasticCloudElasticsearchLogsOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ElasticCloudElasticsearchLogsOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ElasticCloudElasticsearchLogsOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ElasticCloudElasticsearchLogsOutputReference) FilteringTag() ElasticCloudElasticsearchLogsFilteringTagList {
	var returns ElasticCloudElasticsearchLogsFilteringTagList
	_jsii_.Get(
		j,
		"filteringTag",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ElasticCloudElasticsearchLogsOutputReference) FilteringTagInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"filteringTagInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ElasticCloudElasticsearchLogsOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ElasticCloudElasticsearchLogsOutputReference) InternalValue() *ElasticCloudElasticsearchLogs {
	var returns *ElasticCloudElasticsearchLogs
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ElasticCloudElasticsearchLogsOutputReference) SendActivityLogs() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"sendActivityLogs",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ElasticCloudElasticsearchLogsOutputReference) SendActivityLogsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"sendActivityLogsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ElasticCloudElasticsearchLogsOutputReference) SendAzureadLogs() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"sendAzureadLogs",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ElasticCloudElasticsearchLogsOutputReference) SendAzureadLogsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"sendAzureadLogsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ElasticCloudElasticsearchLogsOutputReference) SendSubscriptionLogs() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"sendSubscriptionLogs",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ElasticCloudElasticsearchLogsOutputReference) SendSubscriptionLogsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"sendSubscriptionLogsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ElasticCloudElasticsearchLogsOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ElasticCloudElasticsearchLogsOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewElasticCloudElasticsearchLogsOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) ElasticCloudElasticsearchLogsOutputReference {
	_init_.Initialize()

	if err := validateNewElasticCloudElasticsearchLogsOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_ElasticCloudElasticsearchLogsOutputReference{}

	_jsii_.Create(
		"@cdktn/provider-azurerm.elasticCloudElasticsearch.ElasticCloudElasticsearchLogsOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewElasticCloudElasticsearchLogsOutputReference_Override(e ElasticCloudElasticsearchLogsOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/provider-azurerm.elasticCloudElasticsearch.ElasticCloudElasticsearchLogsOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		e,
	)
}

func (j *jsiiProxy_ElasticCloudElasticsearchLogsOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_ElasticCloudElasticsearchLogsOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_ElasticCloudElasticsearchLogsOutputReference)SetInternalValue(val *ElasticCloudElasticsearchLogs) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_ElasticCloudElasticsearchLogsOutputReference)SetSendActivityLogs(val interface{}) {
	if err := j.validateSetSendActivityLogsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"sendActivityLogs",
		val,
	)
}

func (j *jsiiProxy_ElasticCloudElasticsearchLogsOutputReference)SetSendAzureadLogs(val interface{}) {
	if err := j.validateSetSendAzureadLogsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"sendAzureadLogs",
		val,
	)
}

func (j *jsiiProxy_ElasticCloudElasticsearchLogsOutputReference)SetSendSubscriptionLogs(val interface{}) {
	if err := j.validateSetSendSubscriptionLogsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"sendSubscriptionLogs",
		val,
	)
}

func (j *jsiiProxy_ElasticCloudElasticsearchLogsOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_ElasticCloudElasticsearchLogsOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (e *jsiiProxy_ElasticCloudElasticsearchLogsOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		e,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (e *jsiiProxy_ElasticCloudElasticsearchLogsOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	if err := e.validateGetAnyMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]interface{}

	_jsii_.Invoke(
		e,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_ElasticCloudElasticsearchLogsOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
	if err := e.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		e,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_ElasticCloudElasticsearchLogsOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	if err := e.validateGetBooleanMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*bool

	_jsii_.Invoke(
		e,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_ElasticCloudElasticsearchLogsOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	if err := e.validateGetListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*string

	_jsii_.Invoke(
		e,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_ElasticCloudElasticsearchLogsOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	if err := e.validateGetNumberAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *float64

	_jsii_.Invoke(
		e,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_ElasticCloudElasticsearchLogsOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	if err := e.validateGetNumberListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*float64

	_jsii_.Invoke(
		e,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_ElasticCloudElasticsearchLogsOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	if err := e.validateGetNumberMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*float64

	_jsii_.Invoke(
		e,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_ElasticCloudElasticsearchLogsOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	if err := e.validateGetStringAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		e,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_ElasticCloudElasticsearchLogsOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	if err := e.validateGetStringMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*string

	_jsii_.Invoke(
		e,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_ElasticCloudElasticsearchLogsOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		e,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (e *jsiiProxy_ElasticCloudElasticsearchLogsOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
	if err := e.validateInterpolationForAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		e,
		"interpolationForAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_ElasticCloudElasticsearchLogsOutputReference) PutFilteringTag(value interface{}) {
	if err := e.validatePutFilteringTagParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		e,
		"putFilteringTag",
		[]interface{}{value},
	)
}

func (e *jsiiProxy_ElasticCloudElasticsearchLogsOutputReference) ResetFilteringTag() {
	_jsii_.InvokeVoid(
		e,
		"resetFilteringTag",
		nil, // no parameters
	)
}

func (e *jsiiProxy_ElasticCloudElasticsearchLogsOutputReference) ResetSendActivityLogs() {
	_jsii_.InvokeVoid(
		e,
		"resetSendActivityLogs",
		nil, // no parameters
	)
}

func (e *jsiiProxy_ElasticCloudElasticsearchLogsOutputReference) ResetSendAzureadLogs() {
	_jsii_.InvokeVoid(
		e,
		"resetSendAzureadLogs",
		nil, // no parameters
	)
}

func (e *jsiiProxy_ElasticCloudElasticsearchLogsOutputReference) ResetSendSubscriptionLogs() {
	_jsii_.InvokeVoid(
		e,
		"resetSendSubscriptionLogs",
		nil, // no parameters
	)
}

func (e *jsiiProxy_ElasticCloudElasticsearchLogsOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
	if err := e.validateResolveParameters(context); err != nil {
		panic(err)
	}
	var returns interface{}

	_jsii_.Invoke(
		e,
		"resolve",
		[]interface{}{context},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_ElasticCloudElasticsearchLogsOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		e,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

