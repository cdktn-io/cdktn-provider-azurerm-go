// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

//go:build no_runtime_type_checking

package customprovider

// Building without runtime type checking enabled, so all the below just return nil

func (c *jsiiProxy_CustomProviderValidationList) validateAllWithMapKeyParameters(mapKeyAttributeName *string) error {
	return nil
}

func (c *jsiiProxy_CustomProviderValidationList) validateGetParameters(index *float64) error {
	return nil
}

func (c *jsiiProxy_CustomProviderValidationList) validateResolveParameters(context cdktn.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_CustomProviderValidationList) validateSetInternalValueParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_CustomProviderValidationList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_CustomProviderValidationList) validateSetTerraformResourceParameters(val cdktn.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_CustomProviderValidationList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewCustomProviderValidationListParameters(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}

