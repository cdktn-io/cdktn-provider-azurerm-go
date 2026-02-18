// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

//go:build no_runtime_type_checking

package linuxwebapp

// Building without runtime type checking enabled, so all the below just return nil

func (l *jsiiProxy_LinuxWebAppConnectionStringList) validateAllWithMapKeyParameters(mapKeyAttributeName *string) error {
	return nil
}

func (l *jsiiProxy_LinuxWebAppConnectionStringList) validateGetParameters(index *float64) error {
	return nil
}

func (l *jsiiProxy_LinuxWebAppConnectionStringList) validateResolveParameters(context cdktn.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_LinuxWebAppConnectionStringList) validateSetInternalValueParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_LinuxWebAppConnectionStringList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_LinuxWebAppConnectionStringList) validateSetTerraformResourceParameters(val cdktn.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_LinuxWebAppConnectionStringList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewLinuxWebAppConnectionStringListParameters(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}

