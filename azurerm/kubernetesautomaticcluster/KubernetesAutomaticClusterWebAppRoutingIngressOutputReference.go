// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package kubernetesautomaticcluster

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-provider-azurerm-go/azurerm/v16/jsii"

	"github.com/cdktn-io/cdktn-provider-azurerm-go/azurerm/v16/kubernetesautomaticcluster/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type KubernetesAutomaticClusterWebAppRoutingIngressOutputReference interface {
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
	DefaultNginxController() *string
	SetDefaultNginxController(val *string)
	DefaultNginxControllerInput() *string
	DnsZoneIds() *[]*string
	SetDnsZoneIds(val *[]*string)
	DnsZoneIdsInput() *[]*string
	// Experimental.
	Fqn() *string
	InternalValue() *KubernetesAutomaticClusterWebAppRoutingIngress
	SetInternalValue(val *KubernetesAutomaticClusterWebAppRoutingIngress)
	IstioEnabled() interface{}
	SetIstioEnabled(val interface{})
	IstioEnabledInput() interface{}
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktn.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktn.IInterpolatingParent)
	WebAppRoutingIdentity() KubernetesAutomaticClusterWebAppRoutingIngressWebAppRoutingIdentityList
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
	ResetDefaultNginxController()
	ResetDnsZoneIds()
	ResetIstioEnabled()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktn.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for KubernetesAutomaticClusterWebAppRoutingIngressOutputReference
type jsiiProxy_KubernetesAutomaticClusterWebAppRoutingIngressOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_KubernetesAutomaticClusterWebAppRoutingIngressOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KubernetesAutomaticClusterWebAppRoutingIngressOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KubernetesAutomaticClusterWebAppRoutingIngressOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KubernetesAutomaticClusterWebAppRoutingIngressOutputReference) DefaultNginxController() *string {
	var returns *string
	_jsii_.Get(
		j,
		"defaultNginxController",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KubernetesAutomaticClusterWebAppRoutingIngressOutputReference) DefaultNginxControllerInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"defaultNginxControllerInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KubernetesAutomaticClusterWebAppRoutingIngressOutputReference) DnsZoneIds() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dnsZoneIds",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KubernetesAutomaticClusterWebAppRoutingIngressOutputReference) DnsZoneIdsInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dnsZoneIdsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KubernetesAutomaticClusterWebAppRoutingIngressOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KubernetesAutomaticClusterWebAppRoutingIngressOutputReference) InternalValue() *KubernetesAutomaticClusterWebAppRoutingIngress {
	var returns *KubernetesAutomaticClusterWebAppRoutingIngress
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KubernetesAutomaticClusterWebAppRoutingIngressOutputReference) IstioEnabled() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"istioEnabled",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KubernetesAutomaticClusterWebAppRoutingIngressOutputReference) IstioEnabledInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"istioEnabledInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KubernetesAutomaticClusterWebAppRoutingIngressOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KubernetesAutomaticClusterWebAppRoutingIngressOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KubernetesAutomaticClusterWebAppRoutingIngressOutputReference) WebAppRoutingIdentity() KubernetesAutomaticClusterWebAppRoutingIngressWebAppRoutingIdentityList {
	var returns KubernetesAutomaticClusterWebAppRoutingIngressWebAppRoutingIdentityList
	_jsii_.Get(
		j,
		"webAppRoutingIdentity",
		&returns,
	)
	return returns
}


func NewKubernetesAutomaticClusterWebAppRoutingIngressOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) KubernetesAutomaticClusterWebAppRoutingIngressOutputReference {
	_init_.Initialize()

	if err := validateNewKubernetesAutomaticClusterWebAppRoutingIngressOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_KubernetesAutomaticClusterWebAppRoutingIngressOutputReference{}

	_jsii_.Create(
		"@cdktn/provider-azurerm.kubernetesAutomaticCluster.KubernetesAutomaticClusterWebAppRoutingIngressOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewKubernetesAutomaticClusterWebAppRoutingIngressOutputReference_Override(k KubernetesAutomaticClusterWebAppRoutingIngressOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/provider-azurerm.kubernetesAutomaticCluster.KubernetesAutomaticClusterWebAppRoutingIngressOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		k,
	)
}

