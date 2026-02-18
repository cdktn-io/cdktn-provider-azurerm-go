// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

//go:build no_runtime_type_checking

package privateendpoint

// Building without runtime type checking enabled, so all the below just return nil

func (p *jsiiProxy_PrivateEndpointNetworkInterfaceList) validateAllWithMapKeyParameters(mapKeyAttributeName *string) error {
	return nil
}

func (p *jsiiProxy_PrivateEndpointNetworkInterfaceList) validateGetParameters(index *float64) error {
	return nil
}

func (p *jsiiProxy_PrivateEndpointNetworkInterfaceList) validateResolveParameters(context cdktn.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_PrivateEndpointNetworkInterfaceList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_PrivateEndpointNetworkInterfaceList) validateSetTerraformResourceParameters(val cdktn.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_PrivateEndpointNetworkInterfaceList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewPrivateEndpointNetworkInterfaceListParameters(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}

