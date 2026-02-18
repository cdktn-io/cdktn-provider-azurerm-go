// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

//go:build no_runtime_type_checking

package monitormetricalert

// Building without runtime type checking enabled, so all the below just return nil

func (m *jsiiProxy_MonitorMetricAlertCriteriaList) validateAllWithMapKeyParameters(mapKeyAttributeName *string) error {
	return nil
}

func (m *jsiiProxy_MonitorMetricAlertCriteriaList) validateGetParameters(index *float64) error {
	return nil
}

func (m *jsiiProxy_MonitorMetricAlertCriteriaList) validateResolveParameters(context cdktn.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_MonitorMetricAlertCriteriaList) validateSetInternalValueParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_MonitorMetricAlertCriteriaList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_MonitorMetricAlertCriteriaList) validateSetTerraformResourceParameters(val cdktn.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_MonitorMetricAlertCriteriaList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewMonitorMetricAlertCriteriaListParameters(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}

