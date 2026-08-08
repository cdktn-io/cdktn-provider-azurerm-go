// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package cdnfrontdoorbatchruleset

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-provider-azurerm-go/azurerm/v18/jsii"

	"github.com/cdktn-io/cdktn-provider-azurerm-go/azurerm/v18/cdnfrontdoorbatchruleset/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type CdnFrontdoorBatchRuleSetRuleConditionsOutputReference interface {
	cdktn.ComplexObject
	ClientPort() CdnFrontdoorBatchRuleSetRuleConditionsClientPortList
	ClientPortInput() interface{}
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
	DeviceType() CdnFrontdoorBatchRuleSetRuleConditionsDeviceTypeList
	DeviceTypeInput() interface{}
	// Experimental.
	Fqn() *string
	HostName() CdnFrontdoorBatchRuleSetRuleConditionsHostNameList
	HostNameInput() interface{}
	HttpVersion() CdnFrontdoorBatchRuleSetRuleConditionsHttpVersionList
	HttpVersionInput() interface{}
	InternalValue() *CdnFrontdoorBatchRuleSetRuleConditions
	SetInternalValue(val *CdnFrontdoorBatchRuleSetRuleConditions)
	PostArgument() CdnFrontdoorBatchRuleSetRuleConditionsPostArgumentList
	PostArgumentInput() interface{}
	QueryString() CdnFrontdoorBatchRuleSetRuleConditionsQueryStringList
	QueryStringInput() interface{}
	RemoteAddress() CdnFrontdoorBatchRuleSetRuleConditionsRemoteAddressList
	RemoteAddressInput() interface{}
	RequestBody() CdnFrontdoorBatchRuleSetRuleConditionsRequestBodyList
	RequestBodyInput() interface{}
	RequestCookies() CdnFrontdoorBatchRuleSetRuleConditionsRequestCookiesList
	RequestCookiesInput() interface{}
	RequestFileExtension() CdnFrontdoorBatchRuleSetRuleConditionsRequestFileExtensionList
	RequestFileExtensionInput() interface{}
	RequestFilename() CdnFrontdoorBatchRuleSetRuleConditionsRequestFilenameList
	RequestFilenameInput() interface{}
	RequestHeader() CdnFrontdoorBatchRuleSetRuleConditionsRequestHeaderList
	RequestHeaderInput() interface{}
	RequestMethod() CdnFrontdoorBatchRuleSetRuleConditionsRequestMethodList
	RequestMethodInput() interface{}
	RequestPath() CdnFrontdoorBatchRuleSetRuleConditionsRequestPathList
	RequestPathInput() interface{}
	RequestScheme() CdnFrontdoorBatchRuleSetRuleConditionsRequestSchemeList
	RequestSchemeInput() interface{}
	RequestUrl() CdnFrontdoorBatchRuleSetRuleConditionsRequestUrlList
	RequestUrlInput() interface{}
	ServerPort() CdnFrontdoorBatchRuleSetRuleConditionsServerPortList
	ServerPortInput() interface{}
	SocketAddress() CdnFrontdoorBatchRuleSetRuleConditionsSocketAddressList
	SocketAddressInput() interface{}
	SslProtocol() CdnFrontdoorBatchRuleSetRuleConditionsSslProtocolList
	SslProtocolInput() interface{}
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
	PutClientPort(value interface{})
	PutDeviceType(value interface{})
	PutHostName(value interface{})
	PutHttpVersion(value interface{})
	PutPostArgument(value interface{})
	PutQueryString(value interface{})
	PutRemoteAddress(value interface{})
	PutRequestBody(value interface{})
	PutRequestCookies(value interface{})
	PutRequestFileExtension(value interface{})
	PutRequestFilename(value interface{})
	PutRequestHeader(value interface{})
	PutRequestMethod(value interface{})
	PutRequestPath(value interface{})
	PutRequestScheme(value interface{})
	PutRequestUrl(value interface{})
	PutServerPort(value interface{})
	PutSocketAddress(value interface{})
	PutSslProtocol(value interface{})
	ResetClientPort()
	ResetDeviceType()
	ResetHostName()
	ResetHttpVersion()
	ResetPostArgument()
	ResetQueryString()
	ResetRemoteAddress()
	ResetRequestBody()
	ResetRequestCookies()
	ResetRequestFileExtension()
	ResetRequestFilename()
	ResetRequestHeader()
	ResetRequestMethod()
	ResetRequestPath()
	ResetRequestScheme()
	ResetRequestUrl()
	ResetServerPort()
	ResetSocketAddress()
	ResetSslProtocol()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktn.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for CdnFrontdoorBatchRuleSetRuleConditionsOutputReference
