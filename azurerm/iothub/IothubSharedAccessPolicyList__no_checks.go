// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

//go:build no_runtime_type_checking

package iothub

// Building without runtime type checking enabled, so all the below just return nil

func (i *jsiiProxy_IothubSharedAccessPolicyList) validateAllWithMapKeyParameters(mapKeyAttributeName *string) error {
	return nil
}

func (i *jsiiProxy_IothubSharedAccessPolicyList) validateGetParameters(index *float64) error {
	return nil
}

func (i *jsiiProxy_IothubSharedAccessPolicyList) validateResolveParameters(context cdktn.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_IothubSharedAccessPolicyList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_IothubSharedAccessPolicyList) validateSetTerraformResourceParameters(val cdktn.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_IothubSharedAccessPolicyList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewIothubSharedAccessPolicyListParameters(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}

