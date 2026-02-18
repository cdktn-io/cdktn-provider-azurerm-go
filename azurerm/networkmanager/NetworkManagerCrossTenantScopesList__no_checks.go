// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

//go:build no_runtime_type_checking

package networkmanager

// Building without runtime type checking enabled, so all the below just return nil

func (n *jsiiProxy_NetworkManagerCrossTenantScopesList) validateAllWithMapKeyParameters(mapKeyAttributeName *string) error {
	return nil
}

func (n *jsiiProxy_NetworkManagerCrossTenantScopesList) validateGetParameters(index *float64) error {
	return nil
}

func (n *jsiiProxy_NetworkManagerCrossTenantScopesList) validateResolveParameters(context cdktn.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_NetworkManagerCrossTenantScopesList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_NetworkManagerCrossTenantScopesList) validateSetTerraformResourceParameters(val cdktn.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_NetworkManagerCrossTenantScopesList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewNetworkManagerCrossTenantScopesListParameters(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}

