// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

//go:build no_runtime_type_checking

package frontdoor

// Building without runtime type checking enabled, so all the below just return nil

func (f *jsiiProxy_FrontdoorBackendPoolList) validateAllWithMapKeyParameters(mapKeyAttributeName *string) error {
	return nil
}

func (f *jsiiProxy_FrontdoorBackendPoolList) validateGetParameters(index *float64) error {
	return nil
}

func (f *jsiiProxy_FrontdoorBackendPoolList) validateResolveParameters(context cdktn.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_FrontdoorBackendPoolList) validateSetInternalValueParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_FrontdoorBackendPoolList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_FrontdoorBackendPoolList) validateSetTerraformResourceParameters(val cdktn.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_FrontdoorBackendPoolList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewFrontdoorBackendPoolListParameters(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}

