// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

//go:build no_runtime_type_checking

package automationaccount

// Building without runtime type checking enabled, so all the below just return nil

func (a *jsiiProxy_AutomationAccountEncryptionList) validateAllWithMapKeyParameters(mapKeyAttributeName *string) error {
	return nil
}

func (a *jsiiProxy_AutomationAccountEncryptionList) validateGetParameters(index *float64) error {
	return nil
}

func (a *jsiiProxy_AutomationAccountEncryptionList) validateResolveParameters(context cdktn.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_AutomationAccountEncryptionList) validateSetInternalValueParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_AutomationAccountEncryptionList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_AutomationAccountEncryptionList) validateSetTerraformResourceParameters(val cdktn.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_AutomationAccountEncryptionList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewAutomationAccountEncryptionListParameters(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}

