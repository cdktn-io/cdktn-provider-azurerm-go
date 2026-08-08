// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package recoveryservicesvault

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-provider-azurerm-go/azurerm/v18/jsii"

	"github.com/cdktn-io/cdktn-provider-azurerm-go/azurerm/v18/recoveryservicesvault/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type RecoveryServicesVaultMonitoringOutputReference interface {
	cdktn.ComplexObject
	AlertsForAllFailoverIssuesEnabled() interface{}
	SetAlertsForAllFailoverIssuesEnabled(val interface{})
	AlertsForAllFailoverIssuesEnabledInput() interface{}
	AlertsForAllJobFailuresEnabled() interface{}
	SetAlertsForAllJobFailuresEnabled(val interface{})
	AlertsForAllJobFailuresEnabledInput() interface{}
	AlertsForAllReplicationIssuesEnabled() interface{}
	SetAlertsForAllReplicationIssuesEnabled(val interface{})
	AlertsForAllReplicationIssuesEnabledInput() interface{}
	AlertsForCriticalOperationFailuresEnabled() interface{}
	SetAlertsForCriticalOperationFailuresEnabled(val interface{})
	AlertsForCriticalOperationFailuresEnabledInput() interface{}
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
	EmailNotificationsForSiteRecoveryEnabled() interface{}
	SetEmailNotificationsForSiteRecoveryEnabled(val interface{})
	EmailNotificationsForSiteRecoveryEnabledInput() interface{}
	// Experimental.
	Fqn() *string
	InternalValue() *RecoveryServicesVaultMonitoring
	SetInternalValue(val *RecoveryServicesVaultMonitoring)
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
	ResetAlertsForAllFailoverIssuesEnabled()
	ResetAlertsForAllJobFailuresEnabled()
	ResetAlertsForAllReplicationIssuesEnabled()
	ResetAlertsForCriticalOperationFailuresEnabled()
	ResetEmailNotificationsForSiteRecoveryEnabled()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktn.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for RecoveryServicesVaultMonitoringOutputReference
type jsiiProxy_RecoveryServicesVaultMonitoringOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_RecoveryServicesVaultMonitoringOutputReference) AlertsForAllFailoverIssuesEnabled() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"alertsForAllFailoverIssuesEnabled",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RecoveryServicesVaultMonitoringOutputReference) AlertsForAllFailoverIssuesEnabledInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"alertsForAllFailoverIssuesEnabledInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RecoveryServicesVaultMonitoringOutputReference) AlertsForAllJobFailuresEnabled() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"alertsForAllJobFailuresEnabled",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RecoveryServicesVaultMonitoringOutputReference) AlertsForAllJobFailuresEnabledInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"alertsForAllJobFailuresEnabledInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RecoveryServicesVaultMonitoringOutputReference) AlertsForAllReplicationIssuesEnabled() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"alertsForAllReplicationIssuesEnabled",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RecoveryServicesVaultMonitoringOutputReference) AlertsForAllReplicationIssuesEnabledInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"alertsForAllReplicationIssuesEnabledInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RecoveryServicesVaultMonitoringOutputReference) AlertsForCriticalOperationFailuresEnabled() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"alertsForCriticalOperationFailuresEnabled",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RecoveryServicesVaultMonitoringOutputReference) AlertsForCriticalOperationFailuresEnabledInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"alertsForCriticalOperationFailuresEnabledInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RecoveryServicesVaultMonitoringOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RecoveryServicesVaultMonitoringOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RecoveryServicesVaultMonitoringOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RecoveryServicesVaultMonitoringOutputReference) EmailNotificationsForSiteRecoveryEnabled() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"emailNotificationsForSiteRecoveryEnabled",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RecoveryServicesVaultMonitoringOutputReference) EmailNotificationsForSiteRecoveryEnabledInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"emailNotificationsForSiteRecoveryEnabledInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RecoveryServicesVaultMonitoringOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RecoveryServicesVaultMonitoringOutputReference) InternalValue() *RecoveryServicesVaultMonitoring {
	var returns *RecoveryServicesVaultMonitoring
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RecoveryServicesVaultMonitoringOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RecoveryServicesVaultMonitoringOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewRecoveryServicesVaultMonitoringOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) RecoveryServicesVaultMonitoringOutputReference {
	_init_.Initialize()

	if err := validateNewRecoveryServicesVaultMonitoringOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_RecoveryServicesVaultMonitoringOutputReference{}

	_jsii_.Create(
		"@cdktn/provider-azurerm.recoveryServicesVault.RecoveryServicesVaultMonitoringOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewRecoveryServicesVaultMonitoringOutputReference_Override(r RecoveryServicesVaultMonitoringOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/provider-azurerm.recoveryServicesVault.RecoveryServicesVaultMonitoringOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		r,
	)
}

