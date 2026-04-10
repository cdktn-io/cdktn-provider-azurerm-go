// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

//go:build !no_runtime_type_checking

package manageddevopspool

import (
	"fmt"

	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

func (m *jsiiProxy_ManagedDevopsPoolStatelessAgentManualResourcePredictionOutputReference) validateGetAnyMapAttributeParameters(terraformAttribute *string) error {
	if terraformAttribute == nil {
		return fmt.Errorf("parameter terraformAttribute is required, but nil was provided")
	}

	return nil
}

func (m *jsiiProxy_ManagedDevopsPoolStatelessAgentManualResourcePredictionOutputReference) validateGetBooleanAttributeParameters(terraformAttribute *string) error {
	if terraformAttribute == nil {
		return fmt.Errorf("parameter terraformAttribute is required, but nil was provided")
	}

	return nil
}

func (m *jsiiProxy_ManagedDevopsPoolStatelessAgentManualResourcePredictionOutputReference) validateGetBooleanMapAttributeParameters(terraformAttribute *string) error {
	if terraformAttribute == nil {
		return fmt.Errorf("parameter terraformAttribute is required, but nil was provided")
	}

	return nil
}

func (m *jsiiProxy_ManagedDevopsPoolStatelessAgentManualResourcePredictionOutputReference) validateGetListAttributeParameters(terraformAttribute *string) error {
	if terraformAttribute == nil {
		return fmt.Errorf("parameter terraformAttribute is required, but nil was provided")
	}

	return nil
}

func (m *jsiiProxy_ManagedDevopsPoolStatelessAgentManualResourcePredictionOutputReference) validateGetNumberAttributeParameters(terraformAttribute *string) error {
	if terraformAttribute == nil {
		return fmt.Errorf("parameter terraformAttribute is required, but nil was provided")
	}

	return nil
}

func (m *jsiiProxy_ManagedDevopsPoolStatelessAgentManualResourcePredictionOutputReference) validateGetNumberListAttributeParameters(terraformAttribute *string) error {
	if terraformAttribute == nil {
		return fmt.Errorf("parameter terraformAttribute is required, but nil was provided")
	}

	return nil
}

func (m *jsiiProxy_ManagedDevopsPoolStatelessAgentManualResourcePredictionOutputReference) validateGetNumberMapAttributeParameters(terraformAttribute *string) error {
	if terraformAttribute == nil {
		return fmt.Errorf("parameter terraformAttribute is required, but nil was provided")
	}

	return nil
}

func (m *jsiiProxy_ManagedDevopsPoolStatelessAgentManualResourcePredictionOutputReference) validateGetStringAttributeParameters(terraformAttribute *string) error {
	if terraformAttribute == nil {
		return fmt.Errorf("parameter terraformAttribute is required, but nil was provided")
	}

	return nil
}

func (m *jsiiProxy_ManagedDevopsPoolStatelessAgentManualResourcePredictionOutputReference) validateGetStringMapAttributeParameters(terraformAttribute *string) error {
	if terraformAttribute == nil {
		return fmt.Errorf("parameter terraformAttribute is required, but nil was provided")
	}

	return nil
}

func (m *jsiiProxy_ManagedDevopsPoolStatelessAgentManualResourcePredictionOutputReference) validateInterpolationForAttributeParameters(terraformAttribute *string) error {
	if terraformAttribute == nil {
		return fmt.Errorf("parameter terraformAttribute is required, but nil was provided")
	}

	return nil
}

func (m *jsiiProxy_ManagedDevopsPoolStatelessAgentManualResourcePredictionOutputReference) validatePutFridayScheduleParameters(value interface{}) error {
	if value == nil {
		return fmt.Errorf("parameter value is required, but nil was provided")
	}
	switch value.(type) {
	case cdktn.IResolvable:
		// ok
	case *[]*ManagedDevopsPoolStatelessAgentManualResourcePredictionFridaySchedule:
		value := value.(*[]*ManagedDevopsPoolStatelessAgentManualResourcePredictionFridaySchedule)
		for idx_cd4240, v := range *value {
			if err := _jsii_.ValidateStruct(v, func() string { return fmt.Sprintf("parameter value[%#v]", idx_cd4240) }); err != nil {
				return err
			}
		}
	case []*ManagedDevopsPoolStatelessAgentManualResourcePredictionFridaySchedule:
		value_ := value.([]*ManagedDevopsPoolStatelessAgentManualResourcePredictionFridaySchedule)
		value := &value_
		for idx_cd4240, v := range *value {
			if err := _jsii_.ValidateStruct(v, func() string { return fmt.Sprintf("parameter value[%#v]", idx_cd4240) }); err != nil {
				return err
			}
		}
	default:
		if !_jsii_.IsAnonymousProxy(value) {
			return fmt.Errorf("parameter value must be one of the allowed types: cdktn.IResolvable, *[]*ManagedDevopsPoolStatelessAgentManualResourcePredictionFridaySchedule; received %#v (a %T)", value, value)
		}
	}

	return nil
}

func (m *jsiiProxy_ManagedDevopsPoolStatelessAgentManualResourcePredictionOutputReference) validatePutMondayScheduleParameters(value interface{}) error {
	if value == nil {
		return fmt.Errorf("parameter value is required, but nil was provided")
	}
	switch value.(type) {
	case cdktn.IResolvable:
		// ok
	case *[]*ManagedDevopsPoolStatelessAgentManualResourcePredictionMondaySchedule:
		value := value.(*[]*ManagedDevopsPoolStatelessAgentManualResourcePredictionMondaySchedule)
		for idx_cd4240, v := range *value {
			if err := _jsii_.ValidateStruct(v, func() string { return fmt.Sprintf("parameter value[%#v]", idx_cd4240) }); err != nil {
				return err
			}
		}
	case []*ManagedDevopsPoolStatelessAgentManualResourcePredictionMondaySchedule:
		value_ := value.([]*ManagedDevopsPoolStatelessAgentManualResourcePredictionMondaySchedule)
		value := &value_
		for idx_cd4240, v := range *value {
			if err := _jsii_.ValidateStruct(v, func() string { return fmt.Sprintf("parameter value[%#v]", idx_cd4240) }); err != nil {
				return err
			}
		}
	default:
		if !_jsii_.IsAnonymousProxy(value) {
			return fmt.Errorf("parameter value must be one of the allowed types: cdktn.IResolvable, *[]*ManagedDevopsPoolStatelessAgentManualResourcePredictionMondaySchedule; received %#v (a %T)", value, value)
		}
	}

	return nil
}

func (m *jsiiProxy_ManagedDevopsPoolStatelessAgentManualResourcePredictionOutputReference) validatePutSaturdayScheduleParameters(value interface{}) error {
	if value == nil {
		return fmt.Errorf("parameter value is required, but nil was provided")
	}
	switch value.(type) {
	case cdktn.IResolvable:
		// ok
	case *[]*ManagedDevopsPoolStatelessAgentManualResourcePredictionSaturdaySchedule:
		value := value.(*[]*ManagedDevopsPoolStatelessAgentManualResourcePredictionSaturdaySchedule)
		for idx_cd4240, v := range *value {
			if err := _jsii_.ValidateStruct(v, func() string { return fmt.Sprintf("parameter value[%#v]", idx_cd4240) }); err != nil {
				return err
			}
		}
	case []*ManagedDevopsPoolStatelessAgentManualResourcePredictionSaturdaySchedule:
		value_ := value.([]*ManagedDevopsPoolStatelessAgentManualResourcePredictionSaturdaySchedule)
		value := &value_
		for idx_cd4240, v := range *value {
			if err := _jsii_.ValidateStruct(v, func() string { return fmt.Sprintf("parameter value[%#v]", idx_cd4240) }); err != nil {
				return err
			}
		}
	default:
		if !_jsii_.IsAnonymousProxy(value) {
			return fmt.Errorf("parameter value must be one of the allowed types: cdktn.IResolvable, *[]*ManagedDevopsPoolStatelessAgentManualResourcePredictionSaturdaySchedule; received %#v (a %T)", value, value)
		}
	}

	return nil
}

func (m *jsiiProxy_ManagedDevopsPoolStatelessAgentManualResourcePredictionOutputReference) validatePutSundayScheduleParameters(value interface{}) error {
	if value == nil {
		return fmt.Errorf("parameter value is required, but nil was provided")
	}
	switch value.(type) {
	case cdktn.IResolvable:
		// ok
	case *[]*ManagedDevopsPoolStatelessAgentManualResourcePredictionSundaySchedule:
		value := value.(*[]*ManagedDevopsPoolStatelessAgentManualResourcePredictionSundaySchedule)
		for idx_cd4240, v := range *value {
			if err := _jsii_.ValidateStruct(v, func() string { return fmt.Sprintf("parameter value[%#v]", idx_cd4240) }); err != nil {
				return err
			}
		}
	case []*ManagedDevopsPoolStatelessAgentManualResourcePredictionSundaySchedule:
		value_ := value.([]*ManagedDevopsPoolStatelessAgentManualResourcePredictionSundaySchedule)
		value := &value_
		for idx_cd4240, v := range *value {
			if err := _jsii_.ValidateStruct(v, func() string { return fmt.Sprintf("parameter value[%#v]", idx_cd4240) }); err != nil {
				return err
			}
		}
	default:
		if !_jsii_.IsAnonymousProxy(value) {
			return fmt.Errorf("parameter value must be one of the allowed types: cdktn.IResolvable, *[]*ManagedDevopsPoolStatelessAgentManualResourcePredictionSundaySchedule; received %#v (a %T)", value, value)
		}
	}

	return nil
}

func (m *jsiiProxy_ManagedDevopsPoolStatelessAgentManualResourcePredictionOutputReference) validatePutThursdayScheduleParameters(value interface{}) error {
	if value == nil {
		return fmt.Errorf("parameter value is required, but nil was provided")
	}
	switch value.(type) {
	case cdktn.IResolvable:
		// ok
	case *[]*ManagedDevopsPoolStatelessAgentManualResourcePredictionThursdaySchedule:
		value := value.(*[]*ManagedDevopsPoolStatelessAgentManualResourcePredictionThursdaySchedule)
		for idx_cd4240, v := range *value {
			if err := _jsii_.ValidateStruct(v, func() string { return fmt.Sprintf("parameter value[%#v]", idx_cd4240) }); err != nil {
				return err
			}
		}
	case []*ManagedDevopsPoolStatelessAgentManualResourcePredictionThursdaySchedule:
		value_ := value.([]*ManagedDevopsPoolStatelessAgentManualResourcePredictionThursdaySchedule)
		value := &value_
		for idx_cd4240, v := range *value {
			if err := _jsii_.ValidateStruct(v, func() string { return fmt.Sprintf("parameter value[%#v]", idx_cd4240) }); err != nil {
				return err
			}
		}
	default:
		if !_jsii_.IsAnonymousProxy(value) {
			return fmt.Errorf("parameter value must be one of the allowed types: cdktn.IResolvable, *[]*ManagedDevopsPoolStatelessAgentManualResourcePredictionThursdaySchedule; received %#v (a %T)", value, value)
		}
	}

	return nil
}

func (m *jsiiProxy_ManagedDevopsPoolStatelessAgentManualResourcePredictionOutputReference) validatePutTuesdayScheduleParameters(value interface{}) error {
	if value == nil {
		return fmt.Errorf("parameter value is required, but nil was provided")
	}
	switch value.(type) {
	case cdktn.IResolvable:
		// ok
	case *[]*ManagedDevopsPoolStatelessAgentManualResourcePredictionTuesdaySchedule:
		value := value.(*[]*ManagedDevopsPoolStatelessAgentManualResourcePredictionTuesdaySchedule)
		for idx_cd4240, v := range *value {
			if err := _jsii_.ValidateStruct(v, func() string { return fmt.Sprintf("parameter value[%#v]", idx_cd4240) }); err != nil {
				return err
			}
		}
	case []*ManagedDevopsPoolStatelessAgentManualResourcePredictionTuesdaySchedule:
		value_ := value.([]*ManagedDevopsPoolStatelessAgentManualResourcePredictionTuesdaySchedule)
		value := &value_
		for idx_cd4240, v := range *value {
			if err := _jsii_.ValidateStruct(v, func() string { return fmt.Sprintf("parameter value[%#v]", idx_cd4240) }); err != nil {
				return err
			}
		}
	default:
		if !_jsii_.IsAnonymousProxy(value) {
			return fmt.Errorf("parameter value must be one of the allowed types: cdktn.IResolvable, *[]*ManagedDevopsPoolStatelessAgentManualResourcePredictionTuesdaySchedule; received %#v (a %T)", value, value)
		}
	}

	return nil
}

func (m *jsiiProxy_ManagedDevopsPoolStatelessAgentManualResourcePredictionOutputReference) validatePutWednesdayScheduleParameters(value interface{}) error {
	if value == nil {
		return fmt.Errorf("parameter value is required, but nil was provided")
	}
	switch value.(type) {
	case cdktn.IResolvable:
		// ok
	case *[]*ManagedDevopsPoolStatelessAgentManualResourcePredictionWednesdaySchedule:
		value := value.(*[]*ManagedDevopsPoolStatelessAgentManualResourcePredictionWednesdaySchedule)
		for idx_cd4240, v := range *value {
			if err := _jsii_.ValidateStruct(v, func() string { return fmt.Sprintf("parameter value[%#v]", idx_cd4240) }); err != nil {
				return err
			}
		}
	case []*ManagedDevopsPoolStatelessAgentManualResourcePredictionWednesdaySchedule:
		value_ := value.([]*ManagedDevopsPoolStatelessAgentManualResourcePredictionWednesdaySchedule)
		value := &value_
		for idx_cd4240, v := range *value {
			if err := _jsii_.ValidateStruct(v, func() string { return fmt.Sprintf("parameter value[%#v]", idx_cd4240) }); err != nil {
				return err
			}
		}
	default:
		if !_jsii_.IsAnonymousProxy(value) {
			return fmt.Errorf("parameter value must be one of the allowed types: cdktn.IResolvable, *[]*ManagedDevopsPoolStatelessAgentManualResourcePredictionWednesdaySchedule; received %#v (a %T)", value, value)
		}
	}

	return nil
}

func (m *jsiiProxy_ManagedDevopsPoolStatelessAgentManualResourcePredictionOutputReference) validateResolveParameters(context cdktn.IResolveContext) error {
	if context == nil {
		return fmt.Errorf("parameter context is required, but nil was provided")
	}

	return nil
}

func (j *jsiiProxy_ManagedDevopsPoolStatelessAgentManualResourcePredictionOutputReference) validateSetAllWeekScheduleParameters(val *float64) error {
	if val == nil {
		return fmt.Errorf("parameter val is required, but nil was provided")
	}

	return nil
}

func (j *jsiiProxy_ManagedDevopsPoolStatelessAgentManualResourcePredictionOutputReference) validateSetComplexObjectIndexParameters(val interface{}) error {
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

func (j *jsiiProxy_ManagedDevopsPoolStatelessAgentManualResourcePredictionOutputReference) validateSetComplexObjectIsFromSetParameters(val *bool) error {
	if val == nil {
		return fmt.Errorf("parameter val is required, but nil was provided")
	}

	return nil
}

func (j *jsiiProxy_ManagedDevopsPoolStatelessAgentManualResourcePredictionOutputReference) validateSetInternalValueParameters(val *ManagedDevopsPoolStatelessAgentManualResourcePrediction) error {
	if err := _jsii_.ValidateStruct(val, func() string { return "parameter val" }); err != nil {
		return err
	}

	return nil
}

func (j *jsiiProxy_ManagedDevopsPoolStatelessAgentManualResourcePredictionOutputReference) validateSetTerraformAttributeParameters(val *string) error {
	if val == nil {
		return fmt.Errorf("parameter val is required, but nil was provided")
	}

	return nil
}

func (j *jsiiProxy_ManagedDevopsPoolStatelessAgentManualResourcePredictionOutputReference) validateSetTerraformResourceParameters(val cdktn.IInterpolatingParent) error {
	if val == nil {
		return fmt.Errorf("parameter val is required, but nil was provided")
	}

	return nil
}

func (j *jsiiProxy_ManagedDevopsPoolStatelessAgentManualResourcePredictionOutputReference) validateSetTimeZoneNameParameters(val *string) error {
	if val == nil {
		return fmt.Errorf("parameter val is required, but nil was provided")
	}

	return nil
}

func validateNewManagedDevopsPoolStatelessAgentManualResourcePredictionOutputReferenceParameters(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) error {
	if terraformResource == nil {
		return fmt.Errorf("parameter terraformResource is required, but nil was provided")
	}

	if terraformAttribute == nil {
		return fmt.Errorf("parameter terraformAttribute is required, but nil was provided")
	}

	return nil
}

