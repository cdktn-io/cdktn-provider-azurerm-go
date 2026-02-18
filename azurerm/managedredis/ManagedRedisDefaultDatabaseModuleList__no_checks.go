// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

//go:build no_runtime_type_checking

package managedredis

// Building without runtime type checking enabled, so all the below just return nil

func (m *jsiiProxy_ManagedRedisDefaultDatabaseModuleList) validateAllWithMapKeyParameters(mapKeyAttributeName *string) error {
	return nil
}

func (m *jsiiProxy_ManagedRedisDefaultDatabaseModuleList) validateGetParameters(index *float64) error {
	return nil
}

func (m *jsiiProxy_ManagedRedisDefaultDatabaseModuleList) validateResolveParameters(context cdktn.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_ManagedRedisDefaultDatabaseModuleList) validateSetInternalValueParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_ManagedRedisDefaultDatabaseModuleList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_ManagedRedisDefaultDatabaseModuleList) validateSetTerraformResourceParameters(val cdktn.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_ManagedRedisDefaultDatabaseModuleList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewManagedRedisDefaultDatabaseModuleListParameters(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}

