// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

//go:build no_runtime_type_checking

package frontdoorrulesengine

// Building without runtime type checking enabled, so all the below just return nil

func (f *jsiiProxy_FrontdoorRulesEngineRuleList) validateAllWithMapKeyParameters(mapKeyAttributeName *string) error {
	return nil
}

func (f *jsiiProxy_FrontdoorRulesEngineRuleList) validateGetParameters(index *float64) error {
	return nil
}

func (f *jsiiProxy_FrontdoorRulesEngineRuleList) validateResolveParameters(context cdktn.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_FrontdoorRulesEngineRuleList) validateSetInternalValueParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_FrontdoorRulesEngineRuleList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_FrontdoorRulesEngineRuleList) validateSetTerraformResourceParameters(val cdktn.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_FrontdoorRulesEngineRuleList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewFrontdoorRulesEngineRuleListParameters(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}

