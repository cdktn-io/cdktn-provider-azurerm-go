// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package cdnfrontdoorrule

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-provider-azurerm-go/azurerm/v18/jsii"

	"github.com/cdktn-io/cdktn-provider-azurerm-go/azurerm/v18/cdnfrontdoorrule/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type CdnFrontdoorRuleConditionsOutputReference interface {
	cdktn.ComplexObject
	ClientPort() CdnFrontdoorRuleConditionsClientPortList
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
	DeviceType() CdnFrontdoorRuleConditionsDeviceTypeList
	DeviceTypeInput() interface{}
	// Experimental.
	Fqn() *string
	HostName() CdnFrontdoorRuleConditionsHostNameList
	HostNameInput() interface{}
	HttpVersion() CdnFrontdoorRuleConditionsHttpVersionList
	HttpVersionInput() interface{}
	InternalValue() *CdnFrontdoorRuleConditions
	SetInternalValue(val *CdnFrontdoorRuleConditions)
	PostArgument() CdnFrontdoorRuleConditionsPostArgumentList
	PostArgumentInput() interface{}
	QueryString() CdnFrontdoorRuleConditionsQueryStringList
	QueryStringInput() interface{}
	RemoteAddress() CdnFrontdoorRuleConditionsRemoteAddressList
	RemoteAddressInput() interface{}
	RequestBody() CdnFrontdoorRuleConditionsRequestBodyList
	RequestBodyInput() interface{}
	RequestCookies() CdnFrontdoorRuleConditionsRequestCookiesList
	RequestCookiesInput() interface{}
	RequestFileExtension() CdnFrontdoorRuleConditionsRequestFileExtensionList
	RequestFileExtensionInput() interface{}
	RequestFilename() CdnFrontdoorRuleConditionsRequestFilenameList
	RequestFilenameInput() interface{}
	RequestHeader() CdnFrontdoorRuleConditionsRequestHeaderList
	RequestHeaderInput() interface{}
	RequestMethod() CdnFrontdoorRuleConditionsRequestMethodList
	RequestMethodInput() interface{}
	RequestPath() CdnFrontdoorRuleConditionsRequestPathList
	RequestPathInput() interface{}
	RequestScheme() CdnFrontdoorRuleConditionsRequestSchemeList
	RequestSchemeInput() interface{}
	RequestUrl() CdnFrontdoorRuleConditionsRequestUrlList
	RequestUrlInput() interface{}
	ServerPort() CdnFrontdoorRuleConditionsServerPortList
	ServerPortInput() interface{}
	SocketAddress() CdnFrontdoorRuleConditionsSocketAddressList
	SocketAddressInput() interface{}
	SslProtocol() CdnFrontdoorRuleConditionsSslProtocolList
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

