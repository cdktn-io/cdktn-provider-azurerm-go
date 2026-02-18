// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

//go:build no_runtime_type_checking

package containerapp

// Building without runtime type checking enabled, so all the below just return nil

func (c *jsiiProxy_ContainerAppTemplateContainerList) validateAllWithMapKeyParameters(mapKeyAttributeName *string) error {
	return nil
}

func (c *jsiiProxy_ContainerAppTemplateContainerList) validateGetParameters(index *float64) error {
	return nil
}

func (c *jsiiProxy_ContainerAppTemplateContainerList) validateResolveParameters(context cdktn.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_ContainerAppTemplateContainerList) validateSetInternalValueParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_ContainerAppTemplateContainerList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_ContainerAppTemplateContainerList) validateSetTerraformResourceParameters(val cdktn.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_ContainerAppTemplateContainerList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewContainerAppTemplateContainerListParameters(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}

