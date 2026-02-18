// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

//go:build no_runtime_type_checking

package dataazurermimage

// Building without runtime type checking enabled, so all the below just return nil

func (d *jsiiProxy_DataAzurermImageOsDiskList) validateAllWithMapKeyParameters(mapKeyAttributeName *string) error {
	return nil
}

func (d *jsiiProxy_DataAzurermImageOsDiskList) validateGetParameters(index *float64) error {
	return nil
}

func (d *jsiiProxy_DataAzurermImageOsDiskList) validateResolveParameters(context cdktn.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_DataAzurermImageOsDiskList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_DataAzurermImageOsDiskList) validateSetTerraformResourceParameters(val cdktn.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_DataAzurermImageOsDiskList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewDataAzurermImageOsDiskListParameters(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}