type jsiiProxy_CdnFrontdoorBatchRuleSetRuleConditionsOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_CdnFrontdoorBatchRuleSetRuleConditionsOutputReference) ClientPort() CdnFrontdoorBatchRuleSetRuleConditionsClientPortList {
	var returns CdnFrontdoorBatchRuleSetRuleConditionsClientPortList
	_jsii_.Get(
		j,
		"clientPort",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CdnFrontdoorBatchRuleSetRuleConditionsOutputReference) ClientPortInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"clientPortInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CdnFrontdoorBatchRuleSetRuleConditionsOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CdnFrontdoorBatchRuleSetRuleConditionsOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CdnFrontdoorBatchRuleSetRuleConditionsOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CdnFrontdoorBatchRuleSetRuleConditionsOutputReference) DeviceType() CdnFrontdoorBatchRuleSetRuleConditionsDeviceTypeList {
	var returns CdnFrontdoorBatchRuleSetRuleConditionsDeviceTypeList
	_jsii_.Get(
		j,
		"deviceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CdnFrontdoorBatchRuleSetRuleConditionsOutputReference) DeviceTypeInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"deviceTypeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CdnFrontdoorBatchRuleSetRuleConditionsOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CdnFrontdoorBatchRuleSetRuleConditionsOutputReference) HostName() CdnFrontdoorBatchRuleSetRuleConditionsHostNameList {
	var returns CdnFrontdoorBatchRuleSetRuleConditionsHostNameList
	_jsii_.Get(
		j,
		"hostName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CdnFrontdoorBatchRuleSetRuleConditionsOutputReference) HostNameInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"hostNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CdnFrontdoorBatchRuleSetRuleConditionsOutputReference) HttpVersion() CdnFrontdoorBatchRuleSetRuleConditionsHttpVersionList {
	var returns CdnFrontdoorBatchRuleSetRuleConditionsHttpVersionList
	_jsii_.Get(
		j,
		"httpVersion",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CdnFrontdoorBatchRuleSetRuleConditionsOutputReference) HttpVersionInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"httpVersionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CdnFrontdoorBatchRuleSetRuleConditionsOutputReference) InternalValue() *CdnFrontdoorBatchRuleSetRuleConditions {
	var returns *CdnFrontdoorBatchRuleSetRuleConditions
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CdnFrontdoorBatchRuleSetRuleConditionsOutputReference) PostArgument() CdnFrontdoorBatchRuleSetRuleConditionsPostArgumentList {
	var returns CdnFrontdoorBatchRuleSetRuleConditionsPostArgumentList
	_jsii_.Get(
		j,
		"postArgument",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CdnFrontdoorBatchRuleSetRuleConditionsOutputReference) PostArgumentInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"postArgumentInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CdnFrontdoorBatchRuleSetRuleConditionsOutputReference) QueryString() CdnFrontdoorBatchRuleSetRuleConditionsQueryStringList {
	var returns CdnFrontdoorBatchRuleSetRuleConditionsQueryStringList
	_jsii_.Get(
		j,
		"queryString",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CdnFrontdoorBatchRuleSetRuleConditionsOutputReference) QueryStringInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"queryStringInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CdnFrontdoorBatchRuleSetRuleConditionsOutputReference) RemoteAddress() CdnFrontdoorBatchRuleSetRuleConditionsRemoteAddressList {
	var returns CdnFrontdoorBatchRuleSetRuleConditionsRemoteAddressList
	_jsii_.Get(
		j,
		"remoteAddress",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CdnFrontdoorBatchRuleSetRuleConditionsOutputReference) RemoteAddressInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"remoteAddressInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CdnFrontdoorBatchRuleSetRuleConditionsOutputReference) RequestBody() CdnFrontdoorBatchRuleSetRuleConditionsRequestBodyList {
	var returns CdnFrontdoorBatchRuleSetRuleConditionsRequestBodyList
	_jsii_.Get(
		j,
		"requestBody",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CdnFrontdoorBatchRuleSetRuleConditionsOutputReference) RequestBodyInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"requestBodyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CdnFrontdoorBatchRuleSetRuleConditionsOutputReference) RequestCookies() CdnFrontdoorBatchRuleSetRuleConditionsRequestCookiesList {
	var returns CdnFrontdoorBatchRuleSetRuleConditionsRequestCookiesList
	_jsii_.Get(
		j,
		"requestCookies",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CdnFrontdoorBatchRuleSetRuleConditionsOutputReference) RequestCookiesInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"requestCookiesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CdnFrontdoorBatchRuleSetRuleConditionsOutputReference) RequestFileExtension() CdnFrontdoorBatchRuleSetRuleConditionsRequestFileExtensionList {
	var returns CdnFrontdoorBatchRuleSetRuleConditionsRequestFileExtensionList
	_jsii_.Get(
		j,
		"requestFileExtension",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CdnFrontdoorBatchRuleSetRuleConditionsOutputReference) RequestFileExtensionInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"requestFileExtensionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CdnFrontdoorBatchRuleSetRuleConditionsOutputReference) RequestFilename() CdnFrontdoorBatchRuleSetRuleConditionsRequestFilenameList {
	var returns CdnFrontdoorBatchRuleSetRuleConditionsRequestFilenameList
	_jsii_.Get(
		j,
		"requestFilename",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CdnFrontdoorBatchRuleSetRuleConditionsOutputReference) RequestFilenameInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"requestFilenameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CdnFrontdoorBatchRuleSetRuleConditionsOutputReference) RequestHeader() CdnFrontdoorBatchRuleSetRuleConditionsRequestHeaderList {
	var returns CdnFrontdoorBatchRuleSetRuleConditionsRequestHeaderList
	_jsii_.Get(
		j,
		"requestHeader",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CdnFrontdoorBatchRuleSetRuleConditionsOutputReference) RequestHeaderInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"requestHeaderInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CdnFrontdoorBatchRuleSetRuleConditionsOutputReference) RequestMethod() CdnFrontdoorBatchRuleSetRuleConditionsRequestMethodList {
	var returns CdnFrontdoorBatchRuleSetRuleConditionsRequestMethodList
	_jsii_.Get(
		j,
		"requestMethod",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CdnFrontdoorBatchRuleSetRuleConditionsOutputReference) RequestMethodInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"requestMethodInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CdnFrontdoorBatchRuleSetRuleConditionsOutputReference) RequestPath() CdnFrontdoorBatchRuleSetRuleConditionsRequestPathList {
	var returns CdnFrontdoorBatchRuleSetRuleConditionsRequestPathList
	_jsii_.Get(
		j,
		"requestPath",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CdnFrontdoorBatchRuleSetRuleConditionsOutputReference) RequestPathInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"requestPathInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CdnFrontdoorBatchRuleSetRuleConditionsOutputReference) RequestScheme() CdnFrontdoorBatchRuleSetRuleConditionsRequestSchemeList {
	var returns CdnFrontdoorBatchRuleSetRuleConditionsRequestSchemeList
	_jsii_.Get(
		j,
		"requestScheme",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CdnFrontdoorBatchRuleSetRuleConditionsOutputReference) RequestSchemeInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"requestSchemeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CdnFrontdoorBatchRuleSetRuleConditionsOutputReference) RequestUrl() CdnFrontdoorBatchRuleSetRuleConditionsRequestUrlList {
	var returns CdnFrontdoorBatchRuleSetRuleConditionsRequestUrlList
	_jsii_.Get(
		j,
		"requestUrl",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CdnFrontdoorBatchRuleSetRuleConditionsOutputReference) RequestUrlInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"requestUrlInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CdnFrontdoorBatchRuleSetRuleConditionsOutputReference) ServerPort() CdnFrontdoorBatchRuleSetRuleConditionsServerPortList {
	var returns CdnFrontdoorBatchRuleSetRuleConditionsServerPortList
	_jsii_.Get(
		j,
		"serverPort",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CdnFrontdoorBatchRuleSetRuleConditionsOutputReference) ServerPortInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"serverPortInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CdnFrontdoorBatchRuleSetRuleConditionsOutputReference) SocketAddress() CdnFrontdoorBatchRuleSetRuleConditionsSocketAddressList {
	var returns CdnFrontdoorBatchRuleSetRuleConditionsSocketAddressList
	_jsii_.Get(
		j,
		"socketAddress",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CdnFrontdoorBatchRuleSetRuleConditionsOutputReference) SocketAddressInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"socketAddressInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CdnFrontdoorBatchRuleSetRuleConditionsOutputReference) SslProtocol() CdnFrontdoorBatchRuleSetRuleConditionsSslProtocolList {
	var returns CdnFrontdoorBatchRuleSetRuleConditionsSslProtocolList
	_jsii_.Get(
		j,
		"sslProtocol",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CdnFrontdoorBatchRuleSetRuleConditionsOutputReference) SslProtocolInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"sslProtocolInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CdnFrontdoorBatchRuleSetRuleConditionsOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CdnFrontdoorBatchRuleSetRuleConditionsOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewCdnFrontdoorBatchRuleSetRuleConditionsOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) CdnFrontdoorBatchRuleSetRuleConditionsOutputReference {
	_init_.Initialize()

	if err := validateNewCdnFrontdoorBatchRuleSetRuleConditionsOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_CdnFrontdoorBatchRuleSetRuleConditionsOutputReference{}

	_jsii_.Create(
		"@cdktn/provider-azurerm.cdnFrontdoorBatchRuleSet.CdnFrontdoorBatchRuleSetRuleConditionsOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewCdnFrontdoorBatchRuleSetRuleConditionsOutputReference_Override(c CdnFrontdoorBatchRuleSetRuleConditionsOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/provider-azurerm.cdnFrontdoorBatchRuleSet.CdnFrontdoorBatchRuleSetRuleConditionsOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		c,
	)
}

func (j *jsiiProxy_CdnFrontdoorBatchRuleSetRuleConditionsOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_CdnFrontdoorBatchRuleSetRuleConditionsOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_CdnFrontdoorBatchRuleSetRuleConditionsOutputReference)SetInternalValue(val *CdnFrontdoorBatchRuleSetRuleConditions) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_CdnFrontdoorBatchRuleSetRuleConditionsOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_CdnFrontdoorBatchRuleSetRuleConditionsOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (c *jsiiProxy_CdnFrontdoorBatchRuleSetRuleConditionsOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CdnFrontdoorBatchRuleSetRuleConditionsOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	if err := c.validateGetAnyMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]interface{}

	_jsii_.Invoke(
		c,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CdnFrontdoorBatchRuleSetRuleConditionsOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
	if err := c.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		c,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CdnFrontdoorBatchRuleSetRuleConditionsOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	if err := c.validateGetBooleanMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*bool

	_jsii_.Invoke(
		c,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CdnFrontdoorBatchRuleSetRuleConditionsOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	if err := c.validateGetListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*string

	_jsii_.Invoke(
		c,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CdnFrontdoorBatchRuleSetRuleConditionsOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	if err := c.validateGetNumberAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *float64

	_jsii_.Invoke(
		c,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CdnFrontdoorBatchRuleSetRuleConditionsOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	if err := c.validateGetNumberListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*float64

	_jsii_.Invoke(
		c,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CdnFrontdoorBatchRuleSetRuleConditionsOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	if err := c.validateGetNumberMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*float64

	_jsii_.Invoke(
		c,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CdnFrontdoorBatchRuleSetRuleConditionsOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	if err := c.validateGetStringAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		c,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CdnFrontdoorBatchRuleSetRuleConditionsOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	if err := c.validateGetStringMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*string

	_jsii_.Invoke(
		c,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CdnFrontdoorBatchRuleSetRuleConditionsOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		c,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CdnFrontdoorBatchRuleSetRuleConditionsOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
	if err := c.validateInterpolationForAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		c,
		"interpolationForAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CdnFrontdoorBatchRuleSetRuleConditionsOutputReference) PutClientPort(value interface{}) {
	if err := c.validatePutClientPortParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"putClientPort",
		[]interface{}{value},
	)
}

func (c *jsiiProxy_CdnFrontdoorBatchRuleSetRuleConditionsOutputReference) PutDeviceType(value interface{}) {
	if err := c.validatePutDeviceTypeParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"putDeviceType",
		[]interface{}{value},
	)
}

func (c *jsiiProxy_CdnFrontdoorBatchRuleSetRuleConditionsOutputReference) PutHostName(value interface{}) {
	if err := c.validatePutHostNameParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"putHostName",
		[]interface{}{value},
	)
}

func (c *jsiiProxy_CdnFrontdoorBatchRuleSetRuleConditionsOutputReference) PutHttpVersion(value interface{}) {
	if err := c.validatePutHttpVersionParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"putHttpVersion",
		[]interface{}{value},
	)
}

func (c *jsiiProxy_CdnFrontdoorBatchRuleSetRuleConditionsOutputReference) PutPostArgument(value interface{}) {
	if err := c.validatePutPostArgumentParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"putPostArgument",
		[]interface{}{value},
	)
}

func (c *jsiiProxy_CdnFrontdoorBatchRuleSetRuleConditionsOutputReference) PutQueryString(value interface{}) {
	if err := c.validatePutQueryStringParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"putQueryString",
		[]interface{}{value},
	)
}

func (c *jsiiProxy_CdnFrontdoorBatchRuleSetRuleConditionsOutputReference) PutRemoteAddress(value interface{}) {
	if err := c.validatePutRemoteAddressParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"putRemoteAddress",
		[]interface{}{value},
	)
}

func (c *jsiiProxy_CdnFrontdoorBatchRuleSetRuleConditionsOutputReference) PutRequestBody(value interface{}) {
	if err := c.validatePutRequestBodyParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"putRequestBody",
		[]interface{}{value},
	)
}

