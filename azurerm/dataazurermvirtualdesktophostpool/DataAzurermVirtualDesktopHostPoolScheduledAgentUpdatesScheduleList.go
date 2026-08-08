// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package dataazurermvirtualdesktophostpool

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-provider-azurerm-go/azurerm/v18/jsii"

	"github.com/cdktn-io/cdktn-provider-azurerm-go/azurerm/v18/dataazurermvirtualdesktophostpool/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type DataAzurermVirtualDesktopHostPoolScheduledAgentUpdatesScheduleList interface {
	cdktn.ComplexList
	// The creation stack of this resolvable which will be appended to errors thrown during resolution.
	//
	// If this returns an empty array the stack will not be attached.
	// Experimental.
	CreationStack() *[]*string
	// Experimental.
	Fqn() *string
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktn.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktn.IInterpolatingParent)
	// Experimental.
	WrapsSet() *bool
	// Experimental.
	SetWrapsSet(val *bool)
	// Creating an iterator for this complex list.
	//
	// The list will be converted into a map with the mapKeyAttributeName as the key.
	// Experimental.
	AllWithMapKey(mapKeyAttributeName *string) cdktn.DynamicListTerraformIterator
	// Experimental.
	ComputeFqn() *string
	Get(index *float64) DataAzurermVirtualDesktopHostPoolScheduledAgentUpdatesScheduleOutputReference
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktn.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for DataAzurermVirtualDesktopHostPoolScheduledAgentUpdatesScheduleList
type jsiiProxy_DataAzurermVirtualDesktopHostPoolScheduledAgentUpdatesScheduleList struct {
	internal.Type__cdktnComplexList
}

func (j *jsiiProxy_DataAzurermVirtualDesktopHostPoolScheduledAgentUpdatesScheduleList) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAzurermVirtualDesktopHostPoolScheduledAgentUpdatesScheduleList) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAzurermVirtualDesktopHostPoolScheduledAgentUpdatesScheduleList) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAzurermVirtualDesktopHostPoolScheduledAgentUpdatesScheduleList) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAzurermVirtualDesktopHostPoolScheduledAgentUpdatesScheduleList) WrapsSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"wrapsSet",
		&returns,
	)
	return returns
}


func NewDataAzurermVirtualDesktopHostPoolScheduledAgentUpdatesScheduleList(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) DataAzurermVirtualDesktopHostPoolScheduledAgentUpdatesScheduleList {
	_init_.Initialize()

	if err := validateNewDataAzurermVirtualDesktopHostPoolScheduledAgentUpdatesScheduleListParameters(terraformResource, terraformAttribute, wrapsSet); err != nil {
		panic(err)
	}
	j := jsiiProxy_DataAzurermVirtualDesktopHostPoolScheduledAgentUpdatesScheduleList{}

	_jsii_.Create(
		"@cdktn/provider-azurerm.dataAzurermVirtualDesktopHostPool.DataAzurermVirtualDesktopHostPoolScheduledAgentUpdatesScheduleList",
		[]interface{}{terraformResource, terraformAttribute, wrapsSet},
		&j,
	)

	return &j
}

func NewDataAzurermVirtualDesktopHostPoolScheduledAgentUpdatesScheduleList_Override(d DataAzurermVirtualDesktopHostPoolScheduledAgentUpdatesScheduleList, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/provider-azurerm.dataAzurermVirtualDesktopHostPool.DataAzurermVirtualDesktopHostPoolScheduledAgentUpdatesScheduleList",
		[]interface{}{terraformResource, terraformAttribute, wrapsSet},
		d,
	)
}

func (j *jsiiProxy_DataAzurermVirtualDesktopHostPoolScheduledAgentUpdatesScheduleList)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_DataAzurermVirtualDesktopHostPoolScheduledAgentUpdatesScheduleList)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_DataAzurermVirtualDesktopHostPoolScheduledAgentUpdatesScheduleList)SetWrapsSet(val *bool) {
	if err := j.validateSetWrapsSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"wrapsSet",
		val,
	)
}

func (d *jsiiProxy_DataAzurermVirtualDesktopHostPoolScheduledAgentUpdatesScheduleList) AllWithMapKey(mapKeyAttributeName *string) cdktn.DynamicListTerraformIterator {
	if err := d.validateAllWithMapKeyParameters(mapKeyAttributeName); err != nil {
		panic(err)
	}
	var returns cdktn.DynamicListTerraformIterator

	_jsii_.Invoke(
		d,
		"allWithMapKey",
		[]interface{}{mapKeyAttributeName},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataAzurermVirtualDesktopHostPoolScheduledAgentUpdatesScheduleList) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataAzurermVirtualDesktopHostPoolScheduledAgentUpdatesScheduleList) Get(index *float64) DataAzurermVirtualDesktopHostPoolScheduledAgentUpdatesScheduleOutputReference {
	if err := d.validateGetParameters(index); err != nil {
		panic(err)
	}
	var returns DataAzurermVirtualDesktopHostPoolScheduledAgentUpdatesScheduleOutputReference

	_jsii_.Invoke(
		d,
		"get",
		[]interface{}{index},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataAzurermVirtualDesktopHostPoolScheduledAgentUpdatesScheduleList) Resolve(context cdktn.IResolveContext) interface{} {
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

func (d *jsiiProxy_DataAzurermVirtualDesktopHostPoolScheduledAgentUpdatesScheduleList) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

