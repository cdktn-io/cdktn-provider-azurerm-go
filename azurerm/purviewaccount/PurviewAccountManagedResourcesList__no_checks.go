// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

//go:build no_runtime_type_checking

package purviewaccount

// Building without runtime type checking enabled, so all the below just return nil

func (p *jsiiProxy_PurviewAccountManagedResourcesList) validateAllWithMapKeyParameters(mapKeyAttributeName *string) error {
	return nil
}

func (p *jsiiProxy_PurviewAccountManagedResourcesList) validateGetParameters(index *float64) error {
	return nil
}

func (p *jsiiProxy_PurviewAccountManagedResourcesList) validateResolveParameters(context cdktn.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_PurviewAccountManagedResourcesList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_PurviewAccountManagedResourcesList) validateSetTerraformResourceParameters(val cdktn.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_PurviewAccountManagedResourcesList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewPurviewAccountManagedResourcesListParameters(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}

