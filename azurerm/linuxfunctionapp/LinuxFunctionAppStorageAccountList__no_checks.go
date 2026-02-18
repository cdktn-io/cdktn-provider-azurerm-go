// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

//go:build no_runtime_type_checking

package linuxfunctionapp

// Building without runtime type checking enabled, so all the below just return nil

func (l *jsiiProxy_LinuxFunctionAppStorageAccountList) validateAllWithMapKeyParameters(mapKeyAttributeName *string) error {
	return nil
}

func (l *jsiiProxy_LinuxFunctionAppStorageAccountList) validateGetParameters(index *float64) error {
	return nil
}

func (l *jsiiProxy_LinuxFunctionAppStorageAccountList) validateResolveParameters(context cdktn.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_LinuxFunctionAppStorageAccountList) validateSetInternalValueParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_LinuxFunctionAppStorageAccountList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_LinuxFunctionAppStorageAccountList) validateSetTerraformResourceParameters(val cdktn.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_LinuxFunctionAppStorageAccountList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewLinuxFunctionAppStorageAccountListParameters(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}

