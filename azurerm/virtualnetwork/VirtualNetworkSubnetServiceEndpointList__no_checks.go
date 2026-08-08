// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

//go:build no_runtime_type_checking

package virtualnetwork

// Building without runtime type checking enabled, so all the below just return nil

func (v *jsiiProxy_VirtualNetworkSubnetServiceEndpointList) validateAllWithMapKeyParameters(mapKeyAttributeName *string) error {
	return nil
}

func (v *jsiiProxy_VirtualNetworkSubnetServiceEndpointList) validateGetParameters(index *float64) error {
	return nil
}

func (v *jsiiProxy_VirtualNetworkSubnetServiceEndpointList) validateResolveParameters(context cdktn.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_VirtualNetworkSubnetServiceEndpointList) validateSetInternalValueParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_VirtualNetworkSubnetServiceEndpointList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_VirtualNetworkSubnetServiceEndpointList) validateSetTerraformResourceParameters(val cdktn.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_VirtualNetworkSubnetServiceEndpointList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewVirtualNetworkSubnetServiceEndpointListParameters(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}

