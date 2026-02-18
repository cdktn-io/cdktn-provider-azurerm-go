// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

//go:build no_runtime_type_checking

package dataazurermloadtest

// Building without runtime type checking enabled, so all the below just return nil

func (d *jsiiProxy_DataAzurermLoadTestIdentityList) validateAllWithMapKeyParameters(mapKeyAttributeName *string) error {
	return nil
}

func (d *jsiiProxy_DataAzurermLoadTestIdentityList) validateGetParameters(index *float64) error {
	return nil
}

func (d *jsiiProxy_DataAzurermLoadTestIdentityList) validateResolveParameters(context cdktn.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_DataAzurermLoadTestIdentityList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_DataAzurermLoadTestIdentityList) validateSetTerraformResourceParameters(val cdktn.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_DataAzurermLoadTestIdentityList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewDataAzurermLoadTestIdentityListParameters(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}

