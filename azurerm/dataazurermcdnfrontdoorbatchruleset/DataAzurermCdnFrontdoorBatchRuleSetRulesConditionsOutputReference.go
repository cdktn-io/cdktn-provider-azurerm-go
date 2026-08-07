// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package dataazurermcdnfrontdoorbatchruleset

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-provider-azurerm-go/azurerm/v17/jsii"

	"github.com/cdktn-io/cdktn-provider-azurerm-go/azurerm/v17/dataazurermcdnfrontdoorbatchruleset/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type DataAzurermCdnFrontdoorBatchRuleSetRulesConditionsOutputReference interface {
	cdktn.ComplexObject
	ClientPort() DataAzurermCdnFrontdoorBatchRuleSetRulesConditionsClientPortList
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
	DeviceType() DataAzurermCdnFrontdoorBatchRuleSetRulesConditionsDeviceTypeList
	// Experimental.
	Fqn() *string
	HostName() DataAzurermCdnFrontdoorBatchRuleSetRulesConditionsHostNameList
	HttpVersion() DataAzurermCdnFrontdoorBatchRuleSetRulesConditionsHttpVersionList
	InternalValue() *DataAzurermCdnFrontdoorBatchRuleSetRulesConditions
	SetInternalValue(val *DataAzurermCdnFrontdoorBatchRuleSetRulesConditions)
	PostArgument() DataAzurermCdnFrontdoorBatchRuleSetRulesConditionsPostArgumentList
	QueryString() DataAzurermCdnFrontdoorBatchRuleSetRulesConditionsQueryStringList
	RemoteAddress() DataAzurermCdnFrontdoorBatchRuleSetRulesConditionsRemoteAddressList
	RequestBody() DataAzurermCdnFrontdoorBatchRuleSetRulesConditionsRequestBodyList
	RequestCookies() DataAzurermCdnFrontdoorBatchRuleSetRulesConditionsRequestCookiesList
	RequestFileExtension() DataAzurermCdnFrontdoorBatchRuleSetRulesConditionsRequestFileExtensionList
	RequestFilename() DataAzurermCdnFrontdoorBatchRuleSetRulesConditionsRequestFilenameList
	RequestHeader() DataAzurermCdnFrontdoorBatchRuleSetRulesConditionsRequestHeaderList
	RequestMethod() DataAzurermCdnFrontdoorBatchRuleSetRulesConditionsRequestMethodList
	RequestPath() DataAzurermCdnFrontdoorBatchRuleSetRulesConditionsRequestPathList
	RequestScheme() DataAzurermCdnFrontdoorBatchRuleSetRulesConditionsRequestSchemeList
	RequestUrl() DataAzurermCdnFrontdoorBatchRuleSetRulesConditionsRequestUrlList
	ServerPort() DataAzurermCdnFrontdoorBatchRuleSetRulesConditionsServerPortList
	SocketAddress() DataAzurermCdnFrontdoorBatchRuleSetRulesConditionsSocketAddressList
	SslProtocol() DataAzurermCdnFrontdoorBatchRuleSetRulesConditionsSslProtocolList
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

