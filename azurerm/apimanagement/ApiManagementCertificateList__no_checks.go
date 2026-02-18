// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

//go:build no_runtime_type_checking

package apimanagement

// Building without runtime type checking enabled, so all the below just return nil

func (a *jsiiProxy_ApiManagementCertificateList) validateAllWithMapKeyParameters(mapKeyAttributeName *string) error {
	return nil
}

func (a *jsiiProxy_ApiManagementCertificateList) validateGetParameters(index *float64) error {
	return nil
}

func (a *jsiiProxy_ApiManagementCertificateList) validateResolveParameters(context cdktn.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_ApiManagementCertificateList) validateSetInternalValueParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_ApiManagementCertificateList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_ApiManagementCertificateList) validateSetTerraformResourceParameters(val cdktn.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_ApiManagementCertificateList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewApiManagementCertificateListParameters(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}