func (c *jsiiProxy_CdnFrontdoorBatchRuleSetRuleConditionsOutputReference) PutRequestCookies(value interface{}) {
	if err := c.validatePutRequestCookiesParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"putRequestCookies",
		[]interface{}{value},
	)
}

func (c *jsiiProxy_CdnFrontdoorBatchRuleSetRuleConditionsOutputReference) PutRequestFileExtension(value interface{}) {
	if err := c.validatePutRequestFileExtensionParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"putRequestFileExtension",
		[]interface{}{value},
	)
}

func (c *jsiiProxy_CdnFrontdoorBatchRuleSetRuleConditionsOutputReference) PutRequestFilename(value interface{}) {
	if err := c.validatePutRequestFilenameParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"putRequestFilename",
		[]interface{}{value},
	)
}

func (c *jsiiProxy_CdnFrontdoorBatchRuleSetRuleConditionsOutputReference) PutRequestHeader(value interface{}) {
	if err := c.validatePutRequestHeaderParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"putRequestHeader",
		[]interface{}{value},
	)
}

func (c *jsiiProxy_CdnFrontdoorBatchRuleSetRuleConditionsOutputReference) PutRequestMethod(value interface{}) {
	if err := c.validatePutRequestMethodParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"putRequestMethod",
		[]interface{}{value},
	)
}

