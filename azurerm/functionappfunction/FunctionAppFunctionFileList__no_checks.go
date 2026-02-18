// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

//go:build no_runtime_type_checking

package functionappfunction

// Building without runtime type checking enabled, so all the below just return nil

func (f *jsiiProxy_FunctionAppFunctionFileList) validateAllWithMapKeyParameters(mapKeyAttributeName *string) error {
	return nil
}

func (f *jsiiProxy_FunctionAppFunctionFileList) validateGetParameters(index *float64) error {
	return nil
}

func (f *jsiiProxy_FunctionAppFunctionFileList) validateResolveParameters(context cdktn.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_FunctionAppFunctionFileList) validateSetInternalValueParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_FunctionAppFunctionFileList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_FunctionAppFunctionFileList) validateSetTerraformResourceParameters(val cdktn.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_FunctionAppFunctionFileList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewFunctionAppFunctionFileListParameters(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}

