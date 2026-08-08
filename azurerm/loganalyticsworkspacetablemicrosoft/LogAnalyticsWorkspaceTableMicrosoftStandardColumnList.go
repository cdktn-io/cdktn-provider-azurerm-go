// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package loganalyticsworkspacetablemicrosoft

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-provider-azurerm-go/azurerm/v18/jsii"

	"github.com/cdktn-io/cdktn-provider-azurerm-go/azurerm/v18/loganalyticsworkspacetablemicrosoft/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type LogAnalyticsWorkspaceTableMicrosoftStandardColumnList interface {
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
	Get(index *float64) LogAnalyticsWorkspaceTableMicrosoftStandardColumnOutputReference
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktn.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for LogAnalyticsWorkspaceTableMicrosoftStandardColumnList
type jsiiProxy_LogAnalyticsWorkspaceTableMicrosoftStandardColumnList struct {
	internal.Type__cdktnComplexList
}

func (j *jsiiProxy_LogAnalyticsWorkspaceTableMicrosoftStandardColumnList) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LogAnalyticsWorkspaceTableMicrosoftStandardColumnList) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LogAnalyticsWorkspaceTableMicrosoftStandardColumnList) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LogAnalyticsWorkspaceTableMicrosoftStandardColumnList) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LogAnalyticsWorkspaceTableMicrosoftStandardColumnList) WrapsSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"wrapsSet",
		&returns,
	)
	return returns
}


func NewLogAnalyticsWorkspaceTableMicrosoftStandardColumnList(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) LogAnalyticsWorkspaceTableMicrosoftStandardColumnList {
	_init_.Initialize()

	if err := validateNewLogAnalyticsWorkspaceTableMicrosoftStandardColumnListParameters(terraformResource, terraformAttribute, wrapsSet); err != nil {
		panic(err)
	}
	j := jsiiProxy_LogAnalyticsWorkspaceTableMicrosoftStandardColumnList{}

	_jsii_.Create(
		"@cdktn/provider-azurerm.logAnalyticsWorkspaceTableMicrosoft.LogAnalyticsWorkspaceTableMicrosoftStandardColumnList",
		[]interface{}{terraformResource, terraformAttribute, wrapsSet},
		&j,
	)

	return &j
}

func NewLogAnalyticsWorkspaceTableMicrosoftStandardColumnList_Override(l LogAnalyticsWorkspaceTableMicrosoftStandardColumnList, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/provider-azurerm.logAnalyticsWorkspaceTableMicrosoft.LogAnalyticsWorkspaceTableMicrosoftStandardColumnList",
		[]interface{}{terraformResource, terraformAttribute, wrapsSet},
		l,
	)
}

func (j *jsiiProxy_LogAnalyticsWorkspaceTableMicrosoftStandardColumnList)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_LogAnalyticsWorkspaceTableMicrosoftStandardColumnList)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_LogAnalyticsWorkspaceTableMicrosoftStandardColumnList)SetWrapsSet(val *bool) {
	if err := j.validateSetWrapsSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"wrapsSet",
		val,
	)
}

func (l *jsiiProxy_LogAnalyticsWorkspaceTableMicrosoftStandardColumnList) AllWithMapKey(mapKeyAttributeName *string) cdktn.DynamicListTerraformIterator {
	if err := l.validateAllWithMapKeyParameters(mapKeyAttributeName); err != nil {
		panic(err)
	}
	var returns cdktn.DynamicListTerraformIterator

	_jsii_.Invoke(
		l,
		"allWithMapKey",
		[]interface{}{mapKeyAttributeName},
		&returns,
	)

	return returns
}

func (l *jsiiProxy_LogAnalyticsWorkspaceTableMicrosoftStandardColumnList) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		l,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (l *jsiiProxy_LogAnalyticsWorkspaceTableMicrosoftStandardColumnList) Get(index *float64) LogAnalyticsWorkspaceTableMicrosoftStandardColumnOutputReference {
	if err := l.validateGetParameters(index); err != nil {
		panic(err)
	}
	var returns LogAnalyticsWorkspaceTableMicrosoftStandardColumnOutputReference

	_jsii_.Invoke(
		l,
		"get",
		[]interface{}{index},
		&returns,
	)

	return returns
}

func (l *jsiiProxy_LogAnalyticsWorkspaceTableMicrosoftStandardColumnList) Resolve(context cdktn.IResolveContext) interface{} {
	if err := l.validateResolveParameters(context); err != nil {
		panic(err)
	}
	var returns interface{}

	_jsii_.Invoke(
		l,
		"resolve",
		[]interface{}{context},
		&returns,
	)

	return returns
}

func (l *jsiiProxy_LogAnalyticsWorkspaceTableMicrosoftStandardColumnList) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		l,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