func (c *jsiiProxy_CdnFrontdoorBatchRuleSetRuleConditionsOutputReference) PutRequestPath(value interface{}) {
	if err := c.validatePutRequestPathParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"putRequestPath",
		[]interface{}{value},
	)
}

func (c *jsiiProxy_CdnFrontdoorBatchRuleSetRuleConditionsOutputReference) PutRequestScheme(value interface{}) {
	if err := c.validatePutRequestSchemeParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"putRequestScheme",
		[]interface{}{value},
	)
}

func (c *jsiiProxy_CdnFrontdoorBatchRuleSetRuleConditionsOutputReference) PutRequestUrl(value interface{}) {
	if err := c.validatePutRequestUrlParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"putRequestUrl",
		[]interface{}{value},
	)
}

func (c *jsiiProxy_CdnFrontdoorBatchRuleSetRuleConditionsOutputReference) PutServerPort(value interface{}) {
	if err := c.validatePutServerPortParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"putServerPort",
		[]interface{}{value},
	)
}

func (c *jsiiProxy_CdnFrontdoorBatchRuleSetRuleConditionsOutputReference) PutSocketAddress(value interface{}) {
	if err := c.validatePutSocketAddressParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"putSocketAddress",
		[]interface{}{value},
	)
}

func (c *jsiiProxy_CdnFrontdoorBatchRuleSetRuleConditionsOutputReference) PutSslProtocol(value interface{}) {
	if err := c.validatePutSslProtocolParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"putSslProtocol",
		[]interface{}{value},
	)
}

