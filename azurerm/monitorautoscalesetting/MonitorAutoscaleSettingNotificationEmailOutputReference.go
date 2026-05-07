// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package monitorautoscalesetting

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-provider-azurerm-go/azurerm/v16/jsii"

	"github.com/cdktn-io/cdktn-provider-azurerm-go/azurerm/v16/monitorautoscalesetting/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type MonitorAutoscaleSettingNotificationEmailOutputReference interface {
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
	CustomEmails() *[]*string
	SetCustomEmails(val *[]*string)
	CustomEmailsInput() *[]*string
	// Experimental.
	Fqn() *string
	InternalValue() *MonitorAutoscaleSettingNotificationEmail
	SetInternalValue(val *MonitorAutoscaleSettingNotificationEmail)
	SendToSubscriptionAdministrator() interface{}
	SetSendToSubscriptionAdministrator(val interface{})
	SendToSubscriptionAdministratorInput() interface{}
	SendToSubscriptionCoAdministrator() interface{}
	SetSendToSubscriptionCoAdministrator(val interface{})
	SendToSubscriptionCoAdministratorInput() interface{}
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
	ResetCustomEmails()
	ResetSendToSubscriptionAdministrator()
	ResetSendToSubscriptionCoAdministrator()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktn.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for MonitorAutoscaleSettingNotificationEmailOutputReference
type jsiiProxy_MonitorAutoscaleSettingNotificationEmailOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_MonitorAutoscaleSettingNotificationEmailOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MonitorAutoscaleSettingNotificationEmailOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MonitorAutoscaleSettingNotificationEmailOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MonitorAutoscaleSettingNotificationEmailOutputReference) CustomEmails() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"customEmails",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MonitorAutoscaleSettingNotificationEmailOutputReference) CustomEmailsInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"customEmailsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MonitorAutoscaleSettingNotificationEmailOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MonitorAutoscaleSettingNotificationEmailOutputReference) InternalValue() *MonitorAutoscaleSettingNotificationEmail {
	var returns *MonitorAutoscaleSettingNotificationEmail
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MonitorAutoscaleSettingNotificationEmailOutputReference) SendToSubscriptionAdministrator() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"sendToSubscriptionAdministrator",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MonitorAutoscaleSettingNotificationEmailOutputReference) SendToSubscriptionAdministratorInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"sendToSubscriptionAdministratorInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MonitorAutoscaleSettingNotificationEmailOutputReference) SendToSubscriptionCoAdministrator() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"sendToSubscriptionCoAdministrator",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MonitorAutoscaleSettingNotificationEmailOutputReference) SendToSubscriptionCoAdministratorInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"sendToSubscriptionCoAdministratorInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MonitorAutoscaleSettingNotificationEmailOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MonitorAutoscaleSettingNotificationEmailOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewMonitorAutoscaleSettingNotificationEmailOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) MonitorAutoscaleSettingNotificationEmailOutputReference {
	_init_.Initialize()

	if err := validateNewMonitorAutoscaleSettingNotificationEmailOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_MonitorAutoscaleSettingNotificationEmailOutputReference{}

	_jsii_.Create(
		"@cdktn/provider-azurerm.monitorAutoscaleSetting.MonitorAutoscaleSettingNotificationEmailOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewMonitorAutoscaleSettingNotificationEmailOutputReference_Override(m MonitorAutoscaleSettingNotificationEmailOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/provider-azurerm.monitorAutoscaleSetting.MonitorAutoscaleSettingNotificationEmailOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		m,
	)
}

func (j *jsiiProxy_MonitorAutoscaleSettingNotificationEmailOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_MonitorAutoscaleSettingNotificationEmailOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_MonitorAutoscaleSettingNotificationEmailOutputReference)SetCustomEmails(val *[]*string) {
	if err := j.validateSetCustomEmailsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"customEmails",
		val,
	)
}

func (j *jsiiProxy_MonitorAutoscaleSettingNotificationEmailOutputReference)SetInternalValue(val *MonitorAutoscaleSettingNotificationEmail) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_MonitorAutoscaleSettingNotificationEmailOutputReference)SetSendToSubscriptionAdministrator(val interface{}) {
	if err := j.validateSetSendToSubscriptionAdministratorParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"sendToSubscriptionAdministrator",
		val,
	)
}

func (j *jsiiProxy_MonitorAutoscaleSettingNotificationEmailOutputReference)SetSendToSubscriptionCoAdministrator(val interface{}) {
	if err := j.validateSetSendToSubscriptionCoAdministratorParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"sendToSubscriptionCoAdministrator",
		val,
	)
}

func (j *jsiiProxy_MonitorAutoscaleSettingNotificationEmailOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_MonitorAutoscaleSettingNotificationEmailOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (m *jsiiProxy_MonitorAutoscaleSettingNotificationEmailOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		m,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MonitorAutoscaleSettingNotificationEmailOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	if err := m.validateGetAnyMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]interface{}

	_jsii_.Invoke(
		m,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MonitorAutoscaleSettingNotificationEmailOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
	if err := m.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		m,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MonitorAutoscaleSettingNotificationEmailOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	if err := m.validateGetBooleanMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*bool

	_jsii_.Invoke(
		m,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MonitorAutoscaleSettingNotificationEmailOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	if err := m.validateGetListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*string

	_jsii_.Invoke(
		m,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MonitorAutoscaleSettingNotificationEmailOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	if err := m.validateGetNumberAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *float64

	_jsii_.Invoke(
		m,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MonitorAutoscaleSettingNotificationEmailOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	if err := m.validateGetNumberListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*float64

	_jsii_.Invoke(
		m,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MonitorAutoscaleSettingNotificationEmailOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	if err := m.validateGetNumberMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*float64

	_jsii_.Invoke(
		m,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MonitorAutoscaleSettingNotificationEmailOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	if err := m.validateGetStringAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		m,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MonitorAutoscaleSettingNotificationEmailOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	if err := m.validateGetStringMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*string

	_jsii_.Invoke(
		m,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MonitorAutoscaleSettingNotificationEmailOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		m,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MonitorAutoscaleSettingNotificationEmailOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
	if err := m.validateInterpolationForAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		m,
		"interpolationForAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MonitorAutoscaleSettingNotificationEmailOutputReference) ResetCustomEmails() {
	_jsii_.InvokeVoid(
		m,
		"resetCustomEmails",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MonitorAutoscaleSettingNotificationEmailOutputReference) ResetSendToSubscriptionAdministrator() {
	_jsii_.InvokeVoid(
		m,
		"resetSendToSubscriptionAdministrator",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MonitorAutoscaleSettingNotificationEmailOutputReference) ResetSendToSubscriptionCoAdministrator() {
	_jsii_.InvokeVoid(
		m,
		"resetSendToSubscriptionCoAdministrator",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MonitorAutoscaleSettingNotificationEmailOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
	if err := m.validateResolveParameters(context); err != nil {
		panic(err)
	}
	var returns interface{}

	_jsii_.Invoke(
		m,
		"resolve",
		[]interface{}{context},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MonitorAutoscaleSettingNotificationEmailOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		m,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

