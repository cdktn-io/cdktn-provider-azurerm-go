// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

//go:build no_runtime_type_checking

package providerfunctions

// Building without runtime type checking enabled, so all the below just return nil

func (a *jsiiProxy_AzurermProviderFunctions) validateNormaliseResourceIdParameters(id *string) error {
	return nil
}

func (a *jsiiProxy_AzurermProviderFunctions) validateParseResourceIdParameters(id *string) error {
	return nil
}

func validateNewAzurermProviderFunctionsParameters(providerLocalName *string) error {
	return nil
}

