// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

//go:build no_runtime_type_checking

package virtualhub

// Building without runtime type checking enabled, so all the below just return nil

func (v *jsiiProxy_VirtualHubRouteList) validateAllWithMapKeyParameters(mapKeyAttributeName *string) error {
	return nil
}

func (v *jsiiProxy_VirtualHubRouteList) validateGetParameters(index *float64) error {
	return nil
}

func (v *jsiiProxy_VirtualHubRouteList) validateResolveParameters(context cdktn.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_VirtualHubRouteList) validateSetInternalValueParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_VirtualHubRouteList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_VirtualHubRouteList) validateSetTerraformResourceParameters(val cdktn.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_VirtualHubRouteList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewVirtualHubRouteListParameters(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}