func (j *jsiiProxy_KubernetesAutomaticClusterWebAppRoutingIngressOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_KubernetesAutomaticClusterWebAppRoutingIngressOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_KubernetesAutomaticClusterWebAppRoutingIngressOutputReference)SetDefaultNginxController(val *string) {
	if err := j.validateSetDefaultNginxControllerParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"defaultNginxController",
		val,
	)
}

func (j *jsiiProxy_KubernetesAutomaticClusterWebAppRoutingIngressOutputReference)SetDnsZoneIds(val *[]*string) {
	if err := j.validateSetDnsZoneIdsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"dnsZoneIds",
		val,
	)
}

func (j *jsiiProxy_KubernetesAutomaticClusterWebAppRoutingIngressOutputReference)SetInternalValue(val *KubernetesAutomaticClusterWebAppRoutingIngress) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_KubernetesAutomaticClusterWebAppRoutingIngressOutputReference)SetIstioEnabled(val interface{}) {
	if err := j.validateSetIstioEnabledParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"istioEnabled",
		val,
	)
}

func (j *jsiiProxy_KubernetesAutomaticClusterWebAppRoutingIngressOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_KubernetesAutomaticClusterWebAppRoutingIngressOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (k *jsiiProxy_KubernetesAutomaticClusterWebAppRoutingIngressOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		k,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (k *jsiiProxy_KubernetesAutomaticClusterWebAppRoutingIngressOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	if err := k.validateGetAnyMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]interface{}

	_jsii_.Invoke(
		k,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (k *jsiiProxy_KubernetesAutomaticClusterWebAppRoutingIngressOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
	if err := k.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		k,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (k *jsiiProxy_KubernetesAutomaticClusterWebAppRoutingIngressOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	if err := k.validateGetBooleanMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*bool

	_jsii_.Invoke(
		k,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (k *jsiiProxy_KubernetesAutomaticClusterWebAppRoutingIngressOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	if err := k.validateGetListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*string

	_jsii_.Invoke(
		k,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (k *jsiiProxy_KubernetesAutomaticClusterWebAppRoutingIngressOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	if err := k.validateGetNumberAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *float64

	_jsii_.Invoke(
		k,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (k *jsiiProxy_KubernetesAutomaticClusterWebAppRoutingIngressOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	if err := k.validateGetNumberListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*float64

	_jsii_.Invoke(
		k,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (k *jsiiProxy_KubernetesAutomaticClusterWebAppRoutingIngressOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	if err := k.validateGetNumberMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*float64

	_jsii_.Invoke(
		k,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (k *jsiiProxy_KubernetesAutomaticClusterWebAppRoutingIngressOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	if err := k.validateGetStringAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		k,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (k *jsiiProxy_KubernetesAutomaticClusterWebAppRoutingIngressOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	if err := k.validateGetStringMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*string

	_jsii_.Invoke(
		k,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (k *jsiiProxy_KubernetesAutomaticClusterWebAppRoutingIngressOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		k,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (k *jsiiProxy_KubernetesAutomaticClusterWebAppRoutingIngressOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
	if err := k.validateInterpolationForAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		k,
		"interpolationForAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (k *jsiiProxy_KubernetesAutomaticClusterWebAppRoutingIngressOutputReference) ResetDefaultNginxController() {
	_jsii_.InvokeVoid(
		k,
		"resetDefaultNginxController",
		nil, // no parameters
	)
}

func (k *jsiiProxy_KubernetesAutomaticClusterWebAppRoutingIngressOutputReference) ResetDnsZoneIds() {
	_jsii_.InvokeVoid(
		k,
		"resetDnsZoneIds",
		nil, // no parameters
	)
}

func (k *jsiiProxy_KubernetesAutomaticClusterWebAppRoutingIngressOutputReference) ResetIstioEnabled() {
	_jsii_.InvokeVoid(
		k,
		"resetIstioEnabled",
		nil, // no parameters
	)
}

func (k *jsiiProxy_KubernetesAutomaticClusterWebAppRoutingIngressOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
	if err := k.validateResolveParameters(context); err != nil {
		panic(err)
	}
	var returns interface{}

	_jsii_.Invoke(
		k,
		"resolve",
		[]interface{}{context},
		&returns,
	)

	return returns
}

func (k *jsiiProxy_KubernetesAutomaticClusterWebAppRoutingIngressOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		k,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

