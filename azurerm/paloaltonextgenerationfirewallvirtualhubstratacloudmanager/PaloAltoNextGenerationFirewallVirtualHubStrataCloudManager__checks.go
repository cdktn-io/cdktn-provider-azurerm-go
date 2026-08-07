// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

//go:build !no_runtime_type_checking

package paloaltonextgenerationfirewallvirtualhubstratacloudmanager

import (
	"fmt"

	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/constructs-go/constructs/v10"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

func (p *jsiiProxy_PaloAltoNextGenerationFirewallVirtualHubStrataCloudManager) validateAddMoveTargetParameters(moveTarget *string) error {
	if moveTarget == nil {
		return fmt.Errorf("parameter moveTarget is required, but nil was provided")
	}

	return nil
}

func (p *jsiiProxy_PaloAltoNextGenerationFirewallVirtualHubStrataCloudManager) validateAddOverrideParameters(path *string, value interface{}) error {
	if path == nil {
		return fmt.Errorf("parameter path is required, but nil was provided")
	}

	if value == nil {
		return fmt.Errorf("parameter value is required, but nil was provided")
	}

	return nil
}

func (p *jsiiProxy_PaloAltoNextGenerationFirewallVirtualHubStrataCloudManager) validateGetAnyMapAttributeParameters(terraformAttribute *string) error {
	if terraformAttribute == nil {
		return fmt.Errorf("parameter terraformAttribute is required, but nil was provided")
	}

	return nil
}

func (p *jsiiProxy_PaloAltoNextGenerationFirewallVirtualHubStrataCloudManager) validateGetBooleanAttributeParameters(terraformAttribute *string) error {
	if terraformAttribute == nil {
		return fmt.Errorf("parameter terraformAttribute is required, but nil was provided")
	}

	return nil
}

func (p *jsiiProxy_PaloAltoNextGenerationFirewallVirtualHubStrataCloudManager) validateGetBooleanMapAttributeParameters(terraformAttribute *string) error {
	if terraformAttribute == nil {
		return fmt.Errorf("parameter terraformAttribute is required, but nil was provided")
	}

	return nil
}

func (p *jsiiProxy_PaloAltoNextGenerationFirewallVirtualHubStrataCloudManager) validateGetListAttributeParameters(terraformAttribute *string) error {
	if terraformAttribute == nil {
		return fmt.Errorf("parameter terraformAttribute is required, but nil was provided")
	}

	return nil
}

func (p *jsiiProxy_PaloAltoNextGenerationFirewallVirtualHubStrataCloudManager) validateGetNumberAttributeParameters(terraformAttribute *string) error {
	if terraformAttribute == nil {
		return fmt.Errorf("parameter terraformAttribute is required, but nil was provided")
	}

	return nil
}

func (p *jsiiProxy_PaloAltoNextGenerationFirewallVirtualHubStrataCloudManager) validateGetNumberListAttributeParameters(terraformAttribute *string) error {
	if terraformAttribute == nil {
		return fmt.Errorf("parameter terraformAttribute is required, but nil was provided")
	}

	return nil
}

func (p *jsiiProxy_PaloAltoNextGenerationFirewallVirtualHubStrataCloudManager) validateGetNumberMapAttributeParameters(terraformAttribute *string) error {
	if terraformAttribute == nil {
		return fmt.Errorf("parameter terraformAttribute is required, but nil was provided")
	}

	return nil
}

func (p *jsiiProxy_PaloAltoNextGenerationFirewallVirtualHubStrataCloudManager) validateGetStringAttributeParameters(terraformAttribute *string) error {
	if terraformAttribute == nil {
		return fmt.Errorf("parameter terraformAttribute is required, but nil was provided")
	}

	return nil
}

func (p *jsiiProxy_PaloAltoNextGenerationFirewallVirtualHubStrataCloudManager) validateGetStringMapAttributeParameters(terraformAttribute *string) error {
	if terraformAttribute == nil {
		return fmt.Errorf("parameter terraformAttribute is required, but nil was provided")
	}

	return nil
}

func (p *jsiiProxy_PaloAltoNextGenerationFirewallVirtualHubStrataCloudManager) validateImportFromParameters(id *string) error {
	if id == nil {
		return fmt.Errorf("parameter id is required, but nil was provided")
	}

	return nil
}

func (p *jsiiProxy_PaloAltoNextGenerationFirewallVirtualHubStrataCloudManager) validateInterpolationForAttributeParameters(terraformAttribute *string) error {
	if terraformAttribute == nil {
		return fmt.Errorf("parameter terraformAttribute is required, but nil was provided")
	}

	return nil
}

func (p *jsiiProxy_PaloAltoNextGenerationFirewallVirtualHubStrataCloudManager) validateMarkWriteOnlyAttributeParameters(value interface{}) error {
	if value == nil {
		return fmt.Errorf("parameter value is required, but nil was provided")
	}

	return nil
}

func (p *jsiiProxy_PaloAltoNextGenerationFirewallVirtualHubStrataCloudManager) validateMoveFromIdParameters(id *string) error {
	if id == nil {
		return fmt.Errorf("parameter id is required, but nil was provided")
	}

	return nil
}

func (p *jsiiProxy_PaloAltoNextGenerationFirewallVirtualHubStrataCloudManager) validateMoveToParameters(moveTarget *string, index interface{}) error {
	if moveTarget == nil {
		return fmt.Errorf("parameter moveTarget is required, but nil was provided")
	}

	switch index.(type) {
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
		return fmt.Errorf("parameter index must be one of the allowed types: *string, *float64; received %#v (a %T)", index, index)
	}

	return nil
}

func (p *jsiiProxy_PaloAltoNextGenerationFirewallVirtualHubStrataCloudManager) validateMoveToIdParameters(id *string) error {
	if id == nil {
		return fmt.Errorf("parameter id is required, but nil was provided")
	}

	return nil
}

func (p *jsiiProxy_PaloAltoNextGenerationFirewallVirtualHubStrataCloudManager) validateOverrideLogicalIdParameters(newLogicalId *string) error {
	if newLogicalId == nil {
		return fmt.Errorf("parameter newLogicalId is required, but nil was provided")
	}

	return nil
}

func (p *jsiiProxy_PaloAltoNextGenerationFirewallVirtualHubStrataCloudManager) validatePutDestinationNatParameters(value interface{}) error {
	if value == nil {
		return fmt.Errorf("parameter value is required, but nil was provided")
	}
	switch value.(type) {
	case cdktn.IResolvable:
		// ok
	case *[]*PaloAltoNextGenerationFirewallVirtualHubStrataCloudManagerDestinationNat:
		value := value.(*[]*PaloAltoNextGenerationFirewallVirtualHubStrataCloudManagerDestinationNat)
		for idx_cd4240, v := range *value {
			if err := _jsii_.ValidateStruct(v, func() string { return fmt.Sprintf("parameter value[%#v]", idx_cd4240) }); err != nil {
				return err
			}
		}
	case []*PaloAltoNextGenerationFirewallVirtualHubStrataCloudManagerDestinationNat:
		value_ := value.([]*PaloAltoNextGenerationFirewallVirtualHubStrataCloudManagerDestinationNat)
		value := &value_
		for idx_cd4240, v := range *value {
			if err := _jsii_.ValidateStruct(v, func() string { return fmt.Sprintf("parameter value[%#v]", idx_cd4240) }); err != nil {
				return err
			}
		}
	default:
		if !_jsii_.IsAnonymousProxy(value) {
			return fmt.Errorf("parameter value must be one of the allowed types: cdktn.IResolvable, *[]*PaloAltoNextGenerationFirewallVirtualHubStrataCloudManagerDestinationNat; received %#v (a %T)", value, value)
		}
	}

	return nil
}

func (p *jsiiProxy_PaloAltoNextGenerationFirewallVirtualHubStrataCloudManager) validatePutDnsSettingsParameters(value *PaloAltoNextGenerationFirewallVirtualHubStrataCloudManagerDnsSettings) error {
	if value == nil {
		return fmt.Errorf("parameter value is required, but nil was provided")
	}
	if err := _jsii_.ValidateStruct(value, func() string { return "parameter value" }); err != nil {
		return err
	}

	return nil
}

func (p *jsiiProxy_PaloAltoNextGenerationFirewallVirtualHubStrataCloudManager) validatePutIdentityParameters(value *PaloAltoNextGenerationFirewallVirtualHubStrataCloudManagerIdentity) error {
	if value == nil {
		return fmt.Errorf("parameter value is required, but nil was provided")
	}
	if err := _jsii_.ValidateStruct(value, func() string { return "parameter value" }); err != nil {
		return err
	}

	return nil
}

func (p *jsiiProxy_PaloAltoNextGenerationFirewallVirtualHubStrataCloudManager) validatePutNetworkProfileParameters(value *PaloAltoNextGenerationFirewallVirtualHubStrataCloudManagerNetworkProfile) error {
	if value == nil {
		return fmt.Errorf("parameter value is required, but nil was provided")
	}
	if err := _jsii_.ValidateStruct(value, func() string { return "parameter value" }); err != nil {
		return err
	}

	return nil
}

func (p *jsiiProxy_PaloAltoNextGenerationFirewallVirtualHubStrataCloudManager) validatePutTimeoutsParameters(value *PaloAltoNextGenerationFirewallVirtualHubStrataCloudManagerTimeouts) error {
	if value == nil {
		return fmt.Errorf("parameter value is required, but nil was provided")
	}
	if err := _jsii_.ValidateStruct(value, func() string { return "parameter value" }); err != nil {
		return err
	}

	return nil
}

func (p *jsiiProxy_PaloAltoNextGenerationFirewallVirtualHubStrataCloudManager) validateRegisterProviderFeatureUsageParameters(feature cdktn.ProviderFeature) error {
	if feature == "" {
		return fmt.Errorf("parameter feature is required, but nil was provided")
	}

	return nil
}

func validatePaloAltoNextGenerationFirewallVirtualHubStrataCloudManager_GenerateConfigForImportParameters(scope constructs.Construct, importToId *string, importFromId *string) error {
	if scope == nil {
		return fmt.Errorf("parameter scope is required, but nil was provided")
	}

	if importToId == nil {
		return fmt.Errorf("parameter importToId is required, but nil was provided")
	}

	if importFromId == nil {
		return fmt.Errorf("parameter importFromId is required, but nil was provided")
	}

	return nil
}

func validatePaloAltoNextGenerationFirewallVirtualHubStrataCloudManager_IsConstructParameters(x interface{}) error {
	if x == nil {
		return fmt.Errorf("parameter x is required, but nil was provided")
	}

	return nil
}

func validatePaloAltoNextGenerationFirewallVirtualHubStrataCloudManager_IsTerraformElementParameters(x interface{}) error {
	if x == nil {
		return fmt.Errorf("parameter x is required, but nil was provided")
	}

	return nil
}

func validatePaloAltoNextGenerationFirewallVirtualHubStrataCloudManager_IsTerraformResourceParameters(x interface{}) error {
	if x == nil {
		return fmt.Errorf("parameter x is required, but nil was provided")
	}

	return nil
}

func (j *jsiiProxy_PaloAltoNextGenerationFirewallVirtualHubStrataCloudManager) validateSetConnectionParameters(val interface{}) error {
	switch val.(type) {
	case *cdktn.SSHProvisionerConnection:
		val := val.(*cdktn.SSHProvisionerConnection)
		if err := _jsii_.ValidateStruct(val, func() string { return "parameter val" }); err != nil {
			return err
		}
	case cdktn.SSHProvisionerConnection:
		val_ := val.(cdktn.SSHProvisionerConnection)
		val := &val_
		if err := _jsii_.ValidateStruct(val, func() string { return "parameter val" }); err != nil {
			return err
		}
	case *cdktn.WinrmProvisionerConnection:
		val := val.(*cdktn.WinrmProvisionerConnection)
		if err := _jsii_.ValidateStruct(val, func() string { return "parameter val" }); err != nil {
			return err
		}
	case cdktn.WinrmProvisionerConnection:
		val_ := val.(cdktn.WinrmProvisionerConnection)
		val := &val_
		if err := _jsii_.ValidateStruct(val, func() string { return "parameter val" }); err != nil {
			return err
		}
	default:
		if !_jsii_.IsAnonymousProxy(val) {
			return fmt.Errorf("parameter val must be one of the allowed types: *cdktn.SSHProvisionerConnection, *cdktn.WinrmProvisionerConnection; received %#v (a %T)", val, val)
		}
	}

	return nil
}

func (j *jsiiProxy_PaloAltoNextGenerationFirewallVirtualHubStrataCloudManager) validateSetCountParameters(val interface{}) error {
	switch val.(type) {
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
	case cdktn.TerraformCount:
		// ok
	default:
		if !_jsii_.IsAnonymousProxy(val) {
			return fmt.Errorf("parameter val must be one of the allowed types: *float64, cdktn.TerraformCount; received %#v (a %T)", val, val)
		}
	}

	return nil
}

func (j *jsiiProxy_PaloAltoNextGenerationFirewallVirtualHubStrataCloudManager) validateSetIdParameters(val *string) error {
	if val == nil {
		return fmt.Errorf("parameter val is required, but nil was provided")
	}

	return nil
}

func (j *jsiiProxy_PaloAltoNextGenerationFirewallVirtualHubStrataCloudManager) validateSetLifecycleParameters(val *cdktn.TerraformResourceLifecycle) error {
	if err := _jsii_.ValidateStruct(val, func() string { return "parameter val" }); err != nil {
		return err
	}

	return nil
}

func (j *jsiiProxy_PaloAltoNextGenerationFirewallVirtualHubStrataCloudManager) validateSetLocationParameters(val *string) error {
	if val == nil {
		return fmt.Errorf("parameter val is required, but nil was provided")
	}

	return nil
}

func (j *jsiiProxy_PaloAltoNextGenerationFirewallVirtualHubStrataCloudManager) validateSetMarketplaceOfferIdParameters(val *string) error {
	if val == nil {
		return fmt.Errorf("parameter val is required, but nil was provided")
	}

	return nil
}

func (j *jsiiProxy_PaloAltoNextGenerationFirewallVirtualHubStrataCloudManager) validateSetNameParameters(val *string) error {
	if val == nil {
		return fmt.Errorf("parameter val is required, but nil was provided")
	}

	return nil
}

func (j *jsiiProxy_PaloAltoNextGenerationFirewallVirtualHubStrataCloudManager) validateSetPlanIdParameters(val *string) error {
	if val == nil {
		return fmt.Errorf("parameter val is required, but nil was provided")
	}

	return nil
}

func (j *jsiiProxy_PaloAltoNextGenerationFirewallVirtualHubStrataCloudManager) validateSetProvisionersParameters(val *[]interface{}) error {
	for idx_97dfc6, v := range *val {
		switch v.(type) {
		case *cdktn.FileProvisioner:
			v := v.(*cdktn.FileProvisioner)
			if err := _jsii_.ValidateStruct(v, func() string { return fmt.Sprintf("parameter val[%#v]", idx_97dfc6) }); err != nil {
				return err
			}
		case cdktn.FileProvisioner:
			v_ := v.(cdktn.FileProvisioner)
			v := &v_
			if err := _jsii_.ValidateStruct(v, func() string { return fmt.Sprintf("parameter val[%#v]", idx_97dfc6) }); err != nil {
				return err
			}
		case *cdktn.LocalExecProvisioner:
			v := v.(*cdktn.LocalExecProvisioner)
			if err := _jsii_.ValidateStruct(v, func() string { return fmt.Sprintf("parameter val[%#v]", idx_97dfc6) }); err != nil {
				return err
			}
		case cdktn.LocalExecProvisioner:
			v_ := v.(cdktn.LocalExecProvisioner)
			v := &v_
			if err := _jsii_.ValidateStruct(v, func() string { return fmt.Sprintf("parameter val[%#v]", idx_97dfc6) }); err != nil {
				return err
			}
		case *cdktn.RemoteExecProvisioner:
			v := v.(*cdktn.RemoteExecProvisioner)
			if err := _jsii_.ValidateStruct(v, func() string { return fmt.Sprintf("parameter val[%#v]", idx_97dfc6) }); err != nil {
				return err
			}
		case cdktn.RemoteExecProvisioner:
			v_ := v.(cdktn.RemoteExecProvisioner)
			v := &v_
			if err := _jsii_.ValidateStruct(v, func() string { return fmt.Sprintf("parameter val[%#v]", idx_97dfc6) }); err != nil {
				return err
			}
		default:
			if !_jsii_.IsAnonymousProxy(v) {
				return fmt.Errorf("parameter val[%#v] must be one of the allowed types: *cdktn.FileProvisioner, *cdktn.LocalExecProvisioner, *cdktn.RemoteExecProvisioner; received %#v (a %T)", idx_97dfc6, v, v)
			}
		}
	}

	return nil
}

func (j *jsiiProxy_PaloAltoNextGenerationFirewallVirtualHubStrataCloudManager) validateSetResourceGroupNameParameters(val *string) error {
	if val == nil {
		return fmt.Errorf("parameter val is required, but nil was provided")
	}

	return nil
}

func (j *jsiiProxy_PaloAltoNextGenerationFirewallVirtualHubStrataCloudManager) validateSetStrataCloudManagerTenantNameParameters(val *string) error {
	if val == nil {
		return fmt.Errorf("parameter val is required, but nil was provided")
	}

	return nil
}

func (j *jsiiProxy_PaloAltoNextGenerationFirewallVirtualHubStrataCloudManager) validateSetTagsParameters(val *map[string]*string) error {
	if val == nil {
		return fmt.Errorf("parameter val is required, but nil was provided")
	}

	return nil
}

func validateNewPaloAltoNextGenerationFirewallVirtualHubStrataCloudManagerParameters(scope constructs.Construct, id *string, config *PaloAltoNextGenerationFirewallVirtualHubStrataCloudManagerConfig) error {
	if scope == nil {
		return fmt.Errorf("parameter scope is required, but nil was provided")
	}

	if id == nil {
		return fmt.Errorf("parameter id is required, but nil was provided")
	}

	if config == nil {
		return fmt.Errorf("parameter config is required, but nil was provided")
	}
	if err := _jsii_.ValidateStruct(config, func() string { return "parameter config" }); err != nil {
		return err
	}

	return nil
}

