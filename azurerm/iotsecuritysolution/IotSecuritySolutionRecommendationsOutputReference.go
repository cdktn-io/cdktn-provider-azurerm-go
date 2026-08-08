// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package iotsecuritysolution

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-provider-azurerm-go/azurerm/v18/jsii"

	"github.com/cdktn-io/cdktn-provider-azurerm-go/azurerm/v18/iotsecuritysolution/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type IotSecuritySolutionRecommendationsOutputReference interface {
	cdktn.ComplexObject
	AcrAuthentication() interface{}
	SetAcrAuthentication(val interface{})
	AcrAuthenticationInput() interface{}
	AgentSendUnutilizedMsg() interface{}
	SetAgentSendUnutilizedMsg(val interface{})
	AgentSendUnutilizedMsgInput() interface{}
	Baseline() interface{}
	SetBaseline(val interface{})
	BaselineInput() interface{}
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
	EdgeHubMemOptimize() interface{}
	SetEdgeHubMemOptimize(val interface{})
	EdgeHubMemOptimizeInput() interface{}
	EdgeLoggingOption() interface{}
	SetEdgeLoggingOption(val interface{})
	EdgeLoggingOptionInput() interface{}
	// Experimental.
	Fqn() *string
	InconsistentModuleSettings() interface{}
	SetInconsistentModuleSettings(val interface{})
	InconsistentModuleSettingsInput() interface{}
	InstallAgent() interface{}
	SetInstallAgent(val interface{})
	InstallAgentInput() interface{}
	InternalValue() *IotSecuritySolutionRecommendations
	SetInternalValue(val *IotSecuritySolutionRecommendations)
	IpFilterDenyAll() interface{}
	SetIpFilterDenyAll(val interface{})
	IpFilterDenyAllInput() interface{}
	IpFilterPermissiveRule() interface{}
	SetIpFilterPermissiveRule(val interface{})
	IpFilterPermissiveRuleInput() interface{}
	OpenPorts() interface{}
	SetOpenPorts(val interface{})
	OpenPortsInput() interface{}
	PermissiveFirewallPolicy() interface{}
	SetPermissiveFirewallPolicy(val interface{})
	PermissiveFirewallPolicyInput() interface{}
	PermissiveInputFirewallRules() interface{}
	SetPermissiveInputFirewallRules(val interface{})
	PermissiveInputFirewallRulesInput() interface{}
	PermissiveOutputFirewallRules() interface{}
	SetPermissiveOutputFirewallRules(val interface{})
	PermissiveOutputFirewallRulesInput() interface{}
	PrivilegedDockerOptions() interface{}
	SetPrivilegedDockerOptions(val interface{})
	PrivilegedDockerOptionsInput() interface{}
	SharedCredentials() interface{}
	SetSharedCredentials(val interface{})
	SharedCredentialsInput() interface{}
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktn.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktn.IInterpolatingParent)
	VulnerableTlsCipherSuite() interface{}
	SetVulnerableTlsCipherSuite(val interface{})
	VulnerableTlsCipherSuiteInput() interface{}
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
	ResetAcrAuthentication()
	ResetAgentSendUnutilizedMsg()
	ResetBaseline()
	ResetEdgeHubMemOptimize()
	ResetEdgeLoggingOption()
	ResetInconsistentModuleSettings()
	ResetInstallAgent()
	ResetIpFilterDenyAll()
	ResetIpFilterPermissiveRule()
	ResetOpenPorts()
	ResetPermissiveFirewallPolicy()
	ResetPermissiveInputFirewallRules()
	ResetPermissiveOutputFirewallRules()
	ResetPrivilegedDockerOptions()
	ResetSharedCredentials()
	ResetVulnerableTlsCipherSuite()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktn.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for IotSecuritySolutionRecommendationsOutputReference
