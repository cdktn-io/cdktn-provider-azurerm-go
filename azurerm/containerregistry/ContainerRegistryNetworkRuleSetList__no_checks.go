// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

//go:build no_runtime_type_checking

package containerregistry

// Building without runtime type checking enabled, so all the below just return nil

func (c *jsiiProxy_ContainerRegistryNetworkRuleSetList) validateAllWithMapKeyParameters(mapKeyAttributeName *string) error {
	return nil
}

func (c *jsiiProxy_ContainerRegistryNetworkRuleSetList) validateGetParameters(index *float64) error {
	return nil
}

func (c *jsiiProxy_ContainerRegistryNetworkRuleSetList) validateResolveParameters(context cdktn.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_ContainerRegistryNetworkRuleSetList) validateSetInternalValueParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_ContainerRegistryNetworkRuleSetList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_ContainerRegistryNetworkRuleSetList) validateSetTerraformResourceParameters(val cdktn.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_ContainerRegistryNetworkRuleSetList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewContainerRegistryNetworkRuleSetListParameters(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}

