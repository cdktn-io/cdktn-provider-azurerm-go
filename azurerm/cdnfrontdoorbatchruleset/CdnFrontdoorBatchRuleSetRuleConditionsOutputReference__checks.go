// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

//go:build !no_runtime_type_checking

package cdnfrontdoorbatchruleset

import (
	"fmt"

	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

func (c *jsiiProxy_CdnFrontdoorBatchRuleSetRuleConditionsOutputReference) validateGetAnyMapAttributeParameters(terraformAttribute *string) error {
	if terraformAttribute == nil {
		return fmt.Errorf("parameter terraformAttribute is required, but nil was provided")
	}

	return nil
}

func (c *jsiiProxy_CdnFrontdoorBatchRuleSetRuleConditionsOutputReference) validateGetBooleanAttributeParameters(terraformAttribute *string) error {
	if terraformAttribute == nil {
		return fmt.Errorf("parameter terraformAttribute is required, but nil was provided")
	}

	return nil
}

func (c *jsiiProxy_CdnFrontdoorBatchRuleSetRuleConditionsOutputReference) validateGetBooleanMapAttributeParameters(terraformAttribute *string) error {
	if terraformAttribute == nil {
		return fmt.Errorf("parameter terraformAttribute is required, but nil was provided")
	}

	return nil
}

func (c *jsiiProxy_CdnFrontdoorBatchRuleSetRuleConditionsOutputReference) validateGetListAttributeParameters(terraformAttribute *string) error {
	if terraformAttribute == nil {
		return fmt.Errorf("parameter terraformAttribute is required, but nil was provided")
	}

	return nil
}

func (c *jsiiProxy_CdnFrontdoorBatchRuleSetRuleConditionsOutputReference) validateGetNumberAttributeParameters(terraformAttribute *string) error {
	if terraformAttribute == nil {
		return fmt.Errorf("parameter terraformAttribute is required, but nil was provided")
	}

	return nil
}

func (c *jsiiProxy_CdnFrontdoorBatchRuleSetRuleConditionsOutputReference) validateGetNumberListAttributeParameters(terraformAttribute *string) error {
	if terraformAttribute == nil {
		return fmt.Errorf("parameter terraformAttribute is required, but nil was provided")
	}

	return nil
}

func (c *jsiiProxy_CdnFrontdoorBatchRuleSetRuleConditionsOutputReference) validateGetNumberMapAttributeParameters(terraformAttribute *string) error {
	if terraformAttribute == nil {
		return fmt.Errorf("parameter terraformAttribute is required, but nil was provided")
	}

	return nil
}

func (c *jsiiProxy_CdnFrontdoorBatchRuleSetRuleConditionsOutputReference) validateGetStringAttributeParameters(terraformAttribute *string) error {
	if terraformAttribute == nil {
		return fmt.Errorf("parameter terraformAttribute is required, but nil was provided")
	}

	return nil
}

func (c *jsiiProxy_CdnFrontdoorBatchRuleSetRuleConditionsOutputReference) validateGetStringMapAttributeParameters(terraformAttribute *string) error {
	if terraformAttribute == nil {
		return fmt.Errorf("parameter terraformAttribute is required, but nil was provided")
	}

	return nil
}

func (c *jsiiProxy_CdnFrontdoorBatchRuleSetRuleConditionsOutputReference) validateInterpolationForAttributeParameters(terraformAttribute *string) error {
	if terraformAttribute == nil {
		return fmt.Errorf("parameter terraformAttribute is required, but nil was provided")
	}

	return nil
}

func (c *jsiiProxy_CdnFrontdoorBatchRuleSetRuleConditionsOutputReference) validatePutClientPortParameters(value interface{}) error {
	if value == nil {
		return fmt.Errorf("parameter value is required, but nil was provided")
	}
	switch value.(type) {
	case cdktn.IResolvable:
		// ok
	case *[]*CdnFrontdoorBatchRuleSetRuleConditionsClientPort:
		value := value.(*[]*CdnFrontdoorBatchRuleSetRuleConditionsClientPort)
		for idx_cd4240, v := range *value {
			if err := _jsii_.ValidateStruct(v, func() string { return fmt.Sprintf("parameter value[%#v]", idx_cd4240) }); err != nil {
				return err
			}
		}
	case []*CdnFrontdoorBatchRuleSetRuleConditionsClientPort:
		value_ := value.([]*CdnFrontdoorBatchRuleSetRuleConditionsClientPort)
		value := &value_
		for idx_cd4240, v := range *value {
			if err := _jsii_.ValidateStruct(v, func() string { return fmt.Sprintf("parameter value[%#v]", idx_cd4240) }); err != nil {
				return err
			}
		}
	default:
		if !_jsii_.IsAnonymousProxy(value) {
			return fmt.Errorf("parameter value must be one of the allowed types: cdktn.IResolvable, *[]*CdnFrontdoorBatchRuleSetRuleConditionsClientPort; received %#v (a %T)", value, value)
		}
	}

	return nil
}

func (c *jsiiProxy_CdnFrontdoorBatchRuleSetRuleConditionsOutputReference) validatePutDeviceTypeParameters(value interface{}) error {
	if value == nil {
		return fmt.Errorf("parameter value is required, but nil was provided")
	}
	switch value.(type) {
	case cdktn.IResolvable:
		// ok
	case *[]*CdnFrontdoorBatchRuleSetRuleConditionsDeviceType:
		value := value.(*[]*CdnFrontdoorBatchRuleSetRuleConditionsDeviceType)
		for idx_cd4240, v := range *value {
			if err := _jsii_.ValidateStruct(v, func() string { return fmt.Sprintf("parameter value[%#v]", idx_cd4240) }); err != nil {
				return err
			}
		}
	case []*CdnFrontdoorBatchRuleSetRuleConditionsDeviceType:
		value_ := value.([]*CdnFrontdoorBatchRuleSetRuleConditionsDeviceType)
		value := &value_
		for idx_cd4240, v := range *value {
			if err := _jsii_.ValidateStruct(v, func() string { return fmt.Sprintf("parameter value[%#v]", idx_cd4240) }); err != nil {
				return err
			}
		}
	default:
		if !_jsii_.IsAnonymousProxy(value) {
			return fmt.Errorf("parameter value must be one of the allowed types: cdktn.IResolvable, *[]*CdnFrontdoorBatchRuleSetRuleConditionsDeviceType; received %#v (a %T)", value, value)
		}
	}

	return nil
}

func (c *jsiiProxy_CdnFrontdoorBatchRuleSetRuleConditionsOutputReference) validatePutHostNameParameters(value interface{}) error {
	if value == nil {
		return fmt.Errorf("parameter value is required, but nil was provided")
	}
	switch value.(type) {
	case cdktn.IResolvable:
		// ok
	case *[]*CdnFrontdoorBatchRuleSetRuleConditionsHostName:
		value := value.(*[]*CdnFrontdoorBatchRuleSetRuleConditionsHostName)
		for idx_cd4240, v := range *value {
			if err := _jsii_.ValidateStruct(v, func() string { return fmt.Sprintf("parameter value[%#v]", idx_cd4240) }); err != nil {
				return err
			}
		}
	case []*CdnFrontdoorBatchRuleSetRuleConditionsHostName:
		value_ := value.([]*CdnFrontdoorBatchRuleSetRuleConditionsHostName)
		value := &value_
		for idx_cd4240, v := range *value {
			if err := _jsii_.ValidateStruct(v, func() string { return fmt.Sprintf("parameter value[%#v]", idx_cd4240) }); err != nil {
				return err
			}
		}
	default:
		if !_jsii_.IsAnonymousProxy(value) {
			return fmt.Errorf("parameter value must be one of the allowed types: cdktn.IResolvable, *[]*CdnFrontdoorBatchRuleSetRuleConditionsHostName; received %#v (a %T)", value, value)
		}
	}

	return nil
}

func (c *jsiiProxy_CdnFrontdoorBatchRuleSetRuleConditionsOutputReference) validatePutHttpVersionParameters(value interface{}) error {
	if value == nil {
		return fmt.Errorf("parameter value is required, but nil was provided")
	}
	switch value.(type) {
	case cdktn.IResolvable:
		// ok
	case *[]*CdnFrontdoorBatchRuleSetRuleConditionsHttpVersion:
		value := value.(*[]*CdnFrontdoorBatchRuleSetRuleConditionsHttpVersion)
		for idx_cd4240, v := range *value {
			if err := _jsii_.ValidateStruct(v, func() string { return fmt.Sprintf("parameter value[%#v]", idx_cd4240) }); err != nil {
				return err
			}
		}
	case []*CdnFrontdoorBatchRuleSetRuleConditionsHttpVersion:
		value_ := value.([]*CdnFrontdoorBatchRuleSetRuleConditionsHttpVersion)
		value := &value_
		for idx_cd4240, v := range *value {
			if err := _jsii_.ValidateStruct(v, func() string { return fmt.Sprintf("parameter value[%#v]", idx_cd4240) }); err != nil {
				return err
			}
		}
	default:
		if !_jsii_.IsAnonymousProxy(value) {
			return fmt.Errorf("parameter value must be one of the allowed types: cdktn.IResolvable, *[]*CdnFrontdoorBatchRuleSetRuleConditionsHttpVersion; received %#v (a %T)", value, value)
		}
	}

	return nil
}

func (c *jsiiProxy_CdnFrontdoorBatchRuleSetRuleConditionsOutputReference) validatePutPostArgumentParameters(value interface{}) error {
	if value == nil {
		return fmt.Errorf("parameter value is required, but nil was provided")
	}
	switch value.(type) {
	case cdktn.IResolvable:
		// ok
	case *[]*CdnFrontdoorBatchRuleSetRuleConditionsPostArgument:
		value := value.(*[]*CdnFrontdoorBatchRuleSetRuleConditionsPostArgument)
		for idx_cd4240, v := range *value {
			if err := _jsii_.ValidateStruct(v, func() string { return fmt.Sprintf("parameter value[%#v]", idx_cd4240) }); err != nil {
				return err
			}
		}
	case []*CdnFrontdoorBatchRuleSetRuleConditionsPostArgument:
		value_ := value.([]*CdnFrontdoorBatchRuleSetRuleConditionsPostArgument)
		value := &value_
		for idx_cd4240, v := range *value {
			if err := _jsii_.ValidateStruct(v, func() string { return fmt.Sprintf("parameter value[%#v]", idx_cd4240) }); err != nil {
				return err
			}
		}
	default:
		if !_jsii_.IsAnonymousProxy(value) {
			return fmt.Errorf("parameter value must be one of the allowed types: cdktn.IResolvable, *[]*CdnFrontdoorBatchRuleSetRuleConditionsPostArgument; received %#v (a %T)", value, value)
		}
	}

	return nil
}

func (c *jsiiProxy_CdnFrontdoorBatchRuleSetRuleConditionsOutputReference) validatePutQueryStringParameters(value interface{}) error {
	if value == nil {
		return fmt.Errorf("parameter value is required, but nil was provided")
	}
	switch value.(type) {
	case cdktn.IResolvable:
		// ok
	case *[]*CdnFrontdoorBatchRuleSetRuleConditionsQueryString:
		value := value.(*[]*CdnFrontdoorBatchRuleSetRuleConditionsQueryString)
		for idx_cd4240, v := range *value {
			if err := _jsii_.ValidateStruct(v, func() string { return fmt.Sprintf("parameter value[%#v]", idx_cd4240) }); err != nil {
				return err
			}
		}
	case []*CdnFrontdoorBatchRuleSetRuleConditionsQueryString:
		value_ := value.([]*CdnFrontdoorBatchRuleSetRuleConditionsQueryString)
		value := &value_
		for idx_cd4240, v := range *value {
			if err := _jsii_.ValidateStruct(v, func() string { return fmt.Sprintf("parameter value[%#v]", idx_cd4240) }); err != nil {
				return err
			}
		}
	default:
		if !_jsii_.IsAnonymousProxy(value) {
			return fmt.Errorf("parameter value must be one of the allowed types: cdktn.IResolvable, *[]*CdnFrontdoorBatchRuleSetRuleConditionsQueryString; received %#v (a %T)", value, value)
		}
	}

	return nil
}

func (c *jsiiProxy_CdnFrontdoorBatchRuleSetRuleConditionsOutputReference) validatePutRemoteAddressParameters(value interface{}) error {
	if value == nil {
		return fmt.Errorf("parameter value is required, but nil was provided")
	}
	switch value.(type) {
	case cdktn.IResolvable:
		// ok
	case *[]*CdnFrontdoorBatchRuleSetRuleConditionsRemoteAddress:
		value := value.(*[]*CdnFrontdoorBatchRuleSetRuleConditionsRemoteAddress)
		for idx_cd4240, v := range *value {
			if err := _jsii_.ValidateStruct(v, func() string { return fmt.Sprintf("parameter value[%#v]", idx_cd4240) }); err != nil {
				return err
			}
		}
	case []*CdnFrontdoorBatchRuleSetRuleConditionsRemoteAddress:
		value_ := value.([]*CdnFrontdoorBatchRuleSetRuleConditionsRemoteAddress)
		value := &value_
		for idx_cd4240, v := range *value {
			if err := _jsii_.ValidateStruct(v, func() string { return fmt.Sprintf("parameter value[%#v]", idx_cd4240) }); err != nil {
				return err
			}
		}
	default:
		if !_jsii_.IsAnonymousProxy(value) {
			return fmt.Errorf("parameter value must be one of the allowed types: cdktn.IResolvable, *[]*CdnFrontdoorBatchRuleSetRuleConditionsRemoteAddress; received %#v (a %T)", value, value)
		}
	}

	return nil
}

func (c *jsiiProxy_CdnFrontdoorBatchRuleSetRuleConditionsOutputReference) validatePutRequestBodyParameters(value interface{}) error {
	if value == nil {
		return fmt.Errorf("parameter value is required, but nil was provided")
	}
	switch value.(type) {
	case cdktn.IResolvable:
		// ok
	case *[]*CdnFrontdoorBatchRuleSetRuleConditionsRequestBody:
		value := value.(*[]*CdnFrontdoorBatchRuleSetRuleConditionsRequestBody)
		for idx_cd4240, v := range *value {
			if err := _jsii_.ValidateStruct(v, func() string { return fmt.Sprintf("parameter value[%#v]", idx_cd4240) }); err != nil {
				return err
			}
		}
	case []*CdnFrontdoorBatchRuleSetRuleConditionsRequestBody:
		value_ := value.([]*CdnFrontdoorBatchRuleSetRuleConditionsRequestBody)
		value := &value_
		for idx_cd4240, v := range *value {
			if err := _jsii_.ValidateStruct(v, func() string { return fmt.Sprintf("parameter value[%#v]", idx_cd4240) }); err != nil {
				return err
			}
		}
	default:
		if !_jsii_.IsAnonymousProxy(value) {
			return fmt.Errorf("parameter value must be one of the allowed types: cdktn.IResolvable, *[]*CdnFrontdoorBatchRuleSetRuleConditionsRequestBody; received %#v (a %T)", value, value)
		}
	}

	return nil
}

func (c *jsiiProxy_CdnFrontdoorBatchRuleSetRuleConditionsOutputReference) validatePutRequestCookiesParameters(value interface{}) error {
	if value == nil {
		return fmt.Errorf("parameter value is required, but nil was provided")
	}
	switch value.(type) {
	case cdktn.IResolvable:
		// ok
	case *[]*CdnFrontdoorBatchRuleSetRuleConditionsRequestCookies:
		value := value.(*[]*CdnFrontdoorBatchRuleSetRuleConditionsRequestCookies)
		for idx_cd4240, v := range *value {
			if err := _jsii_.ValidateStruct(v, func() string { return fmt.Sprintf("parameter value[%#v]", idx_cd4240) }); err != nil {
				return err
			}
		}
	case []*CdnFrontdoorBatchRuleSetRuleConditionsRequestCookies:
		value_ := value.([]*CdnFrontdoorBatchRuleSetRuleConditionsRequestCookies)
		value := &value_
		for idx_cd4240, v := range *value {
			if err := _jsii_.ValidateStruct(v, func() string { return fmt.Sprintf("parameter value[%#v]", idx_cd4240) }); err != nil {
				return err
			}
		}
	default:
		if !_jsii_.IsAnonymousProxy(value) {
			return fmt.Errorf("parameter value must be one of the allowed types: cdktn.IResolvable, *[]*CdnFrontdoorBatchRuleSetRuleConditionsRequestCookies; received %#v (a %T)", value, value)
		}
	}

	return nil
}

func (c *jsiiProxy_CdnFrontdoorBatchRuleSetRuleConditionsOutputReference) validatePutRequestFileExtensionParameters(value interface{}) error {
	if value == nil {
		return fmt.Errorf("parameter value is required, but nil was provided")
	}
	switch value.(type) {
	case cdktn.IResolvable:
		// ok
	case *[]*CdnFrontdoorBatchRuleSetRuleConditionsRequestFileExtension:
		value := value.(*[]*CdnFrontdoorBatchRuleSetRuleConditionsRequestFileExtension)
		for idx_cd4240, v := range *value {
			if err := _jsii_.ValidateStruct(v, func() string { return fmt.Sprintf("parameter value[%#v]", idx_cd4240) }); err != nil {
				return err
			}
		}
	case []*CdnFrontdoorBatchRuleSetRuleConditionsRequestFileExtension:
		value_ := value.([]*CdnFrontdoorBatchRuleSetRuleConditionsRequestFileExtension)
		value := &value_
		for idx_cd4240, v := range *value {
			if err := _jsii_.ValidateStruct(v, func() string { return fmt.Sprintf("parameter value[%#v]", idx_cd4240) }); err != nil {
				return err
			}
		}
	default:
		if !_jsii_.IsAnonymousProxy(value) {
			return fmt.Errorf("parameter value must be one of the allowed types: cdktn.IResolvable, *[]*CdnFrontdoorBatchRuleSetRuleConditionsRequestFileExtension; received %#v (a %T)", value, value)
		}
	}

	return nil
}

func (c *jsiiProxy_CdnFrontdoorBatchRuleSetRuleConditionsOutputReference) validatePutRequestFilenameParameters(value interface{}) error {
	if value == nil {
		return fmt.Errorf("parameter value is required, but nil was provided")
	}
	switch value.(type) {
	case cdktn.IResolvable:
		// ok
	case *[]*CdnFrontdoorBatchRuleSetRuleConditionsRequestFilename:
		value := value.(*[]*CdnFrontdoorBatchRuleSetRuleConditionsRequestFilename)
		for idx_cd4240, v := range *value {
			if err := _jsii_.ValidateStruct(v, func() string { return fmt.Sprintf("parameter value[%#v]", idx_cd4240) }); err != nil {
				return err
			}
		}
	case []*CdnFrontdoorBatchRuleSetRuleConditionsRequestFilename:
		value_ := value.([]*CdnFrontdoorBatchRuleSetRuleConditionsRequestFilename)
		value := &value_
		for idx_cd4240, v := range *value {
			if err := _jsii_.ValidateStruct(v, func() string { return fmt.Sprintf("parameter value[%#v]", idx_cd4240) }); err != nil {
				return err
			}
		}
	default:
		if !_jsii_.IsAnonymousProxy(value) {
			return fmt.Errorf("parameter value must be one of the allowed types: cdktn.IResolvable, *[]*CdnFrontdoorBatchRuleSetRuleConditionsRequestFilename; received %#v (a %T)", value, value)
		}
	}

	return nil
}

func (c *jsiiProxy_CdnFrontdoorBatchRuleSetRuleConditionsOutputReference) validatePutRequestHeaderParameters(value interface{}) error {
	if value == nil {
		return fmt.Errorf("parameter value is required, but nil was provided")
	}
	switch value.(type) {
	case cdktn.IResolvable:
		// ok
	case *[]*CdnFrontdoorBatchRuleSetRuleConditionsRequestHeader:
		value := value.(*[]*CdnFrontdoorBatchRuleSetRuleConditionsRequestHeader)
		for idx_cd4240, v := range *value {
			if err := _jsii_.ValidateStruct(v, func() string { return fmt.Sprintf("parameter value[%#v]", idx_cd4240) }); err != nil {
				return err
			}
		}
	case []*CdnFrontdoorBatchRuleSetRuleConditionsRequestHeader:
		value_ := value.([]*CdnFrontdoorBatchRuleSetRuleConditionsRequestHeader)
		value := &value_
		for idx_cd4240, v := range *value {
			if err := _jsii_.ValidateStruct(v, func() string { return fmt.Sprintf("parameter value[%#v]", idx_cd4240) }); err != nil {
				return err
			}
		}
	default:
		if !_jsii_.IsAnonymousProxy(value) {
			return fmt.Errorf("parameter value must be one of the allowed types: cdktn.IResolvable, *[]*CdnFrontdoorBatchRuleSetRuleConditionsRequestHeader; received %#v (a %T)", value, value)
		}
	}

	return nil
}

func (c *jsiiProxy_CdnFrontdoorBatchRuleSetRuleConditionsOutputReference) validatePutRequestMethodParameters(value interface{}) error {
	if value == nil {
		return fmt.Errorf("parameter value is required, but nil was provided")
	}
	switch value.(type) {
	case cdktn.IResolvable:
		// ok
	case *[]*CdnFrontdoorBatchRuleSetRuleConditionsRequestMethod:
		value := value.(*[]*CdnFrontdoorBatchRuleSetRuleConditionsRequestMethod)
		for idx_cd4240, v := range *value {
			if err := _jsii_.ValidateStruct(v, func() string { return fmt.Sprintf("parameter value[%#v]", idx_cd4240) }); err != nil {
				return err
			}
		}
	case []*CdnFrontdoorBatchRuleSetRuleConditionsRequestMethod:
		value_ := value.([]*CdnFrontdoorBatchRuleSetRuleConditionsRequestMethod)
		value := &value_
		for idx_cd4240, v := range *value {
			if err := _jsii_.ValidateStruct(v, func() string { return fmt.Sprintf("parameter value[%#v]", idx_cd4240) }); err != nil {
				return err
			}
		}
	default:
		if !_jsii_.IsAnonymousProxy(value) {
			return fmt.Errorf("parameter value must be one of the allowed types: cdktn.IResolvable, *[]*CdnFrontdoorBatchRuleSetRuleConditionsRequestMethod; received %#v (a %T)", value, value)
		}
	}

	return nil
}

func (c *jsiiProxy_CdnFrontdoorBatchRuleSetRuleConditionsOutputReference) validatePutRequestPathParameters(value interface{}) error {
	if value == nil {
		return fmt.Errorf("parameter value is required, but nil was provided")
	}
	switch value.(type) {
	case cdktn.IResolvable:
		// ok
	case *[]*CdnFrontdoorBatchRuleSetRuleConditionsRequestPath:
		value := value.(*[]*CdnFrontdoorBatchRuleSetRuleConditionsRequestPath)
		for idx_cd4240, v := range *value {
			if err := _jsii_.ValidateStruct(v, func() string { return fmt.Sprintf("parameter value[%#v]", idx_cd4240) }); err != nil {
				return err
			}
		}
	case []*CdnFrontdoorBatchRuleSetRuleConditionsRequestPath:
		value_ := value.([]*CdnFrontdoorBatchRuleSetRuleConditionsRequestPath)
		value := &value_
		for idx_cd4240, v := range *value {
			if err := _jsii_.ValidateStruct(v, func() string { return fmt.Sprintf("parameter value[%#v]", idx_cd4240) }); err != nil {
				return err
			}
		}
	default:
		if !_jsii_.IsAnonymousProxy(value) {
			return fmt.Errorf("parameter value must be one of the allowed types: cdktn.IResolvable, *[]*CdnFrontdoorBatchRuleSetRuleConditionsRequestPath; received %#v (a %T)", value, value)
		}
	}

	return nil
}

func (c *jsiiProxy_CdnFrontdoorBatchRuleSetRuleConditionsOutputReference) validatePutRequestSchemeParameters(value interface{}) error {
	if value == nil {
		return fmt.Errorf("parameter value is required, but nil was provided")
	}
	switch value.(type) {
	case cdktn.IResolvable:
		// ok
	case *[]*CdnFrontdoorBatchRuleSetRuleConditionsRequestScheme:
		value := value.(*[]*CdnFrontdoorBatchRuleSetRuleConditionsRequestScheme)
		for idx_cd4240, v := range *value {
			if err := _jsii_.ValidateStruct(v, func() string { return fmt.Sprintf("parameter value[%#v]", idx_cd4240) }); err != nil {
				return err
			}
		}
	case []*CdnFrontdoorBatchRuleSetRuleConditionsRequestScheme:
		value_ := value.([]*CdnFrontdoorBatchRuleSetRuleConditionsRequestScheme)
		value := &value_
		for idx_cd4240, v := range *value {
			if err := _jsii_.ValidateStruct(v, func() string { return fmt.Sprintf("parameter value[%#v]", idx_cd4240) }); err != nil {
				return err
			}
		}
	default:
		if !_jsii_.IsAnonymousProxy(value) {
			return fmt.Errorf("parameter value must be one of the allowed types: cdktn.IResolvable, *[]*CdnFrontdoorBatchRuleSetRuleConditionsRequestScheme; received %#v (a %T)", value, value)
		}
	}

	return nil
}

func (c *jsiiProxy_CdnFrontdoorBatchRuleSetRuleConditionsOutputReference) validatePutRequestUrlParameters(value interface{}) error {
	if value == nil {
		return fmt.Errorf("parameter value is required, but nil was provided")
	}
	switch value.(type) {
	case cdktn.IResolvable:
		// ok
	case *[]*CdnFrontdoorBatchRuleSetRuleConditionsRequestUrl:
		value := value.(*[]*CdnFrontdoorBatchRuleSetRuleConditionsRequestUrl)
		for idx_cd4240, v := range *value {
			if err := _jsii_.ValidateStruct(v, func() string { return fmt.Sprintf("parameter value[%#v]", idx_cd4240) }); err != nil {
				return err
			}
		}
	case []*CdnFrontdoorBatchRuleSetRuleConditionsRequestUrl:
		value_ := value.([]*CdnFrontdoorBatchRuleSetRuleConditionsRequestUrl)
		value := &value_
		for idx_cd4240, v := range *value {
			if err := _jsii_.ValidateStruct(v, func() string { return fmt.Sprintf("parameter value[%#v]", idx_cd4240) }); err != nil {
				return err
			}
		}
	default:
		if !_jsii_.IsAnonymousProxy(value) {
			return fmt.Errorf("parameter value must be one of the allowed types: cdktn.IResolvable, *[]*CdnFrontdoorBatchRuleSetRuleConditionsRequestUrl; received %#v (a %T)", value, value)
		}
	}

	return nil
}

func (c *jsiiProxy_CdnFrontdoorBatchRuleSetRuleConditionsOutputReference) validatePutServerPortParameters(value interface{}) error {
	if value == nil {
		return fmt.Errorf("parameter value is required, but nil was provided")
	}
	switch value.(type) {
	case cdktn.IResolvable:
		// ok
	case *[]*CdnFrontdoorBatchRuleSetRuleConditionsServerPort:
		value := value.(*[]*CdnFrontdoorBatchRuleSetRuleConditionsServerPort)
		for idx_cd4240, v := range *value {
			if err := _jsii_.ValidateStruct(v, func() string { return fmt.Sprintf("parameter value[%#v]", idx_cd4240) }); err != nil {
				return err
			}
		}
	case []*CdnFrontdoorBatchRuleSetRuleConditionsServerPort:
		value_ := value.([]*CdnFrontdoorBatchRuleSetRuleConditionsServerPort)
		value := &value_
		for idx_cd4240, v := range *value {
			if err := _jsii_.ValidateStruct(v, func() string { return fmt.Sprintf("parameter value[%#v]", idx_cd4240) }); err != nil {
				return err
			}
		}
	default:
		if !_jsii_.IsAnonymousProxy(value) {
			return fmt.Errorf("parameter value must be one of the allowed types: cdktn.IResolvable, *[]*CdnFrontdoorBatchRuleSetRuleConditionsServerPort; received %#v (a %T)", value, value)
		}
	}

	return nil
}

func (c *jsiiProxy_CdnFrontdoorBatchRuleSetRuleConditionsOutputReference) validatePutSocketAddressParameters(value interface{}) error {
	if value == nil {
		return fmt.Errorf("parameter value is required, but nil was provided")
	}
	switch value.(type) {
	case cdktn.IResolvable:
		// ok
	case *[]*CdnFrontdoorBatchRuleSetRuleConditionsSocketAddress:
		value := value.(*[]*CdnFrontdoorBatchRuleSetRuleConditionsSocketAddress)
		for idx_cd4240, v := range *value {
			if err := _jsii_.ValidateStruct(v, func() string { return fmt.Sprintf("parameter value[%#v]", idx_cd4240) }); err != nil {
				return err
			}
		}
	case []*CdnFrontdoorBatchRuleSetRuleConditionsSocketAddress:
		value_ := value.([]*CdnFrontdoorBatchRuleSetRuleConditionsSocketAddress)
		value := &value_
		for idx_cd4240, v := range *value {
			if err := _jsii_.ValidateStruct(v, func() string { return fmt.Sprintf("parameter value[%#v]", idx_cd4240) }); err != nil {
				return err
			}
		}
	default:
		if !_jsii_.IsAnonymousProxy(value) {
			return fmt.Errorf("parameter value must be one of the allowed types: cdktn.IResolvable, *[]*CdnFrontdoorBatchRuleSetRuleConditionsSocketAddress; received %#v (a %T)", value, value)
		}
	}

	return nil
}

func (c *jsiiProxy_CdnFrontdoorBatchRuleSetRuleConditionsOutputReference) validatePutSslProtocolParameters(value interface{}) error {
	if value == nil {
		return fmt.Errorf("parameter value is required, but nil was provided")
	}
	switch value.(type) {
	case cdktn.IResolvable:
		// ok
	case *[]*CdnFrontdoorBatchRuleSetRuleConditionsSslProtocol:
		value := value.(*[]*CdnFrontdoorBatchRuleSetRuleConditionsSslProtocol)
		for idx_cd4240, v := range *value {
			if err := _jsii_.ValidateStruct(v, func() string { return fmt.Sprintf("parameter value[%#v]", idx_cd4240) }); err != nil {
				return err
			}
		}
	case []*CdnFrontdoorBatchRuleSetRuleConditionsSslProtocol:
		value_ := value.([]*CdnFrontdoorBatchRuleSetRuleConditionsSslProtocol)
		value := &value_
		for idx_cd4240, v := range *value {
			if err := _jsii_.ValidateStruct(v, func() string { return fmt.Sprintf("parameter value[%#v]", idx_cd4240) }); err != nil {
				return err
			}
		}
	default:
		if !_jsii_.IsAnonymousProxy(value) {
			return fmt.Errorf("parameter value must be one of the allowed types: cdktn.IResolvable, *[]*CdnFrontdoorBatchRuleSetRuleConditionsSslProtocol; received %#v (a %T)", value, value)
		}
	}

	return nil
}

func (c *jsiiProxy_CdnFrontdoorBatchRuleSetRuleConditionsOutputReference) validateResolveParameters(context cdktn.IResolveContext) error {
	if context == nil {
		return fmt.Errorf("parameter context is required, but nil was provided")
	}

	return nil
}

func (j *jsiiProxy_CdnFrontdoorBatchRuleSetRuleConditionsOutputReference) validateSetComplexObjectIndexParameters(val interface{}) error {
	switch val.(type) {
	case *string:
		// ok
	case string:
		// ok
	case *float64:
		// ok
	case float64:
		// ok
	case *int:
		// ok
	case int:
		// ok
	case *uint:
		// ok
	case uint:
		// ok
	case *int8:
		// ok
	case int8:
		// ok
	case *int16:
		// ok
	case int16:
		// ok
	case *int32:
		// ok
	case int32:
		// ok
	case *int64:
		// ok
	case int64:
		// ok
	case *uint8:
		// ok
	case uint8:
		// ok
	case *uint16:
		// ok
	case uint16:
		// ok
	case *uint32:
		// ok
	case uint32:
		// ok
	case *uint64:
		// ok
	case uint64:
		// ok
	default:
		return fmt.Errorf("parameter val must be one of the allowed types: *string, *float64; received %#v (a %T)", val, val)
	}

	return nil
}

func (j *jsiiProxy_CdnFrontdoorBatchRuleSetRuleConditionsOutputReference) validateSetComplexObjectIsFromSetParameters(val *bool) error {
	if val == nil {
		return fmt.Errorf("parameter val is required, but nil was provided")
	}

	return nil
}

func (j *jsiiProxy_CdnFrontdoorBatchRuleSetRuleConditionsOutputReference) validateSetInternalValueParameters(val *CdnFrontdoorBatchRuleSetRuleConditions) error {
	if err := _jsii_.ValidateStruct(val, func() string { return "parameter val" }); err != nil {
		return err
	}

	return nil
}

func (j *jsiiProxy_CdnFrontdoorBatchRuleSetRuleConditionsOutputReference) validateSetTerraformAttributeParameters(val *string) error {
	if val == nil {
		return fmt.Errorf("parameter val is required, but nil was provided")
	}

	return nil
}

func (j *jsiiProxy_CdnFrontdoorBatchRuleSetRuleConditionsOutputReference) validateSetTerraformResourceParameters(val cdktn.IInterpolatingParent) error {
	if val == nil {
		return fmt.Errorf("parameter val is required, but nil was provided")
	}

	return nil
}

func validateNewCdnFrontdoorBatchRuleSetRuleConditionsOutputReferenceParameters(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) error {
	if terraformResource == nil {
		return fmt.Errorf("parameter terraformResource is required, but nil was provided")
	}

	if terraformAttribute == nil {
		return fmt.Errorf("parameter terraformAttribute is required, but nil was provided")
	}

	return nil
}

