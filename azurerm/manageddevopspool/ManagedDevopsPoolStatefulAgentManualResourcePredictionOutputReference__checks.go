// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

//go:build !no_runtime_type_checking

package manageddevopspool

import (
	"fmt"

	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

func (m *jsiiProxy_ManagedDevopsPoolStatefulAgentManualResourcePredictionOutputReference) validateGetAnyMapAttributeParameters(terraformAttribute *string) error {
	if terraformAttribute == nil {
		return fmt.Errorf("parameter terraformAttribute is required, but nil was provided")
	}

	return nil
}

func (m *jsiiProxy_ManagedDevopsPoolStatefulAgentManualResourcePredictionOutputReference) validateGetBooleanAttributeParameters(terraformAttribute *string) error {
	if terraformAttribute == nil {
		return fmt.Errorf("parameter terraformAttribute is required, but nil was provided")
	}

	return nil
}

func (m *jsiiProxy_ManagedDevopsPoolStatefulAgentManualResourcePredictionOutputReference) validateGetBooleanMapAttributeParameters(terraformAttribute *string) error {
	if terraformAttribute == nil {
		return fmt.Errorf("parameter terraformAttribute is required, but nil was provided")
	}

	return nil
}

func (m *jsiiProxy_ManagedDevopsPoolStatefulAgentManualResourcePredictionOutputReference) validateGetListAttributeParameters(terraformAttribute *string) error {
	if terraformAttribute == nil {
		return fmt.Errorf("parameter terraformAttribute is required, but nil was provided")
	}

	return nil
}

func (m *jsiiProxy_ManagedDevopsPoolStatefulAgentManualResourcePredictionOutputReference) validateGetNumberAttributeParameters(terraformAttribute *string) error {
	if terraformAttribute == nil {
		return fmt.Errorf("parameter terraformAttribute is required, but nil was provided")
	}

	return nil
}

func (m *jsiiProxy_ManagedDevopsPoolStatefulAgentManualResourcePredictionOutputReference) validateGetNumberListAttributeParameters(terraformAttribute *string) error {
	if terraformAttribute == nil {
		return fmt.Errorf("parameter terraformAttribute is required, but nil was provided")
	}

	return nil
}

func (m *jsiiProxy_ManagedDevopsPoolStatefulAgentManualResourcePredictionOutputReference) validateGetNumberMapAttributeParameters(terraformAttribute *string) error {
	if terraformAttribute == nil {
		return fmt.Errorf("parameter terraformAttribute is required, but nil was provided")
	}

	return nil
}

func (m *jsiiProxy_ManagedDevopsPoolStatefulAgentManualResourcePredictionOutputReference) validateGetStringAttributeParameters(terraformAttribute *string) error {
	if terraformAttribute == nil {
		return fmt.Errorf("parameter terraformAttribute is required, but nil was provided")
	}

	return nil
}

func (m *jsiiProxy_ManagedDevopsPoolStatefulAgentManualResourcePredictionOutputReference) validateGetStringMapAttributeParameters(terraformAttribute *string) error {
	if terraformAttribute == nil {
		return fmt.Errorf("parameter terraformAttribute is required, but nil was provided")
	}

	return nil
}

func (m *jsiiProxy_ManagedDevopsPoolStatefulAgentManualResourcePredictionOutputReference) validateInterpolationForAttributeParameters(terraformAttribute *string) error {
	if terraformAttribute == nil {
		return fmt.Errorf("parameter terraformAttribute is required, but nil was provided")
	}

	return nil
}

func (m *jsiiProxy_ManagedDevopsPoolStatefulAgentManualResourcePredictionOutputReference) validatePutFridayScheduleParameters(value interface{}) error {
	if value == nil {
		return fmt.Errorf("parameter value is required, but nil was provided")
	}
	switch value.(type) {
	case cdktn.IResolvable:
		// ok
	case *[]*ManagedDevopsPoolStatefulAgentManualResourcePredictionFridaySchedule:
		value := value.(*[]*ManagedDevopsPoolStatefulAgentManualResourcePredictionFridaySchedule)
		for idx_cd4240, v := range *value {
			if err := _jsii_.ValidateStruct(v, func() string { return fmt.Sprintf("parameter value[%#v]", idx_cd4240) }); err != nil {
				return err
			}
		}
	case []*ManagedDevopsPoolStatefulAgentManualResourcePredictionFridaySchedule:
		value_ := value.([]*ManagedDevopsPoolStatefulAgentManualResourcePredictionFridaySchedule)
		value := &value_
		for idx_cd4240, v := range *value {
			if err := _jsii_.ValidateStruct(v, func() string { return fmt.Sprintf("parameter value[%#v]", idx_cd4240) }); err != nil {
				return err
			}
		}
	default:
		if !_jsii_.IsAnonymousProxy(value) {
			return fmt.Errorf("parameter value must be one of the allowed types: cdktn.IResolvable, *[]*ManagedDevopsPoolStatefulAgentManualResourcePredictionFridaySchedule; received %#v (a %T)", value, value)
		}
	}

	return nil
}

func (m *jsiiProxy_ManagedDevopsPoolStatefulAgentManualResourcePredictionOutputReference) validatePutMondayScheduleParameters(value interface{}) error {
	if value == nil {
		return fmt.Errorf("parameter value is required, but nil was provided")
	}
	switch value.(type) {
	case cdktn.IResolvable:
		// ok
	case *[]*ManagedDevopsPoolStatefulAgentManualResourcePredictionMondaySchedule:
		value := value.(*[]*ManagedDevopsPoolStatefulAgentManualResourcePredictionMondaySchedule)
		for idx_cd4240, v := range *value {
			if err := _jsii_.ValidateStruct(v, func() string { return fmt.Sprintf("parameter value[%#v]", idx_cd4240) }); err != nil {
				return err
			}
		}
	case []*ManagedDevopsPoolStatefulAgentManualResourcePredictionMondaySchedule:
		value_ := value.([]*ManagedDevopsPoolStatefulAgentManualResourcePredictionMondaySchedule)
		value := &value_
		for idx_cd4240, v := range *value {
			if err := _jsii_.ValidateStruct(v, func() string { return fmt.Sprintf("parameter value[%#v]", idx_cd4240) }); err != nil {
				return err
			}
		}
	default:
		if !_jsii_.IsAnonymousProxy(value) {
			return fmt.Errorf("parameter value must be one of the allowed types: cdktn.IResolvable, *[]*ManagedDevopsPoolStatefulAgentManualResourcePredictionMondaySchedule; received %#v (a %T)", value, value)
		}
	}

	return nil
}

func (m *jsiiProxy_ManagedDevopsPoolStatefulAgentManualResourcePredictionOutputReference) validatePutSaturdayScheduleParameters(value interface{}) error {
	if value == nil {
		return fmt.Errorf("parameter value is required, but nil was provided")
	}
	switch value.(type) {
	case cdktn.IResolvable:
		// ok
	case *[]*ManagedDevopsPoolStatefulAgentManualResourcePredictionSaturdaySchedule:
		value := value.(*[]*ManagedDevopsPoolStatefulAgentManualResourcePredictionSaturdaySchedule)
		for idx_cd4240, v := range *value {
			if err := _jsii_.ValidateStruct(v, func() string { return fmt.Sprintf("parameter value[%#v]", idx_cd4240) }); err != nil {
				return err
			}
		}
	case []*ManagedDevopsPoolStatefulAgentManualResourcePredictionSaturdaySchedule:
		value_ := value.([]*ManagedDevopsPoolStatefulAgentManualResourcePredictionSaturdaySchedule)
		value := &value_
		for idx_cd4240, v := range *value {
			if err := _jsii_.ValidateStruct(v, func() string { return fmt.Sprintf("parameter value[%#v]", idx_cd4240) }); err != nil {
				return err
			}
		}
	default:
		if !_jsii_.IsAnonymousProxy(value) {
			return fmt.Errorf("parameter value must be one of the allowed types: cdktn.IResolvable, *[]*ManagedDevopsPoolStatefulAgentManualResourcePredictionSaturdaySchedule; received %#v (a %T)", value, value)
		}
	}

	return nil
}

func (m *jsiiProxy_ManagedDevopsPoolStatefulAgentManualResourcePredictionOutputReference) validatePutSundayScheduleParameters(value interface{}) error {
	if value == nil {
		return fmt.Errorf("parameter value is required, but nil was provided")
	}
	switch value.(type) {
	case cdktn.IResolvable:
		// ok
	case *[]*ManagedDevopsPoolStatefulAgentManualResourcePredictionSundaySchedule:
		value := value.(*[]*ManagedDevopsPoolStatefulAgentManualResourcePredictionSundaySchedule)
		for idx_cd4240, v := range *value {
			if err := _jsii_.ValidateStruct(v, func() string { return fmt.Sprintf("parameter value[%#v]", idx_cd4240) }); err != nil {
				return err
			}
		}
	case []*ManagedDevopsPoolStatefulAgentManualResourcePredictionSundaySchedule:
		value_ := value.([]*ManagedDevopsPoolStatefulAgentManualResourcePredictionSundaySchedule)
		value := &value_
		for idx_cd4240, v := range *value {
			if err := _jsii_.ValidateStruct(v, func() string { return fmt.Sprintf("parameter value[%#v]", idx_cd4240) }); err != nil {
				return err
			}
		}
	default:
		if !_jsii_.IsAnonymousProxy(value) {
			return fmt.Errorf("parameter value must be one of the allowed types: cdktn.IResolvable, *[]*ManagedDevopsPoolStatefulAgentManualResourcePredictionSundaySchedule; received %#v (a %T)", value, value)
		}
	}

	return nil
}

func (m *jsiiProxy_ManagedDevopsPoolStatefulAgentManualResourcePredictionOutputReference) validatePutThursdayScheduleParameters(value interface{}) error {
	if value == nil {
		return fmt.Errorf("parameter value is required, but nil was provided")
	}
	switch value.(type) {
	case cdktn.IResolvable:
		// ok
	case *[]*ManagedDevopsPoolStatefulAgentManualResourcePredictionThursdaySchedule:
		value := value.(*[]*ManagedDevopsPoolStatefulAgentManualResourcePredictionThursdaySchedule)
		for idx_cd4240, v := range *value {
			if err := _jsii_.ValidateStruct(v, func() string { return fmt.Sprintf("parameter value[%#v]", idx_cd4240) }); err != nil {
				return err
			}
		}
	case []*ManagedDevopsPoolStatefulAgentManualResourcePredictionThursdaySchedule:
		value_ := value.([]*ManagedDevopsPoolStatefulAgentManualResourcePredictionThursdaySchedule)
		value := &value_
		for idx_cd4240, v := range *value {
			if err := _jsii_.ValidateStruct(v, func() string { return fmt.Sprintf("parameter value[%#v]", idx_cd4240) }); err != nil {
				return err
			}
		}
	default:
		if !_jsii_.IsAnonymousProxy(value) {
			return fmt.Errorf("parameter value must be one of the allowed types: cdktn.IResolvable, *[]*ManagedDevopsPoolStatefulAgentManualResourcePredictionThursdaySchedule; received %#v (a %T)", value, value)
		}
	}

	return nil
}

func (m *jsiiProxy_ManagedDevopsPoolStatefulAgentManualResourcePredictionOutputReference) validatePutTuesdayScheduleParameters(value interface{}) error {
	if value == nil {
		return fmt.Errorf("parameter value is required, but nil was provided")
	}
	switch value.(type) {
	case cdktn.IResolvable:
		// ok
	case *[]*ManagedDevopsPoolStatefulAgentManualResourcePredictionTuesdaySchedule:
		value := value.(*[]*ManagedDevopsPoolStatefulAgentManualResourcePredictionTuesdaySchedule)
		for idx_cd4240, v := range *value {
			if err := _jsii_.ValidateStruct(v, func() string { return fmt.Sprintf("parameter value[%#v]", idx_cd4240) }); err != nil {
				return err
			}
		}
	case []*ManagedDevopsPoolStatefulAgentManualResourcePredictionTuesdaySchedule:
		value_ := value.([]*ManagedDevopsPoolStatefulAgentManualResourcePredictionTuesdaySchedule)
		value := &value_
		for idx_cd4240, v := range *value {
			if err := _jsii_.ValidateStruct(v, func() string { return fmt.Sprintf("parameter value[%#v]", idx_cd4240) }); err != nil {
				return err
			}
		}
	default:
		if !_jsii_.IsAnonymousProxy(value) {
			return fmt.Errorf("parameter value must be one of the allowed types: cdktn.IResolvable, *[]*ManagedDevopsPoolStatefulAgentManualResourcePredictionTuesdaySchedule; received %#v (a %T)", value, value)
		}
	}

	return nil
}

func (m *jsiiProxy_ManagedDevopsPoolStatefulAgentManualResourcePredictionOutputReference) validatePutWednesdayScheduleParameters(value interface{}) error {
	if value == nil {
		return fmt.Errorf("parameter value is required, but nil was provided")
	}
	switch value.(type) {
	case cdktn.IResolvable:
		// ok
	case *[]*ManagedDevopsPoolStatefulAgentManualResourcePredictionWednesdaySchedule:
		value := value.(*[]*ManagedDevopsPoolStatefulAgentManualResourcePredictionWednesdaySchedule)
		for idx_cd4240, v := range *value {
			if err := _jsii_.ValidateStruct(v, func() string { return fmt.Sprintf("parameter value[%#v]", idx_cd4240) }); err != nil {
				return err
			}
		}
	case []*ManagedDevopsPoolStatefulAgentManualResourcePredictionWednesdaySchedule:
		value_ := value.([]*ManagedDevopsPoolStatefulAgentManualResourcePredictionWednesdaySchedule)
		value := &value_
		for idx_cd4240, v := range *value {
			if err := _jsii_.ValidateStruct(v, func() string { return fmt.Sprintf("parameter value[%#v]", idx_cd4240) }); err != nil {
				return err
			}
		}
	default:
		if !_jsii_.IsAnonymousProxy(value) {
			return fmt.Errorf("parameter value must be one of the allowed types: cdktn.IResolvable, *[]*ManagedDevopsPoolStatefulAgentManualResourcePredictionWednesdaySchedule; received %#v (a %T)", value, value)
		}
	}

	return nil
}

func (m *jsiiProxy_ManagedDevopsPoolStatefulAgentManualResourcePredictionOutputReference) validateResolveParameters(context cdktn.IResolveContext) error {
	if context == nil {
		return fmt.Errorf("parameter context is required, but nil was provided")
	}

	return nil
}

func (j *jsiiProxy_ManagedDevopsPoolStatefulAgentManualResourcePredictionOutputReference) validateSetAllWeekScheduleParameters(val *float64) error {
	if val == nil {
		return fmt.Errorf("parameter val is required, but nil was provided")
	}

	return nil
}

func (j *jsiiProxy_ManagedDevopsPoolStatefulAgentManualResourcePredictionOutputReference) validateSetComplexObjectIndexParameters(val interface{}) error {
	switch val.(type) {
	case *string:
		// ok
	case string:
		// ok
	case *float64:
		// ok
	case float64:
		// ok
	case *int:
		// ok
	case int:
		// ok
	case *uint:
		// ok
	case uint:
		// ok
	case *int8:
		// ok
	case int8:
		// ok
	case *int16:
		// ok
	case int16:
		// ok
	case *int32:
		// ok
	case int32:
		// ok
	case *int64:
		// ok
	case int64:
		// ok
	case *uint8:
		// ok
	case uint8:
		// ok
	case *uint16:
		// ok
	case uint16:
		// ok
	case *uint32:
		// ok
	case uint32:
		// ok
	case *uint64:
		// ok
	case uint64:
		// ok
	default:
		return fmt.Errorf("parameter val must be one of the allowed types: *string, *float64; received %#v (a %T)", val, val)
	}

	return nil
}

func (j *jsiiProxy_ManagedDevopsPoolStatefulAgentManualResourcePredictionOutputReference) validateSetComplexObjectIsFromSetParameters(val *bool) error {
	if val == nil {
		return fmt.Errorf("parameter val is required, but nil was provided")
	}

	return nil
}

func (j *jsiiProxy_ManagedDevopsPoolStatefulAgentManualResourcePredictionOutputReference) validateSetInternalValueParameters(val *ManagedDevopsPoolStatefulAgentManualResourcePrediction) error {
	if err := _jsii_.ValidateStruct(val, func() string { return "parameter val" }); err != nil {
		return err
	}

	return nil
}

func (j *jsiiProxy_ManagedDevopsPoolStatefulAgentManualResourcePredictionOutputReference) validateSetTerraformAttributeParameters(val *string) error {
	if val == nil {
		return fmt.Errorf("parameter val is required, but nil was provided")
	}

	return nil
}

func (j *jsiiProxy_ManagedDevopsPoolStatefulAgentManualResourcePredictionOutputReference) validateSetTerraformResourceParameters(val cdktn.IInterpolatingParent) error {
	if val == nil {
		return fmt.Errorf("parameter val is required, but nil was provided")
	}

	return nil
}

func (j *jsiiProxy_ManagedDevopsPoolStatefulAgentManualResourcePredictionOutputReference) validateSetTimeZoneNameParameters(val *string) error {
	if val == nil {
		return fmt.Errorf("parameter val is required, but nil was provided")
	}

	return nil
}

func validateNewManagedDevopsPoolStatefulAgentManualResourcePredictionOutputReferenceParameters(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) error {
	if terraformResource == nil {
		return fmt.Errorf("parameter terraformResource is required, but nil was provided")
	}

	if terraformAttribute == nil {
		return fmt.Errorf("parameter terraformAttribute is required, but nil was provided")
	}

	return nil
}