type jsiiProxy_IotSecuritySolutionRecommendationsOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_IotSecuritySolutionRecommendationsOutputReference) AcrAuthentication() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"acrAuthentication",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IotSecuritySolutionRecommendationsOutputReference) AcrAuthenticationInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"acrAuthenticationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IotSecuritySolutionRecommendationsOutputReference) AgentSendUnutilizedMsg() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"agentSendUnutilizedMsg",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IotSecuritySolutionRecommendationsOutputReference) AgentSendUnutilizedMsgInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"agentSendUnutilizedMsgInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IotSecuritySolutionRecommendationsOutputReference) Baseline() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"baseline",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IotSecuritySolutionRecommendationsOutputReference) BaselineInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"baselineInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IotSecuritySolutionRecommendationsOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IotSecuritySolutionRecommendationsOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IotSecuritySolutionRecommendationsOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IotSecuritySolutionRecommendationsOutputReference) EdgeHubMemOptimize() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"edgeHubMemOptimize",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IotSecuritySolutionRecommendationsOutputReference) EdgeHubMemOptimizeInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"edgeHubMemOptimizeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IotSecuritySolutionRecommendationsOutputReference) EdgeLoggingOption() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"edgeLoggingOption",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IotSecuritySolutionRecommendationsOutputReference) EdgeLoggingOptionInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"edgeLoggingOptionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IotSecuritySolutionRecommendationsOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IotSecuritySolutionRecommendationsOutputReference) InconsistentModuleSettings() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"inconsistentModuleSettings",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IotSecuritySolutionRecommendationsOutputReference) InconsistentModuleSettingsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"inconsistentModuleSettingsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IotSecuritySolutionRecommendationsOutputReference) InstallAgent() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"installAgent",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IotSecuritySolutionRecommendationsOutputReference) InstallAgentInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"installAgentInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IotSecuritySolutionRecommendationsOutputReference) InternalValue() *IotSecuritySolutionRecommendations {
	var returns *IotSecuritySolutionRecommendations
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IotSecuritySolutionRecommendationsOutputReference) IpFilterDenyAll() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"ipFilterDenyAll",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IotSecuritySolutionRecommendationsOutputReference) IpFilterDenyAllInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"ipFilterDenyAllInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IotSecuritySolutionRecommendationsOutputReference) IpFilterPermissiveRule() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"ipFilterPermissiveRule",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IotSecuritySolutionRecommendationsOutputReference) IpFilterPermissiveRuleInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"ipFilterPermissiveRuleInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IotSecuritySolutionRecommendationsOutputReference) OpenPorts() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"openPorts",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IotSecuritySolutionRecommendationsOutputReference) OpenPortsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"openPortsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IotSecuritySolutionRecommendationsOutputReference) PermissiveFirewallPolicy() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"permissiveFirewallPolicy",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IotSecuritySolutionRecommendationsOutputReference) PermissiveFirewallPolicyInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"permissiveFirewallPolicyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IotSecuritySolutionRecommendationsOutputReference) PermissiveInputFirewallRules() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"permissiveInputFirewallRules",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IotSecuritySolutionRecommendationsOutputReference) PermissiveInputFirewallRulesInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"permissiveInputFirewallRulesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IotSecuritySolutionRecommendationsOutputReference) PermissiveOutputFirewallRules() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"permissiveOutputFirewallRules",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IotSecuritySolutionRecommendationsOutputReference) PermissiveOutputFirewallRulesInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"permissiveOutputFirewallRulesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IotSecuritySolutionRecommendationsOutputReference) PrivilegedDockerOptions() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"privilegedDockerOptions",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IotSecuritySolutionRecommendationsOutputReference) PrivilegedDockerOptionsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"privilegedDockerOptionsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IotSecuritySolutionRecommendationsOutputReference) SharedCredentials() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"sharedCredentials",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IotSecuritySolutionRecommendationsOutputReference) SharedCredentialsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"sharedCredentialsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IotSecuritySolutionRecommendationsOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IotSecuritySolutionRecommendationsOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IotSecuritySolutionRecommendationsOutputReference) VulnerableTlsCipherSuite() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"vulnerableTlsCipherSuite",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IotSecuritySolutionRecommendationsOutputReference) VulnerableTlsCipherSuiteInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"vulnerableTlsCipherSuiteInput",
		&returns,
	)
	return returns
}


func NewIotSecuritySolutionRecommendationsOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) IotSecuritySolutionRecommendationsOutputReference {
	_init_.Initialize()

	if err := validateNewIotSecuritySolutionRecommendationsOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_IotSecuritySolutionRecommendationsOutputReference{}

	_jsii_.Create(
		"@cdktn/provider-azurerm.iotSecuritySolution.IotSecuritySolutionRecommendationsOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewIotSecuritySolutionRecommendationsOutputReference_Override(i IotSecuritySolutionRecommendationsOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/provider-azurerm.iotSecuritySolution.IotSecuritySolutionRecommendationsOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		i,
	)
}

func (j *jsiiProxy_IotSecuritySolutionRecommendationsOutputReference)SetAcrAuthentication(val interface{}) {
	if err := j.validateSetAcrAuthenticationParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"acrAuthentication",
		val,
	)
}

func (j *jsiiProxy_IotSecuritySolutionRecommendationsOutputReference)SetAgentSendUnutilizedMsg(val interface{}) {
	if err := j.validateSetAgentSendUnutilizedMsgParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"agentSendUnutilizedMsg",
		val,
	)
}

func (j *jsiiProxy_IotSecuritySolutionRecommendationsOutputReference)SetBaseline(val interface{}) {
	if err := j.validateSetBaselineParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"baseline",
		val,
	)
}

