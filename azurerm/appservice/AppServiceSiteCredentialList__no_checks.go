// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

//go:build no_runtime_type_checking

package appservice

// Building without runtime type checking enabled, so all the below just return nil

func (a *jsiiProxy_AppServiceSiteCredentialList) validateAllWithMapKeyParameters(mapKeyAttributeName *string) error {
	return nil
}

func (a *jsiiProxy_AppServiceSiteCredentialList) validateGetParameters(index *float64) error {
	return nil
}

func (a *jsiiProxy_AppServiceSiteCredentialList) validateResolveParameters(context cdktn.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_AppServiceSiteCredentialList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_AppServiceSiteCredentialList) validateSetTerraformResourceParameters(val cdktn.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_AppServiceSiteCredentialList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewAppServiceSiteCredentialListParameters(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}