// The jsii proxy struct for CdnFrontdoorRuleConditionsOutputReference
type jsiiProxy_CdnFrontdoorRuleConditionsOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_CdnFrontdoorRuleConditionsOutputReference) ClientPort() CdnFrontdoorRuleConditionsClientPortList {
	var returns CdnFrontdoorRuleConditionsClientPortList
	_jsii_.Get(
		j,
		"clientPort",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CdnFrontdoorRuleConditionsOutputReference) ClientPortInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"clientPortInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CdnFrontdoorRuleConditionsOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CdnFrontdoorRuleConditionsOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CdnFrontdoorRuleConditionsOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CdnFrontdoorRuleConditionsOutputReference) DeviceType() CdnFrontdoorRuleConditionsDeviceTypeList {
	var returns CdnFrontdoorRuleConditionsDeviceTypeList
	_jsii_.Get(
		j,
		"deviceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CdnFrontdoorRuleConditionsOutputReference) DeviceTypeInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"deviceTypeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CdnFrontdoorRuleConditionsOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CdnFrontdoorRuleConditionsOutputReference) HostName() CdnFrontdoorRuleConditionsHostNameList {
	var returns CdnFrontdoorRuleConditionsHostNameList
	_jsii_.Get(
		j,
		"hostName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CdnFrontdoorRuleConditionsOutputReference) HostNameInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"hostNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CdnFrontdoorRuleConditionsOutputReference) HttpVersion() CdnFrontdoorRuleConditionsHttpVersionList {
	var returns CdnFrontdoorRuleConditionsHttpVersionList
	_jsii_.Get(
		j,
		"httpVersion",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CdnFrontdoorRuleConditionsOutputReference) HttpVersionInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"httpVersionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CdnFrontdoorRuleConditionsOutputReference) InternalValue() *CdnFrontdoorRuleConditions {
	var returns *CdnFrontdoorRuleConditions
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CdnFrontdoorRuleConditionsOutputReference) PostArgument() CdnFrontdoorRuleConditionsPostArgumentList {
	var returns CdnFrontdoorRuleConditionsPostArgumentList
	_jsii_.Get(
		j,
		"postArgument",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CdnFrontdoorRuleConditionsOutputReference) PostArgumentInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"postArgumentInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CdnFrontdoorRuleConditionsOutputReference) QueryString() CdnFrontdoorRuleConditionsQueryStringList {
	var returns CdnFrontdoorRuleConditionsQueryStringList
	_jsii_.Get(
		j,
		"queryString",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CdnFrontdoorRuleConditionsOutputReference) QueryStringInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"queryStringInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CdnFrontdoorRuleConditionsOutputReference) RemoteAddress() CdnFrontdoorRuleConditionsRemoteAddressList {
	var returns CdnFrontdoorRuleConditionsRemoteAddressList
	_jsii_.Get(
		j,
		"remoteAddress",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CdnFrontdoorRuleConditionsOutputReference) RemoteAddressInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"remoteAddressInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CdnFrontdoorRuleConditionsOutputReference) RequestBody() CdnFrontdoorRuleConditionsRequestBodyList {
	var returns CdnFrontdoorRuleConditionsRequestBodyList
	_jsii_.Get(
		j,
		"requestBody",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CdnFrontdoorRuleConditionsOutputReference) RequestBodyInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"requestBodyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CdnFrontdoorRuleConditionsOutputReference) RequestCookies() CdnFrontdoorRuleConditionsRequestCookiesList {
	var returns CdnFrontdoorRuleConditionsRequestCookiesList
	_jsii_.Get(
		j,
		"requestCookies",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CdnFrontdoorRuleConditionsOutputReference) RequestCookiesInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"requestCookiesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CdnFrontdoorRuleConditionsOutputReference) RequestFileExtension() CdnFrontdoorRuleConditionsRequestFileExtensionList {
	var returns CdnFrontdoorRuleConditionsRequestFileExtensionList
	_jsii_.Get(
		j,
		"requestFileExtension",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CdnFrontdoorRuleConditionsOutputReference) RequestFileExtensionInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"requestFileExtensionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CdnFrontdoorRuleConditionsOutputReference) RequestFilename() CdnFrontdoorRuleConditionsRequestFilenameList {
	var returns CdnFrontdoorRuleConditionsRequestFilenameList
	_jsii_.Get(
		j,
		"requestFilename",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CdnFrontdoorRuleConditionsOutputReference) RequestFilenameInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"requestFilenameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CdnFrontdoorRuleConditionsOutputReference) RequestHeader() CdnFrontdoorRuleConditionsRequestHeaderList {
	var returns CdnFrontdoorRuleConditionsRequestHeaderList
	_jsii_.Get(
		j,
		"requestHeader",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CdnFrontdoorRuleConditionsOutputReference) RequestHeaderInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"requestHeaderInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CdnFrontdoorRuleConditionsOutputReference) RequestMethod() CdnFrontdoorRuleConditionsRequestMethodList {
	var returns CdnFrontdoorRuleConditionsRequestMethodList
	_jsii_.Get(
		j,
		"requestMethod",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CdnFrontdoorRuleConditionsOutputReference) RequestMethodInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"requestMethodInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CdnFrontdoorRuleConditionsOutputReference) RequestPath() CdnFrontdoorRuleConditionsRequestPathList {
	var returns CdnFrontdoorRuleConditionsRequestPathList
	_jsii_.Get(
		j,
		"requestPath",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CdnFrontdoorRuleConditionsOutputReference) RequestPathInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"requestPathInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CdnFrontdoorRuleConditionsOutputReference) RequestScheme() CdnFrontdoorRuleConditionsRequestSchemeList {
	var returns CdnFrontdoorRuleConditionsRequestSchemeList
	_jsii_.Get(
		j,
		"requestScheme",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CdnFrontdoorRuleConditionsOutputReference) RequestSchemeInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"requestSchemeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CdnFrontdoorRuleConditionsOutputReference) RequestUrl() CdnFrontdoorRuleConditionsRequestUrlList {
	var returns CdnFrontdoorRuleConditionsRequestUrlList
	_jsii_.Get(
		j,
		"requestUrl",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CdnFrontdoorRuleConditionsOutputReference) RequestUrlInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"requestUrlInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CdnFrontdoorRuleConditionsOutputReference) ServerPort() CdnFrontdoorRuleConditionsServerPortList {
	var returns CdnFrontdoorRuleConditionsServerPortList
	_jsii_.Get(
		j,
		"serverPort",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CdnFrontdoorRuleConditionsOutputReference) ServerPortInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"serverPortInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CdnFrontdoorRuleConditionsOutputReference) SocketAddress() CdnFrontdoorRuleConditionsSocketAddressList {
	var returns CdnFrontdoorRuleConditionsSocketAddressList
	_jsii_.Get(
		j,
		"socketAddress",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CdnFrontdoorRuleConditionsOutputReference) SocketAddressInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"socketAddressInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CdnFrontdoorRuleConditionsOutputReference) SslProtocol() CdnFrontdoorRuleConditionsSslProtocolList {
	var returns CdnFrontdoorRuleConditionsSslProtocolList
	_jsii_.Get(
		j,
		"sslProtocol",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CdnFrontdoorRuleConditionsOutputReference) SslProtocolInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"sslProtocolInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CdnFrontdoorRuleConditionsOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CdnFrontdoorRuleConditionsOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewCdnFrontdoorRuleConditionsOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) CdnFrontdoorRuleConditionsOutputReference {
	_init_.Initialize()

	if err := validateNewCdnFrontdoorRuleConditionsOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_CdnFrontdoorRuleConditionsOutputReference{}

	_jsii_.Create(
		"@cdktn/provider-azurerm.cdnFrontdoorRule.CdnFrontdoorRuleConditionsOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewCdnFrontdoorRuleConditionsOutputReference_Override(c CdnFrontdoorRuleConditionsOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/provider-azurerm.cdnFrontdoorRule.CdnFrontdoorRuleConditionsOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		c,
	)
}

func (j *jsiiProxy_CdnFrontdoorRuleConditionsOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_CdnFrontdoorRuleConditionsOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_CdnFrontdoorRuleConditionsOutputReference)SetInternalValue(val *CdnFrontdoorRuleConditions) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_CdnFrontdoorRuleConditionsOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_CdnFrontdoorRuleConditionsOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (c *jsiiProxy_CdnFrontdoorRuleConditionsOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CdnFrontdoorRuleConditionsOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (c *jsiiProxy_CdnFrontdoorRuleConditionsOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (c *jsiiProxy_CdnFrontdoorRuleConditionsOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (c *jsiiProxy_CdnFrontdoorRuleConditionsOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (c *jsiiProxy_CdnFrontdoorRuleConditionsOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (c *jsiiProxy_CdnFrontdoorRuleConditionsOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (c *jsiiProxy_CdnFrontdoorRuleConditionsOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (c *jsiiProxy_CdnFrontdoorRuleConditionsOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (c *jsiiProxy_CdnFrontdoorRuleConditionsOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (c *jsiiProxy_CdnFrontdoorRuleConditionsOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		c,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CdnFrontdoorRuleConditionsOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (c *jsiiProxy_CdnFrontdoorRuleConditionsOutputReference) PutClientPort(value interface{}) {
	if err := c.validatePutClientPortParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"putClientPort",
		[]interface{}{value},
	)
}

func (c *jsiiProxy_CdnFrontdoorRuleConditionsOutputReference) PutDeviceType(value interface{}) {
	if err := c.validatePutDeviceTypeParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"putDeviceType",
		[]interface{}{value},
	)
}

func (c *jsiiProxy_CdnFrontdoorRuleConditionsOutputReference) PutHostName(value interface{}) {
	if err := c.validatePutHostNameParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"putHostName",
		[]interface{}{value},
	)
}

func (c *jsiiProxy_CdnFrontdoorRuleConditionsOutputReference) PutHttpVersion(value interface{}) {
	if err := c.validatePutHttpVersionParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"putHttpVersion",
		[]interface{}{value},
	)
}

func (c *jsiiProxy_CdnFrontdoorRuleConditionsOutputReference) PutPostArgument(value interface{}) {
	if err := c.validatePutPostArgumentParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"putPostArgument",
		[]interface{}{value},
	)
}

func (c *jsiiProxy_CdnFrontdoorRuleConditionsOutputReference) PutQueryString(value interface{}) {
	if err := c.validatePutQueryStringParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"putQueryString",
		[]interface{}{value},
	)
}

func (c *jsiiProxy_CdnFrontdoorRuleConditionsOutputReference) PutRemoteAddress(value interface{}) {
	if err := c.validatePutRemoteAddressParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"putRemoteAddress",
		[]interface{}{value},
	)
}

func (c *jsiiProxy_CdnFrontdoorRuleConditionsOutputReference) PutRequestBody(value interface{}) {
	if err := c.validatePutRequestBodyParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"putRequestBody",
		[]interface{}{value},
	)
}

func (c *jsiiProxy_CdnFrontdoorRuleConditionsOutputReference) PutRequestCookies(value interface{}) {
	if err := c.validatePutRequestCookiesParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"putRequestCookies",
		[]interface{}{value},
	)
}

func (c *jsiiProxy_CdnFrontdoorRuleConditionsOutputReference) PutRequestFileExtension(value interface{}) {
	if err := c.validatePutRequestFileExtensionParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"putRequestFileExtension",
		[]interface{}{value},
	)
}

func (c *jsiiProxy_CdnFrontdoorRuleConditionsOutputReference) PutRequestFilename(value interface{}) {
	if err := c.validatePutRequestFilenameParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"putRequestFilename",
		[]interface{}{value},
	)
}

func (c *jsiiProxy_CdnFrontdoorRuleConditionsOutputReference) PutRequestHeader(value interface{}) {
	if err := c.validatePutRequestHeaderParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"putRequestHeader",
		[]interface{}{value},
	)
}

