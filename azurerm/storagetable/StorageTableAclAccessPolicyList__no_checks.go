// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

//go:build no_runtime_type_checking

package storagetable

// Building without runtime type checking enabled, so all the below just return nil

func (s *jsiiProxy_StorageTableAclAccessPolicyList) validateAllWithMapKeyParameters(mapKeyAttributeName *string) error {
	return nil
}

func (s *jsiiProxy_StorageTableAclAccessPolicyList) validateGetParameters(index *float64) error {
	return nil
}

func (s *jsiiProxy_StorageTableAclAccessPolicyList) validateResolveParameters(context cdktn.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_StorageTableAclAccessPolicyList) validateSetInternalValueParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_StorageTableAclAccessPolicyList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_StorageTableAclAccessPolicyList) validateSetTerraformResourceParameters(val cdktn.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_StorageTableAclAccessPolicyList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewStorageTableAclAccessPolicyListParameters(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}

