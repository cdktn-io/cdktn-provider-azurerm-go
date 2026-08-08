// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package dataazurermsiterecoveryreplicationrecoveryplan

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-provider-azurerm-go/azurerm/v18/jsii"

	"github.com/cdktn-io/cdktn-provider-azurerm-go/azurerm/v18/dataazurermsiterecoveryreplicationrecoveryplan/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type DataAzurermSiteRecoveryReplicationRecoveryPlanRecoveryGroupPreActionListList interface {
	cdktn.MapList
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
	// Experimental.
	ComputeFqn() *string
	Get(index *float64) DataAzurermSiteRecoveryReplicationRecoveryPlanRecoveryGroupPreActionList
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

// The jsii proxy struct for DataAzurermSiteRecoveryReplicationRecoveryPlanRecoveryGroupPreActionListList
type jsiiProxy_DataAzurermSiteRecoveryReplicationRecoveryPlanRecoveryGroupPreActionListList struct {
	internal.Type__cdktnMapList
}

func (j *jsiiProxy_DataAzurermSiteRecoveryReplicationRecoveryPlanRecoveryGroupPreActionListList) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAzurermSiteRecoveryReplicationRecoveryPlanRecoveryGroupPreActionListList) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAzurermSiteRecoveryReplicationRecoveryPlanRecoveryGroupPreActionListList) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAzurermSiteRecoveryReplicationRecoveryPlanRecoveryGroupPreActionListList) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAzurermSiteRecoveryReplicationRecoveryPlanRecoveryGroupPreActionListList) WrapsSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"wrapsSet",
		&returns,
	)
	return returns
}


func NewDataAzurermSiteRecoveryReplicationRecoveryPlanRecoveryGroupPreActionListList(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) DataAzurermSiteRecoveryReplicationRecoveryPlanRecoveryGroupPreActionListList {
	_init_.Initialize()

	if err := validateNewDataAzurermSiteRecoveryReplicationRecoveryPlanRecoveryGroupPreActionListListParameters(terraformResource, terraformAttribute, wrapsSet); err != nil {
		panic(err)
	}
	j := jsiiProxy_DataAzurermSiteRecoveryReplicationRecoveryPlanRecoveryGroupPreActionListList{}

	_jsii_.Create(
		"@cdktn/provider-azurerm.dataAzurermSiteRecoveryReplicationRecoveryPlan.DataAzurermSiteRecoveryReplicationRecoveryPlanRecoveryGroupPreActionListList",
		[]interface{}{terraformResource, terraformAttribute, wrapsSet},
		&j,
	)

	return &j
}

func NewDataAzurermSiteRecoveryReplicationRecoveryPlanRecoveryGroupPreActionListList_Override(d DataAzurermSiteRecoveryReplicationRecoveryPlanRecoveryGroupPreActionListList, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/provider-azurerm.dataAzurermSiteRecoveryReplicationRecoveryPlan.DataAzurermSiteRecoveryReplicationRecoveryPlanRecoveryGroupPreActionListList",
		[]interface{}{terraformResource, terraformAttribute, wrapsSet},
		d,
	)
}

func (j *jsiiProxy_DataAzurermSiteRecoveryReplicationRecoveryPlanRecoveryGroupPreActionListList)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_DataAzurermSiteRecoveryReplicationRecoveryPlanRecoveryGroupPreActionListList)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_DataAzurermSiteRecoveryReplicationRecoveryPlanRecoveryGroupPreActionListList)SetWrapsSet(val *bool) {
	if err := j.validateSetWrapsSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"wrapsSet",
		val,
	)
}

func (d *jsiiProxy_DataAzurermSiteRecoveryReplicationRecoveryPlanRecoveryGroupPreActionListList) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataAzurermSiteRecoveryReplicationRecoveryPlanRecoveryGroupPreActionListList) Get(index *float64) DataAzurermSiteRecoveryReplicationRecoveryPlanRecoveryGroupPreActionList {
	if err := d.validateGetParameters(index); err != nil {
		panic(err)
	}
	var returns DataAzurermSiteRecoveryReplicationRecoveryPlanRecoveryGroupPreActionList

	_jsii_.Invoke(
		d,
		"get",
		[]interface{}{index},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataAzurermSiteRecoveryReplicationRecoveryPlanRecoveryGroupPreActionListList) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (d *jsiiProxy_DataAzurermSiteRecoveryReplicationRecoveryPlanRecoveryGroupPreActionListList) Resolve(context cdktn.IResolveContext) interface{} {
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

func (d *jsiiProxy_DataAzurermSiteRecoveryReplicationRecoveryPlanRecoveryGroupPreActionListList) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