func (c *jsiiProxy_CdnFrontdoorRuleConditionsOutputReference) PutRequestMethod(value interface{}) {
	if err := c.validatePutRequestMethodParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"putRequestMethod",
		[]interface{}{value},
	)
}

func (c *jsiiProxy_CdnFrontdoorRuleConditionsOutputReference) PutRequestPath(value interface{}) {
	if err := c.validatePutRequestPathParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"putRequestPath",
		[]interface{}{value},
	)
}

func (c *jsiiProxy_CdnFrontdoorRuleConditionsOutputReference) PutRequestScheme(value interface{}) {
	if err := c.validatePutRequestSchemeParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"putRequestScheme",
		[]interface{}{value},
	)
}

func (c *jsiiProxy_CdnFrontdoorRuleConditionsOutputReference) PutRequestUrl(value interface{}) {
	if err := c.validatePutRequestUrlParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"putRequestUrl",
		[]interface{}{value},
	)
}

func (c *jsiiProxy_CdnFrontdoorRuleConditionsOutputReference) PutServerPort(value interface{}) {
	if err := c.validatePutServerPortParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"putServerPort",
		[]interface{}{value},
	)
}

func (c *jsiiProxy_CdnFrontdoorRuleConditionsOutputReference) PutSocketAddress(value interface{}) {
	if err := c.validatePutSocketAddressParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"putSocketAddress",
		[]interface{}{value},
	)
}

