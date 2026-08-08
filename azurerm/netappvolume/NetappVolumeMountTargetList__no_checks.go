// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

//go:build no_runtime_type_checking

package netappvolume

// Building without runtime type checking enabled, so all the below just return nil

func (n *jsiiProxy_NetappVolumeMountTargetList) validateAllWithMapKeyParameters(mapKeyAttributeName *string) error {
	return nil
}

func (n *jsiiProxy_NetappVolumeMountTargetList) validateGetParameters(index *float64) error {
	return nil
}

func (n *jsiiProxy_NetappVolumeMountTargetList) validateResolveParameters(context cdktn.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_NetappVolumeMountTargetList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_NetappVolumeMountTargetList) validateSetTerraformResourceParameters(val cdktn.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_NetappVolumeMountTargetList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewNetappVolumeMountTargetListParameters(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}

