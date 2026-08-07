// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package providerfunctions

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-provider-azurerm-go/azurerm/v17/jsii"

	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

// Provider-defined functions of the azurerm provider.
type AzurermProviderFunctions interface {
	// Parses and attempts to normalise the casing on an Azure Resource Manager ID into the correct casing for Terraform.
	NormaliseResourceId(id *string) *string
	// Parses an Azure Resource Manager ID and exposes the contained information.
	ParseResourceId(id *string) cdktn.IResolvable
}

// The jsii proxy struct for AzurermProviderFunctions
type jsiiProxy_AzurermProviderFunctions struct {
	_ byte // padding
}

func NewAzurermProviderFunctions(providerLocalName *string) AzurermProviderFunctions {
	_init_.Initialize()

	if err := validateNewAzurermProviderFunctionsParameters(providerLocalName); err != nil {
		panic(err)
	}
	j := jsiiProxy_AzurermProviderFunctions{}

	_jsii_.Create(
		"@cdktn/provider-azurerm.providerFunctions.AzurermProviderFunctions",
		[]interface{}{providerLocalName},
		&j,
	)

	return &j
}

func NewAzurermProviderFunctions_Override(a AzurermProviderFunctions, providerLocalName *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/provider-azurerm.providerFunctions.AzurermProviderFunctions",
		[]interface{}{providerLocalName},
		a,
	)
}

func (a *jsiiProxy_AzurermProviderFunctions) NormaliseResourceId(id *string) *string {
	if err := a.validateNormaliseResourceIdParameters(id); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		a,
		"normaliseResourceId",
		[]interface{}{id},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AzurermProviderFunctions) ParseResourceId(id *string) cdktn.IResolvable {
	if err := a.validateParseResourceIdParameters(id); err != nil {
		panic(err)
	}
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		a,
		"parseResourceId",
		[]interface{}{id},
		&returns,
	)

	return returns
}