// The jsii proxy struct for DataAzurermCdnFrontdoorBatchRuleSetRulesConditionsOutputReference
type jsiiProxy_DataAzurermCdnFrontdoorBatchRuleSetRulesConditionsOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_DataAzurermCdnFrontdoorBatchRuleSetRulesConditionsOutputReference) ClientPort() DataAzurermCdnFrontdoorBatchRuleSetRulesConditionsClientPortList {
	var returns DataAzurermCdnFrontdoorBatchRuleSetRulesConditionsClientPortList
	_jsii_.Get(
		j,
		"clientPort",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAzurermCdnFrontdoorBatchRuleSetRulesConditionsOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAzurermCdnFrontdoorBatchRuleSetRulesConditionsOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAzurermCdnFrontdoorBatchRuleSetRulesConditionsOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAzurermCdnFrontdoorBatchRuleSetRulesConditionsOutputReference) DeviceType() DataAzurermCdnFrontdoorBatchRuleSetRulesConditionsDeviceTypeList {
	var returns DataAzurermCdnFrontdoorBatchRuleSetRulesConditionsDeviceTypeList
	_jsii_.Get(
		j,
		"deviceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAzurermCdnFrontdoorBatchRuleSetRulesConditionsOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAzurermCdnFrontdoorBatchRuleSetRulesConditionsOutputReference) HostName() DataAzurermCdnFrontdoorBatchRuleSetRulesConditionsHostNameList {
	var returns DataAzurermCdnFrontdoorBatchRuleSetRulesConditionsHostNameList
	_jsii_.Get(
		j,
		"hostName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAzurermCdnFrontdoorBatchRuleSetRulesConditionsOutputReference) HttpVersion() DataAzurermCdnFrontdoorBatchRuleSetRulesConditionsHttpVersionList {
	var returns DataAzurermCdnFrontdoorBatchRuleSetRulesConditionsHttpVersionList
	_jsii_.Get(
		j,
		"httpVersion",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAzurermCdnFrontdoorBatchRuleSetRulesConditionsOutputReference) InternalValue() *DataAzurermCdnFrontdoorBatchRuleSetRulesConditions {
	var returns *DataAzurermCdnFrontdoorBatchRuleSetRulesConditions
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAzurermCdnFrontdoorBatchRuleSetRulesConditionsOutputReference) PostArgument() DataAzurermCdnFrontdoorBatchRuleSetRulesConditionsPostArgumentList {
	var returns DataAzurermCdnFrontdoorBatchRuleSetRulesConditionsPostArgumentList
	_jsii_.Get(
		j,
		"postArgument",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAzurermCdnFrontdoorBatchRuleSetRulesConditionsOutputReference) QueryString() DataAzurermCdnFrontdoorBatchRuleSetRulesConditionsQueryStringList {
	var returns DataAzurermCdnFrontdoorBatchRuleSetRulesConditionsQueryStringList
	_jsii_.Get(
		j,
		"queryString",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAzurermCdnFrontdoorBatchRuleSetRulesConditionsOutputReference) RemoteAddress() DataAzurermCdnFrontdoorBatchRuleSetRulesConditionsRemoteAddressList {
	var returns DataAzurermCdnFrontdoorBatchRuleSetRulesConditionsRemoteAddressList
	_jsii_.Get(
		j,
		"remoteAddress",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAzurermCdnFrontdoorBatchRuleSetRulesConditionsOutputReference) RequestBody() DataAzurermCdnFrontdoorBatchRuleSetRulesConditionsRequestBodyList {
	var returns DataAzurermCdnFrontdoorBatchRuleSetRulesConditionsRequestBodyList
	_jsii_.Get(
		j,
		"requestBody",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAzurermCdnFrontdoorBatchRuleSetRulesConditionsOutputReference) RequestCookies() DataAzurermCdnFrontdoorBatchRuleSetRulesConditionsRequestCookiesList {
	var returns DataAzurermCdnFrontdoorBatchRuleSetRulesConditionsRequestCookiesList
	_jsii_.Get(
		j,
		"requestCookies",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAzurermCdnFrontdoorBatchRuleSetRulesConditionsOutputReference) RequestFileExtension() DataAzurermCdnFrontdoorBatchRuleSetRulesConditionsRequestFileExtensionList {
	var returns DataAzurermCdnFrontdoorBatchRuleSetRulesConditionsRequestFileExtensionList
	_jsii_.Get(
		j,
		"requestFileExtension",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAzurermCdnFrontdoorBatchRuleSetRulesConditionsOutputReference) RequestFilename() DataAzurermCdnFrontdoorBatchRuleSetRulesConditionsRequestFilenameList {
	var returns DataAzurermCdnFrontdoorBatchRuleSetRulesConditionsRequestFilenameList
	_jsii_.Get(
		j,
		"requestFilename",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAzurermCdnFrontdoorBatchRuleSetRulesConditionsOutputReference) RequestHeader() DataAzurermCdnFrontdoorBatchRuleSetRulesConditionsRequestHeaderList {
	var returns DataAzurermCdnFrontdoorBatchRuleSetRulesConditionsRequestHeaderList
	_jsii_.Get(
		j,
		"requestHeader",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAzurermCdnFrontdoorBatchRuleSetRulesConditionsOutputReference) RequestMethod() DataAzurermCdnFrontdoorBatchRuleSetRulesConditionsRequestMethodList {
	var returns DataAzurermCdnFrontdoorBatchRuleSetRulesConditionsRequestMethodList
	_jsii_.Get(
		j,
		"requestMethod",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAzurermCdnFrontdoorBatchRuleSetRulesConditionsOutputReference) RequestPath() DataAzurermCdnFrontdoorBatchRuleSetRulesConditionsRequestPathList {
	var returns DataAzurermCdnFrontdoorBatchRuleSetRulesConditionsRequestPathList
	_jsii_.Get(
		j,
		"requestPath",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAzurermCdnFrontdoorBatchRuleSetRulesConditionsOutputReference) RequestScheme() DataAzurermCdnFrontdoorBatchRuleSetRulesConditionsRequestSchemeList {
	var returns DataAzurermCdnFrontdoorBatchRuleSetRulesConditionsRequestSchemeList
	_jsii_.Get(
		j,
		"requestScheme",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAzurermCdnFrontdoorBatchRuleSetRulesConditionsOutputReference) RequestUrl() DataAzurermCdnFrontdoorBatchRuleSetRulesConditionsRequestUrlList {
	var returns DataAzurermCdnFrontdoorBatchRuleSetRulesConditionsRequestUrlList
	_jsii_.Get(
		j,
		"requestUrl",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAzurermCdnFrontdoorBatchRuleSetRulesConditionsOutputReference) ServerPort() DataAzurermCdnFrontdoorBatchRuleSetRulesConditionsServerPortList {
	var returns DataAzurermCdnFrontdoorBatchRuleSetRulesConditionsServerPortList
	_jsii_.Get(
		j,
		"serverPort",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAzurermCdnFrontdoorBatchRuleSetRulesConditionsOutputReference) SocketAddress() DataAzurermCdnFrontdoorBatchRuleSetRulesConditionsSocketAddressList {
	var returns DataAzurermCdnFrontdoorBatchRuleSetRulesConditionsSocketAddressList
	_jsii_.Get(
		j,
		"socketAddress",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAzurermCdnFrontdoorBatchRuleSetRulesConditionsOutputReference) SslProtocol() DataAzurermCdnFrontdoorBatchRuleSetRulesConditionsSslProtocolList {
	var returns DataAzurermCdnFrontdoorBatchRuleSetRulesConditionsSslProtocolList
	_jsii_.Get(
		j,
		"sslProtocol",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAzurermCdnFrontdoorBatchRuleSetRulesConditionsOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAzurermCdnFrontdoorBatchRuleSetRulesConditionsOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewDataAzurermCdnFrontdoorBatchRuleSetRulesConditionsOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) DataAzurermCdnFrontdoorBatchRuleSetRulesConditionsOutputReference {
	_init_.Initialize()

	if err := validateNewDataAzurermCdnFrontdoorBatchRuleSetRulesConditionsOutputReferenceParameters(terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet); err != nil {
		panic(err)
	}
	j := jsiiProxy_DataAzurermCdnFrontdoorBatchRuleSetRulesConditionsOutputReference{}

	_jsii_.Create(
		"@cdktn/provider-azurerm.dataAzurermCdnFrontdoorBatchRuleSet.DataAzurermCdnFrontdoorBatchRuleSetRulesConditionsOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		&j,
	)

	return &j
}

func NewDataAzurermCdnFrontdoorBatchRuleSetRulesConditionsOutputReference_Override(d DataAzurermCdnFrontdoorBatchRuleSetRulesConditionsOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/provider-azurerm.dataAzurermCdnFrontdoorBatchRuleSet.DataAzurermCdnFrontdoorBatchRuleSetRulesConditionsOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		d,
	)
}

func (j *jsiiProxy_DataAzurermCdnFrontdoorBatchRuleSetRulesConditionsOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_DataAzurermCdnFrontdoorBatchRuleSetRulesConditionsOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_DataAzurermCdnFrontdoorBatchRuleSetRulesConditionsOutputReference)SetInternalValue(val *DataAzurermCdnFrontdoorBatchRuleSetRulesConditions) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_DataAzurermCdnFrontdoorBatchRuleSetRulesConditionsOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_DataAzurermCdnFrontdoorBatchRuleSetRulesConditionsOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (d *jsiiProxy_DataAzurermCdnFrontdoorBatchRuleSetRulesConditionsOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataAzurermCdnFrontdoorBatchRuleSetRulesConditionsOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	if err := d.validateGetAnyMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]interface{}

	_jsii_.Invoke(
		d,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataAzurermCdnFrontdoorBatchRuleSetRulesConditionsOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
	if err := d.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		d,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataAzurermCdnFrontdoorBatchRuleSetRulesConditionsOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	if err := d.validateGetBooleanMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*bool

	_jsii_.Invoke(
		d,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataAzurermCdnFrontdoorBatchRuleSetRulesConditionsOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	if err := d.validateGetListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*string

	_jsii_.Invoke(
		d,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataAzurermCdnFrontdoorBatchRuleSetRulesConditionsOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	if err := d.validateGetNumberAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *float64

	_jsii_.Invoke(
		d,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataAzurermCdnFrontdoorBatchRuleSetRulesConditionsOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	if err := d.validateGetNumberListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*float64

	_jsii_.Invoke(
		d,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataAzurermCdnFrontdoorBatchRuleSetRulesConditionsOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	if err := d.validateGetNumberMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*float64

	_jsii_.Invoke(
		d,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataAzurermCdnFrontdoorBatchRuleSetRulesConditionsOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	if err := d.validateGetStringAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		d,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataAzurermCdnFrontdoorBatchRuleSetRulesConditionsOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	if err := d.validateGetStringMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*string

	_jsii_.Invoke(
		d,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataAzurermCdnFrontdoorBatchRuleSetRulesConditionsOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		d,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataAzurermCdnFrontdoorBatchRuleSetRulesConditionsOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
	if err := d.validateInterpolationForAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		d,
		"interpolationForAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataAzurermCdnFrontdoorBatchRuleSetRulesConditionsOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
	if err := d.validateResolveParameters(context); err != nil {
		panic(err)
	}
	var returns interface{}

	_jsii_.Invoke(
		d,
		"resolve",
		[]interface{}{context},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataAzurermCdnFrontdoorBatchRuleSetRulesConditionsOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

