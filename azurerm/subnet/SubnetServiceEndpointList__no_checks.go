// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

//go:build no_runtime_type_checking

package subnet

// Building without runtime type checking enabled, so all the below just return nil

func (s *jsiiProxy_SubnetServiceEndpointList) validateAllWithMapKeyParameters(mapKeyAttributeName *string) error {
	return nil
}

func (s *jsiiProxy_SubnetServiceEndpointList) validateGetParameters(index *float64) error {
	return nil
}

func (s *jsiiProxy_SubnetServiceEndpointList) validateResolveParameters(context cdktn.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_SubnetServiceEndpointList) validateSetInternalValueParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_SubnetServiceEndpointList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_SubnetServiceEndpointList) validateSetTerraformResourceParameters(val cdktn.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_SubnetServiceEndpointList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewSubnetServiceEndpointListParameters(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}

