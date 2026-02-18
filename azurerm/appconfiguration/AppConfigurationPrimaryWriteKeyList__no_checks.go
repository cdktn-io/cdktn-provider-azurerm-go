// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

//go:build no_runtime_type_checking

package appconfiguration

// Building without runtime type checking enabled, so all the below just return nil

func (a *jsiiProxy_AppConfigurationPrimaryWriteKeyList) validateAllWithMapKeyParameters(mapKeyAttributeName *string) error {
	return nil
}

func (a *jsiiProxy_AppConfigurationPrimaryWriteKeyList) validateGetParameters(index *float64) error {
	return nil
}

func (a *jsiiProxy_AppConfigurationPrimaryWriteKeyList) validateResolveParameters(context cdktn.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_AppConfigurationPrimaryWriteKeyList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_AppConfigurationPrimaryWriteKeyList) validateSetTerraformResourceParameters(val cdktn.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_AppConfigurationPrimaryWriteKeyList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewAppConfigurationPrimaryWriteKeyListParameters(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}