func (j *jsiiProxy_RecoveryServicesVaultMonitoringOutputReference)SetAlertsForAllFailoverIssuesEnabled(val interface{}) {
	if err := j.validateSetAlertsForAllFailoverIssuesEnabledParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"alertsForAllFailoverIssuesEnabled",
		val,
	)
}

func (j *jsiiProxy_RecoveryServicesVaultMonitoringOutputReference)SetAlertsForAllJobFailuresEnabled(val interface{}) {
	if err := j.validateSetAlertsForAllJobFailuresEnabledParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"alertsForAllJobFailuresEnabled",
		val,
	)
}

func (j *jsiiProxy_RecoveryServicesVaultMonitoringOutputReference)SetAlertsForAllReplicationIssuesEnabled(val interface{}) {
	if err := j.validateSetAlertsForAllReplicationIssuesEnabledParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"alertsForAllReplicationIssuesEnabled",
		val,
	)
}

func (j *jsiiProxy_RecoveryServicesVaultMonitoringOutputReference)SetAlertsForCriticalOperationFailuresEnabled(val interface{}) {
	if err := j.validateSetAlertsForCriticalOperationFailuresEnabledParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"alertsForCriticalOperationFailuresEnabled",
		val,
	)
}

func (j *jsiiProxy_RecoveryServicesVaultMonitoringOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_RecoveryServicesVaultMonitoringOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_RecoveryServicesVaultMonitoringOutputReference)SetEmailNotificationsForSiteRecoveryEnabled(val interface{}) {
	if err := j.validateSetEmailNotificationsForSiteRecoveryEnabledParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"emailNotificationsForSiteRecoveryEnabled",
		val,
	)
}

func (j *jsiiProxy_RecoveryServicesVaultMonitoringOutputReference)SetInternalValue(val *RecoveryServicesVaultMonitoring) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_RecoveryServicesVaultMonitoringOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_RecoveryServicesVaultMonitoringOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (r *jsiiProxy_RecoveryServicesVaultMonitoringOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		r,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (r *jsiiProxy_RecoveryServicesVaultMonitoringOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (r *jsiiProxy_RecoveryServicesVaultMonitoringOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (r *jsiiProxy_RecoveryServicesVaultMonitoringOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (r *jsiiProxy_RecoveryServicesVaultMonitoringOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (r *jsiiProxy_RecoveryServicesVaultMonitoringOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (r *jsiiProxy_RecoveryServicesVaultMonitoringOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (r *jsiiProxy_RecoveryServicesVaultMonitoringOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (r *jsiiProxy_RecoveryServicesVaultMonitoringOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (r *jsiiProxy_RecoveryServicesVaultMonitoringOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (r *jsiiProxy_RecoveryServicesVaultMonitoringOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		r,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (r *jsiiProxy_RecoveryServicesVaultMonitoringOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (r *jsiiProxy_RecoveryServicesVaultMonitoringOutputReference) ResetAlertsForAllFailoverIssuesEnabled() {
	_jsii_.InvokeVoid(
		r,
		"resetAlertsForAllFailoverIssuesEnabled",
		nil, // no parameters
	)
}

func (r *jsiiProxy_RecoveryServicesVaultMonitoringOutputReference) ResetAlertsForAllJobFailuresEnabled() {
	_jsii_.InvokeVoid(
		r,
		"resetAlertsForAllJobFailuresEnabled",
		nil, // no parameters
	)
}

func (r *jsiiProxy_RecoveryServicesVaultMonitoringOutputReference) ResetAlertsForAllReplicationIssuesEnabled() {
	_jsii_.InvokeVoid(
		r,
		"resetAlertsForAllReplicationIssuesEnabled",
		nil, // no parameters
	)
}

func (r *jsiiProxy_RecoveryServicesVaultMonitoringOutputReference) ResetAlertsForCriticalOperationFailuresEnabled() {
	_jsii_.InvokeVoid(
		r,
		"resetAlertsForCriticalOperationFailuresEnabled",
		nil, // no parameters
	)
}

func (r *jsiiProxy_RecoveryServicesVaultMonitoringOutputReference) ResetEmailNotificationsForSiteRecoveryEnabled() {
	_jsii_.InvokeVoid(
		r,
		"resetEmailNotificationsForSiteRecoveryEnabled",
		nil, // no parameters
	)
}

func (r *jsiiProxy_RecoveryServicesVaultMonitoringOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
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

func (r *jsiiProxy_RecoveryServicesVaultMonitoringOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		r,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

