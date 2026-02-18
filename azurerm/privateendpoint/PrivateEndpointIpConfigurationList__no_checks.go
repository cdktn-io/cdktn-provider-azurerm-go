// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

//go:build no_runtime_type_checking

package privateendpoint

// Building without runtime type checking enabled, so all the below just return nil

func (p *jsiiProxy_PrivateEndpointIpConfigurationList) validateAllWithMapKeyParameters(mapKeyAttributeName *string) error {
	return nil
}

func (p *jsiiProxy_PrivateEndpointIpConfigurationList) validateGetParameters(index *float64) error {
	return nil
}

func (p *jsiiProxy_PrivateEndpointIpConfigurationList) validateResolveParameters(context cdktn.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_PrivateEndpointIpConfigurationList) validateSetInternalValueParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_PrivateEndpointIpConfigurationList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_PrivateEndpointIpConfigurationList) validateSetTerraformResourceParameters(val cdktn.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_PrivateEndpointIpConfigurationList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewPrivateEndpointIpConfigurationListParameters(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}