func (c *jsiiProxy_CdnFrontdoorRuleConditionsOutputReference) PutSslProtocol(value interface{}) {
	if err := c.validatePutSslProtocolParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"putSslProtocol",
		[]interface{}{value},
	)
}

func (c *jsiiProxy_CdnFrontdoorRuleConditionsOutputReference) ResetClientPort() {
	_jsii_.InvokeVoid(
		c,
		"resetClientPort",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CdnFrontdoorRuleConditionsOutputReference) ResetDeviceType() {
	_jsii_.InvokeVoid(
		c,
		"resetDeviceType",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CdnFrontdoorRuleConditionsOutputReference) ResetHostName() {
	_jsii_.InvokeVoid(
		c,
		"resetHostName",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CdnFrontdoorRuleConditionsOutputReference) ResetHttpVersion() {
	_jsii_.InvokeVoid(
		c,
		"resetHttpVersion",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CdnFrontdoorRuleConditionsOutputReference) ResetPostArgument() {
	_jsii_.InvokeVoid(
		c,
		"resetPostArgument",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CdnFrontdoorRuleConditionsOutputReference) ResetQueryString() {
	_jsii_.InvokeVoid(
		c,
		"resetQueryString",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CdnFrontdoorRuleConditionsOutputReference) ResetRemoteAddress() {
	_jsii_.InvokeVoid(
		c,
		"resetRemoteAddress",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CdnFrontdoorRuleConditionsOutputReference) ResetRequestBody() {
	_jsii_.InvokeVoid(
		c,
		"resetRequestBody",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CdnFrontdoorRuleConditionsOutputReference) ResetRequestCookies() {
	_jsii_.InvokeVoid(
		c,
		"resetRequestCookies",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CdnFrontdoorRuleConditionsOutputReference) ResetRequestFileExtension() {
	_jsii_.InvokeVoid(
		c,
		"resetRequestFileExtension",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CdnFrontdoorRuleConditionsOutputReference) ResetRequestFilename() {
	_jsii_.InvokeVoid(
		c,
		"resetRequestFilename",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CdnFrontdoorRuleConditionsOutputReference) ResetRequestHeader() {
	_jsii_.InvokeVoid(
		c,
		"resetRequestHeader",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CdnFrontdoorRuleConditionsOutputReference) ResetRequestMethod() {
	_jsii_.InvokeVoid(
		c,
		"resetRequestMethod",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CdnFrontdoorRuleConditionsOutputReference) ResetRequestPath() {
	_jsii_.InvokeVoid(
		c,
		"resetRequestPath",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CdnFrontdoorRuleConditionsOutputReference) ResetRequestScheme() {
	_jsii_.InvokeVoid(
		c,
		"resetRequestScheme",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CdnFrontdoorRuleConditionsOutputReference) ResetRequestUrl() {
	_jsii_.InvokeVoid(
		c,
		"resetRequestUrl",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CdnFrontdoorRuleConditionsOutputReference) ResetServerPort() {
	_jsii_.InvokeVoid(
		c,
		"resetServerPort",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CdnFrontdoorRuleConditionsOutputReference) ResetSocketAddress() {
	_jsii_.InvokeVoid(
		c,
		"resetSocketAddress",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CdnFrontdoorRuleConditionsOutputReference) ResetSslProtocol() {
	_jsii_.InvokeVoid(
		c,
		"resetSslProtocol",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CdnFrontdoorRuleConditionsOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
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

func (c *jsiiProxy_CdnFrontdoorRuleConditionsOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

