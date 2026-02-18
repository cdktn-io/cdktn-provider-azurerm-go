// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

//go:build no_runtime_type_checking

package appservice

// Building without runtime type checking enabled, so all the below just return nil

func (a *jsiiProxy_AppServiceStorageAccountList) validateAllWithMapKeyParameters(mapKeyAttributeName *string) error {
	return nil
}

func (a *jsiiProxy_AppServiceStorageAccountList) validateGetParameters(index *float64) error {
	return nil
}

func (a *jsiiProxy_AppServiceStorageAccountList) validateResolveParameters(context cdktn.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_AppServiceStorageAccountList) validateSetInternalValueParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_AppServiceStorageAccountList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_AppServiceStorageAccountList) validateSetTerraformResourceParameters(val cdktn.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_AppServiceStorageAccountList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewAppServiceStorageAccountListParameters(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}

