// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

//go:build no_runtime_type_checking

package lb

// Building without runtime type checking enabled, so all the below just return nil

func (l *jsiiProxy_LbFrontendIpConfigurationList) validateAllWithMapKeyParameters(mapKeyAttributeName *string) error {
	return nil
}

func (l *jsiiProxy_LbFrontendIpConfigurationList) validateGetParameters(index *float64) error {
	return nil
}

func (l *jsiiProxy_LbFrontendIpConfigurationList) validateResolveParameters(context cdktn.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_LbFrontendIpConfigurationList) validateSetInternalValueParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_LbFrontendIpConfigurationList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_LbFrontendIpConfigurationList) validateSetTerraformResourceParameters(val cdktn.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_LbFrontendIpConfigurationList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewLbFrontendIpConfigurationListParameters(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}

