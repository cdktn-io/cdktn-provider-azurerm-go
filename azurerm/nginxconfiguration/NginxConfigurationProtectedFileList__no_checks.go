// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

//go:build no_runtime_type_checking

package nginxconfiguration

// Building without runtime type checking enabled, so all the below just return nil

func (n *jsiiProxy_NginxConfigurationProtectedFileList) validateAllWithMapKeyParameters(mapKeyAttributeName *string) error {
	return nil
}

func (n *jsiiProxy_NginxConfigurationProtectedFileList) validateGetParameters(index *float64) error {
	return nil
}

func (n *jsiiProxy_NginxConfigurationProtectedFileList) validateResolveParameters(context cdktn.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_NginxConfigurationProtectedFileList) validateSetInternalValueParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_NginxConfigurationProtectedFileList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_NginxConfigurationProtectedFileList) validateSetTerraformResourceParameters(val cdktn.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_NginxConfigurationProtectedFileList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewNginxConfigurationProtectedFileListParameters(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}

