// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

//go:build no_runtime_type_checking

package cognitiveaccount

// Building without runtime type checking enabled, so all the below just return nil

func (c *jsiiProxy_CognitiveAccountStorageList) validateAllWithMapKeyParameters(mapKeyAttributeName *string) error {
	return nil
}

func (c *jsiiProxy_CognitiveAccountStorageList) validateGetParameters(index *float64) error {
	return nil
}

func (c *jsiiProxy_CognitiveAccountStorageList) validateResolveParameters(context cdktn.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_CognitiveAccountStorageList) validateSetInternalValueParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_CognitiveAccountStorageList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_CognitiveAccountStorageList) validateSetTerraformResourceParameters(val cdktn.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_CognitiveAccountStorageList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewCognitiveAccountStorageListParameters(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}