func (c *jsiiProxy_CdnFrontdoorBatchRuleSetRuleConditionsOutputReference) ResetClientPort() {
	_jsii_.InvokeVoid(
		c,
		"resetClientPort",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CdnFrontdoorBatchRuleSetRuleConditionsOutputReference) ResetDeviceType() {
	_jsii_.InvokeVoid(
		c,
		"resetDeviceType",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CdnFrontdoorBatchRuleSetRuleConditionsOutputReference) ResetHostName() {
	_jsii_.InvokeVoid(
		c,
		"resetHostName",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CdnFrontdoorBatchRuleSetRuleConditionsOutputReference) ResetHttpVersion() {
	_jsii_.InvokeVoid(
		c,
		"resetHttpVersion",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CdnFrontdoorBatchRuleSetRuleConditionsOutputReference) ResetPostArgument() {
	_jsii_.InvokeVoid(
		c,
		"resetPostArgument",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CdnFrontdoorBatchRuleSetRuleConditionsOutputReference) ResetQueryString() {
	_jsii_.InvokeVoid(
		c,
		"resetQueryString",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CdnFrontdoorBatchRuleSetRuleConditionsOutputReference) ResetRemoteAddress() {
	_jsii_.InvokeVoid(
		c,
		"resetRemoteAddress",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CdnFrontdoorBatchRuleSetRuleConditionsOutputReference) ResetRequestBody() {
	_jsii_.InvokeVoid(
		c,
		"resetRequestBody",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CdnFrontdoorBatchRuleSetRuleConditionsOutputReference) ResetRequestCookies() {
	_jsii_.InvokeVoid(
		c,
		"resetRequestCookies",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CdnFrontdoorBatchRuleSetRuleConditionsOutputReference) ResetRequestFileExtension() {
	_jsii_.InvokeVoid(
		c,
		"resetRequestFileExtension",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CdnFrontdoorBatchRuleSetRuleConditionsOutputReference) ResetRequestFilename() {
	_jsii_.InvokeVoid(
		c,
		"resetRequestFilename",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CdnFrontdoorBatchRuleSetRuleConditionsOutputReference) ResetRequestHeader() {
	_jsii_.InvokeVoid(
		c,
		"resetRequestHeader",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CdnFrontdoorBatchRuleSetRuleConditionsOutputReference) ResetRequestMethod() {
	_jsii_.InvokeVoid(
		c,
		"resetRequestMethod",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CdnFrontdoorBatchRuleSetRuleConditionsOutputReference) ResetRequestPath() {
	_jsii_.InvokeVoid(
		c,
		"resetRequestPath",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CdnFrontdoorBatchRuleSetRuleConditionsOutputReference) ResetRequestScheme() {
	_jsii_.InvokeVoid(
		c,
		"resetRequestScheme",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CdnFrontdoorBatchRuleSetRuleConditionsOutputReference) ResetRequestUrl() {
	_jsii_.InvokeVoid(
		c,
		"resetRequestUrl",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CdnFrontdoorBatchRuleSetRuleConditionsOutputReference) ResetServerPort() {
	_jsii_.InvokeVoid(
		c,
		"resetServerPort",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CdnFrontdoorBatchRuleSetRuleConditionsOutputReference) ResetSocketAddress() {
	_jsii_.InvokeVoid(
		c,
		"resetSocketAddress",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CdnFrontdoorBatchRuleSetRuleConditionsOutputReference) ResetSslProtocol() {
	_jsii_.InvokeVoid(
		c,
		"resetSslProtocol",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CdnFrontdoorBatchRuleSetRuleConditionsOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
	if err := c.validateResolveParameters(context); err != nil {
		panic(err)
	}
	var returns interface{}

	_jsii_.Invoke(
		c,
		"resolve",
		[]interface{}{context},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CdnFrontdoorBatchRuleSetRuleConditionsOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

