// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

//go:build no_runtime_type_checking

package iothub

// Building without runtime type checking enabled, so all the below just return nil

func (i *jsiiProxy_IothubNetworkRuleSetList) validateAllWithMapKeyParameters(mapKeyAttributeName *string) error {
	return nil
}

func (i *jsiiProxy_IothubNetworkRuleSetList) validateGetParameters(index *float64) error {
	return nil
}

func (i *jsiiProxy_IothubNetworkRuleSetList) validateResolveParameters(context cdktn.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_IothubNetworkRuleSetList) validateSetInternalValueParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_IothubNetworkRuleSetList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_IothubNetworkRuleSetList) validateSetTerraformResourceParameters(val cdktn.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_IothubNetworkRuleSetList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewIothubNetworkRuleSetListParameters(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}