func (j *jsiiProxy_IotSecuritySolutionRecommendationsOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_IotSecuritySolutionRecommendationsOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_IotSecuritySolutionRecommendationsOutputReference)SetEdgeHubMemOptimize(val interface{}) {
	if err := j.validateSetEdgeHubMemOptimizeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"edgeHubMemOptimize",
		val,
	)
}

func (j *jsiiProxy_IotSecuritySolutionRecommendationsOutputReference)SetEdgeLoggingOption(val interface{}) {
	if err := j.validateSetEdgeLoggingOptionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"edgeLoggingOption",
		val,
	)
}

func (j *jsiiProxy_IotSecuritySolutionRecommendationsOutputReference)SetInconsistentModuleSettings(val interface{}) {
	if err := j.validateSetInconsistentModuleSettingsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"inconsistentModuleSettings",
		val,
	)
}

func (j *jsiiProxy_IotSecuritySolutionRecommendationsOutputReference)SetInstallAgent(val interface{}) {
	if err := j.validateSetInstallAgentParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"installAgent",
		val,
	)
}

func (j *jsiiProxy_IotSecuritySolutionRecommendationsOutputReference)SetInternalValue(val *IotSecuritySolutionRecommendations) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_IotSecuritySolutionRecommendationsOutputReference)SetIpFilterDenyAll(val interface{}) {
	if err := j.validateSetIpFilterDenyAllParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"ipFilterDenyAll",
		val,
	)
}

func (j *jsiiProxy_IotSecuritySolutionRecommendationsOutputReference)SetIpFilterPermissiveRule(val interface{}) {
	if err := j.validateSetIpFilterPermissiveRuleParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"ipFilterPermissiveRule",
		val,
	)
}

func (j *jsiiProxy_IotSecuritySolutionRecommendationsOutputReference)SetOpenPorts(val interface{}) {
	if err := j.validateSetOpenPortsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"openPorts",
		val,
	)
}

func (j *jsiiProxy_IotSecuritySolutionRecommendationsOutputReference)SetPermissiveFirewallPolicy(val interface{}) {
	if err := j.validateSetPermissiveFirewallPolicyParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"permissiveFirewallPolicy",
		val,
	)
}

func (j *jsiiProxy_IotSecuritySolutionRecommendationsOutputReference)SetPermissiveInputFirewallRules(val interface{}) {
	if err := j.validateSetPermissiveInputFirewallRulesParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"permissiveInputFirewallRules",
		val,
	)
}

func (j *jsiiProxy_IotSecuritySolutionRecommendationsOutputReference)SetPermissiveOutputFirewallRules(val interface{}) {
	if err := j.validateSetPermissiveOutputFirewallRulesParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"permissiveOutputFirewallRules",
		val,
	)
}

func (j *jsiiProxy_IotSecuritySolutionRecommendationsOutputReference)SetPrivilegedDockerOptions(val interface{}) {
	if err := j.validateSetPrivilegedDockerOptionsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"privilegedDockerOptions",
		val,
	)
}

func (j *jsiiProxy_IotSecuritySolutionRecommendationsOutputReference)SetSharedCredentials(val interface{}) {
	if err := j.validateSetSharedCredentialsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"sharedCredentials",
		val,
	)
}

func (j *jsiiProxy_IotSecuritySolutionRecommendationsOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_IotSecuritySolutionRecommendationsOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_IotSecuritySolutionRecommendationsOutputReference)SetVulnerableTlsCipherSuite(val interface{}) {
	if err := j.validateSetVulnerableTlsCipherSuiteParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"vulnerableTlsCipherSuite",
		val,
	)
}

func (i *jsiiProxy_IotSecuritySolutionRecommendationsOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		i,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (i *jsiiProxy_IotSecuritySolutionRecommendationsOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	if err := i.validateGetAnyMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]interface{}

	_jsii_.Invoke(
		i,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (i *jsiiProxy_IotSecuritySolutionRecommendationsOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
	if err := i.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		i,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (i *jsiiProxy_IotSecuritySolutionRecommendationsOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	if err := i.validateGetBooleanMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*bool

	_jsii_.Invoke(
		i,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (i *jsiiProxy_IotSecuritySolutionRecommendationsOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	if err := i.validateGetListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*string

	_jsii_.Invoke(
		i,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (i *jsiiProxy_IotSecuritySolutionRecommendationsOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	if err := i.validateGetNumberAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *float64

	_jsii_.Invoke(
		i,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (i *jsiiProxy_IotSecuritySolutionRecommendationsOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	if err := i.validateGetNumberListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*float64

	_jsii_.Invoke(
		i,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (i *jsiiProxy_IotSecuritySolutionRecommendationsOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	if err := i.validateGetNumberMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*float64

	_jsii_.Invoke(
		i,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (i *jsiiProxy_IotSecuritySolutionRecommendationsOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	if err := i.validateGetStringAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		i,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (i *jsiiProxy_IotSecuritySolutionRecommendationsOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	if err := i.validateGetStringMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*string

	_jsii_.Invoke(
		i,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (i *jsiiProxy_IotSecuritySolutionRecommendationsOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		i,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (i *jsiiProxy_IotSecuritySolutionRecommendationsOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
	if err := i.validateInterpolationForAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		i,
		"interpolationForAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (i *jsiiProxy_IotSecuritySolutionRecommendationsOutputReference) ResetAcrAuthentication() {
	_jsii_.InvokeVoid(
		i,
		"resetAcrAuthentication",
		nil, // no parameters
	)
}

func (i *jsiiProxy_IotSecuritySolutionRecommendationsOutputReference) ResetAgentSendUnutilizedMsg() {
	_jsii_.InvokeVoid(
		i,
		"resetAgentSendUnutilizedMsg",
		nil, // no parameters
	)
}

func (i *jsiiProxy_IotSecuritySolutionRecommendationsOutputReference) ResetBaseline() {
	_jsii_.InvokeVoid(
		i,
		"resetBaseline",
		nil, // no parameters
	)
}

func (i *jsiiProxy_IotSecuritySolutionRecommendationsOutputReference) ResetEdgeHubMemOptimize() {
	_jsii_.InvokeVoid(
		i,
		"resetEdgeHubMemOptimize",
		nil, // no parameters
	)
}

func (i *jsiiProxy_IotSecuritySolutionRecommendationsOutputReference) ResetEdgeLoggingOption() {
	_jsii_.InvokeVoid(
		i,
		"resetEdgeLoggingOption",
		nil, // no parameters
	)
}

func (i *jsiiProxy_IotSecuritySolutionRecommendationsOutputReference) ResetInconsistentModuleSettings() {
	_jsii_.InvokeVoid(
		i,
		"resetInconsistentModuleSettings",
		nil, // no parameters
	)
}

func (i *jsiiProxy_IotSecuritySolutionRecommendationsOutputReference) ResetInstallAgent() {
	_jsii_.InvokeVoid(
		i,
		"resetInstallAgent",
		nil, // no parameters
	)
}

func (i *jsiiProxy_IotSecuritySolutionRecommendationsOutputReference) ResetIpFilterDenyAll() {
	_jsii_.InvokeVoid(
		i,
		"resetIpFilterDenyAll",
		nil, // no parameters
	)
}

func (i *jsiiProxy_IotSecuritySolutionRecommendationsOutputReference) ResetIpFilterPermissiveRule() {
	_jsii_.InvokeVoid(
		i,
		"resetIpFilterPermissiveRule",
		nil, // no parameters
	)
}

func (i *jsiiProxy_IotSecuritySolutionRecommendationsOutputReference) ResetOpenPorts() {
	_jsii_.InvokeVoid(
		i,
		"resetOpenPorts",
		nil, // no parameters
	)
}

func (i *jsiiProxy_IotSecuritySolutionRecommendationsOutputReference) ResetPermissiveFirewallPolicy() {
	_jsii_.InvokeVoid(
		i,
		"resetPermissiveFirewallPolicy",
		nil, // no parameters
	)
}

func (i *jsiiProxy_IotSecuritySolutionRecommendationsOutputReference) ResetPermissiveInputFirewallRules() {
	_jsii_.InvokeVoid(
		i,
		"resetPermissiveInputFirewallRules",
		nil, // no parameters
	)
}

func (i *jsiiProxy_IotSecuritySolutionRecommendationsOutputReference) ResetPermissiveOutputFirewallRules() {
	_jsii_.InvokeVoid(
		i,
		"resetPermissiveOutputFirewallRules",
		nil, // no parameters
	)
}

func (i *jsiiProxy_IotSecuritySolutionRecommendationsOutputReference) ResetPrivilegedDockerOptions() {
	_jsii_.InvokeVoid(
		i,
		"resetPrivilegedDockerOptions",
		nil, // no parameters
	)
}

func (i *jsiiProxy_IotSecuritySolutionRecommendationsOutputReference) ResetSharedCredentials() {
	_jsii_.InvokeVoid(
		i,
		"resetSharedCredentials",
		nil, // no parameters
	)
}

func (i *jsiiProxy_IotSecuritySolutionRecommendationsOutputReference) ResetVulnerableTlsCipherSuite() {
	_jsii_.InvokeVoid(
		i,
		"resetVulnerableTlsCipherSuite",
		nil, // no parameters
	)
}

func (i *jsiiProxy_IotSecuritySolutionRecommendationsOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
	if err := i.validateResolveParameters(context); err != nil {
		panic(err)
	}
	var returns interface{}

	_jsii_.Invoke(
		i,
		"resolve",
		[]interface{}{context},
		&returns,
	)

	return returns
}

func (i *jsiiProxy_IotSecuritySolutionRecommendationsOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		i,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

